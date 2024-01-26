# certlint
Small ZLint API to sit in a side car model with an enrichments service

## Build & Run

```shell
go build
./certlint
```

## API

*Note: The service runs on port 8080 by default.*

* `POST /pem`: PEM certificate handler
* `POST /der`: DER certificate handler

See the `examples/api.py` example for usage.
