package telego

import (
	"os"
	"testing"

	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"

	"github.com/mymmrac/telego/telegoapi"
)

//nolint:funlen
func TestTypesInterfaces(t *testing.T) {
	assert.Implements(t, (*ReplyMarkup)(nil), &ReplyKeyboardMarkup{})
	assert.Equal(t, MarkupTypeReplyKeyboard, (&ReplyKeyboardMarkup{}).ReplyType())

	assert.Implements(t, (*ReplyMarkup)(nil), &ReplyKeyboardRemove{})
	assert.Equal(t, MarkupTypeReplyKeyboardRemove, (&ReplyKeyboardRemove{}).ReplyType())

	assert.Implements(t, (*ReplyMarkup)(nil), &InlineKeyboardMarkup{})
	assert.Equal(t, MarkupTypeInlineKeyboard, (&InlineKeyboardMarkup{}).ReplyType())

	assert.Implements(t, (*ReplyMarkup)(nil), &ForceReply{})
	assert.Equal(t, MarkupTypeForceReply, (&ForceReply{}).ReplyType())

	assert.Implements(t, (*ChatMember)(nil), &ChatMemberOwner{})
	assert.Equal(t, MemberStatusCreator, (&ChatMemberOwner{}).MemberStatus())

	assert.Implements(t, (*ChatMember)(nil), &ChatMemberAdministrator{})
	assert.Equal(t, MemberStatusAdministrator, (&ChatMemberAdministrator{}).MemberStatus())

	assert.Implements(t, (*ChatMember)(nil), &ChatMemberMember{})
	assert.Equal(t, MemberStatusMember, (&ChatMemberMember{}).MemberStatus())

	assert.Implements(t, (*ChatMember)(nil), &ChatMemberRestricted{})
	assert.Equal(t, MemberStatusRestricted, (&ChatMemberRestricted{}).MemberStatus())

	assert.Implements(t, (*ChatMember)(nil), &ChatMemberLeft{})
	assert.Equal(t, MemberStatusLeft, (&ChatMemberLeft{}).MemberStatus())

	assert.Implements(t, (*ChatMember)(nil), &ChatMemberBanned{})
	assert.Equal(t, MemberStatusBanned, (&ChatMemberBanned{}).MemberStatus())

	assert.Implements(t, (*BotCommandScope)(nil), &BotCommandScopeDefault{})
	assert.Equal(t, ScopeTypeDefault, (&BotCommandScopeDefault{}).ScopeType())

	assert.Implements(t, (*BotCommandScope)(nil), &BotCommandScopeAllPrivateChats{})
	assert.Equal(t, ScopeTypeAllPrivateChats, (&BotCommandScopeAllPrivateChats{}).ScopeType())

	assert.Implements(t, (*BotCommandScope)(nil), &BotCommandScopeAllGroupChats{})
	assert.Equal(t, ScopeTypeAllGroupChats, (&BotCommandScopeAllGroupChats{}).ScopeType())

	assert.Implements(t, (*BotCommandScope)(nil), &BotCommandScopeAllChatAdministrators{})
	assert.Equal(t, ScopeTypeAllChatAdministrators, (&BotCommandScopeAllChatAdministrators{}).ScopeType())

	assert.Implements(t, (*BotCommandScope)(nil), &BotCommandScopeChat{})
	assert.Equal(t, ScopeTypeChat, (&BotCommandScopeChat{}).ScopeType())

	assert.Implements(t, (*BotCommandScope)(nil), &BotCommandScopeChatAdministrators{})
	assert.Equal(t, ScopeTypeChatAdministrators, (&BotCommandScopeChatAdministrators{}).ScopeType())

	assert.Implements(t, (*BotCommandScope)(nil), &BotCommandScopeChatMember{})
	assert.Equal(t, ScopeTypeChatMember, (&BotCommandScopeChatMember{}).ScopeType())

	assert.Implements(t, (*MenuButton)(nil), &MenuButtonCommands{})
	assert.Equal(t, ButtonTypeCommands, (&MenuButtonCommands{}).ButtonType())

	assert.Implements(t, (*MenuButton)(nil), &MenuButtonWebApp{})
	assert.Equal(t, ButtonTypeWebApp, (&MenuButtonWebApp{}).ButtonType())

	assert.Implements(t, (*MenuButton)(nil), &MenuButtonDefault{})
	assert.Equal(t, ButtonTypeDefault, (&MenuButtonDefault{}).ButtonType())

	assert.Implements(t, (*InputMedia)(nil), &InputMediaPhoto{})
	assert.Equal(t, MediaTypePhoto, (&InputMediaPhoto{}).MediaType())

	assert.Implements(t, (*InputMedia)(nil), &InputMediaVideo{})
	assert.Equal(t, MediaTypeVideo, (&InputMediaVideo{}).MediaType())

	assert.Implements(t, (*InputMedia)(nil), &InputMediaAnimation{})
	assert.Equal(t, MediaTypeAnimation, (&InputMediaAnimation{}).MediaType())

	assert.Implements(t, (*InputMedia)(nil), &InputMediaAudio{})
	assert.Equal(t, MediaTypeAudio, (&InputMediaAudio{}).MediaType())

	assert.Implements(t, (*InputMedia)(nil), &InputMediaDocument{})
	assert.Equal(t, MediaTypeDocument, (&InputMediaDocument{}).MediaType())

	assert.Implements(t, (*InlineQueryResult)(nil), &InlineQueryResultArticle{})
	assert.Equal(t, ResultTypeArticle, (&InlineQueryResultArticle{}).ResultType())

	assert.Implements(t, (*InlineQueryResult)(nil), &InlineQueryResultPhoto{})
	assert.Equal(t, ResultTypePhoto, (&InlineQueryResultPhoto{}).ResultType())

	assert.Implements(t, (*InlineQueryResult)(nil), &InlineQueryResultGif{})
	assert.Equal(t, ResultTypeGif, (&InlineQueryResultGif{}).ResultType())

	assert.Implements(t, (*InlineQueryResult)(nil), &InlineQueryResultMpeg4Gif{})
	assert.Equal(t, ResultTypeMpeg4Gif, (&InlineQueryResultMpeg4Gif{}).ResultType())

	assert.Implements(t, (*InlineQueryResult)(nil), &InlineQueryResultVideo{})
	assert.Equal(t, ResultTypeVideo, (&InlineQueryResultVideo{}).ResultType())

	assert.Implements(t, (*InlineQueryResult)(nil), &InlineQueryResultAudio{})
	assert.Equal(t, ResultTypeAudio, (&InlineQueryResultAudio{}).ResultType())

	assert.Implements(t, (*InlineQueryResult)(nil), &InlineQueryResultVoice{})
	assert.Equal(t, ResultTypeVoice, (&InlineQueryResultVoice{}).ResultType())

	assert.Implements(t, (*InlineQueryResult)(nil), &InlineQueryResultDocument{})
	assert.Equal(t, ResultTypeDocument, (&InlineQueryResultDocument{}).ResultType())

	assert.Implements(t, (*InlineQueryResult)(nil), &InlineQueryResultLocation{})
	assert.Equal(t, ResultTypeLocation, (&InlineQueryResultLocation{}).ResultType())

	assert.Implements(t, (*InlineQueryResult)(nil), &InlineQueryResultVenue{})
	assert.Equal(t, ResultTypeVenue, (&InlineQueryResultVenue{}).ResultType())

	assert.Implements(t, (*InlineQueryResult)(nil), &InlineQueryResultContact{})
	assert.Equal(t, ResultTypeContact, (&InlineQueryResultContact{}).ResultType())

	assert.Implements(t, (*InlineQueryResult)(nil), &InlineQueryResultGame{})
	assert.Equal(t, ResultTypeGame, (&InlineQueryResultGame{}).ResultType())

	assert.Implements(t, (*InlineQueryResult)(nil), &InlineQueryResultCachedPhoto{})
	assert.Equal(t, ResultTypePhoto, (&InlineQueryResultCachedPhoto{}).ResultType())

	assert.Implements(t, (*InlineQueryResult)(nil), &InlineQueryResultCachedGif{})
	assert.Equal(t, ResultTypeGif, (&InlineQueryResultCachedGif{}).ResultType())

	assert.Implements(t, (*InlineQueryResult)(nil), &InlineQueryResultCachedMpeg4Gif{})
	assert.Equal(t, ResultTypeMpeg4Gif, (&InlineQueryResultCachedMpeg4Gif{}).ResultType())

	assert.Implements(t, (*InlineQueryResult)(nil), &InlineQueryResultCachedSticker{})
	assert.Equal(t, ResultTypeSticker, (&InlineQueryResultCachedSticker{}).ResultType())

	assert.Implements(t, (*InlineQueryResult)(nil), &InlineQueryResultCachedDocument{})
	assert.Equal(t, ResultTypeDocument, (&InlineQueryResultCachedDocument{}).ResultType())

	assert.Implements(t, (*InlineQueryResult)(nil), &InlineQueryResultCachedVideo{})
	assert.Equal(t, ResultTypeVideo, (&InlineQueryResultCachedVideo{}).ResultType())

	assert.Implements(t, (*InlineQueryResult)(nil), &InlineQueryResultCachedVoice{})
	assert.Equal(t, ResultTypeVoice, (&InlineQueryResultCachedVoice{}).ResultType())

	assert.Implements(t, (*InlineQueryResult)(nil), &InlineQueryResultCachedAudio{})
	assert.Equal(t, ResultTypeAudio, (&InlineQueryResultCachedAudio{}).ResultType())

	assert.Implements(t, (*InputMessageContent)(nil), &InputTextMessageContent{})
	assert.Equal(t, ContentTypeText, (&InputTextMessageContent{}).ContentType())

	assert.Implements(t, (*InputMessageContent)(nil), &InputLocationMessageContent{})
	assert.Equal(t, ContentTypeLocation, (&InputLocationMessageContent{}).ContentType())

	assert.Implements(t, (*InputMessageContent)(nil), &InputVenueMessageContent{})
	assert.Equal(t, ContentTypeVenue, (&InputVenueMessageContent{}).ContentType())

	assert.Implements(t, (*InputMessageContent)(nil), &InputContactMessageContent{})
	assert.Equal(t, ContentTypeContact, (&InputContactMessageContent{}).ContentType())

	assert.Implements(t, (*InputMessageContent)(nil), &InputInvoiceMessageContent{})
	assert.Equal(t, ContentTypeInvoice, (&InputInvoiceMessageContent{}).ContentType())

	assert.Implements(t, (*PassportElementError)(nil), &PassportElementErrorDataField{})
	assert.Equal(t, ErrorSourceDataField, (&PassportElementErrorDataField{}).ErrorSource())

	assert.Implements(t, (*PassportElementError)(nil), &PassportElementErrorFrontSide{})
	assert.Equal(t, ErrorSourceFrontSide, (&PassportElementErrorFrontSide{}).ErrorSource())

	assert.Implements(t, (*PassportElementError)(nil), &PassportElementErrorReverseSide{})
	assert.Equal(t, ErrorSourceReverseSide, (&PassportElementErrorReverseSide{}).ErrorSource())

	assert.Implements(t, (*PassportElementError)(nil), &PassportElementErrorSelfie{})
	assert.Equal(t, ErrorSourceSelfie, (&PassportElementErrorSelfie{}).ErrorSource())

	assert.Implements(t, (*PassportElementError)(nil), &PassportElementErrorFile{})
	assert.Equal(t, ErrorSourceFile, (&PassportElementErrorFile{}).ErrorSource())

	assert.Implements(t, (*PassportElementError)(nil), &PassportElementErrorFiles{})
	assert.Equal(t, ErrorSourceFiles, (&PassportElementErrorFiles{}).ErrorSource())

	assert.Implements(t, (*PassportElementError)(nil), &PassportElementErrorTranslationFile{})
	assert.Equal(t, ErrorSourceTranslationFile, (&PassportElementErrorTranslationFile{}).ErrorSource())

	assert.Implements(t, (*PassportElementError)(nil), &PassportElementErrorTranslationFiles{})
	assert.Equal(t, ErrorSourceTranslationFiles, (&PassportElementErrorTranslationFiles{}).ErrorSource())

	assert.Implements(t, (*PassportElementError)(nil), &PassportElementErrorUnspecified{})
	assert.Equal(t, ErrorSourceUnspecified, (&PassportElementErrorUnspecified{}).ErrorSource())
}

func Test_chatMemberData_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		data    ChatMember
		isError bool
	}{
		{
			name: "success_creator",
			json: `{"status": "creator"}`,
			data: &ChatMemberOwner{
				Status: MemberStatusCreator,
			},
			isError: false,
		},
		{
			name: "success_administrator",
			json: `{"status": "administrator"}`,
			data: &ChatMemberAdministrator{
				Status: MemberStatusAdministrator,
			},
			isError: false,
		},
		{
			name: "success_member",
			json: `{"status": "member"}`,
			data: &ChatMemberMember{
				Status: MemberStatusMember,
			},
			isError: false,
		},
		{
			name: "success_restricted",
			json: `{"status": "restricted"}`,
			data: &ChatMemberRestricted{
				Status: MemberStatusRestricted,
			},
			isError: false,
		},
		{
			name: "success_left",
			json: `{"status": "left"}`,
			data: &ChatMemberLeft{
				Status: MemberStatusLeft,
			},
			isError: false,
		},
		{
			name: "success_kicked",
			json: `{"status": "kicked"}`,
			data: &ChatMemberBanned{
				Status: MemberStatusBanned,
			},
			isError: false,
		},
		{
			name:    "error_unknown_status",
			json:    `{"status": "test status"}`,
			data:    nil,
			isError: true,
		},
		{
			name:    "error_no_status",
			json:    "",
			data:    nil,
			isError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &chatMemberData{}
			err := c.UnmarshalJSON([]byte(tt.json))
			if tt.isError {
				assert.Error(t, err)
				assert.Nil(t, c.Data)
				return
			}
			assert.NoError(t, err)
			assert.EqualValues(t, tt.data, c.Data)
		})
	}
}

func TestChatMember_MemberUser(t *testing.T) {
	members := []ChatMember{
		&ChatMemberOwner{
			User: User{ID: 1},
		},
		&ChatMemberAdministrator{
			User: User{ID: 2},
		},
		&ChatMemberMember{
			User: User{ID: 3},
		},
		&ChatMemberRestricted{
			User: User{ID: 4},
		},
		&ChatMemberLeft{
			User: User{ID: 5},
		},
		&ChatMemberBanned{
			User: User{ID: 6},
		},
	}

	for i, cm := range members {
		assert.EqualValues(t, User{ID: int64(i) + 1}, cm.MemberUser())
	}
}

func TestChatMemberUpdated_UnmarshalJSON(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedCMU := &ChatMemberUpdated{
			Chat:          Chat{},
			From:          User{},
			Date:          0,
			OldChatMember: &ChatMemberOwner{Status: MemberStatusCreator},
			NewChatMember: &ChatMemberMember{Status: MemberStatusMember},
			InviteLink:    nil,
		}
		jsonData, err := json.Marshal(expectedCMU)
		assert.NoError(t, err)

		cmu := &ChatMemberUpdated{}
		err = cmu.UnmarshalJSON(jsonData)
		assert.NoError(t, err)
		assert.EqualValues(t, expectedCMU, cmu)
	})

	t.Run("error", func(t *testing.T) {
		cmu := &ChatMemberUpdated{}
		err := cmu.UnmarshalJSON([]byte("test"))
		assert.Error(t, err)
	})
}

func Test_menuButtonData_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		data    MenuButton
		isError bool
	}{
		{
			name: "success_commands",
			json: `{"type": "commands"}`,
			data: &MenuButtonCommands{
				Type: ButtonTypeCommands,
			},
			isError: false,
		},
		{
			name: "success_web_app",
			json: `{"type": "web_app"}`,
			data: &MenuButtonWebApp{
				Type: ButtonTypeWebApp,
			},
			isError: false,
		},
		{
			name: "success_default",
			json: `{"type": "default"}`,
			data: &MenuButtonDefault{
				Type: ButtonTypeDefault,
			},
			isError: false,
		},
		{
			name:    "error_unknown_type",
			json:    `{"type": "test type"}`,
			data:    nil,
			isError: true,
		},
		{
			name:    "error_no_type",
			json:    "",
			data:    nil,
			isError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &menuButtonData{}
			err := m.UnmarshalJSON([]byte(tt.json))
			if tt.isError {
				assert.Error(t, err)
				assert.Nil(t, m.Data)
				return
			}
			assert.NoError(t, err)
			assert.EqualValues(t, tt.data, m.Data)
		})
	}
}

func TestChatID_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		chatID   ChatID
		jsonData string
		isError  bool
	}{
		{
			name:     "empty",
			chatID:   ChatID{},
			jsonData: `""`,
			isError:  false,
		},
		{
			name: "success_id",
			chatID: ChatID{
				ID: 123,
			},
			jsonData: "123",
			isError:  false,
		},
		{
			name: "success_username",
			chatID: ChatID{
				Username: "test",
			},
			jsonData: `"test"`,
			isError:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.chatID.MarshalJSON()
			if tt.isError {
				assert.Error(t, err)
				assert.Nil(t, data)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.jsonData, string(data))
		})
	}
}

func TestInputFile_MarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		inputFile InputFile
		jsonData  string
		isError   bool
	}{
		{
			name: "success_file",
			inputFile: InputFile{
				File: &os.File{},
			},
			jsonData: `""`,
			isError:  false,
		},
		{
			name: "success_file_need_attach",
			inputFile: InputFile{
				File:       testNamedReade{},
				needAttach: true,
			},
			jsonData: `"` + attachFile + `test"`,
			isError:  false,
		},
		{
			name: "success_id",
			inputFile: InputFile{
				FileID: "fileID",
			},
			jsonData: `"fileID"`,
			isError:  false,
		},
		{
			name: "success_url",
			inputFile: InputFile{
				URL: "url",
			},
			jsonData: `"url"`,
			isError:  false,
		},
		{
			name:      "error",
			inputFile: InputFile{},
			jsonData:  "",
			isError:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.inputFile.MarshalJSON()
			if tt.isError {
				assert.Error(t, err)
				assert.Nil(t, data)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.jsonData, string(data))
		})
	}
}

func TestInputMedia_fileParameters(t *testing.T) {
	im := &InputMediaPhoto{
		Media: testInputFile,
	}

	assert.Equal(t, map[string]telegoapi.NamedReader{
		"media": testNamedReade{},
	}, im.fileParameters())
	assert.True(t, im.Media.needAttach)
}

func TestInputMediaVideo_fileParameters(t *testing.T) {
	im := &InputMediaVideo{
		Media: testInputFile,
		Thumb: &testInputFile,
	}

	assert.Equal(t, map[string]telegoapi.NamedReader{
		"media": testNamedReade{},
		"thumb": testNamedReade{},
	}, im.fileParameters())
	assert.True(t, im.Media.needAttach)
}

func TestInputMediaAnimation_fileParameters(t *testing.T) {
	im := &InputMediaAnimation{
		Media: testInputFile,
		Thumb: &testInputFile,
	}

	assert.Equal(t, map[string]telegoapi.NamedReader{
		"media": testNamedReade{},
		"thumb": testNamedReade{},
	}, im.fileParameters())
	assert.True(t, im.Media.needAttach)
}

func TestInputMediaAudio_fileParameters(t *testing.T) {
	im := &InputMediaAudio{
		Media: testInputFile,
		Thumb: &testInputFile,
	}

	assert.Equal(t, map[string]telegoapi.NamedReader{
		"media": testNamedReade{},
		"thumb": testNamedReade{},
	}, im.fileParameters())
	assert.True(t, im.Media.needAttach)
}

func TestInputMediaDocument_fileParameters(t *testing.T) {
	im := &InputMediaDocument{
		Media: testInputFile,
		Thumb: &testInputFile,
	}

	assert.Equal(t, map[string]telegoapi.NamedReader{
		"media": testNamedReade{},
		"thumb": testNamedReade{},
	}, im.fileParameters())
	assert.True(t, im.Media.needAttach)
}

func TestTypesConstants(t *testing.T) {
	tests := [][]string{
		{
			ChatTypeSender, ChatTypePrivate, ChatTypeGroup, ChatTypeSupergroup, ChatTypeChannel,
		},
		{
			EntityTypeMention, EntityTypeHashtag, EntityTypeCashtag, EntityTypeBotCommand, EntityTypeURL,
			EntityTypeEmail, EntityTypePhoneNumber, EntityTypeBold, EntityTypeItalic, EntityTypeUnderline,
			EntityTypeStrikethrough, EntityTypeSpoiler, EntityTypeCode, EntityTypePre, EntityTypeTextLink,
			EntityTypeTextMention,
		},
		{
			EmojiDice, EmojiDarts, EmojiBowling, EmojiBasketball, EmojiSoccer, EmojiSlotMachine,
		},
		{
			PollTypeRegular, PollTypeQuiz,
		},
		{
			MarkupTypeReplyKeyboard, MarkupTypeReplyKeyboardRemove, MarkupTypeInlineKeyboard, MarkupTypeForceReply,
		},
		{
			MemberStatusCreator, MemberStatusAdministrator, MemberStatusMember, MemberStatusRestricted,
			MemberStatusLeft, MemberStatusBanned,
		},
		{
			ScopeTypeDefault, ScopeTypeAllPrivateChats, ScopeTypeAllGroupChats, ScopeTypeAllChatAdministrators,
			ScopeTypeChat, ScopeTypeChatAdministrators, ScopeTypeChatMember,
		},
		{
			ButtonTypeCommands, ButtonTypeWebApp, ButtonTypeDefault,
		},
		{
			MediaTypePhoto, MediaTypeVideo, MediaTypeAnimation, MediaTypeAudio, MediaTypeDocument,
		},
		{
			PointForehead, PointEyes, PointMouth, PointChin,
		},
		{
			ResultTypeArticle, ResultTypePhoto, ResultTypeGif, ResultTypeMpeg4Gif, ResultTypeVideo, ResultTypeAudio,
			ResultTypeVoice, ResultTypeDocument, ResultTypeLocation, ResultTypeVenue, ResultTypeContact,
			ResultTypeGame, ResultTypeSticker,
		},
		{
			MimeTypeImageJpeg, MimeTypeImageGif, MimeTypeVideoMp4, MimeTypeTextHTML, MimeTypeApplicationPDF,
			MimeTypeApplicationZip,
		},
		{
			ContentTypeText, ContentTypeLocation, ContentTypeVenue, ContentTypeContact, ContentTypeInvoice,
		},
		{
			ElementTypePersonalDetails, ElementTypePassport, ElementTypeDriverLicense, ElementTypeIdentityCard,
			ElementTypeInternalPassport, ElementTypeAddress, ElementTypeUtilityBill, ElementTypeBankStatement,
			ElementTypeRentalAgreement, ElementTypePassportRegistration, ElementTypeTemporaryRegistration,
			ElementTypePhoneNumber, ElementTypeEmail,
		},
		{
			ErrorSourceDataField, ErrorSourceFrontSide, ErrorSourceReverseSide, ErrorSourceSelfie, ErrorSourceFile,
			ErrorSourceFiles, ErrorSourceTranslationFile, ErrorSourceTranslationFiles, ErrorSourceUnspecified,
		},
	}

	for _, tt := range tests {
		assert.True(t, len(tt) > 0)
		for _, ct := range tt {
			assert.True(t, len(ct) > 0)
		}
	}
}
