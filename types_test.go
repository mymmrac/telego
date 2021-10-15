package telego

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mymmrac/telego/api"
)

// TODO: Check t.Parallel() with -parallel

func TestReplyKeyboardMarkup_ReplyType(t *testing.T) {
	assert.Equal(t, MarkupTypeReplyKeyboard, (&ReplyKeyboardMarkup{}).ReplyType())
}

func TestReplyKeyboardRemove_ReplyType(t *testing.T) {
	assert.Equal(t, MarkupTypeReplyKeyboardRemove, (&ReplyKeyboardRemove{}).ReplyType())
}

func TestInlineKeyboardMarkup_ReplyType(t *testing.T) {
	assert.Equal(t, MarkupTypeInlineKeyboard, (&InlineKeyboardMarkup{}).ReplyType())
}

func TestForceReply_ReplyType(t *testing.T) {
	assert.Equal(t, MarkupTypeForceReply, (&ForceReply{}).ReplyType())
}

func TestChatMemberOwner_MemberStatus(t *testing.T) {
	assert.Equal(t, MemberStatusCreator, (&ChatMemberOwner{}).MemberStatus())
}

func TestChatMemberAdministrator_MemberStatus(t *testing.T) {
	assert.Equal(t, MemberStatusAdministrator, (&ChatMemberAdministrator{}).MemberStatus())
}

func TestChatMemberMember_MemberStatus(t *testing.T) {
	assert.Equal(t, MemberStatusMember, (&ChatMemberMember{}).MemberStatus())
}

func TestChatMemberRestricted_MemberStatus(t *testing.T) {
	assert.Equal(t, MemberStatusRestricted, (&ChatMemberRestricted{}).MemberStatus())
}

func TestChatMemberLeft_MemberStatus(t *testing.T) {
	assert.Equal(t, MemberStatusLeft, (&ChatMemberLeft{}).MemberStatus())
}

func TestChatMemberBanned_MemberStatus(t *testing.T) {
	assert.Equal(t, MemberStatusKicked, (&ChatMemberBanned{}).MemberStatus())
}

func TestBotCommandScopeDefault_ScopeType(t *testing.T) {
	assert.Equal(t, ScopeTypeDefault, (&BotCommandScopeDefault{}).ScopeType())
}

func TestBotCommandScopeAllPrivateChats_ScopeType(t *testing.T) {
	assert.Equal(t, ScopeTypeAllPrivateChats, (&BotCommandScopeAllPrivateChats{}).ScopeType())
}

func TestBotCommandScopeAllGroupChats_ScopeType(t *testing.T) {
	assert.Equal(t, ScopeTypeAllGroupChats, (&BotCommandScopeAllGroupChats{}).ScopeType())
}

func TestBotCommandScopeAllChatAdministrators_ScopeType(t *testing.T) {
	assert.Equal(t, ScopeTypeAllChatAdministrators, (&BotCommandScopeAllChatAdministrators{}).ScopeType())
}

func TestBotCommandScopeChat_ScopeType(t *testing.T) {
	assert.Equal(t, ScopeTypeChat, (&BotCommandScopeChat{}).ScopeType())
}

func TestBotCommandScopeChatAdministrators_ScopeType(t *testing.T) {
	assert.Equal(t, ScopeTypeChatAdministrators, (&BotCommandScopeChatAdministrators{}).ScopeType())
}

func TestBotCommandScopeChatMember_ScopeType(t *testing.T) {
	assert.Equal(t, ScopeTypeChatMember, (&BotCommandScopeChatMember{}).ScopeType())
}

func TestInputMediaPhoto_MediaType(t *testing.T) {
	assert.Equal(t, MediaTypePhoto, (&InputMediaPhoto{}).MediaType())
}

func TestInputMediaVideo_MediaType(t *testing.T) {
	assert.Equal(t, MediaTypeVideo, (&InputMediaVideo{}).MediaType())
}

func TestInputMediaAnimation_MediaType(t *testing.T) {
	assert.Equal(t, MediaTypeAnimation, (&InputMediaAnimation{}).MediaType())
}

func TestInputMediaAudio_MediaType(t *testing.T) {
	assert.Equal(t, MediaTypeAudio, (&InputMediaAudio{}).MediaType())
}

func TestInputMediaDocument_MediaType(t *testing.T) {
	assert.Equal(t, MediaTypeDocument, (&InputMediaDocument{}).MediaType())
}

func TestInlineQueryResultArticle_ResultType(t *testing.T) {
	assert.Equal(t, ResultTypeArticle, (&InlineQueryResultArticle{}).ResultType())
}

func TestInlineQueryResultPhoto_ResultType(t *testing.T) {
	assert.Equal(t, ResultTypePhoto, (&InlineQueryResultPhoto{}).ResultType())
}

func TestInlineQueryResultGif_ResultType(t *testing.T) {
	assert.Equal(t, ResultTypeGif, (&InlineQueryResultGif{}).ResultType())
}

func TestInlineQueryResultMpeg4Gif_ResultType(t *testing.T) {
	assert.Equal(t, ResultTypeMpeg4Gif, (&InlineQueryResultMpeg4Gif{}).ResultType())
}

func TestInlineQueryResultVideo_ResultType(t *testing.T) {
	assert.Equal(t, ResultTypeVideo, (&InlineQueryResultVideo{}).ResultType())
}

func TestInlineQueryResultAudio_ResultType(t *testing.T) {
	assert.Equal(t, ResultTypeAudio, (&InlineQueryResultAudio{}).ResultType())
}

func TestInlineQueryResultVoice_ResultType(t *testing.T) {
	assert.Equal(t, ResultTypeVoice, (&InlineQueryResultVoice{}).ResultType())
}

func TestInlineQueryResultDocument_ResultType(t *testing.T) {
	assert.Equal(t, ResultTypeDocument, (&InlineQueryResultDocument{}).ResultType())
}

func TestInlineQueryResultLocation_ResultType(t *testing.T) {
	assert.Equal(t, ResultTypeLocation, (&InlineQueryResultLocation{}).ResultType())
}

func TestInlineQueryResultVenue_ResultType(t *testing.T) {
	assert.Equal(t, ResultTypeVenue, (&InlineQueryResultVenue{}).ResultType())
}

func TestInlineQueryResultContact_ResultType(t *testing.T) {
	assert.Equal(t, ResultTypeContact, (&InlineQueryResultContact{}).ResultType())
}

func TestInlineQueryResultGame_ResultType(t *testing.T) {
	assert.Equal(t, ResultTypeGame, (&InlineQueryResultGame{}).ResultType())
}

func TestInlineQueryResultCachedPhoto_ResultType(t *testing.T) {
	assert.Equal(t, ResultTypePhoto, (&InlineQueryResultCachedPhoto{}).ResultType())
}

func TestInlineQueryResultCachedGif_ResultType(t *testing.T) {
	assert.Equal(t, ResultTypeGif, (&InlineQueryResultCachedGif{}).ResultType())
}

func TestInlineQueryResultCachedMpeg4Gif_ResultType(t *testing.T) {
	assert.Equal(t, ResultTypeMpeg4Gif, (&InlineQueryResultCachedMpeg4Gif{}).ResultType())
}

func TestInlineQueryResultCachedSticker_ResultType(t *testing.T) {
	assert.Equal(t, ResultTypeSticker, (&InlineQueryResultCachedSticker{}).ResultType())
}

func TestInlineQueryResultCachedDocument_ResultType(t *testing.T) {
	assert.Equal(t, ResultTypeDocument, (&InlineQueryResultCachedDocument{}).ResultType())
}

func TestInlineQueryResultCachedVideo_ResultType(t *testing.T) {
	assert.Equal(t, ResultTypeVideo, (&InlineQueryResultCachedVideo{}).ResultType())
}

func TestInlineQueryResultCachedVoice_ResultType(t *testing.T) {
	assert.Equal(t, ResultTypeVoice, (&InlineQueryResultCachedVoice{}).ResultType())
}

func TestInlineQueryResultCachedAudio_ResultType(t *testing.T) {
	assert.Equal(t, ResultTypeAudio, (&InlineQueryResultCachedAudio{}).ResultType())
}

func TestInputTextMessageContent_ContentType(t *testing.T) {
	assert.Equal(t, ContentTypeText, (&InputTextMessageContent{}).ContentType())
}

func TestInputLocationMessageContent_ContentType(t *testing.T) {
	assert.Equal(t, ContentTypeLocation, (&InputLocationMessageContent{}).ContentType())
}

func TestInputVenueMessageContent_ContentType(t *testing.T) {
	assert.Equal(t, ContentTypeVenue, (&InputVenueMessageContent{}).ContentType())
}

func TestInputContactMessageContent_ContentType(t *testing.T) {
	assert.Equal(t, ContentTypeContact, (&InputContactMessageContent{}).ContentType())
}

func TestInputInvoiceMessageContent_ContentType(t *testing.T) {
	assert.Equal(t, ContentTypeInvoice, (&InputInvoiceMessageContent{}).ContentType())
}

func TestPassportElementErrorDataField_ErrorSource(t *testing.T) {
	assert.Equal(t, ErrorSourceDataField, (&PassportElementErrorDataField{}).ErrorSource())
}

func TestPassportElementErrorFrontSide_ErrorSource(t *testing.T) {
	assert.Equal(t, ErrorSourceFrontSide, (&PassportElementErrorFrontSide{}).ErrorSource())
}

func TestPassportElementErrorReverseSide_ErrorSource(t *testing.T) {
	assert.Equal(t, ErrorSourceReverseSide, (&PassportElementErrorReverseSide{}).ErrorSource())
}

func TestPassportElementErrorSelfie_ErrorSource(t *testing.T) {
	assert.Equal(t, ErrorSourceSelfie, (&PassportElementErrorSelfie{}).ErrorSource())
}

func TestPassportElementErrorFile_ErrorSource(t *testing.T) {
	assert.Equal(t, ErrorSourceFile, (&PassportElementErrorFile{}).ErrorSource())
}

func TestPassportElementErrorFiles_ErrorSource(t *testing.T) {
	assert.Equal(t, ErrorSourceFiles, (&PassportElementErrorFiles{}).ErrorSource())
}

func TestPassportElementErrorTranslationFile_ErrorSource(t *testing.T) {
	assert.Equal(t, ErrorSourceTranslationFile, (&PassportElementErrorTranslationFile{}).ErrorSource())
}

func TestPassportElementErrorTranslationFiles_ErrorSource(t *testing.T) {
	assert.Equal(t, ErrorSourceTranslationFiles, (&PassportElementErrorTranslationFiles{}).ErrorSource())
}

func TestPassportElementErrorUnspecified_ErrorSource(t *testing.T) {
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
				Status: MemberStatusKicked,
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

func TestChatID_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		chatID   ChatID
		jsonData string
		isError  bool
	}{
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
		{
			name:     "error",
			chatID:   ChatID{},
			jsonData: "",
			isError:  true,
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
			jsonData: ``,
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
		Media: InputFile{
			File: testNamedReade{},
		},
	}

	assert.Equal(t, map[string]api.NamedReader{
		"media": testNamedReade{},
	}, im.fileParameters())
	assert.True(t, im.Media.needAttach)
}

func TestInputMediaVideo_fileParameters(t *testing.T) {
	im := &InputMediaVideo{
		Media: InputFile{
			File: testNamedReade{},
		},
		Thumb: &InputFile{
			File: testNamedReade{},
		},
	}

	assert.Equal(t, map[string]api.NamedReader{
		"media": testNamedReade{},
		"thumb": testNamedReade{},
	}, im.fileParameters())
	assert.True(t, im.Media.needAttach)
}

func TestInputMediaAnimation_fileParameters(t *testing.T) {
	im := &InputMediaAnimation{
		Media: InputFile{
			File: testNamedReade{},
		},
		Thumb: &InputFile{
			File: testNamedReade{},
		},
	}

	assert.Equal(t, map[string]api.NamedReader{
		"media": testNamedReade{},
		"thumb": testNamedReade{},
	}, im.fileParameters())
	assert.True(t, im.Media.needAttach)
}

func TestInputMediaAudio_fileParameters(t *testing.T) {
	im := &InputMediaAudio{
		Media: InputFile{
			File: testNamedReade{},
		},
		Thumb: &InputFile{
			File: testNamedReade{},
		},
	}

	assert.Equal(t, map[string]api.NamedReader{
		"media": testNamedReade{},
		"thumb": testNamedReade{},
	}, im.fileParameters())
	assert.True(t, im.Media.needAttach)
}

func TestInputMediaDocument_fileParameters(t *testing.T) {
	im := &InputMediaDocument{
		Media: InputFile{
			File: testNamedReade{},
		},
		Thumb: &InputFile{
			File: testNamedReade{},
		},
	}

	assert.Equal(t, map[string]api.NamedReader{
		"media": testNamedReade{},
		"thumb": testNamedReade{},
	}, im.fileParameters())
	assert.True(t, im.Media.needAttach)
}
