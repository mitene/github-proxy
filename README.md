# github-proxy

github-proxy is a proxy that can archive any directory in the specified GitHub repository tree and deliver it via HTTP.

# Installation


```
$ go get -u github.com/mitene/github-proxy
```

# Usage

```
$ github-proxy [options]
```
Command line option

| Parameter | Description | Default |
|-----------|---------------------------------------------|---------|
| -port | HTTP server port | 8080 |


## Example to get

```
$ wget --content-disposition "http://localhost:8080/get/mitene/github-proxy?path=cmd"
```
This gets cmd directory of mitene repository in the form of tgz file.

GET params

| Parameter | Description | Default |
|-----------|---------------------------------------------|---------|
| ref | git commit hash or branch name | master |
| type | archive formats. supported zip, tgz, tar,gz | tgz |
| path | path in GitHub repository | / |

## Rules for downloaded file name

[Owner]-[Repository]-[Ref].[Type]

## Working with private repositories

Set personal access token in environment variable `GITHUB_ACCESS_TOKEN`.

### Example

mitene-github-proxy-master.tgz

# License

The source code is licensed MIT.
