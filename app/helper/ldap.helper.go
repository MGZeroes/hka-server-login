package helper

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"reflect"
)

func FormatLdapUser(userKey, username, userDN string) string {
	return fmt.Sprintf("%s=%s,%s", userKey, username, userDN)
}

// GetLdapAttributes retrieves LDAP attribute names from the struct tags
func GetLdapAttributes(v interface{}) []string {
	val := reflect.ValueOf(v).Elem()
	typ := val.Type()
	var attributes []string
	for i := 0; i < val.NumField(); i++ {
		ldapTag := typ.Field(i).Tag.Get("ldap")
		if ldapTag != "" {
			attributes = append(attributes, ldapTag)
		}
	}
	return attributes
}

// SetLdapAttributes populates the struct fields with LDAP attribute values
func SetLdapAttributes(entry *ldap.Entry, v interface{}) error {
	val := reflect.ValueOf(v).Elem()
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		ldapTag := typ.Field(i).Tag.Get("ldap")
		if ldapTag != "" {
			value := entry.GetAttributeValue(ldapTag)
			field := val.Field(i)
			if field.CanSet() {
				field.SetString(value)
			}
		}
	}
	return nil
}
