package cmdservice

import (
	"fmt"
	"time"

	"github.com/anicoll/gosungrow/pkg/cmdlog"
	"github.com/anicoll/gosungrow/pkg/only"
	"github.com/kardianos/service"
)

func (c *Service) Control(action string) error {
	for range only.Once {

		// errs := make(chan error, 5)
		// logger, err = s.Logger(errs)
		// if err != nil {
		// 	break
		// }
		//
		// go func() {
		// 	for {
		// 		err := <-errs
		// 		if err != nil {
		// 			log.Print(err)
		// 		}
		// 	}
		// }()

		if action == "" {
			fmt.Printf("Valid actions: %q\n", service.ControlAction)
			break
		}

		c.Error = service.Control(c.service, action)
		if c.Error != nil {
			fmt.Printf("Valid actions: %q\n", service.ControlAction)
			break
		}
		fmt.Printf("Service action '%s' OK.\n", action)

		// err = s.Run()
		// if err != nil {
		// 	break
		// }
	}

	return c.Error
}

// func ServiceStart() error {
// 	var err error
//
// 	for range only.Once {
//
// 	}
//
// 	return err
// }
//
// func ServiceStop() error {
// 	var err error
//
// 	for range only.Once {
//
// 	}
//
// 	return err
// }
//
// func ServiceState() error {
// 	var err error
//
// 	for range only.Once {
//
// 	}
//
// 	return err
// }

func (p *program) Start(s service.Service) error {
	if service.Interactive() {
		cmdlog.Printf("Running in terminal.")
	} else {
		cmdlog.Printf("Running under service manager.")
	}

	p.exit = make(chan struct{})

	// Start should not block. Do the actual work async.
	go p.exec()
	return nil
}

func (p *program) exec() {
	cmdlog.Printf("I'm running %v.", service.Platform())

	cmdlog.Printf("app.StartApp()")
	time.Sleep(time.Second * 120)

	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case tm := <-ticker.C:
			cmdlog.Printf("Still running at %v...", tm)
		case <-p.exit:
			ticker.Stop()
			return
		}
	}
}

func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	cmdlog.Printf("I'm Stopping!")
	close(p.exit)
	return nil
}
