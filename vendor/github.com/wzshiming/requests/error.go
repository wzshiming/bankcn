package requests

import "errors"

var (
	ErrNoRedirect   = errors.New("No Redirect")
	ErrNotTransport = errors.New("not a *http.Transport")
)
