package cloudpayments

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const (
	baseURL = "https://api.cloudpayments.ru/"
)

var (
	// ErrInvalidResponse возвращается, когда ответ от API не может быть разобран.
	ErrInvalidResponse = errors.New("invalid response from API")
)

// sendRequest отправляет HTTP-запрос к API CloudPayments.
// Входные параметры: endpoint - конечная точка API, params - данные запроса в формате JSON, requestID - идентификатор запроса.
// Возвращаемое значение: тело ответа и ошибка, если она возникла.
func (c *Client) sendRequest(endpoint string, params []byte, requestID *string) ([]byte, error) {
	req, err := http.NewRequest("POST", baseURL+endpoint, bytes.NewBuffer(params))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.config.PublicId, c.config.ApiSecret)
	req.Header.Set("Content-Type", "application/json")

	if requestID != nil {
		req.Header.Set("X-Request-ID", *requestID)
	}

	client := &http.Client{Timeout: c.config.Timeout}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Ping проверяет доступность API CloudPayments. [https://developers.cloudpayments.ru/#testovyy-metod]
// Возвращаемое значение: ответ от API и ошибка, если она возникла.
func (c *Client) Ping() (*BaseResponse, error) {
	var data *BaseResponse
	response, err := c.sendRequest("test", nil, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &data); err != nil {
		return nil, ErrInvalidResponse
	}

	return data, nil
}

// ChargeCryptogramPayment выполняет оплату по криптограмме. [https://developers.cloudpayments.ru/#oplata-po-kriptogramme]
// Входные параметры: cpr - запрос на оплату по криптограмме.
// Возвращаемое значение: ответ на оплату и ошибка, если она возникла.
func (c *Client) ChargeCryptogramPayment(cpr *CryptogramPaymentRequest) (*PaymentResponse, error) {
	paymentResponse := &PaymentResponse{}

	params, err := json.Marshal(cpr)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("payments/cards/charge", params, cpr.RequestID)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &paymentResponse); err != nil {
		return nil, ErrInvalidResponse
	}

	return paymentResponse, nil
}

// AuthorizeCryptogramPayment выполняет двухстадийную оплату по криптограмме. [https://developers.cloudpayments.ru/#oplata-po-kriptogramme]
// Входные параметры: cpr - запрос на авторизацию по криптограмме.
// Возвращаемое значение: ответ на авторизацию и ошибка, если она возникла.
func (c *Client) AuthorizeCryptogramPayment(cpr *CryptogramPaymentRequest) (*Payment3DSResponse, error) {
	paymentResponse := &Payment3DSResponse{}

	params, err := json.Marshal(cpr)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("payments/cards/auth", params, cpr.RequestID)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &paymentResponse); err != nil {
		return nil, ErrInvalidResponse
	}

	return paymentResponse, nil
}

// Confirm3DSPayment подтверждает оплату с использованием 3DS. [https://developers.cloudpayments.ru/#obrabotka-3-d-secure]
// Входные параметры: confirm3DS - запрос на подтверждение 3DS оплаты.
// Возвращаемое значение: ответ на подтверждение и ошибка, если она возникла.
func (c *Client) Confirm3DSPayment(confirm3DS *Confirm3DSRequest) (*PaymentResponse, error) {
	paymentResponse := &PaymentResponse{}

	params, err := json.Marshal(confirm3DS)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("payments/cards/post3ds", params, confirm3DS.RequestID)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &paymentResponse); err != nil {
		return nil, ErrInvalidResponse
	}

	return paymentResponse, nil
}

// ChargeTokenPayment выполняет оплату по токену. [https://developers.cloudpayments.ru/#oplata-po-tokenu-rekarring]
// Входные параметры: tpr - запрос на оплату по токену.
// Возвращаемое значение: ответ на оплату и ошибка, если она возникла.
func (c *Client) ChargeTokenPayment(tpr *TokenPaymentRequest) (*PaymentResponse, error) {
	paymentResponse := &PaymentResponse{}

	params, err := json.Marshal(tpr)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("payments/tokens/charge", params, tpr.RequestID)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &paymentResponse); err != nil {
		return nil, ErrInvalidResponse
	}

	return paymentResponse, nil
}

// AuthorizeTokenPayment выполняет авторизацию оплаты по токену. [https://developers.cloudpayments.ru/#oplata-po-tokenu-rekarring]
// Входные параметры: tpr - запрос на авторизацию по токену.
// Возвращаемое значение: ответ на авторизацию и ошибка, если она возникла.
func (c *Client) AuthorizeTokenPayment(tpr *TokenPaymentRequest) (*PaymentResponse, error) {
	paymentResponse := &PaymentResponse{}

	params, err := json.Marshal(tpr)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("payments/tokens/auth", params, tpr.RequestID)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &paymentResponse); err != nil {
		return nil, ErrInvalidResponse
	}

	return paymentResponse, nil
}

// ConfirmPayment подтверждает авторизованную оплату. [https://developers.cloudpayments.ru/#podtverzhdenie-oplaty]
// Входные параметры: confirm - запрос на подтверждение оплаты.
// Возвращаемое значение: базовый ответ и ошибка, если она возникла.
func (c *Client) ConfirmPayment(confirm *ConfirmPaymentRequest) (*BaseResponse, error) {
	baseResponse := &BaseResponse{}

	params, err := json.Marshal(confirm)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("payments/confirm", params, confirm.RequestID)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &baseResponse); err != nil {
		return nil, ErrInvalidResponse
	}

	return baseResponse, nil
}

// VoidPayment отменяет оплату. [https://developers.cloudpayments.ru/#otmena-oplaty]
// Входные параметры: vpr - запрос на отмену оплаты.
// Возвращаемое значение: базовый ответ и ошибка, если она возникла.
func (c *Client) VoidPayment(vpr *VoidPaymentRequest) (*BaseResponse, error) {
	baseResponse := &BaseResponse{}

	params, err := json.Marshal(vpr)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("payments/void", params, vpr.RequestID)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &baseResponse); err != nil {
		return nil, ErrInvalidResponse
	}

	return baseResponse, nil
}

// RefundPayment выполняет возврат средств. [https://developers.cloudpayments.ru/#vozvrat-deneg]
// Входные параметры: rpp - запрос на возврат средств.
// Возвращаемое значение: базовый ответ и ошибка, если она возникла.
func (c *Client) RefundPayment(rpp *RefundPaymentRequest) (*BaseResponse, error) {
	baseResponse := &BaseResponse{}

	params, err := json.Marshal(rpp)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("payments/refund", params, rpp.RequestID)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &baseResponse); err != nil {
		return nil, ErrInvalidResponse
	}

	return baseResponse, nil
}

// ChargeCryptogramPayout выполняет выплату по криптограмме. [https://developers.cloudpayments.ru/#vyplata-po-kriptogramme]
// Входные параметры: cpr - запрос на выплату по криптограмме.
// Возвращаемое значение: ответ на выплату и ошибка, если она возникла.
func (c *Client) ChargeCryptogramPayout(cpr *CryptogramPayoutRequest) (*PaymentResponse, error) {
	paymentResponse := &PaymentResponse{}

	params, err := json.Marshal(cpr)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("payments/cards/topup", params, cpr.RequestID)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &paymentResponse); err != nil {
		return nil, ErrInvalidResponse
	}

	return paymentResponse, nil
}

// ChargeTokenPayout выполняет выплату по токену. [https://developers.cloudpayments.ru/#vyplata-po-tokenu]
// Входные параметры: tpr - запрос на выплату по токену.
// Возвращаемое значение: ответ на выплату и ошибка, если она возникла.
func (c *Client) ChargeTokenPayout(tpr *TokenPayoutRequest) (*PaymentResponse, error) {
	paymentResponse := &PaymentResponse{}

	params, err := json.Marshal(tpr)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("payments/token/topup", params, tpr.RequestID)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &paymentResponse); err != nil {
		return nil, ErrInvalidResponse
	}

	return paymentResponse, nil
}

// Payment получает информацию о платеже по идентификатору транзакции. [https://developers.cloudpayments.ru/#prosmotr-tranzaktsii]
// Входные параметры: gpr - запрос на получение информации о платеже.
// Возвращаемое значение: ответ с информацией о платеже и ошибка, если она возникла.
func (c *Client) Payment(gpr *GetPaymentRequest) (*PaymentResponse, error) {
	paymentResponse := &PaymentResponse{}

	params, err := json.Marshal(gpr)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("payments/get", params, gpr.RequestID)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &paymentResponse); err != nil {
		return nil, ErrInvalidResponse
	}

	return paymentResponse, nil
}

// FindPaymentByInvoiceID ищет платеж по идентификатору счета. [https://developers.cloudpayments.ru/#proverka-statusa-platezha]
// Входные параметры: fpr - запрос на поиск платежа.
// Возвращаемое значение: ответ с информацией о платеже и ошибка, если она возникла.
func (c *Client) FindPaymentByInvoiceID(fpr *FindPaymentRequest) (*PaymentResponse, error) {
	paymentResponse := &PaymentResponse{}

	params, err := json.Marshal(fpr)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("v2/payments/find", params, fpr.RequestID)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &paymentResponse); err != nil {
		return nil, ErrInvalidResponse
	}

	return paymentResponse, nil
}

// PaymentsList получает список транзакций за день. [https://developers.cloudpayments.ru/#vygruzka-spiska-tranzaktsiy]
// Входные параметры: plr - запрос на получение списка транзакций.
// Возвращаемое значение: ответ с информацией о транзакциях и ошибка, если она возникла.
func (c *Client) PaymentsList(plr *PaymentsListRequest) (*PaymentsListResponse, error) {
	list := &PaymentsListResponse{}

	params, err := json.Marshal(plr)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("payments/list", params, plr.RequestID)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &list); err != nil {
		return nil, ErrInvalidResponse
	}

	return list, nil
}

// PaymentsListV2 получает список транзакций за произвольный период. [https://developers.cloudpayments.ru/#vygruzka-spiska-tranzaktsiy-za-proizvolnyy-period]
// Входные параметры: plr - запрос на получение списка транзакций.
// Возвращаемое значение: ответ с информацией о транзакциях и ошибка, если она возникла.
func (c *Client) PaymentsListV2(plr *PaymentsListV2Request) (*PaymentsListV2Response, error) {
	list := &PaymentsListV2Response{}

	params, err := json.Marshal(plr)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("v2/payments/list", params, plr.RequestID)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &list); err != nil {
		return nil, ErrInvalidResponse
	}

	return list, nil
}

// TokensList получает список токенов. [https://developers.cloudpayments.ru/#vygruzka-tokenov]
// Входные параметры: tlr - запрос на получение списка токенов.
// Возвращаемое значение: ответ с информацией о токенах и ошибка, если она возникла.
func (c *Client) TokensList(tlr *TokensListRequest) (*TokensListResponse, error) {
	list := &TokensListResponse{}

	params, err := json.Marshal(tlr)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("payments/tokens/list", params, tlr.RequestID)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &list); err != nil {
		return nil, ErrInvalidResponse
	}

	return list, nil
}

// CreateSubscription создает подписку. [https://developers.cloudpayments.ru/#sozdanie-podpiski-na-rekurrentnye-platezhi]
// Входные параметры: scr - запрос на создание подписки.
// Возвращаемое значение: ответ с информацией о подписке и ошибка, если она возникла.
func (c *Client) CreateSubscription(scr *SubscriptionCreateRequest) (*SubscriptionResponse, error) {
	subscriptionResponse := &SubscriptionResponse{}

	params, err := json.Marshal(scr)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("subscriptions/create", params, scr.RequestID)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &subscriptionResponse); err != nil {
		return nil, ErrInvalidResponse
	}

	return subscriptionResponse, nil
}

// GetSubscription получает информацию о подписке по идентификатору транзакции. [https://developers.cloudpayments.ru/#zapros-informatsii-o-podpiske]
// Входные параметры: sgr - запрос на получение информации о подписке.
// Возвращаемое значение: ответ с информацией о подписке и ошибка, если она возникла.
func (c *Client) GetSubscription(sgr *SubscriptionGetRequest) (*SubscriptionResponse, error) {
	subscriptionResponse := &SubscriptionResponse{}

	params, err := json.Marshal(sgr)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("subscriptions/get", params, sgr.RequestID)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &subscriptionResponse); err != nil {
		return nil, ErrInvalidResponse
	}

	return subscriptionResponse, nil
}

// GetSubscriptionsList получает список подписок по идентификатору аккаунта. [https://developers.cloudpayments.ru/#poisk-podpisok]
// Входные параметры: slr - запрос на получение списка подписок.
// Возвращаемое значение: ответ с информацией о подписках и ошибка, если она возникла.
func (c *Client) GetSubscriptionsList(slr *SubscriptionListRequest) (*SubscriptionsListResponse, error) {
	list := &SubscriptionsListResponse{}

	params, err := json.Marshal(slr)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("subscriptions/find", params, slr.RequestID)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &list); err != nil {
		return nil, ErrInvalidResponse
	}

	return list, nil
}

// UpdateSubscription обновляет подписку. [https://developers.cloudpayments.ru/#izmenenie-podpiski-na-rekurrentnye-platezhi]
// Входные параметры: sur - запрос на обновление подписки.
// Возвращаемое значение: ответ с информацией о подписке и ошибка, если она возникла.
func (c *Client) UpdateSubscription(sur *SubscriptionUpdateRequest) (*SubscriptionResponse, error) {
	subscriptionResponse := &SubscriptionResponse{}

	params, err := json.Marshal(sur)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("subscriptions/update", params, sur.RequestID)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &subscriptionResponse); err != nil {
		return nil, ErrInvalidResponse
	}

	return subscriptionResponse, nil
}

// CancelSubscription отменяет подписку. [https://developers.cloudpayments.ru/#otmena-podpiski-na-rekurrentnye-platezhi]
// Входные параметры: sur - запрос на отмену подписки.
// Возвращаемое значение: базовый ответ и ошибка, если она возникла.
func (c *Client) CancelSubscription(sur *SubscriptionCancelRequest) (*BaseResponse, error) {
	baseResponse := &BaseResponse{}

	params, err := json.Marshal(sur)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("subscriptions/cancel", params, sur.RequestID)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &baseResponse); err != nil {
		return nil, ErrInvalidResponse
	}

	return baseResponse, nil
}

// CreateOrder создает заказ для отправки по email. [https://developers.cloudpayments.ru/#sozdanie-scheta-dlya-otpravki-po-pochte]
// Входные параметры: lpr - запрос на создание заказа.
// Возвращаемое значение: базовый ответ и ошибка, если она возникла.
func (c *Client) CreateOrder(lpr *LinkPaymentRequest) (*OrderResponse, error) {
	orderResponse := &OrderResponse{}

	params, err := json.Marshal(lpr)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("orders/create", params, lpr.RequestID)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &orderResponse); err != nil {
		return nil, ErrInvalidResponse
	}

	return orderResponse, nil
}

// CancelOrder отменяет созданный заказ. [https://developers.cloudpayments.ru/#otmena-sozdannogo-scheta]
// Входные параметры: cor - запрос на отмену заказа.
// Возвращаемое значение: базовый ответ и ошибка, если она возникла.
func (c *Client) CancelOrder(cor *CancelOrderRequest) (*BaseResponse, error) {
	baseResponse := &BaseResponse{}

	params, err := json.Marshal(cor)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("orders/cancel", params, cor.RequestID)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &baseResponse); err != nil {
		return nil, ErrInvalidResponse
	}

	return baseResponse, nil
}

// GetNotificationSettings получает настройки уведомлений для указанного типа. [https://developers.cloudpayments.ru/#prosmotr-nastroek-uvedomleniy]
// Входные параметры: notificationType - тип уведомления (например, "pay", "fail").
// Возвращаемое значение: настройки уведомлений и ошибка, если она возникла.
func (c *Client) GetNotificationSettings(notificationType NotificationType) (*NotificationSettingsResponse, error) {
	notificationSettings := &NotificationSettingsResponse{}

	response, err := c.sendRequest("site/notifications/"+string(notificationType)+"/get", nil, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &notificationSettings); err != nil {
		return nil, ErrInvalidResponse
	}

	return notificationSettings, nil
}

// UpdateNotificationSettings обновляет настройки уведомлений для указанного типа. [https://developers.cloudpayments.ru/#izmenenie-nastroek-uvedomleniy]
// Входные параметры: notificationType - тип уведомления, settings - новые настройки.
// Возвращаемое значение: базовый ответ и ошибка, если она возникла.
func (c *Client) UpdateNotificationSettings(notificationType NotificationType, settings *NotificationSettings) (*BaseResponse, error) {
	baseResponse := &BaseResponse{}

	params, err := json.Marshal(settings)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("site/notifications/"+string(notificationType)+"/update", params, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &baseResponse); err != nil {
		return nil, ErrInvalidResponse
	}

	return baseResponse, nil
}

// StartApplePaySession запускает сессию для оплаты через Apple Pay. [https://developers.cloudpayments.ru/#zapusk-sessii-dlya-oplaty-cherez-apple-pay]
// Входные параметры: validationUrl - URL для валидации сессии, paymentUrl - URL для оплаты.
// Возвращаемое значение: данные сессии Apple Pay и ошибка, если она возникла.
func (c *Client) StartApplePaySession(validationUrl, paymentUrl string) (*ApplePaySessionResponse, error) {
	applePaySession := &ApplePaySessionResponse{}

	requestBody := map[string]string{
		"ValidationUrl": validationUrl,
		"PaymentUrl":    paymentUrl,
	}

	params, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	response, err := c.sendRequest("applepay/startsession", params, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &applePaySession); err != nil {
		return nil, ErrInvalidResponse
	}

	return applePaySession, nil
}
