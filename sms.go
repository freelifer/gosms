package gosms

import "fmt"

//secretId string
//secretKey string
//region string
type SMSClient interface {
	InitWithAccessKey(region string, secretId string, secretKey string)

	//func sendSMS(number string, templateCode string) {}
	SendSMS(req *Request) (*Response, error)
}

var adapters = make(map[string]SMSClient)

// Register makes a cache adapter available by the adapter name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, adapter SMSClient) {
	if adapter == nil {
		panic("cache: Register adapter is nil")
	}
	if _, dup := adapters[name]; dup {
		panic("cache: Register called twice for adapter " + name)
	}
	adapters[name] = adapter
}

// Create a new cache driver by adapter and config string.
// config need to be correct JSON as string: {"interval":360}.
// it will start gc automatically.
func NewSMSClient(adapterName, region string, secretId string, secretKey string) (SMSClient, error) {
	adapter, ok := adapters[adapterName]
	if !ok {
		return nil, fmt.Errorf("sms: unknown adaptername %q (forgotten import?)", adapterName)
	}
	adapter.InitWithAccessKey(region, secretId, secretKey)
	return adapter, nil
}

type RequestParam struct {
	Key   string
	Value string
}

type Request struct {
	AppId         string
	Number        string
	SignName      string
	TemplateCode  string
	TemplateParam []*RequestParam
}

type Response struct {
	Source    string
	RequestId string
	BizId     string
	Code      string
	Message   string
}

func NewSMSRequest() *Request {
	return &Request{}
}
