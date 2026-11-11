# hex-loader

[![CI](https://github.com/anitouchq/hex-loader/actions/workflows/ci.yml/badge.svg)](https://github.com/anitouchq/hex-loader/actions)

Учебное GUI-приложение на Go и Fyne для открытия и проверки файлов Intel HEX. Приложение проверяет структуру записей, количество байтов, контрольные суммы и обязательную EOF-запись.

## Возможности

- Открытие файлов через системный диалог.
- Проверка префикса и длины каждой записи.
- Проверка контрольной суммы Intel HEX.
- Обработка Data, EOF, Extended Segment Address, Extended Linear Address и Start Address records.
- Отображение количества записей и байтов данных.
- Сообщение об ошибке с номером строки.

Приложение проверяет файл, но не прошивает микроконтроллер и не заменяет программатор.

## CI checks

- `gofmt` — форматирование;
- `go vet ./...` — статический анализ;
- `go test ./... -v` — тесты парсера;
- `go build` — проверочная сборка GUI.

Fyne использует CGO и системные графические библиотеки, поэтому Release собирается на нативных runner'ах Linux, macOS и Windows.

## Release artifacts

| Платформа | Артефакт |
|---|---|
| Linux x64 | `hex-loader-linux-x64` |
| macOS | `hex-loader-macos` |
| Windows x64 | `hex-loader-windows-x64.exe` |

Чтобы запустить публикацию, отправьте тег версии:

```bash
git tag v1.0.0
git push origin v1.0.0
```

Бинарники будут опубликованы в [GitHub Releases](https://github.com/anitouchq/hex-loader/releases).

## Локальный запуск

Требуются Go 1.27+, GCC и графические библиотеки Fyne. Для Ubuntu/Debian:

```bash
sudo apt-get install gcc libgl1-mesa-dev xorg-dev libwayland-dev libxkbcommon-dev
go test ./... -v
go run .
```

## Структура проекта

| Файл или каталог | Назначение |
|---|---|
| `.github/workflows/ci.yml` | Проверки, нативная сборка под три ОС и публикация Release |
| `main.go` | Запуск GUI и взаимодействие с пользователем |
| `hex.go` | Разбор и проверка записей Intel HEX |
| `hex_test.go` | Тесты корректных и ошибочных HEX-файлов |
| `go.mod` | Имя Go-модуля и прямые зависимости |
| `go.sum` | Контрольные суммы модулей |
| `.gitignore` | Локальные файлы, исключаемые из Git |

## Результаты

- [GitHub Actions](https://github.com/anitouchq/hex-loader/actions)
- [Опубликованные релизы](https://github.com/anitouchq/hex-loader/releases)


## 📸 Скриншоты

### 1. Работающее приложение

![Окно Hex Loader с открытым Intel HEX файлом](screenshots/application.png)

### 2. GitHub Actions — успешный CI/CD

![GitHub Actions — успешная сборка](screenshots/github-actions-release.png)

### 3. GitHub Release v1.0.4 — 3 бинарника

![GitHub Release с бинарными файлами](screenshots/github-release.png)

[Главный каталог CI/CD проектов](https://github.com/anitouchq/CI-CD-Pipelines)
