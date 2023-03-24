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

// GlobalConfig represents elements of the Global Configuration Descriptor.
// Lots of these aren't documented but have been mapped from the
// XML schema at https://www.jfrog.com/public/xsd/artifactory-v2_2_5.xsd
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File
type GlobalConfig struct {
	Revision                  *int                       `yaml:"-" xml:"revision,omitempty"`
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
	DownloadRedirectConfig    *DownloadRedirectConfig    `yaml:"downloadRedirectConfig,omitempty" xml:"downloadRedirectConfig,omitempty"`
	Security                  *Security                  `yaml:"security,omitempty" xml:"security,omitempty"`
	Backups                   *Backups                   `yaml:"backups,omitempty" xml:"backups>backup,omitempty"`
	Proxies                   *Proxies                   `yaml:"proxies,omitempty" xml:"proxies>proxy,omitempty"`
	ReverseProxies            *ReverseProxies            `yaml:"reverseProxies,omitempty" xml:"reverseProxies>reverseProxy,omitempty"`
	PropertySets              *PropertySets              `yaml:"propertySets,omitempty" xml:"propertySets>propertySet,omitempty"`
	RepoLayouts               *RepoLayouts               `yaml:"repoLayouts,omitempty" xml:"repoLayouts>repoLayout,omitempty"`
	BintrayApplications       *BintrayApplications       `yaml:"bintrayApplications,omitempty" xml:"bintrayApplications>bintrayApplication,omitempty"`
	LocalRepositories         *[]LocalRepository         `yaml:"-" xml:"localRepositories>localRepository,omitempty"`
	RemoteRepositories        *[]RemoteRepository        `yaml:"-" xml:"remoteRepositories>remoteRepository,omitempty"`
	VirtualRepositories       *[]VirtualRepository       `yaml:"-" xml:"virtualRepositories>virtualRepository,omitempty"`
	LocalReplications         *[]Replication             `yaml:"-" xml:"localReplications>localReplication,omitempty"`
	RemoteReplications        *[]Replication             `yaml:"-" xml:"remoteReplications>remoteReplication,omitempty"`
}

// AddonsConfig represents Addons-related configuration.
// This is undocumented in YAML Configuration File.
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
	Reset          *bool   `yaml:"-" xml:"-"`
}

// MarshalYAML implements the Marshaller interface.
func (m MailServer) MarshalYAML() (interface{}, error) {
	if m.Reset != nil && *m.Reset {
		return nil, nil
	}
	return m, nil
}

// XrayConfig represents Xray related settings in Artifactory's Configuration
// Descriptor. This is undocumented in YAML Configuration File.
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
type BintrayConfig struct {
	UserName        *string `yaml:"userName,omitempty" xml:"userName,omitempty"`
	ApiKey          *string `yaml:"apiKey,omitempty" xml:"apiKey,omitempty"`
	FileUploadLimit *int    `yaml:"fileUploadLimit,omitempty" xml:"fileUploadLimit,omitempty"`
}

// Proxy represents a Proxy setting in Artifactory's Global Configuration Descriptor.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-General(General,PropertySets,Proxy,Mail)
type Proxy struct {
	Key             *string `yaml:"-" xml:"key,omitempty"`
	Domain          *string `yaml:"domain,omitempty" xml:"domain,omitempty"`
	Host            *string `yaml:"host,omitempty" xml:"host,omitempty"`
	NtHost          *string `yaml:"ntHost,omitempty" xml:"ntHost,omitempty"`
	Password        *string `yaml:"password,omitempty" xml:"password,omitempty"`
	Port            *int    `yaml:"port,omitempty" xml:"port,omitempty"`
	RedirectToHosts *string `yaml:"redirectedToHosts,omitempty" xml:"redirectedToHosts,omitempty"`
	Username        *string `yaml:"username,omitempty" xml:"username,omitempty"`
	DefaultProxy    *bool   `yaml:"defaultProxy,omitempty" xml:"defaultProxy,omitempty"`
}

// Proxies is an alias for a slice of Proxy that can be
// properly marshaled to/from YAML.
type Proxies []*Proxy

// MarshalYAML implements the yaml.Marshaller interface for Proxies.
func (p Proxies) MarshalYAML() (interface{}, error) {
	proxiesMap := make(map[string]*Proxy)
	for _, proxy := range p {
		if *proxy == (Proxy{Key: proxy.Key}) {
			proxiesMap[*proxy.Key] = nil
		} else {
			proxiesMap[*proxy.Key] = proxy
		}

	}

	return proxiesMap, nil
}

// UnmarshalYAML implements yaml.Unmarshaler for Proxies.
func (p *Proxies) UnmarshalYAML(unmarshal func(interface{}) error) error {
	proxiesMap := make(map[*string]*Proxy)
	if err := unmarshal(proxiesMap); err != nil {
		return err
	}

	var proxiesSlice Proxies
	for proxyKey, proxy := range proxiesMap {
		if proxy == nil {
			continue
		}
		proxy.Key = proxyKey
		proxiesSlice = append(proxiesSlice, proxy)
	}

	*p = proxiesSlice
	return nil
}

// ReverseProxy represents a Reverse Proxy configuration in Artifactory's Global Configuration Descriptor.
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

// ReverseProxies is an alias for a slice of ReverseProxy that can be
// properly marshaled to/from YAML.
type ReverseProxies []*ReverseProxy

// MarshalYAML implements the yaml.Marshaller interface for ReverseProxies.
func (r ReverseProxies) MarshalYAML() (interface{}, error) {
	reverseProxiesMap := make(map[string]*ReverseProxy)
	for _, reverseProxy := range r {
		if *reverseProxy == (ReverseProxy{Key: reverseProxy.Key}) {
			reverseProxiesMap[*reverseProxy.Key] = nil
		} else {
			reverseProxiesMap[*reverseProxy.Key] = reverseProxy
		}
	}

	return reverseProxiesMap, nil
}

// UnmarshalYAML implements yaml.Unmarshaler for ReverseProxies.
func (r *ReverseProxies) UnmarshalYAML(unmarshal func(interface{}) error) error {
	reverseProxiesMap := make(map[*string]*ReverseProxy)
	if err := unmarshal(reverseProxiesMap); err != nil {
		return err
	}

	var reverseProxiesSlice ReverseProxies
	for reverseProxyKey, reverseProxy := range reverseProxiesMap {
		if reverseProxy == nil {
			continue
		}
		reverseProxy.Key = reverseProxyKey
		reverseProxiesSlice = append(reverseProxiesSlice, reverseProxy)
	}

	*r = reverseProxiesSlice
	return nil
}

// PropertySet represents a Property Set in Artifactory's Global Configuration Descriptor.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-General(General,PropertySets,Proxy,Mail)
type PropertySet struct {
	Name       *string     `yaml:"-" xml:"name,omitempty"`
	Properties *Properties `yaml:"properties,omitempty" xml:"properties>property,omitempty"`
	Visible    *bool       `yaml:"visible,omitempty" xml:"visible,omitempty"`
}

// Properties is an alias for a slice of Property in a PropertySet that can be
// properly marshaled to/from YAML.
type Properties []*Property

// MarshalYAML implements the yaml.Marshaller interface for Properties.
func (p Properties) MarshalYAML() (interface{}, error) {
	propertiesMap := make(map[string]*Property)
	for _, property := range p {
		if *property == (Property{Name: property.Name}) {
			propertiesMap[*property.Name] = nil
		} else {
			propertiesMap[*property.Name] = property
		}
	}

	return propertiesMap, nil
}

// UnmarshalYAML implements yaml.Unmarshaler for Properties.
func (p *Properties) UnmarshalYAML(unmarshal func(interface{}) error) error {
	propertiesMap := make(map[*string]*Property)
	if err := unmarshal(propertiesMap); err != nil {
		return err
	}

	var propertiesSlice Properties
	for propertyName, property := range propertiesMap {
		if property == nil {
			continue
		}
		property.Name = propertyName
		propertiesSlice = append(propertiesSlice, property)
	}

	*p = propertiesSlice
	return nil
}

type Property struct {
	Name                   *string           `yaml:"-" xml:"name,omitempty"`
	PredefinedValues       *PredefinedValues `yaml:"predefinedValues,omitempty" xml:"predefinedValues>predefinedValue,omitempty"`
	ClosedPredefinedValues *bool             `yaml:"closedPredefinedValues,omitempty" xml:"closedPredefinedValues,omitempty"`
	MultipleChoice         *bool             `yaml:"multipleChoice,omitempty" xml:"multipleChoice,omitempty"`
}

type PredefinedValue struct {
	Value        *string `yaml:"-" xml:"value,omitempty"`
	DefaultValue *bool   `yaml:"defaultValue,omitempty" xml:"defaultValue,omitempty"`
}

// PredefinedValues is an alias for a slice of PredefinedValue in a
// PropertySet's Property that can be properly marshaled to/from YAML.
type PredefinedValues []*PredefinedValue

// MarshalYAML implements the yaml.Marshaller interface for PredefinedValues.
func (p PredefinedValues) MarshalYAML() (interface{}, error) {
	predefinedValuesMap := make(map[string]*PredefinedValue)
	for _, predefinedValue := range p {
		if *predefinedValue == (PredefinedValue{Value: predefinedValue.Value}) {
			predefinedValuesMap[*predefinedValue.Value] = nil
		} else {
			predefinedValuesMap[*predefinedValue.Value] = predefinedValue
		}
	}

	return predefinedValuesMap, nil
}

// UnmarshalYAML implements yaml.Unmarshaler for PredefinedValues.
func (p *PredefinedValues) UnmarshalYAML(unmarshal func(interface{}) error) error {
	predefinedValuesMap := make(map[*string]*PredefinedValue)
	if err := unmarshal(predefinedValuesMap); err != nil {
		return err
	}

	var predefinedValuesSlice PredefinedValues
	for predefinedValueName, predefinedValue := range predefinedValuesMap {
		if predefinedValue == nil {
			continue
		}
		predefinedValue.Value = predefinedValueName
		predefinedValuesSlice = append(predefinedValuesSlice, predefinedValue)
	}

	*p = predefinedValuesSlice
	return nil
}

// PropertySets is an alias for a slice of PropertySet that can be
// properly marshaled to/from YAML.
type PropertySets []*PropertySet

// MarshalYAML implements the yaml.Marshaller interface for PropertySets.
func (p PropertySets) MarshalYAML() (interface{}, error) {
	propertySetsMap := make(map[string]*PropertySet)
	for _, propertySet := range p {
		if *propertySet == (PropertySet{Name: propertySet.Name}) {
			propertySetsMap[*propertySet.Name] = nil
		} else {
			propertySetsMap[*propertySet.Name] = propertySet
		}
	}

	return propertySetsMap, nil
}

// UnmarshalYAML implements yaml.Unmarshaler for PropertySets.
func (p *PropertySets) UnmarshalYAML(unmarshal func(interface{}) error) error {
	propertySetsMap := make(map[*string]*PropertySet)
	if err := unmarshal(propertySetsMap); err != nil {
		return err
	}

	var propertySetsSlice PropertySets
	for propertySetName, propertySet := range propertySetsMap {
		if propertySet == nil {
			continue
		}
		propertySet.Name = propertySetName
		propertySetsSlice = append(propertySetsSlice, propertySet)
	}

	*p = propertySetsSlice
	return nil
}

type Security struct {
	AnonAccessEnabled                *bool                 `yaml:"anonAccessEnabled,omitempty" xml:"anonAccessEnabled,omitempty"`
	HideUnauthorizedResources        *bool                 `yaml:"hideUnauthorizedResources,omitempty" xml:"hideUnauthorizedResources,omitempty"`
	UserLockPolicy                   *UserLockPolicy       `yaml:"userLockPolicy,omitempty" xml:"userLockPolicy,omitempty"`
	PasswordSettings                 *PasswordSettings     `yaml:"passwordSettings,omitempty" xml:"passwordSettings,omitempty"`
	LdapSettings                     *LdapSettings         `yaml:"ldapSettings,omitempty" xml:"ldapSettings>ldapSetting,omitempty"`
	LdapGroupSettings                *LdapGroupSettings    `yaml:"ldapGroupSettings,omitempty" xml:"ldapGroupSettings>ldapGroupSetting,omitempty"`
	HttpSsoSettings                  *HttpSsoSettings      `yaml:"httpSsoSettings,omitempty" xml:"httpSsoSettings,omitempty"`
	CrowdSettings                    *CrowdSettings        `yaml:"crowdSettings,omitempty" xml:"crowdSettings,omitempty"`
	SamlSettings                     *SamlSettings         `yaml:"samlSettings,omitempty" xml:"samlSettings,omitempty"`
	OauthSettings                    *OauthSettings        `yaml:"oauthSettings,omitempty" xml:"oauthSettings,omitempty"`
	AccessClientSettings             *AccessClientSettings `yaml:"accessClientSettings,omitempty" xml:"accessClientSettings,omitempty"`
	BuildGlobalBasicReadAllowed      *bool                 `yaml:"buildGlobalBasicReadAllowed,omitempty" xml:"buildGlobalBasicReadAllowed,omitempty"`
	BuildGlobalBasicReadForAnonymous *bool                 `yaml:"buildGlobalBasicReadForAnonymous,omitempty" xml:"buildGlobalBasicReadForAnonymous,omitempty"`
}

// PasswordSettings represents the Password settings in Artifactory Security Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Security(Generalsecurity,PasswordPolicy,LDAP,SAML,OAuth,HTTPSSO,Crowd)
type PasswordSettings struct {
	EncryptionPolicy *string           `yaml:"encryptionPolicy,omitempty" xml:"encryptionPolicy,omitempty"`
	ExpirationPolicy *ExpirationPolicy `yaml:"expirationPolicy,omitempty" xml:"expirationPolicy,omitempty"`
	ResetPolicy      *ResetPolicy      `yaml:"resetPolicy,omitempty" xml:"resetPolicy,omitempty"`
	Reset            *bool             `yaml:"-" xml:"-"`
}

// MarshalYAML implements the Marshaller interface.
func (p PasswordSettings) MarshalYAML() (interface{}, error) {
	if p.Reset != nil && *p.Reset {
		return nil, nil
	}
	return p, nil
}

// ExpirationPolicy represents the Password Expiration Policy settings in Artifactory Security Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Security(Generalsecurity,PasswordPolicy,LDAP,SAML,OAuth,HTTPSSO,Crowd)
type ExpirationPolicy struct {
	Enabled        *bool `yaml:"enabled,omitempty" xml:"enabled,omitempty"`
	PasswordMaxAge *int  `yaml:"passwordMaxAge,omitempty" xml:"passwordMaxAge,omitempty"`
	NotifyByEmail  *bool `yaml:"notifyByEmail,omitempty" xml:"notifyByEmail,omitempty"`
	Reset          *bool `yaml:"-" xml:"-"`
}

// MarshalYAML implements the Marshaller interface.
func (e ExpirationPolicy) MarshalYAML() (interface{}, error) {
	if e.Reset != nil && *e.Reset {
		return nil, nil
	}
	return e, nil
}

// ResetPolicy represents the Password Reset Protection policy settings in Artifactory Security Configuration.
// This is undocumented in YAML Configuration File.
type ResetPolicy struct {
	Enabled               *bool `yaml:"enabled,omitempty" xml:"enabled,omitempty"`
	MaxAttemptsPerAddress *int  `yaml:"maxAttemptsPerAddress,omitempty" xml:"maxAttemptsPerAddress,omitempty"`
	TimeToBlockInMinutes  *int  `yaml:"timeToBlockInMinutes,omitempty" xml:"timeToBlockInMinutes,omitempty"`
	Reset                 *bool `yaml:"-" xml:"-"`
}

func (r ResetPolicy) MarshalYAML() (interface{}, error) {
	if r.Reset != nil && *r.Reset {
		return nil, nil
	}
	return r, nil
}

// UserLockPolicy represents the User Lock Policy settings in Artifactory Security Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Security(Generalsecurity,PasswordPolicy,LDAP,SAML,OAuth,HTTPSSO,Crowd)
type UserLockPolicy struct {
	Enabled       *bool `yaml:"enabled,omitempty" xml:"enabled,omitempty"`
	LoginAttempts *int  `yaml:"loginAttempts,omitempty" xml:"loginAttempts,omitempty"`
	Reset         *bool `yaml:"-" xml:"-"`
}

// MarshalYAML implements the Marshaller interface.
func (u UserLockPolicy) MarshalYAML() (interface{}, error) {
	if u.Reset != nil && *u.Reset {
		return nil, nil
	}
	return u, nil
}

// SigningKeysSetting represents the GPG Signing settings in Artifactory Security Configuration.
// This is undocumented in YAML Configuration File.
type SigningKeysSettings struct {
	Passphrase       *string `yaml:"passphrase,omitempty" xml:"passphrase,omitempty"`
	KeyStorePassword *string `yaml:"keyStorePassword,omitempty" xml:"keyStorePassword,omitempty"`
	Reset            *bool   `yaml:"-" xml:"-"`
}

// MarshalYAML implements the Marshaller interface.
func (s SigningKeysSettings) MarshalYAML() (interface{}, error) {
	if s.Reset != nil && *s.Reset {
		return nil, nil
	}
	return s, nil
}

// LdapSetting represents the LDAP settings in Artifactory Security Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Security(Generalsecurity,PasswordPolicy,LDAP,SAML,OAuth,HTTPSSO,Crowd)
type LdapSetting struct {
	Key                      *string            `yaml:"-" xml:"key,omitempty"`
	EmailAttribute           *string            `yaml:"emailAttribute,omitempty" xml:"emailAttribute,omitempty"`
	LdapPoisoningProtection  *bool              `yaml:"ldapPoisoningProtection,omitempty" xml:"ldapPoisoningProtection,omitempty"`
	LdapUrl                  *string            `yaml:"ldapUrl,omitempty" xml:"ldapUrl,omitempty"`
	Search                   *LdapSettingSearch `yaml:"search,omitempty" xml:"search,omitempty"`
	UserDnPattern            *string            `yaml:"userDnPattern,omitempty" xml:"userDnPattern,omitempty"`
	AllowUserToAccessProfile *bool              `yaml:"allowUserToAccessProfile,omitempty" xml:"allowUserToAccessProfile,omitempty"`
	AutoCreateUser           *bool              `yaml:"autoCreateUser,omitempty" xml:"autoCreateUser,omitempty"`
	Enabled                  *bool              `yaml:"enabled,omitempty" xml:"enabled,omitempty"`
}

// LdapSettingSearch represents the Search setting in an LDAPSetting.
type LdapSettingSearch struct {
	ManagerDn       *string `yaml:"managerDn,omitempty" xml:"managerDn,omitempty"`
	ManagerPassword *string `yaml:"managerPassword,omitempty" xml:"managerPassword,omitempty"`
	SearchBase      *string `yaml:"searchBase,omitempty" xml:"searchBase,omitempty"`
	SearchFilter    *string `yaml:"searchFilter,omitempty" xml:"searchFilter,omitempty"`
	SearchSubTree   *bool   `yaml:"searchSubTree,omitempty" xml:"searchSubTree,omitempty"`
}

// LdapSettings is an alias for a slice of LdapSetting that can be
// properly marshaled to/from YAML.
type LdapSettings []*LdapSetting

// MarshalYAML implements the yaml.Marshaller interface for LdapSettings.
func (l LdapSettings) MarshalYAML() (interface{}, error) {
	ldapSettingsMap := make(map[string]*LdapSetting)
	for _, ldapSetting := range l {
		if *ldapSetting == (LdapSetting{Key: ldapSetting.Key}) {
			ldapSettingsMap[*ldapSetting.Key] = nil
		} else {
			ldapSettingsMap[*ldapSetting.Key] = ldapSetting
		}
	}

	return ldapSettingsMap, nil
}

// UnmarshalYAML implements yaml.Unmarshaler for LdapSettings.
func (l *LdapSettings) UnmarshalYAML(unmarshal func(interface{}) error) error {
	ldapSettingsMap := make(map[*string]*LdapSetting)
	if err := unmarshal(ldapSettingsMap); err != nil {
		return err
	}

	var ldapSettingsSlice LdapSettings
	for ldapSettingKey, ldapSetting := range ldapSettingsMap {
		if ldapSetting == nil {
			continue
		}
		ldapSetting.Key = ldapSettingKey
		ldapSettingsSlice = append(ldapSettingsSlice, ldapSetting)
	}

	*l = ldapSettingsSlice
	return nil
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
	SubTree              *bool   `yaml:"subTree,omitempty" xml:"subTree,omitempty"`
}

// LdapGroupSettings is an alias for a slice of LdapGroupSetting that can be
// properly marshaled to/from YAML.
type LdapGroupSettings []*LdapGroupSetting

// MarshalYAML implements the yaml.Marshaller interface for LdapGroupSettings.
func (l LdapGroupSettings) MarshalYAML() (interface{}, error) {
	ldapGroupSettingsMap := make(map[string]*LdapGroupSetting)
	for _, ldapGroupSetting := range l {
		if *ldapGroupSetting == (LdapGroupSetting{Name: ldapGroupSetting.Name}) {
			ldapGroupSettingsMap[*ldapGroupSetting.Name] = nil
		} else {
			ldapGroupSettingsMap[*ldapGroupSetting.Name] = ldapGroupSetting
		}
	}

	return ldapGroupSettingsMap, nil
}

// UnmarshalYAML implements yaml.Unmarshaler for LdapGroupSettings.
func (l *LdapGroupSettings) UnmarshalYAML(unmarshal func(interface{}) error) error {
	ldapGroupSettingsMap := make(map[*string]*LdapGroupSetting)
	if err := unmarshal(ldapGroupSettingsMap); err != nil {
		return err
	}

	var ldapGroupSettingsSlice LdapGroupSettings
	for ldapGroupSettingName, ldapGroupSetting := range ldapGroupSettingsMap {
		if ldapGroupSetting == nil {
			continue
		}
		ldapGroupSetting.Name = ldapGroupSettingName
		ldapGroupSettingsSlice = append(ldapGroupSettingsSlice, ldapGroupSetting)
	}

	*l = ldapGroupSettingsSlice
	return nil
}

// AccessClientSettings represents the Access Client settings in Artifactory
// Configuration Descriptor. This is undocumented in YAML Configuration File.
type AccessClientSettings struct {
	ServerUrl                           *string `yaml:"serverUrl,omitempty" xml:"serverUrl,omitempty"`
	AdminToken                          *string `yaml:"adminToken,omitempty" xml:"adminToken,omitempty"`
	UserTokenMaxExpiresInMinutes        *int    `yaml:"userTokenMaxExpiresInMinutes,omitempty" xml:"userTokenMaxExpiresInMinutes,omitempty"`
	TokenVerifyResultCacheSize          *int    `yaml:"tokenVerifyResultCacheSize,omitempty" xml:"tokenVerifyResultCacheSize,omitempty"`
	TokenVerifyResultCacheExpirySeconds *int    `yaml:"tokenVerifyResultCacheExpirySeconds,omitempty" xml:"tokenVerifyResultCacheExpirySeconds,omitempty"`
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
	Reset                     *bool   `yaml:"-" xml:"-"`
}

// MarshalYAML implements the Marshaller interface.
func (c CrowdSettings) MarshalYAML() (interface{}, error) {
	if c.Reset != nil && *c.Reset {
		return nil, nil
	}
	return c, nil
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
	Reset                    *bool   `yaml:"-" xml:"-"`
}

// MarshalYAML implements the Marshaller interface.
func (s SamlSettings) MarshalYAML() (interface{}, error) {
	if s.Reset != nil && *s.Reset {
		return nil, nil
	}
	return s, nil
}

// OauthSettings represents the OAuth settings in Artifactory Security Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Security(Generalsecurity,PasswordPolicy,LDAP,SAML,OAuth,HTTPSSO,Crowd)
type OauthSettings struct {
	AllowUserToAccessProfile *bool                  `yaml:"allowUserToAccessProfile,omitempty" xml:"allowUserToAccessProfile,omitempty"`
	EnableIntegration        *bool                  `yaml:"enableIntegration,omitempty" xml:"enableIntegration,omitempty"`
	PersistUsers             *bool                  `yaml:"persistUsers,omitempty" xml:"persistUsers,omitempty"`
	OauthProvidersSettings   *OauthProviderSettings `yaml:"oauthProvidersSettings,omitempty" xml:"oauthProvidersSettings>oauthProvidersSettings,omitempty"`
	Reset                    *bool                  `yaml:"-" xml:"-"`
}

// MarshalYAML implements the Marshaller interface.
func (o OauthSettings) MarshalYAML() (interface{}, error) {
	if o.Reset != nil && *o.Reset {
		return nil, nil
	}
	return o, nil
}

// OauthProviderSettings is an alias for a slice of OauthProviderSetting that can be
// properly marshaled to/from YAML.
type OauthProviderSettings []*OauthProviderSetting

// MarshalYAML implements the yaml.Marshaller interface for OauthProviderSettings.
func (o OauthProviderSettings) MarshalYAML() (interface{}, error) {
	oauthProviderSettingsMap := make(map[string]*OauthProviderSetting)
	for _, oauthProviderSetting := range o {
		if *oauthProviderSetting == (OauthProviderSetting{Name: oauthProviderSetting.Name}) {
			oauthProviderSettingsMap[*oauthProviderSetting.Name] = nil
		} else {
			oauthProviderSettingsMap[*oauthProviderSetting.Name] = oauthProviderSetting
		}
	}

	return oauthProviderSettingsMap, nil
}

// UnmarshalYAML implements yaml.Unmarshaler for OauthProviderSettings.
func (o *OauthProviderSettings) UnmarshalYAML(unmarshal func(interface{}) error) error {
	oauthProviderSettingsMap := make(map[*string]*OauthProviderSetting)
	if err := unmarshal(oauthProviderSettingsMap); err != nil {
		return err
	}

	var oauthProviderSettingsSlice OauthProviderSettings
	for oauthProviderSettingName, oauthProviderSetting := range oauthProviderSettingsMap {
		if oauthProviderSetting == nil {
			continue
		}
		oauthProviderSetting.Name = oauthProviderSettingName
		oauthProviderSettingsSlice = append(oauthProviderSettingsSlice, oauthProviderSetting)
	}

	*o = oauthProviderSettingsSlice
	return nil
}

// OauthProviderSetting represents the Oauth Provider settings in Artifactory Security Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Security(Generalsecurity,PasswordPolicy,LDAP,SAML,OAuth,HTTPSSO,Crowd)
type OauthProviderSetting struct {
	Name         *string `yaml:"-" xml:"name,omitempty"`
	Id           *string `yaml:"id,omitempty" xml:"id,omitempty"`
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
	Reset                     *bool   `yaml:"-" xml:"-"`
}

// MarshalYAML implements the Marshaller interface.
func (h HttpSsoSettings) MarshalYAML() (interface{}, error) {
	if h.Reset != nil && *h.Reset {
		return nil, nil
	}
	return h, nil
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

// Backups is an alias for a slice of Backup that can be
// properly marshaled to/from YAML.
type Backups []*Backup

// MarshalYAML implements the yaml.Marshaller interface for Backups.
func (b Backups) MarshalYAML() (interface{}, error) {
	backupsMap := make(map[string]*Backup)
	for _, backup := range b {
		if *backup == (Backup{Key: backup.Key}) {
			backupsMap[*backup.Key] = nil
		} else {
			backupsMap[*backup.Key] = backup
		}
	}

	return backupsMap, nil
}

// UnmarshalYAML implements yaml.Unmarshaler for Backups.
func (b *Backups) UnmarshalYAML(unmarshal func(interface{}) error) error {
	backupsMap := make(map[*string]*Backup)
	if err := unmarshal(backupsMap); err != nil {
		return err
	}

	var backupsSlice Backups
	for backupKey, backupConfig := range backupsMap {
		if backupConfig == nil {
			continue
		}
		backupConfig.Key = backupKey
		backupsSlice = append(backupsSlice, backupConfig)
	}

	*b = backupsSlice
	return nil
}

// Indexer represents the Maven Indexer settings in Artifactory Services Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Servicesconfiguration(Backups,MavenIndexer)
type Indexer struct {
	Enabled              *bool     `yaml:"enabled,omitempty" xml:"enabled,omitempty"`
	CronExp              *string   `yaml:"cronExp,omitempty" xml:"cronExp,omitempty"`
	IncludedRepositories *[]string `yaml:"includedRepositories,omitempty" xml:"includedRepositories>repositoryRef,omitempty"`
	Reset                *bool     `yaml:"-" xml:"-"`
}

// MarshalYAML implements the Marshaller interface.
func (i Indexer) MarshalYAML() (interface{}, error) {
	if i.Reset != nil && *i.Reset {
		return nil, nil
	}
	return i, nil
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

// RepoLayouts is an alias for a slice of RepoLayout that can be
// properly marshaled to/from YAML.
type RepoLayouts []*RepoLayout

// MarshalYAML implements the yaml.Marshaller interface for RepoLayouts.
func (r RepoLayouts) MarshalYAML() (interface{}, error) {
	repoLayoutsMap := make(map[string]*RepoLayout)
	for _, repoLayout := range r {
		if *repoLayout == (RepoLayout{Name: repoLayout.Name}) {
			repoLayoutsMap[*repoLayout.Name] = nil
		} else {
			repoLayoutsMap[*repoLayout.Name] = repoLayout
		}
	}

	return repoLayoutsMap, nil
}

// UnmarshalYAML implements yaml.Unmarshaler for RepoLayouts.
func (r *RepoLayouts) UnmarshalYAML(unmarshal func(interface{}) error) error {
	repoLayoutsMap := make(map[*string]*RepoLayout)
	if err := unmarshal(repoLayoutsMap); err != nil {
		return err
	}

	var repoLayoutsSlice RepoLayouts
	for repoLayoutName, repoLayout := range repoLayoutsMap {
		if repoLayout == nil {
			continue
		}
		repoLayout.Name = repoLayoutName
		repoLayoutsSlice = append(repoLayoutsSlice, repoLayout)
	}

	*r = repoLayoutsSlice
	return nil
}

// GcConfig represents the Garbage Collection settings in Artifactory Maintenance Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Servicesconfiguration(Backups,MavenIndexer)
type GcConfig struct {
	CronExp *string `yaml:"cronExp,omitempty" xml:"cronExp,omitempty"`
	Reset   *bool   `yaml:"-" xml:"-"`
}

// MarshalYAML implements the Marshaller interface.
func (g GcConfig) MarshalYAML() (interface{}, error) {
	if g.Reset != nil && *g.Reset {
		return nil, nil
	}
	return g, nil
}

// CleanupConfig represents the Cleanup Unused Cached Artifacts setting in Artifactory Maintenance Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Servicesconfiguration(Backups,MavenIndexer)
type CleanupConfig struct {
	CronExp *string `yaml:"cronExp,omitempty" xml:"cronExp,omitempty"`
	Reset   *bool   `yaml:"-" xml:"-"`
}

// MarshalYAML implements the Marshaller interface.
func (c CleanupConfig) MarshalYAML() (interface{}, error) {
	if c.Reset != nil && *c.Reset {
		return nil, nil
	}
	return c, nil
}

// VirtualCacheCleanupConfig represents the Cleanup Virtual Repositories
// setting in Artifactory Maintenance Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Servicesconfiguration(Backups,MavenIndexer)
type VirtualCacheCleanupConfig struct {
	CronExp *string `yaml:"cronExp,omitempty" xml:"cronExp,omitempty"`
	Reset   *bool   `yaml:"-" xml:"-"`
}

// MarshalYAML implements the Marshaller interface.
func (v VirtualCacheCleanupConfig) MarshalYAML() (interface{}, error) {
	if v.Reset != nil && *v.Reset {
		return nil, nil
	}
	return v, nil
}

// QuotaConfig represents the Storage Quota settings in Artifactory Maintenance Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Servicesconfiguration(Backups,MavenIndexer)
type QuotaConfig struct {
	DiskSpaceLimitPercentage   *int  `yaml:"diskSpaceLimitPercentage,omitempty" xml:"diskSpaceLimitPercentage,omitempty"`
	DiskSpaceWarningPercentage *int  `yaml:"diskSpaceWarningPercentage,omitempty" xml:"diskSpaceWarningPercentage,omitempty"`
	Enabled                    *bool `yaml:"enabled,omitempty" xml:"enabled,omitempty"`
	Reset                      *bool `yaml:"-" xml:"-"`
}

// MarshalYAML implements the Marshaller interface.
func (q QuotaConfig) MarshalYAML() (interface{}, error) {
	if q.Reset != nil && *q.Reset {
		return nil, nil
	}
	return q, nil
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
	Reset          *bool   `yaml:"-" xml:"-"`
}

// MarshalYAML implements the Marshaller interface.
func (s SystemMessageConfig) MarshalYAML() (interface{}, error) {
	if s.Reset != nil && *s.Reset {
		return nil, nil
	}
	return s, nil
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
	Reset                 *bool `yaml:"-" xml:"-"`
}

// MarshalYAML implements the Marshaller interface.
func (f FolderDownloadConfig) MarshalYAML() (interface{}, error) {
	if f.Reset != nil && *f.Reset {
		return nil, nil
	}
	return f, nil
}

// TrashcanConfig represents Trash Can settings in Artifactory General Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-General(General,PropertySets,Proxy,Mail)
type TrashcanConfig struct {
	Enabled             *bool `yaml:"enabled,omitempty" xml:"enabled,omitempty"`
	RetentionPeriodDays *int  `yaml:"retentionPeriodDays,omitempty" xml:"retentionPeriodDays,omitempty"`
	AllowPermDeletes    *bool `yaml:"allowPermDeletes,omitempty" xml:"allowPermDeletes,omitempty"`
	Reset               *bool `yaml:"-" xml:"-"`
}

// MarshalYAML implements the Marshaller interface.
func (t TrashcanConfig) MarshalYAML() (interface{}, error) {
	if t.Reset != nil && *t.Reset {
		return nil, nil
	}
	return t, nil
}

// ReplicationsConfig represents Global Replication Blocking
// settings in Artifactory General Configuration.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-General(General,PropertySets,Proxy,Mail)
type ReplicationsConfig struct {
	BlockPullReplications *bool `yaml:"blockPullReplications,omitempty" xml:"blockPullReplications,omitempty"`
	BlockPushReplications *bool `yaml:"blockPushReplications,omitempty" xml:"blockPushReplications,omitempty"`
	Reset                 *bool `yaml:"-" xml:"-"`
}

// MarshalYAML implements the Marshaller interface.
func (r ReplicationsConfig) MarshalYAML() (interface{}, error) {
	if r.Reset != nil && *r.Reset {
		return nil, nil
	}
	return r, nil
}

// BintrayApplication represents Bintray Oauth applications configuration.
// This is undocumented in YAML Configuration File.
type BintrayApplication struct {
	Key          *string `yaml:"-" xml:"key,omitempty"`
	ClientId     *string `yaml:"clientId" xml:"clientId,omitempty"`
	Secret       *string `yaml:"secret" xml:"secret,omitempty"`
	Org          *string `yaml:"org" xml:"org,omitempty"`
	Scope        *string `yaml:"scope" xml:"scope,omitempty"`
	RefreshToken *string `yaml:"refreshToken" xml:"refreshToken,omitempty"`
}

// BintrayApplications is an alias for a slice of BintrayApplication that can be
// properly marshaled to/from YAML.
type BintrayApplications []*BintrayApplication

// MarshalYAML implements the yaml.Marshaller interface for RepoLayouts.
func (b BintrayApplications) MarshalYAML() (interface{}, error) {
	bintrayApplicationsMap := make(map[string]*BintrayApplication)
	for _, bintrayApplication := range b {
		if *bintrayApplication == (BintrayApplication{Key: bintrayApplication.Key}) {
			bintrayApplicationsMap[*bintrayApplication.Key] = nil
		} else {
			bintrayApplicationsMap[*bintrayApplication.Key] = bintrayApplication
		}
	}

	return bintrayApplicationsMap, nil
}

// UnmarshalYAML implements yaml.Unmarshaler for RepoLayouts.
func (b *BintrayApplications) UnmarshalYAML(unmarshal func(interface{}) error) error {
	bintrayApplicationsMap := make(map[*string]*BintrayApplication)
	if err := unmarshal(bintrayApplicationsMap); err != nil {
		return err
	}

	var bintrayApplicationsSlice BintrayApplications
	for bintrayApplicationKey, bintrayApplication := range bintrayApplicationsMap {
		if bintrayApplication == nil {
			continue
		}
		bintrayApplication.Key = bintrayApplicationKey
		bintrayApplicationsSlice = append(bintrayApplicationsSlice, bintrayApplication)
	}

	*b = bintrayApplicationsSlice
	return nil
}

// SumoLogicConfig represents Sumo Logic settings in Artifactory's Configuration
// Descriptor. This is undocumented in YAML Configuration File.
type SumoLogicConfig struct {
	Enabled  *bool   `yaml:"enabled,omitempty" xml:"enabled,omitempty"`
	Proxy    *string `yaml:"proxy,omitempty" xml:"proxyRef,omitempty"`
	ClientId *string `yaml:"clientId,omitempty" xml:"clientId,omitempty"`
	Secret   *string `yaml:"secret,omitempty" xml:"secret,omitempty"`
	Reset    *bool   `yaml:"-" xml:"-"`
}

// MarshalYAML implements the Marshaller interface.
func (s SumoLogicConfig) MarshalYAML() (interface{}, error) {
	if s.Reset != nil && *s.Reset {
		return nil, nil
	}
	return s, nil
}

// ReleaseBundlesConfig represents Release Bundle settings in Artifactory's
// Configuration Descriptor. This is undocumented in YAML Configuration File.
type ReleaseBundlesConfig struct {
	IncompleteCleanupPeriodHours *int  `yaml:"incompleteCleanupPeriodHours,omitempty" xml:"incompleteCleanupPeriodHours,omitempty"`
	Reset                        *bool `yaml:"-" xml:"-"`
}

// MarshalYAML implements the Marshaller interface.
func (r ReleaseBundlesConfig) MarshalYAML() (interface{}, error) {
	if r.Reset != nil && *r.Reset {
		return nil, nil
	}
	return r, nil
}

// SignedUrlConfig represents Signed URL settings in Artifactory's Configuration
// Descriptor. This is undocumented in YAML Configuration File.
type SignedUrlConfig struct {
	MaxValidForSeconds *int  `yaml:"maxValidForSeconds,omitempty" xml:"maxValidForSeconds,omitempty"`
	Reset              *bool `yaml:"-" xml:"-"`
}

// MarshalYAML implements the Marshaller interface.
func (s SignedUrlConfig) MarshalYAML() (interface{}, error) {
	if s.Reset != nil && *s.Reset {
		return nil, nil
	}
	return s, nil
}

// DownloadRedirectConfig represents Download Redirect settings in Artifactory's
// Configuration Descriptor. This is undocumented in YAML Configuration File.
type DownloadRedirectConfig struct {
	FileMinimumSize *int  `yaml:"fileMinimumSize,omitempty" xml:"fileMinimumSize,omitempty"`
	Reset           *bool `yaml:"-" xml:"-"`
}

// MarshalYAML implements the Marshaller interface.
func (d DownloadRedirectConfig) MarshalYAML() (interface{}, error) {
	if d.Reset != nil && *d.Reset {
		return nil, nil
	}
	return d, nil
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
func (s *SystemService) GetConfiguration() (*GlobalConfig, *Response, error) {
	u := "/api/system/configuration"
	v := new(bytes.Buffer)

	resp, err := s.client.Call("GET", u, nil, v)

	config := new(GlobalConfig)
	err = xml.Unmarshal(v.Bytes(), config)
	if err != nil {
		return nil, resp, err
	}
	return config, resp, err
}

// UpdateConfiguration applies the provided Global system configuration to Artifactory.
//
// Docs: https://www.jfrog.com/confluence/display/RTF/Artifactory+REST+API#ArtifactoryRESTAPI-GeneralConfiguration
//
//	https://www.jfrog.com/confluence/display/RTF/YAML+Configuration+File#YAMLConfigurationFile-Advanced
func (s *SystemService) UpdateConfiguration(config GlobalConfig) (*string, *Response, error) {
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

	v := new(bytes.Buffer)
	resp, err := s.client.Do(req, v)
	return String(v.String()), resp, err
}
