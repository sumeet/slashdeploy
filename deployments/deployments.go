package deployments

import "fmt"

type Deployment struct {
	// A unique identifier for the deployment request that was created.
	ID string
}

type DeploymentRequest struct {
	// The User or Organization that owns the repository.
	Owner string

	// The name of the repository to deploy.
	Repository string

	// The git ref to deploy.
	Ref string

	// The environment to deploy to.
	Environment string
}

func (d *DeploymentRequest) String() string {
	return fmt.Sprintf("%s/%s@%s to %s", d.Owner, d.Repository, d.Ref, d.Environment)
}

// Deployer represents something that can create deployment requests.
type Deployer interface {
	// When deploy is called, the implementer should create a deployment
	// request then immediately return. The implementer can use the provided
	// Statuses to notify the consumer about status updates.
	Deploy(DeploymentRequest) (*Deployment, error)
}

type DeployerFunc func(DeploymentRequest) (*Deployment, error)

func (fn DeployerFunc) Deploy(req DeploymentRequest) (*Deployment, error) {
	return fn(req)
}

// NullDeployer is a Deployer implementation that does nothing.
var NullDeployer = DeployerFunc(func(req DeploymentRequest) (*Deployment, error) {
	return &Deployment{ID: "1"}, nil
})