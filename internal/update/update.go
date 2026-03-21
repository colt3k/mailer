package update

import (
	log "github.com/colt3k/nglog/ng"
	"github.com/colt3k/utils/netut/hc"
	"github.com/colt3k/utils/updater"
	"github.com/colt3k/utils/updater/artifactory"
	"github.com/colt3k/utils/version"
)

// CheckUpdate checks the configured artifactory source and asks the shared updater to apply a newer release.
func CheckUpdate(appName string) {

	// Build metadata is passed to the updater so it can compare the running binary against remote artifacts.
	v := updater.Version{
		Version:   version.VERSION,
		BuildDate: version.BUILDDATE,
	}

	// If the endpoint probe fails, the connection list stays empty and the updater becomes a no-op.
	var cons []updater.Connection

	// Only register an update source when the artifactory base URL is reachable.
	if b, err := hc.Reachable("http://domain:8081", "main", 2, false); b && err == nil {
		c := updater.Connection{
			Name:               "main",
			User:               "ronly",
			PassOrToken:        "",
			URLPrefix:          "http://domain:8081/artifactory/",
			Repository:         "go-release-local/",
			Path:               "mailer/",
			OnAvailable:        "http://domain:8081",
			OnAvailableViaHTTP: true,
		}

		cons = []updater.Connection{c}
	}

	log.Logln(log.INFO, "checking for update")
	for i, d := range cons {
		log.Logf(log.DEBUG, "%d checking %s for update", i, d.Name)
	}
	artifactory.PerformUpdate(appName, cons, v, true)
}
