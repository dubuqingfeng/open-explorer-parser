package rpc

import (
	"net/http"
	"io/ioutil"
	log "github.com/sirupsen/logrus"
	"encoding/json"
	"time"
	"bytes"
	"fmt"
	"github.com/pkg/errors"
)

type RpcClient struct {
	address    string
	user       string
	passwd     string
	ssl        bool
	httpClient *http.Client
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

type rpcResult struct {
	response *http.Response
	err      error
}

// Need to force open authentication
func newRPCClient(address string, user string, passwd string, ssl bool) *RpcClient {
	httpClient := &http.Client{}
	client := &RpcClient{address: address, user: user, passwd: passwd, httpClient: httpClient}
	return client
}

// Do Request and timeout limit
func (this *RpcClient) DoRequest(timer *time.Timer, req *http.Request) (*http.Response, error) {
	done := make(chan rpcResult, 1)
	go func() {
		resp, err := this.httpClient.Do(req)
		done <- rpcResult{resp, err}
	}()
	select {
	case r := <-done:
		return r.response, r.err
	case <-timer.C:
		return nil, errors.New("Timeout")
	}
}

// need add timeout limit
func (this *RpcClient) call(method string, params interface{}) (response rpcResponse, err error) {
	// build http request
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("rpc client base call")

	request := rpcRequest{time.Now().UnixNano(), "2.0", method, params}
	payload, err := json.Marshal(request)
	fmt.Println(request)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", this.address, bytes.NewBuffer(payload))
	if err != nil {
		log.Error(err)
		return
	}

	//req.SetBasicAuth(this.user, this.passwd)
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Accept", "application/json")
	fmt.Println(this.user)
	fmt.Println(this.passwd)
	// Timer
	resp, err := this.DoRequest(timer, req)

	if err != nil {
		log.Error(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &response)
	return
}

// get sync status, keep sync
type SyncStatus struct {
}
