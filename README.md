# go-departments-project

### Steps to build the project
1. Start containers
```
docker compose -f deployments/docker-compose.yml up -d
```

2. Execute database migrations
```
migrate -path configs/migrations -database "mysql://root:parola123\!@tcp(localhost:3306)/project" -verbose up
```
3. Start the golang webserver
```
make serve
```


### Steps to clean the project

1. Clean (stop and remove) containers
```
docker compose -f deployments/docker-compose.yml down
```


### Extra instructions
* The database migrations need the [migrate](https://github.com/golang-migrate/migrate) tool. If you don't want to use it, all the up/down scripts are written in SQL and available at
```
configs/migrations/*
```
* Postman requests can be found at 
```
configs/postman/postman_collection.json
```

* Roll back the database migrations
```
migrate -path configs/migrations -database "mysql://root:parola123\!@tcp(localhost:3306)/project" -verbose down
```

### Containers
1. Redis - used to store user sessions
2. Redisinsight - client for Redis
3. Mysql - the main SQL database