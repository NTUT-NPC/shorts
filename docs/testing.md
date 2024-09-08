
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
