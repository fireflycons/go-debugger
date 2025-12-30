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
$ gdb main
(gdb) run
Starting program: /.../main
[New LWP 217628]
[New LWP 217629]
[New LWP 217630]
Debugger is attached
[LWP 217630 exited]
[LWP 217628 exited]
[LWP 217547 exited]
[LWP 217629 exited]
[New process 217547]
[Inferior 1 (process 217547) exited normally]
(gdb) q
