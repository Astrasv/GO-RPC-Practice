package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"sync"
)

type Room struct {
	Type     string
	Cost     int
	Numrooms int
	Total    int
}

type API struct {
	mu       sync.Mutex
	database []Room
}

func (a *API) GetDB(_ int, reply *[]Room) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	*reply = a.database
	return nil
}

func (a *API) CheckAvail(roomtype string, reply *Room) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	for _, val := range a.database {
		if val.Type == roomtype {
			if val.Numrooms > 0 {
				fmt.Printf("Rooms of type %s are available !!\n", roomtype)
				*reply = val
				return nil
			}
			return fmt.Errorf("no rooms available for type %s", roomtype)
		}
	}
	return fmt.Errorf("room type %s not found", roomtype)
}

func (a *API) BookRoom(roomtype string, reply *Room) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	for idx := range a.database { 
		if a.database[idx].Type == roomtype {
			if a.database[idx].Numrooms > 0 {
				a.database[idx].Numrooms-- 
				*reply = a.database[idx]
				fmt.Printf("Room of type %s booked !!\n", roomtype)
				return nil
			}
			return fmt.Errorf("no rooms available for type %s", roomtype)
		}
	}
	return fmt.Errorf("room type %s not found", roomtype)
}

func (a *API) CancelRoom(roomtype string, reply *Room) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	for idx := range a.database {
		if a.database[idx].Type == roomtype {
			if a.database[idx].Numrooms < a.database[idx].Total {
				a.database[idx].Numrooms++ 
				*reply = a.database[idx]
				fmt.Printf("Room of type %s cancelled !!\n", roomtype)
				return nil
			}
			return fmt.Errorf("all rooms of type %s are already available", roomtype)
		}
	}
	return fmt.Errorf("room type %s not found", roomtype)
}

func main() {
	api := &API{
		database: []Room{
			{"type0", 1000, 10, 10},
			{"type1", 1500, 20, 20},
			{"type2", 2000, 5, 5},
			{"type3", 3000, 3, 3},
			{"type4", 5000, 2, 2},
		},
	}

	err := rpc.Register(api)
	if err != nil {
		log.Fatal("Registering error: ", err)
	}

	rpc.HandleHTTP()

	ln, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Listening error: ", err)
	}

	fmt.Println("Server Listening on port 4040")

	err = http.Serve(ln, nil)
	if err != nil {
		log.Fatal("Serving error: ", err)
	}
}
