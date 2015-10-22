package main
import (
   	"fmt"
    "httprouter"
    "net/http"
    "encoding/json"
    "log"
)

type RequestData struct{
	Name string
}

type ResponseData struct{
	Greetings string
}

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
		requestData := new(RequestData)
		responseData := new(ResponseData)
		decoder := json.NewDecoder(req.Body)
		error := decoder.Decode(&requestData)
		if error != nil {
			log.Println(error.Error())
			http.Error(rw, error.Error(), http.StatusInternalServerError)
			return
		}
		responseData.Greetings = "Hello," + requestData.Name + "!"
		outgoingJSON, err := json.Marshal(responseData)
		if err != nil {
			log.Println(error.Error())
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		fmt.Fprint(rw, string(outgoingJSON))
}

func main() {
    mux := httprouter.New()
    mux.POST("/hello", hello)
    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: mux,
    }
    server.ListenAndServe()
}
