package servers

import (
	"fmt"
	"os"
	"strings"

	git "github.com/go-git/go-git/v5"
)

func clone(roleFolder string, gitRemoteURL string) error {
	_ = os.MkdirAll(roleFolder, os.ModePerm)
	cloneOpts := &git.CloneOptions{
		URL:      gitRemoteURL,
		Progress: os.Stdout,
	}
	r, err := git.PlainClone(roleFolder, false, cloneOpts)
	if err != nil {
		fmt.Println("Failed to clone repo\n---\n", err.Error())
		return err
	}
	fmt.Println(r)

	return nil
}

func pull(roleFolder string) error {
	r, err := git.PlainOpen(roleFolder)
	if err != nil {
		fmt.Printf("Failed to open local repo: %s\n---\n%s\n", roleFolder, err.Error())
		return err
	}
	w, err := r.Worktree()
	if err != nil {
		fmt.Println("Failed to open parse local git tree\n---\n", err.Error())
		return err
	}
	err = w.Pull(&git.PullOptions{RemoteName: "origin", Progress: os.Stdout})
	if err != nil && !strings.Contains(err.Error(), "already up-to-date") {
		fmt.Println("Failed to update local repo\n---\n", err.Error())
		return err
	}
	return nil
}
