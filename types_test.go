package telego

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego/internal/json"
	ta "github.com/mymmrac/telego/telegoapi"
)

func TestTypesInterfaces(t *testing.T) {
	assert.Implements(t, (*MessageOrigin)(nil), &MessageOriginUser{})
	assert.Equal(t, OriginTypeUser, (&MessageOriginUser{}).OriginType())

	assert.Implements(t, (*MessageOrigin)(nil), &MessageOriginHiddenUser{})
	assert.Equal(t, OriginTypeHiddenUser, (&MessageOriginHiddenUser{}).OriginType())

	assert.Implements(t, (*MessageOrigin)(nil), &MessageOriginChat{})
	assert.Equal(t, OriginTypeChat, (&MessageOriginChat{}).OriginType())

	assert.Implements(t, (*MessageOrigin)(nil), &MessageOriginChannel{})
	assert.Equal(t, OriginTypeChannel, (&MessageOriginChannel{}).OriginType())

	assert.Implements(t, (*PaidMedia)(nil), &PaidMediaPreview{})
	assert.Equal(t, PaidMediaTypePreview, (&PaidMediaPreview{}).MediaType())

	assert.Implements(t, (*PaidMedia)(nil), &PaidMediaPhoto{})
	assert.Equal(t, PaidMediaTypePhoto, (&PaidMediaPhoto{}).MediaType())

	assert.Implements(t, (*PaidMedia)(nil), &PaidMediaVideo{})
	assert.Equal(t, PaidMediaTypeVideo, (&PaidMediaVideo{}).MediaType())

	assert.Implements(t, (*BackgroundFill)(nil), &BackgroundFillSolid{})
	assert.Equal(t, BackgroundFilledSolid, (&BackgroundFillSolid{}).BackgroundFilled())

	assert.Implements(t, (*BackgroundFill)(nil), &BackgroundFillGradient{})
	assert.Equal(t, BackgroundFilledGradient, (&BackgroundFillGradient{}).BackgroundFilled())

	assert.Implements(t, (*BackgroundFill)(nil), &BackgroundFillFreeformGradient{})
	assert.Equal(t, BackgroundFilledFreeformGradient, (&BackgroundFillFreeformGradient{}).BackgroundFilled())

	assert.Implements(t, (*BackgroundType)(nil), &BackgroundTypeFill{})
	assert.Equal(t, BackgroundTypeNameFill, (&BackgroundTypeFill{}).BackgroundType())

	assert.Implements(t, (*BackgroundType)(nil), &BackgroundTypeWallpaper{})
	assert.Equal(t, BackgroundTypeNameWallpaper, (&BackgroundTypeWallpaper{}).BackgroundType())

	assert.Implements(t, (*BackgroundType)(nil), &BackgroundTypePattern{})
	assert.Equal(t, BackgroundTypeNamePattern, (&BackgroundTypePattern{}).BackgroundType())

	assert.Implements(t, (*BackgroundType)(nil), &BackgroundTypeChatTheme{})
	assert.Equal(t, BackgroundTypeNameChatTheme, (&BackgroundTypeChatTheme{}).BackgroundType())

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

	assert.Implements(t, (*ReactionType)(nil), &ReactionTypeEmoji{})
	assert.Equal(t, ReactionEmoji, (&ReactionTypeEmoji{}).ReactionType())

	assert.Implements(t, (*ReactionType)(nil), &ReactionTypeCustomEmoji{})
	assert.Equal(t, ReactionCustomEmoji, (&ReactionTypeCustomEmoji{}).ReactionType())

	assert.Implements(t, (*ReactionType)(nil), &ReactionTypePaid{})
	assert.Equal(t, ReactionPaid, (&ReactionTypePaid{}).ReactionType())

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

	assert.Implements(t, (*ChatBoostSource)(nil), &ChatBoostSourcePremium{})
	assert.Equal(t, BoostSourcePremium, (&ChatBoostSourcePremium{}).BoostSource())

	assert.Implements(t, (*ChatBoostSource)(nil), &ChatBoostSourceGiftCode{})
	assert.Equal(t, BoostSourceGiftCode, (&ChatBoostSourceGiftCode{}).BoostSource())

	assert.Implements(t, (*ChatBoostSource)(nil), &ChatBoostSourceGiveaway{})
	assert.Equal(t, BoostSourceGiveaway, (&ChatBoostSourceGiveaway{}).BoostSource())

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

	assert.Implements(t, (*InputPaidMedia)(nil), &InputPaidMediaPhoto{})
	assert.Equal(t, PaidMediaTypePhoto, (&InputPaidMediaPhoto{}).MediaType())

	assert.Implements(t, (*InputPaidMedia)(nil), &InputPaidMediaVideo{})
	assert.Equal(t, PaidMediaTypeVideo, (&InputPaidMediaVideo{}).MediaType())

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

	assert.Implements(t, (*RevenueWithdrawalState)(nil), &RevenueWithdrawalStatePending{})
	assert.Equal(t, WithdrawalStatePending, (&RevenueWithdrawalStatePending{}).WithdrawalState())

	assert.Implements(t, (*RevenueWithdrawalState)(nil), &RevenueWithdrawalStateSucceeded{})
	assert.Equal(t, WithdrawalStateSucceeded, (&RevenueWithdrawalStateSucceeded{}).WithdrawalState())

	assert.Implements(t, (*RevenueWithdrawalState)(nil), &RevenueWithdrawalStateFailed{})
	assert.Equal(t, WithdrawalStateFailed, (&RevenueWithdrawalStateFailed{}).WithdrawalState())

	assert.Implements(t, (*TransactionPartner)(nil), &TransactionPartnerUser{})
	assert.Equal(t, PartnerTypeUser, (&TransactionPartnerUser{}).PartnerType())

	assert.Implements(t, (*TransactionPartner)(nil), &TransactionPartnerFragment{})
	assert.Equal(t, PartnerTypeFragment, (&TransactionPartnerFragment{}).PartnerType())

	assert.Implements(t, (*TransactionPartner)(nil), &TransactionPartnerTelegramAds{})
	assert.Equal(t, PartnerTypeTelegramAds, (&TransactionPartnerTelegramAds{}).PartnerType())

	assert.Implements(t, (*TransactionPartner)(nil), &TransactionPartnerOther{})
	assert.Equal(t, PartnerTypeOther, (&TransactionPartnerOther{}).PartnerType())

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
				require.Error(t, err)
				assert.Nil(t, c.Data)
				return
			}
			require.NoError(t, err)
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
		require.NoError(t, err)

		cmu := &ChatMemberUpdated{}
		err = cmu.UnmarshalJSON(jsonData)
		require.NoError(t, err)
		assert.EqualValues(t, expectedCMU, cmu)
	})

	t.Run("error", func(t *testing.T) {
		cmu := &ChatMemberUpdated{}
		err := cmu.UnmarshalJSON([]byte("test"))
		require.Error(t, err)
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
				require.Error(t, err)
				assert.Nil(t, m.Data)
				return
			}
			require.NoError(t, err)
			assert.EqualValues(t, tt.data, m.Data)
		})
	}
}

func TestChat_ChatID(t *testing.T) {
	chat := Chat{ID: 1}
	chatID := chat.ChatID()
	assert.Empty(t, chatID.Username)
	assert.Equal(t, chat.ID, chatID.ID)
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
				require.Error(t, err)
				assert.Nil(t, data)
				return
			}
			require.NoError(t, err)
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
				require.Error(t, err)
				assert.Nil(t, data)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.jsonData, string(data))
		})
	}
}

func TestInputMedia_fileParameters(t *testing.T) {
	im := &InputMediaPhoto{
		Media: testInputFile,
	}

	assert.Equal(t, map[string]ta.NamedReader{
		"media": testNamedReade{},
	}, im.fileParameters())
	assert.True(t, im.Media.needAttach)
}

func TestInputMediaVideo_fileParameters(t *testing.T) {
	im := &InputMediaVideo{
		Media:     testInputFile,
		Thumbnail: &testInputFile,
	}

	assert.Equal(t, map[string]ta.NamedReader{
		"media":     testNamedReade{},
		"thumbnail": testNamedReade{},
	}, im.fileParameters())
	assert.True(t, im.Media.needAttach)
}

func TestInputMediaAnimation_fileParameters(t *testing.T) {
	im := &InputMediaAnimation{
		Media:     testInputFile,
		Thumbnail: &testInputFile,
	}

	assert.Equal(t, map[string]ta.NamedReader{
		"media":     testNamedReade{},
		"thumbnail": testNamedReade{},
	}, im.fileParameters())
	assert.True(t, im.Media.needAttach)
}

func TestInputMediaAudio_fileParameters(t *testing.T) {
	im := &InputMediaAudio{
		Media:     testInputFile,
		Thumbnail: &testInputFile,
	}

	assert.Equal(t, map[string]ta.NamedReader{
		"media":     testNamedReade{},
		"thumbnail": testNamedReade{},
	}, im.fileParameters())
	assert.True(t, im.Media.needAttach)
}

func TestInputMediaDocument_fileParameters(t *testing.T) {
	im := &InputMediaDocument{
		Media:     testInputFile,
		Thumbnail: &testInputFile,
	}

	assert.Equal(t, map[string]ta.NamedReader{
		"media":     testNamedReade{},
		"thumbnail": testNamedReade{},
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
			EntityTypeStrikethrough, EntityTypeSpoiler, EntityTypeBlockquote, EntityTypeExpandableBlockquote,
			EntityTypeCode, EntityTypePre, EntityTypeTextLink, EntityTypeTextMention, EntityTypeCustomEmoji,
		},
		{
			OriginTypeUser, OriginTypeHiddenUser, OriginTypeChat, OriginTypeChannel,
		},
		{
			EmojiDice, EmojiDarts, EmojiBowling, EmojiBasketball, EmojiSoccer, EmojiSlotMachine,
		},
		{
			PollTypeRegular, PollTypeQuiz,
		},
		{
			BackgroundFilledSolid, BackgroundFilledGradient, BackgroundFilledFreeformGradient,
		},
		{
			BackgroundTypeNameFill, BackgroundTypeNameWallpaper, BackgroundTypeNamePattern, BackgroundTypeNameChatTheme,
		},
		{
			MarkupTypeReplyKeyboard, MarkupTypeReplyKeyboardRemove, MarkupTypeInlineKeyboard, MarkupTypeForceReply,
		},
		{
			MemberStatusCreator, MemberStatusAdministrator, MemberStatusMember, MemberStatusRestricted,
			MemberStatusLeft, MemberStatusBanned,
		},
		{
			ReactionEmoji, ReactionCustomEmoji, ReactionPaid,
		},
		{
			ScopeTypeDefault, ScopeTypeAllPrivateChats, ScopeTypeAllGroupChats, ScopeTypeAllChatAdministrators,
			ScopeTypeChat, ScopeTypeChatAdministrators, ScopeTypeChatMember,
		},
		{
			ButtonTypeCommands, ButtonTypeWebApp, ButtonTypeDefault,
		},
		{
			BoostSourcePremium, BoostSourceGiftCode, BoostSourceGiveaway,
		},
		{
			MediaTypePhoto, MediaTypeVideo, MediaTypeAnimation, MediaTypeAudio, MediaTypeDocument,
		},
		{
			PaidMediaTypePreview, PaidMediaTypePhoto, PaidMediaTypeVideo,
		},
		{
			StickerTypeRegular, StickerTypeMask, StickerTypeCustomEmoji,
		},
		{
			PointForehead, PointEyes, PointMouth, PointChin,
		},
		{
			StickerStatic, StickerAnimated, StickerVideo,
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
			WithdrawalStatePending, WithdrawalStateSucceeded, WithdrawalStateFailed,
		},
		{
			PartnerTypeUser, PartnerTypeFragment, PartnerTypeTelegramAds, PartnerTypeOther,
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
		assert.NotEmpty(t, tt)
		for _, ct := range tt {
			assert.NotEmpty(t, ct)
		}
	}
}

func TestUpdate_Clone(t *testing.T) {
	u := Update{
		UpdateID: 1,
		Message: &Message{
			Text: "ok",
			Chat: Chat{
				ID: 1,
			},
			Contact: &Contact{
				PhoneNumber: "123",
			},
			ForwardOrigin: &MessageOriginUser{
				Type: OriginTypeUser,
				Date: 123,
				SenderUser: User{
					ID: 1,
				},
			},
			PinnedMessage: &InaccessibleMessage{
				Chat: Chat{
					ID: 1,
				},
				MessageID: 1,
				Date:      0,
			},
		},
	}

	assert.NotPanics(t, func() {
		uc := u.Clone()
		assert.Equal(t, u, uc)
	})

	assert.Panics(t, func() {
		_ = (Update{ChatMember: &ChatMemberUpdated{}}).Clone()
	})
}

func BenchmarkUpdate_Clone(b *testing.B) {
	const n1 = 1
	const s1 = "text"
	const b1 = true

	c1 := Chat{
		ID:        n1,
		Type:      s1,
		Title:     s1,
		Username:  s1,
		FirstName: s1,
		LastName:  s1,
	}

	u1 := User{
		ID:                      n1,
		IsBot:                   b1,
		FirstName:               s1,
		LastName:                s1,
		Username:                s1,
		LanguageCode:            s1,
		IsPremium:               b1,
		AddedToAttachmentMenu:   b1,
		CanJoinGroups:           b1,
		CanReadAllGroupMessages: b1,
		SupportsInlineQueries:   b1,
	}

	u := Update{
		UpdateID: n1,
		Message: &Message{
			MessageID:  n1,
			From:       &u1,
			SenderChat: &c1,
			Date:       n1,
			Chat:       c1,
			ForwardOrigin: &MessageOriginChat{
				Type:            OriginTypeChat,
				Date:            n1,
				SenderChat:      c1,
				AuthorSignature: s1,
			},
			IsAutomaticForward:    b1,
			ViaBot:                &u1,
			EditDate:              n1,
			HasProtectedContent:   b1,
			MediaGroupID:          s1,
			AuthorSignature:       s1,
			Text:                  s1,
			Caption:               s1,
			NewChatTitle:          s1,
			DeleteChatPhoto:       b1,
			GroupChatCreated:      b1,
			SupergroupChatCreated: b1,
			ChannelChatCreated:    b1,
			MigrateToChatID:       n1,
			MigrateFromChatID:     n1,
			ConnectedWebsite:      s1,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = u.Clone()
	}
}

func TestUpdate_CloneSafe(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		u := Update{
			UpdateID: 1,
			Message: &Message{
				Text: "ok",
				Contact: &Contact{
					PhoneNumber: "123",
				},
			},
		}

		uc, err := u.CloneSafe()
		require.NoError(t, err)
		assert.Equal(t, u, uc)
	})

	t.Run("error_unmarshal", func(t *testing.T) {
		uc, err := (Update{ChatMember: &ChatMemberUpdated{}}).CloneSafe()
		require.Error(t, err)
		assert.Zero(t, uc)
	})

	t.Run("error_marshal", func(t *testing.T) {
		u := Update{
			MyChatMember: &ChatMemberUpdated{
				OldChatMember: badChatMember{},
			},
		}
		uc, err := u.CloneSafe()
		require.Error(t, err)
		assert.Zero(t, uc)
	})
}

type badChatMember struct{}

func (b badChatMember) MarshalJSON() ([]byte, error) {
	return nil, errTest
}

func (b badChatMember) MemberStatus() string {
	panic("implement me")
}

func (b badChatMember) MemberUser() User {
	panic("implement me")
}

func (b badChatMember) MemberIsMember() bool {
	panic("implement me")
}

func (b badChatMember) iChatMember() {}

func TestChatID_String(t *testing.T) {
	tests := []struct {
		name        string
		chatID      ChatID
		stringValue string
	}{
		{
			name:        "empty",
			chatID:      ChatID{},
			stringValue: "",
		},
		{
			name: "id",
			chatID: ChatID{
				ID: 123,
			},
			stringValue: "123",
		},
		{
			name: "username",
			chatID: ChatID{
				Username: "test",
			},
			stringValue: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.stringValue, tt.chatID.String())
		})
	}
}

func TestInputFile_String(t *testing.T) {
	tests := []struct {
		name        string
		inputFile   InputFile
		stringValue string
	}{
		{
			name:        "empty",
			inputFile:   InputFile{},
			stringValue: "",
		},
		{
			name: "file",
			inputFile: InputFile{
				File: &testNamedReade{},
			},
			stringValue: "test",
		},
		{
			name: "id",
			inputFile: InputFile{
				FileID: "fileID",
			},
			stringValue: "fileID",
		},
		{
			name: "url",
			inputFile: InputFile{
				URL: "url",
			},
			stringValue: "url",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.stringValue, tt.inputFile.String())
		})
	}
}

func TestUpdate_Context(t *testing.T) {
	u := Update{
		UpdateID: 1,
	}

	assert.NotNil(t, u.Context())

	ctx := context.Background()
	cu := u.WithContext(ctx)
	assert.Equal(t, ctx, cu.Context())
	assert.Equal(t, u.UpdateID, cu.UpdateID)

	assert.Panics(t, func() {
		u.WithContext(nil) //nolint:staticcheck
	})
}

func Test_ChatFullInfo_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		data    *ChatFullInfo
		isError bool
	}{
		{
			name: "success",
			json: `{"id": 1}`,
			data: &ChatFullInfo{
				ID: 1,
			},
			isError: false,
		},
		{
			name:    "error_invalid",
			json:    "",
			data:    nil,
			isError: true,
		},
		{
			name:    "error_invalid_available_reactions",
			json:    `{"available_reactions": [{}]}`,
			data:    nil,
			isError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatFullInfo{}
			err := c.UnmarshalJSON([]byte(tt.json))
			if tt.isError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.EqualValues(t, tt.data, c)
		})
	}
}

func Test_ExternalReplyInfo_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		data    *ExternalReplyInfo
		isError bool
	}{
		{
			name: "success",
			json: `{"message_id": 1, "origin": {"type": "user"}}`,
			data: &ExternalReplyInfo{
				MessageID: 1,
				Origin: &MessageOriginUser{
					Type: OriginTypeUser,
				},
			},
			isError: false,
		},
		{
			name:    "error_invalid",
			json:    "",
			data:    nil,
			isError: true,
		},
		{
			name:    "error_no_origin",
			json:    `{}`,
			data:    nil,
			isError: true,
		},
		{
			name:    "error_invalid_origin",
			json:    `{"origin": {}}`,
			data:    nil,
			isError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ExternalReplyInfo{}
			err := e.UnmarshalJSON([]byte(tt.json))
			if tt.isError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.EqualValues(t, tt.data, e)
		})
	}
}

func Test_CallbackQuery_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		data    *CallbackQuery
		isError bool
	}{
		{
			name: "success",
			json: `{"id": "1"}`,
			data: &CallbackQuery{
				ID: "1",
			},
			isError: false,
		},
		{
			name:    "error_invalid",
			json:    "",
			data:    nil,
			isError: true,
		},
		{
			name:    "error_invalid_message",
			json:    `{"message": {"date": "a"}}`,
			data:    nil,
			isError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CallbackQuery{}
			err := c.UnmarshalJSON([]byte(tt.json))
			if tt.isError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.EqualValues(t, tt.data, c)
		})
	}
}

func Test_ReactionCount_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		data    *ReactionCount
		isError bool
	}{
		{
			name: "success",
			json: `{"total_count": 1, "type": {"type": "emoji"}}`,
			data: &ReactionCount{
				TotalCount: 1,
				Type: &ReactionTypeEmoji{
					Type: ReactionEmoji,
				},
			},
			isError: false,
		},
		{
			name:    "error_invalid",
			json:    "",
			data:    nil,
			isError: true,
		},
		{
			name:    "error_no_type",
			json:    `{}`,
			data:    nil,
			isError: true,
		},
		{
			name:    "error_invalid_type",
			json:    `{"type": {}}`,
			data:    nil,
			isError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ReactionCount{}
			err := c.UnmarshalJSON([]byte(tt.json))
			if tt.isError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.EqualValues(t, tt.data, c)
		})
	}
}

func Test_MessageReactionUpdated_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		data    *MessageReactionUpdated
		isError bool
	}{
		{
			name: "success",
			json: `{"old_reaction": [], "new_reaction": []}`,
			data: &MessageReactionUpdated{
				OldReaction: make([]ReactionType, 0),
				NewReaction: make([]ReactionType, 0),
			},
			isError: false,
		},
		{
			name:    "error_invalid",
			json:    "",
			data:    nil,
			isError: true,
		},
		{
			name:    "error_no_old_reaction",
			json:    `{"new_reaction": []}`,
			data:    nil,
			isError: true,
		},
		{
			name:    "error_no_new_reaction",
			json:    `{"old_reaction": []}`,
			data:    nil,
			isError: true,
		},
		{
			name:    "error_invalid_old_reaction",
			json:    `{"old_reaction": [{"type": 1}], "new_reaction": []}`,
			data:    nil,
			isError: true,
		},
		{
			name:    "error_invalid_new_reaction",
			json:    `{"old_reaction": [],  "new_reaction": [{"type": 1}]}`,
			data:    nil,
			isError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MessageReactionUpdated{}
			err := m.UnmarshalJSON([]byte(tt.json))
			if tt.isError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.EqualValues(t, tt.data, m)
		})
	}
}

func Test_ChatBoost_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		data    *ChatBoost
		isError bool
	}{
		{
			name: "success",
			json: `{"boost_id": "1", "source": {"source": "premium"}}`,
			data: &ChatBoost{
				BoostID: "1",
				Source: &ChatBoostSourcePremium{
					Source: BoostSourcePremium,
				},
			},
			isError: false,
		},
		{
			name:    "error_invalid",
			json:    "",
			data:    nil,
			isError: true,
		},
		{
			name:    "error_no_source",
			json:    `{}`,
			data:    nil,
			isError: true,
		},
		{
			name:    "error_invalid_source",
			json:    `{"source": {"source": ""}`,
			data:    nil,
			isError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatBoost{}
			err := c.UnmarshalJSON([]byte(tt.json))
			if tt.isError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.EqualValues(t, tt.data, c)
		})
	}
}

func Test_ChatBoostRemoved_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		data    *ChatBoostRemoved
		isError bool
	}{
		{
			name: "success",
			json: `{"boost_id": "1", "source": {"source": "premium"}}`,
			data: &ChatBoostRemoved{
				BoostID: "1",
				Source: &ChatBoostSourcePremium{
					Source: BoostSourcePremium,
				},
			},
			isError: false,
		},
		{
			name:    "error_invalid",
			json:    "",
			data:    nil,
			isError: true,
		},
		{
			name:    "error_no_source",
			json:    `{}`,
			data:    nil,
			isError: true,
		},
		{
			name:    "error_invalid_source",
			json:    `{"source": {"source": ""}`,
			data:    nil,
			isError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatBoostRemoved{}
			err := c.UnmarshalJSON([]byte(tt.json))
			if tt.isError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.EqualValues(t, tt.data, c)
		})
	}
}

func Test_BackgroundTypeFill_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		data    *BackgroundTypeFill
		isError bool
	}{
		{
			name: "success",
			json: `{"type": "fill", "fill": {"type": "solid"}}`,
			data: &BackgroundTypeFill{
				Type: BackgroundTypeNameFill,
				Fill: &BackgroundFillSolid{
					Type: BackgroundFilledSolid,
				},
			},
			isError: false,
		},
		{
			name:    "error_invalid",
			json:    "",
			data:    nil,
			isError: true,
		},
		{
			name:    "error_no_fill",
			json:    `{}`,
			data:    nil,
			isError: true,
		},
		{
			name:    "error_invalid_fill",
			json:    `{"fill": {"type": ""}`,
			data:    nil,
			isError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &BackgroundTypeFill{}
			err := c.UnmarshalJSON([]byte(tt.json))
			if tt.isError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.EqualValues(t, tt.data, c)
		})
	}
}

func Test_ChatBackground_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		data    *ChatBackground
		isError bool
	}{
		{
			name: "success",
			json: `{"type": {"type": "wallpaper"}}`,
			data: &ChatBackground{
				Type: &BackgroundTypeWallpaper{
					Type: BackgroundTypeNameWallpaper,
				},
			},
			isError: false,
		},
		{
			name:    "error_invalid",
			json:    "",
			data:    nil,
			isError: true,
		},
		{
			name:    "error_no_type",
			json:    `{}`,
			data:    nil,
			isError: true,
		},
		{
			name:    "error_invalid_type",
			json:    `{"type": {"type": ""}`,
			data:    nil,
			isError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatBackground{}
			err := c.UnmarshalJSON([]byte(tt.json))
			if tt.isError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.EqualValues(t, tt.data, c)
		})
	}
}
