# To-do API Server

Build on Go and PostgreSQL

## Get started

You need to install [Docker](https://docs.docker.com/get-docker/) first.
At the root directory, 
```
docker-compose up -d [--build]
```
You will be able to access  
db http://localhost:5432 (user: postgres, password: password)  
pgadmin http://localhost:5050 (user: test@test.com password: password)  
api server http://localhost:8080 

#### Troubleshooting
When you run application first time, you need to go to pgadmin and create server or use psql in the postgres container manually.
Right click on "Servers" and select "Create" > "Server",
only thing you need to be aware is typing "db" or "Docker Container Alias You Changed" in Host name/address at "Connection" tab.  
After it's done, restart api server container and you are good to go.

Note:  
All the environment variables you can set are all listed in .env.
