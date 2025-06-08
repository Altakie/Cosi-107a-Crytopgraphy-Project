package main

import (
	"log"
	"net/http"
	"strconv"
)
// Need to have a server struct and an initialization func


type Server struct {
	PrivateKey int
	PublicVar int
	N int
	Port string
	SharedSecret int
}

func InitServer(s Server) bool {

	port := s.Port

	router := http.NewServeMux()
	
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Port: " + port + " \n"))
		w.Write([]byte("Public Variable: " + strconv.Itoa(s.PublicVar) + " \n"))
		w.Write([]byte("Upper Bound: " + strconv.Itoa(s.N) + " \n"))
		w.Write([]byte("Private Key: " + strconv.Itoa(s.PrivateKey) + " \n"))
		w.Write([]byte("Shared Secret: " + strconv.Itoa(s.SharedSecret) + " \n"))
	})

	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

    // log is better than fmt for real programs
	log.Println("Starting server on port http://127.0.0.1:" + port)
	server.ListenAndServe()

	return true
}

func main() {
	alice := Server { 
		Port: "8081",
		PrivateKey: -1,
		PublicVar: -1,
		N: -1,
		SharedSecret: -1,
	}
	bob := Server { 
		Port: "8082",
		PrivateKey: -1,
		PublicVar: -1,
		N: -1,
		SharedSecret: -1,
	}

	InitServer(alice)
	log.Println("hello")
	InitServer(bob)
}
