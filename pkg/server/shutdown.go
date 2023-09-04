package server

import (
	"os"
	"os/signal"
	"syscall"
)

// Hook a graceful shutdown hook, default with signals of SIGINT and SIGTERM
type Hook interface {
	// WithSignals add more signals into hook
	WithSignals(signals ...syscall.Signal) Hook

	// Close register shutdown handles
	Close(funcs ...func(sg os.Signal))
}

type hook struct {
	ctx chan os.Signal
}

// NewHook create a Hook instance
func NewHook() Hook {
	hook := &hook{
		ctx: make(chan os.Signal, 1),
	}

	return hook.WithSignals(syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGHUP)
}

func (h *hook) WithSignals(signals ...syscall.Signal) Hook {
	for _, s := range signals {
		signal.Notify(h.ctx, s)
	}

	return h
}

func (h *hook) Close(funcs ...func(sg os.Signal)) {
	select {
	case sg := <-h.ctx:
		signal.Stop(h.ctx)
		close(h.ctx)

		for _, f := range funcs {
			f(sg)
		}
	}

}
