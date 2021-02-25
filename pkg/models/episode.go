package models

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type DialogType string

const (
	DialogTypeUnkown = DialogType("unknown")
	DialogTypeSong   = DialogType("song")
	DialogTypeChat   = DialogType("chat")
)

type MetadataType string

const (
	MetadataTypePublication = MetadataType("publication")
	MetadataTypeSeries      = MetadataType("series")
	MetadataTypeDate        = MetadataType("date")
)

type Dialog struct {
	Type    DialogType `json:"type"`
	Actor   string     `json:"actor"`
	Content string     `json:"content"`
}

type Metadata struct {
	Type  MetadataType `json:"type"`
	Value string       `json:"value"`
}

type Episode struct {
	Source     string     `json:"source"`
	Meta       []Metadata `json:"metadata"`
	Transcript []Dialog   `json:"transcript"`
}

func (e Episode) MetaValue(t MetadataType) string {
	for _, m := range e.Meta {
		if m.Type == t {
			return m.Value
		}
	}
	return "na"
}

func (e Episode) CanonicalName() string {
	date := "na"
	if rawDate := e.MetaValue(MetadataTypeDate); rawDate != "" {
		t, err := time.Parse(time.RFC3339, rawDate)
		if err == nil {
			date = t.Format("Jan-01-2006")
		}
	}

	//avoid overwriting na files
	unique := fmt.Sprintf("%x", sha256.Sum256([]byte(e.Source)))

	return fmt.Sprintf("%s-%s-%s-%s", e.MetaValue(MetadataTypePublication), e.MetaValue(MetadataTypeSeries), date, unique[0:6])
}

