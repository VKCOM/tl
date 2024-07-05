#include <iostream>
#include "cpp/a.top2.hpp"

std::string to_hex(const uint8_t *data, size_t count) {
	static const char hexdigits[] = "0123456789abcdef";

	std::string result(count * 2, char());
	for (size_t i = 0; i != count; ++i) {
		uint8_t ch        = data[i];
		result[i * 2]     = hexdigits[(ch >> 4) & 0xf];
		result[i * 2 + 1] = hexdigits[ch & 0xf];
	}
	return result;
}

int main() {
	basictl::tl_ostream_string str;
	basictl::tl_ostream_string str2;


    tl2::a::Top2 top2;

    top2.write(str);
    auto & buf = str.get_buffer();
    std::cout << top2.tl_name() << ": " << to_hex(reinterpret_cast<const uint8_t *>(buf.data()), buf.size()) << std::endl;

//    top3.write(str2);

//    std::cout << top3.tl_name() << ": " << to_hex(reinterpret_cast<const uint8_t *>(buf2.data()), buf2.size()) << std::endl;
    return 0;
}
