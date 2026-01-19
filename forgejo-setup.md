# install forgejo

```bash
brew install forgejo
```

# create repo

go to forgejo ui and create some repo

# set web hook for repo

```
repo -> settings -> web hooks -> http://127.0.0.1:4000/forgejo_hook
```

# set webhook allowed host

`nano ~/homebrew/var/forgejo/custom/conf/app.ini`


```ini
[webhook]
ALLOWED_HOST_LIST = 127.0.0.1
```

restart forgejo

```bash
brew services restart forgejo
```

# build dsci orchestrator

```bash
git clone https://github.com/melezhik/dsci-runner.git
cd dsci-runner
docker build . -t dsci
```

# run dsci orchestrator

On the same host where forgejo instance runs:

```bash
docker network create dsci || :
docker run \
--network dsci \
--rm --name dsci -it \
-p 4000:4000 \
--env FORGEJO_HOST=$FORGEJO_HOST \
--env FORGEJO_API_TOKEN=$FORGEJO_API_TOKEN \
--env DSCI_FEEDBACK_URL=http://127.0.0.1:4000 \
-v /var/run/docker.sock:/var/run/docker.sock \
dsci
```

On the command above one needs to change FORGEJO_HOST to actual value related to the
running forgejo instance. 

FORGEJO_API_TOKEN needs to be generated before hand:

```
user settings -> applications -> manage access tokens -> generate token
```

and inserted into the docker run command parameters. Permissions for access token needs to be set at least with "repository" permissions

**Attention!** Make it sure DSCI_FEEDBACK_URL is available publicly to see dci reports from forgejo UI

# set up dsci cicd pipeline

go to created repo and create pipeline code under `.dsci` directory, see instructions on the main doc page.

---

Happy hacking with forgejo and dsci! 😄

