package atccmd

import (
	"os"

	"github.com/concourse/atc/builds"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc/dbng"
)

type drainer struct {
	logger  lager.Logger
	drain   chan<- struct{}
	tracker builds.BuildTracker
	bus     dbng.NotificationsBus
}

func (d drainer) Run(signals <-chan os.Signal, ready chan<- struct{}) error {
	close(ready)

	<-signals

	d.logger.Info("releasing-tracker")
	d.tracker.Release()
	d.logger.Info("released-tracker")

	close(d.drain)
	d.logger.Info("sending-atc-shutdown-message")
	d.bus.Notify("atc_shutdown")

	return nil
}
