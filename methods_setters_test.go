package telego

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUpdatesParams_Setters(t *testing.T) {
	g := (&GetUpdatesParams{}).
		WithOffset(1).
		WithLimit(2).
		WithTimeout(3).
		WithAllowedUpdates([]string{"AllowedUpdates"}...)

	assert.Equal(t, &GetUpdatesParams{
		Offset:         1,
		Limit:          2,
		Timeout:        3,
		AllowedUpdates: []string{"AllowedUpdates"},
	}, g)
}

func TestSetWebhookParams_Setters(t *testing.T) {
	s := (&SetWebhookParams{}).
		WithURL("URL").
		WithCertificate(&testInputFile).
		WithIPAddress("IPAddress").
		WithMaxConnections(1).
		WithAllowedUpdates([]string{"AllowedUpdates"}...).
		WithDropPendingUpdates().
		WithSecretToken("SecretToken")

	assert.Equal(t, &SetWebhookParams{
		URL:                "URL",
		Certificate:        &testInputFile,
		IPAddress:          "IPAddress",
		MaxConnections:     1,
		AllowedUpdates:     []string{"AllowedUpdates"},
		DropPendingUpdates: true,
		SecretToken:        "SecretToken",
	}, s)
}

func TestDeleteWebhookParams_Setters(t *testing.T) {
	d := (&DeleteWebhookParams{}).
		WithDropPendingUpdates()

	assert.Equal(t, &DeleteWebhookParams{
		DropPendingUpdates: true,
	}, d)
}

func TestSendMessageParams_Setters(t *testing.T) {
	s := (&SendMessageParams{}).
		WithBusinessConnectionID("BusinessConnectionID").
		WithChatID(ChatID{ID: 1}).
		WithMessageThreadID(2).
		WithText("Text").
		WithParseMode("ParseMode").
		WithEntities([]MessageEntity{{Type: "Entities"}}...).
		WithLinkPreviewOptions(&LinkPreviewOptions{IsDisabled: true}).
		WithDisableNotification().
		WithProtectContent().
		WithReplyParameters(&ReplyParameters{MessageID: 3}).
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendMessageParams{
		BusinessConnectionID: "BusinessConnectionID",
		ChatID:               ChatID{ID: 1},
		MessageThreadID:      2,
		Text:                 "Text",
		ParseMode:            "ParseMode",
		Entities:             []MessageEntity{{Type: "Entities"}},
		LinkPreviewOptions:   &LinkPreviewOptions{IsDisabled: true},
		DisableNotification:  true,
		ProtectContent:       true,
		ReplyParameters:      &ReplyParameters{MessageID: 3},
		ReplyMarkup:          &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestForwardMessageParams_Setters(t *testing.T) {
	f := (&ForwardMessageParams{}).
		WithChatID(ChatID{ID: 4}).
		WithMessageThreadID(1).
		WithFromChatID(ChatID{ID: 2}).
		WithDisableNotification().
		WithProtectContent().
		WithMessageID(3)

	assert.Equal(t, &ForwardMessageParams{
		ChatID:              ChatID{ID: 4},
		MessageThreadID:     1,
		FromChatID:          ChatID{ID: 2},
		DisableNotification: true,
		ProtectContent:      true,
		MessageID:           3,
	}, f)
}

func TestForwardMessagesParams_Setters(t *testing.T) {
	f := (&ForwardMessagesParams{}).
		WithChatID(ChatID{ID: 4}).
		WithMessageThreadID(1).
		WithFromChatID(ChatID{ID: 2}).
		WithMessageIDs([]int{3}...).
		WithDisableNotification().
		WithProtectContent()

	assert.Equal(t, &ForwardMessagesParams{
		ChatID:              ChatID{ID: 4},
		MessageThreadID:     1,
		FromChatID:          ChatID{ID: 2},
		MessageIDs:          []int{3},
		DisableNotification: true,
		ProtectContent:      true,
	}, f)
}

func TestCopyMessageParams_Setters(t *testing.T) {
	c := (&CopyMessageParams{}).
		WithChatID(ChatID{ID: 4}).
		WithMessageThreadID(1).
		WithFromChatID(ChatID{ID: 2}).
		WithMessageID(3).
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithDisableNotification().
		WithProtectContent().
		WithReplyParameters(&ReplyParameters{MessageID: 4}).
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &CopyMessageParams{
		ChatID:              ChatID{ID: 4},
		MessageThreadID:     1,
		FromChatID:          ChatID{ID: 2},
		MessageID:           3,
		Caption:             "Caption",
		ParseMode:           "ParseMode",
		CaptionEntities:     []MessageEntity{{Type: "CaptionEntities"}},
		DisableNotification: true,
		ProtectContent:      true,
		ReplyParameters:     &ReplyParameters{MessageID: 4},
		ReplyMarkup:         &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, c)
}

func TestCopyMessagesParams_Setters(t *testing.T) {
	c := (&CopyMessagesParams{}).
		WithChatID(ChatID{ID: 5}).
		WithMessageThreadID(1).
		WithFromChatID(ChatID{ID: 2}).
		WithMessageIDs([]int{3}...).
		WithDisableNotification().
		WithProtectContent().
		WithRemoveCaption()

	assert.Equal(t, &CopyMessagesParams{
		ChatID:              ChatID{ID: 5},
		MessageThreadID:     1,
		FromChatID:          ChatID{ID: 2},
		MessageIDs:          []int{3},
		DisableNotification: true,
		ProtectContent:      true,
		RemoveCaption:       true,
	}, c)
}

func TestSendPhotoParams_Setters(t *testing.T) {
	s := (&SendPhotoParams{}).
		WithBusinessConnectionID("BusinessConnectionID").
		WithChatID(ChatID{ID: 1}).
		WithMessageThreadID(2).
		WithPhoto(testInputFile).
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithHasSpoiler().
		WithDisableNotification().
		WithProtectContent().
		WithReplyParameters(&ReplyParameters{MessageID: 3}).
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendPhotoParams{
		BusinessConnectionID: "BusinessConnectionID",
		ChatID:               ChatID{ID: 1},
		MessageThreadID:      2,
		Photo:                testInputFile,
		Caption:              "Caption",
		ParseMode:            "ParseMode",
		CaptionEntities:      []MessageEntity{{Type: "CaptionEntities"}},
		HasSpoiler:           true,
		DisableNotification:  true,
		ProtectContent:       true,
		ReplyParameters:      &ReplyParameters{MessageID: 3},
		ReplyMarkup:          &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestSendAudioParams_Setters(t *testing.T) {
	s := (&SendAudioParams{}).
		WithBusinessConnectionID("BusinessConnectionID").
		WithChatID(ChatID{ID: 1}).
		WithMessageThreadID(2).
		WithAudio(testInputFile).
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithDuration(3).
		WithPerformer("Performer").
		WithTitle("Title").
		WithThumbnail(&testInputFile).
		WithDisableNotification().
		WithProtectContent().
		WithReplyParameters(&ReplyParameters{MessageID: 4}).
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendAudioParams{
		BusinessConnectionID: "BusinessConnectionID",
		ChatID:               ChatID{ID: 1},
		MessageThreadID:      2,
		Audio:                testInputFile,
		Caption:              "Caption",
		ParseMode:            "ParseMode",
		CaptionEntities:      []MessageEntity{{Type: "CaptionEntities"}},
		Duration:             3,
		Performer:            "Performer",
		Title:                "Title",
		Thumbnail:            &testInputFile,
		DisableNotification:  true,
		ProtectContent:       true,
		ReplyParameters:      &ReplyParameters{MessageID: 4},
		ReplyMarkup:          &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestSendDocumentParams_Setters(t *testing.T) {
	s := (&SendDocumentParams{}).
		WithBusinessConnectionID("BusinessConnectionID").
		WithChatID(ChatID{ID: 1}).
		WithMessageThreadID(2).
		WithDocument(testInputFile).
		WithThumbnail(&testInputFile).
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithDisableContentTypeDetection().
		WithDisableNotification().
		WithProtectContent().
		WithReplyParameters(&ReplyParameters{MessageID: 3}).
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendDocumentParams{
		BusinessConnectionID:        "BusinessConnectionID",
		ChatID:                      ChatID{ID: 1},
		MessageThreadID:             2,
		Document:                    testInputFile,
		Thumbnail:                   &testInputFile,
		Caption:                     "Caption",
		ParseMode:                   "ParseMode",
		CaptionEntities:             []MessageEntity{{Type: "CaptionEntities"}},
		DisableContentTypeDetection: true,
		DisableNotification:         true,
		ProtectContent:              true,
		ReplyParameters:             &ReplyParameters{MessageID: 3},
		ReplyMarkup:                 &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestSendVideoParams_Setters(t *testing.T) {
	s := (&SendVideoParams{}).
		WithBusinessConnectionID("BusinessConnectionID").
		WithChatID(ChatID{ID: 1}).
		WithMessageThreadID(2).
		WithVideo(testInputFile).
		WithDuration(3).
		WithWidth(4).
		WithHeight(5).
		WithThumbnail(&testInputFile).
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithHasSpoiler().
		WithSupportsStreaming().
		WithDisableNotification().
		WithProtectContent().
		WithReplyParameters(&ReplyParameters{MessageID: 6}).
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendVideoParams{
		BusinessConnectionID: "BusinessConnectionID",
		ChatID:               ChatID{ID: 1},
		MessageThreadID:      2,
		Video:                testInputFile,
		Duration:             3,
		Width:                4,
		Height:               5,
		Thumbnail:            &testInputFile,
		Caption:              "Caption",
		ParseMode:            "ParseMode",
		CaptionEntities:      []MessageEntity{{Type: "CaptionEntities"}},
		HasSpoiler:           true,
		SupportsStreaming:    true,
		DisableNotification:  true,
		ProtectContent:       true,
		ReplyParameters:      &ReplyParameters{MessageID: 6},
		ReplyMarkup:          &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestSendAnimationParams_Setters(t *testing.T) {
	s := (&SendAnimationParams{}).
		WithBusinessConnectionID("BusinessConnectionID").
		WithChatID(ChatID{ID: 1}).
		WithMessageThreadID(2).
		WithAnimation(testInputFile).
		WithDuration(3).
		WithWidth(4).
		WithHeight(5).
		WithThumbnail(&testInputFile).
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithHasSpoiler().
		WithDisableNotification().
		WithProtectContent().
		WithReplyParameters(&ReplyParameters{MessageID: 6}).
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendAnimationParams{
		BusinessConnectionID: "BusinessConnectionID",
		ChatID:               ChatID{ID: 1},
		MessageThreadID:      2,
		Animation:            testInputFile,
		Duration:             3,
		Width:                4,
		Height:               5,
		Thumbnail:            &testInputFile,
		Caption:              "Caption",
		ParseMode:            "ParseMode",
		CaptionEntities:      []MessageEntity{{Type: "CaptionEntities"}},
		HasSpoiler:           true,
		DisableNotification:  true,
		ProtectContent:       true,
		ReplyParameters:      &ReplyParameters{MessageID: 6},
		ReplyMarkup:          &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestSendVoiceParams_Setters(t *testing.T) {
	s := (&SendVoiceParams{}).
		WithBusinessConnectionID("BusinessConnectionID").
		WithChatID(ChatID{ID: 1}).
		WithMessageThreadID(2).
		WithVoice(testInputFile).
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithDuration(3).
		WithDisableNotification().
		WithProtectContent().
		WithReplyParameters(&ReplyParameters{MessageID: 4}).
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendVoiceParams{
		BusinessConnectionID: "BusinessConnectionID",
		ChatID:               ChatID{ID: 1},
		MessageThreadID:      2,
		Voice:                testInputFile,
		Caption:              "Caption",
		ParseMode:            "ParseMode",
		CaptionEntities:      []MessageEntity{{Type: "CaptionEntities"}},
		Duration:             3,
		DisableNotification:  true,
		ProtectContent:       true,
		ReplyParameters:      &ReplyParameters{MessageID: 4},
		ReplyMarkup:          &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestSendVideoNoteParams_Setters(t *testing.T) {
	s := (&SendVideoNoteParams{}).
		WithBusinessConnectionID("BusinessConnectionID").
		WithChatID(ChatID{ID: 1}).
		WithMessageThreadID(2).
		WithVideoNote(testInputFile).
		WithDuration(3).
		WithLength(4).
		WithThumbnail(&testInputFile).
		WithDisableNotification().
		WithProtectContent().
		WithReplyParameters(&ReplyParameters{MessageID: 5}).
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendVideoNoteParams{
		BusinessConnectionID: "BusinessConnectionID",
		ChatID:               ChatID{ID: 1},
		MessageThreadID:      2,
		VideoNote:            testInputFile,
		Duration:             3,
		Length:               4,
		Thumbnail:            &testInputFile,
		DisableNotification:  true,
		ProtectContent:       true,
		ReplyParameters:      &ReplyParameters{MessageID: 5},
		ReplyMarkup:          &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestSendMediaGroupParams_Setters(t *testing.T) {
	s := (&SendMediaGroupParams{}).
		WithBusinessConnectionID("BusinessConnectionID").
		WithChatID(ChatID{ID: 1}).
		WithMessageThreadID(2).
		WithMedia([]InputMedia{&InputMediaAnimation{Type: "Media"}}...).
		WithDisableNotification().
		WithProtectContent().
		WithReplyParameters(&ReplyParameters{MessageID: 3})

	assert.Equal(t, &SendMediaGroupParams{
		BusinessConnectionID: "BusinessConnectionID",
		ChatID:               ChatID{ID: 1},
		MessageThreadID:      2,
		Media:                []InputMedia{&InputMediaAnimation{Type: "Media"}},
		DisableNotification:  true,
		ProtectContent:       true,
		ReplyParameters:      &ReplyParameters{MessageID: 3},
	}, s)
}

func TestSendLocationParams_Setters(t *testing.T) {
	s := (&SendLocationParams{}).
		WithBusinessConnectionID("BusinessConnectionID").
		WithChatID(ChatID{ID: 1}).
		WithMessageThreadID(2).
		WithLivePeriod(3).
		WithHeading(4).
		WithProximityAlertRadius(5).
		WithDisableNotification().
		WithProtectContent().
		WithReplyParameters(&ReplyParameters{MessageID: 6}).
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendLocationParams{
		BusinessConnectionID: "BusinessConnectionID",
		ChatID:               ChatID{ID: 1},
		MessageThreadID:      2,
		LivePeriod:           3,
		Heading:              4,
		ProximityAlertRadius: 5,
		DisableNotification:  true,
		ProtectContent:       true,
		ReplyParameters:      &ReplyParameters{MessageID: 6},
		ReplyMarkup:          &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestSendVenueParams_Setters(t *testing.T) {
	s := (&SendVenueParams{}).
		WithBusinessConnectionID("BusinessConnectionID").
		WithChatID(ChatID{ID: 1}).
		WithMessageThreadID(2).
		WithTitle("Title").
		WithAddress("Address").
		WithFoursquareID("FoursquareID").
		WithFoursquareType("FoursquareType").
		WithGooglePlaceID("GooglePlaceID").
		WithGooglePlaceType("GooglePlaceType").
		WithDisableNotification().
		WithProtectContent().
		WithReplyParameters(&ReplyParameters{MessageID: 3}).
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendVenueParams{
		BusinessConnectionID: "BusinessConnectionID",
		ChatID:               ChatID{ID: 1},
		MessageThreadID:      2,
		Title:                "Title",
		Address:              "Address",
		FoursquareID:         "FoursquareID",
		FoursquareType:       "FoursquareType",
		GooglePlaceID:        "GooglePlaceID",
		GooglePlaceType:      "GooglePlaceType",
		DisableNotification:  true,
		ProtectContent:       true,
		ReplyParameters:      &ReplyParameters{MessageID: 3},
		ReplyMarkup:          &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestSendContactParams_Setters(t *testing.T) {
	s := (&SendContactParams{}).
		WithBusinessConnectionID("BusinessConnectionID").
		WithChatID(ChatID{ID: 1}).
		WithMessageThreadID(2).
		WithPhoneNumber("PhoneNumber").
		WithFirstName("FirstName").
		WithLastName("LastName").
		WithVcard("Vcard").
		WithDisableNotification().
		WithProtectContent().
		WithReplyParameters(&ReplyParameters{MessageID: 3}).
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendContactParams{
		BusinessConnectionID: "BusinessConnectionID",
		ChatID:               ChatID{ID: 1},
		MessageThreadID:      2,
		PhoneNumber:          "PhoneNumber",
		FirstName:            "FirstName",
		LastName:             "LastName",
		Vcard:                "Vcard",
		DisableNotification:  true,
		ProtectContent:       true,
		ReplyParameters:      &ReplyParameters{MessageID: 3},
		ReplyMarkup:          &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestSendPollParams_Setters(t *testing.T) {
	s := (&SendPollParams{}).
		WithBusinessConnectionID("BusinessConnectionID").
		WithChatID(ChatID{ID: 1}).
		WithMessageThreadID(2).
		WithQuestion("Question").
		WithQuestionParseMode("QuestionParseMode").
		WithQuestionEntities([]MessageEntity{{Type: "QuestionEntities"}}...).
		WithOptions([]InputPollOption{{}}...).
		WithIsAnonymous(true).
		WithType("Type").
		WithAllowsMultipleAnswers().
		WithCorrectOptionID(3).
		WithExplanation("Explanation").
		WithExplanationParseMode("ExplanationParseMode").
		WithExplanationEntities([]MessageEntity{{Type: "ExplanationEntities"}}...).
		WithOpenPeriod(4).
		WithIsClosed().
		WithDisableNotification().
		WithProtectContent().
		WithReplyParameters(&ReplyParameters{MessageID: 5}).
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendPollParams{
		BusinessConnectionID:  "BusinessConnectionID",
		ChatID:                ChatID{ID: 1},
		MessageThreadID:       2,
		Question:              "Question",
		QuestionParseMode:     "QuestionParseMode",
		QuestionEntities:      []MessageEntity{{Type: "QuestionEntities"}},
		Options:               []InputPollOption{{}},
		IsAnonymous:           ToPtr(true),
		Type:                  "Type",
		AllowsMultipleAnswers: true,
		CorrectOptionID:       ToPtr(3),
		Explanation:           "Explanation",
		ExplanationParseMode:  "ExplanationParseMode",
		ExplanationEntities:   []MessageEntity{{Type: "ExplanationEntities"}},
		OpenPeriod:            4,
		IsClosed:              true,
		DisableNotification:   true,
		ProtectContent:        true,
		ReplyParameters:       &ReplyParameters{MessageID: 5},
		ReplyMarkup:           &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestSendDiceParams_Setters(t *testing.T) {
	s := (&SendDiceParams{}).
		WithBusinessConnectionID("BusinessConnectionID").
		WithChatID(ChatID{ID: 1}).
		WithMessageThreadID(2).
		WithEmoji("Emoji").
		WithDisableNotification().
		WithProtectContent().
		WithReplyParameters(&ReplyParameters{MessageID: 3}).
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendDiceParams{
		BusinessConnectionID: "BusinessConnectionID",
		ChatID:               ChatID{ID: 1},
		MessageThreadID:      2,
		Emoji:                "Emoji",
		DisableNotification:  true,
		ProtectContent:       true,
		ReplyParameters:      &ReplyParameters{MessageID: 3},
		ReplyMarkup:          &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestSendChatActionParams_Setters(t *testing.T) {
	s := (&SendChatActionParams{}).
		WithBusinessConnectionID("BusinessConnectionID").
		WithChatID(ChatID{ID: 1}).
		WithMessageThreadID(2).
		WithAction("Action")

	assert.Equal(t, &SendChatActionParams{
		BusinessConnectionID: "BusinessConnectionID",
		ChatID:               ChatID{ID: 1},
		MessageThreadID:      2,
		Action:               "Action",
	}, s)
}

func TestSetMessageReactionParams_Setters(t *testing.T) {
	s := (&SetMessageReactionParams{}).
		WithChatID(ChatID{ID: 3}).
		WithMessageID(1).
		WithReaction([]ReactionType{&ReactionTypeEmoji{Type: ReactionEmoji}}...).
		WithIsBig()

	assert.Equal(t, &SetMessageReactionParams{
		ChatID:    ChatID{ID: 3},
		MessageID: 1,
		Reaction:  []ReactionType{&ReactionTypeEmoji{Type: ReactionEmoji}},
		IsBig:     true,
	}, s)
}

func TestGetUserProfilePhotosParams_Setters(t *testing.T) {
	g := (&GetUserProfilePhotosParams{}).
		WithOffset(2).
		WithLimit(1)

	assert.Equal(t, &GetUserProfilePhotosParams{
		Offset: 2,
		Limit:  1,
	}, g)
}

func TestGetFileParams_Setters(t *testing.T) {
	g := (&GetFileParams{}).
		WithFileID("FileID")

	assert.Equal(t, &GetFileParams{
		FileID: "FileID",
	}, g)
}

func TestBanChatMemberParams_Setters(t *testing.T) {
	b := (&BanChatMemberParams{}).
		WithChatID(ChatID{ID: 1}).
		WithRevokeMessages()

	assert.Equal(t, &BanChatMemberParams{
		ChatID:         ChatID{ID: 1},
		RevokeMessages: true,
	}, b)
}

func TestUnbanChatMemberParams_Setters(t *testing.T) {
	u := (&UnbanChatMemberParams{}).
		WithChatID(ChatID{ID: 1}).
		WithOnlyIfBanned()

	assert.Equal(t, &UnbanChatMemberParams{
		ChatID:       ChatID{ID: 1},
		OnlyIfBanned: true,
	}, u)
}

func TestRestrictChatMemberParams_Setters(t *testing.T) {
	r := (&RestrictChatMemberParams{}).
		WithChatID(ChatID{ID: 1}).
		WithPermissions(ChatPermissions{CanSendMessages: ToPtr(true)}).
		WithUseIndependentChatPermissions()

	assert.Equal(t, &RestrictChatMemberParams{
		ChatID:                        ChatID{ID: 1},
		Permissions:                   ChatPermissions{CanSendMessages: ToPtr(true)},
		UseIndependentChatPermissions: true,
	}, r)
}

func TestPromoteChatMemberParams_Setters(t *testing.T) {
	p := (&PromoteChatMemberParams{}).
		WithChatID(ChatID{ID: 1}).
		WithIsAnonymous(true).
		WithCanManageChat(true).
		WithCanDeleteMessages(true).
		WithCanManageVideoChats(true).
		WithCanRestrictMembers(true).
		WithCanPromoteMembers(true).
		WithCanChangeInfo(true).
		WithCanInviteUsers(true).
		WithCanPostStories(true).
		WithCanEditStories(true).
		WithCanDeleteStories(true).
		WithCanPostMessages(true).
		WithCanEditMessages(true).
		WithCanPinMessages(true).
		WithCanManageTopics(true)

	assert.Equal(t, &PromoteChatMemberParams{
		ChatID:              ChatID{ID: 1},
		IsAnonymous:         ToPtr(true),
		CanManageChat:       ToPtr(true),
		CanDeleteMessages:   ToPtr(true),
		CanManageVideoChats: ToPtr(true),
		CanRestrictMembers:  ToPtr(true),
		CanPromoteMembers:   ToPtr(true),
		CanChangeInfo:       ToPtr(true),
		CanInviteUsers:      ToPtr(true),
		CanPostStories:      ToPtr(true),
		CanEditStories:      ToPtr(true),
		CanDeleteStories:    ToPtr(true),
		CanPostMessages:     ToPtr(true),
		CanEditMessages:     ToPtr(true),
		CanPinMessages:      ToPtr(true),
		CanManageTopics:     ToPtr(true),
	}, p)
}

func TestSetChatAdministratorCustomTitleParams_Setters(t *testing.T) {
	s := (&SetChatAdministratorCustomTitleParams{}).
		WithChatID(ChatID{ID: 1}).
		WithCustomTitle("CustomTitle")

	assert.Equal(t, &SetChatAdministratorCustomTitleParams{
		ChatID:      ChatID{ID: 1},
		CustomTitle: "CustomTitle",
	}, s)
}

func TestBanChatSenderChatParams_Setters(t *testing.T) {
	b := (&BanChatSenderChatParams{}).
		WithChatID(ChatID{ID: 1})

	assert.Equal(t, &BanChatSenderChatParams{
		ChatID: ChatID{ID: 1},
	}, b)
}

func TestUnbanChatSenderChatParams_Setters(t *testing.T) {
	u := (&UnbanChatSenderChatParams{}).
		WithChatID(ChatID{ID: 1})

	assert.Equal(t, &UnbanChatSenderChatParams{
		ChatID: ChatID{ID: 1},
	}, u)
}

func TestSetChatPermissionsParams_Setters(t *testing.T) {
	s := (&SetChatPermissionsParams{}).
		WithChatID(ChatID{ID: 1}).
		WithPermissions(ChatPermissions{CanSendMessages: ToPtr(true)}).
		WithUseIndependentChatPermissions()

	assert.Equal(t, &SetChatPermissionsParams{
		ChatID:                        ChatID{ID: 1},
		Permissions:                   ChatPermissions{CanSendMessages: ToPtr(true)},
		UseIndependentChatPermissions: true,
	}, s)
}

func TestExportChatInviteLinkParams_Setters(t *testing.T) {
	e := (&ExportChatInviteLinkParams{}).
		WithChatID(ChatID{ID: 1})

	assert.Equal(t, &ExportChatInviteLinkParams{
		ChatID: ChatID{ID: 1},
	}, e)
}

func TestCreateChatInviteLinkParams_Setters(t *testing.T) {
	c := (&CreateChatInviteLinkParams{}).
		WithChatID(ChatID{ID: 1}).
		WithName("Name").
		WithMemberLimit(1).
		WithCreatesJoinRequest()

	assert.Equal(t, &CreateChatInviteLinkParams{
		ChatID:             ChatID{ID: 1},
		Name:               "Name",
		MemberLimit:        1,
		CreatesJoinRequest: true,
	}, c)
}

func TestEditChatInviteLinkParams_Setters(t *testing.T) {
	e := (&EditChatInviteLinkParams{}).
		WithChatID(ChatID{ID: 2}).
		WithInviteLink("InviteLink").
		WithName("Name").
		WithMemberLimit(1).
		WithCreatesJoinRequest()

	assert.Equal(t, &EditChatInviteLinkParams{
		ChatID:             ChatID{ID: 2},
		InviteLink:         "InviteLink",
		Name:               "Name",
		MemberLimit:        1,
		CreatesJoinRequest: true,
	}, e)
}

func TestRevokeChatInviteLinkParams_Setters(t *testing.T) {
	r := (&RevokeChatInviteLinkParams{}).
		WithChatID(ChatID{ID: 2}).
		WithInviteLink("InviteLink")

	assert.Equal(t, &RevokeChatInviteLinkParams{
		ChatID:     ChatID{ID: 2},
		InviteLink: "InviteLink",
	}, r)
}

func TestApproveChatJoinRequestParams_Setters(t *testing.T) {
	a := (&ApproveChatJoinRequestParams{}).
		WithChatID(ChatID{ID: 1})

	assert.Equal(t, &ApproveChatJoinRequestParams{
		ChatID: ChatID{ID: 1},
	}, a)
}

func TestDeclineChatJoinRequestParams_Setters(t *testing.T) {
	d := (&DeclineChatJoinRequestParams{}).
		WithChatID(ChatID{ID: 1})

	assert.Equal(t, &DeclineChatJoinRequestParams{
		ChatID: ChatID{ID: 1},
	}, d)
}

func TestSetChatPhotoParams_Setters(t *testing.T) {
	s := (&SetChatPhotoParams{}).
		WithChatID(ChatID{ID: 1}).
		WithPhoto(testInputFile)

	assert.Equal(t, &SetChatPhotoParams{
		ChatID: ChatID{ID: 1},
		Photo:  testInputFile,
	}, s)
}

func TestDeleteChatPhotoParams_Setters(t *testing.T) {
	d := (&DeleteChatPhotoParams{}).
		WithChatID(ChatID{ID: 1})

	assert.Equal(t, &DeleteChatPhotoParams{
		ChatID: ChatID{ID: 1},
	}, d)
}

func TestSetChatTitleParams_Setters(t *testing.T) {
	s := (&SetChatTitleParams{}).
		WithChatID(ChatID{ID: 1}).
		WithTitle("Title")

	assert.Equal(t, &SetChatTitleParams{
		ChatID: ChatID{ID: 1},
		Title:  "Title",
	}, s)
}

func TestSetChatDescriptionParams_Setters(t *testing.T) {
	s := (&SetChatDescriptionParams{}).
		WithChatID(ChatID{ID: 1}).
		WithDescription("Description")

	assert.Equal(t, &SetChatDescriptionParams{
		ChatID:      ChatID{ID: 1},
		Description: "Description",
	}, s)
}

func TestPinChatMessageParams_Setters(t *testing.T) {
	p := (&PinChatMessageParams{}).
		WithChatID(ChatID{ID: 1}).
		WithMessageID(1).
		WithDisableNotification()

	assert.Equal(t, &PinChatMessageParams{
		ChatID:              ChatID{ID: 1},
		MessageID:           1,
		DisableNotification: true,
	}, p)
}

func TestUnpinChatMessageParams_Setters(t *testing.T) {
	u := (&UnpinChatMessageParams{}).
		WithChatID(ChatID{ID: 2}).
		WithMessageID(1)

	assert.Equal(t, &UnpinChatMessageParams{
		ChatID:    ChatID{ID: 2},
		MessageID: 1,
	}, u)
}

func TestUnpinAllChatMessagesParams_Setters(t *testing.T) {
	u := (&UnpinAllChatMessagesParams{}).
		WithChatID(ChatID{ID: 2})

	assert.Equal(t, &UnpinAllChatMessagesParams{
		ChatID: ChatID{ID: 2},
	}, u)
}

func TestLeaveChatParams_Setters(t *testing.T) {
	l := (&LeaveChatParams{}).
		WithChatID(ChatID{ID: 1})

	assert.Equal(t, &LeaveChatParams{
		ChatID: ChatID{ID: 1},
	}, l)
}

func TestGetChatParams_Setters(t *testing.T) {
	g := (&GetChatParams{}).
		WithChatID(ChatID{ID: 1})

	assert.Equal(t, &GetChatParams{
		ChatID: ChatID{ID: 1},
	}, g)
}

func TestGetChatAdministratorsParams_Setters(t *testing.T) {
	g := (&GetChatAdministratorsParams{}).
		WithChatID(ChatID{ID: 1})

	assert.Equal(t, &GetChatAdministratorsParams{
		ChatID: ChatID{ID: 1},
	}, g)
}

func TestGetChatMemberCountParams_Setters(t *testing.T) {
	g := (&GetChatMemberCountParams{}).
		WithChatID(ChatID{ID: 1})

	assert.Equal(t, &GetChatMemberCountParams{
		ChatID: ChatID{ID: 1},
	}, g)
}

func TestGetChatMemberParams_Setters(t *testing.T) {
	g := (&GetChatMemberParams{}).
		WithChatID(ChatID{ID: 1})

	assert.Equal(t, &GetChatMemberParams{
		ChatID: ChatID{ID: 1},
	}, g)
}

func TestSetChatStickerSetParams_Setters(t *testing.T) {
	s := (&SetChatStickerSetParams{}).
		WithChatID(ChatID{ID: 1}).
		WithStickerSetName("StickerSetName")

	assert.Equal(t, &SetChatStickerSetParams{
		ChatID:         ChatID{ID: 1},
		StickerSetName: "StickerSetName",
	}, s)
}

func TestDeleteChatStickerSetParams_Setters(t *testing.T) {
	d := (&DeleteChatStickerSetParams{}).
		WithChatID(ChatID{ID: 1})

	assert.Equal(t, &DeleteChatStickerSetParams{
		ChatID: ChatID{ID: 1},
	}, d)
}

func TestCreateForumTopicParams_Setters(t *testing.T) {
	c := (&CreateForumTopicParams{}).
		WithChatID(ChatID{ID: 1}).
		WithName("Name").
		WithIconColor(1).
		WithIconCustomEmojiID("IconCustomEmojiID")

	assert.Equal(t, &CreateForumTopicParams{
		ChatID:            ChatID{ID: 1},
		Name:              "Name",
		IconColor:         1,
		IconCustomEmojiID: "IconCustomEmojiID",
	}, c)
}

func TestEditForumTopicParams_Setters(t *testing.T) {
	e := (&EditForumTopicParams{}).
		WithChatID(ChatID{ID: 2}).
		WithMessageThreadID(1).
		WithName("Name").
		WithIconCustomEmojiID("IconCustomEmojiID")

	assert.Equal(t, &EditForumTopicParams{
		ChatID:            ChatID{ID: 2},
		MessageThreadID:   1,
		Name:              "Name",
		IconCustomEmojiID: ToPtr("IconCustomEmojiID"),
	}, e)
}

func TestCloseForumTopicParams_Setters(t *testing.T) {
	c := (&CloseForumTopicParams{}).
		WithChatID(ChatID{ID: 2}).
		WithMessageThreadID(1)

	assert.Equal(t, &CloseForumTopicParams{
		ChatID:          ChatID{ID: 2},
		MessageThreadID: 1,
	}, c)
}

func TestReopenForumTopicParams_Setters(t *testing.T) {
	r := (&ReopenForumTopicParams{}).
		WithChatID(ChatID{ID: 2}).
		WithMessageThreadID(1)

	assert.Equal(t, &ReopenForumTopicParams{
		ChatID:          ChatID{ID: 2},
		MessageThreadID: 1,
	}, r)
}

func TestDeleteForumTopicParams_Setters(t *testing.T) {
	d := (&DeleteForumTopicParams{}).
		WithChatID(ChatID{ID: 2}).
		WithMessageThreadID(1)

	assert.Equal(t, &DeleteForumTopicParams{
		ChatID:          ChatID{ID: 2},
		MessageThreadID: 1,
	}, d)
}

func TestUnpinAllForumTopicMessagesParams_Setters(t *testing.T) {
	u := (&UnpinAllForumTopicMessagesParams{}).
		WithChatID(ChatID{ID: 2}).
		WithMessageThreadID(1)

	assert.Equal(t, &UnpinAllForumTopicMessagesParams{
		ChatID:          ChatID{ID: 2},
		MessageThreadID: 1,
	}, u)
}

func TestEditGeneralForumTopicParams_Setters(t *testing.T) {
	e := (&EditGeneralForumTopicParams{}).
		WithChatID(ChatID{ID: 2}).
		WithName("Name")

	assert.Equal(t, &EditGeneralForumTopicParams{
		ChatID: ChatID{ID: 2},
		Name:   "Name",
	}, e)
}

func TestCloseGeneralForumTopicParams_Setters(t *testing.T) {
	c := (&CloseGeneralForumTopicParams{}).
		WithChatID(ChatID{ID: 1})

	assert.Equal(t, &CloseGeneralForumTopicParams{
		ChatID: ChatID{ID: 1},
	}, c)
}

func TestReopenGeneralForumTopicParams_Setters(t *testing.T) {
	r := (&ReopenGeneralForumTopicParams{}).
		WithChatID(ChatID{ID: 1})

	assert.Equal(t, &ReopenGeneralForumTopicParams{
		ChatID: ChatID{ID: 1},
	}, r)
}

func TestHideGeneralForumTopicParams_Setters(t *testing.T) {
	h := (&HideGeneralForumTopicParams{}).
		WithChatID(ChatID{ID: 1})

	assert.Equal(t, &HideGeneralForumTopicParams{
		ChatID: ChatID{ID: 1},
	}, h)
}

func TestUnhideGeneralForumTopicParams_Setters(t *testing.T) {
	u := (&UnhideGeneralForumTopicParams{}).
		WithChatID(ChatID{ID: 1})

	assert.Equal(t, &UnhideGeneralForumTopicParams{
		ChatID: ChatID{ID: 1},
	}, u)
}

func TestUnpinAllGeneralForumTopicMessagesParams_Setters(t *testing.T) {
	u := (&UnpinAllGeneralForumTopicMessagesParams{}).
		WithChatID(ChatID{ID: 1})

	assert.Equal(t, &UnpinAllGeneralForumTopicMessagesParams{
		ChatID: ChatID{ID: 1},
	}, u)
}

func TestAnswerCallbackQueryParams_Setters(t *testing.T) {
	a := (&AnswerCallbackQueryParams{}).
		WithCallbackQueryID("CallbackQueryID").
		WithText("Text").
		WithShowAlert().
		WithURL("URL").
		WithCacheTime(1)

	assert.Equal(t, &AnswerCallbackQueryParams{
		CallbackQueryID: "CallbackQueryID",
		Text:            "Text",
		ShowAlert:       true,
		URL:             "URL",
		CacheTime:       1,
	}, a)
}

func TestGetUserChatBoostsParams_Setters(t *testing.T) {
	g := (&GetUserChatBoostsParams{}).
		WithChatID(ChatID{ID: 2})

	assert.Equal(t, &GetUserChatBoostsParams{
		ChatID: ChatID{ID: 2},
	}, g)
}

func TestGetBusinessConnectionParams_Setters(t *testing.T) {
	g := (&GetBusinessConnectionParams{}).
		WithBusinessConnectionID("BusinessConnectionID")

	assert.Equal(t, &GetBusinessConnectionParams{
		BusinessConnectionID: "BusinessConnectionID",
	}, g)
}

func TestSetMyCommandsParams_Setters(t *testing.T) {
	s := (&SetMyCommandsParams{}).
		WithCommands([]BotCommand{{Command: "Commands"}}...).
		WithScope(&BotCommandScopeDefault{Type: "Scope"}).
		WithLanguageCode("LanguageCode")

	assert.Equal(t, &SetMyCommandsParams{
		Commands:     []BotCommand{{Command: "Commands"}},
		Scope:        &BotCommandScopeDefault{Type: "Scope"},
		LanguageCode: "LanguageCode",
	}, s)
}

func TestDeleteMyCommandsParams_Setters(t *testing.T) {
	d := (&DeleteMyCommandsParams{}).
		WithScope(&BotCommandScopeDefault{Type: "Scope"}).
		WithLanguageCode("LanguageCode")

	assert.Equal(t, &DeleteMyCommandsParams{
		Scope:        &BotCommandScopeDefault{Type: "Scope"},
		LanguageCode: "LanguageCode",
	}, d)
}

func TestGetMyCommandsParams_Setters(t *testing.T) {
	g := (&GetMyCommandsParams{}).
		WithScope(&BotCommandScopeDefault{Type: "Scope"}).
		WithLanguageCode("LanguageCode")

	assert.Equal(t, &GetMyCommandsParams{
		Scope:        &BotCommandScopeDefault{Type: "Scope"},
		LanguageCode: "LanguageCode",
	}, g)
}

func TestSetMyNameParams_Setters(t *testing.T) {
	s := (&SetMyNameParams{}).
		WithName("Name").
		WithLanguageCode("LanguageCode")

	assert.Equal(t, &SetMyNameParams{
		Name:         "Name",
		LanguageCode: "LanguageCode",
	}, s)
}

func TestGetMyNameParams_Setters(t *testing.T) {
	g := (&GetMyNameParams{}).
		WithLanguageCode("LanguageCode")

	assert.Equal(t, &GetMyNameParams{
		LanguageCode: "LanguageCode",
	}, g)
}

func TestSetMyDescriptionParams_Setters(t *testing.T) {
	s := (&SetMyDescriptionParams{}).
		WithDescription("Description").
		WithLanguageCode("LanguageCode")

	assert.Equal(t, &SetMyDescriptionParams{
		Description:  "Description",
		LanguageCode: "LanguageCode",
	}, s)
}

func TestGetMyDescriptionParams_Setters(t *testing.T) {
	g := (&GetMyDescriptionParams{}).
		WithLanguageCode("LanguageCode")

	assert.Equal(t, &GetMyDescriptionParams{
		LanguageCode: "LanguageCode",
	}, g)
}

func TestSetMyShortDescriptionParams_Setters(t *testing.T) {
	s := (&SetMyShortDescriptionParams{}).
		WithShortDescription("ShortDescription").
		WithLanguageCode("LanguageCode")

	assert.Equal(t, &SetMyShortDescriptionParams{
		ShortDescription: "ShortDescription",
		LanguageCode:     "LanguageCode",
	}, s)
}

func TestGetMyShortDescriptionParams_Setters(t *testing.T) {
	g := (&GetMyShortDescriptionParams{}).
		WithLanguageCode("LanguageCode")

	assert.Equal(t, &GetMyShortDescriptionParams{
		LanguageCode: "LanguageCode",
	}, g)
}

func TestSetChatMenuButtonParams_Setters(t *testing.T) {
	s := (&SetChatMenuButtonParams{}).
		WithMenuButton(&MenuButtonCommands{Type: "MenuButton"})

	assert.Equal(t, &SetChatMenuButtonParams{
		MenuButton: &MenuButtonCommands{Type: "MenuButton"},
	}, s)
}

func TestSetMyDefaultAdministratorRightsParams_Setters(t *testing.T) {
	s := (&SetMyDefaultAdministratorRightsParams{}).
		WithRights(&ChatAdministratorRights{IsAnonymous: true}).
		WithForChannels()

	assert.Equal(t, &SetMyDefaultAdministratorRightsParams{
		Rights:      &ChatAdministratorRights{IsAnonymous: true},
		ForChannels: true,
	}, s)
}

func TestGetMyDefaultAdministratorRightsParams_Setters(t *testing.T) {
	g := (&GetMyDefaultAdministratorRightsParams{}).
		WithForChannels()

	assert.Equal(t, &GetMyDefaultAdministratorRightsParams{
		ForChannels: true,
	}, g)
}

func TestEditMessageTextParams_Setters(t *testing.T) {
	e := (&EditMessageTextParams{}).
		WithChatID(ChatID{ID: 1}).
		WithMessageID(1).
		WithInlineMessageID("InlineMessageID").
		WithText("Text").
		WithParseMode("ParseMode").
		WithEntities([]MessageEntity{{Type: "Entities"}}...).
		WithLinkPreviewOptions(&LinkPreviewOptions{IsDisabled: true}).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}})

	assert.Equal(t, &EditMessageTextParams{
		ChatID:             ChatID{ID: 1},
		MessageID:          1,
		InlineMessageID:    "InlineMessageID",
		Text:               "Text",
		ParseMode:          "ParseMode",
		Entities:           []MessageEntity{{Type: "Entities"}},
		LinkPreviewOptions: &LinkPreviewOptions{IsDisabled: true},
		ReplyMarkup:        &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
	}, e)
}

func TestEditMessageCaptionParams_Setters(t *testing.T) {
	e := (&EditMessageCaptionParams{}).
		WithChatID(ChatID{ID: 2}).
		WithMessageID(1).
		WithInlineMessageID("InlineMessageID").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}})

	assert.Equal(t, &EditMessageCaptionParams{
		ChatID:          ChatID{ID: 2},
		MessageID:       1,
		InlineMessageID: "InlineMessageID",
		Caption:         "Caption",
		ParseMode:       "ParseMode",
		CaptionEntities: []MessageEntity{{Type: "CaptionEntities"}},
		ReplyMarkup:     &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
	}, e)
}

func TestEditMessageMediaParams_Setters(t *testing.T) {
	e := (&EditMessageMediaParams{}).
		WithChatID(ChatID{ID: 2}).
		WithMessageID(1).
		WithInlineMessageID("InlineMessageID").
		WithMedia(&InputMediaAnimation{Type: "Media"}).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}})

	assert.Equal(t, &EditMessageMediaParams{
		ChatID:          ChatID{ID: 2},
		MessageID:       1,
		InlineMessageID: "InlineMessageID",
		Media:           &InputMediaAnimation{Type: "Media"},
		ReplyMarkup:     &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
	}, e)
}

func TestEditMessageLiveLocationParams_Setters(t *testing.T) {
	e := (&EditMessageLiveLocationParams{}).
		WithChatID(ChatID{ID: 2}).
		WithMessageID(1).
		WithInlineMessageID("InlineMessageID").
		WithLivePeriod(2).
		WithHeading(3).
		WithProximityAlertRadius(4).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}})

	assert.Equal(t, &EditMessageLiveLocationParams{
		ChatID:               ChatID{ID: 2},
		MessageID:            1,
		InlineMessageID:      "InlineMessageID",
		LivePeriod:           2,
		Heading:              3,
		ProximityAlertRadius: 4,
		ReplyMarkup:          &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
	}, e)
}

func TestStopMessageLiveLocationParams_Setters(t *testing.T) {
	s := (&StopMessageLiveLocationParams{}).
		WithChatID(ChatID{ID: 5}).
		WithMessageID(1).
		WithInlineMessageID("InlineMessageID").
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}})

	assert.Equal(t, &StopMessageLiveLocationParams{
		ChatID:          ChatID{ID: 5},
		MessageID:       1,
		InlineMessageID: "InlineMessageID",
		ReplyMarkup:     &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
	}, s)
}

func TestEditMessageReplyMarkupParams_Setters(t *testing.T) {
	e := (&EditMessageReplyMarkupParams{}).
		WithChatID(ChatID{ID: 2}).
		WithMessageID(1).
		WithInlineMessageID("InlineMessageID").
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}})

	assert.Equal(t, &EditMessageReplyMarkupParams{
		ChatID:          ChatID{ID: 2},
		MessageID:       1,
		InlineMessageID: "InlineMessageID",
		ReplyMarkup:     &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
	}, e)
}

func TestStopPollParams_Setters(t *testing.T) {
	s := (&StopPollParams{}).
		WithChatID(ChatID{ID: 2}).
		WithMessageID(1).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}})

	assert.Equal(t, &StopPollParams{
		ChatID:      ChatID{ID: 2},
		MessageID:   1,
		ReplyMarkup: &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
	}, s)
}

func TestDeleteMessageParams_Setters(t *testing.T) {
	d := (&DeleteMessageParams{}).
		WithChatID(ChatID{ID: 2}).
		WithMessageID(1)

	assert.Equal(t, &DeleteMessageParams{
		ChatID:    ChatID{ID: 2},
		MessageID: 1,
	}, d)
}

func TestDeleteMessagesParams_Setters(t *testing.T) {
	d := (&DeleteMessagesParams{}).
		WithChatID(ChatID{ID: 2}).
		WithMessageIDs([]int{1}...)

	assert.Equal(t, &DeleteMessagesParams{
		ChatID:     ChatID{ID: 2},
		MessageIDs: []int{1},
	}, d)
}

func TestSendStickerParams_Setters(t *testing.T) {
	s := (&SendStickerParams{}).
		WithBusinessConnectionID("BusinessConnectionID").
		WithChatID(ChatID{ID: 1}).
		WithMessageThreadID(2).
		WithSticker(testInputFile).
		WithEmoji("Emoji").
		WithDisableNotification().
		WithProtectContent().
		WithReplyParameters(&ReplyParameters{MessageID: 3}).
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendStickerParams{
		BusinessConnectionID: "BusinessConnectionID",
		ChatID:               ChatID{ID: 1},
		MessageThreadID:      2,
		Sticker:              testInputFile,
		Emoji:                "Emoji",
		DisableNotification:  true,
		ProtectContent:       true,
		ReplyParameters:      &ReplyParameters{MessageID: 3},
		ReplyMarkup:          &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestGetStickerSetParams_Setters(t *testing.T) {
	g := (&GetStickerSetParams{}).
		WithName("Name")

	assert.Equal(t, &GetStickerSetParams{
		Name: "Name",
	}, g)
}

func TestGetCustomEmojiStickersParams_Setters(t *testing.T) {
	g := (&GetCustomEmojiStickersParams{}).
		WithCustomEmojiIDs([]string{"CustomEmojiIDs"}...)

	assert.Equal(t, &GetCustomEmojiStickersParams{
		CustomEmojiIDs: []string{"CustomEmojiIDs"},
	}, g)
}

func TestUploadStickerFileParams_Setters(t *testing.T) {
	u := (&UploadStickerFileParams{}).
		WithSticker(testInputFile).
		WithStickerFormat("StickerFormat")

	assert.Equal(t, &UploadStickerFileParams{
		Sticker:       testInputFile,
		StickerFormat: "StickerFormat",
	}, u)
}

func TestCreateNewStickerSetParams_Setters(t *testing.T) {
	c := (&CreateNewStickerSetParams{}).
		WithName("Name").
		WithTitle("Title").
		WithStickers([]InputSticker{{}}...).
		WithStickerType("StickerType").
		WithNeedsRepainting()

	assert.Equal(t, &CreateNewStickerSetParams{
		Name:            "Name",
		Title:           "Title",
		Stickers:        []InputSticker{{}},
		StickerType:     "StickerType",
		NeedsRepainting: true,
	}, c)
}

func TestAddStickerToSetParams_Setters(t *testing.T) {
	a := (&AddStickerToSetParams{}).
		WithName("Name").
		WithSticker(InputSticker{Sticker: testInputFile})

	assert.Equal(t, &AddStickerToSetParams{
		Name:    "Name",
		Sticker: InputSticker{Sticker: testInputFile},
	}, a)
}

func TestSetStickerPositionInSetParams_Setters(t *testing.T) {
	s := (&SetStickerPositionInSetParams{}).
		WithSticker("Sticker").
		WithPosition(1)

	assert.Equal(t, &SetStickerPositionInSetParams{
		Sticker:  "Sticker",
		Position: 1,
	}, s)
}

func TestDeleteStickerFromSetParams_Setters(t *testing.T) {
	d := (&DeleteStickerFromSetParams{}).
		WithSticker("Sticker")

	assert.Equal(t, &DeleteStickerFromSetParams{
		Sticker: "Sticker",
	}, d)
}

func TestReplaceStickerInSetParams_Setters(t *testing.T) {
	r := (&ReplaceStickerInSetParams{}).
		WithName("Name").
		WithOldSticker("OldSticker").
		WithSticker(InputSticker{Sticker: testInputFile})

	assert.Equal(t, &ReplaceStickerInSetParams{
		Name:       "Name",
		OldSticker: "OldSticker",
		Sticker:    InputSticker{Sticker: testInputFile},
	}, r)
}

func TestSetStickerEmojiListParams_Setters(t *testing.T) {
	s := (&SetStickerEmojiListParams{}).
		WithSticker("Sticker").
		WithEmojiList([]string{"EmojiList"}...)

	assert.Equal(t, &SetStickerEmojiListParams{
		Sticker:   "Sticker",
		EmojiList: []string{"EmojiList"},
	}, s)
}

func TestSetStickerKeywordsParams_Setters(t *testing.T) {
	s := (&SetStickerKeywordsParams{}).
		WithSticker("Sticker").
		WithKeywords([]string{"Keywords"}...)

	assert.Equal(t, &SetStickerKeywordsParams{
		Sticker:  "Sticker",
		Keywords: []string{"Keywords"},
	}, s)
}

func TestSetStickerMaskPositionParams_Setters(t *testing.T) {
	s := (&SetStickerMaskPositionParams{}).
		WithSticker("Sticker").
		WithMaskPosition(&MaskPosition{Point: "MaskPosition"})

	assert.Equal(t, &SetStickerMaskPositionParams{
		Sticker:      "Sticker",
		MaskPosition: &MaskPosition{Point: "MaskPosition"},
	}, s)
}

func TestSetStickerSetTitleParams_Setters(t *testing.T) {
	s := (&SetStickerSetTitleParams{}).
		WithName("Name").
		WithTitle("Title")

	assert.Equal(t, &SetStickerSetTitleParams{
		Name:  "Name",
		Title: "Title",
	}, s)
}

func TestSetStickerSetThumbnailParams_Setters(t *testing.T) {
	s := (&SetStickerSetThumbnailParams{}).
		WithName("Name").
		WithThumbnail(&testInputFile).
		WithFormat("Format")

	assert.Equal(t, &SetStickerSetThumbnailParams{
		Name:      "Name",
		Thumbnail: &testInputFile,
		Format:    "Format",
	}, s)
}

func TestSetCustomEmojiStickerSetThumbnailParams_Setters(t *testing.T) {
	s := (&SetCustomEmojiStickerSetThumbnailParams{}).
		WithName("Name").
		WithCustomEmojiID("CustomEmojiID")

	assert.Equal(t, &SetCustomEmojiStickerSetThumbnailParams{
		Name:          "Name",
		CustomEmojiID: "CustomEmojiID",
	}, s)
}

func TestDeleteStickerSetParams_Setters(t *testing.T) {
	d := (&DeleteStickerSetParams{}).
		WithName("Name")

	assert.Equal(t, &DeleteStickerSetParams{
		Name: "Name",
	}, d)
}

func TestAnswerInlineQueryParams_Setters(t *testing.T) {
	a := (&AnswerInlineQueryParams{}).
		WithInlineQueryID("InlineQueryID").
		WithResults([]InlineQueryResult{&InlineQueryResultArticle{Type: "Results"}}...).
		WithCacheTime(1).
		WithIsPersonal().
		WithNextOffset("NextOffset").
		WithButton(&InlineQueryResultsButton{})

	assert.Equal(t, &AnswerInlineQueryParams{
		InlineQueryID: "InlineQueryID",
		Results:       []InlineQueryResult{&InlineQueryResultArticle{Type: "Results"}},
		CacheTime:     1,
		IsPersonal:    true,
		NextOffset:    "NextOffset",
		Button:        &InlineQueryResultsButton{},
	}, a)
}

func TestAnswerWebAppQueryParams_Setters(t *testing.T) {
	a := (&AnswerWebAppQueryParams{}).
		WithWebAppQueryID("WebAppQueryID").
		WithResult(&InlineQueryResultArticle{Type: "Result"})

	assert.Equal(t, &AnswerWebAppQueryParams{
		WebAppQueryID: "WebAppQueryID",
		Result:        &InlineQueryResultArticle{Type: "Result"},
	}, a)
}

func TestSendInvoiceParams_Setters(t *testing.T) {
	s := (&SendInvoiceParams{}).
		WithChatID(ChatID{ID: 1}).
		WithMessageThreadID(1).
		WithTitle("Title").
		WithDescription("Description").
		WithPayload("Payload").
		WithProviderToken("ProviderToken").
		WithCurrency("Currency").
		WithPrices([]LabeledPrice{{Label: "Prices"}}...).
		WithMaxTipAmount(2).
		WithSuggestedTipAmounts([]int{3}...).
		WithStartParameter("StartParameter").
		WithProviderData("ProviderData").
		WithPhotoURL("PhotoURL").
		WithPhotoSize(4).
		WithPhotoWidth(5).
		WithPhotoHeight(6).
		WithNeedName().
		WithNeedPhoneNumber().
		WithNeedEmail().
		WithNeedShippingAddress().
		WithSendPhoneNumberToProvider().
		WithSendEmailToProvider().
		WithIsFlexible().
		WithDisableNotification().
		WithProtectContent().
		WithReplyParameters(&ReplyParameters{MessageID: 7}).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}})

	assert.Equal(t, &SendInvoiceParams{
		ChatID:                    ChatID{ID: 1},
		MessageThreadID:           1,
		Title:                     "Title",
		Description:               "Description",
		Payload:                   "Payload",
		ProviderToken:             "ProviderToken",
		Currency:                  "Currency",
		Prices:                    []LabeledPrice{{Label: "Prices"}},
		MaxTipAmount:              2,
		SuggestedTipAmounts:       []int{3},
		StartParameter:            "StartParameter",
		ProviderData:              "ProviderData",
		PhotoURL:                  "PhotoURL",
		PhotoSize:                 4,
		PhotoWidth:                5,
		PhotoHeight:               6,
		NeedName:                  true,
		NeedPhoneNumber:           true,
		NeedEmail:                 true,
		NeedShippingAddress:       true,
		SendPhoneNumberToProvider: true,
		SendEmailToProvider:       true,
		IsFlexible:                true,
		DisableNotification:       true,
		ProtectContent:            true,
		ReplyParameters:           &ReplyParameters{MessageID: 7},
		ReplyMarkup:               &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
	}, s)
}

func TestCreateInvoiceLinkParams_Setters(t *testing.T) {
	c := (&CreateInvoiceLinkParams{}).
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

	assert.Equal(t, &CreateInvoiceLinkParams{
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
	}, c)
}

func TestAnswerShippingQueryParams_Setters(t *testing.T) {
	a := (&AnswerShippingQueryParams{}).
		WithShippingQueryID("ShippingQueryID").
		WithOk().
		WithShippingOptions([]ShippingOption{{ID: "ShippingOptions"}}...).
		WithErrorMessage("ErrorMessage")

	assert.Equal(t, &AnswerShippingQueryParams{
		ShippingQueryID: "ShippingQueryID",
		Ok:              true,
		ShippingOptions: []ShippingOption{{ID: "ShippingOptions"}},
		ErrorMessage:    "ErrorMessage",
	}, a)
}

func TestAnswerPreCheckoutQueryParams_Setters(t *testing.T) {
	a := (&AnswerPreCheckoutQueryParams{}).
		WithPreCheckoutQueryID("PreCheckoutQueryID").
		WithOk().
		WithErrorMessage("ErrorMessage")

	assert.Equal(t, &AnswerPreCheckoutQueryParams{
		PreCheckoutQueryID: "PreCheckoutQueryID",
		Ok:                 true,
		ErrorMessage:       "ErrorMessage",
	}, a)
}

func TestSetPassportDataErrorsParams_Setters(t *testing.T) {
	s := (&SetPassportDataErrorsParams{}).
		WithErrors([]PassportElementError{&PassportElementErrorDataField{}}...)

	assert.Equal(t, &SetPassportDataErrorsParams{
		Errors: []PassportElementError{&PassportElementErrorDataField{}},
	}, s)
}

func TestSendGameParams_Setters(t *testing.T) {
	s := (&SendGameParams{}).
		WithBusinessConnectionID("BusinessConnectionID").
		WithMessageThreadID(1).
		WithGameShortName("GameShortName").
		WithDisableNotification().
		WithProtectContent().
		WithReplyParameters(&ReplyParameters{MessageID: 2}).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}})

	assert.Equal(t, &SendGameParams{
		BusinessConnectionID: "BusinessConnectionID",
		MessageThreadID:      1,
		GameShortName:        "GameShortName",
		DisableNotification:  true,
		ProtectContent:       true,
		ReplyParameters:      &ReplyParameters{MessageID: 2},
		ReplyMarkup:          &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
	}, s)
}

func TestSetGameScoreParams_Setters(t *testing.T) {
	s := (&SetGameScoreParams{}).
		WithScore(3).
		WithForce().
		WithDisableEditMessage().
		WithMessageID(1).
		WithInlineMessageID("InlineMessageID")

	assert.Equal(t, &SetGameScoreParams{
		Score:              3,
		Force:              true,
		DisableEditMessage: true,
		MessageID:          1,
		InlineMessageID:    "InlineMessageID",
	}, s)
}

func TestGetGameHighScoresParams_Setters(t *testing.T) {
	g := (&GetGameHighScoresParams{}).
		WithMessageID(2).
		WithInlineMessageID("InlineMessageID")

	assert.Equal(t, &GetGameHighScoresParams{
		MessageID:       2,
		InlineMessageID: "InlineMessageID",
	}, g)
}
