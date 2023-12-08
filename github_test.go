package sdk

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const token = ""

func TestClient_StarredRepo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		t.Parallel()

		c := NewClient(token)

		resp, err := c.StarredRepo(
			WithPerPage(1),
		)

		require.Nil(t, err)
		require.Len(t, resp, 1)
	})

	t.Run("fail", func(t *testing.T) {
		t.Parallel()

		c := NewClient("")

		_, err := c.StarredRepo(
			WithPerPage(1),
		)

		require.NotNil(t, err)
	})
}
