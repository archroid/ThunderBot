# ElProfessorBot



<div align="center">
  <br />
  <p>
    <a href="https://discord.js.org"><img src="./docs/assets/BotLogo.png" width="150" alt="discord.js" /></a>
  </p>
  <br />
  <p>
<img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/archroid/ElProfessorBot">
    <img alt="Discord" src="https://img.shields.io/discord/907862584550633492?label=ElProfessorBot">
    <img alt="GitHub" src="https://img.shields.io/github/license/archroid/ElProfessorBot">
    <a href="#"><img src="https://dcbadge.vercel.app/api/shield/782162374890487810?style=flat&compact=true" /></a>
    <a href="https://discord.gg/golang"><img src="https://img.shields.io/badge/Discord%20Gophers-%23discordgo-blue.svg" /></a>
    <img alt="PyPI - Status" src="https://img.shields.io/pypi/status/go">
  </p>
</div>

## 🔥 About

ElProfessorBot is a powerful [Discord](https://discord.com/) bot that allows you to easily manage your discord server and play music from youtube written with [DiscordGo](https://github.com/bwmarrin/discordgo) and [Mongodb](https://github.com/mongodb/mongo-go-driver).

### ⭐️ Features
- Play musics from youtube API
- Polls and predictions
- Delete server messages 
- Welcome and auto role assign system.
- Rules managment

## 🔨 Installation

**Mongodb and ffmpeg are required**  

1. Ensure you have [Go](https://go.dev/dl/) installed.
2. Clone the repo:
 ```sh
git clone https://github.com/archroid/ElProfessorBot.git
```
3. move to the project directory:

```sh
cd ElProfessorBot
```
4. export `DISCORD_BOT_TOKEN` & `YOUTUBE_API_KEY`. 

```sh
export  DISCORD_BOT_TOKEN={DISCORD_BOT_TOKEN}
```

* See [this](https://www.writebots.com/discord-bot-token/) to make a bot and get your token.

* If you don't have the youtube API key, click. [here](https://blog.hubspot.com/website/how-to-get-youtube-api-key) 
 
5. run mongoDb service

```sh
sudo systemctl start mongodb
```
* See [this](https://docs.mongodb.com/manual/installation/) link to install mongoDb if you don't have it installed.
6. Run the project
```sh
go run main.go
```

### 📑 Libraries

- [DiscordGo](https://github.com/bwmarrin/discordgo)  Go package that provides low level bindings to the Discord chat client API (`go get github.com/bwmarrin/discordgo`)
- [disgolink](github.com/DisgoOrg/disgolink) disgolink is a Lavalink Client (`go get github.com/DisgoOrg/disgolink`)
- [logrus](https://github.com/sirupsen/logrus) Structured logger for Go (`go get github.com/sirupsen/logrus`)
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver) The MongoDB supported driver for Go. (`go get go.mongodb.org/mongo-driver/mongo`)
- [Ken](https://github.com/zekrotja/ken) A prototype, object-oriented and highly modular Discord application commands handler for Discordgo. (`go get github.com/zekrotja/ken`)
- [di](https://github.com/sarulabs/di) Dependency injection framework for go programs (golang). (`go get github.com/sarulabs/di/v2`)


### © Copyright Notice
- Some parts of code are originaly from project [github.com/zekroTJA/shinpuru](https://github.com/zekroTJA/shinpuru) licensed under the [MIT License](https://github.com/zekroTJA/shinpuru/blob/master/LICENCE).

