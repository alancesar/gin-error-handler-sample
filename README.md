# Handling errors in gin-gonic

## Starting application
```shell
go mod dowload
go run main.go
```

## Request

```shell
curl --location --request POST 'http://localhost:8080/api' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "some-id",
    "fail_database": false,
    "fail_service": false
}'
```

## Errors
There are two errors struct: `database.Err` and `service.Err`.
They implement `error` interface `Error() string` and `Is(target error) bool`,
returning true if `target` is `pkg.InternalErr`.

## Middleware
The middleware `ErrorHandlerMiddleware()` get the errors from gin-gonic context.
If present, check if it is a `pkg.InternalErr`.
Case true, return an internal server error; otherwise, return bad request.  
This middleware also checks if `err.Err` is a `database.Err` and prints additional logs.

## Handler
The `api.Handler` call run `c.Error(err)` to append errors in gin-gonic context.

## Additional information
- [errors.Is documentation](https://pkg.go.dev/errors#Is)
- [gin-gonic bindings](https://github.com/gin-gonic/gin#model-binding-and-validation)