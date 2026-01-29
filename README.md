# Waifu & Husbando Catcher — Go Port

This repository is a Go port of the WAIFU-HUSBANDO-CATCHER Telegram bot.

Quick start:
1. Copy `configs/.env.example` to `.env` and set TELEGRAM_TOKEN and MONGO_URI.
2. Build: `make build`
3. Run: `make run` or `./bin/bot`

Files & structure:
- cmd/bot: entrypoint
- internal/config: env parsing
- internal/storage: mongo client
- internal/bot: bot init & main loop
- internal/handlers: Telegram command handlers
- internal/scheduler: spawn scheduler
- internal/models: DB models
- internal/utils: helper functions (small-caps, catbox upload)

Notes:
- The port covers core features and scaffolds all modules.
- Many advanced modules contain TODOs to port detailed logic (gift/trade, inline queries, leaderboard aggregates).
- After initial run, please test using a development Telegram bot token.

License: MIT (same as original)