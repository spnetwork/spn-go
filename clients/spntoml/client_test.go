package spntoml

import (
	"net/http"
	"strings"
	"testing"

	"github.com/spn/go/support/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientURL(t *testing.T) {
	//HACK:  we're testing an internal method rather than setting up a http client
	//mock.

	c := &Client{UseHTTP: false}
	assert.Equal(t, "https://spn.org/.well-known/spn.toml", c.url("spn.org"))

	c = &Client{UseHTTP: true}
	assert.Equal(t, "http://spn.org/.well-known/spn.toml", c.url("spn.org"))
}

func TestClient(t *testing.T) {
	h := httptest.NewClient()
	c := &Client{HTTP: h}

	// happy path
	h.
		On("GET", "https://spn.org/.well-known/spn.toml").
		ReturnString(http.StatusOK,
			`FEDERATION_SERVER="https://localhost/federation"`,
		)
	stoml, err := c.GetSpnToml("spn.org")
	require.NoError(t, err)
	assert.Equal(t, "https://localhost/federation", stoml.FederationServer)

	// spn.toml exceeds limit
	h.
		On("GET", "https://toobig.org/.well-known/spn.toml").
		ReturnString(http.StatusOK,
			`FEDERATION_SERVER="https://localhost/federation`+strings.Repeat("0", SpnTomlMaxSize)+`"`,
		)
	stoml, err = c.GetSpnToml("toobig.org")
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "spn.toml response exceeds")
	}

	// not found
	h.
		On("GET", "https://missing.org/.well-known/spn.toml").
		ReturnNotFound()
	stoml, err = c.GetSpnToml("missing.org")
	assert.EqualError(t, err, "http request failed with non-200 status code")

	// invalid toml
	h.
		On("GET", "https://json.org/.well-known/spn.toml").
		ReturnJSON(http.StatusOK, map[string]string{"hello": "world"})
	stoml, err = c.GetSpnToml("json.org")

	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "toml decode failed")
	}
}
