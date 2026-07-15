package telego

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego/internal/json"
	ta "github.com/mymmrac/telego/telegoapi"
)

func TestTypesRichBlocks(t *testing.T) {
	assert.Implements(t, (*MessageOrigin)(nil), &MessageOriginUser{})
	assert.Equal(t, OriginTypeUser, (&MessageOriginUser{}).OriginType())

	assert.Implements(t, (*MessageOrigin)(nil), &MessageOriginHiddenUser{})
	assert.Equal(t, OriginTypeHiddenUser, (&MessageOriginHiddenUser{}).OriginType())

	assert.Implements(t, (*MessageOrigin)(nil), &MessageOriginChat{})
	assert.Equal(t, OriginTypeChat, (&MessageOriginChat{}).OriginType())

	assert.Implements(t, (*MessageOrigin)(nil), &MessageOriginChannel{})
	assert.Equal(t, OriginTypeChannel, (&MessageOriginChannel{}).OriginType())

	assert.Implements(t, (*PaidMedia)(nil), &PaidMediaLivePhoto{})
	assert.Equal(t, PaidMediaTypeLivePhoto, (&PaidMediaLivePhoto{}).MediaType())

	assert.Implements(t, (*PaidMedia)(nil), &PaidMediaPhoto{})
	assert.Equal(t, PaidMediaTypePhoto, (&PaidMediaPhoto{}).MediaType())

	assert.Implements(t, (*PaidMedia)(nil), &PaidMediaPreview{})
	assert.Equal(t, PaidMediaTypePreview, (&PaidMediaPreview{}).MediaType())

	assert.Implements(t, (*PaidMedia)(nil), &PaidMediaVideo{})
	assert.Equal(t, PaidMediaTypeVideo, (&PaidMediaVideo{}).MediaType())

	assert.Implements(t, (*PaidMedia)(nil), &paidMediaOther{})
	assert.Equal(t, paidMediaTypeOther, (&paidMediaOther{}).MediaType())

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

	assert.Implements(t, (*StoryAreaType)(nil), &StoryAreaTypeLocation{})
	assert.Equal(t, StoryAreaLocation, (&StoryAreaTypeLocation{}).StoryAreaType())

	assert.Implements(t, (*StoryAreaType)(nil), &StoryAreaTypeSuggestedReaction{})
	assert.Equal(t, StoryAreaSuggestedReaction, (&StoryAreaTypeSuggestedReaction{}).StoryAreaType())

	assert.Implements(t, (*StoryAreaType)(nil), &StoryAreaTypeLink{})
	assert.Equal(t, StoryAreaLink, (&StoryAreaTypeLink{}).StoryAreaType())

	assert.Implements(t, (*StoryAreaType)(nil), &StoryAreaTypeWeather{})
	assert.Equal(t, StoryAreaWeather, (&StoryAreaTypeWeather{}).StoryAreaType())

	assert.Implements(t, (*StoryAreaType)(nil), &StoryAreaTypeUniqueGift{})
	assert.Equal(t, StoryAreaUniqueGift, (&StoryAreaTypeUniqueGift{}).StoryAreaType())

	assert.Implements(t, (*ReactionType)(nil), &ReactionTypeEmoji{})
	assert.Equal(t, ReactionEmoji, (&ReactionTypeEmoji{}).ReactionType())

	assert.Implements(t, (*ReactionType)(nil), &ReactionTypeCustomEmoji{})
	assert.Equal(t, ReactionCustomEmoji, (&ReactionTypeCustomEmoji{}).ReactionType())

	assert.Implements(t, (*ReactionType)(nil), &ReactionTypePaid{})
	assert.Equal(t, ReactionPaid, (&ReactionTypePaid{}).ReactionType())

	assert.Implements(t, (*OwnedGift)(nil), &OwnedGiftRegular{})
	assert.Equal(t, GiftTypeRegular, (&OwnedGiftRegular{}).GiftType())

	assert.Implements(t, (*OwnedGift)(nil), &OwnedGiftUnique{})
	assert.Equal(t, GiftTypeUnique, (&OwnedGiftUnique{}).GiftType())

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

	assert.Implements(t, (*InputMedia)(nil), &InputMediaAnimation{})
	assert.Implements(t, (*InputPollMedia)(nil), &InputMediaAnimation{})
	assert.Implements(t, (*InputPollOptionMedia)(nil), &InputMediaAnimation{})
	assert.Equal(t, MediaTypeAnimation, (&InputMediaAnimation{}).MediaType())

	assert.Implements(t, (*InputMedia)(nil), &InputMediaAudio{})
	assert.Implements(t, (*InputPollMedia)(nil), &InputMediaAudio{})
	assert.Equal(t, MediaTypeAudio, (&InputMediaAudio{}).MediaType())

	assert.Implements(t, (*InputMedia)(nil), &InputMediaDocument{})
	assert.Implements(t, (*InputPollMedia)(nil), &InputMediaDocument{})
	assert.Equal(t, MediaTypeDocument, (&InputMediaDocument{}).MediaType())

	assert.Implements(t, (*InputPollOptionMedia)(nil), &InputMediaLink{})
	assert.Equal(t, MediaTypeLink, (&InputMediaLink{}).MediaType())

	assert.Implements(t, (*InputMedia)(nil), &InputMediaLivePhoto{})
	assert.Implements(t, (*InputPollMedia)(nil), &InputMediaLivePhoto{})
	assert.Implements(t, (*InputPollOptionMedia)(nil), &InputMediaLivePhoto{})
	assert.Equal(t, MediaTypeLivePhoto, (&InputMediaLivePhoto{}).MediaType())

	assert.Implements(t, (*InputPollMedia)(nil), &InputMediaLocation{})
	assert.Implements(t, (*InputPollOptionMedia)(nil), &InputMediaLocation{})
	assert.Equal(t, MediaTypeLocation, (&InputMediaLocation{}).MediaType())

	assert.Implements(t, (*InputMedia)(nil), &InputMediaPhoto{})
	assert.Implements(t, (*InputPollMedia)(nil), &InputMediaPhoto{})
	assert.Implements(t, (*InputPollOptionMedia)(nil), &InputMediaPhoto{})
	assert.Equal(t, MediaTypePhoto, (&InputMediaPhoto{}).MediaType())

	assert.Implements(t, (*InputPollOptionMedia)(nil), &InputMediaSticker{})
	assert.Equal(t, MediaTypeSticker, (&InputMediaSticker{}).MediaType())

	assert.Implements(t, (*InputPollMedia)(nil), &InputMediaVenue{})
	assert.Implements(t, (*InputPollOptionMedia)(nil), &InputMediaVenue{})
	assert.Equal(t, MediaTypeVenue, (&InputMediaVenue{}).MediaType())

	assert.Implements(t, (*InputMedia)(nil), &InputMediaVideo{})
	assert.Implements(t, (*InputPollMedia)(nil), &InputMediaVideo{})
	assert.Implements(t, (*InputPollOptionMedia)(nil), &InputMediaVideo{})
	assert.Equal(t, MediaTypeVideo, (&InputMediaVideo{}).MediaType())

	assert.Implements(t, (*RichMessageMedia)(nil), &InputMediaVoiceNote{})
	assert.Equal(t, MediaTypeVoiceNote, (&InputMediaVoiceNote{}).MediaType())

	assert.Implements(t, (*InputPaidMedia)(nil), &InputPaidMediaLivePhoto{})
	assert.Equal(t, PaidMediaTypeLivePhoto, (&InputPaidMediaLivePhoto{}).MediaType())

	assert.Implements(t, (*InputPaidMedia)(nil), &InputPaidMediaPhoto{})
	assert.Equal(t, PaidMediaTypePhoto, (&InputPaidMediaPhoto{}).MediaType())

	assert.Implements(t, (*InputPaidMedia)(nil), &InputPaidMediaVideo{})
	assert.Equal(t, PaidMediaTypeVideo, (&InputPaidMediaVideo{}).MediaType())

	assert.Implements(t, (*InputProfilePhoto)(nil), &InputProfilePhotoStatic{})
	assert.Equal(t, PhotoTypeStatic, (&InputProfilePhotoStatic{}).ProfilePhotoType())

	assert.Implements(t, (*InputProfilePhoto)(nil), &InputProfilePhotoAnimated{})
	assert.Equal(t, PhotoTypeAnimated, (&InputProfilePhotoAnimated{}).ProfilePhotoType())

	assert.Implements(t, (*InputStoryContent)(nil), &InputStoryContentPhoto{})
	assert.Equal(t, StoryTypePhoto, (&InputStoryContentPhoto{}).StoryType())

	assert.Implements(t, (*InputStoryContent)(nil), &InputStoryContentVideo{})
	assert.Equal(t, StoryTypeVideo, (&InputStoryContentVideo{}).StoryType())

	assert.Implements(t, (*RichText)(nil), ToPtr(RichTextPlain("")))
	assert.Equal(t, TextTypePlain, (ToPtr(RichTextPlain(""))).TextType())

	assert.Implements(t, (*RichText)(nil), ToPtr(RichTextList{}))
	assert.Equal(t, TextTypeList, (ToPtr(RichTextList{})).TextType())

	assert.Implements(t, (*RichText)(nil), &RichTextBold{})
	assert.Equal(t, TextTypeBold, (&RichTextBold{}).TextType())

	assert.Implements(t, (*RichText)(nil), &RichTextItalic{})
	assert.Equal(t, TextTypeItalic, (&RichTextItalic{}).TextType())

	assert.Implements(t, (*RichText)(nil), &RichTextUnderline{})
	assert.Equal(t, TextTypeUnderline, (&RichTextUnderline{}).TextType())

	assert.Implements(t, (*RichText)(nil), &RichTextStrikethrough{})
	assert.Equal(t, TextTypeStrikethrough, (&RichTextStrikethrough{}).TextType())

	assert.Implements(t, (*RichText)(nil), &RichTextSpoiler{})
	assert.Equal(t, TextTypeSpoiler, (&RichTextSpoiler{}).TextType())

	assert.Implements(t, (*RichText)(nil), &RichTextDateTime{})
	assert.Equal(t, TextTypeDateTime, (&RichTextDateTime{}).TextType())

	assert.Implements(t, (*RichText)(nil), &RichTextTextMention{})
	assert.Equal(t, TextTypeTextMention, (&RichTextTextMention{}).TextType())

	assert.Implements(t, (*RichText)(nil), &RichTextSubscript{})
	assert.Equal(t, TextTypeSubscript, (&RichTextSubscript{}).TextType())

	assert.Implements(t, (*RichText)(nil), &RichTextSuperscript{})
	assert.Equal(t, TextTypeSuperscript, (&RichTextSuperscript{}).TextType())

	assert.Implements(t, (*RichText)(nil), &RichTextMarked{})
	assert.Equal(t, TextTypeMarked, (&RichTextMarked{}).TextType())

	assert.Implements(t, (*RichText)(nil), &RichTextCode{})
	assert.Equal(t, TextTypeCode, (&RichTextCode{}).TextType())

	assert.Implements(t, (*RichText)(nil), &RichTextCustomEmoji{})
	assert.Equal(t, TextTypeCustomEmoji, (&RichTextCustomEmoji{}).TextType())

	assert.Implements(t, (*RichText)(nil), &RichTextMathematicalExpression{})
	assert.Equal(t, TextTypeMathematicalExpression, (&RichTextMathematicalExpression{}).TextType())

	assert.Implements(t, (*RichText)(nil), &RichTextURL{})
	assert.Equal(t, TextTypeURL, (&RichTextURL{}).TextType())

	assert.Implements(t, (*RichText)(nil), &RichTextEmailAddress{})
	assert.Equal(t, TextTypeEmailAddress, (&RichTextEmailAddress{}).TextType())

	assert.Implements(t, (*RichText)(nil), &RichTextPhoneNumber{})
	assert.Equal(t, TextTypePhoneNumber, (&RichTextPhoneNumber{}).TextType())

	assert.Implements(t, (*RichText)(nil), &RichTextBankCardNumber{})
	assert.Equal(t, TextTypeBankCardNumber, (&RichTextBankCardNumber{}).TextType())

	assert.Implements(t, (*RichText)(nil), &RichTextMention{})
	assert.Equal(t, TextTypeMention, (&RichTextMention{}).TextType())

	assert.Implements(t, (*RichText)(nil), &RichTextHashtag{})
	assert.Equal(t, TextTypeHashtag, (&RichTextHashtag{}).TextType())

	assert.Implements(t, (*RichText)(nil), &RichTextCashtag{})
	assert.Equal(t, TextTypeCashtag, (&RichTextCashtag{}).TextType())

	assert.Implements(t, (*RichText)(nil), &RichTextBotCommand{})
	assert.Equal(t, TextTypeBotCommand, (&RichTextBotCommand{}).TextType())

	assert.Implements(t, (*RichText)(nil), &RichTextAnchor{})
	assert.Equal(t, TextTypeAnchor, (&RichTextAnchor{}).TextType())

	assert.Implements(t, (*RichText)(nil), &RichTextAnchorLink{})
	assert.Equal(t, TextTypeAnchorLink, (&RichTextAnchorLink{}).TextType())

	assert.Implements(t, (*RichText)(nil), &RichTextReference{})
	assert.Equal(t, TextTypeReference, (&RichTextReference{}).TextType())

	assert.Implements(t, (*RichText)(nil), &RichTextReferenceLink{})
	assert.Equal(t, TextTypeReferenceLink, (&RichTextReferenceLink{}).TextType())

	assert.Implements(t, (*RichBlock)(nil), &RichBlockParagraph{})
	assert.Equal(t, BlockTypeParagraph, (&RichBlockParagraph{}).BlockType())

	assert.Implements(t, (*RichBlock)(nil), &RichBlockSectionHeading{})
	assert.Equal(t, BlockTypeSectionHeading, (&RichBlockSectionHeading{}).BlockType())

	assert.Implements(t, (*RichBlock)(nil), &RichBlockPreformatted{})
	assert.Equal(t, BlockTypePreformatted, (&RichBlockPreformatted{}).BlockType())

	assert.Implements(t, (*RichBlock)(nil), &RichBlockFooter{})
	assert.Equal(t, BlockTypeFooter, (&RichBlockFooter{}).BlockType())

	assert.Implements(t, (*RichBlock)(nil), &RichBlockDivider{})
	assert.Equal(t, BlockTypeDivider, (&RichBlockDivider{}).BlockType())

	assert.Implements(t, (*RichBlock)(nil), &RichBlockMathematicalExpression{})
	assert.Equal(t, BlockTypeMathematicalExpression, (&RichBlockMathematicalExpression{}).BlockType())

	assert.Implements(t, (*RichBlock)(nil), &RichBlockAnchor{})
	assert.Equal(t, BlockTypeAnchor, (&RichBlockAnchor{}).BlockType())

	assert.Implements(t, (*RichBlock)(nil), &RichBlockList{})
	assert.Equal(t, BlockTypeList, (&RichBlockList{}).BlockType())

	assert.Implements(t, (*RichBlock)(nil), &RichBlockBlockQuotation{})
	assert.Equal(t, BlockTypeBlockQuotation, (&RichBlockBlockQuotation{}).BlockType())

	assert.Implements(t, (*RichBlock)(nil), &RichBlockPullQuotation{})
	assert.Equal(t, BlockTypePullQuotation, (&RichBlockPullQuotation{}).BlockType())

	assert.Implements(t, (*RichBlock)(nil), &RichBlockCollage{})
	assert.Equal(t, BlockTypeCollage, (&RichBlockCollage{}).BlockType())

	assert.Implements(t, (*RichBlock)(nil), &RichBlockSlideshow{})
	assert.Equal(t, BlockTypeSlideshow, (&RichBlockSlideshow{}).BlockType())

	assert.Implements(t, (*RichBlock)(nil), &RichBlockTable{})
	assert.Equal(t, BlockTypeTable, (&RichBlockTable{}).BlockType())

	assert.Implements(t, (*RichBlock)(nil), &RichBlockDetails{})
	assert.Equal(t, BlockTypeDetails, (&RichBlockDetails{}).BlockType())

	assert.Implements(t, (*RichBlock)(nil), &RichBlockMap{})
	assert.Equal(t, BlockTypeMap, (&RichBlockMap{}).BlockType())

	assert.Implements(t, (*RichBlock)(nil), &RichBlockAnimation{})
	assert.Equal(t, BlockTypeAnimation, (&RichBlockAnimation{}).BlockType())

	assert.Implements(t, (*RichBlock)(nil), &RichBlockAudio{})
	assert.Equal(t, BlockTypeAudio, (&RichBlockAudio{}).BlockType())

	assert.Implements(t, (*RichBlock)(nil), &RichBlockPhoto{})
	assert.Equal(t, BlockTypePhoto, (&RichBlockPhoto{}).BlockType())

	assert.Implements(t, (*RichBlock)(nil), &RichBlockVideo{})
	assert.Equal(t, BlockTypeVideo, (&RichBlockVideo{}).BlockType())

	assert.Implements(t, (*RichBlock)(nil), &RichBlockVoiceNote{})
	assert.Equal(t, BlockTypeVoiceNote, (&RichBlockVoiceNote{}).BlockType())

	assert.Implements(t, (*RichBlock)(nil), &RichBlockThinking{})
	assert.Equal(t, BlockTypeThinking, (&RichBlockThinking{}).BlockType())

	assert.Implements(t, (*InputRichBlock)(nil), &InputRichBlockParagraph{})
	assert.Equal(t, BlockTypeParagraph, (&InputRichBlockParagraph{}).BlockType())

	assert.Implements(t, (*InputRichBlock)(nil), &InputRichBlockSectionHeading{})
	assert.Equal(t, BlockTypeSectionHeading, (&InputRichBlockSectionHeading{}).BlockType())

	assert.Implements(t, (*InputRichBlock)(nil), &InputRichBlockPreformatted{})
	assert.Equal(t, BlockTypePreformatted, (&InputRichBlockPreformatted{}).BlockType())

	assert.Implements(t, (*InputRichBlock)(nil), &InputRichBlockFooter{})
	assert.Equal(t, BlockTypeFooter, (&InputRichBlockFooter{}).BlockType())

	assert.Implements(t, (*InputRichBlock)(nil), &InputRichBlockDivider{})
	assert.Equal(t, BlockTypeDivider, (&InputRichBlockDivider{}).BlockType())

	assert.Implements(t, (*InputRichBlock)(nil), &InputRichBlockMathematicalExpression{})
	assert.Equal(t, BlockTypeMathematicalExpression, (&InputRichBlockMathematicalExpression{}).BlockType())

	assert.Implements(t, (*InputRichBlock)(nil), &InputRichBlockAnchor{})
	assert.Equal(t, BlockTypeAnchor, (&InputRichBlockAnchor{}).BlockType())

	assert.Implements(t, (*InputRichBlock)(nil), &InputRichBlockList{})
	assert.Equal(t, BlockTypeList, (&InputRichBlockList{}).BlockType())

	assert.Implements(t, (*InputRichBlock)(nil), &InputRichBlockBlockQuotation{})
	assert.Equal(t, BlockTypeBlockQuotation, (&InputRichBlockBlockQuotation{}).BlockType())

	assert.Implements(t, (*InputRichBlock)(nil), &InputRichBlockPullQuotation{})
	assert.Equal(t, BlockTypePullQuotation, (&InputRichBlockPullQuotation{}).BlockType())

	assert.Implements(t, (*InputRichBlock)(nil), &InputRichBlockCollage{})
	assert.Equal(t, BlockTypeCollage, (&InputRichBlockCollage{}).BlockType())

	assert.Implements(t, (*InputRichBlock)(nil), &InputRichBlockSlideshow{})
	assert.Equal(t, BlockTypeSlideshow, (&InputRichBlockSlideshow{}).BlockType())

	assert.Implements(t, (*InputRichBlock)(nil), &InputRichBlockTable{})
	assert.Equal(t, BlockTypeTable, (&InputRichBlockTable{}).BlockType())

	assert.Implements(t, (*InputRichBlock)(nil), &InputRichBlockDetails{})
	assert.Equal(t, BlockTypeDetails, (&InputRichBlockDetails{}).BlockType())

	assert.Implements(t, (*InputRichBlock)(nil), &InputRichBlockMap{})
	assert.Equal(t, BlockTypeMap, (&InputRichBlockMap{}).BlockType())

	assert.Implements(t, (*InputRichBlock)(nil), &InputRichBlockAnimation{})
	assert.Equal(t, BlockTypeAnimation, (&InputRichBlockAnimation{}).BlockType())

	assert.Implements(t, (*InputRichBlock)(nil), &InputRichBlockAudio{})
	assert.Equal(t, BlockTypeAudio, (&InputRichBlockAudio{}).BlockType())

	assert.Implements(t, (*InputRichBlock)(nil), &InputRichBlockPhoto{})
	assert.Equal(t, BlockTypePhoto, (&InputRichBlockPhoto{}).BlockType())

	assert.Implements(t, (*InputRichBlock)(nil), &InputRichBlockVideo{})
	assert.Equal(t, BlockTypeVideo, (&InputRichBlockVideo{}).BlockType())

	assert.Implements(t, (*InputRichBlock)(nil), &InputRichBlockVoiceNote{})
	assert.Equal(t, BlockTypeVoiceNote, (&InputRichBlockVoiceNote{}).BlockType())

	assert.Implements(t, (*InputRichBlock)(nil), &InputRichBlockThinking{})
	assert.Equal(t, BlockTypeThinking, (&InputRichBlockThinking{}).BlockType())

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

	assert.Implements(t, (*InputMessageContent)(nil), &InputRichMessageContent{})
	assert.Equal(t, ContentTypeRich, (&InputRichMessageContent{}).ContentType())

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

	assert.Implements(t, (*TransactionPartner)(nil), &TransactionPartnerChat{})
	assert.Equal(t, PartnerTypeChat, (&TransactionPartnerChat{}).PartnerType())

	assert.Implements(t, (*TransactionPartner)(nil), &TransactionPartnerAffiliateProgram{})
	assert.Equal(t, PartnerTypeAffiliateProgram, (&TransactionPartnerAffiliateProgram{}).PartnerType())

	assert.Implements(t, (*TransactionPartner)(nil), &TransactionPartnerFragment{})
	assert.Equal(t, PartnerTypeFragment, (&TransactionPartnerFragment{}).PartnerType())

	assert.Implements(t, (*TransactionPartner)(nil), &TransactionPartnerTelegramAds{})
	assert.Equal(t, PartnerTypeTelegramAds, (&TransactionPartnerTelegramAds{}).PartnerType())

	assert.Implements(t, (*TransactionPartner)(nil), &TransactionPartnerTelegramApi{})
	assert.Equal(t, PartnerTypeTelegramApi, (&TransactionPartnerTelegramApi{}).PartnerType())

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
			assert.Equal(t, tt.data, c.Data)
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
		assert.Equal(t, User{ID: int64(i) + 1}, cm.MemberUser())
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
		assert.Equal(t, expectedCMU, cmu)
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
			assert.Equal(t, tt.data, m.Data)
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

func TestChatID_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		chatID   ChatID
		isError  bool
	}{
		{
			name:     "empty",
			jsonData: `""`,
			chatID:   ChatID{},
			isError:  false,
		},
		{
			name:     "success_id",
			jsonData: "123",
			chatID: ChatID{
				ID: 123,
			},
			isError: false,
		},
		{
			name:     "success_username",
			jsonData: `"test"`,
			chatID: ChatID{
				Username: "test",
			},
			isError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var chatID ChatID
			err := chatID.UnmarshalJSON([]byte(tt.jsonData))
			if tt.isError {
				require.Error(t, err)
				assert.Nil(t, data)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.chatID, chatID)
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
				File:       testNamedReader{},
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
		"media": testNamedReader{},
	}, im.fileParameters())
	assert.True(t, im.Media.needAttach)
}

func TestInputMediaVideo_fileParameters(t *testing.T) {
	im := &InputMediaVideo{
		Media:     testInputFile,
		Thumbnail: &testInputFile,
	}

	assert.Equal(t, map[string]ta.NamedReader{
		"media":     testNamedReader{},
		"thumbnail": testNamedReader{},
	}, im.fileParameters())
	assert.True(t, im.Media.needAttach)
}

func TestInputMediaAnimation_fileParameters(t *testing.T) {
	im := &InputMediaAnimation{
		Media:     testInputFile,
		Thumbnail: &testInputFile,
	}

	assert.Equal(t, map[string]ta.NamedReader{
		"media":     testNamedReader{},
		"thumbnail": testNamedReader{},
	}, im.fileParameters())
	assert.True(t, im.Media.needAttach)
}

func TestInputMediaAudio_fileParameters(t *testing.T) {
	im := &InputMediaAudio{
		Media:     testInputFile,
		Thumbnail: &testInputFile,
	}

	assert.Equal(t, map[string]ta.NamedReader{
		"media":     testNamedReader{},
		"thumbnail": testNamedReader{},
	}, im.fileParameters())
	assert.True(t, im.Media.needAttach)
}

func TestInputMediaDocument_fileParameters(t *testing.T) {
	im := &InputMediaDocument{
		Media:     testInputFile,
		Thumbnail: &testInputFile,
	}

	assert.Equal(t, map[string]ta.NamedReader{
		"media":     testNamedReader{},
		"thumbnail": testNamedReader{},
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
			EntityTypeDateTime,
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
			ButtonStyleDanger, ButtonStyleSuccess, ButtonStylePrimary,
		},
		{
			MemberStatusCreator, MemberStatusAdministrator, MemberStatusMember, MemberStatusRestricted,
			MemberStatusLeft, MemberStatusBanned,
		},
		{
			StoryAreaLocation, StoryAreaSuggestedReaction, StoryAreaLink, StoryAreaWeather, StoryAreaUniqueGift,
		},
		{
			ReactionEmoji, ReactionCustomEmoji, ReactionPaid,
		},
		{
			GiftRarityUncommon, GiftRarityRare, GiftRarityEpic, GiftRarityLegendary,
		},
		{
			GiftOriginUpgrade, GiftOriginTransfer, GiftOriginResale, GiftOriginGiftedUpgrade, GiftOriginOffer,
		},
		{
			GiftTypeRegular, GiftTypeUnique,
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
			MediaTypeAnimation, MediaTypeAudio, MediaTypeDocument, MediaTypeLink, MediaTypeLivePhoto, MediaTypeLocation,
			MediaTypePhoto, MediaTypeSticker, MediaTypeVenue, MediaTypeVideo, MediaTypeVoiceNote,
		},
		{
			PaidMediaTypeLivePhoto, PaidMediaTypePhoto, PaidMediaTypePreview, PaidMediaTypeVideo, paidMediaTypeOther,
		},
		{
			PhotoTypeStatic, PhotoTypeAnimated,
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
			TextTypePlain, TextTypeList,
		},
		{
			TextTypeBold, TextTypeItalic, TextTypeUnderline, TextTypeStrikethrough, TextTypeSpoiler, TextTypeDateTime,
			TextTypeTextMention, TextTypeSubscript, TextTypeSuperscript, TextTypeMarked, TextTypeCode,
			TextTypeCustomEmoji, TextTypeMathematicalExpression, TextTypeURL, TextTypeEmailAddress,
			TextTypePhoneNumber, TextTypeBankCardNumber, TextTypeMention, TextTypeHashtag, TextTypeCashtag,
			TextTypeBotCommand, TextTypeAnchor, TextTypeAnchorLink, TextTypeReference, TextTypeReferenceLink,
		},
		{
			BlockTypeParagraph, BlockTypeSectionHeading, BlockTypePreformatted, BlockTypeFooter, BlockTypeDivider,
			BlockTypeMathematicalExpression, BlockTypeAnchor, BlockTypeList, BlockTypeBlockQuotation,
			BlockTypePullQuotation, BlockTypeCollage, BlockTypeSlideshow, BlockTypeTable, BlockTypeDetails,
			BlockTypeMap, BlockTypeAnimation, BlockTypeAudio, BlockTypePhoto, BlockTypeVideo, BlockTypeVoiceNote,
			BlockTypeThinking,
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
			ContentTypeText, ContentTypeRich, ContentTypeLocation, ContentTypeVenue, ContentTypeContact,
			ContentTypeInvoice,
		},
		{
			WithdrawalStatePending, WithdrawalStateSucceeded, WithdrawalStateFailed,
		},
		{
			PartnerTypeUser, PartnerTypeChat, PartnerTypeAffiliateProgram, PartnerTypeFragment, PartnerTypeTelegramAds,
			PartnerTypeTelegramApi, PartnerTypeOther,
		},
		{
			TransactionTypeInvoicePayment, TransactionTypePaidMediaPayment, TransactionTypeGiftPurchase,
			TransactionTypePremiumPurchase, TransactionTypeBusinessAccountTransfer,
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
	const n1 = 1234567890
	const s1 = "Here is some long text used for testing cloning performance."
	const b1 = true

	c1 := Chat{
		ID:               n1,
		Type:             s1,
		Title:            s1,
		Username:         s1,
		FirstName:        s1,
		LastName:         s1,
		IsForum:          b1,
		IsDirectMessages: b1,
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
		CanConnectToBusiness:    b1,
		HasMainWebApp:           b1,
		HasTopicsEnabled:        b1,
	}

	u := Update{
		UpdateID: n1,
		Message: &Message{
			MessageID:       n1,
			MessageThreadID: n1,
			DirectMessagesTopic: &DirectMessagesTopic{
				TopicID: n1,
				User:    &u1,
			},
			From:                 &u1,
			SenderChat:           &c1,
			SenderBoostCount:     n1,
			Date:                 n1,
			BusinessConnectionID: s1,
			Chat:                 c1,
			ForwardOrigin: &MessageOriginChat{
				Type:            OriginTypeChat,
				Date:            n1,
				SenderChat:      c1,
				AuthorSignature: s1,
			},
			IsTopicMessage:         b1,
			IsAutomaticForward:     b1,
			ReplyToChecklistTaskID: n1,
			ViaBot:                 &u1,
			EditDate:               n1,
			HasProtectedContent:    b1,
			IsFromOffline:          b1,
			IsPaidPost:             b1,
			MediaGroupID:           s1,
			AuthorSignature:        s1,
			PaidStarCount:          n1,
			Text:                   s1,
			EffectID:               s1,
			Caption:                s1,
			ShowCaptionAboveMedia:  b1,
			HasMediaSpoiler:        b1,
			NewChatTitle:           s1,
			DeleteChatPhoto:        b1,
			GroupChatCreated:       b1,
			SupergroupChatCreated:  b1,
			ChannelChatCreated:     b1,
			MigrateToChatID:        n1,
			MigrateFromChatID:      n1,
			ConnectedWebsite:       s1,
		},
	}

	for b.Loop() {
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
	panic("unreachable: badChatMember.MemberStatus")
}

func (b badChatMember) MemberUser() User {
	panic("unreachable: badChatMember.MemberUser")
}

func (b badChatMember) MemberIsMember() bool {
	panic("unreachable: badChatMember.MemberIsMember")
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
				File: &testNamedReader{},
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

	ctx := t.Context()
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
			assert.Equal(t, tt.data, c)
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
			assert.Equal(t, tt.data, e)
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
			assert.Equal(t, tt.data, c)
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
			assert.Equal(t, tt.data, c)
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
			assert.Equal(t, tt.data, m)
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
			assert.Equal(t, tt.data, c)
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
			assert.Equal(t, tt.data, c)
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
			assert.Equal(t, tt.data, c)
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
			assert.Equal(t, tt.data, c)
		})
	}
}

func Test_RichTextBold_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		data    *RichTextBold
		isError bool
	}{
		{
			name: "success_text",
			json: `{"text": "hello"}`,
			data: &RichTextBold{
				Text: ToPtr(RichTextPlain("hello")),
			},
			isError: false,
		},
		{
			name: "success_list",
			json: `{"text": [ "hello", "world" ]}`,
			data: &RichTextBold{
				Text: &RichTextList{
					ToPtr(RichTextPlain("hello")),
					ToPtr(RichTextPlain("world")),
				},
			},
			isError: false,
		},
		{
			name: "success_bold",
			json: `{"text": {"type": "bold", "text": "hello"}}`,
			data: &RichTextBold{
				Text: &RichTextBold{
					Type: TextTypeBold,
					Text: ToPtr(RichTextPlain("hello")),
				},
			},
			isError: false,
		},
		{
			name: "success_list_bold",
			json: `{"text": [ "hello", {"type": "bold", "text": "world"} ]}`,
			data: &RichTextBold{
				Text: &RichTextList{
					ToPtr(RichTextPlain("hello")),
					&RichTextBold{
						Type: TextTypeBold,
						Text: ToPtr(RichTextPlain("world")),
					},
				},
			},
			isError: false,
		},
		{
			name: "success_list_list_bold",
			json: `{"text": [ "hello", [ "world", {"type": "bold", "text": "there"} ] ]}`,
			data: &RichTextBold{
				Text: &RichTextList{
					ToPtr(RichTextPlain("hello")),
					&RichTextList{
						ToPtr(RichTextPlain("world")),
						&RichTextBold{
							Type: TextTypeBold,
							Text: ToPtr(RichTextPlain("there")),
						},
					},
				},
			},
			isError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &RichTextBold{}
			err := p.UnmarshalJSON([]byte(tt.json))
			if tt.isError {
				require.Error(t, err)
				return
			}

			formated, _ := json.Marshal(p) //nolint:errcheck
			t.Log(string(formated))

			require.NoError(t, err)
			assert.Equal(t, tt.data, p)
		})
	}
}

//nolint:lll
func Test_Message_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		data    *Message
		isError bool
	}{
		{
			name: "success_rich_text",
			json: `
{
    "message_id": 568,
    "from": {
        "id": 1,
        "is_bot": false,
        "first_name": "A",
        "last_name": "B",
        "username": "C",
        "language_code": "en"
    },
    "date": 1781429550,
    "chat": {
        "id": 1,
        "type": "private",
        "username": "C",
        "first_name": "A",
        "last_name": "B"
    },
    "forward_origin": {
        "type": "user",
        "date": 1781292642,
        "sender_user": {
            "id": 8886479340,
            "is_bot": true,
            "first_name": "Rich Text Demo",
            "username": "richtextdemobot"
        }
    },
    "rich_message": {
        "blocks": [
            {
                "type": "anchor",
                "name": "top"
            },
            {
                "type": "heading",
                "text": "All Types Demo",
                "size": 1
            },
            {
                "type": "divider"
            },
            {
                "type": "paragraph",
                "text": [
                    "Jump to ",
                    {
                        "type": "anchor_link",
                        "text": "text",
                        "anchor_name": "text"
                    },
                    " | ",
                    {
                        "type": "anchor_link",
                        "text": "structure",
                        "anchor_name": "structure"
                    },
                    " | ",
                    {
                        "type": "anchor_link",
                        "text": "media",
                        "anchor_name": "media"
                    },
                    " | ",
                    {
                        "type": "anchor_link",
                        "text": "advanced",
                        "anchor_name": "advanced"
                    }
                ]
            },
            {
                "type": "anchor",
                "name": "text"
            },
            {
                "type": "heading",
                "text": "Text Formatting",
                "size": 2
            },
            {
                "type": "paragraph",
                "text": [
                    "I want to see ",
                    {
                        "type": "bold",
                        "text": "mountains"
                    },
                    " again, Gandalf ",
                    {
                        "type": "italic",
                        "text": "You talking to me?"
                    },
                    " It's alive! ",
                    {
                        "type": "underline",
                        "text": "It's alive!"
                    },
                    " One, two, ",
                    {
                        "type": "strikethrough",
                        "text": "five!"
                    },
                    ", three! The name's Bond. ",
                    {
                        "type": "spoiler",
                        "text": "James Bond"
                    },
                    " You're a ",
                    {
                        "type": "marked",
                        "text": "wizard"
                    },
                    ", Harry ",
                    {
                        "type": "code",
                        "text": "I'm sorry, Dave."
                    }
                ]
            },
            {
                "type": "paragraph",
                "text": {
                    "type": "reference",
                    "text": "I understood that reference",
                    "name": "note1"
                }
            },
            {
                "type": "paragraph",
                "text": [
                    "Show any emotion with custom emoji ",
                    {
                        "type": "custom_emoji",
                        "custom_emoji_id": "5208541126583136130",
                        "alternative_text": "🎉"
                    },
                    {
                        "type": "custom_emoji",
                        "custom_emoji_id": "5384182985224374928",
                        "alternative_text": "🧐"
                    },
                    {
                        "type": "custom_emoji",
                        "custom_emoji_id": "6052851174929860280",
                        "alternative_text": "😓"
                    }
                ]
            },
            {
                "type": "paragraph",
                "text": [
                    "Solve for ",
                    {
                        "type": "mathematical_expression",
                        "expression": "x"
                    },
                    " and other variables ",
                    {
                        "type": "mathematical_expression",
                        "expression": "E = mc^2"
                    },
                    ", ",
                    {
                        "type": "mathematical_expression",
                        "expression": "a^2 + b^2 = c^2"
                    }
                ]
            },
            {
                "type": "paragraph",
                "text": [
                    "Keep track of important dates ",
                    {
                        "type": "date_time",
                        "text": "Aug 13, 2013",
                        "unix_time": 1735689600,
                        "date_time_format": "D"
                    }
                ]
            },
            {
                "type": "anchor",
                "name": "structure"
            },
            {
                "type": "heading",
                "text": "Structure",
                "size": 2
            },
            {
                "type": "blockquote",
                "blocks": [
                    {
                        "type": "paragraph",
                        "text": "That's what she said"
                    }
                ]
            },
            {
                "type": "blockquote",
                "blocks": [
                    {
                        "type": "paragraph",
                        "text": [
                            "Also available in multiple lines With ",
                            {
                                "type": "bold",
                                "text": "formatting"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "text": "for useful context, and not just jokes"
                    }
                ]
            },
            {
                "type": "list",
                "items": [
                    {
                        "label": "•",
                        "blocks": [
                            {
                                "type": "paragraph",
                                "text": "To the point"
                            }
                        ]
                    },
                    {
                        "label": "•",
                        "blocks": [
                            {
                                "type": "paragraph",
                                "text": "Mission Accomplished"
                            }
                        ],
                        "has_checkbox": true,
                        "is_checked": true
                    }
                ]
            },
            {
                "type": "list",
                "items": [
                    {
                        "label": "1.",
                        "blocks": [
                            {
                                "type": "paragraph",
                                "text": "Step 1"
                            }
                        ],
                        "value": 1,
                        "type": "1"
                    },
                    {
                        "label": "2.",
                        "blocks": [
                            {
                                "type": "paragraph",
                                "text": "Step 2"
                            }
                        ],
                        "value": 2,
                        "type": "1"
                    },
                    {
                        "label": "3.",
                        "blocks": [
                            {
                                "type": "paragraph",
                                "text": "???"
                            }
                        ],
                        "value": 3,
                        "type": "1"
                    },
                    {
                        "label": "4.",
                        "blocks": [
                            {
                                "type": "paragraph",
                                "text": "Profit"
                            }
                        ],
                        "value": 4,
                        "type": "1"
                    }
                ]
            },
            {
                "type": "paragraph",
                "text": "Communicate with the machines"
            },
            {
                "type": "pre",
                "text": "echo \"hello\";",
                "language": "php"
            },
            {
                "type": "divider"
            },
            {
                "type": "details",
                "summary": "Show me more",
                "blocks": [
                    {
                        "type": "paragraph",
                        "text": "Oh, you actually opened this section. I guess I should've thought of something clever to put here but I didn't think you'd actually do it."
                    }
                ]
            },
            {
                "type": "anchor",
                "name": "media"
            },
            {
                "type": "heading",
                "text": "Media",
                "size": 2
            },
            {
                "type": "photo",
                "photo": [
                    {
                        "file_id": "AgACAgEAAxUAAWoudS4p50CO_voLJEmVF1xcPxFzAAKLDmsbOIBYRYg7EJIa7FX4AQADAgADcwADPAQ",
                        "file_unique_id": "AQADiw5rGziAWEV4",
                        "width": 90,
                        "height": 90,
                        "file_size": 1842
                    },
                    {
                        "file_id": "AgACAgEAAxUAAWoudS4p50CO_voLJEmVF1xcPxFzAAKLDmsbOIBYRYg7EJIa7FX4AQADAgADbQADPAQ",
                        "file_unique_id": "AQADiw5rGziAWEVy",
                        "width": 320,
                        "height": 320,
                        "file_size": 28310
                    },
                    {
                        "file_id": "AgACAgEAAxUAAWoudS4p50CO_voLJEmVF1xcPxFzAAKLDmsbOIBYRYg7EJIa7FX4AQADAgADeAADPAQ",
                        "file_unique_id": "AQADiw5rGziAWEV9",
                        "width": 800,
                        "height": 800,
                        "file_size": 140648
                    },
                    {
                        "file_id": "AgACAgEAAxUAAWoudS4p50CO_voLJEmVF1xcPxFzAAKLDmsbOIBYRYg7EJIa7FX4AQADAgADeQADPAQ",
                        "file_unique_id": "AQADiw5rGziAWEV-",
                        "width": 1024,
                        "height": 1024,
                        "file_size": 195428
                    }
                ]
            },
            {
                "type": "video",
                "video": {
                    "file_id": "BAACAgEAAxUAAWoudS66mV3Q_3p10EoaSjkwU3QeAAKTBQACOIBYRcgXAAHZYA8vPTwE",
                    "file_unique_id": "AgADkwUAAjiAWEU",
                    "width": 1088,
                    "height": 1088,
                    "duration": 6,
                    "thumbnail": {
                        "file_id": "AAMCAQADFQABai51LrqZXdD_enXQShpKOTBTdB4AApMFAAI4gFhFyBcAAdlgDy89AQAHbQADPAQ",
                        "file_unique_id": "AQADkwUAAjiAWEVy",
                        "width": 320,
                        "height": 320,
                        "file_size": 24168
                    },
                    "file_name": "dubaiVideo.mp4",
                    "mime_type": "video/mp4",
                    "file_size": 5045080
                }
            },
            {
                "type": "audio",
                "audio": {
                    "file_id": "CQACAgEAAxUAAWoudS5SVRzsS6c0eGqZh1j9SQq_AAKNBQACOIBYRbYY6PeTTClJPAQ",
                    "file_unique_id": "AgADjQUAAjiAWEU",
                    "duration": 224,
                    "performer": "alphavano",
                    "title": "Neon Rain Train",
                    "file_name": "Neon Rain Train.mp3",
                    "mime_type": "audio/mpeg",
                    "file_size": 5328656,
                    "thumbnail": {
                        "file_id": "AAMCAQADFQABai51LlJVHOxLpzR4apmHWP1JCr8AAo0FAAI4gFhFthjo95NMKUkBAAdtAAM8BA",
                        "file_unique_id": "AQADjQUAAjiAWEVy",
                        "width": 320,
                        "height": 320,
                        "file_size": 26468
                    }
                }
            },
            {
                "type": "collage",
                "blocks": [
                    {
                        "type": "photo",
                        "photo": [
                            {
                                "file_id": "AgACAgEAAxUAAWoudS6Wk9IK8vQ_tDMCCGoPEg-oAAKIDmsbOIBYRYZeWkDKYjtCAQADAgADcwADPAQ",
                                "file_unique_id": "AQADiA5rGziAWEV4",
                                "width": 90,
                                "height": 48,
                                "file_size": 1359
                            },
                            {
                                "file_id": "AgACAgEAAxUAAWoudS6Wk9IK8vQ_tDMCCGoPEg-oAAKIDmsbOIBYRYZeWkDKYjtCAQADAgADbQADPAQ",
                                "file_unique_id": "AQADiA5rGziAWEVy",
                                "width": 320,
                                "height": 169,
                                "file_size": 25622
                            },
                            {
                                "file_id": "AgACAgEAAxUAAWoudS6Wk9IK8vQ_tDMCCGoPEg-oAAKIDmsbOIBYRYZeWkDKYjtCAQADAgADeAADPAQ",
                                "file_unique_id": "AQADiA5rGziAWEV9",
                                "width": 800,
                                "height": 422,
                                "file_size": 122928
                            },
                            {
                                "file_id": "AgACAgEAAxUAAWoudS6Wk9IK8vQ_tDMCCGoPEg-oAAKIDmsbOIBYRYZeWkDKYjtCAQADAgADeQADPAQ",
                                "file_unique_id": "AQADiA5rGziAWEV-",
                                "width": 1100,
                                "height": 580,
                                "file_size": 176934
                            }
                        ]
                    },
                    {
                        "type": "photo",
                        "photo": [
                            {
                                "file_id": "AgACAgEAAxUAAWoudS51v4w_bPsXHwABPGxSEW6tPwAChg5rGziAWEVkghk4Jtl-5gEAAwIAA3MAAzwE",
                                "file_unique_id": "AQADhg5rGziAWEV4",
                                "width": 90,
                                "height": 60,
                                "file_size": 1851
                            },
                            {
                                "file_id": "AgACAgEAAxUAAWoudS51v4w_bPsXHwABPGxSEW6tPwAChg5rGziAWEVkghk4Jtl-5gEAAwIAA20AAzwE",
                                "file_unique_id": "AQADhg5rGziAWEVy",
                                "width": 320,
                                "height": 213,
                                "file_size": 31488
                            },
                            {
                                "file_id": "AgACAgEAAxUAAWoudS51v4w_bPsXHwABPGxSEW6tPwAChg5rGziAWEVkghk4Jtl-5gEAAwIAA3gAAzwE",
                                "file_unique_id": "AQADhg5rGziAWEV9",
                                "width": 800,
                                "height": 533,
                                "file_size": 140763
                            },
                            {
                                "file_id": "AgACAgEAAxUAAWoudS51v4w_bPsXHwABPGxSEW6tPwAChg5rGziAWEVkghk4Jtl-5gEAAwIAA3kAAzwE",
                                "file_unique_id": "AQADhg5rGziAWEV-",
                                "width": 1170,
                                "height": 780,
                                "file_size": 222750
                            }
                        ]
                    }
                ]
            },
            {
                "type": "map",
                "location": {
                    "latitude": 25.195949,
                    "longitude": 55.273412
                },
                "zoom": 15,
                "width": 400,
                "height": 200,
                "caption": {
                    "text": "Where are we?"
                }
            },
            {
                "type": "anchor",
                "name": "advanced"
            },
            {
                "type": "heading",
                "text": "Advanced",
                "size": 2
            },
            {
                "type": "table",
                "cells": [
                    [
                        {
                            "text": "Type",
                            "is_header": true,
                            "align": "center",
                            "valign": "middle"
                        },
                        {
                            "text": "Supported",
                            "is_header": true,
                            "align": "center",
                            "valign": "middle"
                        },
                        {
                            "text": "Composition",
                            "is_header": true,
                            "align": "center",
                            "valign": "middle"
                        }
                    ],
                    [
                        {
                            "text": "Table",
                            "align": "left",
                            "valign": "middle"
                        },
                        {
                            "text": {
                                "type": "bold",
                                "text": "Yes"
                            },
                            "align": "left",
                            "valign": "middle"
                        },
                        {
                            "text": "100% text",
                            "align": "left",
                            "valign": "middle"
                        }
                    ]
                ],
                "is_bordered": true,
                "is_striped": true
            },
            {
                "type": "paragraph",
                "text": "I don't know what this is but apparently it's also math 👇"
            },
            {
                "type": "mathematical_expression",
                "expression": "\\sum_{i=1}^n i = \\frac{n(n+1)}{2}"
            },
            {
                "type": "pullquote",
                "text": "To be truly free, you should be ready to risk everything for freedom.",
                "credit": "Pavel Durov"
            },
            {
                "type": "paragraph",
                "text": [
                    "Make sure to always cite your sources",
                    {
                        "type": "superscript",
                        "text": [
                            {
                                "type": "anchor",
                                "name": "fnref-1-1"
                            },
                            {
                                "type": "reference_link",
                                "text": "1",
                                "reference_name": "fn-1"
                            }
                        ]
                    },
                    "."
                ]
            },
            {
                "type": "paragraph",
                "text": [
                    {
                        "type": "url",
                        "text": "Link",
                        "url": "https://youtu.be/dQw4w9WgXcQ"
                    },
                    " · ",
                    {
                        "type": "email_address",
                        "text": "Email",
                        "email_address": "user@example.com"
                    },
                    " · ",
                    {
                        "type": "text_mention",
                        "text": "User Mention",
                        "user": {
                            "id": 777000,
                            "is_bot": false,
                            "first_name": "Telegram"
                        }
                    }
                ]
            },
            {
                "type": "footer",
                "text": [
                    "1. ",
                    {
                        "type": "reference",
                        "text": "Source: me, because I said so.",
                        "name": "fn-1"
                    },
                    " ",
                    {
                        "type": "anchor_link",
                        "text": "↩",
                        "anchor_name": "fnref-1-1"
                    }
                ]
            }
        ]
    }
}
`,
			data: &Message{
				MessageID: 568,
				From: &User{
					ID:           1,
					FirstName:    "A",
					LastName:     "B",
					Username:     "C",
					LanguageCode: "en",
				},
				Date: 1781429550,
				Chat: Chat{
					ID:               1,
					Type:             "private",
					Title:            "",
					Username:         "C",
					FirstName:        "A",
					LastName:         "B",
					IsForum:          false,
					IsDirectMessages: false,
				},
				ForwardOrigin: &MessageOriginUser{
					Type: "user",
					Date: 1781292642,
					SenderUser: User{
						ID:        8886479340,
						IsBot:     true,
						FirstName: "Rich Text Demo",
						Username:  "richtextdemobot",
					},
				},
				RichMessage: &RichMessage{
					Blocks: []RichBlock{
						&RichBlockAnchor{
							Type: "anchor",
							Name: "top",
						},
						&RichBlockSectionHeading{
							Type: "heading",
							Text: ToPtr(RichTextPlain("All Types Demo")),
							Size: 1,
						},
						&RichBlockDivider{
							Type: "divider",
						},
						&RichBlockParagraph{
							Type: "paragraph",
							Text: &RichTextList{
								ToPtr(RichTextPlain("Jump to ")),
								&RichTextAnchorLink{
									Type:       "anchor_link",
									Text:       ToPtr(RichTextPlain("text")),
									AnchorName: "text",
								},
								ToPtr(RichTextPlain(" | ")),
								&RichTextAnchorLink{
									Type:       "anchor_link",
									Text:       ToPtr(RichTextPlain("structure")),
									AnchorName: "structure",
								},
								ToPtr(RichTextPlain(" | ")),
								&RichTextAnchorLink{
									Type:       "anchor_link",
									Text:       ToPtr(RichTextPlain("media")),
									AnchorName: "media",
								},
								ToPtr(RichTextPlain(" | ")),
								&RichTextAnchorLink{
									Type:       "anchor_link",
									Text:       ToPtr(RichTextPlain("advanced")),
									AnchorName: "advanced",
								},
							},
						},
						&RichBlockAnchor{
							Type: "anchor",
							Name: "text",
						},
						&RichBlockSectionHeading{
							Type: "heading",
							Text: ToPtr(RichTextPlain("Text Formatting")),
							Size: 2,
						},
						&RichBlockParagraph{
							Type: "paragraph",
							Text: &RichTextList{
								ToPtr(RichTextPlain("I want to see ")),
								&RichTextBold{
									Type: "bold",
									Text: ToPtr(RichTextPlain("mountains")),
								},
								ToPtr(RichTextPlain(" again, Gandalf ")),
								&RichTextItalic{
									Type: "italic",
									Text: ToPtr(RichTextPlain("You talking to me?")),
								},
								ToPtr(RichTextPlain(" It's alive! ")),
								&RichTextUnderline{
									Type: "underline",
									Text: ToPtr(RichTextPlain("It's alive!")),
								},
								ToPtr(RichTextPlain(" One, two, ")),
								&RichTextStrikethrough{
									Type: "strikethrough",
									Text: ToPtr(RichTextPlain("five!")),
								},
								ToPtr(RichTextPlain(", three! The name's Bond. ")),
								&RichTextSpoiler{
									Type: "spoiler",
									Text: ToPtr(RichTextPlain("James Bond")),
								},
								ToPtr(RichTextPlain(" You're a ")),
								&RichTextMarked{
									Type: "marked",
									Text: ToPtr(RichTextPlain("wizard")),
								},
								ToPtr(RichTextPlain(", Harry ")),
								&RichTextCode{
									Type: "code",
									Text: ToPtr(RichTextPlain("I'm sorry, Dave.")),
								},
							},
						},
						&RichBlockParagraph{
							Type: "paragraph",
							Text: &RichTextReference{
								Type: "reference",
								Text: ToPtr(RichTextPlain("I understood that reference")),
								Name: "note1",
							},
						},
						&RichBlockParagraph{
							Type: "paragraph",
							Text: &RichTextList{
								ToPtr(RichTextPlain("Show any emotion with custom emoji ")),
								&RichTextCustomEmoji{
									Type:            "custom_emoji",
									CustomEmojiID:   "5208541126583136130",
									AlternativeText: "🎉",
								},
								&RichTextCustomEmoji{
									Type:            "custom_emoji",
									CustomEmojiID:   "5384182985224374928",
									AlternativeText: "🧐",
								},
								&RichTextCustomEmoji{
									Type:            "custom_emoji",
									CustomEmojiID:   "6052851174929860280",
									AlternativeText: "😓",
								},
							},
						},
						&RichBlockParagraph{
							Type: "paragraph",
							Text: &RichTextList{
								ToPtr(RichTextPlain("Solve for ")),
								&RichTextMathematicalExpression{
									Type:       "mathematical_expression",
									Expression: "x",
								},
								ToPtr(RichTextPlain(" and other variables ")),
								&RichTextMathematicalExpression{
									Type:       "mathematical_expression",
									Expression: "E = mc^2",
								},
								ToPtr(RichTextPlain(", ")),
								&RichTextMathematicalExpression{
									Type:       "mathematical_expression",
									Expression: "a^2 + b^2 = c^2",
								},
							},
						},
						&RichBlockParagraph{
							Type: "paragraph",
							Text: &RichTextList{
								ToPtr(RichTextPlain("Keep track of important dates ")),
								&RichTextDateTime{
									Type:           "date_time",
									Text:           ToPtr(RichTextPlain("Aug 13, 2013")),
									UnixTime:       1735689600,
									DateTimeFormat: "D",
								},
							},
						},
						&RichBlockAnchor{
							Type: "anchor",
							Name: "structure",
						},
						&RichBlockSectionHeading{
							Type: "heading",
							Text: ToPtr(RichTextPlain("Structure")),
							Size: 2,
						},
						&RichBlockBlockQuotation{
							Type: "blockquote",
							Blocks: []RichBlock{
								&RichBlockParagraph{
									Type: "paragraph",
									Text: ToPtr(RichTextPlain("That's what she said")),
								},
							},
							Credit: nil,
						},
						&RichBlockBlockQuotation{
							Type: "blockquote",
							Blocks: []RichBlock{
								&RichBlockParagraph{
									Type: "paragraph",
									Text: &RichTextList{
										ToPtr(RichTextPlain("Also available in multiple lines With ")),
										&RichTextBold{
											Type: "bold",
											Text: ToPtr(RichTextPlain("formatting")),
										},
									},
								},
								&RichBlockParagraph{
									Type: "paragraph",
									Text: ToPtr(RichTextPlain("for useful context, and not just jokes")),
								},
							},
							Credit: nil,
						},
						&RichBlockList{
							Type: "list",
							Items: []RichBlockListItem{
								{
									Label: "•",
									Blocks: []RichBlock{
										&RichBlockParagraph{
											Type: "paragraph",
											Text: ToPtr(RichTextPlain("To the point")),
										},
									},
									HasCheckbox: false,
									IsChecked:   false,
									Value:       0,
									Type:        "",
								},
								{
									Label: "•",
									Blocks: []RichBlock{
										&RichBlockParagraph{
											Type: "paragraph",
											Text: ToPtr(RichTextPlain("Mission Accomplished")),
										},
									},
									HasCheckbox: true,
									IsChecked:   true,
									Value:       0,
									Type:        "",
								},
							},
						},
						&RichBlockList{
							Type: "list",
							Items: []RichBlockListItem{
								{
									Label: "1.",
									Blocks: []RichBlock{
										&RichBlockParagraph{
											Type: "paragraph",
											Text: ToPtr(RichTextPlain("Step 1")),
										},
									},
									HasCheckbox: false,
									IsChecked:   false,
									Value:       1,
									Type:        "1",
								},
								{
									Label: "2.",
									Blocks: []RichBlock{
										&RichBlockParagraph{
											Type: "paragraph",
											Text: ToPtr(RichTextPlain("Step 2")),
										},
									},
									HasCheckbox: false,
									IsChecked:   false,
									Value:       2,
									Type:        "1",
								},
								{
									Label: "3.",
									Blocks: []RichBlock{
										&RichBlockParagraph{
											Type: "paragraph",
											Text: ToPtr(RichTextPlain("???")),
										},
									},
									HasCheckbox: false,
									IsChecked:   false,
									Value:       3,
									Type:        "1",
								},
								{
									Label: "4.",
									Blocks: []RichBlock{
										&RichBlockParagraph{
											Type: "paragraph",
											Text: ToPtr(RichTextPlain("Profit")),
										},
									},
									HasCheckbox: false,
									IsChecked:   false,
									Value:       4,
									Type:        "1",
								},
							},
						},
						&RichBlockParagraph{
							Type: "paragraph",
							Text: ToPtr(RichTextPlain("Communicate with the machines")),
						},
						&RichBlockPreformatted{
							Type:     "pre",
							Text:     ToPtr(RichTextPlain("echo \"hello\";")),
							Language: "php",
						},
						&RichBlockDivider{
							Type: "divider",
						},
						&RichBlockDetails{
							Type:    "details",
							Summary: ToPtr(RichTextPlain("Show me more")),
							Blocks: []RichBlock{
								&RichBlockParagraph{
									Type: "paragraph",
									Text: ToPtr(RichTextPlain("Oh, you actually opened this section. I guess I should've thought of something clever to put here but I didn't think you'd actually do it.")),
								},
							},
							IsOpen: false,
						},
						&RichBlockAnchor{
							Type: "anchor",
							Name: "media",
						},
						&RichBlockSectionHeading{
							Type: "heading",
							Text: ToPtr(RichTextPlain("Media")),
							Size: 2,
						},
						&RichBlockPhoto{
							Type: "photo",
							Photo: []PhotoSize{
								{
									FileID:       "AgACAgEAAxUAAWoudS4p50CO_voLJEmVF1xcPxFzAAKLDmsbOIBYRYg7EJIa7FX4AQADAgADcwADPAQ",
									FileUniqueID: "AQADiw5rGziAWEV4",
									Width:        90,
									Height:       90,
									FileSize:     1842,
								},
								{
									FileID:       "AgACAgEAAxUAAWoudS4p50CO_voLJEmVF1xcPxFzAAKLDmsbOIBYRYg7EJIa7FX4AQADAgADbQADPAQ",
									FileUniqueID: "AQADiw5rGziAWEVy",
									Width:        320,
									Height:       320,
									FileSize:     28310,
								},
								{
									FileID:       "AgACAgEAAxUAAWoudS4p50CO_voLJEmVF1xcPxFzAAKLDmsbOIBYRYg7EJIa7FX4AQADAgADeAADPAQ",
									FileUniqueID: "AQADiw5rGziAWEV9",
									Width:        800,
									Height:       800,
									FileSize:     140648,
								},
								{
									FileID:       "AgACAgEAAxUAAWoudS4p50CO_voLJEmVF1xcPxFzAAKLDmsbOIBYRYg7EJIa7FX4AQADAgADeQADPAQ",
									FileUniqueID: "AQADiw5rGziAWEV-",
									Width:        1024,
									Height:       1024,
									FileSize:     195428,
								},
							},
							HasSpoiler: false,
							Caption:    (*RichBlockCaption)(nil),
						},
						&RichBlockVideo{
							Type: "video",
							Video: Video{
								FileID:       "BAACAgEAAxUAAWoudS66mV3Q_3p10EoaSjkwU3QeAAKTBQACOIBYRcgXAAHZYA8vPTwE",
								FileUniqueID: "AgADkwUAAjiAWEU",
								Width:        1088,
								Height:       1088,
								Duration:     6,
								Thumbnail: &PhotoSize{
									FileID:       "AAMCAQADFQABai51LrqZXdD_enXQShpKOTBTdB4AApMFAAI4gFhFyBcAAdlgDy89AQAHbQADPAQ",
									FileUniqueID: "AQADkwUAAjiAWEVy",
									Width:        320,
									Height:       320,
									FileSize:     24168,
								},
								Cover:          []PhotoSize(nil),
								StartTimestamp: 0,
								Qualities:      []VideoQuality(nil),
								FileName:       "dubaiVideo.mp4",
								MimeType:       "video/mp4",
								FileSize:       5045080,
							},
							HasSpoiler: false,
							Caption:    (*RichBlockCaption)(nil),
						},
						&RichBlockAudio{
							Type: "audio",
							Audio: Audio{
								FileID:       "CQACAgEAAxUAAWoudS5SVRzsS6c0eGqZh1j9SQq_AAKNBQACOIBYRbYY6PeTTClJPAQ",
								FileUniqueID: "AgADjQUAAjiAWEU",
								Duration:     224,
								Performer:    "alphavano",
								Title:        "Neon Rain Train",
								FileName:     "Neon Rain Train.mp3",
								MimeType:     "audio/mpeg",
								FileSize:     5328656,
								Thumbnail: &PhotoSize{
									FileID:       "AAMCAQADFQABai51LlJVHOxLpzR4apmHWP1JCr8AAo0FAAI4gFhFthjo95NMKUkBAAdtAAM8BA",
									FileUniqueID: "AQADjQUAAjiAWEVy",
									Width:        320,
									Height:       320,
									FileSize:     26468,
								},
							},
							Caption: (*RichBlockCaption)(nil),
						},
						&RichBlockCollage{
							Type: "collage",
							Blocks: []RichBlock{
								&RichBlockPhoto{
									Type: "photo",
									Photo: []PhotoSize{
										{
											FileID:       "AgACAgEAAxUAAWoudS6Wk9IK8vQ_tDMCCGoPEg-oAAKIDmsbOIBYRYZeWkDKYjtCAQADAgADcwADPAQ",
											FileUniqueID: "AQADiA5rGziAWEV4",
											Width:        90,
											Height:       48,
											FileSize:     1359,
										},
										{
											FileID:       "AgACAgEAAxUAAWoudS6Wk9IK8vQ_tDMCCGoPEg-oAAKIDmsbOIBYRYZeWkDKYjtCAQADAgADbQADPAQ",
											FileUniqueID: "AQADiA5rGziAWEVy",
											Width:        320,
											Height:       169,
											FileSize:     25622,
										},
										{
											FileID:       "AgACAgEAAxUAAWoudS6Wk9IK8vQ_tDMCCGoPEg-oAAKIDmsbOIBYRYZeWkDKYjtCAQADAgADeAADPAQ",
											FileUniqueID: "AQADiA5rGziAWEV9",
											Width:        800,
											Height:       422,
											FileSize:     122928,
										},
										{
											FileID:       "AgACAgEAAxUAAWoudS6Wk9IK8vQ_tDMCCGoPEg-oAAKIDmsbOIBYRYZeWkDKYjtCAQADAgADeQADPAQ",
											FileUniqueID: "AQADiA5rGziAWEV-",
											Width:        1100,
											Height:       580,
											FileSize:     176934,
										},
									},
									HasSpoiler: false,
									Caption:    (*RichBlockCaption)(nil),
								},
								&RichBlockPhoto{
									Type: "photo",
									Photo: []PhotoSize{
										{
											FileID:       "AgACAgEAAxUAAWoudS51v4w_bPsXHwABPGxSEW6tPwAChg5rGziAWEVkghk4Jtl-5gEAAwIAA3MAAzwE",
											FileUniqueID: "AQADhg5rGziAWEV4",
											Width:        90,
											Height:       60,
											FileSize:     1851,
										},
										{
											FileID:       "AgACAgEAAxUAAWoudS51v4w_bPsXHwABPGxSEW6tPwAChg5rGziAWEVkghk4Jtl-5gEAAwIAA20AAzwE",
											FileUniqueID: "AQADhg5rGziAWEVy",
											Width:        320,
											Height:       213,
											FileSize:     31488,
										},
										{
											FileID:       "AgACAgEAAxUAAWoudS51v4w_bPsXHwABPGxSEW6tPwAChg5rGziAWEVkghk4Jtl-5gEAAwIAA3gAAzwE",
											FileUniqueID: "AQADhg5rGziAWEV9",
											Width:        800,
											Height:       533,
											FileSize:     140763,
										},
										{
											FileID:       "AgACAgEAAxUAAWoudS51v4w_bPsXHwABPGxSEW6tPwAChg5rGziAWEVkghk4Jtl-5gEAAwIAA3kAAzwE",
											FileUniqueID: "AQADhg5rGziAWEV-",
											Width:        1170,
											Height:       780,
											FileSize:     222750,
										},
									},
									HasSpoiler: false,
									Caption:    (*RichBlockCaption)(nil),
								},
							},
							Caption: (*RichBlockCaption)(nil),
						},
						&RichBlockMap{
							Type: "map",
							Location: Location{
								Latitude:             25.195949,
								Longitude:            55.273412,
								HorizontalAccuracy:   0.000000,
								LivePeriod:           0,
								Heading:              0,
								ProximityAlertRadius: 0,
							},
							Zoom:   15,
							Width:  400,
							Height: 200,
							Caption: &RichBlockCaption{
								Text:   ToPtr(RichTextPlain("Where are we?")),
								Credit: nil,
							},
						},
						&RichBlockAnchor{
							Type: "anchor",
							Name: "advanced",
						},
						&RichBlockSectionHeading{
							Type: "heading",
							Text: ToPtr(RichTextPlain("Advanced")),
							Size: 2,
						},
						&RichBlockTable{
							Type: "table",
							Cells: [][]RichBlockTableCell{
								{
									{
										Text:     ToPtr(RichTextPlain("Type")),
										IsHeader: true,
										Colspan:  0,
										Rowspan:  0,
										Align:    "center",
										Valign:   "middle",
									},
									{
										Text:     ToPtr(RichTextPlain("Supported")),
										IsHeader: true,
										Colspan:  0,
										Rowspan:  0,
										Align:    "center",
										Valign:   "middle",
									},
									{
										Text:     ToPtr(RichTextPlain("Composition")),
										IsHeader: true,
										Colspan:  0,
										Rowspan:  0,
										Align:    "center",
										Valign:   "middle",
									},
								},
								{
									{
										Text:     ToPtr(RichTextPlain("Table")),
										IsHeader: false,
										Colspan:  0,
										Rowspan:  0,
										Align:    "left",
										Valign:   "middle",
									},
									{
										Text: &RichTextBold{
											Type: "bold",
											Text: ToPtr(RichTextPlain("Yes")),
										},
										IsHeader: false,
										Colspan:  0,
										Rowspan:  0,
										Align:    "left",
										Valign:   "middle",
									},
									{
										Text:     ToPtr(RichTextPlain("100% text")),
										IsHeader: false,
										Colspan:  0,
										Rowspan:  0,
										Align:    "left",
										Valign:   "middle",
									},
								},
							},
							IsBordered: true,
							IsStriped:  true,
							Caption:    nil,
						},
						&RichBlockParagraph{
							Type: "paragraph",
							Text: ToPtr(RichTextPlain("I don't know what this is but apparently it's also math 👇")),
						},
						&RichBlockMathematicalExpression{
							Type:       "mathematical_expression",
							Expression: "\\sum_{i=1}^n i = \\frac{n(n+1)}{2}",
						},
						&RichBlockPullQuotation{
							Type:   "pullquote",
							Text:   ToPtr(RichTextPlain("To be truly free, you should be ready to risk everything for freedom.")),
							Credit: ToPtr(RichTextPlain("Pavel Durov")),
						},
						&RichBlockParagraph{
							Type: "paragraph",
							Text: &RichTextList{
								ToPtr(RichTextPlain("Make sure to always cite your sources")),
								&RichTextSuperscript{
									Type: "superscript",
									Text: &RichTextList{
										&RichTextAnchor{
											Type: "anchor",
											Name: "fnref-1-1",
										},
										&RichTextReferenceLink{
											Type:          "reference_link",
											Text:          ToPtr(RichTextPlain("1")),
											ReferenceName: "fn-1",
										},
									},
								},
								ToPtr(RichTextPlain(".")),
							},
						},
						&RichBlockParagraph{
							Type: "paragraph",
							Text: &RichTextList{
								&RichTextURL{
									Type: "url",
									Text: ToPtr(RichTextPlain("Link")),
									URL:  "https://youtu.be/dQw4w9WgXcQ",
								},
								ToPtr(RichTextPlain(" · ")),
								&RichTextEmailAddress{
									Type:         "email_address",
									Text:         ToPtr(RichTextPlain("Email")),
									EmailAddress: "user@example.com",
								},
								ToPtr(RichTextPlain(" · ")),
								&RichTextTextMention{
									Type: "text_mention",
									Text: ToPtr(RichTextPlain("User Mention")),
									User: User{
										ID:                         777000,
										IsBot:                      false,
										FirstName:                  "Telegram",
										LastName:                   "",
										Username:                   "",
										LanguageCode:               "",
										IsPremium:                  false,
										AddedToAttachmentMenu:      false,
										CanJoinGroups:              false,
										CanReadAllGroupMessages:    false,
										SupportsGuestQueries:       false,
										SupportsInlineQueries:      false,
										CanConnectToBusiness:       false,
										HasMainWebApp:              false,
										HasTopicsEnabled:           false,
										AllowsUsersToCreateTopics:  false,
										CanManageBots:              false,
										SupportsJoinRequestQueries: false,
									},
								},
							},
						},
						&RichBlockFooter{
							Type: "footer",
							Text: &RichTextList{
								ToPtr(RichTextPlain("1. ")),
								&RichTextReference{
									Type: "reference",
									Text: ToPtr(RichTextPlain("Source: me, because I said so.")),
									Name: "fn-1",
								},
								ToPtr(RichTextPlain(" ")),
								&RichTextAnchorLink{
									Type:       "anchor_link",
									Text:       ToPtr(RichTextPlain("↩")),
									AnchorName: "fnref-1-1",
								},
							},
						},
					},
				},
			},
			isError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{}
			err := m.UnmarshalJSON([]byte(tt.json))
			if tt.isError {
				require.Error(t, err)
				return
			}

			formated, _ := json.Marshal(m) //nolint:errcheck
			t.Log(string(formated))

			require.NoError(t, err)
			assert.Equal(t, tt.data.RichMessage, m.RichMessage)
		})
	}
}
