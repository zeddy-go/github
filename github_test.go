package sdk

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

const token = ""

func TestClient_StarredRepo(t *testing.T) {
	c := NewClient(token)

	resp, err := c.StarredRepo(
		WithPerPage(1),
	)

	require.Nil(t, err)
	fmt.Printf("%+v\n", resp)
}
