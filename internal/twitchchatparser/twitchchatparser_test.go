package twitchchatparser

import (
	"testing"
)

func TestShouldProcessMessage(t *testing.T) {
	message := ":tmi.twitch.tv CAP * NAK :twitch.tv/tagsPASS oauth:123"
	if ShouldProcessMessage(message) {
		t.Errorf("Expected ShouldProcessMessage to return false, got true")
	}

	message = ":tmi.twitch.tv CAP * ACK :twitch.tv/commands"
	if ShouldProcessMessage(message) {
		t.Errorf("Expected ShouldProcessMessage to return false, got true")
	}

	message = ":tmi.twitch.tv 001 justinfan12345 :Welcome, GLHF!"
	if ShouldProcessMessage(message) {
		t.Errorf("Expected ShouldProcessMessage to return false, got true")
	}

	message = ":tmi.twitch.tv 002 justinfan12345 :Your host is tmi.twitch.tv"
	if ShouldProcessMessage(message) {
		t.Errorf("Expected ShouldProcessMessage to return false, got true")
	}

	message = ":tmi.twitch.tv 003 justinfan12345 :This server is rather new"
	if ShouldProcessMessage(message) {
		t.Errorf("Expected ShouldProcessMessage to return false, got true")
	}

	message = ":tmi.twitch.tv 004 justinfan12345 :-"
	if ShouldProcessMessage(message) {
		t.Errorf("Expected ShouldProcessMessage to return false, got true")
	}

	message = ":tmi.twitch.tv 375 justinfan12345 :-"
	if ShouldProcessMessage(message) {
		t.Errorf("Expected ShouldProcessMessage to return false, got true")
	}

	message = ":tmi.twitch.tv 372 justinfan12345 :You are in a maze of twisty passages, all alike."
	if ShouldProcessMessage(message) {
		t.Errorf("Expected ShouldProcessMessage to return false, got true")
	}

	message = ":tmi.twitch.tv 376 justinfan12345 :>"
	if ShouldProcessMessage(message) {
		t.Errorf("Expected ShouldProcessMessage to return false, got true")
	}

	message = ":justinfan12345!justinfan12345@justinfan12345.tmi.twitch.tv JOIN #channel"
	if ShouldProcessMessage(message) {
		t.Errorf("Expected ShouldProcessMessage to return false, got true")
	}

	message = ":tmi.twitch.tv ROOMSTATE #channel"
	if ShouldProcessMessage(message) {
		t.Errorf("Expected ShouldProcessMessage to return false, got true")
	}

	message = ":justinfan12345.tmi.twitch.tv 353 justinfan12345 = #channel :justinfan12345"
	if ShouldProcessMessage(message) {
		t.Errorf("Expected ShouldProcessMessage to return false, got true")
	}

	message = ":justinfan12345.tmi.twitch.tv 366 justinfan12345 #channel :End of /NAMES list"
	if ShouldProcessMessage(message) {
		t.Errorf("Expected ShouldProcessMessage to return false, got true")
	}

	message = ":name!name@name.tmi.twitch.tv PRIVMSG #channel :Here is my example message, I know!"
	if !ShouldProcessMessage(message) {
		t.Errorf("Expected ShouldProcessMessage to return true, got false")
	}

	message = "PING :tmi.twitch.tv"
	if !ShouldProcessMessage(message) {
		t.Errorf("Expected ShouldProcessMessage to return false, got true")
	}
}

func TestParseMessage(t *testing.T) {
	message := "@badge-info=;badges=broadcaster/1;color=#0000FF;display-name=abc;emotes=;id=d3152743-9e76-4c2f-a851-8ec9435472b7;mod=0;room-id=11148817;subscriber=0;tmi-sent-ts=1642782604907;turbo=0;user-id=123456789;user-type= :abc!abc@abc.tmi.twitch.tv PRIVMSG #xyz :-asd"
	parsedMessage := ParseMessage(message)
	if parsedMessage.Tags["badge-info"] != "" {
		t.Errorf("Expected badge-info to be empty, got %s", parsedMessage.Tags["badge-info"])
	}
	if parsedMessage.Tags["badges"] != "map[broadcaster:1]" {
		t.Errorf("Expected badges to be map[broadcaster:1], got %s", parsedMessage.Tags["badges"])
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
	if parsedTags["badges"] != "map[broadcaster:1]" {
		t.Errorf("Expected badges to be map[broadcaster:1], got %s", parsedTags["badges"])
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
