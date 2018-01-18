[![Build Status](https://travis-ci.org/neckhair/crontainer.svg?branch=master)](https://travis-ci.org/neckhair/crontainer)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/ab158c206de34374962341d8e023ac07)](https://www.codacy.com/app/neckhair/crontainer?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=neckhair/crontainer&amp;utm_campaign=Badge_Grade)

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

## Development

Dependencies are managed with [`dep`](https://github.com/golang/dep). Just run `dep ensure` or `make vendor` to install all dependencies.

There is a make file for the most common tasks:

```sh
# install dependencies
$ make vendor

# run tests and build binary
$ make all

# create docker image
$ make docker
```

## License

This project is licensed under the terms of the MIT license.
