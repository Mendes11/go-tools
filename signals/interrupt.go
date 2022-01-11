package signals

import (
	"log"
	"os"
	"os/signal"
)

// CallOnInterrupt will call a function whenever
// we receive a System Interrupt signal
func CallOnInterrupt(callback func()) {
	CallOnSignals(callback, os.Interrupt)
}

// CallOnSignals registers a function that will be called
// whenever one of the signals is triggered
func CallOnSignals(callback func(), signals ...os.Signal) {
	sigCh := make(chan os.Signal, len(signals))
	signal.Notify(sigCh, signals...)

	go callOnSignal(sigCh, callback)
}

func callOnSignal(sigCh chan os.Signal, callback func()) {
	sig := <-sigCh
	log.Printf("%T received. Calling Cancel Function", sig)
	callback()
}
