package main

import (
	"fmt"
	"golang-webRTC/config"
	"log"
	"net/http"
)

func main() {
	config.AllRooms.Init()

	http.HandleFunc("/create", config.CreateRoomRequestHandler)
	http.HandleFunc("/join", config.JoinRoomRequestHandler)

	log.Println("starting Server on  Port 8000")
	fmt.Println(" ")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal((err))
	}

}
