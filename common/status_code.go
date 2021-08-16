package common

type StatusCode int

const (
	InternalError                             StatusCode = -1
	OK                                        StatusCode = 0
	GrafanaWebhookUnmarshalJsonError          StatusCode = 100001
	ClientCallAPIError                        StatusCode = 100002
	WechatWorkCallAPIError                    StatusCode = 200001
	WechatWorkCallAPIWrongJsonFormatWarning   StatusCode = 200002
	WechatWorkParseResponseBodyFailureWarning StatusCode = 200003
)

func (s StatusCode) String() string {
	switch s {
	case InternalError:
		return "[Error] Internal Error"
	case OK:
		return "[OK]"
	case GrafanaWebhookUnmarshalJsonError:
		return "[ERROR] JSON Unmarshal failure when receive Grafana webhook"
	case ClientCallAPIError:
		return "[ERROR] Client call API failure"
	case WechatWorkCallAPIError:
		return "[ERROR] Call Wechat-Work API failure"
	case WechatWorkCallAPIWrongJsonFormatWarning:
		return "[Warning] Wrong json format when call Wechat-WorkC API"
	case WechatWorkParseResponseBodyFailureWarning:
		return "[Warning] Parse response body failure when call Wechat-WorkC API"
	default:
		return "[UNKNOWN]"
	}
}