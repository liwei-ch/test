#include <tr1/functional>
#include <iostream>


   typedef std::tr1::function<void( const std::string&, const std::string&)> F1;

   void test(const std::string& s1, const std::string &s2) {
     std::cout << s1 << " " << s2 << std::endl;
   }
   
int main() {
    std::string s1("s1");
    std::string s2("s2");
   F1 f= std::tr1::bind(test, s1,s2);
   std::cout << "first---->" << f.T1 << std::endl;
   F1();
}
