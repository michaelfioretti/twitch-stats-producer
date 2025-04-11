package twitchchatparser_test

import (
	"testing"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
	models "github.com/michaelfioretti/twitch-stats-producer/internal/models/proto"
	"github.com/michaelfioretti/twitch-stats-producer/internal/twitchchatparser"
	"github.com/stretchr/testify/assert"
)

func TestParseTwitchMessage(t *testing.T) {
	message := twitch.PrivateMessage{
		User: twitch.User{
			Name: "testuser",
		},
		Message: "Hello, world!",
		Tags: map[string]string{
			"badges":     "subscriber/1,moderator/1",
			"bits":       "100",
			"subscriber": "1",
			"mod":        "1",
			"color":      "#1E90FF",
			"room-id":    "123456",
		},
		Channel: "#testchannel",
		Time:    testTime(),
	}

	expected := &models.TwitchMessage{
		Username:   "testuser",
		Channel:    "#testchannel",
		Message:    "Hello, world!",
		Badges:     []string{"subscriber/1", "moderator/1"},
		Bits:       100,
		Mod:        1,
		Subscribed: 1,
		Color:      "#1E90FF",
		RoomID:     "#123456",
		CreatedAt:  int32(testTime().Unix()),
	}

	result := twitchchatparser.ParseTwitchMessage(message)
	assert.Equal(t, expected, result)
}

func TestParseTwitchMessage_ActionMessage(t *testing.T) {
	message := twitch.PrivateMessage{
		User: twitch.User{
			Name: "testuser",
		},
		Message: "\x01ACTION waves hello\x01",
		Tags: map[string]string{
			"badges":     "vip/1",
			"bits":       "0",
			"subscriber": "0",
			"mod":        "0",
			"color":      "#FF4500",
			"room-id":    "654321",
		},
		Channel: "#anotherchannel",
		Time:    testTime(),
	}

	expected := &models.TwitchMessage{
		Username:   "testuser",
		Channel:    "#anotherchannel",
		Message:    "waves hello",
		Badges:     []string{"vip/1"},
		Bits:       0,
		Mod:        0,
		Subscribed: 0,
		Color:      "#FF4500",
		RoomID:     "#654321",
		CreatedAt:  int32(testTime().Unix()),
	}

	result := twitchchatparser.ParseTwitchMessage(message)
	assert.Equal(t, expected, result)
}

func TestParseTwitchMessage_NoBadges(t *testing.T) {
	messageTime := testTime()

	message := twitch.PrivateMessage{
		User: twitch.User{
			Name: "nobadgesuser",
		},
		Message: "No badges here!",
		Tags: map[string]string{
			"badges":     "",
			"bits":       "0",
			"subscriber": "0",
			"mod":        "0",
			"color":      "",
			"room-id":    "789012",
		},
		Channel: "#nobadgeschannel",
		Time:    messageTime,
	}

	var expectedBadges []string

	expected := &models.TwitchMessage{
		Username:   "nobadgesuser",
		Channel:    "#nobadgeschannel",
		Message:    "No badges here!",
		Badges:     expectedBadges,
		Bits:       0,
		Mod:        0,
		Subscribed: 0,
		Color:      "",
		RoomID:     "#789012",
		CreatedAt:  int32(messageTime.Unix()),
	}

	result := twitchchatparser.ParseTwitchMessage(message)
	assert.Equal(t, expected, result)
}

func testTime() time.Time {
	return time.Now()
}
