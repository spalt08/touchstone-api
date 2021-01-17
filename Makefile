include ./config/development.env
export

PID = /tmp/touchstone-api.pid
GO_FILES = cmd/api/server.go

swagger: 
	swagger generate spec -m -o docs/swagger.json

swagger-html: swagger
	redoc-cli bundle docs/swagger.json -o docs/swagger.html 

test:
	go test pkg/user/*.go -v

start:
	go run -mod=vendor $(GO_FILES) & echo $$! > $(PID)
	@echo "STARTED touchstone-api"

kill:
	-kill `pgrep -P \`cat $(PID)\`` && \
	 kill `cat $(PID)`
	@echo "STOPED touchstone-api" && printf '%*s\n' "40" '' | tr ' ' -

restart: kill start

prepare:
	go mod vendor

serve: prepare start
	fswatch -e vendor -or --event=Updated /home/touchstone-api/pkg | xargs -n1 -I {} make restart
