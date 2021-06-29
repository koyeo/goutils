package _robot

type DingTalkMessage struct {
	MsgType    string                   `json:"msgtype"` // required
	At         *DingTalkAt              `json:"at,omitempty"`
	Text       *DingTalkMessageText     `json:"text,omitempty"`
	Link       *DingTalkMessageLink     `json:"link,omitempty"`
	Markdown   *DingTalkMessageMarkdown `json:"markdown,omitempty"`
	ActionCard *DingTalkActionCard      `json:"actionCard,omitempty"`
}

type DingTalkAt struct {
	AtMobiles []string `json:"atMobiles,omitempty"` // @用户的手机号
	AtUserIds []string `json:"atUserIds,omitempty"` // @人的用户ID
	IsAtAll   bool     `json:"isAtAll,omitempty"`   // 是否 @所有人
}

type DingTalkMessageText struct {
	Content string `json:"content"` // required
}

type DingTalkMessageLink struct {
	Title      string `json:"title"`      // required
	Text       string `json:"text"`       // required
	MessageUrl string `json:"messageUrl"` // required
	PicUrl     string `json:"picUrl,omitempty"`
}

type DingTalkMessageMarkdown struct {
	Title string `json:"title"` // required
	Text  string `json:"text"`  // required
}

type DingTalkActionCard struct {
	Title          string `json:"title"`       // required
	Text           string `json:"text"`        // required
	SingleTitle    string `json:"singleTitle"` // required
	SingleURL      string `json:"singleURL"`   // required
	BtnOrientation string `json:"btnOrientation,omitempty"`
}
