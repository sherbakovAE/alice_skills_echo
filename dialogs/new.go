package dialogs

import (
	"../logging"
	json "github.com/pquerna/ffjson/ffjson"
	"net/http"
	"strings"
	"github.com/labstack/echo"

)

type (
	// Questions является вебхук-каналом входящих запросов от пользователя.
	Questions <-chan Question

	// Answers является вебхук-каналом исходящих ответов к пользователям.
	Answers chan Answer
)
var log = logging.GetInstance()
// New создаёт простой роутер для прослушивания входящих данных по вебхуку и
// возвращает два канала: для чтения запросов и отправки ответов соответственно.
// и функцию обработчик для передачи в роут
func New() (Questions, Answers, func(c echo.Context) error)  {

	var err error
	questions := make(chan Question)
	answers := make(chan Answer)

	handleFunc := func(c echo.Context) error {
		log.Println("Тело входящего запроса:")
		//log.Println(c.Request().Body)

		log.Println("Декодируем запрос...")
		var question Question
		if err := c.Bind(&question); err != nil {
			return err
		}
		log.Println("Отправляем запрос в канал...",questions)
		questions <- question

		var answer Answer
		for answer = range answers {
			a := answer.Session
			q := question.Session
			if !strings.EqualFold(a.SessionID, q.SessionID) ||
				!strings.EqualFold(a.UserID, q.UserID) ||
				a.MessageID != q.MessageID {
				log.Println("Это не тот ответ...")
				continue
			}

			log.Println("Обнаружен подходящий запрос! Отвечаем...")
			break
		}

		log.Println("Дождались нужный ответ! Отправляем его...")
		c.Response().Header().Set("Content-Type", "application/json")
		c.Response().WriteHeader(http.StatusOK)

		log.Println("Кодируем ответ...")
		if err = json.NewEncoder(c.Response()).Encode(answer); err != nil {
			log.Println("Ошибка:", err.Error())
			c.Response().WriteHeader(http.StatusInternalServerError)
			return err
		}
		return err
	}
	return questions, answers, handleFunc
}

