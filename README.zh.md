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

### API 功能

#### 嘗試重新導向 API

`/api/try` 端點允許您嘗試重新導向到某個 slug，如果 slug 不存在，則會重新導向到指定的備用 URL。此 API 需要兩個參數：`slug`（嘗試重新導向的目標）和 `fallback`（當 slug 不存在時重新導向的完整 URL）。

```sh
curl -v "localhost:8080/api/try?slug=discord&fallback=not-found"
```

```text
< HTTP/1.1 302 Found
< Content-Type: text/html; charset=utf-8
< Location: https://discord.gg/9yYtgA4HXz
< Date: Sun, 10 Sep 2024 15:30:22 GMT
< Content-Length: 52
< 
<a href="https://discord.gg/9yYtgA4HXz">Found</a>.
```

如果 slug 不存在，將會重新導向到同一域名上的備用路徑：

```sh
curl -v "localhost:8080/api/try?slug=nonexistent&fallback=not-found"
```

```text
< HTTP/1.1 302 Found
< Content-Type: text/html; charset=utf-8
< Location: http://localhost:8080/not-found
< Date: Sun, 10 Sep 2024 15:31:45 GMT
< Content-Length: 56
< 
<a href="http://localhost:8080/not-found">Found</a>.
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

2. 執行開發腳本：

    ```sh
    scripts/dev.sh
    ```

伺服器會在 `http://localhost:8080` 啟動。記得根據[即時重新載入設定](#即時重新載入設定)部分建立並編輯你的 `config/redirects.toml` 文件來設定重新導向。
