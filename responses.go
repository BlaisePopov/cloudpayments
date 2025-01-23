package cloudpayments

// PaymentResponse представляет ответ на оплату.
type PaymentResponse struct {
	Success bool        `json:"Success"`
	Message string      `json:"Message"`
	Model   PaymentModel `json:"Model"`
}

// PaymentModel содержит детали платежа.
type PaymentModel struct {
	TransactionId               int64   `json:"TransactionId"`
	Amount                      float64 `json:"Amount"`
	Currency                    string  `json:"Currency"`
	CurrencyCode                int     `json:"CurrencyCode"`
	PaymentAmount               float64 `json:"PaymentAmount"`
	PaymentCurrency             string  `json:"PaymentCurrency"`
	PaymentCurrencyCode         int     `json:"PaymentCurrencyCode"`
	InvoiceId                   string  `json:"InvoiceId"`
	AccountId                   string  `json:"AccountId"`
	Email                       string  `json:"Email"`
	Description                 string  `json:"Description"`
	JsonData                    string  `json:"JsonData"`
	CreatedDate                 string  `json:"CreatedDate"`
	PayoutDate                  string  `json:"PayoutDate"`
	PayoutDateIso               string  `json:"PayoutDateIso"`
	PayoutAmount                float64 `json:"PayoutAmount"`
	CreatedDateIso              string  `json:"CreatedDateIso"`
	AuthDate                    string  `json:"AuthDate"`
	AuthDateIso                 string  `json:"AuthDateIso"`
	ConfirmDate                 string  `json:"ConfirmDate"`
	ConfirmDateIso              string  `json:"ConfirmDateIso"`
	AuthCode                    string  `json:"AuthCode"`
	TestMode                    bool    `json:"TestMode"`
	Rrn                         string  `json:"Rrn"`
	OriginalTransactionId       int64   `json:"OriginalTransactionId"`
	FallBackScenarioDeclinedTransactionId int64 `json:"FallBackScenarioDeclinedTransactionId"`
	IpAddress                   string  `json:"IpAddress"`
	IpCountry                   string  `json:"IpCountry"`
	IpCity                      string  `json:"IpCity"`
	IpRegion                    string  `json:"IpRegion"`
	IpDistrict                  string  `json:"IpDistrict"`
	IpLatitude                  float64 `json:"IpLatitude"`
	IpLongitude                 float64 `json:"IpLongitude"`
	CardFirstSix                string  `json:"CardFirstSix"`
	CardLastFour                string  `json:"CardLastFour"`
	CardExpDate                 string  `json:"CardExpDate"`
	CardType                    string  `json:"CardType"`
	CardProduct                 string  `json:"CardProduct"`
	CardCategory                string  `json:"CardCategory"`
	EscrowAccumulationId        string  `json:"EscrowAccumulationId"`
	IssuerBankCountry           string  `json:"IssuerBankCountry"`
	Issuer                      string  `json:"Issuer"`
	CardTypeCode                int     `json:"CardTypeCode"`
	Status                      string  `json:"Status"`
	StatusCode                  int     `json:"StatusCode"`
	CultureName                 string  `json:"CultureName"`
	Reason                      string  `json:"Reason"`
	CardHolderMessage           string  `json:"CardHolderMessage"`
	Type                        int     `json:"Type"`
	Refunded                    bool    `json:"Refunded"`
	Name                        string  `json:"Name"`
	Token                       string  `json:"Token"`
	SubscriptionId              string  `json:"SubscriptionId"`
	GatewayName                 string  `json:"GatewayName"`
	ApplePay                    bool    `json:"ApplePay"`
	AndroidPay                  bool    `json:"AndroidPay"`
	WalletType                  string  `json:"WalletType"`
	TotalFee                    float64 `json:"TotalFee"`
}

// Payment3DSResponse представляет ответ на оплату с использованием 3DS.
type Payment3DSResponse struct {
	Success bool        `json:"Success"`
	Message string      `json:"Message"`
	Model   Payment3DSModel `json:"Model"`
}

// Payment3DSModel содержит детали 3DS оплаты.
type Payment3DSModel struct {
	TransactionId       int64   `json:"TransactionId"`
	PaReq              string  `json:"PaReq"`
	GoReq              string  `json:"GoReq"`
	AcsUrl             string  `json:"AcsUrl"`
	ThreeDsSessionData string  `json:"ThreeDsSessionData"`
	IFrameIsAllowed    bool    `json:"IFrameIsAllowed"`
	FrameWidth         int     `json:"FrameWidth"`
	FrameHeight        int     `json:"FrameHeight"`
	ThreeDsCallbackId  string  `json:"ThreeDsCallbackId"`
	EscrowAccumulationId string `json:"EscrowAccumulationId"`
}

// BaseResponse представляет базовый ответ от API.
type BaseResponse struct {
	Success bool        `json:"Success"`
	Message string      `json:"Message"`
}

// SubscriptionsListResponse представляет ответ с информацией о подписках.
type SubscriptionsListResponse struct {
	Success bool           `json:"Success"`
	Message string         `json:"Message"`
	Model   []Subscription `json:"Model"`
}

// SubscriptionResponse представляет ответ на запросы, связанные с подписками.
type SubscriptionResponse struct {
	Success bool           `json:"Success"`
	Message string         `json:"Message"`
	Model   Subscription   `json:"Model"`
}

// Subscription содержит информацию о подписке.
type Subscription struct {
	Id                        string  `json:"Id"`
	AccountId                 string  `json:"AccountId"`
	Description               string  `json:"Description"`
	Email                     string  `json:"Email"`
	Amount                    float64 `json:"Amount"`
	CurrencyCode              int     `json:"CurrencyCode"`
	Currency                  string  `json:"Currency"`
	RequireConfirmation       bool    `json:"RequireConfirmation"`
	StartDate                 string  `json:"StartDate"`
	StartDateIso              string  `json:"StartDateIso"`
	IntervalCode              int     `json:"IntervalCode"`
	Interval                  string  `json:"Interval"`
	Period                    int     `json:"Period"`
	MaxPeriods                int     `json:"MaxPeriods"`
	CultureName               string  `json:"CultureName"`
	StatusCode                int     `json:"StatusCode"`
	Status                    string  `json:"Status"`
	SuccessfulTransactionsNumber int `json:"SuccessfulTransactionsNumber"`
	FailedTransactionsNumber  int    `json:"FailedTransactionsNumber"`
	LastTransactionDate       string `json:"LastTransactionDate"`
	LastTransactionDateIso    string `json:"LastTransactionDateIso"`
	NextTransactionDate       string `json:"NextTransactionDate"`
	NextTransactionDateIso    string `json:"NextTransactionDateIso"`
	Receipt                   string `json:"Receipt"`
	FailoverSchemeId          string `json:"FailoverSchemeId"`
}

// TokensListResponse представляет ответ с информацией о токенах.
type TokensListResponse struct {
	Success bool    `json:"Success"`
	Message string  `json:"Message"`
	Model   []Token `json:"Model"`
}

// Token содержит информацию о токене.
type Token struct {
	Token             string `json:"Token"`
	AccountId         string `json:"AccountId"`
	CardMask          string `json:"CardMask"`
	ExpirationDateMonth int  `json:"ExpirationDateMonth"`
	ExpirationDateYear int  `json:"ExpirationDateYear"`
}

// PaymentsListResponse представляет ответ с информацией о транзакциях за день.
type PaymentsListResponse struct {
	Success bool      `json:"Success"`
	Message string    `json:"Message"`
	Model   []Payment `json:"Model"`
}

// PaymentsListV2Response представляет ответ с информацией о транзакциях за произвольный период.
type PaymentsListV2Response struct {
	Success bool      `json:"Success"`
	Message string    `json:"Message"`
	Model   []Payment `json:"Model"`
}

// Payment содержит информацию о транзакции.
type Payment struct {
	PublicId                        string  `json:"PublicId"`
	TerminalUrl                     string  `json:"TerminalUrl"`
	TransactionId                   int64   `json:"TransactionId"`
	Amount                          float64 `json:"Amount"`
	Currency                        string  `json:"Currency"`
	CurrencyCode                    int     `json:"CurrencyCode"`
	PaymentAmount                   float64 `json:"PaymentAmount"`
	PaymentCurrency                 string  `json:"PaymentCurrency"`
	PaymentCurrencyCode             int     `json:"PaymentCurrencyCode"`
	InvoiceId                       string  `json:"InvoiceId"`
	AccountId                       string  `json:"AccountId"`
	Email                           string  `json:"Email"`
	Description                     string  `json:"Description"`
	JsonData                        string  `json:"JsonData"`
	CreatedDate                     string  `json:"CreatedDate"`
	PayoutDate                      string  `json:"PayoutDate"`
	PayoutDateIso                   string  `json:"PayoutDateIso"`
	PayoutAmount                    float64 `json:"PayoutAmount"`
	CreatedDateIso                  string  `json:"CreatedDateIso"`
	AuthDate                        string  `json:"AuthDate"`
	AuthDateIso                     string  `json:"AuthDateIso"`
	ConfirmDate                     string  `json:"ConfirmDate"`
	ConfirmDateIso                  string  `json:"ConfirmDateIso"`
	AuthCode                        string  `json:"AuthCode"`
	TestMode                        bool    `json:"TestMode"`
	Rrn                             string  `json:"Rrn"`
	OriginalTransactionId           int64   `json:"OriginalTransactionId"`
	FallBackScenarioDeclinedTransactionId int64 `json:"FallBackScenarioDeclinedTransactionId"`
	IpAddress                       string  `json:"IpAddress"`
	IpCountry                       string  `json:"IpCountry"`
	IpCity                          string  `json:"IpCity"`
	IpRegion                        string  `json:"IpRegion"`
	IpDistrict                      string  `json:"IpDistrict"`
	IpLatitude                      float64 `json:"IpLatitude"`
	IpLongitude                     float64 `json:"IpLongitude"`
	CardFirstSix                    string  `json:"CardFirstSix"`
	CardLastFour                    string  `json:"CardLastFour"`
	CardExpDate                     string  `json:"CardExpDate"`
	CardType                        string  `json:"CardType"`
	CardProduct                     string  `json:"CardProduct"`
	CardCategory                    string  `json:"CardCategory"`
	EscrowAccumulationId            string  `json:"EscrowAccumulationId"`
	IssuerBankCountry               string  `json:"IssuerBankCountry"`
	Issuer                          string  `json:"Issuer"`
	CardTypeCode                    int     `json:"CardTypeCode"`
	Status                          string  `json:"Status"`
	StatusCode                      int     `json:"StatusCode"`
	CultureName                     string  `json:"CultureName"`
	Reason                          string  `json:"Reason"`
	CardHolderMessage               string  `json:"CardHolderMessage"`
	Type                            int     `json:"Type"`
	Refunded                        bool    `json:"Refunded"`
	Name                            string  `json:"Name"`
	Token                           string  `json:"Token"`
	SubscriptionId                  string  `json:"SubscriptionId"`
	GatewayName                     string  `json:"GatewayName"`
	ApplePay                        bool    `json:"ApplePay"`
	AndroidPay                      bool    `json:"AndroidPay"`
	WalletType                      string  `json:"WalletType"`
	TotalFee                        float64 `json:"TotalFee"`
}

// OrderResponse представляет ответ на запросы, связанные с заказами.
type OrderResponse struct {
	Success bool        `json:"Success"`
	Message string      `json:"Message"`
	Model   OrderModel  `json:"Model"`
}

// OrderModel содержит детали заказа.
type OrderModel struct {
	Id              string  `json:"Id"`
	Number          int     `json:"Number"`
	Amount          float64 `json:"Amount"`
	Currency        string  `json:"Currency"`
	CurrencyCode    int     `json:"CurrencyCode"`
	Email           string  `json:"Email"`
	Phone           string  `json:"Phone"`
	Description     string  `json:"Description"`
	RequireConfirmation bool `json:"RequireConfirmation"`
	Url             string  `json:"Url"`
	CultureName     string  `json:"CultureName"`
	CreatedDate     string  `json:"CreatedDate"`
	CreatedDateIso  string  `json:"CreatedDateIso"`
	PaymentDate     string  `json:"PaymentDate"`
	PaymentDateIso  string  `json:"PaymentDateIso"`
	StatusCode      int     `json:"StatusCode"`
	Status          string  `json:"Status"`
	InternalId      int     `json:"InternalId"`
}

// NotificationSettingsResponse представляет ответ с настройками уведомлений.
type NotificationSettingsResponse struct {
	Success bool               `json:"Success"`
	Message string             `json:"Message"`
	Model   NotificationSettings `json:"Model"`
}

// NotificationSettings содержит настройки уведомлений.
type NotificationSettings struct {
	IsEnabled  bool   `json:"IsEnabled"`
	Address    string `json:"Address"`
	HttpMethod string `json:"HttpMethod"`
	Encoding   string `json:"Encoding"`
	Format     string `json:"Format"`
}

// ApplePaySessionResponse представляет ответ с данными сессии Apple Pay.
type ApplePaySessionResponse struct {
	Success bool            `json:"Success"`
	Message string          `json:"Message"`
	Model   ApplePaySession `json:"Model"`
}

// ApplePaySession содержит данные сессии Apple Pay.
type ApplePaySession struct {
	EpochTimestamp             int64  `json:"epochTimestamp"`
	ExpiresAt                  int64  `json:"expiresAt"`
	MerchantSessionIdentifier  string `json:"merchantSessionIdentifier"`
	Nonce                      string `json:"nonce"`
	MerchantIdentifier         string `json:"merchantIdentifier"`
	DomainName                 string `json:"domainName"`
	DisplayName                string `json:"displayName"`
	Signature                  string `json:"signature"`
}