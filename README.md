# cogsearchctl
A command line tool for communicating with Azure Cognitive Search

## How to build

### Prerequisites
- [Go](https://golang.org/doc/install) 1.20 or later

```shell
# generate cogsearchctl in ./dist
make build
```

## How to use

Please build cogsearchctl or put a binary from [Releases](https://github.com/ks6088ts-labs/cogsearchctl/releases) in your PATH.

### Create index

```shell
# [Help]
❯ ./dist/cogsearchctl index create --help
Create an index for Azure Cognitive Search.

Usage:
  cogsearchctl index create [flags]

Flags:
  -f, --bodyFilePath string        body file path (default "body.json")
  -h, --help                       help for create
  -k, --searchApiKey string        search api key (default "searchApiKey")
  -i, --searchIndexName string     search index name (default "searchIndexName")
  -s, --searchServiceName string   search service name (default "searchServiceName")

Global Flags:
      --config string   config file (default is $HOME/.cogsearchctl.yaml)

# [Create index]
❯ ./dist/cogsearchctl index create \
    --bodyFilePath ./examples/create_index.json \
    --searchApiKey $searchApiKey \
    --searchIndexName $searchIndexName \
    --searchServiceName $searchServiceName
create called
2023/07/19 18:13:46 status code=201
```

## References

- [nohanaga/Azure-Cognitive-Search-Workshop](https://github.com/nohanaga/Azure-Cognitive-Search-Workshop)
