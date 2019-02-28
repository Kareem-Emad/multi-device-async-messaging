#include <fstream>
#include <stdlib.h>

using namespace std;

int main()
{
    ofstream out;
    out.open("input.txt");

    for(int i=0; i<100000; i++){
        int x = (rand() % 100000) + 1;
        out << x << endl;
    }

    out.close();
}