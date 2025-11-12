# 🎲 Image Gacha Generator

AI 抽卡式图片生成 Web 应用 - 用户输入 prompt，系统随机生成多张不同风格的图片。真正的"抽卡"体验，每次生成都有惊喜！

## 🌟 核心功能

-   **智能图片生成**：用户输入 prompt，系统随机生成多张不同风格的图片
-   **多模型支持**：集成 Stable Diffusion v1.5/v2.1 和 DALL-E 3 等多个 AI 模型
-   **真实抽卡体验**：相同 prompt + 相同模型 = 不同图片结果，每次生成都有惊喜
-   **参数控制**：支持调整 steps、cfg_scale、negative_prompt 等参数
-   **批量生成**：一次可生成 1-10 张图片
-   **实时反馈**：生成过程中显示实时状态

## 🏗️ 技术架构

### 后端 (Go + Gin)

-   **HTTP 服务器**：使用 Gin 框架处理 API 请求
-   **多模型路由**：根据用户选择路由到不同的 AI 模型 API
-   **随机种子控制**：通过随机种子确保相同 prompt 生成不同结果
-   **并发处理**：支持高并发请求
-   **错误重试机制**：实现指数退避策略的重试机制

### 前端 (React + TypeScript)

-   **现代 UI**：基于 React 18 和 TypeScript 的类型安全界面
-   **Vite 构建**：极速开发和构建体验
-   **组件化设计**：清晰的组件结构和代码组织
-   **响应式设计**：完美支持各种屏幕尺寸

## 📦 项目结构

```
img-gacha/
├── backend/
│   ├── main.go                 # 服务器主程序
│   ├── go.mod                  # Go依赖管理
│   ├── handlers/               # HTTP请求处理器
│   ├── services/               # 业务逻辑服务
│   │   └── generator.go        # 图片生成服务
│   ├── models/                 # 数据模型
│   │   └── image.go            # 图片相关数据模型
│   ├── utils/                  # 工具函数
│   │   └── random.go           # 随机数生成工具
│   └── .env.example            # 环境变量示例
│
├── frontend/
│   ├── src/
│   │   ├── components/         # React组件
│   │   │   ├── PromptInput.tsx # Prompt输入组件
│   │   │   └── ImageGallery.tsx # 图片画廊组件
│   │   ├── services/           # 业务逻辑服务
│   │   │   └── api.ts          # API调用服务
│   │   ├── types/              # TypeScript类型定义
│   │   │   └── index.ts        # 类型定义
│   │   ├── App.tsx             # 主应用组件
│   │   ├── main.tsx            # 应用入口
│   │   └── index.css           # 全局样式
│   ├── index.html              # HTML模板
│   ├── package.json            # NPM依赖管理
│   ├── tsconfig.json           # TypeScript配置
│   ├── vite.config.ts          # Vite配置
│   └── .gitignore              # Git忽略配置
│
├── LICENSE                     # MIT许可证
└── README.md                   # 项目文档
```

## 🚀 快速开始

### 前置要求

-   Go 1.21+
-   Node.js 18+
-   npm 9+

### 后端设置

```bash
cd backend

# 下载依赖
go mod download

# 配置环境变量
cp .env.example .env
# 编辑 .env 文件，配置API密钥等

# 运行服务器
go run main.go
```

服务器将在 `http://localhost:8080` 启动

### 前端设置

```bash
cd frontend

# 安装依赖
npm install

# 开发模式运行
npm run dev

# 构建生产版本
npm run build

# 预览生产构建
npm run preview
```

前端将在 `http://localhost:5173` 启动（开发模式）

## 🔑 关键技术点

### 随机化实现

```go
// 通过随机种子控制图片差异
// 相同prompt + 相同模型 + 不同seed = 完全不同的图片

// 基础种子生成
seed := rand.Int63()

// 批量生成时的种子衍生
derivedSeeds := DeriveSeeds(baseSeed, count)
```

### API 设计

**POST /api/generate** - 生成图片

请求体：

```json
{
    "prompt": "a beautiful landscape",
    "model": "stable-diffusion-v2.1",
    "count": 4,
    "seed": 12345,
    "steps": 30,
    "cfg_scale": 7.5,
    "negative_prompt": "ugly, distorted",
    "height": 768,
    "width": 768
}
```

响应：

```json
{
    "task_id": "uuid-string",
    "status": "pending",
    "created_at": "2025-11-12T10:00:00Z"
}
```

**GET /api/status/:taskId** - 查询生成状态

**GET /api/models** - 获取可用模型列表

### 用户体验

1. **实时反馈**：用户可以看到生成进度
2. **失败重试**：自动重试失败的请求
3. **结果缓存**：生成结果可被缓存和复用
4. **批量操作**：支持一次生成多张图片

## 🛠️ 开发指南

### 添加新的 AI 模型

1. 在 `backend/models/image.go` 中更新 `GenerationRequest` 的模型选项
2. 在 `backend/services/generator.go` 中实现模型路由逻辑
3. 创建相应的 API 客户端包

### 扩展前端功能

1. 在 `frontend/src/components` 中创建新组件
2. 在 `frontend/src/services/api.ts` 中添加新的 API 调用
3. 在 `frontend/src/types/index.ts` 中定义类型

## 📝 环境变量配置

参考 `backend/.env.example` 文件，需要配置：

-   `STABLE_DIFFUSION_API_URL` - Stable Diffusion API 地址
-   `STABLE_DIFFUSION_API_KEY` - API 密钥
-   `OPENAI_API_KEY` - OpenAI API 密钥（用于 DALL-E）
-   其他配置项

## 📄 许可证

MIT License - 详见 [LICENSE](LICENSE) 文件

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📧 联系方式

如有问题或建议，欢迎通过 GitHub Issues 联系我们。

---

**Happy Gacha! 🎲✨**
