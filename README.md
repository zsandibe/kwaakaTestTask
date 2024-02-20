# Kwaaka test task

Тестовое задание можно посмотреть в  **TODO.MD**

EST API сервис на языке Go с использованием MongoDB или PostgreSQL. 

Используемые технологии:

+ Go
+ MongoDB
+ Gin
+ Docker (для запуска mongoDB)

Созданный сервис имеет чистую архитектуру, что обеспечивает простое расширение его возможностей и удобное тестирование. Также в нем реализован Graceful Shutdown для правильного завершения работы сервиса.

# Environment

```dotenv
SERVER_HOST=YOUR_SERVER_HOST
SERVER_PORT=YOUR_SERVER_PORT
MONGODB_URL=MONGODB_URL
NAME_DB=DBNAME
NAME_COLLECTION=COLLECTION_NAME
API_URL=API_URL
API_KEY=API_KEY
```


# Usage

Сперва нужно запустить контейнер  mongoDB:
```shell
make docker-run
```
Затем сам сервис:
```shell
make run
```
Чтобы остановить\удалить контейнер:
```shell
make docker-stop
```
```shell
make docker-rm
```

# Routes

* PUT /weather
  * Добавляет/обновляет инфу о погоде
  # Тело запроса
  ```json
    {
        "city":""
    }
  ```
  ```
* GET /weather
  * Возвращает инфу по городу
  ```query_params
    {
        "city":""
    }
  ```
  # Тело ответа
  ```json
    {
        "id": "",
        "city": "Nur-Sultan",
        "temperature": -27.029999,
        "created_at": "2024-02-20T17:40:04.596Z",
        "updated_at": "2024-02-20T17:40:04.596Z"
    }
  ```
Дополнительно добавил route getAllWeather для просмотра всех записей.

* GET /weatherLists
    * Возвращает все погоды
     # Тело ответа
    ```json
    [
    {
        "id": "",
        "city": "Nur-Sultan",
        "temperature": -27.029999,
        "created_at": "2024-02-20T17:40:04.596Z",
        "updated_at": "2024-02-20T17:40:04.596Z"
    },
    {
        "id": "",
        "city": "Aktobe",
        "temperature": -20.949997,
        "created_at": "2024-02-20T17:40:46.809Z",
        "updated_at": "2024-02-20T17:40:46.809Z"
    },
    {
        "id": "",
        "city": "Taraz",
        "temperature": -7.769989,
        "created_at": "2024-02-20T17:44:01.39Z",
        "updated_at": "2024-02-20T17:45:35.179Z"
    },
    {
        "id": "",
        "city": "Madrid",
        "temperature": 17.170013,
        "created_at": "2024-02-20T17:46:59.09Z",
        "updated_at": "2024-02-20T17:46:59.09Z"
    }
    ]
    ```


