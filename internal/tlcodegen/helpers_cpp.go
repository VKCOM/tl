// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

const basicHPPTLCodeHeader = `%s
#pragma once

#include <array>
#include <cstring>
#include <memory>
#include <optional>
#include <ostream>
#include <stddef.h>
#include <stdexcept>
#include <string>
#include <utility>
#include <variant>
#include <vector>
#include <span>

namespace %s {
`

const basicHPPTLCodeFooter = `
} // namespace %s

`

const basicHPPTLCodeBody = `
    enum {
        TL_MAX_TINY_STRING_LEN = 253,
        TL_BIG_STRING_LEN = 0xffffff,
        TL_BIG_STRING_MARKER = 0xfe,
    };

    class tl_istream_interface {
    public:
        virtual ~tl_istream_interface() = default;

        virtual std::span<const std::byte> get_buffer() = 0;
        virtual void release_buffer(size_t size) = 0;
    };

    class tl_ostream_interface {
    public:
        virtual ~tl_ostream_interface() = default;

        virtual std::span<std::byte> get_buffer() = 0;
        virtual void release_buffer(size_t size) = 0;
    };

    class tl_istream {
    public:
        explicit tl_istream(tl_istream_interface* provider);
        tl_istream(const tl_istream&) = delete;
        tl_istream& operator=(const tl_istream&) = delete;

        tl_istream(tl_istream&&) = delete;
        tl_istream& operator=(tl_istream&&) = delete;

        ~tl_istream();

        bool nat_read(uint32_t& value);
        bool nat_read_exact_tag(uint32_t tag);
        bool int_read(int32_t& value);
        bool long_read(int64_t& value);
        bool float_read(float& value);
        bool double_read(double& value);
        bool bool_read(bool& value, uint32_t f, uint32_t t);
        bool string_read(std::string& value);

        void last_release();

        bool has_error();

        bool set_error(const char* e);
        bool set_error_eof();
        bool set_error_sequence_length();
        bool set_error_string_padding();
        bool set_error_bool_tag();
        bool set_error_expected_tag();
        bool set_error_union_tag();

    protected:
        tl_istream_interface* provider;

        bool hasError = false;
        const char* start_block{};
        const char* ptr{};
        const char* end_block{};
    private:
        void grow_buffer();

        bool ensure_byte();

        bool fetch_data(void* vdata, size_t size);
        bool fetch_data2(void* vdata, size_t size);
        bool fetch_data_append(std::string& value, size_t size);
        bool fetch_pad(size_t len);
    };

    class tl_ostream {
    public:
        explicit tl_ostream(tl_ostream_interface* provider);
        tl_ostream(const tl_ostream&) = delete;
        tl_ostream& operator=(const tl_ostream&) = delete;

        tl_ostream(tl_ostream&&) = delete;
        tl_ostream& operator=(tl_ostream&&) = delete;

        ~tl_ostream();

        bool nat_write(uint32_t value);
        bool int_write(int32_t value);
        bool long_write(int64_t value);
        bool float_write(float value);
        bool double_write(double value);
        bool string_write(const std::string& value);

        void last_release();

        bool has_error();
        bool set_error(const char* e);
        bool set_error_eof();
        bool set_error_sequence_length();
    protected:
        tl_ostream_interface* provider;

        bool hasError = false;
        char* start_block{};
        char* ptr{};
        char* end_block{};
    private:
        void grow_buffer();
        bool store_data(const void* vdata, size_t size);
        bool store_pad(size_t size);
    };
`
const basicCPPTLCodeHeader = `%s
#include "%s%s"

#define TLGEN2_UNLIKELY(x)                                                                         \
  (x)                   // __builtin_expect((x), 0) // could improve performance on your platform
#define TLGEN2_NOINLINE // __attribute__ ((noinline)) // could improve performance on your platform

namespace %s {
`

const basicCPPTLCodeBody = `
    tl_istream::tl_istream(tl_istream_interface *provider) : provider(provider) {}

    tl_istream::~tl_istream() {
        this->last_release();
    }

    bool tl_istream::nat_read(uint32_t &value) { return fetch_data(&value, 4); }

    bool tl_istream::nat_read_exact_tag(uint32_t tag) {
        uint32_t actual_tag = 0;
        if (TLGEN2_UNLIKELY(!nat_read(actual_tag))) {
            return false;
        }
        if (TLGEN2_UNLIKELY(tag != actual_tag)) {
            return set_error_expected_tag();
        }
        return true;
    }

    bool tl_istream::int_read(int32_t &value) { return fetch_data(&value, 4); }

    bool tl_istream::long_read(int64_t &value) { return fetch_data(&value, 8); }

    bool tl_istream::float_read(float &value) { return fetch_data(&value, 4); }

    bool tl_istream::double_read(double &value) { return fetch_data(&value, 8); }

    bool tl_istream::bool_read(bool &value, uint32_t f, uint32_t t) {
        uint32_t tag = 0;
        if (TLGEN2_UNLIKELY(!nat_read(tag))) {
            return false;
        }
        if (tag == t) {
            value = true;
            return true;
        }
        if (TLGEN2_UNLIKELY(tag != f)) {
            set_error_bool_tag();
        }
        value = false;
        return true;
    }

    bool tl_istream::string_read(std::string &value) {
        if (TLGEN2_UNLIKELY(!ensure_byte())) {
            return false;
        }
        auto len = size_t(static_cast<unsigned char>(*ptr));
        if (TLGEN2_UNLIKELY(len >= TL_BIG_STRING_MARKER)) {
            if (TLGEN2_UNLIKELY(len > TL_BIG_STRING_MARKER)) {
                return set_error("TODO - huge string");
            }
            uint32_t len32 = 0;
            if (TLGEN2_UNLIKELY(!nat_read(len32))) {
                return false;
            }
            len = len32 >> 8U;
            value.clear();
            if (TLGEN2_UNLIKELY(!fetch_data_append(value, len))) {
                return false;
            }
            if (TLGEN2_UNLIKELY(!fetch_pad((-len) & 3))) {
                return false;
            }
            return true;
        }
        auto pad = ((-(len + 1)) & 3);
        auto fullLen = 1 + len + pad;
        if (TLGEN2_UNLIKELY(ptr + fullLen > end_block)) {
            ptr += 1;
            value.clear();
            if (TLGEN2_UNLIKELY(!fetch_data_append(value, len))) {
                return false;
            }
            if (TLGEN2_UNLIKELY(!fetch_pad(pad))) {
                return false;
            }
            return true;
        }
        // fast path for short strings that fully fit in buffer
        uint32_t x = 0;
        std::memcpy(&x, ptr + fullLen - 4, 4);
        if (TLGEN2_UNLIKELY((x & ~(0xFFFFFFFFU >> (8 * pad))) != 0)) {
            return set_error_string_padding();
        }
        value.assign(ptr + 1, len);
        ptr += fullLen;
        return true;
    }

    void tl_istream::last_release() {
        provider->release_buffer(ptr - start_block);
        start_block = ptr;
    }

    bool tl_istream::has_error() {
        return hasError;
    }

    bool tl_istream::set_error(const char *e) {
        hasError = true;
        return false;
    } // TODO - set error field

    bool tl_istream::set_error_eof() { return set_error("EOF"); }

    bool tl_istream::set_error_sequence_length() { return set_error("sequence_length"); }

    bool tl_istream::set_error_string_padding() { return set_error("string_padding"); }

    bool tl_istream::set_error_bool_tag() { return set_error("bool_tag"); }

    bool tl_istream::set_error_expected_tag() { return set_error("expected_tag"); }

    bool tl_istream::set_error_union_tag() { return set_error("union_tag"); }

    void tl_istream::grow_buffer() {
        ptr = end_block;
        provider->release_buffer(ptr - start_block);
        auto new_buffer = provider->get_buffer();
        ptr = reinterpret_cast<const char *>(new_buffer.data());
        start_block = ptr;
        end_block = ptr + new_buffer.size();
    }

    bool tl_istream::ensure_byte() {
        if (TLGEN2_UNLIKELY(ptr >= end_block)) {
            // assert(ptr <= end)
            grow_buffer();
            // assert(ptr <= end)
            if (TLGEN2_UNLIKELY(ptr == end_block)) {
                return set_error_eof();
            }
        }
        return true;
    }

    bool tl_istream::fetch_data(void *vdata, size_t size) {
        char *data = reinterpret_cast<char *>(vdata);
        if (TLGEN2_UNLIKELY(ptr + size > end_block)) {
            return fetch_data2(vdata, size);
        }
        std::memcpy(data, ptr, size);
        ptr += size;
        return true;
    }

    bool tl_istream::fetch_data2(void *vdata, size_t size) {
        char *data = reinterpret_cast<char *>(vdata);
        for (; TLGEN2_UNLIKELY(ptr + size > end_block);) {
            // assert(ptr <= end)
            std::memcpy(data, ptr, end_block - ptr);
            data += end_block - ptr;
            size -= end_block - ptr;
            grow_buffer();
            // assert(ptr <= end)
            if (TLGEN2_UNLIKELY(ptr == end_block)) {
                return set_error_eof();
            }
        }
        std::memcpy(data, ptr, size);
        ptr += size;
        return true;
    }

    bool tl_istream::fetch_data_append(std::string &value, size_t size) {
        for (; TLGEN2_UNLIKELY(ptr + size > end_block);) {
            // assert(ptr <= end)
            value.append(ptr, end_block - ptr);
            size -= end_block - ptr;
            grow_buffer();
            // assert(ptr <= end)
            if (TLGEN2_UNLIKELY(ptr == end_block)) {
                return set_error_eof();
            }
        }
        value.append(ptr, size);
        ptr += size;
        return true;
    }

    bool tl_istream::fetch_pad(size_t len) {
        uint32_t x = 0;
        if (TLGEN2_UNLIKELY(!fetch_data(&x, len))) {
            return false;
        }
        if (TLGEN2_UNLIKELY(x != 0)) {
            return set_error_string_padding();
        }
        return true;
    }


    tl_ostream::tl_ostream(tl_ostream_interface *provider) {
        this->provider = provider;
    }

    tl_ostream::~tl_ostream() {
        this->last_release();
    }

    bool tl_ostream::nat_write(uint32_t value) { return store_data(&value, 4); }

    bool tl_ostream::int_write(int32_t value) { return store_data(&value, 4); }

    bool tl_ostream::long_write(int64_t value) { return store_data(&value, 8); }

    bool tl_ostream::float_write(float value) { return store_data(&value, 4); }

    bool tl_ostream::double_write(double value) { return store_data(&value, 8); }

    bool tl_ostream::string_write(const std::string &value) {
        auto len = value.size();
        if (TLGEN2_UNLIKELY(len > TL_MAX_TINY_STRING_LEN)) {
            if (TLGEN2_UNLIKELY(len > TL_BIG_STRING_LEN)) {
                return set_error("TODO - huge string");
            }
            uint32_t p = (len << 8U) | TL_BIG_STRING_MARKER;
            if (TLGEN2_UNLIKELY(!store_data(&p, 4))) {
                return false;
            }
            if (TLGEN2_UNLIKELY(!store_data(value.data(), value.size()))) {
                return false;
            }
            if (TLGEN2_UNLIKELY(!store_pad((-len) & 3))) {
                return false;
            }
            return true;
        }
        auto pad = ((-(len + 1)) & 3);
        auto fullLen = 1 + len + pad;
        if (TLGEN2_UNLIKELY(ptr + fullLen > end_block)) {
            auto p = static_cast<unsigned char>(len);
            if (TLGEN2_UNLIKELY(!store_data(&p, 1))) {
                return false;
            }
            if (TLGEN2_UNLIKELY(!store_data(value.data(), value.size()))) {
                return false;
            }
            if (TLGEN2_UNLIKELY(!store_pad(pad))) {
                return false;
            }
            return true;
        }
        // fast path for short strings that fully fit in buffer
        uint32_t x = 0;
        std::memcpy(ptr + fullLen - 4, &x, 4); // padding first
        *ptr = static_cast<char>(len);
        std::memcpy(ptr + 1, value.data(), len);
        ptr += fullLen;
        return true;
    }

    void tl_ostream::last_release() {
        provider->release_buffer(ptr - start_block);
        start_block = ptr;
    }

    bool tl_ostream::has_error() {
        return hasError;
    }

    bool tl_ostream::set_error(const char *e) {
        hasError = true;
        return false;
    } // TODO - set error field

    bool tl_ostream::set_error_eof() { return set_error("EOF"); }

    bool tl_ostream::set_error_sequence_length() { return set_error("sequence_length"); }

    void tl_ostream::grow_buffer() {
        ptr = end_block;
        provider->release_buffer(ptr - start_block);
        auto new_buffer = provider->get_buffer();
        ptr = reinterpret_cast<char *>(new_buffer.data());
        start_block = ptr;
        end_block = ptr + new_buffer.size();
    }

    bool tl_ostream::store_data(const void *vdata, size_t size) {
        const char *data = reinterpret_cast<const char *>(vdata);
        for (; TLGEN2_UNLIKELY(ptr + size > end_block);) {
            // assert(ptr <= end)
            std::memcpy(ptr, data, end_block - ptr);
            data += end_block - ptr;
            size -= end_block - ptr;
            grow_buffer();
            // assert(ptr <= end)
            if (TLGEN2_UNLIKELY(ptr == end_block)) {
                return set_error_eof();
            }
        }
        std::memcpy(ptr, data, size);
        ptr += size;
        return true;
    }

    bool tl_ostream::store_pad(size_t size) {
        for (; TLGEN2_UNLIKELY(ptr + size > end_block);) {
            // assert(ptr <= end)
            std::memset(ptr, 0, end_block - ptr);
            size -= end_block - ptr;
            grow_buffer();
            // assert(ptr <= end)
            if (TLGEN2_UNLIKELY(ptr == end_block)) {
                return set_error_eof();
            }
        }
        if (size != 0) {
            ptr[0] = 0;
            ptr[size - 1] = 0;
            ptr[size / 2] = 0;
            ptr += size;
        }
        return true;
    }
`

const basicCPPTLCodeFooter = `
} // namespace %s


#undef TLGEN2_NOINLINE
#undef TLGEN2_UNLIKELY
`

const basicTLStringsImplHPP = `%[1]s
#pragma once
#include "%[2]s%[3]s"

namespace %[2]s {
    class tl_istream_string : public tl_istream_interface {
    public:
        explicit tl_istream_string(const std::string & buffer) : buffer(buffer) {}

        std::span<const std::byte> get_buffer() override;
        void release_buffer(size_t size) override;
        
        std::span<const std::byte> used_buffer();
    private:
        const std::string & buffer;
        size_t used_size = 0;
    };

    class tl_ostream_string : public tl_ostream_interface {
    public:
        explicit tl_ostream_string(std::string & buffer) : buffer(buffer) {}

        std::span<std::byte> get_buffer() override;
        void release_buffer(size_t size) override;

        std::span<std::byte> used_buffer();
    private:
        std::string & buffer;
        size_t used_size = 0;
    };
};`

const basicTLStringsImplCPP = `%[1]s
#include "%[2]s%[3]s"
#include "string_io%[3]s"

namespace %[2]s {
    std::span<const std::byte> tl_istream_string::get_buffer() {
        return {reinterpret_cast<const std::byte*>(buffer.data()) + used_size, buffer.size() - used_size};
    }

    void tl_istream_string::release_buffer(size_t size) {
        used_size += size;
    }

    std::span<const std::byte> tl_istream_string::used_buffer() {
        return {reinterpret_cast<const std::byte*>(buffer.data()), used_size};
    }

    std::span<std::byte> tl_ostream_string::get_buffer() {
        return {reinterpret_cast<std::byte*>(buffer.data()) + used_size, buffer.size() - used_size};
    }

    void tl_ostream_string::release_buffer(size_t size) {
        used_size += size;
        if (used_size == buffer.size()) {
            buffer.resize(buffer.size() * 3 / 2 + 1024);
        }
    }

    std::span<std::byte> tl_ostream_string::used_buffer() {
        return {reinterpret_cast<std::byte*>(buffer.data()), used_size};
    }
}`
