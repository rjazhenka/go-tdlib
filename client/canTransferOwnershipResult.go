// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/rjazhenka/go-tdlib/v2/tdlib"
)

// CanTransferOwnership Checks whether the current session can be used to transfer a chat ownership to another user
func (client *Client) CanTransferOwnership() (tdlib.CanTransferOwnershipResult, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "canTransferOwnership",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	switch tdlib.CanTransferOwnershipResultEnum(result.Data["@type"].(string)) {

	case tdlib.CanTransferOwnershipResultOkType:
		var canTransferOwnershipResult tdlib.CanTransferOwnershipResultOk
		err = json.Unmarshal(result.Raw, &canTransferOwnershipResult)
		return &canTransferOwnershipResult, err

	case tdlib.CanTransferOwnershipResultPasswordNeededType:
		var canTransferOwnershipResult tdlib.CanTransferOwnershipResultPasswordNeeded
		err = json.Unmarshal(result.Raw, &canTransferOwnershipResult)
		return &canTransferOwnershipResult, err

	case tdlib.CanTransferOwnershipResultPasswordTooFreshType:
		var canTransferOwnershipResult tdlib.CanTransferOwnershipResultPasswordTooFresh
		err = json.Unmarshal(result.Raw, &canTransferOwnershipResult)
		return &canTransferOwnershipResult, err

	case tdlib.CanTransferOwnershipResultSessionTooFreshType:
		var canTransferOwnershipResult tdlib.CanTransferOwnershipResultSessionTooFresh
		err = json.Unmarshal(result.Raw, &canTransferOwnershipResult)
		return &canTransferOwnershipResult, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}
