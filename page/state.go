package page

import "github.com/gopherjs/gopherjs/js"

// State that is passed to the frontend script from the backend handler.
type State struct {
	// TODO: Consider *js.Object.
	Production   bool
	ImportPath   string
	RepoSpec     repoSpec
	ProcessedRev string // ProcessedRev is processed rev; its value is replaced by default branch if empty.
	CommitID     string // TODO: Either get rid of godep or make gopherjs_http.NewFS use Godeps.json versions, then can start using vcs.CommitID directly.
}

// TODO: Dedup.
// repoSpec identifies a repository.
type repoSpec struct {
	VCSType  string
	CloneURL string
}

// TODO: Dedup. It's duplicated because including *js.Object in backend makes it panic because of:
//
//           runtime error: invalid memory address or nil pointer dereference
//
//       It happens in call to html/template.(*Template).ExecuteTemplate().
type StateObject struct {
	*js.Object
	Production   bool           `js:"Production"`
	ImportPath   string         `js:"ImportPath"` // TODO: Consider changing GopherJS so this explicit js tag isn't needed?
	RepoSpec     repoSpecObject `js:"RepoSpec"`
	ProcessedRev string         `js:"ProcessedRev"` // ProcessedRev is processed rev; its value is replaced by default branch if empty.
	CommitID     string         `js:"CommitID"`     // TODO: Either get rid of godep or make gopherjs_http.NewFS use Godeps.json versions, then can start using vcs.CommitID directly.
}

// TODO: Dedup. It's duplicated because including *js.Object in backend makes it panic.
type repoSpecObject struct {
	*js.Object
	VCSType  string `js:"VCSType"`
	CloneURL string `js:"CloneURL"`
}