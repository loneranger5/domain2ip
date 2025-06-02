
# 🕵️‍♂️ Go Domain Resolver

Concurrent domain-to-IP resolver written in Go using `net.LookupIP` and `github.com/sourcegraph/conc` for parallel execution. It reads a list of domain names from a file and outputs resolved IPv4 addresses to both the console and a specified output file.
inspired from 

---

## 🚀 Features

- Concurrent domain resolution using goroutines
- Outputs only valid IPv4 addresses
- Saves results to a file
- Clean and minimal dependencies

---

## 📦 Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/loneranger5/domain2ip.git)
   cd domain2ip
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Build the binary**:
   ```bash
   go build -o domain2ip
   ```

---

## 🛠 Usage

```bash
./domain2ip -domain domains.txt -output result.txt
```

### 📄 Arguments

| Flag         | Description                         | Required | Default     |
|--------------|-------------------------------------|----------|-------------|
| `-domain`    | Path to file containing domains     | ✅       | _None_      |
| `-output`    | Output filename for IP results      | ❌       | `output.txt`|

---

## 📁 Input File Format

The domain file (`domains.txt`) should contain one domain per line, for example:

```
google.com
example.org
openai.com
```

---

## 📝 Sample Output

 `output.txt`:

```
google.com -> 142.250.182.142
example.org -> 93.184.216.34
openai.com -> 104.18.21.213
```

---

## ⚠️ Notes

- Only IPv4 addresses are captured.
- Invalid or unreachable domains are silently ignored.
- The program panics if the domain file is empty or not found.

