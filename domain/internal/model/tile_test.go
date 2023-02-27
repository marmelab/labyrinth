package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreasure(t *testing.T) {
	t.Run("MarshalJSON", func(t *testing.T) {
		{
			bytes, err := json.Marshal(NoTreasure)
			assert.Nil(t, err)
			assert.Equal(t, `"·"`, string(bytes))
		}
		{
			treasure := Treasure('A')
			bytes, err := json.Marshal(treasure)
			assert.Nil(t, err)
			assert.Equal(t, `"A"`, string(bytes))
		}
		{
			treasure := Treasure('B')
			bytes, err := json.Marshal(treasure)
			assert.Nil(t, err)
			assert.Equal(t, `"B"`, string(bytes))
		}
	})

	t.Run("UnmarshalJSON", func(t *testing.T) {
		{
			bytes := []byte(`"·"`)

			var treasure Treasure
			err := json.Unmarshal(bytes, &treasure)
			assert.Nil(t, err)
			assert.Equal(t, NoTreasure, treasure)
		}
		{
			bytes := []byte(`"A"`)

			var treasure Treasure
			err := json.Unmarshal(bytes, &treasure)
			assert.Nil(t, err)
			assert.Equal(t, Treasure('A'), treasure)
		}
		{
			bytes := []byte(`"B"`)

			var treasure Treasure
			err := json.Unmarshal(bytes, &treasure)
			assert.Nil(t, err)
			assert.Equal(t, Treasure('B'), treasure)
		}
	})
}
