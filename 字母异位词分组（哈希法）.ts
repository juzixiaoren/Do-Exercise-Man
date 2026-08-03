function groupAnagrams(strs: string[]): string[][] {
  let map: Map<string, string[]> = new Map();
  for (let i = 0; i < strs.length; i++) {
    let word: string = strs[i];
    word = word.split("").sort().join("");
    if (map.has(word)) {
      map.get(word)!.push(strs[i]);
    } else {
      map.set(word, [strs[i]]);
    }
  }
  return Array.from(map.values());
}
//哈希表，把排序后的字符串作为key，原字符串作为value，最后返回map.values()
