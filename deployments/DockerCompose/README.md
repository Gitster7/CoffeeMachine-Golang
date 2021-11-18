## Docker Compose

Edit .env file to change variables from docker-compose file and then to run in detached mode run:
```
docker-compose up -d
```
to start service coffeemachine. New docker compose command from new docker cli is in experimental stage and still is not fully working so stick to docker-compose command.

To stop docker-compose service run:
```
docker-compose down
```

Variables and their defaults are:
```
COFFEEMACHINE_IMAGE=github.com/kostovic/coffeemachine
COFFEEMACHINE_TAG=restapiv2.0
COFFEEMACHINE_HTTP_PORT=3000
GIN_MODE=release
```