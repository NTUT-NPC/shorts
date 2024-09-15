<p align="center">
  <img src="docs/shorts.svg" alt="Shorts Logo" align="center" width="128" height="128">
</p>

<h1 align="center">短褲 Shorts</h1>

<p align="center">一個使用 Go 語言製作的輕量短網址服務</p>

## 功能

### 即時重新載入設定

編輯 `config/redirects.toml` 文件，Shorts 會自動重新載入設定。

```toml
# config/redirects.toml

[temporary]
"discord" = "https://discord.gg/9yYtgA4HXz"

[permanent]
"google" = "https://www.google.com"
```

### 臨時和永久重新導向

可以將實驗性的連結添加為臨時重新導向（307），然後將它們改成永久重新導向（301）以加快重新導向速度。

基於上述設定：

```sh
curl -v localhost:8080/discord
```

```text
< HTTP/1.1 302 Found
< Content-Type: text/html; charset=utf-8
< Location: https://discord.gg/9yYtgA4HXz
< Date: Sun, 08 Sep 2024 14:28:11 GMT
< Content-Length: 52
< 
<a href="https://discord.gg/9yYtgA4HXz">Found</a>.
```

```sh
curl -v localhost:8080/google
```

```text
< HTTP/1.1 301 Moved Permanently
< Content-Type: text/html; charset=utf-8
< Location: https://www.google.com
< Date: Mon, 09 Sep 2024 08:38:22 GMT
< Content-Length: 57
< 
<a href="https://www.google.com">Moved Permanently</a>.
```

### 查看統計數據

Shorts 會在 `config/stats.json` 中記錄每個重新導向的訪問者數量和最後訪問時間。

```json
{
  "discord": {
    "visitors": 1,
    "last_visited": "2024-09-08T22:28:11.894270007+08:00"
  },
  "google": {
    "visitors": 1,
    "last_visited": "2024-09-09T16:38:22.113075596+08:00"
  }
}
```

## 部署

我們建議使用 Docker 部署 Shorts。

### Docker

```sh
docker run -d -p 8080:8080 \
  -v $PWD/config:/config \
  ghcr.io/ntut-npc/shorts
```

### Docker Compose

參見位於 [docs/compose.yaml](docs/compose.yaml) 的配置範例。

## 開發

本地開發 Shorts 的步驟：

1. 再製這個倉庫：

    ```sh
    git clone https://github.com/ntut-npc/shorts.git
    cd shorts
    ```

2. 安裝依賴：

    ```sh
    go mod download
    ```

3. 執行應用程式：

    ```sh
    go run .
    ```

伺服器會在 `http://localhost:8080` 啟動。

### 開發期間的即時重新載入

為了獲得更好的開發體驗，我們推薦使用 [gow](https://github.com/mitranim/gow)，它會在原始碼更改時自動重啟程式。

記得根據[即時重新載入設定](#即時重新載入設定)部分建立並編輯你的 `config/redirects.toml` 文件來設定重新導向。
