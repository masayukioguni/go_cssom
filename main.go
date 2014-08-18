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
	stylesheet   stylesheet.CSSStyleSheet

	NowSelectorText string
	NowProperty     string
	NowValue        string
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
	conf.stylesheet = make(stylesheet.CSSStyleSheet, 0)

	s := scanner.New("div .a { font-size: 150%}")
	conf.SetPaserState(PARSE_STATE_NONE)

	for {
		token := s.Next()
		//log.Println(token)

		if token.Type == scanner.TokenEOF || token.Type == scanner.TokenError {
			break
		}
		switch token.Type {
		case scanner.TokenIdent:
			//log.Println(token.Value)
			if conf.GetPaserState() == PARSE_STATE_NONE ||
				conf.GetPaserState() == PARSE_STATE_START_SELRCTOR ||
				conf.GetPaserState() == PARSE_STATE_END_DECLARE_BLOCK {
				conf.SetPaserState(PARSE_STATE_START_SELRCTOR)

				conf.NowSelectorText += token.Value

			}

			if conf.GetPaserState() == PARSE_STATE_START_DECLARE_BLOCK {
				conf.NowProperty = token.Value
			}

		case scanner.TokenS:
			if conf.GetPaserState() == PARSE_STATE_START_SELRCTOR {
				if string(' ') == token.Value {
					//log.Println(token.Value)
					conf.NowSelectorText += token.Value
				}
			}

		case scanner.TokenChar:

			if conf.GetPaserState() == PARSE_STATE_START_SELRCTOR {
				if string('.') == token.Value {
					//log.Println(token.Value)
					conf.NowSelectorText += token.Value
				}
			}

			if string('{') == token.Value {
				//log.Println(token.Value)
				conf.SetPaserState(PARSE_STATE_START_DECLARE_BLOCK)

			}

			if string('}') == token.Value {
				//log.Println(token.Value)
				conf.SetPaserState(PARSE_STATE_END_DECLARE_BLOCK)

				rule := stylesheet.NewCSSStyleRule(stylesheet.DOMString(conf.NowSelectorText),
					stylesheet.DOMString(conf.NowProperty),
					stylesheet.DOMString(conf.NowValue))
				log.Println(*rule)
				conf.stylesheet.AddRuleList(rule)
				log.Println(conf.stylesheet)

			}
		case scanner.TokenPercentage:
			if conf.GetPaserState() == PARSE_STATE_START_DECLARE_BLOCK {
				conf.NowValue = token.Value
			}
		}
		rulelist := conf.stylesheet.GetRuleList()

		log.Println(rulelist[0].CssType)

	}

}
