// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/rjazhenka/go-tdlib/v2/tdlib"
)

// GetJsonValue Converts a JSON-serialized string to corresponding JsonValue object. Can be called synchronously
// @param jsonString The JSON-serialized string
func (client *Client) GetJsonValue(jsonString string) (tdlib.JsonValue, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getJsonValue",
		"json":  jsonString,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	switch tdlib.JsonValueEnum(result.Data["@type"].(string)) {

	case tdlib.JsonValueNullType:
		var jsonValue tdlib.JsonValueNull
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case tdlib.JsonValueBooleanType:
		var jsonValue tdlib.JsonValueBoolean
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case tdlib.JsonValueNumberType:
		var jsonValue tdlib.JsonValueNumber
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case tdlib.JsonValueStringType:
		var jsonValue tdlib.JsonValueString
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case tdlib.JsonValueArrayType:
		var jsonValue tdlib.JsonValueArray
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case tdlib.JsonValueObjectType:
		var jsonValue tdlib.JsonValueObject
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// GetApplicationConfig Returns application config, provided by the server. Can be called before authorization
func (client *Client) GetApplicationConfig() (tdlib.JsonValue, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getApplicationConfig",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	switch tdlib.JsonValueEnum(result.Data["@type"].(string)) {

	case tdlib.JsonValueNullType:
		var jsonValue tdlib.JsonValueNull
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case tdlib.JsonValueBooleanType:
		var jsonValue tdlib.JsonValueBoolean
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case tdlib.JsonValueNumberType:
		var jsonValue tdlib.JsonValueNumber
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case tdlib.JsonValueStringType:
		var jsonValue tdlib.JsonValueString
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case tdlib.JsonValueArrayType:
		var jsonValue tdlib.JsonValueArray
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case tdlib.JsonValueObjectType:
		var jsonValue tdlib.JsonValueObject
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}
