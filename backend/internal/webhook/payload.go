package webhook

type PullRequestEvent struct {
	Action      string      `json:"action"`
	Number      int         `json:"number"`
	PullRequest PullRequest `json:"pull_request"`
	Repository  Repository  `json:"repository"`
}

type PullRequest struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Head   Commit `json:"head"`
	User   User   `json:"user"`
}

type Commit struct {
	SHA string `json:"sha"`
	Ref string `json:"ref"`
}

type Repository struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
}

type User struct {
	Login string `json:"login"`
}
