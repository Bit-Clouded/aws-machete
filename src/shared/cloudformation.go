package shared

import (
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

type CfnInterface interface {
	DescribeStacks(input *cloudformation.DescribeStacksInput) (*cloudformation.DescribeStacksOutput, error)
	UpdateStack(input *cloudformation.UpdateStackInput) (*cloudformation.UpdateStackOutput, error)
	CreateChangeSet(input *cloudformation.CreateChangeSetInput) (*cloudformation.CreateChangeSetOutput, error)
	DescribeChangeSet(input *cloudformation.DescribeChangeSetInput) (*cloudformation.DescribeChangeSetOutput, error)
	GetTemplate(input *cloudformation.GetTemplateInput) (*cloudformation.GetTemplateOutput, error)
	ExecuteChangeSet(input *cloudformation.ExecuteChangeSetInput) (*cloudformation.ExecuteChangeSetOutput, error)
	GetTemplateSummary(input *cloudformation.GetTemplateSummaryInput) (*cloudformation.GetTemplateSummaryOutput, error)
	CancelUpdateStack(input *cloudformation.CancelUpdateStackInput) (*cloudformation.CancelUpdateStackOutput, error)
}
