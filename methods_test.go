package telego

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	ta "github.com/mymmrac/telego/telegoapi"
)

var testCtx = context.Background()

func TestBot_GetUpdates(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedUpdates := []Update{
			{UpdateID: 1},
			{UpdateID: 2},
		}
		resp := telegoResponse(t, expectedUpdates)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		updates, err := m.Bot.GetUpdates(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedUpdates, updates)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		updates, err := m.Bot.GetUpdates(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, updates)
	})
}

func TestBot_SetWebhook(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetWebhook(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetWebhook(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_DeleteWebhook(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.DeleteWebhook(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteWebhook(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_GetWebhookInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedWebhookInfo := &WebhookInfo{
			URL: "test",
		}
		resp := telegoResponse(t, expectedWebhookInfo)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		webhookInfo, err := m.Bot.GetWebhookInfo(testCtx)
		require.NoError(t, err)
		assert.Equal(t, expectedWebhookInfo, webhookInfo)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		webhookInfo, err := m.Bot.GetWebhookInfo(testCtx)
		require.Error(t, err)
		assert.Nil(t, webhookInfo)
	})
}

func TestBot_GetMe(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedUser := &User{
			ID: 1,
		}
		resp := telegoResponse(t, expectedUser)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		user, err := m.Bot.GetMe(testCtx)
		require.NoError(t, err)
		assert.Equal(t, expectedUser, user)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		user, err := m.Bot.GetMe(testCtx)
		require.Error(t, err)
		assert.Nil(t, user)
	})
}

func TestBot_LogOut(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.LogOut(testCtx)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.LogOut(testCtx)
		require.Error(t, err)
	})
}

func TestBot_Close(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.Close(testCtx)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.Close(testCtx)
		require.Error(t, err)
	})
}

func TestBot_SendMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		resp := telegoResponse(t, expectedMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendMessage(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendMessage(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_ForwardMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		resp := telegoResponse(t, expectedMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.ForwardMessage(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.ForwardMessage(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_ForwardMessages(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedMessageIDs := []MessageID{
			{MessageID: 3},
			{MessageID: 4},
		}
		resp := telegoResponse(t, expectedMessageIDs)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		messageIDs, err := m.Bot.ForwardMessages(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessageIDs, messageIDs)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		messageIDs, err := m.Bot.ForwardMessages(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, messageIDs)
	})
}

func TestBot_CopyMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedMessageID := &MessageID{
			MessageID: 1,
		}
		resp := telegoResponse(t, expectedMessageID)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		messageID, err := m.Bot.CopyMessage(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessageID, messageID)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		messageID, err := m.Bot.CopyMessage(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, messageID)
	})
}

func TestBot_CopyMessages(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedMessageIDs := []MessageID{
			{MessageID: 1},
			{MessageID: 2},
		}
		resp := telegoResponse(t, expectedMessageIDs)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		messageIDs, err := m.Bot.CopyMessages(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessageIDs, messageIDs)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		messageIDs, err := m.Bot.CopyMessages(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, messageIDs)
	})
}

func TestBot_SendPhoto(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		resp := telegoResponse(t, expectedMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendPhoto(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendPhoto(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_SendAudio(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		resp := telegoResponse(t, expectedMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendAudio(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendAudio(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_SendDocument(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		resp := telegoResponse(t, expectedMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendDocument(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendDocument(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_SendVideo(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		resp := telegoResponse(t, expectedMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendVideo(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendVideo(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_SendAnimation(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		resp := telegoResponse(t, expectedMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendAnimation(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendAnimation(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_SendVoice(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		resp := telegoResponse(t, expectedMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendVoice(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendVoice(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_SendVideoNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		resp := telegoResponse(t, expectedMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendVideoNote(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendVideoNote(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_SendPaidMedia(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		resp := telegoResponse(t, expectedMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendPaidMedia(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendPaidMedia(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_SendMediaGroup(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedMessages := []Message{
			{MessageID: 1},
			{MessageID: 2},
		}
		resp := telegoResponse(t, expectedMessages)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		messages, err := m.Bot.SendMediaGroup(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessages, messages)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		messages, err := m.Bot.SendMediaGroup(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, messages)
	})
}

func TestBot_SendLocation(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		resp := telegoResponse(t, expectedMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendLocation(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendLocation(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_SendVenue(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		resp := telegoResponse(t, expectedMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendVenue(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendVenue(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_SendContact(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		resp := telegoResponse(t, expectedMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendContact(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendContact(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_SendPoll(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		resp := telegoResponse(t, expectedMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendPoll(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendPoll(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_SendDice(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		resp := telegoResponse(t, expectedMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendDice(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendDice(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_SendChatAction(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SendChatAction(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SendChatAction(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_SetMessageReaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetMessageReaction(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetMessageReaction(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_GetUserProfilePhotos(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedUserProfilePhotos := &UserProfilePhotos{
			TotalCount: 1,
		}
		resp := telegoResponse(t, expectedUserProfilePhotos)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		userProfilePhotos, err := m.Bot.GetUserProfilePhotos(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedUserProfilePhotos, userProfilePhotos)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		userProfilePhotos, err := m.Bot.GetUserProfilePhotos(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, userProfilePhotos)
	})
}

func TestBot_SetUserEmojiStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetUserEmojiStatus(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetUserEmojiStatus(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_GetFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedFile := &File{
			FileID: "test",
		}
		resp := telegoResponse(t, expectedFile)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		file, err := m.Bot.GetFile(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedFile, file)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		file, err := m.Bot.GetFile(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, file)
	})
}

func TestBot_BanChatMember(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.BanChatMember(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.BanChatMember(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_UnbanChatMember(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.UnbanChatMember(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.UnbanChatMember(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_RestrictChatMember(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.RestrictChatMember(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.RestrictChatMember(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_PromoteChatMember(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.PromoteChatMember(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.PromoteChatMember(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_SetChatAdministratorCustomTitle(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetChatAdministratorCustomTitle(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetChatAdministratorCustomTitle(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_BanChatSenderChat(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.BanChatSenderChat(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.BanChatSenderChat(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_UnbanChatSenderChat(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.UnbanChatSenderChat(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.UnbanChatSenderChat(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_SetChatPermissions(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetChatPermissions(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetChatPermissions(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_ExportChatInviteLink(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedInviteLink := "InviteLink"
		resp := telegoResponse(t, expectedInviteLink)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		inviteLink, err := m.Bot.ExportChatInviteLink(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, &expectedInviteLink, inviteLink)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		inviteLink, err := m.Bot.ExportChatInviteLink(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, inviteLink)
	})
}

func TestBot_CreateChatInviteLink(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedChatInviteLink := &ChatInviteLink{
			InviteLink: "test",
		}
		resp := telegoResponse(t, expectedChatInviteLink)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		chatInviteLink, err := m.Bot.CreateChatInviteLink(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedChatInviteLink, chatInviteLink)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		chatInviteLink, err := m.Bot.CreateChatInviteLink(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, chatInviteLink)
	})
}

func TestBot_EditChatInviteLink(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedChatInviteLink := &ChatInviteLink{
			InviteLink: "test",
		}
		resp := telegoResponse(t, expectedChatInviteLink)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		chatInviteLink, err := m.Bot.EditChatInviteLink(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedChatInviteLink, chatInviteLink)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		chatInviteLink, err := m.Bot.EditChatInviteLink(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, chatInviteLink)
	})
}

func TestBot_CreateChatSubscriptionInviteLink(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedChatInviteLink := &ChatInviteLink{
			InviteLink: "InviteLink",
		}
		resp := telegoResponse(t, expectedChatInviteLink)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		chatInviteLink, err := m.Bot.CreateChatSubscriptionInviteLink(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedChatInviteLink, chatInviteLink)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		chatInviteLink, err := m.Bot.CreateChatSubscriptionInviteLink(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, chatInviteLink)
	})
}

func TestBot_EditChatSubscriptionInviteLink(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedChatInviteLink := &ChatInviteLink{
			InviteLink: "InviteLink",
		}
		resp := telegoResponse(t, expectedChatInviteLink)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		chatInviteLink, err := m.Bot.EditChatSubscriptionInviteLink(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedChatInviteLink, chatInviteLink)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		chatInviteLink, err := m.Bot.EditChatSubscriptionInviteLink(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, chatInviteLink)
	})
}

func TestBot_RevokeChatInviteLink(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedChatInviteLink := &ChatInviteLink{
			InviteLink: "test",
		}
		resp := telegoResponse(t, expectedChatInviteLink)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		chatInviteLink, err := m.Bot.RevokeChatInviteLink(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedChatInviteLink, chatInviteLink)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		chatInviteLink, err := m.Bot.RevokeChatInviteLink(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, chatInviteLink)
	})
}

func TestBot_ApproveChatJoinRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.ApproveChatJoinRequest(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.ApproveChatJoinRequest(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_DeclineChatJoinRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.DeclineChatJoinRequest(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeclineChatJoinRequest(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_SetChatPhoto(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetChatPhoto(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetChatPhoto(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_DeleteChatPhoto(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.DeleteChatPhoto(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteChatPhoto(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_SetChatTitle(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetChatTitle(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetChatTitle(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_SetChatDescription(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetChatDescription(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetChatDescription(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_PinChatMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.PinChatMessage(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.PinChatMessage(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_UnpinChatMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.UnpinChatMessage(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.UnpinChatMessage(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_UnpinAllChatMessages(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.UnpinAllChatMessages(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.UnpinAllChatMessages(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_LeaveChat(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.LeaveChat(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.LeaveChat(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_GetChat(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedChatFullInfo := &ChatFullInfo{
			ID: 1,
		}
		resp := telegoResponse(t, expectedChatFullInfo)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		chatFullInfo, err := m.Bot.GetChat(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedChatFullInfo, chatFullInfo)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		chatFullInfo, err := m.Bot.GetChat(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, chatFullInfo)
	})
}

func TestBot_GetChatAdministrators(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedChatMembers := []ChatMember{
			&ChatMemberOwner{Status: MemberStatusCreator},
			&ChatMemberMember{Status: MemberStatusMember},
		}
		resp := telegoResponse(t, expectedChatMembers)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		chatMembers, err := m.Bot.GetChatAdministrators(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedChatMembers, chatMembers)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		chatMembers, err := m.Bot.GetChatAdministrators(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, chatMembers)
	})
}

func TestBot_GetChatMemberCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedChatMemberCount := 1
		resp := telegoResponse(t, expectedChatMemberCount)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		chatMemberCount, err := m.Bot.GetChatMemberCount(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, &expectedChatMemberCount, chatMemberCount)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		nt, err := m.Bot.GetChatMemberCount(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, nt)
	})
}

func TestBot_GetChatMember(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedChatMember := &ChatMemberOwner{
			Status: MemberStatusCreator,
		}
		resp := telegoResponse(t, expectedChatMember)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		chatMember, err := m.Bot.GetChatMember(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedChatMember, chatMember)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		chatMember, err := m.Bot.GetChatMember(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, chatMember)
	})
}

func TestBot_SetChatStickerSet(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetChatStickerSet(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetChatStickerSet(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_DeleteChatStickerSet(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.DeleteChatStickerSet(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteChatStickerSet(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_GetForumTopicIconStickers(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedStickers := []Sticker{
			{},
		}
		resp := telegoResponse(t, expectedStickers)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		stickers, err := m.Bot.GetForumTopicIconStickers(testCtx)
		require.NoError(t, err)
		assert.Equal(t, expectedStickers, stickers)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		stickers, err := m.Bot.GetForumTopicIconStickers(testCtx)
		require.Error(t, err)
		assert.Nil(t, stickers)
	})
}

func TestBot_CreateForumTopic(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedForumTopic := &ForumTopic{
			MessageThreadID: 1,
		}
		resp := telegoResponse(t, expectedForumTopic)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		forumTopic, err := m.Bot.CreateForumTopic(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedForumTopic, forumTopic)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		forumTopic, err := m.Bot.CreateForumTopic(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, forumTopic)
	})
}

func TestBot_EditForumTopic(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.EditForumTopic(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.EditForumTopic(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_CloseForumTopic(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.CloseForumTopic(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.CloseForumTopic(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_ReopenForumTopic(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.ReopenForumTopic(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.ReopenForumTopic(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_DeleteForumTopic(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.DeleteForumTopic(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteForumTopic(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_UnpinAllForumTopicMessages(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.UnpinAllForumTopicMessages(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.UnpinAllForumTopicMessages(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_EditGeneralForumTopic(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.EditGeneralForumTopic(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.EditGeneralForumTopic(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_CloseGeneralForumTopic(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.CloseGeneralForumTopic(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.CloseGeneralForumTopic(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_ReopenGeneralForumTopic(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.ReopenGeneralForumTopic(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.ReopenGeneralForumTopic(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_HideGeneralForumTopic(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.HideGeneralForumTopic(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.HideGeneralForumTopic(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_UnhideGeneralForumTopic(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.UnhideGeneralForumTopic(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.UnhideGeneralForumTopic(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_UnpinAllGeneralForumTopicMessages(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.UnpinAllGeneralForumTopicMessages(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.UnpinAllGeneralForumTopicMessages(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_AnswerCallbackQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.AnswerCallbackQuery(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.AnswerCallbackQuery(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_GetUserChatBoosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedUserChatBoosts := &UserChatBoosts{}
		resp := telegoResponse(t, expectedUserChatBoosts)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		userChatBoosts, err := m.Bot.GetUserChatBoosts(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedUserChatBoosts, userChatBoosts)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		userChatBoosts, err := m.Bot.GetUserChatBoosts(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, userChatBoosts)
	})
}

func TestBot_GetBusinessConnection(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedBusinessConnection := &BusinessConnection{}
		resp := telegoResponse(t, expectedBusinessConnection)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		businessConnection, err := m.Bot.GetBusinessConnection(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedBusinessConnection, businessConnection)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		businessConnection, err := m.Bot.GetBusinessConnection(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, businessConnection)
	})
}

func TestBot_SetMyCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetMyCommands(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetMyCommands(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_DeleteMyCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.DeleteMyCommands(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteMyCommands(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_GetMyCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedBotCommands := []BotCommand{
			{Command: "test 1"},
			{Command: "test 2"},
		}
		resp := telegoResponse(t, expectedBotCommands)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		botCommands, err := m.Bot.GetMyCommands(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedBotCommands, botCommands)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		botCommands, err := m.Bot.GetMyCommands(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, botCommands)
	})
}

func TestBot_SetMyName(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetMyName(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetMyName(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_GetMyName(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedBotName := &BotName{
			Name: "Name",
		}
		resp := telegoResponse(t, expectedBotName)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		botName, err := m.Bot.GetMyName(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedBotName, botName)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		botName, err := m.Bot.GetMyName(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, botName)
	})
}

func TestBot_SetMyDescription(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetMyDescription(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetMyDescription(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_GetMyDescription(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedBotDescription := &BotDescription{
			Description: "Description",
		}
		resp := telegoResponse(t, expectedBotDescription)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		botDescription, err := m.Bot.GetMyDescription(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedBotDescription, botDescription)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		botDescription, err := m.Bot.GetMyDescription(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, botDescription)
	})
}

func TestBot_SetMyShortDescription(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetMyShortDescription(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetMyShortDescription(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_GetMyShortDescription(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedBotShortDescription := &BotShortDescription{
			ShortDescription: "ShortDescription",
		}
		resp := telegoResponse(t, expectedBotShortDescription)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		botShortDescription, err := m.Bot.GetMyShortDescription(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedBotShortDescription, botShortDescription)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		botShortDescription, err := m.Bot.GetMyShortDescription(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, botShortDescription)
	})
}

func TestBot_SetChatMenuButton(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetChatMenuButton(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetChatMenuButton(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_GetChatMenuButton(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedMenuButton := &MenuButtonCommands{
			Type: ButtonTypeCommands,
		}
		resp := telegoResponse(t, expectedMenuButton)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		menuButton, err := m.Bot.GetChatMenuButton(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMenuButton, menuButton)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		menuButton, err := m.Bot.GetChatMenuButton(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, menuButton)
	})
}

func TestBot_SetMyDefaultAdministratorRights(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetMyDefaultAdministratorRights(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetMyDefaultAdministratorRights(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_GetMyDefaultAdministratorRights(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedChatAdministratorRights := &ChatAdministratorRights{
			IsAnonymous: true,
		}
		resp := telegoResponse(t, expectedChatAdministratorRights)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		chatAdministratorRights, err := m.Bot.GetMyDefaultAdministratorRights(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedChatAdministratorRights, chatAdministratorRights)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		chatAdministratorRights, err := m.Bot.GetMyDefaultAdministratorRights(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, chatAdministratorRights)
	})
}

func TestBot_EditMessageText(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		resp := telegoResponse(t, expectedMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.EditMessageText(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.EditMessageText(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_EditMessageCaption(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		resp := telegoResponse(t, expectedMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.EditMessageCaption(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.EditMessageCaption(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_EditMessageMedia(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		resp := telegoResponse(t, expectedMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.EditMessageMedia(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.EditMessageMedia(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_EditMessageLiveLocation(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		resp := telegoResponse(t, expectedMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.EditMessageLiveLocation(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.EditMessageLiveLocation(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_StopMessageLiveLocation(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		resp := telegoResponse(t, expectedMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.StopMessageLiveLocation(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.StopMessageLiveLocation(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_EditMessageReplyMarkup(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		resp := telegoResponse(t, expectedMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.EditMessageReplyMarkup(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.EditMessageReplyMarkup(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_StopPoll(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedPoll := &Poll{
			ID: "test",
		}
		resp := telegoResponse(t, expectedPoll)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		poll, err := m.Bot.StopPoll(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedPoll, poll)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		poll, err := m.Bot.StopPoll(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, poll)
	})
}

func TestBot_DeleteMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.DeleteMessage(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteMessage(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_DeleteMessages(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.DeleteMessages(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteMessages(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_SendSticker(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		resp := telegoResponse(t, expectedMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendSticker(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendSticker(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_GetStickerSet(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedStickerSet := &StickerSet{
			Name: "test",
		}
		resp := telegoResponse(t, expectedStickerSet)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		stickerSet, err := m.Bot.GetStickerSet(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedStickerSet, stickerSet)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		stickerSet, err := m.Bot.GetStickerSet(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, stickerSet)
	})
}

func TestBot_GetCustomEmojiStickers(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedStickers := []Sticker{
			{FileID: "FileID"},
		}
		resp := telegoResponse(t, expectedStickers)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		stickers, err := m.Bot.GetCustomEmojiStickers(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedStickers, stickers)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		stickers, err := m.Bot.GetCustomEmojiStickers(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, stickers)
	})
}

func TestBot_UploadStickerFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedFile := &File{
			FileID: "test",
		}
		resp := telegoResponse(t, expectedFile)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		file, err := m.Bot.UploadStickerFile(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedFile, file)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		file, err := m.Bot.UploadStickerFile(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, file)
	})
}

func TestBot_CreateNewStickerSet(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.CreateNewStickerSet(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.CreateNewStickerSet(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_AddStickerToSet(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.AddStickerToSet(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.AddStickerToSet(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_SetStickerPositionInSet(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetStickerPositionInSet(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetStickerPositionInSet(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_DeleteStickerFromSet(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.DeleteStickerFromSet(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteStickerFromSet(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_ReplaceStickerInSet(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.ReplaceStickerInSet(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.ReplaceStickerInSet(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_SetStickerEmojiList(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetStickerEmojiList(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetStickerEmojiList(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_SetStickerKeywords(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetStickerKeywords(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetStickerKeywords(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_SetStickerMaskPosition(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetStickerMaskPosition(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetStickerMaskPosition(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_SetStickerSetTitle(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetStickerSetTitle(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetStickerSetTitle(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_SetStickerSetThumbnail(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetStickerSetThumbnail(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetStickerSetThumbnail(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_SetCustomEmojiStickerSetThumbnail(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetCustomEmojiStickerSetThumbnail(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetCustomEmojiStickerSetThumbnail(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_DeleteStickerSet(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.DeleteStickerSet(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteStickerSet(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_GetAvailableGifts(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedGifts := &Gifts{
			Gifts: []Gift{{}},
		}
		resp := telegoResponse(t, expectedGifts)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		gifts, err := m.Bot.GetAvailableGifts(testCtx)
		require.NoError(t, err)
		assert.Equal(t, expectedGifts, gifts)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		gifts, err := m.Bot.GetAvailableGifts(testCtx)
		require.Error(t, err)
		assert.Nil(t, gifts)
	})
}

func TestBot_SendGift(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SendGift(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SendGift(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_VerifyUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.VerifyUser(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.VerifyUser(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_VerifyChat(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.VerifyChat(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.VerifyChat(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_RemoveUserVerification(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.RemoveUserVerification(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.RemoveUserVerification(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_RemoveChatVerification(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.RemoveChatVerification(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.RemoveChatVerification(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_AnswerInlineQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.AnswerInlineQuery(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.AnswerInlineQuery(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_AnswerWebAppQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedSentWebAppMessage := &SentWebAppMessage{
			InlineMessageID: "InlineMessageID",
		}
		resp := telegoResponse(t, expectedSentWebAppMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		sentWebAppMessage, err := m.Bot.AnswerWebAppQuery(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedSentWebAppMessage, sentWebAppMessage)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		sentWebAppMessage, err := m.Bot.AnswerWebAppQuery(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, sentWebAppMessage)
	})
}

func TestBot_SavePreparedInlineMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedPreparedInlineMessage := &PreparedInlineMessage{
			ID: "123",
		}
		resp := telegoResponse(t, expectedPreparedInlineMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		preparedInlineMessage, err := m.Bot.SavePreparedInlineMessage(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedPreparedInlineMessage, preparedInlineMessage)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		preparedInlineMessage, err := m.Bot.SavePreparedInlineMessage(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, preparedInlineMessage)
	})
}

func TestBot_SendInvoice(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		resp := telegoResponse(t, expectedMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendInvoice(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendInvoice(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_CreateInvoiceLink(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedInvoiceLink := "InvoiceLink"
		resp := telegoResponse(t, expectedInvoiceLink)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		invoiceLink, err := m.Bot.CreateInvoiceLink(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, &expectedInvoiceLink, invoiceLink)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		invoiceLink, err := m.Bot.CreateInvoiceLink(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, invoiceLink)
	})
}

func TestBot_AnswerShippingQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.AnswerShippingQuery(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.AnswerShippingQuery(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_AnswerPreCheckoutQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.AnswerPreCheckoutQuery(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.AnswerPreCheckoutQuery(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_GetStarTransactions(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedStarTransactions := &StarTransactions{
			Transactions: []StarTransaction{
				{},
			},
		}
		resp := telegoResponse(t, expectedStarTransactions)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		starTransactions, err := m.Bot.GetStarTransactions(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedStarTransactions, starTransactions)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		starTransactions, err := m.Bot.GetStarTransactions(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, starTransactions)
	})
}

func TestBot_RefundStarPayment(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.RefundStarPayment(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.RefundStarPayment(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_EditUserStarSubscription(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.EditUserStarSubscription(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.EditUserStarSubscription(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_SetPassportDataErrors(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetPassportDataErrors(testCtx, nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetPassportDataErrors(testCtx, nil)
		require.Error(t, err)
	})
}

func TestBot_SendGame(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		resp := telegoResponse(t, expectedMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendGame(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendGame(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_SetGameScore(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		resp := telegoResponse(t, expectedMessage)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SetGameScore(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SetGameScore(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_GetGameHighScores(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedGameHighScores := []GameHighScore{
			{Score: 1},
			{Score: 2},
		}
		resp := telegoResponse(t, expectedGameHighScores)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		gameHighScores, err := m.Bot.GetGameHighScores(testCtx, nil)
		require.NoError(t, err)
		assert.Equal(t, expectedGameHighScores, gameHighScores)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		gameHighScores, err := m.Bot.GetGameHighScores(testCtx, nil)
		require.Error(t, err)
		assert.Nil(t, gameHighScores)
	})
}

func TestSetWebhookParams_fileParameters(t *testing.T) {
	p := &SetWebhookParams{
		Certificate: &testInputFile,
	}

	assert.Equal(t, map[string]ta.NamedReader{
		"certificate": testNamedReade{},
	}, p.fileParameters())
}

func TestSendPhotoParams_fileParameters(t *testing.T) {
	p := &SendPhotoParams{
		Photo: testInputFile,
	}

	assert.Equal(t, map[string]ta.NamedReader{
		"photo": testNamedReade{},
	}, p.fileParameters())
}

func TestSendAudioParams_fileParameters(t *testing.T) {
	p := &SendAudioParams{
		Audio:     testInputFile,
		Thumbnail: &testInputFile,
	}

	assert.Equal(t, map[string]ta.NamedReader{
		"audio":     testNamedReade{},
		"thumbnail": testNamedReade{},
	}, p.fileParameters())
}

func TestSendDocumentParams_fileParameters(t *testing.T) {
	p := &SendDocumentParams{
		Document:  testInputFile,
		Thumbnail: &testInputFile,
	}

	assert.Equal(t, map[string]ta.NamedReader{
		"document":  testNamedReade{},
		"thumbnail": testNamedReade{},
	}, p.fileParameters())
}

func TestSendVideoParams_fileParameters(t *testing.T) {
	p := &SendVideoParams{
		Video:     testInputFile,
		Thumbnail: &testInputFile,
	}

	assert.Equal(t, map[string]ta.NamedReader{
		"video":     testNamedReade{},
		"thumbnail": testNamedReade{},
	}, p.fileParameters())
}

func TestSendAnimationParams_fileParameters(t *testing.T) {
	p := &SendAnimationParams{
		Animation: testInputFile,
		Thumbnail: &testInputFile,
	}

	assert.Equal(t, map[string]ta.NamedReader{
		"animation": testNamedReade{},
		"thumbnail": testNamedReade{},
	}, p.fileParameters())
}

func TestSendVoiceParams_fileParameters(t *testing.T) {
	p := &SendVoiceParams{
		Voice: testInputFile,
	}

	assert.Equal(t, map[string]ta.NamedReader{
		"voice": testNamedReade{},
	}, p.fileParameters())
}

func TestSendVideoNoteParams_fileParameters(t *testing.T) {
	p := &SendVideoNoteParams{
		VideoNote: testInputFile,
		Thumbnail: &testInputFile,
	}

	assert.Equal(t, map[string]ta.NamedReader{
		"video_note": testNamedReade{},
		"thumbnail":  testNamedReade{},
	}, p.fileParameters())
}

func TestSendMediaGroupParams_fileParameters(t *testing.T) {
	p := &SendMediaGroupParams{
		Media: []InputMedia{
			&InputMediaDocument{
				Media:     testInputFile,
				Thumbnail: &testInputFile,
			},
			&InputMediaVideo{
				Media:     testInputFile,
				Thumbnail: &InputFile{File: nil},
			},
		},
	}

	assert.Equal(t, map[string]ta.NamedReader{
		"test": testNamedReade{},
	}, p.fileParameters())
}

func TestSetChatPhotoParams_fileParameters(t *testing.T) {
	p := &SetChatPhotoParams{
		Photo: testInputFile,
	}

	assert.Equal(t, map[string]ta.NamedReader{
		"photo": testNamedReade{},
	}, p.fileParameters())
}

func TestEditMessageMediaParams_fileParameters(t *testing.T) {
	p := &EditMessageMediaParams{
		Media: &InputMediaVideo{
			Media:     testInputFile,
			Thumbnail: &InputFile{File: nil},
		},
	}

	assert.Equal(t, map[string]ta.NamedReader{
		"test": testNamedReade{},
	}, p.fileParameters())
}

func TestSendStickerParams_fileParameters(t *testing.T) {
	p := &SendStickerParams{
		Sticker: testInputFile,
	}

	assert.Equal(t, map[string]ta.NamedReader{
		"sticker": testNamedReade{},
	}, p.fileParameters())
}

func TestUploadStickerFileParams_fileParameters(t *testing.T) {
	p := &UploadStickerFileParams{
		Sticker: testInputFile,
	}

	assert.Equal(t, map[string]ta.NamedReader{
		"sticker": testNamedReade{},
	}, p.fileParameters())
}

func TestCreateNewStickerSetParams_fileParameters(t *testing.T) {
	p := &CreateNewStickerSetParams{
		Stickers: []InputSticker{
			{Sticker: testInputFile},
			{Sticker: testInputFile},
			{Sticker: InputFile{URL: "url"}},
		},
	}

	assert.Equal(t, map[string]ta.NamedReader{
		"test": testNamedReade{},
	}, p.fileParameters())

	assert.Equal(t, map[string]ta.NamedReader{}, (&CreateNewStickerSetParams{}).fileParameters())
}

func TestAddStickerToSetParams_fileParameters(t *testing.T) {
	p := &AddStickerToSetParams{
		Sticker: InputSticker{Sticker: testInputFile},
	}

	assert.Equal(t, map[string]ta.NamedReader{
		"test": testNamedReade{},
	}, p.fileParameters())

	assert.Equal(t, map[string]ta.NamedReader{}, (&AddStickerToSetParams{}).fileParameters())
}

func TestSetStickerSetThumbnailParams_fileParameters(t *testing.T) {
	p := &SetStickerSetThumbnailParams{
		Thumbnail: &testInputFile,
	}

	assert.Equal(t, map[string]ta.NamedReader{
		"thumbnail": testNamedReade{},
	}, p.fileParameters())
}

func TestMethodsConstants(t *testing.T) {
	tests := [][]string{
		{
			MessageUpdates, EditedMessageUpdates, ChannelPostUpdates, EditedChannelPostUpdates,
			BusinessConnectionUpdates, BusinessMessageUpdates, EditedBusinessMessageUpdates,
			DeletedBusinessMessagesUpdates, MessageReactionUpdates, MessageReactionCountUpdates, InlineQueryUpdates,
			ChosenInlineResultUpdates, CallbackQueryUpdates, ShippingQueryUpdates, PreCheckoutQueryUpdates,
			PurchasedPaidMediaUpdates, PollUpdates, PollAnswerUpdates, MyChatMemberUpdates,
			ChatMemberUpdates, ChatJoinRequestUpdates, ChatBoostUpdates, RemovedChatBoostUpdates,
		},
		{
			ModeHTML, ModeMarkdown, ModeMarkdownV2,
		},
		{
			ChatActionTyping, ChatActionUploadPhoto, ChatActionRecordVideo, ChatActionUploadVideo,
			ChatActionRecordVoice, ChatActionUploadVoice, ChatActionUploadDocument, ChatActionChooseSticker,
			ChatActionFindLocation, ChatActionRecordVideoNote, ChatActionUploadVideoNote,
		},
		{
			StickerFormatStatic, StickerFormatAnimated, StickerFormatVideo,
		},
	}

	for _, tt := range tests {
		assert.NotEmpty(t, tt)
		for _, ct := range tt {
			assert.NotEmpty(t, ct)
		}
	}
}
