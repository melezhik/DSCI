# Bash plugins

DSCI has a reach mechanism allowing to write extension for pipelines, called plugins. In this short post I am going to overview the process of creation and using simple plugin written on Bash ( other supported by SDK languages are also possible )


# Create plugin meta file

The should be named sparrow.json, as underlying DSCI uses Sparrow framework to run tasks and plugins:

```json
{
  "name" : "foo",
  "description" : "Simple foo plugin",
  "version" : "0.0.1",
  "url" : "https://forgejo.sparrowhub.io/root/sparrow-plugins",
  "category" : "demo, utils"
}
```

Let's skim through some of parameters:

* *name* - name of a plugin, this defines how plugin will be called in tasks

* *description* - short description of plugin

* *version* - plugins are versioned inside internal DSCI repository, so when whenever a plugin
has changed an author assign it a new bumped version and upload it to repository, so users will
see this new version

* *url* - and git repo url that presumably keeps a plugin source code

* *category* - an optional parameter which might be used when search plugin withing DSCI repository

# Create plugin task file

Normally it's the same DSCI job file one might be already familiar with by reading [~tasks](/task.md) documentation.

So the same logic applies here.

Let's create a simple Bash task:

task.bash:

```bash
#!/bin/bash
echo "hello from foo plugin"
```

# Plugin input parameters

DSCI provides convenient DSL for tasks and plugins developers. This SDK allow effectively handle input parameter.

Let's say we have an input parameter named `message`, first of all we can defined a default value for the one:

config.yaml

```yaml
message: "hello"
```

Parameter is handled within task like this:

task.bash

```
#!/bin/bash

message=$(config message)
echo "you say: $message"
```

If we now run plugin without specifying parameter, default value will be picked up.

To override parameter from job file we can do so:

jobs.yaml

```
jobs:
  -
    id: FooJob
    plugin: foo
    params:
      message: "Hi!"
```

# Passing results back to pipeline

Plugin uses standard SDK to pass results back to pipeline:

task.bash

```
#!/bin/bash

message=$(config message)
echo "you say: $message"

update_state out "bye bye"
```

Update_state() functions save output parameters into internal dsci cache, so that other
tasks/plugins can access it. Consider there is another python task in the pipeline:

task.py

```python
#!/bin/python

cfg = config()

print(dict["_dsci_"]["FooJob"]["out"])
```

Thus DSCI provides very efficient way to pass results between subsequent tasks and plugins.

# Inline plugins into Bash scenarios

Another recent and cool feature of DSCI engine is to allow inline plugin into Bash scenarios.

Let's change original task.bash a bit:

```
#!/bin/bash

message=$(config message)
echo "you say: $message"

echo "inline: var_out='Bye!'"
```

So instead of returning results through dsci internal cache we use inline syntax, declaring Bash variable. To use the plugin inside any Bash scenarios we just need to call using `s6` internal DSCI tasks cli:

```bash
#!/bin/bash

eval(s6 --plg-run foo@message='Hello from Bash')
echo $var_out # should print Bye!
```

This mechanism effectively allow us plugins that depend on other plugins in runtime. Very cool.

---

That is it


PS This post does not cover the topic how plugins get uploaded into internal DSCI storage,
probably will reveal this next time )
