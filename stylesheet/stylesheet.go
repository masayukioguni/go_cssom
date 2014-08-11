package stylesheet

type DOMString string

const (
	STYLE_RULE   = 1
	CHARSET_RULE = 2
)

type StyleSheet struct {
	Type DOMString
}

type CSSStyleSheet struct {
	Type     DOMString
	ruleList CSSRuleList
}

type StyleSheetList []StyleSheet

type CSSRuleList []*CSSRule

type CSSRule struct {
	CssType int

	style CSSStyleRule
}

type CSSStyleRule struct {
	CssText DOMString
	styles  []*CSSStyleDeclaration
}

type CSSStyleDeclaration struct {
	Property DOMString
	Value    DOMString
}

func (self *CSSStyleDeclaration) Construct() {
	// Some initialize
}

func NewCSSStyleDeclaration(property DOMString, value DOMString) *CSSStyleDeclaration {
	style := &CSSStyleDeclaration{Property: property, Value: value}
	style.Construct()
	return style
}

func NewCSSRule(Type int) *CSSRule {
	rule := &CSSRule{}
	rule.SetType(Type)
	return rule
}

func (self *CSSRule) GetType() int {
	return self.CssType
}

func (self *CSSRule) SetType(Type int) {
	self.CssType = Type
}

func (self *CSSStyleSheet) AddRuleList(rule *CSSRule) {
	self.ruleList = append(self.ruleList, rule)
}

func (self *CSSStyleRule) AddStyleDeclaration(property DOMString, value DOMString) {
	style := NewCSSStyleDeclaration(property, value)
	self.styles = append(self.styles, style)
}
