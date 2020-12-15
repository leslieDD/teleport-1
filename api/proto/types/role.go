package types

// RoleConditionType specifies if it's an allow rule (true) or deny rule (false).
type RoleConditionType bool

// Role contains a set of permissions or settings
type Role interface {
	// Resource provides common resource methods.
	Resource
	// CheckAndSetDefaults checks and set default values for any missing fields.
	CheckAndSetDefaults() error
	// Equals returns true if the roles are equal. Roles are equal if options and
	// conditions match.
	Equals(other Role) bool
	// ApplyTraits applies the passed in traits to any variables within the role
	// and returns itself.
	ApplyTraits(map[string][]string) Role

	// GetOptions gets role options.
	GetOptions() RoleOptions
	// SetOptions sets role options
	SetOptions(opt RoleOptions)

	// GetLogins gets *nix system logins for allow or deny condition.
	GetLogins(RoleConditionType) []string
	// SetLogins sets *nix system logins for allow or deny condition.
	SetLogins(RoleConditionType, []string)

	// GetNamespaces gets a list of namespaces this role is allowed or denied access to.
	GetNamespaces(RoleConditionType) []string
	// GetNamespaces sets a list of namespaces this role is allowed or denied access to.
	SetNamespaces(RoleConditionType, []string)

	// GetNodeLabels gets the map of node labels this role is allowed or denied access to.
	GetNodeLabels(RoleConditionType) Labels
	// SetNodeLabels sets the map of node labels this role is allowed or denied access to.
	SetNodeLabels(RoleConditionType, Labels)

	// GetAppLabels gets the map of app labels this role is allowed or denied access to.
	GetAppLabels(RoleConditionType) Labels
	// SetAppLabels sets the map of app labels this role is allowed or denied access to.
	SetAppLabels(RoleConditionType, Labels)

	// GetClusterLabels gets the map of cluster labels this role is allowed or denied access to.
	GetClusterLabels(RoleConditionType) Labels
	// SetClusterLabels sets the map of cluster labels this role is allowed or denied access to.
	SetClusterLabels(RoleConditionType, Labels)

	// GetKubernetesLabels gets the map of kubernetes labels this role is
	// allowed or denied access to.
	GetKubernetesLabels(RoleConditionType) Labels
	// SetKubernetesLabels sets the map of kubernetes labels this role is
	// allowed or denied access to.
	SetKubernetesLabels(RoleConditionType, Labels)

	// GetRules gets all allow or deny rules.
	GetRules(rct RoleConditionType) []Rule
	// SetRules sets an allow or deny rule.
	SetRules(rct RoleConditionType, rules []Rule)

	// GetKubeGroups returns kubernetes groups
	GetKubeGroups(RoleConditionType) []string
	// SetKubeGroups sets kubernetes groups for allow or deny condition.
	SetKubeGroups(RoleConditionType, []string)

	// GetKubeUsers returns kubernetes users to impersonate
	GetKubeUsers(RoleConditionType) []string
	// SetKubeUsers sets kubernetes users to impersonate for allow or deny condition.
	SetKubeUsers(RoleConditionType, []string)

	// GetAccessRequestConditions gets allow/deny conditions for access requests.
	GetAccessRequestConditions(RoleConditionType) AccessRequestConditions
	// SetAccessRequestConditions sets allow/deny conditions for access requests.
	SetAccessRequestConditions(RoleConditionType, AccessRequestConditions)
}

// GetVersion returns resource version
func (r *RoleV3) GetVersion() string {
	return r.Version
}

// GetKind returns resource kind
func (r *RoleV3) GetKind() string {
	return r.Kind
}

// GetSubKind returns resource sub kind
func (r *RoleV3) GetSubKind() string {
	return r.SubKind
}

// SetSubKind sets resource subkind
func (r *RoleV3) SetSubKind(s string) {
	r.SubKind = s
}

// GetResourceID returns resource ID
func (r *RoleV3) GetResourceID() int64 {
	return r.Metadata.ID
}

// SetResourceID sets resource ID
func (r *RoleV3) SetResourceID(id int64) {
	r.Metadata.ID = id
}
