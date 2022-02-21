package aliyun

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/freelifer/gosms"
)

type SMSClient struct {
	secretId  string
	secretKey string
	region    string
}

func init() {
	gosms.Register("aliyun", &SMSClient{})
}
func (client *SMSClient) InitWithAccessKey(region string, secretId string, secretKey string) {
	client.region = region
	client.secretId = secretId
	client.secretKey = secretKey
}

func (client *SMSClient) SendSMS(req *gosms.Request) (*gosms.Response, error) {
	c, err := dysmsapi.NewClientWithAccessKey(client.region, client.secretId, client.secretKey)

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = req.Number
	request.SignName = req.SignName
	request.TemplateCode = req.TemplateCode
	request.TemplateParam = client.templateParamSet(req)

	response, err := c.SendSms(request)
	if err != nil {
		return nil, err
	}

	resp := &gosms.Response{}
	resp.Source = "aliyun"
	resp.RequestId = response.RequestId
	resp.Code = response.Code
	resp.Message = response.Message
	resp.BizId = response.BizId
	return resp, nil
}

func (client *SMSClient) templateParamSet(req *gosms.Request) string {
	if req.TemplateParam == nil {
		return ""
	}

	if len(req.TemplateParam) <= 0 {
		return ""
	}
	param := make(map[string]string)
	for i := 0; i < len(req.TemplateParam); i++ {
		p := req.TemplateParam[i]
		param[p.Key] = param[p.Value]
	}
	b, _ := json.Marshal(param)
	return string(b)
}

func main() {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "<accessKeyId>", "<accessSecret>")

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = "17633802772"
	request.TemplateCode = "345"

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}
