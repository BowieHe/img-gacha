# 🚀 快速开始指南

## 项目现状

✅ **已完成**:

-   后端框架搭建 (Go + Gin)
-   前端框架搭建 (React + TypeScript + Vite + Tailwind CSS)
-   pnpm 作为包管理器
-   所有文件使用 `@` 别名导入
-   完整的 Tailwind CSS 集成
-   项目结构和配置

## 📦 所有依赖已安装

通过 pnpm 已安装的包：

```
✅ react, react-dom
✅ vite, @vitejs/plugin-react
✅ typescript, @types/*
✅ tailwindcss, postcss, autoprefixer
✅ axios
✅ eslint, @typescript-eslint/*
```

## 🎨 Tailwind 类名示例参考

### 布局

```tsx
// 全屏布局
<div className="min-h-screen">

// 容器和间距
<div className="max-w-4xl mx-auto px-4 py-8">

// 网格布局
<div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">

// Flex 布局
<div className="flex gap-2 items-center justify-center">
```

### 样式

```tsx
// 背景和颜色
className =
    "bg-white bg-gradient-to-br from-primary to-secondary text-gray-800";

// 圆角和阴影
className = "rounded-lg shadow-lg hover:shadow-xl";

// 边框
className = "border-2 border-gray-300 border-primary";

// 响应式
className = "px-4 md:px-8 lg:px-12";
```

### 交互

```tsx
// 悬停效果
className =
    "hover:bg-primary hover:shadow-xl hover:-translate-y-1 transition-all";

// 禁用状态
className = "disabled:opacity-50 disabled:cursor-not-allowed";

// 焦点状态
className =
    "focus:outline-none focus:border-primary focus:ring-2 focus:ring-primary/20";

// 动画
className = "animate-spin animate-float";
```

## 🔄 常见任务

### 启动开发服务器

```bash
cd frontend
pnpm run dev
```

访问 `http://localhost:5173`

### 构建生产版本

```bash
cd frontend
pnpm run build
```

### 添加新的 npm 包

```bash
cd frontend
pnpm add package-name
```

### 添加开发依赖

```bash
cd frontend
pnpm add -D package-name
```

## 📁 导入模板

```tsx
// ✅ 正确的导入方式
import PromptInput from "@/components/PromptInput";
import { apiService } from "@/services/api";
import type { GenerationRequest } from "@/types";
import "@/index.css";

// ❌ 避免相对路径
// import PromptInput from './components/PromptInput'
// import { apiService } from '../../services/api'
```

## 🎯 后续开发步骤

1. **API 集成**

    - 在 `src/services/api.ts` 中实现后端 API 调用
    - 处理生成图片的异步操作
    - 添加错误处理

2. **扩展组件**

    - 添加图片下载功能
    - 添加历史记录功能
    - 添加分享功能

3. **性能优化**

    - 图片懒加载
    - 请求去抖
    - 缓存策略

4. **功能增强**
    - 用户认证
    - 图片保存/收藏
    - 社区分享

## 📚 文档参考

-   [Tailwind CSS 文档](https://tailwindcss.com/docs)
-   [React 文档](https://react.dev)
-   [TypeScript 文档](https://www.typescriptlang.org/docs/)
-   [Vite 文档](https://vitejs.dev)

## ⚠️ 注意事项

1. 修改 Tailwind 配置后需要重启开发服务器
2. `@tailwind` 指令的 CSS 警告在开发时可以忽略
3. Tailwind 类名的优先级规则遵循 CSS 特异性

## 🆘 故障排除

### 开发服务器无法启动

```bash
# 清除依赖缓存
pnpm store prune

# 重新安装依赖
rm -rf node_modules pnpm-lock.yaml
pnpm install
```

### Tailwind 类名不生效

-   确保类名没有拼写错误
-   检查 `tailwind.config.ts` 的 `content` 配置
-   重启开发服务器

### TypeScript 错误

```bash
# 运行类型检查
pnpm run type-check

# 生成类型定义
pnpm run build
```

---

**祝你开发愉快！🎉**
