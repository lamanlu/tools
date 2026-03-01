# tools

一个基于 Cobra 的命令行工具集，用于生成 root/work 密钥并对字符串进行加密。

## 快速开始

```bash
make build
./tools --help
```

或直接构建：

```bash
go build -mod=readonly -trimpath -buildvcs=false -ldflags="-s -w" -o ./tools ./main.go
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
  key-gen     Generate root/work/random key files

Flags:
  -c, --check   Check cmd run conditions
  -h, --help    help for tools

Use "tools [command] --help" for more information about a command.
```

### tools key-gen --help 输出

```text
Generate root keys (rootKey/root_part_*.key + rootKey/root.salt), a work key encrypted by the root key, or a random key file.

Usage:
  tools key-gen [flags]

Flags:
  -d, --dir string    Key base directory. Will create rootKey/workKey under it.
  -f, --force         Force Create RootKey, Ignore Exist key.
  -h, --help          help for key-gen
  -n, --name string   Work Key File Name. eg: work.key (default "work.key")
  -t, --type string   Key Type: root, work, random.
```

### tools encrypt --help 输出

```text
Encrypt a plaintext string using the specified work key file and output Base64.

Usage:
  tools encrypt [flags]

Flags:
  -d, --key-dir string    Key base directory containing rootKey/workKey.
  -h, --help              help for encrypt
  -k, --work-key string   Work key file name, using for encrypt input string.
```

### tools decrypt --help 输出

```text
Decrypt a Base64-encoded ciphertext using the specified work key file and output plaintext.

Usage:
  tools decrypt [flags]

Flags:
  -d, --key-dir string    Key base directory containing rootKey/workKey.
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

指定自定义密钥目录（会在目录下创建 `rootKey/` 与 `workKey/`）：

```bash
./tools key-gen --type root --dir /tmp/keys
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

指定自定义密钥目录：

```bash
./tools key-gen --type work --name work.key --dir /tmp/keys
```

### 3) 生成随机密钥文件（random key file）
随机密钥文件会生成在 `workKey/` 目录下，文件内容是一个 Base64 编码的随机密钥（长度 32 字节）。

```bash
./tools key-gen --type random --name random.key
```

强制重建指定随机密钥文件：

```bash
./tools key-gen --type random --name random.key --force
```

指定自定义密钥目录：

```bash
./tools key-gen --type random --name random.key --dir /tmp/keys
```

### 4) 加密输入字符串
使用工作密钥加密输入字符串，输出为 Base64（不带换行）。

```bash
./tools encrypt --work-key work.key "your-string"
```

使用自定义密钥目录：

```bash
./tools encrypt --work-key work.key --key-dir /tmp/keys "your-string"
```

### 5) 解密 Base64 字符串
使用工作密钥解密 Base64 密文，输出明文（不带换行）。

```bash
./tools decrypt --work-key work.key "base64-cipher-text"
```

使用自定义密钥目录：

```bash
./tools decrypt --work-key work.key --key-dir /tmp/keys "base64-cipher-text"
```

## 完整加解密示例

```bash
# 1. 生成 root key 与 salt（在自定义目录下的 rootKey/）
./tools key-gen --type root --dir /tmp/keys

# 2. 生成工作密钥（在自定义目录下的 workKey/）
./tools key-gen --type work --name work.key --dir /tmp/keys

# 3. 加密字符串（输出 Base64）
cipher=$(./tools encrypt --work-key work.key --key-dir /tmp/keys "hello-world")
echo "$cipher"

# 4. 解密 Base64（输出明文）
plain=$(./tools decrypt --work-key work.key --key-dir /tmp/keys "$cipher")
echo "$plain"
```

## 自验证脚本

运行全流程自检脚本（构建、生成密钥、加解密并校验一致性）：

```bash
./tools_self_check.sh
```
预期执行结果如下：
```
[1/5] build...
Building...
Done
[2/5] generate root key...
Create Root Key Salt Done
Create Root Key Done
[3/5] generate work key...
Create Work Key: work.key Done
[4/5] encrypt/decrypt...
[5/5] ok
```

## 目录结构（核心部分）

- `main.go`：入口与命令注册
- `keys/`：key-gen 命令实现（参数解析与调用）
- `encrypt/`：字符串加密命令
- `common/`：加密、密钥生成/读取与路径等通用逻辑

## 注意事项

- root key 文件为 `root_part_1.key`、`root_part_2.key` 与 `root.salt`，默认位于 `rootKey/` 目录。
- 工作密钥保存在 `workKey/` 目录内，请妥善保护密钥文件与目录权限。
