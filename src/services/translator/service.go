package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"GoBun/srb2kart/network"
)

func serverInfo(w http.ResponseWriter, r *http.Request) {
	serverInfo, _, err := network.GetServerInfo("localhost:5029")
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := json.Marshal(serverInfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(data)
}

func playerInfo(w http.ResponseWriter, r *http.Request) {
	_, playerInfo, err := network.GetServerInfo("localhost:5029")
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := json.Marshal(playerInfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(data)
}

func main() {
	http.HandleFunc("/serverinfo", serverInfo)
	http.HandleFunc("/playerinfo", playerInfo)

	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}
