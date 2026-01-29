# Shivu Bot - Go Version

A Telegram bot converted from Python to Go for better performance and concurrency.

## Features

- Fast and efficient Go-based Telegram bot
- MongoDB integration for data persistence
- Scheduled tasks support
- Caching mechanism
- Rate limiting
- Modular architecture

## Prerequisites

- Go 1.21 or higher
- MongoDB instance
- Telegram Bot Token

## Installation

### Local Development

1. Clone the repository:
```bash
git clone https://github.com/yourusername/shivu-go.git
cd shivu-go
```

2. Install dependencies:
```bash
go mod download
go mod tidy
```

3. Create a `.env` file in the root directory:
```env
BOT_TOKEN=your_telegram_bot_token
MONGO_URI=your_mongodb_connection_string
DATABASE_NAME=shivu
```

4. Run the bot:
```bash
go run cmd/shivu/main.go
```

### Building

Build for your current platform:
```bash
go build -o shivu cmd/shivu/main.go
```

Build for Linux (for deployment):
```bash
GOOS=linux GOARCH=amd64 go build -o shivu cmd/shivu/main.go
```

### Docker

Build and run with Docker:
```bash
docker build -t shivu-bot .
docker run -d --env-file .env shivu-bot
```

## Project Structure

```
shivu-go/
├── cmd/
│   └── shivu/          # Main application entry point
│       └── main.go
├── internal/
│   ├── config/         # Configuration management
│   ├── bot/            # Bot initialization and setup
│   ├── handlers/       # Message and command handlers
│   ├── database/       # Database operations
│   ├── models/         # Data models
│   ├── services/       # Business logic
│   └── utils/          # Utility functions
├── pkg/                # Public packages (if any)
├── .env.example        # Example environment variables
├── .gitignore
├── Dockerfile
├── go.mod
├── go.sum
├── Procfile
└── README.md
```

## Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| BOT_TOKEN | Telegram Bot API Token | Yes |
| MONGO_URI | MongoDB connection string | Yes |
| DATABASE_NAME | MongoDB database name | Yes |
| LOG_LEVEL | Log level (debug, info, warn, error) | No (default: info) |

## Development

### Running Tests

```bash
go test ./...
```

### Running with live reload (using air)

```bash
# Install air
go install github.com/cosmtrek/air@latest

# Run with air
air
```

## Deployment

### Heroku

1. Create a new Heroku app
2. Set environment variables in Heroku dashboard
3. Deploy using Git:
```bash
git push heroku main
```

### Railway/Render

1. Connect your GitHub repository
2. Set environment variables
3. Deploy automatically on push

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

See LICENSE file for details.

## Original Python Version

This is a Go port of the original Python version. The Python version can be found at [original-repo-link].

## Credits

- Original Python version by IzumiCypherX
- Go conversion by [Your Name]
