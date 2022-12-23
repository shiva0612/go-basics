package service

import (
	"fmt"
)

type Msgservice interface {
	Sendmsg(msg string) bool
}

// ----------------------------------------------------------------
type sms struct {
}

func (s *sms) Sendmsg(msg string) bool {
	fmt.Println("sending msg through sms")
	return true
}

//----------------------------------------------------------------

type ecom struct {
	msgservice Msgservice
}

func (e *ecom) Notify(msg string) error {
	e.msgservice.Sendmsg(msg)
	return nil
}

//----------------------------------------------------------------
