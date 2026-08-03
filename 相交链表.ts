function getIntersectionNode(
  headA: ListNode | null,
  headB: ListNode | null,
): ListNode | null {
  const stack_A: ListNode[] = [];
  const stack_B: ListNode[] = [];
  let res: ListNode | null = null;
  let cur_A = headA;
  let cur_B = headB;
  while (cur_A) {
    stack_A.push(cur_A);
    cur_A = cur_A.next;
  }
  while (cur_B) {
    stack_B.push(cur_B);
    cur_B = cur_B.next;
  }
  while (stack_A.length && stack_B.length) {
    const node_A = stack_A.pop();
    const node_B = stack_B.pop();
    if (node_A === node_B) res = node_A as ListNode;
    else break;
  }
  return res;
}

// def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> ListNode:
//     A, B = headA, headB
//     while A != B:
//         A = A.next if A else headB
//         B = B.next if B else headA
//     return A
