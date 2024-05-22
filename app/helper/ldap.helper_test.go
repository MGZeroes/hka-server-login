package helper

import (
	"github.com/go-ldap/ldap/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Username string `ldap:"uid"`
	Email    string `ldap:"mail"`
}

func TestFormatLdapUser(t *testing.T) {
	userKey := "uid"
	username := "jdoe"
	userDN := "ou=users,dc=example,dc=com"

	expected := "uid=jdoe,ou=users,dc=example,dc=com"
	result := FormatLdapUser(userKey, username, userDN)

	assert.Equal(t, expected, result, "they should be equal")
}

func TestGetLdapAttributes(t *testing.T) {
	user := &User{}
	expected := []string{"uid", "mail"}

	result := GetLdapAttributes(user)

	assert.ElementsMatch(t, expected, result, "they should contain the same elements")
}

func TestSetLdapAttributes(t *testing.T) {
	entry := &ldap.Entry{
		DN: "uid=jdoe,ou=users,dc=example,dc=com",
		Attributes: []*ldap.EntryAttribute{
			{Name: "uid", Values: []string{"jdoe"}},
			{Name: "mail", Values: []string{"jdoe@example.com"}},
		},
	}

	user := &User{}

	err := SetLdapAttributes(entry, user)
	assert.NoError(t, err, "should not return an error")
	assert.Equal(t, "jdoe", user.Username, "Username should be set correctly")
	assert.Equal(t, "jdoe@example.com", user.Email, "Email should be set correctly")
}
