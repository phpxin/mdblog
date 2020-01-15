# mdblog
> A open source blog system       
1. 通过固定路径导入 Markdown 文档      
2. 根据指令自动对路径使用路径进行分类，将内容分类生成json保存，并生成主页     

## 编译运行

> govendor sync      
> go      
> go build .       
> ./mdblog conf/conf.json    

## linux 环境发布
> GOARCH=amd64 GOOS=linux go build -o mdblog_amd64 .     
> pkill mdblog_amd64   
> nohup ./mdblog_amd64 ./conf.json >mdblog_amd64.log &    
> 

