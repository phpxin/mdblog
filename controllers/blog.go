package controllers

import (
	"github.com/phpxin/mdblog/conf"
	"github.com/phpxin/mdblog/core"
	model "github.com/phpxin/mdblog/models"
	"github.com/phpxin/mdblog/tools/log"
	"gopkg.in/russross/blackfriday.v2"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type BlogController struct {

}

func (ctrl *BlogController) Index(r *http.Request) (resp *core.HttpResponse) {
	qStr := r.URL.Query()
	subject := qStr.Get("subject")

	obj,ok := core.SubjectIndexer[subject]
	if !ok {
		return core.HtmlResponse("errors/404", nil)
	}

	page := qStr.Get("p")
	pagen,err := strconv.Atoi(page)
	if err!=nil {
		log.Error("get page failed %s", err.Error())
		pagen = 1
	}

	var limit int32 = 5

	subjects := make([]*core.TreeFolder, 0)
	//articles := make([]*core.TreeFolder, 0)

	for _,item := range obj.Children {
		if len(item.Children)>0 {
			subjects = append(subjects, item)
		}else{
			//articles = append(articles, item)
		}
	}

	articles,amount := model.GetDocsBySubject(subject,int32(pagen), limit)
	prevPage := -1
	nextPage := -1

	if pagen>1 {
		prevPage = pagen-1
	}
	if int32(pagen)*limit < amount {
		nextPage = pagen+1
	}

	hot := model.GetHotRanging()

	sidebar := sidebar(subjects, hot)
	nav := nav()
	footer := footer()

	return core.HtmlResponse("subject3", struct{
		Sidebar template.HTML
		Nav template.HTML
		Footer template.HTML
		Articles []*model.Doc
		SubjectHash string
		Title string
		PrevPage int
		NextPage int
	}{
		template.HTML(sidebar) ,
		template.HTML(nav) ,
		template.HTML(footer) ,
		articles,
		subject,
		obj.Title,
		prevPage,
		nextPage,
	})
}

func (ctrl *BlogController) Detail(r *http.Request) (resp *core.HttpResponse) {
	// @todo 全局参数获取、过滤、格式化、校验插件
	qStr := r.URL.Query()
	mdname := qStr.Get("md")

	obj,ok := model.GetDoc(mdname)
	if !ok {
		return core.HtmlResponse("errors/404", nil)
	}

	contents,_ := ioutil.ReadFile(conf.ConfigInst.Docroot+"/"+obj.Path)
	output := blackfriday.Run(contents)

	title:=obj.Title
	title = strings.Replace(title, "-", " ", -1)
	title = strings.Replace(title, ".md", "", -1)

	subjects := make([]*core.TreeFolder, 0)
	for _,v := range core.SubjectIndexer {

		subjects = append(subjects, v)
	}

	hot := model.GetHotRanging()

	sidebar := sidebar(subjects, hot)
	nav := nav()
	footer := footer()

	editedAt := time.Unix(obj.EditedAt, 0).Format("Mon Jan 2,2006 at 15:04")

	base64img := ""

	//if obj.Img!="" {
	//	imgFp,err := os.Open("."+obj.Img)
	//
	//	if err!=nil {
	//		log.Error("", "cut img failed, open file failed, %s", err.Error())
	//	}else{
	//		m,_,err := image.Decode(imgFp)
	//		if err!=nil {
	//			log.Error("", "cut img failed, decode img failed, %s", err.Error())
	//		}else{
	//			rgbImg := m.(*image.YCbCr)
	//			subImg := rgbImg.SubImage(image.Rect(0, 0, 900, 300)).(*image.YCbCr) //图片裁剪x0 y0 x1 y1
	//
	//			emptyBuff := bytes.NewBuffer(nil)                  // 开辟一个新的空buff
	//			err = jpeg.Encode(emptyBuff, subImg, nil)            // img写入到buff
	//			if err!=nil {
	//				log.Error("", "cut img failed, encode img-buf failed, %s", err.Error())
	//			}else{
	//				//dist := make([]byte, 50000)                        //开辟存储空间
	//				//base64.StdEncoding.Encode(dist, emptyBuff.Bytes()) //buff转成base64
	//				//fmt.Println(string(dist))                          //输出图片base64(type = []byte)
	//				//_ = ioutil.WriteFile("/Users/leo/Documents/gopath/src/github.com/phpxin/mdblog/base64pic.txt", dist, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
	//
	//				base64img = base64.StdEncoding.EncodeToString(emptyBuff.Bytes())
	//			}
	//		}
	//	}
	//}

	return core.HtmlResponse("detail3", struct{
		Title string
		Intro string
		Desc string
		Contents template.HTML
		EditedAt string
		Img string
		Sidebar template.HTML
		Nav template.HTML
		Footer template.HTML
	}{
		title,
		obj.Intro,
		obj.Desc,
		template.HTML(string(output)) ,
		editedAt,
		base64img,
		template.HTML(sidebar) ,
		template.HTML(nav) ,
		template.HTML(footer) ,
	})
}