package sm

import (
	"context"
	model "github.com/m1keru/go-atlassian/pkg/infra/models"
)

type InfoConnector interface {

	// Get retrieves information about the Jira Service Management instance such as software version, builds, and related links.
	//
	// GET /rest/servicedeskapi/info
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/info#get-info
	Get(ctx context.Context) (*model.InfoScheme, *model.ResponseScheme, error)
}
