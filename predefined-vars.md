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

git commit sha

## DSCI_SCM

git repo clone url

## DSCI_REF

git reference ( tag or branch name )

## DSCI_MESSAGE

git commit message

# DSCI_FULL_REPO

git full repository name

## DSCI_FORGEJO_HOST

Forgejo host with schema and http port

## DSCI_JOB_ID

dsci job internal id

