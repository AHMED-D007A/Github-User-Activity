package main

type Event struct {
	ID        string  `json:"id"`
	Type      string  `json:"type"`
	Actor     Actor   `json:"actor"`
	Repo      Repo    `json:"repo"`
	Payload   Payload `json:"payload"`
	Public    bool    `json:"public"`
	CreatedAt string  `json:"created_at"`
}

// Serve info about the users.
type Actor struct {
	ID   int64  `json:"id"`
	Name string `json:"login"`
}

// Serve info about the repos.
type Repo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Payload struct {
	Ref string `json:"ref"`
	// These are for CreateEvent Type (Repositories, Tags, Branches).
	Ref_Type      string `json:"ref_type"`
	Master_Branch string `json:"master_branch"`
	Description   string `json:"description"`
	Pusher_Type   string `json:"pusher_type"`
	// These are for PushEvent Type.
	Repository_Id int64    `json:"repository_id"`
	Push_Id       int64    `json:"push_id"`
	Head          string   `json:"head"`
	Before        string   `json:"before"`
	Commit        []Commit `json:"commits"`
}

// Serve info about the pushed commits.
type Commit struct {
	Sha     string `json:"sha"`
	Message string `json:"message"`
}
