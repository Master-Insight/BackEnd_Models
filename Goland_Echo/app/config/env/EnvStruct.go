package env

// Env estructura variable de entorno
type ConfigEnv struct {
	Port                 string
	CorsOrigin           string
	MongoURI             string
	JWTSecretCode        string
	AdminUsers           []string
	AdminPassword        string
	GmailUserName        string
	GmailUserApp         string
	GmailPassApp         string
	LinkedInClientID     string
	LinkedInClientSecret string
	LinkedInRedirectURI  string
	LinkedInScope        string
	CloudinaryName       string
	CloudinaryKey        string
	CloudinarySecret     string
	CloudinaryURL        string
}
