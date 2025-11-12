# 📚 开发指南

## 项目结构详解

### 后端结构 (Go)

```
backend/
├── main.go                    # 服务启动入口
├── go.mod                     # Go模块定义
├── .env.example              # 环境变量示例
├── handlers/                 # HTTP请求处理层
│   ├── generate.go          # 图片生成请求处理
│   ├── status.go            # 状态查询处理
│   └── models.go            # 模型列表处理
├── services/                # 业务逻辑层
│   ├── generator.go         # 图片生成核心服务
│   └── retry.go             # 重试机制实现
├── models/                  # 数据模型
│   └── image.go             # 图片相关的数据结构
└── utils/                   # 工具函数
    └── random.go            # 随机数生成工具
```

#### 核心模块说明

**main.go** - 服务器启动

-   Gin 框架初始化
-   CORS 中间件配置
-   路由注册
-   服务器启动

**services/generator.go** - 核心生成服务

-   `GenerateImages()` - 创建生成任务
-   `GenerateWithSeed()` - 使用种子生成
-   `RetryWithBackoff()` - 指数退避重试

**utils/random.go** - 随机数工具

-   `GenerateRandomSeed()` - 生成单个随机种子
-   `GenerateRandomSeeds()` - 生成多个随机种子
-   `DeriveSeeds()` - 从基种子衍生确定性种子

### 前端结构 (React + TypeScript)

```
frontend/
├── index.html               # HTML入口
├── package.json            # 依赖管理
├── tsconfig.json           # TypeScript配置
├── vite.config.ts          # Vite构建配置
├── src/
│   ├── main.tsx            # React应用入口
│   ├── App.tsx             # 根组件
│   ├── index.css           # 全局样式
│   ├── App.css             # 应用样式
│   ├── components/         # 可复用组件
│   │   ├── PromptInput.tsx # Prompt输入组件
│   │   ├── PromptInput.css # 输入组件样式
│   │   ├── ImageGallery.tsx # 图片画廊组件
│   │   └── ImageGallery.css # 画廊组件样式
│   ├── services/           # 业务逻辑
│   │   └── api.ts          # API客户端
│   └── types/              # TypeScript类型
│       └── index.ts        # 类型定义
└── .gitignore
```

#### 组件说明

**App.tsx** - 根组件

-   管理全局状态（images, loading）
-   处理生成逻辑
-   组织子组件

**PromptInput.tsx** - 输入组件

-   Prompt 文本框
-   模型选择下拉框
-   图片数量调整
-   高级选项（steps、cfg_scale、negative_prompt）
-   生成按钮

**ImageGallery.tsx** - 画廊组件

-   显示生成的图片
-   加载动画
-   空状态提示
-   响应式网格布局

## 🔧 开发流程

### 1. 启动开发环境

```bash
# 终端1 - 后端
cd backend
go run main.go

# 终端2 - 前端
cd frontend
npm install
npm run dev
```

### 2. 添加新的 API 端点

1. 在 `backend/handlers/` 中创建处理函数：

```go
// handlers/new_endpoint.go
package handlers

func HandleNewEndpoint(c *gin.Context) {
    // 实现逻辑
    c.JSON(http.StatusOK, gin.H{"message": "success"})
}
```

2. 在 `backend/main.go` 中注册路由：

```go
api := router.Group("/api")
{
    api.POST("/new-endpoint", handleNewEndpoint)
}
```

3. 在前端 `frontend/src/services/api.ts` 中添加调用：

```typescript
export const apiService = {
    // ... 现有方法
    newEndpoint: async (params): Promise<Response> => {
        const response = await api.post("/new-endpoint", params);
        return response.data;
    },
};
```

### 3. 集成新的 AI 模型

1. 在 `backend/models/image.go` 中添加模型类型

2. 在 `backend/services/generator.go` 中实现模型调用逻辑

3. 在前端 `frontend/src/components/PromptInput.tsx` 中添加模型选项

### 4. 修改样式

前端样式采用 CSS Modules 方式组织：

-   每个组件对应一个 `.css` 文件
-   全局样式在 `src/index.css`
-   应用样式在 `src/App.css`

修改后自动热重载（HMR）

## 📋 常见任务

### 调试后端

使用 `log` 包输出日志：

```go
import "log"

log.Printf("Debug info: %v", value)
```

### 调试前端

在浏览器开发者工具中查看：

-   Console 输出
-   Network 请求
-   React DevTools 组件状态

### 测试 API

使用 curl 或 Postman：

```bash
curl -X POST http://localhost:8080/api/generate \
  -H "Content-Type: application/json" \
  -d '{
    "prompt": "test",
    "model": "stable-diffusion-v2.1",
    "count": 1
  }'
```

## 🚀 部署

### 后端部署

```bash
cd backend
go build -o img-gacha-server
./img-gacha-server
```

### 前端部署

```bash
cd frontend
npm run build
# dist/ 目录包含生产构建文件
```

## 📚 参考资源

-   [Gin Web Framework](https://github.com/gin-gonic/gin)
-   [React Documentation](https://react.dev)
-   [TypeScript Handbook](https://www.typescriptlang.org/docs/)
-   [Vite Documentation](https://vitejs.dev)

## 🤔 常见问题

**Q: 如何修改后端端口？**
A: 在 `backend/main.go` 中修改 `router.Run(":8080")` 的端口号

**Q: 如何修改前端开发服务器端口？**
A: 在 `frontend/vite.config.ts` 中修改 `server.port` 配置

**Q: 如何添加新的环境变量？**
A: 在 `backend/.env` 中添加，然后在代码中使用 `os.Getenv()` 读取

## 📞 支持

遇到问题？请查看：

1. 项目 README
2. 相关代码注释
3. GitHub Issues
