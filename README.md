# Pokemon server

> Ваша задача написать локальный REST API сервер с локальной базой данных PostgreSQL для хранения покемонов. У сервера должен быть всего один эндпоинт `/pokemons`. Этот эндпоинт должен поддерживать следующие операции:
> - Добавить покемона
> - Получить список всех покемонов
> - Получить покемона по его ID

## API

- GET /pokemons
- GET /pokemons/:id
- POST /pokemons

Структура Pokemon описана в [listing/pokemon.go](./listing/pokemon.go)

## Запуск

```
docker compose build
docker compose up
```

## Тестирование

```
docker compose up pokemon-db
POSTGRES_DSN="host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable" && go test ./...
```
