CloudPayments Go Client

[![Go Reference](https://pkg.go.dev/badge/github.com/BlaisePopov/cloudpayments.svg)](https://pkg.go.dev/github.com/BlaisePopov/cloudpayments)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/BlaisePopov/cloudpayments/blob/main/LICENSE)

Go-клиент для работы с [API CloudPayments](https://developers.cloudpayments.ru/#api). Этот пакет предоставляет удобный интерфейс для взаимодействия с платежной системой [CloudPayments](https://cloudpayments.ru/), включая оплату по криптограмме, токену, управление подписками, возвраты и другие операции.

## Оглавление

- [Установка](#установка)
- [Использование](#использование)
- [Поддерживаемые методы](#поддерживаемые-методы)
- [Лицензия](#лицензия)

## Установка

Для установки пакета используйте команду:

```bash
go get github.com/BlaisePopov/cloudpayments
```

## Использование

### Инициализация клиента

Для начала работы с API CloudPayments необходимо создать клиент, передав в него конфигурацию:

```go
import (
    "github.com/BlaisePopov/cloudpayments"
    "time"
)

config := &cloudpayments.Config{
    PublicId:  "your_public_id",
    ApiSecret: "your_api_secret",
    Timeout:   30 * time.Second,
}

client := cloudpayments.NewClient(config)
```

### Пример оплаты по криптограмме

```go
request := &cloudpayments.CryptogramPaymentRequest{
    Amount:              100.0,
    Currency:            "RUB",
    IpAddress:           "123.123.123.123",
    CardCryptogramPacket: "your_cryptogram_packet",
    InvoiceId:           "12345",
    Description:         "Оплата заказа",
}

response, err := client.ChargeCryptogramPayment(request)
if err != nil {
    log.Fatalf("Ошибка при оплате: %v", err)
}

log.Printf("Ответ от API: %+v", response)
```

### Пример получения информации о платеже

```go
request := &cloudpayments.GetPaymentRequest{
    TransactionId: 123456,
}

response, err := client.Payment(request)
if err != nil {
    log.Fatalf("Ошибка при получении информации о платеже: %v", err)
}

log.Printf("Информация о платеже: %+v", response)
```

## Поддерживаемые методы

### Методы для работы с платежами
- **ChargeCryptogramPayment** — оплата по криптограмме.
- **AuthorizeCryptogramPayment** — двухстадийная оплата по криптограмме.
- **ChargeTokenPayment** — оплата по токену.
- **AuthorizeTokenPayment** — двухстадийная оплата по токену.
- **Confirm3DSPayment** — подтверждение 3DS оплаты.
- **ConfirmPayment** — подтверждение оплаты.
- **RefundPayment** — возврат средств.
- **VoidPayment** — отмена оплаты.
- **Payment** — получение информации о платеже.
- **FindPaymentByInvoiceID** — поиск платежа по идентификатору счета.
- **PaymentsList** — получение списка транзакций за день.
- **PaymentsListV2** — получение списка транзакций за произвольный период.

### Методы для работы с ордерами
- **CreateOrder** — создание заказа для отправки по email.
- **CancelOrder** — отмена заказа.

### Методы для работы с выплатами (Charge)
- **ChargeCryptogramPayout** — выплата по криптограмме.
- **ChargeTokenPayout** — выплата по токену.

### Методы для работы с подписками
- **CreateSubscription** — создание подписки.
- **UpdateSubscription** — обновление подписки.
- **CancelSubscription** — отмена подписки.
- **GetSubscription** — получение информации о подписке.
- **GetSubscriptionsList** — получение списка подписок.

### Методы для работы с токенами
- **TokensList** — получение списка токенов.

### Методы для работы с уведомлениями
- **GetNotificationSettings** — получение настроек уведомлений для указанного типа.
- **UpdateNotificationSettings** — обновление настроек уведомлений для указанного типа.

### Методы для работы с Apple Pay
- **StartApplePaySession** — запуск сессии для оплаты через Apple Pay.

### Утилиты для проверки уведомлений
- **IsTrustedIP** — проверка, что IP-адрес находится в списке доверенных диапазонов.
- **VerifyHMAC** — проверка подлинности и целостности уведомления с использованием HMAC.

### Дополнительные методы
- **Ping** — проверка доступности API.

## Лицензия

Этот проект распространяется под лицензией MIT. Подробнее см. в файле [LICENSE](LICENSE).

---