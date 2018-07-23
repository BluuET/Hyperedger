package controllers

import (
	"net/http"
	"encoding/json"
)

func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	blockData, err := app.Fabric.QueryAll()
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}
	type keyData struct {
		key string 'json:"key"'
		value string 'json:"value"'
	}
	var data []keyData
	json.Unmarshal([]byte(blockData), &data)

	returndData := &struct {
		ResponseData []keyData
	}{
		ResponseData: data,
	}
	renderTemplate(w, r, "home.html", returnData)
}