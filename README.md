[![Build Status](https://travis-ci.org/neckhair/gcron.svg?branch=master)](https://travis-ci.org/neckhair/gcron)

# GCron

GCron is inteded to be a replacement for cron to be used in a Docker image. It does not need root privileges, which makes it work for example on Openshift.

Work is heavily in progress. This is how it should work when it's "finished".

The configuration file will be in [YAML](http://yaml.org/):

```yml
logfile: /dev/stdout
command: echo "Hello World"
schedule: '* * * * * *'
```

And this is how the call will look like:

    gcron -c gcron.yml

Or without a config file:

    gcron --command="echo hello world" --schedule="* * * * * *"

The program will always run in the foreground as this is how it has to behave in a container.

## Docker

The tool is built to run inside a Docker container. This is how you use it:

    docker run neckhair/gcron --command="echo 'Hello World'" --schedule="*/5 * * * * *"

Or you use your own Dockerfile and copy the config file in:

```dockerfile
FROM neckhair/gcron:latest
COPY examples/gcron.yml ./
CMD ["gcron", "--config", "gcron.yml"]
```
