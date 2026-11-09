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

[Главный каталог CI/CD проектов](https://github.com/anitouchq/CI-CD-Pipelines)

## Скриншоты

Добавьте изображения в каталог `screenshots/`:

- `screenshots/application.png` — окно приложения с открытым Intel HEX файлом;
- `screenshots/github-actions-release.png` — успешная матричная сборка;
- `screenshots/github-release.png` — опубликованный Release с файлами.

![Окно приложения](screenshots/application.png)
![Сборка GitHub Actions](screenshots/github-actions-release.png)
![Опубликованный GitHub Release](screenshots/github-release.png)

## Как создать Release-сборку

Когда изменения готовы в `main`, выполни:

```bash
git checkout main
git pull origin main
git tag v1.0.0
git push origin v1.0.0
```

Тег формата `v*` запускает CI, тесты парсера Intel HEX, а затем нативную сборку приложения для Linux, macOS и Windows. Отдельный job публикует полученные бинарники в GitHub Releases.

Для следующего релиза используй новый тег, например `v1.0.1`.

- [GitHub Actions](https://github.com/anitouchq/hex-loader/actions)
- [GitHub Releases](https://github.com/anitouchq/hex-loader/releases)
