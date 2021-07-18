package notion

import (
	"testing"

	"github.com/stretchr/testify/require"
)



func TestCreatePageParams_Validate(t *testing.T) {
	workspaceParams := CreatePageParams{ParentType: ParentTypeWorkspace}
	result := workspaceParams.Validate()
	require.NoError(t, result)
}
