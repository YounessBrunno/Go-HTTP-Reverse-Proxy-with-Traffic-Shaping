package (
	"fmt"
	"net/http"
)


func main() {
   http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
   })

   http.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Sleept for 3 seconds!")
   })

   http.ListenAndServe(":9000", nil)
}

