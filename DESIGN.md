# Memory命令设计文档

## 概述

Memory命令是一个全局的命令行工具，允许用户存储和管理需要记住的重要信息。该工具通过集成Claude AI来自动优化和组织存储的信息，提供智能的记忆管理功能。

## 核心需求

- 全局命令：`memory "信息内容"`
- 自动调用 `claude -p` 命令优化存储内容
- 智能存储到合适的位置和数据结构
- 使用Go语言实现

## 系统架构

### 1. 命令行接口层 (CLI Layer)
- 使用 `cobra` 库构建CLI命令
- 支持多种子命令：
  - `memory add "内容"` - 添加新记忆
  - `memory list` - 列出所有记忆
  - `memory search "关键词"` - 搜索记忆
  - `memory delete <id>` - 删除特定记忆
  - `memory export` - 导出记忆数据

### 2. 数据存储层 (Storage Layer)
#### 存储位置选择策略
```
~/.config/memory/           # Linux/macOS 配置目录
├── memories.json          # 主要记忆数据文件
├── config.yaml           # 配置文件
├── cache/                # 缓存目录
└── backups/              # 备份目录
```

#### 数据结构设计
```go
type Memory struct {
    ID          string    `json:"id"`          // UUID
    OriginalText string   `json:"original"`    // 原始输入
    OptimizedText string  `json:"optimized"`   // AI优化后的内容
    Categories  []string  `json:"categories"`  // 自动分类标签
    Keywords    []string  `json:"keywords"`    // 关键词
    CreatedAt   time.Time `json:"created_at"`  // 创建时间
    UpdatedAt   time.Time `json:"updated_at"`  // 更新时间
    Importance  int       `json:"importance"`  // 重要性评分 1-10
    AccessCount int       `json:"access_count"`// 访问次数
}

type MemoryStore struct {
    Memories []Memory `json:"memories"`
    Metadata Metadata `json:"metadata"`
}

type Metadata struct {
    Version     string    `json:"version"`
    TotalCount  int       `json:"total_count"`
    LastBackup  time.Time `json:"last_backup"`
    LastCleanup time.Time `json:"last_cleanup"`
}
```

### 3. AI集成层 (AI Integration Layer)
#### Claude API集成
```go
type ClaudeService struct {
    APIKey    string
    BaseURL   string
    Model     string
}

type OptimizeRequest struct {
    Text   string `json:"text"`
    Prompt string `json:"prompt"`
}

type OptimizeResponse struct {
    OptimizedText string   `json:"optimized_text"`
    Categories    []string `json:"categories"`
    Keywords      []string `json:"keywords"`
    Importance    int      `json:"importance"`
}
```

#### 优化策略
- 使用专门的提示词(prompt)来优化记忆内容
- 自动提取关键词和分类标签
- 评估信息的重要性级别
- 去重和合并相似记忆

### 4. 索引和搜索层 (Indexing & Search Layer)
#### 搜索功能
- 全文搜索
- 标签搜索
- 关键词匹配
- 模糊搜索
- 时间范围搜索

#### 索引策略
```go
type SearchIndex struct {
    FullTextIndex map[string][]string // 全文索引
    KeywordIndex  map[string][]string // 关键词索引
    CategoryIndex map[string][]string // 分类索引
    DateIndex     map[string][]string // 时间索引
}
```

## 技术栈选择

### 核心依赖库
- **CLI框架**: `github.com/spf13/cobra`
- **配置管理**: `github.com/spf13/viper`
- **JSON处理**: 标准库 `encoding/json`
- **HTTP客户端**: 标准库 `net/http`
- **文件系统**: 标准库 `os` 和 `path/filepath`
- **UUID生成**: `github.com/google/uuid`
- **日志记录**: `github.com/sirupsen/logrus`

### 可选增强库
- **全文搜索**: `github.com/blevesearch/bleve`
- **配置加密**: `github.com/99designs/keyring`
- **进度条**: `github.com/schollz/progressbar`

## 配置管理

### 配置文件结构 (config.yaml)
```yaml
claude:
  api_key: "${CLAUDE_API_KEY}"
  base_url: "https://api.anthropic.com"
  model: "claude-3-sonnet-20240229"
  timeout: 30s

storage:
  data_dir: "~/.config/memory"
  backup_enabled: true
  backup_interval: "24h"
  max_backups: 7

search:
  enable_full_text: true
  enable_fuzzy: true
  max_results: 50

performance:
  cache_enabled: true
  cache_ttl: "1h"
  batch_size: 10
```

## 安全考虑

### 1. 敏感信息保护
- API密钥加密存储
- 支持环境变量配置
- 本地文件权限控制 (0600)

### 2. 数据备份和恢复
- 自动定期备份
- 版本化备份
- 灾难恢复机制

### 3. 输入验证
- 内容长度限制
- 特殊字符过滤
- 注入攻击防护

## 性能优化

### 1. 缓存策略
- 内存缓存频繁访问的记忆
- AI响应缓存机制
- 搜索结果缓存

### 2. 批处理
- 批量AI请求处理
- 批量存储更新
- 异步处理非关键操作

### 3. 索引优化
- 增量索引更新
- 延迟索引构建
- 压缩存储格式

## 错误处理

### 1. 网络错误
- 自动重试机制
- 离线模式支持
- 优雅降级

### 2. 存储错误
- 数据完整性检查
- 自动修复机制
- 回滚功能

### 3. AI服务错误
- 本地回退策略
- 错误分类处理
- 用户友好的错误信息

## 扩展性设计

### 1. 插件系统
- 自定义存储后端
- 第三方AI服务集成
- 自定义搜索算法

### 2. 导入导出
- 支持多种格式 (JSON, CSV, Markdown)
- 与其他工具集成
- 数据迁移工具

### 3. API接口
- RESTful API
- GraphQL支持
- Webhook集成

## 项目结构

```
memory/
├── cmd/                    # 命令行入口
│   ├── root.go
│   ├── add.go
│   ├── list.go
│   ├── search.go
│   ├── delete.go
│   └── export.go
├── internal/
│   ├── config/            # 配置管理
│   ├── storage/           # 存储层
│   ├── ai/               # AI集成
│   ├── search/           # 搜索功能
│   └── models/           # 数据模型
├── pkg/
│   └── memory/           # 公共API
├── scripts/              # 构建脚本
├── docs/                 # 文档
├── tests/                # 测试
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## 开发计划

### Phase 1: 核心功能 (2-3周)
- [ ] 基础CLI框架搭建
- [ ] 数据存储实现
- [ ] Claude API集成
- [ ] 基本的增删改查功能

### Phase 2: 搜索和索引 (1-2周)
- [ ] 全文搜索实现
- [ ] 索引系统构建
- [ ] 高级搜索功能

### Phase 3: 优化和增强 (1-2周)
- [ ] 性能优化
- [ ] 错误处理完善
- [ ] 用户体验优化

### Phase 4: 高级功能 (2-3周)
- [ ] 数据导入导出
- [ ] 备份恢复系统
- [ ] 插件系统框架

## 测试策略

### 1. 单元测试
- 核心业务逻辑测试
- 数据存储测试
- AI集成测试

### 2. 集成测试
- 端到端功能测试
- API集成测试
- 性能测试

### 3. 用户测试
- 可用性测试
- 压力测试
- 兼容性测试

## 部署和分发

### 1. 构建系统
- 跨平台编译
- 自动化构建流水线
- 版本管理

### 2. 分发方式
- GitHub Releases
- Homebrew (macOS)
- APT/YUM 包 (Linux)
- Chocolatey (Windows)

### 3. 文档和支持
- 用户手册
- API文档
- 故障排除指南