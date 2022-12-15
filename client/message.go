// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/rjazhenka/go-tdlib/v2/tdlib"
)

// GetMessage Returns information about a message
// @param chatID Identifier of the chat the message belongs to
// @param messageID Identifier of the message to get
func (client *Client) GetMessage(chatID int64, messageID int64) (*tdlib.Message, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":      "getMessage",
		"chat_id":    chatID,
		"message_id": messageID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var messageDummy tdlib.Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// GetMessageLocally Returns information about a message, if it is available locally without sending network request. This is an offline request
// @param chatID Identifier of the chat the message belongs to
// @param messageID Identifier of the message to get
func (client *Client) GetMessageLocally(chatID int64, messageID int64) (*tdlib.Message, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":      "getMessageLocally",
		"chat_id":    chatID,
		"message_id": messageID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var messageDummy tdlib.Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// GetRepliedMessage Returns information about a message that is replied by a given message. Also returns the pinned message, the game message, and the invoice message for messages of the types messagePinMessage, messageGameScore, and messagePaymentSuccessful respectively
// @param chatID Identifier of the chat the message belongs to
// @param messageID Identifier of the message reply to which to get
func (client *Client) GetRepliedMessage(chatID int64, messageID int64) (*tdlib.Message, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":      "getRepliedMessage",
		"chat_id":    chatID,
		"message_id": messageID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var messageDummy tdlib.Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// GetChatPinnedMessage Returns information about a newest pinned message in the chat
// @param chatID Identifier of the chat the message belongs to
func (client *Client) GetChatPinnedMessage(chatID int64) (*tdlib.Message, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":   "getChatPinnedMessage",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var message tdlib.Message
	err = json.Unmarshal(result.Raw, &message)
	return &message, err

}

// GetCallbackQueryMessage Returns information about a message with the callback button that originated a callback query; for bots only
// @param chatID Identifier of the chat the message belongs to
// @param messageID Message identifier
// @param callbackQueryID Identifier of the callback query
func (client *Client) GetCallbackQueryMessage(chatID int64, messageID int64, callbackQueryID *tdlib.JSONInt64) (*tdlib.Message, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":             "getCallbackQueryMessage",
		"chat_id":           chatID,
		"message_id":        messageID,
		"callback_query_id": callbackQueryID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var messageDummy tdlib.Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// GetChatMessageByDate Returns the last message sent in a chat no later than the specified date
// @param chatID Chat identifier
// @param date Point in time (Unix timestamp) relative to which to search for messages
func (client *Client) GetChatMessageByDate(chatID int64, date int32) (*tdlib.Message, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":   "getChatMessageByDate",
		"chat_id": chatID,
		"date":    date,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var message tdlib.Message
	err = json.Unmarshal(result.Raw, &message)
	return &message, err

}

// SendMessage Sends a message. Returns the sent message
// @param chatID Target chat
// @param messageThreadID If not 0, a message thread identifier in which the message will be sent
// @param replyToMessageID Identifier of the message to reply to or 0
// @param options Options to be used to send the message
// @param replyMarkup Markup for replying to the message; for bots only
// @param inputMessageContent The content of the message to be sent
func (client *Client) SendMessage(chatID int64, messageThreadID int64, replyToMessageID int64, options *tdlib.MessageSendOptions, replyMarkup tdlib.ReplyMarkup, inputMessageContent tdlib.InputMessageContent) (*tdlib.Message, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":                 "sendMessage",
		"chat_id":               chatID,
		"message_thread_id":     messageThreadID,
		"reply_to_message_id":   replyToMessageID,
		"options":               options,
		"reply_markup":          replyMarkup,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var messageDummy tdlib.Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// SendBotStartMessage Invites a bot to a chat (if it is not yet a member) and sends it the /start command. Bots can't be invited to a private chat other than the chat with the bot. Bots can't be invited to channels (although they can be added as admins) and secret chats. Returns the sent message
// @param botUserID Identifier of the bot
// @param chatID Identifier of the target chat
// @param parameter A hidden parameter sent to the bot for deep linking purposes (https://core.telegram.org/bots#deep-linking)
func (client *Client) SendBotStartMessage(botUserID int64, chatID int64, parameter string) (*tdlib.Message, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":       "sendBotStartMessage",
		"bot_user_id": botUserID,
		"chat_id":     chatID,
		"parameter":   parameter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var message tdlib.Message
	err = json.Unmarshal(result.Raw, &message)
	return &message, err

}

// SendInlineQueryResultMessage Sends the result of an inline query as a message. Returns the sent message. Always clears a chat draft message
// @param chatID Target chat
// @param messageThreadID If not 0, a message thread identifier in which the message will be sent
// @param replyToMessageID Identifier of a message to reply to or 0
// @param options Options to be used to send the message
// @param queryID Identifier of the inline query
// @param resultID Identifier of the inline result
// @param hideViaBot If true, there will be no mention of a bot, via which the message is sent. Can be used only for bots GetOption("animation_search_bot_username"), GetOption("photo_search_bot_username") and GetOption("venue_search_bot_username")
func (client *Client) SendInlineQueryResultMessage(chatID int64, messageThreadID int64, replyToMessageID int64, options *tdlib.MessageSendOptions, queryID *tdlib.JSONInt64, resultID string, hideViaBot bool) (*tdlib.Message, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":               "sendInlineQueryResultMessage",
		"chat_id":             chatID,
		"message_thread_id":   messageThreadID,
		"reply_to_message_id": replyToMessageID,
		"options":             options,
		"query_id":            queryID,
		"result_id":           resultID,
		"hide_via_bot":        hideViaBot,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var messageDummy tdlib.Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// AddLocalMessage Adds a local message to a chat. The message is persistent across application restarts only if the message database is used. Returns the added message
// @param chatID Target chat
// @param sender The sender sender of the message
// @param replyToMessageID Identifier of the message to reply to or 0
// @param disableNotification Pass true to disable notification for the message
// @param inputMessageContent The content of the message to be added
func (client *Client) AddLocalMessage(chatID int64, sender tdlib.MessageSender, replyToMessageID int64, disableNotification bool, inputMessageContent tdlib.InputMessageContent) (*tdlib.Message, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":                 "addLocalMessage",
		"chat_id":               chatID,
		"sender":                sender,
		"reply_to_message_id":   replyToMessageID,
		"disable_notification":  disableNotification,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var messageDummy tdlib.Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// EditMessageText Edits the text of a message (or a text of a game message). Returns the edited message after the edit is completed on the server side
// @param chatID The chat the message belongs to
// @param messageID Identifier of the message
// @param replyMarkup The new message reply markup; for bots only
// @param inputMessageContent New text content of the message. Should be of type inputMessageText
func (client *Client) EditMessageText(chatID int64, messageID int64, replyMarkup tdlib.ReplyMarkup, inputMessageContent tdlib.InputMessageContent) (*tdlib.Message, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":                 "editMessageText",
		"chat_id":               chatID,
		"message_id":            messageID,
		"reply_markup":          replyMarkup,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var messageDummy tdlib.Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// EditMessageLiveLocation Edits the message content of a live location. Messages can be edited for a limited period of time specified in the live location. Returns the edited message after the edit is completed on the server side
// @param chatID The chat the message belongs to
// @param messageID Identifier of the message
// @param replyMarkup The new message reply markup; for bots only
// @param location New location content of the message; may be null. Pass null to stop sharing the live location
// @param heading The new direction in which the location moves, in degrees; 1-360. Pass 0 if unknown
// @param proximityAlertRadius The new maximum distance for proximity alerts, in meters (0-100000). Pass 0 if the notification is disabled
func (client *Client) EditMessageLiveLocation(chatID int64, messageID int64, replyMarkup tdlib.ReplyMarkup, location *tdlib.Location, heading int32, proximityAlertRadius int32) (*tdlib.Message, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":                  "editMessageLiveLocation",
		"chat_id":                chatID,
		"message_id":             messageID,
		"reply_markup":           replyMarkup,
		"location":               location,
		"heading":                heading,
		"proximity_alert_radius": proximityAlertRadius,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var messageDummy tdlib.Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// EditMessageMedia Edits the content of a message with an animation, an audio, a document, a photo or a video, including message caption. If only the caption needs to be edited, use editMessageCaption instead. The media can't be edited if the message was set to self-destruct or to a self-destructing media. The type of message content in an album can't be changed with exception of replacing a photo with a video or vice versa. Returns the edited message after the edit is completed on the server side
// @param chatID The chat the message belongs to
// @param messageID Identifier of the message
// @param replyMarkup The new message reply markup; for bots only
// @param inputMessageContent New content of the message. Must be one of the following types: inputMessageAnimation, inputMessageAudio, inputMessageDocument, inputMessagePhoto or inputMessageVideo
func (client *Client) EditMessageMedia(chatID int64, messageID int64, replyMarkup tdlib.ReplyMarkup, inputMessageContent tdlib.InputMessageContent) (*tdlib.Message, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":                 "editMessageMedia",
		"chat_id":               chatID,
		"message_id":            messageID,
		"reply_markup":          replyMarkup,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var messageDummy tdlib.Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// EditMessageCaption Edits the message content caption. Returns the edited message after the edit is completed on the server side
// @param chatID The chat the message belongs to
// @param messageID Identifier of the message
// @param replyMarkup The new message reply markup; for bots only
// @param caption New message content caption; 0-GetOption("message_caption_length_max") characters
func (client *Client) EditMessageCaption(chatID int64, messageID int64, replyMarkup tdlib.ReplyMarkup, caption *tdlib.FormattedText) (*tdlib.Message, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":        "editMessageCaption",
		"chat_id":      chatID,
		"message_id":   messageID,
		"reply_markup": replyMarkup,
		"caption":      caption,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var messageDummy tdlib.Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// EditMessageReplyMarkup Edits the message reply markup; for bots only. Returns the edited message after the edit is completed on the server side
// @param chatID The chat the message belongs to
// @param messageID Identifier of the message
// @param replyMarkup The new message reply markup
func (client *Client) EditMessageReplyMarkup(chatID int64, messageID int64, replyMarkup tdlib.ReplyMarkup) (*tdlib.Message, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":        "editMessageReplyMarkup",
		"chat_id":      chatID,
		"message_id":   messageID,
		"reply_markup": replyMarkup,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var messageDummy tdlib.Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// SetGameScore Updates the game score of the specified user in the game; for bots only
// @param chatID The chat to which the message with the game belongs
// @param messageID Identifier of the message
// @param editMessage True, if the message should be edited
// @param userID User identifier
// @param score The new score
// @param force Pass true to update the score even if it decreases. If the score is 0, the user will be deleted from the high score table
func (client *Client) SetGameScore(chatID int64, messageID int64, editMessage bool, userID int64, score int32, force bool) (*tdlib.Message, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":        "setGameScore",
		"chat_id":      chatID,
		"message_id":   messageID,
		"edit_message": editMessage,
		"user_id":      userID,
		"score":        score,
		"force":        force,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var messageDummy tdlib.Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}
