// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/rjazhenka/go-tdlib/v2/tdlib"
)

// TestUseUpdate Does nothing and ensures that the Update object is used; for testing only. This is an offline method. Can be called before authorization
func (client *Client) TestUseUpdate() (tdlib.Update, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "testUseUpdate",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	switch tdlib.UpdateEnum(result.Data["@type"].(string)) {

	case tdlib.UpdateAuthorizationStateType:
		var update tdlib.UpdateAuthorizationState
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateNewMessageType:
		var update tdlib.UpdateNewMessage
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateMessageSendAcknowledgedType:
		var update tdlib.UpdateMessageSendAcknowledged
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateMessageSendSucceededType:
		var update tdlib.UpdateMessageSendSucceeded
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateMessageSendFailedType:
		var update tdlib.UpdateMessageSendFailed
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateMessageContentType:
		var update tdlib.UpdateMessageContent
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateMessageEditedType:
		var update tdlib.UpdateMessageEdited
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateMessageIsPinnedType:
		var update tdlib.UpdateMessageIsPinned
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateMessageInteractionInfoType:
		var update tdlib.UpdateMessageInteractionInfo
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateMessageContentOpenedType:
		var update tdlib.UpdateMessageContentOpened
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateMessageMentionReadType:
		var update tdlib.UpdateMessageMentionRead
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateMessageLiveLocationViewedType:
		var update tdlib.UpdateMessageLiveLocationViewed
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateNewChatType:
		var update tdlib.UpdateNewChat
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateChatTitleType:
		var update tdlib.UpdateChatTitle
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateChatPhotoType:
		var update tdlib.UpdateChatPhoto
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateChatPermissionsType:
		var update tdlib.UpdateChatPermissions
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateChatLastMessageType:
		var update tdlib.UpdateChatLastMessage
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateChatPositionType:
		var update tdlib.UpdateChatPosition
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateChatIsMarkedAsUnreadType:
		var update tdlib.UpdateChatIsMarkedAsUnread
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateChatIsBlockedType:
		var update tdlib.UpdateChatIsBlocked
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateChatHasScheduledMessagesType:
		var update tdlib.UpdateChatHasScheduledMessages
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateChatVoiceChatType:
		var update tdlib.UpdateChatVoiceChat
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateChatDefaultDisableNotificationType:
		var update tdlib.UpdateChatDefaultDisableNotification
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateChatReadInboxType:
		var update tdlib.UpdateChatReadInbox
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateChatReadOutboxType:
		var update tdlib.UpdateChatReadOutbox
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateChatUnreadMentionCountType:
		var update tdlib.UpdateChatUnreadMentionCount
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateChatNotificationSettingsType:
		var update tdlib.UpdateChatNotificationSettings
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateScopeNotificationSettingsType:
		var update tdlib.UpdateScopeNotificationSettings
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateChatMessageTTLSettingType:
		var update tdlib.UpdateChatMessageTTLSetting
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateChatActionBarType:
		var update tdlib.UpdateChatActionBar
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateChatThemeType:
		var update tdlib.UpdateChatTheme
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateChatReplyMarkupType:
		var update tdlib.UpdateChatReplyMarkup
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateChatDraftMessageType:
		var update tdlib.UpdateChatDraftMessage
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateChatFiltersType:
		var update tdlib.UpdateChatFilters
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateChatOnlineMemberCountType:
		var update tdlib.UpdateChatOnlineMemberCount
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateNotificationType:
		var update tdlib.UpdateNotification
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateNotificationGroupType:
		var update tdlib.UpdateNotificationGroup
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateActiveNotificationsType:
		var update tdlib.UpdateActiveNotifications
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateHavePendingNotificationsType:
		var update tdlib.UpdateHavePendingNotifications
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateDeleteMessagesType:
		var update tdlib.UpdateDeleteMessages
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateUserChatActionType:
		var update tdlib.UpdateUserChatAction
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateUserStatusType:
		var update tdlib.UpdateUserStatus
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateUserType:
		var update tdlib.UpdateUser
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateBasicGroupType:
		var update tdlib.UpdateBasicGroup
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateSupergroupType:
		var update tdlib.UpdateSupergroup
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateSecretChatType:
		var update tdlib.UpdateSecretChat
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateUserFullInfoType:
		var update tdlib.UpdateUserFullInfo
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateBasicGroupFullInfoType:
		var update tdlib.UpdateBasicGroupFullInfo
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateSupergroupFullInfoType:
		var update tdlib.UpdateSupergroupFullInfo
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateServiceNotificationType:
		var update tdlib.UpdateServiceNotification
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateFileType:
		var update tdlib.UpdateFile
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateFileGenerationStartType:
		var update tdlib.UpdateFileGenerationStart
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateFileGenerationStopType:
		var update tdlib.UpdateFileGenerationStop
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateCallType:
		var update tdlib.UpdateCall
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateGroupCallType:
		var update tdlib.UpdateGroupCall
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateGroupCallParticipantType:
		var update tdlib.UpdateGroupCallParticipant
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateNewCallSignalingDataType:
		var update tdlib.UpdateNewCallSignalingData
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateUserPrivacySettingRulesType:
		var update tdlib.UpdateUserPrivacySettingRules
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateUnreadMessageCountType:
		var update tdlib.UpdateUnreadMessageCount
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateUnreadChatCountType:
		var update tdlib.UpdateUnreadChatCount
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateOptionType:
		var update tdlib.UpdateOption
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateStickerSetType:
		var update tdlib.UpdateStickerSet
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateInstalledStickerSetsType:
		var update tdlib.UpdateInstalledStickerSets
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateTrendingStickerSetsType:
		var update tdlib.UpdateTrendingStickerSets
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateRecentStickersType:
		var update tdlib.UpdateRecentStickers
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateFavoriteStickersType:
		var update tdlib.UpdateFavoriteStickers
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateSavedAnimationsType:
		var update tdlib.UpdateSavedAnimations
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateSelectedBackgroundType:
		var update tdlib.UpdateSelectedBackground
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateChatThemesType:
		var update tdlib.UpdateChatThemes
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateLanguagePackStringsType:
		var update tdlib.UpdateLanguagePackStrings
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateConnectionStateType:
		var update tdlib.UpdateConnectionState
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateTermsOfServiceType:
		var update tdlib.UpdateTermsOfService
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateUsersNearbyType:
		var update tdlib.UpdateUsersNearby
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateDiceEmojisType:
		var update tdlib.UpdateDiceEmojis
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateAnimationSearchParametersType:
		var update tdlib.UpdateAnimationSearchParameters
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateSuggestedActionsType:
		var update tdlib.UpdateSuggestedActions
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateNewInlineQueryType:
		var update tdlib.UpdateNewInlineQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateNewChosenInlineResultType:
		var update tdlib.UpdateNewChosenInlineResult
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateNewCallbackQueryType:
		var update tdlib.UpdateNewCallbackQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateNewInlineCallbackQueryType:
		var update tdlib.UpdateNewInlineCallbackQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateNewShippingQueryType:
		var update tdlib.UpdateNewShippingQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateNewPreCheckoutQueryType:
		var update tdlib.UpdateNewPreCheckoutQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateNewCustomEventType:
		var update tdlib.UpdateNewCustomEvent
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateNewCustomQueryType:
		var update tdlib.UpdateNewCustomQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdatePollType:
		var update tdlib.UpdatePoll
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdatePollAnswerType:
		var update tdlib.UpdatePollAnswer
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case tdlib.UpdateChatMemberType:
		var update tdlib.UpdateChatMember
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}
