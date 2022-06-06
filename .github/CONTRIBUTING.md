## Commit messages

We are using the [conventional commits specification](https://www.conventionalcommits.org/en/v1.0.0/) to write commit messages.

## Submitting a pull request

1. Create a new branch from the `dev` branch.
2. Make your changes, add tests, and ensure tests pass
3. Submit a pull request from your new branch into `dev`.

## Implementing and testing new commands

Suppose we are implementing a new command `taikun foobar run`.

Tthe command's implementation must be in `cmd/foobar/run/run.go` and the tests
in `cmd/foobar/run/test_spec.sh`.

We are using [ShellSpec](https://github.com/shellspec/shellspec) to write unit
tests. See the project's README for detailed documentation.

All your tests must go in a new context.
The context's name must correspond to the path of the Go file's parent
directory in `cmd/`. In our case, it is `foobar/run`.
```shellspec
Context 'foobar/run'

# Add your tests here
# ...

End
```
