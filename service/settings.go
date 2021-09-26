package service

import (
	"flag"
)

var listenAddress = flag.String("listen", "", "Address and port to bind HTTP server to")

type HttpSettings struct {
	ListenAddress *string
}

func SettingsInit() {
	nums := ":8082"
	listenAddress = &nums
}
