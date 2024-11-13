package definitions

const (
	Gaito SiteType = "gaito"
	Gaigu SiteType = "gaigu"
)

var SiteConfigs = map[SiteType]string{
	Gaito: "https://gaito.love",
	Gaigu: "https://gaigu31.tv",
}
