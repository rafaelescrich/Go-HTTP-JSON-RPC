package main

// Copyright 2011-2014 ThePiachu. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/rafaelescrich/Go-HTTP-JSON-RPC"
)

func main() {
	res, err := httpjsonrpc.Call("http://user:pass@127.0.0.1:8332", "getinfo", 1, []interface{}{})
	if err != nil {
		log.Fatalf("Err: %v", err)
	}
	log.Println(res)
}

func tmp() {
	data, err := json.Marshal(map[string]interface{}{
		"method": "getinfo",
		"id":     1,
		"params": []interface{}{},
	})
	if err != nil {
		log.Fatalf("Marshal: %v", err)
	}
	resp, err := http.Post("http://user:pass@127.0.0.1:8332",
		"application/json", strings.NewReader(string(data)))
	if err != nil {
		log.Fatalf("Post: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ReadAll: %v", err)
	}
	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Println(result)
}
