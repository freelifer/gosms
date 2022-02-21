# gosms

### 阿里云短信费用

![tencent_sms](https://github.com/freelifer/gosms/blob/main/doc/img_2.png?raw=true)

### 腾讯云短信费用

![tencent_sms](https://github.com/freelifer/gosms/blob/main/doc/img.png?raw=true)

![tencent_sms](https://github.com/freelifer/gosms/blob/main/doc/img_1.png?raw=true)

### 使用

#### 安装 SDK

```base
go get -u github.com/freelifer/gosms
```

#### 使用阿里云方案发送短信

```golang
// 阿里云短信
client, err := NewSMSClient("aliyun", "<region>", "<accessKeyId>", "<accessSecret>")
if err != nil {
t.Error("error " + err.Error())
}

req := NewSMSRequest()
req.Number = "xxx"
req.SignName = "阿里云"
req.TemplateCode = "SMS_1530****"
req.TemplateParam = []*RequestParam{{"code", "5764"}, {"phone", "10086"}}
client.SendSMS(req)
```

#### 使用腾讯云方案发送短信

```golang
// 腾讯云短信
client, err := NewSMSClient("tencent", "<region>", "<accessKeyId>", "<accessSecret>")
if err != nil {
t.Error("error " + err.Error())
}

req := NewSMSRequest()
req.AppId = "1400787878"
req.Number = "xxxx"
req.SignName = "腾讯云"
req.TemplateCode = "SMS_1530****"
req.TemplateParam = []*RequestParam{{"code", "5764"}, {"phone", "10086"}}
client.SendSMS(req)
```
