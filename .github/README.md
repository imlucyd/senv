
# Senv

## Install

```bash
# Requires go >=1.18
$ go install github.com/luisnquin/senv@latest
```

## Demo

[![demo](https://asciinema.org/a/eZrIbb4eDxX0tO7fWyFop2Zg8.svg)](https://asciinema.org/a/eZrIbb4eDxX0tO7fWyFop2Zg8)

## Configuration

It requires a `senv.yaml` or `senv.yml` file in your `current directory` or `root working directory`
(you can call the program in any subdirectory of the project but for that it requires a git repository
already initialized).

If the working root directory is not found or the program files are not found, you'll not be able to switch.

### File example

```yaml
# senv.yaml
envFile: ./app/.env # optional absolute/relative path
envs:
- name: dev
  variables:
    DATABASE_USER: admin
    DATABASE_PASSWORD: pwd123
    DATABASE_HOST: localhost
- name: dev-2
  variables:
    DATABASE_USER: admin
    DATABASE_PASSWORD: pwd321
    DATABASE_HOST: test-host
- name: prod
  variables:
    DATABASE_USER: admin
    DATABASE_PASSWORD: test-password
    DATABASE_HOST: test-host
- name: preprod
  variables:
    DATABASE_PASSWORD: root123
    DATABASE_HOST: localhost
defaults:
  DATABASE_USER: admin # both variables will be added to the
  DATABASE_PORT: 5432 # selected environment if not declared
useExportPrefix: false # optional
```

## Completions

Completions are very simple.

### Supported shells

- zsh

```bash
 # Add this line to your .zshrc file
 $ source <(senv completion zsh)
```

- bash

```bash
 # Add this line to your .bashrc file
  $ source <(senv completion bash)
```

## LICENSE

[MIT](../LICENSE)
