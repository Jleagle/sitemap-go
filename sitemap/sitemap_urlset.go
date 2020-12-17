package sitemap

import (
	"encoding/xml"
	"time"
)

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

func (set *urlSet) AddLocation(location string, lastModified time.Time, changeFrequency changeFrequency, priority float32) {
	loc := Location{Location: location, ChangeFrequency: changeFrequency, Priority: priority}
	loc.SetModified(lastModified)
	set.Locations = append(set.Locations, loc)
}

func (set urlSet) String() (string, error) {

	b, err := xml.Marshal(set)
	if err != nil {
		return "", err
	}

	return xml.Header + string(b), nil
}

type Location struct {
	Location        string          `xml:"loc"`
	LastModified    string          `xml:"lastmod,omitempty"`
	ChangeFrequency changeFrequency `xml:"changefreq,omitempty"`
	Priority        float32         `xml:"priority,omitempty"`
}

func (sm *Location) SetModified(t time.Time) {
	if !t.IsZero() {
		sm.LastModified = t.Format(time.RFC3339)
	}
}
