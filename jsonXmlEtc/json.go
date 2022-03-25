package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type RequestContent struct {
	User    string
	Message string `json:"msg"`
}

type Request struct {
	Request RequestContent
	Author  string `json:"user"`
}

func LoadAndParseJson() {
	jsonData, err := ioutil.ReadFile("examples/example.json")
	fmt.Println(err)
	var request Request
	fmt.Println(json.Unmarshal(jsonData, &request))
	fmt.Println(request)
}

func LoadAndParseRawMsgToMap() {
	jsonData, _ := ioutil.ReadFile("examples/example.json")
	var objmap map[string]interface{} //unstructured parse to map[string]interface{}
	json.Unmarshal(jsonData, &objmap)
	fmt.Println(objmap)
}

func LoadAndParseRawMsg() {
	jsonData, _ := ioutil.ReadFile("examples/example.json")
	var objmap map[string]json.RawMessage //unstructured parse of specific fields
	json.Unmarshal(jsonData, &objmap)
	fmt.Println(objmap)
	var internalMap map[string]string
	json.Unmarshal(objmap["request"], &internalMap)
	fmt.Println(internalMap)
}

func main() {
	// LoadAndParseJson()
	LoadAndParseRawMsgToMap()
}
