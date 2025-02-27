// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/rjazhenka/go-tdlib/v2/tdlib"
)

// GetChats Returns an ordered list of chats from the beginning of a chat list. For informational purposes only. Use loadChats instead to maintain chat lists
// @param chatList The chat list in which to return chats
// @param limit The maximum number of chats to be returned
func (client *Client) GetChats(chatList tdlib.ChatList, limit int32) (*tdlib.Chats, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":     "getChats",
		"chat_list": chatList,
		"limit":     limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chats tdlib.Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// SearchPublicChats Searches public chats by looking for specified query in their username and title. Currently only private chats, supergroups and channels can be public. Returns a meaningful number of results. Returns nothing if the length of the searched username prefix is less than 5. Excludes private chats with contacts and chats from the chat list from the results
// @param query Query to search for
func (client *Client) SearchPublicChats(query string) (*tdlib.Chats, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "searchPublicChats",
		"query": query,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chats tdlib.Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// SearchChats Searches for the specified query in the title and username of already known chats, this is an offline request. Returns chats in the order seen in the main chat list
// @param query Query to search for. If the query is empty, returns up to 50 recently found chats
// @param limit The maximum number of chats to be returned
func (client *Client) SearchChats(query string, limit int32) (*tdlib.Chats, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "searchChats",
		"query": query,
		"limit": limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chats tdlib.Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// SearchChatsOnServer Searches for the specified query in the title and username of already known chats via request to the server. Returns chats in the order seen in the main chat list
// @param query Query to search for
// @param limit The maximum number of chats to be returned
func (client *Client) SearchChatsOnServer(query string, limit int32) (*tdlib.Chats, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "searchChatsOnServer",
		"query": query,
		"limit": limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chats tdlib.Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// GetTopChats Returns a list of frequently used chats. Supported only if the chat info database is enabled
// @param category Category of chats to be returned
// @param limit The maximum number of chats to be returned; up to 30
func (client *Client) GetTopChats(category tdlib.TopChatCategory, limit int32) (*tdlib.Chats, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":    "getTopChats",
		"category": category,
		"limit":    limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chats tdlib.Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// GetRecentlyOpenedChats Returns recently opened chats, this is an offline request. Returns chats in the order of last opening
// @param limit The maximum number of chats to be returned
func (client *Client) GetRecentlyOpenedChats(limit int32) (*tdlib.Chats, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getRecentlyOpenedChats",
		"limit": limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chats tdlib.Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// GetCreatedPublicChats Returns a list of public chats of the specified type, owned by the user
// @param typeParam Type of the public chats to return
func (client *Client) GetCreatedPublicChats(typeParam tdlib.PublicChatType) (*tdlib.Chats, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getCreatedPublicChats",
		"type":  typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chats tdlib.Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// GetSuitableDiscussionChats Returns a list of basic group and supergroup chats, which can be used as a discussion group for a channel. Returned basic group chats must be first upgraded to supergroups before they can be set as a discussion group. To set a returned supergroup as a discussion group, access to its old messages must be enabled using toggleSupergroupIsAllHistoryAvailable first
func (client *Client) GetSuitableDiscussionChats() (*tdlib.Chats, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getSuitableDiscussionChats",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chats tdlib.Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// GetInactiveSupergroupChats Returns a list of recently inactive supergroups and channels. Can be used when user reaches limit on the number of joined supergroups and channels and receives CHANNELS_TOO_MUCH error
func (client *Client) GetInactiveSupergroupChats() (*tdlib.Chats, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getInactiveSupergroupChats",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chats tdlib.Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// GetGroupsInCommon Returns a list of common group chats with a given user. Chats are sorted by their type and creation date
// @param userID User identifier
// @param offsetChatID Chat identifier starting from which to return chats; use 0 for the first request
// @param limit The maximum number of chats to be returned; up to 100
func (client *Client) GetGroupsInCommon(userID int64, offsetChatID int64, limit int32) (*tdlib.Chats, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":          "getGroupsInCommon",
		"user_id":        userID,
		"offset_chat_id": offsetChatID,
		"limit":          limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chats tdlib.Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// GetChatNotificationSettingsExceptions Returns list of chats with non-default notification settings
// @param scope If specified, only chats from the specified scope will be returned
// @param compareSound If true, also chats with non-default sound will be returned
func (client *Client) GetChatNotificationSettingsExceptions(scope tdlib.NotificationSettingsScope, compareSound bool) (*tdlib.Chats, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":         "getChatNotificationSettingsExceptions",
		"scope":         scope,
		"compare_sound": compareSound,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chats tdlib.Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}
