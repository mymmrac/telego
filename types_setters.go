package telego

// WithKeyboard adds keyboard parameter
func (r *ReplyKeyboardMarkup) WithKeyboard(keyboard ...[]KeyboardButton) *ReplyKeyboardMarkup {
	r.Keyboard = keyboard
	return r
}

// WithIsPersistent adds is persistent parameter
func (r *ReplyKeyboardMarkup) WithIsPersistent() *ReplyKeyboardMarkup {
	r.IsPersistent = true
	return r
}

// WithResizeKeyboard adds resize keyboard parameter
func (r *ReplyKeyboardMarkup) WithResizeKeyboard() *ReplyKeyboardMarkup {
	r.ResizeKeyboard = true
	return r
}

// WithOneTimeKeyboard adds one time keyboard parameter
func (r *ReplyKeyboardMarkup) WithOneTimeKeyboard() *ReplyKeyboardMarkup {
	r.OneTimeKeyboard = true
	return r
}

// WithInputFieldPlaceholder adds input field placeholder parameter
func (r *ReplyKeyboardMarkup) WithInputFieldPlaceholder(inputFieldPlaceholder string) *ReplyKeyboardMarkup {
	r.InputFieldPlaceholder = inputFieldPlaceholder
	return r
}

// WithSelective adds selective parameter
func (r *ReplyKeyboardMarkup) WithSelective() *ReplyKeyboardMarkup {
	r.Selective = true
	return r
}

// WithText adds text parameter
func (k KeyboardButton) WithText(text string) KeyboardButton {
	k.Text = text
	return k
}

// WithRequestUser adds request user parameter
func (k KeyboardButton) WithRequestUser(requestUser *KeyboardButtonRequestUser) KeyboardButton {
	k.RequestUser = requestUser
	return k
}

// WithRequestChat adds request chat parameter
func (k KeyboardButton) WithRequestChat(requestChat *KeyboardButtonRequestChat) KeyboardButton {
	k.RequestChat = requestChat
	return k
}

// WithRequestContact adds request contact parameter
func (k KeyboardButton) WithRequestContact() KeyboardButton {
	k.RequestContact = true
	return k
}

// WithRequestLocation adds request location parameter
func (k KeyboardButton) WithRequestLocation() KeyboardButton {
	k.RequestLocation = true
	return k
}

// WithRequestPoll adds request poll parameter
func (k KeyboardButton) WithRequestPoll(requestPoll *KeyboardButtonPollType) KeyboardButton {
	k.RequestPoll = requestPoll
	return k
}

// WithWebApp adds web app parameter
func (k KeyboardButton) WithWebApp(webApp *WebAppInfo) KeyboardButton {
	k.WebApp = webApp
	return k
}

// WithUserIsBot adds user is bot parameter
func (k *KeyboardButtonRequestUser) WithUserIsBot(userIsBot bool) *KeyboardButtonRequestUser {
	k.UserIsBot = ToPtr(userIsBot)
	return k
}

// WithUserIsPremium adds user is premium parameter
func (k *KeyboardButtonRequestUser) WithUserIsPremium(userIsPremium bool) *KeyboardButtonRequestUser {
	k.UserIsPremium = ToPtr(userIsPremium)
	return k
}

// WithChatIsChannel adds chat is channel parameter
func (k *KeyboardButtonRequestChat) WithChatIsChannel() *KeyboardButtonRequestChat {
	k.ChatIsChannel = true
	return k
}

// WithChatIsForum adds chat is forum parameter
func (k *KeyboardButtonRequestChat) WithChatIsForum(chatIsForum bool) *KeyboardButtonRequestChat {
	k.ChatIsForum = ToPtr(chatIsForum)
	return k
}

// WithChatHasUsername adds chat has username parameter
func (k *KeyboardButtonRequestChat) WithChatHasUsername(chatHasUsername bool) *KeyboardButtonRequestChat {
	k.ChatHasUsername = ToPtr(chatHasUsername)
	return k
}

// WithChatIsCreated adds chat is created parameter
func (k *KeyboardButtonRequestChat) WithChatIsCreated(chatIsCreated bool) *KeyboardButtonRequestChat {
	k.ChatIsCreated = ToPtr(chatIsCreated)
	return k
}

// WithUserAdministratorRights adds user administrator rights parameter
func (k *KeyboardButtonRequestChat) WithUserAdministratorRights(userAdministratorRights *ChatAdministratorRights,
) *KeyboardButtonRequestChat {
	k.UserAdministratorRights = userAdministratorRights
	return k
}

// WithBotAdministratorRights adds bot administrator rights parameter
func (k *KeyboardButtonRequestChat) WithBotAdministratorRights(botAdministratorRights *ChatAdministratorRights,
) *KeyboardButtonRequestChat {
	k.BotAdministratorRights = botAdministratorRights
	return k
}

// WithBotIsMember adds bot is member parameter
func (k *KeyboardButtonRequestChat) WithBotIsMember(botIsMember bool) *KeyboardButtonRequestChat {
	k.BotIsMember = ToPtr(botIsMember)
	return k
}

// WithRemoveKeyboard adds remove keyboard parameter
func (r *ReplyKeyboardRemove) WithRemoveKeyboard() *ReplyKeyboardRemove {
	r.RemoveKeyboard = true
	return r
}

// WithSelective adds selective parameter
func (r *ReplyKeyboardRemove) WithSelective() *ReplyKeyboardRemove {
	r.Selective = true
	return r
}

// WithInlineKeyboard adds inline keyboard parameter
func (i *InlineKeyboardMarkup) WithInlineKeyboard(inlineKeyboard ...[]InlineKeyboardButton) *InlineKeyboardMarkup {
	i.InlineKeyboard = inlineKeyboard
	return i
}

// WithText adds text parameter
func (i InlineKeyboardButton) WithText(text string) InlineKeyboardButton {
	i.Text = text
	return i
}

// WithURL adds URL parameter
func (i InlineKeyboardButton) WithURL(url string) InlineKeyboardButton {
	i.URL = url
	return i
}

// WithCallbackData adds callback data parameter
func (i InlineKeyboardButton) WithCallbackData(callbackData string) InlineKeyboardButton {
	i.CallbackData = callbackData
	return i
}

// WithWebApp adds web app parameter
func (i InlineKeyboardButton) WithWebApp(webApp *WebAppInfo) InlineKeyboardButton {
	i.WebApp = webApp
	return i
}

// WithLoginURL adds login URL parameter
func (i InlineKeyboardButton) WithLoginURL(loginURL *LoginURL) InlineKeyboardButton {
	i.LoginURL = loginURL
	return i
}

// WithSwitchInlineQuery adds switch inline query parameter
func (i InlineKeyboardButton) WithSwitchInlineQuery(switchInlineQuery string) InlineKeyboardButton {
	i.SwitchInlineQuery = ToPtr(switchInlineQuery)
	return i
}

// WithSwitchInlineQueryCurrentChat adds switch inline query current chat parameter
func (i InlineKeyboardButton) WithSwitchInlineQueryCurrentChat(switchInlineQueryCurrentChat string,
) InlineKeyboardButton {
	i.SwitchInlineQueryCurrentChat = ToPtr(switchInlineQueryCurrentChat)
	return i
}

// WithSwitchInlineQueryChosenChat adds switch inline query chosen chat parameter
func (i InlineKeyboardButton) WithSwitchInlineQueryChosenChat(switchInlineQueryChosenChat *SwitchInlineQueryChosenChat,
) InlineKeyboardButton {
	i.SwitchInlineQueryChosenChat = switchInlineQueryChosenChat
	return i
}

// WithCallbackGame adds callback game parameter
func (i InlineKeyboardButton) WithCallbackGame(callbackGame *CallbackGame) InlineKeyboardButton {
	i.CallbackGame = callbackGame
	return i
}

// WithPay adds pay parameter
func (i InlineKeyboardButton) WithPay() InlineKeyboardButton {
	i.Pay = true
	return i
}

// WithForceReply adds force reply parameter
func (f *ForceReply) WithForceReply() *ForceReply {
	f.ForceReply = true
	return f
}

// WithInputFieldPlaceholder adds input field placeholder parameter
func (f *ForceReply) WithInputFieldPlaceholder(inputFieldPlaceholder string) *ForceReply {
	f.InputFieldPlaceholder = inputFieldPlaceholder
	return f
}

// WithSelective adds selective parameter
func (f *ForceReply) WithSelective() *ForceReply {
	f.Selective = true
	return f
}

// WithText adds text parameter
func (m *MenuButtonWebApp) WithText(text string) *MenuButtonWebApp {
	m.Text = text
	return m
}

// WithWebApp adds web app parameter
func (m *MenuButtonWebApp) WithWebApp(webApp WebAppInfo) *MenuButtonWebApp {
	m.WebApp = webApp
	return m
}

// WithMedia adds media parameter
func (i *InputMediaPhoto) WithMedia(media InputFile) *InputMediaPhoto {
	i.Media = media
	return i
}

// WithCaption adds caption parameter
func (i *InputMediaPhoto) WithCaption(caption string) *InputMediaPhoto {
	i.Caption = caption
	return i
}

// WithParseMode adds parse mode parameter
func (i *InputMediaPhoto) WithParseMode(parseMode string) *InputMediaPhoto {
	i.ParseMode = parseMode
	return i
}

// WithCaptionEntities adds caption entities parameter
func (i *InputMediaPhoto) WithCaptionEntities(captionEntities ...MessageEntity) *InputMediaPhoto {
	i.CaptionEntities = captionEntities
	return i
}

// WithHasSpoiler adds has spoiler parameter
func (i *InputMediaPhoto) WithHasSpoiler() *InputMediaPhoto {
	i.HasSpoiler = true
	return i
}

// WithMedia adds media parameter
func (i *InputMediaVideo) WithMedia(media InputFile) *InputMediaVideo {
	i.Media = media
	return i
}

// WithThumbnail adds thumbnail parameter
func (i *InputMediaVideo) WithThumbnail(thumbnail *InputFile) *InputMediaVideo {
	i.Thumbnail = thumbnail
	return i
}

// WithCaption adds caption parameter
func (i *InputMediaVideo) WithCaption(caption string) *InputMediaVideo {
	i.Caption = caption
	return i
}

// WithParseMode adds parse mode parameter
func (i *InputMediaVideo) WithParseMode(parseMode string) *InputMediaVideo {
	i.ParseMode = parseMode
	return i
}

// WithCaptionEntities adds caption entities parameter
func (i *InputMediaVideo) WithCaptionEntities(captionEntities ...MessageEntity) *InputMediaVideo {
	i.CaptionEntities = captionEntities
	return i
}

// WithWidth adds width parameter
func (i *InputMediaVideo) WithWidth(width int) *InputMediaVideo {
	i.Width = width
	return i
}

// WithHeight adds height parameter
func (i *InputMediaVideo) WithHeight(height int) *InputMediaVideo {
	i.Height = height
	return i
}

// WithDuration adds duration parameter
func (i *InputMediaVideo) WithDuration(duration int) *InputMediaVideo {
	i.Duration = duration
	return i
}

// WithSupportsStreaming adds supports streaming parameter
func (i *InputMediaVideo) WithSupportsStreaming() *InputMediaVideo {
	i.SupportsStreaming = true
	return i
}

// WithHasSpoiler adds has spoiler parameter
func (i *InputMediaVideo) WithHasSpoiler() *InputMediaVideo {
	i.HasSpoiler = true
	return i
}

// WithMedia adds media parameter
func (i *InputMediaAnimation) WithMedia(media InputFile) *InputMediaAnimation {
	i.Media = media
	return i
}

// WithThumbnail adds thumbnail parameter
func (i *InputMediaAnimation) WithThumbnail(thumbnail *InputFile) *InputMediaAnimation {
	i.Thumbnail = thumbnail
	return i
}

// WithCaption adds caption parameter
func (i *InputMediaAnimation) WithCaption(caption string) *InputMediaAnimation {
	i.Caption = caption
	return i
}

// WithParseMode adds parse mode parameter
func (i *InputMediaAnimation) WithParseMode(parseMode string) *InputMediaAnimation {
	i.ParseMode = parseMode
	return i
}

// WithCaptionEntities adds caption entities parameter
func (i *InputMediaAnimation) WithCaptionEntities(captionEntities ...MessageEntity) *InputMediaAnimation {
	i.CaptionEntities = captionEntities
	return i
}

// WithWidth adds width parameter
func (i *InputMediaAnimation) WithWidth(width int) *InputMediaAnimation {
	i.Width = width
	return i
}

// WithHeight adds height parameter
func (i *InputMediaAnimation) WithHeight(height int) *InputMediaAnimation {
	i.Height = height
	return i
}

// WithDuration adds duration parameter
func (i *InputMediaAnimation) WithDuration(duration int) *InputMediaAnimation {
	i.Duration = duration
	return i
}

// WithHasSpoiler adds has spoiler parameter
func (i *InputMediaAnimation) WithHasSpoiler() *InputMediaAnimation {
	i.HasSpoiler = true
	return i
}

// WithMedia adds media parameter
func (i *InputMediaAudio) WithMedia(media InputFile) *InputMediaAudio {
	i.Media = media
	return i
}

// WithThumbnail adds thumbnail parameter
func (i *InputMediaAudio) WithThumbnail(thumbnail *InputFile) *InputMediaAudio {
	i.Thumbnail = thumbnail
	return i
}

// WithCaption adds caption parameter
func (i *InputMediaAudio) WithCaption(caption string) *InputMediaAudio {
	i.Caption = caption
	return i
}

// WithParseMode adds parse mode parameter
func (i *InputMediaAudio) WithParseMode(parseMode string) *InputMediaAudio {
	i.ParseMode = parseMode
	return i
}

// WithCaptionEntities adds caption entities parameter
func (i *InputMediaAudio) WithCaptionEntities(captionEntities ...MessageEntity) *InputMediaAudio {
	i.CaptionEntities = captionEntities
	return i
}

// WithDuration adds duration parameter
func (i *InputMediaAudio) WithDuration(duration int) *InputMediaAudio {
	i.Duration = duration
	return i
}

// WithPerformer adds performer parameter
func (i *InputMediaAudio) WithPerformer(performer string) *InputMediaAudio {
	i.Performer = performer
	return i
}

// WithTitle adds title parameter
func (i *InputMediaAudio) WithTitle(title string) *InputMediaAudio {
	i.Title = title
	return i
}

// WithMedia adds media parameter
func (i *InputMediaDocument) WithMedia(media InputFile) *InputMediaDocument {
	i.Media = media
	return i
}

// WithThumbnail adds thumbnail parameter
func (i *InputMediaDocument) WithThumbnail(thumbnail *InputFile) *InputMediaDocument {
	i.Thumbnail = thumbnail
	return i
}

// WithCaption adds caption parameter
func (i *InputMediaDocument) WithCaption(caption string) *InputMediaDocument {
	i.Caption = caption
	return i
}

// WithParseMode adds parse mode parameter
func (i *InputMediaDocument) WithParseMode(parseMode string) *InputMediaDocument {
	i.ParseMode = parseMode
	return i
}

// WithCaptionEntities adds caption entities parameter
func (i *InputMediaDocument) WithCaptionEntities(captionEntities ...MessageEntity) *InputMediaDocument {
	i.CaptionEntities = captionEntities
	return i
}

// WithDisableContentTypeDetection adds disable content type detection parameter
func (i *InputMediaDocument) WithDisableContentTypeDetection() *InputMediaDocument {
	i.DisableContentTypeDetection = true
	return i
}

// WithSticker adds sticker parameter
func (i *InputSticker) WithSticker(sticker InputFile) *InputSticker {
	i.Sticker = sticker
	return i
}

// WithEmojiList adds emoji list parameter
func (i *InputSticker) WithEmojiList(emojiList ...string) *InputSticker {
	i.EmojiList = emojiList
	return i
}

// WithMaskPosition adds mask position parameter
func (i *InputSticker) WithMaskPosition(maskPosition *MaskPosition) *InputSticker {
	i.MaskPosition = maskPosition
	return i
}

// WithKeywords adds keywords parameter
func (i *InputSticker) WithKeywords(keywords ...string) *InputSticker {
	i.Keywords = keywords
	return i
}

// WithID adds ID parameter
func (i *InlineQueryResultArticle) WithID(iD string) *InlineQueryResultArticle {
	i.ID = iD
	return i
}

// WithTitle adds title parameter
func (i *InlineQueryResultArticle) WithTitle(title string) *InlineQueryResultArticle {
	i.Title = title
	return i
}

// WithInputMessageContent adds input message content parameter
func (i *InlineQueryResultArticle) WithInputMessageContent(inputMessageContent InputMessageContent,
) *InlineQueryResultArticle {
	i.InputMessageContent = inputMessageContent
	return i
}

// WithReplyMarkup adds reply markup parameter
func (i *InlineQueryResultArticle) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *InlineQueryResultArticle {
	i.ReplyMarkup = replyMarkup
	return i
}

// WithURL adds URL parameter
func (i *InlineQueryResultArticle) WithURL(url string) *InlineQueryResultArticle {
	i.URL = url
	return i
}

// WithHideURL adds hide URL parameter
func (i *InlineQueryResultArticle) WithHideURL() *InlineQueryResultArticle {
	i.HideURL = true
	return i
}

// WithDescription adds description parameter
func (i *InlineQueryResultArticle) WithDescription(description string) *InlineQueryResultArticle {
	i.Description = description
	return i
}

// WithThumbnailURL adds thumbnail URL parameter
func (i *InlineQueryResultArticle) WithThumbnailURL(thumbnailURL string) *InlineQueryResultArticle {
	i.ThumbnailURL = thumbnailURL
	return i
}

// WithThumbnailWidth adds thumbnail width parameter
func (i *InlineQueryResultArticle) WithThumbnailWidth(thumbnailWidth int) *InlineQueryResultArticle {
	i.ThumbnailWidth = thumbnailWidth
	return i
}

// WithThumbnailHeight adds thumbnail height parameter
func (i *InlineQueryResultArticle) WithThumbnailHeight(thumbnailHeight int) *InlineQueryResultArticle {
	i.ThumbnailHeight = thumbnailHeight
	return i
}

// WithID adds ID parameter
func (i *InlineQueryResultPhoto) WithID(iD string) *InlineQueryResultPhoto {
	i.ID = iD
	return i
}

// WithPhotoURL adds photo URL parameter
func (i *InlineQueryResultPhoto) WithPhotoURL(photoURL string) *InlineQueryResultPhoto {
	i.PhotoURL = photoURL
	return i
}

// WithThumbnailURL adds thumbnail URL parameter
func (i *InlineQueryResultPhoto) WithThumbnailURL(thumbnailURL string) *InlineQueryResultPhoto {
	i.ThumbnailURL = thumbnailURL
	return i
}

// WithPhotoWidth adds photo width parameter
func (i *InlineQueryResultPhoto) WithPhotoWidth(photoWidth int) *InlineQueryResultPhoto {
	i.PhotoWidth = photoWidth
	return i
}

// WithPhotoHeight adds photo height parameter
func (i *InlineQueryResultPhoto) WithPhotoHeight(photoHeight int) *InlineQueryResultPhoto {
	i.PhotoHeight = photoHeight
	return i
}

// WithTitle adds title parameter
func (i *InlineQueryResultPhoto) WithTitle(title string) *InlineQueryResultPhoto {
	i.Title = title
	return i
}

// WithDescription adds description parameter
func (i *InlineQueryResultPhoto) WithDescription(description string) *InlineQueryResultPhoto {
	i.Description = description
	return i
}

// WithCaption adds caption parameter
func (i *InlineQueryResultPhoto) WithCaption(caption string) *InlineQueryResultPhoto {
	i.Caption = caption
	return i
}

// WithParseMode adds parse mode parameter
func (i *InlineQueryResultPhoto) WithParseMode(parseMode string) *InlineQueryResultPhoto {
	i.ParseMode = parseMode
	return i
}

// WithCaptionEntities adds caption entities parameter
func (i *InlineQueryResultPhoto) WithCaptionEntities(captionEntities ...MessageEntity) *InlineQueryResultPhoto {
	i.CaptionEntities = captionEntities
	return i
}

// WithReplyMarkup adds reply markup parameter
func (i *InlineQueryResultPhoto) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *InlineQueryResultPhoto {
	i.ReplyMarkup = replyMarkup
	return i
}

// WithInputMessageContent adds input message content parameter
func (i *InlineQueryResultPhoto) WithInputMessageContent(inputMessageContent InputMessageContent,
) *InlineQueryResultPhoto {
	i.InputMessageContent = inputMessageContent
	return i
}

// WithID adds ID parameter
func (i *InlineQueryResultGif) WithID(iD string) *InlineQueryResultGif {
	i.ID = iD
	return i
}

// WithGifURL adds gif URL parameter
func (i *InlineQueryResultGif) WithGifURL(gifURL string) *InlineQueryResultGif {
	i.GifURL = gifURL
	return i
}

// WithGifWidth adds gif width parameter
func (i *InlineQueryResultGif) WithGifWidth(gifWidth int) *InlineQueryResultGif {
	i.GifWidth = gifWidth
	return i
}

// WithGifHeight adds gif height parameter
func (i *InlineQueryResultGif) WithGifHeight(gifHeight int) *InlineQueryResultGif {
	i.GifHeight = gifHeight
	return i
}

// WithGifDuration adds gif duration parameter
func (i *InlineQueryResultGif) WithGifDuration(gifDuration int) *InlineQueryResultGif {
	i.GifDuration = gifDuration
	return i
}

// WithThumbnailURL adds thumbnail URL parameter
func (i *InlineQueryResultGif) WithThumbnailURL(thumbnailURL string) *InlineQueryResultGif {
	i.ThumbnailURL = thumbnailURL
	return i
}

// WithThumbnailMimeType adds thumbnail mime type parameter
func (i *InlineQueryResultGif) WithThumbnailMimeType(thumbnailMimeType string) *InlineQueryResultGif {
	i.ThumbnailMimeType = thumbnailMimeType
	return i
}

// WithTitle adds title parameter
func (i *InlineQueryResultGif) WithTitle(title string) *InlineQueryResultGif {
	i.Title = title
	return i
}

// WithCaption adds caption parameter
func (i *InlineQueryResultGif) WithCaption(caption string) *InlineQueryResultGif {
	i.Caption = caption
	return i
}

// WithParseMode adds parse mode parameter
func (i *InlineQueryResultGif) WithParseMode(parseMode string) *InlineQueryResultGif {
	i.ParseMode = parseMode
	return i
}

// WithCaptionEntities adds caption entities parameter
func (i *InlineQueryResultGif) WithCaptionEntities(captionEntities ...MessageEntity) *InlineQueryResultGif {
	i.CaptionEntities = captionEntities
	return i
}

// WithReplyMarkup adds reply markup parameter
func (i *InlineQueryResultGif) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *InlineQueryResultGif {
	i.ReplyMarkup = replyMarkup
	return i
}

// WithInputMessageContent adds input message content parameter
func (i *InlineQueryResultGif) WithInputMessageContent(inputMessageContent InputMessageContent) *InlineQueryResultGif {
	i.InputMessageContent = inputMessageContent
	return i
}

// WithID adds ID parameter
func (i *InlineQueryResultMpeg4Gif) WithID(iD string) *InlineQueryResultMpeg4Gif {
	i.ID = iD
	return i
}

// WithMpeg4URL adds mpeg4 URL parameter
func (i *InlineQueryResultMpeg4Gif) WithMpeg4URL(mpeg4URL string) *InlineQueryResultMpeg4Gif {
	i.Mpeg4URL = mpeg4URL
	return i
}

// WithMpeg4Width adds mpeg4 width parameter
func (i *InlineQueryResultMpeg4Gif) WithMpeg4Width(mpeg4Width int) *InlineQueryResultMpeg4Gif {
	i.Mpeg4Width = mpeg4Width
	return i
}

// WithMpeg4Height adds mpeg4 height parameter
func (i *InlineQueryResultMpeg4Gif) WithMpeg4Height(mpeg4Height int) *InlineQueryResultMpeg4Gif {
	i.Mpeg4Height = mpeg4Height
	return i
}

// WithMpeg4Duration adds mpeg4 duration parameter
func (i *InlineQueryResultMpeg4Gif) WithMpeg4Duration(mpeg4Duration int) *InlineQueryResultMpeg4Gif {
	i.Mpeg4Duration = mpeg4Duration
	return i
}

// WithThumbnailURL adds thumbnail URL parameter
func (i *InlineQueryResultMpeg4Gif) WithThumbnailURL(thumbnailURL string) *InlineQueryResultMpeg4Gif {
	i.ThumbnailURL = thumbnailURL
	return i
}

// WithThumbnailMimeType adds thumbnail mime type parameter
func (i *InlineQueryResultMpeg4Gif) WithThumbnailMimeType(thumbnailMimeType string) *InlineQueryResultMpeg4Gif {
	i.ThumbnailMimeType = thumbnailMimeType
	return i
}

// WithTitle adds title parameter
func (i *InlineQueryResultMpeg4Gif) WithTitle(title string) *InlineQueryResultMpeg4Gif {
	i.Title = title
	return i
}

// WithCaption adds caption parameter
func (i *InlineQueryResultMpeg4Gif) WithCaption(caption string) *InlineQueryResultMpeg4Gif {
	i.Caption = caption
	return i
}

// WithParseMode adds parse mode parameter
func (i *InlineQueryResultMpeg4Gif) WithParseMode(parseMode string) *InlineQueryResultMpeg4Gif {
	i.ParseMode = parseMode
	return i
}

// WithCaptionEntities adds caption entities parameter
func (i *InlineQueryResultMpeg4Gif) WithCaptionEntities(captionEntities ...MessageEntity) *InlineQueryResultMpeg4Gif {
	i.CaptionEntities = captionEntities
	return i
}

// WithReplyMarkup adds reply markup parameter
func (i *InlineQueryResultMpeg4Gif) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *InlineQueryResultMpeg4Gif {
	i.ReplyMarkup = replyMarkup
	return i
}

// WithInputMessageContent adds input message content parameter
func (i *InlineQueryResultMpeg4Gif) WithInputMessageContent(inputMessageContent InputMessageContent,
) *InlineQueryResultMpeg4Gif {
	i.InputMessageContent = inputMessageContent
	return i
}

// WithID adds ID parameter
func (i *InlineQueryResultVideo) WithID(iD string) *InlineQueryResultVideo {
	i.ID = iD
	return i
}

// WithVideoURL adds video URL parameter
func (i *InlineQueryResultVideo) WithVideoURL(videoURL string) *InlineQueryResultVideo {
	i.VideoURL = videoURL
	return i
}

// WithMimeType adds mime type parameter
func (i *InlineQueryResultVideo) WithMimeType(mimeType string) *InlineQueryResultVideo {
	i.MimeType = mimeType
	return i
}

// WithThumbnailURL adds thumbnail URL parameter
func (i *InlineQueryResultVideo) WithThumbnailURL(thumbnailURL string) *InlineQueryResultVideo {
	i.ThumbnailURL = thumbnailURL
	return i
}

// WithTitle adds title parameter
func (i *InlineQueryResultVideo) WithTitle(title string) *InlineQueryResultVideo {
	i.Title = title
	return i
}

// WithCaption adds caption parameter
func (i *InlineQueryResultVideo) WithCaption(caption string) *InlineQueryResultVideo {
	i.Caption = caption
	return i
}

// WithParseMode adds parse mode parameter
func (i *InlineQueryResultVideo) WithParseMode(parseMode string) *InlineQueryResultVideo {
	i.ParseMode = parseMode
	return i
}

// WithCaptionEntities adds caption entities parameter
func (i *InlineQueryResultVideo) WithCaptionEntities(captionEntities ...MessageEntity) *InlineQueryResultVideo {
	i.CaptionEntities = captionEntities
	return i
}

// WithVideoWidth adds video width parameter
func (i *InlineQueryResultVideo) WithVideoWidth(videoWidth int) *InlineQueryResultVideo {
	i.VideoWidth = videoWidth
	return i
}

// WithVideoHeight adds video height parameter
func (i *InlineQueryResultVideo) WithVideoHeight(videoHeight int) *InlineQueryResultVideo {
	i.VideoHeight = videoHeight
	return i
}

// WithVideoDuration adds video duration parameter
func (i *InlineQueryResultVideo) WithVideoDuration(videoDuration int) *InlineQueryResultVideo {
	i.VideoDuration = videoDuration
	return i
}

// WithDescription adds description parameter
func (i *InlineQueryResultVideo) WithDescription(description string) *InlineQueryResultVideo {
	i.Description = description
	return i
}

// WithReplyMarkup adds reply markup parameter
func (i *InlineQueryResultVideo) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *InlineQueryResultVideo {
	i.ReplyMarkup = replyMarkup
	return i
}

// WithInputMessageContent adds input message content parameter
func (i *InlineQueryResultVideo) WithInputMessageContent(inputMessageContent InputMessageContent,
) *InlineQueryResultVideo {
	i.InputMessageContent = inputMessageContent
	return i
}

// WithID adds ID parameter
func (i *InlineQueryResultAudio) WithID(iD string) *InlineQueryResultAudio {
	i.ID = iD
	return i
}

// WithAudioURL adds audio URL parameter
func (i *InlineQueryResultAudio) WithAudioURL(audioURL string) *InlineQueryResultAudio {
	i.AudioURL = audioURL
	return i
}

// WithTitle adds title parameter
func (i *InlineQueryResultAudio) WithTitle(title string) *InlineQueryResultAudio {
	i.Title = title
	return i
}

// WithCaption adds caption parameter
func (i *InlineQueryResultAudio) WithCaption(caption string) *InlineQueryResultAudio {
	i.Caption = caption
	return i
}

// WithParseMode adds parse mode parameter
func (i *InlineQueryResultAudio) WithParseMode(parseMode string) *InlineQueryResultAudio {
	i.ParseMode = parseMode
	return i
}

// WithCaptionEntities adds caption entities parameter
func (i *InlineQueryResultAudio) WithCaptionEntities(captionEntities ...MessageEntity) *InlineQueryResultAudio {
	i.CaptionEntities = captionEntities
	return i
}

// WithPerformer adds performer parameter
func (i *InlineQueryResultAudio) WithPerformer(performer string) *InlineQueryResultAudio {
	i.Performer = performer
	return i
}

// WithAudioDuration adds audio duration parameter
func (i *InlineQueryResultAudio) WithAudioDuration(audioDuration int) *InlineQueryResultAudio {
	i.AudioDuration = audioDuration
	return i
}

// WithReplyMarkup adds reply markup parameter
func (i *InlineQueryResultAudio) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *InlineQueryResultAudio {
	i.ReplyMarkup = replyMarkup
	return i
}

// WithInputMessageContent adds input message content parameter
func (i *InlineQueryResultAudio) WithInputMessageContent(inputMessageContent InputMessageContent,
) *InlineQueryResultAudio {
	i.InputMessageContent = inputMessageContent
	return i
}

// WithID adds ID parameter
func (i *InlineQueryResultVoice) WithID(iD string) *InlineQueryResultVoice {
	i.ID = iD
	return i
}

// WithVoiceURL adds voice URL parameter
func (i *InlineQueryResultVoice) WithVoiceURL(voiceURL string) *InlineQueryResultVoice {
	i.VoiceURL = voiceURL
	return i
}

// WithTitle adds title parameter
func (i *InlineQueryResultVoice) WithTitle(title string) *InlineQueryResultVoice {
	i.Title = title
	return i
}

// WithCaption adds caption parameter
func (i *InlineQueryResultVoice) WithCaption(caption string) *InlineQueryResultVoice {
	i.Caption = caption
	return i
}

// WithParseMode adds parse mode parameter
func (i *InlineQueryResultVoice) WithParseMode(parseMode string) *InlineQueryResultVoice {
	i.ParseMode = parseMode
	return i
}

// WithCaptionEntities adds caption entities parameter
func (i *InlineQueryResultVoice) WithCaptionEntities(captionEntities ...MessageEntity) *InlineQueryResultVoice {
	i.CaptionEntities = captionEntities
	return i
}

// WithVoiceDuration adds voice duration parameter
func (i *InlineQueryResultVoice) WithVoiceDuration(voiceDuration int) *InlineQueryResultVoice {
	i.VoiceDuration = voiceDuration
	return i
}

// WithReplyMarkup adds reply markup parameter
func (i *InlineQueryResultVoice) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *InlineQueryResultVoice {
	i.ReplyMarkup = replyMarkup
	return i
}

// WithInputMessageContent adds input message content parameter
func (i *InlineQueryResultVoice) WithInputMessageContent(inputMessageContent InputMessageContent,
) *InlineQueryResultVoice {
	i.InputMessageContent = inputMessageContent
	return i
}

// WithID adds ID parameter
func (i *InlineQueryResultDocument) WithID(iD string) *InlineQueryResultDocument {
	i.ID = iD
	return i
}

// WithTitle adds title parameter
func (i *InlineQueryResultDocument) WithTitle(title string) *InlineQueryResultDocument {
	i.Title = title
	return i
}

// WithCaption adds caption parameter
func (i *InlineQueryResultDocument) WithCaption(caption string) *InlineQueryResultDocument {
	i.Caption = caption
	return i
}

// WithParseMode adds parse mode parameter
func (i *InlineQueryResultDocument) WithParseMode(parseMode string) *InlineQueryResultDocument {
	i.ParseMode = parseMode
	return i
}

// WithCaptionEntities adds caption entities parameter
func (i *InlineQueryResultDocument) WithCaptionEntities(captionEntities ...MessageEntity) *InlineQueryResultDocument {
	i.CaptionEntities = captionEntities
	return i
}

// WithDocumentURL adds document URL parameter
func (i *InlineQueryResultDocument) WithDocumentURL(documentURL string) *InlineQueryResultDocument {
	i.DocumentURL = documentURL
	return i
}

// WithMimeType adds mime type parameter
func (i *InlineQueryResultDocument) WithMimeType(mimeType string) *InlineQueryResultDocument {
	i.MimeType = mimeType
	return i
}

// WithDescription adds description parameter
func (i *InlineQueryResultDocument) WithDescription(description string) *InlineQueryResultDocument {
	i.Description = description
	return i
}

// WithReplyMarkup adds reply markup parameter
func (i *InlineQueryResultDocument) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *InlineQueryResultDocument {
	i.ReplyMarkup = replyMarkup
	return i
}

// WithInputMessageContent adds input message content parameter
func (i *InlineQueryResultDocument) WithInputMessageContent(inputMessageContent InputMessageContent,
) *InlineQueryResultDocument {
	i.InputMessageContent = inputMessageContent
	return i
}

// WithThumbnailURL adds thumbnail URL parameter
func (i *InlineQueryResultDocument) WithThumbnailURL(thumbnailURL string) *InlineQueryResultDocument {
	i.ThumbnailURL = thumbnailURL
	return i
}

// WithThumbnailWidth adds thumbnail width parameter
func (i *InlineQueryResultDocument) WithThumbnailWidth(thumbnailWidth int) *InlineQueryResultDocument {
	i.ThumbnailWidth = thumbnailWidth
	return i
}

// WithThumbnailHeight adds thumbnail height parameter
func (i *InlineQueryResultDocument) WithThumbnailHeight(thumbnailHeight int) *InlineQueryResultDocument {
	i.ThumbnailHeight = thumbnailHeight
	return i
}

// WithID adds ID parameter
func (i *InlineQueryResultLocation) WithID(iD string) *InlineQueryResultLocation {
	i.ID = iD
	return i
}

// WithTitle adds title parameter
func (i *InlineQueryResultLocation) WithTitle(title string) *InlineQueryResultLocation {
	i.Title = title
	return i
}

// WithLivePeriod adds live period parameter
func (i *InlineQueryResultLocation) WithLivePeriod(livePeriod int) *InlineQueryResultLocation {
	i.LivePeriod = livePeriod
	return i
}

// WithHeading adds heading parameter
func (i *InlineQueryResultLocation) WithHeading(heading int) *InlineQueryResultLocation {
	i.Heading = heading
	return i
}

// WithProximityAlertRadius adds proximity alert radius parameter
func (i *InlineQueryResultLocation) WithProximityAlertRadius(proximityAlertRadius int) *InlineQueryResultLocation {
	i.ProximityAlertRadius = proximityAlertRadius
	return i
}

// WithReplyMarkup adds reply markup parameter
func (i *InlineQueryResultLocation) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *InlineQueryResultLocation {
	i.ReplyMarkup = replyMarkup
	return i
}

// WithInputMessageContent adds input message content parameter
func (i *InlineQueryResultLocation) WithInputMessageContent(inputMessageContent InputMessageContent,
) *InlineQueryResultLocation {
	i.InputMessageContent = inputMessageContent
	return i
}

// WithThumbnailURL adds thumbnail URL parameter
func (i *InlineQueryResultLocation) WithThumbnailURL(thumbnailURL string) *InlineQueryResultLocation {
	i.ThumbnailURL = thumbnailURL
	return i
}

// WithThumbnailWidth adds thumbnail width parameter
func (i *InlineQueryResultLocation) WithThumbnailWidth(thumbnailWidth int) *InlineQueryResultLocation {
	i.ThumbnailWidth = thumbnailWidth
	return i
}

// WithThumbnailHeight adds thumbnail height parameter
func (i *InlineQueryResultLocation) WithThumbnailHeight(thumbnailHeight int) *InlineQueryResultLocation {
	i.ThumbnailHeight = thumbnailHeight
	return i
}

// WithID adds ID parameter
func (i *InlineQueryResultVenue) WithID(iD string) *InlineQueryResultVenue {
	i.ID = iD
	return i
}

// WithTitle adds title parameter
func (i *InlineQueryResultVenue) WithTitle(title string) *InlineQueryResultVenue {
	i.Title = title
	return i
}

// WithAddress adds address parameter
func (i *InlineQueryResultVenue) WithAddress(address string) *InlineQueryResultVenue {
	i.Address = address
	return i
}

// WithFoursquareID adds foursquare ID parameter
func (i *InlineQueryResultVenue) WithFoursquareID(foursquareID string) *InlineQueryResultVenue {
	i.FoursquareID = foursquareID
	return i
}

// WithFoursquareType adds foursquare type parameter
func (i *InlineQueryResultVenue) WithFoursquareType(foursquareType string) *InlineQueryResultVenue {
	i.FoursquareType = foursquareType
	return i
}

// WithGooglePlaceID adds google place ID parameter
func (i *InlineQueryResultVenue) WithGooglePlaceID(googlePlaceID string) *InlineQueryResultVenue {
	i.GooglePlaceID = googlePlaceID
	return i
}

// WithGooglePlaceType adds google place type parameter
func (i *InlineQueryResultVenue) WithGooglePlaceType(googlePlaceType string) *InlineQueryResultVenue {
	i.GooglePlaceType = googlePlaceType
	return i
}

// WithReplyMarkup adds reply markup parameter
func (i *InlineQueryResultVenue) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *InlineQueryResultVenue {
	i.ReplyMarkup = replyMarkup
	return i
}

// WithInputMessageContent adds input message content parameter
func (i *InlineQueryResultVenue) WithInputMessageContent(inputMessageContent InputMessageContent,
) *InlineQueryResultVenue {
	i.InputMessageContent = inputMessageContent
	return i
}

// WithThumbnailURL adds thumbnail URL parameter
func (i *InlineQueryResultVenue) WithThumbnailURL(thumbnailURL string) *InlineQueryResultVenue {
	i.ThumbnailURL = thumbnailURL
	return i
}

// WithThumbnailWidth adds thumbnail width parameter
func (i *InlineQueryResultVenue) WithThumbnailWidth(thumbnailWidth int) *InlineQueryResultVenue {
	i.ThumbnailWidth = thumbnailWidth
	return i
}

// WithThumbnailHeight adds thumbnail height parameter
func (i *InlineQueryResultVenue) WithThumbnailHeight(thumbnailHeight int) *InlineQueryResultVenue {
	i.ThumbnailHeight = thumbnailHeight
	return i
}

// WithID adds ID parameter
func (i *InlineQueryResultContact) WithID(iD string) *InlineQueryResultContact {
	i.ID = iD
	return i
}

// WithPhoneNumber adds phone number parameter
func (i *InlineQueryResultContact) WithPhoneNumber(phoneNumber string) *InlineQueryResultContact {
	i.PhoneNumber = phoneNumber
	return i
}

// WithFirstName adds first name parameter
func (i *InlineQueryResultContact) WithFirstName(firstName string) *InlineQueryResultContact {
	i.FirstName = firstName
	return i
}

// WithLastName adds last name parameter
func (i *InlineQueryResultContact) WithLastName(lastName string) *InlineQueryResultContact {
	i.LastName = lastName
	return i
}

// WithVcard adds vcard parameter
func (i *InlineQueryResultContact) WithVcard(vcard string) *InlineQueryResultContact {
	i.Vcard = vcard
	return i
}

// WithReplyMarkup adds reply markup parameter
func (i *InlineQueryResultContact) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *InlineQueryResultContact {
	i.ReplyMarkup = replyMarkup
	return i
}

// WithInputMessageContent adds input message content parameter
func (i *InlineQueryResultContact) WithInputMessageContent(inputMessageContent InputMessageContent,
) *InlineQueryResultContact {
	i.InputMessageContent = inputMessageContent
	return i
}

// WithThumbnailURL adds thumbnail URL parameter
func (i *InlineQueryResultContact) WithThumbnailURL(thumbnailURL string) *InlineQueryResultContact {
	i.ThumbnailURL = thumbnailURL
	return i
}

// WithThumbnailWidth adds thumbnail width parameter
func (i *InlineQueryResultContact) WithThumbnailWidth(thumbnailWidth int) *InlineQueryResultContact {
	i.ThumbnailWidth = thumbnailWidth
	return i
}

// WithThumbnailHeight adds thumbnail height parameter
func (i *InlineQueryResultContact) WithThumbnailHeight(thumbnailHeight int) *InlineQueryResultContact {
	i.ThumbnailHeight = thumbnailHeight
	return i
}

// WithID adds ID parameter
func (i *InlineQueryResultGame) WithID(iD string) *InlineQueryResultGame {
	i.ID = iD
	return i
}

// WithGameShortName adds game short name parameter
func (i *InlineQueryResultGame) WithGameShortName(gameShortName string) *InlineQueryResultGame {
	i.GameShortName = gameShortName
	return i
}

// WithReplyMarkup adds reply markup parameter
func (i *InlineQueryResultGame) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *InlineQueryResultGame {
	i.ReplyMarkup = replyMarkup
	return i
}

// WithID adds ID parameter
func (i *InlineQueryResultCachedPhoto) WithID(iD string) *InlineQueryResultCachedPhoto {
	i.ID = iD
	return i
}

// WithPhotoFileID adds photo file ID parameter
func (i *InlineQueryResultCachedPhoto) WithPhotoFileID(photoFileID string) *InlineQueryResultCachedPhoto {
	i.PhotoFileID = photoFileID
	return i
}

// WithTitle adds title parameter
func (i *InlineQueryResultCachedPhoto) WithTitle(title string) *InlineQueryResultCachedPhoto {
	i.Title = title
	return i
}

// WithDescription adds description parameter
func (i *InlineQueryResultCachedPhoto) WithDescription(description string) *InlineQueryResultCachedPhoto {
	i.Description = description
	return i
}

// WithCaption adds caption parameter
func (i *InlineQueryResultCachedPhoto) WithCaption(caption string) *InlineQueryResultCachedPhoto {
	i.Caption = caption
	return i
}

// WithParseMode adds parse mode parameter
func (i *InlineQueryResultCachedPhoto) WithParseMode(parseMode string) *InlineQueryResultCachedPhoto {
	i.ParseMode = parseMode
	return i
}

// WithCaptionEntities adds caption entities parameter
func (i *InlineQueryResultCachedPhoto) WithCaptionEntities(captionEntities ...MessageEntity,
) *InlineQueryResultCachedPhoto {
	i.CaptionEntities = captionEntities
	return i
}

// WithReplyMarkup adds reply markup parameter
func (i *InlineQueryResultCachedPhoto) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup,
) *InlineQueryResultCachedPhoto {
	i.ReplyMarkup = replyMarkup
	return i
}

// WithInputMessageContent adds input message content parameter
func (i *InlineQueryResultCachedPhoto) WithInputMessageContent(inputMessageContent InputMessageContent,
) *InlineQueryResultCachedPhoto {
	i.InputMessageContent = inputMessageContent
	return i
}

// WithID adds ID parameter
func (i *InlineQueryResultCachedGif) WithID(iD string) *InlineQueryResultCachedGif {
	i.ID = iD
	return i
}

// WithGifFileID adds gif file ID parameter
func (i *InlineQueryResultCachedGif) WithGifFileID(gifFileID string) *InlineQueryResultCachedGif {
	i.GifFileID = gifFileID
	return i
}

// WithTitle adds title parameter
func (i *InlineQueryResultCachedGif) WithTitle(title string) *InlineQueryResultCachedGif {
	i.Title = title
	return i
}

// WithCaption adds caption parameter
func (i *InlineQueryResultCachedGif) WithCaption(caption string) *InlineQueryResultCachedGif {
	i.Caption = caption
	return i
}

// WithParseMode adds parse mode parameter
func (i *InlineQueryResultCachedGif) WithParseMode(parseMode string) *InlineQueryResultCachedGif {
	i.ParseMode = parseMode
	return i
}

// WithCaptionEntities adds caption entities parameter
func (i *InlineQueryResultCachedGif) WithCaptionEntities(captionEntities ...MessageEntity) *InlineQueryResultCachedGif {
	i.CaptionEntities = captionEntities
	return i
}

// WithReplyMarkup adds reply markup parameter
func (i *InlineQueryResultCachedGif) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *InlineQueryResultCachedGif {
	i.ReplyMarkup = replyMarkup
	return i
}

// WithInputMessageContent adds input message content parameter
func (i *InlineQueryResultCachedGif) WithInputMessageContent(inputMessageContent InputMessageContent,
) *InlineQueryResultCachedGif {
	i.InputMessageContent = inputMessageContent
	return i
}

// WithID adds ID parameter
func (i *InlineQueryResultCachedMpeg4Gif) WithID(iD string) *InlineQueryResultCachedMpeg4Gif {
	i.ID = iD
	return i
}

// WithMpeg4FileID adds mpeg4 file ID parameter
func (i *InlineQueryResultCachedMpeg4Gif) WithMpeg4FileID(mpeg4FileID string) *InlineQueryResultCachedMpeg4Gif {
	i.Mpeg4FileID = mpeg4FileID
	return i
}

// WithTitle adds title parameter
func (i *InlineQueryResultCachedMpeg4Gif) WithTitle(title string) *InlineQueryResultCachedMpeg4Gif {
	i.Title = title
	return i
}

// WithCaption adds caption parameter
func (i *InlineQueryResultCachedMpeg4Gif) WithCaption(caption string) *InlineQueryResultCachedMpeg4Gif {
	i.Caption = caption
	return i
}

// WithParseMode adds parse mode parameter
func (i *InlineQueryResultCachedMpeg4Gif) WithParseMode(parseMode string) *InlineQueryResultCachedMpeg4Gif {
	i.ParseMode = parseMode
	return i
}

// WithCaptionEntities adds caption entities parameter
func (i *InlineQueryResultCachedMpeg4Gif) WithCaptionEntities(captionEntities ...MessageEntity,
) *InlineQueryResultCachedMpeg4Gif {
	i.CaptionEntities = captionEntities
	return i
}

// WithReplyMarkup adds reply markup parameter
func (i *InlineQueryResultCachedMpeg4Gif) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup,
) *InlineQueryResultCachedMpeg4Gif {
	i.ReplyMarkup = replyMarkup
	return i
}

// WithInputMessageContent adds input message content parameter
func (i *InlineQueryResultCachedMpeg4Gif) WithInputMessageContent(inputMessageContent InputMessageContent,
) *InlineQueryResultCachedMpeg4Gif {
	i.InputMessageContent = inputMessageContent
	return i
}

// WithID adds ID parameter
func (i *InlineQueryResultCachedSticker) WithID(iD string) *InlineQueryResultCachedSticker {
	i.ID = iD
	return i
}

// WithStickerFileID adds sticker file ID parameter
func (i *InlineQueryResultCachedSticker) WithStickerFileID(stickerFileID string) *InlineQueryResultCachedSticker {
	i.StickerFileID = stickerFileID
	return i
}

// WithReplyMarkup adds reply markup parameter
func (i *InlineQueryResultCachedSticker) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup,
) *InlineQueryResultCachedSticker {
	i.ReplyMarkup = replyMarkup
	return i
}

// WithInputMessageContent adds input message content parameter
func (i *InlineQueryResultCachedSticker) WithInputMessageContent(inputMessageContent InputMessageContent,
) *InlineQueryResultCachedSticker {
	i.InputMessageContent = inputMessageContent
	return i
}

// WithID adds ID parameter
func (i *InlineQueryResultCachedDocument) WithID(iD string) *InlineQueryResultCachedDocument {
	i.ID = iD
	return i
}

// WithTitle adds title parameter
func (i *InlineQueryResultCachedDocument) WithTitle(title string) *InlineQueryResultCachedDocument {
	i.Title = title
	return i
}

// WithDocumentFileID adds document file ID parameter
func (i *InlineQueryResultCachedDocument) WithDocumentFileID(documentFileID string) *InlineQueryResultCachedDocument {
	i.DocumentFileID = documentFileID
	return i
}

// WithDescription adds description parameter
func (i *InlineQueryResultCachedDocument) WithDescription(description string) *InlineQueryResultCachedDocument {
	i.Description = description
	return i
}

// WithCaption adds caption parameter
func (i *InlineQueryResultCachedDocument) WithCaption(caption string) *InlineQueryResultCachedDocument {
	i.Caption = caption
	return i
}

// WithParseMode adds parse mode parameter
func (i *InlineQueryResultCachedDocument) WithParseMode(parseMode string) *InlineQueryResultCachedDocument {
	i.ParseMode = parseMode
	return i
}

// WithCaptionEntities adds caption entities parameter
func (i *InlineQueryResultCachedDocument) WithCaptionEntities(captionEntities ...MessageEntity,
) *InlineQueryResultCachedDocument {
	i.CaptionEntities = captionEntities
	return i
}

// WithReplyMarkup adds reply markup parameter
func (i *InlineQueryResultCachedDocument) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup,
) *InlineQueryResultCachedDocument {
	i.ReplyMarkup = replyMarkup
	return i
}

// WithInputMessageContent adds input message content parameter
func (i *InlineQueryResultCachedDocument) WithInputMessageContent(inputMessageContent InputMessageContent,
) *InlineQueryResultCachedDocument {
	i.InputMessageContent = inputMessageContent
	return i
}

// WithID adds ID parameter
func (i *InlineQueryResultCachedVideo) WithID(iD string) *InlineQueryResultCachedVideo {
	i.ID = iD
	return i
}

// WithVideoFileID adds video file ID parameter
func (i *InlineQueryResultCachedVideo) WithVideoFileID(videoFileID string) *InlineQueryResultCachedVideo {
	i.VideoFileID = videoFileID
	return i
}

// WithTitle adds title parameter
func (i *InlineQueryResultCachedVideo) WithTitle(title string) *InlineQueryResultCachedVideo {
	i.Title = title
	return i
}

// WithDescription adds description parameter
func (i *InlineQueryResultCachedVideo) WithDescription(description string) *InlineQueryResultCachedVideo {
	i.Description = description
	return i
}

// WithCaption adds caption parameter
func (i *InlineQueryResultCachedVideo) WithCaption(caption string) *InlineQueryResultCachedVideo {
	i.Caption = caption
	return i
}

// WithParseMode adds parse mode parameter
func (i *InlineQueryResultCachedVideo) WithParseMode(parseMode string) *InlineQueryResultCachedVideo {
	i.ParseMode = parseMode
	return i
}

// WithCaptionEntities adds caption entities parameter
func (i *InlineQueryResultCachedVideo) WithCaptionEntities(captionEntities ...MessageEntity,
) *InlineQueryResultCachedVideo {
	i.CaptionEntities = captionEntities
	return i
}

// WithReplyMarkup adds reply markup parameter
func (i *InlineQueryResultCachedVideo) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup,
) *InlineQueryResultCachedVideo {
	i.ReplyMarkup = replyMarkup
	return i
}

// WithInputMessageContent adds input message content parameter
func (i *InlineQueryResultCachedVideo) WithInputMessageContent(inputMessageContent InputMessageContent,
) *InlineQueryResultCachedVideo {
	i.InputMessageContent = inputMessageContent
	return i
}

// WithID adds ID parameter
func (i *InlineQueryResultCachedVoice) WithID(iD string) *InlineQueryResultCachedVoice {
	i.ID = iD
	return i
}

// WithVoiceFileID adds voice file ID parameter
func (i *InlineQueryResultCachedVoice) WithVoiceFileID(voiceFileID string) *InlineQueryResultCachedVoice {
	i.VoiceFileID = voiceFileID
	return i
}

// WithTitle adds title parameter
func (i *InlineQueryResultCachedVoice) WithTitle(title string) *InlineQueryResultCachedVoice {
	i.Title = title
	return i
}

// WithCaption adds caption parameter
func (i *InlineQueryResultCachedVoice) WithCaption(caption string) *InlineQueryResultCachedVoice {
	i.Caption = caption
	return i
}

// WithParseMode adds parse mode parameter
func (i *InlineQueryResultCachedVoice) WithParseMode(parseMode string) *InlineQueryResultCachedVoice {
	i.ParseMode = parseMode
	return i
}

// WithCaptionEntities adds caption entities parameter
func (i *InlineQueryResultCachedVoice) WithCaptionEntities(captionEntities ...MessageEntity,
) *InlineQueryResultCachedVoice {
	i.CaptionEntities = captionEntities
	return i
}

// WithReplyMarkup adds reply markup parameter
func (i *InlineQueryResultCachedVoice) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup,
) *InlineQueryResultCachedVoice {
	i.ReplyMarkup = replyMarkup
	return i
}

// WithInputMessageContent adds input message content parameter
func (i *InlineQueryResultCachedVoice) WithInputMessageContent(inputMessageContent InputMessageContent,
) *InlineQueryResultCachedVoice {
	i.InputMessageContent = inputMessageContent
	return i
}

// WithID adds ID parameter
func (i *InlineQueryResultCachedAudio) WithID(iD string) *InlineQueryResultCachedAudio {
	i.ID = iD
	return i
}

// WithAudioFileID adds audio file ID parameter
func (i *InlineQueryResultCachedAudio) WithAudioFileID(audioFileID string) *InlineQueryResultCachedAudio {
	i.AudioFileID = audioFileID
	return i
}

// WithCaption adds caption parameter
func (i *InlineQueryResultCachedAudio) WithCaption(caption string) *InlineQueryResultCachedAudio {
	i.Caption = caption
	return i
}

// WithParseMode adds parse mode parameter
func (i *InlineQueryResultCachedAudio) WithParseMode(parseMode string) *InlineQueryResultCachedAudio {
	i.ParseMode = parseMode
	return i
}

// WithCaptionEntities adds caption entities parameter
func (i *InlineQueryResultCachedAudio) WithCaptionEntities(captionEntities ...MessageEntity,
) *InlineQueryResultCachedAudio {
	i.CaptionEntities = captionEntities
	return i
}

// WithReplyMarkup adds reply markup parameter
func (i *InlineQueryResultCachedAudio) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup,
) *InlineQueryResultCachedAudio {
	i.ReplyMarkup = replyMarkup
	return i
}

// WithInputMessageContent adds input message content parameter
func (i *InlineQueryResultCachedAudio) WithInputMessageContent(inputMessageContent InputMessageContent,
) *InlineQueryResultCachedAudio {
	i.InputMessageContent = inputMessageContent
	return i
}

// WithMessageText adds message text parameter
func (i *InputTextMessageContent) WithMessageText(messageText string) *InputTextMessageContent {
	i.MessageText = messageText
	return i
}

// WithParseMode adds parse mode parameter
func (i *InputTextMessageContent) WithParseMode(parseMode string) *InputTextMessageContent {
	i.ParseMode = parseMode
	return i
}

// WithEntities adds entities parameter
func (i *InputTextMessageContent) WithEntities(entities ...MessageEntity) *InputTextMessageContent {
	i.Entities = entities
	return i
}

// WithDisableWebPagePreview adds disable web page preview parameter
func (i *InputTextMessageContent) WithDisableWebPagePreview() *InputTextMessageContent {
	i.DisableWebPagePreview = true
	return i
}

// WithLivePeriod adds live period parameter
func (i *InputLocationMessageContent) WithLivePeriod(livePeriod int) *InputLocationMessageContent {
	i.LivePeriod = livePeriod
	return i
}

// WithHeading adds heading parameter
func (i *InputLocationMessageContent) WithHeading(heading int) *InputLocationMessageContent {
	i.Heading = heading
	return i
}

// WithProximityAlertRadius adds proximity alert radius parameter
func (i *InputLocationMessageContent) WithProximityAlertRadius(proximityAlertRadius int) *InputLocationMessageContent {
	i.ProximityAlertRadius = proximityAlertRadius
	return i
}

// WithTitle adds title parameter
func (i *InputVenueMessageContent) WithTitle(title string) *InputVenueMessageContent {
	i.Title = title
	return i
}

// WithAddress adds address parameter
func (i *InputVenueMessageContent) WithAddress(address string) *InputVenueMessageContent {
	i.Address = address
	return i
}

// WithFoursquareID adds foursquare ID parameter
func (i *InputVenueMessageContent) WithFoursquareID(foursquareID string) *InputVenueMessageContent {
	i.FoursquareID = foursquareID
	return i
}

// WithFoursquareType adds foursquare type parameter
func (i *InputVenueMessageContent) WithFoursquareType(foursquareType string) *InputVenueMessageContent {
	i.FoursquareType = foursquareType
	return i
}

// WithGooglePlaceID adds google place ID parameter
func (i *InputVenueMessageContent) WithGooglePlaceID(googlePlaceID string) *InputVenueMessageContent {
	i.GooglePlaceID = googlePlaceID
	return i
}

// WithGooglePlaceType adds google place type parameter
func (i *InputVenueMessageContent) WithGooglePlaceType(googlePlaceType string) *InputVenueMessageContent {
	i.GooglePlaceType = googlePlaceType
	return i
}

// WithPhoneNumber adds phone number parameter
func (i *InputContactMessageContent) WithPhoneNumber(phoneNumber string) *InputContactMessageContent {
	i.PhoneNumber = phoneNumber
	return i
}

// WithFirstName adds first name parameter
func (i *InputContactMessageContent) WithFirstName(firstName string) *InputContactMessageContent {
	i.FirstName = firstName
	return i
}

// WithLastName adds last name parameter
func (i *InputContactMessageContent) WithLastName(lastName string) *InputContactMessageContent {
	i.LastName = lastName
	return i
}

// WithVcard adds vcard parameter
func (i *InputContactMessageContent) WithVcard(vcard string) *InputContactMessageContent {
	i.Vcard = vcard
	return i
}

// WithTitle adds title parameter
func (i *InputInvoiceMessageContent) WithTitle(title string) *InputInvoiceMessageContent {
	i.Title = title
	return i
}

// WithDescription adds description parameter
func (i *InputInvoiceMessageContent) WithDescription(description string) *InputInvoiceMessageContent {
	i.Description = description
	return i
}

// WithPayload adds payload parameter
func (i *InputInvoiceMessageContent) WithPayload(payload string) *InputInvoiceMessageContent {
	i.Payload = payload
	return i
}

// WithProviderToken adds provider token parameter
func (i *InputInvoiceMessageContent) WithProviderToken(providerToken string) *InputInvoiceMessageContent {
	i.ProviderToken = providerToken
	return i
}

// WithCurrency adds currency parameter
func (i *InputInvoiceMessageContent) WithCurrency(currency string) *InputInvoiceMessageContent {
	i.Currency = currency
	return i
}

// WithPrices adds prices parameter
func (i *InputInvoiceMessageContent) WithPrices(prices ...LabeledPrice) *InputInvoiceMessageContent {
	i.Prices = prices
	return i
}

// WithMaxTipAmount adds max tip amount parameter
func (i *InputInvoiceMessageContent) WithMaxTipAmount(maxTipAmount int) *InputInvoiceMessageContent {
	i.MaxTipAmount = maxTipAmount
	return i
}

// WithSuggestedTipAmounts adds suggested tip amounts parameter
func (i *InputInvoiceMessageContent) WithSuggestedTipAmounts(suggestedTipAmounts ...int) *InputInvoiceMessageContent {
	i.SuggestedTipAmounts = suggestedTipAmounts
	return i
}

// WithProviderData adds provider data parameter
func (i *InputInvoiceMessageContent) WithProviderData(providerData string) *InputInvoiceMessageContent {
	i.ProviderData = providerData
	return i
}

// WithPhotoURL adds photo URL parameter
func (i *InputInvoiceMessageContent) WithPhotoURL(photoURL string) *InputInvoiceMessageContent {
	i.PhotoURL = photoURL
	return i
}

// WithPhotoSize adds photo size parameter
func (i *InputInvoiceMessageContent) WithPhotoSize(photoSize int) *InputInvoiceMessageContent {
	i.PhotoSize = photoSize
	return i
}

// WithPhotoWidth adds photo width parameter
func (i *InputInvoiceMessageContent) WithPhotoWidth(photoWidth int) *InputInvoiceMessageContent {
	i.PhotoWidth = photoWidth
	return i
}

// WithPhotoHeight adds photo height parameter
func (i *InputInvoiceMessageContent) WithPhotoHeight(photoHeight int) *InputInvoiceMessageContent {
	i.PhotoHeight = photoHeight
	return i
}

// WithNeedName adds need name parameter
func (i *InputInvoiceMessageContent) WithNeedName() *InputInvoiceMessageContent {
	i.NeedName = true
	return i
}

// WithNeedPhoneNumber adds need phone number parameter
func (i *InputInvoiceMessageContent) WithNeedPhoneNumber() *InputInvoiceMessageContent {
	i.NeedPhoneNumber = true
	return i
}

// WithNeedEmail adds need email parameter
func (i *InputInvoiceMessageContent) WithNeedEmail() *InputInvoiceMessageContent {
	i.NeedEmail = true
	return i
}

// WithNeedShippingAddress adds need shipping address parameter
func (i *InputInvoiceMessageContent) WithNeedShippingAddress() *InputInvoiceMessageContent {
	i.NeedShippingAddress = true
	return i
}

// WithSendPhoneNumberToProvider adds send phone number to provider parameter
func (i *InputInvoiceMessageContent) WithSendPhoneNumberToProvider() *InputInvoiceMessageContent {
	i.SendPhoneNumberToProvider = true
	return i
}

// WithSendEmailToProvider adds send email to provider parameter
func (i *InputInvoiceMessageContent) WithSendEmailToProvider() *InputInvoiceMessageContent {
	i.SendEmailToProvider = true
	return i
}

// WithIsFlexible adds is flexible parameter
func (i *InputInvoiceMessageContent) WithIsFlexible() *InputInvoiceMessageContent {
	i.IsFlexible = true
	return i
}
