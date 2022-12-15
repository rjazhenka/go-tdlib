// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/rjazhenka/go-tdlib/v2/tdlib"
)

// GetChatEventLog Returns a list of service actions taken by chat members and administrators in the last 48 hours. Available only for supergroups and channels. Requires administrator rights. Returns results in reverse chronological order (i. e., in order of decreasing event_id)
// @param chatID Chat identifier
// @param query Search query by which to filter events
// @param fromEventID Identifier of an event from which to return results. Use 0 to get results from the latest events
// @param limit The maximum number of events to return; up to 100
// @param filters The types of events to return. By default, all types will be returned
// @param userIDs User identifiers by which to filter events. By default, events relating to all users will be returned
func (client *Client) GetChatEventLog(chatID int64, query string, fromEventID *tdlib.JSONInt64, limit int32, filters *tdlib.ChatEventLogFilters, userIDs []int64) (*tdlib.ChatEvents, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":         "getChatEventLog",
		"chat_id":       chatID,
		"query":         query,
		"from_event_id": fromEventID,
		"limit":         limit,
		"filters":       filters,
		"user_ids":      userIDs,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chatEvents tdlib.ChatEvents
	err = json.Unmarshal(result.Raw, &chatEvents)
	return &chatEvents, err

}
