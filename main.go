package main

import (
	"encoding/json"
	"fmt"
	"log"
	"tianiste/blog-aggregator/internal/config"
)

func main() {
	var c config.Config
	c = config.Read()
	err := c.SetUser("lane")
	if err != nil {
		fmt.Errorf("error while setting user name: %s", err)
	}
	c = config.Read()
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bytes))

}
