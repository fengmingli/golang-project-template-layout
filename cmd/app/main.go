package main

import (
	"context"
	"log"
	"net/http"

	"github.com/fengmingli/golang-project-template-layout/internal/adapter/controller"
	"github.com/fengmingli/golang-project-template-layout/internal/adapter/repository"
	"github.com/fengmingli/golang-project-template-layout/internal/domain/service"
	"github.com/fengmingli/golang-project-template-layout/internal/infrastructure/config"
	"github.com/fengmingli/golang-project-template-layout/internal/infrastructure/db"
	"github.com/fengmingli/golang-project-template-layout/internal/infrastructure/router"
	"github.com/fengmingli/golang-project-template-layout/internal/usecase"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.LoadConfig,
			db.NewDatabase,
			repository.NewUserRepository,
			service.NewUserService,
			usecase.NewCreateUserUseCase,
			usecase.NewGetUserUseCase,
			usecase.NewListUsersUseCase,
			usecase.NewUpdateUserUseCase,
			controller.NewUserController,
			router.NewRouter,
		),
		fx.Invoke(startServer),
	)

	app.Run()
}

func startServer(lc fx.Lifecycle, r *mux.Router) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				log.Println("Starting server at :8080")
				if err := http.ListenAndServe(":8080", r); err != nil {
					log.Fatal(err)
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			log.Println("Stopping server")
			return nil
		},
	})
}
