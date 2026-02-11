# testmergerfix2

## hello-web service

Simple web app that serves a hello response and a health check endpoint.

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

Default local URL:

```text
http://localhost:8080
```

Quick verification:

```bash
curl http://localhost:8080/hello
curl http://localhost:8080/healthz
```

Endpoints:
- `GET /healthz`
- `GET /hello`

Optional Docker flow:

```bash
make docker-run-hello-web
```
