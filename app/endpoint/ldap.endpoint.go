package endpoint

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/labstack/echo/v4"
	"hka-server-login/config"
	"hka-server-login/dto"
	"hka-server-login/service"
	"net/http"
)

const CHECK_CREDENTIALS string = "/api/check_credentials"

type LdapEndpoint struct {
	cfg config.LdapConfig
}

type LdapEndpointImpl interface {
	CheckCredentials(echo.Context) error
}

func NewLdapEndpoint(cfg config.LdapConfig) LdapEndpointImpl {
	return &LdapEndpoint{cfg}
}

func (endpoint *LdapEndpoint) CheckCredentials(c echo.Context) error {

	c.Logger().Print()

	// Bind the request body to the credentials struct
	var creds dto.Credentials
	if err := c.Bind(&creds); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	ldapConn, err := ldap.DialURL(fmt.Sprintf("%s:%d", endpoint.cfg.Host, endpoint.cfg.Port))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to connect to LDAP server")
	}
	defer ldapConn.Close()

	ldapService := service.NewLdapService(endpoint.cfg, ldapConn)

	user, err := ldapService.SearchUser(creds, creds.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp := dto.User{}
	resp.FromModel(user)
	return c.JSON(http.StatusOK, resp)
}
