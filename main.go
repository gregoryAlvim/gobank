package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gregoryAlvim/gobank/database"
	"github.com/gregoryAlvim/gobank/handlers"
	"github.com/gregoryAlvim/gobank/repositories"
	"github.com/gregoryAlvim/gobank/services"
)

// @title Bank API
// @version 1.0
// @description This is a sample server for a banking application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	// Database connection
	database.InitDB(os.Getenv("DATABASE_URL"))

	// Initialize repository and service
	accountRepo := repositories.NewPsqlAccountRepository()
	accountService := services.NewAccountService(accountRepo)
	accountHandler := handlers.NewAccountHandler(accountService)

	// Router
	r := mux.NewRouter()

	// Swagger
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Handlers
	r.HandleFunc("/account", accountHandler.CreateAccount).Methods("POST")
	r.HandleFunc("/account/{id}/balance", accountHandler.GetBalance).Methods("GET")
	r.HandleFunc("/account/{id}/deposit", accountHandler.Deposit).Methods("POST")
	r.HandleFunc("/account/{id}/withdraw", accountHandler.Withdraw).Methods("POST")
	r.HandleFunc("/account/transfer", accountHandler.Transfer).Methods("POST")
	r.HandleFunc("/account/{id}", accountHandler.CloseAccount).Methods("DELETE")

	// CSRF protection
	csrfMiddleware := csrf.Protect([]byte("32-byte-long-auth-key"))

	// Start server
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", csrfMiddleware(r)))
}
