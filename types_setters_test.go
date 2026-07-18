package telego

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplyParameters_Setters(t *testing.T) {
	r := (&ReplyParameters{}).
		WithMessageID(1).
		WithChatID(ChatID{ID: 2}).
		WithEphemeralMessageID(3).
		WithAllowSendingWithoutReply().
		WithQuote("Quote").
		WithQuoteParseMode("QuoteParseMode").
		WithQuoteEntities([]MessageEntity{{Type: "QuoteEntities"}}...).
		WithQuotePosition(4).
		WithChecklistTaskID(5).
		WithPollOptionID("PollOptionID")

	assert.Equal(t, &ReplyParameters{
		MessageID:                1,
		ChatID:                   ChatID{ID: 2},
		EphemeralMessageID:       3,
		AllowSendingWithoutReply: true,
		Quote:                    "Quote",
		QuoteParseMode:           "QuoteParseMode",
		QuoteEntities:            []MessageEntity{{Type: "QuoteEntities"}},
		QuotePosition:            4,
		ChecklistTaskID:          5,
		PollOptionID:             "PollOptionID",
	}, r)
}

func TestInputPollOption_Setters(t *testing.T) {
	i := (&InputPollOption{}).
		WithText("Text").
		WithTextParseMode("TextParseMode").
		WithTextEntities([]MessageEntity{{Type: "TextEntities"}}...).
		WithMedia(&InputMediaAnimation{Type: "Media"})

	assert.Equal(t, &InputPollOption{
		Text:          "Text",
		TextParseMode: "TextParseMode",
		TextEntities:  []MessageEntity{{Type: "TextEntities"}},
		Media:         &InputMediaAnimation{Type: "Media"},
	}, i)
}

func TestInputChecklistTask_Setters(t *testing.T) {
	i := (&InputChecklistTask{}).
		WithID(1).
		WithText("Text").
		WithParseMode("ParseMode").
		WithTextEntities([]MessageEntity{{Type: "TextEntities"}}...)

	assert.Equal(t, &InputChecklistTask{
		ID:           1,
		Text:         "Text",
		ParseMode:    "ParseMode",
		TextEntities: []MessageEntity{{Type: "TextEntities"}},
	}, i)
}

func TestInputChecklist_Setters(t *testing.T) {
	i := (&InputChecklist{}).
		WithTitle("Title").
		WithParseMode("ParseMode").
		WithTitleEntities([]MessageEntity{{Type: "TitleEntities"}}...).
		WithTasks([]InputChecklistTask{{}}...).
		WithOthersCanAddTasks().
		WithOthersCanMarkTasksAsDone()

	assert.Equal(t, &InputChecklist{
		Title:                    "Title",
		ParseMode:                "ParseMode",
		TitleEntities:            []MessageEntity{{Type: "TitleEntities"}},
		Tasks:                    []InputChecklistTask{{}},
		OthersCanAddTasks:        true,
		OthersCanMarkTasksAsDone: true,
	}, i)
}

func TestReplyKeyboardMarkup_Setters(t *testing.T) {
	r := (&ReplyKeyboardMarkup{}).
		WithKeyboard([][]KeyboardButton{{}}...).
		WithIsPersistent().
		WithResizeKeyboard().
		WithOneTimeKeyboard().
		WithInputFieldPlaceholder("InputFieldPlaceholder").
		WithSelective()

	assert.Equal(t, &ReplyKeyboardMarkup{
		Keyboard:              [][]KeyboardButton{{}},
		IsPersistent:          true,
		ResizeKeyboard:        true,
		OneTimeKeyboard:       true,
		InputFieldPlaceholder: "InputFieldPlaceholder",
		Selective:             true,
	}, r)
}

func TestKeyboardButton_Setters(t *testing.T) {
	k := (KeyboardButton{}).
		WithText("Text").
		WithIconCustomEmojiID("IconCustomEmojiID").
		WithStyle("Style").
		WithRequestUsers(&KeyboardButtonRequestUsers{RequestID: 1}).
		WithRequestChat(&KeyboardButtonRequestChat{RequestID: 2}).
		WithRequestManagedBot(&KeyboardButtonRequestManagedBot{RequestID: 3}).
		WithRequestContact().
		WithRequestLocation().
		WithRequestPoll(&KeyboardButtonPollType{Type: "RequestPoll"}).
		WithWebApp(&WebAppInfo{})

	assert.Equal(t, KeyboardButton{
		Text:              "Text",
		IconCustomEmojiID: "IconCustomEmojiID",
		Style:             "Style",
		RequestUsers:      &KeyboardButtonRequestUsers{RequestID: 1},
		RequestChat:       &KeyboardButtonRequestChat{RequestID: 2},
		RequestManagedBot: &KeyboardButtonRequestManagedBot{RequestID: 3},
		RequestContact:    true,
		RequestLocation:   true,
		RequestPoll:       &KeyboardButtonPollType{Type: "RequestPoll"},
		WebApp:            &WebAppInfo{},
	}, k)
}

func TestKeyboardButtonRequestUsers_Setters(t *testing.T) {
	k := (&KeyboardButtonRequestUsers{}).
		WithRequestID(4).
		WithUserIsBot(true).
		WithUserIsPremium(true).
		WithMaxQuantity(1).
		WithRequestName(true).
		WithRequestUsername(true).
		WithRequestPhoto(true)

	assert.Equal(t, &KeyboardButtonRequestUsers{
		RequestID:       4,
		UserIsBot:       ToPtr(true),
		UserIsPremium:   ToPtr(true),
		MaxQuantity:     1,
		RequestName:     ToPtr(true),
		RequestUsername: ToPtr(true),
		RequestPhoto:    ToPtr(true),
	}, k)
}

func TestKeyboardButtonRequestChat_Setters(t *testing.T) {
	k := (&KeyboardButtonRequestChat{}).
		WithRequestID(2).
		WithChatIsChannel().
		WithChatIsForum(true).
		WithChatHasUsername(true).
		WithChatIsCreated(true).
		WithUserAdministratorRights(&ChatAdministratorRights{IsAnonymous: true}).
		WithBotAdministratorRights(&ChatAdministratorRights{IsAnonymous: true}).
		WithBotIsMember(true).
		WithRequestTitle(true).
		WithRequestUsername(true).
		WithRequestPhoto(true)

	assert.Equal(t, &KeyboardButtonRequestChat{
		RequestID:               2,
		ChatIsChannel:           true,
		ChatIsForum:             ToPtr(true),
		ChatHasUsername:         ToPtr(true),
		ChatIsCreated:           ToPtr(true),
		UserAdministratorRights: &ChatAdministratorRights{IsAnonymous: true},
		BotAdministratorRights:  &ChatAdministratorRights{IsAnonymous: true},
		BotIsMember:             ToPtr(true),
		RequestTitle:            ToPtr(true),
		RequestUsername:         ToPtr(true),
		RequestPhoto:            ToPtr(true),
	}, k)
}

func TestKeyboardButtonRequestManagedBot_Setters(t *testing.T) {
	k := (&KeyboardButtonRequestManagedBot{}).
		WithRequestID(1).
		WithSuggestedName("SuggestedName").
		WithSuggestedUsername("SuggestedUsername")

	assert.Equal(t, &KeyboardButtonRequestManagedBot{
		RequestID:         1,
		SuggestedName:     "SuggestedName",
		SuggestedUsername: "SuggestedUsername",
	}, k)
}

func TestReplyKeyboardRemove_Setters(t *testing.T) {
	r := (&ReplyKeyboardRemove{}).
		WithRemoveKeyboard().
		WithSelective()

	assert.Equal(t, &ReplyKeyboardRemove{
		RemoveKeyboard: true,
		Selective:      true,
	}, r)
}

func TestInlineKeyboardMarkup_Setters(t *testing.T) {
	i := (&InlineKeyboardMarkup{}).
		WithInlineKeyboard([][]InlineKeyboardButton{{}}...)

	assert.Equal(t, &InlineKeyboardMarkup{
		InlineKeyboard: [][]InlineKeyboardButton{{}},
	}, i)
}

func TestInlineKeyboardButton_Setters(t *testing.T) {
	i := (InlineKeyboardButton{}).
		WithText("Text").
		WithIconCustomEmojiID("IconCustomEmojiID").
		WithStyle("Style").
		WithURL("URL").
		WithCallbackData("CallbackData").
		WithWebApp(&WebAppInfo{}).
		WithLoginURL(&LoginURL{URL: "LoginURL"}).
		WithSwitchInlineQuery("SwitchInlineQuery").
		WithSwitchInlineQueryCurrentChat("SwitchInlineQueryCurrentChat").
		WithSwitchInlineQueryChosenChat(&SwitchInlineQueryChosenChat{AllowUserChats: true}).
		WithCopyText(&CopyTextButton{}).
		WithCallbackGame(&CallbackGame{}).
		WithPay()

	assert.Equal(t, InlineKeyboardButton{
		Text:                         "Text",
		IconCustomEmojiID:            "IconCustomEmojiID",
		Style:                        "Style",
		URL:                          "URL",
		CallbackData:                 "CallbackData",
		WebApp:                       &WebAppInfo{},
		LoginURL:                     &LoginURL{URL: "LoginURL"},
		SwitchInlineQuery:            ToPtr("SwitchInlineQuery"),
		SwitchInlineQueryCurrentChat: ToPtr("SwitchInlineQueryCurrentChat"),
		SwitchInlineQueryChosenChat:  &SwitchInlineQueryChosenChat{AllowUserChats: true},
		CopyText:                     &CopyTextButton{},
		CallbackGame:                 &CallbackGame{},
		Pay:                          true,
	}, i)
}

func TestForceReply_Setters(t *testing.T) {
	f := (&ForceReply{}).
		WithForceReply().
		WithInputFieldPlaceholder("InputFieldPlaceholder").
		WithSelective()

	assert.Equal(t, &ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: "InputFieldPlaceholder",
		Selective:             true,
	}, f)
}

func TestMenuButtonWebApp_Setters(t *testing.T) {
	m := (&MenuButtonWebApp{}).
		WithText("Text").
		WithWebApp(WebAppInfo{})

	assert.Equal(t, &MenuButtonWebApp{
		Text:   "Text",
		WebApp: WebAppInfo{},
	}, m)
}

func TestInputMediaAnimation_Setters(t *testing.T) {
	i := (&InputMediaAnimation{}).
		WithMedia(testInputFile).
		WithThumbnail(&testInputFile).
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithShowCaptionAboveMedia().
		WithWidth(1).
		WithHeight(2).
		WithDuration(3).
		WithHasSpoiler()

	assert.Equal(t, &InputMediaAnimation{
		Media:                 testInputFile,
		Thumbnail:             &testInputFile,
		Caption:               "Caption",
		ParseMode:             "ParseMode",
		CaptionEntities:       []MessageEntity{{Type: "CaptionEntities"}},
		ShowCaptionAboveMedia: true,
		Width:                 1,
		Height:                2,
		Duration:              3,
		HasSpoiler:            true,
	}, i)
}

func TestInputMediaAudio_Setters(t *testing.T) {
	i := (&InputMediaAudio{}).
		WithMedia(testInputFile).
		WithThumbnail(&testInputFile).
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithDuration(1).
		WithPerformer("Performer").
		WithTitle("Title")

	assert.Equal(t, &InputMediaAudio{
		Media:           testInputFile,
		Thumbnail:       &testInputFile,
		Caption:         "Caption",
		ParseMode:       "ParseMode",
		CaptionEntities: []MessageEntity{{Type: "CaptionEntities"}},
		Duration:        1,
		Performer:       "Performer",
		Title:           "Title",
	}, i)
}

func TestInputMediaDocument_Setters(t *testing.T) {
	i := (&InputMediaDocument{}).
		WithMedia(testInputFile).
		WithThumbnail(&testInputFile).
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithDisableContentTypeDetection()

	assert.Equal(t, &InputMediaDocument{
		Media:                       testInputFile,
		Thumbnail:                   &testInputFile,
		Caption:                     "Caption",
		ParseMode:                   "ParseMode",
		CaptionEntities:             []MessageEntity{{Type: "CaptionEntities"}},
		DisableContentTypeDetection: true,
	}, i)
}

func TestInputMediaLivePhoto_Setters(t *testing.T) {
	i := (&InputMediaLivePhoto{}).
		WithMedia(testInputFile).
		WithPhoto(testInputFile).
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithShowCaptionAboveMedia().
		WithHasSpoiler()

	assert.Equal(t, &InputMediaLivePhoto{
		Media:                 testInputFile,
		Photo:                 testInputFile,
		Caption:               "Caption",
		ParseMode:             "ParseMode",
		CaptionEntities:       []MessageEntity{{Type: "CaptionEntities"}},
		ShowCaptionAboveMedia: true,
		HasSpoiler:            true,
	}, i)
}

func TestInputMediaLocation_Setters(t *testing.T) {
	i := (&InputMediaLocation{}).
		WithLatitude(1.0).
		WithLongitude(1.0).
		WithHorizontalAccuracy(2.0)

	assert.Equal(t, &InputMediaLocation{
		Latitude:           1.0,
		Longitude:          1.0,
		HorizontalAccuracy: 2.0,
	}, i)
}

func TestInputMediaPhoto_Setters(t *testing.T) {
	i := (&InputMediaPhoto{}).
		WithMedia(testInputFile).
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithShowCaptionAboveMedia().
		WithHasSpoiler()

	assert.Equal(t, &InputMediaPhoto{
		Media:                 testInputFile,
		Caption:               "Caption",
		ParseMode:             "ParseMode",
		CaptionEntities:       []MessageEntity{{Type: "CaptionEntities"}},
		ShowCaptionAboveMedia: true,
		HasSpoiler:            true,
	}, i)
}

func TestInputMediaSticker_Setters(t *testing.T) {
	i := (&InputMediaSticker{}).
		WithMedia(testInputFile).
		WithEmoji("Emoji")

	assert.Equal(t, &InputMediaSticker{
		Media: testInputFile,
		Emoji: "Emoji",
	}, i)
}

func TestInputMediaVenue_Setters(t *testing.T) {
	i := (&InputMediaVenue{}).
		WithLatitude(1.0).
		WithLongitude(1.0).
		WithTitle("Title").
		WithAddress("Address").
		WithFoursquareID("FoursquareID").
		WithFoursquareType("FoursquareType").
		WithGooglePlaceID("GooglePlaceID").
		WithGooglePlaceType("GooglePlaceType")

	assert.Equal(t, &InputMediaVenue{
		Latitude:        1.0,
		Longitude:       1.0,
		Title:           "Title",
		Address:         "Address",
		FoursquareID:    "FoursquareID",
		FoursquareType:  "FoursquareType",
		GooglePlaceID:   "GooglePlaceID",
		GooglePlaceType: "GooglePlaceType",
	}, i)
}

func TestInputMediaVideo_Setters(t *testing.T) {
	i := (&InputMediaVideo{}).
		WithMedia(testInputFile).
		WithThumbnail(&testInputFile).
		WithCover(&testInputFile).
		WithStartTimestamp(1).
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithShowCaptionAboveMedia().
		WithWidth(2).
		WithHeight(3).
		WithDuration(4).
		WithSupportsStreaming().
		WithHasSpoiler()

	assert.Equal(t, &InputMediaVideo{
		Media:                 testInputFile,
		Thumbnail:             &testInputFile,
		Cover:                 &testInputFile,
		StartTimestamp:        1,
		Caption:               "Caption",
		ParseMode:             "ParseMode",
		CaptionEntities:       []MessageEntity{{Type: "CaptionEntities"}},
		ShowCaptionAboveMedia: true,
		Width:                 2,
		Height:                3,
		Duration:              4,
		SupportsStreaming:     true,
		HasSpoiler:            true,
	}, i)
}

func TestInputPaidMediaLivePhoto_Setters(t *testing.T) {
	i := (&InputPaidMediaLivePhoto{}).
		WithMedia(testInputFile).
		WithPhoto(testInputFile)

	assert.Equal(t, &InputPaidMediaLivePhoto{
		Media: testInputFile,
		Photo: testInputFile,
	}, i)
}

func TestInputPaidMediaPhoto_Setters(t *testing.T) {
	i := (&InputPaidMediaPhoto{}).
		WithMedia(testInputFile)

	assert.Equal(t, &InputPaidMediaPhoto{
		Media: testInputFile,
	}, i)
}

func TestInputPaidMediaVideo_Setters(t *testing.T) {
	i := (&InputPaidMediaVideo{}).
		WithMedia(testInputFile).
		WithThumbnail(&testInputFile).
		WithCover(&testInputFile).
		WithStartTimestamp(1).
		WithWidth(2).
		WithHeight(3).
		WithDuration(4).
		WithSupportsStreaming()

	assert.Equal(t, &InputPaidMediaVideo{
		Media:             testInputFile,
		Thumbnail:         &testInputFile,
		Cover:             &testInputFile,
		StartTimestamp:    1,
		Width:             2,
		Height:            3,
		Duration:          4,
		SupportsStreaming: true,
	}, i)
}

func TestInputProfilePhotoStatic_Setters(t *testing.T) {
	i := (&InputProfilePhotoStatic{}).
		WithPhoto(testInputFile)

	assert.Equal(t, &InputProfilePhotoStatic{
		Photo: testInputFile,
	}, i)
}

func TestInputProfilePhotoAnimated_Setters(t *testing.T) {
	i := (&InputProfilePhotoAnimated{}).
		WithAnimation(testInputFile).
		WithMainFrameTimestamp(1.0)

	assert.Equal(t, &InputProfilePhotoAnimated{
		Animation:          testInputFile,
		MainFrameTimestamp: 1.0,
	}, i)
}

func TestInputStoryContentPhoto_Setters(t *testing.T) {
	i := (&InputStoryContentPhoto{}).
		WithPhoto(testInputFile)

	assert.Equal(t, &InputStoryContentPhoto{
		Photo: testInputFile,
	}, i)
}

func TestInputStoryContentVideo_Setters(t *testing.T) {
	i := (&InputStoryContentVideo{}).
		WithVideo(testInputFile).
		WithDuration(1.0).
		WithCoverFrameTimestamp(2.0).
		WithIsAnimation()

	assert.Equal(t, &InputStoryContentVideo{
		Video:               testInputFile,
		Duration:            1.0,
		CoverFrameTimestamp: 2.0,
		IsAnimation:         true,
	}, i)
}

func TestInputSticker_Setters(t *testing.T) {
	i := (&InputSticker{}).
		WithSticker(testInputFile).
		WithFormat("Format").
		WithEmojiList([]string{"EmojiList"}...).
		WithMaskPosition(&MaskPosition{Point: "MaskPosition"}).
		WithKeywords([]string{"Keywords"}...)

	assert.Equal(t, &InputSticker{
		Sticker:      testInputFile,
		Format:       "Format",
		EmojiList:    []string{"EmojiList"},
		MaskPosition: &MaskPosition{Point: "MaskPosition"},
		Keywords:     []string{"Keywords"},
	}, i)
}

func TestInputRichMessage_Setters(t *testing.T) {
	i := (&InputRichMessage{}).
		WithBlocks([]InputRichBlock{&InputRichBlockParagraph{}}...).
		WithHTML("HTML").
		WithMarkdown("Markdown").
		WithMedia([]InputRichMessageMedia{{}}...).
		WithIsRtl().
		WithSkipEntityDetection()

	assert.Equal(t, &InputRichMessage{
		Blocks:              []InputRichBlock{&InputRichBlockParagraph{}},
		HTML:                "HTML",
		Markdown:            "Markdown",
		Media:               []InputRichMessageMedia{{}},
		IsRtl:               true,
		SkipEntityDetection: true,
	}, i)
}

func TestRichTextBold_Setters(t *testing.T) {
	r := (&RichTextBold{}).
		WithText(&RichTextBold{})

	assert.Equal(t, &RichTextBold{
		Text: &RichTextBold{},
	}, r)
}

func TestRichTextItalic_Setters(t *testing.T) {
	r := (&RichTextItalic{}).
		WithText(&RichTextBold{})

	assert.Equal(t, &RichTextItalic{
		Text: &RichTextBold{},
	}, r)
}

func TestRichTextUnderline_Setters(t *testing.T) {
	r := (&RichTextUnderline{}).
		WithText(&RichTextBold{})

	assert.Equal(t, &RichTextUnderline{
		Text: &RichTextBold{},
	}, r)
}

func TestRichTextStrikethrough_Setters(t *testing.T) {
	r := (&RichTextStrikethrough{}).
		WithText(&RichTextBold{})

	assert.Equal(t, &RichTextStrikethrough{
		Text: &RichTextBold{},
	}, r)
}

func TestRichTextSpoiler_Setters(t *testing.T) {
	r := (&RichTextSpoiler{}).
		WithText(&RichTextBold{})

	assert.Equal(t, &RichTextSpoiler{
		Text: &RichTextBold{},
	}, r)
}

func TestRichTextDateTime_Setters(t *testing.T) {
	r := (&RichTextDateTime{}).
		WithText(&RichTextBold{}).
		WithUnixTime(1).
		WithDateTimeFormat("DateTimeFormat")

	assert.Equal(t, &RichTextDateTime{
		Text:           &RichTextBold{},
		UnixTime:       1,
		DateTimeFormat: "DateTimeFormat",
	}, r)
}

func TestRichTextTextMention_Setters(t *testing.T) {
	r := (&RichTextTextMention{}).
		WithText(&RichTextBold{}).
		WithUser(User{ID: 1})

	assert.Equal(t, &RichTextTextMention{
		Text: &RichTextBold{},
		User: User{ID: 1},
	}, r)
}

func TestRichTextSubscript_Setters(t *testing.T) {
	r := (&RichTextSubscript{}).
		WithText(&RichTextBold{})

	assert.Equal(t, &RichTextSubscript{
		Text: &RichTextBold{},
	}, r)
}

func TestRichTextSuperscript_Setters(t *testing.T) {
	r := (&RichTextSuperscript{}).
		WithText(&RichTextBold{})

	assert.Equal(t, &RichTextSuperscript{
		Text: &RichTextBold{},
	}, r)
}

func TestRichTextMarked_Setters(t *testing.T) {
	r := (&RichTextMarked{}).
		WithText(&RichTextBold{})

	assert.Equal(t, &RichTextMarked{
		Text: &RichTextBold{},
	}, r)
}

func TestRichTextCode_Setters(t *testing.T) {
	r := (&RichTextCode{}).
		WithText(&RichTextBold{})

	assert.Equal(t, &RichTextCode{
		Text: &RichTextBold{},
	}, r)
}

func TestRichTextCustomEmoji_Setters(t *testing.T) {
	r := (&RichTextCustomEmoji{}).
		WithCustomEmojiID("CustomEmojiID").
		WithAlternativeText("AlternativeText")

	assert.Equal(t, &RichTextCustomEmoji{
		CustomEmojiID:   "CustomEmojiID",
		AlternativeText: "AlternativeText",
	}, r)
}

func TestRichTextMathematicalExpression_Setters(t *testing.T) {
	r := (&RichTextMathematicalExpression{}).
		WithExpression("Expression")

	assert.Equal(t, &RichTextMathematicalExpression{
		Expression: "Expression",
	}, r)
}

func TestRichTextURL_Setters(t *testing.T) {
	r := (&RichTextURL{}).
		WithText(&RichTextBold{}).
		WithURL("URL")

	assert.Equal(t, &RichTextURL{
		Text: &RichTextBold{},
		URL:  "URL",
	}, r)
}

func TestRichTextEmailAddress_Setters(t *testing.T) {
	r := (&RichTextEmailAddress{}).
		WithText(&RichTextBold{}).
		WithEmailAddress("EmailAddress")

	assert.Equal(t, &RichTextEmailAddress{
		Text:         &RichTextBold{},
		EmailAddress: "EmailAddress",
	}, r)
}

func TestRichTextPhoneNumber_Setters(t *testing.T) {
	r := (&RichTextPhoneNumber{}).
		WithText(&RichTextBold{}).
		WithPhoneNumber("PhoneNumber")

	assert.Equal(t, &RichTextPhoneNumber{
		Text:        &RichTextBold{},
		PhoneNumber: "PhoneNumber",
	}, r)
}

func TestRichTextBankCardNumber_Setters(t *testing.T) {
	r := (&RichTextBankCardNumber{}).
		WithText(&RichTextBold{}).
		WithBankCardNumber("BankCardNumber")

	assert.Equal(t, &RichTextBankCardNumber{
		Text:           &RichTextBold{},
		BankCardNumber: "BankCardNumber",
	}, r)
}

func TestRichTextMention_Setters(t *testing.T) {
	r := (&RichTextMention{}).
		WithText(&RichTextBold{}).
		WithUsername("Username")

	assert.Equal(t, &RichTextMention{
		Text:     &RichTextBold{},
		Username: "Username",
	}, r)
}

func TestRichTextHashtag_Setters(t *testing.T) {
	r := (&RichTextHashtag{}).
		WithText(&RichTextBold{}).
		WithHashtag("Hashtag")

	assert.Equal(t, &RichTextHashtag{
		Text:    &RichTextBold{},
		Hashtag: "Hashtag",
	}, r)
}

func TestRichTextCashtag_Setters(t *testing.T) {
	r := (&RichTextCashtag{}).
		WithText(&RichTextBold{}).
		WithCashtag("Cashtag")

	assert.Equal(t, &RichTextCashtag{
		Text:    &RichTextBold{},
		Cashtag: "Cashtag",
	}, r)
}

func TestRichTextBotCommand_Setters(t *testing.T) {
	r := (&RichTextBotCommand{}).
		WithText(&RichTextBold{}).
		WithBotCommand("BotCommand")

	assert.Equal(t, &RichTextBotCommand{
		Text:       &RichTextBold{},
		BotCommand: "BotCommand",
	}, r)
}

func TestRichTextAnchor_Setters(t *testing.T) {
	r := (&RichTextAnchor{}).
		WithName("Name")

	assert.Equal(t, &RichTextAnchor{
		Name: "Name",
	}, r)
}

func TestRichTextAnchorLink_Setters(t *testing.T) {
	r := (&RichTextAnchorLink{}).
		WithText(&RichTextBold{}).
		WithAnchorName("AnchorName")

	assert.Equal(t, &RichTextAnchorLink{
		Text:       &RichTextBold{},
		AnchorName: "AnchorName",
	}, r)
}

func TestRichTextReference_Setters(t *testing.T) {
	r := (&RichTextReference{}).
		WithText(&RichTextBold{}).
		WithName("Name")

	assert.Equal(t, &RichTextReference{
		Text: &RichTextBold{},
		Name: "Name",
	}, r)
}

func TestRichTextReferenceLink_Setters(t *testing.T) {
	r := (&RichTextReferenceLink{}).
		WithText(&RichTextBold{}).
		WithReferenceName("ReferenceName")

	assert.Equal(t, &RichTextReferenceLink{
		Text:          &RichTextBold{},
		ReferenceName: "ReferenceName",
	}, r)
}

func TestRichBlockCaption_Setters(t *testing.T) {
	r := (&RichBlockCaption{}).
		WithText(&RichTextBold{}).
		WithCredit(&RichTextBold{})

	assert.Equal(t, &RichBlockCaption{
		Text:   &RichTextBold{},
		Credit: &RichTextBold{},
	}, r)
}

func TestRichBlockTableCell_Setters(t *testing.T) {
	r := (RichBlockTableCell{}).
		WithText(&RichTextBold{}).
		WithIsHeader().
		WithColspan(1).
		WithRowspan(2).
		WithAlign("Align").
		WithValign("Valign")

	assert.Equal(t, RichBlockTableCell{
		Text:     &RichTextBold{},
		IsHeader: true,
		Colspan:  1,
		Rowspan:  2,
		Align:    "Align",
		Valign:   "Valign",
	}, r)
}

func TestInputRichBlockListItem_Setters(t *testing.T) {
	i := (InputRichBlockListItem{}).
		WithBlocks([]InputRichBlock{&InputRichBlockParagraph{}}...).
		WithHasCheckbox().
		WithIsChecked().
		WithValue(1).
		WithType("Type")

	assert.Equal(t, InputRichBlockListItem{
		Blocks:      []InputRichBlock{&InputRichBlockParagraph{}},
		HasCheckbox: true,
		IsChecked:   true,
		Value:       1,
		Type:        "Type",
	}, i)
}

func TestInputRichBlockParagraph_Setters(t *testing.T) {
	i := (&InputRichBlockParagraph{}).
		WithText(&RichTextBold{})

	assert.Equal(t, &InputRichBlockParagraph{
		Text: &RichTextBold{},
	}, i)
}

func TestInputRichBlockSectionHeading_Setters(t *testing.T) {
	i := (&InputRichBlockSectionHeading{}).
		WithText(&RichTextBold{}).
		WithSize(1)

	assert.Equal(t, &InputRichBlockSectionHeading{
		Text: &RichTextBold{},
		Size: 1,
	}, i)
}

func TestInputRichBlockPreformatted_Setters(t *testing.T) {
	i := (&InputRichBlockPreformatted{}).
		WithText(&RichTextBold{}).
		WithLanguage("Language")

	assert.Equal(t, &InputRichBlockPreformatted{
		Text:     &RichTextBold{},
		Language: "Language",
	}, i)
}

func TestInputRichBlockFooter_Setters(t *testing.T) {
	i := (&InputRichBlockFooter{}).
		WithText(&RichTextBold{})

	assert.Equal(t, &InputRichBlockFooter{
		Text: &RichTextBold{},
	}, i)
}

func TestInputRichBlockMathematicalExpression_Setters(t *testing.T) {
	i := (&InputRichBlockMathematicalExpression{}).
		WithExpression("Expression")

	assert.Equal(t, &InputRichBlockMathematicalExpression{
		Expression: "Expression",
	}, i)
}

func TestInputRichBlockAnchor_Setters(t *testing.T) {
	i := (&InputRichBlockAnchor{}).
		WithName("Name")

	assert.Equal(t, &InputRichBlockAnchor{
		Name: "Name",
	}, i)
}

func TestInputRichBlockList_Setters(t *testing.T) {
	i := (&InputRichBlockList{}).
		WithItems([]InputRichBlockListItem{{}}...)

	assert.Equal(t, &InputRichBlockList{
		Items: []InputRichBlockListItem{{}},
	}, i)
}

func TestInputRichBlockBlockQuotation_Setters(t *testing.T) {
	i := (&InputRichBlockBlockQuotation{}).
		WithBlocks([]InputRichBlock{&InputRichBlockParagraph{}}...).
		WithCredit(&RichTextBold{})

	assert.Equal(t, &InputRichBlockBlockQuotation{
		Blocks: []InputRichBlock{&InputRichBlockParagraph{}},
		Credit: &RichTextBold{},
	}, i)
}

func TestInputRichBlockPullQuotation_Setters(t *testing.T) {
	i := (&InputRichBlockPullQuotation{}).
		WithText(&RichTextBold{}).
		WithCredit(&RichTextBold{})

	assert.Equal(t, &InputRichBlockPullQuotation{
		Text:   &RichTextBold{},
		Credit: &RichTextBold{},
	}, i)
}

func TestInputRichBlockCollage_Setters(t *testing.T) {
	i := (&InputRichBlockCollage{}).
		WithBlocks([]InputRichBlock{&InputRichBlockParagraph{}}...).
		WithCaption(&RichBlockCaption{})

	assert.Equal(t, &InputRichBlockCollage{
		Blocks:  []InputRichBlock{&InputRichBlockParagraph{}},
		Caption: &RichBlockCaption{},
	}, i)
}

func TestInputRichBlockSlideshow_Setters(t *testing.T) {
	i := (&InputRichBlockSlideshow{}).
		WithBlocks([]InputRichBlock{&InputRichBlockParagraph{}}...).
		WithCaption(&RichBlockCaption{})

	assert.Equal(t, &InputRichBlockSlideshow{
		Blocks:  []InputRichBlock{&InputRichBlockParagraph{}},
		Caption: &RichBlockCaption{},
	}, i)
}

func TestInputRichBlockTable_Setters(t *testing.T) {
	i := (&InputRichBlockTable{}).
		WithCells([][]RichBlockTableCell{{}}...).
		WithIsBordered().
		WithIsStriped().
		WithCaption(&RichTextBold{})

	assert.Equal(t, &InputRichBlockTable{
		Cells:      [][]RichBlockTableCell{{}},
		IsBordered: true,
		IsStriped:  true,
		Caption:    &RichTextBold{},
	}, i)
}

func TestInputRichBlockDetails_Setters(t *testing.T) {
	i := (&InputRichBlockDetails{}).
		WithSummary(&RichTextBold{}).
		WithBlocks([]InputRichBlock{&InputRichBlockParagraph{}}...).
		WithIsOpen()

	assert.Equal(t, &InputRichBlockDetails{
		Summary: &RichTextBold{},
		Blocks:  []InputRichBlock{&InputRichBlockParagraph{}},
		IsOpen:  true,
	}, i)
}

func TestInputRichBlockMap_Setters(t *testing.T) {
	i := (&InputRichBlockMap{}).
		WithLocation(Location{Latitude: 1, Longitude: 1}).
		WithZoom(1).
		WithWidth(2).
		WithHeight(3).
		WithCaption(&RichBlockCaption{})

	assert.Equal(t, &InputRichBlockMap{
		Location: Location{Latitude: 1, Longitude: 1},
		Zoom:     1,
		Width:    2,
		Height:   3,
		Caption:  &RichBlockCaption{},
	}, i)
}

func TestInputRichBlockAnimation_Setters(t *testing.T) {
	i := (&InputRichBlockAnimation{}).
		WithAnimation(InputMediaAnimation{}).
		WithCaption(&RichBlockCaption{})

	assert.Equal(t, &InputRichBlockAnimation{
		Animation: InputMediaAnimation{},
		Caption:   &RichBlockCaption{},
	}, i)
}

func TestInputRichBlockAudio_Setters(t *testing.T) {
	i := (&InputRichBlockAudio{}).
		WithAudio(InputMediaAudio{}).
		WithCaption(&RichBlockCaption{})

	assert.Equal(t, &InputRichBlockAudio{
		Audio:   InputMediaAudio{},
		Caption: &RichBlockCaption{},
	}, i)
}

func TestInputRichBlockPhoto_Setters(t *testing.T) {
	i := (&InputRichBlockPhoto{}).
		WithPhoto(InputMediaPhoto{}).
		WithCaption(&RichBlockCaption{})

	assert.Equal(t, &InputRichBlockPhoto{
		Photo:   InputMediaPhoto{},
		Caption: &RichBlockCaption{},
	}, i)
}

func TestInputRichBlockVideo_Setters(t *testing.T) {
	i := (&InputRichBlockVideo{}).
		WithVideo(InputMediaVideo{}).
		WithCaption(&RichBlockCaption{})

	assert.Equal(t, &InputRichBlockVideo{
		Video:   InputMediaVideo{},
		Caption: &RichBlockCaption{},
	}, i)
}

func TestInputRichBlockVoiceNote_Setters(t *testing.T) {
	i := (&InputRichBlockVoiceNote{}).
		WithVoiceNote(InputMediaVoiceNote{}).
		WithCaption(&RichBlockCaption{})

	assert.Equal(t, &InputRichBlockVoiceNote{
		VoiceNote: InputMediaVoiceNote{},
		Caption:   &RichBlockCaption{},
	}, i)
}

func TestInputRichBlockThinking_Setters(t *testing.T) {
	i := (&InputRichBlockThinking{}).
		WithText(&RichTextBold{})

	assert.Equal(t, &InputRichBlockThinking{
		Text: &RichTextBold{},
	}, i)
}

func TestInlineQueryResultArticle_Setters(t *testing.T) {
	i := (&InlineQueryResultArticle{}).
		WithID("ID").
		WithTitle("Title").
		WithInputMessageContent(&InputTextMessageContent{}).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithURL("URL").
		WithDescription("Description").
		WithThumbnailURL("ThumbnailURL").
		WithThumbnailWidth(1).
		WithThumbnailHeight(2)

	assert.Equal(t, &InlineQueryResultArticle{
		ID:                  "ID",
		Title:               "Title",
		InputMessageContent: &InputTextMessageContent{},
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		URL:                 "URL",
		Description:         "Description",
		ThumbnailURL:        "ThumbnailURL",
		ThumbnailWidth:      1,
		ThumbnailHeight:     2,
	}, i)
}

func TestInlineQueryResultPhoto_Setters(t *testing.T) {
	i := (&InlineQueryResultPhoto{}).
		WithID("ID").
		WithPhotoURL("PhotoURL").
		WithThumbnailURL("ThumbnailURL").
		WithPhotoWidth(1).
		WithPhotoHeight(2).
		WithTitle("Title").
		WithDescription("Description").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithShowCaptionAboveMedia().
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultPhoto{
		ID:                    "ID",
		PhotoURL:              "PhotoURL",
		ThumbnailURL:          "ThumbnailURL",
		PhotoWidth:            1,
		PhotoHeight:           2,
		Title:                 "Title",
		Description:           "Description",
		Caption:               "Caption",
		ParseMode:             "ParseMode",
		CaptionEntities:       []MessageEntity{{Type: "CaptionEntities"}},
		ShowCaptionAboveMedia: true,
		ReplyMarkup:           &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent:   &InputTextMessageContent{},
	}, i)
}

func TestInlineQueryResultGif_Setters(t *testing.T) {
	i := (&InlineQueryResultGif{}).
		WithID("ID").
		WithGifURL("GifURL").
		WithGifWidth(1).
		WithGifHeight(2).
		WithGifDuration(3).
		WithThumbnailURL("ThumbnailURL").
		WithThumbnailMimeType("ThumbnailMimeType").
		WithTitle("Title").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithShowCaptionAboveMedia().
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultGif{
		ID:                    "ID",
		GifURL:                "GifURL",
		GifWidth:              1,
		GifHeight:             2,
		GifDuration:           3,
		ThumbnailURL:          "ThumbnailURL",
		ThumbnailMimeType:     "ThumbnailMimeType",
		Title:                 "Title",
		Caption:               "Caption",
		ParseMode:             "ParseMode",
		CaptionEntities:       []MessageEntity{{Type: "CaptionEntities"}},
		ShowCaptionAboveMedia: true,
		ReplyMarkup:           &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent:   &InputTextMessageContent{},
	}, i)
}

func TestInlineQueryResultMpeg4Gif_Setters(t *testing.T) {
	i := (&InlineQueryResultMpeg4Gif{}).
		WithID("ID").
		WithMpeg4URL("Mpeg4URL").
		WithMpeg4Width(1).
		WithMpeg4Height(2).
		WithMpeg4Duration(3).
		WithThumbnailURL("ThumbnailURL").
		WithThumbnailMimeType("ThumbnailMimeType").
		WithTitle("Title").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithShowCaptionAboveMedia().
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultMpeg4Gif{
		ID:                    "ID",
		Mpeg4URL:              "Mpeg4URL",
		Mpeg4Width:            1,
		Mpeg4Height:           2,
		Mpeg4Duration:         3,
		ThumbnailURL:          "ThumbnailURL",
		ThumbnailMimeType:     "ThumbnailMimeType",
		Title:                 "Title",
		Caption:               "Caption",
		ParseMode:             "ParseMode",
		CaptionEntities:       []MessageEntity{{Type: "CaptionEntities"}},
		ShowCaptionAboveMedia: true,
		ReplyMarkup:           &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent:   &InputTextMessageContent{},
	}, i)
}

func TestInlineQueryResultVideo_Setters(t *testing.T) {
	i := (&InlineQueryResultVideo{}).
		WithID("ID").
		WithVideoURL("VideoURL").
		WithMimeType("MimeType").
		WithThumbnailURL("ThumbnailURL").
		WithTitle("Title").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithShowCaptionAboveMedia().
		WithVideoWidth(1).
		WithVideoHeight(2).
		WithVideoDuration(3).
		WithDescription("Description").
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultVideo{
		ID:                    "ID",
		VideoURL:              "VideoURL",
		MimeType:              "MimeType",
		ThumbnailURL:          "ThumbnailURL",
		Title:                 "Title",
		Caption:               "Caption",
		ParseMode:             "ParseMode",
		CaptionEntities:       []MessageEntity{{Type: "CaptionEntities"}},
		ShowCaptionAboveMedia: true,
		VideoWidth:            1,
		VideoHeight:           2,
		VideoDuration:         3,
		Description:           "Description",
		ReplyMarkup:           &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent:   &InputTextMessageContent{},
	}, i)
}

func TestInlineQueryResultAudio_Setters(t *testing.T) {
	i := (&InlineQueryResultAudio{}).
		WithID("ID").
		WithAudioURL("AudioURL").
		WithTitle("Title").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithPerformer("Performer").
		WithAudioDuration(1).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultAudio{
		ID:                  "ID",
		AudioURL:            "AudioURL",
		Title:               "Title",
		Caption:             "Caption",
		ParseMode:           "ParseMode",
		CaptionEntities:     []MessageEntity{{Type: "CaptionEntities"}},
		Performer:           "Performer",
		AudioDuration:       1,
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent: &InputTextMessageContent{},
	}, i)
}

func TestInlineQueryResultVoice_Setters(t *testing.T) {
	i := (&InlineQueryResultVoice{}).
		WithID("ID").
		WithVoiceURL("VoiceURL").
		WithTitle("Title").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithVoiceDuration(1).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultVoice{
		ID:                  "ID",
		VoiceURL:            "VoiceURL",
		Title:               "Title",
		Caption:             "Caption",
		ParseMode:           "ParseMode",
		CaptionEntities:     []MessageEntity{{Type: "CaptionEntities"}},
		VoiceDuration:       1,
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent: &InputTextMessageContent{},
	}, i)
}

func TestInlineQueryResultDocument_Setters(t *testing.T) {
	i := (&InlineQueryResultDocument{}).
		WithID("ID").
		WithTitle("Title").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithDocumentURL("DocumentURL").
		WithMimeType("MimeType").
		WithDescription("Description").
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{}).
		WithThumbnailURL("ThumbnailURL").
		WithThumbnailWidth(1).
		WithThumbnailHeight(2)

	assert.Equal(t, &InlineQueryResultDocument{
		ID:                  "ID",
		Title:               "Title",
		Caption:             "Caption",
		ParseMode:           "ParseMode",
		CaptionEntities:     []MessageEntity{{Type: "CaptionEntities"}},
		DocumentURL:         "DocumentURL",
		MimeType:            "MimeType",
		Description:         "Description",
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent: &InputTextMessageContent{},
		ThumbnailURL:        "ThumbnailURL",
		ThumbnailWidth:      1,
		ThumbnailHeight:     2,
	}, i)
}

func TestInlineQueryResultLocation_Setters(t *testing.T) {
	i := (&InlineQueryResultLocation{}).
		WithID("ID").
		WithLatitude(1.0).
		WithLongitude(2.0).
		WithTitle("Title").
		WithHorizontalAccuracy(3.0).
		WithLivePeriod(4).
		WithHeading(5).
		WithProximityAlertRadius(6).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{}).
		WithThumbnailURL("ThumbnailURL").
		WithThumbnailWidth(7).
		WithThumbnailHeight(8)

	assert.Equal(t, &InlineQueryResultLocation{
		ID:                   "ID",
		Latitude:             1.0,
		Longitude:            2.0,
		Title:                "Title",
		HorizontalAccuracy:   3.0,
		LivePeriod:           4,
		Heading:              5,
		ProximityAlertRadius: 6,
		ReplyMarkup:          &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent:  &InputTextMessageContent{},
		ThumbnailURL:         "ThumbnailURL",
		ThumbnailWidth:       7,
		ThumbnailHeight:      8,
	}, i)
}

func TestInlineQueryResultVenue_Setters(t *testing.T) {
	i := (&InlineQueryResultVenue{}).
		WithID("ID").
		WithLatitude(1.0).
		WithLongitude(2.0).
		WithTitle("Title").
		WithAddress("Address").
		WithFoursquareID("FoursquareID").
		WithFoursquareType("FoursquareType").
		WithGooglePlaceID("GooglePlaceID").
		WithGooglePlaceType("GooglePlaceType").
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{}).
		WithThumbnailURL("ThumbnailURL").
		WithThumbnailWidth(3).
		WithThumbnailHeight(4)

	assert.Equal(t, &InlineQueryResultVenue{
		ID:                  "ID",
		Latitude:            1.0,
		Longitude:           2.0,
		Title:               "Title",
		Address:             "Address",
		FoursquareID:        "FoursquareID",
		FoursquareType:      "FoursquareType",
		GooglePlaceID:       "GooglePlaceID",
		GooglePlaceType:     "GooglePlaceType",
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent: &InputTextMessageContent{},
		ThumbnailURL:        "ThumbnailURL",
		ThumbnailWidth:      3,
		ThumbnailHeight:     4,
	}, i)
}

func TestInlineQueryResultContact_Setters(t *testing.T) {
	i := (&InlineQueryResultContact{}).
		WithID("ID").
		WithPhoneNumber("PhoneNumber").
		WithFirstName("FirstName").
		WithLastName("LastName").
		WithVcard("Vcard").
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{}).
		WithThumbnailURL("ThumbnailURL").
		WithThumbnailWidth(1).
		WithThumbnailHeight(2)

	assert.Equal(t, &InlineQueryResultContact{
		ID:                  "ID",
		PhoneNumber:         "PhoneNumber",
		FirstName:           "FirstName",
		LastName:            "LastName",
		Vcard:               "Vcard",
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent: &InputTextMessageContent{},
		ThumbnailURL:        "ThumbnailURL",
		ThumbnailWidth:      1,
		ThumbnailHeight:     2,
	}, i)
}

func TestInlineQueryResultGame_Setters(t *testing.T) {
	i := (&InlineQueryResultGame{}).
		WithID("ID").
		WithGameShortName("GameShortName").
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}})

	assert.Equal(t, &InlineQueryResultGame{
		ID:            "ID",
		GameShortName: "GameShortName",
		ReplyMarkup:   &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
	}, i)
}

func TestInlineQueryResultCachedPhoto_Setters(t *testing.T) {
	i := (&InlineQueryResultCachedPhoto{}).
		WithID("ID").
		WithPhotoFileID("PhotoFileID").
		WithTitle("Title").
		WithDescription("Description").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithShowCaptionAboveMedia().
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultCachedPhoto{
		ID:                    "ID",
		PhotoFileID:           "PhotoFileID",
		Title:                 "Title",
		Description:           "Description",
		Caption:               "Caption",
		ParseMode:             "ParseMode",
		CaptionEntities:       []MessageEntity{{Type: "CaptionEntities"}},
		ShowCaptionAboveMedia: true,
		ReplyMarkup:           &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent:   &InputTextMessageContent{},
	}, i)
}

func TestInlineQueryResultCachedGif_Setters(t *testing.T) {
	i := (&InlineQueryResultCachedGif{}).
		WithID("ID").
		WithGifFileID("GifFileID").
		WithTitle("Title").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithShowCaptionAboveMedia().
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultCachedGif{
		ID:                    "ID",
		GifFileID:             "GifFileID",
		Title:                 "Title",
		Caption:               "Caption",
		ParseMode:             "ParseMode",
		CaptionEntities:       []MessageEntity{{Type: "CaptionEntities"}},
		ShowCaptionAboveMedia: true,
		ReplyMarkup:           &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent:   &InputTextMessageContent{},
	}, i)
}

func TestInlineQueryResultCachedMpeg4Gif_Setters(t *testing.T) {
	i := (&InlineQueryResultCachedMpeg4Gif{}).
		WithID("ID").
		WithMpeg4FileID("Mpeg4FileID").
		WithTitle("Title").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithShowCaptionAboveMedia().
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultCachedMpeg4Gif{
		ID:                    "ID",
		Mpeg4FileID:           "Mpeg4FileID",
		Title:                 "Title",
		Caption:               "Caption",
		ParseMode:             "ParseMode",
		CaptionEntities:       []MessageEntity{{Type: "CaptionEntities"}},
		ShowCaptionAboveMedia: true,
		ReplyMarkup:           &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent:   &InputTextMessageContent{},
	}, i)
}

func TestInlineQueryResultCachedSticker_Setters(t *testing.T) {
	i := (&InlineQueryResultCachedSticker{}).
		WithID("ID").
		WithStickerFileID("StickerFileID").
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultCachedSticker{
		ID:                  "ID",
		StickerFileID:       "StickerFileID",
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent: &InputTextMessageContent{},
	}, i)
}

func TestInlineQueryResultCachedDocument_Setters(t *testing.T) {
	i := (&InlineQueryResultCachedDocument{}).
		WithID("ID").
		WithTitle("Title").
		WithDocumentFileID("DocumentFileID").
		WithDescription("Description").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultCachedDocument{
		ID:                  "ID",
		Title:               "Title",
		DocumentFileID:      "DocumentFileID",
		Description:         "Description",
		Caption:             "Caption",
		ParseMode:           "ParseMode",
		CaptionEntities:     []MessageEntity{{Type: "CaptionEntities"}},
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent: &InputTextMessageContent{},
	}, i)
}

func TestInlineQueryResultCachedVideo_Setters(t *testing.T) {
	i := (&InlineQueryResultCachedVideo{}).
		WithID("ID").
		WithVideoFileID("VideoFileID").
		WithTitle("Title").
		WithDescription("Description").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithShowCaptionAboveMedia().
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultCachedVideo{
		ID:                    "ID",
		VideoFileID:           "VideoFileID",
		Title:                 "Title",
		Description:           "Description",
		Caption:               "Caption",
		ParseMode:             "ParseMode",
		CaptionEntities:       []MessageEntity{{Type: "CaptionEntities"}},
		ShowCaptionAboveMedia: true,
		ReplyMarkup:           &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent:   &InputTextMessageContent{},
	}, i)
}

func TestInlineQueryResultCachedVoice_Setters(t *testing.T) {
	i := (&InlineQueryResultCachedVoice{}).
		WithID("ID").
		WithVoiceFileID("VoiceFileID").
		WithTitle("Title").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultCachedVoice{
		ID:                  "ID",
		VoiceFileID:         "VoiceFileID",
		Title:               "Title",
		Caption:             "Caption",
		ParseMode:           "ParseMode",
		CaptionEntities:     []MessageEntity{{Type: "CaptionEntities"}},
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent: &InputTextMessageContent{},
	}, i)
}

func TestInlineQueryResultCachedAudio_Setters(t *testing.T) {
	i := (&InlineQueryResultCachedAudio{}).
		WithID("ID").
		WithAudioFileID("AudioFileID").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultCachedAudio{
		ID:                  "ID",
		AudioFileID:         "AudioFileID",
		Caption:             "Caption",
		ParseMode:           "ParseMode",
		CaptionEntities:     []MessageEntity{{Type: "CaptionEntities"}},
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent: &InputTextMessageContent{},
	}, i)
}

func TestInputTextMessageContent_Setters(t *testing.T) {
	i := (&InputTextMessageContent{}).
		WithMessageText("MessageText").
		WithParseMode("ParseMode").
		WithEntities([]MessageEntity{{Type: "Entities"}}...).
		WithLinkPreviewOptions(&LinkPreviewOptions{IsDisabled: true})

	assert.Equal(t, &InputTextMessageContent{
		MessageText:        "MessageText",
		ParseMode:          "ParseMode",
		Entities:           []MessageEntity{{Type: "Entities"}},
		LinkPreviewOptions: &LinkPreviewOptions{IsDisabled: true},
	}, i)
}

func TestInputRichMessageContent_Setters(t *testing.T) {
	i := (&InputRichMessageContent{}).
		WithRichMessage(InputRichMessage{})

	assert.Equal(t, &InputRichMessageContent{
		RichMessage: InputRichMessage{},
	}, i)
}

func TestInputLocationMessageContent_Setters(t *testing.T) {
	i := (&InputLocationMessageContent{}).
		WithLatitude(1.0).
		WithLongitude(1.0).
		WithHorizontalAccuracy(2.0).
		WithLivePeriod(3).
		WithHeading(4).
		WithProximityAlertRadius(5)

	assert.Equal(t, &InputLocationMessageContent{
		Latitude:             1.0,
		Longitude:            1.0,
		HorizontalAccuracy:   2.0,
		LivePeriod:           3,
		Heading:              4,
		ProximityAlertRadius: 5,
	}, i)
}

func TestInputVenueMessageContent_Setters(t *testing.T) {
	i := (&InputVenueMessageContent{}).
		WithLatitude(6.0).
		WithLongitude(1.0).
		WithTitle("Title").
		WithAddress("Address").
		WithFoursquareID("FoursquareID").
		WithFoursquareType("FoursquareType").
		WithGooglePlaceID("GooglePlaceID").
		WithGooglePlaceType("GooglePlaceType")

	assert.Equal(t, &InputVenueMessageContent{
		Latitude:        6.0,
		Longitude:       1.0,
		Title:           "Title",
		Address:         "Address",
		FoursquareID:    "FoursquareID",
		FoursquareType:  "FoursquareType",
		GooglePlaceID:   "GooglePlaceID",
		GooglePlaceType: "GooglePlaceType",
	}, i)
}

func TestInputContactMessageContent_Setters(t *testing.T) {
	i := (&InputContactMessageContent{}).
		WithPhoneNumber("PhoneNumber").
		WithFirstName("FirstName").
		WithLastName("LastName").
		WithVcard("Vcard")

	assert.Equal(t, &InputContactMessageContent{
		PhoneNumber: "PhoneNumber",
		FirstName:   "FirstName",
		LastName:    "LastName",
		Vcard:       "Vcard",
	}, i)
}

func TestInputInvoiceMessageContent_Setters(t *testing.T) {
	i := (&InputInvoiceMessageContent{}).
		WithTitle("Title").
		WithDescription("Description").
		WithPayload("Payload").
		WithProviderToken("ProviderToken").
		WithCurrency("Currency").
		WithPrices([]LabeledPrice{{Label: "Prices"}}...).
		WithMaxTipAmount(1).
		WithSuggestedTipAmounts([]int{2}...).
		WithProviderData("ProviderData").
		WithPhotoURL("PhotoURL").
		WithPhotoSize(3).
		WithPhotoWidth(4).
		WithPhotoHeight(5).
		WithNeedName().
		WithNeedPhoneNumber().
		WithNeedEmail().
		WithNeedShippingAddress().
		WithSendPhoneNumberToProvider().
		WithSendEmailToProvider().
		WithIsFlexible()

	assert.Equal(t, &InputInvoiceMessageContent{
		Title:                     "Title",
		Description:               "Description",
		Payload:                   "Payload",
		ProviderToken:             "ProviderToken",
		Currency:                  "Currency",
		Prices:                    []LabeledPrice{{Label: "Prices"}},
		MaxTipAmount:              1,
		SuggestedTipAmounts:       []int{2},
		ProviderData:              "ProviderData",
		PhotoURL:                  "PhotoURL",
		PhotoSize:                 3,
		PhotoWidth:                4,
		PhotoHeight:               5,
		NeedName:                  true,
		NeedPhoneNumber:           true,
		NeedEmail:                 true,
		NeedShippingAddress:       true,
		SendPhoneNumberToProvider: true,
		SendEmailToProvider:       true,
		IsFlexible:                true,
	}, i)
}
