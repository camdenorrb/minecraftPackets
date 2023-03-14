package common

type Chat struct {
	*ChatComponent
	Extra []ChatComponent `json:"extra"`
}

type ChatComponent struct {
	Text          string         `json:"text"`
	Bold          bool           `json:"bold"`
	Italic        bool           `json:"italic"`
	Underlined    bool           `json:"underlined"`
	Strikethrough bool           `json:"strikethrough"`
	Obfuscated    bool           `json:"obfuscated"`
	Font          string         `json:"font"`
	Color         string         `json:"color"`
	Insertion     string         `json:"insertion"`
	ClickEvent    ChatClickEvent `json:"clickEvent"`
	HoverEvent    ChatHoverEvent `json:"hoverEvent"`
}

type ChatClickEvent struct {
	OpenURL         string `json:"open_url"`
	OpenFile        string `json:"open_file"`
	RunCommand      string `json:"run_command"`
	TwitchUserInfo  string `json:"twitch_user_info"`
	SuggestCommand  string `json:"suggest_command"`
	ChangePage      string `json:"change_page"`
	CopyToClipboard string `json:"copy_to_clipboard"`
}

type ChatHoverEvent struct {
	ShowText        string `json:"show_text"`
	ShowItem        string `json:"show_item"`
	ShowEntity      string `json:"show_entity"`
	ShowAchievement string `json:"show_achievement"`
}
