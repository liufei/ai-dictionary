package httputils

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"github.com/privatesquare/bkst-go-utils/utils/logger"
	"strings"
)

const (
	AuthUserKey              = "username"
	AuthPassKey              = "password"
	BasicAuthRequiredErrMsg  = "401 unauthorized: Basic authentication is required"
	BasicAuthFailedErrMsg    = "401 unauthorized: username or password is incorrect"
	authenticationSuccessMsg = "Authenticated successfully"
)

// BasicAuthRequiredError represents a error when basic authentication is not provided while making
// a http request to the server.
type BasicAuthRequiredError struct{}

// Error returns the formatted BasicAuthRequiredError
func (e BasicAuthRequiredError) Error() string {
	return BasicAuthRequiredErrMsg
}

// BasicAuthRequired is a gin middleware for checking if basic authentication is provided in the request
// The method writes the basic auth to the gin context
// The method returns an error if basic authentication is not set
func BasicAuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := GetBasicAuthFromHeader(ctx); err != nil {
			BasicAuthError(ctx)
			return
		}
		ctx.Next()
	}
}

// BasicAuth is a gin middleware for validation if basic authentication is provided in the request
// and the auth user and password matches with the stored user accounts.
// The method writes the basic auth to the gin context
// The method returns an error if basic authentication is not set and
// if the authentication fails to match with a user account.
func BasicAuth(accounts map[string]string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := GetBasicAuthFromHeader(ctx); err != nil {
			BasicAuthError(ctx)
			return
		}

		if val, ok := accounts[ctx.GetString(AuthUserKey)]; ok {
			if val == ctx.GetString(AuthPassKey) {
				logger.Info(authenticationSuccessMsg)
			} else {
				BasicAuthFailed(ctx)
				return
			}
		} else {
			BasicAuthFailed(ctx)
			return
		}
	}
}

// GetBasicAuthFromHeader gets basics authentication from the Authorization header.
func GetBasicAuthFromHeader(ctx *gin.Context) error {

	if ctx.Request.Header.Get(AuthorizationHeaderKey) == "" {
		return BasicAuthRequiredError{}
	}

	auth := strings.SplitN(ctx.Request.Header.Get(AuthorizationHeaderKey), " ", 2)

	if len(auth) != 2 || auth[0] != "Basic" {
		return BasicAuthRequiredError{}
	}

	dAuth, err := base64.StdEncoding.DecodeString(auth[1])
	if err != nil {
		return BasicAuthRequiredError{}
	}

	cred := strings.SplitN(string(dAuth), ":", 2)

	if len(cred) != 2 {
		return BasicAuthRequiredError{}
	}

	ctx.Set(AuthUserKey, cred[0])
	ctx.Set(AuthPassKey, cred[1])

	return nil
}

// BasicAuthError writes an error to the gin context if basic authentication is not provided
func BasicAuthError(ctx *gin.Context) {
	err := errors.UnauthorizedError(BasicAuthRequiredErrMsg)
	ctx.JSON(err.StatusCode, RestErrMsg{Error: err.Message})
	logger.Info(err.Message)
	ctx.Abort()
}

// BasicAuthFailed writes a error to the gin context if basic authentication fails
func BasicAuthFailed(ctx *gin.Context) {
	err := errors.UnauthorizedError(BasicAuthFailedErrMsg)
	ctx.JSON(err.StatusCode, RestErrMsg{Error: err.Message})
	logger.Info(err.Message)
	ctx.Abort()
}
