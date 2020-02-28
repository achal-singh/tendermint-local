package model

type TransactionCommit struct {
	ResponseCode float64 `json:"responseCode"`
	Hash         string  `json:"hash"`
	Height       string  `json:"height"`
}

type TransactionData struct {
	Exists bool   `json:"isKeyPresent"`
	Key    string `json:"key"`
	Value  string `json:"value"`
	Error  string `json:"error"`
}

type Configuration struct {
	Port          string
	SEND_TX_URL   string
	QUERY_KEY_URL string
}
