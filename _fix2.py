import os

p = r"e:\做题家\Do-Exercise-Man\C++\走迷宫.cpp"
raw = open(p, "rb").read()

# 原始为 GBK + CRLF
text = raw.decode("gbk", errors="replace")

# 写回 UTF-8 + CRLF（保持原换行风格），用 newline='' 保留原 \r\n
with open(p, "w", encoding="utf-8", newline="") as f:
    f.write(text)

# 验证
out = open(p, "rb").read()
print("CRLFCRLF", out.count(b"\r\r\n"), "CRLF", out.count(b"\r\n"), "LF", out.count(b"\n"))
t = out.decode("utf-8")
# 抽样中文
import re
samples = re.findall(r"[一-鿿]{4,}", t)[:5]
print("中文抽样:", samples)
print("size", len(out))
