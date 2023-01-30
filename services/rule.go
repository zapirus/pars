package services

import (
	"log"
	"strings"

	"golang.org/x/net/html"
)

const rTagLongest = "longest"

type Rule struct {
	actionForText   string // Правило заполнения текстовых полей
	actionForRadio  string // Правило выбора радиобаттона
	actionForSelect string // Правило выбора дропдауна
}

func (r Rule) Validate() bool {
	if r.actionForText == "" {
		return false
	}

	if r.actionForRadio != rTagLongest {
		return false
	}

	if r.actionForSelect != rTagLongest {
		return false
	}

	return true
}

func ruleParser(page string) Rule {
	var rule Rule
	htmlTokens := html.NewTokenizer(strings.NewReader(page))

loop:
	for {
		tt := htmlTokens.Next()
		switch tt {
		case html.ErrorToken:
			log.Println("parsing rule ended")
			break loop
		case html.TextToken:
			t := htmlTokens.Token()
			switch t.Data {
			case "INPUT[@type=text]":
				htmlTokens.Next()
				inputTt := htmlTokens.Next()
				switch inputTt {
				case html.TextToken:
					inputT := htmlTokens.Token()
					rule.actionForText = strings.TrimSuffix(strings.TrimPrefix(inputT.Data, " field must be filled in with \""), "\" value")
				}
			case "INPUT[@type=radio]":
				htmlTokens.Next()
				inputTt := htmlTokens.Next()
				switch inputTt {
				case html.TextToken:
					inputT := htmlTokens.Token()
					rule.actionForRadio = strings.TrimSuffix(strings.TrimPrefix(inputT.Data, " field must be selected with the "), " value")
				}
			case "SELECT":
				htmlTokens.Next()
				inputTt := htmlTokens.Next()
				switch inputTt {
				case html.TextToken:
					inputT := htmlTokens.Token()
					rule.actionForSelect = strings.TrimSuffix(strings.TrimPrefix(inputT.Data, " field must be selected with the "), " value")
				}
			}
		case html.StartTagToken:
			t := htmlTokens.Token()
			isAnchor := t.Data == "a"
			if isAnchor {
				log.Println("We found an anchor!")
			}
		}
	}

	return rule
}
