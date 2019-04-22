package routes

import (
	"fmt"
	"go-iris-curd/main/web/supports"
	"io"
	"os"
	"os/exec"
	"strings"

	conf "go-iris-curd/main/inits/parse"

	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)

// CommonHub 通用路由
func CommonHub(party iris.Party) {
	// demo测试API模块
	common := party.Party("/common")
	{
		upload := common.Party("/upload")

		upload.Post("/", hero.Handler(CreateUpload)) // 新增一个
		// upload.Get("/", hero.Handler(ProductList))                   // 产品列表
		// upload.Get("/{pid:long}", hero.Handler(GetOneProduct))       // 获得一个
		// upload.Put("/", hero.Handler(UpdateProduct))                 // 更新类别
		// upload.Delete("/{uids:string}", hero.Handler(DeleteProduct)) // 删除产品
		// upload.Delete("/del", hero.Handler(DeleteProducts))          // 批量删除产品
	}

}

// CreateUpload 增加产品
func CreateUpload(ctx iris.Context) {

	file, info, err := ctx.FormFile("uploadfile")

	if err != nil {
		ctx.Application().Logger().Errorf(err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, supports.OptionFailur, nil)
		return
	}

	defer file.Close()
	fname := info.Filename

	path := conf.O.UploadFile
	fmt.Println(path)

	// Create a file with the same name
	// assuming that you have a folder named 'uploads'
	out, err := os.OpenFile(path+fname,
		os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}
	defer out.Close()

	io.Copy(out, file)

	supports.OkR(ctx, supports.OptionSuccess)

}

func getCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	checkErr(err)
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
