# GoGetEnv
Just a simple go program to echo basic HTTP request details and serve environment variables of host over port 80.
Can be useful for testing or troubleshooting load balancer, proxy, container setups etc.

## Warning
**Don't use on production environment**
Environment variables of a machine can contain sensetive information like Database credentials, Secrets etc. Use with caution.

## Example
Request:
```
$ curl -H "X-Test-Id: 1" "http://localhost/foo?a=b&c=1&c=2"
```
Response:
```
{
	"host": "localhost",
	"path": "/foo",
	"url_params": {
		"a": [
			"b"
		],
		"c": [
			"1",
			"2"
		]
	},
	"header": {
		"Accept": [
			"*/*"
		],
		"User-Agent": [
			"curl/7.64.0"
		],
		"X-Test-Id": [
			"1"
		]
	},
	"environment": [
		"PATH=...",
		"FOO=bar"
	]
}
```

## Docker image
You may use the Docker image of `gogetenv` to test or debug your container setup:

```
https://hub.docker.com/r/jtogrul/gogetenv
```

## Example on Kubernetes
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: gogetenv
  labels:
    app: gogetenv
spec:
  replicas: 1
  selector:
    matchLabels:
      app: gogetenv
  template:
    metadata:
      labels:
        app: gogetenv
    spec:
      containers:
      - name: gogetenv
        image: jtogrul/gogetenv:latest
        ports:
        - containerPort: 80
```