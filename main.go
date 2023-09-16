/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/

//go:generate go run main.go db migrate up
//go:generate go run github.com/ogen-go/ogen/cmd/ogen -allow-remote -package oas -target ./internal/api/oas -clean ./api/oas.yml
//go:generate sqlboiler psql

package main

import "github.com/seanpar203/go-api/cmd"

func main() {
	cmd.Execute()
}
