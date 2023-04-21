package main

import (
	"log"

	"github.com/rmaylor-base/iconik_csv/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("error encountered: %s", err)
	}
}
