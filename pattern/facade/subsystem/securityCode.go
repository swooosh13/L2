package subsystem

import (
	"errors"
	"fmt"
)

type securityCode struct {
	code int
}

func NewSecurityCode(code int) *securityCode {
	return &securityCode{code}
}

func (s *securityCode) checkCode(incomingCode int) error {
	if s.code != incomingCode {
		return errors.New("security code is incorrect")
	}

	fmt.Println("SecurityCode verified successfully")
	return nil
}
