package main

import (
	"github.com/pkg/errors"
	"github.com/project-planton/neo4j-kubernetes-pulumi-module/pkg"
	neo4jkubernetesv1 "github.com/project-planton/project-planton/apis/go/project/planton/provider/kubernetes/neo4jkubernetes/v1"
	"github.com/project-planton/project-planton/pkg/pulmod/stackinput"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		stackInput := &neo4jkubernetesv1.Neo4JKubernetesStackInput{}

		if err := stackinput.LoadStackInput(ctx, stackInput); err != nil {
			return errors.Wrap(err, "failed to load stack-input")
		}

		return pkg.Resources(ctx, stackInput)
	})
}
