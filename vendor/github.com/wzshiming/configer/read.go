package configer

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Read(cfgpath string) ([]byte, error) {
	u, err := url.Parse(cfgpath)
	if err != nil {
		return nil, err
	}

	switch u.Scheme {
	case "http", "https":
		resp, err := http.Get(u.String())
		if err != nil {
			return nil, err
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return data, nil

	case "file", "":
		data, err := ioutil.ReadFile(u.Path)
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	return nil, errors.New("error config path " + cfgpath)
}
