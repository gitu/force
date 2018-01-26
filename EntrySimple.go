package main

import (
	"encoding/json"
	"errors"
	"github.com/gitu/force/goraml"
	"log"
)

type EntrySimple struct {
	Start goraml.DateTime `json:"start"`
	End   goraml.DateTime `json:"end"`
	Title string          `json:"title"`
	Group string          `json:"group"`
	Type  EnumEntryType   `json:"type"`
}

var _allGroups = map[string]string{
	"abst":    "Absturzsicherung",
	"alle":    "Alle",
	"as":      "Atemschutz",
	"elektro": "Elektro Gruppe",
	"hng":     "Herznotfall Gruppe",
	"journ":   "Journalführer",
	"kader":   "Kader",
	"kdo":     "Kommando",
	"kurs":    "Kurse",
	"mwd":     "MWD",
	"mwd.1":   "MWD Gruppe 1",
	"mwd.2":   "MWD Gruppe 2",
	"mwd.3":   "MWD Gruppe 3",
	"mwd.4":   "MWD Gruppe 4",
	"neu":     "Neueingeteilte",
	"pn.1":    "Pikett Nacht 1",
	"pn.2":    "Pikett Nacht 2",
	"pt.1":    "Pikett Tag 1",
	"pt.2":    "Pikett Tag 2",
	"san":     "Sanität",
	"stab":    "Stab",
	"tank":    "Tanklagergruppe",
	"vkd":     "VKD",
	"z.1":     "Zug 1",
	"z.2":     "Zug 2",
	"z.3":     "Zug 3",
	"z.4":     "Zug 4",
	"info":    "Zur Information",
}

func (s EntrySimple) ToEntry() (*Entry, error) {
	group, err := GetGroup(s.Group)
	if err != nil {
		return nil, err
	}
	return &Entry{
		Start: s.Start,
		End:   s.End,
		Name:  s.Title,
		Group: *group,
		Type:  s.Type,
	}, nil
}

func GetAllGroups() *[]Group {
	var groups []Group
	for key, value := range _allGroups {
		groups = append(groups, Group{
			Id:   key,
			Name: value,
		})
	}
	return &groups
}

var allGroups = GetAllGroups()
var allEntries = GetAllEntries()

func GetAllEntries() *[]Entry {
	data := MustAsset("entries.json")
	entries := make([]EntrySimple, 0)
	err := json.Unmarshal(data, &entries)

	if err != nil {
		log.Fatal("FAILED laoding entries", err)
	}
	var target []Entry

	for _, e := range entries {
		en, err := e.ToEntry()
		if err != nil {
			log.Fatal("FAILED laoding entries", err)
		}
		target = append(target, *en)
	}
	return &target
}

func GetGroup(s string) (*Group, error) {
	value, found := _allGroups[s]
	if found {
		return &Group{
			Id:   s,
			Name: value,
		}, nil
	} else {
		return nil, errors.New("not found")
	}
}

func GetEntries(groupIds []string) []Entry {
	var entries []Entry
	selectedGroups := make(map[string]string)
	for _, groupId := range groupIds {
		selectedGroups[groupId] = groupId
	}
	for _, e := range *allEntries {
		if _, ok := selectedGroups[e.Group.Id]; ok {
			entries = append(entries, e)
		}
	}
	return entries
}

func GetGroups() Groups {
	var groups Groups
	for _, g := range *allGroups {
		groups = append(groups, g)
	}
	return groups
}
