package requests

import (
	"bytes"
	"encoding/base64"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strings"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

// Common HTTP methods.
//
// Unless otherwise noted, these are defined in RFC 7231 section 4.3.
const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH" // RFC 5789
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"

	CharsetUTF8     = "; charset=utf-8"
	MimeJSON        = "application/json" + CharsetUTF8
	MimeXML         = "application/xml" + CharsetUTF8
	MimeTextPlain   = "text/plain" + CharsetUTF8
	MimeOctetStream = "application/octet-stream" + CharsetUTF8
	MimeURLEncoded  = "application/x-www-form-urlencoded" + CharsetUTF8
	MimeFormData    = "multipart/form-data" + CharsetUTF8

	HeaderUserAgent       = "User-Agent"
	HeaderAccept          = "Accept"
	HeaderContentType     = "Content-Type"
	HeaderContentLength   = "Content-Length"
	HeaderContentEncoding = "Content-Encoding"
	HeaderAuthorization   = "Authorization"
)

// Default
var (
	DefaultPrefix         = "REQUESTS"
	DefaultVersion        = "1.0"
	DefaultUserAgentValue = "Mozilla/5.0 (compatible; " + DefaultPrefix + "/" + DefaultVersion + "; +https://github.com/wzshiming/requests)"
)

// paramPair represent custom data part for header path query form
type paramPair struct {
	Param string
	Value string
}

type paramPairs []*paramPair

func (t *paramPairs) Clone() paramPairs {
	n := make(paramPairs, len(*t))
	copy(n, *t)
	return n
}

func (t *paramPairs) add(i int, n *paramPair) {
	*t = append(*t, nil)
	l := len(*t)
	copy((*t)[i+1:l], (*t)[i:l-1])
	(*t)[i] = n
}

func (t *paramPairs) Add(param, value string) {
	i := t.SearchIndex(param)
	t.add(i, &paramPair{
		Param: param,
		Value: value,
	})
}

func (t *paramPairs) AddReplace(param, value string) {
	i := t.SearchIndex(param)
	tt := t.Index(i - 1)
	if tt == nil || tt.Param != param {
		t.add(i, &paramPair{
			Param: param,
			Value: value,
		})
	} else {
		tt.Value = value
	}
	return
}

func (t *paramPairs) AddNoRepeat(param, value string) {
	i := t.SearchIndex(param)
	tt := t.Index(i - 1)
	if tt == nil || tt.Param != param {
		t.add(i, &paramPair{
			Param: param,
			Value: value,
		})
	}
	return
}

func (t *paramPairs) Search(name string) (*paramPair, bool) {
	i := t.SearchIndex(name)
	if i == 0 {
		return nil, false
	}
	tt := t.Index(i - 1)
	if tt == nil || tt.Param != name {
		return nil, false
	}
	return tt, true
}

func (t *paramPairs) SearchIndex(name string) int {
	i := sort.Search(t.Len(), func(i int) bool {
		d := t.Index(i)
		if d == nil {
			return false
		}
		return d.Param < name
	})
	return i
}

func (t *paramPairs) Index(i int) *paramPair {
	if i >= t.Len() || i < 0 {
		return nil
	}
	return (*t)[i]
}

func (t *paramPairs) Len() int {
	return len(*t)
}

// multiFile represent custom data part for multipart request
type multiFile struct {
	Param       string
	FileName    string
	ContentType string
	io.Reader
}

type multiFiles []*multiFile

func toHeader(header http.Header, p paramPairs) (http.Header, error) {
	for _, v := range p {
		header[v.Param] = append(header[v.Param], v.Value)
	}
	return header, nil
}

func toQuery(rawQuery string, p paramPairs) (string, error) {
	param := url.Values{}
	for _, v := range p {
		param[v.Param] = append(param[v.Param], v.Value)
	}
	param0, _ := url.ParseQuery(rawQuery)
	for k, v := range param0 {
		_, ok := param[k]
		if !ok {
			param[k] = v
		}
	}
	return param.Encode(), nil
}

var toPathCompile = regexp.MustCompile(`\{[^}]*\}`)

func toPath(path string, p paramPairs) (string, error) {
	path = toPathCompile.ReplaceAllStringFunc(path, func(s string) string {
		k := s[1 : len(s)-1]
		// Because the number is small, it's faster to use the loop directly
		for _, v := range p {
			if v.Param == k {
				return v.Value
			}
		}
		return s
	})
	return path, nil
}

func toForm(p paramPairs) (io.Reader, error) {
	vs := url.Values{}
	for _, v := range p {
		vs.Add(v.Param, v.Value)
	}
	return bytes.NewBufferString(vs.Encode()), nil
}

func toMulti(p paramPairs, m multiFiles) (io.Reader, string, error) {
	buf := bytes.NewBuffer(nil)
	mw := multipart.NewWriter(buf)

	for _, v := range p {
		err := mw.WriteField(v.Param, v.Value)
		if err != nil {
			return nil, "", err
		}
	}

	for _, v := range m {
		w, err := mw.CreateFormFile(v.Param, v.FileName)
		if err != nil {
			return nil, "", err
		}
		_, err = io.Copy(w, v.Reader)
		if err != nil {
			return nil, "", err
		}
	}

	err := mw.Close()
	if err != nil {
		return nil, "", err
	}
	return buf, mw.FormDataContentType(), nil
}

// See 2 (end of page 4) http://www.ietf.org/rfc/rfc2617.txt
// "To receive authorization, the client sends the userid and password,
// separated by a single colon (":") character, within a base64
// encoded string in the credentials."
// It is not meant to be urlencoded.
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func readCookies(line string) (cookies []*http.Cookie) {
	parts := strings.Split(strings.TrimSpace(line), ";")
	if len(parts) == 1 && parts[0] == "" {
		return
	}
	// Per-line attributes
	for i := 0; i < len(parts); i++ {
		parts[i] = strings.TrimSpace(parts[i])
		if len(parts[i]) == 0 {
			continue
		}
		name, val := parts[i], ""
		if j := strings.Index(name, "="); j >= 0 {
			name, val = name[:j], name[j+1:]
		}

		// Strip the quotes, if present.
		if len(val) > 1 && val[0] == '"' && val[len(val)-1] == '"' {
			val = val[1 : len(val)-1]
		}

		cookies = append(cookies, &http.Cookie{Name: name, Value: val})
	}

	return cookies
}

// Cookies raw to Cookies.
func Cookies(raw interface{}) []*http.Cookie {
	switch t := raw.(type) {
	case []*http.Cookie:
		return t
	case *http.Cookie:
		return []*http.Cookie{t}
	case http.Cookie:
		return []*http.Cookie{&t}
	case string:
		return readCookies(t)
	}
	return nil
}

// URL raw to URL structure.
func URL(raw interface{}) *url.URL {
	switch t := raw.(type) {
	case *url.URL:
		return t
	case url.URL:
		return &t
	case string:
		r, _ := url.Parse(t)
		return r
	}
	return nil
}

// TryCharset try charset
func TryCharset(r io.Reader, contentType string) io.Reader {
	if _, params, err := mime.ParseMediaType(contentType); err == nil {
		if cs, ok := params["charset"]; ok {
			if e, _ := charset.Lookup(cs); e != nil && e != encoding.Nop {
				r = transform.NewReader(r, e.NewDecoder())
			}
		}
	}
	return r
}
