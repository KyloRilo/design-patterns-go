package facade

import "fmt"

type Account struct {
	name string
}

func (a *Account) checkAccount(name string) error {
	if a.name != name {
		return fmt.Errorf("Account Name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}

func newAccount(name string) *Account {
	return &Account{
		name: name,
	}
}

type SecurityCode struct {
	code int
}

func (s *SecurityCode) checkCode(incomingCode int) error {
	if s.code != incomingCode {
		return fmt.Errorf("Security Code is incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		code: code,
	}
}

type Wallet struct {
	balance int
}

func (w *Wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
	return
}

func (w *Wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("Balance is not sufficient")
	}
	fmt.Println("Wallet balance is Sufficient")
	w.balance = w.balance - amount
	return nil
}

func newWallet() *Wallet {
	return &Wallet{
		balance: 0,
	}
}

type Ledger struct {
	entries []string
}

func (s *Ledger) makeEntry(accountID, txnType string, amount int) {
	entry := fmt.Sprintf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountID, txnType, amount)
	s.entries = append(s.entries, entry)

	fmt.Print(entry)
}

type Notification struct{}

func (n *Notification) sendCreditNotify() {
	fmt.Println("Sending wallet credit notification")
}

func (n *Notification) sendDebitNotify() {
	fmt.Println("Sending wallet debit notification")
}

type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
	notification *Notification
	ledger       *Ledger
}

func NewWalletFacade(accountId string, code int) *WalletFacade {
	fmt.Println("Creating wallet")

	walletFacade := &WalletFacade{
		account:      newAccount(accountId),
		securityCode: newSecurityCode(code),
		wallet:       newWallet(),
		notification: &Notification{},
		ledger:       &Ledger{},
	}

	fmt.Println("Wallet created")
	return walletFacade
}

func (w *WalletFacade) AddMoney(accountId string, code int, amount int) error {
	fmt.Println("Adding money to wallet")
	err := w.account.checkAccount(accountId)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(code)
	if err != nil {
		return err
	}

	w.wallet.creditBalance(amount)
	w.notification.sendCreditNotify()
	w.ledger.makeEntry(accountId, "credit", amount)
	return nil
}

func (w *WalletFacade) DeductMoney(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.sendDebitNotify()
	w.ledger.makeEntry(accountID, "debit", amount)
	return nil
}
