package main

import "fmt"

type client struct {
}

//Ожидает на вход реализацию интерфейса
func (c *client) insertLightningConnectorIntoComputer(com computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.insertIntoLightningPort()
}
