package api

import (
	"fmt"
	"github.com/NextCodeX/desktop-application-pedestal/util"
	. "github.com/changlie/go-common/a"
	"github.com/changlie/go-common/files"
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
	fmt.Println(arg.Val)

	list2 := files.FileList2(arg.Val)

	return success1(list2)
}
