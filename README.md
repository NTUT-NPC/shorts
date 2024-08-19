<p align="center">
  <img src="docs/shorts.png" align="center" width="100" height="100">
</p>

<h1 align="center">短褲</h1>

<p align="center">一個極簡的短網址伺服器</p>

## 用

1. 設

```toml
# config/redirects.toml

[temporary]
"discord" = "https://discord.gg/9yYtgA4HXz"

[permanent]
"tat/android" = "https://play.google.com/store/apps/details?id=club.ntut.npc.tat"
"tat/ios" = "https://apps.apple.com/tw/app/id1513875597"
```

2. 開

```sh
go run main.go
```

3. 試

```text
$ http :8080/discord
HTTP/1.1 302 Found
Content-Length: 52
Content-Type: text/html; charset=utf-8
Date: Mon, 19 Aug 2024 15:04:49 GMT
Location: https://discord.gg/9yYtgA4HXz

<a href="https://discord.gg/9yYtgA4HXz">Found</a>..
```

```text
$ http :8080/tat/android
HTTP/1.1 301 Moved Permanently
Content-Length: 98
Content-Type: text/html; charset=utf-8
Date: Mon, 19 Aug 2024 15:04:03 GMT
Location: https://play.google.com/store/apps/details?id=club.ntut.npc.tat

<a href="https://play.google.com/store/apps/details?id=club.ntut.npc.tat">Moved Permanently</a>.
```

## 架

1. 做

```sh
docker build -t shorts .
```

2. 起

```sh
docker run -p 80:8080 -v ./config:/config shorts
```

或看 [compose.yaml](docs/compose.yaml)。
