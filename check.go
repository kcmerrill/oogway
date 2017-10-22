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

func (c *check) isMuted() bool {
	return c.Instructions.Muted
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
		if c.Instructions.Check.exec("check", c).ok() {
			// were we in critical mode?
			if c.Attempts >= c.Instructions.Try() && c.Instructions.Try() != 0 {
				// yes ...lets recove
				cLog.Info("Recovering")
				if !c.Instructions.Recover.exec("recover", c).ok() {
					cLog.WithFields(log.Fields{"instruction": "recover"}).Error(c.Instructions.Recover.Error.Error())
				}
			}

			// reset ...
			c.Attempts = 0
			c.Status = "OK"
			cLog.Info("Ok")
			if !c.Instructions.OK.exec("ok", c).ok() {
				cLog.WithFields(log.Fields{"instruction": "ok"}).Error(c.Instructions.OK.Error.Error())
			}
			return
		}

		// increase our attempts
		c.Attempts++

		if c.Attempts <= c.Instructions.Try() {
			// where we at in regards to attempts vs tries
			if c.Attempts >= c.Instructions.Try() {
				// we need to error out :(
				cLog.Error("Check failed. Upgrading status to critical")
				c.Status = "Critical"
				c.LastCritical = time.Now()
				if !c.Instructions.Critical.exec("critical", c).ok() {
					cLog.WithFields(log.Fields{"instruction": "critical"}).Error(c.Instructions.Critical.Error.Error())
				}
			} else {
				// not yet critical
				cLog.Warn("Check failed")
				c.Status = "Warning"
				if !c.Instructions.Warning.exec("warn", c).ok() {
					cLog.WithFields(log.Fields{"instruction": "warning"}).Error(c.Instructions.Warning.Error.Error())
				}
			}

			// ok, so we failed ... but not quite at error levels
			if c.Attempts >= c.Instructions.Fix.After && c.Attempts <= c.Instructions.Try() {
				if c.Instructions.Fix.okToExec() {
					// regardless, try to fix ...
					cLog.Info("Attempting fix")
					if !c.Instructions.Fix.exec("fix", c).ok() {
						cLog.WithFields(log.Fields{"instruction": "fix"}).Error(c.Instructions.Fix.Error.Error())
					}
				}
			}
		} else {
			// Ok, so we are over our check attempts ....
			if c.LastCritical.Add(c.Instructions.Reset).Before(time.Now()) && c.Instructions.Try() != 0 && c.Instructions.Reset != 0*time.Second {
				c.Attempts = 0
				cLog.Info("Reset check")
			}
		}
	}()
}
