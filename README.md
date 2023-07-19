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

### Create data source

```shell
# [Help]
❯ ./dist/cogsearchctl datasource create --help
Create a datasource for Azure Cognitive Search.

Usage:
  cogsearchctl datasource create [flags]

Flags:
  -f, --bodyFilePath string        body file path (default "body.json")
  -d, --dataSourceName string      data source name (default "dataSourceName")
  -h, --help                       help for create
  -k, --searchApiKey string        search api key (default "searchApiKey")
  -s, --searchServiceName string   search service name (default "searchServiceName")

Global Flags:
      --config string   config file (default is $HOME/.cogsearchctl.yaml)

# [Create data source]
# edit {{env_storage_connection_string}} and {{env_storage_container}} in ./examples/create_datasource.json
❯ ./dist/cogsearchctl datasource create \
    --bodyFilePath ./examples/create_datasource.json \
    --dataSourceName $dataSourceName \
    --searchApiKey $searchApiKey \
    --searchServiceName $searchServiceName
2023/07/19 18:43:21 status code=201
```

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
2023/07/19 18:13:46 status code=201
```

### Create skill set

```shell
# [Help]
❯ ./dist/cogsearchctl skillset create --help
Create a skill set for Azure Cognitive Search.

Usage:
  cogsearchctl skillset create [flags]

Flags:
  -f, --bodyFilePath string        body file path (default "body.json")
  -h, --help                       help for create
  -k, --searchApiKey string        search api key (default "searchApiKey")
  -s, --searchServiceName string   search service name (default "searchServiceName")
  -n, --skillSetName string        skill set name (default "skillSetName")

Global Flags:
      --config string   config file (default is $HOME/.cogsearchctl.yaml)

# [Create skill set]
❯ ./dist/cogsearchctl skillset create \
    --bodyFilePath ./examples/create_skillset.json \
    --searchApiKey $searchApiKey \
    --searchServiceName $searchServiceName \
    --skillSetName $skillSetName
2023/07/19 19:08:03 status code=201
```


## References

- [nohanaga/Azure-Cognitive-Search-Workshop](https://github.com/nohanaga/Azure-Cognitive-Search-Workshop)
