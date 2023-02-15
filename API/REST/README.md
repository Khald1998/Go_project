# Create a module in which you can manage dependencies

## To install Gorilla mux, write the command:
"go get -u github.com/gin-gonic/gin"

## To install CompileDaemon, write the command: "https://github.com/githubnemo/CompileDaemon"
go get github.com/githubnemo/CompileDaemon
go install github.com/githubnemo/CompileDaemon

## To install godotenv, write the command: "https://github.com/joho/godotenv"
go get github.com/joho/godotenv


## To install GORM, write the command
"go get -u github.com/gorilla/mux"

## To install GORM, write the command:
"go get -u gorm.io/gorm"
## To install the MySQL driver, write the command:
"go get -u gorm.io/driver/mysql"


## We will have a little understanding of the Gorilla mux and GORM package.

### gorilla/mux will be used to route the incoming HTTP request to the correct method that handles the specific operation. For instance, when a client initiates a POST request to the endpoint, routing makes the machine recognize where the request should be routed.

### GORM is an ORM with helper functions that carry queries or execute commands over a specific database.
### Now, we will create the main.go file in the rest_api directory and write the basic syntax in the main.go file.



# How to install 

## step one

### Run the go mod init command, giving it the path of the module your code will be in.
<code>go mod init REST</code>
test

