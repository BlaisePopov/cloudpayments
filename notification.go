package cloudpayments

import (
	"time"
)

// CheckNotification представляет уведомление типа Check. [https://developers.cloudpayments.ru/#check]
type CheckNotification struct {
	TransactionId     int64   `json:"TransactionId"`
	Amount            float64 `json:"Amount"`
	Currency          string  `json:"Currency"`
	PaymentAmount     string  `json:"PaymentAmount"`
	PaymentCurrency   string  `json:"PaymentCurrency"`
	DateTime          string  `json:"DateTime"` // Формат: yyyy-MM-dd HH:mm:ss
	CardId            string  `json:"CardId,omitempty"`
	CardFirstSix      string  `json:"CardFirstSix"`
	CardLastFour      string  `json:"CardLastFour"`
	CardType          string  `json:"CardType"` // Visa, Mastercard, Maestro, "МИР"
	CardExpDate       string  `json:"CardExpDate"`
	TestMode          int     `json:"TestMode"` // 1 или 0
	Status            string  `json:"Status"`   // Completed или Authorized
	OperationType     string  `json:"OperationType"`
	InvoiceId         string  `json:"InvoiceId,omitempty"`
	AccountId         string  `json:"AccountId,omitempty"`
	SubscriptionId    string  `json:"SubscriptionId,omitempty"`
	TokenRecipient    string  `json:"TokenRecipient,omitempty"`
	Name              string  `json:"Name,omitempty"`
	Email             string  `json:"Email,omitempty"`
	IpAddress         string  `json:"IpAddress,omitempty"`
	IpCountry         string  `json:"IpCountry,omitempty"`
	IpCity            string  `json:"IpCity,omitempty"`
	IpRegion          string  `json:"IpRegion,omitempty"`
	IpDistrict        string  `json:"IpDistrict,omitempty"`
	IpLatitude        string  `json:"IpLatitude,omitempty"`
	IpLongitude       string  `json:"IpLongitude,omitempty"`
	Issuer            string  `json:"Issuer,omitempty"`
	IssuerBankCountry string  `json:"IssuerBankCountry,omitempty"`
	Description       string  `json:"Description,omitempty"`
	CardProduct       string  `json:"CardProduct,omitempty"`
	PaymentMethod     string  `json:"PaymentMethod,omitempty"`
	Data              string  `json:"Data,omitempty"` // JSON-строка
}

// PayNotification представляет уведомление типа Pay. [https://developers.cloudpayments.ru/#pay]
type PayNotification struct {
	TransactionId                         int64   `json:"TransactionId"`
	Amount                                float64 `json:"Amount"`
	Currency                              string  `json:"Currency"`
	PaymentAmount                         string  `json:"PaymentAmount"`
	PaymentCurrency                       string  `json:"PaymentCurrency"`
	DateTime                              string  `json:"DateTime"` // Формат: yyyy-MM-dd HH:mm:ss
	CardId                                string  `json:"CardId,omitempty"`
	CardFirstSix                          string  `json:"CardFirstSix"`
	CardLastFour                          string  `json:"CardLastFour"`
	CardType                              string  `json:"CardType"` // Visa, Mastercard, Maestro, "МИР"
	CardExpDate                           string  `json:"CardExpDate"`
	TestMode                              int     `json:"TestMode"` // 1 или 0
	Status                                string  `json:"Status"`   // Completed или Authorized
	OperationType                         string  `json:"OperationType"`
	GatewayName                           string  `json:"GatewayName"`
	InvoiceId                             string  `json:"InvoiceId,omitempty"`
	AccountId                             string  `json:"AccountId,omitempty"`
	SubscriptionId                        string  `json:"SubscriptionId,omitempty"`
	Name                                  string  `json:"Name,omitempty"`
	Email                                 string  `json:"Email,omitempty"`
	IpAddress                             string  `json:"IpAddress,omitempty"`
	IpCountry                             string  `json:"IpCountry,omitempty"`
	IpCity                                string  `json:"IpCity,omitempty"`
	IpRegion                              string  `json:"IpRegion,omitempty"`
	IpDistrict                            string  `json:"IpDistrict,omitempty"`
	IpLatitude                            string  `json:"IpLatitude,omitempty"`
	IpLongitude                           string  `json:"IpLongitude,omitempty"`
	Issuer                                string  `json:"Issuer,omitempty"`
	IssuerBankCountry                     string  `json:"IssuerBankCountry,omitempty"`
	Description                           string  `json:"Description,omitempty"`
	AuthCode                              string  `json:"AuthCode,omitempty"`
	Data                                  string  `json:"Data,omitempty"` // JSON-строка
	Token                                 string  `json:"Token,omitempty"`
	TotalFee                              float64 `json:"TotalFee,omitempty"`
	CardProduct                           string  `json:"CardProduct,omitempty"`
	PaymentMethod                         string  `json:"PaymentMethod,omitempty"`
	FallBackScenarioDeclinedTransactionId int64   `json:"FallBackScenarioDeclinedTransactionId,omitempty"`
	Rrn                                   string  `json:"Rrn,omitempty"`
	CustomFields                          string  `json:"CustomFields,omitempty"` // JSON-строка
}

// FailNotification представляет уведомление типа Fail. [https://developers.cloudpayments.ru/#fail]
type FailNotification struct {
	TransactionId                         int64   `json:"TransactionId"`
	Amount                                float64 `json:"Amount"`
	Currency                              string  `json:"Currency"`
	PaymentAmount                         string  `json:"PaymentAmount"`
	PaymentCurrency                       string  `json:"PaymentCurrency"`
	DateTime                              string  `json:"DateTime"` // Формат: yyyy-MM-dd HH:mm:ss
	CardId                                string  `json:"CardId,omitempty"`
	CardFirstSix                          string  `json:"CardFirstSix"`
	CardLastFour                          string  `json:"CardLastFour"`
	CardType                              string  `json:"CardType"` // Visa, Mastercard, Maestro, "МИР"
	CardExpDate                           string  `json:"CardExpDate"`
	TestMode                              int     `json:"TestMode"` // 1 или 0
	Reason                                string  `json:"Reason"`
	ReasonCode                            int     `json:"ReasonCode"`
	OperationType                         string  `json:"OperationType"`
	InvoiceId                             string  `json:"InvoiceId,omitempty"`
	AccountId                             string  `json:"AccountId,omitempty"`
	SubscriptionId                        string  `json:"SubscriptionId,omitempty"`
	Name                                  string  `json:"Name,omitempty"`
	Email                                 string  `json:"Email,omitempty"`
	IpAddress                             string  `json:"IpAddress,omitempty"`
	IpCountry                             string  `json:"IpCountry,omitempty"`
	IpCity                                string  `json:"IpCity,omitempty"`
	IpRegion                              string  `json:"IpRegion,omitempty"`
	IpDistrict                            string  `json:"IpDistrict,omitempty"`
	IpLatitude                            string  `json:"IpLatitude,omitempty"`
	IpLongitude                           string  `json:"IpLongitude,omitempty"`
	Issuer                                string  `json:"Issuer,omitempty"`
	IssuerBankCountry                     string  `json:"IssuerBankCountry,omitempty"`
	Description                           string  `json:"Description,omitempty"`
	Data                                  string  `json:"Data,omitempty"` // JSON-строка
	Token                                 string  `json:"Token,omitempty"`
	PaymentMethod                         string  `json:"PaymentMethod,omitempty"`
	FallBackScenarioDeclinedTransactionId int64   `json:"FallBackScenarioDeclinedTransactionId,omitempty"`
	Rrn                                   string  `json:"Rrn,omitempty"`
	CustomFields                          string  `json:"CustomFields,omitempty"` // JSON-строка
}

// ConfirmNotification представляет уведомление типа Confirm. [https://developers.cloudpayments.ru/#confirm]
type ConfirmNotification struct {
	TransactionId     int64   `json:"TransactionId"`
	Amount            float64 `json:"Amount"`
	Currency          string  `json:"Currency"`
	PaymentAmount     string  `json:"PaymentAmount"`
	PaymentCurrency   string  `json:"PaymentCurrency"`
	DateTime          string  `json:"DateTime"` // Формат: yyyy-MM-dd HH:mm:ss
	CardFirstSix      string  `json:"CardFirstSix"`
	CardLastFour      string  `json:"CardLastFour"`
	CardType          string  `json:"CardType"` // Visa, Mastercard, Maestro, "МИР"
	CardExpDate       string  `json:"CardExpDate"`
	TestMode          int     `json:"TestMode"` // 1 или 0
	Status            string  `json:"Status"`   // Completed
	InvoiceId         string  `json:"InvoiceId,omitempty"`
	AccountId         string  `json:"AccountId,omitempty"`
	SubscriptionId    string  `json:"SubscriptionId,omitempty"`
	Name              string  `json:"Name,omitempty"`
	Email             string  `json:"Email,omitempty"`
	IpAddress         string  `json:"IpAddress,omitempty"`
	IpCountry         string  `json:"IpCountry,omitempty"`
	IpCity            string  `json:"IpCity,omitempty"`
	IpRegion          string  `json:"IpRegion,omitempty"`
	IpDistrict        string  `json:"IpDistrict,omitempty"`
	IpLatitude        string  `json:"IpLatitude,omitempty"`
	IpLongitude       string  `json:"IpLongitude,omitempty"`
	Issuer            string  `json:"Issuer,omitempty"`
	IssuerBankCountry string  `json:"IssuerBankCountry,omitempty"`
	Description       string  `json:"Description,omitempty"`
	AuthCode          string  `json:"AuthCode,omitempty"`
	Data              string  `json:"Data,omitempty"` // JSON-строка
	Token             string  `json:"Token,omitempty"`
	PaymentMethod     string  `json:"PaymentMethod,omitempty"`
	Rrn               string  `json:"Rrn,omitempty"`
}

// RefundNotification представляет уведомление типа Refund. [https://developers.cloudpayments.ru/#refund]
type RefundNotification struct {
	TransactionId        int64   `json:"TransactionId"`
	PaymentTransactionId int64   `json:"PaymentTransactionId"`
	Amount               float64 `json:"Amount"`
	DateTime             string  `json:"DateTime"` // Формат: yyyy-MM-dd HH:mm:ss
	OperationType        string  `json:"OperationType"`
	InvoiceId            string  `json:"InvoiceId,omitempty"`
	AccountId            string  `json:"AccountId,omitempty"`
	Email                string  `json:"Email,omitempty"`
	Data                 string  `json:"Data,omitempty"` // JSON-строка
	Rrn                  string  `json:"Rrn,omitempty"`
	CustomFields         string  `json:"CustomFields,omitempty"` // JSON-строка
}

// RecurrentNotification представляет уведомление типа Recurrent. [https://developers.cloudpayments.ru/#recurrent]
type RecurrentNotification struct {
	Id                           string    `json:"Id"`
	AccountId                    string    `json:"AccountId"`
	Description                  string    `json:"Description"`
	Email                        string    `json:"Email"`
	Amount                       float64   `json:"Amount"`
	Currency                     string    `json:"Currency"`
	RequireConfirmation          bool      `json:"RequireConfirmation"`
	StartDate                    time.Time `json:"StartDate"`
	Interval                     string    `json:"Interval"` // Week, Month
	Period                       int       `json:"Period"`
	Status                       string    `json:"Status"`
	SuccessfulTransactionsNumber int       `json:"SuccessfulTransactionsNumber"`
	FailedTransactionsNumber     int       `json:"FailedTransactionsNumber"`
	MaxPeriods                   int       `json:"MaxPeriods,omitempty"`
	LastTransactionDate          time.Time `json:"LastTransactionDate,omitempty"`
	NextTransactionDate          time.Time `json:"NextTransactionDate,omitempty"`
}

// CancelNotification представляет уведомление типа Cancel. [https://developers.cloudpayments.ru/#cancel]
type CancelNotification struct {
	TransactionId int64   `json:"TransactionId"`
	Amount        float64 `json:"Amount"`
	DateTime      string  `json:"DateTime"` // Формат: yyyy-MM-dd HH:mm:ss
	InvoiceId     string  `json:"InvoiceId,omitempty"`
	AccountId     string  `json:"AccountId,omitempty"`
	Email         string  `json:"Email,omitempty"`
	Data          string  `json:"Data,omitempty"` // JSON-строка
	Rrn           string  `json:"Rrn,omitempty"`
}
