package sitemap

import (
	"encoding/xml"
	"io"
	"time"
)

// NewIndex creates a new Index.
func NewIndex() *Index {
	return &Index{
		Namespace: namespace,
	}
}

// Index represents a sitemap index.
type Index struct {
	XMLName   xml.Name       `xml:"sitemapindex"`
	Namespace string         `xml:"xmlns,attr"`
	SiteMaps  []IndexSitemap `xml:"sitemap"`
}

// AddSitemap adds a new sitemap to the index.
func (i *Index) AddSitemap(location string, lastModified time.Time) {
	sm := IndexSitemap{Location: location}
	sm.SetModified(lastModified)
	i.SiteMaps = append(i.SiteMaps, sm)
}

// String returns the XML string representation of the index.
func (i *Index) String() (string, error) {
	b, err := xml.Marshal(i)
	if err != nil {
		return "", err
	}

	return xml.Header + string(b), nil
}

// WriteTo writes the index XML to the provided writer.
func (i *Index) WriteTo(w io.Writer) (int64, error) {
	_, err := w.Write([]byte(xml.Header))
	if err != nil {
		return 0, err
	}

	enc := xml.NewEncoder(w)
	defer enc.Close()

	err = enc.Encode(i)
	if err != nil {
		return 0, err
	}

	return 0, nil
}

type IndexSitemap struct {
	Location     string `xml:"loc"`
	LastModified string `xml:"lastmod,omitempty"`
}

func (sm *IndexSitemap) SetModified(t time.Time) {
	if !t.IsZero() {
		sm.LastModified = t.Format(time.RFC3339)
	}
}
