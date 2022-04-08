package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var apis map[int]string
var c chan map[int]interface{}

// type Result struct {
// 	Success   bool
// 	Timestamp int
// 	Base      string
// 	Date      string
// 	Rates     map[string]float64
// }

// type Error struct {
// 	Success bool
// 	Error   struct {
// 		Code int
// 		Type string
// 		Info string
// 	}
// }

func fetchData(API int) {
	url := apis[API]
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var result map[string]interface{}
			json.Unmarshal([]byte(body), &result)

			var re = make(map[int]interface{})

			switch API {

			case 1:
				if result["success"] == true {
					// fmt.Println(result["rates"].(map[string]interface{})["USD"])
					re[API] = result["rates"].(map[string]interface{})["USD"]
				} else {
					// fmt.Println(result["error"].(map[string]interface{})["info"])
					re[API] = result["error"].(map[string]interface{})["info"]
				}

				c <- re

			case 2:
				if result["main"] != nil {
					// fmt.Println(result["main"].(map[string]interface{})["temp"])
					re[API] = result["main"].(map[string]interface{})["temp"]
				} else {
					// fmt.Println(result["message"])
					re[API] = result["message"]
				}

				c <- re
			}
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}

func main() {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(dir + string(os.PathSeparator) + ".keys")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var data []string
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		log.Fatal(err)
	}

	c = make(chan map[int]interface{})

	apis = make(map[int]string)

	apis[1] = "http://data.fixer.io/api/latest?access_key=" + data[0]
	apis[2] = "http://api.openweathermap.org/data/2.5/weather?id=" + data[1]

	go fetchData(1)
	go fetchData(2)

	for i := 0; i < 2; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Done!!!!!!!!!")
	fmt.Scanln()

	// if resp, err := http.Get(url); err == nil {
	// 	defer resp.Body.Close()
	// 	if body, err := ioutil.ReadAll(resp.Body); err == nil {
	// 		var result Result
	// 		json.Unmarshal([]byte(body), &result)
	// 		if result.Success {
	// 			keys := make([]string, 0, len(result.Rates))
	// 			for k := range result.Rates {
	// 				keys = append(keys, k)
	// 			}
	// 			sort.Strings(keys)
	// 			for _, k := range keys {
	// 				fmt.Println(k, result.Rates[k])
	// 			}
	// 		} else {
	// 			var err Error
	// 			json.Unmarshal([]byte(body), &err)
	// 			fmt.Println(err.Error.Info)
	// 		}

	// 	} else {
	// 		log.Fatal(err)
	// 	}
	// } else {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Done")

}
