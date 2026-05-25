# Go 自动化单元测试教程

## 目录

1. [什么是自动化单元测试](#第-1-章)
2. [表驱动测试](#第-2-章)
3. [子测试 `t.Run`](#第-3-章)
4. [黑盒测试 vs 白盒测试](#第-4-章)
5. [基准测试](#第-5-章)
6. [示例测试](#第-6-章)
7. [测试覆盖率](#第-7-章)
8. [Makefile 自动化](#第-8-章)
9. [CI/CD 流水线](#第-9-章)
10. [Pre-commit 钩子](#第-10-章)

附录 A: [命令速查](#附录-a命令速查)
附录 B: [项目测试清单](#附录-b项目测试清单)

---

## 第 1 章：什么是自动化单元测试

### 为什么需要自动化测试

自动化单元测试是指编写代码来验证另一段代码的正确性，并在每次代码变更时自动运行这些验证。Go 语言将测试作为一等公民内置在工具链中——不需要任何第三方测试框架。

Go 测试的核心约定：

| 约定 | 说明 |
|------|------|
| 测试文件以 `_test.go` 结尾 | Go 编译器自动识别 |
| 测试函数以 `Test` 开头 | 签名为 `func TestXxx(t *testing.T)` |
| 运行测试 | `go test ./...` |

### 心智模型

每个测试就是一个断言：**给定某个输入，期望得到特定的输出**。测试要么通过（PASS），要么失败（FAIL）。没有中间状态。

```go
// 来自本项目的真实测试 — search/binarysearch_test.go
func TestBinarySearch(t *testing.T) {
    tests := []struct {
        name     string   // 用例名称
        arr      []int    // 输入数组
        target   int      // 查找目标
        expected int      // 期望结果
    }{
        {name: "在中间找到", arr: []int{1, 3, 5, 7, 9}, target: 5, expected: 2},
        {name: "未找到",     arr: []int{1, 3, 5, 7, 9}, target: 4, expected: -1},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := search.BinarySearch(tt.arr, tt.target)
            if got != tt.expected {
                t.Errorf("BinarySearch(%v, %d) = %d, want %d", tt.arr, tt.target, got, tt.expected)
            }
        })
    }
}
```

### 运行测试的基础命令

```bash
# 运行所有包的测试
go test ./...

# 详细输出
go test -v ./...

# 运行指定包
go test -v ./search/

# 运行指定测试函数
go test -v -run TestBinarySearch ./search/

# 绕过测试缓存（强制重新运行）
go test -count=1 ./...
```

---

## 第 2 章：表驱动测试

### 什么是表驱动测试

表驱动测试（Table-Driven Tests）是 Go 社区公认的最佳测试模式。它将测试用例组织为匿名结构体切片，每个用例包含 `name`（名称）、`input`（输入）、`expected`（期望）等字段。

**为什么使用表驱动测试？**

- **减少样板代码**：一个 for 循环处理所有用例，而不是为每个用例写一个测试函数
- **可读性强**：一眼看清所有被测试的场景，新增用例只需在表中追加一行
- **易于维护**：修改测试逻辑只需改一处，所有用例自动受益
- **失败定位精确**：结合 `t.Run` 子测试，能精确定位到哪个用例失败了

### 完整示例

来自 `mathutil/fibonacci_test.go`：

```go
func TestFibonacci(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected int
    }{
        {name: "负数输入", input: -1, expected: -1},
        {name: "F(0)",    input: 0,  expected: 0},
        {name: "F(1)",    input: 1,  expected: 1},
        {name: "F(2)",    input: 2,  expected: 1},
        {name: "F(10)",   input: 10, expected: 55},
        {name: "F(20)",   input: 20, expected: 6765},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := mathutil.Fibonacci(tt.input)
            if got != tt.expected {
                t.Errorf("Fibonacci(%d) = %d, want %d", tt.input, got, tt.expected)
            }
        })
    }
}
```

### 测试场景覆盖清单

每个函数至少应考虑以下场景：

| 场景类别 | 说明 | 本项目示例 |
|---------|------|-----------|
| 正常输入 | 典型的日常输入值 | 所有测试函数均覆盖 |
| 空输入 | 空切片、空字符串、空数据结构 | `[]int{}`, 空栈/空队列 |
| 单元素 | 长度为 1 的切片或链表 | sort、search、ds 包 |
| 边界条件 | 第一个/最后一个元素、最大值/最小值 | search 包 |
| 错误路径 | 非法操作（如空栈 Pop） | ds 包 |
| 重复值 | 排序中的重复元素 | sort 包 |
| 负数 | 负数输入的处理 | mathutil 包 |
| 不变量 | 函数不应修改原始数据 | `TestMergeSortImmutability` |

### 表驱动测试的最佳实践

1. **始终包含 `name` 字段**：描述性名称让 CI 日志中的失败信息一目了然
2. **用 `t.Errorf` 而非 `t.Fatalf`**：`Errorf` 报告失败但继续运行其他用例，`Fatalf` 会立即停止整个测试
3. **用 `reflect.DeepEqual` 比较切片和映射**：Go 的 `==` 不能直接比较引用类型
4. **每个用例使用独立数据**：避免用例之间共享可变数据导致交叉污染

---

## 第 3 章：子测试 `t.Run`

### `t.Run` 的作用

`t.Run(name, func)` 将每个测试用例注册为独立的**子测试**。这带来了四个关键好处：

1. **层次化输出**：每个子测试独立显示 PASS/FAIL，结构清晰
2. **独立运行**：可以用 `go test -run TestFibonacci/F(10)` 单独运行一个子测试
3. **独立失败**：一个用例失败不影响其他用例继续执行
4. **CI 日志友好**：失败信息精确到用例名称，便于快速定位

### 子测试的输出格式

```
=== RUN   TestFibonacci
=== RUN   TestFibonacci/负数输入
=== RUN   TestFibonacci/F(0)
=== RUN   TestFibonacci/F(1)
=== RUN   TestFibonacci/F(10)
--- PASS: TestFibonacci (0.00s)
    --- PASS: TestFibonacci/负数输入 (0.00s)
    --- PASS: TestFibonacci/F(0) (0.00s)
    --- PASS: TestFibonacci/F(1) (0.00s)
    --- PASS: TestFibonacci/F(10) (0.00s)
```

缩进层次清楚地展示了父子关系，失败时一目了然。

### 运行特定子测试

```bash
# 只运行 Fibonacci 中 F(10) 这个子测试
go test -v -run "TestFibonacci/F\(10\)" ./mathutil/

# 运行所有包含 "empty" 的测试用例
go test -v -run "empty" ./...
```

### 并行子测试

对于互相独立的子测试，可以用 `t.Parallel()` 并行运行以加速测试：

```go
for _, tt := range tests {
    tt := tt  // Go 1.22 之前需要捕获循环变量
    t.Run(tt.name, func(t *testing.T) {
        t.Parallel()  // 标记为可并行
        // ... 测试逻辑
    })
}
```

> **注意**：`t.Parallel()` 在子测试间共享父测试的生命周期。如果测试数据较大或涉及共享状态，需谨慎使用。

---

## 第 4 章：黑盒测试 vs 白盒测试

### 两种测试包声明

Go 允许两种测试包声明方式：

| 声明 | 类型 | 可访问 | 适用场景 |
|------|------|--------|---------|
| `package foo` | 白盒测试（内部） | 可访问未导出的标识符 | 需要测试内部实现细节 |
| `package foo_test` | 黑盒测试（外部） | 只能访问导出的 API | 模拟真实调用者 |

### 本项目为什么全部使用黑盒测试

本项目所有测试文件都使用 `package foo_test` 外部测试包。这是有意为之：

1. **模拟真实调用者**：测试代码的视角和外部用户完全一致
2. **关注行为而非实现**：测试「做什么」（what），而非「怎么做」（how）
3. **防止测试脆弱**：重构内部实现时，测试不会因为改了内部细节而失效
4. **避免导入循环**：`package foo_test` 可以安全导入其他包

### 示例：测试 MergeSort 的不变性

来自 `sort/mergesort_test.go`：

```go
func TestMergeSortImmutability(t *testing.T) {
    original := []int{3, 1, 2}
    copied := make([]int, len(original))
    copy(copied, original)  // 保存原始数据的副本

    sort.MergeSort(original)  // 执行排序

    // 验证：原始切片不应被修改
    if !reflect.DeepEqual(original, copied) {
        t.Errorf("MergeSort 修改了原切片: %v -> %v", copied, original)
    }
}
```

这个测试验证的是**行为保证**（不修改原数据），不需要访问任何内部状态。

### 何时需要白盒测试

- 测试复杂的未导出辅助函数
- 需要验证内部状态转换
- 需要设置或检查未导出的字段

你可以在同一目录下同时放 `package foo` 和 `package foo_test` 的 `_test.go` 文件，Go 完全支持这种混合模式。

---

## 第 5 章：基准测试

### 什么是基准测试

Go 的基准测试用于测量代码性能。函数签名为 `func BenchmarkXxx(b *testing.B)`，Go 运行时会自动决定迭代次数 `b.N` 以获得稳定的测量结果。

```bash
# 运行所有基准测试
go test -bench=. ./...

# 同时统计内存分配
go test -bench=. -benchmem ./...

# 多次运行以获得更稳定的结果
go test -bench=. -count=5 ./...
```

### 基本写法

来自 `mathutil/benchmark_test.go`：

```go
func BenchmarkFibonacci(b *testing.B) {
    b.Run("n=10", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            mathutil.Fibonacci(10)
        }
    })
    b.Run("n=50", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            mathutil.Fibonacci(50)
        }
    })
}
```

通过子基准测试（`b.Run`），可以在一次运行中对比不同输入规模下的性能差异。

### 计时器控制

当每轮迭代需要准备工作时，使用 `StopTimer` / `StartTimer` 排除准备时间：

来自 `sort/benchmark_test.go`：

```go
func BenchmarkBubbleSort(b *testing.B) {
    original := generateRandomSlice(1000)
    b.Run("size=1000", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            b.StopTimer()                        // 暂停计时
            arr := make([]int, len(original))
            copy(arr, original)                  // 准备工作不计入
            b.StartTimer()                       // 恢复计时
            sort.BubbleSort(arr)                 // 只有这部分被计时
        }
    })
}
```

关键 API：

| API | 用途 |
|-----|------|
| `b.ResetTimer()` | 在循环**之前**调用一次，清除之前的计时 |
| `b.StopTimer()` | 暂停计时 |
| `b.StartTimer()` | 恢复计时 |

### 算法性能对比

运行 `make bench` 可以直观看到不同算法的性能差异：

| 算法 | 时间复杂度 | 1000 个元素 | 10000 个元素 |
|------|---------|------------|-------------|
| BubbleSort | O(n²) | ~1.6ms | ~160ms |
| QuickSort | O(n log n) | ~0.1ms | ~1ms |
| BinarySearch | O(log n) | ~11ns | ~13ns |
| LinearSearch | O(n) | ~10μs | ~100μs |

### 理解基准测试输出

```
BenchmarkBubbleSort/size=1000-8    634    1635969 ns/op    0 B/op    0 allocs/op
```

| 字段 | 含义 |
|------|------|
| `634` | 运行了 634 轮迭代 |
| `1635969 ns/op` | 每轮操作平均耗时 1.6ms |
| `0 B/op` | 每轮操作分配的内存字节数 |
| `0 allocs/op` | 每轮操作的内存分配次数 |

### 注意事项

- **编译器优化可能消除无用代码**：确保基准结果被实际使用（赋值给包级别变量等）
- **系统负载影响结果**：在安静的环境下运行以获得可重复的结果
- **`b.ResetTimer()` 放在循环外**：每轮迭代前的重置应该在循环内部通过 StopTimer/StartTimer 处理

---

## 第 6 章：示例测试

### Go 独有的「文档即测试」模式

Go 的 `Example` 测试同时做两件事：

1. **可运行的测试**：随 `go test` 一起执行并验证输出是否正确
2. **嵌入文档**：`go doc` 命令会展示示例代码作为函数的用法文档

这是 Go 生态中非常独特的特性——文档不会过时，因为它本身就是会运行的测试。

### 命名约定

| 命名格式 | 用途 | 示例 |
|---------|------|------|
| `Example()` | 整个包的示例 | `func Example() { ... }` |
| `ExampleF()` | 函数 `F` 的示例 | `func ExampleGreeting() { ... }` |
| `ExampleT()` | 类型 `T` 的示例 | `func ExampleStack() { ... }` |
| `ExampleT_M()` | 类型 `T` 的方法 `M` | `func ExampleStack_Push() { ... }` |
| `ExampleF_suffix()` | 函数 `F` 的附加示例 | `func ExampleFibonacci_negative() { ... }` |

### 最简单的示例

来自 `tttys/example_test.go`——本项目最简单的示例测试：

```go
func ExampleGreeting() {
    fmt.Println(tttys.Greeting())
    // Output: Hello from tttys!
}
```

### 多行输出示例

来自 `search/example_test.go`：

```go
func ExampleBinarySearch() {
    arr := []int{1, 3, 5, 7, 9}
    fmt.Println(search.BinarySearch(arr, 7))
    fmt.Println(search.BinarySearch(arr, 4))
    // Output:
    // 3
    // -1
}
```

### 无序输出

如果输出行的顺序不固定（例如遍历 map），使用 `// Unordered output:`：

```go
func ExampleMap() {
    m := map[string]int{"a": 1, "b": 2}
    for k, v := range m {
        fmt.Println(k, v)
    }
    // Unordered output:
    // a 1
    // b 2
}
```

### 类型方法示例

来自 `ds/example_test.go`：

```go
func ExampleStack_Push() {
    s := ds.NewStack()
    s.Push(10)
    s.Push(20)
    val, _ := s.Peek()
    fmt.Println(val)
    // Output: 20
}
```

### 格式注意事项

- Go 切片的 `fmt.Println` 输出是 `[1 2 3]`（空格分隔，没有逗号），新手容易写错
- `// Output:` 注释必须与标准输出**逐字节完全匹配**，包括空格和换行
- `// Output:` 必须紧跟在函数体最后一个语句之后，中间不能有空行

### 本项目所有示例测试总览

| 文件 | 示例函数 | 演示特性 |
|------|---------|---------|
| `tttys/example_test.go` | `ExampleGreeting`, `ExampleInitialize` | 最简单的入门示例 |
| `utils/example_test.go` | `ExampleConfigMessage` | 中文/Unicode 输出验证 |
| `mathutil/example_test.go` | `ExampleFibonacci`, `ExampleFibonacci_negative` 等 | 多后缀示例、多行输出 |
| `sort/example_test.go` | `ExampleBubbleSort`, `ExampleQuickSort` 等 | 切片输出格式 |
| `search/example_test.go` | `ExampleBinarySearch` | 多行输出 |
| `ds/example_test.go` | `ExampleStack_Push`, `ExampleQueue_Dequeue` 等 | `Type_Method` 命名 |

---

## 第 7 章：测试覆盖率

### 覆盖率命令

```bash
# 显示每个包的覆盖率百分比
go test -cover ./...

# 生成覆盖率文件
go test -coverprofile=coverage.out ./...

# 按函数查看覆盖率详情
go tool cover -func=coverage.out

# 生成 HTML 可视化报告（在浏览器中查看）
go tool cover -html=coverage.out -o coverage.html
```

### 本项目的覆盖率

```
$ go test -cover ./...
DeepTest/mathutil    coverage: 100.0% of statements
DeepTest/sort        coverage: 100.0% of statements
DeepTest/search      coverage: 100.0% of statements
DeepTest/ds          coverage: 100.0% of statements
DeepTest/tttys       coverage: 100.0% of statements
DeepTest/utils       coverage: 100.0% of statements
```

### 100% 覆盖率就是目标吗？

**不是。** 100% 语句覆盖率只意味着每行代码都被执行过，但绝不意味着：

- 所有可能的输入组合都被测试了
- 所有边界条件都被覆盖了（例如整数溢出）
- 并发安全性被验证了
- 逻辑错误被发现了（代码可以"跑过"但结果是错的）

正确的心态是：**覆盖率是地板，不是天花板。** 它是质量的最低保障，不能替代好的测试用例设计。

### 覆盖率模式对比

| 模式 | 含义 | 适用场景 |
|------|------|---------|
| `-covermode=set` | 每行是否被执行过（是/否） | 日常开发，默认值 |
| `-covermode=count` | 每行被执行了多少次 | 分析热点路径 |
| `-covermode=atomic` | 同 count，但协程安全 | CI 流水线推荐 |

---

## 第 8 章：Makefile 自动化

### 为什么在 Go 项目中使用 Makefile

Go 工具链本身很强大，但 Makefile 提供了额外的价值：

- **标准化操作**：团队成员运行相同的命令，避免每个人记住不同的参数
- **减少记忆负担**：`make test` 比 `go test -v -count=1 -race ./...` 好记得多
- **编排复杂流程**：`make all` 一次性运行 vet → test → cover → bench → build
- **自文档化**：Makefile 本身就是一个可执行的操作清单

### 本项目的 Makefile

```makefile
.PHONY: test cover htmlcover bench vet lint build test-fast clean all

test:        # 运行所有测试（详细输出）
test-fast:   # 运行所有测试（简要输出，更快）
cover:       # 运行测试并打印覆盖率摘要
htmlcover:   # 生成 HTML 覆盖率报告
bench:       # 运行所有基准测试（含内存统计）
vet:         # 静态分析
lint:        # 运行 golangci-lint（需要安装）
build:       # 验证所有包可以编译
clean:       # 清理生成的文件
all:         # 完整流水线：vet → test → cover → bench → build
```

### 日常使用场景

```bash
make test       # 最常用 —— 每次改完代码跑一次
make bench      # 性能优化时使用
make all        # 提交代码前跑一次，确保一切正常
make htmlcover  # 在浏览器中查看哪些代码没被测试到
make clean      # 清理生成的临时文件
```

---

## 第 9 章：CI/CD 流水线

### GitHub Actions 工作流

本项目包含一个完整的 CI 流水线（`.github/workflows/ci.yml`），在每次 push 和 PR 时自动触发：

```yaml
name: CI

on:
  push:
    branches: [main, master]
  pull_request:
    branches: [main, master]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4       # 拉取代码
      - uses: actions/setup-go@v5       # 安装 Go
        with:
          go-version: '1.24'
      - run: go vet ./...               # 静态分析
      - run: go test -v ./...           # 运行测试
      - run: |                          # 覆盖率报告
          go test -coverprofile=coverage.out -covermode=atomic ./...
          go tool cover -func=coverage.out
      - run: go test -bench=. -benchmem ./...  # 基准测试
      - run: go build ./...             # 验证编译
```

### 各步骤的作用

| 步骤 | 目的 | 失败意味着 |
|------|------|-----------|
| `go vet` | 捕获常见错误（如不可达代码、错误的 Printf 格式） | 代码有静态问题 |
| `go test -v` | 运行所有单元测试和示例测试 | 功能回归 |
| `go test -cover` | 生成覆盖率报告 | 不阻断，仅报告 |
| `go test -bench` | 运行基准测试 | 不阻断，但可发现极端性能退化 |
| `go build` | 验证所有包能编译 | 代码无法构建 |

### README 状态徽章

```markdown
[![CI](https://github.com/用户名/仓库名/actions/workflows/ci.yml/badge.svg)](https://github.com/用户名/仓库名/actions)
```

徽章会根据最新的 CI 运行状态自动变为绿色（通过）或红色（失败）。

### 进阶：多版本矩阵测试

如果需要在多个 Go 版本上同时测试：

```yaml
strategy:
  matrix:
    go-version: ['1.22', '1.23', '1.24']
```

GitHub Actions 会为每个版本并行创建一个运行实例。

---

## 第 10 章：Pre-commit 钩子

### 什么是 Git Hooks

Git hooks 是在特定 Git 事件发生时自动运行的脚本。**pre-commit** 钩子在 `git commit` 之前运行——如果钩子以非零退出码结束，commit 会被阻止。

这个机制确保了：**有问题的代码永远不会进入仓库历史**。

### 本项目的 Pre-commit 钩子

`.githooks/pre-commit`：

```sh
#!/bin/sh
echo "=== 运行 go vet ==="
go vet ./...
if [ $? -ne 0 ]; then
    echo "失败: go vet 发现问题，请修复后再提交。"
    exit 1
fi

echo "=== 运行测试 ==="
go test ./...
if [ $? -ne 0 ]; then
    echo "失败: 测试必须全部通过才能提交。"
    exit 1
fi

echo "=== 全部检查通过 ==="
```

### 启用钩子

```bash
# 方法一：运行设置脚本
bash scripts/setup-hooks.sh

# 方法二：手动配置
git config core.hooksPath .githooks
```

### 跳过钩子（紧急情况）

```bash
git commit --no-verify -m "紧急修复"
```

### 钩子设计原则

| 原则 | 说明 |
|------|------|
| **保持轻量** | 钩子中只跑 vet + test，不跑 bench（太慢） |
| **快速反馈** | 在提交前发现问题，比等到 CI 失败再修快得多 |
| **可跳过** | 提供 `--no-verify` 作为紧急逃生通道 |
| **提交到仓库** | 钩子脚本放在 `.githooks/` 目录下，跟随版本控制 |
| **团队共享** | 每个开发者克隆后运行一次 `setup-hooks.sh` 即可启用 |

---

## 附录 A：命令速查

### `go test` 常用标志一览

```bash
go test ./...                        # 运行所有测试
go test -v ./...                     # 详细输出模式
go test -run TestName ./...          # 运行匹配的测试
go test -run "TestName/子测试" ./...  # 运行特定子测试
go test -count=1 ./...               # 绕过缓存（强制重新执行）
go test -cover ./...                 # 显示覆盖率百分比
go test -coverprofile=c.out ./...    # 生成覆盖率文件
go test -bench=. ./...               # 运行所有基准测试
go test -bench=. -benchmem ./...     # 基准测试 + 内存分配统计
go test -bench=. -count=5 ./...      # 多次运行基准测试（更稳定）
go test -timeout 30s ./...           # 设置超时时间
go test -short ./...                 # 跳过长时间运行的测试
go test -race ./...                  # 启用竞态检测
```

### 本项目 Makefile 命令

```bash
make test        # 运行所有测试（详细输出）
make test-fast   # 运行所有测试（简要输出）
make bench       # 运行所有基准测试
make cover       # 覆盖率摘要
make htmlcover   # 生成 HTML 覆盖率报告
make vet         # 静态分析
make lint        # golangci-lint 检查
make build       # 验证编译
make all         # 完整流水线
make clean       # 清理生成文件
```

### 测试函数命名约定速查

| 前缀 | 函数签名 | 用途 |
|------|---------|------|
| `Test` | `func TestXxx(t *testing.T)` | 单元测试 |
| `Benchmark` | `func BenchmarkXxx(b *testing.B)` | 基准测试 |
| `Example` | `func ExampleXxx()` | 示例测试（需 `// Output:` 注释） |
| `ExampleXxx_suffix` | `func ExampleXxx_suffix()` | 函数 Xxx 的附加示例 |
| `ExampleT_M` | `func ExampleT_M()` | 类型 T 的方法 M 的示例 |

### 常见问题排查

| 问题 | 解决方案 |
|------|---------|
| 测试被缓存了，修改后仍显示 PASS | 使用 `go test -count=1` |
| `go test` 报 "no test files" | 检查文件名是否以 `_test.go` 结尾 |
| Example 测试输出不匹配 | 检查空格和换行是否完全一致 |
| 基准测试运行时间太长 | 使用 `-benchtime=1x` 限制运行时间 |
| `reflect.DeepEqual` 总是 false | 检查类型是否匹配（如 `[]int` vs `[]int64`） |

---

## 附录 B：项目测试清单

### 测试统计

| 包 | 源文件 | 单元测试 | 基准测试 | 示例测试 | 覆盖率 |
|-----|-------|---------|---------|---------|--------|
| `mathutil` | 4 | 4 | 4 | 6 | 100% |
| `sort` | 3 | 4 | 3 | 3 | 100% |
| `search` | 2 | 2 | 2 | 2 | 100% |
| `ds` | 3 | 11 | 4 | 6 | 100% |
| `tttys` | 1 | 2 | — | 2 | 100% |
| `utils` | 1 | 1 | — | 1 | 100% |
| **合计** | **14** | **24** | **13** | **20** | **100%** |

### 各包测试覆盖详情

**mathutil 包**

| 测试文件 | 包含的测试 |
|---------|-----------|
| `fibonacci_test.go` | 负数、F(0)、F(1)、F(2)、F(10)、F(20) |
| `gcd_test.go` | 双零、单零、正常、互质、相同数、负数 |
| `prime_test.go` | 负数、0、1、2、3、4、大质数、大合数 |
| `factorial_test.go` | 负数、0!、1!、5!、10! |
| `benchmark_test.go` | Fibonacci(n=10/20/50), Factorial(n=10/20), GCD(small/large), IsPrime(small/large/composite) |
| `example_test.go` | ExampleFibonacci, ExampleFibonacci_negative, ExampleFactorial, ExampleFactorial_zero, ExampleGCD, ExampleIsPrime |

**sort 包**

| 测试文件 | 包含的测试 |
|---------|-----------|
| `bubblesort_test.go` | 空切片、单元素、已排序、逆序、重复、负数、全等值 |
| `quicksort_test.go` | 同上 7 种场景 |
| `mergesort_test.go` | 同上 7 种场景 + 不修改原切片验证 |
| `benchmark_test.go` | BubbleSort/QuickSort/MergeSort 在 10/100/1000/10000 规模下的性能 |
| `example_test.go` | ExampleBubbleSort, ExampleQuickSort, ExampleMergeSort |

**search 包**

| 测试文件 | 包含的测试 |
|---------|-----------|
| `binarysearch_test.go` | 空切片、单元素找到/未找到、首/尾/中间找到、三种未找到、重复元素 |
| `linearsearch_test.go` | 空切片、单元素、首次出现、未找到、负数、未排序也能用 |
| `benchmark_test.go` | BinarySearch(首/尾/未找到) vs LinearSearch(首/尾/未找到)，直观对比 O(log n) vs O(n) |
| `example_test.go` | ExampleBinarySearch(多行输出), ExampleLinearSearch |

**ds 包**

| 测试文件 | 包含的测试 |
|---------|-----------|
| `stack_test.go` | 新建、Push、Pop、Pop 空栈报错、Peek 空栈报错、LIFO 顺序 |
| `queue_test.go` | 新建、Enqueue、Dequeue、Dequeue 空队列报错、Peek 空队列报错、FIFO 顺序 |
| `linkedlist_test.go` | 新建、头插、尾插、删头、删中间、删尾、删不存在、删空列表、搜索、搜索空列表、Size |
| `benchmark_test.go` | StackPush, StackPop, QueueEnqueue, QueueDequeue |
| `example_test.go` | 6 个示例覆盖 Stack/Queue 的核心操作 |
