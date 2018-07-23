package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)
func (app *Application) QueryHandler(w http.ResponseWriter, r *http.Request) {
	QueryValue := r.FormValue("key")
	fmt.Println(QueryValue)
	blockData, txnID, err := app.Fabric.QueryOne(QueryValue)

	fmt.Println("#### Query One ###")
	fmt.Printf("%v", blockData)
if err != nil {
	http.Error(w, "Unable to query the blockchain", 500)
}

type KeyData struct {
	Key string `json:"key"`
	Value string `json:"value"`
}
var data KeyData
json.Unmarshal([]byte(blockData), &data)

returnData := &struct {
	ResponseData  KeyData
	TransactionID string
}{
	ResponseData:  data,
	TransactionID: txnID,
}

returnData.TransactionID = txnID

	fmt.Println("######## ResponseData")
	fmt.Printf("%v", returnData)

	renderTemplate(w, r, "query.html", returnData)
}