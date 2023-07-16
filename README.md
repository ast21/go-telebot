# Learning Assistant telegram bot

Telegram bot for learning

## Migrations

Для работы с миграциями, нужно установить [GOOSE](https://github.com/pressly/goose#install), 
и создать файл настройки `.gooserc`
```bash
cp .gooserc.example .gooserc
```

Отредактировать DSN и выполнить эту команду
```bash
source .gooserc
```

Провести миграции:
```bash
goose up
```
