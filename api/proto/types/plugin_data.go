package types

import (
	fmt "fmt"
	time "time"

	"github.com/gravitational/trace"
	"github.com/jonboulle/clockwork"
)

// PluginData is used by plugins to store per-resource state.  An instance of PluginData
// corresponds to a resource which may be managed by one or more plugins.  Data is stored
// as a mapping of the form `plugin -> key -> val`, effectively giving each plugin its own
// key-value store.  Importantly, an instance of PluginData can only be created for a resource
// which currently exist, and automatically expires shortly after the corresponding resource.
// Currently, only the AccessRequest resource is supported.
type PluginData interface {
	Resource
	// Entries gets all entries.
	Entries() map[string]*PluginDataEntry
	// Update attempts to apply an update.
	Update(params PluginDataUpdateParams) error
	// CheckAndSetDefaults validates the plugin data
	// and supplies default values where appropriate.
	CheckAndSetDefaults() error
}

func (r *PluginDataV3) CheckAndSetDefaults() error {
	if err := r.Metadata.CheckAndSetDefaults(); err != nil {
		return trace.Wrap(err)
	}
	if r.SubKind == "" {
		return trace.BadParameter("plugin data missing subkind")
	}
	return nil
}

func (r *PluginDataV3) Entries() map[string]*PluginDataEntry {
	if r.Spec.Entries == nil {
		r.Spec.Entries = make(map[string]*PluginDataEntry)
	}
	return r.Spec.Entries
}

func (r *PluginDataV3) Update(params PluginDataUpdateParams) error {
	// See #3286 for a complete discussion of the design constraints at play here.

	if params.Kind != r.GetSubKind() {
		return trace.BadParameter("resource kind mismatch in update params")
	}

	if params.Resource != r.GetName() {
		return trace.BadParameter("resource name mismatch in update params")
	}

	// If expectations were given, ensure that they are met before continuing
	if params.Expect != nil {
		if err := r.checkExpectations(params.Plugin, params.Expect); err != nil {
			return trace.Wrap(err)
		}
	}
	// Ensure that Entries has been initialized
	if r.Spec.Entries == nil {
		r.Spec.Entries = make(map[string]*PluginDataEntry, 1)
	}
	// Ensure that the specific Plugin has been initialized
	if r.Spec.Entries[params.Plugin] == nil {
		r.Spec.Entries[params.Plugin] = &PluginDataEntry{
			Data: make(map[string]string, len(params.Set)),
		}
	}
	entry := r.Spec.Entries[params.Plugin]
	for key, val := range params.Set {
		// Keys which are explicitly set to the empty string are
		// treated as DELETE operations.
		if val == "" {
			delete(entry.Data, key)
			continue
		}
		entry.Data[key] = val
	}
	// Its possible that this update was simply clearing all data;
	// if that is the case, remove the entry.
	if len(entry.Data) == 0 {
		delete(r.Spec.Entries, params.Plugin)
	}
	return nil
}

// checkExpectations verifies that the data for `plugin` matches the expected
// state described by `expect`.  This function implements the behavior of the
// `PluginDataUpdateParams.Expect` mapping.
func (r *PluginDataV3) checkExpectations(plugin string, expect map[string]string) error {
	var entry *PluginDataEntry
	if r.Spec.Entries != nil {
		entry = r.Spec.Entries[plugin]
	}
	if entry == nil {
		// If no entry currently exists, then the only expectation that can
		// match is one which only specifies fields which shouldn't exist.
		for key, val := range expect {
			if val != "" {
				return trace.CompareFailed("expectations not met for field %q", key)
			}
		}
		return nil
	}
	for key, val := range expect {
		if entry.Data[key] != val {
			return trace.CompareFailed("expectations not met for field %q", key)

		}
	}
	return nil
}

func (f *PluginDataFilter) Match(data PluginData) bool {
	if f.Kind != "" && f.Kind != data.GetSubKind() {
		return false
	}
	if f.Resource != "" && f.Resource != data.GetName() {
		return false
	}
	if f.Plugin != "" {
		if _, ok := data.Entries()[f.Plugin]; !ok {
			return false
		}
	}
	return true
}

func (d *PluginDataEntry) Equals(other *PluginDataEntry) bool {
	if other == nil {
		return false
	}
	if len(d.Data) != len(other.Data) {
		return false
	}
	for key, val := range d.Data {
		if other.Data[key] != val {
			return false
		}
	}
	return true
}

func (r *PluginDataV3) GetKind() string {
	return r.Kind
}

func (r *PluginDataV3) GetSubKind() string {
	return r.SubKind
}

func (r *PluginDataV3) SetSubKind(subKind string) {
	r.SubKind = subKind
}

func (r *PluginDataV3) GetVersion() string {
	return r.Version
}

func (r *PluginDataV3) GetName() string {
	return r.Metadata.Name
}

func (r *PluginDataV3) SetName(name string) {
	r.Metadata.Name = name
}

func (r *PluginDataV3) Expiry() time.Time {
	return r.Metadata.Expiry()
}

func (r *PluginDataV3) SetExpiry(expiry time.Time) {
	r.Metadata.SetExpiry(expiry)
}

func (r *PluginDataV3) SetTTL(clock clockwork.Clock, ttl time.Duration) {
	r.Metadata.SetTTL(clock, ttl)
}

func (r *PluginDataV3) GetMetadata() Metadata {
	return r.Metadata
}

func (r *PluginDataV3) GetResourceID() int64 {
	return r.Metadata.GetID()
}

func (r *PluginDataV3) SetResourceID(id int64) {
	r.Metadata.SetID(id)
}

func (r *PluginDataV3) String() string {
	return fmt.Sprintf("PluginData(kind=%s,resource=%s,entries=%d)", r.GetSubKind(), r.GetName(), len(r.Spec.Entries))
}
