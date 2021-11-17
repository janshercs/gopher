package main

import (
	"log"
	"net/http"

	"poker"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.Alerter), store)
	server, _ := poker.NewPlayerServer(store, game)

	log.Fatal(http.ListenAndServe(":5000", server))
}
