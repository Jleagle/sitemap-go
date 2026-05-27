package sitemap_go

const (
	namespace      = "http://www.sitemaps.org/schemas/sitemap/0.9"
	imageNamespace = "http://www.google.com/schemas/sitemap-image/1.1"
	videoNamespace = "http://www.google.com/schemas/sitemap-video/1.1"
)

// ChangeFrequency represents the change frequency of a page.
type ChangeFrequency string

const (
	FrequencyAlways  ChangeFrequency = "always"
	FrequencyHourly  ChangeFrequency = "hourly"
	FrequencyDaily   ChangeFrequency = "daily"
	FrequencyWeekly  ChangeFrequency = "weekly"
	FrequencyMonthly ChangeFrequency = "monthly"
	FrequencyYearly  ChangeFrequency = "yearly"
	FrequencyNever   ChangeFrequency = "never"
)
