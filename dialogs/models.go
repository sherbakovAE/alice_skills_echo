package dialogs

type (
	// Question содержит информацию о входящем запросе пользователя.
	Question struct {
		// Информация об устройстве, с помощью которого пользователь
		// разговаривает с Алисой.
		Meta Meta `json:"meta"`

		// Данные, полученные от пользователя.
		Request Request `json:"request"`

		// Данные о сессии.
		Session Session `json:"session"`

		// Версия протокола.
		Version string `json:"version"`
	}

	// Answer представляет собой ответ пользователю.
	Answer struct {
		// Версия протокола.
		Version string `json:"version"`

		// Данные о сессии.
		Session Session `json:"session"`

		// Данные для ответа пользователю.
		Response Response `json:"response"`
	}

	// Response содержит данные для ответа пользователю.
	Response struct {
		// Текст, который следует показать пользователю.
		Text string `json:"text"`

		// Ответ в формате TTS (text-to-speech).
		TTS string `json:"tts"`

		// Кнопки, которые следует показать пользователю.
		Buttons []Button `json:"buttons"`

		// Признак конца разговора (получая этот флаг в ответе, Алиса
		// автоматически завершает работу навыка в приложении).
		EndSession bool `json:"end_session"`
	}

	// Button представляет собой произвольную кнопку в диалоге.
	Button struct {
		// Текст кнопки.
		Title string `json:"title"`

		// Произвольный JSON, который Яндекс.Диалоги должны отправить
		// обработчику, если данная кнопка будет нажата.
		Payload Payload `json:"payload,omitempty"`

		// URL, который должна открывать кнопка.
		URL string `json:"url"`

		// Признак того, что кнопку нужно убрать после следующего запроса
		// пользователя. Допустимые значения:
		// * false — кнопка должна оставаться активной (значение по умолчанию);
		// * true — кнопку нужно скрывать после нажатия.
		Hide bool `json:"hide"`
	}

	// Meta содержит информацию об устройстве, с помощью которого пользователь
	// разговаривает с Алисой.
	Meta struct {
		// Язык в POSIX-формате.
		Locale string `json:"locale"`

		// Название часового пояса, включая алиасы.
		TimeZone string `json:"timezone"`

		// Идентификатор устройства и приложения, в котором идет разговор.
		ClientID string `json:"client_id"`
	}

	// Request содержит данные, полученные от пользователя.
	Request struct {
		// Тип ввода:
		// * SimpleUtterance — голосовой ввод;
		// * ButtonPressed — нажатие кнопки.
		Type string `json:"type"`

		// Формальные характеристики реплики, которые удалось выделить
		// Яндекс.Диалогам. Отсутствует, если ни одно из вложенных свойств не
		// применимо.
		Markup Markup `json:"markup,omitempty"`

		// Текст пользовательского запроса без активационных фраз Алисы и
		// конкретного навыка.
		Command string `json:"command"`

		// Полный текст пользовательского запроса.
		OriginalUtterance string `json:"original_utterance"`

		// JSON, полученный с нажатой кнопкой от обработчика навыка (в ответе на
		// предыдущий запрос).
		Payload Payload `json:"payload,omitempty"`

		// Слова и именованные сущности, которые Диалоги извлекли из запроса пользователя.
		Nlu Nlu `json:"nlu,omitempty"`
	}
	Nlu struct {
		// Массив слов из произнесенной пользователем фразы.
		Tokens []string `json:"tokens,omitempty"`
		// Массив именованных сущностей.
		Entities []struct {
			//Обозначение начала и конца именованной сущности в массиве слов. Нумерация слов в массиве начинается с 0.
			Tokens struct {
				Start int `json:"start,omitempty"` // Первое слово именованной сущности
				End   int `json:"end,omitempty"`   //Первое слово после именованной сущности.
			} `json:"tokens"`
			Type string `json:"type,omitempty"` // Тип именованной сущности. Возможные значения:
			//Value string  `json:"value,omitempty"`

			//YANDEX.DATETIME — дата и время, абсолютные или относительные.
			//YANDEX.FIO — фамилия, имя и отчество.
			//YANDEX.GEO — местоположение (адрес или аэропорт).
			//YANDEX.NUMBER — число, целое или с плавающей точкой.
		} `json:"entities,omitempty"`
	}
	// Markup содержит формальные характеристики реплики, которые удалось
	// выделить Яндекс.Диалогам. Отсутствует, если ни одно из вложенных свойств
	// не применимо.
	Markup struct {
		// Признак реплики, которая содержит криминальный подтекст (самоубийство,
		// разжигание ненависти, угрозы). Вы можете настроить навык на
		// определенную реакцию для таких случаев — например, отвечать "Не
		// понимаю, о чем вы. Пожалуйста, переформулируйте вопрос."
		//
		// Возможно только значение true. Если признак не применим, это свойство
		// не включается в ответ.
		DangerousContext bool `json:"dangerous_context,omitempty"`
	}

	// Session содержит данные о сессии.
	Session struct {
		// Признак новой сессии. Возможные значения:
		// * true — пользователь начал новый разговор с навыком;
		// * false — запрос отправлен в рамках уже начатого разговора.
		New bool `json:"new"`

		// Уникальный идентификатор сессии, 64 байта.
		SessionID string `json:"session_id"`

		// Идентификатор сообщения в рамках сессии. Инкрементируется с каждым
		// следующим запросом.
		MessageID int64 `json:"message_id"`

		// Идентификатор вызываемого навыка.
		SkillID string `json:"skill_id"`

		// Обфусцированный идентификатор пользователя.
		UserID string `json:"user_id"`
	}

	// Payload представляет собой произвольные JSON данные, идущие c кнопкой.
	Payload interface{}
)

const (
	// TypeSimpleUtterance является идентификатором события голосового ввода
	TypeSimpleUtterance = "SimpleUtterance"

	// TypeButtonPressed является идентификатором события нажатия на кнопку
	TypeButtonPressed = "ButtonPressed"
)
