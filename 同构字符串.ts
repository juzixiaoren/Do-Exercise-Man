function isIsomorphic(s: string, t: string): boolean {
  if (s.length !== t.length) return false;
  else {
    let map: Map<string, string> = new Map();
    let set: Set<string> = new Set();
    for (let i = 0; i < s.length; i++) {
      if (map.has(s[i])) {
        if (map.get(s[i]) !== t[i]) return false;
      } else if (set.has(t[i])) {
        return false;
      } else {
        map.set(s[i], t[i]);
        set.add(t[i]);
      }
    }
    return true;
  }
}
