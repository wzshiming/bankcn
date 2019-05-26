package requests

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

// Request type is used to compose and send individual request from client
type Request struct {
	baseURL         *url.URL
	method          string
	headerParam     paramPairs
	queryParam      paramPairs
	pathParam       paramPairs
	formParam       paramPairs
	multiFiles      multiFiles
	body            io.Reader
	sendAt          time.Time
	rawRequest      *http.Request
	client          *Client
	ctx             context.Context
	discardResponse bool
	noCache         bool
}

func newRequest(c *Client) *Request {
	return &Request{
		client: c,
		method: MethodGet,
	}
}

// Clone returns clone the request
func (r *Request) Clone() *Request {
	n := &Request{}
	*n = *r
	if n.baseURL != nil {
		bu := *n.baseURL
		n.baseURL = &bu
	}
	return n
}

// AddCookies adds cookie to the client.
func (r *Request) AddCookies(cookies []*http.Cookie) *Request {
	r.client.AddCookies(r.baseURL, cookies)
	return r
}

// Client returns the client of the request
func (r *Request) Client() *Client {
	return r.client
}

// SetURL sets URL in the client instance.
func (r *Request) SetURL(u *url.URL) *Request {
	if u == nil {
		r.baseURL = nil
		return r
	}
	r.baseURL = u
	if user := r.baseURL.User; user != nil {
		pwd, _ := user.Password()
		r.SetBasicAuth(user.Username(), pwd)
		r.baseURL.User = nil
	}

	if r.baseURL.RawQuery != "" {
		qs, _ := url.ParseQuery(r.baseURL.RawQuery)
		for k, v := range qs {
			for _, v := range v {
				r.SetQuery(k, v)
			}
		}
		r.baseURL.RawQuery = ""
	}
	return r
}

// SetURLByStr sets URL in the client instance.
func (r *Request) SetURLByStr(rawurl string) *Request {
	r.SetURL(r.GetURL(rawurl))
	return r
}

// GetURL gets URL.
func (r *Request) GetURL(rawurl string) *url.URL {
	if rawurl == "" {
		u, _ := r.processURL()
		return u
	}
	var nu *url.URL
	var err error
	if r.baseURL == nil {
		nu, err = url.Parse(rawurl)
	} else {
		nu, err = r.baseURL.Parse(rawurl)
	}
	if err != nil {
		r.client.printError(err)
	}
	return nu
}

// SetContext sets context.Context for current Request.
func (r *Request) SetContext(ctx context.Context) *Request {
	r.ctx = ctx
	return r
}

// SetTimeout sets timeout for current Request.
func (r *Request) SetTimeout(timeout time.Duration) *Request {
	return r.SetDeadline(time.Now().Add(timeout))
}

// SetDeadline sets deadline for current Request.
func (r *Request) SetDeadline(d time.Time) *Request {
	if r.ctx == nil {
		r.ctx = context.TODO()
	}
	r.ctx, _ = context.WithDeadline(r.ctx, d)
	return r
}

func (r *Request) withContext() {
	if r.ctx != nil {
		r.rawRequest = r.rawRequest.WithContext(r.ctx)
	}
}

func (r *Request) isCancelled() bool {
	return r.ctx != nil && r.ctx.Err() != nil
}

// SetHeader sets header field and its value in the current request.
func (r *Request) SetHeader(param, value string) *Request {
	//	param = textproto.CanonicalMIMEHeaderKey(param)
	r.headerParam.AddReplace(param, value)
	return r
}

// AddHeader adds header field and its value in the current request.
func (r *Request) AddHeader(param, value string) *Request {
	//	param = textproto.CanonicalMIMEHeaderKey(param)
	r.headerParam.Add(param, value)
	return r
}

// AddHeaderIfNot adds header field and its value in the current request if not.
func (r *Request) AddHeaderIfNot(param, value string) *Request {
	//	param = textproto.CanonicalMIMEHeaderKey(param)
	r.headerParam.AddNoRepeat(param, value)
	return r
}

// SetPath sets path parameter and its value in the current request.
func (r *Request) SetPath(param, value string) *Request {
	r.pathParam.AddReplace(param, value)
	return r
}

// AddPathIfNot adds path parameter and its value in the current request if not.
func (r *Request) AddPathIfNot(param, value string) *Request {
	r.pathParam.AddNoRepeat(param, value)
	return r
}

// SetQuery sets query parameter and its value in the current request.
func (r *Request) SetQuery(param, value string) *Request {
	r.queryParam.AddReplace(param, value)
	return r
}

// AddQuery adds query field and its value in the current request.
func (r *Request) AddQuery(param, value string) *Request {
	r.queryParam.Add(param, value)
	return r
}

// AddQueryIfNot adds query field and its value in the current request if not.
func (r *Request) AddQueryIfNot(param, value string) *Request {
	r.queryParam.AddNoRepeat(param, value)
	return r
}

// SetForm sets multiple form parameters with multi-value
func (r *Request) SetForm(param, value string) *Request {
	r.formParam.AddReplace(param, value)
	return r
}

// AddForm adds from field and its value in the current request.
func (r *Request) AddForm(param, value string) *Request {
	r.formParam.Add(param, value)
	return r
}

// AddFormIfNot adds from field and its value in the current request if not.
func (r *Request) AddFormIfNot(param, value string) *Request {
	r.formParam.AddNoRepeat(param, value)
	return r
}

// SetFile sets custom data using io.Reader for multipart upload.
func (r *Request) SetFile(param, fileName, contentType string, reader io.Reader) *Request {
	r.multiFiles = append(r.multiFiles, &multiFile{
		Param:       param,
		FileName:    fileName,
		ContentType: contentType,
		Reader:      reader,
	})
	return r
}

// SetJSON sets data encoded by JSON to the request body.
func (r *Request) SetJSON(i interface{}) *Request {
	data, err := json.Marshal(i)
	if err != nil {
		r.client.printError(err)
		return r
	}
	r.body = bytes.NewReader(data)
	r.AddHeaderIfNot(HeaderContentType, MimeJSON)
	return r
}

// SetXML sets data encoded by XML to the request body.
func (r *Request) SetXML(i interface{}) *Request {
	data, err := xml.Marshal(i)
	if err != nil {
		r.client.printError(err)
		return r
	}
	r.body = bytes.NewReader(data)
	r.AddHeaderIfNot(HeaderContentType, MimeXML)
	return r
}

// SetBody sets request body for the request.
func (r *Request) SetBody(body io.Reader) *Request {
	r.body = body
	return r
}

// SetContentType sets content type header in the HTTP request.
func (r *Request) SetContentType(contentType string) *Request {
	r.SetHeader(HeaderContentType, contentType)
	return r
}

// SetBasicAuth sets basic authentication header in the HTTP request.
func (r *Request) SetBasicAuth(username, password string) *Request {
	r.SetHeader(HeaderAuthorization, "Basic "+basicAuth(username, password))
	return r
}

// SetAuthToken sets bearer auth token header in the HTTP request.
func (r *Request) SetAuthToken(token string) *Request {
	r.SetHeader(HeaderAuthorization, "Bearer "+token)
	return r
}

// SetUserAgent sets user agent header in the HTTP request.
func (r *Request) SetUserAgent(ua string) *Request {
	r.SetHeader(HeaderUserAgent, ua)
	return r
}

// SetDiscardResponse sets unread the response body.
func (r *Request) SetDiscardResponse(discard bool) *Request {
	r.discardResponse = discard
	return r
}

// SetMethod sets method in the HTTP request.
func (r *Request) SetMethod(method string) *Request {
	r.method = strings.ToUpper(method)
	return r
}

// Head does HEAD HTTP request.
func (r *Request) Head(url string) (*Response, error) {
	return r.Clone().SetMethod(MethodHead).SetURLByStr(url).do()
}

// Get does GET HTTP request.
func (r *Request) Get(url string) (*Response, error) {
	return r.Clone().SetMethod(MethodGet).SetURLByStr(url).do()
}

// Post does POST HTTP request.
func (r *Request) Post(url string) (*Response, error) {
	return r.Clone().SetMethod(MethodPost).SetURLByStr(url).do()
}

// Put does PUT HTTP request.
func (r *Request) Put(url string) (*Response, error) {
	return r.Clone().SetMethod(MethodPut).SetURLByStr(url).do()
}

// Delete does DELETE HTTP request.
func (r *Request) Delete(url string) (*Response, error) {
	return r.Clone().SetMethod(MethodDelete).SetURLByStr(url).do()
}

// Options does OPTIONS HTTP request.
func (r *Request) Options(url string) (*Response, error) {
	return r.Clone().SetMethod(MethodOptions).SetURLByStr(url).do()
}

// Trace does TRACE HTTP request.
func (r *Request) Trace(url string) (*Response, error) {
	return r.Clone().SetMethod(MethodTrace).SetURLByStr(url).do()
}

// Patch does PATCH HTTP request.
func (r *Request) Patch(url string) (*Response, error) {
	return r.Clone().SetMethod(MethodPatch).SetURLByStr(url).do()
}

// NoCache Clear the cache for this request
func (r *Request) NoCache() *Request {
	r.noCache = true
	return r
}

// Do performs the HTTP request
func (r *Request) Do() (*Response, error) {
	return r.Clone().do()
}

func (r *Request) do() (*Response, error) {
	return r.client.do(r)
}

func (r *Request) processURL() (*url.URL, error) {
	u := r.baseURL
	if u == nil {
		u = &url.URL{}
	}
	q := []string{}
	// fill path
	if len(r.pathParam) != 0 {
		path, err := toPath(u.Path, r.pathParam)
		if err != nil {
			return nil, err
		}
		q = append(q, path)
	} else {
		q = append(q, u.Path)
	}

	// fill query
	if len(r.queryParam) != 0 {
		rq, err := toQuery(u.RawQuery, r.queryParam)
		if err != nil {
			return nil, err
		}
		q = append(q, rq)
	} else if u.RawQuery != "" {
		q = append(q, u.RawQuery)
	}
	return u.Parse(strings.Join(q, "?"))
}

func (r *Request) RawRequest() (*http.Request, error) {
	if r.rawRequest != nil {
		return r.rawRequest, nil
	}
	u, err := r.processURL()
	if err != nil {
		return nil, err
	}
	r.baseURL = u
	if r.body == nil {
		if len(r.multiFiles) != 0 { // fill multpair
			body, contentType, err := toMulti(r.formParam, r.multiFiles)
			if err != nil {
				return nil, err
			}
			r.AddHeaderIfNot(HeaderContentType, contentType)
			r.body = body
		} else if len(r.formParam) != 0 { // fill form
			body, err := toForm(r.formParam)
			if err != nil {
				return nil, err
			}
			r.AddHeaderIfNot(HeaderContentType, MimeURLEncoded)
			r.body = body
		}
	}

	req, err := http.NewRequest(r.method, r.baseURL.String(), r.body)
	if err != nil {
		return nil, err
	}

	// fill header
	r.AddHeaderIfNot(HeaderUserAgent, DefaultUserAgentValue)
	header, err := toHeader(req.Header, r.headerParam)
	if err != nil {
		return nil, err
	}

	if r.client.proxyFromEnv {
		u, err := http.ProxyFromEnvironment(req)
		if err != nil {
			return nil, err
		}
		r.client.SetProxyURL(u)
	}

	req.Header = header
	r.rawRequest = req

	r.withContext()
	return req, nil
}

func (r *Request) messageBody() []byte {
	if r.rawRequest.Body == nil {
		return nil
	}
	body, _ := ioutil.ReadAll(r.rawRequest.Body)
	r.rawRequest.Body.Close()
	r.rawRequest.Body = ioutil.NopCloser(bytes.NewReader(body))
	return body
}

// String returns the HTTP request basic information
func (r *Request) String() string {
	return fmt.Sprintf("%s %s", r.method, r.baseURL.String())
}

// Message returns the HTTP request all information
func (r *Request) Message() string {
	return r.message(true)
}

// MessageHead returns the HTTP request header information
func (r *Request) MessageHead() string {
	return r.message(false)
}

// Unique returns identifies the uniqueness of the request
func (r *Request) Unique() ([]byte, error) {
	req, err := r.Clone().RawRequest()
	if err != nil {
		return nil, err
	}

	b, err := httputil.DumpRequest(req, false)
	if err != nil {
		return nil, err
	}

	b = append(b, r.messageBody()...)
	return b, nil
}

func (r *Request) message(body bool) string {
	req, err := r.Clone().RawRequest()
	if err != nil {
		return err.Error()
	}

	if r.client.cli.Jar != nil {
		for _, v := range r.client.cli.Jar.Cookies(req.URL) {
			req.AddCookie(v)
		}
	}

	b, err := httputil.DumpRequest(req, false)
	if err != nil {
		return err.Error()
	}

	if body {
		b = append(b, r.messageBody()...)
	}
	return string(b)
}
