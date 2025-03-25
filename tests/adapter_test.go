package tests

import (
	"testing"

	"github.com/KyloRilo/design-patterns-go/patterns/adapter"
)

func TestAdapter(t *testing.T) {
	client := &adapter.Client{}
	mac := &adapter.Mac{}

	client.InsertLightningConnector(mac)

	windowsMachine := &adapter.Windows{}
	windowsMachineAdapter := &adapter.WindowsAdapter{
		WindowsMachine: windowsMachine,
	}

	client.InsertLightningConnector(windowsMachineAdapter)
}
