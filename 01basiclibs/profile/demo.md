
https://www.jianshu.com/p/4e4ff6be6af9

# 1.测试用例
go test -bench=. -cpuprofile=cpu.prof

go tool pprof -http=:6060 cpu2.prof

#


Bo-Yi Wu
2400位订阅者

# golang博客：

## Bo-Yi Wu

https://blog.wu-boy.com/2020/12/embedding-files-in-go-1-16/

Download By FindYoutube.net
Go 語言官方維護團隊 rsc 之前在 GitHub Issue 上面提出要在 go command line 直接支援 Embedding Files，沒想到過沒幾個月，就直接實現出來了，並且預計在 2021 的 go 1.16 版本直接支援 embed 套件。有了這個功能，就可以將靜態檔案或專案設定檔直接包起來，這樣部署就更方便了。底下來看看官方怎麼使用。

00:00 GO 1.16 推出 embed 套件
01:33 如何使用 embed 套件
03:37 Gin 搭配 embed 編譯靜態檔案
06:08 Gin 範例演示

中文部落格: https://blog.wu-boy.com/2020/12/embed...
Go 1.16: https://tip.golang.org/doc/go1.16
Go embed package: https://tip.golang.org/pkg/embed/

Go 語言實戰課程: https://www.udemy.com/course/golang-f...
Drone CI/CD 自動化課程: https://www.udemy.com/course/devops-o...
Docker 實戰教學: https://www.udemy.com/course/docker-p...

需要架構，洽談，教學，教育訓練，以下是我聯絡方式
email: appleboy.tw@gmail.com
line ID: appleboy46 或手機 0934353293

小額贊助: http://bit.ly/donate-appleboy
部落格: https://blog.wu-boy.com
Twitter: https://twitter.com/appleboy
Facebook: https://www.facebook.com/appleboy46
投影片: https://speakerdeck.com/appleboy
GitHub: https://github.com/appleboy
Youtube: http://bit.ly/youtube-boy