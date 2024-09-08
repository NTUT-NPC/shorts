<p align="center">
  <img src="docs/shorts.png" alt="Shorts Logo" align="center" width="100" height="100">
</p>

<h1 align="center">短褲 Shorts</h1>

<p align="center">一個極簡的短網址伺服器</p>

## Getting Started

### 安裝

1. 確保 [Docker](https://docs.docker.com/engine/install/) 已安裝

1. 建立目錄 
    ```sh
    mkdir -p /srv/shorts-docker/config
    cd /srv/shorts-docker
    ```
1. 建立並編輯 `compose.yaml`
    ```yaml
    services:
      shorts:
        image: ghcr.io/ntut-npc/shorts:latest
        container_name: shorts-docker
        ports:
          - "<custom-port>:8080"
        volumes:
          - ./config:/config
    ```

1. 啟動服務
    ```sh
    docker compose up -d
    ```

## 設定短網址

照著以下範本編輯 `config/redirects.toml`

```toml
# config/redirects.toml

[temporary]
"discord" = "https://discord.gg/9yYtgA4HXz"
"<custom-string>" = "<url>"

[permanent]
"google" = "https://www.google.com"
```

   短網址為：`http://<custom-domain>:<custom-port>/<custom-string>`

## 查看報告
報告在 `config/stats.json`

```json
{
  "discord": {
    "visitors": 1,
    "last_visited": "2024-08-20T17:49:36.57603941+08:00"
  },
  "google": {
    "visitors": 1,
    "last_visited": "2024-08-20T17:49:42.709932014+08:00"
  }
}
```
    
## 測試
直接使用瀏覽器測試
