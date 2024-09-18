package types

//go:generate stringer -type=Type

type Type int

const (
	Integer Type = iota
)
