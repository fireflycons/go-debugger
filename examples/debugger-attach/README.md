# example

1. Run the example

    ```
    go run examples/debugger-attach/main.go
    Debugger attached: false. Try attaching a debugger. Process PID: 323031
    Press CTRL-C to exit.
    ```

1. Attach the debugger and continue the process

    ```
    dlv attach 323031
    Warning: no debug info found, some functionality will be missing such as stack traces and variable evaluation.
    Type 'help' for list of commands.
    (dlv) c
    ```

1. Observe state change in running example

    ```
    Debugger attached: true
    ```

1. Detach the debugger (Press CTRL-C)

    ```
    received SIGINT, stopping process (will not forward signal)
    > runtime.futex() /.../go/src/runtime/sys_linux_amd64.s:558 (PC: 0x4790e3)
    553:         MOVQ    ts+16(FP), R10
    554:         MOVQ    addr2+24(FP), R8
    555:         MOVL    val3+32(FP), R9
    556:         MOVL    $SYS_futex, AX
    557:         SYSCALL
    => 558:         MOVL    AX, ret+40(FP)
    559:         RET
    560:
    561: // int32 clone(int32 flags, void *stk, M *mp, G *gp, void (*fn)(void));
    562: TEXT runtime·clone(SB),NOSPLIT|NOFRAME,$0
    563:         MOVL    flags+0(FP), DI
    (dlv) q
    Would you like to kill the process? [Y/n] n
    ```

1.  Observe state change again in running example

    ```
    Debugger attached: false
    ```
