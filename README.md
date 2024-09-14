<p align="center">
  <img src="docs/shorts.svg" alt="Shorts Logo" align="center" width="128" height="128">
</p>

<h1 align="center">短褲 Shorts</h1>

<p align="center">A lightweight URL shortener built using Go</p>

[中文 README](README_ZH.md)

## Features

### Hot-Reloading Configuration

Edit the `config/redirects.toml` file and Shorts will automatically reload the configuration.

```toml
# config/redirects.toml

[temporary]
"discord" = "https://discord.gg/9yYtgA4HXz"

[permanent]
"google" = "https://www.google.com"
```

### Temporary and Permanent Redirects

Add experimental redirects as temporary redirects (307) and change them to permanent redirects (301) for faster redirection.

With the above configuration:

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

### Viewing Statistics

Shorts records the number of visitors and the last visited time for each redirect in `config/stats.json`.

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

## Deployment

We recommend deploying Shorts using Docker.

### Docker

```sh
docker run -d -p 8080:8080 \
  -v $PWD/config:/config \
  ghcr.io/ntut-npc/shorts
```

### Docker Compose

See [docs/compose.yaml](docs/compose.yaml) for an example Docker Compose configuration.

## Development

To set up Shorts for local development:

1. Clone the repository:

    ```sh
    git clone https://github.com/ntut-npc/shorts.git
    cd shorts
    ```

2. Install dependencies:

    ```sh
    go mod download
    ```

3. Run the application:

    ```sh
    go run .
    ```

The server will start on `http://localhost:8080`.

### Hot Reloading During Development

For a better development experience, we recommend using [gow](https://github.com/mitranim/gow), which automatically restarts the application when source files change.

Remember to create and configure your `config/redirects.toml` file as described in the Configuration section to set up your redirects.
