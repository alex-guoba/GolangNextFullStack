package main

import (
	"log"

	"github.com/alex-guoba/gin-clean-template/cmd/cmds"
	_ "github.com/alex-guoba/gin-clean-template/server/docs"
)

func main() {
	if err := cmds.Execute(); err != nil {
		log.Fatal(err)
	}
}
