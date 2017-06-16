package test

import (
	"encoding/json"
	"testing"
	d "tryhard-platform/data"
)

func TestMarshalMessage(t *testing.T) {
	msg := d.Message{
		ID:      "TEST",
		Command: "COMMAND",
	}
	result, err := json.Marshal(msg)
	if err != nil {
		t.Error(err)
	}
	t.Log("Result:", result)
	if len(result) == 0 {
		t.Fail()
	}
}

func TestUnMarshalMessage(t *testing.T) {
	msgString := `{ "command": "TEST" }`
	msgByte := []byte(msgString)
	msg := d.Message{}

	err := json.Unmarshal(msgByte, &msg)
	if err != nil {
		t.Error(err)
	}

	if msg.Command != "TEST" {
		t.FailNow()
	}
}

func TestMarshalHeartbeatMessage(t *testing.T) {
	msg := d.HeartbeatMessage{
		Message: d.Message{
			ID:      "123",
			Command: "HEARTBEAT",
		},
		Heartbeat: d.Heartbeat{
			LatencySeconds: 10,
		},
	}
	result, err := json.Marshal(msg)
	if err != nil {
		t.Error(err)
	}
	t.Log("Result:", result)
	if len(result) == 0 {
		t.Fail()
	}
}

func TestUnMarshalHeartbeatMessage(t *testing.T) {
	msgString := `{ "message": { "command": "HEARTBEAT" }, "heartbeat": { "latencySeconds": 10 }}`
	msgByte := []byte(msgString)
	msg := d.HeartbeatMessage{}

	err := json.Unmarshal(msgByte, &msg)
	if err != nil {
		t.Error(err)
	}

	t.Log("result", msg)

	if msg.Message.Command != "HEARTBEAT" {
		t.FailNow()
	}

	if msg.Heartbeat.LatencySeconds != 10 {
		t.FailNow()
	}
}
func TestMarshalPartyMessage(t *testing.T) {
	msg := d.PartyMessage{
		Message: d.Message{
			ID:      "123",
			Command: "HEARTBEAT",
		},
		Player: d.Player{
			Username: "TH_TEST",
		},
		Party: d.Party{
			Code: "123456",
		},
	}
	result, err := json.Marshal(msg)
	if err != nil {
		t.Error(err)
	}
	t.Log("Result:", result)
	if len(result) == 0 {
		t.Fail()
	}
}
func TestUnMarshalPartyMessage(t *testing.T) {
	msgString := `{ 
		"message": { "command": "JOIN" }, 
		"player": { "username": "TEST_USERNAME" },
		"party": { "code": "123456" }
	}`
	msgByte := []byte(msgString)
	msg := d.PartyMessage{}

	err := json.Unmarshal(msgByte, &msg)
	if err != nil {
		t.Error(err)
	}

	t.Log("result", msg)

	if msg.Message.Command != "JOIN" {
		t.FailNow()
	}

	if msg.Player.Username != "TEST_USERNAME" {
		t.FailNow()
	}

	if msg.Party.Code != "123456" {
		t.FailNow()
	}
}

func TestCuriousity(t *testing.T) {
	msgString := `{ 
		"message": {
			"client": { "id": "CLIENT ID" },
			"id": "MSG ID",
			"command": "JOIN"
		},
		"player": { "username": "TEST_USERNAME" },
		"party": { "code": "123456" }
	}`
	msgByte := []byte(msgString)
	msg := d.MessageAdapter{}

	err := json.Unmarshal(msgByte, &msg)
	if err != nil {
		t.Error(err)
	}

	t.Log("result", msg)

	if msg.Command != "JOIN" {
		t.FailNow()
	}
}
