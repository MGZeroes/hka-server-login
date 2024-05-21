package bootstrap

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"hka-server-login/config"
	"hka-server-login/endpoint"
	"log"
	"net/http"
)

type Server struct {
	cfg       config.Config
	webserver *echo.Echo
}

type ServerImpl interface {
	InitLdapEndpoints()
	Start()
}

func NewServer(cfg config.Config) ServerImpl {
	server := new(Server)
	server.cfg = cfg
	server.webserver = echo.New()

	server.webserver.Use(middleware.Logger())

	return server
}

func (server *Server) InitLdapEndpoints() {
	ldapEndpoint := endpoint.NewLdapEndpoint(server.cfg.LDAP)
	server.webserver.POST(endpoint.CHECK_CREDENTIALS, ldapEndpoint.CheckCredentials)
}

func (server *Server) Start() {
	// TODO http.Server
	var err error

	if !server.cfg.Server.EnableTLS {
		err = server.webserver.Start(server.cfg.Server.Addr)
	} else {
		err = server.webserver.StartTLS(
			server.cfg.Server.Addr,
			server.cfg.Server.TLS.Certificate,
			server.cfg.Server.TLS.Key,
		)
	}

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
