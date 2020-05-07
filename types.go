package main


type Response struct {
	Tree []GitFile `json:"tree"`
}

type GitFile struct {
	Name string `json:"path"`
	Type string `json:"type"`
}

type Repository struct {
	User string
	RepoName string
	Branch string
}

func newRepo(user, repoName, branch string) Repository {
	return Repository {
		User: user,
		RepoName: repoName,
		Branch: branch,
	}
}
