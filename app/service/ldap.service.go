package service

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"hka-server-login/config"
	"hka-server-login/dto"
	"hka-server-login/helper"
	"hka-server-login/model"
)

type LdapService struct {
	cfg        config.LdapConfig
	ldapClient ldap.Client
}

type LdapServiceImpl interface {
	SearchUser(dto.Credentials, string) (*model.User, error)
}

func NewLdapService(cfg config.LdapConfig, l ldap.Client) LdapServiceImpl {
	return &LdapService{cfg: cfg, ldapClient: l}
}

func (service *LdapService) validateCredentials(creds dto.Credentials) error {

	if len(creds.Username) == 0 {
		return fmt.Errorf("username is empty")
	}

	if len(creds.Password) == 0 {
		return fmt.Errorf("password is empty")
	}

	return nil
}

func (service *LdapService) bind(creds dto.Credentials) error {
	err := service.validateCredentials(creds)
	if err != nil {
		return err
	}

	err = service.ldapClient.Bind(helper.FormatLdapUser(service.cfg.UserKey, creds.Username, service.cfg.UserDN), creds.Password)
	if err != nil {
		return fmt.Errorf("invalid username or password")
	}

	return nil
}

func (service *LdapService) unbind() error {
	err := service.ldapClient.Unbind()
	if err != nil {
		return fmt.Errorf("refused to unbind")
	}

	return nil
}

// SearchUser searches for a single user
func (service *LdapService) SearchUser(creds dto.Credentials, username string) (*model.User, error) {
	user := &model.User{}

	err := service.bind(creds)
	if err != nil {
		return nil, err
	}
	defer service.unbind()

	attributes := helper.GetLdapAttributes(user)
	searchRequest := ldap.NewSearchRequest(
		service.cfg.UserDN,
		ldap.ScopeSingleLevel, ldap.NeverDerefAliases, 1, 0, false,
		fmt.Sprintf("(uid=%s)", username),
		attributes,
		nil,
	)

	sr, err := service.ldapClient.Search(searchRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to search for user: %v", err)
	}

	if len(sr.Entries) == 0 {
		return nil, fmt.Errorf("user does not exist")
	}

	if len(sr.Entries) > 1 {
		return nil, fmt.Errorf("multiple users found")
	}

	helper.SetLdapAttributes(sr.Entries[0], user)

	return user, nil
}
