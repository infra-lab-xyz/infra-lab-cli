package helm

import (
	"fmt"
)

func ListRepos() (err error) {
	repos, _ := GetRepos()
	fmt.Printf("%-40s %-40s\n", "Name", "URL")
	for _, repo := range repos {
		fmt.Printf("%-40s %-40s\n", repo.Name, repo.Url)
	}

	return nil
}
