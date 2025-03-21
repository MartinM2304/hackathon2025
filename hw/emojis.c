#include <stdint.h>

#define WIDTH 16
#define HEIGHT 16

typedef struct {
    uint8_t r;
    uint8_t g;
    uint8_t b;
} Pixel;

Pixel smile[HEIGHT][WIDTH] = {
    {{0, 0, 0}, {255, 255, 255}, {244, 67, 54}, {233, 30, 99}, {156, 39, 176}, {103, 58, 183}, {63, 81, 181}, {33, 150, 243},
     {3, 169, 244}, {0, 188, 212}, {0, 150, 136}, {76, 175, 80}, {139, 195, 74}, {205, 220, 57}, {255, 235, 59}, {255, 193, 7}},
    {{255, 152, 0}, {255, 87, 34}, {121, 85, 72}, {158, 158, 158}, {96, 125, 139}, {255, 235, 238}, {255, 205, 210}, {239, 154, 154},
     {229, 115, 115}, {239, 83, 80}, {229, 57, 53}, {211, 47, 47}, {198, 40, 40}, {183, 28, 28}, {255, 138, 128}, {255, 82, 82}},
    {{255, 23, 68}, {213, 0, 0}, {252, 228, 236}, {248, 187, 208}, {244, 143, 177}, {240, 98, 146}, {236, 64, 122}, {233, 30, 99},
     {216, 27, 96}, {194, 24, 91}, {173, 20, 87}, {136, 14, 79}, {255, 128, 171}, {255, 64, 129}, {245, 0, 87}, {197, 17, 98}},
    {{243, 229, 245}, {225, 190, 231}, {206, 147, 216}, {186, 104, 200}, {171, 71, 188}, {156, 39, 176}, {142, 36, 170}, {123, 31, 162},
     {106, 27, 154}, {74, 20, 140}, {234, 128, 252}, {224, 64, 251}, {213, 0, 249}, {170, 0, 255}, {238, 232, 246}, {209, 196, 233}},
    {{179, 157, 219}, {149, 117, 205}, {126, 87, 194}, {103, 58, 183}, {94, 53, 177}, {81, 45, 168}, {69, 39, 160}, {49, 27, 146},
     {179, 136, 255}, {124, 77, 255}, {101, 31, 255}, {98, 0, 234}, {232, 234, 246}, {197, 202, 233}, {159, 168, 218}, {121, 134, 203}},
    {{92, 107, 192}, {63, 81, 181}, {57, 73, 171}, {48, 63, 159}, {40, 53, 147}, {26, 35, 126}, {140, 158, 255}, {83, 109, 254},
     {61, 90, 254}, {48, 79, 254}, {227, 242, 253}, {187, 222, 251}, {144, 202, 249}, {100, 181, 246}, {66, 165, 245}, {33, 150, 243}},
    {{30, 136, 229}, {25, 118, 210}, {21, 101, 192}, {13, 71, 161}, {130, 177, 255}, {68, 138, 255}, {41, 121, 255}, {41, 98, 255},
     {225, 245, 254}, {179, 229, 252}, {129, 212, 250}, {79, 195, 247}, {41, 182, 246}, {3, 169, 244}, {3, 155, 229}, {2, 136, 209}},
    {{2, 119, 189}, {1, 87, 155}, {128, 216, 255}, {64, 196, 255}, {0, 176, 255}, {0, 145, 234}, {224, 247, 250}, {178, 235, 242},
     {128, 222, 234}, {77, 208, 225}, {38, 198, 218}, {0, 188, 212}, {0, 172, 193}, {0, 151, 167}, {0, 131, 143}, {0, 96, 100}},
    {{132, 255, 255}, {24, 255, 255}, {0, 229, 255}, {0, 184, 212}, {224, 242, 241}, {178, 223, 219}, {128, 203, 196}, {77, 182, 172},
     {38, 166, 154}, {0, 150, 136}, {0, 137, 123}, {0, 121, 107}, {0, 105, 92}, {0, 77, 64}, {167, 255, 235}, {100, 255, 218}}
};


Pixel sad[16][16] = {
    { {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {34, 177, 76}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {34, 177, 76}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0} },
    { {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {34, 177, 76}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {34, 177, 76}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0} },
    { {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {229, 170, 122}, {229, 170, 122}, {229, 170, 122}, {229, 170, 122}, {229, 170, 122}, {229, 170, 122}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0} },
    { {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {229, 170, 122}, {229, 170, 122}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {229, 170, 122}, {229, 170, 122}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0} },
    { {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {229, 170, 122}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {229, 170, 122}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0} }
};
