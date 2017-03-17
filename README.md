[![Build Status](https://travis-ci.org/neckhair/gcron.svg?branch=master)](https://travis-ci.org/neckhair/gcron)

# GCron

GCron is inteded to be a replacement for cron to be used in a Docker image. It does not need root privileges, which makes it work for example on Openshift.

Work is heavily in progress. This is how it should work when it's "finished".

The configuration file will be in [YAML](http://yaml.org/):

```yml
logfile: /dev/stdout
jobs:
- name: Print something to stdout every second
  command: echo "Hello World"
  pattern: '* * * * * *'

- name: Call a url every 5 minutes.
  command: 'curl https://www.github.com'
  pattern: '* */5 * * * *'
```

And this is how the call will look like:

    gcron -c gcron.yml

Or without a config file:

    gcron --command="echo hello world" --pattern="* * * * * *"

The program will always run in the foreground as this is how it has to behave in a container.
