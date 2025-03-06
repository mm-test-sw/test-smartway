Сборка прокета:
1. ``docker compose up``
2. дождаться запуска сервиса
3. ``docker exec smartway_server goose -dir migrations postgres "user=root password=rpass dbname=smartway_db host=smartway_pg port=5432 sslmode=disable" up``

Сервис поддерживает запросы:
1. POST **/airlines** - Добавление авиакомпании <br/>
Пример: { "code": "TE", "name": "Test" }
2. DELETE **/airlines/{code}** - Удаление авиакомпании по code <br/>
Пример: airlines/TE
3. POST **/providers** - Добавление поставщика <br/>
Пример: { "id": "TE", "name": "Test" }
4. DELETE **/providers/{id}** - Удаление поставщика по id <br/>
Пример: /providers/TE
5. PUT **/airlines/providers** - Изменение списка поставщиков (указывается список id) у указанной
   авиакомпании (указывается code) <br/>
Пример: { "code": "TE", "providersId": ["IF", "RS"] }
6. POST **/schemas** - Добавление схемы поиска <br/>
Пример: { "id": 5, "name": "Test" }
7. GET **/schemas/{name}** - Поиск схемы поиска по названию <br/>
Пример: /schemas/Test
8. PATCH **/schemas** - Изменение схемы поиска <br/>
Пример: { "id": 5, "name": "TesT2", "providers": ["TT", "TY"] }
9. DELETE **schemas/{id}** - Удаление схемы поиска по id <br/>
Пример: /schemas/5
10. POST **/accounts** - Добавление аккаунта <br/>
Пример: { "id": 666, "schemaId": 1 }
11. PUT **/accounts** - Изменение схемы, назначенной аккаунту <br/>
Пример: { "id": 666, "schemaId": 2 }
12. DELETE **/accounts/{id}** - Удаление аккаунта по Id <br/>
Пример: /accounts/666
13. GET **/accounts/{id}/airlines** - Получение списка авиакомпаний по id аккаунта <br/>
Пример: /accounts/3/airlines
14. GET **/providers/{id}/airlines** - Получение списка авиакомпаний по id поставщика <br/>
Пример: /providers/RS/airlines


test edit
