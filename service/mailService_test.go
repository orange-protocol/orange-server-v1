package service

import (
	"fmt"
	"testing"
)

func TestEmailService_GenerateEmailVCode(t *testing.T) {
	es := &EmailService{}

	for i := 0; i < 10; i++ {
		fmt.Printf("%s\n", es.GenerateEmailVCode("", ""))
	}
}
