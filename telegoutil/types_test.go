package telegoutil

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego"
)

const epsilon = 0.00001

type testNamedReader struct{}

func (t testNamedReader) Read(_ []byte) (n int, err error) {
	panic("unreachable: testNamedReader.Read")
}

func (t testNamedReader) Name() string {
	return text1
}

var (
	nr  = testNamedReader{}
	nr2 = testNamedReader{}
)

func TestFile(t *testing.T) {
	f := File(nr)
	assert.Equal(t, nr, f.File)
}

func TestFileFromReader(t *testing.T) {
	f := FileFromReader(nr, nr.Name())
	assert.Equal(t, nr.Name(), f.File.Name())
}

func TestFileFromBytes(t *testing.T) {
	f := FileFromBytes([]byte(text1), text2)
	assert.Equal(t, text2, f.File.Name())
}

func TestFileByID(t *testing.T) {
	f := FileFromID(text1)
	assert.Equal(t, text1, f.FileID)
}

func TestFileByURL(t *testing.T) {
	f := FileFromURL(text1)
	assert.Equal(t, text1, f.URL)
}

func TestDownloadFile(t *testing.T) {
	expectedData := []byte("OK")
	srv := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path == "/error" {
			writer.WriteHeader(http.StatusInternalServerError)
		} else {
			//nolint:errcheck
			_, _ = writer.Write(expectedData)
			writer.WriteHeader(http.StatusOK)
		}
	}))
	defer srv.Close()

	t.Run("success", func(t *testing.T) {
		data, err := DownloadFile(srv.URL + "/")
		require.NoError(t, err)
		assert.Equal(t, expectedData, data)
	})

	t.Run("error_request", func(t *testing.T) {
		data, err := DownloadFile("")
		require.Error(t, err)
		assert.Nil(t, data)
	})

	t.Run("error_status", func(t *testing.T) {
		data, err := DownloadFile(srv.URL + "/error")
		require.Error(t, err)
		assert.Nil(t, data)
	})
}

func TestID(t *testing.T) {
	i := ID(number2)
	assert.Equal(t, number2, i.ID)
}

func TestUsername(t *testing.T) {
	u := Username(text1)
	assert.Equal(t, text1, u.Username)
}

func TestErrorDataField(t *testing.T) {
	e := ErrorDataField(text1, text2, text3, text4)
	assert.Equal(t, telego.ErrorSourceDataField, e.Source)
	assert.Equal(t, text1, e.Type)
	assert.Equal(t, text2, e.Message)
	assert.Equal(t, text3, e.FieldName)
	assert.Equal(t, text4, e.DataHash)
}

func TestErrorFile(t *testing.T) {
	e := ErrorFile(text1, text2, text3)
	assert.Equal(t, telego.ErrorSourceFile, e.Source)
	assert.Equal(t, text1, e.Type)
	assert.Equal(t, text2, e.Message)
	assert.Equal(t, text3, e.FileHash)
}

func TestErrorFiles(t *testing.T) {
	e := ErrorFiles(text1, text2, text3, text4)
	assert.Equal(t, telego.ErrorSourceFiles, e.Source)
	assert.Equal(t, text1, e.Type)
	assert.Equal(t, text2, e.Message)
	assert.Equal(t, []string{text3, text4}, e.FileHashes)
}

func TestErrorFrontSide(t *testing.T) {
	e := ErrorFrontSide(text1, text2, text3)
	assert.Equal(t, telego.ErrorSourceFrontSide, e.Source)
	assert.Equal(t, text1, e.Type)
	assert.Equal(t, text2, e.Message)
	assert.Equal(t, text3, e.FileHash)
}

func TestErrorReverseSide(t *testing.T) {
	e := ErrorReverseSide(text1, text2, text3)
	assert.Equal(t, telego.ErrorSourceReverseSide, e.Source)
	assert.Equal(t, text1, e.Type)
	assert.Equal(t, text2, e.Message)
	assert.Equal(t, text3, e.FileHash)
}

func TestErrorSelfie(t *testing.T) {
	e := ErrorSelfie(text1, text2, text3)
	assert.Equal(t, telego.ErrorSourceSelfie, e.Source)
	assert.Equal(t, text1, e.Type)
	assert.Equal(t, text2, e.Message)
	assert.Equal(t, text3, e.FileHash)
}

func TestErrorTranslationFile(t *testing.T) {
	e := ErrorTranslationFile(text1, text2, text3)
	assert.Equal(t, telego.ErrorSourceTranslationFile, e.Source)
	assert.Equal(t, text1, e.Type)
	assert.Equal(t, text2, e.Message)
	assert.Equal(t, text3, e.FileHash)
}

func TestErrorTranslationFiles(t *testing.T) {
	e := ErrorTranslationFiles(text1, text2, text3, text4)
	assert.Equal(t, telego.ErrorSourceTranslationFiles, e.Source)
	assert.Equal(t, text1, e.Type)
	assert.Equal(t, text2, e.Message)
	assert.Equal(t, []string{text3, text4}, e.FileHashes)
}

func TestErrorUnspecified(t *testing.T) {
	e := ErrorUnspecified(text1, text2, text3)
	assert.Equal(t, telego.ErrorSourceUnspecified, e.Source)
	assert.Equal(t, text1, e.Type)
	assert.Equal(t, text2, e.Message)
	assert.Equal(t, text3, e.ElementHash)
}

func TestInlineKeyboard(t *testing.T) {
	i := InlineKeyboard([]telego.InlineKeyboardButton{}, []telego.InlineKeyboardButton{})
	assert.Len(t, i.InlineKeyboard, 2)
}

func TestInlineKeyboardButton(t *testing.T) {
	i := InlineKeyboardButton(text1)
	assert.Equal(t, text1, i.Text)
}

func TestInlineKeyboardRow(t *testing.T) {
	i := InlineKeyboardRow(telego.InlineKeyboardButton{}, telego.InlineKeyboardButton{})
	assert.Len(t, i, 2)
}

func TestInlineKeyboardGrid(t *testing.T) {
	i := InlineKeyboardGrid([][]telego.InlineKeyboardButton{
		{{}},
		{{}, {}, {}},
	})
	require.Len(t, i.InlineKeyboard, 2)
	assert.Len(t, i.InlineKeyboard[0], 1)
	assert.Len(t, i.InlineKeyboard[1], 3)
}

func TestInlineKeyboardCols(t *testing.T) {
	t.Run("full", func(t *testing.T) {
		b := telego.InlineKeyboardButton{}
		i := InlineKeyboardCols(2, b, b, b, b)
		require.Len(t, i, 2)
		assert.Len(t, i[0], 2)
		assert.Len(t, i[1], 2)
	})

	t.Run("one_off", func(t *testing.T) {
		b := telego.InlineKeyboardButton{}
		i := InlineKeyboardCols(2, b, b, b)
		require.Len(t, i, 2)
		assert.Len(t, i[0], 2)
		assert.Len(t, i[1], 1)
	})

	t.Run("nil", func(t *testing.T) {
		i := InlineKeyboardCols(0)
		assert.Nil(t, i)
	})
}

func TestInlineKeyboardRows(t *testing.T) {
	t.Run("full", func(t *testing.T) {
		b := telego.InlineKeyboardButton{}
		i := InlineKeyboardRows(2, b, b, b, b)
		require.Len(t, i, 2)
		assert.Len(t, i[0], 2)
		assert.Len(t, i[1], 2)
	})

	t.Run("one_off", func(t *testing.T) {
		b := telego.InlineKeyboardButton{}
		i := InlineKeyboardRows(2, b, b, b)
		require.Len(t, i, 2)
		assert.Len(t, i[0], 2)
		assert.Len(t, i[1], 1)
	})

	t.Run("nil", func(t *testing.T) {
		i := InlineKeyboardRows(0)
		assert.Nil(t, i)
	})
}

func TestKeyboard(t *testing.T) {
	k := Keyboard([]telego.KeyboardButton{}, []telego.KeyboardButton{})
	assert.Len(t, k.Keyboard, 2)
}

func TestKeyboardButton(t *testing.T) {
	k := KeyboardButton(text1)
	assert.Equal(t, text1, k.Text)
}

func TestKeyboardRow(t *testing.T) {
	k := KeyboardRow(telego.KeyboardButton{}, telego.KeyboardButton{})
	assert.Len(t, k, 2)
}

func TestKeyboardGrid(t *testing.T) {
	i := KeyboardGrid([][]telego.KeyboardButton{
		{{}},
		{{}, {}, {}},
	})
	require.Len(t, i.Keyboard, 2)
	assert.Len(t, i.Keyboard[0], 1)
	assert.Len(t, i.Keyboard[1], 3)
}

func TestKeyboardCols(t *testing.T) {
	t.Run("full", func(t *testing.T) {
		b := telego.KeyboardButton{}
		i := KeyboardCols(2, b, b, b, b)
		require.Len(t, i, 2)
		assert.Len(t, i[0], 2)
		assert.Len(t, i[1], 2)
	})

	t.Run("one_off", func(t *testing.T) {
		b := telego.KeyboardButton{}
		i := KeyboardCols(2, b, b, b)
		require.Len(t, i, 2)
		assert.Len(t, i[0], 2)
		assert.Len(t, i[1], 1)
	})

	t.Run("nil", func(t *testing.T) {
		i := KeyboardCols(0)
		assert.Nil(t, i)
	})
}

func TestKeyboardRows(t *testing.T) {
	t.Run("full", func(t *testing.T) {
		b := telego.KeyboardButton{}
		i := KeyboardRows(2, b, b, b, b)
		require.Len(t, i, 2)
		assert.Len(t, i[0], 2)
		assert.Len(t, i[1], 2)
	})

	t.Run("one_off", func(t *testing.T) {
		b := telego.KeyboardButton{}
		i := KeyboardRows(2, b, b, b)
		require.Len(t, i, 2)
		assert.Len(t, i[0], 2)
		assert.Len(t, i[1], 1)
	})

	t.Run("nil", func(t *testing.T) {
		i := KeyboardRows(0)
		assert.Nil(t, i)
	})
}

func TestReplyKeyboardRemove(t *testing.T) {
	r := ReplyKeyboardRemove()
	assert.True(t, r.RemoveKeyboard)
}

func TestWebAppInfo(t *testing.T) {
	w := WebAppInfo(text1)
	assert.Equal(t, text1, w.URL)
}

func TestForceReply(t *testing.T) {
	f := ForceReply()
	assert.True(t, f.ForceReply)
}

func TestMediaAnimation(t *testing.T) {
	m := MediaAnimation(telego.InputFile{File: nr})
	assert.Equal(t, telego.MediaTypeAnimation, m.Type)
	assert.Equal(t, nr, m.Media.File)
}

func TestMediaAudio(t *testing.T) {
	m := MediaAudio(telego.InputFile{File: nr})
	assert.Equal(t, telego.MediaTypeAudio, m.Type)
	assert.Equal(t, nr, m.Media.File)
}

func TestMediaDocument(t *testing.T) {
	m := MediaDocument(telego.InputFile{File: nr})
	assert.Equal(t, telego.MediaTypeDocument, m.Type)
	assert.Equal(t, nr, m.Media.File)
}

func TestMediaLivePhoto(t *testing.T) {
	m := MediaLivePhoto(telego.InputFile{File: nr}, telego.InputFile{File: nr2})
	assert.Equal(t, telego.MediaTypeLivePhoto, m.Type)
	assert.Equal(t, nr, m.Media.File)
	assert.Equal(t, nr2, m.Photo.File)
}

func TestMediaLocation(t *testing.T) {
	m := MediaLocation(latitude, longitude)
	assert.Equal(t, telego.MediaTypeLocation, m.Type)
	assert.InEpsilon(t, latitude, m.Latitude, epsilon)
	assert.InEpsilon(t, longitude, m.Longitude, epsilon)
}

func TestMediaPhoto(t *testing.T) {
	m := MediaPhoto(telego.InputFile{File: nr})
	assert.Equal(t, telego.MediaTypePhoto, m.Type)
	assert.Equal(t, nr, m.Media.File)
}

func TestMediaSticker(t *testing.T) {
	m := MediaSticker(telego.InputFile{File: nr})
	assert.Equal(t, telego.MediaTypeSticker, m.Type)
	assert.Equal(t, nr, m.Media.File)
}

func TestMediaVenue(t *testing.T) {
	m := MediaVenue(latitude, longitude, text1, text2)
	assert.Equal(t, telego.MediaTypeVenue, m.Type)
	assert.InEpsilon(t, latitude, m.Latitude, epsilon)
	assert.InEpsilon(t, longitude, m.Longitude, epsilon)
	assert.Equal(t, text1, m.Title)
	assert.Equal(t, text2, m.Address)
}

func TestMediaVideo(t *testing.T) {
	m := MediaVideo(telego.InputFile{File: nr})
	assert.Equal(t, telego.MediaTypeVideo, m.Type)
	assert.Equal(t, nr, m.Media.File)
}

func TestPollTypeAny(t *testing.T) {
	p := PollTypeAny()
	assert.Empty(t, p.Type)
}

func TestPollTypeQuiz(t *testing.T) {
	p := PollTypeQuiz()
	assert.Equal(t, telego.PollTypeQuiz, p.Type)
}

func TestPollTypeRegular(t *testing.T) {
	p := PollTypeRegular()
	assert.Equal(t, telego.PollTypeRegular, p.Type)
}

func TestResultArticle(t *testing.T) {
	r := ResultArticle(text1, text2, &telego.InputTextMessageContent{})
	assert.Equal(t, telego.ResultTypeArticle, r.Type)
	assert.Equal(t, text1, r.ID)
	assert.Equal(t, text2, r.Title)
	assert.Equal(t, &telego.InputTextMessageContent{}, r.InputMessageContent)
}

func TestResultAudio(t *testing.T) {
	r := ResultAudio(text1, text2, text3)
	assert.Equal(t, telego.ResultTypeAudio, r.Type)
	assert.Equal(t, text1, r.ID)
	assert.Equal(t, text2, r.AudioURL)
	assert.Equal(t, text3, r.Title)
}

func TestResultCachedAudio(t *testing.T) {
	r := ResultCachedAudio(text1, text2)
	assert.Equal(t, telego.ResultTypeAudio, r.Type)
	assert.Equal(t, text1, r.ID)
	assert.Equal(t, text2, r.AudioFileID)
}

func TestResultCachedDocument(t *testing.T) {
	r := ResultCachedDocument(text1, text2, text3)
	assert.Equal(t, telego.ResultTypeDocument, r.Type)
	assert.Equal(t, text1, r.ID)
	assert.Equal(t, text2, r.Title)
	assert.Equal(t, text3, r.DocumentFileID)
}

func TestResultCachedGif(t *testing.T) {
	r := ResultCachedGif(text1, text2)
	assert.Equal(t, telego.ResultTypeGif, r.Type)
	assert.Equal(t, text1, r.ID)
	assert.Equal(t, text2, r.GifFileID)
}

func TestResultCachedMpeg4Gif(t *testing.T) {
	r := ResultCachedMpeg4Gif(text1, text2)
	assert.Equal(t, telego.ResultTypeMpeg4Gif, r.Type)
	assert.Equal(t, text1, r.ID)
	assert.Equal(t, text2, r.Mpeg4FileID)
}

func TestResultCachedPhoto(t *testing.T) {
	r := ResultCachedPhoto(text1, text2)
	assert.Equal(t, telego.ResultTypePhoto, r.Type)
	assert.Equal(t, text1, r.ID)
	assert.Equal(t, text2, r.PhotoFileID)
}

func TestResultCachedSticker(t *testing.T) {
	r := ResultCachedSticker(text1, text2)
	assert.Equal(t, telego.ResultTypeSticker, r.Type)
	assert.Equal(t, text1, r.ID)
	assert.Equal(t, text2, r.StickerFileID)
}

func TestResultCachedVideo(t *testing.T) {
	r := ResultCachedVideo(text1, text2, text3)
	assert.Equal(t, telego.ResultTypeVideo, r.Type)
	assert.Equal(t, text1, r.ID)
	assert.Equal(t, text2, r.VideoFileID)
	assert.Equal(t, text3, r.Title)
}

func TestResultCachedVoice(t *testing.T) {
	r := ResultCachedVoice(text1, text2, text3)
	assert.Equal(t, telego.ResultTypeVoice, r.Type)
	assert.Equal(t, text1, r.ID)
	assert.Equal(t, text2, r.VoiceFileID)
	assert.Equal(t, text3, r.Title)
}

func TestResultContact(t *testing.T) {
	r := ResultContact(text1, text2, text3)
	assert.Equal(t, telego.ResultTypeContact, r.Type)
	assert.Equal(t, text1, r.ID)
	assert.Equal(t, text2, r.PhoneNumber)
	assert.Equal(t, text3, r.FirstName)
}

func TestResultDocument(t *testing.T) {
	r := ResultDocument(text1, text2, text3, text4)
	assert.Equal(t, telego.ResultTypeDocument, r.Type)
	assert.Equal(t, text1, r.ID)
	assert.Equal(t, text2, r.Title)
	assert.Equal(t, text3, r.DocumentURL)
	assert.Equal(t, text4, r.MimeType)
}

func TestResultGame(t *testing.T) {
	r := ResultGame(text1, text2)
	assert.Equal(t, telego.ResultTypeGame, r.Type)
	assert.Equal(t, text1, r.ID)
	assert.Equal(t, text2, r.GameShortName)
}

func TestResultGif(t *testing.T) {
	r := ResultGif(text1, text2, text3)
	assert.Equal(t, telego.ResultTypeGif, r.Type)
	assert.Equal(t, text1, r.ID)
	assert.Equal(t, text2, r.GifURL)
	assert.Equal(t, text3, r.ThumbnailURL)
}

func TestResultLocation(t *testing.T) {
	r := ResultLocation(text1, latitude, longitude, text2)
	assert.Equal(t, telego.ResultTypeLocation, r.Type)
	assert.Equal(t, text1, r.ID)
	assert.InEpsilon(t, latitude, r.Latitude, epsilon)
	assert.InEpsilon(t, longitude, r.Longitude, epsilon)
	assert.Equal(t, text2, r.Title)
}

func TestResultMpeg4Gif(t *testing.T) {
	r := ResultMpeg4Gif(text1, text2, text3)
	assert.Equal(t, telego.ResultTypeMpeg4Gif, r.Type)
	assert.Equal(t, text1, r.ID)
	assert.Equal(t, text2, r.Mpeg4URL)
	assert.Equal(t, text3, r.ThumbnailURL)
}

func TestResultPhoto(t *testing.T) {
	r := ResultPhoto(text1, text2, text3)
	assert.Equal(t, telego.ResultTypePhoto, r.Type)
	assert.Equal(t, text1, r.ID)
	assert.Equal(t, text2, r.PhotoURL)
	assert.Equal(t, text3, r.ThumbnailURL)
}

func TestResultVenue(t *testing.T) {
	r := ResultVenue(text1, latitude, longitude, text2, text3)
	assert.Equal(t, telego.ResultTypeVenue, r.Type)
	assert.Equal(t, text1, r.ID)
	assert.InEpsilon(t, latitude, r.Latitude, epsilon)
	assert.InEpsilon(t, longitude, r.Longitude, epsilon)
	assert.Equal(t, text2, r.Title)
	assert.Equal(t, text3, r.Address)
}

func TestResultVideo(t *testing.T) {
	r := ResultVideo(text1, text2, text3, text4, text5)
	assert.Equal(t, telego.ResultTypeVideo, r.Type)
	assert.Equal(t, text1, r.ID)
	assert.Equal(t, text2, r.VideoURL)
	assert.Equal(t, text3, r.MimeType)
	assert.Equal(t, text4, r.ThumbnailURL)
	assert.Equal(t, text5, r.Title)
}

func TestResultVoice(t *testing.T) {
	r := ResultVoice(text1, text2, text3)
	assert.Equal(t, telego.ResultTypeVoice, r.Type)
	assert.Equal(t, text1, r.ID)
	assert.Equal(t, text2, r.VoiceURL)
	assert.Equal(t, text3, r.Title)
}

func TestScopeAllChatAdministrators(t *testing.T) {
	s := ScopeAllChatAdministrators()
	assert.Equal(t, telego.ScopeTypeAllChatAdministrators, s.Type)
}

func TestScopeAllGroupChats(t *testing.T) {
	s := ScopeAllGroupChats()
	assert.Equal(t, telego.ScopeTypeAllGroupChats, s.Type)
}

func TestScopeAllPrivateChats(t *testing.T) {
	s := ScopeAllPrivateChats()
	assert.Equal(t, telego.ScopeTypeAllPrivateChats, s.Type)
}

func TestScopeChat(t *testing.T) {
	s := ScopeChat(id1)
	assert.Equal(t, telego.ScopeTypeChat, s.Type)
	assert.Equal(t, id1, s.ChatID)
}

func TestScopeChatAdministrators(t *testing.T) {
	s := ScopeChatAdministrators(id1)
	assert.Equal(t, telego.ScopeTypeChatAdministrators, s.Type)
	assert.Equal(t, id1, s.ChatID)
}

func TestScopeChatMember(t *testing.T) {
	s := ScopeChatMember(id1, number2)
	assert.Equal(t, telego.ScopeTypeChatMember, s.Type)
	assert.Equal(t, id1, s.ChatID)
	assert.Equal(t, number2, s.UserID)
}

func TestScopeDefault(t *testing.T) {
	s := ScopeDefault()
	assert.Equal(t, telego.ScopeTypeDefault, s.Type)
}

func TestTextMessage(t *testing.T) {
	m := TextMessage(text1)
	assert.Equal(t, text1, m.MessageText)
}

func TestVenueMessage(t *testing.T) {
	m := VenueMessage(latitude, longitude, text1, text2)
	assert.InEpsilon(t, latitude, m.Latitude, epsilon)
	assert.InEpsilon(t, longitude, m.Longitude, epsilon)
	assert.Equal(t, text1, m.Title)
	assert.Equal(t, text2, m.Address)
}

func TestLocationMessage(t *testing.T) {
	m := LocationMessage(latitude, longitude)
	assert.InEpsilon(t, latitude, m.Latitude, epsilon)
	assert.InEpsilon(t, longitude, m.Longitude, epsilon)
}

func TestContactMessage(t *testing.T) {
	m := ContactMessage(text1, text2)
	assert.Equal(t, text1, m.PhoneNumber)
	assert.Equal(t, text2, m.FirstName)
}

func TestInvoiceMessage(t *testing.T) {
	m := InvoiceMessage(text1, text2, text3, text4, text5, prices...)
	assert.Equal(t, text1, m.Title)
	assert.Equal(t, text2, m.Description)
	assert.Equal(t, text3, m.Payload)
	assert.Equal(t, text4, m.ProviderToken)
	assert.Equal(t, text5, m.Currency)
	assert.Equal(t, prices, m.Prices)
}

func TestLabeledPrice(t *testing.T) {
	l := LabeledPrice(text1, number1)
	assert.Equal(t, text1, l.Label)
	assert.Equal(t, number1, l.Amount)
}

func TestShippingOption(t *testing.T) {
	s := ShippingOption(text1, text2, prices...)
	assert.Equal(t, text1, s.ID)
	assert.Equal(t, text2, s.Title)
	assert.Equal(t, prices, s.Prices)
}

func TestReactionEmoji(t *testing.T) {
	r := ReactionEmoji(text1)
	assert.Equal(t, telego.ReactionEmoji, r.Type)
	assert.Equal(t, text1, r.Emoji)
}

func TestReactionCustomEmoji(t *testing.T) {
	r := ReactionCustomEmoji(text1)
	assert.Equal(t, telego.ReactionCustomEmoji, r.Type)
	assert.Equal(t, text1, r.CustomEmojiID)
}

func TestReactionPaid(t *testing.T) {
	r := ReactionPaid()
	assert.Equal(t, telego.ReactionPaid, r.Type)
}

func TestRichBlockParagraph(t *testing.T) {
	r := RichBlockParagraph(richText1)
	assert.Equal(t, telego.BlockTypeParagraph, r.Type)
}

func TestRichBlockSectionHeading(t *testing.T) {
	r := RichBlockSectionHeading(richText1, number1)
	assert.Equal(t, telego.BlockTypeSectionHeading, r.Type)
}

func TestRichBlockPreformatted(t *testing.T) {
	r := RichBlockPreformatted(richText1)
	assert.Equal(t, telego.BlockTypePreformatted, r.Type)
}

func TestRichBlockFooter(t *testing.T) {
	r := RichBlockFooter(richText1)
	assert.Equal(t, telego.BlockTypeFooter, r.Type)
}

func TestRichBlockDivider(t *testing.T) {
	r := RichBlockDivider()
	assert.Equal(t, telego.BlockTypeDivider, r.Type)
}

func TestRichBlockMathematicalExpression(t *testing.T) {
	r := RichBlockMathematicalExpression(text1)
	assert.Equal(t, telego.BlockTypeMathematicalExpression, r.Type)
}

func TestRichBlockAnchor(t *testing.T) {
	r := RichBlockAnchor(text1)
	assert.Equal(t, telego.BlockTypeAnchor, r.Type)
}

func TestRichBlockList(t *testing.T) {
	r := RichBlockList(RichBlockListItem(RichBlockParagraph(richText1)))
	assert.Equal(t, telego.BlockTypeList, r.Type)
}

func TestRichBlockBlockQuotation(t *testing.T) {
	r := RichBlockBlockQuotation(RichBlockParagraph(richText1))
	assert.Equal(t, telego.BlockTypeBlockQuotation, r.Type)
}

func TestRichBlockPullQuotation(t *testing.T) {
	r := RichBlockPullQuotation(richText1)
	assert.Equal(t, telego.BlockTypePullQuotation, r.Type)
}

func TestRichBlockCollage(t *testing.T) {
	r := RichBlockCollage(RichBlockParagraph(richText1))
	assert.Equal(t, telego.BlockTypeCollage, r.Type)
}

func TestRichBlockSlideshow(t *testing.T) {
	r := RichBlockSlideshow(RichBlockParagraph(richText1))
	assert.Equal(t, telego.BlockTypeSlideshow, r.Type)
}

func TestRichBlockTable(t *testing.T) {
	r := RichBlockTable(RichBlockTableRow(RichBlockTableCell(richText1)))
	assert.Equal(t, telego.BlockTypeTable, r.Type)
}

func TestRichBlockTableRow(t *testing.T) {
	i := RichBlockTableRow(telego.RichBlockTableCell{}, telego.RichBlockTableCell{})
	assert.Len(t, i, 2)
}

func TestRichBlockTableGrid(t *testing.T) {
	i := RichBlockTableGrid([][]telego.RichBlockTableCell{
		{{}},
		{{}, {}, {}},
	})
	require.Len(t, i.Cells, 2)
	assert.Len(t, i.Cells[0], 1)
	assert.Len(t, i.Cells[1], 3)
}

func TestRichBlockTableCols(t *testing.T) {
	t.Run("full", func(t *testing.T) {
		b := telego.RichBlockTableCell{}
		i := RichBlockTableCols(2, b, b, b, b)
		require.Len(t, i, 2)
		assert.Len(t, i[0], 2)
		assert.Len(t, i[1], 2)
	})

	t.Run("one_off", func(t *testing.T) {
		b := telego.RichBlockTableCell{}
		i := RichBlockTableCols(2, b, b, b)
		require.Len(t, i, 2)
		assert.Len(t, i[0], 2)
		assert.Len(t, i[1], 1)
	})

	t.Run("nil", func(t *testing.T) {
		i := RichBlockTableCols(0)
		assert.Nil(t, i)
	})
}

func TestRichBlockTableRows(t *testing.T) {
	t.Run("full", func(t *testing.T) {
		b := telego.RichBlockTableCell{}
		i := RichBlockTableRows(2, b, b, b, b)
		require.Len(t, i, 2)
		assert.Len(t, i[0], 2)
		assert.Len(t, i[1], 2)
	})

	t.Run("one_off", func(t *testing.T) {
		b := telego.RichBlockTableCell{}
		i := RichBlockTableRows(2, b, b, b)
		require.Len(t, i, 2)
		assert.Len(t, i[0], 2)
		assert.Len(t, i[1], 1)
	})

	t.Run("nil", func(t *testing.T) {
		i := RichBlockTableRows(0)
		assert.Nil(t, i)
	})
}

func TestRichBlockTableCell(t *testing.T) {
	r := RichBlockTableCell(richText1)
	assert.Equal(t, richText1, r.Text)
}

func TestRichBlockDetails(t *testing.T) {
	r := RichBlockDetails(richText1, RichBlockParagraph(richText1))
	assert.Equal(t, telego.BlockTypeDetails, r.Type)
}

func TestRichBlockMap(t *testing.T) {
	r := RichBlockMap(telego.Location{Latitude: latitude}, number1, number3, number4)
	assert.Equal(t, telego.BlockTypeMap, r.Type)
}

func TestRichBlockAnimation(t *testing.T) {
	r := RichBlockAnimation(telego.InputMediaAnimation{Type: text1})
	assert.Equal(t, telego.BlockTypeAnimation, r.Type)
}

func TestRichBlockAudio(t *testing.T) {
	r := RichBlockAudio(telego.InputMediaAudio{Type: text1})
	assert.Equal(t, telego.BlockTypeAudio, r.Type)
}

func TestRichBlockPhoto(t *testing.T) {
	r := RichBlockPhoto(telego.InputMediaPhoto{Type: text1})
	assert.Equal(t, telego.BlockTypePhoto, r.Type)
}

func TestRichBlockVideo(t *testing.T) {
	r := RichBlockVideo(telego.InputMediaVideo{Type: text1})
	assert.Equal(t, telego.BlockTypeVideo, r.Type)
}

func TestRichBlockVoiceNote(t *testing.T) {
	r := RichBlockVoiceNote(telego.InputMediaVoiceNote{Type: text1})
	assert.Equal(t, telego.BlockTypeVoiceNote, r.Type)
}

func TestRichBlockThinking(t *testing.T) {
	r := RichBlockThinking(richText1)
	assert.Equal(t, telego.BlockTypeThinking, r.Type)
}

func TestRichBlockCaption(t *testing.T) {
	r := RichBlockCaption(richText1)
	assert.Equal(t, richText1, r.Text)
}

func TestRichBlockListItem(t *testing.T) {
	r := RichBlockListItem(RichBlockParagraph(richText1))
	assert.Equal(t, []telego.InputRichBlock{RichBlockParagraph(richText1)}, r.Blocks)
}

func TestRichTextPlain(t *testing.T) {
	r := RichTextPlain(text1)
	assert.EqualValues(t, &text1, r)
}

func TestRichTextList(t *testing.T) {
	r := RichTextList(richText1, richText2)
	assert.EqualValues(t, &[]telego.RichText{richText1, richText2}, r)
}

func TestRichTextBold(t *testing.T) {
	r := RichTextBold(richText1)
	assert.Equal(t, telego.TextTypeBold, r.Type)
}

func TestRichTextItalic(t *testing.T) {
	r := RichTextItalic(richText1)
	assert.Equal(t, telego.TextTypeItalic, r.Type)
}

func TestRichTextUnderline(t *testing.T) {
	r := RichTextUnderline(richText1)
	assert.Equal(t, telego.TextTypeUnderline, r.Type)
}

func TestRichTextStrikethrough(t *testing.T) {
	r := RichTextStrikethrough(richText1)
	assert.Equal(t, telego.TextTypeStrikethrough, r.Type)
}

func TestRichTextSpoiler(t *testing.T) {
	r := RichTextSpoiler(richText1)
	assert.Equal(t, telego.TextTypeSpoiler, r.Type)
}

func TestRichTextDateTime(t *testing.T) {
	r := RichTextDateTime(richText1, number2, text2)
	assert.Equal(t, telego.TextTypeDateTime, r.Type)
}

func TestRichTextTextMention(t *testing.T) {
	r := RichTextTextMention(richText1, telego.User{})
	assert.Equal(t, telego.TextTypeTextMention, r.Type)
}

func TestRichTextSubscript(t *testing.T) {
	r := RichTextSubscript(richText1)
	assert.Equal(t, telego.TextTypeSubscript, r.Type)
}

func TestRichTextSuperscript(t *testing.T) {
	r := RichTextSuperscript(richText1)
	assert.Equal(t, telego.TextTypeSuperscript, r.Type)
}

func TestRichTextMarked(t *testing.T) {
	r := RichTextMarked(richText1)
	assert.Equal(t, telego.TextTypeMarked, r.Type)
}

func TestRichTextCode(t *testing.T) {
	r := RichTextCode(richText1)
	assert.Equal(t, telego.TextTypeCode, r.Type)
}

func TestRichTextCustomEmoji(t *testing.T) {
	r := RichTextCustomEmoji(text1, text2)
	assert.Equal(t, telego.TextTypeCustomEmoji, r.Type)
}

func TestRichTextMathematicalExpression(t *testing.T) {
	r := RichTextMathematicalExpression(text1)
	assert.Equal(t, telego.TextTypeMathematicalExpression, r.Type)
}

func TestRichTextURL(t *testing.T) {
	r := RichTextURL(richText1, text2)
	assert.Equal(t, telego.TextTypeURL, r.Type)
}

func TestRichTextEmailAddress(t *testing.T) {
	r := RichTextEmailAddress(richText1, text2)
	assert.Equal(t, telego.TextTypeEmailAddress, r.Type)
}

func TestRichTextPhoneNumber(t *testing.T) {
	r := RichTextPhoneNumber(richText1, text2)
	assert.Equal(t, telego.TextTypePhoneNumber, r.Type)
}

func TestRichTextBankCardNumber(t *testing.T) {
	r := RichTextBankCardNumber(richText1, text2)
	assert.Equal(t, telego.TextTypeBankCardNumber, r.Type)
}

func TestRichTextMention(t *testing.T) {
	r := RichTextMention(richText1, text2)
	assert.Equal(t, telego.TextTypeMention, r.Type)
}

func TestRichTextHashtag(t *testing.T) {
	r := RichTextHashtag(richText1, text2)
	assert.Equal(t, telego.TextTypeHashtag, r.Type)
}

func TestRichTextCashtag(t *testing.T) {
	r := RichTextCashtag(richText1, text2)
	assert.Equal(t, telego.TextTypeCashtag, r.Type)
}

func TestRichTextBotCommand(t *testing.T) {
	r := RichTextBotCommand(richText1, text2)
	assert.Equal(t, telego.TextTypeBotCommand, r.Type)
}

func TestRichTextAnchor(t *testing.T) {
	r := RichTextAnchor(text1)
	assert.Equal(t, telego.TextTypeAnchor, r.Type)
}

func TestRichTextAnchorLink(t *testing.T) {
	r := RichTextAnchorLink(richText1, text2)
	assert.Equal(t, telego.TextTypeAnchorLink, r.Type)
}

func TestRichTextReference(t *testing.T) {
	r := RichTextReference(richText1, text2)
	assert.Equal(t, telego.TextTypeReference, r.Type)
}

func TestRichTextReferenceLink(t *testing.T) {
	r := RichTextReferenceLink(richText1, text2)
	assert.Equal(t, telego.TextTypeReferenceLink, r.Type)
}

func TestRichMessage(t *testing.T) {
	m := RichMessage(RichBlockParagraph(richText1))
	assert.Equal(t, []telego.InputRichBlock{RichBlockParagraph(richText1)}, m.Blocks)
}

func TestRichMessageHTML(t *testing.T) {
	m := RichMessageHTML(text1)
	assert.Equal(t, text1, m.HTML)
}

func TestRichMessageMarkdown(t *testing.T) {
	m := RichMessageMarkdown(text1)
	assert.Equal(t, text1, m.Markdown)
}
