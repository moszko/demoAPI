package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"vlserver/domain"
	"vlserver/infrastructure"

	"github.com/gorilla/mux"
)

func main() {
	psqlconn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		infrastructure.Host,
		infrastructure.Port,
		infrastructure.User,
		infrastructure.Password,
		infrastructure.Dbname,
	)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	handleRequests(infrastructure.NewTrademarkRepository(db))
}

func handleRequests(trademarkRepository domain.TrademarkRepository) {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/trademarks/word/{name}", wordTrademark(trademarkRepository))
	router.HandleFunc("/api/v1/trademarks/word/similar/{name}", similarWordTrademarks(trademarkRepository))
	log.Fatal(http.ListenAndServe(":80", router))
}

func wordTrademark(trademarkRepository domain.TrademarkRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		trademarkName := vars["name"]
		trademark, _ := trademarkRepository.GetByName(trademarkName)
		json.NewEncoder(w).Encode(trademark)
	}
}

func similarWordTrademarks(trademarkRepository domain.TrademarkRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		trademarkName := vars["name"]
		trademarks, _ := trademarkRepository.GetSimilarByName(trademarkName)
		json.NewEncoder(w).Encode(trademarks)
	}
}
