package handlers

import (
	"github.com/spn/go/clients/federation"
	"github.com/spn/go/clients/horizon"
	"github.com/spn/go/clients/spntoml"
	"github.com/spn/go/services/bridge/internal/config"
	"github.com/spn/go/services/bridge/internal/db"
	"github.com/spn/go/services/bridge/internal/listener"
	"github.com/spn/go/services/bridge/internal/submitter"
	"github.com/spn/go/support/http"
)

// RequestHandler implements bridge server request handlers
type RequestHandler struct {
	Config               *config.Config                          `inject:""`
	Client               http.SimpleHTTPClientInterface          `inject:""`
	Horizon              horizon.ClientInterface                 `inject:""`
	Database             db.Database                             `inject:""`
	SpnTomlResolver  spntoml.ClientInterface             `inject:""`
	FederationResolver   federation.ClientInterface              `inject:""`
	TransactionSubmitter submitter.TransactionSubmitterInterface `inject:""`
	PaymentListener      *listener.PaymentListener               `inject:""`
}

func (rh *RequestHandler) isAssetAllowed(code string, issuer string) bool {
	for _, asset := range rh.Config.Assets {
		if asset.Code == code && asset.Issuer == issuer {
			return true
		}
	}
	return false
}
