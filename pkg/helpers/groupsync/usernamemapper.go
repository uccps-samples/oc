package syncgroups

import (
	"fmt"

	"gopkg.in/ldap.v2"

	"github.com/uccps-samples/library-go/pkg/security/ldaputil"
	"github.com/uccps-samples/oc/pkg/helpers/groupsync/interfaces"
)

// NewUserNameMapper returns a new DefaultLDAPGroupUserNameMapper
func NewUserNameMapper(nameAttributes []string) interfaces.LDAPUserNameMapper {
	return &DefaultLDAPUserNameMapper{
		nameAttributes: nameAttributes,
	}
}

// DefaultLDAPUserNameMapper extracts the OpenShift User name of an LDAP entry representing
// a user in a deterministic manner
type DefaultLDAPUserNameMapper struct {
	nameAttributes []string
}

func (m *DefaultLDAPUserNameMapper) UserNameFor(ldapUser *ldap.Entry) (string, error) {
	openShiftUserName := ldaputil.GetAttributeValue(ldapUser, m.nameAttributes)
	if len(openShiftUserName) == 0 {
		return "", fmt.Errorf("the user entry (%v) does not map to a Uccp User name with the given mapping",
			ldapUser)
	}
	return openShiftUserName, nil
}
