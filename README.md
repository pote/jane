# Jane - A Philote command line client

Jane is a command line client for the [Philote](https://github.com/pote/philote) websockets server, it's useful for debugging purposes or as a reference implementation.

## Basic usage

```bash
$ jane -help
Usage of jane:
  -c string
    	channel you want to connect to (default "test")
  -info
    	If passed, queries Philote for running info and exits
  -p string
    	Philote server (default "localhost:6380")
  -s string
    	Philote server's JWT secret used for authentication.
  -secure
    	Will use wss and https instead of ws and http
  -t string
    	Auth token, if you want to use a specific one
```

### Homebrew

```bash
$ brew install pote/philote/jane
```

### Precompiled binaries

You can find precompiled binaries for your platform of choice in the [latest release](https://github.com/pote/jane/releases) page.

### Install from source

You'll need [gpm](https://github.com/pote/gpm) installed to manage dependencies.

```bash
$ git clone git@github.com:pote/jane.git && cd jane
$ make install
```

## License

Released under MIT License, check LICENSE file for details.
