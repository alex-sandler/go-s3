# GO S3 Service

HTTP-сервис для работы с файлами в S3-совместимом хранилище (MinIO) с поддержкой сжатия изображений.

## Описание

Сервис предоставляет REST API для:
- Загрузки файлов с автоматическим сжатием изображений
- Скачивания файлов
- Просмотра списка файлов
- Удаления файлов
- Проверки состояния сервиса

## Технологии

- **Go 1.23** - основной язык
- **Fiber v2** - HTTP веб-фреймворк
- **MinIO** - S3-совместимое объектное хранилище
- **Zap** - структурированное логирование
- **Docker & Docker Compose** - контейнеризация


## API Endpoints

| Метод | Путь | Описание |
|-------|------|----------|
| `GET` | `/health` | Проверка состояния сервиса |
| `POST` | `/upload` | Загрузка файла |
| `GET` | `/files/:name` | Скачивание файла |
| `GET` | `/files` | Список всех файлов |
| `DELETE` | `/files/:name` | Удаление файла |

## Конфигурация

Скопируйте содержимое `.env.example` в `.env` и настройте переменные окружения:

## Запуск

### С помощью Docker Compose (рекомендуется)

```bash
# Запуск всех сервисов
docker-compose up -d

# Просмотр логов
docker-compose logs -f s3-service

# Остановка сервисов
docker-compose down
```

### Локальная разработка

```bash
# Установка зависимостей
go mod tidy

# Запуск MinIO в Docker
docker-compose up -d minio

# Запуск приложения
go run cmd/go-s3/main.go
```

## Использование

### Загрузка файла

```bash
curl -X POST \
  -F "file=@images/filename.jpeg" \
  http://localhost:8080/upload
```

### Скачивание файла

```bash
curl http://localhost:8080/files/filename.jpg --output downloaded_file.jpeg
```

### Список файлов

```bash
curl http://localhost:8080/files
```

### Удаление файла

```bash
curl -X DELETE http://localhost:8080/files/filename.jpeg
```

### Проверка состояния

```bash
curl http://localhost:8080/health
```

## Особенности

- **Автоматическое сжатие изображений**: JPEG, PNG и другие форматы автоматически сжимаются при загрузке
- **Изменение размера**: Изображения масштабируются до максимальной ширины (настраивается через `IMAGE_MAX_WIDTH`)
- **Временные ссылки**: Поддержка pre-signed URLs для безопасного доступа к файлам
- **Структурированное логирование**: Все операции логируются с помощью Zap
- **Health checks**: Встроенная проверка состояния сервиса и зависимостей

## Разработка

### Сборка

```bash
go build -o bin/go-s3 cmd/go-s3/main.go
```

### Тестирование

```bash
go test ./...
```

### Линтинг

```bash
golangci-lint run
```

## MinIO Web Console

После запуска MinIO будет доступна веб-консоль:
- URL: http://localhost:9001
- Логин: значение `MINIO_ROOT_USER`
- Пароль: значение `MINIO_ROOT_PASSWORD`