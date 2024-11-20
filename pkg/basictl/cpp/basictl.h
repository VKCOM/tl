// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
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

namespace basictl {

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

} // namespace basictl

