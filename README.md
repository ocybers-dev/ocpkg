
# ocpkg

`ocpkg` 是一个为各模块提供通用功能的 Go 包集合，供 [ocybers-dev](https://github.com/ocybers-dev) 组织的项目使用。当前仓库仍在开发中，未正式发布，因此仅支持通过直接引用的方式进行使用。

## 目录
- [ocpkg](#ocpkg)
  - [目录](#目录)
  - [安装](#安装)
    - [注意事项](#注意事项)
  - [使用方法](#使用方法)
  - [示例](#示例)
  - [贡献](#贡献)

## 安装

由于 `ocpkg` 尚未发布到 Go 官方包管理仓库，请通过以下方式在项目中直接引用此仓库。

在你的 Go 模块中使用以下命令将 `ocpkg` 引入为依赖：

```bash
go get github.com/ocybers-dev/ocpkg
```

或在 `go.mod` 文件中直接添加：

```go
module your-module-name

go 1.20

require (
    github.com/ocybers-dev/ocpkg latest
)
```

然后运行以下命令以下载依赖：

```bash
go mod tidy
```

### 注意事项
- 由于 `ocpkg` 未正式发布，建议固定引用的版本，避免不兼容的代码更新带来影响。
- 使用时请确保代码兼容当前 `ocpkg` 的结构和接口。

## 使用方法

在你的代码中引入 `ocpkg` 包：

```go
import "github.com/ocybers-dev/ocpkg"
```

然后可以根据需要使用包内的模块和函数。

## 示例

以下是使用 `ocpkg` 提供的某个通用功能的简单示例：

```go
package main

import (
    "fmt"
    "github.com/ocybers-dev/ocpkg"
)

func main() {
    result := ocpkg.SomeFunction()
    fmt.Println(result)
}
```

请根据具体模块查看 `ocpkg` 中各功能的文档和注释。

## 贡献

欢迎对 `ocpkg` 提出建议或进行贡献！如有意贡献代码，请遵循以下流程：

1. Fork 本仓库。
2. 创建你的功能分支 (`git checkout -b feature/YourFeature`)。
3. 提交你的更改 (`git commit -m 'Add some feature'`)。
4. Push 到分支 (`git push origin feature/YourFeature`)。
5. 创建一个 Pull Request。
