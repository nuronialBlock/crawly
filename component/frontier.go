package component

import (
	"os"
	"strconv"

	"github.com/nuronialBlock/crawly/storage"
	log "github.com/sirupsen/logrus"
)

type Frontier struct {
	Politeness      int
	Delay           int
	FrontierStorage *storage.FrontierStorage
}

func NewFrontier() *Frontier {
	// TODO: read the GetENV values from config.

	// set default value.
	fs := storage.NewFrontierStorage()

	f := &Frontier{
		Politeness:      1,
		Delay:           60,
		FrontierStorage: fs,
	}

	p, err := strconv.Atoi(os.Getenv("politeness"))
	if err != nil {
		log.Warn("Couldn't read politeness, falling back to default value: ", f.Politeness)
	} else {
		f.Politeness = p
		log.Info("Set Politeness to ", f.Politeness)
	}

	d, err := strconv.Atoi(os.Getenv("delay"))
	if err != nil {
		log.Warn("Couldn't read delay, falling back to default value: ", f.Delay)
	} else {
		f.Delay = d
		log.Info("Set Delay to ", f.Delay)
	}

	return f
}
