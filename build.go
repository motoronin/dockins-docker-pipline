package main

import "github.com/docker/engine-api/client"

type Build struct {
	Dockerfile  string
	ContextPath string
	Tags        []string
}

func (build Build) Run(docker *client.Client, p *Pipeline, s *Stage) error {
	/*
	   ctx := context.Background()
	   r := nil // io.Reader
	   docker.ImageBuild(ctx, r, types.ImageBuildOptions{
	       Dockerfile: build.Dockerfile,
	       Tags: build.Tags,
	   })
	*/
	return nil
}

func (cmd Build) String() string {
	return "Build"
}
