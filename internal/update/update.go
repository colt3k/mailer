package update

import (
	log "github.com/colt3k/nglog/ng"
	"github.com/colt3k/utils/netut/hc"
	"github.com/colt3k/utils/updater"
	"github.com/colt3k/utils/updater/artifactory"
	"github.com/colt3k/utils/version"
)

// CheckUpdate check for updates
func CheckUpdate(appName string) {

	v := updater.Version{
		Version:   version.VERSION,
		BuildDate: version.BUILDDATE,
	}

	var cons []updater.Connection

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
	for i,d := range cons {
		log.Logf(log.DEBUG,"%d checking %s for update",i ,d.Name)
	}
	artifactory.PerformUpdate(appName, cons, v, true)
}
