package twitchchatparser

import (
	"fmt"
	"strings"
	"time"
)

type TwitchChatMessage struct {
	Timestamp    time.Time `json:"timestamp"`     // Precise message time (UTC)
	Channel      string    `json:"channel"`       // Channel where the message was sent
	Username     string    `json:"username"`      // Username of the sender
	UserID       string    `json:"user_id"`       // Unique Twitch ID of the sender
	DisplayName  string    `json:"display_name"`  // Display name (may differ from username)
	MessageText  string    `json:"message_text"`  // Actual text content of the message
	IsModerator  bool      `json:"is_moderator"`  // Whether the sender is a moderator
	IsSubscriber bool      `json:"is_subscriber"` // Whether the sender is a subscriber
	Bits         int       `json:"bits"`          // Number of bits used in a cheer (0 if not applicable)
	Emotes       []Emote   `json:"emotes"`        // List of emotes used in the message
	MessageType  string    `json:"message_type"`  // Type of message (e.g., "chat", "whisper", "action")
}

type Emote struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"` // Number of times this emote was used in the message
}

type ParsedTwitchMessage struct {
	Tags       map[string]string
	Source     map[string]string
	Command    map[string]string
	Parameters string
}

func ParseMessage(message string) *ParsedTwitchMessage {
	parsedMessage := &ParsedTwitchMessage{}
	idx := 0

	if message[idx] == '@' {
		endIdx := strings.Index(message, " ")
		rawTagsComponent := message[1:endIdx]
		idx = endIdx + 1
		parsedMessage.Tags = parseTags(rawTagsComponent)
	}

	if message[idx] == ':' {
		idx += 1
		endIdx := strings.Index(message[idx:], " ") + idx
		rawSourceComponent := message[idx:endIdx]
		idx = endIdx + 1
		parsedMessage.Source = parseSource(rawSourceComponent)
	}

	endIdx := strings.Index(message[idx:], ":") + idx
	if endIdx == -1 {
		endIdx = len(message)
	}
	rawCommandComponent := strings.TrimSpace(message[idx:endIdx])
	parsedMessage.Command = parseCommand(rawCommandComponent)

	if endIdx != len(message) {
		idx = endIdx + 1
		rawParametersComponent := message[idx:]
		parsedMessage.Parameters = rawParametersComponent
		if rawParametersComponent[0] == '!' {
			parsedMessage.Command = parseParameters(rawParametersComponent, parsedMessage.Command)
		}
	}

	return parsedMessage
}

func parseTags(tags string) map[string]string {
	tagsToIgnore := map[string]bool{"client-nonce": true, "flags": true}
	dictParsedTags := map[string]string{}

	parsedTags := strings.Split(tags, ";")
	for _, tag := range parsedTags {
		parsedTag := strings.Split(tag, "=")
		tagValue := ""
		if len(parsedTag) > 1 {
			tagValue = parsedTag[1]
		}

		switch parsedTag[0] {
		case "badges":
			fallthrough
		case "badge-info":
			if tagValue != "" {
				dict := map[string]string{}
				badges := strings.Split(tagValue, ",")
				for _, pair := range badges {
					badgeParts := strings.Split(pair, "/")
					dict[badgeParts[0]] = badgeParts[1]
				}
				dictParsedTags[parsedTag[0]] = fmt.Sprintf("%v", dict)
			} else {
				dictParsedTags[parsedTag[0]] = ""
			}
		case "emotes":
			if tagValue != "" {
				dictEmotes := map[string]string{}
				emotes := strings.Split(tagValue, "/")
				for _, emote := range emotes {
					emoteParts := strings.Split(emote, ":")
					textPositions := []map[string]int{}
					positions := strings.Split(emoteParts[1], ",")
					for _, position := range positions {
						positionParts := strings.Split(position, "-")
						pos := map[string]int{"startPosition": int(positionParts[0][0]), "endPosition": int(positionParts[1][0])}
						textPositions = append(textPositions, pos)
					}
					dictEmotes[emoteParts[0]] = fmt.Sprintf("%v", textPositions)
				}
				dictParsedTags[parsedTag[0]] = fmt.Sprintf("%v", dictEmotes)
			} else {
				dictParsedTags[parsedTag[0]] = ""
			}
		case "emote-sets":
			emoteSetIds := strings.Split(tagValue, ",")
			dictParsedTags[parsedTag[0]] = fmt.Sprintf("%v", emoteSetIds)
		default:
			if !tagsToIgnore[parsedTag[0]] {
				dictParsedTags[parsedTag[0]] = tagValue
			}
		}
	}
	return dictParsedTags
}

func parseCommand(rawCommandComponent string) map[string]string {
	commandParts := strings.Split(rawCommandComponent, " ")
	parsedCommand := map[string]string{}

	switch commandParts[0] {
	case "JOIN":
		fallthrough
	case "PART":
		fallthrough
	case "NOTICE":
		fallthrough
	case "CLEARCHAT":
		fallthrough
	case "HOSTTARGET":
		fallthrough
	case "PRIVMSG":
		parsedCommand["command"] = commandParts[0]
		parsedCommand["channel"] = commandParts[1]
	case "PING":
		parsedCommand["command"] = commandParts[0]
	case "CAP":
		parsedCommand["command"] = commandParts[0]
		parsedCommand["isCapRequestEnabled"] = commandParts[2]
	case "GLOBALUSERSTATE":
		parsedCommand["command"] = commandParts[0]
	case "USERSTATE":
		fallthrough
	case "ROOMSTATE":
		parsedCommand["command"] = commandParts[0]
		parsedCommand["channel"] = commandParts[1]
	case "RECONNECT":
		fmt.Println("The Twitch IRC server is about to terminate the connection for maintenance.")
		parsedCommand["command"] = commandParts[0]
	case "421":
		fmt.Printf("Unsupported IRC command: %s\n", commandParts[2])
		return nil
	case "001":
		parsedCommand["command"] = commandParts[0]
		parsedCommand["channel"] = commandParts[1]
	default:
		fmt.Printf("\nUnexpected command: %s\n", commandParts[0])
		return nil
	}
	return parsedCommand
}

func parseSource(rawSourceComponent string) map[string]string {
	if rawSourceComponent == "" {
		return nil
	}
	sourceParts := strings.Split(rawSourceComponent, "!")
	source := map[string]string{}
	if len(sourceParts) == 2 {
		source["nick"] = sourceParts[0]
		source["host"] = sourceParts[1]
	} else {
		source["nick"] = ""
		source["host"] = sourceParts[0]
	}
	return source
}

func parseParameters(rawParametersComponent string, command map[string]string) map[string]string {
	idx := 0
	commandParts := strings.TrimSpace(rawParametersComponent[idx+1:])
	paramsIdx := strings.Index(commandParts, " ")

	if paramsIdx == -1 {
		command["botCommand"] = commandParts
	} else {
		command["botCommand"] = commandParts[:paramsIdx]
		command["botCommandParams"] = strings.TrimSpace(commandParts[paramsIdx:])
	}
	return command
}
