<div align="center">

# miactl

[![Build Status][github-actions-svg]][github-actions]
[![Go Report Card][go-report-card]][go-report-card-link]
[![GoDoc][godoc-svg]][godoc-link]

</div>

**miactl is the cli of the mia-platform DevOps Console**

## Install

### Using brew

```sh
brew install mia-platform/tap/miactl
```

### Using Go

This library requires golang at version >= 1.13

```sh
go get -u github.com/mia-platform/miactl
```

## Example usage

### Get projects

```sh
miactl get projects --apiKey "your-api-key" --apiCookie "sid=your-sid" --apiBaseUrl "https://console.url/"
```

### Projects help

```sh
miactl help
```


## Enabling shell autocompletion

miactl provides autocompletion support for Bash, Zsh and Fish, which can save you a lot of typing.

### Bash

Completion could be generated running the `miactl completion bash` command.
To make this completion work, you should have [bash completion](https://github.com/scop/bash-completion)
correctly installed.

Warning: there are two versions of `bash-completion`, v1 and v2. V1 is for Bash 3.2 (which is the default on macOS), and v2 is for Bash 4.1+. The miactl completion script requires `bash-completion` v2 and Bash 4.1+.

**For linux:**

You could install `bash-completion` running `apt-get install bash-completion`.

The above command create `/usr/share/bash-completion/bash_completion`, which is the main script of bash completion.
Try to run the command `type _init_completion`. If the command succeeds, you're already set. Otherwise, add to the `~/.bashrc` file:

```sh
source /usr/share/bash-completion/bash_completion
```

After the correct installation if bash completion, you could make the completion works running:
```sh
miactl completion bash >/etc/bash_completion.d/miactl
```

**For osx:**

Check the bash version running `echo $BASH_VERSION`. If the version is less than 4.1, you should install a new bash version with Homebrew running

```sh
brew install bash
```

Reload your shell and run the command `echo $BASH_VERSION` to verify the bash version.

Once installed the correct bash version, you should install `bash-completion` to v2. To check if it is already installed, run `type _init_completion`. If not, you can install it with Homebrew

```sh
brew install bash-completion@2
````

And add in your `~/.bashrc` file:
```sh
export BASH_COMPLETION_COMPAT_DIR="/usr/local/etc/bash_completion.d"
[[ -r "/usr/local/etc/profile.d/bash_completion.sh" ]] && . "/usr/local/etc/profile.d/bash_completion.sh"
```

After the correct installation if bash completion, you could make the completion works running:
```sh
miactl completion bash >/etc/bash_completion.d/miactl
```

### Fish

Completion could be generated running the `miactl completion fish` command.

To make this completion work, you should run:
```sh
miactl completion fish >~/.config/fish/completions/miactl.fish
```

### Zsh

Completion could be generated running the `miactl completion zsh` command

The generated completion script should be put somewhere in your $fpath named _miactl.


[github-actions]: https://github.com/mia-platform/miactl/actions
[github-actions-svg]: https://github.com/mia-platform/miactl/workflows/Test%20and%20build/badge.svg
[godoc-svg]: https://godoc.org/github.com/mia-platform/miactl?status.svg
[godoc-link]: https://godoc.org/github.com/mia-platform/miactl
[go-report-card]: https://goreportcard.com/badge/github.com/mia-platform/miactl
[go-report-card-link]: https://goreportcard.com/report/github.com/mia-platform/miactl
[semver]: https://semver.org/
