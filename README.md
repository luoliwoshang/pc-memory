# Memory Manager

一个智能的记忆管理工具，集成Claude AI自动分类和标签功能。

## 功能特性

- 🧠 **智能处理**: 集成Claude AI自动分析和分类记忆内容
- 📁 **自动分类**: 智能识别账号信息、工作任务、个人笔记等
- 🏷️ **自动标签**: 自动生成相关标签便于检索
- 🔍 **强大搜索**: 支持内容、分类、标签全文搜索
- 📊 **统计分析**: 提供记忆统计和分布分析
- 💾 **本地存储**: 数据安全存储在用户目录下

## 安装

### 从源码构建

```bash
git clone <repository>
cd memory-manager
make deps
make build
make install
```

### 直接运行

```bash
go run . add "你要记住的内容"
```

## 使用方法

### 基本命令

#### 添加记忆
```bash
# 添加一条新记忆
memory add "我的GitHub账号密码是abc123"

# 添加工作任务
memory add "明天需要完成项目报告并发送给客户"

# 添加个人笔记
memory add "记住要买牛奶和面包"
```

#### 列出记忆
```bash
# 列出所有记忆
memory list

# 按优先级排序
memory list --sort=priority

# 过滤特定分类
memory list --category="账号信息"

# 限制显示数量
memory list --limit=10
```

#### 搜索记忆
```bash
# 搜索包含特定关键词的记忆
memory search "密码"
memory search "GitHub"
memory search "任务"
```

#### 删除记忆
```bash
# 根据ID删除记忆
memory delete 20240806123456-abc123
```

#### 查看统计
```bash
# 显示记忆统计信息
memory stats
```

## 配置

### Claude集成

确保系统中已安装Claude CLI工具：
```bash
# 检查Claude是否可用
which claude
```

如果Claude不可用，程序会使用基本的分类逻辑。

### 数据存储

记忆数据存储在 `~/.memory-manager/memories.json` 文件中。

## 数据结构

每条记忆包含以下信息：

```json
{
  "id": "20240806123456-abc123",
  "content": "原始内容",
  "processed_by": "claude|basic|none",
  "timestamp": "2024-08-06T12:34:56Z",
  "category": "分类名称",
  "priority": 5,
  "tags": ["标签1", "标签2"],
  "metadata": {
    "summary": "AI生成的摘要"
  }
}
```

## 示例

### 完整使用流程

```bash
# 1. 添加一些记忆
memory add "我的AWS访问密钥: AKIAIOSFODNN7EXAMPLE"
memory add "周五之前完成数据库迁移项目"
memory add "妈妈生日是3月15日，记得买礼物"

# 2. 查看所有记忆
memory list

# 3. 搜索特定内容
memory search "密钥"
memory search "生日"

# 4. 查看统计
memory stats

# 5. 按优先级查看
memory list --sort=priority
```

### 输出示例

```
Memory saved successfully!
ID: 20240806123456-abc123
Content: 我的AWS访问密钥: AKIAIOSFODNN7EXAMPLE
Category: 账号信息
Priority: 5/5
Tags: 账号, 密钥, AWS, 重要
Processed by: claude
AI Summary: AWS访问密钥信息，包含访问密钥ID
```

## 开发

### 项目结构

```
memory-manager/
├── cmd/                 # CLI命令实现
│   ├── root.go
│   ├── add.go
│   ├── list.go
│   ├── search.go
│   ├── delete.go
│   └── stats.go
├── internal/
│   ├── models/          # 数据模型
│   │   └── memory.go
│   ├── storage/         # 存储层
│   │   └── storage.go
│   ├── claude/          # Claude集成
│   │   └── client.go
│   └── memory/          # 业务逻辑
│       └── service.go
├── main.go              # 程序入口
├── go.mod
├── Makefile
└── README.md
```

### 运行测试

```bash
make test
```

### 构建

```bash
make build
```

## 注意事项

1. **安全性**: 敏感信息如密码会被自动标记为高优先级，请妥善保护数据文件
2. **Claude依赖**: Claude AI处理是可选的，没有Claude也能正常使用基本功能
3. **数据备份**: 建议定期备份 `~/.memory-manager/` 目录

## 贡献

欢迎提交Issue和Pull Request！

## 许可证

MIT License
