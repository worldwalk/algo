#用python实现一个链表的反转
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def reverseList(head: ListNode) -> ListNode:
    prev = None
    curr = head
    while curr:
        next = curr.next
        curr.next = prev
        prev = curr
        curr = next
    return prev

def printList(head: ListNode):
    while head:
        print(head.val, end=" ")
        head = head.next
    print()

#测试
head = ListNode(1, ListNode(2, ListNode(3, ListNode(4, ListNode(5)))))
printList(reverseList(head))