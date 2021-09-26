package main

import (
	"fmt"

	"github.com/stuartshome/verify-document/service"
)

func main() {
	service.SettingsInit()
	fmt.Println("Service starting ...")
	service.HttpRun()
}
