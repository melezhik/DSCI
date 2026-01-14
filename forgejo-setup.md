# install forgejo

```bash
brew install forgejo
```

# create repo

go to forgejo ui and create some repo

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

# run dsci runner

On the same host where forgejo instance runs:

```
docker run \
--port 3000:3000,4000:4000 \
--env FORGEJO_HOST=http://127.0.0.1:3000 \
--env FORGEJO_API_TOKEN=token \
dsci
```

On the command above one needs to change FORGEJO_HOST to actual value related to the
running forgejo instance. FORGEJO_API_TOKEN needs to be generated before hand:

```
user settings -> applications -> manage accrss tokens -> generate token
```

and inserted into the docker run command input parameters. Permissions for access token needs to be set at least with "reposiory" permissions

Once dsci container has run, copy the 
content of it's public ssh key from output and save it, you will need it on the next step

Attention! Please make it sure that dsci container is accessible from within your gorgejo instance by http://127.0.0.1:3000 address and dsci web ui is available public by 4000 port.

# set dsci ssh public key

```
user settings -> ssh keys -> add key -> content of public ssh key taken from previous step
```

# set up dsci cicd pipeline

go to created repo and create pipeline code under `.dsci` directory, see instructions on the main doc page.

---

Happy hacking with fogejo and dsci! 😄


