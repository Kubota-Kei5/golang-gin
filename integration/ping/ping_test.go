package integration

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	utils "golang-gin/integration/integutils"
)

func TestPing(t *testing.T) {
	endpoint := utils.GetEndpoint("/ping")
	res, err := http.Get(endpoint)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)
}