# simple

- Go Web框架 - **Simple**





## 读取yaml配置文件
### 使用库[`go-yaml`](https://github.com/go-yaml/yaml)
```bash
go get github.com/go-yaml/yaml
```
### 对外提供的方法
```bash
> go doc github.com/go-yaml/yaml|grep '^func'

func Marshal(in interface{}) (out []byte, err error)
func Unmarshal(in []byte, out interface{}) (err error)
func UnmarshalStrict(in []byte, out interface{}) (err error)
```
- `Marshal`：表示序列化一个结构成为 YAML 格式
- `Unmarshal`：表示反序列化一个 YAML 格式文本成为一个结构
- `UnmarshalStrict`：表示严格反序列化，比如如果 YAML 格式文件中包含重复 key 的字段，那么使用 UnmarshalStrict 函数反序列化会出现错误。


```bash
go get github.com/mitchellh/mapstructure
```