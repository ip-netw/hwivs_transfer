# IVS代理转换工具

华为IVS平台提供的图片的URL为端口IP不固定，在穿越网络时无法确定需要映射的端口

## 安装

直接在[release](https://github.com/lyonsdpy/hwivs_transfer/releases/tag/V0.0.1) 中下载相关可执行文件

也可以clone下载使用go编译

## 选项
win_amd64_hwivs.exe -port=8081 -log=./1.txt -level=debug
默认端口：8080
默认日志文件：无(仅打印终端日志)
默认日志级别：debug

## 使用示例
curl -X 'POST' -d \ 
        \ 'http://43.43.43.43:1183/imgu?
        \ Action=Download&
        \ NvrCode=6df709a6abdb480b8682c90d83fa1d34&
        \ PictureID=MDEjMDAwMyM8HOwpRgiXV57iXfCQ88qPyMZ4' 
        \ -o text.jpg http://127.0.0.1:8080/hwivs
