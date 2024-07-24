//
// Created by fvikhnin on 24.07.24.
//

#ifndef CPP_HEX_UTILS_H

#include <string>
#include <sstream>

namespace hex {
    int get_digit(char &c1) {
        int c = 0;
        if ('0' <= c1 && c1 <= '9') {
            c = c1 - '0';
        } else {
            c = (c1 - 'a') + 10;
        }
        return c;
    }

    std::string parse_hex_to_bytes(const std::string &s) {
        std::string out;
        std::string tmp;

        auto ss = std::stringstream(s);

        while (getline(ss, tmp, ' ')) {
            out += (char) (16 * get_digit(tmp[6]) + get_digit(tmp[7]));
            out += (char) (16 * get_digit(tmp[4]) + get_digit(tmp[5]));
            out += (char) (16 * get_digit(tmp[2]) + get_digit(tmp[3]));
            out += (char) (16 * get_digit(tmp[0]) + get_digit(tmp[1]));
        }

        return out;
    }
}
#define CPP_HEX_UTILS_H

#endif //CPP_HEX_UTILS_H
