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
	"encoding/json"
	"github.com/franela/goblin"
	"github.com/gin-gonic/gin"
	"github.com/target/go-arty/artifactory/fixtures/system"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_System(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create http test server from our fake API handler
	s := httptest.NewServer(system.FakeHandler())

	// Create the client to interact with the http test server
	c, _ := NewClient(s.URL, nil)

	g := goblin.Goblin(t)
	g.Describe("System Service", func() {
		// Close http test server after we're done using it
		g.After(func() {
			s.Close()
		})

		g.Describe("System", func() {

			g.It("- should return valid string for Versions with String()", func() {
				actual := &Versions{
					Version:  String("5.9.5"),
					Revision: String("123456789"),
					Addons:   &[]string{"build", "ldap", "properties"},
				}

				data, _ := ioutil.ReadFile("fixtures/system/version.json")

				var expected Versions
				_ = json.Unmarshal(data, &expected)

				g.Assert(actual.String() == expected.String()).IsTrue()
			})

			g.It("- should return no error with Ping()", func() {
				actual, resp, err := c.System.Ping()
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with Get()", func() {
				actual, resp, err := c.System.Get()
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with GetVersionAndAddOns()", func() {
				actual, resp, err := c.System.GetVersionAndAddOns()
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

			g.It("- should return no error with GetConfiguration()", func() {
				actual, resp, err := c.System.GetConfiguration()
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()

				expected := &GlobalConfigResponse{
					GlobalConfigCommon: &GlobalConfigCommon{
						ServerName:          String("server1"),
						OfflineMode:         Bool(false),
						HelpLinksEnabled:    Bool(true),
						FileUploadMaxSizeMb: Int(100),
						DateFormat:          String("dd-MM-yy HH:mm:ss z"),
						AddonsConfig: &AddonsConfig{
							ShowAddonsInfo:       Bool(true),
							ShowAddonsInfoCookie: String("1574704307382"),
						},
						MailServer: &MailServer{
							Enabled:        Bool(true),
							ArtifactoryUrl: String("http://artifactory.domain.local:80"),
							From:           String("noreply@artifactory.domain.local"),
							Host:           String("mailHost"),
							Username:       String("user1"),
							Password:       String("password1"),
							Port:           Int(25),
							SubjectPrefix:  String("[Artifactory]"),
							Ssl:            Bool(true),
							Tls:            Bool(true),
						},
						XrayConfig: &XrayConfig{
							Enabled:                       Bool(false),
							BaseUrl:                       String("https://xray.domain.local"),
							User:                          String("ATF_xray"),
							Password:                      String("abcde-fghi-jklm-nopq-rstu-vwxyz!"),
							ArtifactoryId:                 String("ATF"),
							XrayId:                        String("0123456789abcdefghijklmnopqrstuvwxyz"),
							AllowDownloadsXrayUnavailable: Bool(true),
							AllowBlockedArtifactsDownload: Bool(true),
							BypassDefaultProxy:            Bool(true),
						},
						BintrayConfig: &BintrayConfig{
							FileUploadLimit: Int(0),
						},
						Indexer: &Indexer{
							Enabled:              Bool(false),
							CronExp:              String("0 23 5 * * ?"),
							IncludedRepositories: nil,
						},
						UrlBase: String("http://myhost.com/artifactory"),
						Logo:    String("http://fileserver/path/to/logo.png"),
						Footer:  String("custom text to appear in the page footer"),
						GcConfig: &GcConfig{
							CronExp: String("0 0 /4 * * ?"),
						},
						CleanupConfig: &CleanupConfig{
							CronExp: String("0 12 5 * * ?"),
						},
						VirtualCacheCleanupConfig: &VirtualCacheCleanupConfig{
							CronExp: String("0 12 0 * * ?"),
						},
						QuotaConfig: &QuotaConfig{
							DiskSpaceLimitPercentage:   Int(95),
							DiskSpaceWarningPercentage: Int(85),
							Enabled:                    Bool(true),
						},
						SystemMessageConfig: &SystemMessageConfig{
							Enabled:        Bool(true),
							Message:        String("Welcome to Artifactory"),
							Title:          String("Hello"),
							TitleColor:     String("#429F46"),
							ShowOnAllPages: Bool(false),
						},
						FolderDownloadConfig: &FolderDownloadConfig{
							Enabled:               Bool(false),
							MaxConcurrentRequests: Int(10),
							MaxDownloadSizeMb:     Int(1024),
							MaxFiles:              Int(5000),
							EnabledForAnonymous:   Bool(false),
						},
						TrashcanConfig: &TrashcanConfig{
							Enabled:             Bool(true),
							RetentionPeriodDays: Int(14),
						},
						ReplicationsConfig: &ReplicationsConfig{
							BlockPullReplications: Bool(false),
							BlockPushReplications: Bool(false),
						},
						SumoLogicConfig: &SumoLogicConfig{
							Enabled:  Bool(false),
							Proxy:    String("proxy1"),
							ClientId: String("abcdef"),
							Secret:   String("mysecret"),
						},
						ReleaseBundlesConfig: &ReleaseBundlesConfig{
							IncompleteCleanupPeriodHours: Int(720),
						},
						SignedUrlConfig: &SignedUrlConfig{
							MaxValidForSeconds: Int(31536000),
						},
					},
					Revision: Int(1311),
					Security: &struct {
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
					}{
						AnonAccessEnabled:         Bool(false),
						HideUnauthorizedResources: Bool(false),
						UserLockPolicy: &UserLockPolicy{
							Enabled:       Bool(false),
							LoginAttempts: Int(5),
						},
						PasswordSettings: &PasswordSettings{
							EncryptionPolicy: String("supported"),
							ExpirationPolicy: &ExpirationPolicy{
								Enabled:        Bool(false),
								PasswordMaxAge: Int(60),
								NotifyByEmail:  Bool(true),
							},
							ResetPolicy: &ResetPolicy{
								Enabled:               Bool(true),
								MaxAttemptsPerAddress: Int(3),
								TimeToBlockInMinutes:  Int(60),
							},
						},
						LdapSettings: &[]LdapSetting{
							{
								Key:                     String("ldap-setting-1"),
								EmailAttribute:          String("mail"),
								LdapPoisoningProtection: Bool(true),
								LdapUrl:                 String("ldap://ldap.domain.local"),
								Search: &struct {
									ManagerDn       *string `yaml:"managerDn,omitempty" xml:"managerDn,omitempty"`
									ManagerPassword *string `yaml:"managerPassword,omitempty" xml:"managerPassword,omitempty"`
									SearchBase      *string `yaml:"searchBase,omitempty" xml:"searchBase,omitempty"`
									SearchFilter    *string `yaml:"searchFilter,omitempty" xml:"searchFilter,omitempty"`
									SearchSubTree   *bool   `yaml:"searchSubTree,omitempty" xml:"searchSubTree,omitempty"`
								}{
									ManagerDn:       String("CN=ldap-user,OU=Services,OU=Root,DC=domain,DC=local"),
									ManagerPassword: String("password"),
									SearchBase:      String("ou=root,dc=domain,dc=local"),
									SearchFilter:    String("sAMAccountName={0}"),
									SearchSubTree:   Bool(true),
								},
								UserDnPattern:            String("uid={0},ou=People"),
								AllowUserToAccessProfile: Bool(false),
								AutoCreateUser:           Bool(true),
								Enabled:                  Bool(true),
							},
						},
						LdapGroupSettings: &[]LdapGroupSetting{
							{
								Name:                 String("ldap-group-setting-1"),
								DescriptionAttribute: String("description"),
								EnabledLdap:          String("ldap-setting-1"),
								Filter:               String("(objectClass=group)"),
								GroupBaseDn:          String("OU=Groups,OU=Root,DC=domain,DC=local"),
								GroupMemberAttribute: String("member:1.2.840.113556.1.4.1941"),
								GroupNameAttribute:   String("cn"),
								Strategy:             String("STATIC"),
								SubTree:              Bool(true),
							},
						},
						HttpSsoSettings: &HttpSsoSettings{
							HttpSsoProxied:            Bool(true),
							RemoteUserRequestVariable: String("REMOTE_USER"),
							AllowUserToAccessProfile:  Bool(true),
							NoAutoUserCreation:        Bool(false),
							SyncLdapGroups:            Bool(true),
						},
						CrowdSettings: &CrowdSettings{
							ApplicationName:           String("Artifactory"),
							Password:                  String("myPassword"),
							ServerUrl:                 String("http://crowd.domain.local"),
							SessionValidationInterval: Int(0),
							EnableIntegration:         Bool(true),
							NoAutoUserCreation:        Bool(false),
							UseDefaultProxy:           Bool(false),
						},
						SamlSettings: &SamlSettings{
							EnableIntegration:        Bool(true),
							Certificate:              String("cert-with-public-key"),
							EmailAttribute:           String("mail"),
							GroupAttribute:           String("group"),
							LoginUrl:                 String("saml.domain.local/login"),
							LogoutUrl:                String("saml.domain.local/logout"),
							NoAutoUserCreation:       Bool(false),
							ServiceProviderName:      String("providerId"),
							AllowUserToAccessProfile: Bool(false),
							AutoRedirect:             Bool(true),
							SyncGroups:               Bool(true),
						},
						OauthSettings: &OauthSettings{
							EnableIntegration:        Bool(true),
							AllowUserToAccessProfile: Bool(false),
							PersistUsers:             Bool(false),
							OauthProvidersSettings: &[]OauthProviderSetting{
								{
									Name:         String("test"),
									Enabled:      Bool(true),
									ProviderType: String("github"),
									Id:           String("test"),
									Secret:       String("secret1"),
									ApiUrl:       String("https://api.github.com/user"),
									AuthUrl:      String("https://github.com/login/oauth/authorize"),
									TokenUrl:     String("https://github.com/login/oauth/access_token"),
									BasicUrl:     String("https://github.com/"),
								},
							},
						},
						AccessClientSettings: &AccessClientSettings{
							ServerUrl:                           nil,
							AdminToken:                          String("admin-token"),
							UserTokenMaxExpiresInMinutes:        Int(60),
							TokenVerifyResultCacheSize:          Int(-1),
							TokenVerifyResultCacheExpirySeconds: Int(-1),
						},
						BuildGlobalBasicReadAllowed:      Bool(false),
						BuildGlobalBasicReadForAnonymous: Bool(false),
					},
					Backups: &[]Backup{
						{
							Key:                    String("backup-daily"),
							Enabled:                Bool(true),
							CronExp:                String("0 0 2 ? * MON-FRI"),
							RetentionPeriodHours:   Int(0),
							CreateArchive:          Bool(false),
							ExcludedRepositories:   &[]string{"example-repo-local"},
							SendMailOnError:        Bool(true),
							ExcludeNewRepositories: Bool(false),
							Precalculate:           Bool(false),
						},
						{
							Key:                    String("backup-weekly"),
							Enabled:                Bool(false),
							CronExp:                String("0 0 2 ? * SAT"),
							RetentionPeriodHours:   Int(336),
							CreateArchive:          Bool(false),
							ExcludedRepositories:   &[]string{"example-repo-local"},
							SendMailOnError:        Bool(true),
							ExcludeNewRepositories: Bool(false),
							Precalculate:           Bool(false),
						},
					},
					LocalRepositories: &[]LocalRepository{
						{
							GenericRepository: &GenericRepository{
								Key:                          String("example-repo-local"),
								PackageType:                  String("generic"),
								Description:                  nil,
								Notes:                        nil,
								IncludesPattern:              String("**/*"),
								ExcludesPattern:              nil,
								LayoutRef:                    String("simple-default"),
								HandleReleases:               Bool(true),
								HandleSnapshots:              Bool(true),
								MaxUniqueSnapshots:           Int(0),
								SuppressPomConsistencyChecks: Bool(true),
								BlackedOut:                   Bool(false),
								PropertySets:                 &[]string{"artifactory"},
							},
							DockerAPIVersion:        String("V2"),
							DebianTrivialLayout:     Bool(false),
							MaxUniqueTags:           Int(0),
							SnapshotVersionBehavior: String("unique"),
							XrayIndex:               Bool(false),
							BlockPushingSchema1:     Bool(true),
							ChecksumPolicyType:      String("client-checksums"),
							CalculateYumMetadata:    Bool(false),
							YumRootDepth:            Int(0),
							EnableFileListsIndexing: Bool(false),
							ArchiveBrowsingEnabled:  Bool(false),
						},
					},
					RemoteRepositories: &[]RemoteRepository{
						{
							GenericRepository: &GenericRepository{
								Key:                          String("docker-remote"),
								PackageType:                  String("docker"),
								Description:                  nil,
								Notes:                        nil,
								IncludesPattern:              String("**/*"),
								ExcludesPattern:              nil,
								LayoutRef:                    String("simple-default"),
								HandleReleases:               Bool(true),
								HandleSnapshots:              Bool(true),
								MaxUniqueSnapshots:           Int(0),
								SuppressPomConsistencyChecks: Bool(true),
								BlackedOut:                   Bool(false),
								PropertySets:                 &[]string{"artifactory"},
							},
							URL:                               String("https://registry-1.docker.io/"),
							Offline:                           Bool(false),
							HardFail:                          Bool(false),
							StoreArtifactsLocally:             Bool(true),
							FetchJarsEagerly:                  Bool(false),
							FetchSourcesEagerly:               Bool(false),
							ExternalDependenciesEnabled:       Bool(true),
							ExternalDependenciesPatterns:      &[]string{"**"},
							RetrievalCachePeriodSecs:          Int(7200),
							AssumedOfflinePeriodSecs:          Int(300),
							MissedRetrievalCachePeriodSecs:    Int(1800),
							RemoteRepoChecksumPolicyType:      String("generate-if-absent"),
							UnusedArtifactsCleanupPeriodHours: Int(0),
							ShareConfiguration:                Bool(false),
							SynchronizeProperties:             Bool(false),
							ListRemoteFolderItems:             Bool(false),
							ContentSynchronisation: &ContentSynchronisation{
								Enabled: Bool(false),
								Properties: &struct {
									Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
								}{
									Enabled: Bool(false),
								},
								Statistics: &struct {
									Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
								}{
									Enabled: Bool(false),
								},
								Source: &struct {
									OriginAbsenceDetection *bool `json:"originAbsenceDetection,omitempty" xml:"originAbsenceDetection,omitempty"`
								}{
									OriginAbsenceDetection: Bool(false),
								},
							},
							BlockPushingSchema1:       Bool(true),
							BlockMismatchingMimeTypes: Bool(true),
							BypassHeadRequests:        Bool(false),
							AllowAnyHostAuth:          Bool(false),
							SocketTimeoutMillis:       Int(15000),
							EnableCookieManagement:    Bool(false),
							EnableTokenAuthentication: Bool(true),
						},
					},
					VirtualRepositories: &[]VirtualRepository{
						{
							GenericRepository: &GenericRepository{
								Key:                          String("generic-virtual"),
								PackageType:                  String("generic"),
								Description:                  nil,
								Notes:                        nil,
								IncludesPattern:              String("**/*"),
								ExcludesPattern:              nil,
								LayoutRef:                    String("simple-default"),
								HandleReleases:               Bool(true),
								HandleSnapshots:              Bool(true),
								MaxUniqueSnapshots:           Int(0),
								SuppressPomConsistencyChecks: Bool(true),
								BlackedOut:                   Bool(false),
								PropertySets:                 &[]string{"artifactory"},
							},
							ArtifactoryRequestsCanRetrieveRemoteArtifacts: Bool(true),
							ResolveDockerTagsByTimestamp:                  Bool(false),
							Repositories:                                  &[]string{"example-repo-local", "generic-remote"},
							PomRepositoryReferencesCleanupPolicy:          String("discard_active_reference"),
							DefaultDeploymentRepo:                         String("example-repo-local"),
							ForceMavenAuthentication:                      Bool(false),
						},
					},
					LocalReplications: &[]Replication{
						{
							Username:                        String("admin"),
							Password:                        String("password1"),
							Url:                             String("https://artifactory-2.domain.local/generic-art1"),
							SocketTimeoutMillis:             Int(15000),
							CronExp:                         String("0 0 12 * * ?"),
							RepoKey:                         String("generic-art1"),
							EnableEventReplication:          Bool(false),
							Enabled:                         Bool(true),
							SyncDeletes:                     Bool(true),
							SyncProperties:                  Bool(true),
							SyncStatistics:                  Bool(true),
							PathPrefix:                      nil,
							CheckBinaryExistenceInFilestore: Bool(false),
						},
					},
					RemoteReplications: &[]Replication{
						{
							CronExp:                         String("0 0 12 * * ?"),
							RepoKey:                         String("generic-art2"),
							EnableEventReplication:          Bool(false),
							Enabled:                         Bool(true),
							SyncDeletes:                     Bool(true),
							SyncProperties:                  Bool(true),
							PathPrefix:                      nil,
							CheckBinaryExistenceInFilestore: Bool(false),
						},
					},
					Proxies: &[]Proxy{
						{
							Key:             String("proxy1"),
							Domain:          String("domain.local"),
							Host:            String("proxy1Host"),
							NtHost:          String("testNtHost"),
							Username:        String("user1"),
							Password:        String("password"),
							Port:            Int(8080),
							RedirectToHosts: String("host1,host2"),
							DefaultProxy:    Bool(false),
						},
					},
					ReverseProxies: &[]ReverseProxy{
						{
							Key:                      String("apache"),
							WebServerType:            String("apache"),
							ArtifactoryAppContext:    String("artifactory"),
							PublicAppContext:         nil,
							ServerName:               String("artifactory.domain.local"),
							ServerNameExpression:     String("*.artifactory.domain.local"),
							SslCertificate:           String(""),
							SslKey:                   String(""),
							DockerReverseProxyMethod: String("subDomain"),
							UseHttps:                 Bool(true),
							UseHttp:                  Bool(true),
							SslPort:                  Int(443),
							HttpPort:                 Int(80),
							ArtifactoryServerName:    String("art1.domain.local"),
							UpStreamName:             String("artifactory"),
							ArtifactoryPort:          Int(8081),
						},
					},
					PropertySets: &[]PropertySetResponse{
						{
							Name: String("artifactory"),
							Properties: &[]struct {
								Name             *string `xml:"name,omitempty"`
								PredefinedValues *[]struct {
									Value        *string `xml:"value,omitempty"`
									DefaultValue *bool   `xml:"defaultValue,omitempty"`
								} `xml:"predefinedValues>predefinedValue,omitempty"`
								ClosedPredefinedValues *bool `xml:"closedPredefinedValues,omitempty"`
								MultipleChoice         *bool `xml:"multipleChoice,omitempty"`
							}{
								{
									Name: String("licenses"),
									PredefinedValues: &[]struct {
										Value        *string `xml:"value,omitempty"`
										DefaultValue *bool   `xml:"defaultValue,omitempty"`
									}{
										{
											Value:        String("AFL-3.0"),
											DefaultValue: Bool(false),
										},
										{
											Value:        String("AGPL-V3"),
											DefaultValue: Bool(false),
										},
									},
									ClosedPredefinedValues: Bool(true),
									MultipleChoice:         Bool(true),
								},
							},
							Visible: Bool(false),
						},
					},
					RepoLayouts: &[]RepoLayout{
						{
							Name:                             String("maven-2-default"),
							ArtifactPathPattern:              String("[orgPath]/[module]/[baseRev](-[folderItegRev])/[module]-[baseRev](-[fileItegRev])(-[classifier]).[ext]"),
							DistinctiveDescriptorPathPattern: Bool(true),
							DescriptorPathPattern:            String("[orgPath]/[module]/[baseRev](-[folderItegRev])/[module]-[baseRev](-[fileItegRev])(-[classifier]).pom"),
							FolderIntegrationRevisionRegExp:  String("SNAPSHOT"),
							FileIntegrationRevisionRegExp:    String("SNAPSHOT|(?:(?:[0-9]{8}.[0-9]{6})-(?:[0-9]+))"),
						},
						{
							Name:                             String("maven-1-default"),
							ArtifactPathPattern:              String("[org]/[type]s/[module]-[baseRev](-[fileItegRev])(-[classifier]).[ext]"),
							DistinctiveDescriptorPathPattern: Bool(true),
							DescriptorPathPattern:            String("[org]/[type]s/[module]-[baseRev](-[fileItegRev]).pom"),
							FolderIntegrationRevisionRegExp:  String(".+"),
							FileIntegrationRevisionRegExp:    String(".+"),
						},
					},
					BintrayApplications: &[]BintrayApplication{
						{
							Key:          String("app1"),
							ClientId:     String("testClientId"),
							Secret:       String("testSecret"),
							Org:          String("testOrg"),
							Scope:        String("testScope"),
							RefreshToken: String("testRefreshToken"),
						},
					},
				}

				g.Assert(actual).Equal(expected)

			})

			g.It("- should return no error with UpdateConfiguration()", func() {
				config := GlobalConfigRequest{
					GlobalConfigCommon: GlobalConfigCommon{
						ServerName: String("server1"),
					},
					Security: nil,
					Backups: &map[string]Backup{
						"backup-daily": {
							Key:                    String("backup-daily"),
							Enabled:                Bool(true),
							CronExp:                String("0 0 2 ? * MON-FRI"),
							RetentionPeriodHours:   Int(0),
							CreateArchive:          Bool(false),
							ExcludedRepositories:   &[]string{"example-repo-local"},
							SendMailOnError:        Bool(true),
							ExcludeNewRepositories: Bool(false),
							Precalculate:           Bool(false),
						},
						"backup-weekly": {
							Key:                    String("backup-weekly"),
							Enabled:                Bool(false),
							CronExp:                String("0 0 2 ? * SAT"),
							RetentionPeriodHours:   Int(336),
							CreateArchive:          Bool(false),
							ExcludedRepositories:   &[]string{"example-repo-local"},
							SendMailOnError:        Bool(true),
							ExcludeNewRepositories: Bool(false),
							Precalculate:           Bool(false),
						},
					},
					Proxies:             nil,
					ReverseProxies:      nil,
					PropertySets:        nil,
					RepoLayouts:         nil,
					BintrayApplications: nil,
				}
				actual, resp, err := c.System.UpdateConfiguration(config)
				g.Assert(actual != nil).IsTrue()
				g.Assert(resp != nil).IsTrue()
				g.Assert(err == nil).IsTrue()
			})

		})

	})

}
