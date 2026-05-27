package sitemap_go

import (
	"bytes"
	"testing"
	"time"
)

func TestIndex(t *testing.T) {

	lm := time.Date(2020, 1, 1, 1, 1, 1, 0, time.UTC)

	sm := NewIndex()
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

	index := New()
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
	expected := `<?xml version="1.0" encoding="UTF-8"?>` + "\n" + `<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:image="http://www.google.com/schemas/sitemap-image/1.1" xmlns:video="http://www.google.com/schemas/sitemap-video/1.1"><url><loc>loc1</loc><lastmod>2020-01-01T01:01:01Z</lastmod><changefreq>always</changefreq><priority>1</priority></url><url><loc>loc2</loc><lastmod>2020-01-01T01:01:01Z</lastmod><changefreq>hourly</changefreq><priority>0.9</priority></url><url><loc>loc3</loc><lastmod>2020-01-01T01:01:01Z</lastmod><changefreq>daily</changefreq><priority>0.8</priority></url><url><loc>loc4</loc><lastmod>2020-01-01T01:01:01Z</lastmod><changefreq>weekly</changefreq><priority>0.7</priority></url><url><loc>loc5</loc><lastmod>2020-01-01T01:01:01Z</lastmod><changefreq>monthly</changefreq><priority>0.6</priority></url><url><loc>loc6</loc><lastmod>2020-01-01T01:01:01Z</lastmod><changefreq>yearly</changefreq><priority>0.5</priority></url><url><loc>loc7</loc><lastmod>2020-01-01T01:01:01Z</lastmod><changefreq>never</changefreq><priority>0.4</priority></url></urlset>`
	if s != expected {
		t.Errorf("expected %s, got %s", expected, s)
	}
}

func TestWriteTo(t *testing.T) {

	lm := time.Date(2020, 1, 1, 1, 1, 1, 0, time.UTC)

	sm := New()
	sm.AddLocation("loc1", lm, FrequencyAlways, 1)

	var buf bytes.Buffer
	_, err := sm.WriteTo(&buf)
	if err != nil {
		t.Error(err)
	}

	expected := `<?xml version="1.0" encoding="UTF-8"?>` + "\n" + `<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:image="http://www.google.com/schemas/sitemap-image/1.1" xmlns:video="http://www.google.com/schemas/sitemap-video/1.1"><url><loc>loc1</loc><lastmod>2020-01-01T01:01:01Z</lastmod><changefreq>always</changefreq><priority>1</priority></url></urlset>`
	if buf.String() != expected {
		t.Errorf("expected %s, got %s", expected, buf.String())
	}
}

func TestExtensions(t *testing.T) {

	lm := time.Date(2020, 1, 1, 1, 1, 1, 0, time.UTC)

	sm := New()
	loc := sm.AddLocation("loc1", lm, FrequencyAlways, 1)

	loc.AddImage(Image{
		Location: "http://example.com/image.jpg",
		Title:    "Image Title",
	})

	loc.AddVideo(Video{
		ThumbnailLocation: "http://example.com/thumb.jpg",
		Title:             "Video Title",
		Description:       "Video Description",
	})

	s, err := sm.String()
	if err != nil {
		t.Error(err)
	}

	expected := `<?xml version="1.0" encoding="UTF-8"?>` + "\n" + `<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:image="http://www.google.com/schemas/sitemap-image/1.1" xmlns:video="http://www.google.com/schemas/sitemap-video/1.1"><url><loc>loc1</loc><lastmod>2020-01-01T01:01:01Z</lastmod><changefreq>always</changefreq><priority>1</priority><image:image><image:loc>http://example.com/image.jpg</image:loc><image:title>Image Title</image:title></image:image><video:video><video:thumbnail_loc>http://example.com/thumb.jpg</video:thumbnail_loc><video:title>Video Title</video:title><video:description>Video Description</video:description></video:video></url></urlset>`
	if s != expected {
		t.Errorf("expected %s, got %s", expected, s)
	}
}
