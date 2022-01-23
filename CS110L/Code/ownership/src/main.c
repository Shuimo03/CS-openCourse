#include<stdio.h>
#include<stdlib.h>
#include<string.h>

void om_nom_nom(char* s) {
    printf("%s\n", s);
    free(s);
}
int main(){
char *s = strdup("hello");
    om_nom_nom(s);
    om_nom_nom(s);
}