# go_parallel_api_requests
Make api calls in parallel using goroutines.
Nothing too fancy. Just a demonstration for learning purposes.

## Example Server
 First, run the example server, and leave it running.
```console
foo@bar:~/go_parallel_api_requests$ ./quickserver
```
Now you have a fake "real world" example in which you have to receive JSON data from three very slow endpoints.

## Testing using curl
> Note: I recommend installing the useful jq utility first.

Open a new terminal window in the project folder.
Now you can test the endpoints in curl like this:
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
