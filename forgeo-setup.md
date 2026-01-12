# install forgejo

```bash
brew install forgejo
```

# create repo

# set web hook for repo

```
repo -> settings -> web hooks -> http://127.0.0.1:3000
```

# set webhook allowed host

`nano ~/homebrew//var/forgejo/custom/conf/app.ini`


```ini
[webhook]
ALLOWED_HOST_LIST = 127.0.0.1
```

restart forgejo

```bash
brew services restart forgejo
```
