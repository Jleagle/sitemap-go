package sitemap

import (
	"encoding/xml"
	"time"
)

const (
	namespace = "http://www.sitemaps.org/schemas/sitemap/0.9"
	iso8601   = time.RFC3339
)

type ChangeFrequency string

const (
	frequencyAlways  ChangeFrequency = "always"
	frequencyHourly  ChangeFrequency = "hourly"
	frequencyDaily   ChangeFrequency = "daily"
	frequencyWeekly  ChangeFrequency = "weekly"
	frequencyMonthly ChangeFrequency = "monthly"
	frequencyYearly  ChangeFrequency = "yearly"
	frequencyNever   ChangeFrequency = "never"
)

// Sitemap index
func NewSiteMapIndex() *sitemapIndex {

	return &sitemapIndex{
		Namespace: namespace,
	}
}

type sitemapIndex struct {
	XMLName   xml.Name  `xml:"sitemapindex"`
	Namespace string    `xml:"xmlns,attr"`
	SiteMaps  []Sitemap `xml:"sitemap"`
}

func (smi *sitemapIndex) AddSitemap(location string, lastModified time.Time) {
	sm := Sitemap{Location: location}
	sm.SetModified(lastModified)
	smi.SiteMaps = append(smi.SiteMaps, sm)
}

func (smi sitemapIndex) Bytes() ([]byte, error) {
	bytes, err := xml.Marshal(smi)
	return []byte(xml.Header + string(bytes)), err
}

type Sitemap struct {
	Location     string `xml:"loc"`
	LastModified string `xml:"lastmod,omitempty"`
}

func (sm *Sitemap) SetModified(time time.Time) {
	if !time.IsZero() {
		sm.LastModified = time.Format(iso8601)
	}
}

// Sitemap
func NewSitemap() *urlSet {

	return &urlSet{
		Namespace: namespace,
	}
}

type urlSet struct {
	XMLName   xml.Name   `xml:"urlset"`
	Namespace string     `xml:"xmlns,attr"`
	Locations []Location `xml:"url"`
}

func (set *urlSet) AddLocation(location string, lastModified time.Time, changeFrequency ChangeFrequency, priority float32) {
	loc := Location{Location: location, ChangeFrequency: changeFrequency, Priority: priority}
	loc.SetModified(lastModified)
	set.Locations = append(set.Locations, loc)
}

func (set urlSet) Bytes() ([]byte, error) {
	bytes, err := xml.Marshal(set)
	return []byte(xml.Header + string(bytes)), err
}

type Location struct {
	Location        string          `xml:"loc"`
	LastModified    string          `xml:"lastmod,omitempty"`
	ChangeFrequency ChangeFrequency `xml:"changefreq,omitempty"`
	Priority        float32         `xml:"priority,omitempty"`
}

func (sm *Location) SetModified(time time.Time) {
	if !time.IsZero() {
		sm.LastModified = time.Format(iso8601)
	}
}
