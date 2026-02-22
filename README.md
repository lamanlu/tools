# tools

一个基于 Cobra 的命令行工具集，目前包含密钥生成与字符串加密功能。

## 快速开始

```bash
make build
./tools --help
```

或直接构建：

```bash
go build -o ./tools ./main.go
```

## 功能与用法

### 1) 生成根密钥（root key）
生成根密钥与盐文件，会在当前目录产生 `root.key` 与 `root.salt`。

```bash
./tools key-gen --type root
```

如需强制覆盖已有文件：

```bash
./tools key-gen --type root --force
```

### 2) 生成工作密钥（work key）
工作密钥会生成在 `workKey/` 目录下，并使用 root key 进行加密存储。

```bash
./tools key-gen --type work --name work.key
```

强制重建指定工作密钥：

```bash
./tools key-gen --type work --name work.key --force
```

### 3) 加密输入字符串
使用工作密钥加密输入字符串，输出为 Base64。

```bash
./tools encrypt --work-key work.key "your-string"
```

## 目录结构（核心部分）

- `main.go`：入口与命令注册
- `keys/`：root/work key 生成与读取逻辑
- `encrypt/`：字符串加密命令
- `common/`：加密与编码等通用工具

## 注意事项

- root key 文件为 `root.key` 与 `root.salt`，默认位于当前目录。
- 工作密钥保存在 `workKey/` 目录内，请妥善保护密钥文件与目录权限。
