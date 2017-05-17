# Jane - A Philote command line client

Jane is a command line client for the [Philote](https://github.com/pote/philote) websockets server, it's useful for debugging purposes or as a reference implementation.

## Basic usage

```bash
$ jane --help
Usage of jane:
  -c string
    	channel you want to connect to (default "test")
  -p string
    	Philote server (default "ws://localhost:6380")
  -s string
    	Philote server's JWT secret used for authentication.
  -t string
    	Auth token, if you want to use a specific one
```

## Install

A Homebrew package is in the works, but for now you'll need to either [install from source](#install-from-source) or head over to the [releases](https://github.com/pote/jane/releases) page and download a precompiled binary.

### Install from source

You'll need [gpm](https://github.com/pote/gpm) installed to manage dependencies.

```bash
$ git clone git@github.com:pote/jane.git && cd jane
$ make install
```

## License

Released under MIT License, check LICENSE file for details.
