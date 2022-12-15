// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/rjazhenka/go-tdlib/v2/tdlib"
)

// TestSquareInt Returns the squared received number; for testing only. This is an offline method. Can be called before authorization
// @param x Number to square
func (client *Client) TestSquareInt(x int32) (*tdlib.TestInt, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "testSquareInt",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var testInt tdlib.TestInt
	err = json.Unmarshal(result.Raw, &testInt)
	return &testInt, err

}
