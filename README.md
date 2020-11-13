# juju-go-example

### Summary

This is a minimal example repository showcasing how to connect to the [Juju](https://juju.is) API and execute commands directly from Go, without relying on the `juju` client.

Running the example program will connect to a Juju controller, retrieve the status (the equivalent of running the `juju status` command) and print the list of applications available from the model.

The default Juju configuration from `.local/share/juju` is used.

### Requirements

- `Go` (tested with 1.15, should work with 1.14)
- A deployed Juju controller, a model with a few deployed applications and user credentials.

### Usage

```bash
$ go run ./main.go
```
