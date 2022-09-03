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
)

type yyStateID = int
type yyRegexID = int

var YYText string

var (
	ErrYYScan = errors.New("failed to scan")
)

// state id to regex id
var yyStateIDToRegexID = []yyRegexID{
	0, // state 0 は BH state
	5,
	11,
	11,
	3,
	3,
	1,
	3,
	5,
	4,
	3,
	7,
	6,
	3,
	10,
	3,
	8,
	9,
	2,
	3,
	3,
	3,
	3,
	3,
	3,
	3,
	3,
	3,
	3,
	3,
	3,
	3,
	3,
	3,
}

var yyFinStates = map[yyStateID]struct{}{
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
}

var yyTransitionTable = map[yyStateID]map[byte]yyStateID{
	1: {
		1:   2,
		2:   2,
		3:   2,
		4:   2,
		5:   2,
		6:   2,
		7:   2,
		8:   2,
		9:   8,
		10:  8,
		11:  2,
		12:  2,
		13:  8,
		14:  2,
		15:  2,
		16:  2,
		17:  2,
		18:  2,
		19:  2,
		20:  2,
		21:  2,
		22:  2,
		23:  2,
		24:  2,
		25:  2,
		26:  2,
		27:  2,
		28:  2,
		29:  2,
		30:  2,
		31:  2,
		32:  8,
		33:  3,
		34:  2,
		35:  2,
		36:  2,
		37:  2,
		38:  2,
		39:  2,
		40:  12,
		41:  11,
		42:  14,
		43:  14,
		44:  2,
		45:  14,
		46:  2,
		47:  14,
		48:  2,
		49:  9,
		50:  9,
		51:  9,
		52:  9,
		53:  9,
		54:  9,
		55:  9,
		56:  9,
		57:  9,
		58:  3,
		59:  2,
		60:  2,
		61:  3,
		62:  2,
		63:  2,
		64:  2,
		65:  4,
		66:  4,
		67:  4,
		68:  4,
		69:  4,
		70:  4,
		71:  4,
		72:  4,
		73:  4,
		74:  4,
		75:  4,
		76:  4,
		77:  4,
		78:  4,
		79:  4,
		80:  4,
		81:  4,
		82:  4,
		83:  4,
		84:  4,
		85:  4,
		86:  4,
		87:  4,
		88:  4,
		89:  4,
		90:  4,
		91:  2,
		92:  2,
		93:  2,
		94:  2,
		95:  2,
		96:  2,
		97:  4,
		98:  4,
		99:  4,
		100: 4,
		101: 4,
		102: 30,
		103: 4,
		104: 4,
		105: 24,
		106: 4,
		107: 4,
		108: 4,
		109: 4,
		110: 4,
		111: 4,
		112: 4,
		113: 4,
		114: 26,
		115: 4,
		116: 4,
		117: 4,
		118: 4,
		119: 25,
		120: 4,
		121: 4,
		122: 4,
		123: 16,
		124: 2,
		125: 17,
		126: 2,
		127: 2,
	},
	3: {
		61: 14,
	},
	4: {
		48:  4,
		49:  4,
		50:  4,
		51:  4,
		52:  4,
		53:  4,
		54:  4,
		55:  4,
		56:  4,
		57:  4,
		65:  4,
		66:  4,
		67:  4,
		68:  4,
		69:  4,
		70:  4,
		71:  4,
		72:  4,
		73:  4,
		74:  4,
		75:  4,
		76:  4,
		77:  4,
		78:  4,
		79:  4,
		80:  4,
		81:  4,
		82:  4,
		83:  4,
		84:  4,
		85:  4,
		86:  4,
		87:  4,
		88:  4,
		89:  4,
		90:  4,
		97:  4,
		98:  4,
		99:  4,
		100: 4,
		101: 4,
		102: 4,
		103: 4,
		104: 4,
		105: 4,
		106: 4,
		107: 4,
		108: 4,
		109: 4,
		110: 4,
		111: 4,
		112: 4,
		113: 4,
		114: 4,
		115: 4,
		116: 4,
		117: 4,
		118: 4,
		119: 4,
		120: 4,
		121: 4,
		122: 4,
	},
	5: {
		48:  4,
		49:  4,
		50:  4,
		51:  4,
		52:  4,
		53:  4,
		54:  4,
		55:  4,
		56:  4,
		57:  4,
		65:  4,
		66:  4,
		67:  4,
		68:  4,
		69:  4,
		70:  4,
		71:  4,
		72:  4,
		73:  4,
		74:  4,
		75:  4,
		76:  4,
		77:  4,
		78:  4,
		79:  4,
		80:  4,
		81:  4,
		82:  4,
		83:  4,
		84:  4,
		85:  4,
		86:  4,
		87:  4,
		88:  4,
		89:  4,
		90:  4,
		97:  4,
		98:  4,
		99:  4,
		100: 4,
		101: 4,
		102: 4,
		103: 4,
		104: 4,
		105: 4,
		106: 4,
		107: 4,
		108: 4,
		109: 4,
		110: 4,
		111: 4,
		112: 4,
		113: 4,
		114: 6,
		115: 4,
		116: 4,
		117: 4,
		118: 4,
		119: 4,
		120: 4,
		121: 4,
		122: 4,
	},
	6: {
		48:  4,
		49:  4,
		50:  4,
		51:  4,
		52:  4,
		53:  4,
		54:  4,
		55:  4,
		56:  4,
		57:  4,
		65:  4,
		66:  4,
		67:  4,
		68:  4,
		69:  4,
		70:  4,
		71:  4,
		72:  4,
		73:  4,
		74:  4,
		75:  4,
		76:  4,
		77:  4,
		78:  4,
		79:  4,
		80:  4,
		81:  4,
		82:  4,
		83:  4,
		84:  4,
		85:  4,
		86:  4,
		87:  4,
		88:  4,
		89:  4,
		90:  4,
		97:  4,
		98:  4,
		99:  4,
		100: 4,
		101: 4,
		102: 4,
		103: 4,
		104: 4,
		105: 4,
		106: 4,
		107: 4,
		108: 4,
		109: 4,
		110: 4,
		111: 4,
		112: 4,
		113: 4,
		114: 4,
		115: 4,
		116: 4,
		117: 4,
		118: 4,
		119: 4,
		120: 4,
		121: 4,
		122: 4,
	},
	7: {
		48:  4,
		49:  4,
		50:  4,
		51:  4,
		52:  4,
		53:  4,
		54:  4,
		55:  4,
		56:  4,
		57:  4,
		65:  4,
		66:  4,
		67:  4,
		68:  4,
		69:  4,
		70:  4,
		71:  4,
		72:  4,
		73:  4,
		74:  4,
		75:  4,
		76:  4,
		77:  4,
		78:  4,
		79:  4,
		80:  4,
		81:  4,
		82:  4,
		83:  4,
		84:  4,
		85:  4,
		86:  4,
		87:  4,
		88:  4,
		89:  4,
		90:  4,
		97:  29,
		98:  4,
		99:  4,
		100: 4,
		101: 4,
		102: 4,
		103: 4,
		104: 4,
		105: 4,
		106: 4,
		107: 4,
		108: 4,
		109: 4,
		110: 4,
		111: 4,
		112: 4,
		113: 4,
		114: 4,
		115: 4,
		116: 4,
		117: 4,
		118: 4,
		119: 4,
		120: 4,
		121: 4,
		122: 4,
	},
	8: {
		9:  8,
		10: 8,
		13: 8,
		32: 8,
	},
	9: {
		48: 9,
		49: 9,
		50: 9,
		51: 9,
		52: 9,
		53: 9,
		54: 9,
		55: 9,
		56: 9,
		57: 9,
	},
	10: {
		48:  4,
		49:  4,
		50:  4,
		51:  4,
		52:  4,
		53:  4,
		54:  4,
		55:  4,
		56:  4,
		57:  4,
		65:  4,
		66:  4,
		67:  4,
		68:  4,
		69:  4,
		70:  4,
		71:  4,
		72:  4,
		73:  4,
		74:  4,
		75:  4,
		76:  4,
		77:  4,
		78:  4,
		79:  4,
		80:  4,
		81:  4,
		82:  4,
		83:  4,
		84:  4,
		85:  4,
		86:  4,
		87:  4,
		88:  4,
		89:  4,
		90:  4,
		97:  4,
		98:  4,
		99:  4,
		100: 4,
		101: 4,
		102: 4,
		103: 4,
		104: 4,
		105: 4,
		106: 4,
		107: 4,
		108: 4,
		109: 4,
		110: 28,
		111: 4,
		112: 4,
		113: 4,
		114: 4,
		115: 4,
		116: 4,
		117: 4,
		118: 4,
		119: 4,
		120: 4,
		121: 4,
		122: 4,
	},
	13: {
		48:  4,
		49:  4,
		50:  4,
		51:  4,
		52:  18,
		53:  4,
		54:  4,
		55:  4,
		56:  4,
		57:  4,
		65:  4,
		66:  4,
		67:  4,
		68:  4,
		69:  4,
		70:  4,
		71:  4,
		72:  4,
		73:  4,
		74:  4,
		75:  4,
		76:  4,
		77:  4,
		78:  4,
		79:  4,
		80:  4,
		81:  4,
		82:  4,
		83:  4,
		84:  4,
		85:  4,
		86:  4,
		87:  4,
		88:  4,
		89:  4,
		90:  4,
		97:  4,
		98:  4,
		99:  4,
		100: 4,
		101: 4,
		102: 4,
		103: 4,
		104: 4,
		105: 4,
		106: 4,
		107: 4,
		108: 4,
		109: 4,
		110: 4,
		111: 4,
		112: 4,
		113: 4,
		114: 4,
		115: 4,
		116: 4,
		117: 4,
		118: 4,
		119: 4,
		120: 4,
		121: 4,
		122: 4,
	},
	15: {
		48:  4,
		49:  4,
		50:  4,
		51:  4,
		52:  4,
		53:  4,
		54:  4,
		55:  4,
		56:  4,
		57:  4,
		65:  4,
		66:  4,
		67:  4,
		68:  4,
		69:  4,
		70:  4,
		71:  4,
		72:  4,
		73:  4,
		74:  4,
		75:  4,
		76:  4,
		77:  4,
		78:  4,
		79:  4,
		80:  4,
		81:  4,
		82:  4,
		83:  4,
		84:  4,
		85:  4,
		86:  4,
		87:  4,
		88:  4,
		89:  4,
		90:  4,
		97:  4,
		98:  4,
		99:  4,
		100: 4,
		101: 6,
		102: 4,
		103: 4,
		104: 4,
		105: 4,
		106: 4,
		107: 4,
		108: 4,
		109: 4,
		110: 4,
		111: 4,
		112: 4,
		113: 4,
		114: 4,
		115: 4,
		116: 4,
		117: 4,
		118: 4,
		119: 4,
		120: 4,
		121: 4,
		122: 4,
	},
	18: {
		48:  4,
		49:  4,
		50:  4,
		51:  4,
		52:  4,
		53:  4,
		54:  4,
		55:  4,
		56:  4,
		57:  4,
		65:  4,
		66:  4,
		67:  4,
		68:  4,
		69:  4,
		70:  4,
		71:  4,
		72:  4,
		73:  4,
		74:  4,
		75:  4,
		76:  4,
		77:  4,
		78:  4,
		79:  4,
		80:  4,
		81:  4,
		82:  4,
		83:  4,
		84:  4,
		85:  4,
		86:  4,
		87:  4,
		88:  4,
		89:  4,
		90:  4,
		97:  4,
		98:  4,
		99:  4,
		100: 4,
		101: 4,
		102: 4,
		103: 4,
		104: 4,
		105: 4,
		106: 4,
		107: 4,
		108: 4,
		109: 4,
		110: 4,
		111: 4,
		112: 4,
		113: 4,
		114: 4,
		115: 4,
		116: 4,
		117: 4,
		118: 4,
		119: 4,
		120: 4,
		121: 4,
		122: 4,
	},
	19: {
		48:  4,
		49:  4,
		50:  4,
		51:  4,
		52:  4,
		53:  4,
		54:  4,
		55:  4,
		56:  4,
		57:  4,
		65:  4,
		66:  4,
		67:  4,
		68:  4,
		69:  4,
		70:  4,
		71:  4,
		72:  4,
		73:  4,
		74:  4,
		75:  4,
		76:  4,
		77:  4,
		78:  4,
		79:  4,
		80:  4,
		81:  4,
		82:  4,
		83:  4,
		84:  4,
		85:  4,
		86:  4,
		87:  4,
		88:  4,
		89:  4,
		90:  4,
		97:  4,
		98:  4,
		99:  4,
		100: 4,
		101: 4,
		102: 4,
		103: 4,
		104: 4,
		105: 4,
		106: 4,
		107: 4,
		108: 4,
		109: 4,
		110: 4,
		111: 4,
		112: 4,
		113: 4,
		114: 4,
		115: 4,
		116: 4,
		117: 21,
		118: 4,
		119: 4,
		120: 4,
		121: 4,
		122: 4,
	},
	20: {
		48:  4,
		49:  4,
		50:  4,
		51:  4,
		52:  4,
		53:  4,
		54:  4,
		55:  4,
		56:  4,
		57:  4,
		65:  4,
		66:  4,
		67:  4,
		68:  4,
		69:  4,
		70:  4,
		71:  4,
		72:  4,
		73:  4,
		74:  4,
		75:  4,
		76:  4,
		77:  4,
		78:  4,
		79:  4,
		80:  4,
		81:  4,
		82:  4,
		83:  4,
		84:  4,
		85:  4,
		86:  4,
		87:  4,
		88:  4,
		89:  4,
		90:  4,
		97:  4,
		98:  4,
		99:  4,
		100: 4,
		101: 4,
		102: 4,
		103: 4,
		104: 4,
		105: 4,
		106: 4,
		107: 4,
		108: 4,
		109: 4,
		110: 4,
		111: 4,
		112: 4,
		113: 4,
		114: 4,
		115: 4,
		116: 19,
		117: 4,
		118: 4,
		119: 4,
		120: 4,
		121: 4,
		122: 4,
	},
	21: {
		48:  4,
		49:  4,
		50:  4,
		51:  4,
		52:  4,
		53:  4,
		54:  4,
		55:  4,
		56:  4,
		57:  4,
		65:  4,
		66:  4,
		67:  4,
		68:  4,
		69:  4,
		70:  4,
		71:  4,
		72:  4,
		73:  4,
		74:  4,
		75:  4,
		76:  4,
		77:  4,
		78:  4,
		79:  4,
		80:  4,
		81:  4,
		82:  4,
		83:  4,
		84:  4,
		85:  4,
		86:  4,
		87:  4,
		88:  4,
		89:  4,
		90:  4,
		97:  4,
		98:  4,
		99:  4,
		100: 4,
		101: 4,
		102: 4,
		103: 4,
		104: 4,
		105: 4,
		106: 4,
		107: 4,
		108: 4,
		109: 4,
		110: 4,
		111: 4,
		112: 4,
		113: 4,
		114: 27,
		115: 4,
		116: 4,
		117: 4,
		118: 4,
		119: 4,
		120: 4,
		121: 4,
		122: 4,
	},
	22: {
		48:  4,
		49:  4,
		50:  4,
		51:  4,
		52:  4,
		53:  4,
		54:  4,
		55:  4,
		56:  4,
		57:  4,
		65:  4,
		66:  4,
		67:  4,
		68:  4,
		69:  4,
		70:  4,
		71:  4,
		72:  4,
		73:  4,
		74:  4,
		75:  4,
		76:  4,
		77:  4,
		78:  4,
		79:  4,
		80:  4,
		81:  4,
		82:  4,
		83:  4,
		84:  4,
		85:  4,
		86:  4,
		87:  4,
		88:  4,
		89:  4,
		90:  4,
		97:  4,
		98:  4,
		99:  4,
		100: 4,
		101: 4,
		102: 4,
		103: 4,
		104: 4,
		105: 33,
		106: 4,
		107: 4,
		108: 4,
		109: 4,
		110: 4,
		111: 4,
		112: 4,
		113: 4,
		114: 4,
		115: 4,
		116: 4,
		117: 4,
		118: 4,
		119: 4,
		120: 4,
		121: 4,
		122: 4,
	},
	23: {
		48:  4,
		49:  4,
		50:  4,
		51:  4,
		52:  4,
		53:  4,
		54:  13,
		55:  4,
		56:  4,
		57:  4,
		65:  4,
		66:  4,
		67:  4,
		68:  4,
		69:  4,
		70:  4,
		71:  4,
		72:  4,
		73:  4,
		74:  4,
		75:  4,
		76:  4,
		77:  4,
		78:  4,
		79:  4,
		80:  4,
		81:  4,
		82:  4,
		83:  4,
		84:  4,
		85:  4,
		86:  4,
		87:  4,
		88:  4,
		89:  4,
		90:  4,
		97:  4,
		98:  4,
		99:  4,
		100: 4,
		101: 4,
		102: 4,
		103: 4,
		104: 4,
		105: 4,
		106: 4,
		107: 4,
		108: 4,
		109: 4,
		110: 4,
		111: 4,
		112: 4,
		113: 4,
		114: 4,
		115: 4,
		116: 4,
		117: 4,
		118: 4,
		119: 4,
		120: 4,
		121: 4,
		122: 4,
	},
	24: {
		48:  4,
		49:  4,
		50:  4,
		51:  4,
		52:  4,
		53:  4,
		54:  4,
		55:  4,
		56:  4,
		57:  4,
		65:  4,
		66:  4,
		67:  4,
		68:  4,
		69:  4,
		70:  4,
		71:  4,
		72:  4,
		73:  4,
		74:  4,
		75:  4,
		76:  4,
		77:  4,
		78:  4,
		79:  4,
		80:  4,
		81:  4,
		82:  4,
		83:  4,
		84:  4,
		85:  4,
		86:  4,
		87:  4,
		88:  4,
		89:  4,
		90:  4,
		97:  4,
		98:  4,
		99:  4,
		100: 4,
		101: 4,
		102: 6,
		103: 4,
		104: 4,
		105: 4,
		106: 4,
		107: 4,
		108: 4,
		109: 4,
		110: 31,
		111: 4,
		112: 4,
		113: 4,
		114: 4,
		115: 4,
		116: 4,
		117: 4,
		118: 4,
		119: 4,
		120: 4,
		121: 4,
		122: 4,
	},
	25: {
		48:  4,
		49:  4,
		50:  4,
		51:  4,
		52:  4,
		53:  4,
		54:  4,
		55:  4,
		56:  4,
		57:  4,
		65:  4,
		66:  4,
		67:  4,
		68:  4,
		69:  4,
		70:  4,
		71:  4,
		72:  4,
		73:  4,
		74:  4,
		75:  4,
		76:  4,
		77:  4,
		78:  4,
		79:  4,
		80:  4,
		81:  4,
		82:  4,
		83:  4,
		84:  4,
		85:  4,
		86:  4,
		87:  4,
		88:  4,
		89:  4,
		90:  4,
		97:  4,
		98:  4,
		99:  4,
		100: 4,
		101: 4,
		102: 4,
		103: 4,
		104: 22,
		105: 4,
		106: 4,
		107: 4,
		108: 4,
		109: 4,
		110: 4,
		111: 4,
		112: 4,
		113: 4,
		114: 4,
		115: 4,
		116: 4,
		117: 4,
		118: 4,
		119: 4,
		120: 4,
		121: 4,
		122: 4,
	},
	26: {
		48:  4,
		49:  4,
		50:  4,
		51:  4,
		52:  4,
		53:  4,
		54:  4,
		55:  4,
		56:  4,
		57:  4,
		65:  4,
		66:  4,
		67:  4,
		68:  4,
		69:  4,
		70:  4,
		71:  4,
		72:  4,
		73:  4,
		74:  4,
		75:  4,
		76:  4,
		77:  4,
		78:  4,
		79:  4,
		80:  4,
		81:  4,
		82:  4,
		83:  4,
		84:  4,
		85:  4,
		86:  4,
		87:  4,
		88:  4,
		89:  4,
		90:  4,
		97:  4,
		98:  4,
		99:  4,
		100: 4,
		101: 20,
		102: 4,
		103: 4,
		104: 4,
		105: 4,
		106: 4,
		107: 4,
		108: 4,
		109: 4,
		110: 4,
		111: 4,
		112: 4,
		113: 4,
		114: 4,
		115: 4,
		116: 4,
		117: 4,
		118: 4,
		119: 4,
		120: 4,
		121: 4,
		122: 4,
	},
	27: {
		48:  4,
		49:  4,
		50:  4,
		51:  4,
		52:  4,
		53:  4,
		54:  4,
		55:  4,
		56:  4,
		57:  4,
		65:  4,
		66:  4,
		67:  4,
		68:  4,
		69:  4,
		70:  4,
		71:  4,
		72:  4,
		73:  4,
		74:  4,
		75:  4,
		76:  4,
		77:  4,
		78:  4,
		79:  4,
		80:  4,
		81:  4,
		82:  4,
		83:  4,
		84:  4,
		85:  4,
		86:  4,
		87:  4,
		88:  4,
		89:  4,
		90:  4,
		97:  4,
		98:  4,
		99:  4,
		100: 4,
		101: 4,
		102: 4,
		103: 4,
		104: 4,
		105: 4,
		106: 4,
		107: 4,
		108: 4,
		109: 4,
		110: 6,
		111: 4,
		112: 4,
		113: 4,
		114: 4,
		115: 4,
		116: 4,
		117: 4,
		118: 4,
		119: 4,
		120: 4,
		121: 4,
		122: 4,
	},
	28: {
		48:  4,
		49:  4,
		50:  4,
		51:  4,
		52:  4,
		53:  4,
		54:  4,
		55:  4,
		56:  4,
		57:  4,
		65:  4,
		66:  4,
		67:  4,
		68:  4,
		69:  4,
		70:  4,
		71:  4,
		72:  4,
		73:  4,
		74:  4,
		75:  4,
		76:  4,
		77:  4,
		78:  4,
		79:  4,
		80:  4,
		81:  4,
		82:  4,
		83:  4,
		84:  4,
		85:  4,
		86:  4,
		87:  4,
		88:  4,
		89:  4,
		90:  4,
		97:  4,
		98:  4,
		99:  6,
		100: 4,
		101: 4,
		102: 4,
		103: 4,
		104: 4,
		105: 4,
		106: 4,
		107: 4,
		108: 4,
		109: 4,
		110: 4,
		111: 4,
		112: 4,
		113: 4,
		114: 4,
		115: 4,
		116: 4,
		117: 4,
		118: 4,
		119: 4,
		120: 4,
		121: 4,
		122: 4,
	},
	29: {
		48:  4,
		49:  4,
		50:  4,
		51:  4,
		52:  4,
		53:  4,
		54:  4,
		55:  4,
		56:  4,
		57:  4,
		65:  4,
		66:  4,
		67:  4,
		68:  4,
		69:  4,
		70:  4,
		71:  4,
		72:  4,
		73:  4,
		74:  4,
		75:  4,
		76:  4,
		77:  4,
		78:  4,
		79:  4,
		80:  4,
		81:  4,
		82:  4,
		83:  4,
		84:  4,
		85:  4,
		86:  4,
		87:  4,
		88:  4,
		89:  4,
		90:  4,
		97:  4,
		98:  4,
		99:  4,
		100: 4,
		101: 4,
		102: 4,
		103: 4,
		104: 4,
		105: 4,
		106: 4,
		107: 4,
		108: 4,
		109: 4,
		110: 4,
		111: 4,
		112: 4,
		113: 4,
		114: 4,
		115: 4,
		116: 23,
		117: 4,
		118: 4,
		119: 4,
		120: 4,
		121: 4,
		122: 4,
	},
	30: {
		48:  4,
		49:  4,
		50:  4,
		51:  4,
		52:  4,
		53:  4,
		54:  4,
		55:  4,
		56:  4,
		57:  4,
		65:  4,
		66:  4,
		67:  4,
		68:  4,
		69:  4,
		70:  4,
		71:  4,
		72:  4,
		73:  4,
		74:  4,
		75:  4,
		76:  4,
		77:  4,
		78:  4,
		79:  4,
		80:  4,
		81:  4,
		82:  4,
		83:  4,
		84:  4,
		85:  4,
		86:  4,
		87:  4,
		88:  4,
		89:  4,
		90:  4,
		97:  4,
		98:  4,
		99:  4,
		100: 4,
		101: 4,
		102: 4,
		103: 4,
		104: 4,
		105: 4,
		106: 4,
		107: 4,
		108: 32,
		109: 4,
		110: 4,
		111: 5,
		112: 4,
		113: 4,
		114: 4,
		115: 4,
		116: 4,
		117: 10,
		118: 4,
		119: 4,
		120: 4,
		121: 4,
		122: 4,
	},
	31: {
		48:  4,
		49:  4,
		50:  4,
		51:  4,
		52:  4,
		53:  4,
		54:  4,
		55:  4,
		56:  4,
		57:  4,
		65:  4,
		66:  4,
		67:  4,
		68:  4,
		69:  4,
		70:  4,
		71:  4,
		72:  4,
		73:  4,
		74:  4,
		75:  4,
		76:  4,
		77:  4,
		78:  4,
		79:  4,
		80:  4,
		81:  4,
		82:  4,
		83:  4,
		84:  4,
		85:  4,
		86:  4,
		87:  4,
		88:  4,
		89:  4,
		90:  4,
		97:  4,
		98:  4,
		99:  4,
		100: 4,
		101: 4,
		102: 4,
		103: 4,
		104: 4,
		105: 4,
		106: 4,
		107: 4,
		108: 4,
		109: 4,
		110: 4,
		111: 4,
		112: 4,
		113: 4,
		114: 4,
		115: 4,
		116: 18,
		117: 4,
		118: 4,
		119: 4,
		120: 4,
		121: 4,
		122: 4,
	},
	32: {
		48:  4,
		49:  4,
		50:  4,
		51:  4,
		52:  4,
		53:  4,
		54:  4,
		55:  4,
		56:  4,
		57:  4,
		65:  4,
		66:  4,
		67:  4,
		68:  4,
		69:  4,
		70:  4,
		71:  4,
		72:  4,
		73:  4,
		74:  4,
		75:  4,
		76:  4,
		77:  4,
		78:  4,
		79:  4,
		80:  4,
		81:  4,
		82:  4,
		83:  4,
		84:  4,
		85:  4,
		86:  4,
		87:  4,
		88:  4,
		89:  4,
		90:  4,
		97:  4,
		98:  4,
		99:  4,
		100: 4,
		101: 4,
		102: 4,
		103: 4,
		104: 4,
		105: 4,
		106: 4,
		107: 4,
		108: 4,
		109: 4,
		110: 4,
		111: 7,
		112: 4,
		113: 4,
		114: 4,
		115: 4,
		116: 4,
		117: 4,
		118: 4,
		119: 4,
		120: 4,
		121: 4,
		122: 4,
	},
	33: {
		48:  4,
		49:  4,
		50:  4,
		51:  4,
		52:  4,
		53:  4,
		54:  4,
		55:  4,
		56:  4,
		57:  4,
		65:  4,
		66:  4,
		67:  4,
		68:  4,
		69:  4,
		70:  4,
		71:  4,
		72:  4,
		73:  4,
		74:  4,
		75:  4,
		76:  4,
		77:  4,
		78:  4,
		79:  4,
		80:  4,
		81:  4,
		82:  4,
		83:  4,
		84:  4,
		85:  4,
		86:  4,
		87:  4,
		88:  4,
		89:  4,
		90:  4,
		97:  4,
		98:  4,
		99:  4,
		100: 4,
		101: 4,
		102: 4,
		103: 4,
		104: 4,
		105: 4,
		106: 4,
		107: 4,
		108: 15,
		109: 4,
		110: 4,
		111: 4,
		112: 4,
		113: 4,
		114: 4,
		115: 4,
		116: 4,
		117: 4,
		118: 4,
		119: 4,
		120: 4,
		121: 4,
		122: 4,
	},
}

func yyNextStep(id yyStateID, b byte) yyStateID {
	if mp, ok := yyTransitionTable[id]; ok {
		return mp[b]
	}

	return 0
}

type yyLexer struct {
	rs          io.ReadSeeker
	beginPos    int
	finPos      int
	currPos     int
	finRegexID  int
	currStateID yyStateID
	YYText      string
}

func New(rs io.ReadSeeker) *yyLexer {
	return &yyLexer{
		rs:          rs,
		beginPos:    0,
		finPos:      0,
		currPos:     0,
		finRegexID:  0,
		currStateID: 1, // init state id is 1.
	}
}

func (yylex *yyLexer) currByte() (byte, error) {
	b := make([]byte, 1)
	if _, err := yylex.rs.Read(b); err != nil {
		return 0, err
	}
	if _, err := yylex.rs.Seek(int64(yylex.currPos), io.SeekStart); err != nil {
		return 0, err
	}

	return b[0], nil
}

func (yylex *yyLexer) Next() (int, error) {
	yyEofCnt := 0
yystart:
	for {
		yyb, err := yylex.currByte()
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
		yyNxStID := yyNextStep(yylex.currStateID, yyb)
		if yyNxStID == 0 {
			yydata := make([]byte, yylex.finPos+1-yylex.beginPos)
			if _, err := yylex.rs.Seek(int64(yylex.beginPos), io.SeekStart); err != nil {
				return 0, err
			}
			if _, err := yylex.rs.Read(yydata); err != nil {
				return 0, err
			}
			yylex.YYText = string(yydata)
			YYText = yylex.YYText
			yyNewCurrPos := yylex.finPos + 1
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
		yylex.currPos++
		if _, err := yylex.rs.Seek(int64(yylex.currPos), io.SeekStart); err != nil {
			return 0, err
		}
	}

	return 0, io.EOF
}

// This part is optional
func main() {
	program := `
func foo123bar() int {
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
