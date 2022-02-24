# EdgeCast CLI

This is the official command line interface (CLI) for EdgeCast.

## Getting Started

### Requirements
- Go 1.17

### Installation
To install the CLI, use the makefile:
```
$ make install
```

### Using the CLI
Commands follow the following format:
```
$ ec <command> <subcommand> --arg1=value1 --arg2=value2`
```

Below is an example:
```
$ ec origin getallorigins --accountnumber=ABCDE --mediatypeid=3
```

In addition to using flags to provide inputs, you may also provide JSON input:
```
$ ec origin getallorigins --input='{AccountNumber:ABCD, MediaTypeID=3}'
```

To view the top-level documentation for the CLI, use the `help` command:
```
$ ec help
```

To get help for a command or subcommand, use the `--help` flag:
```
$ ec origin --help
```

To view your version of the cli:
```
$ ec version
```

## Resources

* [CDN Reference Documentation](https://docs.edgecast.com/cdn/index.html)
    * This is a useful resource for learning about the EdgeCast CDN. It is a good starting point before using this CLI.
* [API Documentation](https://docs.edgecast.com/cdn/index.html#REST-API.htm%3FTocPath%3D_____8)
    * For developers that want to interact directly with the EdgeCast CDN API, refer to this documentation. It contains all of the available operations as well as their inputs and outputs.
* [EdgeCast GO SDK](https://github.com/EdgeCast/ec-sdk-go)
    * For developers that want to use the EdgeCast platform using Go. The EC CLI uses the SDK under the hood.
* [Examples](https://github.com/EdgeCast/ec-cli/tree/main/example)
    * Examples to get started can be found here.
* [Submit an Issue](https://github.com/EdgeCast/ec-cli/issues)
    * Found a bug? Want to request a feature? Please do so here.