package twitchchatparser

import (
	"testing"
)

func TestParseMessage(t *testing.T) {
	message := "@badge-info=;badges=broadcaster/1;color=#0000FF;display-name=abc;emotes=;id=d3152743-9e76-4c2f-a851-8ec9435472b7;mod=0;room-id=11148817;subscriber=0;tmi-sent-ts=1642782604907;turbo=0;user-id=123456789;user-type= :abc!abc@abc.tmi.twitch.tv PRIVMSG #xyz :-asd"
	parsedMessage := ParseMessage(message)
	if parsedMessage.Tags["badge-info"] != "" {
		t.Errorf("Expected badge-info to be empty, got %s", parsedMessage.Tags["badge-info"])
	}
	if parsedMessage.Tags["badges"] != "broadcaster/1" {
		t.Errorf("Expected badges to be broadcaster/1, got %s", parsedMessage.Tags["badges"])
	}
	if parsedMessage.Source["nick"] != "abc" {
		t.Errorf("Expected nick to be abc, got %s", parsedMessage.Source["nick"])
	}
	if parsedMessage.Command["command"] != "PRIVMSG" {
		t.Errorf("Expected command to be PRIVMSG, got %s", parsedMessage.Command["command"])
	}
	if parsedMessage.Parameters != "-asd" {
		t.Errorf("Expected parameters to be -asd, got %s", parsedMessage.Parameters)
	}
}

func TestParseTags(t *testing.T) {
	tags := "badge-info=;badges=broadcaster/1;color=#0000FF;display-name=abc;emotes=;id=d3152743-9e76-4c2f-a851-8ec9435472b7;mod=0;room-id=11148817;subscriber=0;tmi-sent-ts=1642782604907;turbo=0;user-id=123456789;user-type="
	parsedTags := parseTags(tags)
	if parsedTags["badge-info"] != "" {
		t.Errorf("Expected badge-info to be empty, got %s", parsedTags["badge-info"])
	}
	if parsedTags["badges"] != "broadcaster/1" {
		t.Errorf("Expected badges to be broadcaster/1, got %s", parsedTags["badges"])
	}
	if parsedTags["color"] != "#0000FF" {
		t.Errorf("Expected color to be #0000FF, got %s", parsedTags["color"])
	}
}

func TestParseCommand(t *testing.T) {
	command := "PRIVMSG #xyz"
	parsedCommand := parseCommand(command)
	if parsedCommand["command"] != "PRIVMSG" {
		t.Errorf("Expected command to be PRIVMSG, got %s", parsedCommand["command"])
	}
	if parsedCommand["channel"] != "#xyz" {
		t.Errorf("Expected channel to be #xyz, got %s", parsedCommand["channel"])
	}
}

func TestParseSource(t *testing.T) {
	source := "abc!abc@abc.tmi.twitch.tv"
	parsedSource := parseSource(source)
	if parsedSource["nick"] != "abc" {
		t.Errorf("Expected nick to be abc, got %s", parsedSource["nick"])
	}
	if parsedSource["host"] != "abc@abc.tmi.twitch.tv" {
		t.Errorf("Expected host to be abc@abc.tmi.twitch.tv, got %s", parsedSource["host"])
	}
}

func TestParseParameters(t *testing.T) {
	parameters := "!test param1 param2"
	parsedCommand := map[string]string{}
	parsedCommand = parseParameters(parameters, parsedCommand)
	if parsedCommand["botCommand"] != "test" {
		t.Errorf("Expected botCommand to be test, got %s", parsedCommand["botCommand"])
	}
	if parsedCommand["botCommandParams"] != "param1 param2" {
		t.Errorf("Expected botCommandParams to be param1 param2, got %s", parsedCommand["botCommandParams"])
	}
}
