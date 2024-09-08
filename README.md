<p align="center">
  <img src="docs/shorts.png" alt="Shorts Logo" align="center" width="100" height="100">
</p>

<h1 align="center">短褲 Shorts</h1>

<p align="center">A lightweight URL shortener built using Go</p>

## Getting Started

[中文 README](README_ZH.md)

### Prerequisites

1. **Docker**: Ensure Docker is installed on your system. If you need installation instructions, refer to the [Docker installation guide](https://docs.docker.com/engine/install/).

### Building (Optional)

Please see [docs/building.md](docs/building.md)

### Install with docker-compose

1. Create a directory for Docker Compose and configuration files:

    ```sh
    mkdir -p /srv/shorts-docker/config
    cd /srv/shorts-docker
    ```

1. Edit `compose.yaml`:

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

1. Start the service:

    ```sh
    docker compose up -d
    ```

## Configuration

To set up URL redirects, modify the `config/redirects.toml` file:

1. Add your redirects using the following format:

    ```toml
    # config/redirects.toml

    [temporary]
    "discord" = "https://discord.gg/9yYtgA4HXz"
    "<custom-string>" = "<url>"

    [permanent]
    "google" = "https://www.google.com"
    ```

   The URL format will be `http://<custom-domain>:<custom-port>/<custom-string>`.

## Viewing Statistics

Example content of `config/stats.json`:

```json
{
  "discord": {
    "visitors": 1,
    "last_visited": "2024-08-20T17:49:36.57603941+08:00"
  }
}
```

## Test the Connection with cURL

```sh
curl -v localhost:8080/discord
```

```text
> GET /discord HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/8.9.1
> Accept: */*
> 
< HTTP/1.1 302 Found
< Content-Type: text/html; charset=utf-8
< Location: https://discord.gg/9yYtgA4HXz
< Date: Sun, 08 Sep 2024 14:28:11 GMT
< Content-Length: 52
< 
<a href="https://discord.gg/9yYtgA4HXz">Found</a>.
```
