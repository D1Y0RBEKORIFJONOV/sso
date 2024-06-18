package main

import (
	"fmt"
	"sso/internal/config"
)

func main() {
	cnf := config.MustLoad()
	fmt.Println(cnf)
	//TODO Insalization logger
	//TODO Insalization (app);
	//TODO run grpc-server

}
