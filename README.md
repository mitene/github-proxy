# github-proxy

github-proxy is a proxy that can archive any directory in the specified GitHub repository tree and deliver it via HTTP.

## Installation


```
$ go get -u github.com/mitene/github-proxy
```

## Usage

```
$ github-proxy [options]
```

The command starts the HTTP server.

### Command option

| Parameter | Description | Default |
|-----------|---------------------------------------------|---------|
| -port | HTTP server port | 8080 |


## Endpoint

### repo

Download starts after compressing the specified repository path.

```
/repo/{owner}/{repository}
```

#### GET params

| Parameter | Description | Default |
|-----------|---------------------------------------------|---------|
| ref | git commit hash or branch name | master |
| type | archive formats. supported zip, tgz, tar,gz | tgz |
| path | path in GitHub repository | / |


#### Example to get

```
$ wget --content-disposition "http://localhost:8080/repo/mitene/github-proxy?path=cmd"
```

This gets cmd directory of mitene repository in the form of tgz file.

#### Rules for downloaded file name

`Owner`-`Repository`-`Ref`.`Type`

### Example

`mitene-github-proxy-master.tgz`

## Working with private repositories

Set personal access token in environment variable `GITHUB_TOKEN`.

## License

The source code is licensed MIT.
