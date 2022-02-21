package gosms

import (
	"testing"
)

func TestSum(t *testing.T) {

	client, err := NewSMSClient("", "", "", "")
	if err != nil {
		t.Error("error " + err.Error())
	}

	req := NewSMSRequest()
	req.AppId = "1400787878"
	req.Number = "xxx"
	req.SignName = ""
	req.TemplateCode = "111"
	req.TemplateParam = []*RequestParam{{"", ""}, {"", ""}}
	client.SendSMS(req)
}
