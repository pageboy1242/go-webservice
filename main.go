package main
import (
    "net/http"

    "github.com/pageboy1242/go-webservice/controllers"
) 

func main() {
    controllers.RegisterControllers()
    http.ListenAndServe(":3000", nil)
}
