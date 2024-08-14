# runr

Runr is a lightweight command runner developed in Go, designed for simplicity and efficiency. It can be easily configured using `yaml` files, making it a flexible tool for automating and managing command execution.

## Installation

[Download](./build/runr) the executable of Runr and add it to the `$PATH` variable or manually move it to `/usr/bin/runr`.

## Usage

Create a `runr.yaml` configuration file, which Runr will automatically search for in the current working directory. Alternatively, you can set the `RUNR_CONFIG_DIR` environment variable to specify a different directory. To ensure Runr operates correctly, at least one command must be defined in the configuration file. The first entry in the command list must be an executable. Once configured, you can execute commands by running `runr <command-name>`.

## Configuration Reference

```yml
# Anchors are constructs that can be reused throughout the config.
# It's not related to Runr.
default_options: &default_options
  request: "GET"
  header: "Content-Type: text/html"

commands:
  # Specify the command name which can be run using the Runr CLI.
  # This command config is equivalent to: curl https://google.com --request "GET" --header "Content-Type: text/html" --verbose
  request:
    # Specify the command to run. Alternative syntax:
    # command:
    #   - curl
    #   - https://google.com
    command: ["curl", "https://google.com"]
    options:
      # Options maps directly to command line options.
      # Can be either of type boolean, string, integer, or list for adding multiple.
      # Anchor aliases can easily be used to reuse common options.
      <<: *default_options
      # Every option can be specified with prefix if needed Defaults to "--" (e.g. --verbose).
      # Alternative syntax:
      # --verbose: true
      # -verbose: true
      verbose: true
      # Multiple options of the same type can be added using a list.
      header:
        - "Cache-Control: no-cache"
        - "sessionId=3f2a9b10a8c749e78c6f1a234567d890"
    hooks:
      # Runs before. Alternative syntax:
      # pre:
      #   - echo
      #   - "Running pre hook..."
      pre: ["echo", "Running pre hook..."]
      # Runs after. Alternative syntax:
      # post:
      #   - /bin/sh
      #   - -c
      #   - |
      #     echo 'Running post hook...'
      post: ["/bin/sh", "-c", "echo 'Running post hook...'"]
      # Runs only on success. Hooks can also be used to run other Runr commands.
      success: ["/bin/sh", "-c", "runr <command-name>"]
      # Runs only on failure.
      failure: []
```

## Advanced Configuration

It’s also possible to execute complex shell commands that require interpretation by a specific shell like `/bin/sh -c`.

```yml
commands:
  id:
    command: ["/bin/sh", "-c", "echo $(id)"]
```

It's also possible to change the position of the applied command options. By default, they get added at the end of the command automatically. To change it, run the command in a new shell process and access them using the special variable `$@`. Use `--` to signal the end of options for the `/bin/sh` command. Any arguments after `--` will be treated as positional parameters or arguments for the command string executed by `/bin/sh -c`. Try it in the terminal:

```bash
/bin/sh -c 'echo ${@}' -- --option-1 --option-2 --option-3
```

In the example below, during execution `${@}` gets replaced with the actual options.

```yml
commands:
  hello-world:
    command: ["/bin/sh", "-c", "echo ${@}; echo 'world!'", "--"]
    options:
      hello: true
```

To run commands repeatedly or group commands to a workflow, hooks are suitable:

```yml
commands:
  hello-world:
    command: ["/bin/sh", "-c", "echo 'hello world!'"]
    hooks:
      post:
        - /bin/sh
        - -c
        - |
          runr command-1
          runr command-2
          runr command-3
```

**Important:** Please note that invoking `runr` within the command property itself is not feasible. This is due to the fact that the primary command obtains an exclusive lock, thereby preventing concurrent execution of other `runr` commands.

## Contribution

Runr is developed in Go. To build the binary, execute `just go-build`, which will generate the binary at `./build/runr`. To publish a new version, use `just release <patch|minor|major>`. This command will automatically increment the version according to semantic versioning.
