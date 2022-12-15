// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/rjazhenka/go-tdlib/v2/tdlib"
)

// GetUserFullInfo Returns full information about a user by their identifier
// @param userID User identifier
func (client *Client) GetUserFullInfo(userID int64) (*tdlib.UserFullInfo, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":   "getUserFullInfo",
		"user_id": userID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var userFullInfo tdlib.UserFullInfo
	err = json.Unmarshal(result.Raw, &userFullInfo)
	return &userFullInfo, err

}
