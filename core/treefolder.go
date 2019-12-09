package core

import (
	"encoding/json"
	"fmt"
	"github.com/phpxin/mdblog/conf"
	"io/ioutil"
	"os"
	"path/filepath"
)

type TreeFolderIntro struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
	Intro string `json:"intro"`
	Children map[string]*TreeFolderIntro `json:"children"`
}

type TreeFolder struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Title string `json:"title"` // the title
	Desc string `json:"desc"` // for SEO
	Intro string `json:"intro"` // show on index path or article list
	Children []*TreeFolder `json:"children"`
}

var (
	treeFolder *TreeFolder
	DocsIndexer = make(map[string]*TreeFolder)
)

func GetTreeFolder() *TreeFolder {
	return treeFolder
}

func GenerateTreeFolder() error {
	dirPath := conf.ConfigInst.Docroot
	if dirPath=="" {
		return fmt.Errorf("document root dir didn't set")
	}

	finfo,err := os.Lstat(dirPath)
	if err!=nil {
		return err
	}

	if !finfo.IsDir() {
		return fmt.Errorf("%s is not a dir", dirPath)
	}

	treeFolder,err = generateTreeFolder(dirPath)

	if err!=nil {
		return err
	}

	jsonfmt,_ := json.Marshal(treeFolder)
	err = ioutil.WriteFile(conf.ConfigInst.Storagepath+"/database/tree_folder.json", jsonfmt, 0644)
	if err!=nil {
		return err
	}

	return nil
}

func generateTreeFolder(dirPath string) (*TreeFolder, error) {
	finfo,err := os.Lstat(dirPath)
	if err!=nil {
		return nil,err
	}

	treeFolderIntro := new(TreeFolderIntro)

	dirIntroContents,err := ioutil.ReadFile(dirPath+"/introduce.json")
	if err!=nil {
		if !os.IsNotExist(err) {
			return nil,err
		}
	}else{
		err = json.Unmarshal(dirIntroContents, treeFolderIntro)
		if err!=nil {
			return nil,err
		}
	}

	finfos,err := ioutil.ReadDir(dirPath)
	if err!=nil {
		return nil,err
	}

	treeFolder := new(TreeFolder)
	treeFolder.Name = finfo.Name()
	treeFolder.Title = finfo.Name()
	if treeFolderIntro.Title!="" {
		treeFolder.Title = treeFolderIntro.Title
	}
	treeFolder.Desc = treeFolderIntro.Desc
	treeFolder.Intro = treeFolderIntro.Intro
	treeFolder.Path = dirPath

	for _,fitem:=range finfos {
		if fitem.IsDir() {
			cTreeFolder,err := generateTreeFolder(dirPath+"/"+fitem.Name())
			if err!=nil {
				return nil,err
			}

			treeFolder.Children = append(treeFolder.Children, cTreeFolder)

		}else{
			if filepath.Ext(fitem.Name())==".md" {
				theFile := new(TreeFolder)
				theFile.Name = fitem.Name()
				theFile.Title = fitem.Name()
				fileIntro,ok := treeFolderIntro.Children[theFile.Name]
				if ok {
					if fileIntro.Title!="" {
						theFile.Title = fileIntro.Title
					}

					theFile.Desc = fileIntro.Desc
					theFile.Intro = fileIntro.Intro
				}
				theFile.Path = dirPath+"/"+fitem.Name()
				theFile.Children = nil

				treeFolder.Children = append(treeFolder.Children, theFile)
				DocsIndexer[theFile.Name] = theFile
			}

		}
	}

	return treeFolder,nil
}