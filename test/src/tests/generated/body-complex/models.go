package complexgroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
)

// CMYKColors enumerates the values for cmyk colors.
type CMYKColors string

const (
	// BlacK specifies the blac k state for cmyk colors.
	BlacK CMYKColors = "blacK"
	// Cyan specifies the cyan state for cmyk colors.
	Cyan CMYKColors = "cyan"
	// Magenta specifies the magenta state for cmyk colors.
	Magenta CMYKColors = "Magenta"
	// YELLOW specifies the yellow state for cmyk colors.
	YELLOW CMYKColors = "YELLOW"
)

// Fishtype enumerates the values for fishtype.
type Fishtype string

const (
	// FishtypeCookiecuttershark specifies the fishtype cookiecuttershark state for fishtype.
	FishtypeCookiecuttershark Fishtype = "cookiecuttershark"
	// FishtypeFish specifies the fishtype fish state for fishtype.
	FishtypeFish Fishtype = "Fish"
	// FishtypeGoblin specifies the fishtype goblin state for fishtype.
	FishtypeGoblin Fishtype = "goblin"
	// FishtypeSalmon specifies the fishtype salmon state for fishtype.
	FishtypeSalmon Fishtype = "salmon"
	// FishtypeSawshark specifies the fishtype sawshark state for fishtype.
	FishtypeSawshark Fishtype = "sawshark"
	// FishtypeShark specifies the fishtype shark state for fishtype.
	FishtypeShark Fishtype = "shark"
)

// ArrayWrapper
type ArrayWrapper struct {
	autorest.Response `json:"-"`
	Array             *[]string `json:"array,omitempty"`
}

// Basic
type Basic struct {
	autorest.Response `json:"-"`
	ID                *int32     `json:"id,omitempty"`
	Name              *string    `json:"name,omitempty"`
	Color             CMYKColors `json:"color,omitempty"`
}

// BooleanWrapper
type BooleanWrapper struct {
	autorest.Response `json:"-"`
	FieldTrue         *bool `json:"field_true,omitempty"`
	FieldFalse        *bool `json:"field_false,omitempty"`
}

// ByteWrapper
type ByteWrapper struct {
	autorest.Response `json:"-"`
	Field             *[]byte `json:"field,omitempty"`
}

// Cat
type Cat struct {
	ID    *int32  `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Color *string `json:"color,omitempty"`
	Hates *[]Dog  `json:"hates,omitempty"`
}

// Cookiecuttershark
type Cookiecuttershark struct {
	Species  *string    `json:"species,omitempty"`
	Length   *float64   `json:"length,omitempty"`
	Siblings *[]IFish   `json:"siblings,omitempty"`
	Fishtype Fishtype   `json:"fishtype,omitempty"`
	Age      *int32     `json:"age,omitempty"`
	Birthday *date.Time `json:"birthday,omitempty"`
}

// MarshalJSON is the custom marshaler for Cookiecuttershark.
func (c Cookiecuttershark) MarshalJSON() ([]byte, error) {
	c.Fishtype = FishtypeCookiecuttershark
	type Alias Cookiecuttershark
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(c),
	})
}

// AsSalmon is the IFish implementation for Cookiecuttershark.
func (c Cookiecuttershark) AsSalmon() (*Salmon, bool) {
	return nil, false
}

// AsShark is the IFish implementation for Cookiecuttershark.
func (c Cookiecuttershark) AsShark() (*Shark, bool) {
	return nil, false
}

// AsSawshark is the IFish implementation for Cookiecuttershark.
func (c Cookiecuttershark) AsSawshark() (*Sawshark, bool) {
	return nil, false
}

// AsGoblinshark is the IFish implementation for Cookiecuttershark.
func (c Cookiecuttershark) AsGoblinshark() (*Goblinshark, bool) {
	return nil, false
}

// AsCookiecuttershark is the IFish implementation for Cookiecuttershark.
func (c Cookiecuttershark) AsCookiecuttershark() (*Cookiecuttershark, bool) {
	return &c, true
}

// AsFish is the IFish implementation for Cookiecuttershark.
func (c Cookiecuttershark) AsFish() (*Fish, bool) {
	return nil, false
}

// UnmarshalJSON is the custom unmarshaler for Cookiecuttershark struct.
func (c *Cookiecuttershark) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["age"]
	if v != nil {
		var age int32
		err = json.Unmarshal(*m["age"], &age)
		if err != nil {
			return err
		}
		c.Age = &age
	}

	v = m["birthday"]
	if v != nil {
		var birthday date.Time
		err = json.Unmarshal(*m["birthday"], &birthday)
		if err != nil {
			return err
		}
		c.Birthday = &birthday
	}

	v = m["species"]
	if v != nil {
		var species string
		err = json.Unmarshal(*m["species"], &species)
		if err != nil {
			return err
		}
		c.Species = &species
	}

	v = m["length"]
	if v != nil {
		var length float64
		err = json.Unmarshal(*m["length"], &length)
		if err != nil {
			return err
		}
		c.Length = &length
	}

	v = m["siblings"]
	if v != nil {
		siblings, err := unmarshalIFishArray(*m["siblings"])
		if err != nil {
			return err
		}
		c.Siblings = &siblings
	}

	v = m["fishtype"]
	if v != nil {
		var fishtype Fishtype
		err = json.Unmarshal(*m["fishtype"], &fishtype)
		if err != nil {
			return err
		}
		c.Fishtype = fishtype
	}

	return nil
}

// Datetimerfc1123Wrapper
type Datetimerfc1123Wrapper struct {
	autorest.Response `json:"-"`
	Field             *date.TimeRFC1123 `json:"field,omitempty"`
	Now               *date.TimeRFC1123 `json:"now,omitempty"`
}

// DatetimeWrapper
type DatetimeWrapper struct {
	autorest.Response `json:"-"`
	Field             *date.Time `json:"field,omitempty"`
	Now               *date.Time `json:"now,omitempty"`
}

// DateWrapper
type DateWrapper struct {
	autorest.Response `json:"-"`
	Field             *date.Date `json:"field,omitempty"`
	Leap              *date.Date `json:"leap,omitempty"`
}

// DictionaryWrapper
type DictionaryWrapper struct {
	autorest.Response `json:"-"`
	DefaultProgram    *map[string]*string `json:"defaultProgram,omitempty"`
}

// Dog
type Dog struct {
	ID   *int32  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Food *string `json:"food,omitempty"`
}

// DoubleWrapper
type DoubleWrapper struct {
	autorest.Response                                                               `json:"-"`
	Field1                                                                          *float64 `json:"field1,omitempty"`
	Field56ZerosAfterTheDotAndNegativeZeroBeforeDotAndThisIsALongFieldNameOnPurpose *float64 `json:"field_56_zeros_after_the_dot_and_negative_zero_before_dot_and_this_is_a_long_field_name_on_purpose,omitempty"`
}

// DurationWrapper
type DurationWrapper struct {
	autorest.Response `json:"-"`
	Field             *string `json:"field,omitempty"`
}

// Error
type Error struct {
	Status  *int32  `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}

// IFish
type IFish interface {
	AsSalmon() (*Salmon, bool)
	AsShark() (*Shark, bool)
	AsSawshark() (*Sawshark, bool)
	AsGoblinshark() (*Goblinshark, bool)
	AsCookiecuttershark() (*Cookiecuttershark, bool)
	AsFish() (*Fish, bool)
}

// Fish
type Fish struct {
	autorest.Response `json:"-"`
	Species           *string  `json:"species,omitempty"`
	Length            *float64 `json:"length,omitempty"`
	Siblings          *[]IFish `json:"siblings,omitempty"`
	Fishtype          Fishtype `json:"fishtype,omitempty"`
}

func unmarshalIFish(body []byte) (IFish, error) {
	var m map[string]interface{}
	err := json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	switch m["fishtype"] {
	case string(FishtypeSalmon):
		var s Salmon
		err := json.Unmarshal(body, &s)
		return s, err
	case string(FishtypeShark):
		var s Shark
		err := json.Unmarshal(body, &s)
		return s, err
	case string(FishtypeSawshark):
		var s Sawshark
		err := json.Unmarshal(body, &s)
		return s, err
	case string(FishtypeGoblin):
		var g Goblinshark
		err := json.Unmarshal(body, &g)
		return g, err
	case string(FishtypeCookiecuttershark):
		var c Cookiecuttershark
		err := json.Unmarshal(body, &c)
		return c, err
	default:
		var f Fish
		err := json.Unmarshal(body, &f)
		return f, err
	}
}
func unmarshalIFishArray(body []byte) ([]IFish, error) {
	var rawMessages []*json.RawMessage
	err := json.Unmarshal(body, &rawMessages)
	if err != nil {
		return nil, err
	}

	fArray := make([]IFish, len(rawMessages))

	for index, rawMessage := range rawMessages {
		f, err := unmarshalIFish(*rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

// MarshalJSON is the custom marshaler for Fish.
func (f Fish) MarshalJSON() ([]byte, error) {
	f.Fishtype = FishtypeFish
	type Alias Fish
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(f),
	})
}

// AsSalmon is the IFish implementation for Fish.
func (f Fish) AsSalmon() (*Salmon, bool) {
	return nil, false
}

// AsShark is the IFish implementation for Fish.
func (f Fish) AsShark() (*Shark, bool) {
	return nil, false
}

// AsSawshark is the IFish implementation for Fish.
func (f Fish) AsSawshark() (*Sawshark, bool) {
	return nil, false
}

// AsGoblinshark is the IFish implementation for Fish.
func (f Fish) AsGoblinshark() (*Goblinshark, bool) {
	return nil, false
}

// AsCookiecuttershark is the IFish implementation for Fish.
func (f Fish) AsCookiecuttershark() (*Cookiecuttershark, bool) {
	return nil, false
}

// AsFish is the IFish implementation for Fish.
func (f Fish) AsFish() (*Fish, bool) {
	return &f, true
}

// UnmarshalJSON is the custom unmarshaler for Fish struct.
func (f *Fish) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["species"]
	if v != nil {
		var species string
		err = json.Unmarshal(*m["species"], &species)
		if err != nil {
			return err
		}
		f.Species = &species
	}

	v = m["length"]
	if v != nil {
		var length float64
		err = json.Unmarshal(*m["length"], &length)
		if err != nil {
			return err
		}
		f.Length = &length
	}

	v = m["siblings"]
	if v != nil {
		siblings, err := unmarshalIFishArray(*m["siblings"])
		if err != nil {
			return err
		}
		f.Siblings = &siblings
	}

	v = m["fishtype"]
	if v != nil {
		var fishtype Fishtype
		err = json.Unmarshal(*m["fishtype"], &fishtype)
		if err != nil {
			return err
		}
		f.Fishtype = fishtype
	}

	return nil
}

// FishModel
type FishModel struct {
	autorest.Response `json:"-"`
	Value             IFish `json:"value,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for FishModel struct.
func (fm *FishModel) UnmarshalJSON(body []byte) error {
	f, err := unmarshalIFish(body)
	if err != nil {
		return err
	}
	fm.Value = f

	return nil
}

// FloatWrapper
type FloatWrapper struct {
	autorest.Response `json:"-"`
	Field1            *float64 `json:"field1,omitempty"`
	Field2            *float64 `json:"field2,omitempty"`
}

// Goblinshark
type Goblinshark struct {
	Species  *string    `json:"species,omitempty"`
	Length   *float64   `json:"length,omitempty"`
	Siblings *[]IFish   `json:"siblings,omitempty"`
	Fishtype Fishtype   `json:"fishtype,omitempty"`
	Age      *int32     `json:"age,omitempty"`
	Birthday *date.Time `json:"birthday,omitempty"`
	Jawsize  *int32     `json:"jawsize,omitempty"`
}

// MarshalJSON is the custom marshaler for Goblinshark.
func (g Goblinshark) MarshalJSON() ([]byte, error) {
	g.Fishtype = FishtypeGoblin
	type Alias Goblinshark
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(g),
	})
}

// AsSalmon is the IFish implementation for Goblinshark.
func (g Goblinshark) AsSalmon() (*Salmon, bool) {
	return nil, false
}

// AsShark is the IFish implementation for Goblinshark.
func (g Goblinshark) AsShark() (*Shark, bool) {
	return nil, false
}

// AsSawshark is the IFish implementation for Goblinshark.
func (g Goblinshark) AsSawshark() (*Sawshark, bool) {
	return nil, false
}

// AsGoblinshark is the IFish implementation for Goblinshark.
func (g Goblinshark) AsGoblinshark() (*Goblinshark, bool) {
	return &g, true
}

// AsCookiecuttershark is the IFish implementation for Goblinshark.
func (g Goblinshark) AsCookiecuttershark() (*Cookiecuttershark, bool) {
	return nil, false
}

// AsFish is the IFish implementation for Goblinshark.
func (g Goblinshark) AsFish() (*Fish, bool) {
	return nil, false
}

// UnmarshalJSON is the custom unmarshaler for Goblinshark struct.
func (g *Goblinshark) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["jawsize"]
	if v != nil {
		var jawsize int32
		err = json.Unmarshal(*m["jawsize"], &jawsize)
		if err != nil {
			return err
		}
		g.Jawsize = &jawsize
	}

	v = m["age"]
	if v != nil {
		var age int32
		err = json.Unmarshal(*m["age"], &age)
		if err != nil {
			return err
		}
		g.Age = &age
	}

	v = m["birthday"]
	if v != nil {
		var birthday date.Time
		err = json.Unmarshal(*m["birthday"], &birthday)
		if err != nil {
			return err
		}
		g.Birthday = &birthday
	}

	v = m["species"]
	if v != nil {
		var species string
		err = json.Unmarshal(*m["species"], &species)
		if err != nil {
			return err
		}
		g.Species = &species
	}

	v = m["length"]
	if v != nil {
		var length float64
		err = json.Unmarshal(*m["length"], &length)
		if err != nil {
			return err
		}
		g.Length = &length
	}

	v = m["siblings"]
	if v != nil {
		siblings, err := unmarshalIFishArray(*m["siblings"])
		if err != nil {
			return err
		}
		g.Siblings = &siblings
	}

	v = m["fishtype"]
	if v != nil {
		var fishtype Fishtype
		err = json.Unmarshal(*m["fishtype"], &fishtype)
		if err != nil {
			return err
		}
		g.Fishtype = fishtype
	}

	return nil
}

// IntWrapper
type IntWrapper struct {
	autorest.Response `json:"-"`
	Field1            *int32 `json:"field1,omitempty"`
	Field2            *int32 `json:"field2,omitempty"`
}

// LongWrapper
type LongWrapper struct {
	autorest.Response `json:"-"`
	Field1            *int64 `json:"field1,omitempty"`
	Field2            *int64 `json:"field2,omitempty"`
}

// Pet
type Pet struct {
	ID   *int32  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// ReadonlyObj
type ReadonlyObj struct {
	autorest.Response `json:"-"`
	ID                *string `json:"id,omitempty"`
	Size              *int32  `json:"size,omitempty"`
}

// Salmon
type Salmon struct {
	Species  *string  `json:"species,omitempty"`
	Length   *float64 `json:"length,omitempty"`
	Siblings *[]IFish `json:"siblings,omitempty"`
	Fishtype Fishtype `json:"fishtype,omitempty"`
	Location *string  `json:"location,omitempty"`
	Iswild   *bool    `json:"iswild,omitempty"`
}

// MarshalJSON is the custom marshaler for Salmon.
func (s Salmon) MarshalJSON() ([]byte, error) {
	s.Fishtype = FishtypeSalmon
	type Alias Salmon
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(s),
	})
}

// AsSalmon is the IFish implementation for Salmon.
func (s Salmon) AsSalmon() (*Salmon, bool) {
	return &s, true
}

// AsShark is the IFish implementation for Salmon.
func (s Salmon) AsShark() (*Shark, bool) {
	return nil, false
}

// AsSawshark is the IFish implementation for Salmon.
func (s Salmon) AsSawshark() (*Sawshark, bool) {
	return nil, false
}

// AsGoblinshark is the IFish implementation for Salmon.
func (s Salmon) AsGoblinshark() (*Goblinshark, bool) {
	return nil, false
}

// AsCookiecuttershark is the IFish implementation for Salmon.
func (s Salmon) AsCookiecuttershark() (*Cookiecuttershark, bool) {
	return nil, false
}

// AsFish is the IFish implementation for Salmon.
func (s Salmon) AsFish() (*Fish, bool) {
	return nil, false
}

// UnmarshalJSON is the custom unmarshaler for Salmon struct.
func (s *Salmon) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["location"]
	if v != nil {
		var location string
		err = json.Unmarshal(*m["location"], &location)
		if err != nil {
			return err
		}
		s.Location = &location
	}

	v = m["iswild"]
	if v != nil {
		var iswild bool
		err = json.Unmarshal(*m["iswild"], &iswild)
		if err != nil {
			return err
		}
		s.Iswild = &iswild
	}

	v = m["species"]
	if v != nil {
		var species string
		err = json.Unmarshal(*m["species"], &species)
		if err != nil {
			return err
		}
		s.Species = &species
	}

	v = m["length"]
	if v != nil {
		var length float64
		err = json.Unmarshal(*m["length"], &length)
		if err != nil {
			return err
		}
		s.Length = &length
	}

	v = m["siblings"]
	if v != nil {
		siblings, err := unmarshalIFishArray(*m["siblings"])
		if err != nil {
			return err
		}
		s.Siblings = &siblings
	}

	v = m["fishtype"]
	if v != nil {
		var fishtype Fishtype
		err = json.Unmarshal(*m["fishtype"], &fishtype)
		if err != nil {
			return err
		}
		s.Fishtype = fishtype
	}

	return nil
}

// Sawshark
type Sawshark struct {
	Species  *string    `json:"species,omitempty"`
	Length   *float64   `json:"length,omitempty"`
	Siblings *[]IFish   `json:"siblings,omitempty"`
	Fishtype Fishtype   `json:"fishtype,omitempty"`
	Age      *int32     `json:"age,omitempty"`
	Birthday *date.Time `json:"birthday,omitempty"`
	Picture  *[]byte    `json:"picture,omitempty"`
}

// MarshalJSON is the custom marshaler for Sawshark.
func (s Sawshark) MarshalJSON() ([]byte, error) {
	s.Fishtype = FishtypeSawshark
	type Alias Sawshark
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(s),
	})
}

// AsSalmon is the IFish implementation for Sawshark.
func (s Sawshark) AsSalmon() (*Salmon, bool) {
	return nil, false
}

// AsShark is the IFish implementation for Sawshark.
func (s Sawshark) AsShark() (*Shark, bool) {
	return nil, false
}

// AsSawshark is the IFish implementation for Sawshark.
func (s Sawshark) AsSawshark() (*Sawshark, bool) {
	return &s, true
}

// AsGoblinshark is the IFish implementation for Sawshark.
func (s Sawshark) AsGoblinshark() (*Goblinshark, bool) {
	return nil, false
}

// AsCookiecuttershark is the IFish implementation for Sawshark.
func (s Sawshark) AsCookiecuttershark() (*Cookiecuttershark, bool) {
	return nil, false
}

// AsFish is the IFish implementation for Sawshark.
func (s Sawshark) AsFish() (*Fish, bool) {
	return nil, false
}

// UnmarshalJSON is the custom unmarshaler for Sawshark struct.
func (s *Sawshark) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["picture"]
	if v != nil {
		var picture []byte
		err = json.Unmarshal(*m["picture"], &picture)
		if err != nil {
			return err
		}
		s.Picture = &picture
	}

	v = m["age"]
	if v != nil {
		var age int32
		err = json.Unmarshal(*m["age"], &age)
		if err != nil {
			return err
		}
		s.Age = &age
	}

	v = m["birthday"]
	if v != nil {
		var birthday date.Time
		err = json.Unmarshal(*m["birthday"], &birthday)
		if err != nil {
			return err
		}
		s.Birthday = &birthday
	}

	v = m["species"]
	if v != nil {
		var species string
		err = json.Unmarshal(*m["species"], &species)
		if err != nil {
			return err
		}
		s.Species = &species
	}

	v = m["length"]
	if v != nil {
		var length float64
		err = json.Unmarshal(*m["length"], &length)
		if err != nil {
			return err
		}
		s.Length = &length
	}

	v = m["siblings"]
	if v != nil {
		siblings, err := unmarshalIFishArray(*m["siblings"])
		if err != nil {
			return err
		}
		s.Siblings = &siblings
	}

	v = m["fishtype"]
	if v != nil {
		var fishtype Fishtype
		err = json.Unmarshal(*m["fishtype"], &fishtype)
		if err != nil {
			return err
		}
		s.Fishtype = fishtype
	}

	return nil
}

// IShark
type IShark interface {
	AsSawshark() (*Sawshark, bool)
	AsGoblinshark() (*Goblinshark, bool)
	AsCookiecuttershark() (*Cookiecuttershark, bool)
	AsShark() (*Shark, bool)
}

// Shark
type Shark struct {
	Species  *string    `json:"species,omitempty"`
	Length   *float64   `json:"length,omitempty"`
	Siblings *[]IFish   `json:"siblings,omitempty"`
	Fishtype Fishtype   `json:"fishtype,omitempty"`
	Age      *int32     `json:"age,omitempty"`
	Birthday *date.Time `json:"birthday,omitempty"`
}

func unmarshalIShark(body []byte) (IShark, error) {
	var m map[string]interface{}
	err := json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	switch m["fishtype"] {
	case string(FishtypeSawshark):
		var s Sawshark
		err := json.Unmarshal(body, &s)
		return s, err
	case string(FishtypeGoblin):
		var g Goblinshark
		err := json.Unmarshal(body, &g)
		return g, err
	case string(FishtypeCookiecuttershark):
		var c Cookiecuttershark
		err := json.Unmarshal(body, &c)
		return c, err
	default:
		var s Shark
		err := json.Unmarshal(body, &s)
		return s, err
	}
}
func unmarshalISharkArray(body []byte) ([]IShark, error) {
	var rawMessages []*json.RawMessage
	err := json.Unmarshal(body, &rawMessages)
	if err != nil {
		return nil, err
	}

	sArray := make([]IShark, len(rawMessages))

	for index, rawMessage := range rawMessages {
		s, err := unmarshalIShark(*rawMessage)
		if err != nil {
			return nil, err
		}
		sArray[index] = s
	}
	return sArray, nil
}

// MarshalJSON is the custom marshaler for Shark.
func (s Shark) MarshalJSON() ([]byte, error) {
	s.Fishtype = FishtypeShark
	type Alias Shark
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(s),
	})
}

// AsSalmon is the IFish implementation for Shark.
func (s Shark) AsSalmon() (*Salmon, bool) {
	return nil, false
}

// AsShark is the IFish implementation for Shark.
func (s Shark) AsShark() (*Shark, bool) {
	return &s, true
}

// AsSawshark is the IFish implementation for Shark.
func (s Shark) AsSawshark() (*Sawshark, bool) {
	return nil, false
}

// AsGoblinshark is the IFish implementation for Shark.
func (s Shark) AsGoblinshark() (*Goblinshark, bool) {
	return nil, false
}

// AsCookiecuttershark is the IFish implementation for Shark.
func (s Shark) AsCookiecuttershark() (*Cookiecuttershark, bool) {
	return nil, false
}

// AsFish is the IFish implementation for Shark.
func (s Shark) AsFish() (*Fish, bool) {
	return nil, false
}

// UnmarshalJSON is the custom unmarshaler for Shark struct.
func (s *Shark) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["age"]
	if v != nil {
		var age int32
		err = json.Unmarshal(*m["age"], &age)
		if err != nil {
			return err
		}
		s.Age = &age
	}

	v = m["birthday"]
	if v != nil {
		var birthday date.Time
		err = json.Unmarshal(*m["birthday"], &birthday)
		if err != nil {
			return err
		}
		s.Birthday = &birthday
	}

	v = m["species"]
	if v != nil {
		var species string
		err = json.Unmarshal(*m["species"], &species)
		if err != nil {
			return err
		}
		s.Species = &species
	}

	v = m["length"]
	if v != nil {
		var length float64
		err = json.Unmarshal(*m["length"], &length)
		if err != nil {
			return err
		}
		s.Length = &length
	}

	v = m["siblings"]
	if v != nil {
		siblings, err := unmarshalIFishArray(*m["siblings"])
		if err != nil {
			return err
		}
		s.Siblings = &siblings
	}

	v = m["fishtype"]
	if v != nil {
		var fishtype Fishtype
		err = json.Unmarshal(*m["fishtype"], &fishtype)
		if err != nil {
			return err
		}
		s.Fishtype = fishtype
	}

	return nil
}

// Siamese
type Siamese struct {
	autorest.Response `json:"-"`
	ID                *int32  `json:"id,omitempty"`
	Name              *string `json:"name,omitempty"`
	Color             *string `json:"color,omitempty"`
	Hates             *[]Dog  `json:"hates,omitempty"`
	Breed             *string `json:"breed,omitempty"`
}

// StringWrapper
type StringWrapper struct {
	autorest.Response `json:"-"`
	Field             *string `json:"field,omitempty"`
	Empty             *string `json:"empty,omitempty"`
	Null              *string `json:"null,omitempty"`
}
