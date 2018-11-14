# go-telegram-bot-example
Simple example of Telegram Bot written in Go.

# building
You will need Go 1.11+
1. Clone it **outside** the GOPATH
2. `go build -o bot`
3. Create `config.yml` and set database name:
```yaml
database: users.db
```
4. Try it: `./bot -config config.yml`
