// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/risingwavelabs/wavekit/internal/app"
	"github.com/risingwavelabs/wavekit/internal/auth"
	"github.com/risingwavelabs/wavekit/internal/config"
	"github.com/risingwavelabs/wavekit/internal/conn/http"
	"github.com/risingwavelabs/wavekit/internal/conn/meta"
	"github.com/risingwavelabs/wavekit/internal/conn/metricsstore"
	"github.com/risingwavelabs/wavekit/internal/conn/sql"
	"github.com/risingwavelabs/wavekit/internal/controller"
	"github.com/risingwavelabs/wavekit/internal/globalctx"
	"github.com/risingwavelabs/wavekit/internal/macaroons"
	"github.com/risingwavelabs/wavekit/internal/metrics"
	"github.com/risingwavelabs/wavekit/internal/model"
	"github.com/risingwavelabs/wavekit/internal/server"
	"github.com/risingwavelabs/wavekit/internal/service"
	"github.com/risingwavelabs/wavekit/internal/task"
	"github.com/risingwavelabs/wavekit/internal/worker"
	"github.com/risingwavelabs/wavekit/internal/worker/handler"
)

// Injectors from wire.go:

func InitializeApplication() (*app.Application, error) {
	configConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	globalContext := globalctx.New()
	modelInterface, err := model.NewModel(configConfig)
	if err != nil {
		return nil, err
	}
	taskStoreInterface := task.NewTaskStore()
	keyStore := macaroons.NewStore(modelInterface, taskStoreInterface)
	caveatParser := auth.NewCaveatParser()
	macaroonManagerInterface := macaroons.NewMacaroonManager(keyStore, caveatParser)
	authInterface, err := auth.NewAuth(macaroonManagerInterface)
	if err != nil {
		return nil, err
	}
	sqlConnectionManegerInterface := sql.NewSQLConnectionManager(modelInterface)
	risectlManagerInterface, err := meta.NewRisectlManager(configConfig)
	if err != nil {
		return nil, err
	}
	metricsManager, err := metricsstore.NewMetricsManager(modelInterface, configConfig)
	if err != nil {
		return nil, err
	}
	metaHttpManagerInterface := http.NewMetaHttpManager()
	serviceInterface := service.NewService(configConfig, modelInterface, authInterface, sqlConnectionManegerInterface, risectlManagerInterface, metricsManager, metaHttpManagerInterface, taskStoreInterface)
	controllerController := controller.NewController(serviceInterface, authInterface)
	initService := service.NewInitService(modelInterface, serviceInterface)
	serverServer, err := server.NewServer(configConfig, globalContext, controllerController, authInterface, initService)
	if err != nil {
		return nil, err
	}
	metricsServer := metrics.NewMetricsServer(configConfig, globalContext)
	taskHandler := handler.NewTaskHandler(risectlManagerInterface, taskStoreInterface, metaHttpManagerInterface)
	workerWorker, err := worker.NewWorker(globalContext, modelInterface, taskHandler)
	if err != nil {
		return nil, err
	}
	debugServer := app.NewDebugServer(configConfig, globalContext)
	application := app.NewApplication(configConfig, serverServer, metricsServer, workerWorker, debugServer)
	return application, nil
}
