package internal

import (
	"fmt"

	"github.com/unix755/xtools/xRelease/codeberg"
)

type CodebergAPI struct {
	r *codeberg.Release
}

func GetCodebergApiLatest(repo string) (api *CodebergAPI, err error) {
	r, err := codeberg.GetReleaseLatest(repo)
	return &CodebergAPI{r}, err
}

func GetCodebergApiByTagName(repo string, tagName string) (api *CodebergAPI, err error) {
	r, err := codeberg.GetReleaseByTagName(repo, tagName)
	return &CodebergAPI{r}, err
}

func (a *CodebergAPI) GetDownloadLink(includes []string, excludes []string) (downloadLink string, err error) {
	release := a.r.GetAssets(includes, excludes)
	if len(release) <= 0 {
		return "", fmt.Errorf("can not find the release")
	}
	return release[0].BrowserDownloadURL, nil
}
