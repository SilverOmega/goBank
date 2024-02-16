package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(store Storage, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}
	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}
	fmt.Println("New Account => ", acc.Number)
	return acc
}
func seedAccounts(s Storage) {
	seedAccount(s, "Silver", "Omega", "hunter8888")
}
func main() {
	seed := flag.Bool("seed", false, "seed the db")
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Printf("seeding the database")
		seedAccounts(store)
	}
	// seed stuff

	server := NewAPIServer(":3000", store)
	server.Run()
}
