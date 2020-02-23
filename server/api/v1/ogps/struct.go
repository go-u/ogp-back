package ogps

type OgpRequest struct {
	FQDN string `boil:"fqdn" json:"fqdn" toml:"fqdn" yaml:"fqdn"`
}
