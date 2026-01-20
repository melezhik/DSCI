# Jobs

To create a job, create `job.$ext` file inside job folder, where $ext is one of supported by SDK languages extension, for example:

*job_one/job.python*

```python
run_task(
  'task_one', {
      'foo' : 'foo value',
      'bar' : 'bar value'
    }
);

run_task(
  'task_two', {
      'foo' : 'foo value',
      'bar' : 'bar value'
    }
);
```

Following is a list of job file names for different languages:

```
+------------+--------------+
| Language   | File         |
+------------+--------------+
| Raku       | job.raku     |
| Perl       | job.pl       |
| Bash       | job.bash     |
| Python     | job.py       |
| Ruby       | job.rb       |
| Powershell | job.ps1      |
| Golang     | job.go       |
+------------+--------------+
```

Usually job is just a sequence of running [~tasks](/doc/task), but maybe any code:

*job_one/job.python*

```python
print("hello from job")
run_task(
  'task_one', {
      'foo' : 'foo value',
      'bar' : 'bar value'
    }
);
```

Task execution is done by calling `run_task()` function which is supported by all SDK languages, it takes the first mandatory parameter as a path inside `tasks/` directory where task files resides, and optionally tasks variables passed as dictionary

`run_task` function signatures for supported languages:

```
+------------+----------------------------------------------+
| Language   | Signature                                    |
+------------+----------------------------------------------+
| Raku       | run_task(String,HASH)                        |
| Perl       | run_task(SCALAR,HASHREF)                     |
| Bash       | run_task TASK_NANE NAME VAL NAME2 VAL2       |
| Python     | run_task(STRING,DICT)                        |
| Ruby       | run_task(STRING,HASH)                        |
| Powershell | run_task(STRING,HASH)                        |
+------------+----------------------------------------------+
```

## Tasks variables

Tasks variables are handled inside task by using `task_var()` function available by all SDK languages:

Bash:

```bash
foo=$(task_var foo)
# in Bash you may also use shorted form:
echo $bar
```

Python:

```python
foo = task_var["foo"]
```

Raku

```raku
say task_var("foo");
```

`task_var` function signatures for supported languages:

```
+------------------+------------------------------------------------+
| Language         | Signature                                      |
+------------------+------------------------------------------------+
| Raku             | task_var(STRING)                               |
| Perl             | task_var(SCALAR)                               |
| Python           | task_var(STRING)                               |
| Ruby             | task_var(STRING)                               |
| Bash (1-st way)  | $foo                                           |
| Bash (2-nd way)  | $(task_var foo)                                |
| Powershell       | task_var(STRING)                               |
+------------------+------------------------------------------------+
```

# Default job parameters

If there is file named `config.yaml` inside job directory it sets jobs default parameters:

```yaml
name: Alexey
occupation: IT
residency: Russia
```

Job parameters are available within tasks via `config()` function which is supported for all SDK languages:

```python
dict = config()
print(dict["name"])
print(dict["occupation"])
print(dict["residency"])
```


Job default parameters might be overridden by job parameters inside pipeline

# Further reading 

* [~Job Conditions](/doc/job-cond)

* [~Pipelines](/doc/pipeline)

* [~Tasks](/doc/task)
