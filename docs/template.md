模板引擎可以合并模板和上下文数据，产生最终的html

无逻辑模板引擎
逻辑嵌入模板引擎

解析模板源   输入 字符串或者模板文件
执行解析好的模板

解析模板
ParseGlob  *.html
Parse
Loopup
Must

执行模板
Execute
ExecuteTemplate


action {{ xxx }}
### 条件类

```
{{ if arg }}
some content
{{ end }}

{{ if arg }}
some content
{{ else }}
other content
{{ end }}
```
### 迭代、遍历类
遍历数组 slice map channel等数据结构
{{ range array }}
dot is set to the element
{{ end }}
回落机制
{{ else }}
{{ end }}


### 设置类
{{ with arg }}
dot is set to arg
{{ end }}
回落机制

### 包含类
在模板里边包含其他的模板

{{ template "name" }}
传递数据
{{ template "name" arg }}
### 定义类
define action


## 函数和管道
action中设置变量，变量以$开头


管道是按顺序连接到一起的参数，函数和方法

{{ p1 | p2 | p3 }}
{{ 12.3456 | printf "%.2f" }}

自定义函数

可以接收任意数量的输入参数
返回
- 一个值
- 一个值 + 一个错误

内置函数
define template block
html js urlquery 对字串进行转义，防止安全问题
index
print printf println
len 
with

template.Funcs(funcMap FuncMap) *Template

type FuncMap map[string]interface{}

1、创建一个FuncMap
2、函数附加到FuncMap

## 组合模板


逻辑运算符

eq ne
lt gt
le ge
and or not
