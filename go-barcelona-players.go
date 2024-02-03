package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Player struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Position string `json:"position"`
}

var players []Player

func init() {
	players = []Player{
		{Name: "Neymar JR", Age: 28, Position: "Forward"},
		{Name: "Luis Suarez", Age: 33, Position: "Forward"},
		{Name: "Leo Messi", Age: 33, Position: "Forward"},
		{Name: "Iniesta", Age: 35, Position: "Midfielder"},
		{Name: "Busquets", Age: 29, Position: "Midfielder"},
		{Name: "Rakitic", Age: 32, Position: "Midfielder"},
		{Name: "Alba", Age: 32, Position: "Defender"},
		{Name: "Pique", Age: 34, Position: "Defender"},
		{Name: "Mascherano", Age: 35, Position: "Defender"},
		{Name: "Dani Alves", Age: 34, Position: "Defender"},
		{Name: "Ter Stegen", Age: 31, Position: "Goalkeeper"},
	}
}

func getPlayers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(players)
}

func getPlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, player := range players {
		if player.Name == params["name"] {
			json.NewEncoder(w).Encode(player)
			return
		}
	}
	json.NewEncoder(w).Encode(&Player{})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/players", getPlayers).Methods("GET")
	router.HandleFunc("/players/{name}", getPlayer).Methods("GET")
	http.ListenAndServe(":8080", router)
}
