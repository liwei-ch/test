#include <iostream>


template<class T>

struct OpNewCreator {
    static T* Create() {
        return new T;
    }
}

int main() {

}
