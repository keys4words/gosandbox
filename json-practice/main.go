package main

import (
	"fmt"
	"encoding/json"
	"os"
	"log"
	"io/ioutil"
)

type Ninja struct {
	Name string
	Weapon Weapon
	Level int
}

type Weapon struct {
	Name string
	Damage int
}

//Social block representation
type Social struct {
	Vkontakte string `json:"vkontakte"`
	Facebook string `json:"facebook"`
}

//Internal user representation
type User struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Age int `json:"age"`
	Social Social `json:"social"`
}

//Struct fo representation total slice
type Users struct {
	Users []User `json:"users"`
}

func PrintUser(u *User) {
	fmt.Printf("Name: %s\n", u.Name)
	fmt.Printf("Type: %s\n", u.Type)
	fmt.Printf("Age: %d\n", u.Age)
	fmt.Printf("Social: VK: %s and FB: %s\n", u.Social.Vkontakte, u.Social.Facebook)
}

func main() {
	sFrom := `{"name": "Mikle Dudikoff", "weapon": {"name": "Sword", "damage": 2}, "level": 3}`
	fmt.Println(sFrom)
	var ninja Ninja

	err := json.Unmarshal([]byte(sFrom), &ninja)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ninja.Level)
	}
	
	s2, err := json.Marshal(ninja)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%T", s2)
		fmt.Printf("\n%s", s2)
	}

	fmt.Println("\n*******************")
	jsonFile, err := os.Open("users.json")
	if err != nil {
		log.Fatal(err)
	}
	var users Users
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(byteValue, &users)
	// for _, u := range users.Users {
	// 	fmt.Println("================")
	// 	PrintUser(&u)
	// }
	///unmarshal to interface
	var res map[string]interface{}
	json.Unmarshal(byteValue, &res)
	fmt.Println(res["users"])
	defer jsonFile.Close()
}