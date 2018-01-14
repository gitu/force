package main

type EnumEntryType string

const (
	EnumEntryTypemandatory EnumEntryType = "mandatory"
	EnumEntryTypeoptional  EnumEntryType = "optional"
	EnumEntryTypestandby   EnumEntryType = "standby"
)
