package main

import (
	"encoding/json"
	"log"
	"net/http"
	"testagile/services"
	"github.com/gorilla/mux"
	"strconv"
)

var tradingService = services.NewTradingService()

func getAccountBalance(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(services.AccountBalance.Credit)
}

func getTransactionsHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	txs := tradingService.GetTransactionHistory()

	json.NewEncoder(w).Encode(txs)
}

func createTransaction(w http.ResponseWriter, r *http.Request) {

	var txDto services.TransactionBody
	decoder := json.NewDecoder(r.Body)
	decErr := decoder.Decode(&txDto)

	if decErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if txDto.Amount == 0  || (txDto.Type != "CREDIT" &&  txDto.Type !="DEBIT"){ 
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	account, errResp := tradingService.CreateTransaction(txDto)
	if errResp != nil && errResp.Code == services.INSUFICIENT_CREDIT{
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

func getTransactionsById(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	if val, ok := pathParams["id"]; ok {
		_, errNum := strconv.Atoi(val)
		if errNum != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		tx, errTx:= (tradingService.GetTransaction(val))
		if errTx != nil{
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(tx)
	}
}

func main() {
	r := mux.NewRouter()
	
	api := r.PathPrefix("/transactiontrade").Subrouter()
	api.HandleFunc("/accountbalance", getAccountBalance).Methods(http.MethodGet)
	api.HandleFunc("/transactions", getTransactionsHistory).Methods(http.MethodGet)
	api.HandleFunc("/transactions/{id}", getTransactionsById).Methods(http.MethodGet)
	api.HandleFunc("/transactions/create", createTransaction).Methods(http.MethodPost)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	log.Fatal(http.ListenAndServe(":8080", r))
}
