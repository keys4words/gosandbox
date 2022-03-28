package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v3"
)

type SlicesTags []string

func (tags *SlicesTags) UnmarshalYAML(value *yaml.Node) error { //custom yaml
	if value != nil {
		*tags = strings.Split(value.Value, ",")
	}
	return nil
}

type Messages struct {
	Tags SlicesTags `yaml:"tags"`
}

type Subs struct {
	Messages Messages `yaml:"messages"`
}

func ParseYamlWithCustomStruct() {
	config := &Subs{}

	yamlFile, err := ioutil.ReadFile("examples/yaml_subs_example.yml")
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
	ParseYamlWithCustomStruct()
}
