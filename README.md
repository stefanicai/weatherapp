# Weatherapp

Simple API to provide basic weather info

## How to setup and use
### Dependencies
- Golang version 1.3 (install if not available). Older version will work too, just change version in `src/go.mod`

### Starting the app
- clone repo to your local
- change current directory to be `<project path>/src` (basically `cd weatherapp/src`)
- run the following command `go run main.go`.

### Using the app
Go to `http://localhost:8080/` - that will give you the weather for Melbourne,AU

If you'd like the weather for another city, use the `query` parameter. E.g. `http://localhost:8080?query=Sydney,AU` will give you the weather for Sydney.


## Implementation comments

## Notes
- `weatherstack` implementation uses an invalid `access_key`, thus it won't work. That's intentional to show the failover. But the implementation is done, so if you have a valid key, it should work.
- For `openweather` I haven't handled the case where server returns an error. An example of that is done for `weatherstack` since that call will fail. That would need to be implemented for a commercial solution
- I've used the standard `log` package even though it's pretty limited. Thought it'd do for this test. For a commercial solution that might not be a good fit, depends on how you manage logs (e.g. splunk/graylog, simple files etc).
- The http server code is pretty basic. E.g. `CORS` would need to be handled if this is to be exported as an API. As well as proper/appropriate HTTP codes and error messages to be returned. As well as potentially a healthcheck to be used by kubernetes etc
- For a commercial implementation, I'd probably use `GraphQL` to give the client the possibility to select what info to receive.
- Caching (distributed) might be useful if the service is highly used. It could also cover for cases when both services are down.

### Some more about the tech
- json parsing is done differently for the two services:
    - query based parsing (this was the most time consuming part of the exercise due to `gojsonq` not working as expected):
        - For `openweather` I use that. It could be easier to read than replicating the whole message structure as `structs`, when message is very complex. I used `ajson` framework for that, though I initially tried with `gojsonq` which I find it has a nicer interface. Unfortunately it doesn't seem to support our scenario or I didn't figure out how (a PR to the project shouldn't be too complex - basically the `Only` or `Select` methods need to support regular props not just arrays)
    - standard `json` package
        - This is probably a better fit for this test as the message doesn't have as much nesting and only requires a few fields.
        
        
### Mention
I have much more experience with Java (over 10 years) than Go (less than 1 year), but since you're using Go, I thought I'd write this in Go. I enjoyed it more than if I'd have done it in Java too!

### Questions/updates
You can add Github issues with questions or if you'd like me to do something else on it. I enjoyed writing it, happy to add more to it if you'd like.
