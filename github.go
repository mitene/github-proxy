package proxy

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/go-github/v29/github"
	"github.com/mholt/archiver"
	"golang.org/x/oauth2"
)

type GithubClient interface {
	MakeArchive(owner string, repo string, ref string, repoPath string, dest string, ext string) (string, error)
}

type GithubClientImpl struct {
	client *github.Client
}

func NewGithubClient() GithubClient {
	var tc *http.Client
	token, ok := os.LookupEnv("GITHUB_ACCESS_TOKEN")
	if ok {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc = oauth2.NewClient(context.Background(), ts)
	}
	return &GithubClientImpl{
		client: github.NewClient(tc),
	}
}

func (c *GithubClientImpl) MakeArchive(owner string, repo string, ref string, repoPath string, dest string, ext string) (string, error) {
	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		return "", err
	}
	downloadDir := filepath.Join(tempDir, repo)
	err = os.Mkdir(downloadDir, os.ModePerm)
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(tempDir)

	file, err := c.downloadRepoZip(owner, repo, ref, downloadDir)
	if err != nil {
		return "", err
	}

	err = c.extract(file, downloadDir)
	if err != nil {
		return "", err
	}

	targetPath := filepath.Join(downloadDir, repoPath)
	_, err = os.Stat(targetPath)
	if err != nil {
		return "", err
	}

	archiveFile := filepath.Join(dest, fmt.Sprintf("%s-%s-%s.%s", owner, repo, ref, ext))

	err = archiver.Archive(
		[]string{targetPath},
		archiveFile,
	)
	if err != nil {
		return "", err
	}

	return archiveFile, err
}

func (c *GithubClientImpl) downloadRepoZip(owner string, repo string, ref string, dest string) (string, error) {

	opt := github.RepositoryContentGetOptions{
		Ref: ref,
	}
	url, _, err := c.client.Repositories.GetArchiveLink(
		context.Background(),
		owner,
		repo,
		github.Zipball,
		&opt,
		true,
	)
	if err != nil {
		return "", err
	}

	resp, err := http.Get(url.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	tempFile, err := ioutil.TempFile(dest, "")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, resp.Body)
	if err != nil {
		return "", err
	}

	return tempFile.Name(), nil
}

func (c *GithubClientImpl) extract(file string, dest string) error {
	r, err := zip.OpenReader(file)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		filename := filepath.Join(strings.Split(f.Name, string(os.PathSeparator))[1:]...)
		path := filepath.Join(dest, filename)

		if f.FileInfo().IsDir() {
			err = os.MkdirAll(path, f.Mode())
			if err != nil {
				return err
			}
		} else {
			err = func() error {
				rc, err := f.Open()
				if err != nil {
					return err
				}
				defer rc.Close()

				d, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, f.Mode())
				if err != nil {
					return err
				}
				defer d.Close()

				if _, err = io.Copy(d, rc); err != nil {
					return err
				}
				return nil
			}()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
