class ListNode {
  val: number;
  next: ListNode | null;
  constructor(val?: number, next?: ListNode | null) {
    this.val = val === undefined ? 0 : val;
    this.next = next === undefined ? null : next;
  }
}

function isPalindrome(head: ListNode | null): boolean {
  const list: number[] = [];
  let slow = head;
  let fast = head;
  while (fast && fast.next) {
    slow = slow?.next as ListNode;
    fast = fast?.next.next as ListNode;
  }
  let cur = slow;
  let pre = head;
  while (cur) {
    list.push(cur.val);
    cur = cur.next;
  }
  while (list.length) {
    if (list.pop() !== pre?.val) return false;
    pre = pre?.next as ListNode;
  }
  return true;
}
/**
 原题： 回文链表
 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
 例子： head = [1,2,2,1] → true
 * 解题思路：
 * 1. 找到链表的中点，使用快慢指针快速找到
 * 2. 将后半部分压入栈中
 * 3. 栈不为空时，头指针和栈顶元素比较
 */
