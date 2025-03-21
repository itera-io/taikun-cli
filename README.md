# Taikun Command line interface (CLI)
Manage resources in [Taikun](https://taikun.cloud) from the command line.

## Getting started
You can take a look in the quickstart folder for some inspiration on some Taikun CLI scripting inspiration. 

### Downloading the binary
To download the CLI, head to the [latest release page](https://github.com/itera-io/taikun-cli/releases/latest).
Scroll down to the Assets section and select the binary for your architecture.

### Signing in to Taikun
The Taikun CLI reads environment variables to authenticate to Taikun.

To authenticate with your Taikun account, set the following environment variables:
```
TAIKUN_EMAIL
TAIKUN_PASSWORD
```

To authenticate with Keycloak, set the following environment variables:
```
TAIKUN_KEYCLOAK_EMAIL
TAIKUN_KEYCLOAK_PASSWORD
```

The default API host is `api.taikun.cloud`.
To override it, set the following environment variable:
```
TAIKUN_API_HOST (default value is: api.taikun.cloud)
```

Run the following command to check whether you are properly authenticated.
```sh
taikun whoami
```

### Setting up autocompletion
Autocompletion is available for the following shells.
- Bash
- Zsh
- Fish
- PowerShell

The command `taikun completion <shell>` generates an autocompletion script for
the specified shell. For instructions on how to use the generated script, see
the help command of the corresponding shell.

For example, `taikun completion bash -h` provides instructions on how to set up
autocompletion for the Bash shell.

## Command overview
To have an overview of all the commands available, see [the generated command
tree](COMMAND_TREE.md).

## Help
To get information on how to use a command, type `taikun [command] --help` or
`taikun [command] -h` for short.

## Contributing
Please check out the [contributing page](.github/CONTRIBUTING.md) for instructions on how
to contribute to this project.
