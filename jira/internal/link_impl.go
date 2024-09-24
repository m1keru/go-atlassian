package internal

import (
	model "github.com/m1keru/go-atlassian/pkg/infra/models"
	"github.com/m1keru/go-atlassian/service"
)

// NewLinkService creates new instances of LinkADFService and LinkRichTextService.
func NewLinkService(client service.Connector, version string, linkType *LinkTypeService, remote *RemoteLinkService) (*LinkADFService, *LinkRichTextService, error) {

	if version == "" {
		return nil, nil, model.ErrNoVersionProvided
	}

	adfService := &LinkADFService{
		internalClient: &internalLinkADFServiceImpl{
			c:       client,
			version: version,
		},
		Type:   linkType,
		Remote: remote,
	}

	richTextService := &LinkRichTextService{
		internalClient: &internalLinkRichTextServiceImpl{
			c:       client,
			version: version,
		},
		Type:   linkType,
		Remote: remote,
	}

	return adfService, richTextService, nil
}
