package rfind

import (
	"errors"
	"math"
	"os"
	"path/filepath"
)

type Rfind struct {
	MaxUpwardDepth   uint
	MaxDepthFromRoot uint
	Limit            uint
	IsFile           bool
	IsDir            bool
}

func (r *Rfind) Find(originPath string, targets []string) ([]string, error) {
	var searchPaths []string
	for {
		searchPaths = append(searchPaths, originPath)
		parentDir := filepath.Dir(originPath)
		if parentDir == originPath {
			// reaches to the root dir
			break
		}
		originPath = parentDir
	}

	maxUpwardDepth := math.MaxInt
	if r.MaxUpwardDepth > 0 {
		maxUpwardDepth = int(r.MaxUpwardDepth)
	}

	maxDepthFromRoot := math.MaxInt
	if r.MaxDepthFromRoot > 0 {
		maxDepthFromRoot = len(searchPaths) - int(r.MaxDepthFromRoot)
	}

	limit := math.MaxInt
	if r.Limit > 0 {
		limit = int(r.Limit)
	}
	numOfFound := 0

	var foundItems []string
searchLoop:
	for _, searchPath := range searchPaths {
		maxUpwardDepth--
		if maxUpwardDepth < 0 {
			break
		}

		maxDepthFromRoot--
		if maxDepthFromRoot < 0 {
			break
		}

		for _, target := range targets {
			targetFullPath := filepath.Join(searchPath, target)
			stat, err := os.Stat(targetFullPath)
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					continue
				}
				return nil, err
			}

			if (r.IsFile && stat.Mode().IsRegular()) || (r.IsDir && stat.Mode().IsDir()) {
				foundItems = append(foundItems, targetFullPath)
				numOfFound++
				if numOfFound >= limit {
					break searchLoop
				}
			}
		}
	}

	return foundItems, nil
}
