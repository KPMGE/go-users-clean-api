package helpers

import "github.com/KPMGE/go-users-clean-api/src/presentation/protocols"

func Ok(jsonData string) *protocols.HttpResponse {
	return &protocols.HttpResponse{
		StatusCode: 200,
		JsonBody:   jsonData,
	}
}

func BadRequest(err error) *protocols.HttpResponse {
	return &protocols.HttpResponse{
		StatusCode: 400,
		JsonBody:   err.Error(),
	}
}

func ServerError(err error) *protocols.HttpResponse {
	return &protocols.HttpResponse{
		StatusCode: 500,
		JsonBody:   err.Error(),
	}
}
