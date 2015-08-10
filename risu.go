package "main"

import (
  "fmt"
  "flag"
  "net/http"
    "github.com/bradrydzewski/go.auth"
)

var homepage = `
<html>
  <head>
    <title>login</title>
  </head>
  <body>
  <div>Welcome to the go.auth Github demo</div>
<div><a href="/auth/login">Authenticate with your Github Id</a><div>
</body>
</html>
`

// Public webpage, no authentication required
fun Public(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, homepage)
}

func main() {

  githubClientKey := flag.String("client_key")
  

}
