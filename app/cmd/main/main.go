package main

import (
	"flag"
	"github.com/labstack/echo/v4"
	"hka-server-login/bootstrap"
	"log"
)

var configPath = flag.String("configPath", ".\\config\\config.json", "Sets absolute path to config files from root")

func main() {
	flag.Parse()

	var _ = echo.CONNECT

	cfg, err := bootstrap.LoadConfig(*configPath)

	if err != nil {
		log.Fatalln("failed to load config: ", err.Error())
	}

	server := bootstrap.NewServer(*cfg)

	//ldapEndpoint := endpoint.NewLdapEndpoint(cfg)
	//server.AddEndpoint(endpoint.CHECK_CREDENTIALS, ldapEndpoint.CheckCredentials)
	server.InitLdapEndpoints()

	server.Start()
}
