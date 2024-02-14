# golangApi
Create a basic api to Crud on DynamoDb

 <img style="display:block; margin:auto; border-radius: 10px; padding: 4px; background: rgba(0, 0, 0, 0.2); box-shadow: 0 0 20px rgba(0, 0, 0, 0.5);" src="images/entities.png"  />

#### Initialize project

```sh
 go init github.com/name 
```

#### Simple example
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```
#### How debug?
```json
{
  "version": "2.0",
  "configurations": [
    {
      "name": "Crud",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "env": {
        "API_ENV": "qa", 
      },
       "args": ["api"],
      "program": "${workspaceFolder}/src/main.go"
    },

  ]
}
```



