// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/rjazhenka/go-tdlib/v2/tdlib"
)

// GetInternalLinkType Returns information about the type of an internal link. Returns a 404 error if the link is not internal. Can be called before authorization
// @param link The link
func (client *Client) GetInternalLinkType(link string) (tdlib.InternalLinkType, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getInternalLinkType",
		"link":  link,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	switch tdlib.InternalLinkTypeEnum(result.Data["@type"].(string)) {

	case tdlib.InternalLinkTypeActiveSessionsType:
		var internalLinkType tdlib.InternalLinkTypeActiveSessions
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case tdlib.InternalLinkTypeAuthenticationCodeType:
		var internalLinkType tdlib.InternalLinkTypeAuthenticationCode
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case tdlib.InternalLinkTypeBackgroundType:
		var internalLinkType tdlib.InternalLinkTypeBackground
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case tdlib.InternalLinkTypeBotStartType:
		var internalLinkType tdlib.InternalLinkTypeBotStart
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case tdlib.InternalLinkTypeBotStartInGroupType:
		var internalLinkType tdlib.InternalLinkTypeBotStartInGroup
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case tdlib.InternalLinkTypeChangePhoneNumberType:
		var internalLinkType tdlib.InternalLinkTypeChangePhoneNumber
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case tdlib.InternalLinkTypeChatInviteType:
		var internalLinkType tdlib.InternalLinkTypeChatInvite
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case tdlib.InternalLinkTypeFilterSettingsType:
		var internalLinkType tdlib.InternalLinkTypeFilterSettings
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case tdlib.InternalLinkTypeGameType:
		var internalLinkType tdlib.InternalLinkTypeGame
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case tdlib.InternalLinkTypeLanguagePackType:
		var internalLinkType tdlib.InternalLinkTypeLanguagePack
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case tdlib.InternalLinkTypeMessageType:
		var internalLinkType tdlib.InternalLinkTypeMessage
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case tdlib.InternalLinkTypeMessageDraftType:
		var internalLinkType tdlib.InternalLinkTypeMessageDraft
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case tdlib.InternalLinkTypePassportDataRequestType:
		var internalLinkType tdlib.InternalLinkTypePassportDataRequest
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case tdlib.InternalLinkTypePhoneNumberConfirmationType:
		var internalLinkType tdlib.InternalLinkTypePhoneNumberConfirmation
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case tdlib.InternalLinkTypeProxyType:
		var internalLinkType tdlib.InternalLinkTypeProxy
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case tdlib.InternalLinkTypePublicChatType:
		var internalLinkType tdlib.InternalLinkTypePublicChat
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case tdlib.InternalLinkTypeQrCodeAuthenticationType:
		var internalLinkType tdlib.InternalLinkTypeQrCodeAuthentication
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case tdlib.InternalLinkTypeSettingsType:
		var internalLinkType tdlib.InternalLinkTypeSettings
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case tdlib.InternalLinkTypeStickerSetType:
		var internalLinkType tdlib.InternalLinkTypeStickerSet
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case tdlib.InternalLinkTypeThemeType:
		var internalLinkType tdlib.InternalLinkTypeTheme
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case tdlib.InternalLinkTypeThemeSettingsType:
		var internalLinkType tdlib.InternalLinkTypeThemeSettings
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case tdlib.InternalLinkTypeUnknownDeepLinkType:
		var internalLinkType tdlib.InternalLinkTypeUnknownDeepLink
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	case tdlib.InternalLinkTypeVoiceChatType:
		var internalLinkType tdlib.InternalLinkTypeVoiceChat
		err = json.Unmarshal(result.Raw, &internalLinkType)
		return &internalLinkType, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}
