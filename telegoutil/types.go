package telegoutil

import (
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegoapi"
)

// ID creates telego.ChatID from user's identifier
func ID(id int64) telego.ChatID {
	return telego.ChatID{ID: id}
}

// Username creates telego.ChatID from username
func Username(username string) telego.ChatID {
	return telego.ChatID{Username: username}
}

// File creates telego.InputFile from telegoapi.NamedReader
func File(file telegoapi.NamedReader) telego.InputFile {
	return telego.InputFile{
		File: file,
	}
}

// FileByURL creates telego.InputFile from URL
func FileByURL(url string) telego.InputFile {
	return telego.InputFile{
		URL: url,
	}
}

// FileByID creates telego.InputFile from file ID
func FileByID(id string) telego.InputFile {
	return telego.InputFile{
		FileID: id,
	}
}

// Keyboard creates telego.ReplyKeyboardMarkup from slice of keyboard buttons
func Keyboard(rows ...[]telego.KeyboardButton) *telego.ReplyKeyboardMarkup {
	return &telego.ReplyKeyboardMarkup{
		Keyboard: rows,
	}
}

// KeyboardRow creates slice of telego.KeyboardButton
func KeyboardRow(buttons ...telego.KeyboardButton) []telego.KeyboardButton {
	return buttons
}

// KeyboardButton creates telego.KeyboardButton with required fields
func KeyboardButton(text string) telego.KeyboardButton {
	return telego.KeyboardButton{
		Text: text,
	}
}

// PollTypeAny creates telego.KeyboardButtonPollType with any type
func PollTypeAny() *telego.KeyboardButtonPollType {
	return &telego.KeyboardButtonPollType{}
}

// PollTypeRegular creates telego.KeyboardButtonPollType with type regular
func PollTypeRegular() *telego.KeyboardButtonPollType {
	return &telego.KeyboardButtonPollType{Type: telego.PollTypeRegular}
}

// PollTypeQuiz creates telego.KeyboardButtonPollType with type quiz
func PollTypeQuiz() *telego.KeyboardButtonPollType {
	return &telego.KeyboardButtonPollType{Type: telego.PollTypeQuiz}
}

// InlineKeyboard creates telego.InlineKeyboardMarkup from slice of keyboard buttons
func InlineKeyboard(rows ...[]telego.InlineKeyboardButton) *telego.InlineKeyboardMarkup {
	return &telego.InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}
}

// InlineKeyboardRow creates slice of telego.InlineKeyboardButton
func InlineKeyboardRow(buttons ...telego.InlineKeyboardButton) []telego.InlineKeyboardButton {
	return buttons
}

// InlineKeyboardButton creates telego.InlineKeyboardButton with required fields
func InlineKeyboardButton(text string) telego.InlineKeyboardButton {
	return telego.InlineKeyboardButton{
		Text: text,
	}
}

// ResultCachedAudio creates telego.InlineQueryResultCachedAudio with required fields
func ResultCachedAudio(id, audioFileID string) *telego.InlineQueryResultCachedAudio {
	return &telego.InlineQueryResultCachedAudio{
		Type:        telego.ResultTypeAudio,
		ID:          id,
		AudioFileID: audioFileID,
	}
}

// ResultCachedDocument creates telego.InlineQueryResultCachedDocument with required fields
func ResultCachedDocument(id, title, documentFileID string) *telego.InlineQueryResultCachedDocument {
	return &telego.InlineQueryResultCachedDocument{
		Type:           telego.ResultTypeDocument,
		ID:             id,
		Title:          title,
		DocumentFileID: documentFileID,
	}
}

// ResultCachedGif creates telego.InlineQueryResultCachedGif with required fields
func ResultCachedGif(id, gifFileID string) *telego.InlineQueryResultCachedGif {
	return &telego.InlineQueryResultCachedGif{
		Type:      telego.ResultTypeGif,
		ID:        id,
		GifFileID: gifFileID,
	}
}

// ResultCachedMpeg4Gif creates telego.InlineQueryResultCachedMpeg4Gif with required fields
func ResultCachedMpeg4Gif(id, mpeg4FileID string) *telego.InlineQueryResultCachedMpeg4Gif {
	return &telego.InlineQueryResultCachedMpeg4Gif{
		Type:        telego.ResultTypeMpeg4Gif,
		ID:          id,
		Mpeg4FileID: mpeg4FileID,
	}
}

// ResultCachedPhoto creates telego.InlineQueryResultCachedPhoto with required fields
func ResultCachedPhoto(id, photoFileID string) *telego.InlineQueryResultCachedPhoto {
	return &telego.InlineQueryResultCachedPhoto{
		Type:        telego.ResultTypePhoto,
		ID:          id,
		PhotoFileID: photoFileID,
	}
}

// ResultCachedSticker creates telego.InlineQueryResultCachedSticker with required fields
func ResultCachedSticker(id, stickerFileID string) *telego.InlineQueryResultCachedSticker {
	return &telego.InlineQueryResultCachedSticker{
		Type:          telego.ResultTypeSticker,
		ID:            id,
		StickerFileID: stickerFileID,
	}
}

// ResultCachedVideo creates telego.InlineQueryResultCachedVideo with required fields
func ResultCachedVideo(id, videoFileID, title string) *telego.InlineQueryResultCachedVideo {
	return &telego.InlineQueryResultCachedVideo{
		Type:        telego.ResultTypeVideo,
		ID:          id,
		VideoFileID: videoFileID,
		Title:       title,
	}
}

// ResultCachedVoice creates telego.InlineQueryResultCachedVoice with required fields
func ResultCachedVoice(id, voiceFileID, title string) *telego.InlineQueryResultCachedVoice {
	return &telego.InlineQueryResultCachedVoice{
		Type:        telego.ResultTypeVoice,
		ID:          id,
		VoiceFileID: voiceFileID,
		Title:       title,
	}
}

// ResultArticle creates telego.InlineQueryResultArticle with required fields
func ResultArticle(id, title string, inputMessageContent telego.InputMessageContent,
) *telego.InlineQueryResultArticle {
	return &telego.InlineQueryResultArticle{
		Type:                telego.ResultTypeArticle,
		ID:                  id,
		Title:               title,
		InputMessageContent: inputMessageContent,
	}
}

// ResultAudio creates telego.InlineQueryResultAudio with required fields
func ResultAudio(id, audioURL, title string) *telego.InlineQueryResultAudio {
	return &telego.InlineQueryResultAudio{
		Type:     telego.ResultTypeAudio,
		ID:       id,
		AudioURL: audioURL,
		Title:    title,
	}
}

// ResultContact creates telego.InlineQueryResultContact with required fields
func ResultContact(id, phoneNumber, firstName string) *telego.InlineQueryResultContact {
	return &telego.InlineQueryResultContact{
		Type:        telego.ResultTypeContact,
		ID:          id,
		PhoneNumber: phoneNumber,
		FirstName:   firstName,
	}
}

// ResultGame creates telego.InlineQueryResultGame with required fields
func ResultGame(id, gameShortName string) *telego.InlineQueryResultGame {
	return &telego.InlineQueryResultGame{
		Type:          telego.ResultTypeGame,
		ID:            id,
		GameShortName: gameShortName,
	}
}

// ResultDocument creates telego.InlineQueryResultDocument with required fields
func ResultDocument(id, title, documentURL, mimeType string) *telego.InlineQueryResultDocument {
	return &telego.InlineQueryResultDocument{
		Type:        telego.ResultTypeDocument,
		ID:          id,
		Title:       title,
		DocumentURL: documentURL,
		MimeType:    mimeType,
	}
}

// ResultGif creates telego.InlineQueryResultGif with required fields
func ResultGif(id, gifURL, thumbURL string) *telego.InlineQueryResultGif {
	return &telego.InlineQueryResultGif{
		Type:     telego.ResultTypeGif,
		ID:       id,
		GifURL:   gifURL,
		ThumbURL: thumbURL,
	}
}

// ResultLocation creates telego.InlineQueryResultLocation with required fields
func ResultLocation(id string, latitude, longitude float64, title string) *telego.InlineQueryResultLocation {
	return &telego.InlineQueryResultLocation{
		Type:      telego.ResultTypeLocation,
		ID:        id,
		Latitude:  latitude,
		Longitude: longitude,
		Title:     title,
	}
}

// ResultMpeg4Gif creates telego.InlineQueryResultMpeg4Gif with required fields
func ResultMpeg4Gif(id, mpeg4URL, thumbURL string) *telego.InlineQueryResultMpeg4Gif {
	return &telego.InlineQueryResultMpeg4Gif{
		Type:     telego.ResultTypeMpeg4Gif,
		ID:       id,
		Mpeg4URL: mpeg4URL,
		ThumbURL: thumbURL,
	}
}

// ResultPhoto creates telego.InlineQueryResultPhoto with required fields
func ResultPhoto(id, photoURL, thumbURL string) *telego.InlineQueryResultPhoto {
	return &telego.InlineQueryResultPhoto{
		Type:     telego.ResultTypePhoto,
		ID:       id,
		PhotoURL: photoURL,
		ThumbURL: thumbURL,
	}
}

// ResultVenue creates telego.InlineQueryResultVenue with required fields
func ResultVenue(id string, latitude, longitude float64, title, address string,
) *telego.InlineQueryResultVenue {
	return &telego.InlineQueryResultVenue{
		Type:      telego.ResultTypeVenue,
		ID:        id,
		Latitude:  latitude,
		Longitude: longitude,
		Title:     title,
		Address:   address,
	}
}

// ResultVideo creates telego.InlineQueryResultVideo with required fields
func ResultVideo(id, videoURL, mimeType, thumbURL, title string) *telego.InlineQueryResultVideo {
	return &telego.InlineQueryResultVideo{
		Type:     telego.ResultTypeVideo,
		ID:       id,
		VideoURL: videoURL,
		MimeType: mimeType,
		ThumbURL: thumbURL,
		Title:    title,
	}
}

// ResultVoice creates telego.InlineQueryResultVoice with required fields
func ResultVoice(id, voiceURL, title string) *telego.InlineQueryResultVoice {
	return &telego.InlineQueryResultVoice{
		Type:     telego.ResultTypeVoice,
		ID:       id,
		VoiceURL: voiceURL,
		Title:    title,
	}
}

// TextMessage creates telego.InputTextMessageContent with required fields
func TextMessage(messageText string) *telego.InputTextMessageContent {
	return &telego.InputTextMessageContent{
		MessageText: messageText,
	}
}

// LocationMessage creates telego.InputLocationMessageContent with required fields
func LocationMessage(latitude, longitude float64) *telego.InputLocationMessageContent {
	return &telego.InputLocationMessageContent{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// VenueMessage creates telego.InputVenueMessageContent with required fields
func VenueMessage(latitude, longitude float64, title, address string) *telego.InputVenueMessageContent {
	return &telego.InputVenueMessageContent{
		Latitude:  latitude,
		Longitude: longitude,
		Title:     title,
		Address:   address,
	}
}

// ContactMessage creates telego.InputContactMessageContent with required fields
func ContactMessage(phoneNumber, firstName string) *telego.InputContactMessageContent {
	return &telego.InputContactMessageContent{
		PhoneNumber: phoneNumber,
		FirstName:   firstName,
	}
}

// InvoiceMessage creates telego.InputInvoiceMessageContent with required fields
func InvoiceMessage(title, description, payload, providerToken, currency string, prices ...telego.LabeledPrice,
) *telego.InputInvoiceMessageContent {
	return &telego.InputInvoiceMessageContent{
		Title:         title,
		Description:   description,
		Payload:       payload,
		ProviderToken: providerToken,
		Currency:      currency,
		Prices:        prices,
	}
}

// MediaAnimation creates telego.InputMediaAnimation with required fields
func MediaAnimation(media telego.InputFile) *telego.InputMediaAnimation {
	return &telego.InputMediaAnimation{
		Type:  telego.MediaTypeAnimation,
		Media: media,
	}
}

// MediaDocument creates telego.InputMediaDocument with required fields
func MediaDocument(media telego.InputFile) *telego.InputMediaDocument {
	return &telego.InputMediaDocument{
		Type:  telego.MediaTypeDocument,
		Media: media,
	}
}

// MediaAudio creates telego.InputMediaAudio with required fields
func MediaAudio(media telego.InputFile) *telego.InputMediaAudio {
	return &telego.InputMediaAudio{
		Type:  telego.MediaTypeAudio,
		Media: media,
	}
}

// MediaPhoto creates telego.InputMediaPhoto with required fields
func MediaPhoto(media telego.InputFile) *telego.InputMediaPhoto {
	return &telego.InputMediaPhoto{
		Type:  telego.MediaTypePhoto,
		Media: media,
	}
}

// MediaVideo creates telego.InputMediaVideo with required fields
func MediaVideo(media telego.InputFile) *telego.InputMediaVideo {
	return &telego.InputMediaVideo{
		Type:  telego.MediaTypeVideo,
		Media: media,
	}
}
