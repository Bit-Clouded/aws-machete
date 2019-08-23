package cmd

import (
	"aws-machete/src/shared"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

type cfnClient struct {
	cfn shared.CfnInterface
}

func newCfnClient() *cfnClient {
	var sess = session.Must(session.NewSession())

	return &cfnClient{
		cfn: cloudformation.New(sess),
	}
}

var cfnClientInstance *cfnClient = newCfnClient()

func (client *cfnClient) update(
	stackName *string,
	params []*cloudformation.Parameter,
	tags []*cloudformation.Tag,
	templateBody *string) (*cloudformation.UpdateStackOutput, error) {

	capability := cloudformation.CapabilityCapabilityIam
	usePreviousTemplate := templateBody == nil

	csInput := &cloudformation.UpdateStackInput{
		StackName:           stackName,
		Capabilities:        []*string{&capability},
		Parameters:          params,
		Tags:                tags,
		TemplateBody:        templateBody,
		UsePreviousTemplate: &usePreviousTemplate,
	}
	return client.cfn.UpdateStack(csInput)
}
