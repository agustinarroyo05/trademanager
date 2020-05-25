package services

import (
	"fmt"
	"sync"
	"time"
)

type TxType int

const (
	CREDIT TxType = iota
	DEBIT
	INITIAL_CREDIT float64 = 100000
)

const (
	INSUFICIENT_CREDIT int = 0
)

type Errors struct{
	Code int
	Description string 
}

type TradingService struct {
	txId               int
	mu                 sync.RWMutex
	transactionHistory map[string]Transaction
}
type Account struct {
	AccoudnId   float64
	AccountName string
	Credit      float64
}
type Transaction struct {
	Id           string `json:"id"`
	Amount       float64 `json:"amount"`
	TxType       TxType `json:"type"`
	EfectiveDate time.Time `json:"efectiveDate"`
}

type TransactionBody struct {
	Amount float64 `json : "amount"`
	Type string `json : "type"`
}

func NewTradingService() *TradingService {
	t := new(TradingService)
	t.transactionHistory = make(map[string]Transaction)
	return t
}

var AccountBalance = Account{1, "INITIAL_ACCOUNT", INITIAL_CREDIT} 
var txId int

func (t TradingService) CreateTransaction(txDTO TransactionBody) (*Account, *Errors) {

	t.mu.RLock()
	newCredit := AccountBalance.Credit
	var tx = Transaction{}
	switch txDTO.Type {
	case "DEBIT":
		if AccountBalance.Credit-txDTO.Amount < 0 {

			var e Errors = Errors{INSUFICIENT_CREDIT, "Insuficient Credit" }
			return nil, &e
		}
		newCredit = AccountBalance.Credit - txDTO.Amount
		tx.TxType = DEBIT
		break
	case "CREDIT":
		newCredit = AccountBalance.Credit + txDTO.Amount
		tx.TxType = CREDIT
		break
	}
	AccountBalance.Credit = newCredit
	tx.Amount = txDTO.Amount
	txId++
	tx.Id = fmt.Sprintf("%d", txId)
	tx.EfectiveDate = time.Now()
	t.transactionHistory[tx.Id] = tx
	t.mu.RUnlock()
	return &AccountBalance, nil

}

func (t TradingService) GetTransactionHistory() []Transaction {

	var copyTransactionHistroy []Transaction
	t.mu.Lock()
	for _, value := range t.transactionHistory {
		copyTransactionHistroy = append(copyTransactionHistroy, value)
	}
	t.mu.Unlock()
	return copyTransactionHistroy
}

func (t TradingService) GetTransaction(id string) *Transaction {
	t.mu.RLock()
	tx := t.transactionHistory[id]
	t.mu.RUnlock()
	return &tx
}
