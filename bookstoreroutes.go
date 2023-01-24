package routes
                                        //this are all routing part
import(

	"github.com/gorilla/mux"
	"github.com/balaram/bookstore/pkg/controllers"
)


var RsgisterBookstoreRoutes = func (router *mux.Router){
       router.HandleFunc("/book/",controllers.CreateBook).Methods("POST")
	   router.HandleFunc("/book/",controllers.GetBook).Methods("GET")
	   router.HandleFunc("/book/{bookId}",controllers.GetBookById).Methods("GET")
	   router.HandleFunc("/book/{bookId}",controllers.UpdateBook).Methods("PUT")
	   router.HandleFunc("/book/{bookId}",controllers.DeleteBook).Methods("DELETE")
	   
}