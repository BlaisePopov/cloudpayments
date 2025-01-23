package cloudpayments

const (
	ReferToCardIssuer             = 5001 // Отказ эмитента проводить онлайн-операцию. Свяжитесь с вашим банком или воспользуйтесь другой картой.
	InvalidMerchant               = 5003 // Отказ эмитента проводить онлайн-операцию. Свяжитесь с вашим банком или воспользуйтесь другой картой.
	PickUpCard                    = 5004 // Карта потеряна. Свяжитесь с вашим банком или воспользуйтесь другой картой.
	DoNotHonor                    = 5005 // Отказ эмитента без объяснения причин.
	Error                         = 5006 // Отказ сети проводить операцию или неправильный CVV-код. Проверьте правильность введенных данных карты или воспользуйтесь другой картой.
	PickUpCardSpecialConditions   = 5007 // Карта потеряна. Свяжитесь с вашим банком или воспользуйтесь другой картой.
	InvalidTransaction            = 5012 // Карта не предназначена для онлайн-платежей. Воспользуйтесь другой картой или свяжитесь с банком, выпустившим карту.
	AmountError                   = 5013 // Слишком маленькая или слишком большая сумма операции. Проверьте корректность суммы.
	InvalidCardNumber             = 5014 // Некорректный номер карты. Проверьте правильность введенных данных карты или воспользуйтесь другой картой.
	NoSuchIssuer                  = 5015 // Эмитент не найден. Проверьте правильность введенных данных карты или воспользуйтесь другой картой.
	CustomerCancellation          = 5017 // Отказ по желанию держателя карты. Воспользуйтесь другой картой.
	TransactionError              = 5019 // Отказ эмитента без объяснения причин.
	FormatError                   = 5030 // Ошибка на стороне эквайера — неверно сформирована транзакция. Повторите попытку позже.
	BankNotSupportedBySwitch      = 5031 // Неизвестный эмитент карты. Воспользуйтесь другой картой.
	ExpiredCardPickup             = 5033 // Истек срок утери карты. Свяжитесь с вашим банком или воспользуйтесь другой картой.
	SuspectedFraud                = 5034 // Отказ эмитента — подозрение на мошенничество. Свяжитесь с вашим банком или воспользуйтесь другой картой.
	RestrictedCard                = 5036 // Карта не предназначена для платежей. Платежи для этой карты запрещены. Попробуйте другую карту.
	LostCard                      = 5041 // Карта потеряна. Свяжитесь с вашим банком или воспользуйтесь другой картой.
	StolenCard                    = 5043 // Карта украдена. Свяжитесь с вашим банком или воспользуйтесь другой картой.
	InsufficientFunds             = 5051 // Недостаточно средств. Недостаточно средств на карте.
	ExpiredCard                   = 5054 // Карта просрочена или неверно указан срок действия. Проверьте правильность введенных данных карты или воспользуйтесь другой картой.
	InvalidPIN                    = 5055 // Неверный PIN-код. Воспользуйтесь другой картой.
	TransactionNotPermitted       = 5057 // Ограничение на карте.
	TransactionNotPermittedToCard = 5058 // Транзакция не разрешена по карте. Свяжитесь с вашим банком или воспользуйтесь другой картой.
	SuspectedFraudDecline         = 5059 // Транзакция была отклонена банком по подозрению в мошенничестве. Свяжитесь с банком или воспользуйтесь другой картой.
	ExceedsApprovalAmount         = 5061 // Превышена сумма по карте. Превышение лимита оплаты по карте. Измените настройки лимита или оплатите другой картой.
	RestrictedCard2               = 5062 // Карта не предназначена для платежей. Платежи для этой карты запрещены. Попробуйте другую карту.
	SecurityViolation             = 5063 // Карта заблокирована из-за нарушений безопасности. Воспользуйтесь другой картой.
	ExceedWithdrawalFrequency     = 5065 // Превышен лимит операций по карте. Свяжитесь с вашим банком или воспользуйтесь другой картой.
	IncorrectCVV                  = 5082 // Неверный CVV-код. Неверно указан код CVV.
	Timeout                       = 5091 // Эмитент недоступен. Повторите попытку позже или воспользуйтесь другой картой.
	CannotReachNetwork            = 5092 // Эмитент недоступен. Повторите попытку позже или воспользуйтесь другой картой.
	SystemError                   = 5096 // Ошибка банка-эквайера или сети. Повторите попытку позже.
	GatewaysDoNotSupportCard      = 5113 // Шлюз не поддерживает эмитента неРФ. Воспользуйтесь другой картой.
	UnableToProcess               = 5204 // Операция не может быть обработана по прочим причинам. Свяжитесь с вашим банком или воспользуйтесь другой картой.
	AuthenticationFailed          = 5206 // 3-D Secure авторизация не пройдена. Свяжитесь с вашим банком или воспользуйтесь другой картой.
	AuthenticationUnavailable     = 5207 // 3-D Secure авторизация недоступна. Свяжитесь с вашим банком или воспользуйтесь другой картой.
	AntiFraud                     = 5300 // Лимиты эквайера на проведение операций. Воспользуйтесь другой картой.
	NoPhone                       = 5761 // Отсутствует номер телефона. Воспользуйтесь другой картой.
	InvalidPhone                  = 5762 // Некорректный номер телефона. Воспользуйтесь другой картой.
	DifferentPhones               = 5763 // Номер телефона в запросе отличается. Воспользуйтесь другой картой.
	SaveCardRefused               = 6001 // Отказ в привязке счета. Невозможно провести оплату.
	InvalidErrorCode              = 6002 // Некорректное значение code. Невозможно провести оплату.
	AgentIdNotExists              = 6003 // Значение agentId отсутствует в ОПКЦ СБП. Невозможно провести оплату.
	CurrencyNotValid              = 6004 // Невозможно провести операцию с валютой {currency}. Обратитесь в поддержку сайта.
	InvalidPaymentLink            = 6005 // Неверные данные платежной ссылки. Невозможно провести оплату.
	AmountCurrencyNotValid        = 6006 // amount и currency должны одновременно присутствовать или одновременно отсутствовать при qrcType = 01. Невозможно провести оплату.
	InvalidRedirectUrl            = 6007 // Некорректный формат redirectUrl. Невозможно провести оплату.
	InvalidParamsId               = 6008 // Неверные данные paramsId. Невозможно провести оплату.
	InvalidTtl                    = 6009 // Период использования динамической Платежной ссылки СБП в поле qrTtl выходит за допустимый диапазон. Невозможно провести оплату.
	SubscriptionNotFound          = 6010 // Привязка счета не найдена. Невозможно провести оплату.
	InvalidSubscription           = 6011 // Неверные данные привязки счета. Невозможно провести оплату.
	InvalidRefund                 = 6013 // Дублирование идентификатора запроса, назначаемого ТСП или Агентом ТСП, - agentRefundRequestId. Невозможно провести оплату.
	OriginalDetailsIncorrect      = 6014 // Параметры запроса отличаются от параметров исходной Операции СБП C2B. Невозможно провести оплату.
	TransactionNotFound           = 6015 // Запрашиваемая транзакция не найдена. Невозможно провести оплату.
	InvalidMemberId               = 6016 // Значение memberId отсутствует в ОПКЦ СБП. Невозможно провести оплату.
	IncorrectMCCategory           = 6017 // Выполнение операций в данной категории ТСП запрещено. Невозможно провести оплату.
	CrossTermi                    = 6025 // Неполный код. Требуется уточнение.
)

// Коды обработки ответов на check уведомление
const (
	CheckInvalidInvoiceId = 3001 // Обработчик уведомления вернул {"code":10}
	CheckInvalidAccountId = 3002 // Обработчик уведомления вернул {"code":11}
	CheckInvalidAmount    = 3003 // Обработчик уведомления вернул {"code":12}
	CheckOutOfDate        = 3004 // Обработчик уведомления вернул {"code":20}
	CheckFormatError      = 3005 // Обработчик уведомления вернул код, отличный от ожидаемого
	CheckUnavailable      = 3006 // Сервис недоступен
	CheckUnableToConnect  = 3007 // Обработчик не отвечает (404, 504, 508 и т.д.)
	CheckNotAccepted      = 3008 // Обработчик уведомления вернул {"code":13}
)

// Типы операций
type OperationType string

var (
	OperationTypePayment    OperationType = "Payment"    // Оплата
	OperationTypeRefund     OperationType = "Refund"     // Возврат
	OperationTypeCardPayout OperationType = "CardPayout" // Выплата на карту
)

// TransactionStatus представляет статус транзакции.
type TransactionStatus string

var (
	TransactionStatusAwaitingAuthentication TransactionStatus = "AwaitingAuthentication" // Ожидает аутентификации
	TransactionStatusAuthorized             TransactionStatus = "Authorized"             // Авторизована
	TransactionStatusCompleted              TransactionStatus = "Completed"              // Завершена
	TransactionStatusCancelled              TransactionStatus = "Cancelled"              // Отменена
	TransactionStatusDeclined               TransactionStatus = "Declined"               // Отклонена
)

type PaymentStatus string

var (
	PaymentStatusAwaitingAuthentication PaymentStatus = "AwaitingAuthentication" // Ожидание аутентификации
	PaymentStatusAuthorized             PaymentStatus = "Authorized"             // Авторизован
	PaymentStatusCompleted              PaymentStatus = "Completed"              // Завершен
	PaymentStatusCancelled              PaymentStatus = "Cancelled"              // Отменен
	PaymentStatusDeclined               PaymentStatus = "Declined"               // Отклонен
)

// Статусы подписок (рекуррент)
type SubscriptionStatus string

var (
	SubscriptionStatusActive    SubscriptionStatus = "Active"    // Подписка активна
	SubscriptionStatusPastDue   SubscriptionStatus = "PastDue"   // Подписка просрочена
	SubscriptionStatusCancelled SubscriptionStatus = "Cancelled" // Подписка отменена
	SubscriptionStatusRejected  SubscriptionStatus = "Rejected"  // Подписка отклонена
	SubscriptionStatusExpired   SubscriptionStatus = "Expired"   // Подписка завершена
)

// CurrencyCode представляет код валюты.
type CurrencyCode string

var (
	CurrencyKZT CurrencyCode = "KZT" // Казахстанский тенге
	CurrencyRUB CurrencyCode = "RUB" // Российский рубль
	CurrencyEUR CurrencyCode = "EUR" // Евро
	CurrencyUSD CurrencyCode = "USD" // Доллар США
	CurrencyGBP CurrencyCode = "GBP" // Фунт стерлингов
	CurrencyUAH CurrencyCode = "UAH" // Украинская гривна
	CurrencyBYR CurrencyCode = "BYR" // Белорусский рубль (устаревший)
	CurrencyBYN CurrencyCode = "BYN" // Белорусский рубль
	CurrencyAZN CurrencyCode = "AZN" // Азербайджанский манат
	CurrencyCHF CurrencyCode = "CHF" // Швейцарский франк
	CurrencyCZK CurrencyCode = "CZK" // Чешская крона
	CurrencyCAD CurrencyCode = "CAD" // Канадский доллар
	CurrencyPLN CurrencyCode = "PLN" // Польский злотый
	CurrencySEK CurrencyCode = "SEK" // Шведская крона
	CurrencyTRY CurrencyCode = "TRY" // Турецкая лира
	CurrencyCNY CurrencyCode = "CNY" // Китайский юань
	CurrencyINR CurrencyCode = "INR" // Индийская рупия
)

// FiscalDataOperator представляет операторы фискальных данных.
type FiscalDataOperator string

var (
	FiscalDataOperatorPeterService   FiscalDataOperator = "PeterService"   // ООО «ПС СТ»
	FiscalDataOperatorFirstOfd       FiscalDataOperator = "FirstOfd"       // Первый ОФД
	FiscalDataOperatorTaxcom         FiscalDataOperator = "Taxcom"         // Такском
	FiscalDataOperatorPlatformaOfd   FiscalDataOperator = "PlatformaOfd"   // Платформа ОФД
	FiscalDataOperatorOfdYa          FiscalDataOperator = "OfdYa"          // ОФД-Я
	FiscalDataOperatorOfdYandex      FiscalDataOperator = "OfdYandex"      // Яндекс ОФД
	FiscalDataOperatorGarantexpress  FiscalDataOperator = "Garantexpress"  // Гарант ОФД
	FiscalDataOperatorOfdAstralNalog FiscalDataOperator = "OfdAstralNalog" // Астрал ОФД
	FiscalDataOperatorSbis           FiscalDataOperator = "Sbis"           // Компания "ТЕНЗОР", ООО
)

// Коды временных зон
type TimeZone string

var (
	TimeZoneHST  TimeZone = "HST"  // (UTC-10:00) Гавайи
	TimeZoneAKST TimeZone = "AKST" // (UTC-09:00) Аляска
	TimeZonePST  TimeZone = "PST"  // (UTC-08:00) Тихоокеанское время (США и Канада)
	TimeZoneMST  TimeZone = "MST"  // (UTC-07:00) Горное время (США и Канада)
	TimeZoneCST  TimeZone = "CST"  // (UTC-06:00) Центральное время (США и Канада)
	TimeZoneEST  TimeZone = "EST"  // (UTC-05:00) Восточное время (США и Канада)
	TimeZoneAST  TimeZone = "AST"  // (UTC-04:00) Атлантическое время (Канада)
	TimeZoneBRT  TimeZone = "BRT"  // (UTC-03:00) Бразилия
	TimeZoneUTC  TimeZone = "UTC"  // (UTC) Время в формате UTC
	TimeZoneGMT  TimeZone = "GMT"  // (UTC) Дублин, Лиссабон, Лондон, Эдинбург
	TimeZoneCET  TimeZone = "CET"  // (UTC+01:00) Центральная Европа
	TimeZoneEET  TimeZone = "EET"  // (UTC+02:00) Восточная Европа
	TimeZoneMSK  TimeZone = "MSK"  // (UTC+03:00) Москва, Санкт-Петербург
	TimeZoneAZT  TimeZone = "AZT"  // (UTC+04:00) Баку
	TimeZoneAMT  TimeZone = "AMT"  // (UTC+04:00) Ереван
	TimeZoneSAMT TimeZone = "SAMT" // (UTC+04:00) Самара
	TimeZoneGET  TimeZone = "GET"  // (UTC+04:00) Тбилиси
	TimeZoneTJT  TimeZone = "TJT"  // (UTC+05:00) Ташкент
	TimeZoneYEKT TimeZone = "YEKT" // (UTC+05:00) Екатеринбург
	TimeZoneALMT TimeZone = "ALMT" // (UTC+06:00) Алматы
	TimeZoneNOVT TimeZone = "NOVT" // (UTC+06:00) Новосибирск
	TimeZoneKRAT TimeZone = "KRAT" // (UTC+07:00) Красноярск
	TimeZoneHKT  TimeZone = "HKT"  // (UTC+08:00) Гонконг
	TimeZoneIRKT TimeZone = "IRKT" // (UTC+08:00) Иркутск
	TimeZoneSGT  TimeZone = "SGT"  // (UTC+08:00) Сингапур
	TimeZoneULAT TimeZone = "ULAT" // (UTC+08:00) Улан-Батор
	TimeZoneYAKT TimeZone = "YAKT" // (UTC+09:00) Якутск
	TimeZoneVLAT TimeZone = "VLAT" // (UTC+10:00) Владивосток
	TimeZoneSAKT TimeZone = "SAKT" // (UTC+11:00) Чокурдах
	TimeZoneANAT TimeZone = "ANAT" // (UTC+12:00) Анадырь
)

// Типы уведомлений
type NotificationType string

// Коды ответа на запрос от cloudpayments [https://developers.cloudpayments.ru/#uvedomleniya]
var (
	NotificationTypeCheck     NotificationType = "Check"     // Check уведомление
	NotificationTypePay       NotificationType = "Pay"       // Pay уведомление
	NotificationTypeFail      NotificationType = "Fail"      // Fail уведомление
	NotificationTypeConfirm   NotificationType = "Confirm"   // Confirm уведомление
	NotificationTypeRefund    NotificationType = "Refund"    // Refund уведомление
	NotificationTypeRecurrent NotificationType = "Recurrent" // Recurrent уведомление
	NotificationTypeCancel    NotificationType = "Cancel"    // Cancel уведомление
)

const (
	SUCCESS            = 0  // Платеж может быть проведен
	UNKNOWN_INVOICE_ID = 10 // Неверный номер заказа
	INVALID_ACCOUNT_ID = 11 // Некорректный AccountId
	INVALID_AMOUNT     = 12 // Неверная сумма
	REJECTED           = 13 // Платеж не может быть принят
	EXPIRED            = 20 // Платеж просрочен
)

// LanguageCode представляет код языка и временной зоны.
type LanguageCode string

var (
	LanguageRussianRuRU   LanguageCode = "ru-RU" // Русский (MSK)
	LanguageEnglishEnUS   LanguageCode = "en-US" // Английский (CET)
	LanguageLatvianLv     LanguageCode = "lv"    // Латышский (CET)
	LanguageAzerbaijaniAz LanguageCode = "az"    // Азербайджанский (AZT)
	LanguageKazakhKk      LanguageCode = "kk"    // Русский (ALMT)
	LanguageUkrainianUk   LanguageCode = "uk"    // Украинский (EET)
	LanguagePolishPl      LanguageCode = "pl"    // Польский (CET)
	LanguageVietnameseVi  LanguageCode = "vi"    // Вьетнамский (ICT)
	LanguageTurkishTr     LanguageCode = "tr"    // Турецкий (TRT)
)
