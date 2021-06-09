# Этапы выполения

## Создание основы HTTP-сервиса

## Добавление хендлеров и основных сущностей
- InitRoutes() добавление API-адресов, как они должны быть в задании
  pkg/handler/handler.go
- Создание сигнатур хендлеров
  /pkg/handler/user.go
  /pkg/handler/comment.go
- Добавлены основные сущности
  /user.go
  /common.go

## Порядок инициализации: Repository->Service->Handler (DI)
- Описаны интерфейсы
  pkg/service/service.go
  pkg/repository/repository.go
- Хендлер стал ловить сервис
  pkg/handler/handler.go
- В main() определена передача репозитория в сервис, а сервиса в хендлер
  cmd/main.go
  
## Разворачиваем БД в докере (dev-версия)
- Созданы файлы миграций командой
  migrate create -ext sql -dir ./schema -seq init
  Файлы
  schema/000001_init.down.sql
  schema/000001_init.up.sql
- Запустим докер и убедимся что контейнер с postgresql активен
  docker pull postgres
  docker run --name=app-name -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d postgres
  docker exec -it PID /bin/bash
- Запустим миграции
  migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
- В контейнере можно убедиться, что БД с таблицами созданы
```
docker ps
docker exec -it PID /bin/bash

psql -U postgres
\l              # список БД
\c postgres     # выбрать БД postgres
или
\d              # общий список таблиц
 
select * from users;
```

Запуск приложения (в консоли отобразятся слушаемые методы):
```
go run ./cmd/main.go
```

## 
- Логирование с logrus









Документация методов:
http://localhost:8000/swagger/index.html