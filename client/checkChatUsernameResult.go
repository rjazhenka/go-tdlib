// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/Arman92/go-tdlib/v2/tdlib"
)

// CheckChatUsername Checks whether a username can be set for a chat
// @param chatID Chat identifier; should be identifier of a supergroup chat, or a channel chat, or a private chat with self, or zero if the chat is being created
// @param username Username to be checked
func (client *Client) CheckChatUsername(chatID int64, username string) (tdlib.CheckChatUsernameResult, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":    "checkChatUsername",
		"chat_id":  chatID,
		"username": username,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	switch tdlib.CheckChatUsernameResultEnum(result.Data["@type"].(string)) {

	case tdlib.CheckChatUsernameResultOkType:
		var checkChatUsernameResult tdlib.CheckChatUsernameResultOk
		err = json.Unmarshal(result.Raw, &checkChatUsernameResult)
		return &checkChatUsernameResult, err

	case tdlib.CheckChatUsernameResultUsernameInvalidType:
		var checkChatUsernameResult tdlib.CheckChatUsernameResultUsernameInvalid
		err = json.Unmarshal(result.Raw, &checkChatUsernameResult)
		return &checkChatUsernameResult, err

	case tdlib.CheckChatUsernameResultUsernameOccupiedType:
		var checkChatUsernameResult tdlib.CheckChatUsernameResultUsernameOccupied
		err = json.Unmarshal(result.Raw, &checkChatUsernameResult)
		return &checkChatUsernameResult, err

	case tdlib.CheckChatUsernameResultPublicChatsTooMuchType:
		var checkChatUsernameResult tdlib.CheckChatUsernameResultPublicChatsTooMuch
		err = json.Unmarshal(result.Raw, &checkChatUsernameResult)
		return &checkChatUsernameResult, err

	case tdlib.CheckChatUsernameResultPublicGroupsUnavailableType:
		var checkChatUsernameResult tdlib.CheckChatUsernameResultPublicGroupsUnavailable
		err = json.Unmarshal(result.Raw, &checkChatUsernameResult)
		return &checkChatUsernameResult, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}
