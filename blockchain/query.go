package blockchain

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

// QueryHello query the chaincode to get the state of hello
func (setup *FabricSetup) QueryAll() (string, error) {

	// Prepare arguments
	var args []string
	args = append(args, "invoke")
	args = append(args, "query")
	args = append(args, "all")

	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2])}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	return string(response.Payload), nil
}


	// QueryOne query the chaincode to get the record of a specific Key
	func (setup *FabricSetup) QueryOne(value string) (string, error) {

		// Prepare arguments
		var args []string
		args = append(args, "invoke")
		args = append(args, "queryone")
		args = append(args, value)

		response, err := setup.client.Query(chclient.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2])}})
		if err != nil {
			return "", fmt.Errorf("failed to query: %v", err)
		}

		return string(response.Payload), nil
	}
}

func (app *Application) QueryHandler(w http.ResponseWriter, r *http.Request) {

	QueryValue := r.FormValue("key")
	fmt.Println(QueryValue)
	blockData, err := app.Fabric.QueryOne(QueryValue)

	fmt.Println("#### Query One ###")
	fmt.Printf("%v", blockData)

	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}

	type CarData struct {
		key   string `json:"key"`
		value  string `json:"value"`
	}

	var data CarData
	json.Unmarshal([]byte(blockData), &data)

	returnData := &struct {
		ResponseData  keyData
		TransactionID string
	}{
		ResponseData: data,
	}

	fmt.Println("######## ResponseData")
	fmt.Printf("%v", returnData)

	renderTemplate(w, r, "query.html", returnData)
}
