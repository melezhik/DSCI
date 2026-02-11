# Predefined variables

Dsci provides some predefined variables under "DSCI_" prefix
for pipeline developers convenience.

Variables are accessible via `config()` function:


`task.py`:


```python
/bin/python

c = config()
ssh = c[DSCI_COMMIT] # commit sha 
```


# Variables List


## DSCI_COMMIT

git commit

## DSCI_SCM

git repo clone url

## DSCI_JOBID

dsci job internal id

