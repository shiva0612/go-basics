package service

import (
	"fmt"
)

type msgservice interface {
	sendmsg(msg string) bool
}

// ----------------------------------------------------------------
type sms struct {
}

func (s *sms) sendmsg(msg string) bool {
	fmt.Println("sending msg through sms")
	return true
}

//----------------------------------------------------------------

type ecom struct {
	msgservice msgservice
}

func (e *ecom) Notify(msg string) error {
	e.msgservice.sendmsg(msg)
	return nil
}

//----------------------------------------------------------------
