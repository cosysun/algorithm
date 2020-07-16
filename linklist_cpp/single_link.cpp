#include <stdio.h>
#include <iostream>
#include <random>

struct ListNode {
    int value;
    ListNode* pNext;
};

ListNode* GenRandomSingleList(int num) 
{
    ListNode* pHead = nullptr;
    ListNode* pPrev = nullptr;
    for (int i = 0; i < num; i++)
    {
        ListNode* pNode = new ListNode();
        pNode->value = i;
        if (i == 0)
        {
            pHead = pNode;
        } else {
            pPrev->pNext = pNode;
        }
        pPrev = pNode;
    }

    return pHead;
}

ListNode* Rervese(ListNode* pHead) {
    if (pHead == NULL)
    {
        return pHead;
    }

    ListNode* pCur = pHead;
    ListNode* pPrev = nullptr;
    while (pCur ->pNext!= NULL)
    {
        ListNode* pTmp = pCur->pNext;
        pCur->pNext = pPrev; 
        pPrev = pCur;
        pCur = pTmp;
    }
    if (pCur->pNext == nullptr)
    {
        pCur->pNext = pPrev;
    }
   return pCur; 
}

int PrintKNode(ListNode* pHead, int k) {
    // check input valid
    if (pHead == NULL || k <= 0)
    {
        std::cout << "input param valid" << std::endl;
        return -1;
    }

    // check k
    int len = 0;
    ListNode* pNode = pHead;
    while (pNode->pNext != NULL)
    {
        ++len;
        pNode = pNode->pNext;
    }
    if (k > len)
    {
         std::cout << "k:" << k << " is bigger than:"<< len << std::endl;
        return -1;
    }
    ListNode* pSlow = pHead;
    ListNode* pQuick = pHead;
    for (int i = 0; i < k - 1; i++)
    {
        pQuick = pQuick->pNext;
    }
    while (pQuick->pNext != NULL)
    {
        pSlow = pSlow->pNext;
        pQuick = pQuick->pNext;
    }
    std::cout << pSlow->value << std::endl;
    return 0;
}

void PrintListNode(ListNode* pHead) {
    if (pHead == NULL)
    {
        return;
    }

    ListNode* pCur = pHead;
    std::string res = "";
    while (pCur->pNext!= NULL)
    {
        res += std::to_string(pCur->value) + "->";
        pCur = pCur->pNext;
    }
    res += std::to_string(pCur->value);

    std::cout << res << std::endl;
}

ListNode* MergeList(ListNode* pHead1, ListNode*pHead2) {
    if (pHead1 == NULL || pHead2 == NULL)
    {
        return NULL;
    }

    ListNode* pCur1 = pHead1;
    ListNode* pCur2 = pHead2;
    ListNode* pMerge = new ListNode();
    ListNode* pCurMerge = pMerge;
    while (pCur1 != NULL && pCur2 != NULL)
    {
        if (pCur1->value <= pCur2->value)
        {
            pCurMerge->pNext = pCur1;
            pCur1 = pCur1->pNext;
        } else {
            pCurMerge->pNext = pCur2;
            pCur2 = pCur2->pNext;
        }
        pCurMerge = pCurMerge->pNext;
    }
    if (pCur1 == NULL)
    {
        pCurMerge->pNext = pCur2;
    }

    if (pCur2 == NULL) {
        pCurMerge->pNext = pCur1;
    }
   return pMerge;
}

ListNode* MergeList2(ListNode* pHead1, ListNode*pHead2) {
    if (pHead1 == NULL)
    {
        return pHead2;
    }

    if (pHead2 == NULL)
    {
        return pHead1;
    }
    if (pHead1->value <= pHead2->value)
    {
       pHead1->pNext =  MergeList2(pHead1->pNext, pHead2);
       return pHead1;
    } else {
       pHead2->pNext =  MergeList2(pHead1, pHead2->pNext); 
       return pHead2;
    }
    return nullptr;
}

int main() {
    // ListNode* pNode = GenRandomSingleList(10);
    // PrintListNode(pNode);

    // ListNode* pResv = Rervese(pNode);
    // PrintListNode(pResv);

    // PrintKNode(pResv, 1);

    ListNode* pNode1 = GenRandomSingleList(10);
    PrintListNode(pNode1);
    ListNode* pNode2 = GenRandomSingleList(5);
    PrintListNode(pNode2);

    ListNode* mergeNode = MergeList2(pNode1, pNode2);
    PrintListNode(mergeNode);
}