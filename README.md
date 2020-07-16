# Watch'n'Go

 * Watch a single file
 * Watch files recursively in a directory, with an optional pattern
 * Store configuration in INI file or use only the command line
 * Run a command on modifications through `/bin/sh -c <command>` by default…
 * …or output on stdout so you do whatever you want (`fswatch`-like)

## Usage

```
watchngo [-conf watchngo.ini] [-command <your command> -match <match> [-filter <filter>] [-debug] [-output unixshell|raw|stdout]]
```

When using `-command -match -filter` options, configuration will be ignored. This makes it possible to use `watchngo` without writing a configuration file.

### Configuration

See [watchngo.sample.ini](watchngo.sample.ini) configuration example.

## Install

### Binary

Checkout the [releases](https://github.com/LTGRP/watchngo/releases) binaries and put it somewhere in your `$PATH`.

Quick win with the latest release:

```
wget https://github.com/LTGRP/watchngo/releases/download/1.2.0/watchngo -O /usr/local/bin/watchngo
chmod 755 /usr/local/bin/watchngo
```

### Build from sources

```
mkdir -p $GOPATH/src/github.com/LTGRP
cd $GOPATH/src/github.com/LTGRP
git clone https://github.com/LTGRP/watchngo

cd watchngo

glide install
cd cmd/watchngo

make
make install
```

