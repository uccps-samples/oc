package syncgroups

// These constants contain values for annotations and labels affixed to Groups by the LDAP sync job
const (
	// LDAPHostLabel is the Label value that stores the host of the LDAP server
	// TODO: we don't store port here because labels don't allow for colons. We might want to add this back
	// with a different separator
	LDAPHostLabel string = "uccp.io/ldap.host"

	// LDAPURLAnnotation is the Annotation value that stores the host:port of the LDAP server
	LDAPURLAnnotation string = "uccp.io/ldap.url"
	// LDAPUIDAnnotation is the Annotation value that stores the corresponding LDAP group UID for the Group
	LDAPUIDAnnotation string = "uccp.io/ldap.uid"
	// LDAPSyncTime is the Annotation value that stores the last time this Group was synced with LDAP
	LDAPSyncTimeAnnotation string = "uccp.io/ldap.sync-time"
)
