[![Build Status](https://travis-ci.org/kcmerrill/oogway.svg?branch=master)](https://travis-ci.org/kcmerrill/oogway) [![Go Report Card](https://goreportcard.com/badge/github.com/kcmerrill/oogway)](https://goreportcard.com/report/github.com/kcmerrill/oogway)

![oogway](assets/oogway.png "oogway")

**Oogway** is simple, yet flexible, monitoring tool. At it's core, **Oogway** is a coordinated shell runner. You can determine how often checks are run, what to do when they fail, when they succeed. What you do is completely up to you. Send stats, send notifications etc ...

## Usage

```bash
$ oogway --check-interjval 10s --check-dir /path/to/checks --check-extension oog  
```

* `--check-interval` How often the checks should be reloaded? *(default 1s)*
* `--check-dir` Directory where checks are stored *(default ".")*
* `--check-extension` Extension of check yaml files *(default "oog")*

## Configuring Checks

Checks are yaml files with a given extension. By default, that extension is `.oog`. Within these yaml files will contain your checks. Below is what a full **Oogway** file looks like. 
```yaml

kcmerrill.com: # the check name, unique 
    check: # the main command that starts off the chain reaction
        summary: A short description
        Tries: 5 # how many times we should try before the check goes critical
        Every: 30s # the interval in which to check
        cmd: curl --fail https://kcmerrill.com
    ok: # when the check is ok
        cmd: |
            echo -n "kcmerrill.com.ok:60|g|#shell" | nc -4u -w0 127.0.0.1 8125
    warning: # when the check has failed, but not yet critical
        cmd: |
            echo -n "kcmerrill.com.warning:1|g|#shell" | nc -4u -w0 127.0.0.1 8125
    critical: # what takes place when the check goes critical. Logging? Notifications?
        cmd: |
            echo -n "kcmerrill.com.critical:1|g|#shell" | nc -4u -w0 127.0.0.1 8125
            echo "kcmerrill.com failed" | mail -s "Something failed!" kcmerrill@gmail.com
    fix: # Oogway can attempt to "fix" your issue. It might be another script/command to run 
        after: 3 # after how many attempts should we let pass before attempting to fix?
        cmd: |
            ssh me@kcmerill.com docker run -d -P kcmerrill/kcmerrill.com
            echo -n "kcmerrill.com.fixed:1|g|#shell" | nc -4u -w0 127.0.0.1 8125
    recover: # If the check goes critical, and then recovers(say after fix) Oogway can recover
        cmd: |
            echo "Check is ok ...go back to bed" | mail -s "kcmerrill.com is ok" kcmerrill@gmail.com
```

## Binaries || Installation

[![MacOSX](https://raw.githubusercontent.com/kcmerrill/go-dist/master/assets/apple_logo.png "Mac OSX")](http://go-dist.kcmerrill.com/kcmerrill/oogway/mac/amd64) [![Linux](https://raw.githubusercontent.com/kcmerrill/go-dist/master/assets/linux_logo.png "Linux")](http://go-dist.kcmerrill.com/kcmerrill/oogway/linux/amd64)

via go:

`$ go get -u github.com/kcmerrill/oogway`
