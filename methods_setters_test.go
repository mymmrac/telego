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
		WithChatID(ChatID{ID: 1}).
		WithText("Text").
		WithParseMode("ParseMode").
		WithEntities([]MessageEntity{{Type: "Entities"}}...).
		WithDisableWebPagePreview().
		WithDisableNotification().
		WithProtectContent().
		WithReplyToMessageID(1).
		WithAllowSendingWithoutReply().
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendMessageParams{
		ChatID:                   ChatID{ID: 1},
		Text:                     "Text",
		ParseMode:                "ParseMode",
		Entities:                 []MessageEntity{{Type: "Entities"}},
		DisableWebPagePreview:    true,
		DisableNotification:      true,
		ProtectContent:           true,
		ReplyToMessageID:         1,
		AllowSendingWithoutReply: true,
		ReplyMarkup:              &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestForwardMessageParams_Setters(t *testing.T) {
	f := (&ForwardMessageParams{}).
		WithChatID(ChatID{ID: 2}).
		WithFromChatID(ChatID{ID: 1}).
		WithDisableNotification().
		WithProtectContent().
		WithMessageID(2)

	assert.Equal(t, &ForwardMessageParams{
		ChatID:              ChatID{ID: 2},
		FromChatID:          ChatID{ID: 1},
		DisableNotification: true,
		ProtectContent:      true,
		MessageID:           2,
	}, f)
}

func TestCopyMessageParams_Setters(t *testing.T) {
	c := (&CopyMessageParams{}).
		WithChatID(ChatID{ID: 3}).
		WithFromChatID(ChatID{ID: 1}).
		WithMessageID(2).
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithDisableNotification().
		WithProtectContent().
		WithReplyToMessageID(3).
		WithAllowSendingWithoutReply().
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &CopyMessageParams{
		ChatID:                   ChatID{ID: 3},
		FromChatID:               ChatID{ID: 1},
		MessageID:                2,
		Caption:                  "Caption",
		ParseMode:                "ParseMode",
		CaptionEntities:          []MessageEntity{{Type: "CaptionEntities"}},
		DisableNotification:      true,
		ProtectContent:           true,
		ReplyToMessageID:         3,
		AllowSendingWithoutReply: true,
		ReplyMarkup:              &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, c)
}

func TestSendPhotoParams_Setters(t *testing.T) {
	s := (&SendPhotoParams{}).
		WithChatID(ChatID{ID: 4}).
		WithPhoto(testInputFile).
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithDisableNotification().
		WithProtectContent().
		WithReplyToMessageID(1).
		WithAllowSendingWithoutReply().
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendPhotoParams{
		ChatID:                   ChatID{ID: 4},
		Photo:                    testInputFile,
		Caption:                  "Caption",
		ParseMode:                "ParseMode",
		CaptionEntities:          []MessageEntity{{Type: "CaptionEntities"}},
		DisableNotification:      true,
		ProtectContent:           true,
		ReplyToMessageID:         1,
		AllowSendingWithoutReply: true,
		ReplyMarkup:              &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestSendAudioParams_Setters(t *testing.T) {
	s := (&SendAudioParams{}).
		WithChatID(ChatID{ID: 2}).
		WithAudio(testInputFile).
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithDuration(1).
		WithPerformer("Performer").
		WithTitle("Title").
		WithThumb(&testInputFile).
		WithDisableNotification().
		WithProtectContent().
		WithReplyToMessageID(2).
		WithAllowSendingWithoutReply().
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendAudioParams{
		ChatID:                   ChatID{ID: 2},
		Audio:                    testInputFile,
		Caption:                  "Caption",
		ParseMode:                "ParseMode",
		CaptionEntities:          []MessageEntity{{Type: "CaptionEntities"}},
		Duration:                 1,
		Performer:                "Performer",
		Title:                    "Title",
		Thumb:                    &testInputFile,
		DisableNotification:      true,
		ProtectContent:           true,
		ReplyToMessageID:         2,
		AllowSendingWithoutReply: true,
		ReplyMarkup:              &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestSendDocumentParams_Setters(t *testing.T) {
	s := (&SendDocumentParams{}).
		WithChatID(ChatID{ID: 3}).
		WithDocument(testInputFile).
		WithThumb(&testInputFile).
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithDisableContentTypeDetection().
		WithDisableNotification().
		WithProtectContent().
		WithReplyToMessageID(1).
		WithAllowSendingWithoutReply().
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendDocumentParams{
		ChatID:                      ChatID{ID: 3},
		Document:                    testInputFile,
		Thumb:                       &testInputFile,
		Caption:                     "Caption",
		ParseMode:                   "ParseMode",
		CaptionEntities:             []MessageEntity{{Type: "CaptionEntities"}},
		DisableContentTypeDetection: true,
		DisableNotification:         true,
		ProtectContent:              true,
		ReplyToMessageID:            1,
		AllowSendingWithoutReply:    true,
		ReplyMarkup:                 &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestSendVideoParams_Setters(t *testing.T) {
	s := (&SendVideoParams{}).
		WithChatID(ChatID{ID: 2}).
		WithVideo(testInputFile).
		WithDuration(1).
		WithWidth(2).
		WithHeight(3).
		WithThumb(&testInputFile).
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithSupportsStreaming().
		WithDisableNotification().
		WithProtectContent().
		WithReplyToMessageID(4).
		WithAllowSendingWithoutReply().
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendVideoParams{
		ChatID:                   ChatID{ID: 2},
		Video:                    testInputFile,
		Duration:                 1,
		Width:                    2,
		Height:                   3,
		Thumb:                    &testInputFile,
		Caption:                  "Caption",
		ParseMode:                "ParseMode",
		CaptionEntities:          []MessageEntity{{Type: "CaptionEntities"}},
		SupportsStreaming:        true,
		DisableNotification:      true,
		ProtectContent:           true,
		ReplyToMessageID:         4,
		AllowSendingWithoutReply: true,
		ReplyMarkup:              &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestSendAnimationParams_Setters(t *testing.T) {
	s := (&SendAnimationParams{}).
		WithChatID(ChatID{ID: 5}).
		WithAnimation(testInputFile).
		WithDuration(1).
		WithWidth(2).
		WithHeight(3).
		WithThumb(&testInputFile).
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithDisableNotification().
		WithProtectContent().
		WithReplyToMessageID(4).
		WithAllowSendingWithoutReply().
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendAnimationParams{
		ChatID:                   ChatID{ID: 5},
		Animation:                testInputFile,
		Duration:                 1,
		Width:                    2,
		Height:                   3,
		Thumb:                    &testInputFile,
		Caption:                  "Caption",
		ParseMode:                "ParseMode",
		CaptionEntities:          []MessageEntity{{Type: "CaptionEntities"}},
		DisableNotification:      true,
		ProtectContent:           true,
		ReplyToMessageID:         4,
		AllowSendingWithoutReply: true,
		ReplyMarkup:              &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestSendVoiceParams_Setters(t *testing.T) {
	s := (&SendVoiceParams{}).
		WithChatID(ChatID{ID: 5}).
		WithVoice(testInputFile).
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithDuration(1).
		WithDisableNotification().
		WithProtectContent().
		WithReplyToMessageID(2).
		WithAllowSendingWithoutReply().
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendVoiceParams{
		ChatID:                   ChatID{ID: 5},
		Voice:                    testInputFile,
		Caption:                  "Caption",
		ParseMode:                "ParseMode",
		CaptionEntities:          []MessageEntity{{Type: "CaptionEntities"}},
		Duration:                 1,
		DisableNotification:      true,
		ProtectContent:           true,
		ReplyToMessageID:         2,
		AllowSendingWithoutReply: true,
		ReplyMarkup:              &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestSendVideoNoteParams_Setters(t *testing.T) {
	s := (&SendVideoNoteParams{}).
		WithChatID(ChatID{ID: 3}).
		WithVideoNote(testInputFile).
		WithDuration(1).
		WithLength(2).
		WithThumb(&testInputFile).
		WithDisableNotification().
		WithProtectContent().
		WithReplyToMessageID(3).
		WithAllowSendingWithoutReply().
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendVideoNoteParams{
		ChatID:                   ChatID{ID: 3},
		VideoNote:                testInputFile,
		Duration:                 1,
		Length:                   2,
		Thumb:                    &testInputFile,
		DisableNotification:      true,
		ProtectContent:           true,
		ReplyToMessageID:         3,
		AllowSendingWithoutReply: true,
		ReplyMarkup:              &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestSendMediaGroupParams_Setters(t *testing.T) {
	s := (&SendMediaGroupParams{}).
		WithChatID(ChatID{ID: 4}).
		WithMedia([]InputMedia{&InputMediaAnimation{Type: "Media"}}...).
		WithDisableNotification().
		WithProtectContent().
		WithReplyToMessageID(1).
		WithAllowSendingWithoutReply()

	assert.Equal(t, &SendMediaGroupParams{
		ChatID:                   ChatID{ID: 4},
		Media:                    []InputMedia{&InputMediaAnimation{Type: "Media"}},
		DisableNotification:      true,
		ProtectContent:           true,
		ReplyToMessageID:         1,
		AllowSendingWithoutReply: true,
	}, s)
}

func TestSendLocationParams_Setters(t *testing.T) {
	s := (&SendLocationParams{}).
		WithChatID(ChatID{ID: 2}).
		WithLivePeriod(1).
		WithHeading(2).
		WithProximityAlertRadius(3).
		WithDisableNotification().
		WithProtectContent().
		WithReplyToMessageID(4).
		WithAllowSendingWithoutReply().
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendLocationParams{
		ChatID:                   ChatID{ID: 2},
		LivePeriod:               1,
		Heading:                  2,
		ProximityAlertRadius:     3,
		DisableNotification:      true,
		ProtectContent:           true,
		ReplyToMessageID:         4,
		AllowSendingWithoutReply: true,
		ReplyMarkup:              &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestEditMessageLiveLocationParams_Setters(t *testing.T) {
	e := (&EditMessageLiveLocationParams{}).
		WithChatID(ChatID{ID: 5}).
		WithMessageID(1).
		WithInlineMessageID("InlineMessageID").
		WithHeading(2).
		WithProximityAlertRadius(3).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}})

	assert.Equal(t, &EditMessageLiveLocationParams{
		ChatID:               ChatID{ID: 5},
		MessageID:            1,
		InlineMessageID:      "InlineMessageID",
		Heading:              2,
		ProximityAlertRadius: 3,
		ReplyMarkup:          &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
	}, e)
}

func TestStopMessageLiveLocationParams_Setters(t *testing.T) {
	s := (&StopMessageLiveLocationParams{}).
		WithChatID(ChatID{ID: 4}).
		WithMessageID(1).
		WithInlineMessageID("InlineMessageID").
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}})

	assert.Equal(t, &StopMessageLiveLocationParams{
		ChatID:          ChatID{ID: 4},
		MessageID:       1,
		InlineMessageID: "InlineMessageID",
		ReplyMarkup:     &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
	}, s)
}

func TestSendVenueParams_Setters(t *testing.T) {
	s := (&SendVenueParams{}).
		WithChatID(ChatID{ID: 2}).
		WithTitle("Title").
		WithAddress("Address").
		WithFoursquareID("FoursquareID").
		WithFoursquareType("FoursquareType").
		WithGooglePlaceID("GooglePlaceID").
		WithGooglePlaceType("GooglePlaceType").
		WithDisableNotification().
		WithProtectContent().
		WithReplyToMessageID(1).
		WithAllowSendingWithoutReply().
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendVenueParams{
		ChatID:                   ChatID{ID: 2},
		Title:                    "Title",
		Address:                  "Address",
		FoursquareID:             "FoursquareID",
		FoursquareType:           "FoursquareType",
		GooglePlaceID:            "GooglePlaceID",
		GooglePlaceType:          "GooglePlaceType",
		DisableNotification:      true,
		ProtectContent:           true,
		ReplyToMessageID:         1,
		AllowSendingWithoutReply: true,
		ReplyMarkup:              &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestSendContactParams_Setters(t *testing.T) {
	s := (&SendContactParams{}).
		WithChatID(ChatID{ID: 2}).
		WithPhoneNumber("PhoneNumber").
		WithFirstName("FirstName").
		WithLastName("LastName").
		WithVcard("Vcard").
		WithDisableNotification().
		WithProtectContent().
		WithReplyToMessageID(1).
		WithAllowSendingWithoutReply().
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendContactParams{
		ChatID:                   ChatID{ID: 2},
		PhoneNumber:              "PhoneNumber",
		FirstName:                "FirstName",
		LastName:                 "LastName",
		Vcard:                    "Vcard",
		DisableNotification:      true,
		ProtectContent:           true,
		ReplyToMessageID:         1,
		AllowSendingWithoutReply: true,
		ReplyMarkup:              &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestSendPollParams_Setters(t *testing.T) {
	s := (&SendPollParams{}).
		WithChatID(ChatID{ID: 2}).
		WithQuestion("Question").
		WithOptions([]string{"Options"}...).
		WithIsAnonymous().
		WithType("Type").
		WithAllowsMultipleAnswers().
		WithCorrectOptionID(1).
		WithExplanation("Explanation").
		WithExplanationParseMode("ExplanationParseMode").
		WithExplanationEntities([]MessageEntity{{Type: "ExplanationEntities"}}...).
		WithOpenPeriod(2).
		WithIsClosed().
		WithDisableNotification().
		WithProtectContent().
		WithReplyToMessageID(3).
		WithAllowSendingWithoutReply().
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendPollParams{
		ChatID:                   ChatID{ID: 2},
		Question:                 "Question",
		Options:                  []string{"Options"},
		IsAnonymous:              true,
		Type:                     "Type",
		AllowsMultipleAnswers:    true,
		CorrectOptionID:          1,
		Explanation:              "Explanation",
		ExplanationParseMode:     "ExplanationParseMode",
		ExplanationEntities:      []MessageEntity{{Type: "ExplanationEntities"}},
		OpenPeriod:               2,
		IsClosed:                 true,
		DisableNotification:      true,
		ProtectContent:           true,
		ReplyToMessageID:         3,
		AllowSendingWithoutReply: true,
		ReplyMarkup:              &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestSendDiceParams_Setters(t *testing.T) {
	s := (&SendDiceParams{}).
		WithChatID(ChatID{ID: 4}).
		WithEmoji("Emoji").
		WithDisableNotification().
		WithProtectContent().
		WithReplyToMessageID(1).
		WithAllowSendingWithoutReply().
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendDiceParams{
		ChatID:                   ChatID{ID: 4},
		Emoji:                    "Emoji",
		DisableNotification:      true,
		ProtectContent:           true,
		ReplyToMessageID:         1,
		AllowSendingWithoutReply: true,
		ReplyMarkup:              &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestSendChatActionParams_Setters(t *testing.T) {
	s := (&SendChatActionParams{}).
		WithChatID(ChatID{ID: 2}).
		WithAction("Action")

	assert.Equal(t, &SendChatActionParams{
		ChatID: ChatID{ID: 2},
		Action: "Action",
	}, s)
}

func TestGetUserProfilePhotosParams_Setters(t *testing.T) {
	g := (&GetUserProfilePhotosParams{}).
		WithOffset(1).
		WithLimit(1)

	assert.Equal(t, &GetUserProfilePhotosParams{
		Offset: 1,
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
		WithPermissions(ChatPermissions{CanSendMessages: true})

	assert.Equal(t, &RestrictChatMemberParams{
		ChatID:      ChatID{ID: 1},
		Permissions: ChatPermissions{CanSendMessages: true},
	}, r)
}

func TestPromoteChatMemberParams_Setters(t *testing.T) {
	p := (&PromoteChatMemberParams{}).
		WithChatID(ChatID{ID: 1}).
		WithIsAnonymous().
		WithCanManageChat().
		WithCanPostMessages().
		WithCanEditMessages().
		WithCanDeleteMessages().
		WithCanManageVideoChats().
		WithCanRestrictMembers().
		WithCanPromoteMembers().
		WithCanChangeInfo().
		WithCanInviteUsers().
		WithCanPinMessages()

	assert.Equal(t, &PromoteChatMemberParams{
		ChatID:              ChatID{ID: 1},
		IsAnonymous:         true,
		CanManageChat:       true,
		CanPostMessages:     true,
		CanEditMessages:     true,
		CanDeleteMessages:   true,
		CanManageVideoChats: true,
		CanRestrictMembers:  true,
		CanPromoteMembers:   true,
		CanChangeInfo:       true,
		CanInviteUsers:      true,
		CanPinMessages:      true,
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
		WithPermissions(ChatPermissions{CanSendMessages: true})

	assert.Equal(t, &SetChatPermissionsParams{
		ChatID:      ChatID{ID: 1},
		Permissions: ChatPermissions{CanSendMessages: true},
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
		WithDisableWebPagePreview().
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}})

	assert.Equal(t, &EditMessageTextParams{
		ChatID:                ChatID{ID: 1},
		MessageID:             1,
		InlineMessageID:       "InlineMessageID",
		Text:                  "Text",
		ParseMode:             "ParseMode",
		Entities:              []MessageEntity{{Type: "Entities"}},
		DisableWebPagePreview: true,
		ReplyMarkup:           &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
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

func TestSendStickerParams_Setters(t *testing.T) {
	s := (&SendStickerParams{}).
		WithChatID(ChatID{ID: 2}).
		WithSticker(testInputFile).
		WithDisableNotification().
		WithProtectContent().
		WithReplyToMessageID(1).
		WithAllowSendingWithoutReply().
		WithReplyMarkup(&ReplyKeyboardRemove{RemoveKeyboard: true})

	assert.Equal(t, &SendStickerParams{
		ChatID:                   ChatID{ID: 2},
		Sticker:                  testInputFile,
		DisableNotification:      true,
		ProtectContent:           true,
		ReplyToMessageID:         1,
		AllowSendingWithoutReply: true,
		ReplyMarkup:              &ReplyKeyboardRemove{RemoveKeyboard: true},
	}, s)
}

func TestGetStickerSetParams_Setters(t *testing.T) {
	g := (&GetStickerSetParams{}).
		WithName("Name")

	assert.Equal(t, &GetStickerSetParams{
		Name: "Name",
	}, g)
}

func TestUploadStickerFileParams_Setters(t *testing.T) {
	u := (&UploadStickerFileParams{}).
		WithPngSticker(testInputFile)

	assert.Equal(t, &UploadStickerFileParams{
		PngSticker: testInputFile,
	}, u)
}

func TestCreateNewStickerSetParams_Setters(t *testing.T) {
	c := (&CreateNewStickerSetParams{}).
		WithName("Name").
		WithTitle("Title").
		WithPngSticker(&testInputFile).
		WithTgsSticker(&testInputFile).
		WithWebmSticker(&testInputFile).
		WithEmojis("Emojis").
		WithContainsMasks().
		WithMaskPosition(&MaskPosition{Point: "MaskPosition"})

	assert.Equal(t, &CreateNewStickerSetParams{
		Name:          "Name",
		Title:         "Title",
		PngSticker:    &testInputFile,
		TgsSticker:    &testInputFile,
		WebmSticker:   &testInputFile,
		Emojis:        "Emojis",
		ContainsMasks: true,
		MaskPosition:  &MaskPosition{Point: "MaskPosition"},
	}, c)
}

func TestAddStickerToSetParams_Setters(t *testing.T) {
	a := (&AddStickerToSetParams{}).
		WithName("Name").
		WithPngSticker(&testInputFile).
		WithTgsSticker(&testInputFile).
		WithWebmSticker(&testInputFile).
		WithEmojis("Emojis").
		WithMaskPosition(&MaskPosition{Point: "MaskPosition"})

	assert.Equal(t, &AddStickerToSetParams{
		Name:         "Name",
		PngSticker:   &testInputFile,
		TgsSticker:   &testInputFile,
		WebmSticker:  &testInputFile,
		Emojis:       "Emojis",
		MaskPosition: &MaskPosition{Point: "MaskPosition"},
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

func TestSetStickerSetThumbParams_Setters(t *testing.T) {
	s := (&SetStickerSetThumbParams{}).
		WithName("Name").
		WithThumb(&testInputFile)

	assert.Equal(t, &SetStickerSetThumbParams{
		Name:  "Name",
		Thumb: &testInputFile,
	}, s)
}

func TestAnswerInlineQueryParams_Setters(t *testing.T) {
	a := (&AnswerInlineQueryParams{}).
		WithInlineQueryID("InlineQueryID").
		WithResults([]InlineQueryResult{&InlineQueryResultArticle{Type: "Results"}}...).
		WithCacheTime(1).
		WithIsPersonal().
		WithNextOffset("NextOffset").
		WithSwitchPmText("SwitchPmText").
		WithSwitchPmParameter("SwitchPmParameter")

	assert.Equal(t, &AnswerInlineQueryParams{
		InlineQueryID:     "InlineQueryID",
		Results:           []InlineQueryResult{&InlineQueryResultArticle{Type: "Results"}},
		CacheTime:         1,
		IsPersonal:        true,
		NextOffset:        "NextOffset",
		SwitchPmText:      "SwitchPmText",
		SwitchPmParameter: "SwitchPmParameter",
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
		WithTitle("Title").
		WithDescription("Description").
		WithPayload("Payload").
		WithProviderToken("ProviderToken").
		WithCurrency("Currency").
		WithPrices([]LabeledPrice{{Label: "Prices"}}...).
		WithMaxTipAmount(1).
		WithSuggestedTipAmounts([]int{2}...).
		WithStartParameter("StartParameter").
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
		WithIsFlexible().
		WithDisableNotification().
		WithProtectContent().
		WithReplyToMessageID(6).
		WithAllowSendingWithoutReply().
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}})

	assert.Equal(t, &SendInvoiceParams{
		ChatID:                    ChatID{ID: 1},
		Title:                     "Title",
		Description:               "Description",
		Payload:                   "Payload",
		ProviderToken:             "ProviderToken",
		Currency:                  "Currency",
		Prices:                    []LabeledPrice{{Label: "Prices"}},
		MaxTipAmount:              1,
		SuggestedTipAmounts:       []int{2},
		StartParameter:            "StartParameter",
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
		DisableNotification:       true,
		ProtectContent:            true,
		ReplyToMessageID:          6,
		AllowSendingWithoutReply:  true,
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
		WithGameShortName("GameShortName").
		WithDisableNotification().
		WithProtectContent().
		WithReplyToMessageID(1).
		WithAllowSendingWithoutReply().
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}})

	assert.Equal(t, &SendGameParams{
		GameShortName:            "GameShortName",
		DisableNotification:      true,
		ProtectContent:           true,
		ReplyToMessageID:         1,
		AllowSendingWithoutReply: true,
		ReplyMarkup:              &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
	}, s)
}

func TestSetGameScoreParams_Setters(t *testing.T) {
	s := (&SetGameScoreParams{}).
		WithScore(2).
		WithForce().
		WithDisableEditMessage().
		WithMessageID(1).
		WithInlineMessageID("InlineMessageID")

	assert.Equal(t, &SetGameScoreParams{
		Score:              2,
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
