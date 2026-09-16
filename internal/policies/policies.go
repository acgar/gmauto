package policies

import (
	"bufio"
	"fmt"
	"gmauto/internal/config"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const readPoliciesDirName = "policies/read"
const trashPoliciesDirName = "policies/trash"

type Policy struct {
	TagName          string
	ExpirationInDays int
}

func LoadReadPolicies() []*Policy {
	return loadPoliciesDir(path.Join(config.GetSettingsPath(), readPoliciesDirName))
}

func LoadTrashPolicies() []*Policy {
	return loadPoliciesDir(path.Join(config.GetSettingsPath(), trashPoliciesDirName))
}

func loadPoliciesDir(dir string) []*Policy {
	paths, err := filepath.Glob(dir + "*.tag")
	if err != nil {
		panic(err)
	}

	var tagPolicies []*Policy
	for _, path := range paths {
		tagPolicy, err := loadTagPolicyFile(path)
		if err != nil {
			panic(err)
		}
		tagPolicies = append(tagPolicies, tagPolicy)
	}
	return tagPolicies
}

func loadTagPolicyFile(path string) (*Policy, error) {
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

	tagPolicy := &Policy{TagName: tagName}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Discard comments and empty lines
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		re := regexp.MustCompile(`^([0-9]+)d$`)
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			tagPolicy.ExpirationInDays, err = strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return tagPolicy, nil
}
