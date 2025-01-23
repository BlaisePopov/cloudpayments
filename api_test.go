package cloudpayments

import (
	"os"
	"testing"
	"time"
)

func TestClient_ChargeCryptogramPayment(t *testing.T) {
	client := NewClient(Config{
		ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
		PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
		Timeout:   30 * time.Second,
	})
	request := &CryptogramPaymentRequest{
		Amount:               100.0,
		Currency:             &CurrencyRUB,
		IpAddress:            "127.0.0.1",
		CardCryptogramPacket: "test_cryptogram",
	}
	response, err := client.ChargeCryptogramPayment(request)
	if err != nil {
		t.Fatalf("ChargeCryptogramPayment failed: %v", err)
	}
	if !response.Success {
		t.Errorf("Expected success response, got: %v", response)
	}
}

func TestClient_AuthorizeCryptogramPayment(t *testing.T) {
	client := NewClient(Config{
		ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
		PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
		Timeout:   30 * time.Second,
	})
	request := &CryptogramPaymentRequest{
		Amount:               100.0,
		Currency:             &CurrencyRUB,
		IpAddress:            "127.0.0.1",
		CardCryptogramPacket: "test_cryptogram",
	}
	response, err := client.AuthorizeCryptogramPayment(request)
	if err != nil {
		t.Fatalf("AuthorizeCryptogramPayment failed: %v", err)
	}
	if !response.Success {
		t.Errorf("Expected success response, got: %v", response)
	}
}

func TestClient_Confirm3DSPayment(t *testing.T) {
	client := NewClient(Config{
		ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
		PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
		Timeout:   30 * time.Second,
	})
	request := &Confirm3DSRequest{
		TransactionId: 12345,
		PaRes:         "test_pares",
	}
	response, err := client.Confirm3DSPayment(request)
	if err != nil {
		t.Fatalf("Confirm3DSPayment failed: %v", err)
	}
	if !response.Success {
		t.Errorf("Expected success response, got: %v", response)
	}
}

func TestClient_ChargeTokenPayment(t *testing.T) {
	client := NewClient(Config{
		ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
		PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
		Timeout:   30 * time.Second,
	})
	request := &TokenPaymentRequest{
		Amount:          100.0,
		Currency:        &CurrencyRUB,
		AccountId:       "test_account",
		Token:           "test_token",
		TrInitiatorCode: 1,
	}
	response, err := client.ChargeTokenPayment(request)
	if err != nil {
		t.Fatalf("ChargeTokenPayment failed: %v", err)
	}
	if !response.Success {
		t.Errorf("Expected success response, got: %v", response)
	}
}

func TestClient_AuthorizeTokenPayment(t *testing.T) {
	client := NewClient(Config{
		ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
		PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
		Timeout:   30 * time.Second,
	})
	request := &TokenPaymentRequest{
		Amount:          100.0,
		Currency:        &CurrencyRUB,
		AccountId:       "test_account",
		Token:           "test_token",
		TrInitiatorCode: 1,
	}
	response, err := client.AuthorizeTokenPayment(request)
	if err != nil {
		t.Fatalf("AuthorizeTokenPayment failed: %v", err)
	}
	if !response.Success {
		t.Errorf("Expected success response, got: %v", response)
	}
}

func TestClient_ConfirmPayment(t *testing.T) {
	client := NewClient(Config{
		ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
		PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
		Timeout:   30 * time.Second,
	})
	request := &ConfirmPaymentRequest{
		TransactionId: 12345,
		Amount:        100.0,
	}
	response, err := client.ConfirmPayment(request)
	if err != nil {
		t.Fatalf("ConfirmPayment failed: %v", err)
	}
	if !response.Success {
		t.Errorf("Expected success response, got: %v", response)
	}
}

func TestClient_VoidPayment(t *testing.T) {
	client := NewClient(Config{
		ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
		PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
		Timeout:   30 * time.Second,
	})
	request := &VoidPaymentRequest{
		TransactionId: 12345,
	}
	response, err := client.VoidPayment(request)
	if err != nil {
		t.Fatalf("VoidPayment failed: %v", err)
	}
	if !response.Success {
		t.Errorf("Expected success response, got: %v", response)
	}
}

func TestClient_RefundPayment(t *testing.T) {
	client := NewClient(Config{
		ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
		PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
		Timeout:   30 * time.Second,
	})
	request := &RefundPaymentRequest{
		TransactionId: 12345,
		Amount:        100.0,
	}
	response, err := client.RefundPayment(request)
	if err != nil {
		t.Fatalf("RefundPayment failed: %v", err)
	}
	if !response.Success {
		t.Errorf("Expected success response, got: %v", response)
	}
}

func TestClient_ChargeCryptogramPayout(t *testing.T) {
	client := NewClient(Config{
		ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
		PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
		Timeout:   30 * time.Second,
	})
	request := &CryptogramPayoutRequest{
		CardCryptogramPacket: "test_cryptogram",
		Amount:               100.0,
		Currency:             &CurrencyRUB,
	}
	response, err := client.ChargeCryptogramPayout(request)
	if err != nil {
		t.Fatalf("ChargeCryptogramPayout failed: %v", err)
	}
	if !response.Success {
		t.Errorf("Expected success response, got: %v", response)
	}
}

func TestClient_ChargeTokenPayout(t *testing.T) {
	client := NewClient(Config{
		ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
		PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
		Timeout:   30 * time.Second,
	})
	request := &TokenPayoutRequest{
		Token:     "test_token",
		Amount:    100.0,
		AccountId: "test_account",
		Currency:  &CurrencyRUB,
	}
	response, err := client.ChargeTokenPayout(request)
	if err != nil {
		t.Fatalf("ChargeTokenPayout failed: %v", err)
	}
	if !response.Success {
		t.Errorf("Expected success response, got: %v", response)
	}
}

func TestClient_Payment(t *testing.T) {
	client := NewClient(Config{
		ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
		PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
		Timeout:   30 * time.Second,
	})
	request := &GetPaymentRequest{
		TransactionId: 12345,
	}
	response, err := client.Payment(request)
	if err != nil {
		t.Fatalf("Payment failed: %v", err)
	}
	if !response.Success {
		t.Errorf("Expected success response, got: %v", response)
	}
}

func TestClient_FindPaymentByInvoiceID(t *testing.T) {
	client := NewClient(Config{
		ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
		PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
		Timeout:   30 * time.Second,
	})
	request := &FindPaymentRequest{
		InvoiceId: "test_invoice",
	}
	response, err := client.FindPaymentByInvoiceID(request)
	if err != nil {
		t.Fatalf("FindPaymentByInvoiceID failed: %v", err)
	}
	if !response.Success {
		t.Errorf("Expected success response, got: %v", response)
	}
}

func TestClient_PaymentsList(t *testing.T) {
	client := NewClient(Config{
		ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
		PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
		Timeout:   30 * time.Second,
	})
	request := &PaymentsListRequest{
		Date: "2025-01-01",
	}
	response, err := client.PaymentsList(request)
	if err != nil {
		t.Fatalf("PaymentsList failed: %v", err)
	}
	if !response.Success {
		t.Errorf("Expected success response, got: %v", response)
	}
}

func TestClient_PaymentsListV2(t *testing.T) {
	client := NewClient(Config{
		ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
		PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
		Timeout:   30 * time.Second,
	})
	request := &PaymentsListV2Request{
		CreatedDateGte: "2025-01-01",
		CreatedDateLte: "2025-01-01",
		PageNumber:     1,
	}
	response, err := client.PaymentsListV2(request)
	if err != nil {
		t.Fatalf("PaymentsListV2 failed: %v", err)
	}
	if !response.Success {
		t.Errorf("Expected success response, got: %v", response)
	}
}

func TestClient_TokensList(t *testing.T) {
	client := NewClient(Config{
		ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
		PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
		Timeout:   30 * time.Second,
	})
	request := &TokensListRequest{
		PageNumber: 1,
	}
	response, err := client.TokensList(request)
	if err != nil {
		t.Fatalf("TokensList failed: %v", err)
	}
	if !response.Success {
		t.Errorf("Expected success response, got: %v", response)
	}
}

func TestClient_CreateSubscription(t *testing.T) {
	client := NewClient(Config{
		ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
		PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
		Timeout:   30 * time.Second,
	})
	request := &SubscriptionCreateRequest{
		Token:               "test_token",
		AccountId:           "test_account",
		Description:         "test_description",
		Amount:              100.0,
		Currency:            CurrencyRUB,
		RequireConfirmation: true,
		StartDate:           "2025-01-01",
		Interval:            "Month",
		Period:              1,
	}
	response, err := client.CreateSubscription(request)
	if err != nil {
		t.Fatalf("CreateSubscription failed: %v", err)
	}
	if !response.Success {
		t.Errorf("Expected success response, got: %v", response)
	}
}

func TestClient_GetSubscription(t *testing.T) {
	client := NewClient(Config{
		ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
		PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
		Timeout:   30 * time.Second,
	})
	request := &SubscriptionGetRequest{
		Id: "test_subscription_id",
	}
	response, err := client.GetSubscription(request)
	if err != nil {
		t.Fatalf("GetSubscription failed: %v", err)
	}
	if !response.Success {
		t.Errorf("Expected success response, got: %v", response)
	}
}

func TestClient_GetSubscriptionsList(t *testing.T) {
	client := NewClient(Config{
		ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
		PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
		Timeout:   30 * time.Second,
	})
	request := &SubscriptionListRequest{
		AccountId: "test_account",
	}
	response, err := client.GetSubscriptionsList(request)
	if err != nil {
		t.Fatalf("GetSubscriptionsList failed: %v", err)
	}
	if !response.Success {
		t.Errorf("Expected success response, got: %v", response)
	}
}

func TestClient_UpdateSubscription(t *testing.T) {
	client := NewClient(Config{
		ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
		PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
		Timeout:   30 * time.Second,
	})
	request := &SubscriptionUpdateRequest{
		Id:          "test_subscription_id",
		Description: "updated_description",
		Amount:      200.0,
	}
	response, err := client.UpdateSubscription(request)
	if err != nil {
		t.Fatalf("UpdateSubscription failed: %v", err)
	}
	if !response.Success {
		t.Errorf("Expected success response, got: %v", response)
	}
}

func TestClient_CancelSubscription(t *testing.T) {
	client := NewClient(Config{
		ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
		PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
		Timeout:   30 * time.Second,
	})
	request := &SubscriptionCancelRequest{
		Id: "test_subscription_id",
	}
	response, err := client.CancelSubscription(request)
	if err != nil {
		t.Fatalf("CancelSubscription failed: %v", err)
	}
	if !response.Success {
		t.Errorf("Expected success response, got: %v", response)
	}
}

func TestClient_CreateOrder(t *testing.T) {
	client := NewClient(Config{
		ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
		PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
		Timeout:   30 * time.Second,
	})
	request := &LinkPaymentRequest{
		Amount:      100.0,
		Currency:    &CurrencyRUB,
		Description: "test_order",
		Email:       "test@example.com",
	}
	response, err := client.CreateOrder(request)
	if err != nil {
		t.Fatalf("CreateOrder failed: %v", err)
	}
	if !response.Success {
		t.Errorf("Expected success response, got: %v", response)
	}
}

func TestClient_CancelOrder(t *testing.T) {
	client := NewClient(Config{
		ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
		PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
		Timeout:   30 * time.Second,
	})
	request := &CancelOrderRequest{
		Id: "test_order_id",
	}
	response, err := client.CancelOrder(request)
	if err != nil {
		t.Fatalf("CancelOrder failed: %v", err)
	}
	if !response.Success {
		t.Errorf("Expected success response, got: %v", response)
	}
}

func TestClient_GetNotificationSettings(t *testing.T) {
	client := NewClient(Config{
		ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
		PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
		Timeout:   30 * time.Second,
	})
	response, err := client.GetNotificationSettings(NotificationTypePay)
	if err != nil {
		t.Fatalf("GetNotificationSettings failed: %v", err)
	}
	if !response.Success {
		t.Errorf("Expected success response, got: %v", response)
	}
}

func TestClient_UpdateNotificationSettings(t *testing.T) {
	client := NewClient(Config{
		ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
		PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
		Timeout:   30 * time.Second,
	})
	settings := &NotificationSettings{
		IsEnabled:  true,
		Address:    "test@example.com",
		HttpMethod: "POST",
		Encoding:   "UTF-8",
		Format:     "JSON",
	}
	response, err := client.UpdateNotificationSettings(NotificationTypePay, settings)
	if err != nil {
		t.Fatalf("UpdateNotificationSettings failed: %v", err)
	}
	if !response.Success {
		t.Errorf("Expected success response, got: %v", response)
	}
}

func TestClient_StartApplePaySession(t *testing.T) {
	client := NewClient(Config{
		ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
		PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
		Timeout:   30 * time.Second,
	})
	response, err := client.StartApplePaySession("https://validation.url", "https://payment.url")
	if err != nil {
		t.Fatalf("StartApplePaySession failed: %v", err)
	}
	if !response.Success {
		t.Errorf("Expected success response, got: %v", response)
	}
}
