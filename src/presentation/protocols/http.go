package protocols

type HttpResponse struct {
	StatusCode int
	Body       any
}

type HttpRequest struct {
	Params []byte
	Body   []byte
}

func NewHttpRequest(body []byte, params []byte) *HttpRequest {
	return &HttpRequest{
		Body:   body,
		Params: params,
	}
}
