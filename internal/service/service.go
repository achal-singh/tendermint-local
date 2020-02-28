package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"tendermint-local/internal/model"

	"github.com/tkanos/gonfig"
)

var config model.Configuration

func SendTx(key string, value string) model.TransactionCommit {
	config = model.Configuration{}
	err := gonfig.GetConf("./config.json", &config)
	if err != nil {
		fmt.Printf("Error Occurred while loading env %s\n", err)
	}
	var url = strings.Replace(config.SEND_TX_URL, "{key}", key, -1)
	url = strings.Replace(url, "{value}", value, -1)

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return model.TransactionCommit{}
	}
	data, _ := ioutil.ReadAll(response.Body)

	var jsonResult map[string]interface{}
	json.Unmarshal(data, &jsonResult)

	res := jsonResult["result"].(map[string]interface{})
	hash := res["hash"].(string)
	height := res["height"].(string)

	return model.TransactionCommit{ResponseCode: 0, Hash: hash, Height: height}
}

func QueryKey(_key string) model.TransactionData {
	config = model.Configuration{}
	err := gonfig.GetConf("./config.json", &config)
	if err != nil {
		fmt.Printf("Error Occurred while loading env %s\n", err)
	}
	var url = strings.Replace(config.QUERY_KEY_URL, "{key}", _key, -1)

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return model.TransactionData{Error: "HTTP Request Failed"}
	}
	data, _ := ioutil.ReadAll(response.Body)

	var jsonResult map[string]interface{}
	json.Unmarshal(data, &jsonResult)

	result := jsonResult["result"].(map[string]interface{})
	res := result["response"].(map[string]interface{})
	log := res["log"].(string)

	key, errKey := base64.StdEncoding.DecodeString(res["key"].(string))
	if errKey != nil {
		fmt.Println("error:", err)
		return model.TransactionData{Error: "Key Decoding Error!"}
	}
	value, errVal := base64.StdEncoding.DecodeString(res["value"].(string))
	if errVal != nil {
		fmt.Println("error:", err)
		return model.TransactionData{Error: "Value Decoding Error!"}
	}
	if log == "does not exist" {
		return model.TransactionData{Error: _key + "Does not exist"}
	}

	return model.TransactionData{Exists: true, Key: bytes.NewBuffer(key).String(), Value: bytes.NewBuffer(value).String()}
}
