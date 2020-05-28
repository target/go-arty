// Copyright (c) 2016 John E. Vincent
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Copyright (c) 2018 Target Brands, Inc.

package artifactory

import (
	"bytes"
	"encoding/xml"
	"gopkg.in/yaml.v2"
	"net/http"
)

// SystemService handles communication with the system related
// methods of the Artifactory API.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-SYSTEM&CONFIGURATION
type SystemService service

// Versions represents the version information about Artifactory.
type Versions struct {
	Version  *string   `json:"version,omitempty"`
	Revision *string   `json:"revision,omitempty"`
	Addons   *[]string `json:"addons,omitempty"`
}

func (v Versions) String() string {
	return Stringify(v)
}

// GlobalConfigCommon represents elements of the Global Configuration Descriptor
// that are common between a GlobalConfigRequest and GlobalConfigResponse.
// Lots of elements aren't documented but have been mapped from the
// XML schema at https://www.jfrog.com/public/xsd/artifactory-v2_2_5.xsd
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File
type GlobalConfigCommon struct {
	ServerName                *string                    `yaml:"serverName,omitempty" xml:"serverName,omitempty"`
	OfflineMode               *bool                      `yaml:"offlineMode,omitempty" xml:"offlineMode,omitempty"`
	HelpLinksEnabled          *bool                      `yaml:"helpLinksEnabled,omitempty" xml:"helpLinksEnabled,omitempty"`
	FileUploadMaxSizeMb       *int                       `yaml:"fileUploadMaxSizeMb,omitempty" xml:"fileUploadMaxSizeMb,omitempty"`
	DateFormat                *string                    `yaml:"dateFormat,omitempty" xml:"dateFormat,omitempty"`
	AddonsConfig              *AddonsConfig              `yaml:"addons,omitempty" xml:"addons,omitempty"`
	MailServer                *MailServer                `yaml:"mailServer,omitempty" xml:"mailServer,omitempty"`
	XrayConfig                *XrayConfig                `yaml:"xrayConfig,omitempty" xml:"xrayConfig,omitempty"`
	BintrayConfig             *BintrayConfig             `yaml:"bintrayConfig,omitempty" xml:"bintrayConfig,omitempty"`
	Indexer                   *Indexer                   `yaml:"indexer,omitempty" xml:"indexer,omitempty"`
	UrlBase                   *string                    `yaml:"urlBase,omitempty" xml:"urlBase,omitempty"`
	Logo                      *string                    `yaml:"logo,omitempty" xml:"logo,omitempty"`
	Footer                    *string                    `yaml:"footer,omitempty" xml:"footer,omitempty"`
	GcConfig                  *GcConfig                  `yaml:"gcConfig,omitempty" xml:"gcConfig,omitempty"`
	CleanupConfig             *CleanupConfig             `yaml:"cleanupConfig,omitempty" xml:"cleanupConfig,omitempty"`
	VirtualCacheCleanupConfig *VirtualCacheCleanupConfig `yaml:"virtualCacheCleanupConfig,omitempty" xml:"virtualCacheCleanupConfig,omitempty"`
	QuotaConfig               *QuotaConfig               `yaml:"quotaConfig,omitempty" xml:"quotaConfig,omitempty"`
	SystemMessageConfig       *SystemMessageConfig       `yaml:"systemMessageConfig,omitempty" xml:"systemMessageConfig,omitempty"`
	FolderDownloadConfig      *FolderDownloadConfig      `yaml:"folderDownloadConfig,omitempty" xml:"folderDownloadConfig,omitempty"`
	TrashcanConfig            *TrashcanConfig            `yaml:"trashcanConfig,omitempty" xml:"trashcanConfig,omitempty"`
	ReplicationsConfig        *ReplicationsConfig        `yaml:"replicationsConfig,omitempty" xml:"replicationsConfig,omitempty"`
	SumoLogicConfig           *SumoLogicConfig           `yaml:"sumoLogicConfig,omitempty" xml:"sumoLogicConfig,omitempty"`
	ReleaseBundlesConfig      *ReleaseBundlesConfig      `yaml:"releaseBundlesConfig,omitempty" xml:"releaseBundlesConfig,omitempty"`
	SignedUrlConfig           *SignedUrlConfig           `yaml:"signedUrlConfig,omitempty" xml:"signedUrlConfig,omitempty"`
}

// GlobalConfigRequest represents elements of the Global Configuration Descriptor
// that can be updated in a PATCH request.
// Limitations:
// 1) Repository and repository replication configuration is omitted as the
//    Repositories service methods should be used instead.
// 2) It isn't possible to reset config values to their defaults using nil values due
//    to the inclusion of the `omitempty` yaml tag on fields.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File
type GlobalConfigRequest struct {
	GlobalConfigCommon `yaml:",inline"`
	Security           *struct {
		AnonAccessEnabled                *bool                             `yaml:"anonAccessEnabled,omitempty"`
		UserLockPolicy                   *UserLockPolicy                   `yaml:"userLockPolicy,omitempty"`
		PasswordSettings                 *PasswordSettings                 `yaml:"passwordSettings,omitempty"`
		LdapSettings                     *map[string]*LdapSetting          `yaml:"ldapSettings,omitempty"`
		LdapGroupSettings                *map[string]*LdapGroupSetting     `yaml:"ldapGroupSettings,omitempty"`
		HttpSsoSettings                  *HttpSsoSettings                  `yaml:"httpSsoSettings,omitempty"`
		CrowdSettings                    *CrowdSettings                    `yaml:"crowdSettings,omitempty"`
		SamlSettings                     *SamlSettings                     `yaml:"samlSettings,omitempty"`
		OauthSettings                    *OauthSettings                    `yaml:"oauthSettings,omitempty"`
		AccessClientSettings             *AccessClientSettings             `yaml:"accessClientSettings,omitempty"`
		BuildGlobalBasicReadAllowed      *BuildGlobalBasicReadAllowed      `yaml:"buildGlobalBasicReadAllowed,omitempty"`
		BuildGlobalBasicReadForAnonymous *BuildGlobalBasicReadForAnonymous `yaml:"buildGlobalBasicReadForAnonymous,omitempty"`
	} `yaml:"security,omitempty"`
	Backups             *map[string]*Backup             `yaml:"backups,omitempty"`
	Proxies             *map[string]*Proxy              `yaml:"proxies,omitempty"`
	ReverseProxies      *map[string]*ReverseProxy       `yaml:"reverseProxies,omitempty"`
	PropertySets        *map[string]*PropertySetRequest `yaml:"propertySets,omitempty"`
	RepoLayouts         *map[string]*RepoLayout         `yaml:"repoLayouts,omitempty"`
	BintrayApplications *map[string]*BintrayApplication `yaml:"bintrayApplications,omitempty"`
}

func (g GlobalConfigRequest) String() string {
	return Stringify(g)
}

// GlobalConfigResponse represents the response to a GET request for the Global Configuration Descriptor.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File
type GlobalConfigResponse struct {
	*GlobalConfigCommon
	Revision *int `xml:"revision,omitempty"`
	Security *struct {
		AnonAccessEnabled                *bool                 `xml:"anonAccessEnabled,omitempty"`
		HideUnauthorizedResources        *bool                 `xml:"hideUnauthorizedResources,omitempty"`
		UserLockPolicy                   *UserLockPolicy       `xml:"userLockPolicy,omitempty"`
		PasswordSettings                 *PasswordSettings     `xml:"passwordSettings,omitempty"`
		LdapSettings                     *[]LdapSetting        `xml:"ldapSettings>ldapSetting,omitempty"`
		LdapGroupSettings                *[]LdapGroupSetting   `xml:"ldapGroupSettings>ldapGroupSetting,omitempty"`
		HttpSsoSettings                  *HttpSsoSettings      `xml:"httpSsoSettings,omitempty"`
		CrowdSettings                    *CrowdSettings        `xml:"crowdSettings,omitempty"`
		SamlSettings                     *SamlSettings         `xml:"samlSettings,omitempty"`
		OauthSettings                    *OauthSettings        `xml:"oauthSettings,omitempty"`
		AccessClientSettings             *AccessClientSettings `xml:"accessClientSettings,omitempty"`
		BuildGlobalBasicReadAllowed      *bool                 `xml:"buildGlobalBasicReadAllowed,omitempty"`
		BuildGlobalBasicReadForAnonymous *bool                 `xml:"buildGlobalBasicReadForAnonymous,omitempty"`
	} `xml:"security,omitempty"`
	Backups             *[]Backup              `xml:"backups>backup,omitempty"`
	LocalRepositories   *[]LocalRepository     `xml:"localRepositories>localRepository,omitempty"`
	RemoteRepositories  *[]RemoteRepository    `xml:"remoteRepositories>remoteRepository,omitempty"`
	VirtualRepositories *[]VirtualRepository   `xml:"virtualRepositories>virtualRepository,omitempty"`
	LocalReplications   *[]Replication         `xml:"localReplications>localReplication,omitempty"`
	RemoteReplications  *[]Replication         `xml:"remoteReplications>remoteReplication,omitempty"`
	Proxies             *[]Proxy               `xml:"proxies>proxy,omitempty"`
	ReverseProxies      *[]ReverseProxy        `xml:"reverseProxies>reverseProxy,omitempty"`
	PropertySets        *[]PropertySetResponse `xml:"propertySets>propertySet,omitempty"`
	RepoLayouts         *[]RepoLayout          `xml:"repoLayouts>repoLayout,omitempty"`
	BintrayApplications *[]BintrayApplication  `xml:"bintrayApplications>bintrayApplication,omitempty"`
}

func (g GlobalConfigResponse) String() string {
	return Stringify(g)
}

// AddonsConfig represents Addons-related configuration.
// This is undocumented in YAML Configuration File.
//
type AddonsConfig struct {
	ShowAddonsInfo       *bool   `yaml:"showAddonsInfo,omitempty" xml:"showAddonsInfo,omitempty"`
	ShowAddonsInfoCookie *string `yaml:"showAddonsInfoCookie,omitempty" xml:"showAddonsInfoCookie,omitempty"`
}

// MailServer represents a Mail Server setting in Artifactory General Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-General(General,PropertySets,Proxy,Mail)
type MailServer struct {
	Enabled        *bool   `yaml:"enabled,omitempty" xml:"enabled,omitempty"`
	ArtifactoryUrl *string `yaml:"artifactoryUrl,omitempty" xml:"artifactoryUrl,omitempty"`
	From           *string `yaml:"from,omitempty" xml:"from,omitempty"`
	Host           *string `yaml:"host,omitempty" xml:"host,omitempty"`
	Username       *string `yaml:"username,omitempty" xml:"username,omitempty"`
	Password       *string `yaml:"password,omitempty" xml:"password,omitempty"`
	Port           *int    `yaml:"port,omitempty" xml:"port,omitempty"`
	SubjectPrefix  *string `yaml:"subjectPrefix,omitempty" xml:"subjectPrefix,omitempty"`
	Ssl            *bool   `yaml:"ssl,omitempty" xml:"ssl,omitempty"`
	Tls            *bool   `yaml:"tls,omitempty" xml:"tls,omitempty"`
}

// XrayConfig represents Xray related settings in Artifactory's Configuration
// Descriptor. This is undocumented in YAML Configuration File.
//
type XrayConfig struct {
	Enabled                       *bool   `yaml:"enabled,omitempty" xml:"enabled,omitempty"`
	BaseUrl                       *string `yaml:"baseUrl,omitempty" xml:"baseUrl,omitempty"`
	User                          *string `yaml:"user,omitempty" xml:"user,omitempty"`
	Password                      *string `yaml:"password,omitempty" xml:"password,omitempty"`
	ArtifactoryId                 *string `yaml:"artifactoryId,omitempty" xml:"artifactoryId,omitempty"`
	XrayId                        *string `yaml:"xrayId,omitempty" xml:"xrayId,omitempty"`
	AllowDownloadsXrayUnavailable *bool   `yaml:"allowDownloadsXrayUnavailable,omitempty" xml:"allowDownloadsXrayUnavailable,omitempty"`
	AllowBlockedArtifactsDownload *bool   `yaml:"allowBlockedArtifactsDownload,omitempty" xml:"allowBlockedArtifactsDownload,omitempty"`
	BypassDefaultProxy            *bool   `yaml:"bypassDefaultProxy,omitempty" xml:"bypassDefaultProxy,omitempty"`
	Proxy                         *string `yaml:"proxy,omitempty" xml:"proxy,omitempty"`
}

// BintrayConfig represents Bintray settings in Artifactory's Configuration
// Descriptor. This is undocumented in YAML Configuration File.
//
type BintrayConfig struct {
	FileUploadLimit *int `yaml:"fileUploadLimit,omitempty" xml:"fileUploadLimit,omitempty"`
}

// Proxy represents a Proxy setting in Artifactory General Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-General(General,PropertySets,Proxy,Mail)
type Proxy struct {
	Key             *string `yaml:"-" xml:"key,omitempty"`
	Domain          *string `yaml:"domain,omitempty" xml:"domain,omitempty"`
	Host            *string `yaml:"host,omitempty" xml:"host,omitempty"`
	NtHost          *string `yaml:"ntHost,omitempty" xml:"ntHost,omitempty"`
	Password        *string `yaml:"password,omitempty" xml:"password,omitempty"`
	Port            *int    `yaml:"port,omitempty" xml:"port,omitempty"`
	RedirectToHosts *string `yaml:"redirectToHosts,omitempty" xml:"redirectedToHosts,omitempty"`
	Username        *string `yaml:"username,omitempty" xml:"username,omitempty"`
	DefaultProxy    *bool   `yaml:"defaultProxy,omitempty" xml:"defaultProxy,omitempty"`
}

// ReverseProxy represents a Reverse Proxy configuration in Artifactory HTTP Settings.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-GetReverseProxyConfiguration
type ReverseProxy struct {
	Key                      *string `yaml:"-" xml:"key,omitempty"`
	WebServerType            *string `yaml:"webServerType,omitempty" xml:"webServerType,omitempty"`
	ArtifactoryAppContext    *string `yaml:"artifactoryAppContext,omitempty" xml:"artifactoryAppContext,omitempty"`
	PublicAppContext         *string `yaml:"publicAppContext,omitempty" xml:"publicAppContext,omitempty"`
	ServerName               *string `yaml:"serverName,omitempty" xml:"serverName,omitempty"`
	ServerNameExpression     *string `yaml:"serverNameExpression,omitempty" xml:"serverNameExpression,omitempty"`
	SslCertificate           *string `yaml:"sslCertificate,omitempty" xml:"sslCertificate,omitempty"`
	SslKey                   *string `yaml:"sslKey,omitempty" xml:"sslKey,omitempty"`
	DockerReverseProxyMethod *string `yaml:"dockerReverseProxyMethod,omitempty" xml:"dockerReverseProxyMethod,omitempty"`
	UseHttps                 *bool   `yaml:"useHttps,omitempty" xml:"useHttps,omitempty"`
	UseHttp                  *bool   `yaml:"useHttp,omitempty" xml:"useHttp,omitempty"`
	SslPort                  *int    `yaml:"sslPort,omitempty" xml:"sslPort,omitempty"`
	HttpPort                 *int    `yaml:"httpPort,omitempty" xml:"httpPort,omitempty"`
	ArtifactoryServerName    *string `yaml:"artifactoryServerName,omitempty" xml:"artifactoryServerName,omitempty"`
	UpStreamName             *string `yaml:"upStreamName,omitempty" xml:"upStreamName,omitempty"`
	ArtifactoryPort          *int    `yaml:"artifactoryPort,omitempty" xml:"artifactoryPort,omitempty"`
}

// PropertySetResponse represents a Property Set in a response to a GET request for Artifactory General Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-General(General,PropertySets,Proxy,Mail)
type PropertySetResponse struct {
	Name       *string `xml:"name,omitempty"`
	Properties *[]struct {
		Name             *string `xml:"name,omitempty"`
		PredefinedValues *[]struct {
			Value        *string `xml:"value,omitempty"`
			DefaultValue *bool   `xml:"defaultValue,omitempty"`
		} `xml:"predefinedValues>predefinedValue,omitempty"`
		ClosedPredefinedValues *bool `xml:"closedPredefinedValues,omitempty"`
		MultipleChoice         *bool `xml:"multipleChoice,omitempty"`
	} `xml:"properties>property,omitempty"`
	Visible *bool `xml:"visible,omitempty"`
}

// PropertySetRequest represents a Property Set in a PATCH request to update Artifactory General Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-General(General,PropertySets,Proxy,Mail)
type PropertySetRequest struct {
	Properties *[]struct {
		Name             *string `yaml:"name,omitempty"`
		PredefinedValues *map[string]struct {
			DefaultValue *bool `yaml:"defaultValue,omitempty"`
		} `yaml:"predefinedValues,omitempty"`
		ClosedPredefinedValues *bool `yaml:"closedPredefinedValues,omitempty"`
		MultipleChoice         *bool `yaml:"multipleChoice,omitempty"`
	} `yaml:"properties,omitempty"`
	Visible *bool `yaml:"visible,omitempty"`
}

// PasswordSettings represents the Password settings in Artifactory Security Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Security(Generalsecurity,PasswordPolicy,LDAP,SAML,OAuth,HTTPSSO,Crowd)
type PasswordSettings struct {
	EncryptionPolicy *string           `yaml:"encryptionPolicy,omitempty" xml:"encryptionPolicy,omitempty"`
	ExpirationPolicy *ExpirationPolicy `yaml:"expirationPolicy,omitempty" xml:"expirationPolicy,omitempty"`
	ResetPolicy      *ResetPolicy      `yaml:"resetPolicy,omitempty" xml:"resetPolicy,omitempty"`
}

// ExpirationPolicy represents the Password Expiration Policy settings in Artifactory Security Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Security(Generalsecurity,PasswordPolicy,LDAP,SAML,OAuth,HTTPSSO,Crowd)
type ExpirationPolicy struct {
	Enabled        *bool `yaml:"enabled,omitempty" xml:"enabled,omitempty"`
	PasswordMaxAge *int  `yaml:"passwordMaxAge,omitempty" xml:"passwordMaxAge,omitempty"`
	NotifyByEmail  *bool `yaml:"notifyByEmail,omitempty" xml:"notifyByEmail,omitempty"`
}

// ResetPolicy represents the Password Reset Protection policy settings in Artifactory Security Configuration.
// This is undocumented in YAML Configuration File.
//
type ResetPolicy struct {
	Enabled               *bool `yaml:"enabled,omitempty" xml:"enabled,omitempty"`
	MaxAttemptsPerAddress *int  `yaml:"maxAttemptsPerAddress,omitempty" xml:"maxAttemptsPerAddress,omitempty"`
	TimeToBlockInMinutes  *int  `yaml:"timeToBlockInMinutes,omitempty" xml:"timeToBlockInMinutes,omitempty"`
}

// UserLockPolicy represents the User Lock Policy settings in Artifactory Security Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Security(Generalsecurity,PasswordPolicy,LDAP,SAML,OAuth,HTTPSSO,Crowd)
type UserLockPolicy struct {
	Enabled       *bool `yaml:"enabled,omitempty" xml:"enabled,omitempty"`
	LoginAttempts *int  `yaml:"loginAttempts,omitempty" xml:"loginAttempts,omitempty"`
}

// SigningKeysSetting represents the GPG Signing settings in Artifactory Security Configuration.
// This is undocumented in YAML Configuration File.
//
type SigningKeysSettings struct {
	Passphrase       *string `yaml:"passphrase,omitempty" xml:"passphrase,omitempty"`
	KeyStorePassword *string `yaml:"keyStorePassword,omitempty" xml:"keyStorePassword,omitempty"`
}

// LdapSetting represents the LDAP settings in Artifactory Security Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Security(Generalsecurity,PasswordPolicy,LDAP,SAML,OAuth,HTTPSSO,Crowd)
type LdapSetting struct {
	Key                     *string `yaml:"-" xml:"key,omitempty"`
	EmailAttribute          *string `yaml:"emailAttribute,omitempty" xml:"emailAttribute,omitempty"`
	LdapPoisoningProtection *bool   `yaml:"ldapPoisoningProtection,omitempty" xml:"ldapPoisoningProtection,omitempty"`
	LdapUrl                 *string `yaml:"ldapUrl,omitempty" xml:"ldapUrl,omitempty"`
	Search                  *struct {
		ManagerDn       *string `yaml:"managerDn,omitempty" xml:"managerDn,omitempty"`
		ManagerPassword *string `yaml:"managerPassword,omitempty" xml:"managerPassword,omitempty"`
		SearchBase      *string `yaml:"searchBase,omitempty" xml:"searchBase,omitempty"`
		SearchFilter    *string `yaml:"searchFilter,omitempty" xml:"searchFilter,omitempty"`
		SearchSubTree   *bool   `yaml:"searchSubTree,omitempty" xml:"searchSubTree,omitempty"`
	} `yaml:"search,omitempty" xml:"search,omitempty"`
	UserDnPattern            *string `yaml:"userDnPattern,omitempty" xml:"userDnPattern,omitempty"`
	AllowUserToAccessProfile *bool   `yaml:"allowUserToAccessProfile,omitempty" xml:"allowUserToAccessProfile,omitempty"`
	AutoCreateUser           *bool   `yaml:"autoCreateUser,omitempty" xml:"autoCreateUser,omitempty"`
	Enabled                  *bool   `yaml:"enabled,omitempty" xml:"enabled,omitempty"`
}

// LdapGroupSetting represents the LDAP Group settings in Artifactory Security Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Security(Generalsecurity,PasswordPolicy,LDAP,SAML,OAuth,HTTPSSO,Crowd)
type LdapGroupSetting struct {
	Name                 *string `yaml:"-" xml:"name,omitempty"`
	DescriptionAttribute *string `yaml:"descriptionAttribute,omitempty" xml:"descriptionAttribute,omitempty"`
	EnabledLdap          *string `yaml:"enabledLdap,omitempty" xml:"enabledLdap,omitempty"`
	Filter               *string `yaml:"filter,omitempty" xml:"filter,omitempty"`
	GroupBaseDn          *string `yaml:"groupBaseDn,omitempty" xml:"groupBaseDn,omitempty"`
	GroupMemberAttribute *string `yaml:"groupMemberAttribute,omitempty" xml:"groupMemberAttribute,omitempty"`
	GroupNameAttribute   *string `yaml:"groupNameAttribute,omitempty" xml:"groupNameAttribute,omitempty"`
	Strategy             *string `yaml:"strategy,omitempty" xml:"strategy,omitempty"`
	SubTree              *bool   `yaml:"subtree,omitempty" xml:"subTree,omitempty"`
}

// AccessClientSettings represents the Access Client settings in Artifactory
// Configuration Descriptor. This is undocumented in YAML Configuration File.
//
type AccessClientSettings struct {
	ServerUrl                           *string `yaml:"serverUrl,omitempty" xml:"serverUrl,omitempty"`
	AdminToken                          *string `yaml:"adminToken,omitempty" xml:"adminToken,omitempty"`
	UserTokenMaxExpiresInMinutes        *int    `yaml:"userTokenMaxExpiresInMinutes,omitempty" xml:"userTokenMaxExpiresInMinutes,omitempty"`
	TokenVerifyResultCacheSize          *int    `yaml:"tokenVerifyResultCacheSize,omitempty" xml:"tokenVerifyResultCacheSize,omitempty"`
	TokenVerifyResultCacheExpirySeconds *int    `yaml:"tokenVerifyResultCacheExpirySeconds,omitempty" xml:"tokenVerifyResultCacheExpirySeconds,omitempty"`
}

// BuildGlobalBasicReadAllowed represents the Build Global Basic Read Information permission
// settings in Artifactory Security Configuration. This is undocumented in YAML Configuration File.
//
type BuildGlobalBasicReadAllowed struct {
	Enabled *bool `yaml:"enabled,omitempty" xml:"enabled,omitempty"`
}

// BuildGlobalBasicReadForAnonymous represents the Build Global Anonymous Read Information permission
// settings in Artifactory Security Configuration. This is undocumented in YAML Configuration File.
//
type BuildGlobalBasicReadForAnonymous struct {
	Enabled *bool `yaml:"enabled,omitempty" xml:"enabled,omitempty"`
}

// CrowdSettings represents the Crowd settings in Artifactory Security Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Security(Generalsecurity,PasswordPolicy,LDAP,SAML,OAuth,HTTPSSO,Crowd)
type CrowdSettings struct {
	ApplicationName           *string `yaml:"applicationName,omitempty" xml:"applicationName,omitempty"`
	Password                  *string `yaml:"password,omitempty" xml:"password,omitempty"`
	ServerUrl                 *string `yaml:"serverUrl,omitempty" xml:"serverUrl,omitempty"`
	SessionValidationInterval *int    `yaml:"sessionValidationInterval,omitempty" xml:"sessionValidationInterval,omitempty"`
	EnableIntegration         *bool   `yaml:"enableIntegration,omitempty" xml:"enableIntegration,omitempty"`
	NoAutoUserCreation        *bool   `yaml:"noAutoUserCreation,omitempty" xml:"noAutoUserCreation,omitempty"`
	UseDefaultProxy           *bool   `yaml:"useDefaultProxy,omitempty" xml:"useDefaultProxy,omitempty"`
}

// SamlSettings represents the SAML settings in Artifactory Security Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Security(Generalsecurity,PasswordPolicy,LDAP,SAML,OAuth,HTTPSSO,Crowd)
type SamlSettings struct {
	EnableIntegration        *bool   `yaml:"enableIntegration,omitempty" xml:"enableIntegration,omitempty"`
	Certificate              *string `yaml:"certificate,omitempty" xml:"certificate,omitempty"`
	EmailAttribute           *string `yaml:"emailAttribute,omitempty" xml:"emailAttribute,omitempty"`
	GroupAttribute           *string `yaml:"groupAttribute,omitempty" xml:"groupAttribute,omitempty"`
	LoginUrl                 *string `yaml:"loginUrl,omitempty" xml:"loginUrl,omitempty"`
	LogoutUrl                *string `yaml:"logoutUrl,omitempty" xml:"logoutUrl,omitempty"`
	NoAutoUserCreation       *bool   `yaml:"noAutoUserCreation,omitempty" xml:"noAutoUserCreation,omitempty"`
	ServiceProviderName      *string `yaml:"serviceProviderName,omitempty" xml:"serviceProviderName,omitempty"`
	AllowUserToAccessProfile *bool   `yaml:"allowUserToAccessProfile,omitempty" xml:"allowUserToAccessProfile,omitempty"`
	AutoRedirect             *bool   `yaml:"autoRedirect,omitempty" xml:"autoRedirect,omitempty"`
	SyncGroups               *bool   `yaml:"syncGroups,omitempty" xml:"syncGroups,omitempty"`
}

// OauthSettings represents the OAuth settings in Artifactory Security Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Security(Generalsecurity,PasswordPolicy,LDAP,SAML,OAuth,HTTPSSO,Crowd)
type OauthSettings struct {
	AllowUserToAccessProfile *bool                   `yaml:"allowUserToAccessProfile,omitempty" xml:"allowUserToAccessProfile,omitempty"`
	EnableIntegration        *bool                   `yaml:"enableIntegration,omitempty" xml:"enableIntegration,omitempty"`
	PersistUsers             *bool                   `yaml:"persistUsers,omitempty" xml:"persistUsers,omitempty"`
	OauthProvidersSettings   *[]OauthProviderSetting `yaml:"oauthProvidersSettings,omitempty" xml:"oauthProvidersSettings>oauthProvidersSettings,omitempty"`
}

// OauthProviderSetting represents the Oauth Provider settings in Artifactory Security Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Security(Generalsecurity,PasswordPolicy,LDAP,SAML,OAuth,HTTPSSO,Crowd)
type OauthProviderSetting struct {
	Name         *string `yaml:"-" xml:"name,omitempty"`
	Id           *string `yaml:"-" xml:"id,omitempty"`
	ApiUrl       *string `yaml:"apiUrl,omitempty" xml:"apiUrl,omitempty"`
	AuthUrl      *string `yaml:"authUrl,omitempty" xml:"authUrl,omitempty"`
	BasicUrl     *string `yaml:"basicUrl,omitempty" xml:"basicUrl,omitempty"`
	Enabled      *bool   `yaml:"enabled,omitempty" xml:"enabled,omitempty"`
	ProviderType *string `yaml:"providerType,omitempty" xml:"providerType,omitempty"`
	Secret       *string `yaml:"secret,omitempty" xml:"secret,omitempty"`
	TokenUrl     *string `yaml:"tokenUrl,omitempty" xml:"tokenUrl,omitempty"`
}

// HttpSsoSettings represents the HTTP SSO settings in Artifactory Security Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Security(Generalsecurity,PasswordPolicy,LDAP,SAML,OAuth,HTTPSSO,Crowd)
type HttpSsoSettings struct {
	HttpSsoProxied            *bool   `yaml:"httpSsoProxied,omitempty" xml:"httpSsoProxied,omitempty"`
	RemoteUserRequestVariable *string `yaml:"remoteUserRequestVariable,omitempty" xml:"remoteUserRequestVariable,omitempty"`
	AllowUserToAccessProfile  *bool   `yaml:"allowUserToAccessProfile,omitempty" xml:"allowUserToAccessProfile,omitempty"`
	NoAutoUserCreation        *bool   `yaml:"noAutoUserCreation,omitempty" xml:"noAutoUserCreation,omitempty"`
	SyncLdapGroups            *bool   `yaml:"syncLdapGroups,omitempty" xml:"syncLdapGroups,omitempty"`
}

// Backup represents the Backup settings in Artifactory Services Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Servicesconfiguration(Backups,MavenIndexer)
type Backup struct {
	Key                    *string   `yaml:"-" xml:"key,omitempty"`
	CronExp                *string   `yaml:"cronExp,omitempty" xml:"cronExp,omitempty"`
	ExcludedRepositories   *[]string `yaml:"excludedRepositories,omitempty" xml:"excludedRepositories>repositoryRef,omitempty"`
	RetentionPeriodHours   *int      `yaml:"retentionPeriodHours,omitempty" xml:"retentionPeriodHours,omitempty"`
	CreateArchive          *bool     `yaml:"createArchive,omitempty" xml:"createArchive,omitempty"`
	Enabled                *bool     `yaml:"enabled,omitempty" xml:"enabled,omitempty"`
	ExcludeBuilds          *bool     `yaml:"excludeBuilds,omitempty" xml:"excludeBuilds,omitempty"`
	ExcludeNewRepositories *bool     `yaml:"excludeNewRepositories,omitempty" xml:"excludeNewRepositories,omitempty"`
	SendMailOnError        *bool     `yaml:"sendMailOnError,omitempty" xml:"sendMailOnError,omitempty"`
	Precalculate           *bool     `yaml:"precalculate,omitempty" xml:"precalculate,omitempty"`
}

// Indexer represents the Maven Indexer settings in Artifactory Services Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Servicesconfiguration(Backups,MavenIndexer)
type Indexer struct {
	Enabled              *bool     `yaml:"enabled,omitempty" xml:"enabled,omitempty"`
	CronExp              *string   `yaml:"cronExp,omitempty" xml:"cronExp,omitempty"`
	IncludedRepositories *[]string `yaml:"includedRepositories,omitempty" xml:"includedRepositories,omitempty"`
}

// RepoLayout represents a Repository Layout setting in Artifactory.
// This is undocumented in YAML Configuration File.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Repository+Layouts
type RepoLayout struct {
	Name                             *string `yaml:"-" xml:"name,omitempty"`
	ArtifactPathPattern              *string `yaml:"artifactPathPattern,omitempty" xml:"artifactPathPattern,omitempty"`
	DistinctiveDescriptorPathPattern *bool   `yaml:"distinctiveDescriptorPathPattern,omitempty" xml:"distinctiveDescriptorPathPattern,omitempty"`
	DescriptorPathPattern            *string `yaml:"descriptorPathPattern,omitempty" xml:"descriptorPathPattern,omitempty"`
	FolderIntegrationRevisionRegExp  *string `yaml:"folderIntegrationRevisionRegExp,omitempty" xml:"folderIntegrationRevisionRegExp,omitempty"`
	FileIntegrationRevisionRegExp    *string `yaml:"fileIntegrationRevisionRegExp,omitempty" xml:"fileIntegrationRevisionRegExp,omitempty"`
}

// GcConfig represents the Garbage Collection settings in Artifactory Maintenance Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Servicesconfiguration(Backups,MavenIndexer)
type GcConfig struct {
	CronExp *string `yaml:"cronExp,omitempty" xml:"cronExp,omitempty"`
}

// CleanupConfig represents the Cleanup Unused Cached Artifacts setting in Artifactory Maintenance Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Servicesconfiguration(Backups,MavenIndexer)
type CleanupConfig struct {
	CronExp *string `yaml:"cronExp,omitempty" xml:"cronExp,omitempty"`
}

// VirtualCacheCleanupConfig represents the Cleanup Virtual Repositories
// setting in Artifactory Maintenance Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Servicesconfiguration(Backups,MavenIndexer)
type VirtualCacheCleanupConfig struct {
	CronExp *string `yaml:"cronExp,omitempty" xml:"cronExp,omitempty"`
}

// QuotaConfig represents the Storage Quota settings in Artifactory Maintenance Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Servicesconfiguration(Backups,MavenIndexer)
type QuotaConfig struct {
	DiskSpaceLimitPercentage   *int  `yaml:"diskSpaceLimitPercentage,omitempty" xml:"diskSpaceLimitPercentage,omitempty"`
	DiskSpaceWarningPercentage *int  `yaml:"diskSpaceWarningPercentage,omitempty" xml:"diskSpaceWarningPercentage,omitempty"`
	Enabled                    *bool `yaml:"enabled,omitempty" xml:"enabled,omitempty"`
}

// SystemMessageConfig represents Custom Message settings in Artifactory General Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-General(General,PropertySets,Proxy,Mail)
type SystemMessageConfig struct {
	Enabled        *bool   `yaml:"enabled,omitempty" xml:"enabled,omitempty"`
	Message        *string `yaml:"message,omitempty" xml:"message,omitempty"`
	Title          *string `yaml:"title,omitempty" xml:"title,omitempty"`
	TitleColor     *string `yaml:"titleColor,omitempty" xml:"titleColor,omitempty"`
	ShowOnAllPages *bool   `yaml:"showOnAllPages,omitempty" xml:"showOnAllPages,omitempty"`
}

// FolderDownloadConfig represents Folder Download settings in Artifactory General Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-General(General,PropertySets,Proxy,Mail)
type FolderDownloadConfig struct {
	Enabled               *bool `yaml:"enabled,omitempty" xml:"enabled,omitempty"`
	MaxConcurrentRequests *int  `yaml:"maxConcurrentRequests,omitempty" xml:"maxConcurrentRequests,omitempty"`
	MaxDownloadSizeMb     *int  `yaml:"maxDownloadSizeMb,omitempty" xml:"maxDownloadSizeMb,omitempty"`
	MaxFiles              *int  `yaml:"maxFiles,omitempty" xml:"maxFiles,omitempty"`
	EnabledForAnonymous   *bool `yaml:"enabledForAnonymous,omitempty" xml:"enabledForAnonymous,omitempty"`
}

// TrashcanConfig represents Trash Can settings in Artifactory General Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-General(General,PropertySets,Proxy,Mail)
type TrashcanConfig struct {
	Enabled             *bool `yaml:"enabled,omitempty" xml:"enabled,omitempty"`
	RetentionPeriodDays *int  `yaml:"retentionPeriodDays,omitempty" xml:"retentionPeriodDays,omitempty"`
}

// ReplicationsConfig represents Global Replication Blocking
// settings in Artifactory General Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-General(General,PropertySets,Proxy,Mail)
type ReplicationsConfig struct {
	BlockPullReplications *bool `yaml:"blockPullReplications,omitempty" xml:"blockPullReplications,omitempty"`
	BlockPushReplications *bool `yaml:"blockPushReplications,omitempty" xml:"blockPushReplications,omitempty"`
}

// BintrayApplication represents Bintray Oauth applications configuration.
// This is undocumented in YAML Configuration File.
//
type BintrayApplication struct {
	Key          *string `yaml:"-" xml:"key,omitempty"`
	ClientId     *string `yaml:"clientId" xml:"clientId,omitempty"`
	Secret       *string `yaml:"secret" xml:"secret,omitempty"`
	Org          *string `yaml:"org" xml:"org,omitempty"`
	Scope        *string `yaml:"scope" xml:"scope,omitempty"`
	RefreshToken *string `yaml:"refreshToken" xml:"refreshToken,omitempty"`
}

// SumoLogicConfig represents Sumo Logic settings in Artifactory's Configuration
// Descriptor. This is undocumented in YAML Configuration File.
//
type SumoLogicConfig struct {
	Enabled  *bool   `yaml:"enabled,omitempty" xml:"enabled,omitempty"`
	Proxy    *string `yaml:"proxy,omitempty" xml:"proxyRef,omitempty"`
	ClientId *string `yaml:"clientId,omitempty" xml:"clientId,omitempty"`
	Secret   *string `yaml:"secret,omitempty" xml:"secret,omitempty"`
}

// ReleaseBundlesConfig represents Release Bundle settings in Artifactory's
// Configuration Descriptor. This is undocumented in YAML Configuration File.
//
type ReleaseBundlesConfig struct {
	IncompleteCleanupPeriodHours *int `yaml:"incompleteCleanupPeriodHours,omitempty" xml:"incompleteCleanupPeriodHours,omitempty"`
}

// SignedUrlConfig represents Signed URL settings in Artifactory's Configuration
// Descriptor. This is undocumented in YAML Configuration File.
//
type SignedUrlConfig struct {
	MaxValidForSeconds *int `yaml:"maxValidForSeconds,omitempty" xml:"maxValidForSeconds,omitempty"`
}

// DownloadRedirectConfig represents Download Redirect settings in Artifactory's
// Configuration Descriptor. This is undocumented in YAML Configuration File.
//
type DownloadRedirectConfig struct {
	FileMinimumSize *int `yaml:"fileMinimumSize,omitempty" xml:"fileMinimumSize,omitempty"`
}

// Ping returns a simple status response.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-SystemHealthPing
func (s *SystemService) Ping() (*string, *Response, error) {
	u := "/api/system/ping"
	v := new(string)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// Get returns the general system information.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-SystemInfo
func (s *SystemService) Get() (*string, *Response, error) {
	u := "/api/system"
	v := new(string)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// GetVersionAndAddOns returns information about the current version, revision, and installed add-ons.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-VersionandAdd-onsinformation
func (s *SystemService) GetVersionAndAddOns() (*Versions, *Response, error) {
	u := "/api/system/version"
	v := new(Versions)

	resp, err := s.client.Call("GET", u, nil, v)
	return v, resp, err
}

// GetConfiguration returns the Global Artifactory Configuration Descriptor (artifactory.config.xml).
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-GeneralConfiguration
func (s *SystemService) GetConfiguration() (*GlobalConfigResponse, *Response, error) {
	u := "/api/system/configuration"
	v := new(bytes.Buffer)

	resp, err := s.client.Call("GET", u, nil, v)

	config := new(GlobalConfigResponse)
	err = xml.Unmarshal(v.Bytes(), config)
	if err != nil {
		return nil, resp, err
	}
	return config, resp, err
}

// UpdateConfiguration applies the provided Global system configuration to Artifactory.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-GeneralConfiguration
//       https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Advanced
func (s *SystemService) UpdateConfiguration(config GlobalConfigRequest) (*string, *Response, error) {
	u, err := s.client.buildURLForRequest("/api/system/configuration")
	if err != nil {
		return nil, nil, err
	}

	buf := new(bytes.Buffer)
	err = yaml.NewEncoder(buf).Encode(config)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest("PATCH", u, buf)
	if err != nil {
		return nil, nil, err
	}

	// Apply authentication
	if s.client.Authentication.HasAuth() {
		s.client.addAuthentication(req)
	}

	// Set Content-Type header for YAML
	req.Header.Add("Content-Type", "application/yaml")

	v := new(string)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return v, resp, err
	}

	return v, resp, err
}
