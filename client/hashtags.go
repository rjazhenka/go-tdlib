// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/rjazhenka/go-tdlib/v2/tdlib"
)

// SearchHashtags Searches for recently used hashtags by their prefix
// @param prefix Hashtag prefix to search for
// @param limit The maximum number of hashtags to be returned
func (client *Client) SearchHashtags(prefix string, limit int32) (*tdlib.Hashtags, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":  "searchHashtags",
		"prefix": prefix,
		"limit":  limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var hashtags tdlib.Hashtags
	err = json.Unmarshal(result.Raw, &hashtags)
	return &hashtags, err

}
