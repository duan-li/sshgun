package bullet

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPaserNotWorking(t *testing.T) {
	ycontent := []byte(`raincoat: 1
coins: 5
books: 23
spectacles: 2
chairs: 12
pens: 6
`)
	y, err := Paser(ycontent)
	assert.NotNil(t, err)
	assert.Nil(t, y)
}

func TestPaser(t *testing.T) {
	yfile, err := os.ReadFile("../bullet_template.yaml")

	assert.Nil(t, err)
	assert.NotNil(t, yfile)

	y, err := Paser(yfile)
	assert.Nil(t, err)
	assert.NotNil(t, y)

	assert.Equal(t, 1, len(y))
	assert.Equal(t, 2, len(y["bullet"].Servers))
	assert.Equal(t, 2, len(y["bullet"].Commands))

	s1 := y["bullet"].Servers[0]
	assert.Equal(t, "127.0.0.1", s1.Ip)
	assert.Equal(t, 22, s1.Port)
	assert.Equal(t, "root", s1.Username)
	assert.Equal(t, "root", s1.Password)
	assert.Equal(t, "root", s1.Supassword)

	s2 := y["bullet"].Servers[1]
	assert.Equal(t, "127.0.0.2", s2.Ip)
	assert.Equal(t, 22, s2.Port)
	assert.Equal(t, "root", s2.Username)
	assert.Equal(t, "root", s2.Password)
	assert.Empty(t, s2.Supassword)

	assert.Equal(t, "echo \"Hello World\"", y["bullet"].Commands[0])
	assert.Equal(t, "ls -l ~", y["bullet"].Commands[1])

}
