package data

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestHttpDataProvider_CheckUrl(t *testing.T) {
	httpDataProvider := &HttpDataProvider{
		did:    "",
		entris: nil,
		client: &http.Client{Timeout: 60 * time.Minute},
	}
	path := "www.baidu.com"
	err := httpDataProvider.CheckUrl(path)
	assert.Nil(t, err)

	flag, err := httpDataProvider.CheckHttpStatus("GET", path, "")
	assert.Nil(t, err)
	assert.Equal(t, true, flag)
}
