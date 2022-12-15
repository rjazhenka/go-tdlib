// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/rjazhenka/go-tdlib/v2/tdlib"
)

// CreateVoiceChat Creates a voice chat (a group call bound to a chat). Available only for basic groups, supergroups and channels; requires can_manage_voice_chats rights
// @param chatID Chat identifier, in which the voice chat will be created
// @param title Group call title; if empty, chat title will be used
// @param startDate Point in time (Unix timestamp) when the group call is supposed to be started by an administrator; 0 to start the voice chat immediately. The date must be at least 10 seconds and at most 8 days in the future
func (client *Client) CreateVoiceChat(chatID int64, title string, startDate int32) (*tdlib.GroupCallID, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":      "createVoiceChat",
		"chat_id":    chatID,
		"title":      title,
		"start_date": startDate,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var groupCallID tdlib.GroupCallID
	err = json.Unmarshal(result.Raw, &groupCallID)
	return &groupCallID, err

}
