package main

import (
	"context"

	"go.einride.tech/sage/sg"
	"go.einride.tech/sage/sgtool"
	"go.einride.tech/sage/tools/sgconvco"
	"go.einride.tech/sage/tools/sggit"
	"go.einride.tech/sage/tools/sggo"
	"go.einride.tech/sage/tools/sggolangcilint"
	"go.einride.tech/sage/tools/sggolines"
	"go.einride.tech/sage/tools/sgmdformat"
	"go.einride.tech/sage/tools/sgyamlfmt"
)

func main() {
	sg.GenerateMakefiles(
		sg.Makefile{
			Path:          sg.FromGitRoot("Makefile"),
			DefaultTarget: Default,
		},
		sg.Makefile{
			Path:          sg.FromGitRoot("proto", "Makefile"),
			DefaultTarget: Proto.Default,
			Namespace:     Proto{},
		},
	)
}

func Default(ctx context.Context) error {
	sg.Deps(ctx, ConvcoCheck)
	sg.Deps(ctx, FormatMarkdown, FormatYaml)
	sg.Deps(ctx, Proto.Default)
	sg.Deps(ctx, WireGenerate)
	sg.Deps(ctx, SpannerGenerate, GoGenerate)
	sg.Deps(ctx, GoLint)
	sg.Deps(ctx, GoTest)
	sg.Deps(ctx, GoModTidy)
	sg.Deps(ctx, GitVerifyNoDiff)
	return nil
}

func Run(ctx context.Context) error {
	cmd := sg.Command(ctx, "go", "run", "./cmd/server", "-config", "config/local.yaml")
	cmd.Dir = sg.FromGitRoot()
	return cmd.Run()
}

func GoGenerate(ctx context.Context) error {
	sg.Logger(ctx).Println("generating Go code...")
	cmd := sg.Command(ctx, "go", "generate", "./internal/app/server/")
	cmd.Dir = sg.FromGitRoot()
	return cmd.Run()
}

func SpannerGenerate(ctx context.Context) error {
	sg.Logger(ctx).Println("generating Spanner code...")
	cmd := sg.Command(ctx, "go", "run", "go.einride.tech/spanner-aip", "generate")
	cmd.Dir = sg.FromGitRoot()
	return cmd.Run()
}

func WireGenerate(ctx context.Context) error {
	sg.Logger(ctx).Println("generating wire code...")
	lnk, err := sgtool.GoInstallWithModfile(ctx, "github.com/google/wire/cmd/wire", sg.FromGitRoot(".", "go.mod"))
	if err != nil {
		return err
	}
	cmd := sg.Command(ctx, lnk, "gen", "./internal/app/server")
	cmd.Dir = sg.FromGitRoot(".")
	return cmd.Run()
}

func GoModTidy(ctx context.Context) error {
	sg.Logger(ctx).Println("tidying Go module files...")
	return sg.Command(ctx, "go", "mod", "tidy", "-v").Run()
}

func GoTest(ctx context.Context) error {
	sg.Logger(ctx).Println("running Go tests...")
	return sggo.TestCommand(ctx).Run()
}

func GoLint(ctx context.Context) error {
	sg.Logger(ctx).Println("linting Go files...")
	return sggolangcilint.Run(ctx)
}

func GoLintFix(ctx context.Context) error {
	sg.Logger(ctx).Println("fixing Go files...")
	return sggolangcilint.Fix(ctx)
}

func GoFormat(ctx context.Context) error {
	sg.Logger(ctx).Println("formatting Go files...")
	return sggolines.Run(ctx)
}

func FormatMarkdown(ctx context.Context) error {
	sg.Logger(ctx).Println("formatting Markdown files...")
	return sgmdformat.Command(ctx).Run()
}

func FormatYaml(ctx context.Context) error {
	sg.Logger(ctx).Println("formatting Yaml files...")
	return sgyamlfmt.Run(ctx)
}

func ConvcoCheck(ctx context.Context) error {
	sg.Logger(ctx).Println("checking git commits...")
	return sgconvco.Command(ctx, "check", "origin/main..HEAD").Run()
}

func GitVerifyNoDiff(ctx context.Context) error {
	sg.Logger(ctx).Println("verifying that git has no diff...")
	return sggit.VerifyNoDiff(ctx)
}
