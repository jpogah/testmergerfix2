# testmergerfix2

## hello-web service

Run locally with one command:

```bash
make run-hello-web
```

Defaults:
- `HTTP_PORT=8080`
- `APP_ENV=development`
- `HELLO_MESSAGE="hello from hello-web"`
- `READ_HEADER_TIMEOUT_SECONDS=5`
- `SHUTDOWN_TIMEOUT_SECONDS=10`

Endpoints:
- `GET /healthz`
- `GET /hello`

Optional Docker flow:

```bash
make docker-run-hello-web
```
