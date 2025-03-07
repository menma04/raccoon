package app

import (
	"context"
	"raccoon/logger"
)

//Run the server
func Run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	//@TODO - init config

	//start server
	StartServer(ctx, cancel)
	logger.Info("App.Run --> Complete")
	<-ctx.Done()
	return nil
}
