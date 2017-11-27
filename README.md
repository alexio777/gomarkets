# gomarkets

Библиотека торговли из Go (*в активной разработке, не использовать*)

## Connector

* Универсальный интерфейс к биржам
* Для каждой биржи свое расширение connector
* Используется REST API и WebSocket API
* Существует два распространненых типа подключения к биржам. REST(HTTP) API и WebSockets. REST API не подразумевает постоянного подключения. WebSockets соединение поднимается  автоматически.

## OrderBook

* Запускается с указанием connector
* Подписывается на обновление книги по websockets
* Используются числа с абсолютной точностью (https://github.com/shopspring/decimal)
* Стаканы разряженный и сжатый (разряженный под вопросом)

## ChartData

* График

## TODO

* Примеры
* Skiplist uses global rand source in math/rand by default. This rand source acquires a lock when generating random number. Replacing it with a lock-free rand source can provide slightly better performance. Use SkipList.SetRandSource to change rand source.