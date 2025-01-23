package cloudpayments

import (
	"os"
	"testing"
	"time"
)

func TestClient_Ping(t *testing.T) {

	client := NewClient(Config{
		ApiSecret: os.Getenv("CLOUD_PAYMENTS_API_SECRET"),
		PublicId:  os.Getenv("CLOUD_PAYMENTS_PUBLIC_ID"),
		Timeout:   30 * time.Second,
	})

	response, err := client.Ping()

	if err != nil || response.Success != true {
		t.Error(err)
	}
}
