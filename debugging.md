# Dsci job dispatcher logs

```bash
podman logs -f dsci-dispatch
```

# Jobs queue

This command shows how many jobs are queue:

```bash
podman exec -it dsci-dispatch bash -c "ls -l ~/.sparky/projects/dsci/.triggers/"
```

# Job agent

This commands shows container executing dsci jobs

```bash
podman exec -it dsci-dispatch bash -c "podman ps"
```
