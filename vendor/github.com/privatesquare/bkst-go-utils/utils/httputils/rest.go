package httputils

import (
	"github.com/jarcoal/httpmock"
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"net/http"
)

const (
	InternalServerErrMsg      = "Unable to process the request due to an internal error. Please contact the system administrator"
	InvalidPayloadErrMsg      = "Payload is not valid"
	ServerStartupMsg          = "Starting the API server..."
	ServerStartupSuccessMsg   = "The API server has started and is listening on %s"
	ServerUrlMsg              = "API server url: %s"
	ServerAPIDocsMsg          = "API server swagger docs url: %s"
	ServerStartupErrMsg       = "Unable to run the web server"
	NotImplementedYetMsg      = "Not Implemented"
	AuthorizationErrMsg       = "Insufficient privileges"
	ApiHealthPath             = "/health"
	ContentTypeHeaderKey      = "Content-Type"
	AcceptHeaderKey           = "Accept"
	AuthorizationHeaderKey    = "Authorization"
	ApplicationJsonMIMEType   = "application/json"
	TextPlainMIMEType         = "text/plain"
	XwwwFromUrlencodeMIMEType = "application/x-www-form-urlencoded"
	SwaggerPath               = "/swagger/*any"
	SwaggerSpecPathFormat     = "%s/swagger/doc.json"
	SwaggerUriPathFormat      = "%s/swagger/index.html"
)

type RestMsg struct {
	Message string `json:"message"`
}

type RestErrMsg struct {
	Error string `json:"error"`
}

// CheckHTTPErrorStatusCode checks for default errors like 401, 403 and API error status codes that are not defined.
func CheckHTTPErrorStatusCode(statusCode int, errMsg string) *errors.RestErr {
	switch statusCode {
	case http.StatusBadRequest:
		return errors.BadRequestError(errMsg)
	case http.StatusUnauthorized:
		return errors.UnauthorizedError(errMsg)
	case http.StatusForbidden:
		return errors.ForbiddenError(errMsg)
	default:
		return errors.InternalServerError(errMsg)
	}
}

// NewStringToJsonResponder is a custom httpmock.Responder that takes the status code and a json string body
// and creates a responder for a http mock. This is a useful function when unit testing rest API responses.
func NewStringToJsonResponder(statusCode int, body string) httpmock.Responder {
	response := httpmock.NewBytesResponse(statusCode, []byte(body))
	response.Header.Set(ContentTypeHeaderKey, ApplicationJsonMIMEType)
	return httpmock.ResponderFromResponse(response)
}
