package api

import (
	"github.com/NextCodeX/desktop-application-pedestal/util"
	. "github.com/changlie/go-common/a"
	"github.com/changlie/go-common/files"
	"path/filepath"
	"sort"
	"strings"
)

func init() {
	obj := &FileApi{}
	infoMap := CollectFunctionInfo(&obj)
	allModInfo["file"] = infoMap
}

type FileApi Void

func (f *FileApi) DiskPartitions(c Ctx) *Resp {
	return success1(util.DiskPartitions())
}

func (f *FileApi) AppList(c Ctx) *Resp {
	list := files.FileList2(ProgramDir())
	var res []string
	for _, info := range list {
		if info.Dir {
			res = append(res, info.Name)
		}
	}
	return success1(res)
}

func (f *FileApi) FileList(c Ctx) *Resp {
	var arg Req
	c.BindArgs(&arg)
	//fmt.Println(arg.Val)
	targetPath := arg.Val
	if !files.Exists(targetPath) {
		return fail1("invalid path")
	}
	if !files.IsDir(targetPath) {
		targetPath = filepath.Dir(targetPath)
	}

	list2 := files.FileList2(targetPath)
	if len(list2) < 1 {
		return success1(list2)
	}

	return success1(fileListSort(list2))
}

func fileListSort(list2 []*files.FileInfo2) []*files.FileInfo2 {
	var dirs []*files.FileInfo2
	var dirKeys []string
	var fs []*files.FileInfo2
	var fKeys []string
	for _, info2 := range list2 {
		if strings.ToUpper(info2.Name) == strings.ToUpper("$RECYCLE.BIN") ||
			strings.ToUpper(info2.Name) == strings.ToUpper("System Volume Information") ||
			strings.ToUpper(info2.Name) == strings.ToUpper("$WinREAgent") {
			continue
		}
		if info2.Dir {
			dirs = append(dirs, info2)
			dirKeys = append(dirKeys, info2.Name)
		} else {
			fs = append(fs, info2)
			fKeys = append(fKeys, info2.Name)
		}
	}

	sort.Strings(dirKeys)
	sort.Strings(fKeys)

	itemGet := func(kw string, fList []*files.FileInfo2) *files.FileInfo2 {
		for _, finfo := range fList {
			if finfo.Name == kw {
				return finfo
			}
		}
		return nil
	}
	dirsSize := len(dirKeys)
	size := dirsSize + len(fKeys)

	res := make([]*files.FileInfo2, size)
	for i, key := range dirKeys {
		res[i] = itemGet(key, dirs)
	}
	for i, key := range fKeys {
		res[i+dirsSize] = itemGet(key, fs)
	}
	return res
}
