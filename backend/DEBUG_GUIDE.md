# 后端调试指南

## 环境设置

### 1. 获取 API 密钥

#### Google Gemini API Key

1. 访问 [Google AI Studio](https://aistudio.google.com)
2. 点击 "Get API key"
3. 创建新的 API key
4. 复制 API key 到 `.env` 文件中的 `GOOGLE_API_KEY`

#### OpenAI API Key (可选，用于 GPT-4 提示词增强)

1. 访问 [OpenAI Platform](https://platform.openai.com)
2. 点击你的账户 -> API keys
3. 创建新的 API key
4. 复制 API key 到 `.env` 文件中的 `OPENAI_API_KEY`

### 2. 配置环境变量

编辑 `.env` 文件，填入你的 API 密钥：

```bash
GOOGLE_API_KEY=your_actual_google_api_key
OPENAI_API_KEY=your_actual_openai_api_key
PORT=8080
ENV=development
```

## 运行服务器

### 方式 1: 直接运行编译后的二进制

```bash
export $(cat .env | grep -v '#' | xargs)
./img-gacha
```

### 方式 2: 使用启动脚本

```bash
chmod +x run.sh
./run.sh
```

### 方式 3: 使用 Go 命令直接运行

```bash
go run main.go
```

## 测试 API

### 使用测试脚本

```bash
chmod +x test_api.sh
./test_api.sh
```

### 手动测试端点

#### 1. 健康检查

```bash
curl -X GET http://localhost:8080/health
```

响应示例：

```json
{
	"status": "ok"
}
```

#### 2. 获取可用模型

```bash
curl -X GET http://localhost:8080/api/models
```

响应示例：

```json
{
	"models": [
		{
			"id": "gemini-2.0-flash",
			"name": "Gemini 2.0 Flash",
			"description": "Fast and powerful image generation with Gemini",
			"provider": "Google Gemini"
		},
		{
			"id": "gpt-4-image",
			"name": "GPT-4 Enhanced Generation",
			"description": "Uses GPT-4 to enhance prompts, then generates with Gemini",
			"provider": "OpenAI + Google Gemini"
		}
	]
}
```

#### 3. 生成图像 (Gemini 直接生成)

```bash
curl -X POST http://localhost:8080/api/generate \
  -H "Content-Type: application/json" \
  -d '{
    "prompt": "A nano banana dish in a fancy restaurant with a Gemini theme",
    "model": "gemini-2.0-flash",
    "count": 1,
    "steps": 50,
    "cfg_scale": 7.5
  }'
```

#### 4. 生成图像 (GPT-4 增强提示词)

```bash
curl -X POST http://localhost:8080/api/generate \
  -H "Content-Type: application/json" \
  -d '{
    "prompt": "A nano banana dish",
    "model": "gpt-4-image",
    "count": 1
  }'
```

## 常见问题排查

### 1. 错误: "GOOGLE_API_KEY not set"

-   确保 `.env` 文件存在并包含有效的 `GOOGLE_API_KEY`
-   检查环境变量是否正确加载

### 2. 错误: "failed to create Gemini client"

-   验证 API key 是否有效
-   检查网络连接

### 3. 错误: "no images generated"

-   Gemini API 可能没有返回图像数据
-   尝试使用不同的提示词
-   检查 API 的使用配额

### 4. 服务器无法启动

-   确保端口 8080 没有被占用
-   运行 `lsof -i :8080` 检查端口占用情况
-   如果被占用，修改 `.env` 中的 `PORT` 值

## 日志调试

如果遇到问题，检查服务器输出日志中的错误信息。常见的错误信息包括：

-   API key 相关错误
-   网络连接错误
-   JSON 解析错误
-   模型不存在错误

## 下一步

1. 确认 API 端点能正常工作
2. 连接前端应用
3. 测试端到端的图像生成流程
