FROM golang:1.21.6-bullseye

WORKDIR /songs_api

COPY . ./

# Установка зависимостей проекта
RUN go mod download

# Сборка проекта
RUN go build -o ./bin/api ./cmd/songs_api/main.go

EXPOSE 5001

# Команда для запуска контейнера
CMD ["./bin/api"]