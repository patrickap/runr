# runr

`Runr` is a small command runner written in Go. Can be configured with `yaml`.

## Installation

[Download](./build/runr) the executable of `runr` and add it to the `$PATH` variable or manually move it to `/usr/bin/runr`.

## Usage

Create a `runr.yaml` configuration file. By default `Runr` will search for it in the current working directory. Alternatively the `RUNR_CONFIG_DIR` environment variable can be set. For `Runr` to function properbly at least one command must be defined. The first entry in the command list must be an exectubale. Then simply run `runr <command-name>`.

## Overview

```yaml
commands:
  <command-name>:
    command: []
    options: {}
    hooks:
      pre: []
      post: []
      success: []
      failure: []
```

### Command

```yaml
command: [<executable>, <command>, <command>]

command:
  - <executable>
  - <command>
  - <command>
```

### Options

```yaml
options:
  <option>: <value>
  <option>:
    - <value>
    - <value>
  -<option>: <value>
  --<option>: <value>
```

### Hooks

```yaml
hooks:
  pre: []
  post: []
  success: []
  failure: []
```

## Advanced Configuration

...
