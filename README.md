# GO S3 Service

HTTP-сервис для работы с JPEG-файлами в S3-совместимом хранилище (MinIO) с поддержкой сжатия изображений.

## Описание

Сервис предоставляет REST API для управления JPEG-файлами в S3-совместимом хранилище. Основные функции:

- Загрузка отдельных JPEG-файлов с автоматическим сжатием.
- Загрузка пачек JPEG-файлов с сжатием и без.
- Скачивание отдельных JPEG-файлов.
- Скачивание пачек JPEG-файлов в ZIP-архиве с сжатием и без.
- Просмотр списка файлов.
- Удаление файлов.
- Проверка состояния сервиса.

Для тестирования загрузки файлов используйте папку `images` в корне проекта для всех методов.

## Технологии

- **Go 1.23** — основной язык программирования.
- **Fiber v2** — HTTP веб-фреймворк.
- **MinIO** — S3-совместимое объектное хранилище.
- **Zap** — структурированное логирование.
- **Docker & Docker Compose** — контейнеризация.

## API Endpoints

| Метод   | Путь                          | Описание                                      |
|---------|-------------------------------|-----------------------------------------------|
| `GET`   | `/health`                     | Проверка состояния сервиса                    |
| `POST`  | `/upload`                     | Загрузка одного JPEG-файла с сжатием          |
| `POST`  | `/upload-batch`               | Загрузка пачки JPEG-файлов без сжатия         |
| `POST`  | `/upload-batch-compress`      | Загрузка пачки JPEG-файлов с сжатием          |
| `GET`   | `/files/:name`                | Скачивание одного JPEG-файла                  |
| `GET`   | `/files`                      | Список всех файлов                            |
| `GET`   | `/download-batch`             | Скачивание пачки JPEG-файлов в ZIP без сжатия |
| `GET`   | `/download-batch-compress`    | Скачивание пачки JPEG-файлов в ZIP с сжатием  |
| `DELETE`| `/files/:name`                | Удаление файла                                |

## Конфигурация

1. Скопируйте содержимое `.env.example` в `.env`.
2. Настройте переменные окружения в `.env`:
   - `MINIO_ROOT_USER` — логин для MinIO.
   - `MINIO_ROOT_PASSWORD` — пароль для MinIO.
   - `IMAGE_MAX_WIDTH` — максимальная ширина для масштабирования JPEG-изображений.
   - `IMAGE_QUALITY` — качество сжатия JPEG (от 1 до 100).

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

### Загрузка одного JPEG-файла (с сжатием)

```bash
curl -X POST \
  -F "file=@images/file1.jpeg" \
  http://localhost:8080/upload
```

### Загрузка пачки JPEG-файлов (без сжатия)

```bash
curl -X POST \
  -F "files=@images/file1.jpeg" \
  -F "files=@images/file2.jpeg" \
  http://localhost:8080/upload-batch
```

### Загрузка пачки JPEG-файлов (с сжатием)

```bash
curl -X POST \
  -F "files=@images/file1.jpeg" \
  -F "files=@images/file2.jpeg" \
  http://localhost:8080/upload-batch-compress
```

### Скачивание одного JPEG-файла

```bash
curl http://localhost:8080/files/file1.jpeg --output downloaded_file1.jpeg
```

### Скачивание пачки JPEG-файлов (без сжатия)

```bash
curl "http://localhost:8080/download-batch?names=file1.jpeg,file2.jpeg" --output images.zip
```

### Скачивание пачки JPEG-файлов (с сжатием)

```bash
curl "http://localhost:8080/download-batch-compress?names=file1.jpeg,file2.jpeg" --output images_compressed.zip
```

### Список файлов

```bash
curl http://localhost:8080/files
```

### Удаление файла

```bash
curl -X DELETE http://localhost:8080/files/file1.jpeg
```

### Проверка состояния

```bash
curl http://localhost:8080/health
```

## Особенности

- **Автоматическое сжатие JPEG**: Изображения сжимаются при использовании `/upload` или `/upload-batch-compress`.
- **Изменение размера**: JPEG-изображения масштабируются до максимальной ширины, заданной в `IMAGE_MAX_WIDTH`.
- **Пакетная обработка**: Поддержка загрузки и скачивания пачек JPEG-файлов с сжатием и без, с возвратом ZIP-архивов для скачивания.
- **Временные ссылки**: Поддержка pre-signed URLs для безопасного доступа к файлам.
- **Структурированное логирование**: Все операции логируются с помощью Zap.
- **Health checks**: Встроенная проверка состояния сервиса и зависимостей.

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