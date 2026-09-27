# telepeek

`telepeek` is a lightweight command-line tool for inspecting information exposed through the Telegram Bot API. It can retrieve bot details, webhook configuration, chat information, and chat administrators using a Telegram bot token.

Use this tool only with bots and chats you own or are explicitly authorized to inspect.

## Features

- Retrieve bot metadata with `getMe`
- Inspect webhook configuration with `getWebhookInfo`
- Retrieve chat information with `getChat`
- List chat administrators with `getChatAdministrators`
- Run all available checks with a single command
- Pretty-print Telegram API responses as JSON
- Accept bot tokens with or without the `bot` prefix

## Requirements

- Go 1.20 or later
- A valid Telegram bot token
- Network access to `https://api.telegram.org/`

## Installation

### Install via go

```bash
go install github.com/xenophonsec/telepeek@latest
```

Ensure, that the go bin path is included in the `PATH` variable on your machine, for this to work properly.

### Build from source
Clone the repository and build the binary:

```bash
git clone https://github.com/your-username/telepeek.git
cd telepeek
go build -o telepeek .
```

You can optionally install it into your Go binary directory:

```bash
go install .
```

Again, make sure your Go bin directory is included in your `PATH` if you want to run `telepeek` globally.

## Command Reference

| Command | Required arguments | Description |
|---|---|---|
| `bot info` | Bot token | Get bot information |
| `bot webhook` | Bot token | Get webhook information |
| `bot chat` | Bot token, `--chat` | Get information about a chat |
| `bot chatadmins` | Bot token, `--chat` | Get chat administrators |
| `bot all` | Bot token | Get bot and webhook information |
| `bot all` | Bot token, `--chat` | Get all available information, including chat data |

## Example Output

Telegram responses are formatted as indented JSON:

```json
{
    "ok": true,
    "result": {
        "id": 123456789,
        "is_bot": true,
        "first_name": "Example Bot",
        "last_name": "Example Bot",
        "username": "example_bot"
        ...
    }
}
```

If Telegram returns an error, the API response is printed as received.


## Usage

```text
telepeek bot <command> <bot-token> [options]
```

The `bot` command also has the alias `b`.

```bash
telepeek b info <bot-token>
```

The tool automatically adds the `bot` prefix when it is missing. These forms are equivalent:

```bash
telepeek bot info bot123456:ABCDEF...
telepeek bot info 123456:ABCDEF...
```

## Commands

### Get bot information

Retrieves basic information about the bot, including its ID, display name, username, and capabilities.

```bash
telepeek bot info <bot-token>
```

Example:

```bash
telepeek bot info 123456789:ABCDEF123456
```

This calls the Telegram Bot API endpoint:

```text
/getMe
```

### Get webhook information

Retrieves the bot's current webhook configuration.

```bash
telepeek bot webhook <bot-token>
```

Example:

```bash
telepeek bot webhook 123456789:ABCDEF123456
```

This calls:

```text
/getWebhookInfo
```

### Get chat information

Retrieves information about a chat accessible to the bot.

```bash
telepeek bot chat <bot-token> --chat <chat-id>
```

Example:

```bash
telepeek bot chat 123456789:ABCDEF123456 --chat -1001234567890
```

This calls:

```text
/getChat?chat_id=<chat-id>
```

The `--chat` option is required.

### Get chat administrators

Retrieves the administrators of a chat accessible to the bot.

```bash
telepeek bot chatadmins <bot-token> --chat <chat-id>
```

Example:

```bash
telepeek bot chatadmins 123456789:ABCDEF123456 --chat -1001234567890
```

This calls:

```text
/getChatAdministrators?chat_id=<chat-id>
```

The `--chat` option is required.

### Retrieve all available information

Runs the bot information and webhook checks. If a chat ID is provided, it also retrieves chat information and administrators.

```bash
telepeek bot all <bot-token>
```

With chat information:

```bash
telepeek bot all <bot-token> --chat <chat-id>
```

Example:

```bash
telepeek bot all 123456789:ABCDEF123456 --chat -1001234567890
```

The `all` command calls:

```text
/getMe
/getWebhookInfo
/getChat?chat_id=<chat-id>
/getChatAdministrators?chat_id=<chat-id>
```

## Building a Release Binary

Build for Linux:

```bash
GOOS=linux GOARCH=amd64 go build -o telepeek-linux-amd64 .
```

Build for macOS on Apple Silicon:

```bash
GOOS=darwin GOARCH=arm64 go build -o telepeek-darwin-arm64 .
```

Build for Windows:

```bash
GOOS=windows GOARCH=amd64 go build -o telepeek-windows-amd64.exe .
```

## Project Dependencies

`telepeek` uses:

- [`github.com/urfave/cli/v3`](https://github.com/urfave/cli) for command-line parsing
- [`github.com/asmcos/requests`](https://github.com/asmcos/requests) for HTTP requests
- Telegram Bot API for bot and chat data

Install or update dependencies with:

```bash
go mod tidy
```
