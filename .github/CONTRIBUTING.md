## Commit messages

We are using the [conventional commits specification](https://www.conventionalcommits.org/en/v1.0.0/) to write commit messages.

## Submitting a pull request

1. Create a new branch from the `dev` branch.
2. Make your changes, add tests, and ensure tests pass
3. Submit a pull request from your new branch into `dev`.

## Implementing and testing new commands

Suppose we are implementing a new command `taikun foobar`.

The command's implementation (`NewCmdFoobar()`), must be in
`cmd/foobar/foobar.go` and the tests in `cmd/foobar/test_spec.sh`.

To add the `foobar` command to the list of subcommands, it must be added in the
[root command](./cmd/root/root.go) (`taikun`) using the following syntax.
```go
cmd.AddCommand(foobar.NewCmdFoobar())
```

We are using [ShellSpec](https://github.com/shellspec/shellspec) to write unit
tests. See the project's README for detailed documentation.

All your tests must go in a new context.
The context's name must correspond to the path of the Go file's parent
directory in `cmd/`. In our case, it is `foobar`.
```shellspec
Context 'foobar'

# Add your tests here
# ...

End
```

## Moving and/or renaming commands

Suppose we would like to rename `taikun test foobar` to `taikun foobar`.
`foobar` is a child command of `test`, we want it to be a child command of the
root command, i.e. `taikun`.

The code for the `foobar` command would be in `cmd/test/foobar/foobar.go` and
the unit tests in `cmd/test/foobar/test_spec.sh`.

We would first need to move the `cmd/test/foobar/` directory to `cmd/foobar/`.

Since `foobar` was initially a child command of `test`, we need to remove the
following line from `cmd/test/test.go` and add it to `cmd/root/root.go`.
```go
cmd.AddCommand(foobar.NewCmdFoobar())
```

Finally, since the path of `foobar` has changed, the tests' context must be
renamed from `test/foobar` to `foobar`.

In other words, we would need to make the following change in
`cmd/foobar/test_spec.sh`.
```diff
-Context 'test/foobar'
+Context 'foobar'
```
