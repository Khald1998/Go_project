# What is it
<p> A REST api that do an add operations and store the state</p>

Here is the state:
```golang
type event struct {
	ID  int `json:"id"`
	X   int `json:"x"`
	Y   int `json:"y"`
	RES int `json:"res"`
}
```

and it is going to be running on endpoint of port: 
<code>router.Run("localhost:8080")</code>

with those options:
```golang
	router.GET("/events/show", getEvents)
	router.GET("/events/show/:ID", getEventsByID)
	router.POST("/events/addJ", addJEvents)
	router.POST("/events/addQ", addQEvents)
	router.PUT("/events/updata", updateEvent)
	router.DELETE("/events/delete", deleteEvent)

```


## How to install 

### Step one
<p> Run the go mod init command, giving it the path of the module your code will be in, in ouer case it is REST.</p>
<code>go mod init REST</code>

### Step two
<p> Install Gin Web Framework, write the command</p>
<code>go get -u github.com/gin-gonic/gin</code>

