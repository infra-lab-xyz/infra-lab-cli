package helm

func RepoDelete(repoName string) (err error) {
	repo := getRepoByName(repoName)
	if repoName != "" {
		schema, err := parseSchema(repo.Url)
		if err != nil {
			return err
		}

		switch schema {
		case "http", "https":
			err = httpRepoDelete(repo)
		case "oci":
			err = ociRepoDelete(repoName)
		}
		if err != nil {
			return err
		}
	}

	return nil
}
