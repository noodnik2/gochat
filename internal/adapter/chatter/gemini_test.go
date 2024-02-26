package chatter

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewGeminiChatter(t *testing.T) {
	cases := []struct {
		name      string
		cfg       Gemini
		assertion func(*testing.T, *GeminiChatter, error)
	}{
		{
			name: "no APIKEY",
			assertion: func(t *testing.T, chatter *GeminiChatter, err error) {
				t.Helper()
				require.Error(t, err)
				assert.Contains(t, err.Error(), "credentials")
				assert.Nil(t, chatter)
			},
		},
		{
			name: "have APIKEY",
			cfg:  Gemini{APIKey: "some-key", Model: "some-model"},
			assertion: func(t *testing.T, chatter *GeminiChatter, err error) {
				t.Helper()
				require.NoError(t, err)
				require.NotNil(t, chatter)
				assert.Equal(t, "some-model", chatter.Model())
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			chatter, err := NewGeminiChatter(context.Background(), tc.cfg)
			tc.assertion(t, chatter, err)
		})
	}
}
