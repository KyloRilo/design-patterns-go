package tests

import (
	"testing"

	"github.com/KyloRilo/design-patterns-go/patterns/facade"
)

func TestFacade(t *testing.T) {
	walletFacade := facade.NewWalletFacade("abc", 1234)

	err := walletFacade.AddMoney("abc", 1234, 10)
	if err != nil {
		t.Fatalf("Error: %s\n", err.Error())
	}

	err = walletFacade.DeductMoney("abc", 1234, 5)
	if err != nil {
		t.Fatalf("Error: %s\n", err.Error())
	}
}
