# Install

Download and run dsci main binary

```
curl http://dsci.sparrowhub.io/distors/dsci_1.0.1 -o dsci
./dsci
```

Go to web interface and run bootstrap button to install forgejo and dsci job orchestrator

http://127.0.0.1:8080


# Setup some infrastructure

* Go to forgejo interface

* Login using admin/admin credentials (should be changed afterwords)

* Go to `infra` repository, edit and run example pipeline to install some basic services (ssh/mysql/etc)

http://127.0.0.1:3000

## Note

 127.0.0.1 should be replaced by your server public IP address

# Further setup

Normally everything is done via dsci [pipelines](~/doc/pipeline) which in nutshell are just Bash or Python scripts,
so developers just need to write them using guidelines or use some existing dsci plugins



