.PHONY: ydb
ydb:
	docker run -d --rm --name ydb-local -h localhost \
  --platform linux/amd64 \
  -p 2135:2135 -p 2136:2136 -p 8765:8765 \
  -v ~/ydb_certs:/ydb_certs -v ~/ydb_data:/ydb_data \
  -e GRPC_TLS_PORT=2135 -e GRPC_PORT=2136 -e MON_PORT=8765 \
  -e YDB_USE_IN_MEMORY_PDISKS=true \
  --cpus="2" \
  --memory="4G" \
  cr.yandex/yc/yandex-docker-local-ydb:latest

.PHONY: ydb-stop
ydb-stop:
	docker kill ydb-local

.PHONY: ydb-clean
ydb-clean:
	rm -rf ~/ydb_certs/* \
  rm -rf ~/ydb_data/*

.PHONY: restart
restart: ydb-stop ydb-clean ydb

.PHONY: build-datagen
build-datagen:
	go build -o ./app/bin/datagen app/cmd/datagen

.PHONY: build-app
build-app:
	go build -o ./app/bin/employees app/cmd/employees

