package sitemap_go

import (
	"encoding/xml"
	"io"
	"time"
)

// New creates a new Sitemap.
func New() *Sitemap {
	return &Sitemap{
		Namespace:      namespace,
		ImageNamespace: imageNamespace,
		VideoNamespace: videoNamespace,
	}
}

// Sitemap represents a sitemap urlset.
type Sitemap struct {
	XMLName        xml.Name   `xml:"urlset"`
	Namespace      string     `xml:"xmlns,attr"`
	ImageNamespace string     `xml:"xmlns:image,attr,omitempty"`
	VideoNamespace string     `xml:"xmlns:video,attr,omitempty"`
	Locations      []Location `xml:"url"`
}

// AddLocation adds a new location to the sitemap.
func (s *Sitemap) AddLocation(location string, lastModified time.Time, changeFrequency ChangeFrequency, priority float32) *Location {

	if priority < 0 {
		priority = 0
	}
	if priority > 1 {
		priority = 1
	}

	loc := Location{Location: location, ChangeFrequency: changeFrequency, Priority: priority}
	loc.SetModified(lastModified)
	s.Locations = append(s.Locations, loc)

	return &s.Locations[len(s.Locations)-1]
}

// String returns the XML string representation of the sitemap.
func (s *Sitemap) String() (string, error) {
	b, err := xml.Marshal(s)
	if err != nil {
		return "", err
	}

	return xml.Header + string(b), nil
}

// WriteTo writes the sitemap XML to the provided writer.
func (s *Sitemap) WriteTo(w io.Writer) (int64, error) {
	_, err := w.Write([]byte(xml.Header))
	if err != nil {
		return 0, err
	}

	enc := xml.NewEncoder(w)
	defer enc.Close()

	err = enc.Encode(s)
	if err != nil {
		return 0, err
	}

	return 0, nil // io.Writer doesn't easily return total bytes for multiple writes without a wrapper
}

type Location struct {
	Location        string          `xml:"loc"`
	LastModified    string          `xml:"lastmod,omitempty"`
	ChangeFrequency ChangeFrequency `xml:"changefreq,omitempty"`
	Priority        float32         `xml:"priority,omitempty"`
	Images          []Image         `xml:"image:image,omitempty"`
	Videos          []Video         `xml:"video:video,omitempty"`
}

func (sm *Location) SetModified(t time.Time) {
	if !t.IsZero() {
		sm.LastModified = t.Format(time.RFC3339)
	}
}

// AddImage adds an image to the location.
func (sm *Location) AddImage(image Image) {
	sm.Images = append(sm.Images, image)
}

// AddVideo adds a video to the location.
func (sm *Location) AddVideo(video Video) {
	sm.Videos = append(sm.Videos, video)
}

// Image represents a sitemap image.
type Image struct {
	Location    string `xml:"image:loc"`
	Caption     string `xml:"image:caption,omitempty"`
	GeoLocation string `xml:"image:geo_location,omitempty"`
	Title       string `xml:"image:title,omitempty"`
	License     string `xml:"image:license,omitempty"`
}

// Video represents a sitemap video.
type Video struct {
	ThumbnailLocation    string     `xml:"video:thumbnail_loc"`
	Title                string     `xml:"video:title"`
	Description          string     `xml:"video:description"`
	ContentLocation      string     `xml:"video:content_loc,omitempty"`
	PlayerLocation       string     `xml:"video:player_loc,omitempty"`
	Duration             int        `xml:"video:duration,omitempty"`
	ExpirationDate       *time.Time `xml:"video:expiration_date,omitempty"`
	Rating               float32    `xml:"video:rating,omitempty"`
	ViewCount            int        `xml:"video:view_count,omitempty"`
	PublicationDate      *time.Time `xml:"video:publication_date,omitempty"`
	FamilyFriendly       string     `xml:"video:family_friendly,omitempty"`
	Restriction          string     `xml:"video:restriction,omitempty"`
	GalleryLocation      string     `xml:"video:gallery_loc,omitempty"`
	Price                string     `xml:"video:price,omitempty"`
	RequiresSubscription string     `xml:"video:requires_subscription,omitempty"`
	Uploader             string     `xml:"video:uploader,omitempty"`
	Platform             string     `xml:"video:platform,omitempty"`
	Live                 string     `xml:"video:live,omitempty"`
}
