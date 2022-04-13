package helpers

import "github.com/KPMGE/go-users-clean-api/src/presentation/protocols"

func Ok(data interface{}) *protocols.HttpResponse {
	return &protocols.HttpResponse{
		StatusCode: 200,
		Body:       data,
	}
}

func BadRequest(err error) *protocols.HttpResponse {
	return &protocols.HttpResponse{
		StatusCode: 400,
		Body:       err,
	}
}
