
// 

#include<iostream>

using namespace std;

struct Siswa {
    string nama;
    string sekolah;
    unsigned int uangsaku;
};

int main()
{
    struct Siswa siswa1,siswa2,siswa3,siswa4,siswa5,siswa6;
    
    siswa1.nama = "syamil abdurrahman";
    siswa1.sekolah = "katod";
    siswa1.uangsaku = 55000;
    

    
    siswa2.nama = "gogops";
    siswa2.sekolah = "katod";
    siswa2.uangsaku = 55000;
    

    
    siswa3.nama = "gogops";
    siswa3.sekolah = "katod";
    siswa3.uangsaku = 55000;

    
    siswa4.nama = "gogops";
    siswa4.sekolah = "katod";
    siswa4.uangsaku = 123000;

    
    siswa5.nama = "satria";
    siswa5.sekolah = "katod";
    siswa5.uangsaku = 324000;

    
    siswa6.nama = "moler";
    siswa6.sekolah = "katod";
    siswa6.uangsaku = 5500000;
    
    cout << siswa4.nama << " sekolah di " << siswa4.sekolah << endl;
    cout << "dengan duit jajan " << siswa4.uangsaku << " per harinya..." << endl;
    cout << endl;
    cout << siswa6.nama << " sekolah di " << siswa6.sekolah << endl;
    cout << "dengan duit jajan " << siswa6.uangsaku << " per harinya..." << endl;
    cout << endl;
    cout << siswa2.nama << " sekolah di " << siswa2.sekolah << endl;
    cout << "dengan duit jajan " << siswa2.uangsaku << " per harinya..." << endl;
    cout << endl;
    cout << siswa5.nama << " sekolah di " << siswa5.sekolah << endl;
    cout << "dengan duit jajan " << siswa5.uangsaku << " per harinya..." << endl;
    return 0;
    
}
