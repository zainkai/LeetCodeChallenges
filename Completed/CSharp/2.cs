/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public ListNode AddTwoNumbers (ListNode l1, ListNode l2) {
        int carry = 0;
        int sum = 0;
        ListNode prevTemp, temp;
        ListNode head = new ListNode(0);
        prevTemp = head;

        while (l1 != null || l2 != null) {
            sum = 0;
            if(l1 != null) {
                sum += l1.val;
                l1 = l1.next;
            }
            if(l2 != null) {
                sum += l2.val;
                l2 = l2.next;
            }

            if(sum + carry > 9) {
                sum = sum -10 < 0 ? 0 : sum;
                temp = new ListNode(sum - 10);
                carry = 1;
            } else {
                temp = new ListNode(sum + carry);
                carry = 0;
            }

            
            prevTemp.next = temp;
            prevTemp = prevTemp.next;
            temp = temp.next;
        }

        if ((l1 == null && l2 == null && carry == 1) {
            temp = new ListNode(1);
            prevTemp.next = temp;
        }

        return head.next;
    }
};