package requests

var defaul = NewClient()

// NewRequest returns a default request
func NewRequest() *Request {
	return defaul.NewRequest()
}
