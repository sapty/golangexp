package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	logging "github.com/op/go-logging"
)

func main() {



	var logger_s = logging.MustGetLogger("Vendor-Name-Authentication-Log")

	valueMap := make(map[string]interface{})

	detMap := make(map[string]interface{})

	// detTemplate := `[""selector":{"DCDocumentData":{"$exists":false}}"]
	//	`

	detString := "2018-08-29-09.08.45"

	detMap["fcn"] = "VQP_SEARCH"

	A := "A"

	B := "B"

	args := []string{detString, A, B}

	detMap["args"] = args

	valueMap["queryReq"] = detMap

	jsonValue, _ := json.Marshal(valueMap)

	serverURL := "https://bpspreprod.us-south.containers.mybluemix.net/spv11/chaincode/query"

	resp, err := http.Post(serverURL, "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		logger_s.Debugf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		logger_s.Debugf(string(data))
	}

	logger_s.Debugf("test")

}

