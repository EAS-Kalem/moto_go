package main

import ( 
	"flag"
		 "log"
		 "fmt"
		 "api/server.go"
	)

func main () {
	listenAddr := flag.String("listenaddr", ":3000", "the server address")
	flag.Parse()

	server := api.NewServer(*listenAddr)
	fmt.Println("server running on port :3000")
	log.fatal(server.Start())
}