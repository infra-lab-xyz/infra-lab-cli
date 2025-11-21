package helm

import (
	"net/url"
)

func RepoAdd(repoName, repoUrl string, forceUpdate bool) {
	// TODO:
	//      1. Check by name if repo already added to any of http or OCI repos list
	//      2. Determine repository type (necessary on add, delete, update, list chart versions)
	var schema string

	parsedUrl, err := url.Parse(repoUrl)
	if err == nil && parsedUrl.Scheme != "" {
		schema = parsedUrl.Scheme
	} else {
		// TODO: maybe show an error and stop?
		schema = ""
	}

	switch schema {
	case "http", "https":
		_ = httpRepoAdd(repoName, repoUrl, forceUpdate)
	case "oci":
		_ = ociRepoAdd(repoName, repoUrl, forceUpdate)
	}
}
