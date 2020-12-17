package sitemap

const namespace = "http://www.sitemaps.org/schemas/sitemap/0.9"

type changeFrequency string

const (
	FrequencyAlways  changeFrequency = "always"
	FrequencyHourly  changeFrequency = "hourly"
	FrequencyDaily   changeFrequency = "daily"
	FrequencyWeekly  changeFrequency = "weekly"
	FrequencyMonthly changeFrequency = "monthly"
	FrequencyYearly  changeFrequency = "yearly"
	FrequencyNever   changeFrequency = "never"
)
