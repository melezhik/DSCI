Prerequisites for self hosted installation

* Forgejo
* docker
* golang ( to compile runner, pre-maid binaries comming soon )

# install forgejo

```bash
brew install forgejo
```

# create repo

go to forgejo ui and create some repo

# set web hook for repo

```
repo -> settings -> web hooks -> http://127.0.0.1:8080/forgejo_hook
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

# build dsci runner

```bash
git clone https://github.com/melezhik/dsci-runner.git
cd dsci-runner
go mod tidy
go build -o dsci_runner main.go
```

# configure dsci runner

`nano ~/.dsci.toml`:

```toml
ForgejoHost     = "http://127.0.0.1:3000"
ForgejoApiToken = "foobarbaz"
DsciFeedbackUrl  = "http://127.0.0.1:8080"
DsciAgentSkipBootstrap  = false
DsciAgentImage = "alpine:latest"
```

On the command above one needs to change `ForgejoHost` to actual value related to the running forgejo instance. 

`ForgejoApiToken` needs to be generated before hand:

```
user settings -> applications -> manage access tokens -> generate token
```

and inserted into the docker run command parameters. Permissions for access token needs to be set at least with "repository" permissions

**Attention!** Make it sure DsciFeedbackUrl is available publicly to see dsci reports from forgejo UI

# run dsci runner

```bash
./dsci_runner
```

# set up dsci cicd pipeline

go to created repo and create pipeline code under `.dsci` directory, see instructions on the main doc page.

---

Happy hacking with forgejo and dsci! 😄

# Further reading 

* [~Dsci runner configuration](/doc/configuration)

* [~Maintenance](/doc/maintenance)
