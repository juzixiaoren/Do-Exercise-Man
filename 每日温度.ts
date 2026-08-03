function dailyTemperatures(temperatures: number[]): number[] {
  const stack: number[] = [];
  const res: number[] = new Array(temperatures.length).fill(0);
  for (let i = 0; i < temperatures.length; i++) {
    while (
      stack.length &&
      temperatures[i] > temperatures[stack[stack.length - 1]]
    ) {
      const idx = stack.pop() as number;
      res[idx] = i - idx;
    }
    stack.push(i);
  }
  return res;
}
//题目描述
// 请根据每日 气温 列表，重新生成一个列表。对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。如果气温在这之后都不会升高，请在该位置用 0 来代替。
// 1. 用栈记录下标
// 2. 如果当前温度大于栈顶温度，则弹出栈顶，计算距离
// 3. 否则入栈
