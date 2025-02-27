// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/rjazhenka/go-tdlib/v2/tdlib"
)

// TestReturnError Returns the specified error and ensures that the Error object is used; for testing only. Can be called synchronously
// @param error The error to be returned
func (client *Client) TestReturnError(error *tdlib.Error) (*tdlib.Error, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "testReturnError",
		"error": error,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var errorDummy tdlib.Error
	err = json.Unmarshal(result.Raw, &errorDummy)
	return &errorDummy, err

}
