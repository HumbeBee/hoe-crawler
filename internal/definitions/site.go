package definitions

type SiteType string

type SiteConfig struct {
	BaseURL string
}

type SiteConfigMap map[SiteType]SiteConfig
