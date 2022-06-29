/*

Facade представляет собой интерфейс со своей подсистемой.
В нашем случае мы хотим добавлять и вычитать деньги, и ничего более.
При этом мы скрываем работу с подсистемными классами.

+ предоставляет простой и урезанный интерфейс к сложной подсистеме
+ когда хотим разложить подсистему на отдельные слои

- может стать "божественным объектом" (объект, который хранит слишком
много или делает слишком много)

*/

package subsystem

import "fmt"

type WalletSystem interface {
	AddMoneyToWallet(accountId string, securityCode int, amount int) error
	DeductMoneyFromWallet(accountId string, securityCode int, amount int) error
}

type walletFacade struct {
	account      *account
	wallet       *wallet
	securityCode *securityCode
	notification *notification
	ledger       *ledger
}

func NewWalletFacade(accountId string, code int) WalletSystem {
	fmt.Println("Starting create account")
	walletSystem := &walletFacade{
		wallet:       NewWallet(),
		account:      NewAccount(accountId),
		securityCode: NewSecurityCode(code),
		notification: &notification{},
		ledger:       &ledger{},
	}

	fmt.Println("Wallet System created")
	return walletSystem
}

func (w *walletFacade) AddMoneyToWallet(accountId string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	if err := w.account.checkAccount(accountId); err != nil {
		return err
	}

	if err := w.securityCode.checkCode(securityCode); err != nil {
		return err
	}

	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountId, "credit", amount)

	return nil
}

func (w *walletFacade) DeductMoneyFromWallet(accountId string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	if err := w.account.checkAccount(accountId); err != nil {
		return err
	}

	if err := w.securityCode.checkCode(securityCode); err != nil {
		return err
	}

	if err := w.wallet.debitBalance(amount); err != nil {
		return err
	}

	w.notification.sendWalletDebitNotification()
	w.ledger.makeEntry(accountId, "credit", amount)

	return nil
}
