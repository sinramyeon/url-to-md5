# url-to-md5
lets make a tool that takes a http requests and prints the address of the request along with the MD5 hash of the response!

## requirements 
 - build a tool which makes http requests and prints the address of the request along with the MD5 hash of the response.
-  The tool must be able to perform the requests in parallel so that the tool can
complete sooner. The order in which addresses are printed is not important.
- The tool must be able to limit the number of parallel requests, to prevent
exhausting local resources. The tool must accept a flag to indicate this limit, and it
should default to 10 if the flag is not provided.
- The tool must have unit tests
- A README.md must be included describing the usage of this tool.

## how to run
```
build:
	go build -o md5hasher

run:
     ./md5hasher -parallel 3 google.com facebook.com yahoo.com yandex.com
     ./md5hasher google.com facebook.com yahoo.com yandex.com

unit-test:
	go test -test.v -run ''
```