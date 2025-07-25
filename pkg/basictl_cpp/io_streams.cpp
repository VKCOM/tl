/** TLGEN: CPP INCLUDES */
#include "io_streams.h"
#include "io_throwable_streams.h"
/** TLGEN: CPP INCLUDES END */

namespace basictl {
    tl_istream::tl_istream(tl_input_connector &provider) : provider(&provider) {}

    tl_istream::tl_istream(tl_throwable_istream &from) : provider(nullptr) {
        from.pass_data(*this);
    }

    bool tl_istream::string_read(std::string &value) noexcept {
        if (!ensure_byte()) [[unlikely]] {
            return false;
        }
        auto len = size_t(static_cast<unsigned char>(*ptr));
        if (len >= basictl::TL_BIG_STRING_MARKER) [[unlikely]] {
            if (len > basictl::TL_BIG_STRING_MARKER) [[unlikely]] {
                return set_error_sequence_length();
            }
            uint32_t len32 = 0;
            if (!nat_read(len32)) [[unlikely]] {
                return false;
            }
            len = len32 >> 8U;
            value.clear();
            if (!fetch_data_append(value, len)) [[unlikely]] {
                return false;
            }
            if (!fetch_pad((-len) & 3)) [[unlikely]] {
                return false;
            }
            return true;
        }
        auto pad = ((-(len + 1)) & 3);
        auto fullLen = 1 + len + pad;
        if (ptr + fullLen > end_block) [[unlikely]] {
            ptr += 1;
            value.clear();
            if (!fetch_data_append(value, len)) [[unlikely]] {
                return false;
            }
            if (!fetch_pad(pad)) [[unlikely]] {
                return false;
            }
            return true;
        }
        // fast path for short strings that fully fit in buffer
        uint32_t x = 0;
        std::memcpy(&x, ptr + fullLen - 4, 4);
        if ((x & ~(0xFFFFFFFFU >> (8 * pad))) != 0) [[unlikely]] {
            return set_error_string_padding();
        }
        value.assign(reinterpret_cast<const char *>(ptr + 1), len);
        ptr += fullLen;
        return true;
    }

    void tl_istream::sync() noexcept {
        provider->advance(ptr - start_block);
        start_block = ptr;
    }

    bool tl_istream::has_error() const noexcept {
        return error.has_value();
    }

    std::optional<tl_error> &tl_istream::get_error() noexcept {
        return error;
    }

    bool tl_istream::set_error(tl_error_type type, const char *what) noexcept {
        if (!error.has_value()) {
            error = tl_stream_error(type, what);
        }
        return false;
    }

    bool tl_istream::set_error_eof() noexcept { return set_error(tl_error_type::STREAM_EOF, "EOF"); }

    bool tl_istream::set_error_sequence_length() noexcept {
        return set_error(tl_error_type::INCORRECT_SEQUENCE_LENGTH, "sequence_length");
    }

    bool tl_istream::set_error_string_padding() noexcept {
        return set_error(tl_error_type::INCORRECT_STRING_PADDING, "string_padding");
    }

    bool tl_istream::set_error_expected_tag() noexcept {
        return set_error(tl_error_type::UNEXPECTED_TAG, "expected_tag");
    }

    bool tl_istream::set_error_union_tag() noexcept { return set_error(tl_error_type::UNEXPECTED_TAG, "union_tag"); };

    bool tl_istream::set_error_unknown_scenario() noexcept {
        return set_error(tl_error_type::UNKNOWN_SCENARIO, "union_tag");
    };

    void tl_istream::grow_buffer() noexcept {
        ptr = end_block;
        provider->advance(ptr - start_block);
        auto new_buffer_request = provider->get_buffer();
        if (new_buffer_request) {
            auto new_buffer = new_buffer_request.value();
            ptr = new_buffer.data();
            start_block = ptr;
            end_block = ptr + new_buffer.size();
        } else {
            error = new_buffer_request.error();
        }
    }

    bool tl_istream::ensure_byte() noexcept {
        if (ptr >= end_block) [[unlikely]] {
            grow_buffer();
            if (ptr == end_block) [[unlikely]] {
                return set_error_eof();
            }
        }
        return true;
    }

    bool tl_istream::fetch_data(void *vdata, size_t size) noexcept {
        if (ptr + size > end_block) [[unlikely]] {
            return fetch_data2(vdata, size);
        }
        std::memcpy(reinterpret_cast<char *>(vdata), ptr, size);
        ptr += size;
        return true;
    }

    bool tl_istream::fetch_data2(void *vdata, size_t size) noexcept {
        char *data = reinterpret_cast<char *>(vdata);
        for (; ptr + size > end_block;) [[unlikely]] {
            std::memcpy(data, ptr, end_block - ptr);
            data += end_block - ptr;
            size -= end_block - ptr;
            grow_buffer();
            if (ptr == end_block) [[unlikely]] {
                return set_error_eof();
            }
        }
        std::memcpy(data, ptr, size);
        ptr += size;
        return true;
    }

    bool tl_istream::fetch_data_append(std::string &value, size_t size) noexcept {
        for (; ptr + size > end_block;) [[unlikely]] {
            // assert(ptr <= end)
            value.append(reinterpret_cast<const char *>(ptr), end_block - ptr);
            size -= end_block - ptr;
            grow_buffer();
            // assert(ptr <= end)
            if (ptr == end_block) [[unlikely]] {
                return set_error_eof();
            }
        }
        value.append(reinterpret_cast<const char *>(ptr), size);
        ptr += size;
        return true;
    }

    bool tl_istream::fetch_pad(size_t len) noexcept {
        uint32_t x = 0;
        if (!fetch_data(&x, len)) [[unlikely]] {
            return false;
        }
        if (x != 0) [[unlikely]] {
            return set_error_string_padding();
        }
        return true;
    }

    void static throw_exception_from_tl_stream(tl_error &error) {
        switch (error.index()) {
            case 0:
                throw std::get<0>(error);
            case 1:
                throw std::get<1>(error);
        }
    }

    void tl_istream::pass_data(tl_throwable_istream &to) {
        to.provider = provider;
        to.ptr = ptr;
        to.start_block = start_block;
        to.end_block = end_block;

        if (has_error()) {
            ::basictl::throw_exception_from_tl_stream(error.value());
        }
    }

    tl_ostream::tl_ostream(tl_output_connector &provider) : provider(&provider) {}

    tl_ostream::tl_ostream(tl_throwable_ostream &from) : provider(nullptr) {
        from.pass_data(*this);
    }

    bool tl_ostream::string_write(const std::string &value) {
        auto len = value.size();
        if (len > basictl::TL_MAX_TINY_STRING_LEN) [[unlikely]] {
            if (len > basictl::TL_BIG_STRING_LEN) [[unlikely]] {
                return set_error_sequence_length();
            }
            uint32_t p = (len << 8U) | basictl::TL_BIG_STRING_MARKER;
            if (!store_data(&p, 4)) [[unlikely]] {
                return false;
            }
            if (!store_data(value.data(), value.size())) [[unlikely]] {
                return false;
            }
            if (!store_pad((-len) & 3)) [[unlikely]] {
                return false;
            }
            return true;
        }
        auto pad = ((-(len + 1)) & 3);
        auto fullLen = 1 + len + pad;
        if (ptr + fullLen > end_block) [[unlikely]] {
            auto p = static_cast<unsigned char>(len);
            if (!store_data(&p, 1)) [[unlikely]] {
                return false;
            }
            if (!store_data(value.data(), value.size())) [[unlikely]] {
                return false;
            }
            if (!store_pad(pad)) [[unlikely]] {
                return false;
            }
            return true;
        }
        // fast path for short strings that fully fit in buffer
        uint32_t x = 0;
        std::memcpy(ptr + fullLen - 4, &x, 4); // padding first
        *ptr = static_cast<std::byte>(len);
        std::memcpy(ptr + 1, value.data(), len);
        ptr += fullLen;
        return true;
    }

    void tl_ostream::sync() noexcept {
        provider->advance(ptr - start_block);
        start_block = ptr;
    }

    bool tl_ostream::has_error() const noexcept {
        return error.has_value();
    }

    std::optional<tl_error> &tl_ostream::get_error() noexcept {
        return error;
    }

    bool tl_ostream::set_error(tl_error_type type, const char *what) noexcept {
        if (!error.has_value()) {
            error = tl_stream_error(type, what);
        }
        return false;
    }

    bool tl_ostream::set_error_eof() noexcept { return set_error(tl_error_type::STREAM_EOF, "EOF"); }

    bool tl_ostream::set_error_sequence_length() noexcept {
        return set_error(tl_error_type::INCORRECT_SEQUENCE_LENGTH, "sequence_length");
    }

    bool tl_ostream::set_error_string_padding() noexcept {
        return set_error(tl_error_type::INCORRECT_STRING_PADDING, "string_padding");
    }

    bool tl_ostream::set_error_bool_tag() noexcept { return set_error(tl_error_type::UNEXPECTED_TAG, "bool_tag"); }

    bool tl_ostream::set_error_expected_tag() noexcept {
        return set_error(tl_error_type::UNEXPECTED_TAG, "expected_tag");
    }

    bool tl_ostream::set_error_union_tag() noexcept { return set_error(tl_error_type::UNEXPECTED_TAG, "union_tag"); }

    bool tl_ostream::set_error_unknown_scenario() noexcept {
        return set_error(tl_error_type::UNKNOWN_SCENARIO, "union_tag");
    };

    void tl_ostream::grow_buffer() {
        ptr = end_block;
        provider->advance(ptr - start_block);
        auto new_buffer_request = provider->get_buffer();
        if (new_buffer_request) {
            auto new_buffer = new_buffer_request.value();
            ptr = new_buffer.data();
            start_block = ptr;
            end_block = ptr + new_buffer.size();
        } else {
            error = new_buffer_request.error();
        }
    }

    bool tl_ostream::store_data(const void *vdata, size_t size) {
        if (ptr + size > end_block) [[unlikely]] {
            return store_data2(vdata, size);
        }
        std::memcpy(ptr, reinterpret_cast<const char *>(vdata), size);
        ptr += size;
        return true;
    }

    bool tl_ostream::store_data2(const void *vdata, size_t size) {
        const char *data = reinterpret_cast<const char *>(vdata);
        for (; ptr + size > end_block;) [[unlikely]] {
            std::memcpy(ptr, data, end_block - ptr);
            data += end_block - ptr;
            size -= end_block - ptr;
            grow_buffer();
            if (ptr == end_block) [[unlikely]] {
                return set_error_eof();
            }
        }
        std::memcpy(ptr, data, size);
        ptr += size;
        return true;
    }

    bool tl_ostream::store_pad(size_t size) {
        for (; ptr + size > end_block;) [[unlikely]] {
            // assert(ptr <= end)
            std::memset(ptr, 0, end_block - ptr);
            size -= end_block - ptr;
            grow_buffer();
            // assert(ptr <= end)
            if (ptr == end_block) [[unlikely]] {
                return set_error_eof();
            }
        }
        if (size != 0) {
            ptr[0] = static_cast<std::byte>(0);
            ptr[size - 1] = static_cast<std::byte>(0);
            ptr[size / 2] = static_cast<std::byte>(0);
            ptr += size;
        }
        return true;
    }

    void tl_ostream::pass_data(tl_throwable_ostream &to) {
        to.provider = provider;
        to.ptr = ptr;
        to.start_block = start_block;
        to.end_block = end_block;

        if (has_error()) {
            ::basictl::throw_exception_from_tl_stream(error.value());
        }
    }
} // namespace basictl