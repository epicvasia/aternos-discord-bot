# Aternos Discord Bot (Modified Fork)

## !! NOT AN AFK BOT !!
**This is NOT an AFK bot and never will be.**

## !! RISK OF BAN !!
**Use of this bot may lead to bans from Aternos. Use at your own risk.**

> **⚠ Use this software at your own discretion. Read the [TOS](https://github.com/epicvasia/aternos-discord-bot/blob/main/TOS.md) for more information.**

If you're just looking for easy hosting, consider [exaroton](https://exaroton.com/) instead — a premium solution by the Aternos team with an official Discord bot.

---

## About
A Discord bot that can start your Aternos Minecraft server with a simple command.

Built on top of [aternos-api](https://github.com/sleeyax/aternos-api) and forked from the original [sleeyax/aternos-discord-bot](https://github.com/sleeyax/aternos-discord-bot) project.

**Note:** This fork only supports **starting** the server, not stopping it. This version also allows members to start the server themselves.

## Commands

| Command      | Description                                                  |
|--------------|--------------------------------------------------------------|
| `/start`     | Starts your Aternos server asynchronously.                   |
| `/status`    | Shows the current server status (online/offline/starting).   |
| `/info`      | Detailed server info.                                        |
| `/players`   | Lists current players.                                       |
| `/ping`      | Simple connectivity test.                                    |
| `/help`      | Returns helpful resources.                                   |
| `/configure` | Save your Aternos credentials (if not using MongoDB).        |

<details>
<summary>How to get your Aternos credentials</summary>

1. Go to your [Aternos server page](https://aternos.org/server/).
2. Open Developer Tools (Ctrl+Shift+I).
3. Go to the Application/Storage tab and find the cookies.
4. Copy values for `ATERNOS_SERVER` and `ATERNOS_SESSION`.
</details>

---

## Hosting the Bot

### Requirements

- Discord bot token
- Hosting service (e.g., [Railway](https://railway.app/))
- MongoDB URI (optional, only needed for multi-server support)

### Environment Variables

| Variable           | Required | Description                                 |
|--------------------|----------|---------------------------------------------|
| `DISCORD_TOKEN`    | Yes      | Your bot token from [Discord Developer Portal](https://discord.com/developers/applications) |
| `ATERNOS_SESSION`  | Yes*     | Aternos session cookie (if no MongoDB)      |
| `ATERNOS_SERVER`   | Yes*     | Aternos server cookie (if no MongoDB)       |
| `MONGO_DB_URI`     | No       | MongoDB connection string for multi-server  |

---

## Self-Hosting Options

### Railway (Recommended)
Railway is the easiest free option for hosting this bot. You can deploy via GitHub or by creating your files directly in their web interface.

### Docker
Run the bot in a Docker container:
```bash
docker run -d --name aternos-discord-bot \
-e DISCORD_TOKEN="" \
-e ATERNOS_SESSION="" \
-e ATERNOS_SERVER="" \
epicvasia/aternos-discord-bot
```

### Go Source Build
Compile and run the bot from source (requires Go 1.18+):
```bash
git clone https://github.com/epicvasia/aternos-discord-bot.git
cd aternos-discord-bot
go build -o aternos-discord-bot ./cmd/main.go
./aternos-discord-bot
```

### Kubernetes
If you're running Kubernetes:
```bash
kubectl create ns aternos-discord-bot
kubectl create secret generic aternos-secrets \
--from-literal=DISCORD_TOKEN=<> \
--from-literal=ATERNOS_SESSION=<> \
--from-literal=ATERNOS_SERVER=<> \
--from-literal=MONGO_DB_URI=<>
kubectl apply -n aternos-discord-bot -f ./kubernetes.yaml
```

---

## Advanced Integration (Go Developers)
You can import and use the bot as a package:

```go
import (
	"github.com/sleeyax/aternos-discord-bot"
	"github.com/sleeyax/aternos-discord-bot/database"
)

bot := discord.Bot{
	DiscordToken: "<your token>",
	Database: &database.MemoryDatabase{},
}

bot.Start()
defer bot.Stop()
```

---

## License
Licensed under `MIT License`.

[TL;DR](https://tldrlegal.com/license/mit-license):
> A short, permissive software license.  
> Basically, you can do whatever you want as long as you include the original copyright and license notice in any copy of the software/source.  
> There are many variations of this license in use.

---

Maintained by [TheEpicVasia](https://github.com/epicvasia)  
Based on work by [sleeyax](https://github.com/sleeyax/aternos-discord-bot)
