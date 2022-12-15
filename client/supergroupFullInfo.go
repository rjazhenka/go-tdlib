// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/rjazhenka/go-tdlib/v2/tdlib"
)

// GetSupergroupFullInfo Returns full information about a supergroup or a channel by its identifier, cached for up to 1 minute
// @param supergroupID Supergroup or channel identifier
func (client *Client) GetSupergroupFullInfo(supergroupID int64) (*tdlib.SupergroupFullInfo, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":         "getSupergroupFullInfo",
		"supergroup_id": supergroupID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var supergroupFullInfo tdlib.SupergroupFullInfo
	err = json.Unmarshal(result.Raw, &supergroupFullInfo)
	return &supergroupFullInfo, err

}
