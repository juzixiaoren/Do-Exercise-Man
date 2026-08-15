function groupAnagrams1(strs: string[]): string[][] {
  let ans: string[][] = [];
  for (let i = 0; i < strs.length; i++) {
    let word: string = strs[i];
    word = word.split("").sort().join("");
    let flag: boolean = false;
    for (let j = 0; j < ans.length; j++) {
      if (ans[j][0] === word) {
        ans[j].push(strs[i]);
        flag = true;
        break;
      }
    }
    if (!flag) {
      ans.push([word, strs[i]]);
    }
  }
  for (let i = 0; i < ans.length; i++) {
    ans[i].shift();
  }
  return ans;
}
//用数组代替哈希表
