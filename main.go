package main

import (
	"context"
	"os"

	"github.com/galayx-future/costpilot/internal/domain"

	"github.com/galayx-future/costpilot/internal/config"
)

func main() {
	ctx := context.Background()
	printVersion()
	if err := config.Init(); err != nil {
		os.Exit(1)
	}

	a := domain.NewCostAnalysisDomain()
	if err := a.RunPipeline(ctx); err != nil {
		os.Exit(1)
	}

	if err := output(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
