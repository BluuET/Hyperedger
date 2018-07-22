package controllers

import (
	"net/http"
	"encoding/json"
)

func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	helloValue, err := app.Fabric.QueryAll()
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}
	type keyData struct {
		key string 'json:"key"'
		value string 'json:"value"'
	}
	data := &struct {
		key string
	}{
		key: keyValue,
	}
	renderTemplate(w, r, "home.html", data)
}