#include <stdio.h>
#include <iostream>
#include <random>
#include <stack>

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

ListNode* GenLoopSingleList(int num, int point) 
{
    if (num <=0 || point >= num - 1 || point < 0)
    {
        return nullptr;
    }
    
    ListNode* pHead = nullptr;
    ListNode* pPrev = nullptr;
    ListNode* pPoint = nullptr;
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
        if (i == point)
        {
            pPoint = pNode;
        }
        if (i == (num - 1))
        {
            pNode->pNext = pPoint;
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

void PrintReverseListNode(ListNode* pHead) {
    if (pHead == NULL)
    {
        return;
    }
    std::stack<ListNode*> nodeStack;
    ListNode* pCur = pHead;
    while (pCur != NULL)
    {
        nodeStack.push(pCur);
        pCur = pCur->pNext;
    }

    std::string res = "";
    int len = nodeStack.size();
    int i = 0;
    while (!nodeStack.empty())
    {
        if (i == len - 1)
        {
            res += std::to_string(nodeStack.top()->value);
        } else {
            res += std::to_string(nodeStack.top()->value) + "->";
        }
        
        nodeStack.pop();
        i++;
    }
    std::cout << res << std::endl;
}

int GetListLen(ListNode* pHead) {
    if (pHead == NULL)
    {
        return 0;
    }
    ListNode* pCur = pHead;
    int len = 0;
    while (pCur != NULL)
    {
        ++len;
        pCur = pCur->pNext;
    }
    return len;
}

ListNode* FindFirstCommonNode(ListNode* pHead1, ListNode* pHead2) {
    if (pHead1 == NULL || pHead2 == NULL)
    {
        return NULL;
    }
    int len1 = GetListLen(pHead1);
    int len2 = GetListLen(pHead2);
    ListNode* pLong = pHead1;
    ListNode* pShort = pHead2;
    int step_diff = len1 - len2;
    if (len2 > len1)
    {
        pLong = pHead2;
        pShort = pHead1;
        step_diff = len2- len1;
    }
    for (int i = 0; i < step_diff; i++)
    {
        pLong = pLong->pNext;
    }

    while (pShort != NULL && pLong != NULL)
    {
        if(pShort == pLong) {
            return pLong;
        }
        pLong = pLong->pNext;
        pShort = pShort->pNext;
    }
    return nullptr; 
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

bool HasListLoop(ListNode* pHead) {
    if (pHead == NULL)
    {
        return false;
    }
    ListNode* pQuick = pHead;
    ListNode* pSlow = pHead;
    while (pSlow != nullptr && pQuick != nullptr && pQuick->pNext != nullptr)
    {
        pSlow = pSlow->pNext;
        pQuick = pQuick->pNext->pNext; 
        if (pQuick == pSlow)
        {
           return true; 
        }
    }
    return false;
} 

ListNode* InsertTail(ListNode* pHead, int value) {
    if (pHead == NULL)
    {
        return nullptr;
    }

    ListNode* pCur = pHead; 
    while (pCur->pNext != NULL)
    {
        pCur = pCur->pNext;
    }
    ListNode* pNode = new ListNode();
    pNode->value = value;
    pNode->pNext = nullptr;
    pCur->pNext = pNode;
    return  pHead;
}
    
ListNode* Insert(ListNode* pHead, int value, int insertValue) {
    if (pHead == NULL)
    {
        return nullptr;
    }
    ListNode* pCur = pHead; 
    ListNode* pInsertNode = new ListNode();
    pInsertNode->value = insertValue;
    
    while (pCur != NULL)
    {
        if(pCur->value == value) {
            ListNode* pTmp = pCur->pNext;
            pCur->pNext = pInsertNode;
            pInsertNode->pNext = pTmp;
        }
        pCur = pCur->pNext;
    }
    return  pHead;
}

ListNode* Delete(ListNode* pHead, int value) {
    if (pHead == NULL)
    {
        return nullptr;
    }
    ListNode* pCur = pHead; 
    ListNode* pPrev = nullptr;
    while (pCur != NULL)
    {
        if(pCur->value == value) {
            if (pPrev == nullptr)
            {
                pHead = pCur->pNext;
            } else {
                pPrev->pNext = pCur->pNext;
            }
            ListNode* pTmp = pCur;
            pCur = pCur->pNext;

            //释放
            delete pTmp;
        }  else {
            pPrev = pCur;
            pCur = pCur->pNext;
        }
    }
    return  pHead;
}

ListNode* ReverseChange(ListNode* pHead, int k) {
    if (pHead == NULL || k <= 1)
    {
        return pHead;
    }
    int len =0;
    ListNode* pNode = pHead;
    while (pNode != NULL)
    {
        len++;
        pNode = pNode->pNext;
    }
    int n = len / k;
    int skipNum = len % k;
    pNode = pHead;
    for (int i = 0; i < skipNum - 2; i++)
    {
        pNode = pNode->pNext;
    }
    ListNode* pTop = pNode;
    ListNode* pTail = pNode;
    int step = k;
    while (pTail != NULL)
    {
        if ((--step) != 0)
        {
            ListNode *pTmp = pTop->pNext;
            ListNode *pTmp1 = pTail->pNext;
            pTop->pNext = pTail;
            pTail->pNext = pTmp;
            pTail = pTmp1;
        } else {
            step = k;
            pTop = pTail;
        }
    }
    
    
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

    ListNode* pLoopNode = GenLoopSingleList(10, 4);
    ListNode* pNode = GenRandomSingleList(10);
    if (HasListLoop(pNode)) {
        std::cout << "has loop" << std::endl;
    } else {
        std::cout << "no loop" << std::endl;
    }
    
    InsertTail(pNode, 11);
    PrintListNode(pNode);
    Insert(pNode, 1, 11);
    PrintListNode(pNode);

    ListNode* t = Delete(pNode, 11);
     PrintListNode(t);
    
    PrintReverseListNode(t);
}
