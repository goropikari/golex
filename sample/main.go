// Code generated by tlex. DO NOT EDIT.

package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

// generated lexer returned types are (int, error).
const (
	Keyword int = iota + 1
	Type
	Identifier
	Digit
	Whitespace
	LParen
	RParen
	LBracket
	RBracket
	Operator
	Hiragana
)

type yyStateID = int
type yyRegexID = int

var YYText string

var (
	ErrYYScan = errors.New("failed to scan")
)

// state id to regex id
var yyStateIDToRegexID = []yyRegexID{
	0, // state 0 is dead state
	5,
	3,
	10,
	3,
	3,
	3,
	3,
	3,
	3,
	1,
	3,
	11,
	3,
	2,
	3,
	4,
	3,
	3,
	5,
	12,
	3,
	7,
	8,
	12,
	3,
	3,
	3,
	6,
	9,
	3,
	3,
	3,
	3,
	3,
}

var yyFinStates = map[yyStateID]struct{}{
	0:  {},
	1:  {},
	2:  {},
	3:  {},
	4:  {},
	5:  {},
	6:  {},
	7:  {},
	8:  {},
	9:  {},
	10: {},
	11: {},
	12: {},
	13: {},
	14: {},
	15: {},
	16: {},
	17: {},
	18: {},
	19: {},
	20: {},
	21: {},
	22: {},
	23: {},
	24: {},
	25: {},
	26: {},
	27: {},
	28: {},
	29: {},
	30: {},
	31: {},
	32: {},
	33: {},
	34: {},
}

type yyinterval struct {
	l int
	r int
}

func (x yyinterval) overlap(y yyinterval) bool {
	return y.l <= x.r && x.l <= y.r
}

var yyTransitionTable = map[yyStateID]map[yyinterval]yyStateID{
	1: {
		yyinterval{l: 13, r: 13}:         19,
		yyinterval{l: 33, r: 33}:         24,
		yyinterval{l: 109, r: 109}:       2,
		yyinterval{l: 14, r: 31}:         20,
		yyinterval{l: 34, r: 39}:         20,
		yyinterval{l: 91, r: 96}:         20,
		yyinterval{l: 40, r: 40}:         28,
		yyinterval{l: 58, r: 58}:         24,
		yyinterval{l: 49, r: 51}:         16,
		yyinterval{l: 10, r: 10}:         19,
		yyinterval{l: 12437, r: 1114111}: 20,
		yyinterval{l: 118, r: 118}:       2,
		yyinterval{l: 53, r: 53}:         16,
		yyinterval{l: 106, r: 107}:       2,
		yyinterval{l: 44, r: 44}:         20,
		yyinterval{l: 116, r: 116}:       2,
		yyinterval{l: 119, r: 119}:       25,
		yyinterval{l: 120, r: 122}:       2,
		yyinterval{l: 61, r: 61}:         24,
		yyinterval{l: 11, r: 12}:         20,
		yyinterval{l: 48, r: 48}:         20,
		yyinterval{l: 46, r: 46}:         20,
		yyinterval{l: 123, r: 123}:       23,
		yyinterval{l: 111, r: 111}:       2,
		yyinterval{l: 110, r: 110}:       2,
		yyinterval{l: 47, r: 47}:         3,
		yyinterval{l: 55, r: 57}:         16,
		yyinterval{l: 12353, r: 12436}:   12,
		yyinterval{l: 125, r: 125}:       29,
		yyinterval{l: 42, r: 42}:         3,
		yyinterval{l: 9, r: 9}:           19,
		yyinterval{l: 32, r: 32}:         19,
		yyinterval{l: 0, r: 8}:           20,
		yyinterval{l: 105, r: 105}:       27,
		yyinterval{l: 102, r: 102}:       26,
		yyinterval{l: 54, r: 54}:         16,
		yyinterval{l: 59, r: 60}:         20,
		yyinterval{l: 52, r: 52}:         16,
		yyinterval{l: 62, r: 64}:         20,
		yyinterval{l: 101, r: 101}:       2,
		yyinterval{l: 65, r: 90}:         2,
		yyinterval{l: 108, r: 108}:       2,
		yyinterval{l: 112, r: 113}:       2,
		yyinterval{l: 99, r: 99}:         2,
		yyinterval{l: 114, r: 114}:       21,
		yyinterval{l: 45, r: 45}:         3,
		yyinterval{l: 124, r: 124}:       20,
		yyinterval{l: 97, r: 97}:         2,
		yyinterval{l: 100, r: 100}:       2,
		yyinterval{l: 115, r: 115}:       2,
		yyinterval{l: 98, r: 98}:         2,
		yyinterval{l: 126, r: 12352}:     20,
		yyinterval{l: 103, r: 103}:       2,
		yyinterval{l: 117, r: 117}:       2,
		yyinterval{l: 41, r: 41}:         22,
		yyinterval{l: 104, r: 104}:       2,
		yyinterval{l: 43, r: 43}:         3,
	},
	2: {
		yyinterval{l: 104, r: 104}: 2,
		yyinterval{l: 98, r: 98}:   2,
		yyinterval{l: 97, r: 97}:   2,
		yyinterval{l: 120, r: 122}: 2,
		yyinterval{l: 49, r: 51}:   2,
		yyinterval{l: 103, r: 103}: 2,
		yyinterval{l: 106, r: 107}: 2,
		yyinterval{l: 48, r: 48}:   2,
		yyinterval{l: 105, r: 105}: 2,
		yyinterval{l: 55, r: 57}:   2,
		yyinterval{l: 101, r: 101}: 2,
		yyinterval{l: 114, r: 114}: 2,
		yyinterval{l: 99, r: 99}:   2,
		yyinterval{l: 52, r: 52}:   2,
		yyinterval{l: 102, r: 102}: 2,
		yyinterval{l: 100, r: 100}: 2,
		yyinterval{l: 111, r: 111}: 2,
		yyinterval{l: 53, r: 53}:   2,
		yyinterval{l: 54, r: 54}:   2,
		yyinterval{l: 115, r: 115}: 2,
		yyinterval{l: 119, r: 119}: 2,
		yyinterval{l: 108, r: 108}: 2,
		yyinterval{l: 109, r: 109}: 2,
		yyinterval{l: 117, r: 117}: 2,
		yyinterval{l: 116, r: 116}: 2,
		yyinterval{l: 110, r: 110}: 2,
		yyinterval{l: 118, r: 118}: 2,
		yyinterval{l: 65, r: 90}:   2,
		yyinterval{l: 112, r: 113}: 2,
	},
	4: {
		yyinterval{l: 49, r: 51}:   2,
		yyinterval{l: 97, r: 97}:   2,
		yyinterval{l: 101, r: 101}: 2,
		yyinterval{l: 109, r: 109}: 2,
		yyinterval{l: 111, r: 111}: 2,
		yyinterval{l: 100, r: 100}: 2,
		yyinterval{l: 118, r: 118}: 2,
		yyinterval{l: 52, r: 52}:   2,
		yyinterval{l: 99, r: 99}:   2,
		yyinterval{l: 53, r: 53}:   2,
		yyinterval{l: 119, r: 119}: 2,
		yyinterval{l: 115, r: 115}: 2,
		yyinterval{l: 55, r: 57}:   2,
		yyinterval{l: 116, r: 116}: 2,
		yyinterval{l: 104, r: 104}: 2,
		yyinterval{l: 106, r: 107}: 2,
		yyinterval{l: 102, r: 102}: 2,
		yyinterval{l: 98, r: 98}:   2,
		yyinterval{l: 112, r: 113}: 2,
		yyinterval{l: 110, r: 110}: 2,
		yyinterval{l: 65, r: 90}:   2,
		yyinterval{l: 108, r: 108}: 2,
		yyinterval{l: 54, r: 54}:   2,
		yyinterval{l: 114, r: 114}: 2,
		yyinterval{l: 117, r: 117}: 13,
		yyinterval{l: 120, r: 122}: 2,
		yyinterval{l: 48, r: 48}:   2,
		yyinterval{l: 105, r: 105}: 2,
		yyinterval{l: 103, r: 103}: 2,
	},
	5: {
		yyinterval{l: 104, r: 104}: 2,
		yyinterval{l: 106, r: 107}: 2,
		yyinterval{l: 97, r: 97}:   2,
		yyinterval{l: 102, r: 102}: 2,
		yyinterval{l: 120, r: 122}: 2,
		yyinterval{l: 49, r: 51}:   2,
		yyinterval{l: 118, r: 118}: 2,
		yyinterval{l: 55, r: 57}:   2,
		yyinterval{l: 116, r: 116}: 2,
		yyinterval{l: 65, r: 90}:   2,
		yyinterval{l: 114, r: 114}: 2,
		yyinterval{l: 112, r: 113}: 2,
		yyinterval{l: 53, r: 53}:   2,
		yyinterval{l: 54, r: 54}:   2,
		yyinterval{l: 103, r: 103}: 2,
		yyinterval{l: 100, r: 100}: 2,
		yyinterval{l: 108, r: 108}: 2,
		yyinterval{l: 111, r: 111}: 2,
		yyinterval{l: 115, r: 115}: 2,
		yyinterval{l: 48, r: 48}:   2,
		yyinterval{l: 98, r: 98}:   2,
		yyinterval{l: 105, r: 105}: 2,
		yyinterval{l: 110, r: 110}: 2,
		yyinterval{l: 52, r: 52}:   14,
		yyinterval{l: 99, r: 99}:   2,
		yyinterval{l: 101, r: 101}: 2,
		yyinterval{l: 117, r: 117}: 2,
		yyinterval{l: 119, r: 119}: 2,
		yyinterval{l: 109, r: 109}: 2,
	},
	6: {
		yyinterval{l: 102, r: 102}: 2,
		yyinterval{l: 118, r: 118}: 2,
		yyinterval{l: 52, r: 52}:   2,
		yyinterval{l: 116, r: 116}: 2,
		yyinterval{l: 100, r: 100}: 2,
		yyinterval{l: 101, r: 101}: 2,
		yyinterval{l: 49, r: 51}:   2,
		yyinterval{l: 108, r: 108}: 2,
		yyinterval{l: 111, r: 111}: 2,
		yyinterval{l: 98, r: 98}:   2,
		yyinterval{l: 48, r: 48}:   2,
		yyinterval{l: 117, r: 117}: 2,
		yyinterval{l: 53, r: 53}:   2,
		yyinterval{l: 65, r: 90}:   2,
		yyinterval{l: 54, r: 54}:   2,
		yyinterval{l: 115, r: 115}: 2,
		yyinterval{l: 109, r: 109}: 2,
		yyinterval{l: 105, r: 105}: 2,
		yyinterval{l: 119, r: 119}: 2,
		yyinterval{l: 106, r: 107}: 2,
		yyinterval{l: 120, r: 122}: 2,
		yyinterval{l: 55, r: 57}:   2,
		yyinterval{l: 114, r: 114}: 2,
		yyinterval{l: 112, r: 113}: 2,
		yyinterval{l: 99, r: 99}:   2,
		yyinterval{l: 103, r: 103}: 2,
		yyinterval{l: 104, r: 104}: 2,
		yyinterval{l: 110, r: 110}: 2,
		yyinterval{l: 97, r: 97}:   15,
	},
	7: {
		yyinterval{l: 54, r: 54}:   2,
		yyinterval{l: 109, r: 109}: 2,
		yyinterval{l: 99, r: 99}:   10,
		yyinterval{l: 111, r: 111}: 2,
		yyinterval{l: 116, r: 116}: 2,
		yyinterval{l: 117, r: 117}: 2,
		yyinterval{l: 119, r: 119}: 2,
		yyinterval{l: 112, r: 113}: 2,
		yyinterval{l: 105, r: 105}: 2,
		yyinterval{l: 115, r: 115}: 2,
		yyinterval{l: 48, r: 48}:   2,
		yyinterval{l: 102, r: 102}: 2,
		yyinterval{l: 100, r: 100}: 2,
		yyinterval{l: 114, r: 114}: 2,
		yyinterval{l: 104, r: 104}: 2,
		yyinterval{l: 106, r: 107}: 2,
		yyinterval{l: 101, r: 101}: 2,
		yyinterval{l: 55, r: 57}:   2,
		yyinterval{l: 49, r: 51}:   2,
		yyinterval{l: 52, r: 52}:   2,
		yyinterval{l: 120, r: 122}: 2,
		yyinterval{l: 108, r: 108}: 2,
		yyinterval{l: 53, r: 53}:   2,
		yyinterval{l: 98, r: 98}:   2,
		yyinterval{l: 103, r: 103}: 2,
		yyinterval{l: 97, r: 97}:   2,
		yyinterval{l: 65, r: 90}:   2,
		yyinterval{l: 110, r: 110}: 2,
		yyinterval{l: 118, r: 118}: 2,
	},
	8: {
		yyinterval{l: 118, r: 118}: 2,
		yyinterval{l: 100, r: 100}: 2,
		yyinterval{l: 53, r: 53}:   2,
		yyinterval{l: 102, r: 102}: 2,
		yyinterval{l: 119, r: 119}: 2,
		yyinterval{l: 117, r: 117}: 2,
		yyinterval{l: 106, r: 107}: 2,
		yyinterval{l: 111, r: 111}: 2,
		yyinterval{l: 110, r: 110}: 2,
		yyinterval{l: 55, r: 57}:   2,
		yyinterval{l: 108, r: 108}: 2,
		yyinterval{l: 54, r: 54}:   2,
		yyinterval{l: 99, r: 99}:   2,
		yyinterval{l: 103, r: 103}: 2,
		yyinterval{l: 115, r: 115}: 2,
		yyinterval{l: 112, r: 113}: 2,
		yyinterval{l: 114, r: 114}: 2,
		yyinterval{l: 48, r: 48}:   2,
		yyinterval{l: 52, r: 52}:   2,
		yyinterval{l: 97, r: 97}:   2,
		yyinterval{l: 104, r: 104}: 2,
		yyinterval{l: 105, r: 105}: 2,
		yyinterval{l: 98, r: 98}:   2,
		yyinterval{l: 116, r: 116}: 2,
		yyinterval{l: 101, r: 101}: 10,
		yyinterval{l: 109, r: 109}: 2,
		yyinterval{l: 120, r: 122}: 2,
		yyinterval{l: 49, r: 51}:   2,
		yyinterval{l: 65, r: 90}:   2,
	},
	9: {
		yyinterval{l: 116, r: 116}: 2,
		yyinterval{l: 104, r: 104}: 2,
		yyinterval{l: 100, r: 100}: 2,
		yyinterval{l: 101, r: 101}: 2,
		yyinterval{l: 111, r: 111}: 2,
		yyinterval{l: 118, r: 118}: 2,
		yyinterval{l: 54, r: 54}:   2,
		yyinterval{l: 119, r: 119}: 2,
		yyinterval{l: 109, r: 109}: 2,
		yyinterval{l: 49, r: 51}:   2,
		yyinterval{l: 52, r: 52}:   2,
		yyinterval{l: 106, r: 107}: 2,
		yyinterval{l: 110, r: 110}: 10,
		yyinterval{l: 48, r: 48}:   2,
		yyinterval{l: 117, r: 117}: 2,
		yyinterval{l: 103, r: 103}: 2,
		yyinterval{l: 53, r: 53}:   2,
		yyinterval{l: 99, r: 99}:   2,
		yyinterval{l: 55, r: 57}:   2,
		yyinterval{l: 98, r: 98}:   2,
		yyinterval{l: 114, r: 114}: 2,
		yyinterval{l: 105, r: 105}: 2,
		yyinterval{l: 112, r: 113}: 2,
		yyinterval{l: 115, r: 115}: 2,
		yyinterval{l: 65, r: 90}:   2,
		yyinterval{l: 102, r: 102}: 2,
		yyinterval{l: 120, r: 122}: 2,
		yyinterval{l: 97, r: 97}:   2,
		yyinterval{l: 108, r: 108}: 2,
	},
	10: {
		yyinterval{l: 100, r: 100}: 2,
		yyinterval{l: 116, r: 116}: 2,
		yyinterval{l: 120, r: 122}: 2,
		yyinterval{l: 49, r: 51}:   2,
		yyinterval{l: 102, r: 102}: 2,
		yyinterval{l: 117, r: 117}: 2,
		yyinterval{l: 115, r: 115}: 2,
		yyinterval{l: 111, r: 111}: 2,
		yyinterval{l: 104, r: 104}: 2,
		yyinterval{l: 118, r: 118}: 2,
		yyinterval{l: 53, r: 53}:   2,
		yyinterval{l: 114, r: 114}: 2,
		yyinterval{l: 99, r: 99}:   2,
		yyinterval{l: 52, r: 52}:   2,
		yyinterval{l: 54, r: 54}:   2,
		yyinterval{l: 105, r: 105}: 2,
		yyinterval{l: 65, r: 90}:   2,
		yyinterval{l: 109, r: 109}: 2,
		yyinterval{l: 103, r: 103}: 2,
		yyinterval{l: 119, r: 119}: 2,
		yyinterval{l: 97, r: 97}:   2,
		yyinterval{l: 101, r: 101}: 2,
		yyinterval{l: 106, r: 107}: 2,
		yyinterval{l: 108, r: 108}: 2,
		yyinterval{l: 98, r: 98}:   2,
		yyinterval{l: 112, r: 113}: 2,
		yyinterval{l: 48, r: 48}:   2,
		yyinterval{l: 55, r: 57}:   2,
		yyinterval{l: 110, r: 110}: 2,
	},
	11: {
		yyinterval{l: 48, r: 48}:   2,
		yyinterval{l: 116, r: 116}: 14,
		yyinterval{l: 103, r: 103}: 2,
		yyinterval{l: 97, r: 97}:   2,
		yyinterval{l: 119, r: 119}: 2,
		yyinterval{l: 65, r: 90}:   2,
		yyinterval{l: 99, r: 99}:   2,
		yyinterval{l: 100, r: 100}: 2,
		yyinterval{l: 120, r: 122}: 2,
		yyinterval{l: 110, r: 110}: 2,
		yyinterval{l: 54, r: 54}:   2,
		yyinterval{l: 104, r: 104}: 2,
		yyinterval{l: 118, r: 118}: 2,
		yyinterval{l: 102, r: 102}: 2,
		yyinterval{l: 49, r: 51}:   2,
		yyinterval{l: 112, r: 113}: 2,
		yyinterval{l: 55, r: 57}:   2,
		yyinterval{l: 98, r: 98}:   2,
		yyinterval{l: 114, r: 114}: 2,
		yyinterval{l: 108, r: 108}: 2,
		yyinterval{l: 52, r: 52}:   2,
		yyinterval{l: 105, r: 105}: 2,
		yyinterval{l: 115, r: 115}: 2,
		yyinterval{l: 101, r: 101}: 2,
		yyinterval{l: 109, r: 109}: 2,
		yyinterval{l: 117, r: 117}: 2,
		yyinterval{l: 111, r: 111}: 2,
		yyinterval{l: 106, r: 107}: 2,
		yyinterval{l: 53, r: 53}:   2,
	},
	12: {
		yyinterval{l: 12353, r: 12436}: 12,
	},
	13: {
		yyinterval{l: 97, r: 97}:   2,
		yyinterval{l: 119, r: 119}: 2,
		yyinterval{l: 52, r: 52}:   2,
		yyinterval{l: 102, r: 102}: 2,
		yyinterval{l: 48, r: 48}:   2,
		yyinterval{l: 100, r: 100}: 2,
		yyinterval{l: 99, r: 99}:   2,
		yyinterval{l: 118, r: 118}: 2,
		yyinterval{l: 105, r: 105}: 2,
		yyinterval{l: 116, r: 116}: 2,
		yyinterval{l: 54, r: 54}:   2,
		yyinterval{l: 55, r: 57}:   2,
		yyinterval{l: 120, r: 122}: 2,
		yyinterval{l: 98, r: 98}:   2,
		yyinterval{l: 101, r: 101}: 2,
		yyinterval{l: 110, r: 110}: 2,
		yyinterval{l: 103, r: 103}: 2,
		yyinterval{l: 114, r: 114}: 9,
		yyinterval{l: 112, r: 113}: 2,
		yyinterval{l: 65, r: 90}:   2,
		yyinterval{l: 111, r: 111}: 2,
		yyinterval{l: 106, r: 107}: 2,
		yyinterval{l: 117, r: 117}: 2,
		yyinterval{l: 53, r: 53}:   2,
		yyinterval{l: 49, r: 51}:   2,
		yyinterval{l: 109, r: 109}: 2,
		yyinterval{l: 104, r: 104}: 2,
		yyinterval{l: 108, r: 108}: 2,
		yyinterval{l: 115, r: 115}: 2,
	},
	14: {
		yyinterval{l: 120, r: 122}: 2,
		yyinterval{l: 100, r: 100}: 2,
		yyinterval{l: 117, r: 117}: 2,
		yyinterval{l: 114, r: 114}: 2,
		yyinterval{l: 98, r: 98}:   2,
		yyinterval{l: 108, r: 108}: 2,
		yyinterval{l: 118, r: 118}: 2,
		yyinterval{l: 99, r: 99}:   2,
		yyinterval{l: 53, r: 53}:   2,
		yyinterval{l: 54, r: 54}:   2,
		yyinterval{l: 97, r: 97}:   2,
		yyinterval{l: 102, r: 102}: 2,
		yyinterval{l: 48, r: 48}:   2,
		yyinterval{l: 116, r: 116}: 2,
		yyinterval{l: 115, r: 115}: 2,
		yyinterval{l: 106, r: 107}: 2,
		yyinterval{l: 104, r: 104}: 2,
		yyinterval{l: 65, r: 90}:   2,
		yyinterval{l: 55, r: 57}:   2,
		yyinterval{l: 52, r: 52}:   2,
		yyinterval{l: 112, r: 113}: 2,
		yyinterval{l: 110, r: 110}: 2,
		yyinterval{l: 119, r: 119}: 2,
		yyinterval{l: 49, r: 51}:   2,
		yyinterval{l: 103, r: 103}: 2,
		yyinterval{l: 105, r: 105}: 2,
		yyinterval{l: 109, r: 109}: 2,
		yyinterval{l: 101, r: 101}: 2,
		yyinterval{l: 111, r: 111}: 2,
	},
	15: {
		yyinterval{l: 52, r: 52}:   2,
		yyinterval{l: 54, r: 54}:   2,
		yyinterval{l: 104, r: 104}: 2,
		yyinterval{l: 49, r: 51}:   2,
		yyinterval{l: 112, r: 113}: 2,
		yyinterval{l: 110, r: 110}: 2,
		yyinterval{l: 103, r: 103}: 2,
		yyinterval{l: 117, r: 117}: 2,
		yyinterval{l: 109, r: 109}: 2,
		yyinterval{l: 65, r: 90}:   2,
		yyinterval{l: 55, r: 57}:   2,
		yyinterval{l: 118, r: 118}: 2,
		yyinterval{l: 99, r: 99}:   2,
		yyinterval{l: 114, r: 114}: 2,
		yyinterval{l: 105, r: 105}: 2,
		yyinterval{l: 116, r: 116}: 18,
		yyinterval{l: 119, r: 119}: 2,
		yyinterval{l: 53, r: 53}:   2,
		yyinterval{l: 101, r: 101}: 2,
		yyinterval{l: 108, r: 108}: 2,
		yyinterval{l: 100, r: 100}: 2,
		yyinterval{l: 115, r: 115}: 2,
		yyinterval{l: 111, r: 111}: 2,
		yyinterval{l: 120, r: 122}: 2,
		yyinterval{l: 48, r: 48}:   2,
		yyinterval{l: 98, r: 98}:   2,
		yyinterval{l: 102, r: 102}: 2,
		yyinterval{l: 106, r: 107}: 2,
		yyinterval{l: 97, r: 97}:   2,
	},
	16: {
		yyinterval{l: 52, r: 52}: 16,
		yyinterval{l: 53, r: 53}: 16,
		yyinterval{l: 54, r: 54}: 16,
		yyinterval{l: 55, r: 57}: 16,
		yyinterval{l: 48, r: 48}: 16,
		yyinterval{l: 49, r: 51}: 16,
	},
	17: {
		yyinterval{l: 99, r: 99}:   2,
		yyinterval{l: 52, r: 52}:   2,
		yyinterval{l: 108, r: 108}: 2,
		yyinterval{l: 117, r: 117}: 2,
		yyinterval{l: 100, r: 100}: 2,
		yyinterval{l: 115, r: 115}: 2,
		yyinterval{l: 118, r: 118}: 2,
		yyinterval{l: 101, r: 101}: 2,
		yyinterval{l: 103, r: 103}: 2,
		yyinterval{l: 119, r: 119}: 2,
		yyinterval{l: 102, r: 102}: 2,
		yyinterval{l: 65, r: 90}:   2,
		yyinterval{l: 97, r: 97}:   2,
		yyinterval{l: 49, r: 51}:   2,
		yyinterval{l: 106, r: 107}: 2,
		yyinterval{l: 120, r: 122}: 2,
		yyinterval{l: 112, r: 113}: 2,
		yyinterval{l: 98, r: 98}:   2,
		yyinterval{l: 110, r: 110}: 2,
		yyinterval{l: 48, r: 48}:   2,
		yyinterval{l: 53, r: 53}:   2,
		yyinterval{l: 109, r: 109}: 2,
		yyinterval{l: 105, r: 105}: 2,
		yyinterval{l: 54, r: 54}:   2,
		yyinterval{l: 114, r: 114}: 2,
		yyinterval{l: 55, r: 57}:   2,
		yyinterval{l: 104, r: 104}: 2,
		yyinterval{l: 111, r: 111}: 2,
		yyinterval{l: 116, r: 116}: 4,
	},
	18: {
		yyinterval{l: 102, r: 102}: 2,
		yyinterval{l: 117, r: 117}: 2,
		yyinterval{l: 53, r: 53}:   2,
		yyinterval{l: 111, r: 111}: 2,
		yyinterval{l: 120, r: 122}: 2,
		yyinterval{l: 98, r: 98}:   2,
		yyinterval{l: 114, r: 114}: 2,
		yyinterval{l: 105, r: 105}: 2,
		yyinterval{l: 103, r: 103}: 2,
		yyinterval{l: 55, r: 57}:   2,
		yyinterval{l: 49, r: 51}:   2,
		yyinterval{l: 118, r: 118}: 2,
		yyinterval{l: 108, r: 108}: 2,
		yyinterval{l: 54, r: 54}:   5,
		yyinterval{l: 52, r: 52}:   2,
		yyinterval{l: 106, r: 107}: 2,
		yyinterval{l: 100, r: 100}: 2,
		yyinterval{l: 112, r: 113}: 2,
		yyinterval{l: 65, r: 90}:   2,
		yyinterval{l: 101, r: 101}: 2,
		yyinterval{l: 109, r: 109}: 2,
		yyinterval{l: 115, r: 115}: 2,
		yyinterval{l: 99, r: 99}:   2,
		yyinterval{l: 48, r: 48}:   2,
		yyinterval{l: 119, r: 119}: 2,
		yyinterval{l: 116, r: 116}: 2,
		yyinterval{l: 104, r: 104}: 2,
		yyinterval{l: 97, r: 97}:   2,
		yyinterval{l: 110, r: 110}: 2,
	},
	19: {
		yyinterval{l: 9, r: 9}:   19,
		yyinterval{l: 10, r: 10}: 19,
		yyinterval{l: 13, r: 13}: 19,
		yyinterval{l: 32, r: 32}: 19,
	},
	21: {
		yyinterval{l: 52, r: 52}:   2,
		yyinterval{l: 118, r: 118}: 2,
		yyinterval{l: 117, r: 117}: 2,
		yyinterval{l: 48, r: 48}:   2,
		yyinterval{l: 104, r: 104}: 2,
		yyinterval{l: 114, r: 114}: 2,
		yyinterval{l: 102, r: 102}: 2,
		yyinterval{l: 101, r: 101}: 17,
		yyinterval{l: 55, r: 57}:   2,
		yyinterval{l: 65, r: 90}:   2,
		yyinterval{l: 105, r: 105}: 2,
		yyinterval{l: 115, r: 115}: 2,
		yyinterval{l: 119, r: 119}: 2,
		yyinterval{l: 110, r: 110}: 2,
		yyinterval{l: 49, r: 51}:   2,
		yyinterval{l: 116, r: 116}: 2,
		yyinterval{l: 100, r: 100}: 2,
		yyinterval{l: 53, r: 53}:   2,
		yyinterval{l: 54, r: 54}:   2,
		yyinterval{l: 108, r: 108}: 2,
		yyinterval{l: 106, r: 107}: 2,
		yyinterval{l: 111, r: 111}: 2,
		yyinterval{l: 98, r: 98}:   2,
		yyinterval{l: 99, r: 99}:   2,
		yyinterval{l: 103, r: 103}: 2,
		yyinterval{l: 109, r: 109}: 2,
		yyinterval{l: 97, r: 97}:   2,
		yyinterval{l: 120, r: 122}: 2,
		yyinterval{l: 112, r: 113}: 2,
	},
	24: {
		yyinterval{l: 61, r: 61}: 3,
	},
	25: {
		yyinterval{l: 108, r: 108}: 2,
		yyinterval{l: 120, r: 122}: 2,
		yyinterval{l: 53, r: 53}:   2,
		yyinterval{l: 52, r: 52}:   2,
		yyinterval{l: 110, r: 110}: 2,
		yyinterval{l: 49, r: 51}:   2,
		yyinterval{l: 112, r: 113}: 2,
		yyinterval{l: 55, r: 57}:   2,
		yyinterval{l: 106, r: 107}: 2,
		yyinterval{l: 118, r: 118}: 2,
		yyinterval{l: 109, r: 109}: 2,
		yyinterval{l: 48, r: 48}:   2,
		yyinterval{l: 117, r: 117}: 2,
		yyinterval{l: 105, r: 105}: 2,
		yyinterval{l: 114, r: 114}: 2,
		yyinterval{l: 111, r: 111}: 2,
		yyinterval{l: 65, r: 90}:   2,
		yyinterval{l: 101, r: 101}: 2,
		yyinterval{l: 99, r: 99}:   2,
		yyinterval{l: 119, r: 119}: 2,
		yyinterval{l: 97, r: 97}:   2,
		yyinterval{l: 102, r: 102}: 2,
		yyinterval{l: 104, r: 104}: 34,
		yyinterval{l: 54, r: 54}:   2,
		yyinterval{l: 98, r: 98}:   2,
		yyinterval{l: 103, r: 103}: 2,
		yyinterval{l: 100, r: 100}: 2,
		yyinterval{l: 115, r: 115}: 2,
		yyinterval{l: 116, r: 116}: 2,
	},
	26: {
		yyinterval{l: 112, r: 113}: 2,
		yyinterval{l: 110, r: 110}: 2,
		yyinterval{l: 55, r: 57}:   2,
		yyinterval{l: 100, r: 100}: 2,
		yyinterval{l: 109, r: 109}: 2,
		yyinterval{l: 97, r: 97}:   2,
		yyinterval{l: 106, r: 107}: 2,
		yyinterval{l: 105, r: 105}: 2,
		yyinterval{l: 102, r: 102}: 2,
		yyinterval{l: 103, r: 103}: 2,
		yyinterval{l: 101, r: 101}: 2,
		yyinterval{l: 65, r: 90}:   2,
		yyinterval{l: 53, r: 53}:   2,
		yyinterval{l: 111, r: 111}: 32,
		yyinterval{l: 114, r: 114}: 2,
		yyinterval{l: 48, r: 48}:   2,
		yyinterval{l: 119, r: 119}: 2,
		yyinterval{l: 115, r: 115}: 2,
		yyinterval{l: 98, r: 98}:   2,
		yyinterval{l: 104, r: 104}: 2,
		yyinterval{l: 54, r: 54}:   2,
		yyinterval{l: 108, r: 108}: 33,
		yyinterval{l: 120, r: 122}: 2,
		yyinterval{l: 118, r: 118}: 2,
		yyinterval{l: 49, r: 51}:   2,
		yyinterval{l: 99, r: 99}:   2,
		yyinterval{l: 52, r: 52}:   2,
		yyinterval{l: 117, r: 117}: 31,
		yyinterval{l: 116, r: 116}: 2,
	},
	27: {
		yyinterval{l: 120, r: 122}: 2,
		yyinterval{l: 103, r: 103}: 2,
		yyinterval{l: 54, r: 54}:   2,
		yyinterval{l: 55, r: 57}:   2,
		yyinterval{l: 101, r: 101}: 2,
		yyinterval{l: 114, r: 114}: 2,
		yyinterval{l: 99, r: 99}:   2,
		yyinterval{l: 53, r: 53}:   2,
		yyinterval{l: 111, r: 111}: 2,
		yyinterval{l: 108, r: 108}: 2,
		yyinterval{l: 97, r: 97}:   2,
		yyinterval{l: 52, r: 52}:   2,
		yyinterval{l: 117, r: 117}: 2,
		yyinterval{l: 109, r: 109}: 2,
		yyinterval{l: 119, r: 119}: 2,
		yyinterval{l: 105, r: 105}: 2,
		yyinterval{l: 112, r: 113}: 2,
		yyinterval{l: 102, r: 102}: 10,
		yyinterval{l: 106, r: 107}: 2,
		yyinterval{l: 116, r: 116}: 2,
		yyinterval{l: 49, r: 51}:   2,
		yyinterval{l: 118, r: 118}: 2,
		yyinterval{l: 115, r: 115}: 2,
		yyinterval{l: 100, r: 100}: 2,
		yyinterval{l: 110, r: 110}: 11,
		yyinterval{l: 98, r: 98}:   2,
		yyinterval{l: 65, r: 90}:   2,
		yyinterval{l: 48, r: 48}:   2,
		yyinterval{l: 104, r: 104}: 2,
	},
	30: {
		yyinterval{l: 116, r: 116}: 2,
		yyinterval{l: 111, r: 111}: 2,
		yyinterval{l: 105, r: 105}: 2,
		yyinterval{l: 102, r: 102}: 2,
		yyinterval{l: 103, r: 103}: 2,
		yyinterval{l: 106, r: 107}: 2,
		yyinterval{l: 99, r: 99}:   2,
		yyinterval{l: 117, r: 117}: 2,
		yyinterval{l: 49, r: 51}:   2,
		yyinterval{l: 114, r: 114}: 2,
		yyinterval{l: 65, r: 90}:   2,
		yyinterval{l: 108, r: 108}: 8,
		yyinterval{l: 118, r: 118}: 2,
		yyinterval{l: 48, r: 48}:   2,
		yyinterval{l: 112, r: 113}: 2,
		yyinterval{l: 120, r: 122}: 2,
		yyinterval{l: 55, r: 57}:   2,
		yyinterval{l: 100, r: 100}: 2,
		yyinterval{l: 98, r: 98}:   2,
		yyinterval{l: 115, r: 115}: 2,
		yyinterval{l: 110, r: 110}: 2,
		yyinterval{l: 109, r: 109}: 2,
		yyinterval{l: 53, r: 53}:   2,
		yyinterval{l: 101, r: 101}: 2,
		yyinterval{l: 52, r: 52}:   2,
		yyinterval{l: 104, r: 104}: 2,
		yyinterval{l: 97, r: 97}:   2,
		yyinterval{l: 54, r: 54}:   2,
		yyinterval{l: 119, r: 119}: 2,
	},
	31: {
		yyinterval{l: 65, r: 90}:   2,
		yyinterval{l: 54, r: 54}:   2,
		yyinterval{l: 109, r: 109}: 2,
		yyinterval{l: 101, r: 101}: 2,
		yyinterval{l: 117, r: 117}: 2,
		yyinterval{l: 49, r: 51}:   2,
		yyinterval{l: 100, r: 100}: 2,
		yyinterval{l: 55, r: 57}:   2,
		yyinterval{l: 102, r: 102}: 2,
		yyinterval{l: 97, r: 97}:   2,
		yyinterval{l: 110, r: 110}: 7,
		yyinterval{l: 48, r: 48}:   2,
		yyinterval{l: 118, r: 118}: 2,
		yyinterval{l: 112, r: 113}: 2,
		yyinterval{l: 120, r: 122}: 2,
		yyinterval{l: 52, r: 52}:   2,
		yyinterval{l: 105, r: 105}: 2,
		yyinterval{l: 114, r: 114}: 2,
		yyinterval{l: 106, r: 107}: 2,
		yyinterval{l: 104, r: 104}: 2,
		yyinterval{l: 115, r: 115}: 2,
		yyinterval{l: 108, r: 108}: 2,
		yyinterval{l: 119, r: 119}: 2,
		yyinterval{l: 98, r: 98}:   2,
		yyinterval{l: 99, r: 99}:   2,
		yyinterval{l: 111, r: 111}: 2,
		yyinterval{l: 103, r: 103}: 2,
		yyinterval{l: 53, r: 53}:   2,
		yyinterval{l: 116, r: 116}: 2,
	},
	32: {
		yyinterval{l: 99, r: 99}:   2,
		yyinterval{l: 118, r: 118}: 2,
		yyinterval{l: 53, r: 53}:   2,
		yyinterval{l: 98, r: 98}:   2,
		yyinterval{l: 108, r: 108}: 2,
		yyinterval{l: 112, r: 113}: 2,
		yyinterval{l: 115, r: 115}: 2,
		yyinterval{l: 102, r: 102}: 2,
		yyinterval{l: 100, r: 100}: 2,
		yyinterval{l: 52, r: 52}:   2,
		yyinterval{l: 116, r: 116}: 2,
		yyinterval{l: 114, r: 114}: 10,
		yyinterval{l: 106, r: 107}: 2,
		yyinterval{l: 101, r: 101}: 2,
		yyinterval{l: 103, r: 103}: 2,
		yyinterval{l: 110, r: 110}: 2,
		yyinterval{l: 117, r: 117}: 2,
		yyinterval{l: 109, r: 109}: 2,
		yyinterval{l: 120, r: 122}: 2,
		yyinterval{l: 65, r: 90}:   2,
		yyinterval{l: 119, r: 119}: 2,
		yyinterval{l: 55, r: 57}:   2,
		yyinterval{l: 104, r: 104}: 2,
		yyinterval{l: 54, r: 54}:   2,
		yyinterval{l: 49, r: 51}:   2,
		yyinterval{l: 105, r: 105}: 2,
		yyinterval{l: 97, r: 97}:   2,
		yyinterval{l: 48, r: 48}:   2,
		yyinterval{l: 111, r: 111}: 2,
	},
	33: {
		yyinterval{l: 99, r: 99}:   2,
		yyinterval{l: 102, r: 102}: 2,
		yyinterval{l: 109, r: 109}: 2,
		yyinterval{l: 55, r: 57}:   2,
		yyinterval{l: 101, r: 101}: 2,
		yyinterval{l: 114, r: 114}: 2,
		yyinterval{l: 98, r: 98}:   2,
		yyinterval{l: 119, r: 119}: 2,
		yyinterval{l: 117, r: 117}: 2,
		yyinterval{l: 115, r: 115}: 2,
		yyinterval{l: 48, r: 48}:   2,
		yyinterval{l: 97, r: 97}:   2,
		yyinterval{l: 108, r: 108}: 2,
		yyinterval{l: 104, r: 104}: 2,
		yyinterval{l: 52, r: 52}:   2,
		yyinterval{l: 116, r: 116}: 2,
		yyinterval{l: 65, r: 90}:   2,
		yyinterval{l: 106, r: 107}: 2,
		yyinterval{l: 105, r: 105}: 2,
		yyinterval{l: 111, r: 111}: 6,
		yyinterval{l: 103, r: 103}: 2,
		yyinterval{l: 54, r: 54}:   2,
		yyinterval{l: 120, r: 122}: 2,
		yyinterval{l: 110, r: 110}: 2,
		yyinterval{l: 112, r: 113}: 2,
		yyinterval{l: 118, r: 118}: 2,
		yyinterval{l: 100, r: 100}: 2,
		yyinterval{l: 49, r: 51}:   2,
		yyinterval{l: 53, r: 53}:   2,
	},
	34: {
		yyinterval{l: 48, r: 48}:   2,
		yyinterval{l: 102, r: 102}: 2,
		yyinterval{l: 120, r: 122}: 2,
		yyinterval{l: 111, r: 111}: 2,
		yyinterval{l: 99, r: 99}:   2,
		yyinterval{l: 114, r: 114}: 2,
		yyinterval{l: 65, r: 90}:   2,
		yyinterval{l: 117, r: 117}: 2,
		yyinterval{l: 119, r: 119}: 2,
		yyinterval{l: 106, r: 107}: 2,
		yyinterval{l: 116, r: 116}: 2,
		yyinterval{l: 54, r: 54}:   2,
		yyinterval{l: 110, r: 110}: 2,
		yyinterval{l: 100, r: 100}: 2,
		yyinterval{l: 109, r: 109}: 2,
		yyinterval{l: 52, r: 52}:   2,
		yyinterval{l: 115, r: 115}: 2,
		yyinterval{l: 53, r: 53}:   2,
		yyinterval{l: 97, r: 97}:   2,
		yyinterval{l: 103, r: 103}: 2,
		yyinterval{l: 108, r: 108}: 2,
		yyinterval{l: 105, r: 105}: 30,
		yyinterval{l: 101, r: 101}: 2,
		yyinterval{l: 118, r: 118}: 2,
		yyinterval{l: 98, r: 98}:   2,
		yyinterval{l: 49, r: 51}:   2,
		yyinterval{l: 112, r: 113}: 2,
		yyinterval{l: 104, r: 104}: 2,
		yyinterval{l: 55, r: 57}:   2,
	},
}

func yyNextStep(id yyStateID, r rune) yyStateID {
	if mp, ok := yyTransitionTable[id]; ok {
		t := yyinterval{l: int(r), r: int(r)}
		for intv, sid := range mp {
			if intv.overlap(t) {
				return sid
			}
		}
	}

	return 0
}

type yyLexer struct {
	rs          RuneReadSeeker
	beginPos    int
	finPos      int
	currPos     int
	finRegexID  int
	currStateID yyStateID
	YYText      string
}

type RuneReadSeeker interface {
	io.ReadSeeker
	io.RuneScanner
}

func New(rs RuneReadSeeker) *yyLexer {
	return &yyLexer{
		rs:          rs,
		beginPos:    0,
		finPos:      0,
		currPos:     0,
		finRegexID:  0,
		currStateID: 1, // init state id is 1.
	}
}

func (yylex *yyLexer) currRune() (rune, int, error) {
	ru, size, err := yylex.rs.ReadRune()
	if err != nil {
		return 0, 0, err
	}
	if err := yylex.rs.UnreadRune(); err != nil {
		return 0, 0, err
	}
	return ru, size, nil
}

func (yylex *yyLexer) Next() (int, error) {
	yyEofCnt := 0
yystart:
	for {
		yyr, yysize, err := yylex.currRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				yyEofCnt++
				if yyEofCnt == 1 {
					goto finProcess
				}
			}
			return 0, err
		}
	finProcess:
		yyNxStID := yyNextStep(yylex.currStateID, yyr)
		if yyNxStID == 0 {
			if _, err := yylex.rs.Seek(int64(yylex.finPos), io.SeekStart); err != nil {
				return 0, err
			}
			_, lastSize, err := yylex.currRune()
			if err != nil {
				return 0, err
			}
			yydata := make([]byte, yylex.finPos+lastSize-yylex.beginPos)
			if _, err := yylex.rs.Seek(int64(yylex.beginPos), io.SeekStart); err != nil {
				return 0, err
			}
			if _, err := yylex.rs.Read(yydata); err != nil {
				return 0, err
			}
			yylex.YYText = string(yydata)
			YYText = yylex.YYText
			yyNewCurrPos := yylex.finPos + lastSize
			yylex.beginPos = yyNewCurrPos
			yylex.finPos = yyNewCurrPos
			yylex.currPos = yyNewCurrPos
			yylex.currStateID = 1

			regexID := yylex.finRegexID
			yylex.finRegexID = 0
			switch regexID {
			case 0:
				return 0, ErrYYScan
			case 1:
				{
					return Keyword, nil
				}
				goto yystart
			case 2:
				{
					return Type, nil
				}
				goto yystart
			case 3:
				{
					return Identifier, nil
				}
				goto yystart
			case 4:
				{
					return Digit, nil
				}
				goto yystart
			case 5:
				{
				}
				goto yystart
			case 6:
				{
					return LParen, nil
				}
				goto yystart
			case 7:
				{
					return RParen, nil
				}
				goto yystart
			case 8:
				{
					return LBracket, nil
				}
				goto yystart
			case 9:
				{
					return RBracket, nil
				}
				goto yystart
			case 10:
				{
					return Operator, nil
				}
				goto yystart
			case 11:
				{
					return Hiragana, nil
				}
				goto yystart
			case 12:
				{
				}
				goto yystart

			default:
				return 0, ErrYYScan
			}
		}
		if _, ok := yyFinStates[yyNxStID]; ok {
			yylex.finPos = yylex.currPos
			yylex.finRegexID = yyStateIDToRegexID[yyNxStID]
		}
		yylex.currStateID = yyNxStID
		yylex.currPos += yysize
		if _, err := yylex.rs.Seek(int64(yylex.currPos), io.SeekStart); err != nil {
			return 0, err
		}
	}

	return 0, io.EOF
}

// This part is optional
func main() {
	program := `
func foo123barあいう () int {
    x := 1 * 10 + 123 - 1000 / 5432
    y := float64(x)

    return x + y
}
`
	fmt.Println(program)
	fmt.Println("-----------------")

	lex := New(bytes.NewReader([]byte(program)))
	for {
		n, err := lex.Next()
		if err != nil {
			if err == io.EOF {
				return
			}
			panic(err)
		}
		switch n {
		case Keyword:
			fmt.Println("Keyword")
		case Hiragana:
			fmt.Println("Hiragana")
		case Type:
			fmt.Println("Type")
		case Identifier:
			fmt.Println("Identifier")
		case Digit:
			fmt.Println("Digit")
		case Whitespace:
			fmt.Println("Whitespace")
		case LParen:
			fmt.Println("LParen")
		case RParen:
			fmt.Println("RParen")
		case LBracket:
			fmt.Println("LBracket")
		case RBracket:
			fmt.Println("RBracket")
		case Operator:
			fmt.Println("Operator")
		}
		fmt.Printf("\t %#v\n", lex.YYText)
	}
}
