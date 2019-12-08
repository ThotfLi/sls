Small Log System(SLS)

定位
按照不同等级记录不同类型日志。

运行方式
将SLS项目文件复制到项目目录
import "SLS"
生成NewLogMessage(flag int,messagepath string,messagename string)
flag 代表日志等级
messagepath 日志文件位置
messagenam 日志文件名

直接调用Loging(str string,flag)
str 要写入的内容
flag 设置记录内容的日志等级
如果flag < LogMessage.Flag，将不能被记录在日志文件中只会输出在控制台
eg：





设置Sflag 可以设置日志输出格式
Sflag.TIME     记录调用时间
Sflage.STACK   记录调用栈
Sflage.LINE    记录调用行
Sflage.FILE    记录调用文件
