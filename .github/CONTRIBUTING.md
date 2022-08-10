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
[root command](../cmd/root/root.go) (`taikun`) using the following syntax.
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

## A testing example

### The context

Suppose we would like to test the `taikun billing-rule delete` command.

Here is its usage menu:
```sh
sh$ taikun billing rule delete --help
Delete one or more billing rules

Usage:
  taikun billing rule delete <billing-rule-id>... [flags]

Aliases:
  delete, rm
```

The tests for this command are written in the
`cmd/billing/rule/remove/test_spec.sh` file. All Shellspec test cases must be
written in a file named `test_spec.sh` in order to be run by `shellspec`. We
place the file in the same directory as the command's implementation
(`cmd/billing/rule/remove/remove.go`).

You can see the entire testing specification of the billing rule deletion
command [here](../cmd/billing/rule/remove/test_spec.sh).

We first declare a context for the test cases in
`cmd/billing/rule/remove/test_spec.sh` with the `Context` keyword and the
closing keyword `End`.
```sh
Context 'billing/rule/remove'

# Our test cases will go here...

End
```

The name of this context is the path of the parent directory of `test_spec.sh`
starting from the `./cmd` directory (not included). Thus here it is
`billing/rule/remove`.

### The setup function

In order to test the deletion of billing rules, we will first need to create
some. Billing rules require billing credentials, we will thus need to create
a billing credential beforehand.

By convention, commands that must be run before the actual test cases are
called in a `setup` function (functions are declared using Bash syntax).

We can tell Shellspec to run this `setup` function before all the test
cases using the `BeforeAll` directive.
```sh
Context 'billing/rule/remove'

  setup() {
    name=$(_rnd_name)
    cname=$(_rnd_name)
    cid=$(taikun billing credential add -p $PROMETHEUS_PASSWORD -u $PROMETHEUS_URL -l $PROMETHEUS_USERNAME $cname -I)
  }
  BeforeAll 'setup'

  # Our test cases will go here...

End
```

Here we define some useful variables:
- `name` will be the name of the billing rules we create
- `cname` is the name of the billing credential of the billing rules.

We use the helper function `_rnd_name` (defined in
[spec_helper.sh](../.spec/spec_helper.sh)) to generate random names.
```sh
    name=$(_rnd_name)
    cname=$(_rnd_name)
```


We then create the billing credential we need for the billing rules and store
its ID in the `cid` variable. The `-I` flag (or `--id-only`) makes it so only
the ID of the newly created resource is outputted by `taikun billing credential
add`. The required Prometheus credentials are defined in the GitHub secrets of
the test action.
```
cid=$(taikun billing credential add -p $PROMETHEUS_PASSWORD -u $PROMETHEUS_URL -l $PROMETHEUS_USERNAME $cname -I)
```

### The cleanup function
Whenever we create resources as dependencies for test cases, we need to delete
them afterwards, this is done in the `cleanup` function.

In our case, we need to delete the billing credential with the ID stored in
the `cid` variable.

The `AfterAll` directive tells Shellspec to run it after all the test cases
defined in this file.
```sh
Context 'billing/rule/remove'
  
  setup() {
    # ...
  }
  BeforeAll 'setup'

  cleanup() {
    taikun billing credential delete $cid -q 2>/dev/null || true
  }
  AfterAll 'cleanup'

  # Our test cases will go here...

End
```

We use `|| true` so that Shellspec won't consider this a failed test if the
`taikun billing credential delete` command fails.  As we are only testing
billing rule deletion in this specification, billing credential deletion should
not impact the result of this test.  Billing credential deletion is tested in
its own file.

We also redirect all eventual error output using the redirection `2>/dev/null`
and all standard output using the `-q` (or `--quiet`) flag.

### BeforeEach and AfterEach directives
We will be testing three cases:
1. Deleting a nonexistent rule should fail
2. Deleting an existing rule should succeed
3. Deleting an existing rule and a non-existing rule with the same command should fail

As a prerequisite, we need to create a billing rule before each test and
to delete it afterwards in case the deletion failed.

We first define the `add_rule` function which creates a billing rule and stores
its ID in the `id` variable. The `BeforeEach` directive tells Shellspec to run
`add_rule` before each test case.
```sh
  add_rule() {
    id=$(taikun billing rule add $name -b $cid -l foo=bar -m foo --price 1 --price-rate 5 -t count -I)
  }
  BeforeEach 'add_rule'
```

We then define the `rm_rule` function which delete the billing rule with the ID
stored in the `id` variable. The `AfterEach` directive tells Shellspec to run
`rm_rule` after each test case.
```sh
  rm_rule() {
    taikun billing rule delete $id -q 2>/dev/null || true
  }
  AfterEach 'rm_rule'
```

Once again, we use the `--quiet` flag and an stderr redirection to avoid
polluting our test suite's output. `|| true` ensures that this won't be
considered a failed test if `taikun billing rule delete $id` fails (in fact, if
the deletion is properly implemented, we expect it to fail).

### Defining the test cases
We can now define our three test cases.

Test cases are declared using the `Example` keyword followed by a string as a
description and `End` as a closing keyword.
```sh
Example 'delete nonexistent billing rule'
  # ...
End
```

Let's take test case #3 as an example: deleting an existing and a nonexistent
billing rule in the same command.

The command we are testing is `taikun billing rule delete 0 $id` as no billing rule
will ever have the ID `0` and `id` is the ID of the billing rule created in the
`add_rule` function.

We use the syntax `When call` followed by a command or function.
```sh
  Example 'delete nonexistent billing rule'
    When call taikun billing rule delete 0 $id
    # checks go here...
  End
```

We can then check the status code, stdout and stderr.
- The status code should be 1 as `taikun` should fail to delete the billing rule with ID `0`.
- The standard output (stdout) should include a message of successful deletion for the billing rule with ID `$id`.
- The error output (stderr) should include a 404 error message.

```sh
  Example 'delete existing and nonexistent billing rules'
    When call taikun billing rule delete 0 $id
    The status should equal 1
    The output should include 'was deleted successfully'
    The output should include "$id"
    The stderr should include 404
    The stderr should include 'Error: Failed to delete one or more resources'
  End
```
Checks begin with the keyword `The` followed by a _Subject_ (`status`, `output`,
`stderr` or a variable).
Following the keyword `should` is a _Matcher_ such as `equal` or `include` and
one or more arguments such as `1` or `"$id"`.

_Modifiers_ can also be added for more precision. For example, if we expect the
length of stderr to be exactly 80 characters, we can use the following syntax.
```
The length of stderr should equal 80
```

The [Shellspec
documentation](https://github.com/shellspec/shellspec#dsl-syntax)
goes further into detail about the DSL syntax.

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
