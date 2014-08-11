package main

import (
	"./stylesheet"
	"github.com/gorilla/css/scanner"
	"log"
)

type ParserState int

const (
	PARSE_STATE_NONE ParserState = iota
	PARSE_STATE_START_SELRCTOR
	PARSE_STATE_END_SELECTOR
	PARSE_STATE_START_DECLARE_BLOCK
	PARSE_STATE_END_DECLARE_BLOCK
	PARSE_STATE_START_PROPERTY
	PARSE_STATE_END_PROPERTY
	PARSE_STATE_START_VALUE
	PARSE_STATE_END_VALUE
)

type ParserConf struct {
	currentState ParserState
	cssRuleList  stylesheet.CSSRuleList
}

func (c *ParserConf) SetPaserState(state ParserState) {
	c.currentState = state
}

func (c *ParserConf) GetPaserState() ParserState {
	return c.currentState
}

/*
func main() {
	cssRuleList := make(stylesheet.CSSRuleList, 0)

	for i := 0; i < 256; i++ {
		rule := new(stylesheet.CSSRule)
		log.Println(i, ":", *rule)
		rule.CssType = i
		cssRuleList = append(cssRuleList, rule)
	}
	for i := 0; i < 256; i++ {
		log.Println(i, ":", *cssRuleList[i])
	}

	log.Println("size", len(cssRuleList))
	a := cssRuleList[0:10]
	log.Println("silce ", *a[9])

}
*/

func main() {
	conf := new(ParserConf)
	conf.cssRuleList = make(stylesheet.CSSRuleList, 0)

	s := scanner.New("div .a { font-size: 150%%} ")
	conf.SetPaserState(PARSE_STATE_NONE)

	for {
		token := s.Next()
		if token.Type == scanner.TokenEOF || token.Type == scanner.TokenError {
			break
		}
		switch token.Type {
		case scanner.TokenIdent:
			log.Println(token.Value)
		case scanner.TokenS:
			log.Println(token.Value)

		case scanner.TokenChar:
			if string('{') == token.Value {
				log.Println(token.Value)
			}

			if string('.') == token.Value {
				log.Println(token.Value)
			}
			if string('}') == token.Value {
				log.Println(token.Value)
			}
		case scanner.TokenPercentage:
			log.Println(token.Value)

		}

	}
}
