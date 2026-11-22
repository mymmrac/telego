package telegoutil

import (
	"context"
	"slices"

	"github.com/mymmrac/telego"
)

// Settings contains all available parameters to set bot settings.
type Settings struct {
	// Name of the bot in different languages.
	Name []*telego.SetMyNameParams
	// Description of the bot in different languages.
	Description []*telego.SetMyDescriptionParams
	// Short description of the bot in different languages.
	ShortDescription []*telego.SetMyShortDescriptionParams
	// Commands of the bot in different scopes and languages.
	// Note: If command list is empty, commands will be deleted in that scope and language.
	Commands []*telego.SetMyCommandsParams
	// Default administrator rights of the bot in different groups and channels.
	DefaultAdministratorRights []*telego.SetMyDefaultAdministratorRightsParams
	// Global settings of the bot.
	// Note: Some settings are not known by the bot, so they are always set if provided regardless of current value.
	Settings *telego.SetMySettingsParams
	// Profile photo of the bot.
	// Note: Since it's not possible to know if the current profile photo is the same as the provided one, it's always
	// set if provided.
	// Note: If photo (inside params) is nil, it will delete profile photo.
	ProfilePhoto *telego.SetMyProfilePhotoParams
}

// ApplySettings applies all available settings to the bot in efficient way (usually can be called multiple times in a
// row without hitting rate limits, but for some settings it's not guaranteed).
// Note: Does not reset any settings to default values or clears any settings for different scopes.
// Note: Some settings are not known by the bot, so they are always set if provided regardless of current value.
func ApplySettings(ctx context.Context, bot *telego.Bot, settings *Settings) error {
	for _, params := range settings.Name {
		if err := SetMyName(ctx, bot, params); err != nil {
			return err
		}
	}

	for _, params := range settings.Description {
		if err := SetMyDescription(ctx, bot, params); err != nil {
			return err
		}
	}

	for _, params := range settings.ShortDescription {
		if err := SetMyShortDescription(ctx, bot, params); err != nil {
			return err
		}
	}

	for _, params := range settings.Commands {
		if err := SetMyCommands(ctx, bot, params); err != nil {
			return err
		}
	}

	for _, params := range settings.DefaultAdministratorRights {
		if err := SetMyDefaultAdministratorRights(ctx, bot, params); err != nil {
			return err
		}
	}

	if settings.Settings != nil {
		if err := SetMySettings(ctx, bot, settings.Settings); err != nil {
			return err
		}
	}

	if settings.ProfilePhoto != nil {
		if err := SetMyProfilePhoto(ctx, bot, settings.ProfilePhoto); err != nil {
			return err
		}
	}

	return nil
}

// SetMyName sets the bot name if needed.
func SetMyName(ctx context.Context, bot *telego.Bot, params *telego.SetMyNameParams) error {
	current, err := bot.GetMyName(ctx, &telego.GetMyNameParams{
		LanguageCode: params.LanguageCode,
	})
	if err != nil {
		return err
	}

	if current.Name != params.Name {
		if err = bot.SetMyName(ctx, params); err != nil {
			return err
		}
	}

	return nil
}

// SetMyDescription sets the bot description if needed.
func SetMyDescription(ctx context.Context, bot *telego.Bot, params *telego.SetMyDescriptionParams) error {
	current, err := bot.GetMyDescription(ctx, &telego.GetMyDescriptionParams{
		LanguageCode: params.LanguageCode,
	})
	if err != nil {
		return err
	}

	if current.Description != params.Description {
		if err = bot.SetMyDescription(ctx, params); err != nil {
			return err
		}
	}

	return nil
}

// SetMyShortDescription sets the bot short description if needed.
func SetMyShortDescription(ctx context.Context, bot *telego.Bot, params *telego.SetMyShortDescriptionParams) error {
	current, err := bot.GetMyShortDescription(ctx, &telego.GetMyShortDescriptionParams{
		LanguageCode: params.LanguageCode,
	})
	if err != nil {
		return err
	}

	if current.ShortDescription != params.ShortDescription {
		if err = bot.SetMyShortDescription(ctx, params); err != nil {
			return err
		}
	}

	return nil
}

// SetMyCommands sets the bot commands if needed. If command list is empty it will delete commands instead.
func SetMyCommands(ctx context.Context, bot *telego.Bot, params *telego.SetMyCommandsParams) error {
	current, err := bot.GetMyCommands(ctx, &telego.GetMyCommandsParams{
		Scope:        params.Scope,
		LanguageCode: params.LanguageCode,
	})
	if err != nil {
		return err
	}

	if slices.Equal(current, params.Commands) {
		return nil
	}

	if len(params.Commands) == 0 {
		err = bot.DeleteMyCommands(ctx, &telego.DeleteMyCommandsParams{
			Scope:        params.Scope,
			LanguageCode: params.LanguageCode,
		})
		if err != nil {
			return err
		}
	} else {
		if err = bot.SetMyCommands(ctx, params); err != nil {
			return err
		}
	}

	return nil
}

// SetMyDefaultAdministratorRights sets the bot default administrator rights if needed.
func SetMyDefaultAdministratorRights(
	ctx context.Context, bot *telego.Bot, params *telego.SetMyDefaultAdministratorRightsParams,
) error {
	current, err := bot.GetMyDefaultAdministratorRights(ctx, &telego.GetMyDefaultAdministratorRightsParams{
		ForChannels: params.ForChannels,
	})
	if err != nil {
		return err
	}

	if current == nil && params.Rights != nil ||
		current != nil && (params.Rights == nil || *current != *params.Rights) {
		if err = bot.SetMyDefaultAdministratorRights(ctx, params); err != nil {
			return err
		}
	}

	return nil
}

// SetMySettings sets the bot settings if needed.
// Note: Some settings are not known by the bot, so they are always considered not equal if provided.
func SetMySettings(ctx context.Context, bot *telego.Bot, params *telego.SetMySettingsParams) error {
	me, err := bot.GetMe(ctx)
	if err != nil {
		return err
	}

	if !IsMySettingsEqual(me, params) {
		if err = bot.SetMySettings(ctx, params); err != nil {
			return err
		}
	}

	return nil
}

// IsMySettingsEqual checks if the bot settings are equal to the given parameters. If the setting is not provided (nil)
// it is consigned to be equal.
// Note: Some settings are not known by the bot, so they are always considered not equal if provided.
func IsMySettingsEqual(me *telego.BotFullInfo, params *telego.SetMySettingsParams) bool {
	switch {
	case
		params.SupportsInlineQueries != nil && me.SupportsInlineQueries != *params.SupportsInlineQueries,
		params.InlineQueryPlaceholder != "" && me.InlineQueryPlaceholder != params.InlineQueryPlaceholder,
		params.NeedsInlineQueryLocation != nil && me.NeedsInlineQueryLocation != *params.NeedsInlineQueryLocation,
		params.InlineFeedbackShareDenominator != nil, // Current value not known, always not equal
		params.CanJoinGroups != nil && me.CanJoinGroups != *params.CanJoinGroups,
		params.CanReadAllGroupMessages != nil && me.CanReadAllGroupMessages != *params.CanReadAllGroupMessages,
		params.CanConnectToBusiness != nil && me.CanConnectToBusiness != *params.CanConnectToBusiness,
		params.CanManageBots != nil && me.CanManageBots != *params.CanManageBots,
		params.SupportsGuestQueries != nil && me.SupportsGuestQueries != *params.SupportsGuestQueries,
		params.SupportsJoinRequestQueries != nil && me.SupportsJoinRequestQueries != *params.SupportsJoinRequestQueries,
		params.HasTopicsEnabled != nil && me.HasTopicsEnabled != *params.HasTopicsEnabled,
		params.AllowsUsersToCreateTopics != nil && me.AllowsUsersToCreateTopics != *params.AllowsUsersToCreateTopics,
		params.HasTopicSidebar != nil && me.HasTopicSidebar != *params.HasTopicSidebar,
		params.CanCommunicateWithBots != nil, // Current value not known, always not equal
		params.CallbackQueryMode != nil:      // Current value not known, always not equal
		return false
	default:
		return true
	}
}

// SetMyProfilePhoto sets the bot profile photo. If photo is nil it will delete profile photo instead.
func SetMyProfilePhoto(ctx context.Context, bot *telego.Bot, params *telego.SetMyProfilePhotoParams) error {
	if params.Photo == nil {
		if err := bot.RemoveMyProfilePhoto(ctx); err != nil {
			return err
		}
	} else {
		if err := bot.SetMyProfilePhoto(ctx, params); err != nil {
			return err
		}
	}
	return nil
}
