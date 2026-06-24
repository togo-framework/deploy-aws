// Package aws deploys a togo app to AWS (App Runner) by driving the `aws` CLI.
// Select with deploy.provider=aws; requires the aws CLI authenticated.
package aws

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/togo-framework/deploy"
	"github.com/togo-framework/togo"
)

func init() { deploy.RegisterDriver("aws", New) }

// New checks the aws CLI is present.
func New(_ *togo.Kernel) (deploy.Deployer, error) {
	if _, err := exec.LookPath("aws"); err != nil {
		return nil, fmt.Errorf("deploy-aws: the %q CLI is required (install + authenticate it)", "aws")
	}
	return &driver{}, nil
}

type driver struct{}

func run(ctx context.Context, name string, args ...string) (string, error) {
	out, err := exec.CommandContext(ctx, name, args...).CombinedOutput()
	return string(out), err
}

func region(s deploy.Spec, def string) string {
	if s.Region != "" {
		return s.Region
	}
	return def
}

func (d *driver) Provision(ctx context.Context, spec deploy.Spec) (*deploy.Result, error) {
	return d.Deploy(ctx, spec)
}

func (d *driver) Deploy(ctx context.Context, spec deploy.Spec) (*deploy.Result, error) {
	src := fmt.Sprintf("{\"ImageRepository\":{\"ImageIdentifier\":%q,\"ImageRepositoryType\":\"ECR_PUBLIC\",\"ImageConfiguration\":{\"Port\":\"8080\"}}}", spec.Image)
	out, err := run(ctx, "aws", "apprunner", "create-service", "--service-name", spec.App, "--source-configuration", src, "--region", region(spec, "us-east-1"))
	if err != nil {
		return nil, fmt.Errorf("aws apprunner create-service: %v: %s", err, out)
	}
	return &deploy.Result{Message: "App Runner service created", Raw: map[string]any{"out": out}}, nil
}

func (d *driver) Destroy(ctx context.Context, spec deploy.Spec) error {
	_, err := run(ctx, "aws", "apprunner", "delete-service", "--service-arn", spec.App, "--region", region(spec, "us-east-1"))
	return err
}

func (d *driver) Status(ctx context.Context, spec deploy.Spec) (*deploy.Status, error) {
	out, err := run(ctx, "aws", "apprunner", "list-services", "--region", region(spec, "us-east-1"))
	return &deploy.Status{Healthy: err == nil, Detail: out}, err
}
