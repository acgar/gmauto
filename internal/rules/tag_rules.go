package rules

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const tagRulesDir = "./config/tag-rules/"

type TagRules struct {
	TagName string
	rules   []*Rule
}

type Rule struct {
	fromConditions    []string
	subjectConditions []string
}

func LoadRules() []*TagRules {
	return loadRulesDir(tagRulesDir)
}

func loadRulesDir(dir string) []*TagRules {
	paths, err := filepath.Glob(dir + "*.tag")
	if err != nil {
		panic(err)
	}

	var tagRulesArray []*TagRules
	for _, path := range paths {
		tagRule, err := loadTagRulesFile(path)
		if err != nil {
			panic(err)
		}
		tagRulesArray = append(tagRulesArray, tagRule)
	}
	return tagRulesArray
}

func loadTagRulesFile(path string) (*TagRules, error) {
	if filepath.Ext(path) != ".tag" {
		return nil, fmt.Errorf("wrong file extension. expected .tag")
	}

	tagName := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	tagName = strings.ReplaceAll(tagName, "-", "/")

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	tagRules := &TagRules{TagName: tagName}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Discard comments and empty lines
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		rule := Rule{}
		for _, word := range strings.Fields(line) {
			if strings.HasPrefix(word, "from:") {
				rule.fromConditions = append(rule.fromConditions, strings.TrimPrefix(word, "from:"))
			} else if strings.HasPrefix(word, "subject:") {
				rule.subjectConditions = append(rule.subjectConditions, strings.TrimPrefix(word, "subject:"))
			}
		}
		if len(rule.fromConditions) > 0 || len(rule.subjectConditions) > 0 {
			tagRules.rules = append(tagRules.rules, &rule)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return tagRules, nil
}

// CheckTagRules /*
/*
Checks if any tag rule match (OR)
*/
func CheckTagRules(fromText string, subjectText string, tagRules *TagRules) bool {
	for _, rule := range tagRules.rules {
		if checkRule(fromText, subjectText, rule) {
			return true
		}
	}
	return false
}

/*
Checks that all rule conditions are met (AND)
*/
func checkRule(fromText string, subjectText string, rule *Rule) bool {
	for _, fromCondition := range rule.fromConditions {
		if !checkTextPattern(fromText, fromCondition) {
			return false
		}
	}
	for _, subjectCondition := range rule.subjectConditions {
		if !checkTextPattern(subjectText, subjectCondition) {
			return false
		}
	}
	return true
}

func checkTextPattern(input string, pattern string) bool {
	if !strings.Contains(pattern, "*") {
		return input == pattern
	}

	// Process * format
	literalChunks := strings.Split(pattern, "*")
	pos := 0
	for i, chunk := range literalChunks {
		if chunk == "" {
			continue
		}

		indice := strings.Index(input[pos:], chunk)
		if indice == -1 {
			return false
		}

		indice += pos

		// If pattern does not start with *, input text must start identically to pattern
		if i == 0 && indice != 0 {
			return false
		}

		pos = indice + len(chunk)
	}

	// If pattern does not end with *, input text must end identically to pattern
	if !strings.HasSuffix(pattern, "*") {
		return pos == len(input)
	}

	return true
}

func Display(tagRulesArray []*TagRules) {
	fmt.Println("📐 Tag rules:")
	for _, tagRules := range tagRulesArray {
		fmt.Println("  - 🏷️ " + tagRules.TagName)
		for _, rule := range tagRules.rules {
			if len(rule.subjectConditions) > 0 || len(rule.fromConditions) > 0 {
				fmt.Print("        | ")
			}
			if len(rule.subjectConditions) > 0 {
				fmt.Print(" SUBJECT: (" + strings.Join(rule.subjectConditions, " AND ") + ")")
			}
			if len(rule.fromConditions) > 0 {
				fmt.Print(" FROM: (" + strings.Join(rule.fromConditions, " AND ") + ")")
			}
			if len(rule.subjectConditions) > 0 || len(rule.fromConditions) > 0 {
				fmt.Print("\n")
			}
		}
	}

}
