package services

import (
	"github.com/gravitational/teleport/api"
	"github.com/gravitational/teleport/api/proto/types"
)

// The following types are implemented in /api/types, and imported/wrapped here.
// The new structs are used to wrap the imported types with additional methods.
// The other types are basic imports and can be removed if their references are updated.

type AccessRequest = types.AccessRequest
type AccessRequestV3 = types.AccessRequestV3
type AccessRequestSpecV3 = types.AccessRequestSpecV3
type AccessRequestFilter = types.AccessRequestFilter
type AccessRequestConditions struct{ types.AccessRequestConditions }
type RequestState = types.RequestState

type Duration = api.Duration

type PluginData = types.PluginData
type PluginDataV3 = types.PluginDataV3
type PluginDataSpecV3 = types.PluginDataSpecV3
type PluginDataFilter = types.PluginDataFilter
type PluginDataEntry = types.PluginDataEntry
type PluginDataUpdateParams = types.PluginDataUpdateParams

type ProvisionTokenV1 struct{ types.ProvisionTokenV1 }
type ProvisionTokenV2 struct{ types.ProvisionTokenV2 }
type ProvisionTokenSpecV2 = types.ProvisionTokenSpecV2

type RemoteClusterV3 struct{ types.RemoteClusterV3 }

type Resource = types.Resource
type ResourceHeader struct{ types.ResourceHeader }
type Metadata = types.Metadata

type Role = types.Role
type RoleV3 struct{ types.RoleV3 }
type RoleSpecV3 = types.RoleSpecV3
type RoleConditions struct{ types.RoleConditions }
type RoleConditionType = types.RoleConditionType
type RoleOptions struct{ types.RoleOptions }
type Rule struct{ types.Rule }

type WebSessionV2 struct{ types.WebSessionV2 }
type WebSessionSpecV2 = types.WebSessionSpecV2

// Some functions and variables also need to be imported from the types package
var (
	MaxDuration           = api.MaxDuration
	NewDuration           = api.NewDuration
	IsValidLabelKey       = types.IsValidLabelKey
	MetadataSchema        = types.MetadataSchema
	RequestState_NONE     = types.RequestState_NONE
	RequestState_PENDING  = types.RequestState_PENDING
	RequestState_APPROVED = types.RequestState_APPROVED
	RequestState_DENIED   = types.RequestState_DENIED
)

// The following Constants are imported from api to simplify
// refactoring. These could be removed and their references updated.
const (
	DefaultAPIGroup               = api.DefaultAPIGroup
	ActionRead                    = api.ActionRead
	ActionWrite                   = api.ActionWrite
	Wildcard                      = api.Wildcard
	KindNamespace                 = api.KindNamespace
	KindUser                      = api.KindUser
	KindKeyPair                   = api.KindKeyPair
	KindHostCert                  = api.KindHostCert
	KindJWT                       = api.KindJWT
	KindLicense                   = api.KindLicense
	KindRole                      = api.KindRole
	KindAccessRequest             = api.KindAccessRequest
	KindPluginData                = api.KindPluginData
	KindOIDC                      = api.KindOIDC
	KindSAML                      = api.KindSAML
	KindGithub                    = api.KindGithub
	KindOIDCRequest               = api.KindOIDCRequest
	KindSAMLRequest               = api.KindSAMLRequest
	KindGithubRequest             = api.KindGithubRequest
	KindSession                   = api.KindSession
	KindSSHSession                = api.KindSSHSession
	KindWebSession                = api.KindWebSession
	KindAppSession                = api.KindAppSession
	KindEvent                     = api.KindEvent
	KindAuthServer                = api.KindAuthServer
	KindProxy                     = api.KindProxy
	KindNode                      = api.KindNode
	KindAppServer                 = api.KindAppServer
	KindToken                     = api.KindToken
	KindCertAuthority             = api.KindCertAuthority
	KindReverseTunnel             = api.KindReverseTunnel
	KindOIDCConnector             = api.KindOIDCConnector
	KindSAMLConnector             = api.KindSAMLConnector
	KindGithubConnector           = api.KindGithubConnector
	KindConnectors                = api.KindConnectors
	KindClusterAuthPreference     = api.KindClusterAuthPreference
	MetaNameClusterAuthPreference = api.MetaNameClusterAuthPreference
	KindClusterConfig             = api.KindClusterConfig
	KindSemaphore                 = api.KindSemaphore
	MetaNameClusterConfig         = api.MetaNameClusterConfig
	KindClusterName               = api.KindClusterName
	MetaNameClusterName           = api.MetaNameClusterName
	KindStaticTokens              = api.KindStaticTokens
	MetaNameStaticTokens          = api.MetaNameStaticTokens
	KindTrustedCluster            = api.KindTrustedCluster
	KindAuthConnector             = api.KindAuthConnector
	KindTunnelConnection          = api.KindTunnelConnection
	KindRemoteCluster             = api.KindRemoteCluster
	KindResetPasswordToken        = api.KindResetPasswordToken
	KindResetPasswordTokenSecrets = api.KindResetPasswordTokenSecrets
	KindIdentity                  = api.KindIdentity
	KindState                     = api.KindState
	KindKubeService               = api.KindKubeService
	V3                            = api.V3
	V2                            = api.V2
	V1                            = api.V1
	VerbList                      = api.VerbList
	VerbCreate                    = api.VerbCreate
	VerbRead                      = api.VerbRead
	VerbReadNoSecrets             = api.VerbReadNoSecrets
	VerbUpdate                    = api.VerbUpdate
	VerbDelete                    = api.VerbDelete
	VerbRotate                    = api.VerbRotate
)
