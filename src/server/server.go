package server

import (
	"Avito-Challenge/src/database/pgdb"
	"Avito-Challenge/src/handlers"
	"Avito-Challenge/src/repositories/repo_implementation"
	"Avito-Challenge/src/usecases/uc_implementation"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func RunServer(address string, connectionString string) error {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	manager := pgdb.NewPGDBManager()
	err := manager.Connect(connectionString)
	if err != nil {
		fmt.Print("Failed to connect to database")
	} else {
		fmt.Println("Successfully connected to postgres database")
	}

	br := repo_implementation.NewBalanceRepository(manager)
	ar := repo_implementation.NewAccountingRepository(manager)
	ac := uc_implementation.NewAccountingUsecases(ar)
	bc := uc_implementation.NewBalanceUsecases(br, ac)
	bh := handlers.NewBalanceHandlers(bc)

	apiRouter.HandleFunc("/balances", bh.GetByUserId).Methods(http.MethodGet)
	apiRouter.HandleFunc("/balances", bh.AddByUserId).Methods(http.MethodPatch)
	apiRouter.HandleFunc("/reservations", bh.AddReservation).Methods(http.MethodPut)
	apiRouter.HandleFunc("/reservations", bh.CommitReservation).Methods(http.MethodDelete)

	server := http.Server{
		Addr:    address,
		Handler: apiRouter,
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		manager.Disconnect()
		os.Exit(0)
	}()

	fmt.Printf("Balance system server is running on %s\n", address)
	return server.ListenAndServe()
}