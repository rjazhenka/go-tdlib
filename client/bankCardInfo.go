// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/rjazhenka/go-tdlib/v2/tdlib"
)

// GetBankCardInfo Returns information about a bank card
// @param bankCardNumber The bank card number
func (client *Client) GetBankCardInfo(bankCardNumber string) (*tdlib.BankCardInfo, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":            "getBankCardInfo",
		"bank_card_number": bankCardNumber,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var bankCardInfo tdlib.BankCardInfo
	err = json.Unmarshal(result.Raw, &bankCardInfo)
	return &bankCardInfo, err

}
