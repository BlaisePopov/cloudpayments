package cloudpayments

// CryptogramPaymentRequest представляет запрос на оплату по криптограмме.
type CryptogramPaymentRequest struct {
	Amount               float64                `json:"Amount"`
	Currency             *CurrencyCode          `json:"Currency"`
	IpAddress            string                 `json:"IpAddress"`
	CardCryptogramPacket string                 `json:"CardCryptogramPacket"`
	Name                 string                 `json:"Name,omitempty"`
	PaymentUrl           string                 `json:"PaymentUrl,omitempty"`
	InvoiceId            string                 `json:"InvoiceId,omitempty"`
	Description          string                 `json:"Description,omitempty"`
	CultureName          *LanguageCode          `json:"CultureName,omitempty"`
	AccountId            string                 `json:"AccountId,omitempty"`
	Email                string                 `json:"Email,omitempty"`
	Payer                *Payer                 `json:"Payer,omitempty"`
	JsonData             map[string]interface{} `json:"JsonData,omitempty"`
	RequestID            *string                `json:"-"`
}

// TokenPaymentRequest представляет запрос на оплату по токену.
type TokenPaymentRequest struct {
	Amount           float64                `json:"Amount"`
	Currency         *CurrencyCode          `json:"Currency"`
	AccountId        string                 `json:"AccountId"`
	Token            string                 `json:"Token"`
	TrInitiatorCode  int                    `json:"TrInitiatorCode"`
	PaymentScheduled int                    `json:"PaymentScheduled,omitempty"`
	InvoiceId        string                 `json:"InvoiceId,omitempty"`
	Description      string                 `json:"Description,omitempty"`
	IpAddress        string                 `json:"IpAddress,omitempty"`
	Email            string                 `json:"Email,omitempty"`
	Payer            *Payer                 `json:"Payer,omitempty"`
	JsonData         map[string]interface{} `json:"JsonData,omitempty"`
	RequestID        *string                `json:"-"`
}

// Confirm3DSRequest представляет запрос на подтверждение 3DS оплаты.
type Confirm3DSRequest struct {
	TransactionId int64   `json:"TransactionId"`
	PaRes         string  `json:"PaRes"`
	RequestID     *string `json:"-"`
}

// ConfirmPaymentRequest представляет запрос на подтверждение оплаты.
type ConfirmPaymentRequest struct {
	TransactionId int64                  `json:"TransactionId"`
	Amount        float64                `json:"Amount"`
	JsonData      map[string]interface{} `json:"JsonData,omitempty"`
	RequestID     *string                `json:"-"`
}

// RefundPaymentRequest представляет запрос на возврат средств.
type RefundPaymentRequest struct {
	TransactionId int64                  `json:"TransactionId"`
	Amount        float64                `json:"Amount"`
	JsonData      map[string]interface{} `json:"JsonData,omitempty"`
	RequestID     *string                `json:"-"`
}

// VoidPaymentRequest представляет запрос на отмену оплаты.
type VoidPaymentRequest struct {
	TransactionId int64   `json:"TransactionId"`
	RequestID     *string `json:"-"`
}

// GetPaymentRequest представляет запрос на получение информации о платеже.
type GetPaymentRequest struct {
	TransactionId int64   `json:"TransactionId"`
	RequestID     *string `json:"-"`
}

// FindPaymentRequest представляет запрос на поиск платежа по идентификатору счета.
type FindPaymentRequest struct {
	InvoiceId string  `json:"InvoiceId"`
	RequestID *string `json:"-"`
}

// PaymentsListRequest представляет запрос на получение списка транзакций за день.
type PaymentsListRequest struct {
	Date      string  `json:"Date"`
	TimeZone  string  `json:"TimeZone,omitempty"`
	RequestID *string `json:"-"`
}

// PaymentsListV2Request представляет запрос на получение списка транзакций за произвольный период.
type PaymentsListV2Request struct {
	CreatedDateGte string   `json:"CreatedDateGte"`
	CreatedDateLte string   `json:"CreatedDateLte"`
	PageNumber     int      `json:"PageNumber"`
	TimeZone       string   `json:"TimeZone,omitempty"`
	Statuses       []string `json:"Statuses,omitempty"`
	RequestID      *string  `json:"-"`
}

// TokensListRequest представляет запрос на получение списка токенов.
type TokensListRequest struct {
	PageNumber int     `json:"PageNumber"`
	RequestID  *string `json:"-"`
}

// LinkPaymentRequest представляет запрос на создание заказа для отправки по email.
type LinkPaymentRequest struct {
	Amount               float64                `json:"Amount"`
	Currency             *CurrencyCode          `json:"Currency"`
	Description          string                 `json:"Description"`
	Email                string                 `json:"Email,omitempty"`
	RequireConfirmation  bool                   `json:"RequireConfirmation,omitempty"`
	SendEmail            bool                   `json:"SendEmail,omitempty"`
	InvoiceId            string                 `json:"InvoiceId,omitempty"`
	AccountId            string                 `json:"AccountId,omitempty"`
	OfferUri             string                 `json:"OfferUri,omitempty"`
	Phone                string                 `json:"Phone,omitempty"`
	SendSms              bool                   `json:"SendSms,omitempty"`
	SendViber            bool                   `json:"SendViber,omitempty"`
	CultureName          *LanguageCode          `json:"CultureName,omitempty"`
	SubscriptionBehavior string                 `json:"SubscriptionBehavior,omitempty"`
	SuccessRedirectUrl   string                 `json:"SuccessRedirectUrl,omitempty"`
	FailRedirectUrl      string                 `json:"FailRedirectUrl,omitempty"`
	JsonData             map[string]interface{} `json:"JsonData,omitempty"`
	RequestID            *string                `json:"-"`
}

// CancelOrderRequest представляет запрос на отмену заказа.
type CancelOrderRequest struct {
	Id        string  `json:"Id"`
	RequestID *string `json:"-"`
}

// SubscriptionCreateRequest представляет запрос на создание подписки.
type SubscriptionCreateRequest struct {
	Token               string                 `json:"Token"`
	AccountId           string                 `json:"AccountId"`
	Description         string                 `json:"Description"`
	Email               string                 `json:"Email,omitempty"`
	Amount              float64                `json:"Amount"`
	Currency            CurrencyCode           `json:"Currency"`
	RequireConfirmation bool                   `json:"RequireConfirmation"`
	StartDate           string                 `json:"StartDate"`
	Interval            string                 `json:"Interval"`
	Period              int                    `json:"Period"`
	MaxPeriods          int                    `json:"MaxPeriods,omitempty"`
	CustomerReceipt     map[string]interface{} `json:"CustomerReceipt,omitempty"`
	RequestID           *string                `json:"-"`
}

// SubscriptionUpdateRequest представляет запрос на обновление подписки.
type SubscriptionUpdateRequest struct {
	Id                  string                 `json:"Id"`
	Description         string                 `json:"Description,omitempty"`
	Amount              float64                `json:"Amount,omitempty"`
	Currency            *CurrencyCode          `json:"Currency,omitempty"`
	RequireConfirmation bool                   `json:"RequireConfirmation,omitempty"`
	StartDate           string                 `json:"StartDate,omitempty"`
	Interval            string                 `json:"Interval,omitempty"`
	Period              int                    `json:"Period,omitempty"`
	MaxPeriods          int                    `json:"MaxPeriods,omitempty"`
	CustomerReceipt     map[string]interface{} `json:"CustomerReceipt,omitempty"`
	CultureName         *LanguageCode          `json:"CultureName,omitempty"`
	RequestID           *string                `json:"-"`
}

// SubscriptionGetRequest представляет запрос на получение информации о подписке.
type SubscriptionGetRequest struct {
	Id        string  `json:"Id"`
	RequestID *string `json:"-"`
}

// SubscriptionCancelRequest представляет запрос на отмену подписки.
type SubscriptionCancelRequest struct {
	Id        string  `json:"Id"`
	RequestID *string `json:"-"`
}

// SubscriptionListRequest представляет запрос на получение списка подписок.
type SubscriptionListRequest struct {
	AccountId string  `json:"AccountId"`
	RequestID *string `json:"-"`
}

// CryptogramPayoutRequest представляет запрос на выплату по криптограмме.
type CryptogramPayoutRequest struct {
	CardCryptogramPacket string                 `json:"CardCryptogramPacket"`
	Amount               float64                `json:"Amount"`
	Currency             *CurrencyCode          `json:"Currency"`
	Name                 string                 `json:"Name,omitempty"`
	AccountId            string                 `json:"AccountId,omitempty"`
	Email                string                 `json:"Email,omitempty"`
	JsonData             map[string]interface{} `json:"JsonData,omitempty"`
	InvoiceId            string                 `json:"InvoiceId,omitempty"`
	Description          string                 `json:"Description,omitempty"`
	Payer                *Payer                 `json:"Payer,omitempty"`
	Receiver             *Receiver              `json:"Receiver,omitempty"`
	RequestID            *string                `json:"-"`
}

// TokenPayoutRequest представляет запрос на выплату по токену.
type TokenPayoutRequest struct {
	Token     string        `json:"Token"`
	Amount    float64       `json:"Amount"`
	AccountId string        `json:"AccountId"`
	Currency  *CurrencyCode `json:"Currency"`
	InvoiceId string        `json:"InvoiceId,omitempty"`
	Payer     *Payer        `json:"Payer,omitempty"`
	Receiver  *Receiver     `json:"Receiver,omitempty"`
	RequestID *string       `json:"-"`
}

// Payer представляет информацию о плательщике.
type Payer struct {
	FirstName  string `json:"FirstName,omitempty"`
	LastName   string `json:"LastName,omitempty"`
	MiddleName string `json:"MiddleName,omitempty"`
	Birth      string `json:"Birth,omitempty"`
	Street     string `json:"Street,omitempty"`
	Address    string `json:"Address,omitempty"`
	City       string `json:"City,omitempty"`
	Country    string `json:"Country,omitempty"`
	Phone      string `json:"Phone,omitempty"`
	Postcode   string `json:"Postcode,omitempty"`
}

// Receiver представляет информацию о получателе.
type Receiver struct {
	FirstName  string `json:"FirstName,omitempty"`
	LastName   string `json:"LastName,omitempty"`
	MiddleName string `json:"MiddleName,omitempty"`
	Birth      string `json:"Birth,omitempty"`
	Street     string `json:"Street,omitempty"`
	Address    string `json:"Address,omitempty"`
	City       string `json:"City,omitempty"`
	Country    string `json:"Country,omitempty"`
	Phone      string `json:"Phone,omitempty"`
	Postcode   string `json:"Postcode,omitempty"`
}
