# actgroup

`actgroup` is a command line utility to wrap a command with log grouping for GitHub Actions
as described here:
https://docs.github.com/en/actions/using-workflows/workflow-commands-for-github-actions?tool=bash#grouping-log-lines

## Installation

```sh
go install github.com/ronny/actgroup/cmd/actgroup@latest
```

Or install a prebuilt binary from [the Releases page](https://github.com/ronny/actgroup/releases).

## Usage

### With an explicit group title

```sh
actgroup -title "Show current time in UTC" date -u
```

Output:

```
::group::Show current time in UTC
Wed 10 Apr 2024 01:02:03 UTC
::endgroup::
```

### With an implicit group title

When the title is not specified, it's inferred from the command and the args.

```sh
actgroup date -u
```

Output:

```
::group::date -u
Wed 10 Apr 2024 10:03:46 UTC
::endgroup::
```

### Auto detect GitHub Actions

By default `actgroup` auto detects GitHub Actions and use the correct log grouping format.
based on the existence of `GITHUB_ACTIONS=true` env var.

To force it to always output using the GitHub Actions format, use `--force`:

```sh
actgroup -force -title "Show time" date -u
```
