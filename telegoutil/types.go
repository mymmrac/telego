package telegoutil

import (
	"fmt"
	"io"

	"github.com/valyala/fasthttp"

	"github.com/mymmrac/telego"
	ta "github.com/mymmrac/telego/telegoapi"
)

// ID creates [telego.ChatID] from user's identifier
func ID(id int64) telego.ChatID {
	return telego.ChatID{
		ID: id,
	}
}

// Username creates [telego.ChatID] from username
func Username(username string) telego.ChatID {
	return telego.ChatID{
		Username: username,
	}
}

// File creates [telego.InputFile] from [ta.NamedReader]
func File(file ta.NamedReader) telego.InputFile {
	return telego.InputFile{
		File: file,
	}
}

// FileFromReader creates [telego.InputFile] from [io.Reader] and name
func FileFromReader(reader io.Reader, name string) telego.InputFile {
	return telego.InputFile{
		File: NameReader(reader, name),
	}
}

// FileFromBytes creates [telego.InputFile] from slice of bytes and name
func FileFromBytes(data []byte, name string) telego.InputFile {
	return telego.InputFile{
		File: NameBytes(data, name),
	}
}

// FileFromURL creates [telego.InputFile] from URL
func FileFromURL(url string) telego.InputFile {
	return telego.InputFile{
		URL: url,
	}
}

// FileFromID creates [telego.InputFile] from file ID
func FileFromID(id string) telego.InputFile {
	return telego.InputFile{
		FileID: id,
	}
}

// DownloadFile returns downloaded file bytes or error
func DownloadFile(url string) ([]byte, error) {
	var file []byte
	status, file, err := fasthttp.Get(file, url)
	if err != nil {
		return nil, fmt.Errorf("telego: %w", err)
	}

	if status != fasthttp.StatusOK {
		return nil, fmt.Errorf("telego: http status: %d", status)
	}

	return file, nil
}

// Keyboard creates [telego.ReplyKeyboardMarkup] from slice of keyboard buttons
func Keyboard(rows ...[]telego.KeyboardButton) *telego.ReplyKeyboardMarkup {
	return &telego.ReplyKeyboardMarkup{
		Keyboard: rows,
	}
}

// KeyboardRow creates a slice of [telego.KeyboardButton]
func KeyboardRow(buttons ...telego.KeyboardButton) []telego.KeyboardButton {
	return buttons
}

// KeyboardGrid creates a [telego.ReplyKeyboardMarkup] from grid of buttons
func KeyboardGrid(buttons [][]telego.KeyboardButton) *telego.ReplyKeyboardMarkup {
	return &telego.ReplyKeyboardMarkup{
		Keyboard: buttons,
	}
}

// KeyboardCols creates a grid of buttons containing specified number of columns
func KeyboardCols(cols int, buttons ...telego.KeyboardButton) [][]telego.KeyboardButton {
	if cols <= 0 {
		return nil
	}

	col := 0
	row := 0

	rows := len(buttons) / cols
	if len(buttons)%cols != 0 {
		rows++
	}

	grid := make([][]telego.KeyboardButton, 0, rows)
	for i := range buttons {
		if col >= cols {
			col = 0
			row++
		}
		if col == 0 {
			grid = append(grid, make([]telego.KeyboardButton, 0, cols))
		}
		grid[row] = append(grid[row], buttons[i])
		col++
	}

	return grid
}

// KeyboardRows creates a grid of buttons containing specified number of rows
func KeyboardRows(rows int, buttons ...telego.KeyboardButton) [][]telego.KeyboardButton {
	if rows <= 0 {
		return nil
	}

	col := 0
	row := 0

	cols := len(buttons) / rows
	if len(buttons)%rows != 0 {
		cols++
	}

	grid := make([][]telego.KeyboardButton, 0, rows)
	for i := range buttons {
		if col >= cols {
			col = 0
			row++
		}
		if col == 0 {
			grid = append(grid, make([]telego.KeyboardButton, 0, cols))
		}
		grid[row] = append(grid[row], buttons[i])
		col++
	}

	return grid
}

// KeyboardButton creates [telego.KeyboardButton] with required fields
func KeyboardButton(text string) telego.KeyboardButton {
	return telego.KeyboardButton{
		Text: text,
	}
}

// ReplyKeyboardRemove creates [telego.ReplyKeyboardRemove] with required fields
func ReplyKeyboardRemove() *telego.ReplyKeyboardRemove {
	return &telego.ReplyKeyboardRemove{
		RemoveKeyboard: true,
	}
}

// WebAppInfo creates [telego.WebAppInfo] with required fields
func WebAppInfo(url string) *telego.WebAppInfo {
	return &telego.WebAppInfo{
		URL: url,
	}
}

// ForceReply creates [telego.ForceReply] with required fields
func ForceReply() *telego.ForceReply {
	return &telego.ForceReply{
		ForceReply: true,
	}
}

// PollTypeAny creates [telego.KeyboardButtonPollType] with any type
func PollTypeAny() *telego.KeyboardButtonPollType {
	return &telego.KeyboardButtonPollType{}
}

// PollTypeRegular creates [telego.KeyboardButtonPollType] with type regular
func PollTypeRegular() *telego.KeyboardButtonPollType {
	return &telego.KeyboardButtonPollType{
		Type: telego.PollTypeRegular,
	}
}

// PollTypeQuiz creates [telego.KeyboardButtonPollType] with type quiz
func PollTypeQuiz() *telego.KeyboardButtonPollType {
	return &telego.KeyboardButtonPollType{
		Type: telego.PollTypeQuiz,
	}
}

// InlineKeyboard creates [telego.InlineKeyboardMarkup] from slice of keyboard buttons rows
func InlineKeyboard(rows ...[]telego.InlineKeyboardButton) *telego.InlineKeyboardMarkup {
	return &telego.InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}
}

// InlineKeyboardRow creates a slice of telego.InlineKeyboardButton
func InlineKeyboardRow(buttons ...telego.InlineKeyboardButton) []telego.InlineKeyboardButton {
	return buttons
}

// InlineKeyboardGrid creates a [telego.InlineKeyboardMarkup] from grid of buttons
func InlineKeyboardGrid(buttons [][]telego.InlineKeyboardButton) *telego.InlineKeyboardMarkup {
	return &telego.InlineKeyboardMarkup{
		InlineKeyboard: buttons,
	}
}

// InlineKeyboardCols creates a grid of buttons containing specified number of columns
func InlineKeyboardCols(cols int, buttons ...telego.InlineKeyboardButton) [][]telego.InlineKeyboardButton {
	if cols <= 0 {
		return nil
	}

	col := 0
	row := 0

	rows := len(buttons) / cols
	if len(buttons)%cols != 0 {
		rows++
	}

	grid := make([][]telego.InlineKeyboardButton, 0, rows)
	for i := range buttons {
		if col >= cols {
			col = 0
			row++
		}
		if col == 0 {
			grid = append(grid, make([]telego.InlineKeyboardButton, 0, cols))
		}
		grid[row] = append(grid[row], buttons[i])
		col++
	}

	return grid
}

// InlineKeyboardRows creates a grid of buttons containing specified number of rows
func InlineKeyboardRows(rows int, buttons ...telego.InlineKeyboardButton) [][]telego.InlineKeyboardButton {
	if rows <= 0 {
		return nil
	}

	col := 0
	row := 0

	cols := len(buttons) / rows
	if len(buttons)%rows != 0 {
		cols++
	}

	grid := make([][]telego.InlineKeyboardButton, 0, rows)
	for i := range buttons {
		if col >= cols {
			col = 0
			row++
		}
		if col == 0 {
			grid = append(grid, make([]telego.InlineKeyboardButton, 0, cols))
		}
		grid[row] = append(grid[row], buttons[i])
		col++
	}

	return grid
}

// InlineKeyboardButton creates [telego.InlineKeyboardButton] with required fields
func InlineKeyboardButton(text string) telego.InlineKeyboardButton {
	return telego.InlineKeyboardButton{
		Text: text,
	}
}

// ResultCachedAudio creates [telego.InlineQueryResultCachedAudio] with required fields
func ResultCachedAudio(id, audioFileID string) *telego.InlineQueryResultCachedAudio {
	return &telego.InlineQueryResultCachedAudio{
		Type:        telego.ResultTypeAudio,
		ID:          id,
		AudioFileID: audioFileID,
	}
}

// ResultCachedDocument creates [telego.InlineQueryResultCachedDocument] with required fields
func ResultCachedDocument(id, title, documentFileID string) *telego.InlineQueryResultCachedDocument {
	return &telego.InlineQueryResultCachedDocument{
		Type:           telego.ResultTypeDocument,
		ID:             id,
		Title:          title,
		DocumentFileID: documentFileID,
	}
}

// ResultCachedGif creates [telego.InlineQueryResultCachedGif] with required fields
func ResultCachedGif(id, gifFileID string) *telego.InlineQueryResultCachedGif {
	return &telego.InlineQueryResultCachedGif{
		Type:      telego.ResultTypeGif,
		ID:        id,
		GifFileID: gifFileID,
	}
}

// ResultCachedMpeg4Gif creates [telego.InlineQueryResultCachedMpeg4Gif] with required fields
func ResultCachedMpeg4Gif(id, mpeg4FileID string) *telego.InlineQueryResultCachedMpeg4Gif {
	return &telego.InlineQueryResultCachedMpeg4Gif{
		Type:        telego.ResultTypeMpeg4Gif,
		ID:          id,
		Mpeg4FileID: mpeg4FileID,
	}
}

// ResultCachedPhoto creates [telego.InlineQueryResultCachedPhoto] with required fields
func ResultCachedPhoto(id, photoFileID string) *telego.InlineQueryResultCachedPhoto {
	return &telego.InlineQueryResultCachedPhoto{
		Type:        telego.ResultTypePhoto,
		ID:          id,
		PhotoFileID: photoFileID,
	}
}

// ResultCachedSticker creates [telego.InlineQueryResultCachedSticker] with required fields
func ResultCachedSticker(id, stickerFileID string) *telego.InlineQueryResultCachedSticker {
	return &telego.InlineQueryResultCachedSticker{
		Type:          telego.ResultTypeSticker,
		ID:            id,
		StickerFileID: stickerFileID,
	}
}

// ResultCachedVideo creates [telego.InlineQueryResultCachedVideo] with required fields
func ResultCachedVideo(id, videoFileID, title string) *telego.InlineQueryResultCachedVideo {
	return &telego.InlineQueryResultCachedVideo{
		Type:        telego.ResultTypeVideo,
		ID:          id,
		VideoFileID: videoFileID,
		Title:       title,
	}
}

// ResultCachedVoice creates [telego.InlineQueryResultCachedVoice] with required fields
func ResultCachedVoice(id, voiceFileID, title string) *telego.InlineQueryResultCachedVoice {
	return &telego.InlineQueryResultCachedVoice{
		Type:        telego.ResultTypeVoice,
		ID:          id,
		VoiceFileID: voiceFileID,
		Title:       title,
	}
}

// ResultArticle creates [telego.InlineQueryResultArticle] with required fields
func ResultArticle(id, title string, inputMessageContent telego.InputMessageContent,
) *telego.InlineQueryResultArticle {
	return &telego.InlineQueryResultArticle{
		Type:                telego.ResultTypeArticle,
		ID:                  id,
		Title:               title,
		InputMessageContent: inputMessageContent,
	}
}

// ResultAudio creates [telego.InlineQueryResultAudio] with required fields
func ResultAudio(id, audioURL, title string) *telego.InlineQueryResultAudio {
	return &telego.InlineQueryResultAudio{
		Type:     telego.ResultTypeAudio,
		ID:       id,
		AudioURL: audioURL,
		Title:    title,
	}
}

// ResultContact creates [telego.InlineQueryResultContact] with required fields
func ResultContact(id, phoneNumber, firstName string) *telego.InlineQueryResultContact {
	return &telego.InlineQueryResultContact{
		Type:        telego.ResultTypeContact,
		ID:          id,
		PhoneNumber: phoneNumber,
		FirstName:   firstName,
	}
}

// ResultGame creates [telego.InlineQueryResultGame] with required fields
func ResultGame(id, gameShortName string) *telego.InlineQueryResultGame {
	return &telego.InlineQueryResultGame{
		Type:          telego.ResultTypeGame,
		ID:            id,
		GameShortName: gameShortName,
	}
}

// ResultDocument creates [telego.InlineQueryResultDocument] with required fields
func ResultDocument(id, title, documentURL, mimeType string) *telego.InlineQueryResultDocument {
	return &telego.InlineQueryResultDocument{
		Type:        telego.ResultTypeDocument,
		ID:          id,
		Title:       title,
		DocumentURL: documentURL,
		MimeType:    mimeType,
	}
}

// ResultGif creates [telego.InlineQueryResultGif] with required fields
func ResultGif(id, gifURL, thumbnailURL string) *telego.InlineQueryResultGif {
	return &telego.InlineQueryResultGif{
		Type:         telego.ResultTypeGif,
		ID:           id,
		GifURL:       gifURL,
		ThumbnailURL: thumbnailURL,
	}
}

// ResultLocation creates [telego.InlineQueryResultLocation] with required fields
func ResultLocation(id string, latitude, longitude float64, title string) *telego.InlineQueryResultLocation {
	return &telego.InlineQueryResultLocation{
		Type:      telego.ResultTypeLocation,
		ID:        id,
		Latitude:  latitude,
		Longitude: longitude,
		Title:     title,
	}
}

// ResultMpeg4Gif creates [telego.InlineQueryResultMpeg4Gif] with required fields
func ResultMpeg4Gif(id, mpeg4URL, thumbnailURL string) *telego.InlineQueryResultMpeg4Gif {
	return &telego.InlineQueryResultMpeg4Gif{
		Type:         telego.ResultTypeMpeg4Gif,
		ID:           id,
		Mpeg4URL:     mpeg4URL,
		ThumbnailURL: thumbnailURL,
	}
}

// ResultPhoto creates [telego.InlineQueryResultPhoto] with required fields
func ResultPhoto(id, photoURL, thumbnailURL string) *telego.InlineQueryResultPhoto {
	return &telego.InlineQueryResultPhoto{
		Type:         telego.ResultTypePhoto,
		ID:           id,
		PhotoURL:     photoURL,
		ThumbnailURL: thumbnailURL,
	}
}

// ResultVenue creates [telego.InlineQueryResultVenue] with required fields
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

// ResultVideo creates [telego.InlineQueryResultVideo] with required fields
func ResultVideo(id, videoURL, mimeType, thumbnailURL, title string) *telego.InlineQueryResultVideo {
	return &telego.InlineQueryResultVideo{
		Type:         telego.ResultTypeVideo,
		ID:           id,
		VideoURL:     videoURL,
		MimeType:     mimeType,
		ThumbnailURL: thumbnailURL,
		Title:        title,
	}
}

// ResultVoice creates [telego.InlineQueryResultVoice] with required fields
func ResultVoice(id, voiceURL, title string) *telego.InlineQueryResultVoice {
	return &telego.InlineQueryResultVoice{
		Type:     telego.ResultTypeVoice,
		ID:       id,
		VoiceURL: voiceURL,
		Title:    title,
	}
}

// TextMessage creates [telego.InputTextMessageContent] with required fields
func TextMessage(messageText string) *telego.InputTextMessageContent {
	return &telego.InputTextMessageContent{
		MessageText: messageText,
	}
}

// LocationMessage creates [telego.InputLocationMessageContent] with required fields
func LocationMessage(latitude, longitude float64) *telego.InputLocationMessageContent {
	return &telego.InputLocationMessageContent{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// VenueMessage creates [telego.InputVenueMessageContent] with required fields
func VenueMessage(latitude, longitude float64, title, address string) *telego.InputVenueMessageContent {
	return &telego.InputVenueMessageContent{
		Latitude:  latitude,
		Longitude: longitude,
		Title:     title,
		Address:   address,
	}
}

// ContactMessage creates [telego.InputContactMessageContent] with required fields
func ContactMessage(phoneNumber, firstName string) *telego.InputContactMessageContent {
	return &telego.InputContactMessageContent{
		PhoneNumber: phoneNumber,
		FirstName:   firstName,
	}
}

// InvoiceMessage creates [telego.InputInvoiceMessageContent] with required fields
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

// MediaAnimation creates [telego.InputMediaAnimation] with required fields
func MediaAnimation(media telego.InputFile) *telego.InputMediaAnimation {
	return &telego.InputMediaAnimation{
		Type:  telego.MediaTypeAnimation,
		Media: media,
	}
}

// MediaAudio creates [telego.InputMediaAudio] with required fields
func MediaAudio(media telego.InputFile) *telego.InputMediaAudio {
	return &telego.InputMediaAudio{
		Type:  telego.MediaTypeAudio,
		Media: media,
	}
}

// MediaDocument creates [telego.InputMediaDocument] with required fields
func MediaDocument(media telego.InputFile) *telego.InputMediaDocument {
	return &telego.InputMediaDocument{
		Type:  telego.MediaTypeDocument,
		Media: media,
	}
}

// MediaLivePhoto creates [telego.InputMediaLivePhoto] with required fields
func MediaLivePhoto(media, photo telego.InputFile) *telego.InputMediaLivePhoto {
	return &telego.InputMediaLivePhoto{
		Type:  telego.MediaTypeLivePhoto,
		Media: media,
		Photo: photo,
	}
}

// MediaLocation creates [telego.InputMediaLocation] with required fields
func MediaLocation(latitude, longitude float64) *telego.InputMediaLocation {
	return &telego.InputMediaLocation{
		Type:      telego.MediaTypeLocation,
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// MediaPhoto creates [telego.InputMediaPhoto] with required fields
func MediaPhoto(media telego.InputFile) *telego.InputMediaPhoto {
	return &telego.InputMediaPhoto{
		Type:  telego.MediaTypePhoto,
		Media: media,
	}
}

// MediaSticker creates [telego.InputMediaSticker] with required fields
func MediaSticker(media telego.InputFile) *telego.InputMediaSticker {
	return &telego.InputMediaSticker{
		Type:  telego.MediaTypeSticker,
		Media: media,
	}
}

// MediaVenue creates [telego.InputMediaVenue] with required fields
func MediaVenue(latitude, longitude float64, title, address string) *telego.InputMediaVenue {
	return &telego.InputMediaVenue{
		Type:      telego.MediaTypeVenue,
		Latitude:  latitude,
		Longitude: longitude,
		Title:     title,
		Address:   address,
	}
}

// MediaVideo creates [telego.InputMediaVideo] with required fields
func MediaVideo(media telego.InputFile) *telego.InputMediaVideo {
	return &telego.InputMediaVideo{
		Type:  telego.MediaTypeVideo,
		Media: media,
	}
}

// ScopeDefault creates [telego.BotCommandScopeDefault] with required fields
func ScopeDefault() *telego.BotCommandScopeDefault {
	return &telego.BotCommandScopeDefault{
		Type: telego.ScopeTypeDefault,
	}
}

// ScopeAllPrivateChats creates [telego.BotCommandScopeAllPrivateChats] with required fields
func ScopeAllPrivateChats() *telego.BotCommandScopeAllPrivateChats {
	return &telego.BotCommandScopeAllPrivateChats{
		Type: telego.ScopeTypeAllPrivateChats,
	}
}

// ScopeAllGroupChats creates [telego.BotCommandScopeAllGroupChats] with required fields
func ScopeAllGroupChats() *telego.BotCommandScopeAllGroupChats {
	return &telego.BotCommandScopeAllGroupChats{
		Type: telego.ScopeTypeAllGroupChats,
	}
}

// ScopeAllChatAdministrators creates [telego.BotCommandScopeAllChatAdministrators] with required fields
func ScopeAllChatAdministrators() *telego.BotCommandScopeAllChatAdministrators {
	return &telego.BotCommandScopeAllChatAdministrators{
		Type: telego.ScopeTypeAllChatAdministrators,
	}
}

// ScopeChat creates [telego.BotCommandScopeChat] with required fields
func ScopeChat(chatID telego.ChatID) *telego.BotCommandScopeChat {
	return &telego.BotCommandScopeChat{
		Type:   telego.ScopeTypeChat,
		ChatID: chatID,
	}
}

// ScopeChatAdministrators creates [telego.BotCommandScopeChatAdministrators] with required fields
func ScopeChatAdministrators(chatID telego.ChatID) *telego.BotCommandScopeChatAdministrators {
	return &telego.BotCommandScopeChatAdministrators{
		Type:   telego.ScopeTypeChatAdministrators,
		ChatID: chatID,
	}
}

// ScopeChatMember creates [telego.BotCommandScopeChatMember] with required fields
func ScopeChatMember(chatID telego.ChatID, userID int64) *telego.BotCommandScopeChatMember {
	return &telego.BotCommandScopeChatMember{
		Type:   telego.ScopeTypeChatMember,
		ChatID: chatID,
		UserID: userID,
	}
}

// ErrorDataField creates [telego.PassportElementErrorDataField] with required fields
func ErrorDataField(sourceType, message, fieldName, dataHash string) *telego.PassportElementErrorDataField {
	return &telego.PassportElementErrorDataField{
		Source:    telego.ErrorSourceDataField,
		Type:      sourceType,
		FieldName: fieldName,
		DataHash:  dataHash,
		Message:   message,
	}
}

// ErrorFrontSide creates [telego.PassportElementErrorFrontSide] with required fields
func ErrorFrontSide(sourceType, message, fileHash string) *telego.PassportElementErrorFrontSide {
	return &telego.PassportElementErrorFrontSide{
		Source:   telego.ErrorSourceFrontSide,
		Type:     sourceType,
		FileHash: fileHash,
		Message:  message,
	}
}

// ErrorReverseSide creates [telego.PassportElementErrorReverseSide] with required fields
func ErrorReverseSide(sourceType, message, fileHash string) *telego.PassportElementErrorReverseSide {
	return &telego.PassportElementErrorReverseSide{
		Source:   telego.ErrorSourceReverseSide,
		Type:     sourceType,
		FileHash: fileHash,
		Message:  message,
	}
}

// ErrorSelfie creates [telego.PassportElementErrorSelfie] with required fields
func ErrorSelfie(sourceType, message, fileHash string) *telego.PassportElementErrorSelfie {
	return &telego.PassportElementErrorSelfie{
		Source:   telego.ErrorSourceSelfie,
		Type:     sourceType,
		FileHash: fileHash,
		Message:  message,
	}
}

// ErrorFile creates [telego.PassportElementErrorFile] with required fields
func ErrorFile(sourceType, message, fileHash string) *telego.PassportElementErrorFile {
	return &telego.PassportElementErrorFile{
		Source:   telego.ErrorSourceFile,
		Type:     sourceType,
		FileHash: fileHash,
		Message:  message,
	}
}

// ErrorFiles creates [telego.PassportElementErrorFiles] with required fields
func ErrorFiles(sourceType, message string, fileHashes ...string) *telego.PassportElementErrorFiles {
	return &telego.PassportElementErrorFiles{
		Source:     telego.ErrorSourceFiles,
		Type:       sourceType,
		FileHashes: fileHashes,
		Message:    message,
	}
}

// ErrorTranslationFile creates [telego.PassportElementErrorTranslationFile] with required fields
func ErrorTranslationFile(sourceType, message, fileHash string) *telego.PassportElementErrorTranslationFile {
	return &telego.PassportElementErrorTranslationFile{
		Source:   telego.ErrorSourceTranslationFile,
		Type:     sourceType,
		FileHash: fileHash,
		Message:  message,
	}
}

// ErrorTranslationFiles creates [telego.PassportElementErrorTranslationFiles] with required fields
func ErrorTranslationFiles(sourceType, message string, fileHashes ...string,
) *telego.PassportElementErrorTranslationFiles {
	return &telego.PassportElementErrorTranslationFiles{
		Source:     telego.ErrorSourceTranslationFiles,
		Type:       sourceType,
		FileHashes: fileHashes,
		Message:    message,
	}
}

// ErrorUnspecified creates [telego.PassportElementErrorUnspecified] with required fields
func ErrorUnspecified(sourceType, message, elementHash string) *telego.PassportElementErrorUnspecified {
	return &telego.PassportElementErrorUnspecified{
		Source:      telego.ErrorSourceUnspecified,
		Type:        sourceType,
		ElementHash: elementHash,
		Message:     message,
	}
}

// LabeledPrice creates [telego.LabeledPrice] with required fields
func LabeledPrice(label string, amount int) telego.LabeledPrice {
	return telego.LabeledPrice{
		Label:  label,
		Amount: amount,
	}
}

// ShippingOption creates [telego.ShippingOption] with required fields
func ShippingOption(id, title string, prices ...telego.LabeledPrice) telego.ShippingOption {
	return telego.ShippingOption{
		ID:     id,
		Title:  title,
		Prices: prices,
	}
}

// ReactionEmoji creates [telego.ReactionTypeEmoji] with required fields
func ReactionEmoji(emoji string) *telego.ReactionTypeEmoji {
	return &telego.ReactionTypeEmoji{
		Type:  telego.ReactionEmoji,
		Emoji: emoji,
	}
}

// ReactionCustomEmoji creates [telego.ReactionTypeCustomEmoji] with required fields
func ReactionCustomEmoji(customEmojiID string) *telego.ReactionTypeCustomEmoji {
	return &telego.ReactionTypeCustomEmoji{
		Type:          telego.ReactionCustomEmoji,
		CustomEmojiID: customEmojiID,
	}
}

// ReactionPaid creates [telego.ReactionTypePaid] with required fields
func ReactionPaid() *telego.ReactionTypePaid {
	return &telego.ReactionTypePaid{
		Type: telego.ReactionPaid,
	}
}

// RichBlockParagraph creates [telego.InputRichBlockParagraph] with required fields
func RichBlockParagraph(text telego.RichText) *telego.InputRichBlockParagraph {
	return &telego.InputRichBlockParagraph{
		Type: telego.BlockTypeParagraph,
		Text: text,
	}
}

// RichBlockSectionHeading creates [telego.InputRichBlockSectionHeading] with required fields
func RichBlockSectionHeading(text telego.RichText, size int) *telego.InputRichBlockSectionHeading {
	return &telego.InputRichBlockSectionHeading{
		Type: telego.BlockTypeSectionHeading,
		Text: text,
		Size: size,
	}
}

// RichBlockPreformatted creates [telego.InputRichBlockPreformatted] with required fields
func RichBlockPreformatted(text telego.RichText) *telego.InputRichBlockPreformatted {
	return &telego.InputRichBlockPreformatted{
		Type: telego.BlockTypePreformatted,
		Text: text,
	}
}

// RichBlockFooter creates [telego.InputRichBlockFooter] with required fields
func RichBlockFooter(text telego.RichText) *telego.InputRichBlockFooter {
	return &telego.InputRichBlockFooter{
		Type: telego.BlockTypeFooter,
		Text: text,
	}
}

// RichBlockDivider creates [telego.InputRichBlockDivider] with required fields
func RichBlockDivider() *telego.InputRichBlockDivider {
	return &telego.InputRichBlockDivider{
		Type: telego.BlockTypeDivider,
	}
}

// RichBlockMathematicalExpression creates [telego.InputRichBlockMathematicalExpression] with required fields
func RichBlockMathematicalExpression(expression string) *telego.InputRichBlockMathematicalExpression {
	return &telego.InputRichBlockMathematicalExpression{
		Type:       telego.BlockTypeMathematicalExpression,
		Expression: expression,
	}
}

// RichBlockAnchor creates [telego.InputRichBlockAnchor] with required fields
func RichBlockAnchor(name string) *telego.InputRichBlockAnchor {
	return &telego.InputRichBlockAnchor{
		Type: telego.BlockTypeAnchor,
		Name: name,
	}
}

// RichBlockList creates [telego.InputRichBlockList] with required fields
func RichBlockList(items ...telego.InputRichBlockListItem) *telego.InputRichBlockList {
	return &telego.InputRichBlockList{
		Type:  telego.BlockTypeList,
		Items: items,
	}
}

// RichBlockBlockQuotation creates [telego.InputRichBlockBlockQuotation] with required fields
func RichBlockBlockQuotation(blocks ...telego.InputRichBlock) *telego.InputRichBlockBlockQuotation {
	return &telego.InputRichBlockBlockQuotation{
		Type:   telego.BlockTypeBlockQuotation,
		Blocks: blocks,
	}
}

// RichBlockPullQuotation creates [telego.InputRichBlockPullQuotation] with required fields
func RichBlockPullQuotation(text telego.RichText) *telego.InputRichBlockPullQuotation {
	return &telego.InputRichBlockPullQuotation{
		Type: telego.BlockTypePullQuotation,
		Text: text,
	}
}

// RichBlockCollage creates [telego.InputRichBlockCollage] with required fields
func RichBlockCollage(blocks ...telego.InputRichBlock) *telego.InputRichBlockCollage {
	return &telego.InputRichBlockCollage{
		Type:   telego.BlockTypeCollage,
		Blocks: blocks,
	}
}

// RichBlockSlideshow creates [telego.InputRichBlockSlideshow] with required fields
func RichBlockSlideshow(blocks ...telego.InputRichBlock) *telego.InputRichBlockSlideshow {
	return &telego.InputRichBlockSlideshow{
		Type:   telego.BlockTypeSlideshow,
		Blocks: blocks,
	}
}

// RichBlockTable creates [telego.InputRichBlockTable] with required fields
func RichBlockTable(cells ...[]telego.RichBlockTableCell) *telego.InputRichBlockTable {
	return &telego.InputRichBlockTable{
		Type:  telego.BlockTypeTable,
		Cells: cells,
	}
}

// RichBlockTableRow creates a slice of telego.RichBlockTableCell
func RichBlockTableRow(cells ...telego.RichBlockTableCell) []telego.RichBlockTableCell {
	return cells
}

// RichBlockTableGrid creates a [telego.InputRichBlockTable] from grid of cells
func RichBlockTableGrid(cells [][]telego.RichBlockTableCell) *telego.InputRichBlockTable {
	return &telego.InputRichBlockTable{
		Type:  telego.BlockTypeTable,
		Cells: cells,
	}
}

// RichBlockTableCols creates a grid of cells containing specified number of columns
func RichBlockTableCols(cols int, cells ...telego.RichBlockTableCell) [][]telego.RichBlockTableCell {
	if cols <= 0 {
		return nil
	}

	col := 0
	row := 0

	rows := len(cells) / cols
	if len(cells)%cols != 0 {
		rows++
	}

	grid := make([][]telego.RichBlockTableCell, 0, rows)
	for i := range cells {
		if col >= cols {
			col = 0
			row++
		}
		if col == 0 {
			grid = append(grid, make([]telego.RichBlockTableCell, 0, cols))
		}
		grid[row] = append(grid[row], cells[i])
		col++
	}

	return grid
}

// RichBlockTableRows creates a grid of cells containing specified number of rows
func RichBlockTableRows(rows int, cells ...telego.RichBlockTableCell) [][]telego.RichBlockTableCell {
	if rows <= 0 {
		return nil
	}

	col := 0
	row := 0

	cols := len(cells) / rows
	if len(cells)%rows != 0 {
		cols++
	}

	grid := make([][]telego.RichBlockTableCell, 0, rows)
	for i := range cells {
		if col >= cols {
			col = 0
			row++
		}
		if col == 0 {
			grid = append(grid, make([]telego.RichBlockTableCell, 0, cols))
		}
		grid[row] = append(grid[row], cells[i])
		col++
	}

	return grid
}

// RichBlockTableCell creates [telego.RichBlockTableCell] with required fields
func RichBlockTableCell(text telego.RichText) telego.RichBlockTableCell {
	return telego.RichBlockTableCell{
		Text: text,
	}
}

// RichBlockDetails creates [telego.InputRichBlockDetails] with required fields
func RichBlockDetails(summary telego.RichText, blocks ...telego.InputRichBlock) *telego.InputRichBlockDetails {
	return &telego.InputRichBlockDetails{
		Type:    telego.BlockTypeDetails,
		Summary: summary,
		Blocks:  blocks,
	}
}

// RichBlockMap creates [telego.InputRichBlockMap] with required fields
func RichBlockMap(location telego.Location, zoom, width, height int) *telego.InputRichBlockMap {
	return &telego.InputRichBlockMap{
		Type:     telego.BlockTypeMap,
		Location: location,
		Zoom:     zoom,
		Width:    width,
		Height:   height,
	}
}

// RichBlockAnimation creates [telego.InputRichBlockAnimation] with required fields
func RichBlockAnimation(animation telego.InputMediaAnimation) *telego.InputRichBlockAnimation {
	return &telego.InputRichBlockAnimation{
		Type:      telego.BlockTypeAnimation,
		Animation: animation,
	}
}

// RichBlockAudio creates [telego.InputRichBlockAudio] with required fields
func RichBlockAudio(audio telego.InputMediaAudio) *telego.InputRichBlockAudio {
	return &telego.InputRichBlockAudio{
		Type:  telego.BlockTypeAudio,
		Audio: audio,
	}
}

// RichBlockPhoto creates [telego.InputRichBlockPhoto] with required fields
func RichBlockPhoto(photo telego.InputMediaPhoto) *telego.InputRichBlockPhoto {
	return &telego.InputRichBlockPhoto{
		Type:  telego.BlockTypePhoto,
		Photo: photo,
	}
}

// RichBlockVideo creates [telego.InputRichBlockVideo] with required fields
func RichBlockVideo(video telego.InputMediaVideo) *telego.InputRichBlockVideo {
	return &telego.InputRichBlockVideo{
		Type:  telego.BlockTypeVideo,
		Video: video,
	}
}

// RichBlockVoiceNote creates [telego.InputRichBlockVoiceNote] with required fields
func RichBlockVoiceNote(voiceNote telego.InputMediaVoiceNote) *telego.InputRichBlockVoiceNote {
	return &telego.InputRichBlockVoiceNote{
		Type:      telego.BlockTypeVoiceNote,
		VoiceNote: voiceNote,
	}
}

// RichBlockThinking creates [telego.InputRichBlockThinking] with required fields
func RichBlockThinking(text telego.RichText) *telego.InputRichBlockThinking {
	return &telego.InputRichBlockThinking{
		Type: telego.BlockTypeThinking,
		Text: text,
	}
}

// RichBlockCaption creates [telego.RichBlockCaption] with required fields
func RichBlockCaption(text telego.RichText) *telego.RichBlockCaption {
	return &telego.RichBlockCaption{
		Text: text,
	}
}

// RichBlockListItem creates [telego.InputRichBlockListItem] with required fields
func RichBlockListItem(blocks ...telego.InputRichBlock) telego.InputRichBlockListItem {
	return telego.InputRichBlockListItem{
		Blocks: blocks,
	}
}

// RichTextPlain creates [telego.RichTextPlain] from text
func RichTextPlain(text string) *telego.RichTextPlain {
	return telego.ToPtr(telego.RichTextPlain(text))
}

// RichTextList creates [telego.RichTextList] from the list of texts
func RichTextList(texts ...telego.RichText) *telego.RichTextList {
	return telego.ToPtr(telego.RichTextList(texts))
}

// RichTextBold creates [telego.RichTextBold] with required fields
func RichTextBold(text telego.RichText) *telego.RichTextBold {
	return &telego.RichTextBold{
		Type: telego.TextTypeBold,
		Text: text,
	}
}

// RichTextItalic creates [telego.RichTextItalic] with required fields
func RichTextItalic(text telego.RichText) *telego.RichTextItalic {
	return &telego.RichTextItalic{
		Type: telego.TextTypeItalic,
		Text: text,
	}
}

// RichTextUnderline creates [telego.RichTextUnderline] with required fields
func RichTextUnderline(text telego.RichText) *telego.RichTextUnderline {
	return &telego.RichTextUnderline{
		Type: telego.TextTypeUnderline,
		Text: text,
	}
}

// RichTextStrikethrough creates [telego.RichTextStrikethrough] with required fields
func RichTextStrikethrough(text telego.RichText) *telego.RichTextStrikethrough {
	return &telego.RichTextStrikethrough{
		Type: telego.TextTypeStrikethrough,
		Text: text,
	}
}

// RichTextSpoiler creates [telego.RichTextSpoiler] with required fields
func RichTextSpoiler(text telego.RichText) *telego.RichTextSpoiler {
	return &telego.RichTextSpoiler{
		Type: telego.TextTypeSpoiler,
		Text: text,
	}
}

// RichTextDateTime creates [telego.RichTextDateTime] with required fields
func RichTextDateTime(text telego.RichText, unixTime int64, dateTimeFormat string) *telego.RichTextDateTime {
	return &telego.RichTextDateTime{
		Type:           telego.TextTypeDateTime,
		Text:           text,
		UnixTime:       unixTime,
		DateTimeFormat: dateTimeFormat,
	}
}

// RichTextTextMention creates [telego.RichTextTextMention] with required fields
func RichTextTextMention(text telego.RichText, user telego.User) *telego.RichTextTextMention {
	return &telego.RichTextTextMention{
		Type: telego.TextTypeTextMention,
		Text: text,
		User: user,
	}
}

// RichTextSubscript creates [telego.RichTextSubscript] with required fields
func RichTextSubscript(text telego.RichText) *telego.RichTextSubscript {
	return &telego.RichTextSubscript{
		Type: telego.TextTypeSubscript,
		Text: text,
	}
}

// RichTextSuperscript creates [telego.RichTextSuperscript] with required fields
func RichTextSuperscript(text telego.RichText) *telego.RichTextSuperscript {
	return &telego.RichTextSuperscript{
		Type: telego.TextTypeSuperscript,
		Text: text,
	}
}

// RichTextMarked creates [telego.RichTextMarked] with required fields
func RichTextMarked(text telego.RichText) *telego.RichTextMarked {
	return &telego.RichTextMarked{
		Type: telego.TextTypeMarked,
		Text: text,
	}
}

// RichTextCode creates [telego.RichTextCode] with required fields
func RichTextCode(text telego.RichText) *telego.RichTextCode {
	return &telego.RichTextCode{
		Type: telego.TextTypeCode,
		Text: text,
	}
}

// RichTextCustomEmoji creates [telego.RichTextCustomEmoji] with required fields
func RichTextCustomEmoji(customEmojiID, alternativeText string) *telego.RichTextCustomEmoji {
	return &telego.RichTextCustomEmoji{
		Type:            telego.TextTypeCustomEmoji,
		CustomEmojiID:   customEmojiID,
		AlternativeText: alternativeText,
	}
}

// RichTextMathematicalExpression creates [telego.RichTextMathematicalExpression] with required fields
func RichTextMathematicalExpression(expression string) *telego.RichTextMathematicalExpression {
	return &telego.RichTextMathematicalExpression{
		Type:       telego.TextTypeMathematicalExpression,
		Expression: expression,
	}
}

// RichTextURL creates [telego.RichTextURL] with required fields
func RichTextURL(text telego.RichText, url string) *telego.RichTextURL {
	return &telego.RichTextURL{
		Type: telego.TextTypeURL,
		Text: text,
		URL:  url,
	}
}

// RichTextEmailAddress creates [telego.RichTextEmailAddress] with required fields
func RichTextEmailAddress(text telego.RichText, email string) *telego.RichTextEmailAddress {
	return &telego.RichTextEmailAddress{
		Type:         telego.TextTypeEmailAddress,
		Text:         text,
		EmailAddress: email,
	}
}

// RichTextPhoneNumber creates [telego.RichTextPhoneNumber] with required fields
func RichTextPhoneNumber(text telego.RichText, phoneNumber string) *telego.RichTextPhoneNumber {
	return &telego.RichTextPhoneNumber{
		Type:        telego.TextTypePhoneNumber,
		Text:        text,
		PhoneNumber: phoneNumber,
	}
}

// RichTextBankCardNumber creates [telego.RichTextBankCardNumber] with required fields
func RichTextBankCardNumber(text telego.RichText, bankCardNumber string) *telego.RichTextBankCardNumber {
	return &telego.RichTextBankCardNumber{
		Type:           telego.TextTypeBankCardNumber,
		Text:           text,
		BankCardNumber: bankCardNumber,
	}
}

// RichTextMention creates [telego.RichTextMention] with required fields
func RichTextMention(text telego.RichText, username string) *telego.RichTextMention {
	return &telego.RichTextMention{
		Type:     telego.TextTypeMention,
		Text:     text,
		Username: username,
	}
}

// RichTextHashtag creates [telego.RichTextHashtag] with required fields
func RichTextHashtag(text telego.RichText, hashtag string) *telego.RichTextHashtag {
	return &telego.RichTextHashtag{
		Type:    telego.TextTypeHashtag,
		Text:    text,
		Hashtag: hashtag,
	}
}

// RichTextCashtag creates [telego.RichTextCashtag] with required fields
func RichTextCashtag(text telego.RichText, cashtag string) *telego.RichTextCashtag {
	return &telego.RichTextCashtag{
		Type:    telego.TextTypeCashtag,
		Text:    text,
		Cashtag: cashtag,
	}
}

// RichTextBotCommand creates [telego.RichTextBotCommand] with required fields
func RichTextBotCommand(text telego.RichText, botCommand string) *telego.RichTextBotCommand {
	return &telego.RichTextBotCommand{
		Type:       telego.TextTypeBotCommand,
		Text:       text,
		BotCommand: botCommand,
	}
}

// RichTextAnchor creates [telego.RichTextAnchor] with required fields
func RichTextAnchor(name string) *telego.RichTextAnchor {
	return &telego.RichTextAnchor{
		Type: telego.TextTypeAnchor,
		Name: name,
	}
}

// RichTextAnchorLink creates [telego.RichTextAnchorLink] with required fields
func RichTextAnchorLink(text telego.RichText, anchorName string) *telego.RichTextAnchorLink {
	return &telego.RichTextAnchorLink{
		Type:       telego.TextTypeAnchorLink,
		Text:       text,
		AnchorName: anchorName,
	}
}

// RichTextReference creates [telego.RichTextReference] with required fields
func RichTextReference(text telego.RichText, name string) *telego.RichTextReference {
	return &telego.RichTextReference{
		Type: telego.TextTypeReference,
		Text: text,
		Name: name,
	}
}

// RichTextReferenceLink creates [telego.RichTextReferenceLink] with required fields
func RichTextReferenceLink(text telego.RichText, referenceName string) *telego.RichTextReferenceLink {
	return &telego.RichTextReferenceLink{
		Type:          telego.TextTypeReferenceLink,
		Text:          text,
		ReferenceName: referenceName,
	}
}

// RichMessage creates [telego.InputRichMessage] with required fields from blocks
func RichMessage(blocks ...telego.InputRichBlock) telego.InputRichMessage {
	return telego.InputRichMessage{
		Blocks: blocks,
	}
}

// RichMessageHTML creates [telego.InputRichMessage] with required fields from HTML
func RichMessageHTML(html string) telego.InputRichMessage {
	return telego.InputRichMessage{
		HTML: html,
	}
}

// RichMessageMarkdown creates [telego.InputRichMessage] with required fields from markdown
func RichMessageMarkdown(markdown string) telego.InputRichMessage {
	return telego.InputRichMessage{
		Markdown: markdown,
	}
}
