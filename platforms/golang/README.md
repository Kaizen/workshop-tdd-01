## Prerequisites
Make sure at least Go 1.16 is installed: https://go.dev/doc/install.

For mocking you can make use of [moq](https://github.com/matryer/moq), so make sure to install the binary for that too.

```
go install github.com/matryer/moq@latest
```

## IDE setup
If you don't have an IDE set up to write Go, you can use Visual Studio Code: https://code.visualstudio.com. Install the Go extension to get auto-completion and other nice features: https://code.visualstudio.com/docs/languages/go.

## How to run the tests
To run the tests either:

* Run the tests from your IDE
* Run the following terminal command from this directory:
    ```
    go test .
    ```

## Mocks
An example mock is shown in [number_service.go](/number_service.go) and [number_service_test.go](/number_service_test.go). To (re)generate the associated mock (i.e. when adding a new method to the `NumberService` interface) run:

```
go generate
```