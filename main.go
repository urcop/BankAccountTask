package main

import "github.com/urcop/BankAccountTask/server"

func main() {
	server.NewAccountServer().Start()
}
