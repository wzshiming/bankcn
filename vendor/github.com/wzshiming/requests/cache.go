package requests

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
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

type fileCacheDir string

func (f fileCacheDir) Hash(r *Request) string {
	msg := r.messageHash()
	data := md5.Sum([]byte(msg))
	name := hex.EncodeToString(data[:])
	return name
}

type cacheMod struct {
	Body        []byte
	ContentType string
}

func (f fileCacheDir) Load(name string) (*Response, bool) {
	data, err := ioutil.ReadFile(path.Join(string(f), name))
	if err != nil {
		return nil, false
	}
	m := cacheMod{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return nil, false
	}
	resp := Response{
		contentType: m.ContentType,
		body:        m.Body,
	}
	return &resp, true
}

func (f fileCacheDir) Save(name string, resp *Response) {
	m := cacheMod{
		Body:        resp.Body(),
		ContentType: resp.ContentType(),
	}
	data, _ := json.Marshal(m)
	ioutil.WriteFile(path.Join(string(f), name), data, 0666)
	return
}

func (f fileCacheDir) Del(name string) {
	os.Remove(path.Join(string(f), name))
	return
}
