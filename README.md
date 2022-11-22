# Chatty

Чат-сервер для обмена сообщениями между пользователями


## Содержание
- [Запуск](#запуск)
- [Предоставляемый API](#предоставляемый-api)
- [Функционал](#функционал)

## Запуск
- Склонировать репозиторий
```sh
git clone https://github.com/Gusyatnikova/chatty.git
```
- (Опционально) изменить переменные окружения в файле [`.env`](.env)
- Запустить сервис
```sh
docker-compose -f chatty/docker-compose.yml up
```
- Проверить состояние сервиса
```sh
curl http://0.0.0.0:8888/health
```

## Предоставляемый API
- [Swagger](http://0.0.0.0:8888/swagger/index.html)

## Функционал
- Работа с пользователем
  - [x] Зарегистрировать пользователя
  - [x] Авторизовать пользователя