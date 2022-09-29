package factories

import (
	"github.com/KPMGE/go-users-clean-api/src/application/services"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
)

func MakeLoginController() protocols.Controller {
  tokenProvider := 
  service := services.NewLoginService()
}
