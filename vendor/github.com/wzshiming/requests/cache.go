package requests

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"sync"
)

type Cache interface {
	Hash(*Request) string
	Load(name string) (*Response, bool)
	Save(name string, resp *Response)
	Del(name string)
}

func FileCacheDir(s string) fileCacheDir {
	os.MkdirAll(s, 0755)
	return fileCacheDir(s)
}

func MemoryCache() memoryCacheDir {
	return memoryCacheDir{}
}

func Hash(r *Request) string {
	msg, err := r.Unique()
	if err != nil {
		return ""
	}
	data := md5.Sum(msg)
	name := hex.EncodeToString(data[:])
	return name
}

type CacheModel struct {
	Location    *url.URL
	StatusCode  int
	Header      http.Header
	Body        []byte
	ContentType string
}

func (c *CacheModel) Decode(resp *Response) error {
	c.StatusCode = resp.statusCode
	c.Header = resp.header
	c.Location = resp.location
	c.Body = resp.body
	c.ContentType = resp.contentType
	return nil
}

func (c *CacheModel) Encode(resp *Response) error {
	resp.statusCode = c.StatusCode
	resp.header = c.Header
	resp.location = c.Location
	resp.contentType = c.ContentType
	resp.body = c.Body
	return nil
}

type memoryCacheDir struct {
	m sync.Map
}

func (f memoryCacheDir) Hash(r *Request) string {
	return Hash(r)
}

func (f memoryCacheDir) Load(name string) (*Response, bool) {
	d, ok := f.m.Load(name)
	if !ok {
		return nil, false
	}
	data, ok := d.(*Response)
	if !ok {
		return nil, false
	}
	return data, ok
}

func (f memoryCacheDir) Save(name string, resp *Response) {
	f.m.Store(name, resp)
	return
}

func (f memoryCacheDir) Del(name string) {
	f.m.Delete(name)
	return
}

type fileCacheDir string

func (f fileCacheDir) Hash(r *Request) string {
	return Hash(r)
}

func (f fileCacheDir) Load(name string) (*Response, bool) {
	data, err := ioutil.ReadFile(path.Join(string(f), name))
	if err != nil {
		return nil, false
	}

	m := CacheModel{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return nil, false
	}
	resp := &Response{}
	m.Encode(resp)
	return resp, true
}

func (f fileCacheDir) Save(name string, resp *Response) {
	m := &CacheModel{}
	m.Decode(resp)
	data, _ := json.Marshal(m)
	ioutil.WriteFile(path.Join(string(f), name), data, 0666)
	return
}

func (f fileCacheDir) Del(name string) {
	os.Remove(path.Join(string(f), name))
	return
}
