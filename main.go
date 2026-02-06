package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/installable-sh/install/internal/install"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cmd := install.New(os.Args[1:])
	if err := cmd.Exec(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "[install] error: %v\n", err)
		os.Exit(1)
	}
}
