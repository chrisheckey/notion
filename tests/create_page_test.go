package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"

	"github.com/chrisheckey/notion"
)

func TestClient_CreatePage(t *testing.T) {

	ctx := context.Background()

	client := notion.NewClient(testToken(t), &notion.ClientOptions{HTTPClient: &http.Client{}})

	t.Run("create workspace page", func(t *testing.T) {

		opts := notion.CreatePageParams{
			ParentType:             notion.ParentTypeWorkspace,
			//ParentID:               "",
			//DatabasePageProperties: nil,
			Title:                  []notion.RichText{
				{
					Type:        notion.RichTextTypeText,
					//Annotations: nil,
					PlainText:   "Test API Title",
					//HRef:        nil,
					//Text:        nil,
					//Mention:     nil,
					//Equation:    nil,
				},
			},
			//Children:               nil,
			Workspace:              true,
		}

		page, err := client.CreatePage(ctx, opts)
		require.NoError(t, err)
		require.NotNil(t, page)

	})

}
