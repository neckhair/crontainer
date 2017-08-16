[![Build Status](https://travis-ci.org/neckhair/crontainer.svg?branch=master)](https://travis-ci.org/neckhair/crontainer)

# Crontainer

Crontainer is inteded to be a replacement for cron to be used in a Docker image. It does not need root privileges, which makes it work for example on Openshift.

Work is heavily in progress. This is how it should work when it's "finished".

The configuration file will be in [YAML](http://yaml.org/):

```yml
logfile: /dev/stdout
command: echo "Hello World"
schedule: '* * * * * *'
```

And this is how the call will look like:

    crontainer -c crontainer.yml

Or without a config file:

    crontainer --command="echo hello world" --schedule="* * * * * *"

The program will always run in the foreground as this is how it has to behave in a container.

## Docker

The tool is built to run inside a Docker container. This is how you use it:

    docker run neckhair/crontainer --command="echo 'Hello World'" --schedule="*/5 * * * * *"

Or you use your own Dockerfile and copy the config file in:

```dockerfile
FROM neckhair/crontainer:latest
COPY examples/crontainer.yml ./
CMD ["crontainer", "--config", "crontainer.yml"]
```

## License

This project is licensed under the terms of the MIT license.
