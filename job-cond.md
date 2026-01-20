# Job conditions

Job conditions allow to skip or pass jobs on certain criteria.

Passing rules operate on job environment parameters is in `.<param> $condition` form, for example:


Skip job_one is commit message is equal to 'skip'

```yaml
jobs:
  -
    id: job_one
    path: .
    skip: .<message> eq "skip"
``` 

One can combine multiple job env parameters using standard AND/OR/||/&& operators from Boolean logic:

```yaml
jobs:
  -
    id: job_one
    path: .
    skip: .<message> eq "skip" or .<ref> eq "refs/heads/dev"
```

# Job environment parameters

## ref

Branch/Tag name. Example:

    refs/heads/main

## message

Commit message

## repo_full_name

Full repository name. Example:

    root/go-build-example

## sha

Commit SHA. Example:

    689b7f1e30537514759ce40086fd2218fb95793d

## scm

Git repository URL. Example:

    http://127.0.0.1:3000/root/go-build-example.git


## Regular expressions

One may use Raku regular expression to match against them.


Following example only passes job with branch name dev or master:

```yaml
jobs:
  -
    id: job_one
    path: .
    only: .<ref> ~~ "refs/heads/" [ master || main ]
```

## Parenthesis

One can use Parenthesis for grouping of logical expressions.

```yaml
jobs:
  -
    id: job_one
    path: .
    only: > 
      ( .<ref> eq "refs/heads/main" and .<message> eq "release_prod" ) or 
      ( .<ref> eq "refs/heads/dev" and .<message> eq "release_dev" )
```

## Operators

Following the list of operators one may use:

## eq

equivalent to string

## nq

not equivalent to string

## <, >, >=, <=, != 

less then, greater then, less or equal, greater or equal, not equal then an Int

## ~~

regular expression matching

## !~~

regular expression matching, negation form


## Regular expression modifiers

One can modify regular expression modifiers, to change matching logic.

Following example skip job if commit message is `skip_ci` or `Skip_Ci` or `SKIP_CI` - case insensitive 
modifier

Follow [~Raku regexs](https://www.google.com/url?sa=t&source=web&rct=j&opi=89978449&url=https://docs.raku.org/language/regexes)
for more details

```

```yaml
jobs:
  -
    id: job_one
    path: .
    skip: m:i/skip_ci/ 
```

## Global jobs conditions

To globally skip/pass all jobs use `global:` key in the root of pipeline configuration:


```yaml
global:
  skip: .<message> ~~ m:/skip/
jobs:
  -
    id: job_one
    path: .
```
