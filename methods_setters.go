package telego

// WithOffset adds offset parameter
func (p *GetUpdatesParams) WithOffset(offset int) *GetUpdatesParams {
	p.Offset = offset
	return p
}

// WithLimit adds limit parameter
func (p *GetUpdatesParams) WithLimit(limit int) *GetUpdatesParams {
	p.Limit = limit
	return p
}

// WithTimeout adds timeout parameter
func (p *GetUpdatesParams) WithTimeout(timeout int) *GetUpdatesParams {
	p.Timeout = timeout
	return p
}

// WithAllowedUpdates adds allowed updates parameter
func (p *GetUpdatesParams) WithAllowedUpdates(allowedUpdates ...string) *GetUpdatesParams {
	p.AllowedUpdates = allowedUpdates
	return p
}

// WithURL adds URL parameter
func (p *SetWebhookParams) WithURL(url string) *SetWebhookParams {
	p.URL = url
	return p
}

// WithCertificate adds certificate parameter
func (p *SetWebhookParams) WithCertificate(certificate *InputFile) *SetWebhookParams {
	p.Certificate = certificate
	return p
}

// WithIPAddress adds ip address parameter
func (p *SetWebhookParams) WithIPAddress(ipAddress string) *SetWebhookParams {
	p.IPAddress = ipAddress
	return p
}

// WithMaxConnections adds max connections parameter
func (p *SetWebhookParams) WithMaxConnections(maxConnections int) *SetWebhookParams {
	p.MaxConnections = maxConnections
	return p
}

// WithAllowedUpdates adds allowed updates parameter
func (p *SetWebhookParams) WithAllowedUpdates(allowedUpdates ...string) *SetWebhookParams {
	p.AllowedUpdates = allowedUpdates
	return p
}

// WithDropPendingUpdates adds drop pending updates parameter
func (p *SetWebhookParams) WithDropPendingUpdates() *SetWebhookParams {
	p.DropPendingUpdates = true
	return p
}

// WithSecretToken adds secret token parameter
func (p *SetWebhookParams) WithSecretToken(secretToken string) *SetWebhookParams {
	p.SecretToken = secretToken
	return p
}

// WithDropPendingUpdates adds drop pending updates parameter
func (p *DeleteWebhookParams) WithDropPendingUpdates() *DeleteWebhookParams {
	p.DropPendingUpdates = true
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *SendMessageParams) WithBusinessConnectionID(businessConnectionID string) *SendMessageParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithChatID adds chat ID parameter
func (p *SendMessageParams) WithChatID(chatID ChatID) *SendMessageParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *SendMessageParams) WithMessageThreadID(messageThreadID int) *SendMessageParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithText adds text parameter
func (p *SendMessageParams) WithText(text string) *SendMessageParams {
	p.Text = text
	return p
}

// WithParseMode adds parse mode parameter
func (p *SendMessageParams) WithParseMode(parseMode string) *SendMessageParams {
	p.ParseMode = parseMode
	return p
}

// WithEntities adds entities parameter
func (p *SendMessageParams) WithEntities(entities ...MessageEntity) *SendMessageParams {
	p.Entities = entities
	return p
}

// WithLinkPreviewOptions adds link preview options parameter
func (p *SendMessageParams) WithLinkPreviewOptions(linkPreviewOptions *LinkPreviewOptions) *SendMessageParams {
	p.LinkPreviewOptions = linkPreviewOptions
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendMessageParams) WithDisableNotification() *SendMessageParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendMessageParams) WithProtectContent() *SendMessageParams {
	p.ProtectContent = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendMessageParams) WithMessageEffectID(messageEffectID string) *SendMessageParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithReplyParameters adds reply parameters parameter
func (p *SendMessageParams) WithReplyParameters(replyParameters *ReplyParameters) *SendMessageParams {
	p.ReplyParameters = replyParameters
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendMessageParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendMessageParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *ForwardMessageParams) WithChatID(chatID ChatID) *ForwardMessageParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *ForwardMessageParams) WithMessageThreadID(messageThreadID int) *ForwardMessageParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithFromChatID adds from chat ID parameter
func (p *ForwardMessageParams) WithFromChatID(fromChatID ChatID) *ForwardMessageParams {
	p.FromChatID = fromChatID
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *ForwardMessageParams) WithDisableNotification() *ForwardMessageParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *ForwardMessageParams) WithProtectContent() *ForwardMessageParams {
	p.ProtectContent = true
	return p
}

// WithMessageID adds message ID parameter
func (p *ForwardMessageParams) WithMessageID(messageID int) *ForwardMessageParams {
	p.MessageID = messageID
	return p
}

// WithChatID adds chat ID parameter
func (p *ForwardMessagesParams) WithChatID(chatID ChatID) *ForwardMessagesParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *ForwardMessagesParams) WithMessageThreadID(messageThreadID int) *ForwardMessagesParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithFromChatID adds from chat ID parameter
func (p *ForwardMessagesParams) WithFromChatID(fromChatID ChatID) *ForwardMessagesParams {
	p.FromChatID = fromChatID
	return p
}

// WithMessageIDs adds message ids parameter
func (p *ForwardMessagesParams) WithMessageIDs(messageIDs ...int) *ForwardMessagesParams {
	p.MessageIDs = messageIDs
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *ForwardMessagesParams) WithDisableNotification() *ForwardMessagesParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *ForwardMessagesParams) WithProtectContent() *ForwardMessagesParams {
	p.ProtectContent = true
	return p
}

// WithChatID adds chat ID parameter
func (p *CopyMessageParams) WithChatID(chatID ChatID) *CopyMessageParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *CopyMessageParams) WithMessageThreadID(messageThreadID int) *CopyMessageParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithFromChatID adds from chat ID parameter
func (p *CopyMessageParams) WithFromChatID(fromChatID ChatID) *CopyMessageParams {
	p.FromChatID = fromChatID
	return p
}

// WithMessageID adds message ID parameter
func (p *CopyMessageParams) WithMessageID(messageID int) *CopyMessageParams {
	p.MessageID = messageID
	return p
}

// WithCaption adds caption parameter
func (p *CopyMessageParams) WithCaption(caption string) *CopyMessageParams {
	p.Caption = caption
	return p
}

// WithParseMode adds parse mode parameter
func (p *CopyMessageParams) WithParseMode(parseMode string) *CopyMessageParams {
	p.ParseMode = parseMode
	return p
}

// WithCaptionEntities adds caption entities parameter
func (p *CopyMessageParams) WithCaptionEntities(captionEntities ...MessageEntity) *CopyMessageParams {
	p.CaptionEntities = captionEntities
	return p
}

// WithShowCaptionAboveMedia adds show caption above media parameter
func (p *CopyMessageParams) WithShowCaptionAboveMedia() *CopyMessageParams {
	p.ShowCaptionAboveMedia = true
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *CopyMessageParams) WithDisableNotification() *CopyMessageParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *CopyMessageParams) WithProtectContent() *CopyMessageParams {
	p.ProtectContent = true
	return p
}

// WithReplyParameters adds reply parameters parameter
func (p *CopyMessageParams) WithReplyParameters(replyParameters *ReplyParameters) *CopyMessageParams {
	p.ReplyParameters = replyParameters
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *CopyMessageParams) WithReplyMarkup(replyMarkup ReplyMarkup) *CopyMessageParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *CopyMessagesParams) WithChatID(chatID ChatID) *CopyMessagesParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *CopyMessagesParams) WithMessageThreadID(messageThreadID int) *CopyMessagesParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithFromChatID adds from chat ID parameter
func (p *CopyMessagesParams) WithFromChatID(fromChatID ChatID) *CopyMessagesParams {
	p.FromChatID = fromChatID
	return p
}

// WithMessageIDs adds message ids parameter
func (p *CopyMessagesParams) WithMessageIDs(messageIDs ...int) *CopyMessagesParams {
	p.MessageIDs = messageIDs
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *CopyMessagesParams) WithDisableNotification() *CopyMessagesParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *CopyMessagesParams) WithProtectContent() *CopyMessagesParams {
	p.ProtectContent = true
	return p
}

// WithRemoveCaption adds remove caption parameter
func (p *CopyMessagesParams) WithRemoveCaption() *CopyMessagesParams {
	p.RemoveCaption = true
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *SendPhotoParams) WithBusinessConnectionID(businessConnectionID string) *SendPhotoParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithChatID adds chat ID parameter
func (p *SendPhotoParams) WithChatID(chatID ChatID) *SendPhotoParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *SendPhotoParams) WithMessageThreadID(messageThreadID int) *SendPhotoParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithPhoto adds photo parameter
func (p *SendPhotoParams) WithPhoto(photo InputFile) *SendPhotoParams {
	p.Photo = photo
	return p
}

// WithCaption adds caption parameter
func (p *SendPhotoParams) WithCaption(caption string) *SendPhotoParams {
	p.Caption = caption
	return p
}

// WithParseMode adds parse mode parameter
func (p *SendPhotoParams) WithParseMode(parseMode string) *SendPhotoParams {
	p.ParseMode = parseMode
	return p
}

// WithCaptionEntities adds caption entities parameter
func (p *SendPhotoParams) WithCaptionEntities(captionEntities ...MessageEntity) *SendPhotoParams {
	p.CaptionEntities = captionEntities
	return p
}

// WithShowCaptionAboveMedia adds show caption above media parameter
func (p *SendPhotoParams) WithShowCaptionAboveMedia() *SendPhotoParams {
	p.ShowCaptionAboveMedia = true
	return p
}

// WithHasSpoiler adds has spoiler parameter
func (p *SendPhotoParams) WithHasSpoiler() *SendPhotoParams {
	p.HasSpoiler = true
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendPhotoParams) WithDisableNotification() *SendPhotoParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendPhotoParams) WithProtectContent() *SendPhotoParams {
	p.ProtectContent = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendPhotoParams) WithMessageEffectID(messageEffectID string) *SendPhotoParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithReplyParameters adds reply parameters parameter
func (p *SendPhotoParams) WithReplyParameters(replyParameters *ReplyParameters) *SendPhotoParams {
	p.ReplyParameters = replyParameters
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendPhotoParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendPhotoParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *SendAudioParams) WithBusinessConnectionID(businessConnectionID string) *SendAudioParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithChatID adds chat ID parameter
func (p *SendAudioParams) WithChatID(chatID ChatID) *SendAudioParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *SendAudioParams) WithMessageThreadID(messageThreadID int) *SendAudioParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithAudio adds audio parameter
func (p *SendAudioParams) WithAudio(audio InputFile) *SendAudioParams {
	p.Audio = audio
	return p
}

// WithCaption adds caption parameter
func (p *SendAudioParams) WithCaption(caption string) *SendAudioParams {
	p.Caption = caption
	return p
}

// WithParseMode adds parse mode parameter
func (p *SendAudioParams) WithParseMode(parseMode string) *SendAudioParams {
	p.ParseMode = parseMode
	return p
}

// WithCaptionEntities adds caption entities parameter
func (p *SendAudioParams) WithCaptionEntities(captionEntities ...MessageEntity) *SendAudioParams {
	p.CaptionEntities = captionEntities
	return p
}

// WithDuration adds duration parameter
func (p *SendAudioParams) WithDuration(duration int) *SendAudioParams {
	p.Duration = duration
	return p
}

// WithPerformer adds performer parameter
func (p *SendAudioParams) WithPerformer(performer string) *SendAudioParams {
	p.Performer = performer
	return p
}

// WithTitle adds title parameter
func (p *SendAudioParams) WithTitle(title string) *SendAudioParams {
	p.Title = title
	return p
}

// WithThumbnail adds thumbnail parameter
func (p *SendAudioParams) WithThumbnail(thumbnail *InputFile) *SendAudioParams {
	p.Thumbnail = thumbnail
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendAudioParams) WithDisableNotification() *SendAudioParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendAudioParams) WithProtectContent() *SendAudioParams {
	p.ProtectContent = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendAudioParams) WithMessageEffectID(messageEffectID string) *SendAudioParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithReplyParameters adds reply parameters parameter
func (p *SendAudioParams) WithReplyParameters(replyParameters *ReplyParameters) *SendAudioParams {
	p.ReplyParameters = replyParameters
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendAudioParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendAudioParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *SendDocumentParams) WithBusinessConnectionID(businessConnectionID string) *SendDocumentParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithChatID adds chat ID parameter
func (p *SendDocumentParams) WithChatID(chatID ChatID) *SendDocumentParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *SendDocumentParams) WithMessageThreadID(messageThreadID int) *SendDocumentParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithDocument adds document parameter
func (p *SendDocumentParams) WithDocument(document InputFile) *SendDocumentParams {
	p.Document = document
	return p
}

// WithThumbnail adds thumbnail parameter
func (p *SendDocumentParams) WithThumbnail(thumbnail *InputFile) *SendDocumentParams {
	p.Thumbnail = thumbnail
	return p
}

// WithCaption adds caption parameter
func (p *SendDocumentParams) WithCaption(caption string) *SendDocumentParams {
	p.Caption = caption
	return p
}

// WithParseMode adds parse mode parameter
func (p *SendDocumentParams) WithParseMode(parseMode string) *SendDocumentParams {
	p.ParseMode = parseMode
	return p
}

// WithCaptionEntities adds caption entities parameter
func (p *SendDocumentParams) WithCaptionEntities(captionEntities ...MessageEntity) *SendDocumentParams {
	p.CaptionEntities = captionEntities
	return p
}

// WithDisableContentTypeDetection adds disable content type detection parameter
func (p *SendDocumentParams) WithDisableContentTypeDetection() *SendDocumentParams {
	p.DisableContentTypeDetection = true
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendDocumentParams) WithDisableNotification() *SendDocumentParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendDocumentParams) WithProtectContent() *SendDocumentParams {
	p.ProtectContent = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendDocumentParams) WithMessageEffectID(messageEffectID string) *SendDocumentParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithReplyParameters adds reply parameters parameter
func (p *SendDocumentParams) WithReplyParameters(replyParameters *ReplyParameters) *SendDocumentParams {
	p.ReplyParameters = replyParameters
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendDocumentParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendDocumentParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *SendVideoParams) WithBusinessConnectionID(businessConnectionID string) *SendVideoParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithChatID adds chat ID parameter
func (p *SendVideoParams) WithChatID(chatID ChatID) *SendVideoParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *SendVideoParams) WithMessageThreadID(messageThreadID int) *SendVideoParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithVideo adds video parameter
func (p *SendVideoParams) WithVideo(video InputFile) *SendVideoParams {
	p.Video = video
	return p
}

// WithDuration adds duration parameter
func (p *SendVideoParams) WithDuration(duration int) *SendVideoParams {
	p.Duration = duration
	return p
}

// WithWidth adds width parameter
func (p *SendVideoParams) WithWidth(width int) *SendVideoParams {
	p.Width = width
	return p
}

// WithHeight adds height parameter
func (p *SendVideoParams) WithHeight(height int) *SendVideoParams {
	p.Height = height
	return p
}

// WithThumbnail adds thumbnail parameter
func (p *SendVideoParams) WithThumbnail(thumbnail *InputFile) *SendVideoParams {
	p.Thumbnail = thumbnail
	return p
}

// WithCaption adds caption parameter
func (p *SendVideoParams) WithCaption(caption string) *SendVideoParams {
	p.Caption = caption
	return p
}

// WithParseMode adds parse mode parameter
func (p *SendVideoParams) WithParseMode(parseMode string) *SendVideoParams {
	p.ParseMode = parseMode
	return p
}

// WithCaptionEntities adds caption entities parameter
func (p *SendVideoParams) WithCaptionEntities(captionEntities ...MessageEntity) *SendVideoParams {
	p.CaptionEntities = captionEntities
	return p
}

// WithShowCaptionAboveMedia adds show caption above media parameter
func (p *SendVideoParams) WithShowCaptionAboveMedia() *SendVideoParams {
	p.ShowCaptionAboveMedia = true
	return p
}

// WithHasSpoiler adds has spoiler parameter
func (p *SendVideoParams) WithHasSpoiler() *SendVideoParams {
	p.HasSpoiler = true
	return p
}

// WithSupportsStreaming adds supports streaming parameter
func (p *SendVideoParams) WithSupportsStreaming() *SendVideoParams {
	p.SupportsStreaming = true
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendVideoParams) WithDisableNotification() *SendVideoParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendVideoParams) WithProtectContent() *SendVideoParams {
	p.ProtectContent = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendVideoParams) WithMessageEffectID(messageEffectID string) *SendVideoParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithReplyParameters adds reply parameters parameter
func (p *SendVideoParams) WithReplyParameters(replyParameters *ReplyParameters) *SendVideoParams {
	p.ReplyParameters = replyParameters
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendVideoParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendVideoParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *SendAnimationParams) WithBusinessConnectionID(businessConnectionID string) *SendAnimationParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithChatID adds chat ID parameter
func (p *SendAnimationParams) WithChatID(chatID ChatID) *SendAnimationParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *SendAnimationParams) WithMessageThreadID(messageThreadID int) *SendAnimationParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithAnimation adds animation parameter
func (p *SendAnimationParams) WithAnimation(animation InputFile) *SendAnimationParams {
	p.Animation = animation
	return p
}

// WithDuration adds duration parameter
func (p *SendAnimationParams) WithDuration(duration int) *SendAnimationParams {
	p.Duration = duration
	return p
}

// WithWidth adds width parameter
func (p *SendAnimationParams) WithWidth(width int) *SendAnimationParams {
	p.Width = width
	return p
}

// WithHeight adds height parameter
func (p *SendAnimationParams) WithHeight(height int) *SendAnimationParams {
	p.Height = height
	return p
}

// WithThumbnail adds thumbnail parameter
func (p *SendAnimationParams) WithThumbnail(thumbnail *InputFile) *SendAnimationParams {
	p.Thumbnail = thumbnail
	return p
}

// WithCaption adds caption parameter
func (p *SendAnimationParams) WithCaption(caption string) *SendAnimationParams {
	p.Caption = caption
	return p
}

// WithParseMode adds parse mode parameter
func (p *SendAnimationParams) WithParseMode(parseMode string) *SendAnimationParams {
	p.ParseMode = parseMode
	return p
}

// WithCaptionEntities adds caption entities parameter
func (p *SendAnimationParams) WithCaptionEntities(captionEntities ...MessageEntity) *SendAnimationParams {
	p.CaptionEntities = captionEntities
	return p
}

// WithShowCaptionAboveMedia adds show caption above media parameter
func (p *SendAnimationParams) WithShowCaptionAboveMedia() *SendAnimationParams {
	p.ShowCaptionAboveMedia = true
	return p
}

// WithHasSpoiler adds has spoiler parameter
func (p *SendAnimationParams) WithHasSpoiler() *SendAnimationParams {
	p.HasSpoiler = true
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendAnimationParams) WithDisableNotification() *SendAnimationParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendAnimationParams) WithProtectContent() *SendAnimationParams {
	p.ProtectContent = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendAnimationParams) WithMessageEffectID(messageEffectID string) *SendAnimationParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithReplyParameters adds reply parameters parameter
func (p *SendAnimationParams) WithReplyParameters(replyParameters *ReplyParameters) *SendAnimationParams {
	p.ReplyParameters = replyParameters
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendAnimationParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendAnimationParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *SendVoiceParams) WithBusinessConnectionID(businessConnectionID string) *SendVoiceParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithChatID adds chat ID parameter
func (p *SendVoiceParams) WithChatID(chatID ChatID) *SendVoiceParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *SendVoiceParams) WithMessageThreadID(messageThreadID int) *SendVoiceParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithVoice adds voice parameter
func (p *SendVoiceParams) WithVoice(voice InputFile) *SendVoiceParams {
	p.Voice = voice
	return p
}

// WithCaption adds caption parameter
func (p *SendVoiceParams) WithCaption(caption string) *SendVoiceParams {
	p.Caption = caption
	return p
}

// WithParseMode adds parse mode parameter
func (p *SendVoiceParams) WithParseMode(parseMode string) *SendVoiceParams {
	p.ParseMode = parseMode
	return p
}

// WithCaptionEntities adds caption entities parameter
func (p *SendVoiceParams) WithCaptionEntities(captionEntities ...MessageEntity) *SendVoiceParams {
	p.CaptionEntities = captionEntities
	return p
}

// WithDuration adds duration parameter
func (p *SendVoiceParams) WithDuration(duration int) *SendVoiceParams {
	p.Duration = duration
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendVoiceParams) WithDisableNotification() *SendVoiceParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendVoiceParams) WithProtectContent() *SendVoiceParams {
	p.ProtectContent = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendVoiceParams) WithMessageEffectID(messageEffectID string) *SendVoiceParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithReplyParameters adds reply parameters parameter
func (p *SendVoiceParams) WithReplyParameters(replyParameters *ReplyParameters) *SendVoiceParams {
	p.ReplyParameters = replyParameters
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendVoiceParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendVoiceParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *SendVideoNoteParams) WithBusinessConnectionID(businessConnectionID string) *SendVideoNoteParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithChatID adds chat ID parameter
func (p *SendVideoNoteParams) WithChatID(chatID ChatID) *SendVideoNoteParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *SendVideoNoteParams) WithMessageThreadID(messageThreadID int) *SendVideoNoteParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithVideoNote adds video note parameter
func (p *SendVideoNoteParams) WithVideoNote(videoNote InputFile) *SendVideoNoteParams {
	p.VideoNote = videoNote
	return p
}

// WithDuration adds duration parameter
func (p *SendVideoNoteParams) WithDuration(duration int) *SendVideoNoteParams {
	p.Duration = duration
	return p
}

// WithLength adds length parameter
func (p *SendVideoNoteParams) WithLength(length int) *SendVideoNoteParams {
	p.Length = length
	return p
}

// WithThumbnail adds thumbnail parameter
func (p *SendVideoNoteParams) WithThumbnail(thumbnail *InputFile) *SendVideoNoteParams {
	p.Thumbnail = thumbnail
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendVideoNoteParams) WithDisableNotification() *SendVideoNoteParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendVideoNoteParams) WithProtectContent() *SendVideoNoteParams {
	p.ProtectContent = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendVideoNoteParams) WithMessageEffectID(messageEffectID string) *SendVideoNoteParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithReplyParameters adds reply parameters parameter
func (p *SendVideoNoteParams) WithReplyParameters(replyParameters *ReplyParameters) *SendVideoNoteParams {
	p.ReplyParameters = replyParameters
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendVideoNoteParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendVideoNoteParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *SendMediaGroupParams) WithBusinessConnectionID(businessConnectionID string) *SendMediaGroupParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithChatID adds chat ID parameter
func (p *SendMediaGroupParams) WithChatID(chatID ChatID) *SendMediaGroupParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *SendMediaGroupParams) WithMessageThreadID(messageThreadID int) *SendMediaGroupParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithMedia adds media parameter
func (p *SendMediaGroupParams) WithMedia(media ...InputMedia) *SendMediaGroupParams {
	p.Media = media
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendMediaGroupParams) WithDisableNotification() *SendMediaGroupParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendMediaGroupParams) WithProtectContent() *SendMediaGroupParams {
	p.ProtectContent = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendMediaGroupParams) WithMessageEffectID(messageEffectID string) *SendMediaGroupParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithReplyParameters adds reply parameters parameter
func (p *SendMediaGroupParams) WithReplyParameters(replyParameters *ReplyParameters) *SendMediaGroupParams {
	p.ReplyParameters = replyParameters
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *SendLocationParams) WithBusinessConnectionID(businessConnectionID string) *SendLocationParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithChatID adds chat ID parameter
func (p *SendLocationParams) WithChatID(chatID ChatID) *SendLocationParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *SendLocationParams) WithMessageThreadID(messageThreadID int) *SendLocationParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithLivePeriod adds live period parameter
func (p *SendLocationParams) WithLivePeriod(livePeriod int) *SendLocationParams {
	p.LivePeriod = livePeriod
	return p
}

// WithHeading adds heading parameter
func (p *SendLocationParams) WithHeading(heading int) *SendLocationParams {
	p.Heading = heading
	return p
}

// WithProximityAlertRadius adds proximity alert radius parameter
func (p *SendLocationParams) WithProximityAlertRadius(proximityAlertRadius int) *SendLocationParams {
	p.ProximityAlertRadius = proximityAlertRadius
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendLocationParams) WithDisableNotification() *SendLocationParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendLocationParams) WithProtectContent() *SendLocationParams {
	p.ProtectContent = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendLocationParams) WithMessageEffectID(messageEffectID string) *SendLocationParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithReplyParameters adds reply parameters parameter
func (p *SendLocationParams) WithReplyParameters(replyParameters *ReplyParameters) *SendLocationParams {
	p.ReplyParameters = replyParameters
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendLocationParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendLocationParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *SendVenueParams) WithBusinessConnectionID(businessConnectionID string) *SendVenueParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithChatID adds chat ID parameter
func (p *SendVenueParams) WithChatID(chatID ChatID) *SendVenueParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *SendVenueParams) WithMessageThreadID(messageThreadID int) *SendVenueParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithTitle adds title parameter
func (p *SendVenueParams) WithTitle(title string) *SendVenueParams {
	p.Title = title
	return p
}

// WithAddress adds address parameter
func (p *SendVenueParams) WithAddress(address string) *SendVenueParams {
	p.Address = address
	return p
}

// WithFoursquareID adds foursquare ID parameter
func (p *SendVenueParams) WithFoursquareID(foursquareID string) *SendVenueParams {
	p.FoursquareID = foursquareID
	return p
}

// WithFoursquareType adds foursquare type parameter
func (p *SendVenueParams) WithFoursquareType(foursquareType string) *SendVenueParams {
	p.FoursquareType = foursquareType
	return p
}

// WithGooglePlaceID adds google place ID parameter
func (p *SendVenueParams) WithGooglePlaceID(googlePlaceID string) *SendVenueParams {
	p.GooglePlaceID = googlePlaceID
	return p
}

// WithGooglePlaceType adds google place type parameter
func (p *SendVenueParams) WithGooglePlaceType(googlePlaceType string) *SendVenueParams {
	p.GooglePlaceType = googlePlaceType
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendVenueParams) WithDisableNotification() *SendVenueParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendVenueParams) WithProtectContent() *SendVenueParams {
	p.ProtectContent = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendVenueParams) WithMessageEffectID(messageEffectID string) *SendVenueParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithReplyParameters adds reply parameters parameter
func (p *SendVenueParams) WithReplyParameters(replyParameters *ReplyParameters) *SendVenueParams {
	p.ReplyParameters = replyParameters
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendVenueParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendVenueParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *SendContactParams) WithBusinessConnectionID(businessConnectionID string) *SendContactParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithChatID adds chat ID parameter
func (p *SendContactParams) WithChatID(chatID ChatID) *SendContactParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *SendContactParams) WithMessageThreadID(messageThreadID int) *SendContactParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithPhoneNumber adds phone number parameter
func (p *SendContactParams) WithPhoneNumber(phoneNumber string) *SendContactParams {
	p.PhoneNumber = phoneNumber
	return p
}

// WithFirstName adds first name parameter
func (p *SendContactParams) WithFirstName(firstName string) *SendContactParams {
	p.FirstName = firstName
	return p
}

// WithLastName adds last name parameter
func (p *SendContactParams) WithLastName(lastName string) *SendContactParams {
	p.LastName = lastName
	return p
}

// WithVcard adds vcard parameter
func (p *SendContactParams) WithVcard(vcard string) *SendContactParams {
	p.Vcard = vcard
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendContactParams) WithDisableNotification() *SendContactParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendContactParams) WithProtectContent() *SendContactParams {
	p.ProtectContent = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendContactParams) WithMessageEffectID(messageEffectID string) *SendContactParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithReplyParameters adds reply parameters parameter
func (p *SendContactParams) WithReplyParameters(replyParameters *ReplyParameters) *SendContactParams {
	p.ReplyParameters = replyParameters
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendContactParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendContactParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *SendPollParams) WithBusinessConnectionID(businessConnectionID string) *SendPollParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithChatID adds chat ID parameter
func (p *SendPollParams) WithChatID(chatID ChatID) *SendPollParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *SendPollParams) WithMessageThreadID(messageThreadID int) *SendPollParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithQuestion adds question parameter
func (p *SendPollParams) WithQuestion(question string) *SendPollParams {
	p.Question = question
	return p
}

// WithQuestionParseMode adds question parse mode parameter
func (p *SendPollParams) WithQuestionParseMode(questionParseMode string) *SendPollParams {
	p.QuestionParseMode = questionParseMode
	return p
}

// WithQuestionEntities adds question entities parameter
func (p *SendPollParams) WithQuestionEntities(questionEntities ...MessageEntity) *SendPollParams {
	p.QuestionEntities = questionEntities
	return p
}

// WithOptions adds options parameter
func (p *SendPollParams) WithOptions(options ...InputPollOption) *SendPollParams {
	p.Options = options
	return p
}

// WithIsAnonymous adds is anonymous parameter
func (p *SendPollParams) WithIsAnonymous(isAnonymous bool) *SendPollParams {
	p.IsAnonymous = ToPtr(isAnonymous)
	return p
}

// WithType adds type parameter
func (p *SendPollParams) WithType(pollType string) *SendPollParams {
	p.Type = pollType
	return p
}

// WithAllowsMultipleAnswers adds allows multiple answers parameter
func (p *SendPollParams) WithAllowsMultipleAnswers() *SendPollParams {
	p.AllowsMultipleAnswers = true
	return p
}

// WithCorrectOptionID adds correct option ID parameter
func (p *SendPollParams) WithCorrectOptionID(correctOptionID int) *SendPollParams {
	p.CorrectOptionID = ToPtr(correctOptionID)
	return p
}

// WithExplanation adds explanation parameter
func (p *SendPollParams) WithExplanation(explanation string) *SendPollParams {
	p.Explanation = explanation
	return p
}

// WithExplanationParseMode adds explanation parse mode parameter
func (p *SendPollParams) WithExplanationParseMode(explanationParseMode string) *SendPollParams {
	p.ExplanationParseMode = explanationParseMode
	return p
}

// WithExplanationEntities adds explanation entities parameter
func (p *SendPollParams) WithExplanationEntities(explanationEntities ...MessageEntity) *SendPollParams {
	p.ExplanationEntities = explanationEntities
	return p
}

// WithOpenPeriod adds open period parameter
func (p *SendPollParams) WithOpenPeriod(openPeriod int) *SendPollParams {
	p.OpenPeriod = openPeriod
	return p
}

// WithIsClosed adds is closed parameter
func (p *SendPollParams) WithIsClosed() *SendPollParams {
	p.IsClosed = true
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendPollParams) WithDisableNotification() *SendPollParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendPollParams) WithProtectContent() *SendPollParams {
	p.ProtectContent = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendPollParams) WithMessageEffectID(messageEffectID string) *SendPollParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithReplyParameters adds reply parameters parameter
func (p *SendPollParams) WithReplyParameters(replyParameters *ReplyParameters) *SendPollParams {
	p.ReplyParameters = replyParameters
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendPollParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendPollParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *SendDiceParams) WithBusinessConnectionID(businessConnectionID string) *SendDiceParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithChatID adds chat ID parameter
func (p *SendDiceParams) WithChatID(chatID ChatID) *SendDiceParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *SendDiceParams) WithMessageThreadID(messageThreadID int) *SendDiceParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithEmoji adds emoji parameter
func (p *SendDiceParams) WithEmoji(emoji string) *SendDiceParams {
	p.Emoji = emoji
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendDiceParams) WithDisableNotification() *SendDiceParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendDiceParams) WithProtectContent() *SendDiceParams {
	p.ProtectContent = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendDiceParams) WithMessageEffectID(messageEffectID string) *SendDiceParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithReplyParameters adds reply parameters parameter
func (p *SendDiceParams) WithReplyParameters(replyParameters *ReplyParameters) *SendDiceParams {
	p.ReplyParameters = replyParameters
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendDiceParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendDiceParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *SendChatActionParams) WithBusinessConnectionID(businessConnectionID string) *SendChatActionParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithChatID adds chat ID parameter
func (p *SendChatActionParams) WithChatID(chatID ChatID) *SendChatActionParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *SendChatActionParams) WithMessageThreadID(messageThreadID int) *SendChatActionParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithAction adds action parameter
func (p *SendChatActionParams) WithAction(action string) *SendChatActionParams {
	p.Action = action
	return p
}

// WithChatID adds chat ID parameter
func (p *SetMessageReactionParams) WithChatID(chatID ChatID) *SetMessageReactionParams {
	p.ChatID = chatID
	return p
}

// WithMessageID adds message ID parameter
func (p *SetMessageReactionParams) WithMessageID(messageID int) *SetMessageReactionParams {
	p.MessageID = messageID
	return p
}

// WithReaction adds reaction parameter
func (p *SetMessageReactionParams) WithReaction(reaction ...ReactionType) *SetMessageReactionParams {
	p.Reaction = reaction
	return p
}

// WithIsBig adds is big parameter
func (p *SetMessageReactionParams) WithIsBig() *SetMessageReactionParams {
	p.IsBig = true
	return p
}

// WithOffset adds offset parameter
func (p *GetUserProfilePhotosParams) WithOffset(offset int) *GetUserProfilePhotosParams {
	p.Offset = offset
	return p
}

// WithLimit adds limit parameter
func (p *GetUserProfilePhotosParams) WithLimit(limit int) *GetUserProfilePhotosParams {
	p.Limit = limit
	return p
}

// WithFileID adds file ID parameter
func (p *GetFileParams) WithFileID(fileID string) *GetFileParams {
	p.FileID = fileID
	return p
}

// WithChatID adds chat ID parameter
func (p *BanChatMemberParams) WithChatID(chatID ChatID) *BanChatMemberParams {
	p.ChatID = chatID
	return p
}

// WithRevokeMessages adds revoke messages parameter
func (p *BanChatMemberParams) WithRevokeMessages() *BanChatMemberParams {
	p.RevokeMessages = true
	return p
}

// WithChatID adds chat ID parameter
func (p *UnbanChatMemberParams) WithChatID(chatID ChatID) *UnbanChatMemberParams {
	p.ChatID = chatID
	return p
}

// WithOnlyIfBanned adds only if banned parameter
func (p *UnbanChatMemberParams) WithOnlyIfBanned() *UnbanChatMemberParams {
	p.OnlyIfBanned = true
	return p
}

// WithChatID adds chat ID parameter
func (p *RestrictChatMemberParams) WithChatID(chatID ChatID) *RestrictChatMemberParams {
	p.ChatID = chatID
	return p
}

// WithPermissions adds permissions parameter
func (p *RestrictChatMemberParams) WithPermissions(permissions ChatPermissions) *RestrictChatMemberParams {
	p.Permissions = permissions
	return p
}

// WithUseIndependentChatPermissions adds use independent chat permissions parameter
func (p *RestrictChatMemberParams) WithUseIndependentChatPermissions() *RestrictChatMemberParams {
	p.UseIndependentChatPermissions = true
	return p
}

// WithChatID adds chat ID parameter
func (p *PromoteChatMemberParams) WithChatID(chatID ChatID) *PromoteChatMemberParams {
	p.ChatID = chatID
	return p
}

// WithIsAnonymous adds is anonymous parameter
func (p *PromoteChatMemberParams) WithIsAnonymous(isAnonymous bool) *PromoteChatMemberParams {
	p.IsAnonymous = ToPtr(isAnonymous)
	return p
}

// WithCanManageChat adds can manage chat parameter
func (p *PromoteChatMemberParams) WithCanManageChat(canManageChat bool) *PromoteChatMemberParams {
	p.CanManageChat = ToPtr(canManageChat)
	return p
}

// WithCanDeleteMessages adds can delete messages parameter
func (p *PromoteChatMemberParams) WithCanDeleteMessages(canDeleteMessages bool) *PromoteChatMemberParams {
	p.CanDeleteMessages = ToPtr(canDeleteMessages)
	return p
}

// WithCanManageVideoChats adds can manage video chats parameter
func (p *PromoteChatMemberParams) WithCanManageVideoChats(canManageVideoChats bool) *PromoteChatMemberParams {
	p.CanManageVideoChats = ToPtr(canManageVideoChats)
	return p
}

// WithCanRestrictMembers adds can restrict members parameter
func (p *PromoteChatMemberParams) WithCanRestrictMembers(canRestrictMembers bool) *PromoteChatMemberParams {
	p.CanRestrictMembers = ToPtr(canRestrictMembers)
	return p
}

// WithCanPromoteMembers adds can promote members parameter
func (p *PromoteChatMemberParams) WithCanPromoteMembers(canPromoteMembers bool) *PromoteChatMemberParams {
	p.CanPromoteMembers = ToPtr(canPromoteMembers)
	return p
}

// WithCanChangeInfo adds can change info parameter
func (p *PromoteChatMemberParams) WithCanChangeInfo(canChangeInfo bool) *PromoteChatMemberParams {
	p.CanChangeInfo = ToPtr(canChangeInfo)
	return p
}

// WithCanInviteUsers adds can invite users parameter
func (p *PromoteChatMemberParams) WithCanInviteUsers(canInviteUsers bool) *PromoteChatMemberParams {
	p.CanInviteUsers = ToPtr(canInviteUsers)
	return p
}

// WithCanPostStories adds can post stories parameter
func (p *PromoteChatMemberParams) WithCanPostStories(canPostStories bool) *PromoteChatMemberParams {
	p.CanPostStories = ToPtr(canPostStories)
	return p
}

// WithCanEditStories adds can edit stories parameter
func (p *PromoteChatMemberParams) WithCanEditStories(canEditStories bool) *PromoteChatMemberParams {
	p.CanEditStories = ToPtr(canEditStories)
	return p
}

// WithCanDeleteStories adds can delete stories parameter
func (p *PromoteChatMemberParams) WithCanDeleteStories(canDeleteStories bool) *PromoteChatMemberParams {
	p.CanDeleteStories = ToPtr(canDeleteStories)
	return p
}

// WithCanPostMessages adds can post messages parameter
func (p *PromoteChatMemberParams) WithCanPostMessages(canPostMessages bool) *PromoteChatMemberParams {
	p.CanPostMessages = ToPtr(canPostMessages)
	return p
}

// WithCanEditMessages adds can edit messages parameter
func (p *PromoteChatMemberParams) WithCanEditMessages(canEditMessages bool) *PromoteChatMemberParams {
	p.CanEditMessages = ToPtr(canEditMessages)
	return p
}

// WithCanPinMessages adds can pin messages parameter
func (p *PromoteChatMemberParams) WithCanPinMessages(canPinMessages bool) *PromoteChatMemberParams {
	p.CanPinMessages = ToPtr(canPinMessages)
	return p
}

// WithCanManageTopics adds can manage topics parameter
func (p *PromoteChatMemberParams) WithCanManageTopics(canManageTopics bool) *PromoteChatMemberParams {
	p.CanManageTopics = ToPtr(canManageTopics)
	return p
}

// WithChatID adds chat ID parameter
func (p *SetChatAdministratorCustomTitleParams) WithChatID(chatID ChatID) *SetChatAdministratorCustomTitleParams {
	p.ChatID = chatID
	return p
}

// WithCustomTitle adds custom title parameter
func (p *SetChatAdministratorCustomTitleParams) WithCustomTitle(customTitle string,
) *SetChatAdministratorCustomTitleParams {
	p.CustomTitle = customTitle
	return p
}

// WithChatID adds chat ID parameter
func (p *BanChatSenderChatParams) WithChatID(chatID ChatID) *BanChatSenderChatParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *UnbanChatSenderChatParams) WithChatID(chatID ChatID) *UnbanChatSenderChatParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *SetChatPermissionsParams) WithChatID(chatID ChatID) *SetChatPermissionsParams {
	p.ChatID = chatID
	return p
}

// WithPermissions adds permissions parameter
func (p *SetChatPermissionsParams) WithPermissions(permissions ChatPermissions) *SetChatPermissionsParams {
	p.Permissions = permissions
	return p
}

// WithUseIndependentChatPermissions adds use independent chat permissions parameter
func (p *SetChatPermissionsParams) WithUseIndependentChatPermissions() *SetChatPermissionsParams {
	p.UseIndependentChatPermissions = true
	return p
}

// WithChatID adds chat ID parameter
func (p *ExportChatInviteLinkParams) WithChatID(chatID ChatID) *ExportChatInviteLinkParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *CreateChatInviteLinkParams) WithChatID(chatID ChatID) *CreateChatInviteLinkParams {
	p.ChatID = chatID
	return p
}

// WithName adds name parameter
func (p *CreateChatInviteLinkParams) WithName(name string) *CreateChatInviteLinkParams {
	p.Name = name
	return p
}

// WithMemberLimit adds member limit parameter
func (p *CreateChatInviteLinkParams) WithMemberLimit(memberLimit int) *CreateChatInviteLinkParams {
	p.MemberLimit = memberLimit
	return p
}

// WithCreatesJoinRequest adds creates join request parameter
func (p *CreateChatInviteLinkParams) WithCreatesJoinRequest() *CreateChatInviteLinkParams {
	p.CreatesJoinRequest = true
	return p
}

// WithChatID adds chat ID parameter
func (p *EditChatInviteLinkParams) WithChatID(chatID ChatID) *EditChatInviteLinkParams {
	p.ChatID = chatID
	return p
}

// WithInviteLink adds invite link parameter
func (p *EditChatInviteLinkParams) WithInviteLink(inviteLink string) *EditChatInviteLinkParams {
	p.InviteLink = inviteLink
	return p
}

// WithName adds name parameter
func (p *EditChatInviteLinkParams) WithName(name string) *EditChatInviteLinkParams {
	p.Name = name
	return p
}

// WithMemberLimit adds member limit parameter
func (p *EditChatInviteLinkParams) WithMemberLimit(memberLimit int) *EditChatInviteLinkParams {
	p.MemberLimit = memberLimit
	return p
}

// WithCreatesJoinRequest adds creates join request parameter
func (p *EditChatInviteLinkParams) WithCreatesJoinRequest() *EditChatInviteLinkParams {
	p.CreatesJoinRequest = true
	return p
}

// WithChatID adds chat ID parameter
func (p *RevokeChatInviteLinkParams) WithChatID(chatID ChatID) *RevokeChatInviteLinkParams {
	p.ChatID = chatID
	return p
}

// WithInviteLink adds invite link parameter
func (p *RevokeChatInviteLinkParams) WithInviteLink(inviteLink string) *RevokeChatInviteLinkParams {
	p.InviteLink = inviteLink
	return p
}

// WithChatID adds chat ID parameter
func (p *ApproveChatJoinRequestParams) WithChatID(chatID ChatID) *ApproveChatJoinRequestParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *DeclineChatJoinRequestParams) WithChatID(chatID ChatID) *DeclineChatJoinRequestParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *SetChatPhotoParams) WithChatID(chatID ChatID) *SetChatPhotoParams {
	p.ChatID = chatID
	return p
}

// WithPhoto adds photo parameter
func (p *SetChatPhotoParams) WithPhoto(photo InputFile) *SetChatPhotoParams {
	p.Photo = photo
	return p
}

// WithChatID adds chat ID parameter
func (p *DeleteChatPhotoParams) WithChatID(chatID ChatID) *DeleteChatPhotoParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *SetChatTitleParams) WithChatID(chatID ChatID) *SetChatTitleParams {
	p.ChatID = chatID
	return p
}

// WithTitle adds title parameter
func (p *SetChatTitleParams) WithTitle(title string) *SetChatTitleParams {
	p.Title = title
	return p
}

// WithChatID adds chat ID parameter
func (p *SetChatDescriptionParams) WithChatID(chatID ChatID) *SetChatDescriptionParams {
	p.ChatID = chatID
	return p
}

// WithDescription adds description parameter
func (p *SetChatDescriptionParams) WithDescription(description string) *SetChatDescriptionParams {
	p.Description = description
	return p
}

// WithChatID adds chat ID parameter
func (p *PinChatMessageParams) WithChatID(chatID ChatID) *PinChatMessageParams {
	p.ChatID = chatID
	return p
}

// WithMessageID adds message ID parameter
func (p *PinChatMessageParams) WithMessageID(messageID int) *PinChatMessageParams {
	p.MessageID = messageID
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *PinChatMessageParams) WithDisableNotification() *PinChatMessageParams {
	p.DisableNotification = true
	return p
}

// WithChatID adds chat ID parameter
func (p *UnpinChatMessageParams) WithChatID(chatID ChatID) *UnpinChatMessageParams {
	p.ChatID = chatID
	return p
}

// WithMessageID adds message ID parameter
func (p *UnpinChatMessageParams) WithMessageID(messageID int) *UnpinChatMessageParams {
	p.MessageID = messageID
	return p
}

// WithChatID adds chat ID parameter
func (p *UnpinAllChatMessagesParams) WithChatID(chatID ChatID) *UnpinAllChatMessagesParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *LeaveChatParams) WithChatID(chatID ChatID) *LeaveChatParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *GetChatParams) WithChatID(chatID ChatID) *GetChatParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *GetChatAdministratorsParams) WithChatID(chatID ChatID) *GetChatAdministratorsParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *GetChatMemberCountParams) WithChatID(chatID ChatID) *GetChatMemberCountParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *GetChatMemberParams) WithChatID(chatID ChatID) *GetChatMemberParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *SetChatStickerSetParams) WithChatID(chatID ChatID) *SetChatStickerSetParams {
	p.ChatID = chatID
	return p
}

// WithStickerSetName adds sticker set name parameter
func (p *SetChatStickerSetParams) WithStickerSetName(stickerSetName string) *SetChatStickerSetParams {
	p.StickerSetName = stickerSetName
	return p
}

// WithChatID adds chat ID parameter
func (p *DeleteChatStickerSetParams) WithChatID(chatID ChatID) *DeleteChatStickerSetParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *CreateForumTopicParams) WithChatID(chatID ChatID) *CreateForumTopicParams {
	p.ChatID = chatID
	return p
}

// WithName adds name parameter
func (p *CreateForumTopicParams) WithName(name string) *CreateForumTopicParams {
	p.Name = name
	return p
}

// WithIconColor adds icon color parameter
func (p *CreateForumTopicParams) WithIconColor(iconColor int) *CreateForumTopicParams {
	p.IconColor = iconColor
	return p
}

// WithIconCustomEmojiID adds icon custom emoji ID parameter
func (p *CreateForumTopicParams) WithIconCustomEmojiID(iconCustomEmojiID string) *CreateForumTopicParams {
	p.IconCustomEmojiID = iconCustomEmojiID
	return p
}

// WithChatID adds chat ID parameter
func (p *EditForumTopicParams) WithChatID(chatID ChatID) *EditForumTopicParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *EditForumTopicParams) WithMessageThreadID(messageThreadID int) *EditForumTopicParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithName adds name parameter
func (p *EditForumTopicParams) WithName(name string) *EditForumTopicParams {
	p.Name = name
	return p
}

// WithIconCustomEmojiID adds icon custom emoji ID parameter
func (p *EditForumTopicParams) WithIconCustomEmojiID(iconCustomEmojiID string) *EditForumTopicParams {
	p.IconCustomEmojiID = ToPtr(iconCustomEmojiID)
	return p
}

// WithChatID adds chat ID parameter
func (p *CloseForumTopicParams) WithChatID(chatID ChatID) *CloseForumTopicParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *CloseForumTopicParams) WithMessageThreadID(messageThreadID int) *CloseForumTopicParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithChatID adds chat ID parameter
func (p *ReopenForumTopicParams) WithChatID(chatID ChatID) *ReopenForumTopicParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *ReopenForumTopicParams) WithMessageThreadID(messageThreadID int) *ReopenForumTopicParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithChatID adds chat ID parameter
func (p *DeleteForumTopicParams) WithChatID(chatID ChatID) *DeleteForumTopicParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *DeleteForumTopicParams) WithMessageThreadID(messageThreadID int) *DeleteForumTopicParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithChatID adds chat ID parameter
func (p *UnpinAllForumTopicMessagesParams) WithChatID(chatID ChatID) *UnpinAllForumTopicMessagesParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *UnpinAllForumTopicMessagesParams) WithMessageThreadID(messageThreadID int) *UnpinAllForumTopicMessagesParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithChatID adds chat ID parameter
func (p *EditGeneralForumTopicParams) WithChatID(chatID ChatID) *EditGeneralForumTopicParams {
	p.ChatID = chatID
	return p
}

// WithName adds name parameter
func (p *EditGeneralForumTopicParams) WithName(name string) *EditGeneralForumTopicParams {
	p.Name = name
	return p
}

// WithChatID adds chat ID parameter
func (p *CloseGeneralForumTopicParams) WithChatID(chatID ChatID) *CloseGeneralForumTopicParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *ReopenGeneralForumTopicParams) WithChatID(chatID ChatID) *ReopenGeneralForumTopicParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *HideGeneralForumTopicParams) WithChatID(chatID ChatID) *HideGeneralForumTopicParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *UnhideGeneralForumTopicParams) WithChatID(chatID ChatID) *UnhideGeneralForumTopicParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *UnpinAllGeneralForumTopicMessagesParams) WithChatID(chatID ChatID) *UnpinAllGeneralForumTopicMessagesParams {
	p.ChatID = chatID
	return p
}

// WithCallbackQueryID adds callback query ID parameter
func (p *AnswerCallbackQueryParams) WithCallbackQueryID(callbackQueryID string) *AnswerCallbackQueryParams {
	p.CallbackQueryID = callbackQueryID
	return p
}

// WithText adds text parameter
func (p *AnswerCallbackQueryParams) WithText(text string) *AnswerCallbackQueryParams {
	p.Text = text
	return p
}

// WithShowAlert adds show alert parameter
func (p *AnswerCallbackQueryParams) WithShowAlert() *AnswerCallbackQueryParams {
	p.ShowAlert = true
	return p
}

// WithURL adds URL parameter
func (p *AnswerCallbackQueryParams) WithURL(url string) *AnswerCallbackQueryParams {
	p.URL = url
	return p
}

// WithCacheTime adds cache time parameter
func (p *AnswerCallbackQueryParams) WithCacheTime(cacheTime int) *AnswerCallbackQueryParams {
	p.CacheTime = cacheTime
	return p
}

// WithChatID adds chat ID parameter
func (p *GetUserChatBoostsParams) WithChatID(chatID ChatID) *GetUserChatBoostsParams {
	p.ChatID = chatID
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *GetBusinessConnectionParams) WithBusinessConnectionID(businessConnectionID string,
) *GetBusinessConnectionParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithCommands adds commands parameter
func (p *SetMyCommandsParams) WithCommands(commands ...BotCommand) *SetMyCommandsParams {
	p.Commands = commands
	return p
}

// WithScope adds scope parameter
func (p *SetMyCommandsParams) WithScope(scope BotCommandScope) *SetMyCommandsParams {
	p.Scope = scope
	return p
}

// WithLanguageCode adds language code parameter
func (p *SetMyCommandsParams) WithLanguageCode(languageCode string) *SetMyCommandsParams {
	p.LanguageCode = languageCode
	return p
}

// WithScope adds scope parameter
func (p *DeleteMyCommandsParams) WithScope(scope BotCommandScope) *DeleteMyCommandsParams {
	p.Scope = scope
	return p
}

// WithLanguageCode adds language code parameter
func (p *DeleteMyCommandsParams) WithLanguageCode(languageCode string) *DeleteMyCommandsParams {
	p.LanguageCode = languageCode
	return p
}

// WithScope adds scope parameter
func (p *GetMyCommandsParams) WithScope(scope BotCommandScope) *GetMyCommandsParams {
	p.Scope = scope
	return p
}

// WithLanguageCode adds language code parameter
func (p *GetMyCommandsParams) WithLanguageCode(languageCode string) *GetMyCommandsParams {
	p.LanguageCode = languageCode
	return p
}

// WithName adds name parameter
func (p *SetMyNameParams) WithName(name string) *SetMyNameParams {
	p.Name = name
	return p
}

// WithLanguageCode adds language code parameter
func (p *SetMyNameParams) WithLanguageCode(languageCode string) *SetMyNameParams {
	p.LanguageCode = languageCode
	return p
}

// WithLanguageCode adds language code parameter
func (p *GetMyNameParams) WithLanguageCode(languageCode string) *GetMyNameParams {
	p.LanguageCode = languageCode
	return p
}

// WithDescription adds description parameter
func (p *SetMyDescriptionParams) WithDescription(description string) *SetMyDescriptionParams {
	p.Description = description
	return p
}

// WithLanguageCode adds language code parameter
func (p *SetMyDescriptionParams) WithLanguageCode(languageCode string) *SetMyDescriptionParams {
	p.LanguageCode = languageCode
	return p
}

// WithLanguageCode adds language code parameter
func (p *GetMyDescriptionParams) WithLanguageCode(languageCode string) *GetMyDescriptionParams {
	p.LanguageCode = languageCode
	return p
}

// WithShortDescription adds short description parameter
func (p *SetMyShortDescriptionParams) WithShortDescription(shortDescription string) *SetMyShortDescriptionParams {
	p.ShortDescription = shortDescription
	return p
}

// WithLanguageCode adds language code parameter
func (p *SetMyShortDescriptionParams) WithLanguageCode(languageCode string) *SetMyShortDescriptionParams {
	p.LanguageCode = languageCode
	return p
}

// WithLanguageCode adds language code parameter
func (p *GetMyShortDescriptionParams) WithLanguageCode(languageCode string) *GetMyShortDescriptionParams {
	p.LanguageCode = languageCode
	return p
}

// WithMenuButton adds menu button parameter
func (p *SetChatMenuButtonParams) WithMenuButton(menuButton MenuButton) *SetChatMenuButtonParams {
	p.MenuButton = menuButton
	return p
}

// WithRights adds rights parameter
func (p *SetMyDefaultAdministratorRightsParams) WithRights(rights *ChatAdministratorRights,
) *SetMyDefaultAdministratorRightsParams {
	p.Rights = rights
	return p
}

// WithForChannels adds for channels parameter
func (p *SetMyDefaultAdministratorRightsParams) WithForChannels() *SetMyDefaultAdministratorRightsParams {
	p.ForChannels = true
	return p
}

// WithForChannels adds for channels parameter
func (p *GetMyDefaultAdministratorRightsParams) WithForChannels() *GetMyDefaultAdministratorRightsParams {
	p.ForChannels = true
	return p
}

// WithChatID adds chat ID parameter
func (p *EditMessageTextParams) WithChatID(chatID ChatID) *EditMessageTextParams {
	p.ChatID = chatID
	return p
}

// WithMessageID adds message ID parameter
func (p *EditMessageTextParams) WithMessageID(messageID int) *EditMessageTextParams {
	p.MessageID = messageID
	return p
}

// WithInlineMessageID adds inline message ID parameter
func (p *EditMessageTextParams) WithInlineMessageID(inlineMessageID string) *EditMessageTextParams {
	p.InlineMessageID = inlineMessageID
	return p
}

// WithText adds text parameter
func (p *EditMessageTextParams) WithText(text string) *EditMessageTextParams {
	p.Text = text
	return p
}

// WithParseMode adds parse mode parameter
func (p *EditMessageTextParams) WithParseMode(parseMode string) *EditMessageTextParams {
	p.ParseMode = parseMode
	return p
}

// WithEntities adds entities parameter
func (p *EditMessageTextParams) WithEntities(entities ...MessageEntity) *EditMessageTextParams {
	p.Entities = entities
	return p
}

// WithLinkPreviewOptions adds link preview options parameter
func (p *EditMessageTextParams) WithLinkPreviewOptions(linkPreviewOptions *LinkPreviewOptions) *EditMessageTextParams {
	p.LinkPreviewOptions = linkPreviewOptions
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *EditMessageTextParams) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *EditMessageTextParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *EditMessageCaptionParams) WithChatID(chatID ChatID) *EditMessageCaptionParams {
	p.ChatID = chatID
	return p
}

// WithMessageID adds message ID parameter
func (p *EditMessageCaptionParams) WithMessageID(messageID int) *EditMessageCaptionParams {
	p.MessageID = messageID
	return p
}

// WithInlineMessageID adds inline message ID parameter
func (p *EditMessageCaptionParams) WithInlineMessageID(inlineMessageID string) *EditMessageCaptionParams {
	p.InlineMessageID = inlineMessageID
	return p
}

// WithCaption adds caption parameter
func (p *EditMessageCaptionParams) WithCaption(caption string) *EditMessageCaptionParams {
	p.Caption = caption
	return p
}

// WithParseMode adds parse mode parameter
func (p *EditMessageCaptionParams) WithParseMode(parseMode string) *EditMessageCaptionParams {
	p.ParseMode = parseMode
	return p
}

// WithCaptionEntities adds caption entities parameter
func (p *EditMessageCaptionParams) WithCaptionEntities(captionEntities ...MessageEntity) *EditMessageCaptionParams {
	p.CaptionEntities = captionEntities
	return p
}

// WithShowCaptionAboveMedia adds show caption above media parameter
func (p *EditMessageCaptionParams) WithShowCaptionAboveMedia() *EditMessageCaptionParams {
	p.ShowCaptionAboveMedia = true
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *EditMessageCaptionParams) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *EditMessageCaptionParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *EditMessageMediaParams) WithChatID(chatID ChatID) *EditMessageMediaParams {
	p.ChatID = chatID
	return p
}

// WithMessageID adds message ID parameter
func (p *EditMessageMediaParams) WithMessageID(messageID int) *EditMessageMediaParams {
	p.MessageID = messageID
	return p
}

// WithInlineMessageID adds inline message ID parameter
func (p *EditMessageMediaParams) WithInlineMessageID(inlineMessageID string) *EditMessageMediaParams {
	p.InlineMessageID = inlineMessageID
	return p
}

// WithMedia adds media parameter
func (p *EditMessageMediaParams) WithMedia(media InputMedia) *EditMessageMediaParams {
	p.Media = media
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *EditMessageMediaParams) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *EditMessageMediaParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *EditMessageLiveLocationParams) WithChatID(chatID ChatID) *EditMessageLiveLocationParams {
	p.ChatID = chatID
	return p
}

// WithMessageID adds message ID parameter
func (p *EditMessageLiveLocationParams) WithMessageID(messageID int) *EditMessageLiveLocationParams {
	p.MessageID = messageID
	return p
}

// WithInlineMessageID adds inline message ID parameter
func (p *EditMessageLiveLocationParams) WithInlineMessageID(inlineMessageID string) *EditMessageLiveLocationParams {
	p.InlineMessageID = inlineMessageID
	return p
}

// WithLivePeriod adds live period parameter
func (p *EditMessageLiveLocationParams) WithLivePeriod(livePeriod int) *EditMessageLiveLocationParams {
	p.LivePeriod = livePeriod
	return p
}

// WithHeading adds heading parameter
func (p *EditMessageLiveLocationParams) WithHeading(heading int) *EditMessageLiveLocationParams {
	p.Heading = heading
	return p
}

// WithProximityAlertRadius adds proximity alert radius parameter
func (p *EditMessageLiveLocationParams) WithProximityAlertRadius(proximityAlertRadius int,
) *EditMessageLiveLocationParams {
	p.ProximityAlertRadius = proximityAlertRadius
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *EditMessageLiveLocationParams) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup,
) *EditMessageLiveLocationParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *StopMessageLiveLocationParams) WithChatID(chatID ChatID) *StopMessageLiveLocationParams {
	p.ChatID = chatID
	return p
}

// WithMessageID adds message ID parameter
func (p *StopMessageLiveLocationParams) WithMessageID(messageID int) *StopMessageLiveLocationParams {
	p.MessageID = messageID
	return p
}

// WithInlineMessageID adds inline message ID parameter
func (p *StopMessageLiveLocationParams) WithInlineMessageID(inlineMessageID string) *StopMessageLiveLocationParams {
	p.InlineMessageID = inlineMessageID
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *StopMessageLiveLocationParams) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup,
) *StopMessageLiveLocationParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *EditMessageReplyMarkupParams) WithChatID(chatID ChatID) *EditMessageReplyMarkupParams {
	p.ChatID = chatID
	return p
}

// WithMessageID adds message ID parameter
func (p *EditMessageReplyMarkupParams) WithMessageID(messageID int) *EditMessageReplyMarkupParams {
	p.MessageID = messageID
	return p
}

// WithInlineMessageID adds inline message ID parameter
func (p *EditMessageReplyMarkupParams) WithInlineMessageID(inlineMessageID string) *EditMessageReplyMarkupParams {
	p.InlineMessageID = inlineMessageID
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *EditMessageReplyMarkupParams) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup,
) *EditMessageReplyMarkupParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *StopPollParams) WithChatID(chatID ChatID) *StopPollParams {
	p.ChatID = chatID
	return p
}

// WithMessageID adds message ID parameter
func (p *StopPollParams) WithMessageID(messageID int) *StopPollParams {
	p.MessageID = messageID
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *StopPollParams) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *StopPollParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *DeleteMessageParams) WithChatID(chatID ChatID) *DeleteMessageParams {
	p.ChatID = chatID
	return p
}

// WithMessageID adds message ID parameter
func (p *DeleteMessageParams) WithMessageID(messageID int) *DeleteMessageParams {
	p.MessageID = messageID
	return p
}

// WithChatID adds chat ID parameter
func (p *DeleteMessagesParams) WithChatID(chatID ChatID) *DeleteMessagesParams {
	p.ChatID = chatID
	return p
}

// WithMessageIDs adds message ids parameter
func (p *DeleteMessagesParams) WithMessageIDs(messageIDs ...int) *DeleteMessagesParams {
	p.MessageIDs = messageIDs
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *SendStickerParams) WithBusinessConnectionID(businessConnectionID string) *SendStickerParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithChatID adds chat ID parameter
func (p *SendStickerParams) WithChatID(chatID ChatID) *SendStickerParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *SendStickerParams) WithMessageThreadID(messageThreadID int) *SendStickerParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithSticker adds sticker parameter
func (p *SendStickerParams) WithSticker(sticker InputFile) *SendStickerParams {
	p.Sticker = sticker
	return p
}

// WithEmoji adds emoji parameter
func (p *SendStickerParams) WithEmoji(emoji string) *SendStickerParams {
	p.Emoji = emoji
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendStickerParams) WithDisableNotification() *SendStickerParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendStickerParams) WithProtectContent() *SendStickerParams {
	p.ProtectContent = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendStickerParams) WithMessageEffectID(messageEffectID string) *SendStickerParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithReplyParameters adds reply parameters parameter
func (p *SendStickerParams) WithReplyParameters(replyParameters *ReplyParameters) *SendStickerParams {
	p.ReplyParameters = replyParameters
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendStickerParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendStickerParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithName adds name parameter
func (p *GetStickerSetParams) WithName(name string) *GetStickerSetParams {
	p.Name = name
	return p
}

// WithCustomEmojiIDs adds custom emoji ids parameter
func (p *GetCustomEmojiStickersParams) WithCustomEmojiIDs(customEmojiIDs ...string) *GetCustomEmojiStickersParams {
	p.CustomEmojiIDs = customEmojiIDs
	return p
}

// WithSticker adds sticker parameter
func (p *UploadStickerFileParams) WithSticker(sticker InputFile) *UploadStickerFileParams {
	p.Sticker = sticker
	return p
}

// WithStickerFormat adds sticker format parameter
func (p *UploadStickerFileParams) WithStickerFormat(stickerFormat string) *UploadStickerFileParams {
	p.StickerFormat = stickerFormat
	return p
}

// WithName adds name parameter
func (p *CreateNewStickerSetParams) WithName(name string) *CreateNewStickerSetParams {
	p.Name = name
	return p
}

// WithTitle adds title parameter
func (p *CreateNewStickerSetParams) WithTitle(title string) *CreateNewStickerSetParams {
	p.Title = title
	return p
}

// WithStickers adds stickers parameter
func (p *CreateNewStickerSetParams) WithStickers(stickers ...InputSticker) *CreateNewStickerSetParams {
	p.Stickers = stickers
	return p
}

// WithStickerType adds sticker type parameter
func (p *CreateNewStickerSetParams) WithStickerType(stickerType string) *CreateNewStickerSetParams {
	p.StickerType = stickerType
	return p
}

// WithNeedsRepainting adds needs repainting parameter
func (p *CreateNewStickerSetParams) WithNeedsRepainting() *CreateNewStickerSetParams {
	p.NeedsRepainting = true
	return p
}

// WithName adds name parameter
func (p *AddStickerToSetParams) WithName(name string) *AddStickerToSetParams {
	p.Name = name
	return p
}

// WithSticker adds sticker parameter
func (p *AddStickerToSetParams) WithSticker(sticker InputSticker) *AddStickerToSetParams {
	p.Sticker = sticker
	return p
}

// WithSticker adds sticker parameter
func (p *SetStickerPositionInSetParams) WithSticker(sticker string) *SetStickerPositionInSetParams {
	p.Sticker = sticker
	return p
}

// WithPosition adds position parameter
func (p *SetStickerPositionInSetParams) WithPosition(position int) *SetStickerPositionInSetParams {
	p.Position = position
	return p
}

// WithSticker adds sticker parameter
func (p *DeleteStickerFromSetParams) WithSticker(sticker string) *DeleteStickerFromSetParams {
	p.Sticker = sticker
	return p
}

// WithName adds name parameter
func (p *ReplaceStickerInSetParams) WithName(name string) *ReplaceStickerInSetParams {
	p.Name = name
	return p
}

// WithOldSticker adds old sticker parameter
func (p *ReplaceStickerInSetParams) WithOldSticker(oldSticker string) *ReplaceStickerInSetParams {
	p.OldSticker = oldSticker
	return p
}

// WithSticker adds sticker parameter
func (p *ReplaceStickerInSetParams) WithSticker(sticker InputSticker) *ReplaceStickerInSetParams {
	p.Sticker = sticker
	return p
}

// WithSticker adds sticker parameter
func (p *SetStickerEmojiListParams) WithSticker(sticker string) *SetStickerEmojiListParams {
	p.Sticker = sticker
	return p
}

// WithEmojiList adds emoji list parameter
func (p *SetStickerEmojiListParams) WithEmojiList(emojiList ...string) *SetStickerEmojiListParams {
	p.EmojiList = emojiList
	return p
}

// WithSticker adds sticker parameter
func (p *SetStickerKeywordsParams) WithSticker(sticker string) *SetStickerKeywordsParams {
	p.Sticker = sticker
	return p
}

// WithKeywords adds keywords parameter
func (p *SetStickerKeywordsParams) WithKeywords(keywords ...string) *SetStickerKeywordsParams {
	p.Keywords = keywords
	return p
}

// WithSticker adds sticker parameter
func (p *SetStickerMaskPositionParams) WithSticker(sticker string) *SetStickerMaskPositionParams {
	p.Sticker = sticker
	return p
}

// WithMaskPosition adds mask position parameter
func (p *SetStickerMaskPositionParams) WithMaskPosition(maskPosition *MaskPosition) *SetStickerMaskPositionParams {
	p.MaskPosition = maskPosition
	return p
}

// WithName adds name parameter
func (p *SetStickerSetTitleParams) WithName(name string) *SetStickerSetTitleParams {
	p.Name = name
	return p
}

// WithTitle adds title parameter
func (p *SetStickerSetTitleParams) WithTitle(title string) *SetStickerSetTitleParams {
	p.Title = title
	return p
}

// WithName adds name parameter
func (p *SetStickerSetThumbnailParams) WithName(name string) *SetStickerSetThumbnailParams {
	p.Name = name
	return p
}

// WithThumbnail adds thumbnail parameter
func (p *SetStickerSetThumbnailParams) WithThumbnail(thumbnail *InputFile) *SetStickerSetThumbnailParams {
	p.Thumbnail = thumbnail
	return p
}

// WithFormat adds format parameter
func (p *SetStickerSetThumbnailParams) WithFormat(format string) *SetStickerSetThumbnailParams {
	p.Format = format
	return p
}

// WithName adds name parameter
func (p *SetCustomEmojiStickerSetThumbnailParams) WithName(name string) *SetCustomEmojiStickerSetThumbnailParams {
	p.Name = name
	return p
}

// WithCustomEmojiID adds custom emoji ID parameter
func (p *SetCustomEmojiStickerSetThumbnailParams) WithCustomEmojiID(customEmojiID string,
) *SetCustomEmojiStickerSetThumbnailParams {
	p.CustomEmojiID = customEmojiID
	return p
}

// WithName adds name parameter
func (p *DeleteStickerSetParams) WithName(name string) *DeleteStickerSetParams {
	p.Name = name
	return p
}

// WithInlineQueryID adds inline query ID parameter
func (p *AnswerInlineQueryParams) WithInlineQueryID(inlineQueryID string) *AnswerInlineQueryParams {
	p.InlineQueryID = inlineQueryID
	return p
}

// WithResults adds results parameter
func (p *AnswerInlineQueryParams) WithResults(results ...InlineQueryResult) *AnswerInlineQueryParams {
	p.Results = results
	return p
}

// WithCacheTime adds cache time parameter
func (p *AnswerInlineQueryParams) WithCacheTime(cacheTime int) *AnswerInlineQueryParams {
	p.CacheTime = cacheTime
	return p
}

// WithIsPersonal adds is personal parameter
func (p *AnswerInlineQueryParams) WithIsPersonal() *AnswerInlineQueryParams {
	p.IsPersonal = true
	return p
}

// WithNextOffset adds next offset parameter
func (p *AnswerInlineQueryParams) WithNextOffset(nextOffset string) *AnswerInlineQueryParams {
	p.NextOffset = nextOffset
	return p
}

// WithButton adds button parameter
func (p *AnswerInlineQueryParams) WithButton(button *InlineQueryResultsButton) *AnswerInlineQueryParams {
	p.Button = button
	return p
}

// WithWebAppQueryID adds web app query ID parameter
func (p *AnswerWebAppQueryParams) WithWebAppQueryID(webAppQueryID string) *AnswerWebAppQueryParams {
	p.WebAppQueryID = webAppQueryID
	return p
}

// WithResult adds result parameter
func (p *AnswerWebAppQueryParams) WithResult(result InlineQueryResult) *AnswerWebAppQueryParams {
	p.Result = result
	return p
}

// WithChatID adds chat ID parameter
func (p *SendInvoiceParams) WithChatID(chatID ChatID) *SendInvoiceParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *SendInvoiceParams) WithMessageThreadID(messageThreadID int) *SendInvoiceParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithTitle adds title parameter
func (p *SendInvoiceParams) WithTitle(title string) *SendInvoiceParams {
	p.Title = title
	return p
}

// WithDescription adds description parameter
func (p *SendInvoiceParams) WithDescription(description string) *SendInvoiceParams {
	p.Description = description
	return p
}

// WithPayload adds payload parameter
func (p *SendInvoiceParams) WithPayload(payload string) *SendInvoiceParams {
	p.Payload = payload
	return p
}

// WithProviderToken adds provider token parameter
func (p *SendInvoiceParams) WithProviderToken(providerToken string) *SendInvoiceParams {
	p.ProviderToken = providerToken
	return p
}

// WithCurrency adds currency parameter
func (p *SendInvoiceParams) WithCurrency(currency string) *SendInvoiceParams {
	p.Currency = currency
	return p
}

// WithPrices adds prices parameter
func (p *SendInvoiceParams) WithPrices(prices ...LabeledPrice) *SendInvoiceParams {
	p.Prices = prices
	return p
}

// WithMaxTipAmount adds max tip amount parameter
func (p *SendInvoiceParams) WithMaxTipAmount(maxTipAmount int) *SendInvoiceParams {
	p.MaxTipAmount = maxTipAmount
	return p
}

// WithSuggestedTipAmounts adds suggested tip amounts parameter
func (p *SendInvoiceParams) WithSuggestedTipAmounts(suggestedTipAmounts ...int) *SendInvoiceParams {
	p.SuggestedTipAmounts = suggestedTipAmounts
	return p
}

// WithStartParameter adds start parameter parameter
func (p *SendInvoiceParams) WithStartParameter(startParameter string) *SendInvoiceParams {
	p.StartParameter = startParameter
	return p
}

// WithProviderData adds provider data parameter
func (p *SendInvoiceParams) WithProviderData(providerData string) *SendInvoiceParams {
	p.ProviderData = providerData
	return p
}

// WithPhotoURL adds photo URL parameter
func (p *SendInvoiceParams) WithPhotoURL(photoURL string) *SendInvoiceParams {
	p.PhotoURL = photoURL
	return p
}

// WithPhotoSize adds photo size parameter
func (p *SendInvoiceParams) WithPhotoSize(photoSize int) *SendInvoiceParams {
	p.PhotoSize = photoSize
	return p
}

// WithPhotoWidth adds photo width parameter
func (p *SendInvoiceParams) WithPhotoWidth(photoWidth int) *SendInvoiceParams {
	p.PhotoWidth = photoWidth
	return p
}

// WithPhotoHeight adds photo height parameter
func (p *SendInvoiceParams) WithPhotoHeight(photoHeight int) *SendInvoiceParams {
	p.PhotoHeight = photoHeight
	return p
}

// WithNeedName adds need name parameter
func (p *SendInvoiceParams) WithNeedName() *SendInvoiceParams {
	p.NeedName = true
	return p
}

// WithNeedPhoneNumber adds need phone number parameter
func (p *SendInvoiceParams) WithNeedPhoneNumber() *SendInvoiceParams {
	p.NeedPhoneNumber = true
	return p
}

// WithNeedEmail adds need email parameter
func (p *SendInvoiceParams) WithNeedEmail() *SendInvoiceParams {
	p.NeedEmail = true
	return p
}

// WithNeedShippingAddress adds need shipping address parameter
func (p *SendInvoiceParams) WithNeedShippingAddress() *SendInvoiceParams {
	p.NeedShippingAddress = true
	return p
}

// WithSendPhoneNumberToProvider adds send phone number to provider parameter
func (p *SendInvoiceParams) WithSendPhoneNumberToProvider() *SendInvoiceParams {
	p.SendPhoneNumberToProvider = true
	return p
}

// WithSendEmailToProvider adds send email to provider parameter
func (p *SendInvoiceParams) WithSendEmailToProvider() *SendInvoiceParams {
	p.SendEmailToProvider = true
	return p
}

// WithIsFlexible adds is flexible parameter
func (p *SendInvoiceParams) WithIsFlexible() *SendInvoiceParams {
	p.IsFlexible = true
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendInvoiceParams) WithDisableNotification() *SendInvoiceParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendInvoiceParams) WithProtectContent() *SendInvoiceParams {
	p.ProtectContent = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendInvoiceParams) WithMessageEffectID(messageEffectID string) *SendInvoiceParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithReplyParameters adds reply parameters parameter
func (p *SendInvoiceParams) WithReplyParameters(replyParameters *ReplyParameters) *SendInvoiceParams {
	p.ReplyParameters = replyParameters
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendInvoiceParams) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *SendInvoiceParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithTitle adds title parameter
func (p *CreateInvoiceLinkParams) WithTitle(title string) *CreateInvoiceLinkParams {
	p.Title = title
	return p
}

// WithDescription adds description parameter
func (p *CreateInvoiceLinkParams) WithDescription(description string) *CreateInvoiceLinkParams {
	p.Description = description
	return p
}

// WithPayload adds payload parameter
func (p *CreateInvoiceLinkParams) WithPayload(payload string) *CreateInvoiceLinkParams {
	p.Payload = payload
	return p
}

// WithProviderToken adds provider token parameter
func (p *CreateInvoiceLinkParams) WithProviderToken(providerToken string) *CreateInvoiceLinkParams {
	p.ProviderToken = providerToken
	return p
}

// WithCurrency adds currency parameter
func (p *CreateInvoiceLinkParams) WithCurrency(currency string) *CreateInvoiceLinkParams {
	p.Currency = currency
	return p
}

// WithPrices adds prices parameter
func (p *CreateInvoiceLinkParams) WithPrices(prices ...LabeledPrice) *CreateInvoiceLinkParams {
	p.Prices = prices
	return p
}

// WithMaxTipAmount adds max tip amount parameter
func (p *CreateInvoiceLinkParams) WithMaxTipAmount(maxTipAmount int) *CreateInvoiceLinkParams {
	p.MaxTipAmount = maxTipAmount
	return p
}

// WithSuggestedTipAmounts adds suggested tip amounts parameter
func (p *CreateInvoiceLinkParams) WithSuggestedTipAmounts(suggestedTipAmounts ...int) *CreateInvoiceLinkParams {
	p.SuggestedTipAmounts = suggestedTipAmounts
	return p
}

// WithProviderData adds provider data parameter
func (p *CreateInvoiceLinkParams) WithProviderData(providerData string) *CreateInvoiceLinkParams {
	p.ProviderData = providerData
	return p
}

// WithPhotoURL adds photo URL parameter
func (p *CreateInvoiceLinkParams) WithPhotoURL(photoURL string) *CreateInvoiceLinkParams {
	p.PhotoURL = photoURL
	return p
}

// WithPhotoSize adds photo size parameter
func (p *CreateInvoiceLinkParams) WithPhotoSize(photoSize int) *CreateInvoiceLinkParams {
	p.PhotoSize = photoSize
	return p
}

// WithPhotoWidth adds photo width parameter
func (p *CreateInvoiceLinkParams) WithPhotoWidth(photoWidth int) *CreateInvoiceLinkParams {
	p.PhotoWidth = photoWidth
	return p
}

// WithPhotoHeight adds photo height parameter
func (p *CreateInvoiceLinkParams) WithPhotoHeight(photoHeight int) *CreateInvoiceLinkParams {
	p.PhotoHeight = photoHeight
	return p
}

// WithNeedName adds need name parameter
func (p *CreateInvoiceLinkParams) WithNeedName() *CreateInvoiceLinkParams {
	p.NeedName = true
	return p
}

// WithNeedPhoneNumber adds need phone number parameter
func (p *CreateInvoiceLinkParams) WithNeedPhoneNumber() *CreateInvoiceLinkParams {
	p.NeedPhoneNumber = true
	return p
}

// WithNeedEmail adds need email parameter
func (p *CreateInvoiceLinkParams) WithNeedEmail() *CreateInvoiceLinkParams {
	p.NeedEmail = true
	return p
}

// WithNeedShippingAddress adds need shipping address parameter
func (p *CreateInvoiceLinkParams) WithNeedShippingAddress() *CreateInvoiceLinkParams {
	p.NeedShippingAddress = true
	return p
}

// WithSendPhoneNumberToProvider adds send phone number to provider parameter
func (p *CreateInvoiceLinkParams) WithSendPhoneNumberToProvider() *CreateInvoiceLinkParams {
	p.SendPhoneNumberToProvider = true
	return p
}

// WithSendEmailToProvider adds send email to provider parameter
func (p *CreateInvoiceLinkParams) WithSendEmailToProvider() *CreateInvoiceLinkParams {
	p.SendEmailToProvider = true
	return p
}

// WithIsFlexible adds is flexible parameter
func (p *CreateInvoiceLinkParams) WithIsFlexible() *CreateInvoiceLinkParams {
	p.IsFlexible = true
	return p
}

// WithShippingQueryID adds shipping query ID parameter
func (p *AnswerShippingQueryParams) WithShippingQueryID(shippingQueryID string) *AnswerShippingQueryParams {
	p.ShippingQueryID = shippingQueryID
	return p
}

// WithOk adds ok parameter
func (p *AnswerShippingQueryParams) WithOk() *AnswerShippingQueryParams {
	p.Ok = true
	return p
}

// WithShippingOptions adds shipping options parameter
func (p *AnswerShippingQueryParams) WithShippingOptions(shippingOptions ...ShippingOption) *AnswerShippingQueryParams {
	p.ShippingOptions = shippingOptions
	return p
}

// WithErrorMessage adds error message parameter
func (p *AnswerShippingQueryParams) WithErrorMessage(errorMessage string) *AnswerShippingQueryParams {
	p.ErrorMessage = errorMessage
	return p
}

// WithPreCheckoutQueryID adds pre checkout query ID parameter
func (p *AnswerPreCheckoutQueryParams) WithPreCheckoutQueryID(preCheckoutQueryID string) *AnswerPreCheckoutQueryParams {
	p.PreCheckoutQueryID = preCheckoutQueryID
	return p
}

// WithOk adds ok parameter
func (p *AnswerPreCheckoutQueryParams) WithOk() *AnswerPreCheckoutQueryParams {
	p.Ok = true
	return p
}

// WithErrorMessage adds error message parameter
func (p *AnswerPreCheckoutQueryParams) WithErrorMessage(errorMessage string) *AnswerPreCheckoutQueryParams {
	p.ErrorMessage = errorMessage
	return p
}

// WithTelegramPaymentChargeID adds telegram payment charge ID parameter
func (p *RefundStarPaymentParams) WithTelegramPaymentChargeID(telegramPaymentChargeID string) *RefundStarPaymentParams {
	p.TelegramPaymentChargeID = telegramPaymentChargeID
	return p
}

// WithErrors adds errors parameter
func (p *SetPassportDataErrorsParams) WithErrors(errors ...PassportElementError) *SetPassportDataErrorsParams {
	p.Errors = errors
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *SendGameParams) WithBusinessConnectionID(businessConnectionID string) *SendGameParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *SendGameParams) WithMessageThreadID(messageThreadID int) *SendGameParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithGameShortName adds game short name parameter
func (p *SendGameParams) WithGameShortName(gameShortName string) *SendGameParams {
	p.GameShortName = gameShortName
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendGameParams) WithDisableNotification() *SendGameParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendGameParams) WithProtectContent() *SendGameParams {
	p.ProtectContent = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendGameParams) WithMessageEffectID(messageEffectID string) *SendGameParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithReplyParameters adds reply parameters parameter
func (p *SendGameParams) WithReplyParameters(replyParameters *ReplyParameters) *SendGameParams {
	p.ReplyParameters = replyParameters
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendGameParams) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *SendGameParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithScore adds score parameter
func (p *SetGameScoreParams) WithScore(score int) *SetGameScoreParams {
	p.Score = score
	return p
}

// WithForce adds force parameter
func (p *SetGameScoreParams) WithForce() *SetGameScoreParams {
	p.Force = true
	return p
}

// WithDisableEditMessage adds disable edit message parameter
func (p *SetGameScoreParams) WithDisableEditMessage() *SetGameScoreParams {
	p.DisableEditMessage = true
	return p
}

// WithMessageID adds message ID parameter
func (p *SetGameScoreParams) WithMessageID(messageID int) *SetGameScoreParams {
	p.MessageID = messageID
	return p
}

// WithInlineMessageID adds inline message ID parameter
func (p *SetGameScoreParams) WithInlineMessageID(inlineMessageID string) *SetGameScoreParams {
	p.InlineMessageID = inlineMessageID
	return p
}

// WithMessageID adds message ID parameter
func (p *GetGameHighScoresParams) WithMessageID(messageID int) *GetGameHighScoresParams {
	p.MessageID = messageID
	return p
}

// WithInlineMessageID adds inline message ID parameter
func (p *GetGameHighScoresParams) WithInlineMessageID(inlineMessageID string) *GetGameHighScoresParams {
	p.InlineMessageID = inlineMessageID
	return p
}
