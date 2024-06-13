package sales

import (
	userapp "github.com/ardanlabs/encore/app/domain/userapp"
	"github.com/ardanlabs/encore/business/domain/userbus"
)

type appDomain struct {
	userApp *userapp.App
}

type busDomain struct {
	userBus *userbus.Business
}
