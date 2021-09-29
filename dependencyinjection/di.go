package dependencyinjection

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	//Greet(os.Stdout, "Long")
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreetHandler)))
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Long")
}
