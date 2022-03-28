package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type BuildConf struct {
	Definitions map[string]interface{} `yaml:"definitions"`
	Pipelines   map[string]interface{} `yaml:"pipelines"`
}

type Conf struct {
	Hits int64 `yaml:"hits"`
	Time int64 `yaml:"time"`
}

func GetConf() {
	config := &Conf{}

	yamlFile, err := ioutil.ReadFile("examples/example.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	fmt.Println(config)
}

func CreateConf() {
	config := &Conf{
		Hits: 11,
		Time: 1635859067,
	}

	out, _ := yaml.Marshal(config)

	fmt.Println(string(out))
}

func ReadBigConf() {
	config := &BuildConf{} //read yaml into map[string]interface{}

	yamlFile, err := ioutil.ReadFile("examples/anchor_example.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	fmt.Println(config)
}

func main() {
	// GetConf()
	// CreateConf()
	ReadBigConf()
}
