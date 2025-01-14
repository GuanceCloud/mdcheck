// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package check

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	T "testing"

	"github.com/stretchr/testify/assert"
)

func Test_doMatch(t *T.T) {
	t.Run(`section`, func(t *T.T) {
		txt := `
### 前提条件

- [x] 安装 nginx (>=1.9.13)

***该模块只支持 linux 操作系统***

### 安装 Nginx OpenTracing 插件 {#install-plugin}

Nginx OpenTracing 插件是 OpenTracing 开源的链路追踪插件，基于 C++ 编写，可以工作于

- 配置插件

### 安装 DDAgent Nginx OpenTracing 插件

### {{some-template-name}}
`

		f := filepath.Join(t.TempDir(), "some.txt")
		assert.NoError(t, os.WriteFile(f, []byte(txt), 0600))

		co := defaultOpt()
		res, fix := co.doMatch(f)
		t.Logf("res: %+#v, fix: %s", res, fix)

		assert.Len(t, res, 2)
	})
}

type mdref struct {
	text, md, section, url string
}

func parseMarkdownLinks(markdown string) (refs []*mdref) {
	// 1. 创建正则表达式
	re := regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)

	matches := re.FindAllStringSubmatch(markdown, -1)
	for _, match := range matches {
		if len(match) != 3 {
			continue
		}
		// 提取链接文字和链接目标
		linkText := match[1]
		linkTarget := match[2]

		if strings.HasPrefix(linkTarget, "http://") || strings.HasPrefix(linkTarget, "https://") { // seems a web URL
			refs = append(refs, &mdref{
				text: linkText,
				url:  linkTarget,
			})
		} else {
			arr := strings.SplitN(linkTarget, "#", 2)
			switch len(arr) {
			case 1: // no section
				refs = append(refs, &mdref{
					text: linkText,
					md:   arr[0],
				})
			case 2: // with section
				refs = append(refs, &mdref{
					text:    linkText,
					md:      arr[0],
					section: arr[1],
				})
			}
		}
	}

	return
}

func Test_parseMarkdownLinks(t *T.T) {
	markdown := `[here](http://example.com/abc/def){:target="_blank"}
[here](some.md#some-section)
[中文链接](http://example.com/中文)
[another link](some/other.md#another-section)
[simple link](some-page.md)
<!-- markdownlint-disable MD404 -->
[中文链接](http://example.com/中文)
[another link](some/other.md#another-section)
[simple link](some-page.md)
<!-- markdownlint-enable -->
[xxx link](some-page.md)`

	ignoreMD404 := false
	for idx, ln := range strings.Split(markdown, "\n") {
		if ln == "<!-- markdownlint-disable MD404 -->" {
			t.Logf("ignoreMD404 set")
			ignoreMD404 = true
			continue
		}

		if ln == "<!-- markdownlint-enable -->" {
			t.Logf("ignoreMD404 unset")
			ignoreMD404 = false
			continue
		}

		if ignoreMD404 {
			t.Logf("ignore line %d: %s", idx, ln)
			continue
		}

		refs := parseMarkdownLinks(ln)
		if len(refs) == 0 {
			t.Logf("no reference in ln: %s", ln)
		} else {
			for _, ref := range refs {
				t.Logf("[%d] %+#v", idx, ref)
			}
		}
	}
}
