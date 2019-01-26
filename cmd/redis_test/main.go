package main

import (
	"log"

	"github.com/shjp/shjp-auth/redis"
)

func main() {
	log.Println("----------------------------------------------")
	log.Println("[ Starting Redis Client Testing ]")

	options := redis.Options{
		Network:  "tcp",
		Address:  "localhost:6379",
		Reusable: true,
	}

	client, err := options.NewClient()
	if err != nil {
		log.Fatalf("Error creating client: %s", err)
		return
	}

	if err = client.Set("foo", []byte("bar")); err != nil {
		log.Fatalf("Error setting key/value: %s", err)
		return
	}

	val, err := client.Get("foo")
	if err != nil {
		log.Fatalf("Error getting value: %s", err)
		return
	}

	if string(val) != "bar" {
		log.Fatalf("Expected the value to be 'bar' but got '%s'", val)
	} else {
		log.Println("   Test 1 Passed")
	}

	log.Println("[ Completed Redis Client Testing ]")
	log.Println("----------------------------------------------")
}
