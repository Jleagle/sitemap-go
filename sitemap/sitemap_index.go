package sitemap

import (
	"encoding/xml"
	"time"
)

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

func (smi sitemapIndex) String() (string, error) {

	b, err := xml.Marshal(smi)
	if err != nil {
		return "", err
	}

	return xml.Header + string(b), nil
}

type Sitemap struct {
	Location     string `xml:"loc"`
	LastModified string `xml:"lastmod,omitempty"`
}

func (sm *Sitemap) SetModified(t time.Time) {
	if !t.IsZero() {
		sm.LastModified = t.Format(time.RFC3339)
	}
}
