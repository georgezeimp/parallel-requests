# parallel-requests
## Description
This is a tool written in the Go programming language, and is designed to do a series of simple tasks as described below:
1. Accept input from the user with a series of web addresses, and optionally an integer using the **-parallel** flag.
1. Spawn a series of goroutines in parallel each of which will send a request to one of the addresses given as input and calculate the md5 hash of the body of the response. No more goroutines will be run in parallel than the value of the **-parallel** flag. If the user did not pass this flag, the tool will default to a maximum of **10** goroutines in parallel.
1. Each goroutine will then create a "result" string which will consist of the the address to which the request was sent, and the md5 hash of the body of the response, separated by a space.
1. Once all goroutines have been completed and returned their result, the list of all results in presented as output to the user.

## Examples
* `./parallel-requests -parallel 3 https://www.company.com http://google.com`
* `./parallel-requests https://www.company.com http://google.com`
* `./parallel-requests` <- This will not return any results.
