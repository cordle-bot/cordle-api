package models

import "fmt"

type Role struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Permissions    int64  `json:"permissions"`
	PermissionsNew string `json:"permissions_new"`
	Position       int    `json:"position"`
	Color          int    `json:"color"`
	Hoist          bool   `json:"hoist"`
	Managed        bool   `json:"managed"`
	Mentionable    bool   `json:"mentionable"`
	Icon           string `json:"icon"`
	UnicodeEmoji   string `json:"unicode_emoji"`
	Flags          int    `json:"flags"`
}

type Guild struct {
	ID                          string   `json:"id"`
	Name                        string   `json:"name"`
	Icon                        string   `json:"icon"`
	Description                 string   `json:"description"`
	HomeHeader                  string   `json:"home_header"`
	Splash                      string   `json:"splash"`
	DiscoverySplash             string   `json:"discovery_splash"`
	Features                    []string `json:"features"`
	Banner                      string   `json:"banner"`
	OwnerID                     string   `json:"owner_id"`
	ApplicationID               string   `json:"application_id"`
	Region                      string   `json:"region"`
	AfkChannelID                string   `json:"afk_channel_id"`
	AfkTimeout                  int      `json:"afk_timeout"`
	SystemChannelID             string   `json:"system_channel_id"`
	SystemChannelFlags          int      `json:"system_channel_flags"`
	WidgetEnabled               bool     `json:"widget_enabled"`
	WidgetChannelID             string   `json:"widget_channel_id"`
	VerificationLevel           int      `json:"verification_level"`
	Roles                       []Role   `json:"roles"`
	DefaultMessageNotifications int      `json:"default_message_notifications"`
	MFALevel                    int      `json:"mfa_level"`
	ExplicitContentFilter       int      `json:"explicit_content_filter"`
	MaxPresences                string   `json:"max_presences"`
	MaxMembers                  int      `json:"max_members"`
	MaxStageVideoChannelUsers   int      `json:"max_stage_video_channel_users"`
	MaxVideoChannelUsers        int      `json:"max_video_channel_users"`
	VanityURLCode               string   `json:"vanity_url_code"`
	PremiumTier                 int      `json:"premium_tier"`
	PremiumSubscriptionCount    int      `json:"premium_subscription_count"`
	PreferredLocale             string   `json:"preferred_locale"`
	RulesChannelID              string   `json:"rules_channel_id"`
	SafetyAlertsChannelID       string   `json:"safety_alerts_channel_id"`
	PublicUpdatesChannelID      string   `json:"public_updates_channel_id"`
	HubType                     string   `json:"hub_type"`
	PremiumProgressBarEnabled   bool     `json:"premium_progress_bar_enabled"`
	LatestOnboardingQuestionID  string   `json:"latest_onboarding_question_id"`
	NSFW                        bool     `json:"nsfw"`
	NSFWLevel                   int      `json:"nsfw_level"`
	Emojis                      []string `json:"emojis"`
	Stickers                    []string `json:"stickers"`
	IncidentsData               string   `json:"incidents_data"`
	InventorySettings           string   `json:"inventory_settings"`
	EmbedEnabled                bool     `json:"embed_enabled"`
	EmbedChannelID              string   `json:"embed_channel_id"`
	ApproximateMemberCount      int      `json:"approximate_member_count"`
	ApproximatePresenceCount    int      `json:"approximate_presence_count"`
}

func (g *Guild) String() string {
	return fmt.Sprintf(
		"%s\n%s\n%d\n%d\n",
		g.ID,
		g.Name,
		g.ApproximateMemberCount,
		g.ApproximatePresenceCount,
	)
}
