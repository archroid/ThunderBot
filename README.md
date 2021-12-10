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

## About

ElProfessorBot is a powerful [Discord](https://discord.com/) bot that allows you to easily manage your discord server and play music from youtube written with [DiscordGo](https://github.com/bwmarrin/discordgo) and [Mongodb](https://github.com/mongodb/mongo-go-driver).

- Play musics from youtube API
- Polls and predictions
- Delete server messages 
- Welcome and auto role assign system.
- Rules managment

## Installation

**Mongodb and ffmpeg are required**  

```sh
git clone https://github.com/archroid/ElProfessorBot.git
cd ElProfessorBot
export export DISCORD_BOT_TOKEN={DISCORD_BOT_TOKEN} && export YOUTUBE_API_KEY={YOUTUBE_API_KEY}
sudo systemctl start mongodb
go run main.go
```

### Libraries

- [DiscordGo](https://github.com/bwmarrin/discordgo)  Go package that provides low level bindings to the Discord chat client API (`go get github.com/bwmarrin/discordgo`)
- [ca](https://github.com/jonas747/dca) An audio file format that uses opus audio packets and json metadata. (`go get github.com/jonas747/dca/cmd/dca`)
- [logrus](https://github.com/sirupsen/logrus) Structured logger for Go (`go get github.com/sirupsen/logrus`)
- [YouTube Data API v3](https://pkg.go.dev/google.golang.org/api/youtube/v3) Package youtube provides access to the YouTube Data API v3.(`go get google.golang.org/api/youtube/v3`)
- [Youtube in Golang](https://github.com/kkdai/youtube) Youtube video download package (`go get github.com/kkdai/youtube/v2`)
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver) The MongoDB supported driver for Go. (`go.mongodb.org/mongo-driver/mongo`)

## List of Discord APIs

See [this chart](https://abal.moe/Discord/Libraries.html) for a feature 
comparison and list of other Discord API libraries.

