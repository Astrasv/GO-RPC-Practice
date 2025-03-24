package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Room struct {
	Type     string
	Cost     int
	Numrooms int
	Total    int
}

func main() {
	var reply Room
	var db []Room

	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	err = client.Call("API.GetDB", 1, &db)
	if err != nil {
		log.Println("Error getting database:", err)
	} else {
		fmt.Println("Initial Database:", db)
	}

	err = client.Call("API.BookRoom", "type4", &reply)
	if err != nil {
		log.Println("Error booking room:", err)
	} else {
		fmt.Printf("Booked Room: %+v\n", reply)
	}

	err = client.Call("API.BookRoom", "type4", &reply)
	if err != nil {
		log.Println("Error booking room:", err)
	} else {
		fmt.Printf("Booked Room: %+v\n", reply)
	}

	err = client.Call("API.GetDB", 1, &db)
	if err != nil {
		log.Println("Error getting database:", err)
	} else {
		fmt.Println("Database after booking:", db)
	}

	err = client.Call("API.BookRoom", "type4", &reply)
	if err != nil {
		log.Println("Error booking room:", err)
	}

	err = client.Call("API.GetDB", 1, &db)
	if err != nil {
		log.Println("Error getting database:", err)
	} else {
		fmt.Println("Final Database:", db)
	}
}
