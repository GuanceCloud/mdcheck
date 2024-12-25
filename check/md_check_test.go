// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package check

import (
	"os"
	"path/filepath"
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

### 安装 DDAgent Nginx OpenTracing 插件`

		f := filepath.Join(t.TempDir(), "some.txt")
		assert.NoError(t, os.WriteFile(f, []byte(txt), 0600))

		res, fix := doMatch(f, false)
		t.Logf("res: %+#v, fix: %s", res, fix)

		assert.Len(t, res, 2)
	})
}
