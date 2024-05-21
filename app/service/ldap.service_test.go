package service

import (
	"hka-server-login/config"
	"hka-server-login/dto"
	"hka-server-login/helper"
	mock_ldap "hka-server-login/lib/ldap"
	"testing"

	"github.com/go-ldap/ldap/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSearchUser(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLdapClient := mock_ldap.NewMockClient(ctrl)

	// Define the expected LDAP search result
	ldapEntry := ldap.NewEntry("uid=testuser,ou=users,dc=example,dc=com", map[string][]string{
		"cn":        {"testuser"},
		"givenName": {"Test"},
		"sn":        {"User"},
		"mail":      {"testuser@example.com"},
	})
	searchResult := &ldap.SearchResult{
		Entries: []*ldap.Entry{ldapEntry},
	}

	// Create a config object
	cnf := &config.Config{
		LDAPServer: "ldap.example.com",
		LDAPPort:   389,
		UserDN:     "ou=users,dc=example,dc=com",
	}

	creds := dto.Credentials{
		Username: "testuser",
		Password: "testpw",
	}

	// Set up the mock expectation
	mockLdapClient.EXPECT().Bind(gomock.Eq(helper.FormatLdapUser("uid", creds.Username, cnf.UserDN)), gomock.Eq(creds.Password)).Return(nil)
	mockLdapClient.EXPECT().Search(gomock.Any()).Return(searchResult, nil)
	mockLdapClient.EXPECT().Unbind().Return(nil)

	// Create the service
	ldapService := NewLdapService(cnf, mockLdapClient)

	// Call the method to test
	user, err := ldapService.SearchUser(creds, creds.Username)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "Test", user.FirstName)
	assert.Equal(t, "User", user.LastName)
	assert.Equal(t, "testuser@example.com", user.Email)
}
