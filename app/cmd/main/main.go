package main

import (
	"flag"
	"hka-server-login/bootstrap"
	"log"
)

var configPath = flag.String("configPath", ".\\config\\config.json", "Sets absolute path to config files from root")

func main() {
	flag.Parse()

	cfg, err := bootstrap.LoadConfig(*configPath)

	if err != nil {
		log.Fatalln("failed to load config: ", err.Error())
	}

	server := bootstrap.NewServer(*cfg)
	server.InitLdapEndpoints()
	server.Start()
}
