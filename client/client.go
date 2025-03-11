// This file is auto-generated, don't edit it. Thanks.
package client

import (
	rpcutil "github.com/alibabacloud-go/tea-rpc-utils/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	antchainutil "github.com/antchain-openapi-sdk-go/antchain-util/service"
)

/**
 * Model for initing client
 */
type Config struct {
	// accesskey id
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// accesskey secret
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	// security token
	SecurityToken *string `json:"securityToken,omitempty" xml:"securityToken,omitempty"`
	// http protocol
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// read timeout
	ReadTimeout *int `json:"readTimeout,omitempty" xml:"readTimeout,omitempty"`
	// connect timeout
	ConnectTimeout *int `json:"connectTimeout,omitempty" xml:"connectTimeout,omitempty"`
	// http proxy
	HttpProxy *string `json:"httpProxy,omitempty" xml:"httpProxy,omitempty"`
	// https proxy
	HttpsProxy *string `json:"httpsProxy,omitempty" xml:"httpsProxy,omitempty"`
	// endpoint
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// proxy white list
	NoProxy *string `json:"noProxy,omitempty" xml:"noProxy,omitempty"`
	// max idle conns
	MaxIdleConns *int `json:"maxIdleConns,omitempty" xml:"maxIdleConns,omitempty"`
	// user agent
	UserAgent *string `json:"userAgent,omitempty" xml:"userAgent,omitempty"`
	// socks5 proxy
	Socks5Proxy *string `json:"socks5Proxy,omitempty" xml:"socks5Proxy,omitempty"`
	// socks5 network
	Socks5NetWork *string `json:"socks5NetWork,omitempty" xml:"socks5NetWork,omitempty"`
	// 长链接最大空闲时长
	MaxIdleTimeMillis *int `json:"maxIdleTimeMillis,omitempty" xml:"maxIdleTimeMillis,omitempty"`
	// 长链接最大连接时长
	KeepAliveDurationMillis *int `json:"keepAliveDurationMillis,omitempty" xml:"keepAliveDurationMillis,omitempty"`
	// 最大连接数（长链接最大总数）
	MaxRequests *int `json:"maxRequests,omitempty" xml:"maxRequests,omitempty"`
	// 每个目标主机的最大连接数（分主机域名的长链接最大总数
	MaxRequestsPerHost *int `json:"maxRequestsPerHost,omitempty" xml:"maxRequestsPerHost,omitempty"`
}

func (s Config) String() string {
	return tea.Prettify(s)
}

func (s Config) GoString() string {
	return s.String()
}

func (s *Config) SetAccessKeyId(v string) *Config {
	s.AccessKeyId = &v
	return s
}

func (s *Config) SetAccessKeySecret(v string) *Config {
	s.AccessKeySecret = &v
	return s
}

func (s *Config) SetSecurityToken(v string) *Config {
	s.SecurityToken = &v
	return s
}

func (s *Config) SetProtocol(v string) *Config {
	s.Protocol = &v
	return s
}

func (s *Config) SetReadTimeout(v int) *Config {
	s.ReadTimeout = &v
	return s
}

func (s *Config) SetConnectTimeout(v int) *Config {
	s.ConnectTimeout = &v
	return s
}

func (s *Config) SetHttpProxy(v string) *Config {
	s.HttpProxy = &v
	return s
}

func (s *Config) SetHttpsProxy(v string) *Config {
	s.HttpsProxy = &v
	return s
}

func (s *Config) SetEndpoint(v string) *Config {
	s.Endpoint = &v
	return s
}

func (s *Config) SetNoProxy(v string) *Config {
	s.NoProxy = &v
	return s
}

func (s *Config) SetMaxIdleConns(v int) *Config {
	s.MaxIdleConns = &v
	return s
}

func (s *Config) SetUserAgent(v string) *Config {
	s.UserAgent = &v
	return s
}

func (s *Config) SetSocks5Proxy(v string) *Config {
	s.Socks5Proxy = &v
	return s
}

func (s *Config) SetSocks5NetWork(v string) *Config {
	s.Socks5NetWork = &v
	return s
}

func (s *Config) SetMaxIdleTimeMillis(v int) *Config {
	s.MaxIdleTimeMillis = &v
	return s
}

func (s *Config) SetKeepAliveDurationMillis(v int) *Config {
	s.KeepAliveDurationMillis = &v
	return s
}

func (s *Config) SetMaxRequests(v int) *Config {
	s.MaxRequests = &v
	return s
}

func (s *Config) SetMaxRequestsPerHost(v int) *Config {
	s.MaxRequestsPerHost = &v
	return s
}

type FeedbackReportDataRequest struct {
	// OAuth模式下的授权token
	AuthToken *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	// 广告主账号ID
	AccountId *string `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	// 报表类型级别
	Level *string `json:"level,omitempty" xml:"level,omitempty" require:"true"`
	// 回传数据明细，类型json array
	FeedbackData *string `json:"feedback_data,omitempty" xml:"feedback_data,omitempty" require:"true"`
}

func (s FeedbackReportDataRequest) String() string {
	return tea.Prettify(s)
}

func (s FeedbackReportDataRequest) GoString() string {
	return s.String()
}

func (s *FeedbackReportDataRequest) SetAuthToken(v string) *FeedbackReportDataRequest {
	s.AuthToken = &v
	return s
}

func (s *FeedbackReportDataRequest) SetAccountId(v string) *FeedbackReportDataRequest {
	s.AccountId = &v
	return s
}

func (s *FeedbackReportDataRequest) SetLevel(v string) *FeedbackReportDataRequest {
	s.Level = &v
	return s
}

func (s *FeedbackReportDataRequest) SetFeedbackData(v string) *FeedbackReportDataRequest {
	s.FeedbackData = &v
	return s
}

type FeedbackReportDataResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s FeedbackReportDataResponse) String() string {
	return tea.Prettify(s)
}

func (s FeedbackReportDataResponse) GoString() string {
	return s.String()
}

func (s *FeedbackReportDataResponse) SetReqMsgId(v string) *FeedbackReportDataResponse {
	s.ReqMsgId = &v
	return s
}

func (s *FeedbackReportDataResponse) SetResultCode(v string) *FeedbackReportDataResponse {
	s.ResultCode = &v
	return s
}

func (s *FeedbackReportDataResponse) SetResultMsg(v string) *FeedbackReportDataResponse {
	s.ResultMsg = &v
	return s
}

func (s *FeedbackReportDataResponse) SetSuccess(v bool) *FeedbackReportDataResponse {
	s.Success = &v
	return s
}

type ConvertAdDataRequest struct {
	// OAuth模式下的授权token
	AuthToken *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	// 广告主id
	AccountId *int64 `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	// ios/android
	DeviceOsType *string `json:"device_os_type,omitempty" xml:"device_os_type,omitempty" require:"true"`
	// 设备ID（imei或idfa的加密值）
	Muid *string `json:"muid,omitempty" xml:"muid,omitempty" require:"true"`
	// 点击ID
	ClickId *string `json:"click_id,omitempty" xml:"click_id,omitempty" require:"true"`
	// 转化时间
	ConvTime *int64 `json:"conv_time,omitempty" xml:"conv_time,omitempty" require:"true"`
	// 点击时间
	ClickTime *int64 `json:"click_time,omitempty" xml:"click_time,omitempty" require:"true"`
	// 曝光时间
	ImpressionTime *string `json:"impression_time,omitempty" xml:"impression_time,omitempty" require:"true"`
	// 投放日期年月日时分秒（准确到秒），格式为 yyyyMMddhhmmss
	Dt *string `json:"dt,omitempty" xml:"dt,omitempty" require:"true"`
	// 手机号MD5
	MobileMd5 *string `json:"mobile_md5,omitempty" xml:"mobile_md5,omitempty" require:"true"`
	// 是否提单标签0,1
	LabelSubmit *int64 `json:"label_submit,omitempty" xml:"label_submit,omitempty" require:"true"`
	// 是否支付标签0,1
	LabelPay *int64 `json:"label_pay,omitempty" xml:"label_pay,omitempty" require:"true"`
	// 是否升级标签0,1
	LabelUp *int64 `json:"label_up,omitempty" xml:"label_up,omitempty"`
	// m2是否续期
	LabelM2Renewal *int64 `json:"label_m2_renewal,omitempty" xml:"label_m2_renewal,omitempty" require:"true"`
	// 是否退保
	LabelSurrender *int64 `json:"label_surrender,omitempty" xml:"label_surrender,omitempty"`
	// 区分投放渠道来源weixin\youlianghui\chuanshanjia\douyin
	Platform *int64 `json:"platform,omitempty" xml:"platform,omitempty"`
}

func (s ConvertAdDataRequest) String() string {
	return tea.Prettify(s)
}

func (s ConvertAdDataRequest) GoString() string {
	return s.String()
}

func (s *ConvertAdDataRequest) SetAuthToken(v string) *ConvertAdDataRequest {
	s.AuthToken = &v
	return s
}

func (s *ConvertAdDataRequest) SetAccountId(v int64) *ConvertAdDataRequest {
	s.AccountId = &v
	return s
}

func (s *ConvertAdDataRequest) SetDeviceOsType(v string) *ConvertAdDataRequest {
	s.DeviceOsType = &v
	return s
}

func (s *ConvertAdDataRequest) SetMuid(v string) *ConvertAdDataRequest {
	s.Muid = &v
	return s
}

func (s *ConvertAdDataRequest) SetClickId(v string) *ConvertAdDataRequest {
	s.ClickId = &v
	return s
}

func (s *ConvertAdDataRequest) SetConvTime(v int64) *ConvertAdDataRequest {
	s.ConvTime = &v
	return s
}

func (s *ConvertAdDataRequest) SetClickTime(v int64) *ConvertAdDataRequest {
	s.ClickTime = &v
	return s
}

func (s *ConvertAdDataRequest) SetImpressionTime(v string) *ConvertAdDataRequest {
	s.ImpressionTime = &v
	return s
}

func (s *ConvertAdDataRequest) SetDt(v string) *ConvertAdDataRequest {
	s.Dt = &v
	return s
}

func (s *ConvertAdDataRequest) SetMobileMd5(v string) *ConvertAdDataRequest {
	s.MobileMd5 = &v
	return s
}

func (s *ConvertAdDataRequest) SetLabelSubmit(v int64) *ConvertAdDataRequest {
	s.LabelSubmit = &v
	return s
}

func (s *ConvertAdDataRequest) SetLabelPay(v int64) *ConvertAdDataRequest {
	s.LabelPay = &v
	return s
}

func (s *ConvertAdDataRequest) SetLabelUp(v int64) *ConvertAdDataRequest {
	s.LabelUp = &v
	return s
}

func (s *ConvertAdDataRequest) SetLabelM2Renewal(v int64) *ConvertAdDataRequest {
	s.LabelM2Renewal = &v
	return s
}

func (s *ConvertAdDataRequest) SetLabelSurrender(v int64) *ConvertAdDataRequest {
	s.LabelSurrender = &v
	return s
}

func (s *ConvertAdDataRequest) SetPlatform(v int64) *ConvertAdDataRequest {
	s.Platform = &v
	return s
}

type ConvertAdDataResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 调用是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回码描述
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求的唯一id
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

func (s ConvertAdDataResponse) String() string {
	return tea.Prettify(s)
}

func (s ConvertAdDataResponse) GoString() string {
	return s.String()
}

func (s *ConvertAdDataResponse) SetReqMsgId(v string) *ConvertAdDataResponse {
	s.ReqMsgId = &v
	return s
}

func (s *ConvertAdDataResponse) SetResultCode(v string) *ConvertAdDataResponse {
	s.ResultCode = &v
	return s
}

func (s *ConvertAdDataResponse) SetResultMsg(v string) *ConvertAdDataResponse {
	s.ResultMsg = &v
	return s
}

func (s *ConvertAdDataResponse) SetSuccess(v bool) *ConvertAdDataResponse {
	s.Success = &v
	return s
}

func (s *ConvertAdDataResponse) SetCode(v string) *ConvertAdDataResponse {
	s.Code = &v
	return s
}

func (s *ConvertAdDataResponse) SetMessage(v string) *ConvertAdDataResponse {
	s.Message = &v
	return s
}

func (s *ConvertAdDataResponse) SetRequestId(v string) *ConvertAdDataResponse {
	s.RequestId = &v
	return s
}

type Client struct {
	Endpoint                *string
	RegionId                *string
	AccessKeyId             *string
	AccessKeySecret         *string
	Protocol                *string
	UserAgent               *string
	ReadTimeout             *int
	ConnectTimeout          *int
	HttpProxy               *string
	HttpsProxy              *string
	Socks5Proxy             *string
	Socks5NetWork           *string
	NoProxy                 *string
	MaxIdleConns            *int
	SecurityToken           *string
	MaxIdleTimeMillis       *int
	KeepAliveDurationMillis *int
	MaxRequests             *int
	MaxRequestsPerHost      *int
}

/**
 * Init client with Config
 * @param config config contains the necessary information to create a client
 */
func NewClient(config *Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *Config) (_err error) {
	if tea.BoolValue(util.IsUnset(tea.ToMap(config))) {
		_err = tea.NewSDKError(map[string]interface{}{
			"code":    "ParameterMissing",
			"message": "'config' can not be unset",
		})
		return _err
	}

	client.AccessKeyId = config.AccessKeyId
	client.AccessKeySecret = config.AccessKeySecret
	client.SecurityToken = config.SecurityToken
	client.Endpoint = config.Endpoint
	client.Protocol = config.Protocol
	client.UserAgent = config.UserAgent
	client.ReadTimeout = util.DefaultNumber(config.ReadTimeout, tea.Int(20000))
	client.ConnectTimeout = util.DefaultNumber(config.ConnectTimeout, tea.Int(20000))
	client.HttpProxy = config.HttpProxy
	client.HttpsProxy = config.HttpsProxy
	client.NoProxy = config.NoProxy
	client.Socks5Proxy = config.Socks5Proxy
	client.Socks5NetWork = config.Socks5NetWork
	client.MaxIdleConns = util.DefaultNumber(config.MaxIdleConns, tea.Int(60000))
	client.MaxIdleTimeMillis = util.DefaultNumber(config.MaxIdleTimeMillis, tea.Int(5))
	client.KeepAliveDurationMillis = util.DefaultNumber(config.KeepAliveDurationMillis, tea.Int(5000))
	client.MaxRequests = util.DefaultNumber(config.MaxRequests, tea.Int(100))
	client.MaxRequestsPerHost = util.DefaultNumber(config.MaxRequestsPerHost, tea.Int(100))
	return nil
}

/**
 * Encapsulate the request and invoke the network
 * @param action api name
 * @param protocol http or https
 * @param method e.g. GET
 * @param pathname pathname of every api
 * @param request which contains request params
 * @param runtime which controls some details of call api, such as retry times
 * @return the response
 */
func (client *Client) DoRequest(version *string, action *string, protocol *string, method *string, pathname *string, request map[string]interface{}, headers map[string]*string, runtime *util.RuntimeOptions) (_result map[string]interface{}, _err error) {
	_err = tea.Validate(runtime)
	if _err != nil {
		return _result, _err
	}
	_runtime := map[string]interface{}{
		"timeouted":          "retry",
		"readTimeout":        tea.IntValue(util.DefaultNumber(runtime.ReadTimeout, client.ReadTimeout)),
		"connectTimeout":     tea.IntValue(util.DefaultNumber(runtime.ConnectTimeout, client.ConnectTimeout)),
		"httpProxy":          tea.StringValue(util.DefaultString(runtime.HttpProxy, client.HttpProxy)),
		"httpsProxy":         tea.StringValue(util.DefaultString(runtime.HttpsProxy, client.HttpsProxy)),
		"noProxy":            tea.StringValue(util.DefaultString(runtime.NoProxy, client.NoProxy)),
		"maxIdleConns":       tea.IntValue(util.DefaultNumber(runtime.MaxIdleConns, client.MaxIdleConns)),
		"maxIdleTimeMillis":  tea.IntValue(client.MaxIdleTimeMillis),
		"keepAliveDuration":  tea.IntValue(client.KeepAliveDurationMillis),
		"maxRequests":        tea.IntValue(client.MaxRequests),
		"maxRequestsPerHost": tea.IntValue(client.MaxRequestsPerHost),
		"retry": map[string]interface{}{
			"retryable":   tea.BoolValue(runtime.Autoretry),
			"maxAttempts": tea.IntValue(util.DefaultNumber(runtime.MaxAttempts, tea.Int(3))),
		},
		"backoff": map[string]interface{}{
			"policy": tea.StringValue(util.DefaultString(runtime.BackoffPolicy, tea.String("no"))),
			"period": tea.IntValue(util.DefaultNumber(runtime.BackoffPeriod, tea.Int(1))),
		},
		"ignoreSSL": tea.BoolValue(runtime.IgnoreSSL),
	}

	_resp := make(map[string]interface{})
	for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
		if _retryTimes > 0 {
			_backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
			if tea.IntValue(_backoffTime) > 0 {
				tea.Sleep(_backoffTime)
			}
		}

		_resp, _err = func() (map[string]interface{}, error) {
			request_ := tea.NewRequest()
			request_.Protocol = util.DefaultString(client.Protocol, protocol)
			request_.Method = method
			request_.Pathname = pathname
			request_.Query = map[string]*string{
				"method":           action,
				"version":          version,
				"sign_type":        tea.String("HmacSHA1"),
				"req_time":         antchainutil.GetTimestamp(),
				"req_msg_id":       antchainutil.GetNonce(),
				"access_key":       client.AccessKeyId,
				"base_sdk_version": tea.String("TeaSDK-2.0"),
				"sdk_version":      tea.String("2.0.2"),
				"_prod_code":       tea.String("MORSERTA"),
				"_prod_channel":    tea.String("default"),
			}
			if !tea.BoolValue(util.Empty(client.SecurityToken)) {
				request_.Query["security_token"] = client.SecurityToken
			}

			request_.Headers = tea.Merge(map[string]*string{
				"host":       util.DefaultString(client.Endpoint, tea.String("openapi.antchain.antgroup.com")),
				"user-agent": util.GetUserAgent(client.UserAgent),
			}, headers)
			tmp := util.AnyifyMapValue(rpcutil.Query(request))
			request_.Body = tea.ToReader(util.ToFormString(tmp))
			request_.Headers["content-type"] = tea.String("application/x-www-form-urlencoded")
			signedParam := tea.Merge(request_.Query,
				rpcutil.Query(request))
			request_.Query["sign"] = antchainutil.GetSignature(signedParam, client.AccessKeySecret)
			response_, _err := tea.DoRequest(request_, _runtime)
			if _err != nil {
				return _result, _err
			}
			raw, _err := util.ReadAsString(response_.Body)
			if _err != nil {
				return _result, _err
			}

			obj := util.ParseJSON(raw)
			res, _err := util.AssertAsMap(obj)
			if _err != nil {
				return _result, _err
			}

			resp, _err := util.AssertAsMap(res["response"])
			if _err != nil {
				return _result, _err
			}

			if tea.BoolValue(antchainutil.HasError(raw, client.AccessKeySecret)) {
				_err = tea.NewSDKError(map[string]interface{}{
					"message": resp["result_msg"],
					"data":    resp,
					"code":    resp["result_code"],
				})
				return _result, _err
			}

			_result = resp
			return _result, _err
		}()
		if !tea.BoolValue(tea.Retryable(_err)) {
			break
		}
	}

	return _resp, _err
}

/**
 * Description: RTA广告主数据回传
 * Summary: RTA广告主数据回传
 */
func (client *Client) FeedbackReportData(request *FeedbackReportDataRequest) (_result *FeedbackReportDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FeedbackReportDataResponse{}
	_body, _err := client.FeedbackReportDataEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: RTA广告主数据回传
 * Summary: RTA广告主数据回传
 */
func (client *Client) FeedbackReportDataEx(request *FeedbackReportDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FeedbackReportDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &FeedbackReportDataResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("antcloud.morserta.report.data.feedback"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 摩斯RTA提供的转化回传接口
 * Summary: 摩斯RTA提供的转化回传接口
 */
func (client *Client) ConvertAdData(request *ConvertAdDataRequest) (_result *ConvertAdDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ConvertAdDataResponse{}
	_body, _err := client.ConvertAdDataEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 摩斯RTA提供的转化回传接口
 * Summary: 摩斯RTA提供的转化回传接口
 */
func (client *Client) ConvertAdDataEx(request *ConvertAdDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ConvertAdDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ConvertAdDataResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("antcloud.morserta.ad.data.convert"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
