package twitchhelper

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
)

// mockRoundTripper implements http.RoundTripper
type mockRoundTripper struct {
	fn func(*http.Request) *http.Response
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.fn(req), nil
}

func newMockClient(fn func(*http.Request) *http.Response) *http.Client {
	return &http.Client{
		Transport: &mockRoundTripper{fn},
	}
}

func TestGetHttpClient(t *testing.T) {
	client := GetHttpClient()
	if client == nil {
		t.Fatal("Expected HTTP client to be initialized, got nil")
	}
}
func TestSetHttpClient(t *testing.T) {
	client := &http.Client{}
	SetHttpClient(client)
	if HTTPClient != client {
		t.Fatal("Expected HTTP client to be set, but it was not")
	}
}

func TestSendOAuthRequest(t *testing.T) {
	mockResponse := `{
		"access_token": "mock-token",
		"expires_in": 3600,
		"token_type": "bearer"
	}`

	os.Setenv("TWITCH_CLIENT_ID", "fake-client-id")
	os.Setenv("TWITCH_CLIENT_SECRET", "fake-client-secret")

	expectedUrl := "https://id.twitch.tv/oauth2/token?client_id=fake-client-id&client_secret=fake-client-secret&grant_type=client_credentials"

	// Set up the mock HTTP client
	SetHttpClient(newMockClient(func(req *http.Request) *http.Response {
		if req.URL.String() != expectedUrl {
			t.Fatalf("Unexpected URL: %s instead of %s", req.URL.String(), expectedUrl)
		}

		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(mockResponse)),
			Header:     make(http.Header),
		}
	}))

	res := SendOAuthRequest()

	if res.AccessToken != "mock-token" {
		t.Errorf("Expected access token to be 'mock-token', got '%s'", res.AccessToken)
	}
}

func TestGetTop100ChannelsByStreamViewCount(t *testing.T) {
	callCount := 0
	HTTPClient = newMockClient(func(req *http.Request) *http.Response {
		callCount++
		if strings.Contains(req.URL.String(), "oauth2/token") {
			return &http.Response{
				StatusCode: 200,
				Body: io.NopCloser(strings.NewReader(`{
					"access_token": "mock-token",
					"expires_in": 3600,
					"token_type": "bearer"
				}`)),
				Header: make(http.Header),
			}
		}

		// Simulate stream data
		return &http.Response{
			StatusCode: 200,
			Body: io.NopCloser(bytes.NewBufferString(`{
				"data": [
					{"user_name": "StreamerA"},
					{"user_name": "StreamerB"}
				]
			}`)),
			Header: make(http.Header),
		}
	})

	os.Setenv("TWITCH_CLIENT_ID", "id")
	os.Setenv("TWITCH_CLIENT_SECRET", "secret")

	streamers := GetTop100ChannelsByStreamViewCount()
	expected := []string{"StreamerA", "StreamerB"}

	if len(streamers) != len(expected) {
		t.Fatalf("Expected %d streamers, got %d", len(expected), len(streamers))
	}

	for i, name := range expected {
		if streamers[i] != name {
			t.Errorf("Expected streamer[%d] = %s, got %s", i, name, streamers[i])
		}
	}
}
