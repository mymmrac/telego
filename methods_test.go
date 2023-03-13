//nolint:dupl
package telego

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/mymmrac/telego/telegoapi"
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		updates, err := m.Bot.GetUpdates(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedUpdates, updates)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		updates, err := m.Bot.GetUpdates(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetWebhook(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetWebhook(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.DeleteWebhook(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteWebhook(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		webhookInfo, err := m.Bot.GetWebhookInfo()
		assert.NoError(t, err)
		assert.Equal(t, expectedWebhookInfo, webhookInfo)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		webhookInfo, err := m.Bot.GetWebhookInfo()
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		user, err := m.Bot.GetMe()
		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		user, err := m.Bot.GetMe()
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.LogOut()
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.LogOut()
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.Close()
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.Close()
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendMessage(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendMessage(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.ForwardMessage(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.ForwardMessage(nil)
		assert.Error(t, err)
		assert.Nil(t, message)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		messageID, err := m.Bot.CopyMessage(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessageID, messageID)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		messageID, err := m.Bot.CopyMessage(nil)
		assert.Error(t, err)
		assert.Nil(t, messageID)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendPhoto(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendPhoto(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendAudio(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendAudio(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendDocument(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendDocument(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendVideo(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendVideo(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendAnimation(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendAnimation(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendVoice(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendVoice(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendVideoNote(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendVideoNote(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		messages, err := m.Bot.SendMediaGroup(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessages, messages)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		messages, err := m.Bot.SendMediaGroup(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendLocation(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendLocation(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.EditMessageLiveLocation(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.EditMessageLiveLocation(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.StopMessageLiveLocation(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.StopMessageLiveLocation(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendVenue(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendVenue(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendContact(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendContact(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendPoll(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendPoll(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendDice(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendDice(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SendChatAction(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SendChatAction(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		userProfilePhotos, err := m.Bot.GetUserProfilePhotos(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedUserProfilePhotos, userProfilePhotos)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		userProfilePhotos, err := m.Bot.GetUserProfilePhotos(nil)
		assert.Error(t, err)
		assert.Nil(t, userProfilePhotos)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		file, err := m.Bot.GetFile(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedFile, file)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		file, err := m.Bot.GetFile(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.BanChatMember(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.BanChatMember(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.UnbanChatMember(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.UnbanChatMember(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.RestrictChatMember(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.RestrictChatMember(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.PromoteChatMember(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.PromoteChatMember(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetChatAdministratorCustomTitle(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetChatAdministratorCustomTitle(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.BanChatSenderChat(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.BanChatSenderChat(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.UnbanChatSenderChat(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.UnbanChatSenderChat(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetChatPermissions(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetChatPermissions(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		inviteLink, err := m.Bot.ExportChatInviteLink(nil)
		assert.NoError(t, err)
		assert.Equal(t, &expectedInviteLink, inviteLink)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		inviteLink, err := m.Bot.ExportChatInviteLink(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		chatInviteLink, err := m.Bot.CreateChatInviteLink(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedChatInviteLink, chatInviteLink)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		chatInviteLink, err := m.Bot.CreateChatInviteLink(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		chatInviteLink, err := m.Bot.EditChatInviteLink(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedChatInviteLink, chatInviteLink)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		chatInviteLink, err := m.Bot.EditChatInviteLink(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		chatInviteLink, err := m.Bot.RevokeChatInviteLink(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedChatInviteLink, chatInviteLink)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		chatInviteLink, err := m.Bot.RevokeChatInviteLink(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.ApproveChatJoinRequest(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.ApproveChatJoinRequest(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.DeclineChatJoinRequest(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeclineChatJoinRequest(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetChatPhoto(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetChatPhoto(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.DeleteChatPhoto(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteChatPhoto(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetChatTitle(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetChatTitle(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetChatDescription(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetChatDescription(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.PinChatMessage(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.PinChatMessage(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.UnpinChatMessage(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.UnpinChatMessage(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.UnpinAllChatMessages(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.UnpinAllChatMessages(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.LeaveChat(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.LeaveChat(nil)
		assert.Error(t, err)
	})
}

func TestBot_GetChat(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := newMockedBot(ctrl)

	t.Run("success", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil)

		expectedChat := &Chat{
			ID: 1,
		}
		resp := telegoResponse(t, expectedChat)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		chat, err := m.Bot.GetChat(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedChat, chat)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		chat, err := m.Bot.GetChat(nil)
		assert.Error(t, err)
		assert.Nil(t, chat)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		chatMembers, err := m.Bot.GetChatAdministrators(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedChatMembers, chatMembers)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		chatMembers, err := m.Bot.GetChatAdministrators(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		chatMemberCount, err := m.Bot.GetChatMemberCount(nil)
		assert.NoError(t, err)
		assert.Equal(t, &expectedChatMemberCount, chatMemberCount)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		nt, err := m.Bot.GetChatMemberCount(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		chatMember, err := m.Bot.GetChatMember(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedChatMember, chatMember)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		chatMember, err := m.Bot.GetChatMember(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetChatStickerSet(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetChatStickerSet(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.DeleteChatStickerSet(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteChatStickerSet(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		stickers, err := m.Bot.GetForumTopicIconStickers()
		assert.NoError(t, err)
		assert.Equal(t, expectedStickers, stickers)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		stickers, err := m.Bot.GetForumTopicIconStickers()
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		forumTopic, err := m.Bot.CreateForumTopic(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedForumTopic, forumTopic)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		forumTopic, err := m.Bot.CreateForumTopic(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.EditForumTopic(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.EditForumTopic(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.CloseForumTopic(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.CloseForumTopic(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.ReopenForumTopic(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.ReopenForumTopic(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.DeleteForumTopic(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteForumTopic(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.UnpinAllForumTopicMessages(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.UnpinAllForumTopicMessages(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.EditGeneralForumTopic(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.EditGeneralForumTopic(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.CloseGeneralForumTopic(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.CloseGeneralForumTopic(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.ReopenGeneralForumTopic(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.ReopenGeneralForumTopic(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.HideGeneralForumTopic(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.HideGeneralForumTopic(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.UnhideGeneralForumTopic(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.UnhideGeneralForumTopic(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.AnswerCallbackQuery(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.AnswerCallbackQuery(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetMyCommands(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetMyCommands(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.DeleteMyCommands(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteMyCommands(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		botCommands, err := m.Bot.GetMyCommands(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedBotCommands, botCommands)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		botCommands, err := m.Bot.GetMyCommands(nil)
		assert.Error(t, err)
		assert.Nil(t, botCommands)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetChatMenuButton(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetChatMenuButton(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		menuButton, err := m.Bot.GetChatMenuButton(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMenuButton, menuButton)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		menuButton, err := m.Bot.GetChatMenuButton(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetMyDefaultAdministratorRights(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetMyDefaultAdministratorRights(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		chatAdministratorRights, err := m.Bot.GetMyDefaultAdministratorRights(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedChatAdministratorRights, chatAdministratorRights)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		chatAdministratorRights, err := m.Bot.GetMyDefaultAdministratorRights(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.EditMessageText(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.EditMessageText(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.EditMessageCaption(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.EditMessageCaption(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.EditMessageMedia(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.EditMessageMedia(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.EditMessageReplyMarkup(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.EditMessageReplyMarkup(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		poll, err := m.Bot.StopPoll(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedPoll, poll)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		poll, err := m.Bot.StopPoll(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.DeleteMessage(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteMessage(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendSticker(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendSticker(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		stickerSet, err := m.Bot.GetStickerSet(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedStickerSet, stickerSet)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		stickerSet, err := m.Bot.GetStickerSet(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		stickers, err := m.Bot.GetCustomEmojiStickers(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedStickers, stickers)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		stickers, err := m.Bot.GetCustomEmojiStickers(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		file, err := m.Bot.UploadStickerFile(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedFile, file)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		file, err := m.Bot.UploadStickerFile(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.CreateNewStickerSet(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.CreateNewStickerSet(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.AddStickerToSet(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.AddStickerToSet(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetStickerPositionInSet(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetStickerPositionInSet(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.DeleteStickerFromSet(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.DeleteStickerFromSet(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetStickerSetThumb(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetStickerSetThumb(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.AnswerInlineQuery(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.AnswerInlineQuery(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		sentWebAppMessage, err := m.Bot.AnswerWebAppQuery(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedSentWebAppMessage, sentWebAppMessage)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		sentWebAppMessage, err := m.Bot.AnswerWebAppQuery(nil)
		assert.Error(t, err)
		assert.Nil(t, sentWebAppMessage)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendInvoice(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendInvoice(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		invoiceLink, err := m.Bot.CreateInvoiceLink(nil)
		assert.NoError(t, err)
		assert.Equal(t, &expectedInvoiceLink, invoiceLink)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		invoiceLink, err := m.Bot.CreateInvoiceLink(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.AnswerShippingQuery(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.AnswerShippingQuery(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.AnswerPreCheckoutQuery(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.AnswerPreCheckoutQuery(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(emptyResp, nil)

		err := m.Bot.SetPassportDataErrors(nil)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		err := m.Bot.SetPassportDataErrors(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SendGame(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SendGame(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		message, err := m.Bot.SetGameScore(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, message)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		message, err := m.Bot.SetGameScore(nil)
		assert.Error(t, err)
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
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil)

		gameHighScores, err := m.Bot.GetGameHighScores(nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedGameHighScores, gameHighScores)
	})

	t.Run("error", func(t *testing.T) {
		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest)

		gameHighScores, err := m.Bot.GetGameHighScores(nil)
		assert.Error(t, err)
		assert.Nil(t, gameHighScores)
	})
}

func TestSetWebhookParams_fileParameters(t *testing.T) {
	p := &SetWebhookParams{
		Certificate: &testInputFile,
	}

	assert.Equal(t, map[string]telegoapi.NamedReader{
		"certificate": testNamedReade{},
	}, p.fileParameters())
}

func TestSendPhotoParams_fileParameters(t *testing.T) {
	p := &SendPhotoParams{
		Photo: testInputFile,
	}

	assert.Equal(t, map[string]telegoapi.NamedReader{
		"photo": testNamedReade{},
	}, p.fileParameters())
}

func TestSendAudioParams_fileParameters(t *testing.T) {
	p := &SendAudioParams{
		Audio: testInputFile,
		Thumb: &testInputFile,
	}

	assert.Equal(t, map[string]telegoapi.NamedReader{
		"audio": testNamedReade{},
		"thumb": testNamedReade{},
	}, p.fileParameters())
}

func TestSendDocumentParams_fileParameters(t *testing.T) {
	p := &SendDocumentParams{
		Document: testInputFile,
		Thumb:    &testInputFile,
	}

	assert.Equal(t, map[string]telegoapi.NamedReader{
		"document": testNamedReade{},
		"thumb":    testNamedReade{},
	}, p.fileParameters())
}

func TestSendVideoParams_fileParameters(t *testing.T) {
	p := &SendVideoParams{
		Video: testInputFile,
		Thumb: &testInputFile,
	}

	assert.Equal(t, map[string]telegoapi.NamedReader{
		"video": testNamedReade{},
		"thumb": testNamedReade{},
	}, p.fileParameters())
}

func TestSendAnimationParams_fileParameters(t *testing.T) {
	p := &SendAnimationParams{
		Animation: testInputFile,
		Thumb:     &testInputFile,
	}

	assert.Equal(t, map[string]telegoapi.NamedReader{
		"animation": testNamedReade{},
		"thumb":     testNamedReade{},
	}, p.fileParameters())
}

func TestSendVoiceParams_fileParameters(t *testing.T) {
	p := &SendVoiceParams{
		Voice: testInputFile,
	}

	assert.Equal(t, map[string]telegoapi.NamedReader{
		"voice": testNamedReade{},
	}, p.fileParameters())
}

func TestSendVideoNoteParams_fileParameters(t *testing.T) {
	p := &SendVideoNoteParams{
		VideoNote: testInputFile,
		Thumb:     &testInputFile,
	}

	assert.Equal(t, map[string]telegoapi.NamedReader{
		"video_note": testNamedReade{},
		"thumb":      testNamedReade{},
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

	assert.Equal(t, map[string]telegoapi.NamedReader{
		"test": testNamedReade{},
	}, p.fileParameters())
}

func TestSetChatPhotoParams_fileParameters(t *testing.T) {
	p := &SetChatPhotoParams{
		Photo: testInputFile,
	}

	assert.Equal(t, map[string]telegoapi.NamedReader{
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

	assert.Equal(t, map[string]telegoapi.NamedReader{
		"test": testNamedReade{},
	}, p.fileParameters())
}

func TestSendStickerParams_fileParameters(t *testing.T) {
	p := &SendStickerParams{
		Sticker: testInputFile,
	}

	assert.Equal(t, map[string]telegoapi.NamedReader{
		"sticker": testNamedReade{},
	}, p.fileParameters())
}

func TestUploadStickerFileParams_fileParameters(t *testing.T) {
	p := &UploadStickerFileParams{
		PngSticker: testInputFile,
	}

	assert.Equal(t, map[string]telegoapi.NamedReader{
		"png_sticker": testNamedReade{},
	}, p.fileParameters())
}

func TestCreateNewStickerSetParams_fileParameters(t *testing.T) {
	p := &CreateNewStickerSetParams{
		PngSticker:  &testInputFile,
		TgsSticker:  &testInputFile,
		WebmSticker: &testInputFile,
	}

	assert.Equal(t, map[string]telegoapi.NamedReader{
		"png_sticker":  testNamedReade{},
		"tgs_sticker":  testNamedReade{},
		"webm_sticker": testNamedReade{},
	}, p.fileParameters())
}

func TestAddStickerToSetParams_fileParameters(t *testing.T) {
	p := &AddStickerToSetParams{
		PngSticker:  &testInputFile,
		TgsSticker:  &testInputFile,
		WebmSticker: &testInputFile,
	}

	assert.Equal(t, map[string]telegoapi.NamedReader{
		"png_sticker":  testNamedReade{},
		"tgs_sticker":  testNamedReade{},
		"webm_sticker": testNamedReade{},
	}, p.fileParameters())
}

func TestSetStickerSetThumbnailParams_fileParameters(t *testing.T) {
	p := &SetStickerSetThumbParams{
		Thumb: &testInputFile,
	}

	assert.Equal(t, map[string]telegoapi.NamedReader{
		"thumb": testNamedReade{},
	}, p.fileParameters())
}

func TestMethodsConstants(t *testing.T) {
	tests := [][]string{
		{
			MessageUpdates, EditedMessageUpdates, ChannelPostUpdates, EditedChannelPostUpdates, InlineQueryUpdates,
			ChosenInlineResultUpdates, CallbackQueryUpdates, ShippingQueryUpdates, PreCheckoutQueryUpdates,
			PollUpdates, PollAnswerUpdates, MyChatMemberUpdates, ChatMemberUpdates, ChatJoinRequestUpdates,
		},
		{
			ModeHTML, ModeMarkdown, ModeMarkdownV2,
		},
		{
			ChatActionTyping, ChatActionUploadPhoto, ChatActionRecordVideo, ChatActionUploadVideo,
			ChatActionRecordVoice, ChatActionUploadVoice, ChatActionUploadDocument, ChatActionChooseSticker,
			ChatActionFindLocation, ChatActionRecordVideoNote, ChatActionUploadVideoNote,
		},
	}

	for _, tt := range tests {
		assert.True(t, len(tt) > 0)
		for _, ct := range tt {
			assert.True(t, len(ct) > 0)
		}
	}
}
