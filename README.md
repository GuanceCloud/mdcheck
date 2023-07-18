# mdcheck

本工具用于检测 mkdocs 中的 Markdown 文档，用以规范 Markdown 的编写：

- 中英文混写的空格是否合规，如（空格用特殊字符展示以区分效果）：

```markdown
<!-- 错误写法 -->
中文with␣english
english␣with中文
中文with␣english然后又是中文
数字123然后又是中文
行内代码␣`abc_`然后又是中文
行内代码`_abc`␣然后又是中文


<!-- 正确写法 -->
中文␣with␣english
english␣with␣中文
中文␣with␣english␣然后又是中文
数字␣123␣然后又是中文
行内代码␣`abc_`␣然后又是中文
行内代码␣`_abc`␣然后又是中文
```

- 章节是否都手动加 ID，如

```markdown
<!-- 错误写法 -->
## 示例章节

<!-- 正确写法 -->
## 示例章节 {#some-section}
```

- 检查外链是否加了 tab 跳转，如

```markdown
<!-- 错误写法 --->
参见这里的 [Wiki 链接](https://wikipedia.org)

<!-- 正确写法 --->
参见这里的 [Wiki 链接](https://wikipedia.org){:target="_blank"}
```

## 编译

```shell
$ go build
...
```

## 运行

```shell
$ mdcheck -h
Usage of ./mdcheck:
  -autofix
        auto fix error
  -json
        set output in json format(2 space indent)
  -md-dir string
        markdown dirs
  -meta-dir string
        markdown meta dir
```

## TODO

- [ ] 整合已有的三方检测工具
    - [ ] 单词拼写（cspell）
    - [ ] Markdown 格式检测（markdownlint）
    - [ ] 404 检测（markdown-link-check）
- [ ] 工程建设
    - [ ] 添加 lint/makefile 等一系列周边设施
    - [ ] 构建 goreleaser 工作流
