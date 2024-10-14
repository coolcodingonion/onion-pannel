package main

import (
	"onion-pannel/routes"
	"onion-pannel/utils"
)

func main() {
	routes.InitRouter().Run(utils.HttpPort)
}