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