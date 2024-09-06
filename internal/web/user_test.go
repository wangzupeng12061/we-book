package web

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	t.Log("TestEncrypt")
	password := "123456"
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	err = bcrypt.CompareHashAndPassword(fromPassword, []byte(password))
	assert.NoError(t, err)
}
func TestNil(t *testing.T) {
	testTypeAssert(nil)
}
func testTypeAssert(c any) {
	claims := c.(*UserClaims)
	fmt.Println(claims.Uid)
}
