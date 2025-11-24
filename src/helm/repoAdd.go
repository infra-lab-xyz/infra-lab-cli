package helm

import (
	"fmt"
)

func RepoAdd(repoName, repoUrl string, forceUpdate bool) (err error) {
	schema, err := parseSchema(repoUrl)
	if err != nil {
		return err
	}

	switch schema {
	case "http", "https":
		err = httpRepoAdd(repoName, repoUrl, forceUpdate)
	case "oci":
		err = ociRepoAdd(repoName, repoUrl, forceUpdate)
	default:
		fmt.Printf("Schema '%s' is unknown: ", schema)
	}

	if err != nil {
		return err
	}

	return nil
}
