package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	blockData, err := app.Fabric.QueryAll()
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}
	fmt.Println(blockData)
	type KeyData struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	var data []KeyData
	json.Unmarshal([]byte(blockData), &data)

	returnData := &struct {
		ResponseData []KeyData
	}{
		ResponseData: data,
	}
	renderTemplate(w, r, "home.html", returnData)
}
