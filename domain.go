package publicsuffix

import (
	"strings"
)

type Domain struct {
	tld string
	sld string
	trd string
}

func NameToLabels(name string) []string {
	return strings.Split(name, ".")
}

func NewDomain(args ...string) *Domain {
	d := &Domain{}
	if len(args) > 0 {
		d.tld = args[0]
	}
	if len(args) > 1 {
		d.sld = args[1]
	}
	if len(args) > 2 {
		d.trd = args[2]
	}
	return d
}

func (d *Domain) String() string {
	return d.Name()
}

func (d *Domain) ToArray() []string {
	return []string{d.trd, d.sld, d.tld}
}

func (d *Domain) Name() string {
	return strings.Join([]string{d.trd, d.sld, d.tld}, ".")
}

func (d *Domain) Domain() string {
	if d.DomainCheck() {
		return strings.Join([]string{d.sld, d.tld}, ".")
	}
	return ""
}

func (d *Domain) Subdomain() string {
	if d.SubdomainCheck() {
		return strings.Join([]string{d.trd, d.sld, d.tld}, ".")
	}
	return ""
}

func (d *Domain) DomainCheck() bool {
	return !(d.tld == "" || d.sld == "")
}

func (d *Domain) SubdomainCheck() bool {
	return !(d.tld == "" || d.sld == "" || d.trd == "")
}
