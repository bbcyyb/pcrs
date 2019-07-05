# 如何设计基于Golang的Restful风格后台架构

## 在设计的过程中，可以遵循如下约束

- 独立于框架。该架构不会依赖于某些功能强大的软件库存在。这可以让你使用这样的框架作为工具，而不是让你的系统陷入到框架的限制的约束中。

- 可测试性。业务规则可以在没有 UI， 数据库，Web 服务或其他外部元素的情况下进行测试。

- 独立于 UI 。在无需改变系统的其他部分情况下， UI 可以轻松的改变。例如，在没有改变业务规则的情况下，Web UI 可以替换为控制台 UI。

- 独立于数据库。你可以用 Mongo， BigTable， CouchDB 或者其他数据库来替换 Oracle 或 SQL Server，你的业务规则不要绑定到数据库。

- 独立于外部媒介。 实际上，你的业务规则可以简单到根本不去了解外部世界。

## Golang单元测试

### 使用 Table Driven 的方式写测试代码

```go
func TestMod(t * testing.T) {
        tests: = [] struct {
            a int
            b int
            r int
            hasErr bool
        } {
            {
                a: 42,
                b: 9,
                r: 6,
                hasErr: false
            }, {
                a: -1,
                b: 9,
                r: 8,
                hasErr: false
            }, {
                a: -1,
                b: -9,
                r: -1,
                hasErr: false
            }, {
                a: 42,
                b: 0,
                r: 0,
                hasErr: true
            },
        }
    
        for row, test: = range tests {
            r, err: = Mod(test.a, test.b)
            if test.hasError {
                if err == nil {
                    t.Errorf("should have error, row: %d", row)
                }
                continue
            }
            if err != nil {
                t.Errorf("should not have error, row: %d", row)
            }
            if r != test.r {
                t.Errorf("r is expected to be %d but now %d, row: %d", test.r, r, row)
            }
        }
    }
```

### 使用 testify/assert 简化条件判断

```go
import "github.com/stretchr/testify/assert"

for row, test := range tests {
    r, err := Mod(test.a, test.b)
    if test.hasError {
        assert.Error(t, err, "row %d", row)
        continue
    }
    assert.NoError(t, err, "row %d", row)
    assert.Equal(t, test.r, r, "row %d", row)
}
```

### 使用 testify/mock 隔离第三方依赖或者复杂调用

很多时候，测试环境不具备 routine 执行的必要条件。比如查询 consul 里的 KV，即使准备了测试consul，也要先往里面塞测试数据，十分麻烦。又比如查询 AWS S3 的文件列表，每个开发人员一个测试 bucket 太混乱，大家用同一个测试 bucket 更混乱。必须找个方式伪造 consul client 和 AWS S3 client。通过伪造 consul client 查询 KV 的方法，免去连接 consul， 直接返回预设的结果。

testfiy/mock 使得伪造对象的输入输出值可以在运行时决定。更多技巧可看 testify/mock 的文档。