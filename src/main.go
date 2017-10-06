package breyta

import (
	"net/http"

	"strings"

	"errors"

	"io/ioutil"

	"github.com/sethgrid/pester"
)

/**
*  Created by Galileo on 5/10/17.
*
*  Breyta is a [json-xls](http://www.json-xls.com/api) REST client.
*  Retries if json-xls server isn't responding with exponential jitter backoff strategy.
*
*  Breyta requires a Mashape API key.
*  Check out the pricing for the plans here https://market.mashape.com/json-xls-com/json2xls/pricing
 */

var _JSON_URL = "https://json2xls-json-xls-v1.p.mashape.com/ConvertJsonRaw?format={format}&layout=Auto&view={view}&InternalIDs={internalIDs}"
var _XML_URL = "https://json2xls-json-xls-v1.p.mashape.com/ConvertXmlRaw?format={format}&layout={layout}&view={view}&InternalIDs={internalIDs}"

const (
	FormatXlsx = "XLSX" // Request XLSX format
	FormatXls  = "XLS"  // Request XLS format

	LayoutAuto      = "Auto"
	LayoutPortrait  = "Portrait"
	LayoutLandscape = "Landscape"

	ViewHierarchy = "Hierarchy"
	ViewPlain     = "Plain"

	// Both or None. Unique internal IDs are generated for each JSON token (XML element).
	// They might help to lookup corresponding records(rows) between Excel sheets (CSV files)
	Both = "Both"
	None = "None"
)

type breyta struct {
	format      string
	view        string
	internalIDs string
	mashapeKey  string
	c           *pester.Client
}

// NewJSONClient returns a new breyta client for json-2-xls ConvertJsonRaw API for `ConvertJsonRaw`
// panics if format, view, layout and internalIDs are invalid.
func NewJSONClient(format, view, layout, internalIDs, key string) *breyta {
	if err := validateURLParams(format, view, layout, internalIDs); err != nil {
		panic(err)
	}
	c := pester.New()
	c.Concurrency = 3
	c.MaxRetries = 5
	c.Backoff = pester.ExponentialBackoff
	c.KeepLog = true
	r := strings.NewReplacer("{format}", format, "{view}", view, "{layout}", layout, "{internalIDs}", internalIDs)
	_JSON_URL = r.Replace(_JSON_URL)
	return &breyta{format: format, view: view, internalIDs: internalIDs, mashapeKey: key, c: c}
}

// NewXMLClient returns a new breyta client for json-2-xls ConvertJsonRaw API for `ConvertXmlRaw`
// panics if format, view, layout and internalIDs are invalid.
func NewXMLClient(format, view, layout, internalIDs, key string) *breyta {
	if err := validateURLParams(format, view, layout, internalIDs); err != nil {
		panic(err)
	}
	c := pester.New()
	c.Concurrency = 3
	c.MaxRetries = 5
	c.Backoff = pester.ExponentialBackoff
	c.KeepLog = true
	r := strings.NewReplacer("{format}", format, "{view}", view, "{layout}", layout, "{internalIDs}", internalIDs)
	_JSON_URL = r.Replace(_XML_URL)
	return &breyta{format: format, view: view, internalIDs: internalIDs, mashapeKey: key, c: c}
}

// ConvertJSON makes a new HTTP POST Request to json2xls REST API with the json string
// and returns the converted XLSX or XLS data as byte slice.
func (b *breyta) ConvertJSON(json string) ([]byte, error) {
	req, err := http.NewRequest("POST", _JSON_URL, strings.NewReader(json))
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Mashape-Key", b.mashapeKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := b.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ConvertXML makes a new HTTP POST Request to json2xls REST API with the xml string
// and returns the converted XLSX or XLS data as byte slice.
func (b *breyta) ConvertXML(xml string) ([]byte, error) {
	req, err := http.NewRequest("POST", _XML_URL, strings.NewReader(xml))
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Mashape-Key", b.mashapeKey)
	req.Header.Add("Content-Type", "application/xml")

	resp, err := b.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func validateURLParams(format, view, layout, internalIDs string) error {
	if !isFormatValid(format) {
		return errors.New("invalid format")
	}
	if !isViewValid(view) {
		return errors.New("invalid view")
	}
	if !isInteralIdsValid(internalIDs) {
		return errors.New("invalid internalIds")
	}
	if !isLayoutValid(layout) {
		return errors.New("invalid layout")
	}
	return nil
}

func isFormatValid(format string) bool {
	return format == FormatXls || format == FormatXlsx
}

func isViewValid(view string) bool {
	return view == ViewHierarchy || view == ViewPlain || view == Both
}

func isLayoutValid(layout string) bool {
	return layout == LayoutAuto || layout == LayoutLandscape || layout == LayoutPortrait
}

func isInteralIdsValid(internalIds string) bool {
	return internalIds == Both || internalIds == None
}
