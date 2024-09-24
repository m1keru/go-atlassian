package jira

import (
	"context"
	model "github.com/m1keru/go-atlassian/pkg/infra/models"
)

type MySelfConnector interface {

	// Details returns details for the current user.
	//
	// GET /rest/api/{2-3}/myself
	//
	// https://docs.go-atlassian.io/jira-software-cloud/myself#get-current-user
	Details(ctx context.Context, expand []string) (*model.UserScheme, *model.ResponseScheme, error)
}
