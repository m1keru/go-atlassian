package internal

import (
	model "github.com/m1keru/go-atlassian/pkg/infra/models"
	"github.com/m1keru/go-atlassian/service"
)

// NewCommentService creates a new instance of CommentADFService and CommentRichTextService.
// It takes a service.Connector and a version string as input.
// Returns pointers to CommentADFService and CommentRichTextService, and an error if the version is not provided.
func NewCommentService(client service.Connector, version string) (*CommentADFService, *CommentRichTextService, error) {

	if version == "" {
		return nil, nil, model.ErrNoVersionProvided
	}

	adfService := &CommentADFService{
		internalClient: &internalAdfCommentImpl{
			c:       client,
			version: version,
		},
	}

	richTextService := &CommentRichTextService{
		internalClient: &internalRichTextCommentImpl{
			c:       client,
			version: version,
		},
	}

	return adfService, richTextService, nil
}
