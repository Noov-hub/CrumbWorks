#include <iostream>

using namespace std;

int main ()
{
    char karakter1, karakter2, karakter3;
    

    cout << "input 3 karakter sembarang" << endl;
    cout << "=====================" << endl;
    cout << endl;

    cout << "karakter nya 1 digit aja y..." << endl ;
    cout << "karakter pertama: ";
    
    cin >> karakter1;

    cout << "karakter kedua: ";
    cin >> karakter2;

    cout << "karakter ketiga: ";
    cin >> karakter3;

    cout << endl;

    cout << "karakter yang di input :";
    cout << karakter1 << ", " << karakter2 << ", dan" << karakter3;
    cout << endl;
    
    return 0;
}