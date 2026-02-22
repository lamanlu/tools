# tools

一个基于 Cobra 的命令行工具集，用于生成 root/work 密钥并对字符串进行加密。

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

### tools --help 输出

```text
CLI utilities for generating root/work keys and encrypting strings using work keys.

Usage:
  tools [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  decrypt     Decrypt a Base64 string with a work key
  encrypt     Encrypt a string with a work key
  help        Help about any command
  key-gen     Generate root or work keys

Flags:
  -c, --check   Check cmd run conditions
  -h, --help    help for tools

Use "tools [command] --help" for more information about a command.
```

### tools key-gen --help 输出

```text
Generate root keys (rootKey/root_part_*.key + rootKey/root.salt) or a work key encrypted by the root key.

Usage:
  tools key-gen [flags]

Flags:
  -f, --force         Force Create RootKey, Ignore Exist key.
  -h, --help          help for key-gen
  -n, --name string   Work Key File Name. eg: work.key (default "work.key")
  -t, --type string   Key Type: root, work.
```

### tools encrypt --help 输出

```text
Encrypt a plaintext string using the specified work key file and output Base64.

Usage:
  tools encrypt [flags]

Flags:
  -h, --help              help for encrypt
  -k, --work-key string   Work key file name, using for encrypt input string.
```

### tools decrypt --help 输出

```text
Decrypt a Base64-encoded ciphertext using the specified work key file and output plaintext.

Usage:
  tools decrypt [flags]

Flags:
  -h, --help              help for decrypt
  -k, --work-key string   Work key file name, using for decrypt input string.
```

### 1) 生成根密钥（root key）
生成根密钥与盐文件，会在 `rootKey/` 目录产生 `root_part_1.key`、`root_part_2.key` 与 `root.salt`。

```bash
./tools key-gen --type root
```

注意：`--type` 为必填参数。

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
使用工作密钥加密输入字符串，输出为 Base64（不带换行）。

```bash
./tools encrypt --work-key work.key "your-string"
```

### 4) 解密 Base64 字符串
使用工作密钥解密 Base64 密文，输出明文（不带换行）。

```bash
./tools decrypt --work-key work.key "base64-cipher-text"
```

## 完整加解密示例

```bash
# 1. 生成 root key 与 salt（在 rootKey/ 目录）
./tools key-gen --type root

# 2. 生成工作密钥（在 workKey/ 目录）
./tools key-gen --type work --name work.key

# 3. 加密字符串（输出 Base64）
cipher=$(./tools encrypt --work-key work.key "hello-world")
echo "$cipher"

# 4. 解密 Base64（输出明文）
plain=$(./tools decrypt --work-key work.key "$cipher")
echo "$plain"
```

## 目录结构（核心部分）

- `main.go`：入口与命令注册
- `keys/`：key-gen 命令实现（参数解析与调用）
- `encrypt/`：字符串加密命令
- `common/`：加密、密钥生成/读取与路径等通用逻辑

## 注意事项

- root key 文件为 `root_part_1.key`、`root_part_2.key` 与 `root.salt`，默认位于 `rootKey/` 目录。
- 工作密钥保存在 `workKey/` 目录内，请妥善保护密钥文件与目录权限。
