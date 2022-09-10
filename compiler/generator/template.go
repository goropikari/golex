package generator

const tmpl = `// Code generated by tlex. DO NOT EDIT.

package {{ .PackageName }}

import (
	"errors"
	"fmt"
	"io"
)

{{ .EmbeddedTmpl }}

type yyStateID = int
type yyRegexID = int
var YYText string

var (
	ErrYYScan = errors.New("failed to scan")
)

// state id to regex id
var yyStateIDToRegexID = []yyRegexID{
	0, // state 0 is dead state
	{{ .StateIDToRegexIDTmpl }}
}

var yyFinStates = map[yyStateID]struct{}{
	{{ .FinStatesTmpl }}
}

type yyinterval struct {
	l int
	r int
}

func (x yyinterval) overlap(y yyinterval) bool {
	return y.l <= x.r && x.l <= y.r
}

var yyTransitionTable = map[yyStateID]map[yyinterval]yyStateID{
	{{ .TransitionTableTmpl }}
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
	for  {
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
			{{ .RegexActionsTmpl }}
			default:
				return 0, ErrYYScan
			}
		}
		if _, ok := yyFinStates[yyNxStID]; ok {
			yylex.finPos = yylex.currPos
			yylex.finRegexID = yyStateIDToRegexID[yyNxStID]
		}
		yylex.currStateID = yyNxStID
		yylex.currPos+=yysize
		if _, err := yylex.rs.Seek(int64(yylex.currPos), io.SeekStart); err != nil {
			return 0, err
		}
	}

	return 0, io.EOF
}

{{ .UserCodeTmpl }}
`
