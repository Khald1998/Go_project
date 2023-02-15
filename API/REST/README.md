# What is it
<p> A REST api that do an add operations and store the state</p>
```
type event struct {
	ID  int `json:"id"`
	X   int `json:"x"`
	Y   int `json:"y"`
	RES int `json:"res"`
}
```


## How to install 

### Step one
<p> Run the go mod init command, giving it the path of the module your code will be in, in ouer case it is REST.</p>
<code>go mod init REST</code>

### Step two
<p> Install Gorilla mux, write the command</p>
<code>go get -u github.com/gin-gonic/gin</code>

