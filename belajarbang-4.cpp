//Date Of Creation 3/11/2023

#include <iostream>
using namespace std;

int main()
{
    string nama;
    int umur;
    
    cout << "masukan nama mu!: ";
    getline(cin,nama);
    
    cout << "masukan umur mu! :";
    cin >> umur;
    
    cout << "\n";
    
    cout << "selamat datang " << nama << ",semangat!!" << endl;
    cout << "usia lu sekarang " << umur << " tahun," << "besok lu udah mau kuliah..." <<endl;
    
    return 0;
}
