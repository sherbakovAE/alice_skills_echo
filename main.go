package main

import (
	"database/sql"
	"os"

	"./dialogs"
	"./logging"
	"./skills/aecho"
	mathem "./skills/matematica"
	memory "./skills/memory"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)



var (
	answer dialogs.Answer
	//textButton, urlButton, dataButton dialogs.Button
	//buttons                           []dialogs.Button
)

type State int

var db *sql.DB
var err error

var logConfig = middleware.LoggerConfig{Skipper: middleware.DefaultSkipper,
	Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
		`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
		`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
		`"bytes_out":${bytes_out}}` + "\n",
	Output: os.Stdout,
}

var numbers = map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}
var log = logging.GetInstance()



func main() {

	log.Debugf("Стартуем!..")

	// создание и запуск сервера
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(logConfig))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	//подключение тестового файла для yandex-dialogs-client (https://github.com/popstas/yandex-dialogs-client)
	//e.File("/alice/scenarios.yml", "static/audiomemory.yml")

	//создать каналы вопросов и ответов для каждого навыка


	//// запуск навыка Эхо
	var handlerEcho func(c echo.Context) error
	handlerEcho, Echo := aecho.Run()

	go Echo.Start()
	e.POST("/echo", handlerEcho)

	//// запуск навыка "Счёт в уме"
	var handlerMath func(c echo.Context) error
	handlerMath, Mathem := mathem.Run()
	go Mathem.Start()
	e.POST("/math", handlerMath)


	//// запуск навыка "Повторение слов"
	var handlerMemory func(c echo.Context) error
	handlerMemory, Memory := memory.Run()

	// определение используемой БД для использования в навыке
	db, err := sql.Open("sqlite3", "words.db")
	if err != nil {
		log.Fatal(err.Error())
	} else {
		Memory.DB = db
	}
	defer db.Close()

	go Memory.Start()
	e.POST("/memory", handlerMemory)

	e.Logger.Fatal(e.Start(":1323"))
}
