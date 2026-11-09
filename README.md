# 🔎 Hex Loader — Go + Fyne → GitHub Releases

Учебное GUI-приложение на Go и Fyne для открытия и проверки файлов Intel HEX. Приложение проверяет структуру записей, количество байтов, контрольные суммы и обязательную EOF-запись.

## Возможности

- Открытие файлов через системный диалог.
- Проверка префикса и длины каждой записи.
- Проверка Intel HEX checksum.
- Обработка Data, EOF, Extended Segment Address, Extended Linear Address и Start Address records.
- Отображение количества записей и байтов данных.
- Понятное сообщение об ошибке с номером строки.

Приложение проверяет файл, но не прошивает микроконтроллер и не заменяет программатор.

## CI checks

- `gofmt` — форматирование;
- `go vet ./...` — статический анализ;
- `go test ./... -v` — тесты парсера;
- `go build` — сборка GUI.

Fyne использует CGO и системные графические библиотеки. Поэтому Release собирается на нативных runner'ах Linux, macOS и Windows.

## Release artifacts

| Платформа | Артефакт |
|---|---|
| Linux x64 | `hex-loader-linux-x64` |
| macOS | `hex-loader-macos` |
| Windows x64 | `hex-loader-windows-x64.exe` |

Чтобы создать Release, отправьте тег:

```bash
git tag v1.0.0
git push origin v1.0.0
```

## Локальный запуск

Требуется Go 1.27+, GCC и графические библиотеки Fyne. Для Ubuntu/Debian:

```bash
sudo apt-get install gcc libgl1-mesa-dev xorg-dev libwayland-dev libxkbcommon-dev
go mod tidy
go test ./... -v
go run .
```

[Главный каталог CI/CD проектов](https://github.com/anitouchq/CI-CD-Pipelines)
