package jira

import (
	"context"
	model "github.com/m1keru/go-atlassian/pkg/infra/models"
)

type LabelConnector interface {

	// Gets returns a paginated list of labels.
	//
	// GET /rest/api/{2-3}/label
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/labels#get-all-labels
	Gets(ctx context.Context, startAt, maxResults int) (*model.IssueLabelsScheme, *model.ResponseScheme, error)
}
