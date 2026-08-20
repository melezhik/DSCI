# Job Artifacts


DSCI dramatically simplify working with artifacts
process inside pipelines.

To create artifact file just create it in `~/artifacts/`
directory in some task file, for example:


```bash
#!/bin/bash

echo "DSCI is cool" > ~/artifacts/message.txt
```

The the next job will see it in the same directory:


```python
#!/usr/bin/python3

from pathlib import Path

file_path = Path.home() / "artifacts/message.txt"

# Read the file created by previous job(s)
try:
    with open(file_path, "r", encoding="utf-8") as file:
        content = file.read()
        print(content)
except FileNotFoundError:
    print(f"Error: The file at {file_path} was not found.")
```

---

That's it. Artifacts de-facto is any file located in `~/artifacts` dir,

If some job decides to remove a file from `~/artifacts/` it won't be available for the next
jobs.

So artifacts works as a pipeline data buffer


