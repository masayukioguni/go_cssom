package stylesheet

import (
	//"log"
	"testing"
)

func TestCSSRuleSetGetType(t *testing.T) {
	rule := new(CSSRule)
	rule.SetType(100)

	if 100 != rule.GetType() {
		t.Errorf("GetType() = %d , want 100.", rule.GetType())
	}

}

func TestAddRuleList(t *testing.T) {
	styleSheet := new(CSSStyleSheet)
	styleSheet.ruleList = make(CSSRuleList, 0)

	rule1 := NewCSSRule(STYLE_RULE)
	rule2 := NewCSSRule(STYLE_RULE)

	styleSheet.AddRuleList(rule1)
	styleSheet.AddRuleList(rule2)

	actual := len(styleSheet.ruleList)

	if 2 != actual {
		t.Errorf("len(styleSheet.ruleList) = %d , want 2.", actual)
	}

}

func TestAddStyleDeclaration(t *testing.T) {
	rule := NewCSSRule(STYLE_RULE)
	styleRule := rule.style

	styleRule.styles = make([]*CSSStyleDeclaration, 0)

	styleRule.AddStyleDeclaration("test", "test")
	styleRule.AddStyleDeclaration("test2", "test2")

	actual := len(styleRule.styles)

	if 2 != actual {
		t.Errorf("len(rule.styles) = %d , want 2.", actual)
	}

	style := styleRule.styles[0]
	if "test" != style.Property {
		t.Errorf("style.Property = %s , want test.", style.Property)
	}
	if "test" != style.Value {
		t.Errorf("style.Value = %s , want test.", style.Value)
	}

}
