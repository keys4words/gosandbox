package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`   //define tag name before
	Id      int      `xml:"id,attr"` // id is attribute
	Name    string   `xml:"name"`    // name is field
	Origin  []string `xml:"origin"`  //slice of tags
}

func (p Plant) String() string { //redefine string method
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func CreateXml() string {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	out, _ := xml.MarshalIndent(coffee, "", "  ") //format xml with indent
	return string(out)
}

func DecodeXml(input string) {
	var p Plant
	if err := xml.Unmarshal([]byte(input), &p); err != nil {
		panic(err)
	} //decode as json - structure & unmarshal
	fmt.Println(p)
}

func NestedXml() {
	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"` // may point parent
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ := xml.MarshalIndent(nesting, "", "  ")
	fmt.Println(string(out))

	res := &Nesting{}

	xml.Unmarshal(out, res)

	fmt.Println(res)
}

func main() {
	// DecodeXml(CreateXml())
	NestedXml()
}
