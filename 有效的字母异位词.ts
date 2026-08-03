function isAnagram(s: string, t: string): boolean {
  let map: Map<string, number> = new Map();
  if (s.length !== t.length) return false;
  for (let i = 0; i < s.length; i++) {
    if (map.has(s[i])) {
      map.set(s[i], map.get(s[i])! + 1);
    } else {
      map.set(s[i], 1);
    }
  }
  for (let i = 0; i < t.length; i++) {
    if (map.has(t[i]) && map.get(t[i])! > 0) {
      map.set(t[i], map.get(t[i])! - 1);
    } else {
      return false;
    }
  }
  return true;
}
