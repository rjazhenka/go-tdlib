// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/rjazhenka/go-tdlib/v2/tdlib"
)

// ChangePhoneNumber Changes the phone number of the user and sends an authentication code to the user's new phone number. On success, returns information about the sent code
// @param phoneNumber The new phone number of the user in international format
// @param settings Settings for the authentication of the user's phone number
func (client *Client) ChangePhoneNumber(phoneNumber string, settings *tdlib.PhoneNumberAuthenticationSettings) (*tdlib.AuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":        "changePhoneNumber",
		"phone_number": phoneNumber,
		"settings":     settings,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var authenticationCodeInfo tdlib.AuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &authenticationCodeInfo)
	return &authenticationCodeInfo, err

}

// ResendChangePhoneNumberCode Re-sends the authentication code sent to confirm a new phone number for the current user. Works only if the previously received authenticationCodeInfo next_code_type was not null and the server-specified timeout has passed
func (client *Client) ResendChangePhoneNumberCode() (*tdlib.AuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "resendChangePhoneNumberCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var authenticationCodeInfo tdlib.AuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &authenticationCodeInfo)
	return &authenticationCodeInfo, err

}

// SendPhoneNumberVerificationCode Sends a code to verify a phone number to be added to a user's Telegram Passport
// @param phoneNumber The phone number of the user, in international format
// @param settings Settings for the authentication of the user's phone number
func (client *Client) SendPhoneNumberVerificationCode(phoneNumber string, settings *tdlib.PhoneNumberAuthenticationSettings) (*tdlib.AuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":        "sendPhoneNumberVerificationCode",
		"phone_number": phoneNumber,
		"settings":     settings,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var authenticationCodeInfo tdlib.AuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &authenticationCodeInfo)
	return &authenticationCodeInfo, err

}

// ResendPhoneNumberVerificationCode Re-sends the code to verify a phone number to be added to a user's Telegram Passport
func (client *Client) ResendPhoneNumberVerificationCode() (*tdlib.AuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "resendPhoneNumberVerificationCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var authenticationCodeInfo tdlib.AuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &authenticationCodeInfo)
	return &authenticationCodeInfo, err

}

// SendPhoneNumberConfirmationCode Sends phone number confirmation code to handle links of the type internalLinkTypePhoneNumberConfirmation
// @param hash Hash value from the link
// @param phoneNumber Phone number value from the link
// @param settings Settings for the authentication of the user's phone number
func (client *Client) SendPhoneNumberConfirmationCode(hash string, phoneNumber string, settings *tdlib.PhoneNumberAuthenticationSettings) (*tdlib.AuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":        "sendPhoneNumberConfirmationCode",
		"hash":         hash,
		"phone_number": phoneNumber,
		"settings":     settings,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var authenticationCodeInfo tdlib.AuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &authenticationCodeInfo)
	return &authenticationCodeInfo, err

}

// ResendPhoneNumberConfirmationCode Resends phone number confirmation code
func (client *Client) ResendPhoneNumberConfirmationCode() (*tdlib.AuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "resendPhoneNumberConfirmationCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var authenticationCodeInfo tdlib.AuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &authenticationCodeInfo)
	return &authenticationCodeInfo, err

}
