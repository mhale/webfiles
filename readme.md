# webfiles 

A simple wrapper around [Go's built in HTTP file server](http://golang.org/pkg/net/http/#FileServer) which logs requests and allows the directory to serve and port to listen on to be supplied on the command line.

## Usage

When run without any parameters, the default action is to serve the current directory on port 8080. For example, running the server with no parameters and visiting http://localhost:8080/ would result in the following output:

```
$ webfiles 
Serving directory /Users/mhale/go/src/github.com/mhale/webfiles over HTTP on port 8080...
2014/09/01 21:13:00 GET /
2014/09/01 21:13:01 GET /favicon.ico
```

Stop the server with Ctrl+C. You can optionally apply two parameters: the directory to serve and the port, in that order. For example:

```
$ webfiles ~/files 8888
Serving directory /Users/mhale/files over HTTP on port 8888...
```

Note that you will need to run the executable with superuser or administrator privileges if you want to serve files on port 80.

## Licensing

This code is in the public domain.
