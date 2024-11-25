package kotg_impl

import (
	"github.com/labstack/echo/v4"

	api "github.com/portierglobal/vision-online-companion/api/internal/gen/kotg"
	resp "github.com/portierglobal/vision-online-companion/api/internal/impl/response"
	kotg "github.com/portierglobal/vision-online-companion/business/keyonthego"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ api.ServerInterface = (*Server)(nil)

type Server struct{}

func NewServer() Server {
	return Server{}
}

// List sign requests
// (GET /auth/key-otg/sign)
func (s Server) GetAuthKeyOtgSign(ctx echo.Context) error {
	kotg.ListSignRequest()

	response := []api.SignResponse{}
	return resp.SuccessOk(ctx, response)
}

// Create a new signing request using requestID
// (POST /auth/key-otg/sign/create)
func (s Server) PostAuthKeyOtgSignCreate(ctx echo.Context, params api.PostAuthKeyOtgSignCreateParams) error {
	response := api.CreateSignResponse{}
	return resp.SuccessCreated(ctx, response)
}

// Sign the request using requestID
// (POST /auth/key-otg/sign/submit/{requestID})
func (s Server) PostAuthKeyOtgSignSubmitRequestID(ctx echo.Context, requestID api.RequestID) error {
	response := api.SignResponse{}
	return resp.SuccessOk(ctx, response)
}

// Get signed request using requestID
// (GET /auth/key-otg/sign/{requestID})
func (s Server) GetAuthKeyOtgSignRequestID(ctx echo.Context, requestID api.RequestID) error {
	response := api.SignResponse{}
	return resp.SuccessOk(ctx, response)
}

// Create a new signing request
// (POST /key-otg/sign)
func (s Server) PostKeyOtgSign(ctx echo.Context, params api.PostKeyOtgSignParams) error {
	response := api.CreateSignResponse{}
	return resp.SuccessCreated(ctx, response)
}

// Get signed request
// (GET /key-otg/sign/{requestID})
func (s Server) GetKeyOtgSignRequestID(ctx echo.Context, requestID api.RequestID, params api.GetKeyOtgSignRequestIDParams) error {
	response := api.SignResponse{}
	return resp.SuccessOk(ctx, response)
}

// Sign the request
// (POST /key-otg/sign/{requestID})
func (s Server) PostKeyOtgSignRequestID(ctx echo.Context, requestID api.RequestID, params api.PostKeyOtgSignRequestIDParams) error {
	response := api.SignResponse{}
	return resp.SuccessOk(ctx, response)
}

// Get QR code for this request
// (GET /key-otg/sign/{requestID}/qr)
func (s Server) GetKeyOtgSignRequestIDQr(ctx echo.Context, requestID api.RequestID, params api.GetKeyOtgSignRequestIDQrParams) error {
	response := api.SignResponse{}
	return resp.SuccessOk(ctx, response)
}

// Shutdown the server
// (POST /shutdown)
func (s Server) PostShutdown(ctx echo.Context) error {
	response := api.SignResponse{}
	return resp.SuccessOk(ctx, response)
}
