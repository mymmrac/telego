package telegoutil

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegoapi"
)

var errSettingsTest = errors.New("test error")

type settingsTestResponse struct {
	result string
	err    error
}

type settingsTestCaller struct {
	t         *testing.T
	responses []settingsTestResponse
	calls     []string
}

func (c *settingsTestCaller) Call(
	_ context.Context, url string, _ *telegoapi.RequestData,
) (*telegoapi.Response, error) {
	c.t.Helper()
	c.calls = append(c.calls, url[strings.LastIndex(url, "/")+1:])

	if len(c.responses) == 0 {
		c.t.Fatalf("unexpected API call to %q", url)
	}

	response := c.responses[0]
	c.responses = c.responses[1:]
	if response.err != nil {
		return nil, response.err
	}

	return &telegoapi.Response{Ok: true, Result: []byte(response.result)}, nil
}

func newSettingsTestBot(t *testing.T, responses ...settingsTestResponse) (*telego.Bot, *settingsTestCaller) {
	t.Helper()

	caller := &settingsTestCaller{t: t, responses: responses}
	bot, err := telego.NewBot(
		"1234567890:aaaabbbbaaaabbbbaaaabbbbaaaabbbbccc",
		telego.WithAPICaller(caller),
		telego.WithDiscardLogger(),
	)
	require.NoError(t, err)

	return bot, caller
}

func assertSettingsCalls(t *testing.T, caller *settingsTestCaller, want []string) {
	t.Helper()

	assert.Equal(t, want, caller.calls)
	assert.Empty(t, caller.responses)
}

func TestSetMyName(t *testing.T) {
	tests := []struct {
		name      string
		responses []settingsTestResponse
		wantCalls []string
		wantErr   error
	}{
		{
			name:      "unchanged",
			responses: []settingsTestResponse{{result: `{"name":"same"}`}},
			wantCalls: []string{"getMyName"},
		},
		{
			name:      "changed",
			responses: []settingsTestResponse{{result: `{"name":"old"}`}, {}},
			wantCalls: []string{"getMyName", "setMyName"},
		},
		{
			name:      "get_error",
			responses: []settingsTestResponse{{err: errSettingsTest}},
			wantCalls: []string{"getMyName"},
			wantErr:   errSettingsTest,
		},
		{
			name:      "set_error",
			responses: []settingsTestResponse{{result: `{"name":"old"}`}, {err: errSettingsTest}},
			wantCalls: []string{"getMyName", "setMyName"},
			wantErr:   errSettingsTest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bot, caller := newSettingsTestBot(t, tt.responses...)
			err := SetMyName(t.Context(), bot, &telego.SetMyNameParams{Name: "same"})

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
			assertSettingsCalls(t, caller, tt.wantCalls)
		})
	}
}

func TestSetMyDescription(t *testing.T) {
	tests := []struct {
		name      string
		responses []settingsTestResponse
		wantCalls []string
		wantErr   error
	}{
		{
			name:      "unchanged",
			responses: []settingsTestResponse{{result: `{"description":"same"}`}},
			wantCalls: []string{"getMyDescription"},
		},
		{
			name:      "changed",
			responses: []settingsTestResponse{{result: `{"description":"old"}`}, {}},
			wantCalls: []string{"getMyDescription", "setMyDescription"},
		},
		{
			name:      "get_error",
			responses: []settingsTestResponse{{err: errSettingsTest}},
			wantCalls: []string{"getMyDescription"},
			wantErr:   errSettingsTest,
		},
		{
			name:      "set_error",
			responses: []settingsTestResponse{{result: `{"description":"old"}`}, {err: errSettingsTest}},
			wantCalls: []string{"getMyDescription", "setMyDescription"},
			wantErr:   errSettingsTest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bot, caller := newSettingsTestBot(t, tt.responses...)
			err := SetMyDescription(t.Context(), bot, &telego.SetMyDescriptionParams{Description: "same"})

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
			assertSettingsCalls(t, caller, tt.wantCalls)
		})
	}
}

func TestSetMyShortDescription(t *testing.T) {
	tests := []struct {
		name      string
		responses []settingsTestResponse
		wantCalls []string
		wantErr   error
	}{
		{
			name:      "unchanged",
			responses: []settingsTestResponse{{result: `{"short_description":"same"}`}},
			wantCalls: []string{"getMyShortDescription"},
		},
		{
			name:      "changed",
			responses: []settingsTestResponse{{result: `{"short_description":"old"}`}, {}},
			wantCalls: []string{"getMyShortDescription", "setMyShortDescription"},
		},
		{
			name:      "get_error",
			responses: []settingsTestResponse{{err: errSettingsTest}},
			wantCalls: []string{"getMyShortDescription"},
			wantErr:   errSettingsTest,
		},
		{
			name:      "set_error",
			responses: []settingsTestResponse{{result: `{"short_description":"old"}`}, {err: errSettingsTest}},
			wantCalls: []string{"getMyShortDescription", "setMyShortDescription"},
			wantErr:   errSettingsTest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bot, caller := newSettingsTestBot(t, tt.responses...)
			err := SetMyShortDescription(t.Context(), bot, &telego.SetMyShortDescriptionParams{
				ShortDescription: "same",
			})

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
			assertSettingsCalls(t, caller, tt.wantCalls)
		})
	}
}

func TestSetMyCommands(t *testing.T) {
	command := telego.BotCommand{Command: "start", Description: "Start"}
	params := &telego.SetMyCommandsParams{Commands: []telego.BotCommand{command}}

	tests := []struct {
		name      string
		params    *telego.SetMyCommandsParams
		responses []settingsTestResponse
		wantCalls []string
		wantErr   error
	}{
		{
			name:      "unchanged",
			params:    params,
			responses: []settingsTestResponse{{result: `[{"command":"start","description":"Start"}]`}},
			wantCalls: []string{"getMyCommands"},
		},
		{
			name:      "changed",
			params:    params,
			responses: []settingsTestResponse{{result: `[]`}, {}},
			wantCalls: []string{"getMyCommands", "setMyCommands"},
		},
		{
			name:      "delete_empty_commands",
			params:    &telego.SetMyCommandsParams{},
			responses: []settingsTestResponse{{result: `[{"command":"start","description":"Start"}]`}, {}},
			wantCalls: []string{"getMyCommands", "deleteMyCommands"},
		},
		{
			name:      "empty_commands_unchanged",
			params:    &telego.SetMyCommandsParams{},
			responses: []settingsTestResponse{{result: `[]`}},
			wantCalls: []string{"getMyCommands"},
		},
		{
			name:      "get_error",
			params:    params,
			responses: []settingsTestResponse{{err: errSettingsTest}},
			wantCalls: []string{"getMyCommands"},
			wantErr:   errSettingsTest,
		},
		{
			name:      "set_error",
			params:    params,
			responses: []settingsTestResponse{{result: `[]`}, {err: errSettingsTest}},
			wantCalls: []string{"getMyCommands", "setMyCommands"},
			wantErr:   errSettingsTest,
		},
		{
			name:   "delete_error",
			params: &telego.SetMyCommandsParams{},
			responses: []settingsTestResponse{
				{result: `[{"command":"start","description":"Start"}]`},
				{err: errSettingsTest},
			},
			wantCalls: []string{"getMyCommands", "deleteMyCommands"},
			wantErr:   errSettingsTest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bot, caller := newSettingsTestBot(t, tt.responses...)
			err := SetMyCommands(t.Context(), bot, tt.params)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
			assertSettingsCalls(t, caller, tt.wantCalls)
		})
	}
}

func TestSetMyDefaultAdministratorRights(t *testing.T) {
	rights := &telego.ChatAdministratorRights{CanDeleteMessages: true}

	tests := []struct {
		name      string
		params    *telego.SetMyDefaultAdministratorRightsParams
		responses []settingsTestResponse
		wantCalls []string
		wantErr   error
	}{
		{
			name:      "unchanged_non_nil_rights",
			params:    &telego.SetMyDefaultAdministratorRightsParams{Rights: rights},
			responses: []settingsTestResponse{{result: `{"can_delete_messages":true}`}},
			wantCalls: []string{"getMyDefaultAdministratorRights"},
		},
		{
			name:      "changed_rights",
			params:    &telego.SetMyDefaultAdministratorRightsParams{Rights: rights},
			responses: []settingsTestResponse{{result: `{}`}, {}},
			wantCalls: []string{"getMyDefaultAdministratorRights", "setMyDefaultAdministratorRights"},
		},
		{
			name:      "unchanged_nil_rights",
			params:    &telego.SetMyDefaultAdministratorRightsParams{},
			responses: []settingsTestResponse{{result: `null`}},
			wantCalls: []string{"getMyDefaultAdministratorRights"},
		},
		{
			name:      "get_error",
			params:    &telego.SetMyDefaultAdministratorRightsParams{},
			responses: []settingsTestResponse{{err: errSettingsTest}},
			wantCalls: []string{"getMyDefaultAdministratorRights"},
			wantErr:   errSettingsTest,
		},
		{
			name:      "set_error",
			params:    &telego.SetMyDefaultAdministratorRightsParams{Rights: rights},
			responses: []settingsTestResponse{{result: `{}`}, {err: errSettingsTest}},
			wantCalls: []string{"getMyDefaultAdministratorRights", "setMyDefaultAdministratorRights"},
			wantErr:   errSettingsTest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bot, caller := newSettingsTestBot(t, tt.responses...)
			err := SetMyDefaultAdministratorRights(t.Context(), bot, tt.params)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
			assertSettingsCalls(t, caller, tt.wantCalls)
		})
	}
}

func TestIsMySettingsEqual(t *testing.T) {
	trueValue := true
	falseValue := false
	denominator := 10
	autoMode := "auto"

	tests := []struct {
		name   string
		me     *telego.BotFullInfo
		params *telego.SetMySettingsParams
		want   bool
	}{
		{
			name:   "empty_settings",
			me:     &telego.BotFullInfo{},
			params: &telego.SetMySettingsParams{},
			want:   true,
		},
		{
			name:   "nil_settings_are_ignored",
			me:     &telego.BotFullInfo{SupportsInlineQueries: true},
			params: &telego.SetMySettingsParams{InlineQueryPlaceholder: ""},
			want:   true,
		},
		{
			name:   "supports_inline_queries_match",
			me:     &telego.BotFullInfo{SupportsInlineQueries: true},
			params: &telego.SetMySettingsParams{SupportsInlineQueries: &trueValue},
			want:   true,
		},
		{
			name:   "supports_inline_queries_differ",
			me:     &telego.BotFullInfo{},
			params: &telego.SetMySettingsParams{SupportsInlineQueries: &trueValue},
		},
		{
			name:   "inline_query_placeholder_match",
			me:     &telego.BotFullInfo{InlineQueryPlaceholder: "Search"},
			params: &telego.SetMySettingsParams{InlineQueryPlaceholder: "Search"},
			want:   true,
		},
		{
			name:   "inline_query_placeholder_differs",
			me:     &telego.BotFullInfo{InlineQueryPlaceholder: "Search"},
			params: &telego.SetMySettingsParams{InlineQueryPlaceholder: "Find"},
		},
		{
			name:   "needs_inline_query_location_differs",
			me:     &telego.BotFullInfo{},
			params: &telego.SetMySettingsParams{NeedsInlineQueryLocation: &trueValue},
		},
		{
			name:   "can_join_groups_differs",
			me:     &telego.BotFullInfo{},
			params: &telego.SetMySettingsParams{CanJoinGroups: &trueValue},
		},
		{
			name:   "can_read_all_group_messages_differs",
			me:     &telego.BotFullInfo{},
			params: &telego.SetMySettingsParams{CanReadAllGroupMessages: &trueValue},
		},
		{
			name:   "can_connect_to_business_differs",
			me:     &telego.BotFullInfo{},
			params: &telego.SetMySettingsParams{CanConnectToBusiness: &trueValue},
		},
		{
			name:   "can_manage_bots_differs",
			me:     &telego.BotFullInfo{},
			params: &telego.SetMySettingsParams{CanManageBots: &trueValue},
		},
		{
			name:   "supports_guest_queries_differs",
			me:     &telego.BotFullInfo{},
			params: &telego.SetMySettingsParams{SupportsGuestQueries: &trueValue},
		},
		{
			name:   "supports_join_request_queries_differs",
			me:     &telego.BotFullInfo{},
			params: &telego.SetMySettingsParams{SupportsJoinRequestQueries: &trueValue},
		},
		{
			name:   "has_topics_enabled_differs",
			me:     &telego.BotFullInfo{},
			params: &telego.SetMySettingsParams{HasTopicsEnabled: &trueValue},
		},
		{
			name:   "allows_users_to_create_topics_differs",
			me:     &telego.BotFullInfo{},
			params: &telego.SetMySettingsParams{AllowsUsersToCreateTopics: &trueValue},
		},
		{
			name:   "has_topic_sidebar_differs",
			me:     &telego.BotFullInfo{},
			params: &telego.SetMySettingsParams{HasTopicSidebar: &trueValue},
		},
		{
			name:   "unknown_inline_feedback_share_denominator",
			me:     &telego.BotFullInfo{},
			params: &telego.SetMySettingsParams{InlineFeedbackShareDenominator: &denominator},
		},
		{
			name:   "unknown_can_communicate_with_bots",
			me:     &telego.BotFullInfo{},
			params: &telego.SetMySettingsParams{CanCommunicateWithBots: &falseValue},
		},
		{
			name:   "unknown_callback_query_mode",
			me:     &telego.BotFullInfo{},
			params: &telego.SetMySettingsParams{CallbackQueryMode: &autoMode},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsMySettingsEqual(tt.me, tt.params))
		})
	}
}

func TestSetMySettings(t *testing.T) {
	trueValue := true

	tests := []struct {
		name      string
		params    *telego.SetMySettingsParams
		responses []settingsTestResponse
		wantCalls []string
		wantErr   error
	}{
		{
			name:      "unchanged",
			params:    &telego.SetMySettingsParams{},
			responses: []settingsTestResponse{{result: `{}`}},
			wantCalls: []string{"getMe"},
		},
		{
			name:      "changed",
			params:    &telego.SetMySettingsParams{SupportsInlineQueries: &trueValue},
			responses: []settingsTestResponse{{result: `{}`}, {}},
			wantCalls: []string{"getMe", "setMySettings"},
		},
		{
			name:      "unknown_setting_is_always_set",
			params:    &telego.SetMySettingsParams{InlineFeedbackShareDenominator: new(int)},
			responses: []settingsTestResponse{{result: `{}`}, {}},
			wantCalls: []string{"getMe", "setMySettings"},
		},
		{
			name:      "get_error",
			params:    &telego.SetMySettingsParams{},
			responses: []settingsTestResponse{{err: errSettingsTest}},
			wantCalls: []string{"getMe"},
			wantErr:   errSettingsTest,
		},
		{
			name:      "set_error",
			params:    &telego.SetMySettingsParams{SupportsInlineQueries: &trueValue},
			responses: []settingsTestResponse{{result: `{}`}, {err: errSettingsTest}},
			wantCalls: []string{"getMe", "setMySettings"},
			wantErr:   errSettingsTest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bot, caller := newSettingsTestBot(t, tt.responses...)
			err := SetMySettings(t.Context(), bot, tt.params)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
			assertSettingsCalls(t, caller, tt.wantCalls)
		})
	}
}

func TestSetMyProfilePhoto(t *testing.T) {
	params := &telego.SetMyProfilePhotoParams{
		Photo: &telego.InputProfilePhotoStatic{
			Type:  telego.PhotoTypeStatic,
			Photo: telego.InputFile{URL: "https://example.com/photo.jpg"},
		},
	}

	tests := []struct {
		name      string
		params    *telego.SetMyProfilePhotoParams
		responses []settingsTestResponse
		wantCalls []string
		wantErr   error
	}{
		{
			name:      "set_photo",
			params:    params,
			responses: []settingsTestResponse{{}},
			wantCalls: []string{"setMyProfilePhoto"},
		},
		{
			name:      "remove_photo",
			params:    &telego.SetMyProfilePhotoParams{},
			responses: []settingsTestResponse{{}},
			wantCalls: []string{"removeMyProfilePhoto"},
		},
		{
			name:      "set_error",
			params:    params,
			responses: []settingsTestResponse{{err: errSettingsTest}},
			wantCalls: []string{"setMyProfilePhoto"},
			wantErr:   errSettingsTest,
		},
		{
			name:      "remove_error",
			params:    &telego.SetMyProfilePhotoParams{},
			responses: []settingsTestResponse{{err: errSettingsTest}},
			wantCalls: []string{"removeMyProfilePhoto"},
			wantErr:   errSettingsTest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bot, caller := newSettingsTestBot(t, tt.responses...)
			err := SetMyProfilePhoto(t.Context(), bot, tt.params)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
			assertSettingsCalls(t, caller, tt.wantCalls)
		})
	}
}

func TestApplySettings(t *testing.T) {
	trueValue := true
	command := telego.BotCommand{Command: "start", Description: "Start"}
	settings := &Settings{
		Name:             []*telego.SetMyNameParams{{Name: "new"}},
		Description:      []*telego.SetMyDescriptionParams{{Description: "new"}},
		ShortDescription: []*telego.SetMyShortDescriptionParams{{ShortDescription: "new"}},
		Commands:         []*telego.SetMyCommandsParams{{Commands: []telego.BotCommand{command}}},
		DefaultAdministratorRights: []*telego.SetMyDefaultAdministratorRightsParams{{
			Rights: &telego.ChatAdministratorRights{CanDeleteMessages: true},
		}},
		Settings:     &telego.SetMySettingsParams{SupportsInlineQueries: &trueValue},
		ProfilePhoto: &telego.SetMyProfilePhotoParams{},
	}

	t.Run("applies_all_settings_in_order", func(t *testing.T) {
		bot, caller := newSettingsTestBot(t,
			settingsTestResponse{result: `{"name":"old"}`}, settingsTestResponse{},
			settingsTestResponse{result: `{"description":"old"}`}, settingsTestResponse{},
			settingsTestResponse{result: `{"short_description":"old"}`}, settingsTestResponse{},
			settingsTestResponse{result: `[]`}, settingsTestResponse{},
			settingsTestResponse{result: `{}`}, settingsTestResponse{},
			settingsTestResponse{result: `{}`}, settingsTestResponse{},
			settingsTestResponse{},
		)

		err := ApplySettings(t.Context(), bot, settings)

		require.NoError(t, err)
		assertSettingsCalls(t, caller, []string{
			"getMyName", "setMyName",
			"getMyDescription", "setMyDescription",
			"getMyShortDescription", "setMyShortDescription",
			"getMyCommands", "setMyCommands",
			"getMyDefaultAdministratorRights", "setMyDefaultAdministratorRights",
			"getMe", "setMySettings",
			"removeMyProfilePhoto",
		})
	})

	tests := []struct {
		name     string
		settings *Settings
		wantCall string
	}{
		{
			name:     "empty_settings",
			settings: &Settings{},
		},
		{
			name:     "name_error",
			settings: &Settings{Name: []*telego.SetMyNameParams{{Name: "name"}}},
			wantCall: "getMyName",
		},
		{
			name:     "description_error",
			settings: &Settings{Description: []*telego.SetMyDescriptionParams{{Description: "description"}}},
			wantCall: "getMyDescription",
		},
		{
			name:     "short_description_error",
			settings: &Settings{ShortDescription: []*telego.SetMyShortDescriptionParams{{ShortDescription: "short"}}},
			wantCall: "getMyShortDescription",
		},
		{
			name:     "commands_error",
			settings: &Settings{Commands: []*telego.SetMyCommandsParams{{Commands: []telego.BotCommand{command}}}},
			wantCall: "getMyCommands",
		},
		{
			name: "administrator_rights_error",
			settings: &Settings{DefaultAdministratorRights: []*telego.SetMyDefaultAdministratorRightsParams{{
				Rights: &telego.ChatAdministratorRights{CanDeleteMessages: true},
			}}},
			wantCall: "getMyDefaultAdministratorRights",
		},
		{
			name:     "settings_error",
			settings: &Settings{Settings: &telego.SetMySettingsParams{}},
			wantCall: "getMe",
		},
		{
			name:     "profile_photo_error",
			settings: &Settings{ProfilePhoto: &telego.SetMyProfilePhotoParams{}},
			wantCall: "removeMyProfilePhoto",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var responses []settingsTestResponse
			if tt.wantCall != "" {
				responses = append(responses, settingsTestResponse{err: errSettingsTest})
			}

			bot, caller := newSettingsTestBot(t, responses...)

			err := ApplySettings(t.Context(), bot, tt.settings)

			if tt.wantCall == "" {
				require.NoError(t, err)
				assertSettingsCalls(t, caller, nil)
				return
			}

			require.ErrorIs(t, err, errSettingsTest)
			assertSettingsCalls(t, caller, []string{tt.wantCall})
		})
	}
}
