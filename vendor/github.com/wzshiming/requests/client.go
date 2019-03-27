package requests

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"runtime"
	"time"

	"golang.org/x/net/publicsuffix"
)

type logLevel uint8

const (
	// LogIgnore is ignore all information
	LogIgnore logLevel = iota
	// LogError is show Errors Only
	LogError
	// LogInfo is show basic information
	LogInfo
	// LogMessageHead is show header information
	LogMessageHead
	// LogMessageAll is show all information
	LogMessageAll
)

// NewClient the create a client
func NewClient() *Client {
	c := &Client{
		cli: &http.Client{},
	}
	c.SetSkipVerify(true).
		WithLogger().
		SetLogLevel(LogInfo)
	return c
}

// Client contains basic
type Client struct {
	cli          *http.Client
	log          *log.Logger
	logLevel     logLevel
	proxy        *url.URL
	proxyFromEnv bool
	cache        Cache
}

// NewRequest creates a request instance.
func (c *Client) NewRequest() *Request {
	return newRequest(c)
}

// AddCookies adds cookie to the client.
func (c *Client) AddCookies(u *url.URL, cookies []*http.Cookie) *Client {
	c.cli.Jar.SetCookies(u, cookies)
	return c
}

// GetCookie gets cookie.
func (c *Client) GetCookie(u *url.URL, name string) *http.Cookie {
	for _, cookie := range c.cli.Jar.Cookies(u) {
		if cookie.Name == name {
			return cookie
		}
	}
	return nil
}

// SetCookieJar sets cookie jar.
func (c *Client) SetCookieJar(jar http.CookieJar) *Client {
	c.cli.Jar = jar
	return c
}

// GetCookieJar gets cookie jar.
func (c *Client) GetCookieJar() http.CookieJar {
	return c.cli.Jar
}

// WithCookieJar with default cookie jar.
func (c *Client) WithCookieJar() *Client {
	if c.cli.Jar != nil {
		return c
	}
	jar, err := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	if err != nil {
		c.printError(err)
	}
	return c.SetCookieJar(jar)
}

// SetLogLevel sets log level.
func (c *Client) SetLogLevel(l logLevel) *Client {
	c.logLevel = l
	return c
}

// SetLogger sets given writer for logging.
func (c *Client) SetLogger(w io.Writer) *Client {
	c.log = log.New(w, "["+DefaultPrefix+"] ", 0)
	return c
}

// WithLogger with logger.
func (c *Client) WithLogger() *Client {
	if c.log != nil {
		return c
	}
	return c.SetLogger(os.Stdout)
}

// SetTimeout sets timeout for request raised from client.
func (c *Client) SetTimeout(timeout time.Duration) *Client {
	c.cli.Timeout = timeout
	return c
}

// SetTLSClientConfig sets TLSClientConfig.
func (c *Client) SetTLSClientConfig(config *tls.Config) *Client {
	transport, err := c.getTransport()
	if err != nil {
		c.printError(err)
		return c
	}
	transport.TLSClientConfig = config
	return c
}

// SetKeepAlives sets the keep alives.
func (c *Client) SetKeepAlives(enable bool) *Client {
	transport, err := c.getTransport()
	if err != nil {
		c.printError(err)
		return c
	}
	transport.DisableKeepAlives = !enable
	return c
}

// SetProxyFunc sets the Proxy function.
func (c *Client) SetProxyFunc(proxy func(*http.Request) (*url.URL, error)) *Client {
	transport, err := c.getTransport()
	if err != nil {
		c.printError(err)
		return c
	}
	transport.Proxy = proxy
	return c
}

// SetProxyURL sets the Proxy URL.
func (c *Client) SetProxyURL(u *url.URL) *Client {
	return c.SetProxyFunc(http.ProxyURL(u))
}

// SetProxyURLByStr sets the Proxy URL.
func (c *Client) SetProxyURLByStr(rawurl string) *Client {
	u, err := url.Parse(rawurl)
	if err != nil {
		c.printError(err)
		return c
	}
	return c.SetProxyURL(u)
}

// SetProxyFromEnvironment sets the Proxy URL.
func (c *Client) SetProxyFromEnvironment(u bool) *Client {
	c.proxyFromEnv = u
	return c
}

// AddRootCert adds one or more root certificates into requests client
func (c *Client) AddRootCert(cert *x509.Certificate) *Client {
	config, err := c.getTLSConfig()
	if err != nil {
		c.printError(err)
		return c
	}
	if config.RootCAs == nil {
		config.RootCAs = x509.NewCertPool()
	}
	config.RootCAs.AddCert(cert)
	return c
}

// WithSystemCertPool sets system cert poll
func (c *Client) WithSystemCertPool() *Client {
	config, err := c.getTLSConfig()
	if err != nil {
		c.printError(err)
		return c
	}

	if runtime.GOOS != "windows" {
		ca, err := x509.SystemCertPool()
		if err != nil {
			c.printError(err)
			return c
		}
		config.RootCAs = ca
	}
	return c
}

// SetSkipVerify sets skip ca verify
func (c *Client) SetSkipVerify(b bool) *Client {
	config, err := c.getTLSConfig()
	if err != nil {
		c.printError(err)
		return c
	}
	config.InsecureSkipVerify = b
	return c
}

// SetCache sets cache
func (c *Client) SetCache(cache Cache) *Client {
	c.cache = cache
	return c
}

// SetCheckRedirect sets specifies the policy for handling redirects
func (c *Client) SetCheckRedirect(f func(req *http.Request, via []*http.Request) error) *Client {
	c.cli.CheckRedirect = f
	return c
}

// NoRedirect disable redirects
func (c *Client) NoRedirect() *Client {
	return c.SetCheckRedirect(func(req *http.Request, via []*http.Request) error {
		return ErrNoRedirect
	})
}

// getTLSConfig returns a TLS config
func (c *Client) getTLSConfig() (*tls.Config, error) {
	transport, err := c.getTransport()
	if err != nil {
		return nil, err
	}
	if transport.TLSClientConfig == nil {
		transport.TLSClientConfig = &tls.Config{}
	}
	return transport.TLSClientConfig, nil
}

// getTransport returns a transport
func (c *Client) getTransport() (*http.Transport, error) {
	if c.cli.Transport == nil {
		c.cli.Transport = &http.Transport{}
	}

	if transport, ok := c.cli.Transport.(*http.Transport); ok {
		return transport, nil
	}
	return nil, ErrNotTransport
}

// do executes and returns response
func (c *Client) do(req *Request) (*Response, error) {
	_, err := req.process()
	if err != nil {
		return nil, err
	}

	hash := ""
	if c.cache != nil {
		hash = c.cache.Hash(req)
		if req.noCache {
			c.cache.Del(hash)
		} else if resp, ok := c.cache.Load(hash); ok {
			c.printCacheHit(req)
			return resp, nil
		}
	}
	c.printRequest(req)
	req.sendAt = time.Now()
	resp, err := c.cli.Do(req.rawRequest)
	if err != nil {
		return nil, err
	}
	response := &Response{
		request:     req,
		rawResponse: resp,
		recvAt:      time.Now(),
	}
	err = response.process()
	if err != nil {
		return nil, err
	}
	c.printResponse(response)
	if c.cache != nil {
		if code := response.StatusCode(); code >= 200 && code < 300 {
			c.cache.Save(hash, response)
		}
	}
	return response, nil
}

func (c *Client) printError(err error) {
	if c.log != nil && c.logLevel >= LogError {
		c.log.Printf("Error: %v", err.Error())
	}
}

func (c *Client) printCacheHit(r *Request) {
	if c.log != nil {
		switch c.logLevel {
		case LogInfo:
			c.log.Printf("CacheHit: %s", r.String())
		case LogMessageHead:
			c.log.Printf("CacheHit: %s", r.MessageHead())
		case LogMessageAll:
			c.log.Printf("CacheHit: %s", r.Message())
		}
	}
}

func (c *Client) printRequest(r *Request) {
	if c.log != nil {
		switch c.logLevel {
		case LogInfo:
			c.log.Printf("Request:  %s", r.String())
		case LogMessageHead:
			c.log.Printf("Request:  %s", r.MessageHead())
		case LogMessageAll:
			c.log.Printf("Request:  %s", r.Message())
		}
	}
}

func (c *Client) printResponse(r *Response) {
	if c.log != nil {
		switch c.logLevel {
		case LogInfo:
			c.log.Printf("Response: %s", r.String())
		case LogMessageHead:
			c.log.Printf("Response: %s", r.MessageHead())
		case LogMessageAll:
			c.log.Printf("Response: %s", r.Message())
		}
	}
}
