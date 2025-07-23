
# 🛡️ Post-Quantum Signer (Go)

A command-line tool for generating post-quantum keypairs and signing messages, with future plans for Ethereum integration.

## 🚀 Features

- 🔐 Generate keypairs using NIST-approved PQC algorithms (e.g., Dilithium2)
- ✍️ Sign messages using PQ-safe digital signatures
- ✅ Verify signatures against original messages
- 📂 Store keys in a secure directory
- 📦 Written in Go using `liboqs-go` and `cobra`

## 📂 Project Structure

```
pq-signer/
├── main.go                # CLI commands (genkey, sign, verify, etc.)
├── pqcrypto/              # Post-Quantum Crypto implementation
└── keys/                  # Directory for storing generated keys
```

## 🧱 Requirements

- Go >= 1.21
- C compiler (gcc/clang)
- `liboqs` (post-quantum crypto C library)

## ⚙️ Installation

### Step 1: Install dependencies

```bash
sudo apt update
sudo apt install cmake gcc ninja-build pkg-config libssl-dev
```

### Step 2: Install `liboqs-go`

```bash
cd ~
git clone https://github.com/open-quantum-safe/liboqs-go.git
cd liboqs-go
```
  
### Step 3: Update `liboqs-go` to use `liboqs`

```bash
cd ~/liboqs-go/oqs
nano oqs/oqs.go
```

Open `oqs.go` file and change this directive:

```
#cgo pkg-config: liboqs-go
```

to:

```
#cgo pkg-config: liboqs
```

## 🛠️ Clone and Build

Make sure you have Go installed and configured. Then clone the repository:

```bash
git clone https://github.com/freitasgouvea/pq-signer.git
cd pq-signer
go mod tidy
```

## 🧪 Usage

### 🔐 `genkey`

Create folder `keys/` if it doesn't exist:

```bash
mkdir -p keys
``` 
Generate a new keypair:

```bash
go run main.go genkey
```

### ✍️ `sign`

Sign a message using the generated keypair. The message can be any string, e.g., "Hello PQ Ethereum".

```bash
go run main.go sign "Hello PQ Signer"
```

### ✅ `verify`

Verify a signature against the original message. You will need to provide the signature output from the `sign` command.

```bash
go run main.go verify "Hello PQ Signer" <signature>
```

## 📌 Roadmap

* [x] Basic key pair generation
* [x] Message signing and verification
* [ ] Export/import keys from file (with encryption)

## 📚 References

* [Open Quantum Safe (liboqs)](https://github.com/open-quantum-safe/liboqs)
* [liboqs-go](https://github.com/open-quantum-safe/liboqs-go)
* [NIST PQC Standardization](https://csrc.nist.gov/Projects/post-quantum-cryptography)

## 🛡️ Disclaimer

This project is a **research prototype**. Not intended for production use or to manage real funds. Always review crypto libraries and security assumptions before deployment.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
