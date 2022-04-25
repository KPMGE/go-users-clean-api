package protocols

type Controller interface {
	Handle(request *HttpRequest) *HttpResponse
}
