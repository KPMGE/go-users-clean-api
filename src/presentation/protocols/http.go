package protocols

type HttpResponse struct {
	StatusCode int
	JsonBody   []byte
}

type HttpRequest struct {
	Params []byte
	Body   []byte
}

func NewHtppRequest(body []byte, params []byte) *HttpRequest {
	return &HttpRequest{
		Body:   body,
		Params: params,
	}
}
