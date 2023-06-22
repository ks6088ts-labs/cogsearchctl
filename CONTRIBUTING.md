# Contributing guidelines

## Prerequisites

- [Go](https://golang.org/)
- [Git](https://git-scm.com/)
- GNU Make
- [goreleaser](https://github.com/goreleaser/goreleaser)

## Build

```shell
$ make build
```

## Test

```shell
$ make build
```

## Release

Tagging a commit and push it to the repo will create a new release with configured [GitHub action](https://github.com/ks6088ts-labs/cogsearchctl/actions), under the [Releases](https://github.com/ks6088ts-labs/cogsearchctl/releases/) section.

```console
$ git tag v0.1.2
$ git push --tags
```

Before pushing your new tag, please test it locally with:

```console
$ make snapshot
```

If no error, the action should work.
