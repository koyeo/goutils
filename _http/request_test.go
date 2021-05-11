package _http

import (
	"testing"
)

func TestNewRequest(t *testing.T) {
	req := NewRequest()
	rsp, err := req.SetContentType("application/json").Post("http://192.168.1.148:1234/rpc/v0", `{
    "jsonrpc": "2.0",
    "method": "Filecoin.ChainHead",
    "params": [],
    "id": 1
}`)
	if err != nil {
		t.Error(err)
		return
	}
	json, err := rsp.SimpleJson()
	if err != nil {
		t.Errorf("parse json error: %s", err)
		return
	}
	t.Log(json.Get("result").Get("Height").Int64())
}
