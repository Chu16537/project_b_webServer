package main

import (
	"fmt"
	"project_b_webserver/conf"
	"project_b_webserver/webserver"
)

func main() {

	env := new(conf.Config)
	err := conf.ReadEnv("./env.yaml", env)

	if err != nil {
		fmt.Println(err)
		return
	}

	webserver.New(env.WebServer)
}
