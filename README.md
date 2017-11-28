# gomarkets

(*в активной разработке, не использовать*)

Библиотека для торговли на биржах. Разрабатывается для высоко-скоростных приложений.

## Connector

* Универсальный интерфейс к биржам
* Для каждой биржи свое расширение connector
* Используется REST API и WebSocket API
* Существует два распространненых типа подключения к биржам. REST(HTTP) API и WebSockets. REST API не подразумевает постоянного подключения. WebSockets соединение поднимается автоматически

## OrderBook

* Запускается с указанием connector
* Подписывается на обновление книги по websockets
* Используются числа с фиксированной точностью <https://github.com/shopspring/decimal>
* Используется skiplist <http://github.com/huandu/skiplist>
* Стаканы разряженный и сжатый (разряженный под вопросом)

## ChartData

* График

## TODO

* Примеры
* Skiplist uses global rand source in math/rand by default. This rand source acquires a lock when generating random number. Replacing it with a lock-free rand source can provide slightly better performance. Use SkipList.SetRandSource to change rand source