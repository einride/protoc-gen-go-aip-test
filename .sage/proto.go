package main

import (
	"context"

	"go.einride.tech/sage/sg"
	"go.einride.tech/sage/sgtool"
	"go.einride.tech/sage/tools/sgapilinter"
	"go.einride.tech/sage/tools/sgbuf"
)

type Proto sg.Namespace

const (
	// Version of github.com/googleapis/googleapis to generate from.
	// Since there are no tags or branches to generate from in that repository,
	// this will fail once more than 1_000 commits have been added in that repository
	// since this commit.
	googleapisRef = "8d73440c77f9c1a877a31b2d617c5ea07e488892"
)

func (Proto) All(ctx context.Context) error {
	sg.SerialDeps(ctx, Proto.BufFormat, Proto.BufLint, Proto.APILinterLint)
	sg.SerialDeps(ctx, Proto.BufGenerate, Proto.BufGenerateGoogleapis)
	return nil
}

func (Proto) BufLint(ctx context.Context) error {
	sg.Logger(ctx).Println("linting proto files...")
	cmd := sgbuf.Command(ctx, "lint")
	cmd.Dir = sg.FromGitRoot("proto")
	return cmd.Run()
}

func (Proto) APILinterLint(ctx context.Context) error {
	sg.Logger(ctx).Println("linting gRPC APIs...")
	return sgapilinter.Run(ctx)
}

func (Proto) BufFormat(ctx context.Context) error {
	sg.Logger(ctx).Println("formatting proto files...")
	cmd := sgbuf.Command(ctx, "format", "--write")
	cmd.Dir = sg.FromGitRoot("proto")
	return cmd.Run()
}

func (Proto) ProtocGenGo(ctx context.Context) error {
	sg.Logger(ctx).Println("installing...")
	_, err := sgtool.GoInstallWithModfile(ctx, "google.golang.org/protobuf/cmd/protoc-gen-go", "go.mod")
	return err
}

func (Proto) ProtocGenGoGRPC(ctx context.Context) error {
	sg.Logger(ctx).Println("installing...")
	_, err := sgtool.GoInstall(ctx, "google.golang.org/grpc/cmd/protoc-gen-go-grpc", "v1.2.0")
	return err
}

func (Proto) ProtocGenGoAIPTest(ctx context.Context) error {
	sg.Logger(ctx).Println("building binary...")
	return sg.Command(
		ctx,
		"go",
		"build",
		"-o",
		sg.FromBinDir("protoc-gen-go-aip-test"),
		"./cmd/protoc-gen-go-aip-test",
	).Run()
}

func (Proto) BufGenerate(ctx context.Context) error {
	sg.Deps(ctx, Proto.ProtocGenGo, Proto.ProtocGenGoGRPC, Proto.ProtocGenGoAIPTest)
	sg.Logger(ctx).Println("generating proto stubs...")
	cmd := sgbuf.Command(ctx, "generate", "--template", "buf.gen.yaml", "--path", "einride")
	cmd.Dir = sg.FromGitRoot("proto")
	return cmd.Run()
}

func (Proto) BufGenerateGoogleapis(ctx context.Context) error {
	sg.Deps(ctx, Proto.ProtocGenGo, Proto.ProtocGenGoGRPC, Proto.ProtocGenGoAIPTest)
	sg.Logger(ctx).Println("generating example proto stubs...")
	cmd := sgbuf.Command(
		ctx,
		"generate",
		"https://github.com/googleapis/googleapis.git#depth=1000,ref="+googleapisRef,
		"--template", "buf.gen.googleapis.yaml",
		"--path", "google/area120/tables/v1alpha1",
		"--path", "google/cloud/aiplatform/v1",
		"--path", "google/cloud/gsuiteaddons/v1",
		"--path", "google/cloud/scheduler/v1",
		"--path", "google/pubsub/v1",
		"--path", "google/spanner",
	)
	cmd.Dir = sg.FromGitRoot("proto")
	return cmd.Run()
}
