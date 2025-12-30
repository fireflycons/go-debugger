# example

```
$ go build example/main.go
$ ./main
Debugger is not attached
$ dlv exec main
Type 'help' for list of commands.
(dlv) c
Debugger is attached
Process 56758 has exited with status 0
(dlv) q
