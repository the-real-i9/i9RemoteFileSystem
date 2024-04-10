package main

import (
	"i9Packages/i9helpers"
	appprocs "i9rfs/server/procs/app"
	authprocs "i9rfs/server/procs/auth"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	if err := i9helpers.LoadEnv(".env"); err != nil {
		log.Fatal(err)
	}

	if err := i9helpers.InitDBPool(); err != nil {
		log.Fatal(err)
	}

	authSignup := new(authprocs.AuthSignup)
	auth := new(authprocs.Auth)
	rfs := new(appprocs.RFS)

	rpc.Register(auth)
	rpc.Register(authSignup)
	rpc.Register(rfs)

	rpc.HandleHTTP()

	listn, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	go http.Serve(listn, nil)
}
