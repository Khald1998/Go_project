# What is it
<p> A REST api that do an add and other operations and store the state on an RDS</p>

Here is the state:
```golang
type State struct {
	ID  int    `json:"ID"`
	X   int    `json:"X"`
	Y   int    `json:"Y"`
	OP  string `json:"OP"`
	RES int    `json:"RES"`
}
```
Here is the inputs:
```golang
var (
	username = "x"
	password = "x"
	endpoint = "x"
	port     = "x"
	db_name  = "x"
)
```

and it is going to be running on endpoint of port: 
<code>router.Run("localhost:8080")</code>

with those options:
```golang
	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/state/get", GetState).Methods("GET")
	router.HandleFunc("/state/get/{ID}", GetStateID).Methods("GET")
	router.HandleFunc("/state/add", CreateState).Methods("GET")

```


## How to install 

### Step one
<p> Run the go mod init command, giving it the path of the module your code will be in, in ouer case it is REST.</p>
<code>go mod init REST</code>

### Step two
<p> Install Gin Web Framework , write the command</p>
<code>go get -u github.com/gin-gonic/gin</code>

### Step three
<p> Install Gorilla mux, write the command</p>
<code>go get -u github.com/gorilla/mux</code>

### Step five
<p> install the MySQL driver, write the command</p>
<code>go get -u gorm.io/driver/mysql</code>


## We will have a little understanding of the Gorilla mux and GORM package.

<p> gorilla/mux will be used to route the incoming HTTP request to the correct method that handles the specific operation. For instance, when a client initiates a POST request to the endpoint, routing makes the machine recognize where the request should be routed.</p>

<p> GORM is an ORM with helper functions that carry queries or execute commands over a specific database.</p>
<p> Now, we will create the main.go file in the rest_api directory and write the basic syntax in the main.go file.</p>


