package common

const (
	// SlackPrefix is the beginning of all of Slack kinds.
	SlackPrefix = "slack/"
)

const (
	// KindSlackUser is for resources that represent Slack users.
	KindSlackUser = SlackPrefix + "user"

	// KindSlackConversation represents a Slack chat.
	KindSlackConversation = SlackPrefix + "conversation"
)

const (
	// SlackConversationTypePrivateChannel represents a private Slack channel.
	SlackConversationTypePrivateChannel = "private_channel"

	// SlackConversationTypePublicChannel represents a public Slack channel.
	SlackConversationTypePublicChannel = "public_channel"
)

const (
	// LabelSlackManager contains the Slack ID of a user's manager.
	LabelSlackManager = SlackPrefix + LabelManager
)

const (
	// LabelTeamID contains the ID of the related team.
	LabelTeamID = SlackPrefix + "team/id"

	// LabelSlackIsAdmin indicates whether a Slack user is a workspace admin.
	LabelSlackIsAdmin = SlackPrefix + "is_admin"
	// LabelSlackIsOwner indicates whether a Slack user is a workspace owner.
	LabelSlackIsOwner = SlackPrefix + "is_owner"
	// LabelSlackIsPrimaryOwner indicates whether a Slack user is the primary workspace owner.
	LabelSlackIsPrimaryOwner = SlackPrefix + "is_primary_owner"
	// LabelSlackIsRestricted indicates whether a Slack user is restricted.
	LabelSlackIsRestricted = SlackPrefix + "is_restricted"
	// LabelSlackIsUltraRestricted indicates whether a Slack user is ultra-restricted.
	LabelSlackIsUltraRestricted = SlackPrefix + "is_ultra_restricted"
	// LabelSlackIsDeactivated indicates whether a Slack user has been deactivated
	// or deleted; Slack uses these terms interchangeably: api.slack.com/types/user
	LabelSlackIsDeactivated = SlackPrefix + "is_deactivated"
	// LabelSlackIsBot indicates whether a Slack user is a bot.
	LabelSlackIsBot = SlackPrefix + "is_bot"
	// LabelSlackHas2FA indicates whether a Slack user has 2FA configured.
	LabelSlackHas2FA = SlackPrefix + "has_2fa"
)

const (
	// SourceSlack is the Source when a Petition is received by Slack.
	SourceSlack = "Slack"
)

const (
	// ResourceIDSlackBot is the ID of the Slack bot.
	ResourceIDSlackBot = "USLACKBOT"
)
