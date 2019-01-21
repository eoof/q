# modifications to original q

* `-ql` and `-qo` are automatically added to `flag` similar to `glog`
* `-ql` can be used to set debug `level`:
  1. `all` which prints all q.Q() output.
  2. `pkg` which prints q.Q() in packages named `pkg`
  3. `func` which prints q.Q() in functions named `func`
* `-qo` can be used to set output to `stderr` or `stdout` or a file other than `q`.

## examples 

```
go run demo/main.go -ql all -qo stdout
```

This will print all q.Q() output to stdout.

Default is to not output anything. `all` makes every q.Q() output to appear.
Not printing anything by default is intentional.

```
go run demo/main.go -ql demopkg -qo demopkg.log
```

This will print q.Q() output in package named `demopkg`. The output will go to a file 
named $TMPDIR/demopkg.log

```
go run demo/main.go -ql demo2 -qo stderr
```

The output of q.Q() in package or function containing `demo2` will go to `stderr`.

```
go run demo/main.go -ql main.main 
```

The output of q.Q() in `main` function in `main` package will go to a file $TMPDIR/q.

