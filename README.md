# go_parallel_api_requests
Make api calls in parallel using goroutines.

> Note: I recommend installing the useful jq utility first.

## Testing using curl

 First, run the example server, leave it running, then open a new terminal window and get back to the project folder.
```console
foo@bar:~/go_parallel_api_requests$ ./quickserver
```
Now you have a fake "real world" example in which you have to receive JSON data from three very slow endpoints.

You can test them out in curl like this:
```console
foo@bar:~/go_parallel_api_requests$ curl localhost:4501/number1 | jq
foo@bar:~/go_parallel_api_requests$ curl localhost:4501/number2 | jq
foo@bar:~/go_parallel_api_requests$ curl localhost:4501/number3 | jq
```
You will notice that each of them takes three seconds to respond, so the total execution time will be around 9 seconds.

But if you run them in parallel like this:
```console
foo@bar:~/go_parallel_api_requests$  curl --parallel --parallel-immediate --config websites.txt | jq
```
You will get the job done in about three seconds.

## Example in GO

In our very simple program we want to get the price value from each of these endpoints and add them up to get a total price.

### Go example without go routines
```console
foo@bar:~/go_parallel_api_requests$ git checkout main
foo@bar:~/go_parallel_api_requests$ go run main.go
```
### Go example with go routines
```console
foo@bar:~/go_parallel_api_requests$ git checkout parallel
foo@bar:~/go_parallel_api_requests$ go run main.go
```
