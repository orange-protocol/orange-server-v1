/*
 *
 *  * Copyright (C) 2022 The orange protocol Authors
 *  * This file is part of The orange library.
 *  *
 *  * The Orange is free software: you can redistribute it and/or modify
 *  * it under the terms of the GNU Lesser General Public License as published by
 *  * the Free Software Foundation, either version 3 of the License, or
 *  * (at your option) any later version.
 *  *
 *  * The orange is distributed in the hope that it will be useful,
 *  * but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  * GNU Lesser General Public License for more details.
 *  *
 *  * You should have received a copy of the GNU Lesser General Public License
 *  * along with The orange.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package service

import (
	"fmt"
	"math/rand"
	"strings"

	mail "github.com/xhit/go-simple-mail/v2"

	"github.com/orange-protocol/orange-server-v1/config"
	"github.com/orange-protocol/orange-server-v1/store"
)

var GlobalEmailService *EmailService

type EmailService struct {
	//MailAddress string
	//SMTPPort int
	//Password string
	MailContent string
	MailSubject string
	SelfMail    string
	Server      *mail.SMTPServer
	//SmtpClient  *mail.SMTPClient
}

func InitEmailService(mailConfig *config.EmailConfig) error {
	es, err := NewEmailService(mailConfig)
	if err != nil {
		return err
	}
	GlobalEmailService = es
	return nil
}

func NewEmailService(mailConfig *config.EmailConfig) (*EmailService, error) {
	server := mail.NewSMTPClient()
	server.Host = mailConfig.Host
	server.Port = mailConfig.SmtpPort
	server.Username = mailConfig.MailAddress
	server.Password = mailConfig.Password
	server.Encryption = mail.EncryptionTLS

	//client, err := server.Connect()
	//if err != nil {
	//	return nil, err
	//}

	return &EmailService{
		MailContent: mailConfig.Content,
		MailSubject: mailConfig.Subject,
		SelfMail:    mailConfig.MailAddress,
		Server:      server,
		//SmtpClient:  client,
	}, nil
}

func (es *EmailService) SendVerificationCode(to string, code string) error {
	email := mail.NewMSG()
	email.SetFrom(fmt.Sprintf("Orange Support <%s>", es.SelfMail))
	email.AddTo(to)
	email.SetSubject(es.MailSubject)
	email.SetBody(mail.TextHTML, strings.ReplaceAll(es.MailContent, "{{VERIFICATION_CODE}}", code))

	client, err := es.Server.Connect()
	if err != nil {
		return err
	}

	err = email.Send(client)
	if err != nil {
		return err
	}
	client.Close()
	return nil
}

func (es *EmailService) GenerateEmailVCode(did, email string) string {
	i := rand.Int31n(1000000)
	return fmt.Sprintf("%06d", i)
}

func (es *EmailService) RequestEmailVCode(did, email string) (string, error) {
	code := es.GenerateEmailVCode(did, email)
	old, err := store.MySqlDB.GetEmailVerificationCode(did, email)
	if err != nil {
		return "", err
	}
	if len(old) == 0 {
		err = store.MySqlDB.AddNewEmailVerificationCode(did, email, code)
	} else {
		err = store.MySqlDB.UpdateEmailVerificationCode(did, email, code)
	}

	return code, err
}
