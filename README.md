# ydb-client
Go client for ydb test. 
 * Create database with make
 * Generate data to database
 * Make index for data and execute operation.


 # Table generation

Код для создания таблиц и их наполнения случайными данными лежит в `app/cmd/datagen`.

Для запуска `datagen` выполните следующую комманду:

```bash
make build-datagen

./app/bin/datagen -dsn grpc://localhost:2136/local -emp-count 100000
```

## Info. Select from web:

```sql
 // http://localhost:8765/monitoring/
 PRAGMA TablePathPrefix("local/native/query");
 SELECT * FROM seasons;

```