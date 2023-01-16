# todoapp
 
1. First download and install [Go SDK](https://go.dev/doc/install).
2. Clone this project
```sh
$ git clone https://github.com/aliavd1/todoapp.git
```

3. Go to workdir and download dependency packages
```sh
$ cd todoapp
$ go mod download
```

4. After installation, run todoapp project
```sh
$ go run main.go
```
You can see todoapp web page on http://localhost:8000/todos

Also, you can use docker
```sh
$ cd todoapp
$ docker-compose up -d
```
You can see todoapp web page on http://localhost/todos
