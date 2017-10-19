package main

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type check struct {
	Name         string
	ExecLock     *sync.Mutex
	Instructions *instructions
	LastChecked  time.Time
	LastCritical time.Time
	Results      string
	Error        error
	Attempts     int
	Status       string
}

func (c *check) every() time.Duration {
	return c.Instructions.every()
}

func (c *check) id() string {
	return c.Instructions.id()
}

func (c *check) check() {
	go func() {
		// lock it all up ... at least so we don't run the same check twice
		c.ExecLock.Lock()
		defer c.ExecLock.Unlock()

		// make sure we update the last checked date
		c.LastChecked = time.Now()

		// setup our logger
		cLog := log.WithFields(log.Fields{"check": c.Name, "attempt": c.Attempts})

		// was it ok?
		if c.Instructions.Check.exec().ok() {
			// were we in critical mode?
			if c.Attempts >= c.Instructions.Try() && c.Instructions.Try() != 0 {
				// yes ...lets recover
				c.Instructions.Recover.exec()
				cLog.Info("Recovering")
			}

			// reset ...
			c.Attempts = 0
			c.Instructions.OK.exec()
			cLog.Info("Ok")
			c.Status = "OK"
			return
		}

		// increase our attempts
		c.Attempts++

		if c.Attempts <= c.Instructions.Try() {
			// where we at in regards to attempts vs tries
			if c.Attempts >= c.Instructions.Try() {
				// we need to error out :(
				cLog.Error("Check failed. Upgrading status to critical")
				c.Instructions.Critical.exec()
				c.LastCritical = time.Now()
				c.Status = "Critical"
			} else {
				// not yet critical
				cLog.Warn("Check failed")
				c.Instructions.Warning.exec()
				c.Status = "Warning"
			}

			// ok, so we failed ... but not quite at error levels
			if c.Attempts >= c.Instructions.Fix.After && c.Attempts <= c.Instructions.Try() {
				if c.Instructions.Fix.Cmd != "" {
					// regardless, try to fix ...
					c.Instructions.Fix.exec()
					cLog.Info("Attempting fix")
				}
			}
		} else {
			// Ok, so we are over our check attempts ....
			if c.LastCritical.Add(c.Instructions.Reset).Before(time.Now()) && c.Instructions.Try() != 0 {
				c.Attempts = 0
				cLog.Info("Reset check")
			}
		}
	}()
}
