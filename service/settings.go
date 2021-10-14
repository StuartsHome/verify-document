package service

import (
	"flag"
)

var listenAddress = flag.String("listen", "", "Address and port to bind HTTP server to")

type HttpSettings struct {
	ListenAddress *string
}

const Address = ":8082"

func SettingsInit() {
	nums := Address
	listenAddress = &nums
}
