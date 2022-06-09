# to-do-app
An app to track your tasks. Built using Golang and GoFiber.

## Start-up Instructions

1. Make a clone of the repository
2. Spin up the Docker containers for the database and cache using

    ```docker run -it --name todo-postgres -e POSTGRES_USER=todo  -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=todo -p 5432:5432 -d postgres```

    and

    ```docker run -it --name todo-redis -p 6379:6379 -d redis```

    If you change the user or password, be sure to update those in the `.env` file.

3. In the project root directory, run the application using `go run main.go`