package test

// import (
// 	"encoding/json"
// 	"testing"
// 	"tryhard-platform/model"
// )

// func TestMessageEncoding(t *testing.T) {
// 	jsonMsg := `{
// 		"action": "ACTION"
// 	}`

// 	var message model.Message
// 	err := json.Unmarshal([]byte(jsonMsg), &message)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if message.Action != "ACTION" {
// 		t.Fatal("invalid action", message)
// 	}
// }

// func TestHeartbeatMessageEncoding(t *testing.T) {
// 	jsonMsg := `{
// 		"action": "ACTION"
// 	}`

// 	var message model.PartyMessage
// 	err := json.Unmarshal([]byte(jsonMsg), &message)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if message.Action != "ACTION" {
// 		t.Fatal("invalid action", message)
// 	}
// }

// func TestPartyMessageEncoding(t *testing.T) {
// 	jsonMsg := `{
// 		"action": "ACTION",
// 		"party": {
// 			"id": "TEST_ID",
// 			"code": "TEST_CODE",
// 			"status": {
// 				"id": "TEST_ID",
// 				"value": "TEST_STATUS",
// 				"description": "TEST_DESCRIPTION"
// 			}
// 		},
// 		"users":[
// 			{
// 				"client_id": "CLIENT_1",
// 				"username": "TEST_USER"
// 			}
// 		]
// 	}`

// 	var message model.PartyMessage
// 	err := json.Unmarshal([]byte(jsonMsg), &message)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if message.Action != "ACTION" {
// 		t.Fatal("invalid action", message)
// 	}

// 	if len(message.Users) != 1 {
// 		t.Fatal("invalid user count", message)
// 	}
// }
