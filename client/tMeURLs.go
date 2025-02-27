// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/rjazhenka/go-tdlib/v2/tdlib"
)

// GetRecentlyVisitedTMeURLs Returns t.me URLs recently visited by a newly registered user
// @param referrer Google Play referrer to identify the user
func (client *Client) GetRecentlyVisitedTMeURLs(referrer string) (*tdlib.TMeURLs, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":    "getRecentlyVisitedTMeUrls",
		"referrer": referrer,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var tMeURLs tdlib.TMeURLs
	err = json.Unmarshal(result.Raw, &tMeURLs)
	return &tMeURLs, err

}
