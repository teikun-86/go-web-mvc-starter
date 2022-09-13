# GoWebBoilerplate
Golang Web MVC Boilerplate

## Installation
1. Clone this repository `git clone https://github.com/teikun-86/go-web-mvc-starter`
2. CD into `go-web-mvc-starter`
3. Run `go run .` command to start web development server or use `github.com/pilu/fresh` to restart development server automatically.
4. Happy Coding!

## Usage
### Routing
To create routing, just go to `routes/routes.go` and add your own routes below this code
```go
// Create your routes here...
mux.HandleFunc("/", controllers.IndexHandler)
mux.HandleFunc("/foo", controllers.FooHandler)
mux.HandleFunc("/bar", controllers.BarHandler)
```
### Controllers
To handle action in each routes, you can create function in `controllers/MainController.go` or create new file in `controllers` directory with package `controllers` example:
```go
//controllers/FooController.go
package controllers
import ...

func FooHandler(w http.ResponseWriter, r *http.Request) {
   // Controller action here
}
```
