package bootstrap

import (
	"latihan/app/api"
	netHttp "net/http"
	"os"
)

func InitHttp() *netHttp.Server {
	echo := api.SetupRouter()
	addr := os.Getenv("address") + ":" + os.Getenv("port")

	server := &netHttp.Server{
		Addr:    addr,
		Handler: echo,
	}

	return server
}
