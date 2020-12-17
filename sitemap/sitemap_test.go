package sitemap

import (
	"testing"
	"time"
)

func TestIndex(t *testing.T) {

	lm := time.Date(2020, 1, 1, 1, 1, 1, 0, time.UTC)

	sm := NewSiteMapIndex()
	sm.AddSitemap("/sm1", lm)
	sm.AddSitemap("/sm2", lm)

	s, err := sm.String()
	if err != nil {
		t.Error(err)
	}
	if s != `<?xml version="1.0" encoding="UTF-8"?>`+"\n"+`<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"><sitemap><loc>/sm1</loc><lastmod>2020-01-01T01:01:01Z</lastmod></sitemap><sitemap><loc>/sm2</loc><lastmod>2020-01-01T01:01:01Z</lastmod></sitemap></sitemapindex>` {
		t.Error("invalid output")
	}
}

func TestSitemap(t *testing.T) {

	lm := time.Date(2020, 1, 1, 1, 1, 1, 0, time.UTC)

	index := NewSitemap()
	index.AddLocation("loc1", lm, FrequencyAlways, 1)
	index.AddLocation("loc2", lm, FrequencyHourly, .9)
	index.AddLocation("loc3", lm, FrequencyDaily, .8)
	index.AddLocation("loc4", lm, FrequencyWeekly, .7)
	index.AddLocation("loc5", lm, FrequencyMonthly, .6)
	index.AddLocation("loc6", lm, FrequencyYearly, .5)
	index.AddLocation("loc7", lm, FrequencyNever, .4)

	s, err := index.String()
	if err != nil {
		t.Error(err)
	}
	if s != `<?xml version="1.0" encoding="UTF-8"?>`+"\n"+`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"><url><loc>loc1</loc><lastmod>2020-01-01T01:01:01Z</lastmod><changefreq>always</changefreq><priority>1</priority></url><url><loc>loc2</loc><lastmod>2020-01-01T01:01:01Z</lastmod><changefreq>hourly</changefreq><priority>0.9</priority></url><url><loc>loc3</loc><lastmod>2020-01-01T01:01:01Z</lastmod><changefreq>daily</changefreq><priority>0.8</priority></url><url><loc>loc4</loc><lastmod>2020-01-01T01:01:01Z</lastmod><changefreq>weekly</changefreq><priority>0.7</priority></url><url><loc>loc5</loc><lastmod>2020-01-01T01:01:01Z</lastmod><changefreq>monthly</changefreq><priority>0.6</priority></url><url><loc>loc6</loc><lastmod>2020-01-01T01:01:01Z</lastmod><changefreq>yearly</changefreq><priority>0.5</priority></url><url><loc>loc7</loc><lastmod>2020-01-01T01:01:01Z</lastmod><changefreq>never</changefreq><priority>0.4</priority></url></urlset>` {
		t.Error("invalid output")
	}
}
