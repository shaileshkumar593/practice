package main


import(
	"log"
	"net/http"
	"api/repository"
	"api/service"
)


func main(){
	repo :=  repository.NewEmployeeRepository()

	svc := service.NewEmployeeService(repo)
	hanlerlayer := 
}