# deploy-aws — docs

**AWS deploy.** Deploy to AWS App Runner by driving the authenticated `aws` CLI.

## Install

```bash
togo install togo-framework/deploy-aws
```

Registers on the [`deploy`](https://github.com/togo-framework/deploy) base; select it with **deploy.provider in togo.yaml (or DEPLOY_PROVIDER)**, then use **`togo deploy`**.

## Interface

`Deployer` — `Provision`/`Deploy`/`Destroy`/`Status` over a `Spec{App,Dir,BuildCmd,Host,User,Image,Region,Domain}` built from your `togo.yaml`.

## Usage & notes

Requires the `aws` CLI installed + authenticated. Deploys `spec.Image` to App Runner on port 8080; `region` from `togo.yaml` `deploy.region`.

## Example

```bash
togo deploy --provider aws --dry-run   # preview the plan
togo deploy --provider aws
```

## Links

- [AWS App Runner](https://docs.aws.amazon.com/apprunner/)
- [Marketplace](https://to-go.dev/marketplace)
- [Source](https://github.com/togo-framework/deploy-aws)
