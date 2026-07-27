# Dsci job dispatcher logs

```bash
podman logs -f dsci-dispatch
```

# Jobs queues 

This command shows how many jobs are queue:

```bash
podman exec -it dsci-dispatch bash -c "ls -l ~/.sparky/projects/dsci/.triggers/"
```
