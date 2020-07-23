#include <stdio.h>
#include <iostream>
#include <vector>
   
using namespace std;

// bubble sort
std::vector<int> bublleSort(std::vector<int> arry) {
    if (arry.size() <= 1)
    {
        return arry;
    }
    size_t size = arry.size();
    for (size_t i = 0; i < size; i++)
    {
        for (size_t j = 0; j < size - i - 1; j++)
        {
                if (arry[j+1] < arry[j])
                {
                    int tmp = arry[j + 1];
                    arry[j + 1] = arry[j];
                    arry[j] = tmp;
                }
        }
        
    }
    return arry;
}

int main() {
    std::vector<int> arry;
    arry.push_back(4);
    arry.push_back(1);
    arry.push_back(5);
    arry.push_back(3);
    arry.push_back(8);
    arry.push_back(6);

    arry = bublleSort(arry);
    std::cout << "sort:";
    for (size_t i = 0; i < arry.size(); i++)
    {
        std::cout << arry[i] << std::endl;
    }
    
    return 0;
}