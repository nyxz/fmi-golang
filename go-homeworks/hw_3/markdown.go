package main

import (
	"fmt"
	"regexp"
)

type MarkdownParser struct {
	text string
}

func NewMarkdownParser(text string) *MarkdownParser {
	mp := new(MarkdownParser)
	mp.text = text
	return mp
}

func (mp *MarkdownParser) Headers() []string {
	result := make([]string, 0)

	fillWithAllStringSubmatches(`(.*)\n(=)+\s`, mp.text, &result, 1)	
	fillWithAllStringSubmatches(`#\s.*`, mp.text, &result, 0)	

	return result
}

func (mp *MarkdownParser) SubHeadersOf(header string) []string {
	result := make([]string, 0)

	// fillWithAllStringSubmatches(`(?m)(?P<g1>(.*)\n(=)+\s*)`, mp.text, &result, 2)	

	exp := regexp.MustCompile(`(?m)(?P<g1>(.*)\n(?P<g2>=)+\s*)`)
	names := exp.SubexpNames()
	for _, v := range names {
		fmt.Println("---", v, "---")
	}
	// fmt.Println("group 1: ", exp.SubexpNames())
	substrings := exp.FindAllStringSubmatch(mp.text, -1)
	for _, v := range substrings {
		result = append(result, v[1])
	}

	return result
}

func (mp *MarkdownParser) Names() []string {
	// TODO impl
	// re := regexp.MustCompile("(?P<first>[a-zA-Z]+) (?P<last>[a-zA-Z]+)")
	// fmt.Println(re.MatchString("Alan Turing"))
	// fmt.Printf("%q\n", re.SubexpNames())
	// reversed := fmt.Sprintf("${%s} ${%s}", re.SubexpNames()[2], re.SubexpNames()[1])
	// fmt.Println(reversed)
	// fmt.Println(re.ReplaceAllString("Alan Turing", reversed))
	return nil
}

func (mp *MarkdownParser) PhoneNumbers() []string {
	// TODO impl
	return nil
}

func (mp *MarkdownParser) Links() []string {
	// TODO impl
	return nil
}

func (mp *MarkdownParser) Emails() []string {
	// TODO impl
	return nil
}

func (mp *MarkdownParser) GenerateTableOfContents() string {
	// TODO impl
	return ""
}

func fillWithAllStringSubmatches(regex string, text string, result *[]string, subIdx int) {
	exp := regexp.MustCompile(regex)
	substrings := exp.FindAllStringSubmatch(text, -1)
	for _, v := range substrings {
		*result 	= append(*result, v[subIdx])
	}
}

func main() {
	text := `This is a header
=
please handle it.
here is some other text agains
here is some other text agains
--
here is some other text agains
here is some other text agains
Just another row.
=
# Another header here`
	mp := NewMarkdownParser(text)
	headers := mp.SubHeadersOf(text)
	fmt.Println("=========================================")
	for i, v := range headers {
		fmt.Println(i, ">", v, "<")
	}
}