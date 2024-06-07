package twitchhelper

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendOauthRequest(t *testing.T) {
	t.Run("TestSendOauthRequest", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		defer ts.Close()

		oauthResponse := SendOauthRequest()
		assert.NotNil(t, oauthResponse)
	})
}
