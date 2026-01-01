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

// WithDirectMessagesTopicID adds direct messages topic ID parameter
func (p *SendMessageParams) WithDirectMessagesTopicID(directMessagesTopicID int) *SendMessageParams {
	p.DirectMessagesTopicID = directMessagesTopicID
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

// WithAllowPaidBroadcast adds allow paid broadcast parameter
func (p *SendMessageParams) WithAllowPaidBroadcast() *SendMessageParams {
	p.AllowPaidBroadcast = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendMessageParams) WithMessageEffectID(messageEffectID string) *SendMessageParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithSuggestedPostParameters adds suggested post parameters parameter
func (p *SendMessageParams) WithSuggestedPostParameters(suggestedPostParameters *SuggestedPostParameters,
) *SendMessageParams {
	p.SuggestedPostParameters = suggestedPostParameters
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

// WithDirectMessagesTopicID adds direct messages topic ID parameter
func (p *ForwardMessageParams) WithDirectMessagesTopicID(directMessagesTopicID int) *ForwardMessageParams {
	p.DirectMessagesTopicID = directMessagesTopicID
	return p
}

// WithFromChatID adds from chat ID parameter
func (p *ForwardMessageParams) WithFromChatID(fromChatID ChatID) *ForwardMessageParams {
	p.FromChatID = fromChatID
	return p
}

// WithVideoStartTimestamp adds video start timestamp parameter
func (p *ForwardMessageParams) WithVideoStartTimestamp(videoStartTimestamp int) *ForwardMessageParams {
	p.VideoStartTimestamp = videoStartTimestamp
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

// WithMessageEffectID adds message effect ID parameter
func (p *ForwardMessageParams) WithMessageEffectID(messageEffectID string) *ForwardMessageParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithSuggestedPostParameters adds suggested post parameters parameter
func (p *ForwardMessageParams) WithSuggestedPostParameters(suggestedPostParameters *SuggestedPostParameters,
) *ForwardMessageParams {
	p.SuggestedPostParameters = suggestedPostParameters
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

// WithDirectMessagesTopicID adds direct messages topic ID parameter
func (p *ForwardMessagesParams) WithDirectMessagesTopicID(directMessagesTopicID int) *ForwardMessagesParams {
	p.DirectMessagesTopicID = directMessagesTopicID
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

// WithDirectMessagesTopicID adds direct messages topic ID parameter
func (p *CopyMessageParams) WithDirectMessagesTopicID(directMessagesTopicID int) *CopyMessageParams {
	p.DirectMessagesTopicID = directMessagesTopicID
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

// WithVideoStartTimestamp adds video start timestamp parameter
func (p *CopyMessageParams) WithVideoStartTimestamp(videoStartTimestamp int) *CopyMessageParams {
	p.VideoStartTimestamp = videoStartTimestamp
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

// WithAllowPaidBroadcast adds allow paid broadcast parameter
func (p *CopyMessageParams) WithAllowPaidBroadcast() *CopyMessageParams {
	p.AllowPaidBroadcast = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *CopyMessageParams) WithMessageEffectID(messageEffectID string) *CopyMessageParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithSuggestedPostParameters adds suggested post parameters parameter
func (p *CopyMessageParams) WithSuggestedPostParameters(suggestedPostParameters *SuggestedPostParameters,
) *CopyMessageParams {
	p.SuggestedPostParameters = suggestedPostParameters
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

// WithDirectMessagesTopicID adds direct messages topic ID parameter
func (p *CopyMessagesParams) WithDirectMessagesTopicID(directMessagesTopicID int) *CopyMessagesParams {
	p.DirectMessagesTopicID = directMessagesTopicID
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

// WithDirectMessagesTopicID adds direct messages topic ID parameter
func (p *SendPhotoParams) WithDirectMessagesTopicID(directMessagesTopicID int) *SendPhotoParams {
	p.DirectMessagesTopicID = directMessagesTopicID
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

// WithAllowPaidBroadcast adds allow paid broadcast parameter
func (p *SendPhotoParams) WithAllowPaidBroadcast() *SendPhotoParams {
	p.AllowPaidBroadcast = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendPhotoParams) WithMessageEffectID(messageEffectID string) *SendPhotoParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithSuggestedPostParameters adds suggested post parameters parameter
func (p *SendPhotoParams) WithSuggestedPostParameters(suggestedPostParameters *SuggestedPostParameters,
) *SendPhotoParams {
	p.SuggestedPostParameters = suggestedPostParameters
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

// WithDirectMessagesTopicID adds direct messages topic ID parameter
func (p *SendAudioParams) WithDirectMessagesTopicID(directMessagesTopicID int) *SendAudioParams {
	p.DirectMessagesTopicID = directMessagesTopicID
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

// WithAllowPaidBroadcast adds allow paid broadcast parameter
func (p *SendAudioParams) WithAllowPaidBroadcast() *SendAudioParams {
	p.AllowPaidBroadcast = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendAudioParams) WithMessageEffectID(messageEffectID string) *SendAudioParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithSuggestedPostParameters adds suggested post parameters parameter
func (p *SendAudioParams) WithSuggestedPostParameters(suggestedPostParameters *SuggestedPostParameters,
) *SendAudioParams {
	p.SuggestedPostParameters = suggestedPostParameters
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

// WithDirectMessagesTopicID adds direct messages topic ID parameter
func (p *SendDocumentParams) WithDirectMessagesTopicID(directMessagesTopicID int) *SendDocumentParams {
	p.DirectMessagesTopicID = directMessagesTopicID
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

// WithAllowPaidBroadcast adds allow paid broadcast parameter
func (p *SendDocumentParams) WithAllowPaidBroadcast() *SendDocumentParams {
	p.AllowPaidBroadcast = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendDocumentParams) WithMessageEffectID(messageEffectID string) *SendDocumentParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithSuggestedPostParameters adds suggested post parameters parameter
func (p *SendDocumentParams) WithSuggestedPostParameters(suggestedPostParameters *SuggestedPostParameters,
) *SendDocumentParams {
	p.SuggestedPostParameters = suggestedPostParameters
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

// WithDirectMessagesTopicID adds direct messages topic ID parameter
func (p *SendVideoParams) WithDirectMessagesTopicID(directMessagesTopicID int) *SendVideoParams {
	p.DirectMessagesTopicID = directMessagesTopicID
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

// WithCover adds cover parameter
func (p *SendVideoParams) WithCover(cover *InputFile) *SendVideoParams {
	p.Cover = cover
	return p
}

// WithStartTimestamp adds start timestamp parameter
func (p *SendVideoParams) WithStartTimestamp(startTimestamp int) *SendVideoParams {
	p.StartTimestamp = startTimestamp
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

// WithAllowPaidBroadcast adds allow paid broadcast parameter
func (p *SendVideoParams) WithAllowPaidBroadcast() *SendVideoParams {
	p.AllowPaidBroadcast = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendVideoParams) WithMessageEffectID(messageEffectID string) *SendVideoParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithSuggestedPostParameters adds suggested post parameters parameter
func (p *SendVideoParams) WithSuggestedPostParameters(suggestedPostParameters *SuggestedPostParameters,
) *SendVideoParams {
	p.SuggestedPostParameters = suggestedPostParameters
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

// WithDirectMessagesTopicID adds direct messages topic ID parameter
func (p *SendAnimationParams) WithDirectMessagesTopicID(directMessagesTopicID int) *SendAnimationParams {
	p.DirectMessagesTopicID = directMessagesTopicID
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

// WithAllowPaidBroadcast adds allow paid broadcast parameter
func (p *SendAnimationParams) WithAllowPaidBroadcast() *SendAnimationParams {
	p.AllowPaidBroadcast = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendAnimationParams) WithMessageEffectID(messageEffectID string) *SendAnimationParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithSuggestedPostParameters adds suggested post parameters parameter
func (p *SendAnimationParams) WithSuggestedPostParameters(suggestedPostParameters *SuggestedPostParameters,
) *SendAnimationParams {
	p.SuggestedPostParameters = suggestedPostParameters
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

// WithDirectMessagesTopicID adds direct messages topic ID parameter
func (p *SendVoiceParams) WithDirectMessagesTopicID(directMessagesTopicID int) *SendVoiceParams {
	p.DirectMessagesTopicID = directMessagesTopicID
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

// WithAllowPaidBroadcast adds allow paid broadcast parameter
func (p *SendVoiceParams) WithAllowPaidBroadcast() *SendVoiceParams {
	p.AllowPaidBroadcast = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendVoiceParams) WithMessageEffectID(messageEffectID string) *SendVoiceParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithSuggestedPostParameters adds suggested post parameters parameter
func (p *SendVoiceParams) WithSuggestedPostParameters(suggestedPostParameters *SuggestedPostParameters,
) *SendVoiceParams {
	p.SuggestedPostParameters = suggestedPostParameters
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

// WithDirectMessagesTopicID adds direct messages topic ID parameter
func (p *SendVideoNoteParams) WithDirectMessagesTopicID(directMessagesTopicID int) *SendVideoNoteParams {
	p.DirectMessagesTopicID = directMessagesTopicID
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

// WithAllowPaidBroadcast adds allow paid broadcast parameter
func (p *SendVideoNoteParams) WithAllowPaidBroadcast() *SendVideoNoteParams {
	p.AllowPaidBroadcast = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendVideoNoteParams) WithMessageEffectID(messageEffectID string) *SendVideoNoteParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithSuggestedPostParameters adds suggested post parameters parameter
func (p *SendVideoNoteParams) WithSuggestedPostParameters(suggestedPostParameters *SuggestedPostParameters,
) *SendVideoNoteParams {
	p.SuggestedPostParameters = suggestedPostParameters
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
func (p *SendPaidMediaParams) WithBusinessConnectionID(businessConnectionID string) *SendPaidMediaParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithChatID adds chat ID parameter
func (p *SendPaidMediaParams) WithChatID(chatID ChatID) *SendPaidMediaParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *SendPaidMediaParams) WithMessageThreadID(messageThreadID int) *SendPaidMediaParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithDirectMessagesTopicID adds direct messages topic ID parameter
func (p *SendPaidMediaParams) WithDirectMessagesTopicID(directMessagesTopicID int) *SendPaidMediaParams {
	p.DirectMessagesTopicID = directMessagesTopicID
	return p
}

// WithStarCount adds star count parameter
func (p *SendPaidMediaParams) WithStarCount(starCount int) *SendPaidMediaParams {
	p.StarCount = starCount
	return p
}

// WithMedia adds media parameter
func (p *SendPaidMediaParams) WithMedia(media ...InputPaidMedia) *SendPaidMediaParams {
	p.Media = media
	return p
}

// WithPayload adds payload parameter
func (p *SendPaidMediaParams) WithPayload(payload string) *SendPaidMediaParams {
	p.Payload = payload
	return p
}

// WithCaption adds caption parameter
func (p *SendPaidMediaParams) WithCaption(caption string) *SendPaidMediaParams {
	p.Caption = caption
	return p
}

// WithParseMode adds parse mode parameter
func (p *SendPaidMediaParams) WithParseMode(parseMode string) *SendPaidMediaParams {
	p.ParseMode = parseMode
	return p
}

// WithCaptionEntities adds caption entities parameter
func (p *SendPaidMediaParams) WithCaptionEntities(captionEntities ...MessageEntity) *SendPaidMediaParams {
	p.CaptionEntities = captionEntities
	return p
}

// WithShowCaptionAboveMedia adds show caption above media parameter
func (p *SendPaidMediaParams) WithShowCaptionAboveMedia() *SendPaidMediaParams {
	p.ShowCaptionAboveMedia = true
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendPaidMediaParams) WithDisableNotification() *SendPaidMediaParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendPaidMediaParams) WithProtectContent() *SendPaidMediaParams {
	p.ProtectContent = true
	return p
}

// WithAllowPaidBroadcast adds allow paid broadcast parameter
func (p *SendPaidMediaParams) WithAllowPaidBroadcast() *SendPaidMediaParams {
	p.AllowPaidBroadcast = true
	return p
}

// WithSuggestedPostParameters adds suggested post parameters parameter
func (p *SendPaidMediaParams) WithSuggestedPostParameters(suggestedPostParameters *SuggestedPostParameters,
) *SendPaidMediaParams {
	p.SuggestedPostParameters = suggestedPostParameters
	return p
}

// WithReplyParameters adds reply parameters parameter
func (p *SendPaidMediaParams) WithReplyParameters(replyParameters *ReplyParameters) *SendPaidMediaParams {
	p.ReplyParameters = replyParameters
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendPaidMediaParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendPaidMediaParams {
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

// WithDirectMessagesTopicID adds direct messages topic ID parameter
func (p *SendMediaGroupParams) WithDirectMessagesTopicID(directMessagesTopicID int) *SendMediaGroupParams {
	p.DirectMessagesTopicID = directMessagesTopicID
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

// WithAllowPaidBroadcast adds allow paid broadcast parameter
func (p *SendMediaGroupParams) WithAllowPaidBroadcast() *SendMediaGroupParams {
	p.AllowPaidBroadcast = true
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

// WithDirectMessagesTopicID adds direct messages topic ID parameter
func (p *SendLocationParams) WithDirectMessagesTopicID(directMessagesTopicID int) *SendLocationParams {
	p.DirectMessagesTopicID = directMessagesTopicID
	return p
}

// WithLatitude adds latitude parameter
func (p *SendLocationParams) WithLatitude(latitude float64) *SendLocationParams {
	p.Latitude = latitude
	return p
}

// WithLongitude adds longitude parameter
func (p *SendLocationParams) WithLongitude(longitude float64) *SendLocationParams {
	p.Longitude = longitude
	return p
}

// WithHorizontalAccuracy adds horizontal accuracy parameter
func (p *SendLocationParams) WithHorizontalAccuracy(horizontalAccuracy float64) *SendLocationParams {
	p.HorizontalAccuracy = horizontalAccuracy
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

// WithAllowPaidBroadcast adds allow paid broadcast parameter
func (p *SendLocationParams) WithAllowPaidBroadcast() *SendLocationParams {
	p.AllowPaidBroadcast = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendLocationParams) WithMessageEffectID(messageEffectID string) *SendLocationParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithSuggestedPostParameters adds suggested post parameters parameter
func (p *SendLocationParams) WithSuggestedPostParameters(suggestedPostParameters *SuggestedPostParameters,
) *SendLocationParams {
	p.SuggestedPostParameters = suggestedPostParameters
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

// WithDirectMessagesTopicID adds direct messages topic ID parameter
func (p *SendVenueParams) WithDirectMessagesTopicID(directMessagesTopicID int) *SendVenueParams {
	p.DirectMessagesTopicID = directMessagesTopicID
	return p
}

// WithLatitude adds latitude parameter
func (p *SendVenueParams) WithLatitude(latitude float64) *SendVenueParams {
	p.Latitude = latitude
	return p
}

// WithLongitude adds longitude parameter
func (p *SendVenueParams) WithLongitude(longitude float64) *SendVenueParams {
	p.Longitude = longitude
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

// WithAllowPaidBroadcast adds allow paid broadcast parameter
func (p *SendVenueParams) WithAllowPaidBroadcast() *SendVenueParams {
	p.AllowPaidBroadcast = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendVenueParams) WithMessageEffectID(messageEffectID string) *SendVenueParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithSuggestedPostParameters adds suggested post parameters parameter
func (p *SendVenueParams) WithSuggestedPostParameters(suggestedPostParameters *SuggestedPostParameters,
) *SendVenueParams {
	p.SuggestedPostParameters = suggestedPostParameters
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

// WithDirectMessagesTopicID adds direct messages topic ID parameter
func (p *SendContactParams) WithDirectMessagesTopicID(directMessagesTopicID int) *SendContactParams {
	p.DirectMessagesTopicID = directMessagesTopicID
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

// WithAllowPaidBroadcast adds allow paid broadcast parameter
func (p *SendContactParams) WithAllowPaidBroadcast() *SendContactParams {
	p.AllowPaidBroadcast = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendContactParams) WithMessageEffectID(messageEffectID string) *SendContactParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithSuggestedPostParameters adds suggested post parameters parameter
func (p *SendContactParams) WithSuggestedPostParameters(suggestedPostParameters *SuggestedPostParameters,
) *SendContactParams {
	p.SuggestedPostParameters = suggestedPostParameters
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
	p.IsAnonymous = &isAnonymous
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
	p.CorrectOptionID = &correctOptionID
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

// WithCloseDate adds close date parameter
func (p *SendPollParams) WithCloseDate(closeDate int64) *SendPollParams {
	p.CloseDate = closeDate
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

// WithAllowPaidBroadcast adds allow paid broadcast parameter
func (p *SendPollParams) WithAllowPaidBroadcast() *SendPollParams {
	p.AllowPaidBroadcast = true
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
func (p *SendChecklistParams) WithBusinessConnectionID(businessConnectionID string) *SendChecklistParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithChatID adds chat ID parameter
func (p *SendChecklistParams) WithChatID(chatID int64) *SendChecklistParams {
	p.ChatID = chatID
	return p
}

// WithChecklist adds checklist parameter
func (p *SendChecklistParams) WithChecklist(checklist InputChecklist) *SendChecklistParams {
	p.Checklist = checklist
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendChecklistParams) WithDisableNotification() *SendChecklistParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendChecklistParams) WithProtectContent() *SendChecklistParams {
	p.ProtectContent = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendChecklistParams) WithMessageEffectID(messageEffectID string) *SendChecklistParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithReplyParameters adds reply parameters parameter
func (p *SendChecklistParams) WithReplyParameters(replyParameters *ReplyParameters) *SendChecklistParams {
	p.ReplyParameters = replyParameters
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendChecklistParams) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *SendChecklistParams {
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

// WithDirectMessagesTopicID adds direct messages topic ID parameter
func (p *SendDiceParams) WithDirectMessagesTopicID(directMessagesTopicID int) *SendDiceParams {
	p.DirectMessagesTopicID = directMessagesTopicID
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

// WithAllowPaidBroadcast adds allow paid broadcast parameter
func (p *SendDiceParams) WithAllowPaidBroadcast() *SendDiceParams {
	p.AllowPaidBroadcast = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendDiceParams) WithMessageEffectID(messageEffectID string) *SendDiceParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithSuggestedPostParameters adds suggested post parameters parameter
func (p *SendDiceParams) WithSuggestedPostParameters(suggestedPostParameters *SuggestedPostParameters) *SendDiceParams {
	p.SuggestedPostParameters = suggestedPostParameters
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

// WithChatID adds chat ID parameter
func (p *SendMessageDraftParams) WithChatID(chatID int64) *SendMessageDraftParams {
	p.ChatID = chatID
	return p
}

// WithMessageThreadID adds message thread ID parameter
func (p *SendMessageDraftParams) WithMessageThreadID(messageThreadID int) *SendMessageDraftParams {
	p.MessageThreadID = messageThreadID
	return p
}

// WithDraftID adds draft ID parameter
func (p *SendMessageDraftParams) WithDraftID(draftID int) *SendMessageDraftParams {
	p.DraftID = draftID
	return p
}

// WithText adds text parameter
func (p *SendMessageDraftParams) WithText(text string) *SendMessageDraftParams {
	p.Text = text
	return p
}

// WithParseMode adds parse mode parameter
func (p *SendMessageDraftParams) WithParseMode(parseMode string) *SendMessageDraftParams {
	p.ParseMode = parseMode
	return p
}

// WithEntities adds entities parameter
func (p *SendMessageDraftParams) WithEntities(entities ...MessageEntity) *SendMessageDraftParams {
	p.Entities = entities
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

// WithUserID adds user ID parameter
func (p *GetUserProfilePhotosParams) WithUserID(userID int64) *GetUserProfilePhotosParams {
	p.UserID = userID
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

// WithUserID adds user ID parameter
func (p *SetUserEmojiStatusParams) WithUserID(userID int64) *SetUserEmojiStatusParams {
	p.UserID = userID
	return p
}

// WithEmojiStatusCustomEmojiID adds emoji status custom emoji ID parameter
func (p *SetUserEmojiStatusParams) WithEmojiStatusCustomEmojiID(emojiStatusCustomEmojiID string,
) *SetUserEmojiStatusParams {
	p.EmojiStatusCustomEmojiID = emojiStatusCustomEmojiID
	return p
}

// WithEmojiStatusExpirationDate adds emoji status expiration date parameter
func (p *SetUserEmojiStatusParams) WithEmojiStatusExpirationDate(emojiStatusExpirationDate int64,
) *SetUserEmojiStatusParams {
	p.EmojiStatusExpirationDate = emojiStatusExpirationDate
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

// WithUserID adds user ID parameter
func (p *BanChatMemberParams) WithUserID(userID int64) *BanChatMemberParams {
	p.UserID = userID
	return p
}

// WithUntilDate adds until date parameter
func (p *BanChatMemberParams) WithUntilDate(untilDate int64) *BanChatMemberParams {
	p.UntilDate = untilDate
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

// WithUserID adds user ID parameter
func (p *UnbanChatMemberParams) WithUserID(userID int64) *UnbanChatMemberParams {
	p.UserID = userID
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

// WithUserID adds user ID parameter
func (p *RestrictChatMemberParams) WithUserID(userID int64) *RestrictChatMemberParams {
	p.UserID = userID
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

// WithUntilDate adds until date parameter
func (p *RestrictChatMemberParams) WithUntilDate(untilDate int64) *RestrictChatMemberParams {
	p.UntilDate = untilDate
	return p
}

// WithChatID adds chat ID parameter
func (p *PromoteChatMemberParams) WithChatID(chatID ChatID) *PromoteChatMemberParams {
	p.ChatID = chatID
	return p
}

// WithUserID adds user ID parameter
func (p *PromoteChatMemberParams) WithUserID(userID int64) *PromoteChatMemberParams {
	p.UserID = userID
	return p
}

// WithIsAnonymous adds is anonymous parameter
func (p *PromoteChatMemberParams) WithIsAnonymous(isAnonymous bool) *PromoteChatMemberParams {
	p.IsAnonymous = &isAnonymous
	return p
}

// WithCanManageChat adds can manage chat parameter
func (p *PromoteChatMemberParams) WithCanManageChat(canManageChat bool) *PromoteChatMemberParams {
	p.CanManageChat = &canManageChat
	return p
}

// WithCanDeleteMessages adds can delete messages parameter
func (p *PromoteChatMemberParams) WithCanDeleteMessages(canDeleteMessages bool) *PromoteChatMemberParams {
	p.CanDeleteMessages = &canDeleteMessages
	return p
}

// WithCanManageVideoChats adds can manage video chats parameter
func (p *PromoteChatMemberParams) WithCanManageVideoChats(canManageVideoChats bool) *PromoteChatMemberParams {
	p.CanManageVideoChats = &canManageVideoChats
	return p
}

// WithCanRestrictMembers adds can restrict members parameter
func (p *PromoteChatMemberParams) WithCanRestrictMembers(canRestrictMembers bool) *PromoteChatMemberParams {
	p.CanRestrictMembers = &canRestrictMembers
	return p
}

// WithCanPromoteMembers adds can promote members parameter
func (p *PromoteChatMemberParams) WithCanPromoteMembers(canPromoteMembers bool) *PromoteChatMemberParams {
	p.CanPromoteMembers = &canPromoteMembers
	return p
}

// WithCanChangeInfo adds can change info parameter
func (p *PromoteChatMemberParams) WithCanChangeInfo(canChangeInfo bool) *PromoteChatMemberParams {
	p.CanChangeInfo = &canChangeInfo
	return p
}

// WithCanInviteUsers adds can invite users parameter
func (p *PromoteChatMemberParams) WithCanInviteUsers(canInviteUsers bool) *PromoteChatMemberParams {
	p.CanInviteUsers = &canInviteUsers
	return p
}

// WithCanPostStories adds can post stories parameter
func (p *PromoteChatMemberParams) WithCanPostStories(canPostStories bool) *PromoteChatMemberParams {
	p.CanPostStories = &canPostStories
	return p
}

// WithCanEditStories adds can edit stories parameter
func (p *PromoteChatMemberParams) WithCanEditStories(canEditStories bool) *PromoteChatMemberParams {
	p.CanEditStories = &canEditStories
	return p
}

// WithCanDeleteStories adds can delete stories parameter
func (p *PromoteChatMemberParams) WithCanDeleteStories(canDeleteStories bool) *PromoteChatMemberParams {
	p.CanDeleteStories = &canDeleteStories
	return p
}

// WithCanPostMessages adds can post messages parameter
func (p *PromoteChatMemberParams) WithCanPostMessages(canPostMessages bool) *PromoteChatMemberParams {
	p.CanPostMessages = &canPostMessages
	return p
}

// WithCanEditMessages adds can edit messages parameter
func (p *PromoteChatMemberParams) WithCanEditMessages(canEditMessages bool) *PromoteChatMemberParams {
	p.CanEditMessages = &canEditMessages
	return p
}

// WithCanPinMessages adds can pin messages parameter
func (p *PromoteChatMemberParams) WithCanPinMessages(canPinMessages bool) *PromoteChatMemberParams {
	p.CanPinMessages = &canPinMessages
	return p
}

// WithCanManageTopics adds can manage topics parameter
func (p *PromoteChatMemberParams) WithCanManageTopics(canManageTopics bool) *PromoteChatMemberParams {
	p.CanManageTopics = &canManageTopics
	return p
}

// WithCanManageDirectMessages adds can manage direct messages parameter
func (p *PromoteChatMemberParams) WithCanManageDirectMessages(canManageDirectMessages bool) *PromoteChatMemberParams {
	p.CanManageDirectMessages = &canManageDirectMessages
	return p
}

// WithChatID adds chat ID parameter
func (p *SetChatAdministratorCustomTitleParams) WithChatID(chatID ChatID) *SetChatAdministratorCustomTitleParams {
	p.ChatID = chatID
	return p
}

// WithUserID adds user ID parameter
func (p *SetChatAdministratorCustomTitleParams) WithUserID(userID int64) *SetChatAdministratorCustomTitleParams {
	p.UserID = userID
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

// WithSenderChatID adds sender chat ID parameter
func (p *BanChatSenderChatParams) WithSenderChatID(senderChatID int64) *BanChatSenderChatParams {
	p.SenderChatID = senderChatID
	return p
}

// WithChatID adds chat ID parameter
func (p *UnbanChatSenderChatParams) WithChatID(chatID ChatID) *UnbanChatSenderChatParams {
	p.ChatID = chatID
	return p
}

// WithSenderChatID adds sender chat ID parameter
func (p *UnbanChatSenderChatParams) WithSenderChatID(senderChatID int64) *UnbanChatSenderChatParams {
	p.SenderChatID = senderChatID
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

// WithExpireDate adds expire date parameter
func (p *CreateChatInviteLinkParams) WithExpireDate(expireDate int64) *CreateChatInviteLinkParams {
	p.ExpireDate = expireDate
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

// WithExpireDate adds expire date parameter
func (p *EditChatInviteLinkParams) WithExpireDate(expireDate int64) *EditChatInviteLinkParams {
	p.ExpireDate = expireDate
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
func (p *CreateChatSubscriptionInviteLinkParams) WithChatID(chatID ChatID) *CreateChatSubscriptionInviteLinkParams {
	p.ChatID = chatID
	return p
}

// WithName adds name parameter
func (p *CreateChatSubscriptionInviteLinkParams) WithName(name string) *CreateChatSubscriptionInviteLinkParams {
	p.Name = name
	return p
}

// WithSubscriptionPeriod adds subscription period parameter
func (p *CreateChatSubscriptionInviteLinkParams) WithSubscriptionPeriod(subscriptionPeriod int64,
) *CreateChatSubscriptionInviteLinkParams {
	p.SubscriptionPeriod = subscriptionPeriod
	return p
}

// WithSubscriptionPrice adds subscription price parameter
func (p *CreateChatSubscriptionInviteLinkParams) WithSubscriptionPrice(subscriptionPrice int,
) *CreateChatSubscriptionInviteLinkParams {
	p.SubscriptionPrice = subscriptionPrice
	return p
}

// WithChatID adds chat ID parameter
func (p *EditChatSubscriptionInviteLinkParams) WithChatID(chatID ChatID) *EditChatSubscriptionInviteLinkParams {
	p.ChatID = chatID
	return p
}

// WithInviteLink adds invite link parameter
func (p *EditChatSubscriptionInviteLinkParams) WithInviteLink(inviteLink string) *EditChatSubscriptionInviteLinkParams {
	p.InviteLink = inviteLink
	return p
}

// WithName adds name parameter
func (p *EditChatSubscriptionInviteLinkParams) WithName(name string) *EditChatSubscriptionInviteLinkParams {
	p.Name = name
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

// WithUserID adds user ID parameter
func (p *ApproveChatJoinRequestParams) WithUserID(userID int64) *ApproveChatJoinRequestParams {
	p.UserID = userID
	return p
}

// WithChatID adds chat ID parameter
func (p *DeclineChatJoinRequestParams) WithChatID(chatID ChatID) *DeclineChatJoinRequestParams {
	p.ChatID = chatID
	return p
}

// WithUserID adds user ID parameter
func (p *DeclineChatJoinRequestParams) WithUserID(userID int64) *DeclineChatJoinRequestParams {
	p.UserID = userID
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

// WithBusinessConnectionID adds business connection ID parameter
func (p *PinChatMessageParams) WithBusinessConnectionID(businessConnectionID string) *PinChatMessageParams {
	p.BusinessConnectionID = businessConnectionID
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

// WithBusinessConnectionID adds business connection ID parameter
func (p *UnpinChatMessageParams) WithBusinessConnectionID(businessConnectionID string) *UnpinChatMessageParams {
	p.BusinessConnectionID = businessConnectionID
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

// WithUserID adds user ID parameter
func (p *GetChatMemberParams) WithUserID(userID int64) *GetChatMemberParams {
	p.UserID = userID
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
	p.IconCustomEmojiID = &iconCustomEmojiID
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

// WithUserID adds user ID parameter
func (p *GetUserChatBoostsParams) WithUserID(userID int64) *GetUserChatBoostsParams {
	p.UserID = userID
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

// WithChatID adds chat ID parameter
func (p *SetChatMenuButtonParams) WithChatID(chatID int64) *SetChatMenuButtonParams {
	p.ChatID = chatID
	return p
}

// WithMenuButton adds menu button parameter
func (p *SetChatMenuButtonParams) WithMenuButton(menuButton MenuButton) *SetChatMenuButtonParams {
	p.MenuButton = menuButton
	return p
}

// WithChatID adds chat ID parameter
func (p *GetChatMenuButtonParams) WithChatID(chatID int64) *GetChatMenuButtonParams {
	p.ChatID = chatID
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

// WithUserID adds user ID parameter
func (p *SendGiftParams) WithUserID(userID int64) *SendGiftParams {
	p.UserID = userID
	return p
}

// WithChatID adds chat ID parameter
func (p *SendGiftParams) WithChatID(chatID ChatID) *SendGiftParams {
	p.ChatID = chatID
	return p
}

// WithGiftID adds gift ID parameter
func (p *SendGiftParams) WithGiftID(giftID string) *SendGiftParams {
	p.GiftID = giftID
	return p
}

// WithPayForUpgrade adds pay for upgrade parameter
func (p *SendGiftParams) WithPayForUpgrade() *SendGiftParams {
	p.PayForUpgrade = true
	return p
}

// WithText adds text parameter
func (p *SendGiftParams) WithText(text string) *SendGiftParams {
	p.Text = text
	return p
}

// WithTextParseMode adds text parse mode parameter
func (p *SendGiftParams) WithTextParseMode(textParseMode string) *SendGiftParams {
	p.TextParseMode = textParseMode
	return p
}

// WithTextEntities adds text entities parameter
func (p *SendGiftParams) WithTextEntities(textEntities ...MessageEntity) *SendGiftParams {
	p.TextEntities = textEntities
	return p
}

// WithUserID adds user ID parameter
func (p *GiftPremiumSubscriptionParams) WithUserID(userID int64) *GiftPremiumSubscriptionParams {
	p.UserID = userID
	return p
}

// WithMonthCount adds month count parameter
func (p *GiftPremiumSubscriptionParams) WithMonthCount(monthCount int) *GiftPremiumSubscriptionParams {
	p.MonthCount = monthCount
	return p
}

// WithStarCount adds star count parameter
func (p *GiftPremiumSubscriptionParams) WithStarCount(starCount int) *GiftPremiumSubscriptionParams {
	p.StarCount = starCount
	return p
}

// WithText adds text parameter
func (p *GiftPremiumSubscriptionParams) WithText(text string) *GiftPremiumSubscriptionParams {
	p.Text = text
	return p
}

// WithTextParseMode adds text parse mode parameter
func (p *GiftPremiumSubscriptionParams) WithTextParseMode(textParseMode string) *GiftPremiumSubscriptionParams {
	p.TextParseMode = textParseMode
	return p
}

// WithTextEntities adds text entities parameter
func (p *GiftPremiumSubscriptionParams) WithTextEntities(textEntities ...MessageEntity) *GiftPremiumSubscriptionParams {
	p.TextEntities = textEntities
	return p
}

// WithUserID adds user ID parameter
func (p *VerifyUserParams) WithUserID(userID int64) *VerifyUserParams {
	p.UserID = userID
	return p
}

// WithCustomDescription adds custom description parameter
func (p *VerifyUserParams) WithCustomDescription(customDescription string) *VerifyUserParams {
	p.CustomDescription = customDescription
	return p
}

// WithChatID adds chat ID parameter
func (p *VerifyChatParams) WithChatID(chatID ChatID) *VerifyChatParams {
	p.ChatID = chatID
	return p
}

// WithCustomDescription adds custom description parameter
func (p *VerifyChatParams) WithCustomDescription(customDescription string) *VerifyChatParams {
	p.CustomDescription = customDescription
	return p
}

// WithUserID adds user ID parameter
func (p *RemoveUserVerificationParams) WithUserID(userID int64) *RemoveUserVerificationParams {
	p.UserID = userID
	return p
}

// WithChatID adds chat ID parameter
func (p *RemoveChatVerificationParams) WithChatID(chatID ChatID) *RemoveChatVerificationParams {
	p.ChatID = chatID
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *ReadBusinessMessageParams) WithBusinessConnectionID(businessConnectionID string) *ReadBusinessMessageParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithChatID adds chat ID parameter
func (p *ReadBusinessMessageParams) WithChatID(chatID int64) *ReadBusinessMessageParams {
	p.ChatID = chatID
	return p
}

// WithMessageID adds message ID parameter
func (p *ReadBusinessMessageParams) WithMessageID(messageID int) *ReadBusinessMessageParams {
	p.MessageID = messageID
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *DeleteBusinessMessagesParams) WithBusinessConnectionID(businessConnectionID string,
) *DeleteBusinessMessagesParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithMessageIDs adds message ids parameter
func (p *DeleteBusinessMessagesParams) WithMessageIDs(messageIDs ...int) *DeleteBusinessMessagesParams {
	p.MessageIDs = messageIDs
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *SetBusinessAccountNameParams) WithBusinessConnectionID(businessConnectionID string,
) *SetBusinessAccountNameParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithFirstName adds first name parameter
func (p *SetBusinessAccountNameParams) WithFirstName(firstName string) *SetBusinessAccountNameParams {
	p.FirstName = firstName
	return p
}

// WithLastName adds last name parameter
func (p *SetBusinessAccountNameParams) WithLastName(lastName string) *SetBusinessAccountNameParams {
	p.LastName = lastName
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *SetBusinessAccountUsernameParams) WithBusinessConnectionID(businessConnectionID string,
) *SetBusinessAccountUsernameParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithUsername adds username parameter
func (p *SetBusinessAccountUsernameParams) WithUsername(username string) *SetBusinessAccountUsernameParams {
	p.Username = username
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *SetBusinessAccountBioParams) WithBusinessConnectionID(businessConnectionID string,
) *SetBusinessAccountBioParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithBio adds bio parameter
func (p *SetBusinessAccountBioParams) WithBio(bio string) *SetBusinessAccountBioParams {
	p.Bio = bio
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *SetBusinessAccountProfilePhotoParams) WithBusinessConnectionID(businessConnectionID string,
) *SetBusinessAccountProfilePhotoParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithPhoto adds photo parameter
func (p *SetBusinessAccountProfilePhotoParams) WithPhoto(photo InputProfilePhoto,
) *SetBusinessAccountProfilePhotoParams {
	p.Photo = photo
	return p
}

// WithIsPublic adds is public parameter
func (p *SetBusinessAccountProfilePhotoParams) WithIsPublic() *SetBusinessAccountProfilePhotoParams {
	p.IsPublic = true
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *RemoveBusinessAccountProfilePhotoParams) WithBusinessConnectionID(businessConnectionID string,
) *RemoveBusinessAccountProfilePhotoParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithIsPublic adds is public parameter
func (p *RemoveBusinessAccountProfilePhotoParams) WithIsPublic() *RemoveBusinessAccountProfilePhotoParams {
	p.IsPublic = true
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *SetBusinessAccountGiftSettingsParams) WithBusinessConnectionID(businessConnectionID string,
) *SetBusinessAccountGiftSettingsParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithShowGiftButton adds show gift button parameter
func (p *SetBusinessAccountGiftSettingsParams) WithShowGiftButton() *SetBusinessAccountGiftSettingsParams {
	p.ShowGiftButton = true
	return p
}

// WithAcceptedGiftTypes adds accepted gift types parameter
func (p *SetBusinessAccountGiftSettingsParams) WithAcceptedGiftTypes(acceptedGiftTypes AcceptedGiftTypes,
) *SetBusinessAccountGiftSettingsParams {
	p.AcceptedGiftTypes = acceptedGiftTypes
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *GetBusinessAccountStarBalanceParams) WithBusinessConnectionID(businessConnectionID string,
) *GetBusinessAccountStarBalanceParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *TransferBusinessAccountStarsParams) WithBusinessConnectionID(businessConnectionID string,
) *TransferBusinessAccountStarsParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithStarCount adds star count parameter
func (p *TransferBusinessAccountStarsParams) WithStarCount(starCount int) *TransferBusinessAccountStarsParams {
	p.StarCount = starCount
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *GetBusinessAccountGiftsParams) WithBusinessConnectionID(businessConnectionID string,
) *GetBusinessAccountGiftsParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithExcludeUnsaved adds exclude unsaved parameter
func (p *GetBusinessAccountGiftsParams) WithExcludeUnsaved() *GetBusinessAccountGiftsParams {
	p.ExcludeUnsaved = true
	return p
}

// WithExcludeSaved adds exclude saved parameter
func (p *GetBusinessAccountGiftsParams) WithExcludeSaved() *GetBusinessAccountGiftsParams {
	p.ExcludeSaved = true
	return p
}

// WithExcludeUnlimited adds exclude unlimited parameter
func (p *GetBusinessAccountGiftsParams) WithExcludeUnlimited() *GetBusinessAccountGiftsParams {
	p.ExcludeUnlimited = true
	return p
}

// WithExcludeLimitedUpgradable adds exclude limited upgradable parameter
func (p *GetBusinessAccountGiftsParams) WithExcludeLimitedUpgradable() *GetBusinessAccountGiftsParams {
	p.ExcludeLimitedUpgradable = true
	return p
}

// WithExcludeLimitedNonUpgradable adds exclude limited non upgradable parameter
func (p *GetBusinessAccountGiftsParams) WithExcludeLimitedNonUpgradable() *GetBusinessAccountGiftsParams {
	p.ExcludeLimitedNonUpgradable = true
	return p
}

// WithExcludeUnique adds exclude unique parameter
func (p *GetBusinessAccountGiftsParams) WithExcludeUnique() *GetBusinessAccountGiftsParams {
	p.ExcludeUnique = true
	return p
}

// WithExcludeFromBlockchain adds exclude from blockchain parameter
func (p *GetBusinessAccountGiftsParams) WithExcludeFromBlockchain() *GetBusinessAccountGiftsParams {
	p.ExcludeFromBlockchain = true
	return p
}

// WithSortByPrice adds sort by price parameter
func (p *GetBusinessAccountGiftsParams) WithSortByPrice() *GetBusinessAccountGiftsParams {
	p.SortByPrice = true
	return p
}

// WithOffset adds offset parameter
func (p *GetBusinessAccountGiftsParams) WithOffset(offset string) *GetBusinessAccountGiftsParams {
	p.Offset = offset
	return p
}

// WithLimit adds limit parameter
func (p *GetBusinessAccountGiftsParams) WithLimit(limit int) *GetBusinessAccountGiftsParams {
	p.Limit = limit
	return p
}

// WithUserID adds user ID parameter
func (p *GetUserGiftsParams) WithUserID(userID int64) *GetUserGiftsParams {
	p.UserID = userID
	return p
}

// WithExcludeUnlimited adds exclude unlimited parameter
func (p *GetUserGiftsParams) WithExcludeUnlimited() *GetUserGiftsParams {
	p.ExcludeUnlimited = true
	return p
}

// WithExcludeLimitedUpgradable adds exclude limited upgradable parameter
func (p *GetUserGiftsParams) WithExcludeLimitedUpgradable() *GetUserGiftsParams {
	p.ExcludeLimitedUpgradable = true
	return p
}

// WithExcludeLimitedNonUpgradable adds exclude limited non upgradable parameter
func (p *GetUserGiftsParams) WithExcludeLimitedNonUpgradable() *GetUserGiftsParams {
	p.ExcludeLimitedNonUpgradable = true
	return p
}

// WithExcludeFromBlockchain adds exclude from blockchain parameter
func (p *GetUserGiftsParams) WithExcludeFromBlockchain() *GetUserGiftsParams {
	p.ExcludeFromBlockchain = true
	return p
}

// WithExcludeUnique adds exclude unique parameter
func (p *GetUserGiftsParams) WithExcludeUnique() *GetUserGiftsParams {
	p.ExcludeUnique = true
	return p
}

// WithSortByPrice adds sort by price parameter
func (p *GetUserGiftsParams) WithSortByPrice() *GetUserGiftsParams {
	p.SortByPrice = true
	return p
}

// WithOffset adds offset parameter
func (p *GetUserGiftsParams) WithOffset(offset string) *GetUserGiftsParams {
	p.Offset = offset
	return p
}

// WithLimit adds limit parameter
func (p *GetUserGiftsParams) WithLimit(limit int) *GetUserGiftsParams {
	p.Limit = limit
	return p
}

// WithChatID adds chat ID parameter
func (p *GetChatGiftsParams) WithChatID(chatID ChatID) *GetChatGiftsParams {
	p.ChatID = chatID
	return p
}

// WithExcludeUnsaved adds exclude unsaved parameter
func (p *GetChatGiftsParams) WithExcludeUnsaved() *GetChatGiftsParams {
	p.ExcludeUnsaved = true
	return p
}

// WithExcludeSaved adds exclude saved parameter
func (p *GetChatGiftsParams) WithExcludeSaved() *GetChatGiftsParams {
	p.ExcludeSaved = true
	return p
}

// WithExcludeUnlimited adds exclude unlimited parameter
func (p *GetChatGiftsParams) WithExcludeUnlimited() *GetChatGiftsParams {
	p.ExcludeUnlimited = true
	return p
}

// WithExcludeLimitedUpgradable adds exclude limited upgradable parameter
func (p *GetChatGiftsParams) WithExcludeLimitedUpgradable() *GetChatGiftsParams {
	p.ExcludeLimitedUpgradable = true
	return p
}

// WithExcludeLimitedNonUpgradable adds exclude limited non upgradable parameter
func (p *GetChatGiftsParams) WithExcludeLimitedNonUpgradable() *GetChatGiftsParams {
	p.ExcludeLimitedNonUpgradable = true
	return p
}

// WithExcludeFromBlockchain adds exclude from blockchain parameter
func (p *GetChatGiftsParams) WithExcludeFromBlockchain() *GetChatGiftsParams {
	p.ExcludeFromBlockchain = true
	return p
}

// WithExcludeUnique adds exclude unique parameter
func (p *GetChatGiftsParams) WithExcludeUnique() *GetChatGiftsParams {
	p.ExcludeUnique = true
	return p
}

// WithSortByPrice adds sort by price parameter
func (p *GetChatGiftsParams) WithSortByPrice() *GetChatGiftsParams {
	p.SortByPrice = true
	return p
}

// WithOffset adds offset parameter
func (p *GetChatGiftsParams) WithOffset(offset string) *GetChatGiftsParams {
	p.Offset = offset
	return p
}

// WithLimit adds limit parameter
func (p *GetChatGiftsParams) WithLimit(limit int) *GetChatGiftsParams {
	p.Limit = limit
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *ConvertGiftToStarsParams) WithBusinessConnectionID(businessConnectionID string) *ConvertGiftToStarsParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithOwnedGiftID adds owned gift ID parameter
func (p *ConvertGiftToStarsParams) WithOwnedGiftID(ownedGiftID string) *ConvertGiftToStarsParams {
	p.OwnedGiftID = ownedGiftID
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *UpgradeGiftParams) WithBusinessConnectionID(businessConnectionID string) *UpgradeGiftParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithOwnedGiftID adds owned gift ID parameter
func (p *UpgradeGiftParams) WithOwnedGiftID(ownedGiftID string) *UpgradeGiftParams {
	p.OwnedGiftID = ownedGiftID
	return p
}

// WithKeepOriginalDetails adds keep original details parameter
func (p *UpgradeGiftParams) WithKeepOriginalDetails() *UpgradeGiftParams {
	p.KeepOriginalDetails = true
	return p
}

// WithStarCount adds star count parameter
func (p *UpgradeGiftParams) WithStarCount(starCount int) *UpgradeGiftParams {
	p.StarCount = starCount
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *TransferGiftParams) WithBusinessConnectionID(businessConnectionID string) *TransferGiftParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithOwnedGiftID adds owned gift ID parameter
func (p *TransferGiftParams) WithOwnedGiftID(ownedGiftID string) *TransferGiftParams {
	p.OwnedGiftID = ownedGiftID
	return p
}

// WithNewOwnerChatID adds new owner chat ID parameter
func (p *TransferGiftParams) WithNewOwnerChatID(newOwnerChatID int64) *TransferGiftParams {
	p.NewOwnerChatID = newOwnerChatID
	return p
}

// WithStarCount adds star count parameter
func (p *TransferGiftParams) WithStarCount(starCount int) *TransferGiftParams {
	p.StarCount = starCount
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *PostStoryParams) WithBusinessConnectionID(businessConnectionID string) *PostStoryParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithContent adds content parameter
func (p *PostStoryParams) WithContent(content InputStoryContent) *PostStoryParams {
	p.Content = content
	return p
}

// WithActivePeriod adds active period parameter
func (p *PostStoryParams) WithActivePeriod(activePeriod int) *PostStoryParams {
	p.ActivePeriod = activePeriod
	return p
}

// WithCaption adds caption parameter
func (p *PostStoryParams) WithCaption(caption string) *PostStoryParams {
	p.Caption = caption
	return p
}

// WithParseMode adds parse mode parameter
func (p *PostStoryParams) WithParseMode(parseMode string) *PostStoryParams {
	p.ParseMode = parseMode
	return p
}

// WithCaptionEntities adds caption entities parameter
func (p *PostStoryParams) WithCaptionEntities(captionEntities ...MessageEntity) *PostStoryParams {
	p.CaptionEntities = captionEntities
	return p
}

// WithAreas adds areas parameter
func (p *PostStoryParams) WithAreas(areas ...StoryArea) *PostStoryParams {
	p.Areas = areas
	return p
}

// WithPostToChatPage adds post to chat page parameter
func (p *PostStoryParams) WithPostToChatPage() *PostStoryParams {
	p.PostToChatPage = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *PostStoryParams) WithProtectContent() *PostStoryParams {
	p.ProtectContent = true
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *RepostStoryParams) WithBusinessConnectionID(businessConnectionID string) *RepostStoryParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithFromChatID adds from chat ID parameter
func (p *RepostStoryParams) WithFromChatID(fromChatID int) *RepostStoryParams {
	p.FromChatID = fromChatID
	return p
}

// WithFromStoryID adds from story ID parameter
func (p *RepostStoryParams) WithFromStoryID(fromStoryID int) *RepostStoryParams {
	p.FromStoryID = fromStoryID
	return p
}

// WithActivePeriod adds active period parameter
func (p *RepostStoryParams) WithActivePeriod(activePeriod int) *RepostStoryParams {
	p.ActivePeriod = activePeriod
	return p
}

// WithPostToChatPage adds post to chat page parameter
func (p *RepostStoryParams) WithPostToChatPage() *RepostStoryParams {
	p.PostToChatPage = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *RepostStoryParams) WithProtectContent() *RepostStoryParams {
	p.ProtectContent = true
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *EditStoryParams) WithBusinessConnectionID(businessConnectionID string) *EditStoryParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithStoryID adds story ID parameter
func (p *EditStoryParams) WithStoryID(storyID int) *EditStoryParams {
	p.StoryID = storyID
	return p
}

// WithContent adds content parameter
func (p *EditStoryParams) WithContent(content InputStoryContent) *EditStoryParams {
	p.Content = content
	return p
}

// WithCaption adds caption parameter
func (p *EditStoryParams) WithCaption(caption string) *EditStoryParams {
	p.Caption = caption
	return p
}

// WithParseMode adds parse mode parameter
func (p *EditStoryParams) WithParseMode(parseMode string) *EditStoryParams {
	p.ParseMode = parseMode
	return p
}

// WithCaptionEntities adds caption entities parameter
func (p *EditStoryParams) WithCaptionEntities(captionEntities ...MessageEntity) *EditStoryParams {
	p.CaptionEntities = captionEntities
	return p
}

// WithAreas adds areas parameter
func (p *EditStoryParams) WithAreas(areas ...StoryArea) *EditStoryParams {
	p.Areas = areas
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *DeleteStoryParams) WithBusinessConnectionID(businessConnectionID string) *DeleteStoryParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithStoryID adds story ID parameter
func (p *DeleteStoryParams) WithStoryID(storyID int) *DeleteStoryParams {
	p.StoryID = storyID
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *EditMessageTextParams) WithBusinessConnectionID(businessConnectionID string) *EditMessageTextParams {
	p.BusinessConnectionID = businessConnectionID
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

// WithBusinessConnectionID adds business connection ID parameter
func (p *EditMessageCaptionParams) WithBusinessConnectionID(businessConnectionID string) *EditMessageCaptionParams {
	p.BusinessConnectionID = businessConnectionID
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

// WithBusinessConnectionID adds business connection ID parameter
func (p *EditMessageMediaParams) WithBusinessConnectionID(businessConnectionID string) *EditMessageMediaParams {
	p.BusinessConnectionID = businessConnectionID
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

// WithBusinessConnectionID adds business connection ID parameter
func (p *EditMessageLiveLocationParams) WithBusinessConnectionID(businessConnectionID string,
) *EditMessageLiveLocationParams {
	p.BusinessConnectionID = businessConnectionID
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

// WithLatitude adds latitude parameter
func (p *EditMessageLiveLocationParams) WithLatitude(latitude float64) *EditMessageLiveLocationParams {
	p.Latitude = latitude
	return p
}

// WithLongitude adds longitude parameter
func (p *EditMessageLiveLocationParams) WithLongitude(longitude float64) *EditMessageLiveLocationParams {
	p.Longitude = longitude
	return p
}

// WithLivePeriod adds live period parameter
func (p *EditMessageLiveLocationParams) WithLivePeriod(livePeriod int) *EditMessageLiveLocationParams {
	p.LivePeriod = livePeriod
	return p
}

// WithHorizontalAccuracy adds horizontal accuracy parameter
func (p *EditMessageLiveLocationParams) WithHorizontalAccuracy(horizontalAccuracy float64,
) *EditMessageLiveLocationParams {
	p.HorizontalAccuracy = horizontalAccuracy
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

// WithBusinessConnectionID adds business connection ID parameter
func (p *StopMessageLiveLocationParams) WithBusinessConnectionID(businessConnectionID string,
) *StopMessageLiveLocationParams {
	p.BusinessConnectionID = businessConnectionID
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

// WithBusinessConnectionID adds business connection ID parameter
func (p *EditMessageChecklistParams) WithBusinessConnectionID(businessConnectionID string) *EditMessageChecklistParams {
	p.BusinessConnectionID = businessConnectionID
	return p
}

// WithChatID adds chat ID parameter
func (p *EditMessageChecklistParams) WithChatID(chatID int64) *EditMessageChecklistParams {
	p.ChatID = chatID
	return p
}

// WithMessageID adds message ID parameter
func (p *EditMessageChecklistParams) WithMessageID(messageID int) *EditMessageChecklistParams {
	p.MessageID = messageID
	return p
}

// WithChecklist adds checklist parameter
func (p *EditMessageChecklistParams) WithChecklist(checklist InputChecklist) *EditMessageChecklistParams {
	p.Checklist = checklist
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *EditMessageChecklistParams) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *EditMessageChecklistParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithBusinessConnectionID adds business connection ID parameter
func (p *EditMessageReplyMarkupParams) WithBusinessConnectionID(businessConnectionID string,
) *EditMessageReplyMarkupParams {
	p.BusinessConnectionID = businessConnectionID
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

// WithBusinessConnectionID adds business connection ID parameter
func (p *StopPollParams) WithBusinessConnectionID(businessConnectionID string) *StopPollParams {
	p.BusinessConnectionID = businessConnectionID
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
func (p *ApproveSuggestedPostParams) WithChatID(chatID int64) *ApproveSuggestedPostParams {
	p.ChatID = chatID
	return p
}

// WithMessageID adds message ID parameter
func (p *ApproveSuggestedPostParams) WithMessageID(messageID int) *ApproveSuggestedPostParams {
	p.MessageID = messageID
	return p
}

// WithSendDate adds send date parameter
func (p *ApproveSuggestedPostParams) WithSendDate(sendDate int64) *ApproveSuggestedPostParams {
	p.SendDate = sendDate
	return p
}

// WithChatID adds chat ID parameter
func (p *DeclineSuggestedPostParams) WithChatID(chatID int64) *DeclineSuggestedPostParams {
	p.ChatID = chatID
	return p
}

// WithMessageID adds message ID parameter
func (p *DeclineSuggestedPostParams) WithMessageID(messageID int) *DeclineSuggestedPostParams {
	p.MessageID = messageID
	return p
}

// WithComment adds comment parameter
func (p *DeclineSuggestedPostParams) WithComment(comment string) *DeclineSuggestedPostParams {
	p.Comment = comment
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

// WithDirectMessagesTopicID adds direct messages topic ID parameter
func (p *SendStickerParams) WithDirectMessagesTopicID(directMessagesTopicID int) *SendStickerParams {
	p.DirectMessagesTopicID = directMessagesTopicID
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

// WithAllowPaidBroadcast adds allow paid broadcast parameter
func (p *SendStickerParams) WithAllowPaidBroadcast() *SendStickerParams {
	p.AllowPaidBroadcast = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendStickerParams) WithMessageEffectID(messageEffectID string) *SendStickerParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithSuggestedPostParameters adds suggested post parameters parameter
func (p *SendStickerParams) WithSuggestedPostParameters(suggestedPostParameters *SuggestedPostParameters,
) *SendStickerParams {
	p.SuggestedPostParameters = suggestedPostParameters
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

// WithUserID adds user ID parameter
func (p *UploadStickerFileParams) WithUserID(userID int64) *UploadStickerFileParams {
	p.UserID = userID
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

// WithUserID adds user ID parameter
func (p *CreateNewStickerSetParams) WithUserID(userID int64) *CreateNewStickerSetParams {
	p.UserID = userID
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

// WithUserID adds user ID parameter
func (p *AddStickerToSetParams) WithUserID(userID int64) *AddStickerToSetParams {
	p.UserID = userID
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

// WithUserID adds user ID parameter
func (p *ReplaceStickerInSetParams) WithUserID(userID int64) *ReplaceStickerInSetParams {
	p.UserID = userID
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

// WithUserID adds user ID parameter
func (p *SetStickerSetThumbnailParams) WithUserID(userID int64) *SetStickerSetThumbnailParams {
	p.UserID = userID
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

// WithUserID adds user ID parameter
func (p *SavePreparedInlineMessageParams) WithUserID(userID int64) *SavePreparedInlineMessageParams {
	p.UserID = userID
	return p
}

// WithResult adds result parameter
func (p *SavePreparedInlineMessageParams) WithResult(result InlineQueryResult) *SavePreparedInlineMessageParams {
	p.Result = result
	return p
}

// WithAllowUserChats adds allow user chats parameter
func (p *SavePreparedInlineMessageParams) WithAllowUserChats() *SavePreparedInlineMessageParams {
	p.AllowUserChats = true
	return p
}

// WithAllowBotChats adds allow bot chats parameter
func (p *SavePreparedInlineMessageParams) WithAllowBotChats() *SavePreparedInlineMessageParams {
	p.AllowBotChats = true
	return p
}

// WithAllowGroupChats adds allow group chats parameter
func (p *SavePreparedInlineMessageParams) WithAllowGroupChats() *SavePreparedInlineMessageParams {
	p.AllowGroupChats = true
	return p
}

// WithAllowChannelChats adds allow channel chats parameter
func (p *SavePreparedInlineMessageParams) WithAllowChannelChats() *SavePreparedInlineMessageParams {
	p.AllowChannelChats = true
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

// WithDirectMessagesTopicID adds direct messages topic ID parameter
func (p *SendInvoiceParams) WithDirectMessagesTopicID(directMessagesTopicID int) *SendInvoiceParams {
	p.DirectMessagesTopicID = directMessagesTopicID
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

// WithAllowPaidBroadcast adds allow paid broadcast parameter
func (p *SendInvoiceParams) WithAllowPaidBroadcast() *SendInvoiceParams {
	p.AllowPaidBroadcast = true
	return p
}

// WithMessageEffectID adds message effect ID parameter
func (p *SendInvoiceParams) WithMessageEffectID(messageEffectID string) *SendInvoiceParams {
	p.MessageEffectID = messageEffectID
	return p
}

// WithSuggestedPostParameters adds suggested post parameters parameter
func (p *SendInvoiceParams) WithSuggestedPostParameters(suggestedPostParameters *SuggestedPostParameters,
) *SendInvoiceParams {
	p.SuggestedPostParameters = suggestedPostParameters
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

// WithBusinessConnectionID adds business connection ID parameter
func (p *CreateInvoiceLinkParams) WithBusinessConnectionID(businessConnectionID string) *CreateInvoiceLinkParams {
	p.BusinessConnectionID = businessConnectionID
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

// WithSubscriptionPeriod adds subscription period parameter
func (p *CreateInvoiceLinkParams) WithSubscriptionPeriod(subscriptionPeriod int64) *CreateInvoiceLinkParams {
	p.SubscriptionPeriod = subscriptionPeriod
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

// WithOffset adds offset parameter
func (p *GetStarTransactionsParams) WithOffset(offset int) *GetStarTransactionsParams {
	p.Offset = offset
	return p
}

// WithLimit adds limit parameter
func (p *GetStarTransactionsParams) WithLimit(limit int) *GetStarTransactionsParams {
	p.Limit = limit
	return p
}

// WithUserID adds user ID parameter
func (p *RefundStarPaymentParams) WithUserID(userID int64) *RefundStarPaymentParams {
	p.UserID = userID
	return p
}

// WithTelegramPaymentChargeID adds telegram payment charge ID parameter
func (p *RefundStarPaymentParams) WithTelegramPaymentChargeID(telegramPaymentChargeID string) *RefundStarPaymentParams {
	p.TelegramPaymentChargeID = telegramPaymentChargeID
	return p
}

// WithUserID adds user ID parameter
func (p *EditUserStarSubscriptionParams) WithUserID(userID int64) *EditUserStarSubscriptionParams {
	p.UserID = userID
	return p
}

// WithTelegramPaymentChargeID adds telegram payment charge ID parameter
func (p *EditUserStarSubscriptionParams) WithTelegramPaymentChargeID(telegramPaymentChargeID string,
) *EditUserStarSubscriptionParams {
	p.TelegramPaymentChargeID = telegramPaymentChargeID
	return p
}

// WithIsCanceled adds is canceled parameter
func (p *EditUserStarSubscriptionParams) WithIsCanceled() *EditUserStarSubscriptionParams {
	p.IsCanceled = true
	return p
}

// WithUserID adds user ID parameter
func (p *SetPassportDataErrorsParams) WithUserID(userID int64) *SetPassportDataErrorsParams {
	p.UserID = userID
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

// WithChatID adds chat ID parameter
func (p *SendGameParams) WithChatID(chatID int64) *SendGameParams {
	p.ChatID = chatID
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

// WithAllowPaidBroadcast adds allow paid broadcast parameter
func (p *SendGameParams) WithAllowPaidBroadcast() *SendGameParams {
	p.AllowPaidBroadcast = true
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

// WithUserID adds user ID parameter
func (p *SetGameScoreParams) WithUserID(userID int64) *SetGameScoreParams {
	p.UserID = userID
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

// WithChatID adds chat ID parameter
func (p *SetGameScoreParams) WithChatID(chatID int64) *SetGameScoreParams {
	p.ChatID = chatID
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

// WithUserID adds user ID parameter
func (p *GetGameHighScoresParams) WithUserID(userID int64) *GetGameHighScoresParams {
	p.UserID = userID
	return p
}

// WithChatID adds chat ID parameter
func (p *GetGameHighScoresParams) WithChatID(chatID int64) *GetGameHighScoresParams {
	p.ChatID = chatID
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
