package _robot

import (
	"testing"
)

func TestNewDingTalkRobot(t *testing.T) {
	robot := NewDingTalkRobot(&DingTalkConfig{
		Webhook:    "https://oapi.dingtalk.com/robot/send?access_token=767c18a0c45c610d4b44af399261ee06860b444db4e3de32eea73462a90eb849",
		SignSecret: "SEC93dd4c5817ea6578eaef0b7fcfde7f5802dcf0876444c6fe7d417dae8fb6a58c",
	})
	
	robot.Push(&DingTalkMessage{
		MsgType: "text",
		//At: &DingTalkAt{
		//	AtMobiles: []string{"18817392521"},
		//},
		Text: &DingTalkMessageText{
			Content: "Hello world!",
		},
	})
}
