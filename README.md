## Local Environment

Both tests and dev server requires running database docker container in background:

```bash
docker-compose up -d db
make test
```

Dev server requires fswatch:

```bash
brew install fswatch
make serve
```

## Swagger Documentation

Swagger file generation requires `go-swagger` and `redoc-cli` dependecies.

```bash
brew tap go-swagger/go-swagger
brew install go-swagger
npm install -g redoc-cli
```

Use makefile to generate swagger documents

```bash
make swagger-html
```
