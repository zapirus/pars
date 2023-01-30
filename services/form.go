package services

import (
	"log"
	"strings"

	"golang.org/x/net/html"
)

type Form struct {
	text    []string
	radios  []map[string][]string
	selects []map[string][]string
}

func (f Form) Validate() bool {
	if len(f.text) == 0 &&
		len(f.radios) == 0 &&
		len(f.selects) == 0 {
		return false
	}
	return true
}

func formParser(page string) Form {
	var form Form
	htmlTokens := html.NewTokenizer(strings.NewReader(page))

loop:
	for {
		tt := htmlTokens.Next()
		switch tt {
		case html.ErrorToken:
			log.Println("parsing form ended")
			break loop
		case html.TextToken:
			t1 := htmlTokens.Token()
			log.Println(t1.Data)
		case html.StartTagToken:
			t1 := htmlTokens.Token()
			log.Println(t1.Data)

			if t1.Data == "form" {
			loopForm:
				for {
					htmlTokens.Next()
					inner := htmlTokens.Token()
					switch inner.Data {
					case "form":
						break loopForm
					case "input":
						var ftype string
						var name string
						var value string
						for _, a := range inner.Attr {
							switch a.Key {
							case "type":
								ftype = a.Val
							case "name":
								name = a.Val
							case "value":
								value = a.Val
							}
						}

						switch ftype {
						case "text":
							form.text = append(form.text, name)
						case "radio":
							var isFound bool
							for i := range form.radios {
								if _, ok := form.radios[i][name]; ok {
									form.radios[i][name] = append(form.radios[i][name], value)
									isFound = true
									break
								}
							}

							if !isFound {
								form.radios = append(form.radios, map[string][]string{
									name: {value},
								})
							}
						default:
							log.Fatalln("unexpected field type", ftype)
						}
					case "select":
						var selName string
						var values []string
						for _, a := range inner.Attr {
							switch a.Key {
							case "name":
								selName = a.Val
							}
						}

					loopSelect:
						for {
							htmlTokens.Next()
							inner := htmlTokens.Token()
							switch inner.Data {
							case "select":
								break loopSelect
							case "option":
								for _, a := range inner.Attr {
									switch a.Key {
									case "value":
										values = append(values, a.Val)
									}
								}
							}
						}

						form.selects = append(form.selects, map[string][]string{
							selName: values,
						})
					}

					log.Println(inner.Data)
				}
			}
		}
	}

	return form
}
