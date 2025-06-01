DC = docker compose
EXEC = docker exec -it
LOGS = docker logs
ENV = --env-file .env
APP_FILE = docker-compose.yml
APP_CONTAINER = movie-api

.PHONY: app
app:
	${DC} -f ${APP_FILE} ${ENV} up -d

.PHONY: app-down
app-down:
	${DC} -f ${APP_FILE} down

.PHONY: app-logs
app-logs:
	${LOGS} ${APP_CONTAINER} -f

.PHONY: sqlc
sqlc:
	docker run --rm -v "/Users/User/Desktop/projects/GO/movie-api:/src" -w /src sqlc/sqlc generate


.PHONY: migrate
migrate:
	docker run -v /Users/admin/Desktop/Personal/GO/movie-api/migrations:/migrations migrate/migrate \
      -path=/migrations \
      -database "postgres://user:password@host.docker.internal:5404/movie-api?sslmode=disable" \
      up

.PHONY: downgrade
downgrade:
	docker run -v /Users/admin/Desktop/Personal/GO/movie-api/migrations:/migrations migrate/migrate \
          -path=/migrations \
          -database "postgres://user:password@host.docker.internal:5404/movie-api?sslmode=disable" \
          down -all