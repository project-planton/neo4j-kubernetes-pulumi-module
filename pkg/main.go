package pkg

import (
	neo4jkubernetesv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/kubernetes/neo4jkubernetes/v1"
	"github.com/pkg/errors"
	"github.com/project-planton/pulumi-module-golang-commons/pkg/provider/kubernetes/pulumikubernetesprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *neo4jkubernetesv1.Neo4JKubernetesStackInput) error {
	//create kubernetes-provider from the credential in the stack-input
	_, err := pulumikubernetesprovider.GetWithKubernetesClusterCredential(ctx,
		stackInput.KubernetesCluster, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to setup gcp provider")
	}

	return nil
}
