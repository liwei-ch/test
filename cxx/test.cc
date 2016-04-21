#include <vector>
#include <iostream>

struct A{
    int a;
};

int main() {

    std::vector<struct A*> v;
    std::cout << v.size() << std::endl;
    v.push_back(new struct A);
    struct A *p = v[2];
    std::cout << "p= "<< p  << std::endl;
    std::cout << "v.size="<< v.size() << std::endl;
    return 0;
}

