package utils

import (
	"github.com/xeipuuv/gojsonschema"
)

func VerifyJSonSchema(jschema, jsonstr string) (bool, error) {
	schema, err := gojsonschema.NewSchema(gojsonschema.NewStringLoader(jschema))
	if err != nil {
		return false, err
	}

	dl := gojsonschema.NewStringLoader(jsonstr)
	r, err := schema.Validate(dl)
	if err != nil {
		return false, err
	}
	return r.Valid(), nil

}
