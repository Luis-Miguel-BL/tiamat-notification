package api

type Request struct {
	Method string
	Body   string
}
type Response struct {
	StatusCode int
	Body       string
}

type RequestBody interface {
	Validate() error
}

type ControllerFunc func(Request) Response

type Router interface {
	Handle(path string, req Request) (res Response)
}
