# Go DDD 项目骨架

[![GitHub stars](https://img.shields.io/github/stars/ppd324/go-ddd-layout)](https://github.com/ppd324/go-ddd-layout/stargazers)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

#### 基于领域驱动设计（DDD）的 Golang 项目架构模板，适用于构建中大型复杂业务系统。

## 特性

- **领域驱动设计**：清晰分离领域模型、应用服务与基础设施  
- **编译时依赖注入**：使用 Wire 自动生成初始化代码，无需反射或运行时框架
- **配置管理**：集中管理多环境（dev、staging、prod），支持文件和环境变量混合配置
- **测试支持**：示例单元测试与集成测试，领域层可独立验证业务逻辑  

## 技术栈

- Go (>= 1.21)  
- 依赖注入：Google Wire  
- Web 框架：Gin  
- 配置管理：Viper / Env  

## 目录结构
```
├─app
│  ├─application       # 应用层
│  │  ├─dto            # dto对象
│  │  └─services       # 应用服务
│  ├─domain            # 领域层
│  │  ├─aggregate      # 聚合根
│  │  ├─entity         # 实体对象
│  │  ├─repos          # 依赖repo接口
│  │  └─service        # 领域服务
│  ├─infra             # 基础设施层
│  │  └─database       # 数据库组件
│  │      ├─gorm       # mysql组件
│  │      │  └─models  # 数据库DO对象
│  │      └─memory     # 内存组件
│  └─interfaces        # 接口层
│      └─http          # http接口
│          ├─handlers  # 处理handler
│          ├─middlewares # 中间件
│          ├─response    #http响应
│          └─router     # http路由
├─cmd                   # 程序入口
├─conf                  # 加载配置
├─configs               # 配置文件
└─pkg                   # 公共包

```

## 快速开始

1. 克隆仓库
   ```bash
   git clone https://github.com/ppd324/go-ddd-layout.git
   cd go-ddd-layout
   ```
2. 安装依赖
 ```bash
    go mod download
    go install github.com/google/wire/cmd/wire@latest
   ```
3. 生成依赖注入代码
  ```bash
     wire ./... 
  ```
4. 启动服务
  ```bash
     go run cmd/main.go
  ```

--- 

### 设计模式参考
- [领域驱动设计模式手册](https://domainlanguage.com/ddd/reference/)
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Google Wire 官方文档](https://github.com/google/wire)

---

> 提示：实际使用时请根据项目需求调整配置项和层级细节，建议结合具体业务场景设计聚合根边界。