package telego

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	ta "github.com/mymmrac/telego/telegoapi"
)

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

		updates, err := m.Bot.GetUpdates(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedUpdates, updates)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		updates, err := m.Bot.GetUpdates(t.Context(), nil)
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

		err := m.Bot.SetWebhook(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetWebhook(t.Context(), nil)
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

		err := m.Bot.DeleteWebhook(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteWebhook(t.Context(), nil)
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

		webhookInfo, err := m.Bot.GetWebhookInfo(t.Context())
		require.NoError(t, err)
		assert.Equal(t, expectedWebhookInfo, webhookInfo)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		webhookInfo, err := m.Bot.GetWebhookInfo(t.Context())
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

		user, err := m.Bot.GetMe(t.Context())
		require.NoError(t, err)
		assert.Equal(t, expectedUser, user)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		user, err := m.Bot.GetMe(t.Context())
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

		err := m.Bot.LogOut(t.Context())
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.LogOut(t.Context())
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

		err := m.Bot.Close(t.Context())
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.Close(t.Context())
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

		message, err := m.Bot.SendMessage(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendMessage(t.Context(), nil)
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

		message, err := m.Bot.ForwardMessage(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.ForwardMessage(t.Context(), nil)
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

		messageIDs, err := m.Bot.ForwardMessages(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessageIDs, messageIDs)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		messageIDs, err := m.Bot.ForwardMessages(t.Context(), nil)
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

		messageID, err := m.Bot.CopyMessage(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessageID, messageID)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		messageID, err := m.Bot.CopyMessage(t.Context(), nil)
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

		messageIDs, err := m.Bot.CopyMessages(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessageIDs, messageIDs)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		messageIDs, err := m.Bot.CopyMessages(t.Context(), nil)
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

		message, err := m.Bot.SendPhoto(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendPhoto(t.Context(), nil)
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

		message, err := m.Bot.SendAudio(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendAudio(t.Context(), nil)
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

		message, err := m.Bot.SendDocument(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendDocument(t.Context(), nil)
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

		message, err := m.Bot.SendVideo(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendVideo(t.Context(), nil)
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

		message, err := m.Bot.SendAnimation(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendAnimation(t.Context(), nil)
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

		message, err := m.Bot.SendVoice(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendVoice(t.Context(), nil)
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

		message, err := m.Bot.SendVideoNote(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendVideoNote(t.Context(), nil)
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

		message, err := m.Bot.SendPaidMedia(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendPaidMedia(t.Context(), nil)
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

		messages, err := m.Bot.SendMediaGroup(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessages, messages)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		messages, err := m.Bot.SendMediaGroup(t.Context(), nil)
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

		message, err := m.Bot.SendLocation(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendLocation(t.Context(), nil)
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

		message, err := m.Bot.SendVenue(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendVenue(t.Context(), nil)
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

		message, err := m.Bot.SendContact(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendContact(t.Context(), nil)
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

		message, err := m.Bot.SendPoll(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendPoll(t.Context(), nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_SendChecklist(t *testing.T) {
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

		message, err := m.Bot.SendChecklist(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendChecklist(t.Context(), nil)
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

		message, err := m.Bot.SendDice(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendDice(t.Context(), nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_SendMessageDraft(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SendMessageDraft(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SendMessageDraft(t.Context(), nil)
		require.Error(t, err)
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

		err := m.Bot.SendChatAction(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SendChatAction(t.Context(), nil)
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

		err := m.Bot.SetMessageReaction(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetMessageReaction(t.Context(), nil)
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

		userProfilePhotos, err := m.Bot.GetUserProfilePhotos(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedUserProfilePhotos, userProfilePhotos)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		userProfilePhotos, err := m.Bot.GetUserProfilePhotos(t.Context(), nil)
		require.Error(t, err)
		assert.Nil(t, userProfilePhotos)
	})
}

func TestBot_GetUserProfileAudios(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedUserProfileAudios := &UserProfileAudios{}
		resp := telegoResponse(t, expectedUserProfileAudios)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		userProfileAudios, err := m.Bot.GetUserProfileAudios(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedUserProfileAudios, userProfileAudios)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		userProfileAudios, err := m.Bot.GetUserProfileAudios(t.Context(), nil)
		require.Error(t, err)
		assert.Nil(t, userProfileAudios)
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

		err := m.Bot.SetUserEmojiStatus(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetUserEmojiStatus(t.Context(), nil)
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

		file, err := m.Bot.GetFile(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedFile, file)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		file, err := m.Bot.GetFile(t.Context(), nil)
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

		err := m.Bot.BanChatMember(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.BanChatMember(t.Context(), nil)
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

		err := m.Bot.UnbanChatMember(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.UnbanChatMember(t.Context(), nil)
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

		err := m.Bot.RestrictChatMember(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.RestrictChatMember(t.Context(), nil)
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

		err := m.Bot.PromoteChatMember(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.PromoteChatMember(t.Context(), nil)
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

		err := m.Bot.SetChatAdministratorCustomTitle(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetChatAdministratorCustomTitle(t.Context(), nil)
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

		err := m.Bot.BanChatSenderChat(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.BanChatSenderChat(t.Context(), nil)
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

		err := m.Bot.UnbanChatSenderChat(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.UnbanChatSenderChat(t.Context(), nil)
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

		err := m.Bot.SetChatPermissions(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetChatPermissions(t.Context(), nil)
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

		inviteLink, err := m.Bot.ExportChatInviteLink(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, &expectedInviteLink, inviteLink)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		inviteLink, err := m.Bot.ExportChatInviteLink(t.Context(), nil)
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

		chatInviteLink, err := m.Bot.CreateChatInviteLink(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedChatInviteLink, chatInviteLink)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		chatInviteLink, err := m.Bot.CreateChatInviteLink(t.Context(), nil)
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

		chatInviteLink, err := m.Bot.EditChatInviteLink(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedChatInviteLink, chatInviteLink)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		chatInviteLink, err := m.Bot.EditChatInviteLink(t.Context(), nil)
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

		chatInviteLink, err := m.Bot.CreateChatSubscriptionInviteLink(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedChatInviteLink, chatInviteLink)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		chatInviteLink, err := m.Bot.CreateChatSubscriptionInviteLink(t.Context(), nil)
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

		chatInviteLink, err := m.Bot.EditChatSubscriptionInviteLink(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedChatInviteLink, chatInviteLink)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		chatInviteLink, err := m.Bot.EditChatSubscriptionInviteLink(t.Context(), nil)
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

		chatInviteLink, err := m.Bot.RevokeChatInviteLink(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedChatInviteLink, chatInviteLink)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		chatInviteLink, err := m.Bot.RevokeChatInviteLink(t.Context(), nil)
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

		err := m.Bot.ApproveChatJoinRequest(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.ApproveChatJoinRequest(t.Context(), nil)
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

		err := m.Bot.DeclineChatJoinRequest(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeclineChatJoinRequest(t.Context(), nil)
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

		err := m.Bot.SetChatPhoto(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetChatPhoto(t.Context(), nil)
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

		err := m.Bot.DeleteChatPhoto(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteChatPhoto(t.Context(), nil)
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

		err := m.Bot.SetChatTitle(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetChatTitle(t.Context(), nil)
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

		err := m.Bot.SetChatDescription(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetChatDescription(t.Context(), nil)
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

		err := m.Bot.PinChatMessage(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.PinChatMessage(t.Context(), nil)
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

		err := m.Bot.UnpinChatMessage(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.UnpinChatMessage(t.Context(), nil)
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

		err := m.Bot.UnpinAllChatMessages(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.UnpinAllChatMessages(t.Context(), nil)
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

		err := m.Bot.LeaveChat(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.LeaveChat(t.Context(), nil)
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

		chatFullInfo, err := m.Bot.GetChat(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedChatFullInfo, chatFullInfo)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		chatFullInfo, err := m.Bot.GetChat(t.Context(), nil)
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

		chatMembers, err := m.Bot.GetChatAdministrators(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedChatMembers, chatMembers)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		chatMembers, err := m.Bot.GetChatAdministrators(t.Context(), nil)
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

		chatMemberCount, err := m.Bot.GetChatMemberCount(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, &expectedChatMemberCount, chatMemberCount)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		nt, err := m.Bot.GetChatMemberCount(t.Context(), nil)
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

		chatMember, err := m.Bot.GetChatMember(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedChatMember, chatMember)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		chatMember, err := m.Bot.GetChatMember(t.Context(), nil)
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

		err := m.Bot.SetChatStickerSet(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetChatStickerSet(t.Context(), nil)
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

		err := m.Bot.DeleteChatStickerSet(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteChatStickerSet(t.Context(), nil)
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

		stickers, err := m.Bot.GetForumTopicIconStickers(t.Context())
		require.NoError(t, err)
		assert.Equal(t, expectedStickers, stickers)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		stickers, err := m.Bot.GetForumTopicIconStickers(t.Context())
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

		forumTopic, err := m.Bot.CreateForumTopic(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedForumTopic, forumTopic)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		forumTopic, err := m.Bot.CreateForumTopic(t.Context(), nil)
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

		err := m.Bot.EditForumTopic(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.EditForumTopic(t.Context(), nil)
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

		err := m.Bot.CloseForumTopic(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.CloseForumTopic(t.Context(), nil)
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

		err := m.Bot.ReopenForumTopic(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.ReopenForumTopic(t.Context(), nil)
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

		err := m.Bot.DeleteForumTopic(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteForumTopic(t.Context(), nil)
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

		err := m.Bot.UnpinAllForumTopicMessages(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.UnpinAllForumTopicMessages(t.Context(), nil)
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

		err := m.Bot.EditGeneralForumTopic(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.EditGeneralForumTopic(t.Context(), nil)
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

		err := m.Bot.CloseGeneralForumTopic(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.CloseGeneralForumTopic(t.Context(), nil)
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

		err := m.Bot.ReopenGeneralForumTopic(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.ReopenGeneralForumTopic(t.Context(), nil)
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

		err := m.Bot.HideGeneralForumTopic(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.HideGeneralForumTopic(t.Context(), nil)
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

		err := m.Bot.UnhideGeneralForumTopic(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.UnhideGeneralForumTopic(t.Context(), nil)
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

		err := m.Bot.UnpinAllGeneralForumTopicMessages(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.UnpinAllGeneralForumTopicMessages(t.Context(), nil)
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

		err := m.Bot.AnswerCallbackQuery(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.AnswerCallbackQuery(t.Context(), nil)
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

		userChatBoosts, err := m.Bot.GetUserChatBoosts(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedUserChatBoosts, userChatBoosts)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		userChatBoosts, err := m.Bot.GetUserChatBoosts(t.Context(), nil)
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

		businessConnection, err := m.Bot.GetBusinessConnection(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedBusinessConnection, businessConnection)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		businessConnection, err := m.Bot.GetBusinessConnection(t.Context(), nil)
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

		err := m.Bot.SetMyCommands(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetMyCommands(t.Context(), nil)
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

		err := m.Bot.DeleteMyCommands(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteMyCommands(t.Context(), nil)
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

		botCommands, err := m.Bot.GetMyCommands(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedBotCommands, botCommands)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		botCommands, err := m.Bot.GetMyCommands(t.Context(), nil)
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

		err := m.Bot.SetMyName(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetMyName(t.Context(), nil)
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

		botName, err := m.Bot.GetMyName(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedBotName, botName)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		botName, err := m.Bot.GetMyName(t.Context(), nil)
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

		err := m.Bot.SetMyDescription(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetMyDescription(t.Context(), nil)
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

		botDescription, err := m.Bot.GetMyDescription(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedBotDescription, botDescription)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		botDescription, err := m.Bot.GetMyDescription(t.Context(), nil)
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

		err := m.Bot.SetMyShortDescription(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetMyShortDescription(t.Context(), nil)
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

		botShortDescription, err := m.Bot.GetMyShortDescription(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedBotShortDescription, botShortDescription)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		botShortDescription, err := m.Bot.GetMyShortDescription(t.Context(), nil)
		require.Error(t, err)
		assert.Nil(t, botShortDescription)
	})
}

func TestBot_SetMyProfilePhoto(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetMyProfilePhoto(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetMyProfilePhoto(t.Context(), nil)
		require.Error(t, err)
	})
}

func TestBot_RemoveMyProfilePhoto(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.RemoveMyProfilePhoto(t.Context())
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.RemoveMyProfilePhoto(t.Context())
		require.Error(t, err)
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

		err := m.Bot.SetChatMenuButton(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetChatMenuButton(t.Context(), nil)
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

		menuButton, err := m.Bot.GetChatMenuButton(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMenuButton, menuButton)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		menuButton, err := m.Bot.GetChatMenuButton(t.Context(), nil)
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

		err := m.Bot.SetMyDefaultAdministratorRights(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetMyDefaultAdministratorRights(t.Context(), nil)
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

		chatAdministratorRights, err := m.Bot.GetMyDefaultAdministratorRights(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedChatAdministratorRights, chatAdministratorRights)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		chatAdministratorRights, err := m.Bot.GetMyDefaultAdministratorRights(t.Context(), nil)
		require.Error(t, err)
		assert.Nil(t, chatAdministratorRights)
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
			Gifts: []Gift{
				{},
			},
		}
		resp := telegoResponse(t, expectedGifts)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		gifts, err := m.Bot.GetAvailableGifts(t.Context())
		require.NoError(t, err)
		assert.Equal(t, expectedGifts, gifts)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		gifts, err := m.Bot.GetAvailableGifts(t.Context())
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

		err := m.Bot.SendGift(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SendGift(t.Context(), nil)
		require.Error(t, err)
	})
}

func TestBot_GiftPremiumSubscription(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.GiftPremiumSubscription(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.GiftPremiumSubscription(t.Context(), nil)
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

		err := m.Bot.VerifyUser(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.VerifyUser(t.Context(), nil)
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

		err := m.Bot.VerifyChat(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.VerifyChat(t.Context(), nil)
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

		err := m.Bot.RemoveUserVerification(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.RemoveUserVerification(t.Context(), nil)
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

		err := m.Bot.RemoveChatVerification(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.RemoveChatVerification(t.Context(), nil)
		require.Error(t, err)
	})
}

func TestBot_ReadBusinessMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.ReadBusinessMessage(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.ReadBusinessMessage(t.Context(), nil)
		require.Error(t, err)
	})
}

func TestBot_DeleteBusinessMessages(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.DeleteBusinessMessages(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteBusinessMessages(t.Context(), nil)
		require.Error(t, err)
	})
}

func TestBot_SetBusinessAccountName(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetBusinessAccountName(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetBusinessAccountName(t.Context(), nil)
		require.Error(t, err)
	})
}

func TestBot_SetBusinessAccountUsername(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetBusinessAccountUsername(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetBusinessAccountUsername(t.Context(), nil)
		require.Error(t, err)
	})
}

func TestBot_SetBusinessAccountBio(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetBusinessAccountBio(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetBusinessAccountBio(t.Context(), nil)
		require.Error(t, err)
	})
}

func TestBot_SetBusinessAccountProfilePhoto(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetBusinessAccountProfilePhoto(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetBusinessAccountProfilePhoto(t.Context(), nil)
		require.Error(t, err)
	})
}

func TestBot_RemoveBusinessAccountProfilePhoto(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.RemoveBusinessAccountProfilePhoto(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.RemoveBusinessAccountProfilePhoto(t.Context(), nil)
		require.Error(t, err)
	})
}

func TestBot_SetBusinessAccountGiftSettings(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetBusinessAccountGiftSettings(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetBusinessAccountGiftSettings(t.Context(), nil)
		require.Error(t, err)
	})
}

func TestBot_GetBusinessAccountStarBalance(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedStarAmount := &StarAmount{}
		resp := telegoResponse(t, expectedStarAmount)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		starAmount, err := m.Bot.GetBusinessAccountStarBalance(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedStarAmount, starAmount)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		starAmount, err := m.Bot.GetBusinessAccountStarBalance(t.Context(), nil)
		require.Error(t, err)
		assert.Nil(t, starAmount)
	})
}

func TestBot_TransferBusinessAccountStars(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.TransferBusinessAccountStars(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.TransferBusinessAccountStars(t.Context(), nil)
		require.Error(t, err)
	})
}

func TestBot_GetBusinessAccountGifts(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedOwnedGifts := &OwnedGifts{
			TotalCount: 1,
		}
		resp := telegoResponse(t, expectedOwnedGifts)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		ownedGifts, err := m.Bot.GetBusinessAccountGifts(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedOwnedGifts, ownedGifts)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		ownedGifts, err := m.Bot.GetBusinessAccountGifts(t.Context(), nil)
		require.Error(t, err)
		assert.Nil(t, ownedGifts)
	})
}

func TestBot_GetUserGifts(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedOwnedGifts := &OwnedGifts{
			TotalCount: 1,
		}
		resp := telegoResponse(t, expectedOwnedGifts)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		ownedGifts, err := m.Bot.GetUserGifts(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedOwnedGifts, ownedGifts)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		ownedGifts, err := m.Bot.GetUserGifts(t.Context(), nil)
		require.Error(t, err)
		assert.Nil(t, ownedGifts)
	})
}

func TestBot_GetChatGifts(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedOwnedGifts := &OwnedGifts{
			TotalCount: 1,
		}
		resp := telegoResponse(t, expectedOwnedGifts)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		ownedGifts, err := m.Bot.GetChatGifts(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedOwnedGifts, ownedGifts)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		ownedGifts, err := m.Bot.GetChatGifts(t.Context(), nil)
		require.Error(t, err)
		assert.Nil(t, ownedGifts)
	})
}

func TestBot_ConvertGiftToStars(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.ConvertGiftToStars(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.ConvertGiftToStars(t.Context(), nil)
		require.Error(t, err)
	})
}

func TestBot_UpgradeGift(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.UpgradeGift(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.UpgradeGift(t.Context(), nil)
		require.Error(t, err)
	})
}

func TestBot_TransferGift(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.TransferGift(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.TransferGift(t.Context(), nil)
		require.Error(t, err)
	})
}

func TestBot_PostStory(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedStory := &Story{}
		resp := telegoResponse(t, expectedStory)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		story, err := m.Bot.PostStory(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedStory, story)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		story, err := m.Bot.PostStory(t.Context(), nil)
		require.Error(t, err)
		assert.Nil(t, story)
	})
}

func TestBot_RepostStory(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedStory := &Story{
			ID: 1,
		}
		resp := telegoResponse(t, expectedStory)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		story, err := m.Bot.RepostStory(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedStory, story)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		story, err := m.Bot.RepostStory(t.Context(), nil)
		require.Error(t, err)
		assert.Nil(t, story)
	})
}

func TestBot_EditStory(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedStory := &Story{}
		resp := telegoResponse(t, expectedStory)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		story, err := m.Bot.EditStory(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedStory, story)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		story, err := m.Bot.EditStory(t.Context(), nil)
		require.Error(t, err)
		assert.Nil(t, story)
	})
}

func TestBot_DeleteStory(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.DeleteStory(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteStory(t.Context(), nil)
		require.Error(t, err)
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

		message, err := m.Bot.EditMessageText(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.EditMessageText(t.Context(), nil)
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

		message, err := m.Bot.EditMessageCaption(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.EditMessageCaption(t.Context(), nil)
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

		message, err := m.Bot.EditMessageMedia(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.EditMessageMedia(t.Context(), nil)
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

		message, err := m.Bot.EditMessageLiveLocation(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.EditMessageLiveLocation(t.Context(), nil)
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

		message, err := m.Bot.StopMessageLiveLocation(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.StopMessageLiveLocation(t.Context(), nil)
		require.Error(t, err)
		assert.Nil(t, message)
	})
}

func TestBot_EditMessageChecklist(t *testing.T) {
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

		message, err := m.Bot.EditMessageChecklist(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.EditMessageChecklist(t.Context(), nil)
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

		message, err := m.Bot.EditMessageReplyMarkup(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.EditMessageReplyMarkup(t.Context(), nil)
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

		poll, err := m.Bot.StopPoll(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedPoll, poll)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		poll, err := m.Bot.StopPoll(t.Context(), nil)
		require.Error(t, err)
		assert.Nil(t, poll)
	})
}

func TestBot_ApproveSuggestedPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.ApproveSuggestedPost(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.ApproveSuggestedPost(t.Context(), nil)
		require.Error(t, err)
	})
}

func TestBot_DeclineSuggestedPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.DeclineSuggestedPost(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeclineSuggestedPost(t.Context(), nil)
		require.Error(t, err)
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

		err := m.Bot.DeleteMessage(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteMessage(t.Context(), nil)
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

		err := m.Bot.DeleteMessages(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteMessages(t.Context(), nil)
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

		message, err := m.Bot.SendSticker(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendSticker(t.Context(), nil)
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

		expectedStickerSet := &StickerSet{}
		resp := telegoResponse(t, expectedStickerSet)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		stickerSet, err := m.Bot.GetStickerSet(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedStickerSet, stickerSet)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		stickerSet, err := m.Bot.GetStickerSet(t.Context(), nil)
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
			{Type: StickerTypeRegular},
		}
		resp := telegoResponse(t, expectedStickers)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		stickers, err := m.Bot.GetCustomEmojiStickers(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedStickers, stickers)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		stickers, err := m.Bot.GetCustomEmojiStickers(t.Context(), nil)
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

		expectedFile := &File{}
		resp := telegoResponse(t, expectedFile)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		file, err := m.Bot.UploadStickerFile(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedFile, file)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		file, err := m.Bot.UploadStickerFile(t.Context(), nil)
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

		err := m.Bot.CreateNewStickerSet(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.CreateNewStickerSet(t.Context(), nil)
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

		err := m.Bot.AddStickerToSet(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.AddStickerToSet(t.Context(), nil)
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

		err := m.Bot.SetStickerPositionInSet(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetStickerPositionInSet(t.Context(), nil)
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

		err := m.Bot.DeleteStickerFromSet(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteStickerFromSet(t.Context(), nil)
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

		err := m.Bot.ReplaceStickerInSet(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.ReplaceStickerInSet(t.Context(), nil)
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

		err := m.Bot.SetStickerEmojiList(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetStickerEmojiList(t.Context(), nil)
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

		err := m.Bot.SetStickerKeywords(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetStickerKeywords(t.Context(), nil)
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

		err := m.Bot.SetStickerMaskPosition(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetStickerMaskPosition(t.Context(), nil)
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

		err := m.Bot.SetStickerSetTitle(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetStickerSetTitle(t.Context(), nil)
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

		err := m.Bot.SetStickerSetThumbnail(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetStickerSetThumbnail(t.Context(), nil)
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

		err := m.Bot.SetCustomEmojiStickerSetThumbnail(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetCustomEmojiStickerSetThumbnail(t.Context(), nil)
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

		err := m.Bot.DeleteStickerSet(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteStickerSet(t.Context(), nil)
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

		err := m.Bot.AnswerInlineQuery(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.AnswerInlineQuery(t.Context(), nil)
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

		sentWebAppMessage, err := m.Bot.AnswerWebAppQuery(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedSentWebAppMessage, sentWebAppMessage)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		sentWebAppMessage, err := m.Bot.AnswerWebAppQuery(t.Context(), nil)
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

		preparedInlineMessage, err := m.Bot.SavePreparedInlineMessage(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedPreparedInlineMessage, preparedInlineMessage)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		preparedInlineMessage, err := m.Bot.SavePreparedInlineMessage(t.Context(), nil)
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

		message, err := m.Bot.SendInvoice(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendInvoice(t.Context(), nil)
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

		invoiceLink, err := m.Bot.CreateInvoiceLink(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, &expectedInvoiceLink, invoiceLink)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		invoiceLink, err := m.Bot.CreateInvoiceLink(t.Context(), nil)
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

		err := m.Bot.AnswerShippingQuery(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.AnswerShippingQuery(t.Context(), nil)
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

		err := m.Bot.AnswerPreCheckoutQuery(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.AnswerPreCheckoutQuery(t.Context(), nil)
		require.Error(t, err)
	})
}

func TestBot_GetMyStarBalance(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedStarAmount := &StarAmount{}
		resp := telegoResponse(t, expectedStarAmount)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil)

		starAmount, err := m.Bot.GetMyStarBalance(t.Context())
		require.NoError(t, err)
		assert.Equal(t, expectedStarAmount, starAmount)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		starAmount, err := m.Bot.GetMyStarBalance(t.Context())
		require.Error(t, err)
		assert.Nil(t, starAmount)
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

		starTransactions, err := m.Bot.GetStarTransactions(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedStarTransactions, starTransactions)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		starTransactions, err := m.Bot.GetStarTransactions(t.Context(), nil)
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

		err := m.Bot.RefundStarPayment(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.RefundStarPayment(t.Context(), nil)
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

		err := m.Bot.EditUserStarSubscription(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.EditUserStarSubscription(t.Context(), nil)
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

		err := m.Bot.SetPassportDataErrors(t.Context(), nil)
		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetPassportDataErrors(t.Context(), nil)
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

		message, err := m.Bot.SendGame(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendGame(t.Context(), nil)
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

		message, err := m.Bot.SetGameScore(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SetGameScore(t.Context(), nil)
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

		gameHighScores, err := m.Bot.GetGameHighScores(t.Context(), nil)
		require.NoError(t, err)
		assert.Equal(t, expectedGameHighScores, gameHighScores)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		gameHighScores, err := m.Bot.GetGameHighScores(t.Context(), nil)
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
