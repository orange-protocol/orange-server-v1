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

	"github.com/orange-protocol/orange-server-v1/config"
	"github.com/orange-protocol/orange-server-v1/did"
	"github.com/orange-protocol/orange-server-v1/log"
	"github.com/orange-protocol/orange-server-v1/utils"
)

type DidService struct {
	Resolvers map[string]did.Resolver
}

var SysDidService *DidService

func InitDidService(conf *config.SysConfig) error {
	m := make(map[string]did.Resolver)
	for _, dc := range conf.DidConf {
		switch dc.Chain {
		case "ont":
			ontresolver, err := did.NewOntResolver(dc)
			if err != nil {
				return err
			}
			m["ont"] = ontresolver

		default:
			return fmt.Errorf("not supported yet")
		}
	}

	SysDidService = &DidService{Resolvers: m}
	return nil
}

func (s *DidService) ValidateSig(did string, msg string, sig string) (bool, error) {
	resolver, err := s.GetResolver(did)
	if err != nil {
		return false, err
	}
	return resolver.ValidateSig(did, 1, msg, sig)
}

func (s *DidService) IssueCredential(did string, content string, commit bool) (string, string, error) {
	resolver, err := s.GetResolver(did)
	if err != nil {
		return "", "", err
	}
	return resolver.CreateCredential(did, 0, content, commit)
}

func (s *DidService) RevokeCredential(did string, cred string) (string, error) {
	resolver, err := s.GetResolver(did)
	if err != nil {
		return "", err
	}
	return resolver.RevokeCredential(did, 0, cred)
}

func (s *DidService) getDidChain(did string) (string, error) {
	//arr := strings.Split(did, ":")
	//if len(arr) != 3 {
	//	return "", fmt.Errorf("not a valid did format:%s", did)
	//}
	//return arr[1], nil

	return utils.GetChainFromDID(did)
}

func (s *DidService) GetResolver(did string) (did.Resolver, error) {
	chain, err := s.getDidChain(did)
	if err != nil {
		return nil, err
	}
	resolver, ok := s.Resolvers[chain]
	if !ok {
		return nil, fmt.Errorf("not a supported did chain:%s", chain)
	}
	return resolver, nil
}

func (ds *DidService) EncryptDataWithDID(data []byte, did string) ([]byte, error) {

	resolver, err := ds.GetResolver(did)
	if err != nil {
		log.Errorf("errors on GetResolver:%s", err.Error())
		return nil, err
	}
	return resolver.EncrypteDataWithDID(data, did)
}

func (ds *DidService) SignData(selfdid string, data []byte) ([]byte, error) {
	resolver, err := ds.GetResolver(selfdid)
	if err != nil {
		log.Errorf("errors on GetResolver:%s", err.Error())
		return nil, err
	}
	return resolver.SignData(data)
}
