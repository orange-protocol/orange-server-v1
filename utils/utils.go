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

package utils

import "regexp"

const (
	PUBLISHED = "published"
	UNSUBMIT  = "unsubmit"
	DRAFT     = "draft"
	FAILED    = "failed"
	VERIFYING = "verifying"
	REVOKING  = "revoking"
	DP_LABELS = "DP"
	AP_LABELS = "AP"
)

func checkMethodName(methodName string) (bool, error) {
	return regexp.MatchString("/^[A-Za-z][0-9A-Za-z_]+$/", methodName)
}

func checkDpName(dpName string) (bool, error) {
	return regexp.MatchString("/^[0-9A-Za-z\\s]+$/", dpName)
}

func checkDesc(descName string) (bool, error) {
	return regexp.MatchString("/^[0-9A-Za-z\\s\\x21-\\x2f\\x3a-\\x40\\x5b-\\x60\\x7B-\\x7F]+$/", descName)
}
