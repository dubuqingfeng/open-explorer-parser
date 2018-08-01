package rpc

import (
	"net/http"
	"io/ioutil"
	log "github.com/sirupsen/logrus"
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

type RpcClient struct {
	address string
	user    string
	passwd  string
	ssl     bool
}

type rpcResponse struct {
	ID      int             `json:"id"`
	JSONRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   *error          `json:"error"`
}

type rpcRequest struct {
	// ID: time.Now().UnixNano()
	ID      int64       `json:"id"`
	JSONRPC string      `json:"jsonrpc"` // value: "2.0"
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

func newRPCClient(address string, user string, passwd string, ssl bool) *RpcClient {
	client := &RpcClient{address: address, user: user, passwd: passwd}
	return client
}

// need add timeout limit
func (this *RpcClient) call(method string, params interface{}) (response rpcResponse, err error) {
	// build http request
	fmt.Println("rpc client base call")
	request := rpcRequest{time.Now().UnixNano(), "2.0", method, params}
	payload, err := json.Marshal(request)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", this.address, bytes.NewBuffer(payload))
	if err != nil {
		log.Error(err)
		return
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Error(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &response)
	return
}

// get sync status, keep sync
type SyncStatus struct {
}
