package main

import ( 
	"flag"
		 "log"
		 "fmt"
		 "github.com/EAS-Kalem/moto_go/api"
		 "github.com/EAS-Kalem/moto_go/storage"
	)

func main () {
	listenAddr := flag.String("listenaddr", ":3000", "the server address")
	flag.Parse()

	store := storage.NewMemoryStorage ()
	server := api.NewServer(*listenAddr, store)
	fmt.Println("server running on port :3000")
	log.fatal(server.Start())
}