import os, subprocess

cpp_dir = r"e:\做题家\Do-Exercise-Man\C++"
files = [f for f in os.listdir(cpp_dir) if f.endswith(".cpp")]

print("=== .cpp 文件内容编码检测（含中文注释的才需要转）===")
for f in sorted(files):
    p = os.path.join(cpp_dir, f)
    raw = open(p, "rb").read()
    # 检测是否含非ASCII
    non_ascii = [b for b in raw if b > 127]
    if not non_ascii:
        print(f"{f:30} -> 纯ASCII，无需处理")
        continue
    # 尝试 UTF-8 解码
    try:
        raw.decode("utf-8")
        is_utf8 = True
    except Exception:
        is_utf8 = False
    # 尝试 GBK 解码
    try:
        raw.decode("gbk")
        is_gbk = True
    except Exception:
        is_gbk = False
    tag = "UTF-8(正常)" if is_utf8 else ("GBK(乱码需转)" if is_gbk else "其他")
    print(f"{f:30} -> {tag}  (utf8={is_utf8}, gbk={is_gbk}, bytes={len(raw)})")
