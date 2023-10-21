package main

import (
	"fmt"

	"github.com/Ctere1/golang-api/pkgs/api"
	"github.com/Ctere1/golang-api/pkgs/configs"
	"github.com/Ctere1/golang-api/pkgs/storage"
)

func main() {

	fmt.Print(`
	╔═╗┌─┐  ╔═╗╔═╗╦
	║ ╦│ │  ╠═╣╠═╝║
	╚═╝└─┘  ╩ ╩╩  ╩ 

`)
	//load configs from file
	configs.LoadConfigs()

	//database configuraion
	storage.Initialize()

	//start the router
	api.StartRouter()
}
