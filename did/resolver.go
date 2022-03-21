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

package did

type DataField struct {
	Provider_did string `json:"provider_did"`
	Method       string `json:"method"`
	Name         string `json:"name"`
	UserDid      string `json:"user_did"`
	BindData     string `json:"bind_data"`
	Data         string `json:"data"`
}

type AlgorithmField struct {
	Provider_did string `json:"provider_did"`
	Method       string `json:"method"`
	Name         string `json:"name"`
	Result       string `json:"result"`
}

type OrangeCredential struct {
	Data      *DataField      `json:"data"`
	Algorithm *AlgorithmField `json:"algorithm"`
}

type Resolver interface {
	ValidateSig(did string, idx int, msg string, sig string) (bool, error)
	CreateCredential(did string, idx int, content interface{}, commit bool) (string, string, error)
	RevokeCredential(did string, idx int, cred string) (string, error)
	SignData(data []byte) ([]byte, error)
	EncrypteDataWithDID(data []byte, did string) ([]byte, error)
}
