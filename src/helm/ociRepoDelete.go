package helm

func ociRepoDelete(repoName string) (err error) {
	ociRepos := GetOCIRepos()

	for idx := range ociRepos.Repos {
		if ociRepos.Repos[idx].Name == repoName {
			ociRepos.Repos = append(ociRepos.Repos[:idx], ociRepos.Repos[idx+1:]...)
			break
		}
	}

	err = SaveOCIRepos(ociRepos)

	return err
}
