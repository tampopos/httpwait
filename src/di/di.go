package di

import (
	"github.com/tampopos/httpwait/src/client"
	"github.com/tampopos/httpwait/src/httpwait"
	"github.com/tampopos/httpwait/src/stopwatch"
	"go.uber.org/dig"
)

// CreateContainer はDIContainerを生成します
func CreateContainer() (*dig.Container, error) {
	container := dig.New()
	if err := container.Provide(client.Create); err != nil {
		return nil, err
	}
	if err := container.Provide(stopwatch.New); err != nil {
		return nil, err
	}
	if err := container.Provide(httpwait.CreateUseCase); err != nil {
		return nil, err
	}
	return container, nil
}
