package config

type Config struct {
	Server ServerConfig `json:"server"`
	LDAP   LdapConfig   `json:"ldap"`
}

type ServerConfig struct {
	Addr      string    `json:"addr"`
	EnableTLS bool      `json:"enableTLS"`
	TLS       TLSConfig `json:"tls"`
}

type TLSConfig struct {
	Certificate string `json:"certificate"`
	Key         string `json:"key"`
}

type LdapConfig struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	BaseDN  string `json:"baseDN"`
	UserDN  string `json:"userDN"`
	UserKey string `json:"userKey"`
}
