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
