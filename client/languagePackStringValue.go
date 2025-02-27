// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/rjazhenka/go-tdlib/v2/tdlib"
)

// GetLanguagePackString Returns a string stored in the local database from the specified localization target and language pack by its key. Returns a 404 error if the string is not found. Can be called synchronously
// @param languagePackDatabasePath Path to the language pack database in which strings are stored
// @param localizationTarget Localization target to which the language pack belongs
// @param languagePackID Language pack identifier
// @param key Language pack key of the string to be returned
func (client *Client) GetLanguagePackString(languagePackDatabasePath string, localizationTarget string, languagePackID string, key string) (tdlib.LanguagePackStringValue, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":                       "getLanguagePackString",
		"language_pack_database_path": languagePackDatabasePath,
		"localization_target":         localizationTarget,
		"language_pack_id":            languagePackID,
		"key":                         key,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	switch tdlib.LanguagePackStringValueEnum(result.Data["@type"].(string)) {

	case tdlib.LanguagePackStringValueOrdinaryType:
		var languagePackStringValue tdlib.LanguagePackStringValueOrdinary
		err = json.Unmarshal(result.Raw, &languagePackStringValue)
		return &languagePackStringValue, err

	case tdlib.LanguagePackStringValuePluralizedType:
		var languagePackStringValue tdlib.LanguagePackStringValuePluralized
		err = json.Unmarshal(result.Raw, &languagePackStringValue)
		return &languagePackStringValue, err

	case tdlib.LanguagePackStringValueDeletedType:
		var languagePackStringValue tdlib.LanguagePackStringValueDeleted
		err = json.Unmarshal(result.Raw, &languagePackStringValue)
		return &languagePackStringValue, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}
