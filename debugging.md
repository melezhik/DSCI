# Dsci dispatcher container

Docker

```bash
docker exec -it dsci-dispatch bash
```

Podman

```bash
podman exec -it dsci-dispatch bash
```

# Jobs queue

This command shows how many jobs are queue:

```bash
ls -l ~/.dsci/.sparky/projects/dsci/.triggers/
```

# Job agent

Docker

``bash
docker exec -it dsci-agent bash
```

Podman

```bash
podman exec -it dsci-dispatch
docker exec -it dsci-agent bash
```


# TCP ports

Following ports should be open:

* `8080`
* `8181`




