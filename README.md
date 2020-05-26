Run Trading Manager:

Clone the project to an specific location: 
Open a terminal and change the directory to the specific location.

Run:

$go build main.go
$./main

After that you can open your browser at http://localhost:8080
This Url will show the transactions that you have.
To create a new transaction, go to the browser console and make a post request:

fetch("transactiontrade/transactions/create",{
  method: "POST",
  headers: {
    'Content-Type': 'application/json',
  },
  body : JSON.stringify({amount : 10000, type : 'DEBIT'})})
  
  Or use your prefered client to do the request for example Postman:
  POST with raw body: {"amount" : 50000.00, "type" : "DEBIT"}
  
To get a particular transaction use the browser. open the following URL replacing the id placeholder with the transaction id that you need to query.
http://localhost:8080/transactiontrade/transactions/{id}

To get the account balance use the browser with the following query
http://localhost:8080/transactiontrade/accountbalance

To refresh the transaction history, refresh the browser window at http://localhost:8080:
