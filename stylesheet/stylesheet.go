package stylesheet

import (
	"log"
)

type DOMString string

const (
	STYLE_RULE   = 1
	CHARSET_RULE = 2
)

type StyleSheet struct {
	Type DOMString
}

type CSSStyleSheet struct {
	Type        DOMString
	CssRuleList CSSRuleList
}

type StyleSheetList []StyleSheet

type CSSRuleList struct {
	RuleList []*CSSRule
}

type CSSRule struct {
	CssType int
	Style   CSSStyleRule
}

type CSSStyleRule struct {
	SelectorText DOMString
	Styles       []*CSSStyleDeclaration
}

type CSSStyleDeclaration struct {
	Property DOMString
	Value    DOMString
}

func (self *CSSStyleDeclaration) Construct() {
	// Some initialize
}

func NewCSSStyleDeclaration(property DOMString, value DOMString) *CSSStyleDeclaration {
	style := new(CSSStyleDeclaration)
	style.Property = property
	style.Value = value

	style.Construct()
	return style
}

func NewCSSRule(Type int) *CSSRule {
	rule := &CSSRule{}
	rule.SetType(Type)
	return rule
}

func NewCSSStyleRule(selectorText DOMString, Property DOMString, Value DOMString) *CSSRule {
	rule := new(CSSRule)
	rule.SetType(STYLE_RULE)

	rule.Style.SelectorText = selectorText
	rule.Style.AddStyleDeclaration(Property, Value)
	return rule
}

func (self *CSSRule) GetType() int {
	return self.CssType
}

func (self *CSSRule) SetType(Type int) {
	self.CssType = Type
}

func (self *CSSStyleSheet) AddRuleList(rule *CSSRule) {
	self.CssRuleList.RuleList = append(self.CssRuleList.RuleList, rule)
}

func (self *CSSStyleSheet) GetRuleList() *CSSRuleList {
	return &self.CssRuleList
}

func (self *CSSStyleRule) AddStyleDeclaration(property DOMString, value DOMString) {
	style := NewCSSStyleDeclaration(property, value)
	self.Styles = append(self.Styles, style)
}
