# DeepTest

**Go 算法集合 + 自动化单元测试教程**

[![CI](https://github.com/chengduFa/DeepTest/actions/workflows/ci.yml/badge.svg)](https://github.com/chengduFa/DeepTest/actions/workflows/ci.yml)
![Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)

一个实用的 Go 项目，演示经典算法实现和完整的自动化单元测试体系——包括表驱动测试、子测试、基准测试、示例测试、CI/CD 流水线和 Pre-commit 钩子。

## 快速开始

```bash
git clone https://github.com/Hackercds/Unit-Test-for-Golang-Guide.git
cd DeepTest

# 运行所有测试
make test

# 运行基准测试（观察 O(n²) vs O(n log n) vs O(log n) 的性能差异）
make bench

# 完整检查流水线
make all
```

## 项目结构

| 包 | 内容 | 测试类型 |
|------|---------|----------|
| `mathutil/` | 斐波那契、最大公约数、素数判断、阶乘 | 单元测试、基准测试、示例测试 |
| `sort/` | 冒泡排序、快速排序、归并排序 | 单元测试、基准测试、示例测试 |
| `search/` | 二分查找、线性查找 | 单元测试、基准测试、示例测试 |
| `ds/` | 栈、队列、单向链表 | 单元测试、基准测试、示例测试 |
| `tttys/` | 基础函数示例 | 单元测试、示例测试 |
| `utils/` | 配置信息 | 单元测试、示例测试 |

## 测试教程

详见 [TESTING_STRATEGY.md](TESTING_STRATEGY.md)，完整的中文教程共 10 章，涵盖：

- 表驱动测试（Table-Driven Tests）
- 子测试（`t.Run`）
- 黑盒测试 vs 白盒测试
- 基准测试（Benchmark Tests）
- 示例测试（Example Tests）
- 测试覆盖率
- Makefile 自动化
- GitHub Actions CI/CD 流水线
- Pre-commit 钩子

## Makefile 命令

| 命令 | 说明 |
|------|------|
| `make test` | 运行所有测试（详细输出） |
| `make test-fast` | 运行所有测试（简要输出） |
| `make cover` | 运行测试并打印覆盖率摘要 |
| `make htmlcover` | 生成 HTML 覆盖率报告 |
| `make bench` | 运行所有基准测试 |
| `make vet` | 运行静态分析 |
| `make lint` | 运行 golangci-lint |
| `make build` | 验证所有包可以编译 |
| `make clean` | 清理生成的文件 |
| `make all` | 完整流水线（vet + test + cover + bench + build） |

## 测试统计

| 指标 | 数量 |
|------|------|
| 包数量 | 6 |
| 源文件 | 14 |
| 单元测试函数 | 24 |
| 基准测试函数 | 13 |
| 示例测试函数 | 20 |
| 测试覆盖率 | 100% |
