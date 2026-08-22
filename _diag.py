import subprocess

tree_hash = "18d96f08b43ee00dea41f681d996194f778f2251"
raw = subprocess.check_output(["git", "cat-file", "tree", tree_hash])

i = 0
entries = []
while i < len(raw):
    sp = raw.index(b" ", i)
    mode = raw[i:sp].decode()
    nul = raw.index(b"\x00", sp)
    name = raw[sp + 1:nul]
    h = raw[nul + 1:nul + 21].hex()
    entries.append((mode, name, h))
    i = nul + 21

print("=== git tree stored names (raw bytes) ===")
for mode, name, h in entries:
    # try decode as utf-8 and gbk
    try:
        u = name.decode("utf-8")
    except Exception:
        u = None
    try:
        g = name.decode("gbk")
    except Exception:
        g = None
    print(f"{h[:8]} | utf8={u!r:30} | gbk={g!r}")
