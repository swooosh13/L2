package subsystem

import "fmt"

type ledger struct{}

func (s *ledger) makeEntry(accountId, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountId, txnType, amount)
}
