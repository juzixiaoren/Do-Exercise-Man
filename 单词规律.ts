function wordPattern(pattern: string, s: string): boolean {
  const words: string[] = s.split(" ");
  let map: Map<string, string> = new Map();
  let set: Set<string> = new Set();
  if (pattern.length !== words.length) return false;
  for (let i = 0; i < pattern.length; i++) {
    if (map.has(pattern[i])) {
      if (map.get(pattern[i]) !== words[i]) return false;
    } else if (set.has(words[i])) return false;
    else {
      map.set(pattern[i], words[i]);
      set.add(words[i]);
    }
  }
  return true;
}
/*
原题：单词规律
给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。
这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 s 中的每个非空单词之间存在着双向连接的对应规律。
例子：pattern = "abba", s = "dog cat cat dog" → true
* 解题思路：
* 1. 将s按照空格分割成数组
* 2. 遍历pattern，如果map中没有pattern[i]，则将pattern[i]和words[i]存入map中，如果map中已经有pattern[i]，则判断map.get(pattern[i])是否等于words[i]，如果不等于，则返回false
* 3. 遍历map，如果map中有一个key对应的value在set中不存在，则返回false
* 4. 返回true
*/
