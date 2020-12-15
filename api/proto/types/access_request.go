package types

import (
	fmt "fmt"
	time "time"

	"github.com/gravitational/trace"

	"github.com/jonboulle/clockwork"
	"github.com/pborman/uuid"
)

// AccessRequest is a request for temporarily granted roles
type AccessRequest interface {
	Resource
	// GetUser gets the name of the requesting user
	GetUser() string
	// GetRoles gets the roles being requested by the user
	GetRoles() []string
	// SetRoles overrides the roles being requested by the user
	SetRoles([]string)
	// GetState gets the current state of the request
	GetState() RequestState
	// SetState sets the approval state of the request
	SetState(RequestState) error
	// GetCreationTime gets the time at which the request was
	// originally registered with the auth server.
	GetCreationTime() time.Time
	// SetCreationTime sets the creation time of the request.
	SetCreationTime(time.Time)
	// GetAccessExpiry gets the upper limit for which this request
	// may be considered active.
	GetAccessExpiry() time.Time
	// SetAccessExpiry sets the upper limit for which this request
	// may be considered active.
	SetAccessExpiry(time.Time)
	// GetRequestReason gets the reason for the request's creation.
	GetRequestReason() string
	// SetRequestReason sets the reason for the request's creation.
	SetRequestReason(string)
	// GetResolveReason gets the reasson for the request's resolution.
	GetResolveReason() string
	// SetResolveReason sets the reason for the request's resolution.
	SetResolveReason(string)
	// GetResolveAnnotations gets the annotations associated with
	// the request's resolution.
	GetResolveAnnotations() map[string][]string
	// SetResolveAnnotations sets the annotations associated with
	// the request's resolution.
	SetResolveAnnotations(map[string][]string)
	// GetSystemAnnotations gets the teleport-applied annotations.
	GetSystemAnnotations() map[string][]string
	// SetSystemAnnotations sets the teleport-applied annotations.
	SetSystemAnnotations(map[string][]string)
	// CheckAndSetDefaults validates the access request and
	// supplies default values where appropriate.
	CheckAndSetDefaults() error
	// Equals checks equality between access request values.
	Equals(AccessRequest) bool
}

func (r *AccessRequestV3) GetUser() string {
	return r.Spec.User
}

func (r *AccessRequestV3) GetRoles() []string {
	return r.Spec.Roles
}

func (r *AccessRequestV3) SetRoles(roles []string) {
	r.Spec.Roles = roles
}

func (r *AccessRequestV3) GetState() RequestState {
	return r.Spec.State
}

func (r *AccessRequestV3) SetState(state RequestState) error {
	if r.Spec.State.IsDenied() {
		if state.IsDenied() {
			return nil
		}
		return trace.BadParameter("cannot set request-state %q (already denied)", state.String())
	}
	r.Spec.State = state
	return nil
}

func (r *AccessRequestV3) GetCreationTime() time.Time {
	return r.Spec.Created
}

func (r *AccessRequestV3) SetCreationTime(t time.Time) {
	r.Spec.Created = t
}

func (r *AccessRequestV3) GetAccessExpiry() time.Time {
	return r.Spec.Expires
}

func (r *AccessRequestV3) SetAccessExpiry(expiry time.Time) {
	r.Spec.Expires = expiry
}

func (r *AccessRequestV3) GetRequestReason() string {
	return r.Spec.RequestReason
}

func (r *AccessRequestV3) SetRequestReason(reason string) {
	r.Spec.RequestReason = reason
}

func (r *AccessRequestV3) GetResolveReason() string {
	return r.Spec.ResolveReason
}

func (r *AccessRequestV3) SetResolveReason(reason string) {
	r.Spec.ResolveReason = reason
}

func (r *AccessRequestV3) GetResolveAnnotations() map[string][]string {
	return r.Spec.ResolveAnnotations
}

func (r *AccessRequestV3) SetResolveAnnotations(annotations map[string][]string) {
	r.Spec.ResolveAnnotations = annotations
}

func (r *AccessRequestV3) GetSystemAnnotations() map[string][]string {
	return r.Spec.SystemAnnotations
}

func (r *AccessRequestV3) SetSystemAnnotations(annotations map[string][]string) {
	r.Spec.SystemAnnotations = annotations
}

func (r *AccessRequestV3) CheckAndSetDefaults() error {
	if err := r.Metadata.CheckAndSetDefaults(); err != nil {
		return trace.Wrap(err)
	}
	if r.GetState().IsNone() {
		if err := r.SetState(RequestState_PENDING); err != nil {
			return trace.Wrap(err)
		}
	}
	if err := r.Check(); err != nil {
		return trace.Wrap(err)
	}
	return nil
}

func (r *AccessRequestV3) Check() error {
	if r.Kind == "" {
		return trace.BadParameter("access request kind not set")
	}
	if r.Version == "" {
		return trace.BadParameter("access request version not set")
	}
	if r.GetName() == "" {
		return trace.BadParameter("access request id not set")
	}
	if uuid.Parse(r.GetName()) == nil {
		return trace.BadParameter("invalid access request id %q", r.GetName())
	}
	if r.GetUser() == "" {
		return trace.BadParameter("access request user name not set")
	}
	if len(r.GetRoles()) < 1 {
		return trace.BadParameter("access request does not specify any roles")
	}
	if r.GetState().IsPending() {
		if r.GetResolveReason() != "" {
			return trace.BadParameter("pending requests cannot include resolve reason")
		}
		if len(r.GetResolveAnnotations()) != 0 {
			return trace.BadParameter("pending requests cannot include resolve annotations")
		}
	}
	return nil
}

func (r *AccessRequestV3) GetKind() string {
	return r.Kind
}

func (r *AccessRequestV3) GetSubKind() string {
	return r.SubKind
}

func (r *AccessRequestV3) SetSubKind(subKind string) {
	r.SubKind = subKind
}

func (r *AccessRequestV3) GetVersion() string {
	return r.Version
}

func (r *AccessRequestV3) GetName() string {
	return r.Metadata.Name
}

func (r *AccessRequestV3) SetName(name string) {
	r.Metadata.Name = name
}

func (r *AccessRequestV3) Expiry() time.Time {
	return r.Metadata.Expiry()
}

func (r *AccessRequestV3) SetExpiry(expiry time.Time) {
	r.Metadata.SetExpiry(expiry)
}

func (r *AccessRequestV3) SetTTL(clock clockwork.Clock, ttl time.Duration) {
	r.Metadata.SetTTL(clock, ttl)
}

func (r *AccessRequestV3) GetMetadata() Metadata {
	return r.Metadata
}

func (r *AccessRequestV3) GetResourceID() int64 {
	return r.Metadata.GetID()
}

func (r *AccessRequestV3) SetResourceID(id int64) {
	r.Metadata.SetID(id)
}

func (r *AccessRequestV3) String() string {
	return fmt.Sprintf("AccessRequest(user=%v,roles=%+v)", r.Spec.User, r.Spec.Roles)
}

func (r *AccessRequestV3) Equals(other AccessRequest) bool {
	o, ok := other.(*AccessRequestV3)
	if !ok {
		return false
	}
	if r.GetName() != o.GetName() {
		return false
	}
	return r.Spec.Equals(&o.Spec)
}

// stateVariants allows iteration of the expected variants
// of RequestState.
var stateVariants = [4]RequestState{
	RequestState_NONE,
	RequestState_PENDING,
	RequestState_APPROVED,
	RequestState_DENIED,
}

// Parse attempts to interpret a value as a string representation
// of a RequestState.
func (s *RequestState) Parse(val string) error {
	for _, state := range stateVariants {
		if state.String() == val {
			*s = state
			return nil
		}
	}
	return trace.BadParameter("unknown request state: %q", val)
}

func (s RequestState) IsNone() bool {
	return s == RequestState_NONE
}

func (s RequestState) IsPending() bool {
	return s == RequestState_PENDING
}

func (s RequestState) IsApproved() bool {
	return s == RequestState_APPROVED
}

func (s RequestState) IsDenied() bool {
	return s == RequestState_DENIED
}

func (s RequestState) IsResolved() bool {
	return s.IsApproved() || s.IsDenied()
}

func (s *AccessRequestSpecV3) Equals(other *AccessRequestSpecV3) bool {
	if s.User != other.User {
		return false
	}
	if len(s.Roles) != len(other.Roles) {
		return false
	}
	for i, role := range s.Roles {
		if role != other.Roles[i] {
			return false
		}
	}
	if s.Created != other.Created {
		return false
	}
	if s.Expires != other.Expires {
		return false
	}
	return s.State == other.State
}

// key values for map encoding of request filter
const (
	keyID    = "id"
	keyUser  = "user"
	keyState = "state"
)

func (f *AccessRequestFilter) IntoMap() map[string]string {
	m := make(map[string]string)
	if f.ID != "" {
		m[keyID] = f.ID
	}
	if f.User != "" {
		m[keyUser] = f.User
	}
	if !f.State.IsNone() {
		m[keyState] = f.State.String()
	}
	return m
}

func (f *AccessRequestFilter) FromMap(m map[string]string) error {
	for key, val := range m {
		switch key {
		case keyID:
			f.ID = val
		case keyUser:
			f.User = val
		case keyState:
			if err := f.State.Parse(val); err != nil {
				return trace.Wrap(err)
			}
		default:
			return trace.BadParameter("unknown filter key %s", key)
		}
	}
	return nil
}

// Match checks if a given access request matches this filter.
func (f *AccessRequestFilter) Match(req AccessRequest) bool {
	if f.ID != "" && req.GetName() != f.ID {
		return false
	}
	if f.User != "" && req.GetUser() != f.User {
		return false
	}
	if !f.State.IsNone() && req.GetState() != f.State {
		return false
	}
	return true
}

func (f *AccessRequestFilter) Equals(o AccessRequestFilter) bool {
	return f.ID == o.ID && f.User == o.User && f.State == o.State
}
