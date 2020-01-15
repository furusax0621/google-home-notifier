# Google Home Notifier

## Usage

```bash
notifier <config_file>
```

## Configuration

```toml
[general]
host     = 192.168.100.11 # hostname or ip address for Google Home
port     = 8009           # (optional) connection port. default is 8009.
interval = 3              # (optional) intervals after operation. default is 3 seconds

# configuration to speak text
[notify]
text = "Hello World"
lang = "en"          # (optional) language. default is "en"

# configuration to play sounds
[play]
url = "https://example.com/"
```

## Credit

This application is based on [github.com/kunihiko-t/google-home-notifier-go](https://github.com/kunihiko-t/google-home-notifier-go) by [kunihiko-t](https://github.com/kunihiko-t/).