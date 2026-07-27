# Dsci dispatcher container

```bash
podman exec -it dsci-dispatch bash
ps aux
```

# Jobs queue

This command shows how many jobs are queue:

```bash
ls -l ~/.dsci/.sparky/projects/dsci/.triggers/
```

# Job agent

This commands shows container executing dsci jobs

```bash
podman exec -it dsci-dispatch bash -c "podman ps"
```
