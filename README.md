# Calculator Web Service

Это простой веб-сервис, который вычисляет арифметические выражения, отправленные пользователем по HTTP-запросу.

## Оглавление

- [Описание](#описание)
- [Функциональность](#функциональность)
- [Требования](#требования)
- [Установка](#установка)
- [Запуск сервиса](#запуск-сервиса)
- [Использование](#использование)
  - [API Эндпоинт](#api-эндпоинт)
  - [Примеры запросов](#примеры-запросов)
- [Тестирование](#тестирование)

## Описание

Данное приложение представляет собой веб-сервис на языке Go, который принимает арифметические выражения через HTTP POST-запросы и возвращает результат их вычисления.

## Функциональность

- Вычисление арифметических выражений с поддержкой операций сложения, вычитания, умножения и деления.
- Правильная обработка приоритетов операций.
- Возвращает корректные HTTP-коды ответов в зависимости от результата обработки запроса.
- Обработка ошибок и исключительных ситуаций (например, деление на ноль, некорректный ввод).

## Требования

- Go версии 1.16 или выше
- Git для клонирования репозитория

## Установка

Клонируйте репозиторий на своё устройство:

```bash
git clone https://github.com/MoonLeeg/calc-app.git
cd calc-app
```
Запуск сервиса
Запустите сервер с помощью команды:
```bash

Copy
go run ./cmd/main.go
```
Сервер запустится на порту 8080.

Использование
API Эндпоинт
URL: /api/v1/calculate
Метод: POST
Заголовки:

Content-Type: application/json
Тело запроса:

json

Copy
{
  "expression": "ваше арифметическое выражение"
}
Успешный ответ:

Код состояния: 200 OK

Тело ответа:

json

Copy
{
  "result": "результат вычисления"
}
Ошибки:

Код состояния: 422 Unprocessable Entity

Тело ответа:

json

Copy
{
  "error": "Expression is not valid"
}
Код состояния: 500 Internal Server Error

Тело ответа:

json

Copy
{
  "error": "Internal server error"
}
Примеры запросов
Успешный запрос
bash

Copy
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression": "2+2*2"}'
Ответ:

json

Copy
{
  "result": "6"
}
Ошибка 422
Отправка некорректного выражения (например, с лишними символами):

bash

Copy
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression": "2++2"}'
Ответ:

json

Copy
{
  "error": "Expression is not valid"
}
Ошибка 500
Симуляция внутренней ошибки сервера (например, если вызвать панику в коде):

bash

Copy
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression": "simulate server error"}'
Ответ:

json

Copy
{
  "error": "Internal server error"
}
Примечание: Для достижения кода 500 вы можете временно изменить код сервера, чтобы искусственно вызвать ошибку.

Тестирование
Для запуска тестов выполните команду:

bash

Copy
go test ./...
Это выполнит все тесты в проекте и покажет результаты.
