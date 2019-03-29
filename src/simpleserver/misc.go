package simpleserver

import (
	"jkmisc"
)

var globalWether *jkmisc.JKWether

func GlobalWetherInit() error {
	key := GlobalBaseConfig().WetherKey
	globalWether, _ = jkmisc.JKWetherNew(key)
	return nil
}

func GlobalWether() *jkmisc.JKWether {
	return globalWether
}
