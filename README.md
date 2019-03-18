# ippon-ippan-twitter-bot

## Docker Commands

### Build

```console
$ docker build --force-rm --no-cache -t 178inaba/ippon-ippan-twitter-bot .
```

### Run

```console
$ docker run -d
  -e "ACCESS_TOKEN=<access-token>"
  -e "ACCESS_TOKEN_SECRET=<access-token-secret>"
  -e "CONSUMER_KEY=<consumer-key>"
  -e "CONSUMER_SECRET=<consumer-secret>"
  --name ippon-ippan-twitter-bot 178inaba/ippon-ippan-twitter-bot
```

### Push

```console
$ docker push 178inaba/ippon-ippan-twitter-bot
```
