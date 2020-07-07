package main

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	"net/http"
)

type SayHelloImp struct {
}

func (imp *SayHelloImp) EchoHello(name string, greeting *string) (int32, error) {
	*greeting = "hello " + name
	return 0, nil
}

func main() {
	mux := &tars.TarsHttpMux{}
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("hello world, golang"))
	})

	serverConfig := tars.GetServerConfig()
	fmt.Println(serverConfig)

	tars.AddHttpServant(mux, serverConfig.App+"."+serverConfig.Server+".HttpObj") //Register http server
	tars.Run()
}
