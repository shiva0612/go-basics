package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

type smsmock struct {
	mock.Mock
}

func (sm *smsmock) sendmsg(msg string) bool {
	fmt.Println("faking sending msg")
	args := sm.Called(msg)
	fmt.Println("args = ", args)
	return args.Bool(0)
}

func TestNotify(t *testing.T) {
	service := new(smsmock)
	service.On("sendmsg", "order placed").Return(true)

	e := ecom{msgservice: service}
	e.msgservice.sendmsg("order placed")
	service.AssertExpectations(t)
}
