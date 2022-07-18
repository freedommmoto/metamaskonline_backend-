# metamaskonline_backend
backend for support profile , login , register that connect to front-end of project.

```
how to setup 

install api for setup database first
https://github.com/freedommmoto/metamaskonline_api

vim app.env
DB_DRIVER=postgres
DB_SOUECE=postgresql://root:secret@localhost:5432/metamaskonline?sslmode=disable
PUSHER_KEY=put_you_key_here
SERVER_ADDRESS=0.0.0.0:7777
TOKEN_SYMMETRIC_KEY=put_you_key_here
ACCESS_TOKEN_DURATION_TIME=10m

go build main.go
./main
```