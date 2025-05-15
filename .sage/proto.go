package main

import (
	"context"

	"go.einride.tech/sage/sg"
	"go.einride.tech/sage/tools/sgbuf"
)

type Proto sg.Namespace

func (Proto) Default(ctx context.Context) error {
	sg.Deps(ctx)
	return nil
}

func (Proto) BufModInit(ctx context.Context) error {
	sg.Logger(ctx).Println("initializing buf module...")
	cmd := sgbuf.Command(ctx, "mod", "init")
	cmd.Dir = sg.FromGitRoot("proto")
	return cmd.Run()
}

func (Proto) BufModUpdate(ctx context.Context) error {
	sg.Logger(ctx).Println("updating buf module...")
	cmd := sgbuf.Command(ctx, "mod", "update")
	cmd.Dir = sg.FromGitRoot("proto")
	return cmd.Run()
}
