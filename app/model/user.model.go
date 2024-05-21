package model

type User struct {
	Username    string `ldap:"cn"`
	LastName    string `ldap:"sn"`
	Course      string `ldap:"departmentNumber"`
	Description string `ldap:"description"`
	DisplayName string `ldap:"displayName"`
	Ldap        string `ldap:"distinguishedName"`
	FirstName   string `ldap:"givenName"`
	Email       string `ldap:"mail"`
	Faculty     string `ldap:"o"`
	Role        string `ldap:"title"`
	RoomNumber  string `ldap:"roomNumber"`
}
