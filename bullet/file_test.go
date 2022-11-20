package bullet

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	content := []byte("hello world")
	result, err := WriteBullet("test.yml", content)
	assert.True(t, result)
	assert.Nil(t, err)
	data, err1 := ReadBullet("test.yml")
	assert.Nil(t, err1)
	assert.NotNil(t, data)
	assert.Equal(t, content, data)
	os.Remove("test.yml")
}
