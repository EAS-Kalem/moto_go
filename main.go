package main

import ( 
	"flag"
		 "log"
		 "fmt"
		 "github.com/EAS-Kalem/moto_go/api"
	)

func main () {
	listenAddr := flag.String("listenaddr", ":3000", "the server address")
	flag.Parse()

	server := api.NewServer(*listenAddr)
	fmt.Println("server running on port :3000")
	log.fatal(server.Start())
}