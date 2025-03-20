package main

import (
	"fmt"
	"testing"
)

type Computer interface {
	InsertLightningPort()
}

type Mac struct{}

func (m *Mac) InsertLightningPort() {
	fmt.Println("Lightning connector inserted into Mac machine")
}

type Windows struct{}

func (w *Windows) InsertUsbPort() {
	fmt.Println("USB connector inserted into Windows machine")
}

type Client struct{}

func (c *Client) InsertLightningConnector(com Computer) {
	fmt.Println("Client inserts connector")
	com.InsertLightningPort()
}

type WindowsAdapter struct {
	windowsMachine *Windows
}

func (w *WindowsAdapter) InsertLightningPort() {
	fmt.Println("Adapter converting")
	w.windowsMachine.InsertUsbPort()
}

func TestAdapter(t *testing.T) {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnector(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowsMachine: windowsMachine,
	}

	client.InsertLightningConnector(windowsMachineAdapter)
}
