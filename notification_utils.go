package cloudpayments

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net"
	"strings"
)

// NotificationIPRanges содержит список IP-адресов, с которых CloudPayments отправляет уведомления.
var NotificationIPRanges = []string{
	"185.98.81.0/28",
	"87.251.91.160/27",
	"46.46.175.96/27",
	"46.46.168.160/27",
	"162.55.174.97/32",
	"91.216.178.243/32",
}

// IsTrustedIP проверяет, что IP-адрес находится в списке доверенных диапазонов.
// Входные параметры: ip - IP-адрес.
// Возвращаемое значение: true, если IP-адрес доверенный, иначе false.
func IsTrustedIP(ip string) bool {
	// Убираем порт из IP-адреса, если он есть.
	ip = strings.Split(ip, ":")[0]

	for _, ipRange := range NotificationIPRanges {
		if isIPInRange(ip, ipRange) {
			return true
		}
	}
	return false
}

// VerifyHMAC проверяет подлинность и целостность уведомления с использованием HMAC. [https://developers.cloudpayments.ru/#proverka-uvedomleniy]
// Входные параметры:
//   - method: HTTP-метод (например, "POST" или "GET").
//   - body: тело запроса (для POST) или строка параметров (для GET).
//   - contentHMAC: значение заголовка Content-HMAC, генерируется из URL encoded параметров.
//   - apiSecret: секретный ключ API.
//
// Возвращаемое значение: ошибка, если проверка не прошла.
func VerifyHMAC(method, body, contentHMAC, apiSecret string) error {
	// Проверяем наличие хотя бы одного HMAC-заголовка.
	if contentHMAC == "" {
		return errors.New("missing contentHMAC")
	}

	// Определяем сообщение для вычисления HMAC.
	var message string
	if method == "POST" {
		message = body
	} else if method == "GET" {
		message = body
	} else {
		return errors.New("unsupported HTTP method")
	}

	// Вычисляем HMAC для сообщения.
	computedHMAC, err := computeHMAC(message, apiSecret)
	if err != nil {
		return err
	}

	// Сравниваем вычисленный HMAC с заголовком.
	if contentHMAC != "" && contentHMAC != computedHMAC {
		return errors.New("Content-HMAC does not match computed HMAC")
	}

	return nil
}

// computeHMAC вычисляет HMAC для сообщения с использованием API Secret.
// Входные параметры: message - сообщение, apiSecret - секретный ключ API.
// Возвращаемое значение: HMAC в формате base64 и ошибка, если она возникла.
func computeHMAC(message, apiSecret string) (string, error) {
	mac := hmac.New(sha256.New, []byte(apiSecret))
	_, err := mac.Write([]byte(message))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), nil
}

// isIPInRange проверяет, что IP-адрес находится в указанном диапазоне.
// Входные параметры: ip - IP-адрес, ipRange - диапазон IP-адресов в формате CIDR.
// Возвращаемое значение: true, если IP-адрес находится в диапазоне, иначе false.
func isIPInRange(ip, ipRange string) bool {
	ipAddr := net.ParseIP(ip)
	if ipAddr == nil {
		return false
	}

	_, ipNet, err := net.ParseCIDR(ipRange)
	if err != nil {
		return false
	}

	return ipNet.Contains(ipAddr)
}
