# Test_mysql_aws
# For trst Only 

<code>go get -u github.com/go-sql-driver/mysql</code>
<br>
<code>go get -u github.com/aws/aws-lambda-go</code>


Introduction
------------

This is a simple Go program that connects to a MySQL database using the "go-sql-driver/mysql" package.

Installation and Usage
----------------------

1. Clone this repository to your local machine.
2. Make sure you have Go and MySQL installed on your system.
3. Open the `main.go` file and modify the following variables to match your MySQL database configuration: password = "your-password"endpoint = "your-database-endpoint"port = "your-database-port"db\_name = "your-database-name"

2. Save the changes to `main.go`.
3. In your terminal, navigate to the cloned repository and run the following command:
4. If everything is configured correctly, the program should connect to the MySQL database and output "Success!" to the terminal.
Dependencies
------------

- Go v1.13 or higher
- "go-sql-driver/mysql" package
 
License
-------

This project is licensed under the MIT License - see the LICENSE.md file for details.