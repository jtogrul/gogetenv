# GoGetEnv
Just a simple go program to serve basic request information and environment variables of server over port 80.
Can be useful for testing or troubleshooting clusters, microservices etc.

## Warning
**Don't use on production environment**
Environment variables of a machine can contain sensetive information like Database credentials, Secrets etc. Use with caution.

## Example
Request:
```
$ curl "http://localhost/foo?a=b&c=1&c=2"
```
Response:
```
{
	"url_path": "/foo",
	"url_params": {
		"a": [
			"b"
		],
		"c": [
			"1",
			"2"
		]
	},
	"environment": [
		"PATH=...",
		"FOO=bar"
	]
}
```