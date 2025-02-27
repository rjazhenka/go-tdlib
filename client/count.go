// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/rjazhenka/go-tdlib/v2/tdlib"
)

// GetChatMessageCount Returns approximate number of messages of the specified type in the chat
// @param chatID Identifier of the chat in which to count messages
// @param filter Filter for message content; searchMessagesFilterEmpty is unsupported in this function
// @param returnLocal If true, returns count that is available locally without sending network requests, returning -1 if the number of messages is unknown
func (client *Client) GetChatMessageCount(chatID int64, filter tdlib.SearchMessagesFilter, returnLocal bool) (*tdlib.Count, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":        "getChatMessageCount",
		"chat_id":      chatID,
		"filter":       filter,
		"return_local": returnLocal,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var count tdlib.Count
	err = json.Unmarshal(result.Raw, &count)
	return &count, err

}

// GetFileDownloadedPrefixSize Returns file downloaded prefix size from a given offset, in bytes
// @param fileID Identifier of the file
// @param offset Offset from which downloaded prefix size should be calculated
func (client *Client) GetFileDownloadedPrefixSize(fileID int32, offset int32) (*tdlib.Count, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":   "getFileDownloadedPrefixSize",
		"file_id": fileID,
		"offset":  offset,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var count tdlib.Count
	err = json.Unmarshal(result.Raw, &count)
	return &count, err

}

// GetImportedContactCount Returns the total number of imported contacts
func (client *Client) GetImportedContactCount() (*tdlib.Count, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getImportedContactCount",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var count tdlib.Count
	err = json.Unmarshal(result.Raw, &count)
	return &count, err

}
