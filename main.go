package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
)

type config struct {
	port int
	host string
}

func main() {
	c, err := parseArgs(os.Stderr, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}

	err = runCmd(os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func parseArgs(w io.Writer, args []string) (*config, error) {
	c := config{}
	fs := flag.NewFlagSet("ccnmap", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.IntVar(&c.port, "port", 5000, "Port where ccnmap will scan")
	fs.StringVar(&c.host, "host", "localhost", "Host where ccnmap will scan")
	fs.Usage = func() {
		var usageString = `ccnmap is a port scanner that probes a host to identify open network ports.

		Usage of %s: <options> [value]`
		fmt.Fprintf(w, usageString, fs.Name())
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Options: ")
		fs.PrintDefaults()
	}
	if err := fs.Parse(args); err != nil {
		return &c, err
	}
	return &c, nil
}

func runCmd(w io.Writer, c *config) error {
	// Setup signal handling for graceful shutdown
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	go func() {
		<-signalCh
		// Initiate graceful shutdown
		fmt.Println("Bye bye ;-)")
		os.Exit(0)
	}()

	fmt.Printf("Scanning host: %s port: %d\n", c.host, c.port)
	return nil

}
