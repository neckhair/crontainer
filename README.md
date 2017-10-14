[![Build Status](https://travis-ci.org/neckhair/crontainer.svg?branch=master)](https://travis-ci.org/neckhair/crontainer)

# Crontainer

Crontainer is inteded to be a replacement for cron to be used in a Docker image. It does not need root privileges, which makes it work for example on Openshift.

Work is heavily in progress. This is how it should work when it's "finished".

The configuration file will be in [YAML](http://yaml.org/):

```yml
logfile: /dev/stdout
tasks:
- schedule: '*/5 * * * * *'
  type: command
  command:
    command: 'echo "First task"'
- schedule: '*/10 * * * * *'
  type: command
  command:
    command: 'echo "Second task"'
```

And this is how the call will look like:

    crontainer -c crontainer.yml

## Docker

The tool is built to run inside a Docker container. This is how you use it:

    docker run -v $(pwd)/examples/crontainer.yml:/etc/crontainer.yml neckhair/crontainer

## License

This project is licensed under the terms of the MIT license.
