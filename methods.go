package telego

// getUpdates - Use this method to receive incoming updates using long polling (wiki 
// (https://en.wikipedia.org/wiki/Push_technology#Long_polling)). An Array of Update (#update) objects is 
// returned.
// 
// Parameters:
// offset (optional) - Identifier of the first update to be returned. Must be 
// greater by one than the highest among the identifiers of previously received updates. By default, updates 
// starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as 
// getUpdates (#getupdates) is called with an offset higher than its update_id. The negative offset can be 
// specified to retrieve updates starting from -offset update from the end of the updates queue. All previous 
// updates will forgotten.
// 
// limit (optional) - Limits the number of updates to be retrieved. Values between 
// 1-100 are accepted. Defaults to 100.
// 
// timeout (optional) - Timeout in seconds for long polling. Defaults 
// to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes 
// only.
// 
// allowedUpdates (optional) - A JSON-serialized list of the update types you want your bot to 
// receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only 
// receive updates of these types. See Update (#update) for a complete list of available update types. Specify 
// an empty list to receive all update types except chat_member (default). If not specified, the previous 
// setting will be used.Please note that this parameter doesn't affect updates created before the call to the 
// getUpdates, so unwanted updates may be received for a short period of time.
// 
//  
func getUpdates(offset int, limit int, timeout int, allowedUpdates []string){
	
}

// setWebhook - Use this method to specify a url and receive incoming updates via an outgoing webhook. Whenever 
// there is an update for the bot, we will send an HTTPS POST request to the specified url, containing a 
// JSON-serialized Update (#update). In case of an unsuccessful request, we will give up after a reasonable 
// amount of attempts. Returns True on success.
// 
// Parameters:
//  
func setWebhook(){
	
}

// deleteWebhook - Use this method to remove webhook integration if you decide to switch back to getUpdates 
// (#getupdates). Returns True on success.
// 
// Parameters:
// dropPendingUpdates (optional) - Pass True to drop 
// all pending updates
// 
//  
func deleteWebhook(dropPendingUpdates bool){
	
}

// getWebhookInfo - Use this method to get current webhook status. Requires no parameters. On success, returns a 
// WebhookInfo (#webhookinfo) object. If the bot is using getUpdates (#getupdates), will return an object with 
// the url field empty.
// 
// Parameters:
//  
func getWebhookInfo(){
	
}

// getMe - A simple method for testing your bot's auth token. Requires no parameters. Returns basic information 
// about the bot in form of a User (#user) object.
// 
// Parameters:
//  
func getMe(){
	
}

// logOut - Use this method to log out from the cloud Bot API server before launching the bot locally. You must 
// log out the bot before running it locally, otherwise there is no guarantee that the bot will receive updates. 
// After a successful call, you can immediately log in on a local server, but will not be able to log in back to 
// the cloud Bot API server for 10 minutes. Returns True on success. Requires no parameters.
// 
// Parameters:
//  
func logOut(){
	
}

// close - Use this method to close the bot instance before moving it from one local server to another. You need 
// to delete the webhook before calling this method to ensure that the bot isn't launched again after server 
// restart. The method will return error 429 in the first 10 minutes after the bot is launched. Returns True on 
// success. Requires no parameters.
// 
// Parameters:
//  
func close(){
	
}

// sendMessage - Use this method to send text messages. On success, the sent Message (#message) is 
// returned.
// 
// Parameters:
// chatId - Unique identifier for the target chat or username of the target channel 
// (in the format @channelusername)
// 
// text - Text of the message to be sent, 1-4096 characters after entities 
// parsing
// 
// parseMode (optional) - Mode for parsing entities in the message text. See formatting options 
// (#formatting-options) for more details.
// 
// entities (optional) - List of special entities that appear in 
// message text, which can be specified instead of parse_mode
// 
// disableWebPagePreview (optional) - Disables 
// link previews for links in this message
// 
// disableNotification (optional) - Sends the message silently 
// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no 
// sound.
// 
// replyToMessageId (optional) - If the message is a reply, ID of the original 
// message
// 
// allowSendingWithoutReply (optional) - Pass True, if the message should be sent even if the 
// specified replied-to message is not found
// 
// replyMarkup (optional) - Additional interface options. A 
// JSON-serialized object for an inline keyboard 
// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard 
// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from 
// the user.
// 
//  
func sendMessage(chatId IntOrStringChatID, text string, parseMode string, entities []MessageEntity, disableWebPagePreview bool, disableNotification bool, replyToMessageId int, allowSendingWithoutReply bool, replyMarkup *InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply){
	
}

// forwardMessage - Use this method to forward messages of any kind. Service messages can't be forwarded. On 
// success, the sent Message (#message) is returned.
// 
// Parameters:
// chatId - Unique identifier for the 
// target chat or username of the target channel (in the format @channelusername)
// 
// fromChatId - Unique 
// identifier for the chat where the original message was sent (or channel username in the format 
// @channelusername)
// 
// disableNotification (optional) - Sends the message silently 
// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no 
// sound.
// 
// messageId - Message identifier in the chat specified in from_chat_id
// 
//  
func forwardMessage(chatId IntOrStringChatID, fromChatId IntOrStringChatID, disableNotification bool, messageId int){
	
}

// copyMessage - Use this method to copy messages of any kind. Service messages and invoice messages can't be 
// copied. The method is analogous to the method forwardMessage (#forwardmessage), but the copied message 
// doesn't have a link to the original message. Returns the MessageId (#messageid) of the sent message on 
// success.
// 
// Parameters:
// chatId - Unique identifier for the target chat or username of the target channel 
// (in the format @channelusername)
// 
// fromChatId - Unique identifier for the chat where the original message 
// was sent (or channel username in the format @channelusername)
// 
// messageId - Message identifier in the chat 
// specified in from_chat_id
// 
// caption (optional) - New caption for media, 0-1024 characters after entities 
// parsing. If not specified, the original caption is kept
// 
// parseMode (optional) - Mode for parsing entities 
// in the new caption. See formatting options (#formatting-options) for more details.
// 
// captionEntities 
// (optional) - List of special entities that appear in the new caption, which can be specified instead of 
// parse_mode
// 
// disableNotification (optional) - Sends the message silently 
// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no 
// sound.
// 
// replyToMessageId (optional) - If the message is a reply, ID of the original 
// message
// 
// allowSendingWithoutReply (optional) - Pass True, if the message should be sent even if the 
// specified replied-to message is not found
// 
// replyMarkup (optional) - Additional interface options. A 
// JSON-serialized object for an inline keyboard 
// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard 
// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from 
// the user.
// 
//  
func copyMessage(chatId IntOrStringChatID, fromChatId IntOrStringChatID, messageId int, caption string, parseMode string, captionEntities []MessageEntity, disableNotification bool, replyToMessageId int, allowSendingWithoutReply bool, replyMarkup *InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply){
	
}

// sendPhoto - Use this method to send photos. On success, the sent Message (#message) is 
// returned.
// 
// Parameters:
// chatId - Unique identifier for the target chat or username of the target channel 
// (in the format @channelusername)
// 
// photo - Photo to send. Pass a file_id as String to send a photo that 
// exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a photo from 
// the Internet, or upload a new photo using multipart/form-data. The photo must be at most 10 MB in size. The 
// photo's width and height must not exceed 10000 in total. Width and height ratio must be at most 20. More info 
// on Sending Files » (#sending-files)
// 
// caption (optional) - Photo caption (may also be used when resending 
// photos by file_id), 0-1024 characters after entities parsing
// 
// parseMode (optional) - Mode for parsing 
// entities in the photo caption. See formatting options (#formatting-options) for more 
// details.
// 
// captionEntities (optional) - List of special entities that appear in the caption, which can be 
// specified instead of parse_mode
// 
// disableNotification (optional) - Sends the message silently 
// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no 
// sound.
// 
// replyToMessageId (optional) - If the message is a reply, ID of the original 
// message
// 
// allowSendingWithoutReply (optional) - Pass True, if the message should be sent even if the 
// specified replied-to message is not found
// 
// replyMarkup (optional) - Additional interface options. A 
// JSON-serialized object for an inline keyboard 
// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard 
// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from 
// the user.
// 
//  
func sendPhoto(chatId IntOrStringChatID, photo InputFile or String, caption string, parseMode string, captionEntities []MessageEntity, disableNotification bool, replyToMessageId int, allowSendingWithoutReply bool, replyMarkup *InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply){
	
}

// sendAudio - Use this method to send audio files, if you want Telegram clients to display them in the music 
// player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message (#message) is returned. 
// Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the 
// future.
// 
// Parameters:
//  
func sendAudio(){
	
}

// sendDocument - Use this method to send general files. On success, the sent Message (#message) is returned. 
// Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the 
// future.
// 
// Parameters:
// chatId - Unique identifier for the target chat or username of the target channel 
// (in the format @channelusername)
// 
// document - File to send. Pass a file_id as String to send a file that 
// exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from 
// the Internet, or upload a new one using multipart/form-data. More info on Sending Files » 
// (#sending-files)
// 
// thumb (optional) - Thumbnail of the file sent; can be ignored if thumbnail generation 
// for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A 
// thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using 
// multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass 
// “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under 
// <file_attach_name>. More info on Sending Files » (#sending-files)
// 
// caption (optional) - Document caption 
// (may also be used when resending documents by file_id), 0-1024 characters after entities 
// parsing
// 
// parseMode (optional) - Mode for parsing entities in the document caption. See formatting options 
// (#formatting-options) for more details.
// 
// captionEntities (optional) - List of special entities that 
// appear in the caption, which can be specified instead of parse_mode
// 
// disableContentTypeDetection 
// (optional) - Disables automatic server-side content type detection for files uploaded using 
// multipart/form-data
// 
// disableNotification (optional) - Sends the message silently 
// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no 
// sound.
// 
// replyToMessageId (optional) - If the message is a reply, ID of the original 
// message
// 
// allowSendingWithoutReply (optional) - Pass True, if the message should be sent even if the 
// specified replied-to message is not found
// 
// replyMarkup (optional) - Additional interface options. A 
// JSON-serialized object for an inline keyboard 
// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard 
// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from 
// the user.
// 
//  
func sendDocument(chatId IntOrStringChatID, document InputFile or String, thumb *InputFile or String, caption string, parseMode string, captionEntities []MessageEntity, disableContentTypeDetection bool, disableNotification bool, replyToMessageId int, allowSendingWithoutReply bool, replyMarkup *InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply){
	
}

// sendVideo - Use this method to send video files, Telegram clients support mp4 videos (other formats may be 
// sent as Document (#document)). On success, the sent Message (#message) is returned. Bots can currently send 
// video files of up to 50 MB in size, this limit may be changed in the future.
// 
// Parameters:
// chatId - 
// Unique identifier for the target chat or username of the target channel (in the format 
// @channelusername)
// 
// video - Video to send. Pass a file_id as String to send a video that exists on the 
// Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a video from the Internet, 
// or upload a new video using multipart/form-data. More info on Sending Files » (#sending-files)
// 
// duration 
// (optional) - Duration of sent video in seconds
// 
// width (optional) - Video width
// 
// height (optional) - 
// Video height
// 
// thumb (optional) - Thumbnail of the file sent; can be ignored if thumbnail generation for 
// the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A 
// thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using 
// multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass 
// “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under 
// <file_attach_name>. More info on Sending Files » (#sending-files)
// 
// caption (optional) - Video caption 
// (may also be used when resending videos by file_id), 0-1024 characters after entities parsing
// 
// parseMode 
// (optional) - Mode for parsing entities in the video caption. See formatting options (#formatting-options) for 
// more details.
// 
// captionEntities (optional) - List of special entities that appear in the caption, which 
// can be specified instead of parse_mode
// 
// supportsStreaming (optional) - Pass True, if the uploaded video 
// is suitable for streaming
// 
// disableNotification (optional) - Sends the message silently 
// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no 
// sound.
// 
// replyToMessageId (optional) - If the message is a reply, ID of the original 
// message
// 
// allowSendingWithoutReply (optional) - Pass True, if the message should be sent even if the 
// specified replied-to message is not found
// 
// replyMarkup (optional) - Additional interface options. A 
// JSON-serialized object for an inline keyboard 
// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard 
// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from 
// the user.
// 
//  
func sendVideo(chatId IntOrStringChatID, video InputFile or String, duration int, width int, height int, thumb *InputFile or String, caption string, parseMode string, captionEntities []MessageEntity, supportsStreaming bool, disableNotification bool, replyToMessageId int, allowSendingWithoutReply bool, replyMarkup *InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply){
	
}

// sendAnimation - Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On 
// success, the sent Message (#message) is returned. Bots can currently send animation files of up to 50 MB in 
// size, this limit may be changed in the future.
// 
// Parameters:
// chatId - Unique identifier for the target 
// chat or username of the target channel (in the format @channelusername)
// 
// animation - Animation to send. 
// Pass a file_id as String to send an animation that exists on the Telegram servers (recommended), pass an HTTP 
// URL as a String for Telegram to get an animation from the Internet, or upload a new animation using 
// multipart/form-data. More info on Sending Files » (#sending-files)
// 
// duration (optional) - Duration of 
// sent animation in seconds
// 
// width (optional) - Animation width
// 
// height (optional) - Animation 
// height
// 
// thumb (optional) - Thumbnail of the file sent; can be ignored if thumbnail generation for the 
// file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A 
// thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using 
// multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass 
// “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under 
// <file_attach_name>. More info on Sending Files » (#sending-files)
// 
// caption (optional) - Animation 
// caption (may also be used when resending animation by file_id), 0-1024 characters after entities 
// parsing
// 
// parseMode (optional) - Mode for parsing entities in the animation caption. See formatting 
// options (#formatting-options) for more details.
// 
// captionEntities (optional) - List of special entities 
// that appear in the caption, which can be specified instead of parse_mode
// 
// disableNotification (optional) 
// - Sends the message silently (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a 
// notification with no sound.
// 
// replyToMessageId (optional) - If the message is a reply, ID of the original 
// message
// 
// allowSendingWithoutReply (optional) - Pass True, if the message should be sent even if the 
// specified replied-to message is not found
// 
// replyMarkup (optional) - Additional interface options. A 
// JSON-serialized object for an inline keyboard 
// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard 
// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from 
// the user.
// 
//  
func sendAnimation(chatId IntOrStringChatID, animation InputFile or String, duration int, width int, height int, thumb *InputFile or String, caption string, parseMode string, captionEntities []MessageEntity, disableNotification bool, replyToMessageId int, allowSendingWithoutReply bool, replyMarkup *InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply){
	
}

// sendVoice - Use this method to send audio files, if you want Telegram clients to display the file as a 
// playable voice message. For this to work, your audio must be in an .OGG file encoded with OPUS (other formats 
// may be sent as Audio (#audio) or Document (#document)). On success, the sent Message (#message) is returned. 
// Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the 
// future.
// 
// Parameters:
// chatId - Unique identifier for the target chat or username of the target channel 
// (in the format @channelusername)
// 
// voice - Audio file to send. Pass a file_id as String to send a file 
// that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file 
// from the Internet, or upload a new one using multipart/form-data. More info on Sending Files » 
// (#sending-files)
// 
// caption (optional) - Voice message caption, 0-1024 characters after entities 
// parsing
// 
// parseMode (optional) - Mode for parsing entities in the voice message caption. See formatting 
// options (#formatting-options) for more details.
// 
// captionEntities (optional) - List of special entities 
// that appear in the caption, which can be specified instead of parse_mode
// 
// duration (optional) - Duration 
// of the voice message in seconds
// 
// disableNotification (optional) - Sends the message silently 
// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no 
// sound.
// 
// replyToMessageId (optional) - If the message is a reply, ID of the original 
// message
// 
// allowSendingWithoutReply (optional) - Pass True, if the message should be sent even if the 
// specified replied-to message is not found
// 
// replyMarkup (optional) - Additional interface options. A 
// JSON-serialized object for an inline keyboard 
// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard 
// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from 
// the user.
// 
//  
func sendVoice(chatId IntOrStringChatID, voice InputFile or String, caption string, parseMode string, captionEntities []MessageEntity, duration int, disableNotification bool, replyToMessageId int, allowSendingWithoutReply bool, replyMarkup *InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply){
	
}

// sendVideoNote - As of v.4.0 (https://telegram.org/blog/video-messages-and-telescope), Telegram clients 
// support rounded square mp4 videos of up to 1 minute long. Use this method to send video messages. On success, 
// the sent Message (#message) is returned.
// 
// Parameters:
// chatId - Unique identifier for the target chat or 
// username of the target channel (in the format @channelusername)
// 
// videoNote - Video note to send. Pass a 
// file_id as String to send a video note that exists on the Telegram servers (recommended) or upload a new 
// video using multipart/form-data. More info on Sending Files » (#sending-files). Sending video notes by a URL 
// is currently unsupported
// 
// duration (optional) - Duration of sent video in seconds
// 
// length (optional) 
// - Video width and height, i.e. diameter of the video message
// 
// thumb (optional) - Thumbnail of the file 
// sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be 
// in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if 
// the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a 
// new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using 
// multipart/form-data under <file_attach_name>. More info on Sending Files » 
// (#sending-files)
// 
// disableNotification (optional) - Sends the message silently 
// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no 
// sound.
// 
// replyToMessageId (optional) - If the message is a reply, ID of the original 
// message
// 
// allowSendingWithoutReply (optional) - Pass True, if the message should be sent even if the 
// specified replied-to message is not found
// 
// replyMarkup (optional) - Additional interface options. A 
// JSON-serialized object for an inline keyboard 
// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard 
// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from 
// the user.
// 
//  
func sendVideoNote(chatId IntOrStringChatID, videoNote InputFile or String, duration int, length int, thumb *InputFile or String, disableNotification bool, replyToMessageId int, allowSendingWithoutReply bool, replyMarkup *InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply){
	
}

// sendMediaGroup - Use this method to send a group of photos, videos, documents or audios as an album. 
// Documents and audio files can be only grouped in an album with messages of the same type. On success, an 
// array of Messages (#message) that were sent is returned.
// 
// Parameters:
// chatId - Unique identifier for 
// the target chat or username of the target channel (in the format @channelusername)
// 
// media - A 
// JSON-serialized array describing messages to be sent, must include 2-10 items
// 
// disableNotification 
// (optional) - Sends messages silently (https://telegram.org/blog/channels-2-0#silent-messages). Users will 
// receive a notification with no sound.
// 
// replyToMessageId (optional) - If the messages are a reply, ID of 
// the original message
// 
// allowSendingWithoutReply (optional) - Pass True, if the message should be sent even 
// if the specified replied-to message is not found
// 
//  
func sendMediaGroup(chatId IntOrStringChatID, media []InputMediaAudio, InputMediaDocument, InputMediaPhoto and InputMediaVideo, disableNotification bool, replyToMessageId int, allowSendingWithoutReply bool){
	
}

// sendLocation - Use this method to send point on the map. On success, the sent Message (#message) is 
// returned.
// 
// Parameters:
// chatId - Unique identifier for the target chat or username of the target channel 
// (in the format @channelusername)
// 
// latitude - Latitude of the location
// 
// longitude - Longitude of the 
// location
// 
// horizontalAccuracy (optional) - The radius of uncertainty for the location, measured in meters; 
// 0-1500
// 
// livePeriod (optional) - Period in seconds for which the location will be updated (see Live 
// Locations (https://telegram.org/blog/live-locations), should be between 60 and 86400.
// 
// heading (optional) 
// - For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if 
// specified.
// 
// proximityAlertRadius (optional) - For live locations, a maximum distance for proximity alerts 
// about approaching another chat member, in meters. Must be between 1 and 100000 if 
// specified.
// 
// disableNotification (optional) - Sends the message silently 
// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no 
// sound.
// 
// replyToMessageId (optional) - If the message is a reply, ID of the original 
// message
// 
// allowSendingWithoutReply (optional) - Pass True, if the message should be sent even if the 
// specified replied-to message is not found
// 
// replyMarkup (optional) - Additional interface options. A 
// JSON-serialized object for an inline keyboard 
// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard 
// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from 
// the user.
// 
//  
func sendLocation(chatId IntOrStringChatID, latitude float64, longitude float64, horizontalAccuracy float64, livePeriod int, heading int, proximityAlertRadius int, disableNotification bool, replyToMessageId int, allowSendingWithoutReply bool, replyMarkup *InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply){
	
}

// editMessageLiveLocation - Use this method to edit live location messages. A location can be edited until its 
// live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation 
// (#stopmessagelivelocation). On success, if the edited message is not an inline message, the edited Message 
// (#message) is returned, otherwise True is returned.
// 
// Parameters:
// chatId (optional) - Required if 
// inline_message_id is not specified. Unique identifier for the target chat or username of the target channel 
// (in the format @channelusername)
// 
// messageId (optional) - Required if inline_message_id is not specified. 
// Identifier of the message to edit
// 
// inlineMessageId (optional) - Required if chat_id and message_id are 
// not specified. Identifier of the inline message
// 
// latitude - Latitude of new location
// 
// longitude - 
// Longitude of new location
// 
// horizontalAccuracy (optional) - The radius of uncertainty for the location, 
// measured in meters; 0-1500
// 
// heading (optional) - Direction in which the user is moving, in degrees. Must 
// be between 1 and 360 if specified.
// 
// proximityAlertRadius (optional) - Maximum distance for proximity 
// alerts about approaching another chat member, in meters. Must be between 1 and 100000 if 
// specified.
// 
// replyMarkup (optional) - A JSON-serialized object for a new inline keyboard 
// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating).
// 
//  
func editMessageLiveLocation(chatId *IntOrStringChatID, messageId int, inlineMessageId string, latitude float64, longitude float64, horizontalAccuracy float64, heading int, proximityAlertRadius int, replyMarkup *InlineKeyboardMarkup){
	
}

// stopMessageLiveLocation - Use this method to stop updating a live location message before live_period 
// expires. On success, if the message was sent by the bot, the sent Message (#message) is returned, otherwise 
// True is returned.
// 
// Parameters:
// chatId (optional) - Required if inline_message_id is not specified. 
// Unique identifier for the target chat or username of the target channel (in the format 
// @channelusername)
// 
// messageId (optional) - Required if inline_message_id is not specified. Identifier of 
// the message with live location to stop
// 
// inlineMessageId (optional) - Required if chat_id and message_id 
// are not specified. Identifier of the inline message
// 
// replyMarkup (optional) - A JSON-serialized object 
// for a new inline keyboard (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating).
// 
//  
func stopMessageLiveLocation(chatId *IntOrStringChatID, messageId int, inlineMessageId string, replyMarkup *InlineKeyboardMarkup){
	
}

// sendVenue - Use this method to send information about a venue. On success, the sent Message (#message) is 
// returned.
// 
// Parameters:
// chatId - Unique identifier for the target chat or username of the target channel 
// (in the format @channelusername)
// 
// latitude - Latitude of the venue
// 
// longitude - Longitude of the 
// venue
// 
// title - Name of the venue
// 
// address - Address of the venue
// 
// foursquareId (optional) - 
// Foursquare identifier of the venue
// 
// foursquareType (optional) - Foursquare type of the venue, if known. 
// (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or 
// “food/icecream”.)
// 
// googlePlaceId (optional) - Google Places identifier of the 
// venue
// 
// googlePlaceType (optional) - Google Places type of the venue. (See supported types 
// (https://developers.google.com/places/web-service/supported_types).)
// 
// disableNotification (optional) - 
// Sends the message silently (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a 
// notification with no sound.
// 
// replyToMessageId (optional) - If the message is a reply, ID of the original 
// message
// 
// allowSendingWithoutReply (optional) - Pass True, if the message should be sent even if the 
// specified replied-to message is not found
// 
// replyMarkup (optional) - Additional interface options. A 
// JSON-serialized object for an inline keyboard 
// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard 
// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from 
// the user.
// 
//  
func sendVenue(chatId IntOrStringChatID, latitude float64, longitude float64, title string, address string, foursquareId string, foursquareType string, googlePlaceId string, googlePlaceType string, disableNotification bool, replyToMessageId int, allowSendingWithoutReply bool, replyMarkup *InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply){
	
}

// sendContact - Use this method to send phone contacts. On success, the sent Message (#message) is 
// returned.
// 
// Parameters:
// chatId - Unique identifier for the target chat or username of the target channel 
// (in the format @channelusername)
// 
// phoneNumber - Contact's phone number
// 
// firstName - Contact's first 
// name
// 
// lastName (optional) - Contact's last name
// 
// vcard (optional) - Additional data about the contact 
// in the form of a vCard (https://en.wikipedia.org/wiki/VCard), 0-2048 bytes
// 
// disableNotification 
// (optional) - Sends the message silently (https://telegram.org/blog/channels-2-0#silent-messages). Users will 
// receive a notification with no sound.
// 
// replyToMessageId (optional) - If the message is a reply, ID of the 
// original message
// 
// allowSendingWithoutReply (optional) - Pass True, if the message should be sent even if 
// the specified replied-to message is not found
// 
// replyMarkup (optional) - Additional interface options. A 
// JSON-serialized object for an inline keyboard 
// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard 
// (https://core.telegram.org/bots#keyboards), instructions to remove keyboard or to force a reply from the 
// user.
// 
//  
func sendContact(chatId IntOrStringChatID, phoneNumber string, firstName string, lastName string, vcard string, disableNotification bool, replyToMessageId int, allowSendingWithoutReply bool, replyMarkup *InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply){
	
}

// sendPoll - Use this method to send a native poll. On success, the sent Message (#message) is 
// returned.
// 
// Parameters:
// chatId - Unique identifier for the target chat or username of the target channel 
// (in the format @channelusername)
// 
// question - Poll question, 1-300 characters
// 
// options - A 
// JSON-serialized list of answer options, 2-10 strings 1-100 characters each
// 
// isAnonymous (optional) - 
// True, if the poll needs to be anonymous, defaults to True
// 
// type (optional) - Poll type, “quiz” or 
// “regular”, defaults to “regular”
// 
// allowsMultipleAnswers (optional) - True, if the poll allows 
// multiple answers, ignored for polls in quiz mode, defaults to False
// 
// correctOptionId (optional) - 0-based 
// identifier of the correct answer option, required for polls in quiz mode
// 
// explanation (optional) - Text 
// that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 
// characters with at most 2 line feeds after entities parsing
// 
// explanationParseMode (optional) - Mode for 
// parsing entities in the explanation. See formatting options (#formatting-options) for more 
// details.
// 
// explanationEntities (optional) - List of special entities that appear in the poll explanation, 
// which can be specified instead of parse_mode
// 
// openPeriod (optional) - Amount of time in seconds the poll 
// will be active after creation, 5-600. Can't be used together with close_date.
// 
// closeDate (optional) - 
// Point in time (Unix timestamp) when the poll will be automatically closed. Must be at least 5 and no more 
// than 600 seconds in the future. Can't be used together with open_period.
// 
// isClosed (optional) - Pass 
// True, if the poll needs to be immediately closed. This can be useful for poll 
// preview.
// 
// disableNotification (optional) - Sends the message silently 
// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no 
// sound.
// 
// replyToMessageId (optional) - If the message is a reply, ID of the original 
// message
// 
// allowSendingWithoutReply (optional) - Pass True, if the message should be sent even if the 
// specified replied-to message is not found
// 
// replyMarkup (optional) - Additional interface options. A 
// JSON-serialized object for an inline keyboard 
// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard 
// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from 
// the user.
// 
//  
func sendPoll(chatId IntOrStringChatID, question string, options []string, isAnonymous bool, type string, allowsMultipleAnswers bool, correctOptionId int, explanation string, explanationParseMode string, explanationEntities []MessageEntity, openPeriod int, closeDate int, isClosed bool, disableNotification bool, replyToMessageId int, allowSendingWithoutReply bool, replyMarkup *InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply){
	
}

// sendDice - Use this method to send an animated emoji that will display a random value. On success, the sent 
// Message (#message) is returned.
// 
// Parameters:
// chatId - Unique identifier for the target chat or username 
// of the target channel (in the format @channelusername)
// 
// emoji (optional) - Emoji on which the dice throw 
// animation is based. Currently, must be one of “🎲”, “🎯”, “🏀”, “⚽”, “🎳”, or 
// “🎰”. Dice can have values 1-6 for “🎲”, “🎯” and “🎳”, values 1-5 for “🏀” and 
// “⚽”, and values 1-64 for “🎰”. Defaults to “🎲”
// 
// disableNotification (optional) - Sends 
// the message silently (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a 
// notification with no sound.
// 
// replyToMessageId (optional) - If the message is a reply, ID of the original 
// message
// 
// allowSendingWithoutReply (optional) - Pass True, if the message should be sent even if the 
// specified replied-to message is not found
// 
// replyMarkup (optional) - Additional interface options. A 
// JSON-serialized object for an inline keyboard 
// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard 
// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from 
// the user.
// 
//  
func sendDice(chatId IntOrStringChatID, emoji string, disableNotification bool, replyToMessageId int, allowSendingWithoutReply bool, replyMarkup *InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply){
	
}

// sendChatAction - Use this method when you need to tell the user that something is happening on the bot's 
// side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear 
// its typing status). Returns True on success.
// 
// Parameters:
//  
func sendChatAction(){
	
}

// getUserProfilePhotos - Use this method to get a list of profile pictures for a user. Returns a 
// UserProfilePhotos (#userprofilephotos) object.
// 
// Parameters:
// userId - Unique identifier of the target 
// user
// 
// offset (optional) - Sequential number of the first photo to be returned. By default, all photos are 
// returned.
// 
// limit (optional) - Limits the number of photos to be retrieved. Values between 1-100 are 
// accepted. Defaults to 100.
// 
//  
func getUserProfilePhotos(userId int, offset int, limit int){
	
}

// getFile - Use this method to get basic info about a file and prepare it for downloading. For the moment, bots 
// can download files of up to 20MB in size. On success, a File (#file) object is returned. The file can then be 
// downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from 
// the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new 
// one can be requested by calling getFile (#getfile) again.
// 
// Parameters:
// fileId - File identifier to get 
// info about
// 
//  
func getFile(fileId string){
	
}

// banChatMember - Use this method to ban a user in a group, a supergroup or a channel. In the case of 
// supergroups and channels, the user will not be able to return to the chat on their own using invite links, 
// etc., unless unbanned (#unbanchatmember) first. The bot must be an administrator in the chat for this to work 
// and must have the appropriate admin rights. Returns True on success.
// 
// Parameters:
// chatId - Unique 
// identifier for the target group or username of the target supergroup or channel (in the format 
// @channelusername)
// 
// userId - Unique identifier of the target user
// 
// untilDate (optional) - Date when 
// the user will be unbanned, unix time. If user is banned for more than 366 days or less than 30 seconds from 
// the current time they are considered to be banned forever. Applied for supergroups and channels 
// only.
// 
// revokeMessages (optional) - Pass True to delete all messages from the chat for the user that is 
// being removed. If False, the user will be able to see messages in the group that were sent before the user 
// was removed. Always True for supergroups and channels.
// 
//  
func banChatMember(chatId IntOrStringChatID, userId int, untilDate int, revokeMessages bool){
	
}

// unbanChatMember - Use this method to unban a previously banned user in a supergroup or channel. The user will 
// not return to the group or channel automatically, but will be able to join via link, etc. The bot must be an 
// administrator for this to work. By default, this method guarantees that after the call the user is not a 
// member of the chat, but will be able to join it. So if the user is a member of the chat they will also be 
// removed from the chat. If you don't want this, use the parameter only_if_banned. Returns True on 
// success.
// 
// Parameters:
// chatId - Unique identifier for the target group or username of the target 
// supergroup or channel (in the format @username)
// 
// userId - Unique identifier of the target 
// user
// 
// onlyIfBanned (optional) - Do nothing if the user is not banned
// 
//  
func unbanChatMember(chatId IntOrStringChatID, userId int, onlyIfBanned bool){
	
}

// restrictChatMember - Use this method to restrict a user in a supergroup. The bot must be an administrator in 
// the supergroup for this to work and must have the appropriate admin rights. Pass True for all permissions to 
// lift restrictions from a user. Returns True on success.
// 
// Parameters:
// chatId - Unique identifier for the 
// target chat or username of the target supergroup (in the format @supergroupusername)
// 
// userId - Unique 
// identifier of the target user
// 
// permissions - A JSON-serialized object for new user 
// permissions
// 
// untilDate (optional) - Date when restrictions will be lifted for the user, unix time. If 
// user is restricted for more than 366 days or less than 30 seconds from the current time, they are considered 
// to be restricted forever
// 
//  
func restrictChatMember(chatId IntOrStringChatID, userId int, permissions ChatPermissions, untilDate int){
	
}

// promoteChatMember - Use this method to promote or demote a user in a supergroup or a channel. The bot must be 
// an administrator in the chat for this to work and must have the appropriate admin rights. Pass False for all 
// boolean parameters to demote a user. Returns True on success.
// 
// Parameters:
// chatId - Unique identifier 
// for the target chat or username of the target channel (in the format @channelusername)
// 
// userId - Unique 
// identifier of the target user
// 
// isAnonymous (optional) - Pass True, if the administrator's presence in the 
// chat is hidden
// 
// canManageChat (optional) - Pass True, if the administrator can access the chat event log, 
// chat statistics, message statistics in channels, see channel members, see anonymous administrators in 
// supergroups and ignore slow mode. Implied by any other administrator privilege
// 
// canPostMessages 
// (optional) - Pass True, if the administrator can create channel posts, channels only
// 
// canEditMessages 
// (optional) - Pass True, if the administrator can edit messages of other users and can pin messages, channels 
// only
// 
// canDeleteMessages (optional) - Pass True, if the administrator can delete messages of other 
// users
// 
// canManageVoiceChats (optional) - Pass True, if the administrator can manage voice 
// chats
// 
// canRestrictMembers (optional) - Pass True, if the administrator can restrict, ban or unban chat 
// members
// 
// canPromoteMembers (optional) - Pass True, if the administrator can add new administrators with a 
// subset of their own privileges or demote administrators that he has promoted, directly or indirectly 
// (promoted by administrators that were appointed by him)
// 
// canChangeInfo (optional) - Pass True, if the 
// administrator can change chat title, photo and other settings
// 
// canInviteUsers (optional) - Pass True, if 
// the administrator can invite new users to the chat
// 
// canPinMessages (optional) - Pass True, if the 
// administrator can pin messages, supergroups only
// 
//  
func promoteChatMember(chatId IntOrStringChatID, userId int, isAnonymous bool, canManageChat bool, canPostMessages bool, canEditMessages bool, canDeleteMessages bool, canManageVoiceChats bool, canRestrictMembers bool, canPromoteMembers bool, canChangeInfo bool, canInviteUsers bool, canPinMessages bool){
	
}

// setChatAdministratorCustomTitle - Use this method to set a custom title for an administrator in a supergroup 
// promoted by the bot. Returns True on success.
// 
// Parameters:
// chatId - Unique identifier for the target 
// chat or username of the target supergroup (in the format @supergroupusername)
// 
// userId - Unique identifier 
// of the target user
// 
// customTitle - New custom title for the administrator; 0-16 characters, emoji are not 
// allowed
// 
//  
func setChatAdministratorCustomTitle(chatId IntOrStringChatID, userId int, customTitle string){
	
}

// setChatPermissions - Use this method to set default chat permissions for all members. The bot must be an 
// administrator in the group or a supergroup for this to work and must have the can_restrict_members admin 
// rights. Returns True on success.
// 
// Parameters:
// chatId - Unique identifier for the target chat or 
// username of the target supergroup (in the format @supergroupusername)
// 
// permissions - New default chat 
// permissions
// 
//  
func setChatPermissions(chatId IntOrStringChatID, permissions ChatPermissions){
	
}

// exportChatInviteLink - Use this method to generate a new primary invite link for a chat; any previously 
// generated primary link is revoked. The bot must be an administrator in the chat for this to work and must 
// have the appropriate admin rights. Returns the new invite link as String on 
// success.
// 
// Parameters:
// chatId - Unique identifier for the target chat or username of the target channel 
// (in the format @channelusername)
// 
//  
func exportChatInviteLink(chatId IntOrStringChatID){
	
}

// createChatInviteLink - Use this method to create an additional invite link for a chat. The bot must be an 
// administrator in the chat for this to work and must have the appropriate admin rights. The link can be 
// revoked using the method revokeChatInviteLink (#revokechatinvitelink). Returns the new invite link as 
// ChatInviteLink (#chatinvitelink) object.
// 
// Parameters:
// chatId - Unique identifier for the target chat or 
// username of the target channel (in the format @channelusername)
// 
// expireDate (optional) - Point in time 
// (Unix timestamp) when the link will expire
// 
// memberLimit (optional) - Maximum number of users that can be 
// members of the chat simultaneously after joining the chat via this invite link; 1-99999
// 
//  
func createChatInviteLink(chatId IntOrStringChatID, expireDate int, memberLimit int){
	
}

// editChatInviteLink - Use this method to edit a non-primary invite link created by the bot. The bot must be an 
// administrator in the chat for this to work and must have the appropriate admin rights. Returns the edited 
// invite link as a ChatInviteLink (#chatinvitelink) object.
// 
// Parameters:
// chatId - Unique identifier for 
// the target chat or username of the target channel (in the format @channelusername)
// 
// inviteLink - The 
// invite link to edit
// 
// expireDate (optional) - Point in time (Unix timestamp) when the link will 
// expire
// 
// memberLimit (optional) - Maximum number of users that can be members of the chat simultaneously 
// after joining the chat via this invite link; 1-99999
// 
//  
func editChatInviteLink(chatId IntOrStringChatID, inviteLink string, expireDate int, memberLimit int){
	
}

// revokeChatInviteLink - Use this method to revoke an invite link created by the bot. If the primary link is 
// revoked, a new link is automatically generated. The bot must be an administrator in the chat for this to work 
// and must have the appropriate admin rights. Returns the revoked invite link as ChatInviteLink 
// (#chatinvitelink) object.
// 
// Parameters:
// chatId - Unique identifier of the target chat or username of the 
// target channel (in the format @channelusername)
// 
// inviteLink - The invite link to revoke
// 
//  
func revokeChatInviteLink(chatId IntOrStringChatID, inviteLink string){
	
}

// setChatPhoto - Use this method to set a new profile photo for the chat. Photos can't be changed for private 
// chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin 
// rights. Returns True on success.
// 
// Parameters:
// chatId - Unique identifier for the target chat or 
// username of the target channel (in the format @channelusername)
// 
// photo - New chat photo, uploaded using 
// multipart/form-data
// 
//  
func setChatPhoto(chatId IntOrStringChatID, photo InputFile){
	
}

// deleteChatPhoto - Use this method to delete a chat photo. Photos can't be changed for private chats. The bot 
// must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns 
// True on success.
// 
// Parameters:
// chatId - Unique identifier for the target chat or username of the target 
// channel (in the format @channelusername)
// 
//  
func deleteChatPhoto(chatId IntOrStringChatID){
	
}

// setChatTitle - Use this method to change the title of a chat. Titles can't be changed for private chats. The 
// bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns 
// True on success.
// 
// Parameters:
// chatId - Unique identifier for the target chat or username of the target 
// channel (in the format @channelusername)
// 
// title - New chat title, 1-255 characters
// 
//  
func setChatTitle(chatId IntOrStringChatID, title string){
	
}

// setChatDescription - Use this method to change the description of a group, a supergroup or a channel. The bot 
// must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns 
// True on success.
// 
// Parameters:
// chatId - Unique identifier for the target chat or username of the target 
// channel (in the format @channelusername)
// 
// description (optional) - New chat description, 0-255 
// characters
// 
//  
func setChatDescription(chatId IntOrStringChatID, description string){
	
}

// pinChatMessage - Use this method to add a message to the list of pinned messages in a chat. If the chat is 
// not a private chat, the bot must be an administrator in the chat for this to work and must have the 
// 'can_pin_messages' admin right in a supergroup or 'can_edit_messages' admin right in a channel. Returns True 
// on success.
// 
// Parameters:
// chatId - Unique identifier for the target chat or username of the target 
// channel (in the format @channelusername)
// 
// messageId - Identifier of a message to 
// pin
// 
// disableNotification (optional) - Pass True, if it is not necessary to send a notification to all 
// chat members about the new pinned message. Notifications are always disabled in channels and private 
// chats.
// 
//  
func pinChatMessage(chatId IntOrStringChatID, messageId int, disableNotification bool){
	
}

// unpinChatMessage - Use this method to remove a message from the list of pinned messages in a chat. If the 
// chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 
// 'can_pin_messages' admin right in a supergroup or 'can_edit_messages' admin right in a channel. Returns True 
// on success.
// 
// Parameters:
// chatId - Unique identifier for the target chat or username of the target 
// channel (in the format @channelusername)
// 
// messageId (optional) - Identifier of a message to unpin. If not 
// specified, the most recent pinned message (by sending date) will be unpinned.
// 
//  
func unpinChatMessage(chatId IntOrStringChatID, messageId int){
	
}

// unpinAllChatMessages - Use this method to clear the list of pinned messages in a chat. If the chat is not a 
// private chat, the bot must be an administrator in the chat for this to work and must have the 
// 'can_pin_messages' admin right in a supergroup or 'can_edit_messages' admin right in a channel. Returns True 
// on success.
// 
// Parameters:
// chatId - Unique identifier for the target chat or username of the target 
// channel (in the format @channelusername)
// 
//  
func unpinAllChatMessages(chatId IntOrStringChatID){
	
}

// leaveChat - Use this method for your bot to leave a group, supergroup or channel. Returns True on 
// success.
// 
// Parameters:
// chatId - Unique identifier for the target chat or username of the target 
// supergroup or channel (in the format @channelusername)
// 
//  
func leaveChat(chatId IntOrStringChatID){
	
}

// getChat - Use this method to get up to date information about the chat (current name of the user for 
// one-on-one conversations, current username of a user, group or channel, etc.). Returns a Chat (#chat) object 
// on success.
// 
// Parameters:
// chatId - Unique identifier for the target chat or username of the target 
// supergroup or channel (in the format @channelusername)
// 
//  
func getChat(chatId IntOrStringChatID){
	
}

// getChatAdministrators - Use this method to get a list of administrators in a chat. On success, returns an 
// Array of ChatMember (#chatmember) objects that contains information about all chat administrators except 
// other bots. If the chat is a group or a supergroup and no administrators were appointed, only the creator 
// will be returned.
// 
// Parameters:
// chatId - Unique identifier for the target chat or username of the target 
// supergroup or channel (in the format @channelusername)
// 
//  
func getChatAdministrators(chatId IntOrStringChatID){
	
}

// getChatMemberCount - Use this method to get the number of members in a chat. Returns Int on 
// success.
// 
// Parameters:
// chatId - Unique identifier for the target chat or username of the target 
// supergroup or channel (in the format @channelusername)
// 
//  
func getChatMemberCount(chatId IntOrStringChatID){
	
}

// getChatMember - Use this method to get information about a member of a chat. Returns a ChatMember 
// (#chatmember) object on success.
// 
// Parameters:
// chatId - Unique identifier for the target chat or 
// username of the target supergroup or channel (in the format @channelusername)
// 
// userId - Unique identifier 
// of the target user
// 
//  
func getChatMember(chatId IntOrStringChatID, userId int){
	
}

// setChatStickerSet - Use this method to set a new group sticker set for a supergroup. The bot must be an 
// administrator in the chat for this to work and must have the appropriate admin rights. Use the field 
// can_set_sticker_set optionally returned in getChat (#getchat) requests to check if the bot can use this 
// method. Returns True on success.
// 
// Parameters:
// chatId - Unique identifier for the target chat or 
// username of the target supergroup (in the format @supergroupusername)
// 
// stickerSetName - Name of the 
// sticker set to be set as the group sticker set
// 
//  
func setChatStickerSet(chatId IntOrStringChatID, stickerSetName string){
	
}

// deleteChatStickerSet - Use this method to delete a group sticker set from a supergroup. The bot must be an 
// administrator in the chat for this to work and must have the appropriate admin rights. Use the field 
// can_set_sticker_set optionally returned in getChat (#getchat) requests to check if the bot can use this 
// method. Returns True on success.
// 
// Parameters:
// chatId - Unique identifier for the target chat or 
// username of the target supergroup (in the format @supergroupusername)
// 
//  
func deleteChatStickerSet(chatId IntOrStringChatID){
	
}

// answerCallbackQuery - Use this method to send answers to callback queries sent from inline keyboards 
// (/bots#inline-keyboards-and-on-the-fly-updating). The answer will be displayed to the user as a notification 
// at the top of the chat screen or as an alert. On success, True is returned.
// 
// Parameters:
//  
func answerCallbackQuery(){
	
}

// setMyCommands - Use this method to change the list of the bot's commands. See 
// https://core.telegram.org/bots#commands (https://core.telegram.org/bots#commands) for more details about bot 
// commands. Returns True on success.
// 
// Parameters:
// commands - A JSON-serialized list of bot commands to be 
// set as the list of the bot's commands. At most 100 commands can be specified.
// 
// scope (optional) - A 
// JSON-serialized object, describing scope of users for which the commands are relevant. Defaults to 
// BotCommandScopeDefault (#botcommandscopedefault).
// 
// languageCode (optional) - A two-letter ISO 639-1 
// language code. If empty, commands will be applied to all users from the given scope, for whose language there 
// are no dedicated commands
// 
//  
func setMyCommands(commands []BotCommand, scope *BotCommandScope, languageCode string){
	
}

// deleteMyCommands - Use this method to delete the list of the bot's commands for the given scope and user 
// language. After deletion, higher level commands (#determining-list-of-commands) will be shown to affected 
// users. Returns True on success.
// 
// Parameters:
// scope (optional) - A JSON-serialized object, describing 
// scope of users for which the commands are relevant. Defaults to BotCommandScopeDefault 
// (#botcommandscopedefault).
// 
// languageCode (optional) - A two-letter ISO 639-1 language code. If empty, 
// commands will be applied to all users from the given scope, for whose language there are no dedicated 
// commands
// 
//  
func deleteMyCommands(scope *BotCommandScope, languageCode string){
	
}

// getMyCommands - Use this method to get the current list of the bot's commands for the given scope and user 
// language. Returns Array of BotCommand (#botcommand) on success. If commands aren't set, an empty list is 
// returned.
// 
// Parameters:
// scope (optional) - A JSON-serialized object, describing scope of users. Defaults 
// to BotCommandScopeDefault (#botcommandscopedefault).
// 
// languageCode (optional) - A two-letter ISO 639-1 
// language code or an empty string
// 
//  
func getMyCommands(scope *BotCommandScope, languageCode string){
	
}

// editMessageText - Use this method to edit text and game (#games) messages. On success, if the edited message 
// is not an inline message, the edited Message (#message) is returned, otherwise True is 
// returned.
// 
// Parameters:
// chatId (optional) - Required if inline_message_id is not specified. Unique 
// identifier for the target chat or username of the target channel (in the format 
// @channelusername)
// 
// messageId (optional) - Required if inline_message_id is not specified. Identifier of 
// the message to edit
// 
// inlineMessageId (optional) - Required if chat_id and message_id are not specified. 
// Identifier of the inline message
// 
// text - New text of the message, 1-4096 characters after entities 
// parsing
// 
// parseMode (optional) - Mode for parsing entities in the message text. See formatting options 
// (#formatting-options) for more details.
// 
// entities (optional) - List of special entities that appear in 
// message text, which can be specified instead of parse_mode
// 
// disableWebPagePreview (optional) - Disables 
// link previews for links in this message
// 
// replyMarkup (optional) - A JSON-serialized object for an inline 
// keyboard (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating).
// 
//  
func editMessageText(chatId *IntOrStringChatID, messageId int, inlineMessageId string, text string, parseMode string, entities []MessageEntity, disableWebPagePreview bool, replyMarkup *InlineKeyboardMarkup){
	
}

// editMessageCaption - Use this method to edit captions of messages. On success, if the edited message is not 
// an inline message, the edited Message (#message) is returned, otherwise True is 
// returned.
// 
// Parameters:
// chatId (optional) - Required if inline_message_id is not specified. Unique 
// identifier for the target chat or username of the target channel (in the format 
// @channelusername)
// 
// messageId (optional) - Required if inline_message_id is not specified. Identifier of 
// the message to edit
// 
// inlineMessageId (optional) - Required if chat_id and message_id are not specified. 
// Identifier of the inline message
// 
// caption (optional) - New caption of the message, 0-1024 characters 
// after entities parsing
// 
// parseMode (optional) - Mode for parsing entities in the message caption. See 
// formatting options (#formatting-options) for more details.
// 
// captionEntities (optional) - List of special 
// entities that appear in the caption, which can be specified instead of parse_mode
// 
// replyMarkup (optional) 
// - A JSON-serialized object for an inline keyboard 
// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating).
// 
//  
func editMessageCaption(chatId *IntOrStringChatID, messageId int, inlineMessageId string, caption string, parseMode string, captionEntities []MessageEntity, replyMarkup *InlineKeyboardMarkup){
	
}

// editMessageMedia - Use this method to edit animation, audio, document, photo, or video messages. If a message 
// is part of a message album, then it can be edited only to an audio for audio albums, only to a document for 
// document albums and to a photo or a video otherwise. When an inline message is edited, a new file can't be 
// uploaded. Use a previously uploaded file via its file_id or specify a URL. On success, if the edited message 
// was sent by the bot, the edited Message (#message) is returned, otherwise True is 
// returned.
// 
// Parameters:
// chatId (optional) - Required if inline_message_id is not specified. Unique 
// identifier for the target chat or username of the target channel (in the format 
// @channelusername)
// 
// messageId (optional) - Required if inline_message_id is not specified. Identifier of 
// the message to edit
// 
// inlineMessageId (optional) - Required if chat_id and message_id are not specified. 
// Identifier of the inline message
// 
// media - A JSON-serialized object for a new media content of the 
// message
// 
// replyMarkup (optional) - A JSON-serialized object for a new inline keyboard 
// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating).
// 
//  
func editMessageMedia(chatId *IntOrStringChatID, messageId int, inlineMessageId string, media InputMedia, replyMarkup *InlineKeyboardMarkup){
	
}

// editMessageReplyMarkup - Use this method to edit only the reply markup of messages. On success, if the edited 
// message is not an inline message, the edited Message (#message) is returned, otherwise True is 
// returned.
// 
// Parameters:
// chatId (optional) - Required if inline_message_id is not specified. Unique 
// identifier for the target chat or username of the target channel (in the format 
// @channelusername)
// 
// messageId (optional) - Required if inline_message_id is not specified. Identifier of 
// the message to edit
// 
// inlineMessageId (optional) - Required if chat_id and message_id are not specified. 
// Identifier of the inline message
// 
// replyMarkup (optional) - A JSON-serialized object for an inline 
// keyboard (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating).
// 
//  
func editMessageReplyMarkup(chatId *IntOrStringChatID, messageId int, inlineMessageId string, replyMarkup *InlineKeyboardMarkup){
	
}

// stopPoll - Use this method to stop a poll which was sent by the bot. On success, the stopped Poll (#poll) 
// with the final results is returned.
// 
// Parameters:
// chatId - Unique identifier for the target chat or 
// username of the target channel (in the format @channelusername)
// 
// messageId - Identifier of the original 
// message with the poll
// 
// replyMarkup (optional) - A JSON-serialized object for a new message inline 
// keyboard (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating).
// 
//  
func stopPoll(chatId IntOrStringChatID, messageId int, replyMarkup *InlineKeyboardMarkup){
	
}

// deleteMessage - Use this method to delete a message, including service messages, with the following 
// limitations:- A message can only be deleted if it was sent less than 48 hours ago.- A dice message in a 
// private chat can only be deleted if it was sent more than 24 hours ago.- Bots can delete outgoing messages in 
// private chats, groups, and supergroups.- Bots can delete incoming messages in private chats.- Bots granted 
// can_post_messages permissions can delete outgoing messages in channels.- If the bot is an administrator of a 
// group, it can delete any message there.- If the bot has can_delete_messages permission in a supergroup or a 
// channel, it can delete any message there.Returns True on success.
// 
// Parameters:
// chatId - Unique 
// identifier for the target chat or username of the target channel (in the format 
// @channelusername)
// 
// messageId - Identifier of the message to delete
// 
//  
func deleteMessage(chatId IntOrStringChatID, messageId int){
	
}

// sendSticker - Use this method to send static .WEBP or animated (https://telegram.org/blog/animated-stickers) 
// .TGS stickers. On success, the sent Message (#message) is returned.
// 
// Parameters:
// chatId - Unique 
// identifier for the target chat or username of the target channel (in the format 
// @channelusername)
// 
// sticker - Sticker to send. Pass a file_id as String to send a file that exists on the 
// Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a .WEBP file from the 
// Internet, or upload a new one using multipart/form-data. More info on Sending Files » 
// (#sending-files)
// 
// disableNotification (optional) - Sends the message silently 
// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no 
// sound.
// 
// replyToMessageId (optional) - If the message is a reply, ID of the original 
// message
// 
// allowSendingWithoutReply (optional) - Pass True, if the message should be sent even if the 
// specified replied-to message is not found
// 
// replyMarkup (optional) - Additional interface options. A 
// JSON-serialized object for an inline keyboard 
// (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating), custom reply keyboard 
// (https://core.telegram.org/bots#keyboards), instructions to remove reply keyboard or to force a reply from 
// the user.
// 
//  
func sendSticker(chatId IntOrStringChatID, sticker InputFile or String, disableNotification bool, replyToMessageId int, allowSendingWithoutReply bool, replyMarkup *InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply){
	
}

// getStickerSet - Use this method to get a sticker set. On success, a StickerSet (#stickerset) object is 
// returned.
// 
// Parameters:
// name - Name of the sticker set
// 
//  
func getStickerSet(name string){
	
}

// uploadStickerFile - Use this method to upload a .PNG file with a sticker for later use in createNewStickerSet 
// and addStickerToSet methods (can be used multiple times). Returns the uploaded File (#file) on 
// success.
// 
// Parameters:
// userId - User identifier of sticker file owner
// 
// pngSticker - PNG image with 
// the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or 
// height must be exactly 512px. More info on Sending Files » (#sending-files)
// 
//  
func uploadStickerFile(userId int, pngSticker InputFile){
	
}

// createNewStickerSet - Use this method to create a new sticker set owned by a user. The bot will be able to 
// edit the sticker set thus created. You must use exactly one of the fields png_sticker or tgs_sticker. Returns 
// True on success.
// 
// Parameters:
// userId - User identifier of created sticker set owner
// 
// name - Short 
// name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals). Can contain only english letters, 
// digits and underscores. Must begin with a letter, can't contain consecutive underscores and must end in 
// “_by_<bot username>”. <bot_username> is case insensitive. 1-64 characters.
// 
// title - Sticker set 
// title, 1-64 characters
// 
// pngSticker (optional) - PNG image with the sticker, must be up to 512 kilobytes 
// in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id 
// as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for 
// Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending 
// Files » (#sending-files)
// 
// tgsSticker (optional) - TGS animation with the sticker, uploaded using 
// multipart/form-data. See https://core.telegram.org/animated_stickers#technical-requirements 
// (https://core.telegram.org/animated_stickers#technical-requirements) for technical requirements
// 
// emojis - 
// One or more emoji corresponding to the sticker
// 
// containsMasks (optional) - Pass True, if a set of mask 
// stickers should be created
// 
// maskPosition (optional) - A JSON-serialized object for position where the 
// mask should be placed on faces
// 
//  
func createNewStickerSet(userId int, name string, title string, pngSticker *InputFile or String, tgsSticker *InputFile, emojis string, containsMasks bool, maskPosition *MaskPosition){
	
}

// addStickerToSet - Use this method to add a new sticker to a set created by the bot. You must use exactly one 
// of the fields png_sticker or tgs_sticker. Animated stickers can be added to animated sticker sets and only to 
// them. Animated sticker sets can have up to 50 stickers. Static sticker sets can have up to 120 stickers. 
// Returns True on success.
// 
// Parameters:
// userId - User identifier of sticker set owner
// 
// name - Sticker 
// set name
// 
// pngSticker (optional) - PNG image with the sticker, must be up to 512 kilobytes in size, 
// dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a 
// String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram 
// to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files » 
// (#sending-files)
// 
// tgsSticker (optional) - TGS animation with the sticker, uploaded using 
// multipart/form-data. See https://core.telegram.org/animated_stickers#technical-requirements 
// (https://core.telegram.org/animated_stickers#technical-requirements) for technical requirements
// 
// emojis - 
// One or more emoji corresponding to the sticker
// 
// maskPosition (optional) - A JSON-serialized object for 
// position where the mask should be placed on faces
// 
//  
func addStickerToSet(userId int, name string, pngSticker *InputFile or String, tgsSticker *InputFile, emojis string, maskPosition *MaskPosition){
	
}

// setStickerPositionInSet - Use this method to move a sticker in a set created by the bot to a specific 
// position. Returns True on success.
// 
// Parameters:
// sticker - File identifier of the sticker
// 
// position 
// - New sticker position in the set, zero-based
// 
//  
func setStickerPositionInSet(sticker string, position int){
	
}

// deleteStickerFromSet - Use this method to delete a sticker from a set created by the bot. Returns True on 
// success.
// 
// Parameters:
// sticker - File identifier of the sticker
// 
//  
func deleteStickerFromSet(sticker string){
	
}

// setStickerSetThumb - Use this method to set the thumbnail of a sticker set. Animated thumbnails can be set 
// for animated sticker sets only. Returns True on success.
// 
// Parameters:
// name - Sticker set 
// name
// 
// userId - User identifier of the sticker set owner
// 
// thumb (optional) - A PNG image with the 
// thumbnail, must be up to 128 kilobytes in size and have width and height exactly 100px, or a TGS animation 
// with the thumbnail up to 32 kilobytes in size; see 
// https://core.telegram.org/animated_stickers#technical-requirements 
// (https://core.telegram.org/animated_stickers#technical-requirements) for animated sticker technical 
// requirements. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an 
// HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using 
// multipart/form-data. More info on Sending Files » (#sending-files). Animated sticker set thumbnail can't be 
// uploaded via HTTP URL.
// 
//  
func setStickerSetThumb(name string, userId int, thumb *InputFile or String){
	
}

// answerInlineQuery - Use this method to send answers to an inline query. On success, True is returned.No more 
// than 50 results per query are allowed.
// 
// Parameters:
// inlineQueryId - Unique identifier for the answered 
// query
// 
// results - A JSON-serialized array of results for the inline query
// 
// cacheTime (optional) - The 
// maximum amount of time in seconds that the result of the inline query may be cached on the server. Defaults 
// to 300.
// 
// isPersonal (optional) - Pass True, if results may be cached on the server side only for the user 
// that sent the query. By default, results may be returned to any user who sends the same query
// 
// nextOffset 
// (optional) - Pass the offset that a client should send in the next query with the same text to receive more 
// results. Pass an empty string if there are no more results or if you don't support pagination. Offset length 
// can't exceed 64 bytes.
// 
// switchPmText (optional) - If passed, clients will display a button with specified 
// text that switches the user to a private chat with the bot and sends the bot a start message with the 
// parameter switch_pm_parameter
// 
// switchPmParameter (optional) - Deep-linking (/bots#deep-linking) parameter 
// for the /start message sent to the bot when user presses the switch button. 1-64 characters, only A-Z, a-z, 
// 0-9, _ and - are allowed.Example: An inline bot that sends YouTube videos can ask the user to connect the bot 
// to their YouTube account to adapt search results accordingly. To do this, it displays a 'Connect your YouTube 
// account' button above the results, or even before showing any. The user presses the button, switches to a 
// private chat with the bot and, in doing so, passes a start parameter that instructs the bot to return an 
// oauth link. Once done, the bot can offer a switch_inline (#inlinekeyboardmarkup) button so that the user can 
// easily return to the chat where they wanted to use the bot's inline capabilities.
// 
//  
func answerInlineQuery(inlineQueryId string, results []InlineQueryResult, cacheTime int, isPersonal bool, nextOffset string, switchPmText string, switchPmParameter string){
	
}

// sendInvoice - Use this method to send invoices. On success, the sent Message (#message) is 
// returned.
// 
// Parameters:
// chatId - Unique identifier for the target chat or username of the target channel 
// (in the format @channelusername)
// 
// title - Product name, 1-32 characters
// 
// description - Product 
// description, 1-255 characters
// 
// payload - Bot-defined invoice payload, 1-128 bytes. This will not be 
// displayed to the user, use for your internal processes.
// 
// providerToken - Payments provider token, 
// obtained via Botfather (https://t.me/botfather)
// 
// currency - Three-letter ISO 4217 currency code, see more 
// on currencies (/bots/payments#supported-currencies)
// 
// prices - Price breakdown, a JSON-serialized list of 
// components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.)
// 
// maxTipAmount 
// (optional) - The maximum accepted amount for tips in the smallest units of the currency (integer, not 
// float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in 
// currencies.json (https://core.telegram.org/bots/payments/currencies.json), it shows the number of digits past 
// the decimal point for each currency (2 for the majority of currencies). Defaults to 
// 0
// 
// suggestedTipAmounts (optional) - A JSON-serialized array of suggested amounts of tips in the smallest 
// units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The 
// suggested tip amounts must be positive, passed in a strictly increased order and must not exceed 
// max_tip_amount.
// 
// startParameter (optional) - Unique deep-linking parameter. If left empty, forwarded 
// copies of the sent message will have a Pay button, allowing multiple users to pay directly from the forwarded 
// message, using the same invoice. If non-empty, forwarded copies of the sent message will have a URL button 
// with a deep link to the bot (instead of a Pay button), with the value used as the start 
// parameter
// 
// providerData (optional) - A JSON-serialized data about the invoice, which will be shared with 
// the payment provider. A detailed description of required fields should be provided by the payment 
// provider.
// 
// photoUrl (optional) - URL of the product photo for the invoice. Can be a photo of the goods or 
// a marketing image for a service. People like it better when they see what they are paying for.
// 
// photoSize 
// (optional) - Photo size
// 
// photoWidth (optional) - Photo width
// 
// photoHeight (optional) - Photo 
// height
// 
// needName (optional) - Pass True, if you require the user's full name to complete the 
// order
// 
// needPhoneNumber (optional) - Pass True, if you require the user's phone number to complete the 
// order
// 
// needEmail (optional) - Pass True, if you require the user's email address to complete the 
// order
// 
// needShippingAddress (optional) - Pass True, if you require the user's shipping address to complete 
// the order
// 
// sendPhoneNumberToProvider (optional) - Pass True, if user's phone number should be sent to 
// provider
// 
// sendEmailToProvider (optional) - Pass True, if user's email address should be sent to 
// provider
// 
// isFlexible (optional) - Pass True, if the final price depends on the shipping 
// method
// 
// disableNotification (optional) - Sends the message silently 
// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no 
// sound.
// 
// replyToMessageId (optional) - If the message is a reply, ID of the original 
// message
// 
// allowSendingWithoutReply (optional) - Pass True, if the message should be sent even if the 
// specified replied-to message is not found
// 
// replyMarkup (optional) - A JSON-serialized object for an 
// inline keyboard (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating). If empty, one 'Pay 
// total price' button will be shown. If not empty, the first button must be a Pay button.
// 
//  
func sendInvoice(chatId IntOrStringChatID, title string, description string, payload string, providerToken string, currency string, prices []LabeledPrice, maxTipAmount int, suggestedTipAmounts []int, startParameter string, providerData string, photoUrl string, photoSize int, photoWidth int, photoHeight int, needName bool, needPhoneNumber bool, needEmail bool, needShippingAddress bool, sendPhoneNumberToProvider bool, sendEmailToProvider bool, isFlexible bool, disableNotification bool, replyToMessageId int, allowSendingWithoutReply bool, replyMarkup *InlineKeyboardMarkup){
	
}

// answerShippingQuery - If you sent an invoice requesting a shipping address and the parameter is_flexible was 
// specified, the Bot API will send an Update (#update) with a shipping_query field to the bot. Use this method 
// to reply to shipping queries. On success, True is returned.
// 
// Parameters:
// shippingQueryId - Unique 
// identifier for the query to be answered
// 
// ok - Specify True if delivery to the specified address is 
// possible and False if there are any problems (for example, if delivery to the specified address is not 
// possible)
// 
// shippingOptions (optional) - Required if ok is True. A JSON-serialized array of available 
// shipping options.
// 
// errorMessage (optional) - Required if ok is False. Error message in human readable 
// form that explains why it is impossible to complete the order (e.g. "Sorry, delivery to your desired address 
// is unavailable'). Telegram will display this message to the user.
// 
//  
func answerShippingQuery(shippingQueryId string, ok bool, shippingOptions []ShippingOption, errorMessage string){
	
}

// answerPreCheckoutQuery - Once the user has confirmed their payment and shipping details, the Bot API sends 
// the final confirmation in the form of an Update (#update) with the field pre_checkout_query. Use this method 
// to respond to such pre-checkout queries. On success, True is returned. Note: The Bot API must receive an 
// answer within 10 seconds after the pre-checkout query was sent.
// 
// Parameters:
// preCheckoutQueryId - 
// Unique identifier for the query to be answered
// 
// ok - Specify True if everything is alright (goods are 
// available, etc.) and the bot is ready to proceed with the order. Use False if there are any 
// problems.
// 
// errorMessage (optional) - Required if ok is False. Error message in human readable form that 
// explains the reason for failure to proceed with the checkout (e.g. "Sorry, somebody just bought the last of 
// our amazing black T-shirts while you were busy filling out your payment details. Please choose a different 
// color or garment!"). Telegram will display this message to the user.
// 
//  
func answerPreCheckoutQuery(preCheckoutQueryId string, ok bool, errorMessage string){
	
}

// setPassportDataErrors - Informs a user that some of the Telegram Passport elements they provided contains 
// errors. The user will not be able to re-submit their Passport to you until the errors are fixed (the contents 
// of the field for which you returned the error must change). Returns True on success.
// 
// Parameters:
//  
func setPassportDataErrors(){
	
}

// sendGame - Use this method to send a game. On success, the sent Message (#message) is 
// returned.
// 
// Parameters:
// chatId - Unique identifier for the target chat
// 
// gameShortName - Short name 
// of the game, serves as the unique identifier for the game. Set up your games via Botfather 
// (https://t.me/botfather).
// 
// disableNotification (optional) - Sends the message silently 
// (https://telegram.org/blog/channels-2-0#silent-messages). Users will receive a notification with no 
// sound.
// 
// replyToMessageId (optional) - If the message is a reply, ID of the original 
// message
// 
// allowSendingWithoutReply (optional) - Pass True, if the message should be sent even if the 
// specified replied-to message is not found
// 
// replyMarkup (optional) - A JSON-serialized object for an 
// inline keyboard (https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating). If empty, one 
// 'Play game_title' button will be shown. If not empty, the first button must launch the game.
// 
//  
func sendGame(chatId int, gameShortName string, disableNotification bool, replyToMessageId int, allowSendingWithoutReply bool, replyMarkup *InlineKeyboardMarkup){
	
}

// setGameScore - Use this method to set the score of the specified user in a game. On success, if the message 
// was sent by the bot, returns the edited Message (#message), otherwise returns True. Returns an error, if the 
// new score is not greater than the user's current score in the chat and force is 
// False.
// 
// Parameters:
// userId - User identifier
// 
// score - New score, must be non-negative
// 
// force 
// (optional) - Pass True, if the high score is allowed to decrease. This can be useful when fixing mistakes or 
// banning cheaters
// 
// disableEditMessage (optional) - Pass True, if the game message should not be 
// automatically edited to include the current scoreboard
// 
// chatId (optional) - Required if inline_message_id 
// is not specified. Unique identifier for the target chat
// 
// messageId (optional) - Required if 
// inline_message_id is not specified. Identifier of the sent message
// 
// inlineMessageId (optional) - Required 
// if chat_id and message_id are not specified. Identifier of the inline message
// 
//  
func setGameScore(userId int, score int, force bool, disableEditMessage bool, chatId int, messageId int, inlineMessageId string){
	
}

// getGameHighScores - Use this method to get data for high score tables. Will return the score of the specified 
// user and several of their neighbors in a game. On success, returns an Array of GameHighScore (#gamehighscore) 
// objects.
// 
// Parameters:
//  
func getGameHighScores(){
	
}

