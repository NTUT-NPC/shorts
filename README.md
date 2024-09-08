<p align="center">
  <img src="docs/shorts.png" alt="Shorts Logo" align="center" width="100" height="100">
</p>

<h1 align="center">短褲 Shorts</h1>

<p align="center">A lightweight URL shortener built using Go</p>

## Getting Started

### Prerequisites

1. **Docker**: Ensure Docker is installed on your system. If you need installation instructions, refer to the [Docker installation guide](https://docs.docker.com/engine/install/).

### Building (Optional)

If you wish to build the Docker image yourself, follow these steps:

1. Clone the repository:
    ```sh
    git clone https://github.com/NTUT-NPC/shorts.git
    ```

2. Navigate to the repository directory:
    ```sh
    cd shorts
    ```

3. Build the Docker image:
    ```sh
    docker build -t shorts .
    ```

### Install with docker-compose

If you prefer using the pre-built Docker image, you can pull it from the GitHub Container Registry:

1. Pull the Docker image:
    ```sh 
    docker pull ghcr.io/ntut-npc/shorts:latest
    ```

2. Create a directory for Docker Compose and configuration files:
    ```sh
    mkdir -p /srv/shorts-docker/config
    cd /srv/shorts-docker
    ```

3. Create and edit the `compose.yaml` file:
    ```sh
    touch compose.yaml
    nano compose.yaml
    ```

4. Add the following content to `compose.yaml`, then save and exit:
    ```yaml
    services:
      shorts:
        image: shorts
        container_name: shorts-docker
        ports:
          - "<custom-port>:8080"
        volumes:
          - ./config:/config
    ```

5. Start the service:
    ```sh
    docker compose up -d
    ```

## Configuration

To set up URL redirects, modify the `config/redirects.toml` file:

1. Open the configuration file:
    ```sh
    nano config/redirects.toml
    ```

2. Add your redirects using the following format:
    ```toml
    # config/redirects.toml

    [temporary]
    "discord" = "https://discord.gg/9yYtgA4HXz"
    "<custom-string>" = "<url>"

    [permanent]
    "google" = "https://www.google.com"
    ```

   - **Temporary Redirects**: Short-lived and can be updated or removed as needed.
   - **Permanent Redirects**: Intended for long-term use and stable links.

   The URL format will be `http://<custom-domain>:<custom-port>/<custom-string>`. For example: `https://to.ntut.club/discord`.

   **Note**: You do not need to restart the Docker container after editing `redirects.toml`. Changes will be applied automatically.

   For improved security and performance, consider placing this service behind a reverse proxy.

## Viewing Statistics

To access statistics for your redirects, check the `config/stats.json` file:

1. Open the statistics file:
    ```sh
    less config/stats.json
    ```

2. Example content of `stats.json`:
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

   The file includes visitor counts and the last visit timestamp for each redirect.

## Test the Connection

To verify that your URL shortener is working correctly, you can use the `httpie` command-line or desktop tool. Follow the instructions below to test your setup.

### Installing `httpie`

First, install the `httpie` package. You can find the installation tutorial [here](https://httpie.io/).

For example, on a RHEL-based system using `dnf`:
```sh
sudo dnf install httpie
```

### Testing with `https`

Use the `https` command to test the URL redirection. Replace `to.ntut.club` with your custom domain and `discord` with a test string you've configured.

```sh
https to.ntut.club/discord
```

The expected output should look like this:
```text
HTTP/1.1 302 Found
Alt-Svc: h3=":443"; ma=2592000
Content-Length: 52
Content-Type: text/html; charset=utf-8
Date: Mon, 09 Sep 2024 01:23:45 GMT
Location: https://discord.gg/9yYtgA4HXz
Server: Caddy

<a href="https://discord.gg/9yYtgA4HXz">Found</a>.
```

### Testing with `http`

Alternatively, use the `http` command to test. Replace `8080` with your custom port and `/tat/android` with a test string you've configured.

```sh
http :8080/tat/android
```

The expected output should look like this:
```text
HTTP/1.1 301 Moved Permanently
Content-Length: 98
Content-Type: text/html; charset=utf-8
Date: Mon, 19 Aug 2024 15:04:03 GMT
Location: https://play.google.com/store/apps/details?id=club.ntut.npc.tat

<a href="https://play.google.com/store/apps/details?id=club.ntut.npc.tat">Moved Permanently</a>.
```

By running these commands, you can confirm that your URL shortener is functioning as expected.