package models

// func TestIRCChatMessageData(t *testing.T) {
// 	data := IRCChatMessageData{
// 		Streamer: "testStreamer",
// 		Message:  "testMessage",
// 	}

// 	// Test Streamer field
// 	if data.Streamer != "testStreamer" {
// 		t.Errorf("Expected Streamer to be 'testStreamer', got '%s'", data.Streamer)
// 	}

// 	// Test Message field
// 	if data.Message != "testMessage" {
// 		t.Errorf("Expected Message to be 'testMessage', got '%s'", data.Message)
// 	}
// }

// func TestTop100StreamsResponse(t *testing.T) {
// 	// Test unmarshaling JSON data
// 	jsonData := `{"data":[{"id":"1","user_id":"1","user_name":"testUser","game_id":"1","type":"testType","title":"testTitle","viewer_count":100,"started_at":"2022-01-01","language":"en","thumbnail_url":"testURL","tag_ids":["1"],"is_mature":false}]}`
// 	response := Top100StreamsResponse{}
// 	err := json.Unmarshal([]byte(jsonData), &response)
// 	if err != nil {
// 		t.Errorf("Failed to unmarshal JSON data: %s", err)
// 	}

// 	// Test Data field
// 	if len(response.Data) != 1 {
// 		t.Errorf("Expected Data length to be 1, got %d", len(response.Data))
// 	}

// 	// Test Stream field
// 	stream := response.Data[0]
// 	if stream.ID != "1" {
// 		t.Errorf("Expected Stream ID to be '1', got '%s'", stream.ID)
// 	}
// 	if stream.UserID != "1" {
// 		t.Errorf("Expected Stream UserID to be '1', got '%s'", stream.UserID)
// 	}
// 	if stream.UserName != "testUser" {
// 		t.Errorf("Expected Stream UserName to be 'testUser', got '%s'", stream.UserName)
// 	}
// 	// ... continue testing other fields
// }

// func TestTwitchGame(t *testing.T) {
// 	game := TwitchGame{
// 		ID:          "1",
// 		Name:        "testGame",
// 		BoxArtURL:   "testURL",
// 		ViewerCount: 100,
// 		Popularity:  50,
// 	}

// 	// Test ID field
// 	if game.ID != "1" {
// 		t.Errorf("Expected Game ID to be '1', got '%s'", game.ID)
// 	}

// 	// Test Name field
// 	if game.Name != "testGame" {
// 		t.Errorf("Expected Game Name to be 'testGame', got '%s'", game.Name)
// 	}

// 	// Test BoxArtURL field
// 	if game.BoxArtURL != "testURL" {
// 		t.Errorf("Expected Game BoxArtURL to be 'testURL', got '%s'", game.BoxArtURL)
// 	}

// 	// Test ViewerCount field
// 	if game.ViewerCount != 100 {
// 		t.Errorf("Expected Game ViewerCount to be 100, got %d", game.ViewerCount)
// 	}

// 	// Test Popularity field
// 	if game.Popularity != 50 {
// 		t.Errorf("Expected Game Popularity to be 50, got %d", game.Popularity)
// 	}
// }
