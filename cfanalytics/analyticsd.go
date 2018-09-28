package cfanalytics

import (
	"code.cloudfoundry.org/cfdev/config"
	"code.cloudfoundry.org/cfdev/daemon"
)

const AnalyticsDLabel = "org.cloudfoundry.cfdev.cfanalyticsd"

type AnalyticsD struct {
	Config       config.Config
	DaemonRunner DaemonRunner
}

type DaemonRunner interface {
	AddDaemon(daemon.DaemonSpec) error
	RemoveDaemon(string) error
	Start(string) error
	Stop(string) error
	IsRunning(string) (bool, error)
}

func (a *AnalyticsD) Start() error {
	spec := a.DaemonSpec()

	err := a.DaemonRunner.AddDaemon(spec)
	if err != nil {
		return err
	}

	return a.DaemonRunner.Start(AnalyticsDLabel)
}

func (a *AnalyticsD) Stop() error {
	var reterr error
	if err := a.DaemonRunner.Stop(AnalyticsDLabel); err != nil {
		reterr = err
	}
	return reterr
}

func (a *AnalyticsD) Destroy() error {
	return a.DaemonRunner.RemoveDaemon(AnalyticsDLabel)
}

func (a *AnalyticsD) IsRunning() (bool, error) {
	return a.DaemonRunner.IsRunning(AnalyticsDLabel)
}
