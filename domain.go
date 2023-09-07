package publicsuffix

import (
	"strings"
)

type Domain struct {
	Tld string
	Sld string
	Trd string
}

func NameToLabels(name string) []string {
	return strings.Split(name, ".")
}

func NewDomain(args ...string) *Domain {
	d := &Domain{}
	if len(args) > 0 {
		d.Tld = args[0]
	}
	if len(args) > 1 {
		d.Sld = args[1]
	}
	if len(args) > 2 {
		d.Trd = args[2]
	}
	return d
}

func (d *Domain) String() string {
	return d.Name()
}

func (d *Domain) ToArray() []string {
	return []string{d.Trd, d.Sld, d.Tld}
}

func (d *Domain) Name() string {
	return strings.Join([]string{d.Trd, d.Sld, d.Tld}, ".")
}

func (d *Domain) Domain() string {
	if d.DomainCheck() {
		return strings.Join([]string{d.Sld, d.Tld}, ".")
	}
	return ""
}

func (d *Domain) Subdomain() string {
	if d.SubdomainCheck() {
		return strings.Join([]string{d.Trd, d.Sld, d.Tld}, ".")
	}
	return ""
}

func (d *Domain) DomainCheck() bool {
	return !(d.Tld == "" || d.Sld == "")
}

func (d *Domain) SubdomainCheck() bool {
	return !(d.Tld == "" || d.Sld == "" || d.Trd == "")
}
