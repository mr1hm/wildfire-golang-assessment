# wildfire-golang-assessment

## Time Taken
- ~4 hours

## Inital Setup
- Please clone this repository to a local directory of your choice
- At the root of this project's directory, please add a new `config.json` file that contains this:
```
{
  "port": "6000" // Or a port of your choice that you know to be open
}
```
- The server will pull the port value from this file
#### Build/Run
- In terminal, change your path to the `{project root}/cmd/wildfire-golang-assessment` directory
- You can do 1 of 2 things:
1. (Recommended) - this will get all dependencies (if any) automatically
- In terminal, run `go build`
- Verify that a binary file has been created in `{project root}/cmd/wildfire-golang-assessment`
- Then run the binary by running the command `./wildfire-golang-assessment` in terminal
2.
- In terminal, simply run `go run main.go` from `{project root}/cmd/wildfire-golang-assessment`

## Routes
#### "/joke"
- Responds with a json string that replaces "John Doe" within the joke with the name received by the random name service
*Request Example:*
```
curl http://localhost:6000/joke
```
*Response Example:*  
*Client:*
```
my-mbp :: ~ » curl http://localhost:6000/joke
"Zinab Gallien can write multi-threaded applications with a single thread."
```
*Server Logging:*
```
2022/09/08 20:50:37 Reading config file...
2022/09/08 20:50:37 Server starting on Port 6000...
2022/09/08 20:50:45 Random Full Name: Zinab Gallien
2022/09/08 20:50:45 Response: Zinab Gallien can write multi-threaded applications with a single thread.
2022/09/08 20:50:45 counter is now: 2
```

## Adding additional URLs
Adding additional URLs to fetch data from shouldn't be too difficult, you can simply do the following:
- Add the new URL to the existing `URLs` global variable in the `api` package
- Define a new struct for the response data in the `api` package
- Add a case for `req.Service` for the new URL in the `GetURL()` function in the `api` package
- Add a type case for the type switch in the `Set()` handler in the `api` package
- The logic that handles manipulating the string to replace the name "John Doe" in the `getMessage()` handler may not be suitable for your newly requested data, so be sure to add additional logic if needed

## Additional comments
- While the app runs fine, when making very rapid requests there's an issue that occurs on `L58` in `/internal/api/api.go` randomly (during decode of name response)
- An invalid character 'u' is present somewhere, but I was unable to find out what this was within the time limit. It seems like the name api is throwing an error and sending an unexpected data type back or using some kind of character that Go's JSON decoder is decoding to something that starts with `'u'`, but I was unable to pinpoint the issue unfortunately
- There's commented out code in that same block that I've been using to try and debug the issue

- I wasn't able to implement any unit/integration tests within the time limit, but these can be added (along with benchmarks) in the test files (i.e. `/internal/api/api_test`)
