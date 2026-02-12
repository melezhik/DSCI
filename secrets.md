# Secrets

One may provide secrets accessible from ci pipelines

# Create secret

Ssh to a host where dsci runs and create a secret using dsci cli.

Secrets are defined per repositories in `$user/$repo/$name` format, where
$user - is user owing a repository, $repo - repository name and $name - secret name.

For example for repository `demo-php` owned by user `demo`, we can define 
secret named `password`:

```
$ ./dsci_runner --admin --action create-secret
path: (demo/demo-php/password) 
value: (12345)
```


# Use secret within pipeline

Once secret is set it accessible from pipeline as environment variable, for example
for python task (`task.py`):

```python
#!/bin/python
import os;
print("password: ",os.environ['password']);
```
