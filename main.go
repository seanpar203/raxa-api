/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/

//go:generate go run github.com/ogen-go/ogen/cmd/ogen -allow-remote -target ./internal/api/oas -clean ./api/oas.yml
package main

import "github.com/seanpar203/go-api/cmd"

func main() {
	cmd.Execute()
}
