package service

import (
	"shiva/testing-framework/using_mockery/mocks"
	"testing"
)

func TestNotify(t *testing.T) {
	service := new(mocks.Msgservice)
	service.On("Sendmsg", "order placed").Return(true)

	e := ecom{msgservice: service}
	e.msgservice.Sendmsg("order placed")
	service.AssertExpectations(t)
}
