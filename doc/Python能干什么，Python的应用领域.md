# Python能干什么，Python的应用领域

[Python](http://c.biancheng.net/python/) 作为一种功能强大的编程语言，因其简单易学而受到很多开发者的青睐。那么，Python 的应用领域有哪些呢？

Python 的应用领域非常广泛，几乎所有大中型互联网企业都在使用 Python 完成各种各样的任务，例如国外的 Google、Youtube、Dropbox，国内的百度、新浪、搜狐、腾讯、阿里、网易、淘宝、知乎、豆瓣、汽车之家、美团等等。

概括起来，Python 的应用领域主要有如下几个。

## Web应用开发

Python 经常被用于 Web 开发，尽管目前 [PHP](http://c.biancheng.net/php/)、JS 依然是 Web 开发的主流语言，但 Python 上升势头更猛劲。尤其随着 Python 的 Web 开发框架逐渐成熟（比如 Django、flask、TurboGears、web2py 等等），程序员可以更轻松地开发和管理复杂的 Web 程序。

例如，通过 mod_wsgi 模块，Apache 可以运行用 Python 编写的 Web 程序。Python 定义了 WSGI 标准应用接口来协调 HTTP 服务器与基于 Python 的 Web 程序之间的通信。

举个最直观的例子，全球最大的搜索引擎 Google，在其网络搜索系统中就广泛使用 Python 语言。另外，我们经常访问的集电影、读书、音乐于一体的豆瓣网（如图 1 所示），也是使用 Python 实现的。

不仅如此，全球最大的视频网站 Youtube 以及 Dropbox（一款网络文件同步工具)也都是用 Python 开发的。

## 自动化运维

很多操作系统中，Python 是标准的系统组件，大多数 Linux 发行版以及 NetBSD、OpenBSD 和 Mac OS X 都集成了 Python，可以在终端下直接运行 Python。

有一些 Linux 发行版的安装器使用 Python 语言编写，例如 Ubuntu 的 Ubiquity 安装器、Red Hat Linux 和 Fedora 的 Anaconda 安装器等等。

另外，Python 标准库中包含了多个可用来调用操作系统功能的库。例如，通过 pywin32 这个软件包，我们能访问 Windows 的 COM 服务以及其他 Windows API；使用 IronPython，我们能够直接调用 .Net Framework。

通常情况下，Python 编写的系统管理脚本，无论是可读性，还是性能、代码重用度以及扩展性方面，都优于普通的 shell 脚本。

## 人工智能领域

人工智能是项目非常火的一个研究方向，如果要评选当前最热、工资最高的 IT 职位，那么人工智能领域的工程师最有话语权。而 Python 在人工智能领域内的机器学习、神经网络、深度学习等方面，都是主流的编程语言。

可以这么说，基于[大数据](http://c.biancheng.net/big_data/)分析和深度学习发展而来的人工智能，其本质上已经无法离开 Python 的支持了，原因至少有以下几点：

1. 目前世界上优秀的人工智能学习框架，比如 Google 的 TransorFlow（神经网络框架）、FaceBook 的 PyTorch（神经网络框架）以及开源社区的 Karas 神经网络库等，都是用 Python 实现的；
2. 微软的 CNTK（认知工具包）也完全支持 Python，并且该公司开发的 VS Code，也已经把 Python 作为第一级语言进行支持。
3. Python 擅长进行科学计算和数据分析，支持各种数学运算，可以绘制出更高质量的 2D 和 3D 图像。

> VS Code 是微软推出的一款代码编辑工具（IDE），有关它的下载、安装和使用，后续章节会做详细介绍。

总之，AI 时代的来临，使得 Python 从众多编程语言中脱颖而出，Python 作为 AI 时代头牌语言的位置，基本无人可撼动！

## 网路爬虫

Python 语言很早就用来编写网络爬虫。Google 等搜索引擎公司大量地使用 Python 语言编写网络爬虫。

从技术层面上将，Python 提供有很多服务于编写网络爬虫的工具，例如 urllib、Selenium 和 BeautifulSoup 等，还提供了一个网络爬虫框架 Scrapy。

## 科学计算

自 1997 年，NASA 就大量使用 Python 进行各种复杂的科学运算。

并且，和其它解释型语言（如 shell、js、PHP）相比，Python 在数据分析、可视化方面有相当完善和优秀的库，例如 NumPy、SciPy、Matplotlib、pandas 等，这可以满足 Python 程序员编写科学计算程序。

## 游戏开发

很多游戏使用 [C++](http://c.biancheng.net/cplus/) 编写图形显示等高性能模块，而使用 Python 或 Lua 编写游戏的逻辑。和 Python 相比，Lua 的功能更简单，体积更小；而 Python 则支持更多的特性和数据类型。

比如说，国际上指明的游戏 Sid Meier's Civilization就是使用 Python 实现的。


除此之外，Python 可以直接调用 Open GL 实现 3D 绘制，这是高性能游戏引擎的技术基础。事实上，有很多 Python 语言实现的游戏引擎，例如 Pygame、Pyglet 以及 Cocos 2d 等。

以上也仅是介绍了 Python 应用领域的“冰山一角”，例如，还可以利用 Pygame 进行游戏编程；用 PIL 和其他的一些工具进行图像处理；用 PyRo 工具包进行机器人控制编程，等等。有兴趣的读者，可自行搜索资料进行详细了解。

