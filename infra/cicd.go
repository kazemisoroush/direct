package main

// This file defines the CI/CD trust that lets GitHub Actions deploy via OIDC.

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdklabs/cdk-nag-go/cdknag/v2"
)

// gitHubAudience is the audience GitHub sets when requesting AWS credentials.
const gitHubAudience = "sts.amazonaws.com"

// deploySubject pins the trust to a push on the main branch of the direct repo.
const deploySubject = "repo:kazemisoroush/direct:ref:refs/heads/main"

// gitHubOIDCProviderArn is the account-level GitHub Actions OIDC provider. AWS allows only
// one provider per issuer URL per account, so Direct references the shared provider (created
// by the Vault CI/CD stack) rather than creating a second one.
func gitHubOIDCProviderArn(account string) string {
	return "arn:aws:iam::" + account + ":oidc-provider/token.actions.githubusercontent.com"
}

// NewDirectCICDStack defines the GitHub Actions deploy role trusting the shared OIDC provider.
func NewDirectCICDStack(scope constructs.Construct, id string, props *awscdk.StackProps) awscdk.Stack {
	stack := awscdk.NewStack(scope, &id, props)

	principal := awsiam.NewFederatedPrincipal(
		jsii.String(gitHubOIDCProviderArn(*stack.Account())),
		&map[string]any{
			"StringEquals": map[string]any{
				"token.actions.githubusercontent.com:aud": gitHubAudience,
				"token.actions.githubusercontent.com:sub": deploySubject,
			},
		},
		jsii.String("sts:AssumeRoleWithWebIdentity"),
	)

	role := awsiam.NewRole(stack, jsii.String("GithubActionsDeploy"), &awsiam.RoleProps{
		RoleName:    jsii.String("direct-github-actions-deploy"),
		AssumedBy:   principal,
		Description: jsii.String("GitHub Actions assumes this via OIDC to deploy DirectStack."),
	})

	bootstrapRoles := "arn:aws:iam::" + *stack.Account() + ":role/cdk-hnb659fds-*"
	role.AddToPolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Actions:   jsii.Strings("sts:AssumeRole"),
		Resources: jsii.Strings(bootstrapRoles),
	}))

	cdknag.NagSuppressions_AddResourceSuppressions(role, &[]*cdknag.NagPackSuppression{
		{
			Id:     jsii.String("AwsSolutions-IAM5"),
			Reason: jsii.String("The deploy role may only assume the CDK bootstrap roles, which share the cdk-hnb659fds-* name prefix; the wildcard is scoped to those roles in this account."),
		},
	}, jsii.Bool(true))

	awscdk.NewCfnOutput(stack, jsii.String("DeployRoleArn"), &awscdk.CfnOutputProps{Value: role.RoleArn()})

	return stack
}
