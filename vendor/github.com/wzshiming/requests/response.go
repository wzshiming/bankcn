package requests

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"net/http/httputil"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

// Response is an object represents executed request and its values.
type Response struct {
	contentType string
	request     *Request
	rawResponse *http.Response
	body        []byte
	recvAt      time.Time
}

// WriteFile is writes the response body to file.
func (r *Response) WriteFile(file string) error {
	return ioutil.WriteFile(file, r.body, 0666)
}

// Body returns HTTP response as []byte array for the executed request.
func (r *Response) Body() []byte {
	return r.body
}

// ContentType returns HTTP response content type
func (r *Response) ContentType() string {
	return r.contentType
}

// Status returns the HTTP status string for the executed request.
func (r *Response) Status() string {
	if r.rawResponse == nil {
		return "from cache"
	}
	return r.rawResponse.Status
}

// StatusCode returns the HTTP status code for the executed request.
func (r *Response) StatusCode() int {
	if r.rawResponse == nil {
		return 200
	}
	return r.rawResponse.StatusCode
}

// Header returns the response headers
func (r *Response) Header() http.Header {
	if r.rawResponse == nil {
		return nil
	}
	return r.rawResponse.Header
}

// Cookies to access all the response cookies
func (r *Response) Cookies() []*http.Cookie {
	if r.rawResponse == nil {
		return nil
	}
	return r.rawResponse.Cookies()
}

// Time returns the time of HTTP response time that from request we sent and received a request.
func (r *Response) Time() time.Duration {
	return r.recvAt.Sub(r.request.sendAt)
}

// ReceivedAt returns when response got recevied from server for the request.
func (r *Response) ReceivedAt() time.Time {
	return r.recvAt
}

// Size returns the HTTP response size in bytes.
func (r *Response) Size() int {
	return len(r.body)
}

// RawBody returns the HTTP raw response body.
func (r *Response) RawBody() io.Reader {
	return bytes.NewReader(r.body)
}

// String returns the HTTP response basic information
func (r *Response) String() string {
	return fmt.Sprintf("%s %s %d %d %s", r.request.method, r.request.baseURL.String(), r.StatusCode(), r.Size(), r.Time())
}

// Message returns the HTTP response all information
func (r *Response) Message() string {
	return r.message(true)
}

// MessageHead returns the HTTP response header information
func (r *Response) MessageHead() string {
	return r.message(false)
}

func (r *Response) message(body bool) string {
	b, err := httputil.DumpResponse(r.rawResponse, false)
	if err != nil {
		return err.Error()
	}
	if body {
		b = append(b, r.Body()...)
	}
	return string(b)
}

func (r *Response) process() (err error) {
	resp := r.rawResponse
	if resp.Body == nil {
		return nil
	}
	if r.request.discardResponse {
		return resp.Body.Close()
	}

	r.contentType = resp.Header.Get(HeaderContentType)

	body := tryCharset(resp.Body, r.contentType)

	r.body, _ = ioutil.ReadAll(body)
	if err := resp.Body.Close(); err != nil {
		return err
	}
	return nil
}

func tryCharset(r io.Reader, contentType string) io.Reader {
	if _, params, err := mime.ParseMediaType(contentType); err == nil {
		if cs, ok := params["charset"]; ok {
			if e, _ := charset.Lookup(cs); e != nil && e != encoding.Nop {
				r = transform.NewReader(r, e.NewDecoder())
			}
		}
	}
	return r
}
