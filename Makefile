include .env
export

export PROJECT_ROOT=${shell pwd}

env-up:
	@docker compose up -d project-postgres

env-down:
	@docker compose down project-postgres

env-cleanup:
	@read -p "Очистить все volume файлы окружения? Опастность утери данных. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down project-postgres port-forwarder && \
		sudo rm -rf ${PROJECT_ROOT}/out/pgdata && \
		echo "Файлы окружения очищены"; \
	else \
		echo "Очистка окружения отменена"; \
	fi

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder

redis-up:
	@docker compose up -d project-redis

redis-down:
	@docker compose down project-redis

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутствует необходимый параметр seq. Пример: make migrate-create seq=init"; \
		exit 1; \
	fi; \
	docker compose run --rm --user "$$(id -u):$$(id -g)" project-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Отсутствует необходимый параметр action. Пример: make migrate-action action='down 1'"; \
		exit 1; \
	fi; \
	docker compose run --rm project-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@project-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		$(action)

migrate-version:
	@docker compose run --rm project-postgres-migrate \
		-path /migrations \
		-database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@project-postgres:5432/${POSTGRES_DB}?sslmode=disable" \
		version


migrate-force:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутствует необходимый параметр seq. Пример: make migrate-force seq=4"; \
		exit 1; \
	fi; \
	if ! printf '%s\n' "$(seq)" | grep -Eq '^[0-9]+$$'; then \
		echo "Версия должна содержать только цифры: seq=$(seq)"; \
		exit 1; \
	fi; \
	docker compose run --rm project-postgres-migrate \
		-path /migrations \
		-database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@project-postgres:5432/${POSTGRES_DB}?sslmode=disable" \
		force $(seq)

test-db-up:
	@docker run -d --name test-postgres \
		-e POSTGRES_DB=test \
		-e POSTGRES_USER=test \
		-e POSTGRES_PASSWORD=test \
		-p 5433:5432 \
		postgres:16-alpine

test-db-down:
	@docker rm -f test-postgres

migrate-cycle-test:
	@make migrate-action-test action=up POSTGRES_USER=test POSTGRES_PASSWORD=test POSTGRES_DB=test && \
	make migrate-action-test action=down POSTGRES_USER=test POSTGRES_PASSWORD=test POSTGRES_DB=test && \
	make migrate-action-test action=up POSTGRES_USER=test POSTGRES_PASSWORD=test POSTGRES_DB=test

migrate-action-test:
	@if [ -z "$(action)" ]; then \
		echo "Отсутствует необходимый параметр action. Пример: make migrate-action action='down 1'"; \
		exit 1; \
	fi; \
	docker compose run --rm project-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@host.docker.internal:5433/${POSTGRES_DB}?sslmode=disable \
		$(action)

logs-cleanup:
	@read -p "Очистить все log файлы? Опастность утери логов. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		sudo rm -rf ${PROJECT_ROOT}/out/logs && \
		echo "Файлы логов очищены"; \
	else \
		echo "Очистка логов отменена"; \
	fi

project-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost && \
	go run ${PROJECT_ROOT}/cmd/learning_progress/main.go

progress-deploy:
	@docker compose up -d --build progress

progress-undeploy:
	@docker compose down progress

swagger-gen:
	@docker compose run --rm swagger \
		init \
		-g cmd/learning_progress/main.go \
		-o docs \
		--parseInternal \
		--parseDependency 

ps:
	@docker compose ps

make-memprofile:
	@if [ -z "$(action)" ]; then \
		echo "Отсутствует необходимый параметр action. Пример: make make-memprofile action='mem_profile'"; \
		exit 1; \
	fi; \
	go test -bench . -benchmem -memprofile=$(action).pprof

tool-pprof:
	@if [ -z "$(action)" ]; then \
		echo "Отсутствует необходимый параметр action. Пример: make tool-pprof action='mem_profile'"; \
		exit 1; \
	fi; \
	go tool pprof $(action).pprof

tool-trace:
	@if [ -z "$(action)" ]; then \
		echo "Отсутствует необходимый параметр action. Пример: make tool-trace action='trace_old'"; \
		exit 1; \
	fi; \
	go tool trace $(action).log

cover:
	@go test -short -count=1 -race -coverprofile=coverage.out ./...; \
	go tool cover -html=coverage.out -o coverage.html; \
	rm coverage.out

test:
	@go test -race ./...