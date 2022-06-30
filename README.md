# Go Webserver
## Using GoFiber & GORM
## Installation

Install the dependencies.
```sh
go get ./..
```

Start the server.
```sh
go run main.go
```

Run the test.
```sh
go test
```

# Available endpoints
```sh
POST | http://127.0.0.1:3000/api/report/post
POST | http://127.0.0.1:3000/api/report/comment
```

## Example Request
Post to http://127.0.0.1:3000/api/report/post
```javascript
{
    "reason": "not suitable for world NSFW",
    "post_id": 2
}
```