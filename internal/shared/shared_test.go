package shared

import (
	"sync"
	"testing"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/michaelfioretti/twitch-stats-producer/internal/constants"
	models "github.com/michaelfioretti/twitch-stats-producer/internal/models/proto"
	"github.com/stretchr/testify/assert"
)

func TestShared(t *testing.T) {
	t.Run("TestMessageChannel", func(t *testing.T) {
		assert.NotNil(t, MessageChannel)
		assert.IsType(t, make(chan *models.TwitchMessage, constants.MESSAGES_PER_BATCH), MessageChannel)
	})

	t.Run("TestTwitchClient", func(t *testing.T) {
		assert.Nil(t, TwitchClient)
		assert.IsType(t, (*twitch.Client)(nil), TwitchClient)
	})

	t.Run("TestLastUpdatedTopStreamers", func(t *testing.T) {
		assert.Nil(t, LastUpdatedTopStreamers)
		assert.IsType(t, []string{}, LastUpdatedTopStreamers)
	})

	t.Run("TestLastUpdatedTopStreamersMutex", func(t *testing.T) {
		assert.NotNil(t, LastUpdatedTopStreamersMutex)
		assert.IsType(t, &sync.RWMutex{}, LastUpdatedTopStreamersMutex)
	})
}
