package helpers

import (
	"net/http"

	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
)

func Ok(data any) *protocols.HttpResponse {
	return &protocols.HttpResponse{
		StatusCode: http.StatusOK,
		Body:       data,
	}
}

func BadRequest(err error) *protocols.HttpResponse {
	return &protocols.HttpResponse{
		StatusCode: http.StatusBadRequest,
		Body:       err.Error(),
	}
}

func NotFound(err error) *protocols.HttpResponse {
	return &protocols.HttpResponse{
		StatusCode: http.StatusNotFound,
		Body:       err.Error(),
	}
}

func ServerError(err error) *protocols.HttpResponse {
	return &protocols.HttpResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       err.Error(),
	}
}
