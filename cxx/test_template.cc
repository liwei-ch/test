#include <iostream>


template<class T>

struct OpNewCreator {
    static T* Create() {
        return new T;
    }
};

class Widget{
};

template <template <class Created> class CreationPolicy>
class WidgetManager : public CreationPolicy< Widget> {};


template <template <class> class CreationPolicy = OpNewCreator>
class WidgetManager2 : public CreationPolicy< Widget> {};

typedef WidgetManager<OpNewCreator> MyWidgetMgr;

int main() {

}
