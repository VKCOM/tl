/** TLGEN: CPP INCLUDES */
#include "io_streams.h"
#include "io_throwable_streams.h"
/** TLGEN: CPP INCLUDES END */

namespace basictl {
    tl_throwable_istream::tl_throwable_istream(tl_input_connector &provider) : provider(&provider) {}

    void tl_throwable_istream::string_read(std::string &value) {
        ensure_byte();
        auto len = size_t(static_cast<unsigned char>(*ptr));
        if (len >= TL_BIG_STRING_MARKER) [[unlikely]] {
            if (len > TL_BIG_STRING_MARKER) [[unlikely]] {
                throw tl_error(tl_error_type::INCORRECT_SEQUENCE_LENGTH, "TODO - huge string");
            }
            uint32_t len32 = 0;
            nat_read(len32);
            len = len32 >> 8U;
            value.clear();
            fetch_data_append(value, len);
            fetch_pad((-len) & 3);
            return;
        }
        auto pad = ((-(len + 1)) & 3);
        auto fullLen = 1 + len + pad;
        if (ptr + fullLen > end_block) [[unlikely]] {
            ptr += 1;
            value.clear();
            fetch_data_append(value, len);
            fetch_pad(pad);
            return;
        }
        // fast path for short strings that fully fit in buffer
        uint32_t x = 0;
        std::memcpy(&x, ptr + fullLen - 4, 4);
        if ((x & ~(0xFFFFFFFFU >> (8 * pad))) != 0) [[unlikely]] {
            throw tl_error(tl_error_type::INCORRECT_STRING_PADDING, "incorrect string padding");
        }
        value.assign(reinterpret_cast<const char*>(ptr + 1), len);
        ptr += fullLen;
   }

    void tl_throwable_istream::last_release() noexcept {
        provider->release_buffer(ptr - start_block);
        start_block = ptr;
    }

    void tl_throwable_istream::grow_buffer() {
        ptr = end_block;
        provider->release_buffer(ptr - start_block);
        auto new_buffer_request = provider->get_buffer();
        if (new_buffer_request) {
            auto new_buffer = new_buffer_request.value();
            ptr = new_buffer.data();
            start_block = ptr;
            end_block = ptr + new_buffer.size();
        } else {
            throw tl_connector_error(new_buffer_request.error());
        }
    }

    void tl_throwable_istream::ensure_byte() {
        if (ptr >= end_block) [[unlikely]] {
            grow_buffer();
            if (ptr == end_block) [[unlikely]] {
                throw tl_error(tl_error_type::STREAM_EOF, "eof");
            }
        }
    }

    void tl_throwable_istream::fetch_data(void *vdata, size_t size) {
        if (ptr + size > end_block) [[unlikely]] {
            return fetch_data2(vdata, size);
        }
        std::memcpy(reinterpret_cast<char *>(vdata), ptr, size);
        ptr += size;
    }

    void tl_throwable_istream::fetch_data2(void *vdata, size_t size) {
        char *data = reinterpret_cast<char *>(vdata);
        for (; ptr + size > end_block;) [[unlikely]] {
            std::memcpy(data, ptr, end_block - ptr);
            data += end_block - ptr;
            size -= end_block - ptr;
            grow_buffer();
            if (ptr == end_block) [[unlikely]] {
                throw tl_error(tl_error_type::STREAM_EOF, "eof");
            }
        }
        std::memcpy(data, ptr, size);
        ptr += size;
    }

    void tl_throwable_istream::fetch_data_append(std::string &value, size_t size) {
        for (;ptr + size > end_block;) [[unlikely]] {
            // assert(ptr <= end)
            value.append(reinterpret_cast<const char*>(ptr), end_block - ptr);
            size -= end_block - ptr;
            grow_buffer();
            // assert(ptr <= end)
            if (ptr == end_block) [[unlikely]] {
                throw tl_error(tl_error_type::STREAM_EOF, "eof");
            }
        }
        value.append(reinterpret_cast<const char*>(ptr), size);
        ptr += size;
    }

    void tl_throwable_istream::fetch_pad(size_t len) {
        uint32_t x = 0;
        fetch_data(&x, len);
        if (x != 0) [[unlikely]] {
            throw tl_error(tl_error_type::INCORRECT_STRING_PADDING, "incorrect string padding");
        }
    }

    void tl_throwable_istream::pass_data(tl_istream& to) noexcept {
        to.provider = provider;
        to.ptr = ptr;
        to.start_block = start_block;
        to.end_block = end_block;
    }

    tl_throwable_ostream::tl_throwable_ostream(tl_output_connector& provider) {
        this->provider = &provider;
    }

    void tl_throwable_ostream::string_write(const std::string &value) {
        auto len = value.size();
        if (len > TL_MAX_TINY_STRING_LEN) [[unlikely]] {
            if (len > TL_BIG_STRING_LEN) [[unlikely]] {
                throw tl_error (tl_error_type::INCORRECT_SEQUENCE_LENGTH, "TODO - huge string");
            }
            uint32_t p = (len << 8U) | TL_BIG_STRING_MARKER;
            store_data(&p, 4);
            store_data(value.data(), value.size());
            store_pad((-len) & 3);
            return;
        }
        auto pad = ((-(len + 1)) & 3);
        auto fullLen = 1 + len + pad;
        if (ptr + fullLen > end_block) [[unlikely]] {
            auto p = static_cast<unsigned char>(len);
            store_data(&p, 1);
            store_data(value.data(), value.size());
            store_pad(pad);
            return;
        }
        // fast path for short strings that fully fit in buffer
        uint32_t x = 0;
        std::memcpy(ptr + fullLen - 4, &x, 4); // padding first
        *ptr = static_cast<std::byte>(len);
        std::memcpy(ptr + 1, value.data(), len);
        ptr += fullLen;
   }

    void tl_throwable_ostream::last_release() noexcept {
        provider->release_buffer(ptr - start_block);
        start_block = ptr;
    }


    void tl_throwable_ostream::grow_buffer() {
        ptr = end_block;
        provider->release_buffer(ptr - start_block);
        auto new_buffer_request = provider->get_buffer();
        if (new_buffer_request) {
            auto new_buffer = new_buffer_request.value();
            ptr = new_buffer.data();
            start_block = ptr;
            end_block = ptr + new_buffer.size();
        } else {
            throw tl_connector_error(new_buffer_request.error());
        }
    }

    void tl_throwable_ostream::store_data(const void *vdata, size_t size) {
        if (ptr + size > end_block) [[unlikely]] {
            return store_data2(vdata, size);
        }
        std::memcpy(ptr, reinterpret_cast<const char *>(vdata), size);
        ptr += size;
    }

    void tl_throwable_ostream::store_data2(const void *vdata, size_t size) {
        const char *data = reinterpret_cast<const char *>(vdata);
        for (; ptr + size > end_block;) [[unlikely]] {
            std::memcpy(ptr, data, end_block - ptr);
            data += end_block - ptr;
            size -= end_block - ptr;
            grow_buffer();
            if (ptr == end_block) [[unlikely]] {
                throw tl_error(tl_error_type::STREAM_EOF, "eof");
            }
        }
        std::memcpy(ptr, data, size);
        ptr += size;
    }

    void tl_throwable_ostream::store_pad(size_t size) {
        for (; ptr + size > end_block;) [[unlikely]] {
            // assert(ptr <= end)
            std::memset(ptr, 0, end_block - ptr);
            size -= end_block - ptr;
            grow_buffer();
            // assert(ptr <= end)
            if (ptr == end_block) [[unlikely]] {
                throw tl_error(tl_error_type::STREAM_EOF, "eof");
            }
        }
        if (size != 0) {
            ptr[0] = static_cast<std::byte>(0);
            ptr[size - 1] = static_cast<std::byte>(0);
            ptr[size / 2] = static_cast<std::byte>(0);
            ptr += size;
        }
    }

    void tl_throwable_ostream::pass_data(tl_ostream& to) noexcept {
        to.provider = provider;
        to.ptr = ptr;
        to.start_block = start_block;
        to.end_block = end_block;
    }
} // namespace basictl