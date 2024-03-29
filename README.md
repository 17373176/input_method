# 项目名称

输入法

创建输入法时，字典程序会从命令行参数中获得一系列的字典路径，字典路径表示一个本地文件路径或者是 HTTP URL。
如果路径不带 <http://> 或 <https://> 前缀，则从本地路径加载，其中文件名为拼音本身，后缀为 .dat。
如果路径带有 <http://> 或 <https://> 前缀，则需要通过 HTTP GET 请求从服务器加载，使用 GET 方法，返回的数据格式为 text/plain，其 HTTP Body 是字典数据，HTTP Path 的最后一段为拼音本身，后缀为.dat。
应采用并行方式加载词典，并且程序应支持词典文件的热更新。
无论是本地文件还是服务端远程文件，字典内的格式都是相同的：每个文件中包含多行内容，每一行包含一个汉字和这个汉字的出现频次得分，频次分最小为1，最大为10，均为整数，得分越大表示越是高频词。
输入文件中格式如果有误的话，忽略这个有误的文件；但输入文件中的空行不应算作格式有误。如果发生其它错误，例如文件打开失败等，也做忽略处理。

spell是一个字符串，可能是一个完整的拼音，也可能不是。例如下述之一："zhang", "zha", "zhan", "zh"。
查找规则为：

* 如果输入是一个完整的拼音（判断标准为有对应的词典文件），则返回该拼音下所有的汉字，按照高频分数从高到低，若两个字高频分数一样，则按照词典文件中的顺利。
例如输入"de"，则返回 ["的", "得", "地", "德"]。
* 如果输入不是一个完整的拼音（判断标准为没有对应的词典文件），则返回所有前缀与输入相同的拼音的汉字中最高频次的10个，具体排序为：
* 不同频次的汉字，频次越高的排在越前面
* 相同频次的汉字，根据对应的拼音的字母序排列，字母序越小的排在越前面。
* 相同频次的汉字，对应的拼音字母序也相同，则根据文件中的顺序排列。
例如输入两个词典，分别是zhang.dat和zhan.dat，其中zhang.dat为[长: 10, 张: 9]，zhan.dat为[展: 9, 战: 6, 站: 6]。
输入"zha"，则返回：["长", "展", "张", "战", "站"]。原因是：“长”为10分，排第一；“张”和“展”都为9分，“zhan”比“zhang”要短，排序靠前；“战”和“站”都是6分，根据文件中的顺序排。
* 结果中的多音字，只返回其最高权重的。

* 代码必须能够正常编译和运行。
* 要求实现的是一个输入法核心库，不要写成一个 HTTP API 应用。
* 如果你没有找到更好的数据结构，那么请使用 TrieTree 实现核心查找逻辑。不要用 map，map 不适合做前缀匹配。
* 如果使用并发，请仔细分析并发是否真的能带来提升，不要滥用自己不熟悉的特性和功能。
* 仔细处理每一处资源的使用，了解自己用到的每一个方法的注意事项（API 文档上都会有描述），不能出现并发死锁，不能有 Goroutine 泄漏，不能有 FD 泄漏。
* 代码需要在Agile流水上线添加测试覆盖率统计，单测覆盖率不能低于95%，应当尽可能接近100%。
* 单测必须包含有效的case，并且对结果做assert，不能只在单测中运行代码，对运行结果却不做任何对错判断。
* 测试代码应该与被测试代码放在同一目录，但是单独的测试数据应该放在类似 testdata 之类单独的目录。
* 不可以依赖外部实际存在的 HTTP Server,可以使用 net/http/httptest 创建一个 test server。
* 代码需要开启 golint 和 govet 静态检查，并且不可出现中高危代码问题。
* 项目需要开启 go mod 进行依赖管理。

## 快速开始

启动程序：命令行执行 go run main.go ./data/a.dat ./data/ai.dat ./data/an.dat ./data/ang.dat ./data/de.dat ./data/den.dat ./data/deng.dat ./data/zhan.dat ./data/zhang.dat

输入对应拼音，回车得到对应候选字

## 测试数据

data 目录为测试数据，主要有坑的测试点：

1. 由于字典树同一节点层是无序的，不能单独利用自上而下的已有顺序来判断字典序
2. 前缀匹配只返回符合匹配的，不匹配的不返回，例如输入"zh"，应返回 zhan,zhang，而不应该返回 za,zan,ze
3. 精确匹配也有可能词典里没有词，因此不能通过判断精确匹配返回是否有词作为前缀匹配的条件
4. 注意排序算法是否是稳定排序
5. 多音字需要注意去重，"的"，"地"

## 测试

自动化测试，采用 gotests 工具生成单测框架，补充测试用例

安装 gotest
go install github.com/cweill/gotests/...

对指定文件的指定函数生成测试函数，输出到控制台
gotests -only swap ./model/service/ime.go

-w 输出到：源文件名称_test.go 的文件
gotests -only -w swap ./model/service/ime.go

执行单测并计算覆盖率: go test -v -cover
执行 go test -v -coverprofile=c.out && go tool cover -html=c.out -o=tag.html 或 直接在测试函数上方点击 run test, go-ut:run test

代码静态检查/扫描 golangci-lint

<https://golangci-lint.run/>

安装命令：
go install github.com/golangci/golangci-lint/cmd/golangci-lint@master

把如下代码加入命名为.golangci.yml的文件，放入根目录

```
output:
  format: json
  print-issued-lines: true
linters:
  # enable-all: true
  # disable:
  #   - deadcode
  disable-all: true
  enable:
    - stylecheck
    - revive
    - gosimple
    - gofmt
    - lll
    - errcheck
    - errorlint
    - govet
    - gocyclo
    - goimports
linters-settings:
  lll:
    # max line length, lines longer will be reported. Default is 120.
    # '\t' is counted as 1 character by default, and can be changed with the tab-width option
    line-length: 160
    # tab width in spaces. Default to 1.
    tab-width: 1
  errcheck:
    # report about not checking of errors in type assertions: `a := b.(MyStruct)`;
    # default is false: such cases aren't reported by default.
    check-type-assertions: false

    # report about assignment of errors to blank identifier: `num, _ := strconv.Atoi(numStr)`;
    # default is false: such cases aren't reported by default.
    check-blank: false
  gocyclo:
    # Minimal code complexity to report.
    # Default: 30 (but we recommend 10-20)
    min-complexity: 30
```

在代码库根目录执行：
golangci-lint run
