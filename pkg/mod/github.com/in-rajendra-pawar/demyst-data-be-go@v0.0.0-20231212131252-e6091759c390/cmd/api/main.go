package main
import(
	//"net/http"
	"log"
	"github.com/in-rajendra-pawar/demyst-data-be-go/internal/server"
)

func main(){
	log.Println("Starting api server")
	s := server.NewServer()
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}