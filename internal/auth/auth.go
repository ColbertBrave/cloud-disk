package auth

import (
	"io"
	"net/http"

	"github.com/cloud-disk/internal/log"
	"github.com/cloud-disk/app/common"
	"github.com/cloud-disk/internal/config"
)

var Auth HmacAuthenticator

func InitAuth() {
	Auth.SecretKey = []byte(config.AppCfg.AuthCfg.SecretKey)
}

func VerifyRequest(authenticator Authenticator, request *http.Request) error {
	sign, isOk := request.Header["Authorization"]
	if !isOk || len(sign) == 0 {
		return common.ErrNoAuthorization
	}

	bytes, err := io.ReadAll(request.Body)
	if err != nil {
		log.Error("read the request body error:%s", err)
		return err
	}

	return authenticator.Verify(string(bytes), sign[0])
}
