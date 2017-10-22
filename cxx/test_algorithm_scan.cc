

/*
原型：
输入时具有n个浮点数的向量x，输出是输入向量的任何连续的子向量中的最大和；

分治算法
将n个向量画分为a和b两个子向量，递归找出a和b中的最大子向量；找出跨越a和b之间的边界的最大子向量；返回3个中的最大者

扫描算法：
前i个原始中，最大的总和子数组要么在前i-1个元素中（maxsofar），要么在其结束位置（maxendinghere）
 * */

#include <vector>
#include <iostream>

// g++ test_algorithm_scan.cc -std=c++11
int main() {
    std::vector<int> num ={
        31,-41,59,26,-53,58,97,-93,-23,84
    };
    
    int maxendinghere; //前i个之后
    int maxsofar =0; //
    for (int i =0; i<num.size();i++) {
        if (i==0) {
            maxendinghere = num[i];
        }
        else {
            maxendinghere = maxendinghere + num[i];
        }

        //前i个之和为负数时
        if (maxendinghere<0) {
            maxendinghere =0;
        }

        maxsofar = std::max(maxendinghere, maxsofar);
    }
    
    std::cout << "maxsofar:" << maxsofar << std::endl;
    std::cout << "maxendinghere:" << maxendinghere << std::endl;
}
