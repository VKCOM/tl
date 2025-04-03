#ifndef BASICTL_CPP_IO_STREAMS_H
#define BASICTL_CPP_IO_STREAMS_H

#include <array>
#include <optional>
#include <ostream>
#include <memory>
#include <variant>
#include <vector>

#include "constants.h"
#include "io_connectors.h"

namespace basictl {

    class tl_istream {
    public:
        explicit tl_istream(tl_input_connector& provider);
        tl_istream(const tl_istream&) = delete;
        tl_istream& operator=(const tl_istream&) = delete;

        tl_istream(tl_istream&&) = delete;
        tl_istream& operator=(tl_istream&&) = delete;

        ~tl_istream() { last_release(); };

        bool nat_read(uint32_t& value) noexcept {
            if (ptr + TL_UINT32_SIZE > end_block) [[unlikely]] {
                return fetch_data2(&value, TL_UINT32_SIZE);
            }
            std::memcpy(reinterpret_cast<char *>(&value), ptr, TL_UINT32_SIZE);
            ptr += TL_UINT32_SIZE;
            return true;
        }
        bool nat_read_exact_tag(uint32_t tag) noexcept {
            uint32_t actual_tag = 0;
            if (!nat_read(actual_tag)) [[unlikely]] {
                return false;
            }
            if (tag != actual_tag) [[unlikely]] {
                return set_error_expected_tag();
            }
            return true;
        }
        bool int_read(int32_t& value) noexcept {
            if (ptr + TL_INT32_SIZE > end_block) [[unlikely]] {
                return fetch_data2(&value, TL_INT32_SIZE);
            }
            std::memcpy(reinterpret_cast<char *>(&value), ptr, TL_INT32_SIZE);
            ptr += TL_INT32_SIZE;
            return true;
        };
        bool long_read(int64_t& value) noexcept {
            if (ptr + TL_INT64_SIZE > end_block) [[unlikely]] {
                return fetch_data2(&value, TL_INT64_SIZE);
            }
            std::memcpy(reinterpret_cast<char *>(&value), ptr, TL_INT64_SIZE);
            ptr += TL_INT64_SIZE;
            return true;
        };
        bool float_read(float& value) noexcept {
            if (ptr + TL_FLOAT32_SIZE > end_block) [[unlikely]] {
                return fetch_data2(&value, TL_FLOAT32_SIZE);
            }
            std::memcpy(reinterpret_cast<char *>(&value), ptr, TL_FLOAT32_SIZE);
            ptr += TL_FLOAT32_SIZE;
            return true;
        };
        bool double_read(double& value) noexcept {
            if (ptr + TL_FLOAT64_SIZE > end_block) [[unlikely]] {
                return fetch_data2(&value, TL_FLOAT64_SIZE);
            }
            std::memcpy(reinterpret_cast<char *>(&value), ptr, TL_FLOAT64_SIZE);
            ptr += TL_FLOAT64_SIZE;
            return true;
        }
        bool bool_read(bool &value, uint32_t f, uint32_t t) noexcept {
            uint32_t tag = 0;
            if (!nat_read(tag)) [[unlikely]] {
                return false;
            }
            if (tag == t) {
                value = true;
                return true;
            }
            if (tag != f) [[unlikely]] {
                set_error(tl_error_type::UNEXPECTED_TAG, "unexpected bool tag");
            }
            value = false;
            return true;
        }

        bool string_read(std::string& value) noexcept;

        void last_release() noexcept;

        [[nodiscard]] bool has_error() const;
        [[nodiscard]] std::optional<tl_stream_error>& get_error();

        bool set_error(tl_error_type type, const char* what) noexcept;
        bool set_error_eof() noexcept;
        bool set_error_sequence_length() noexcept;
        bool set_error_string_padding() noexcept;
        bool set_error_expected_tag() noexcept;
    private:
        tl_input_connector* provider;
        std::optional<tl_stream_error> error;

        const std::byte* start_block{};
        const std::byte* ptr{};
        const std::byte* end_block{};

        void grow_buffer() noexcept;

        bool ensure_byte() noexcept;

        bool fetch_data(void* vdata, size_t size) noexcept;
        bool fetch_data2(void* vdata, size_t size) noexcept;
        bool fetch_data_append(std::string& value, size_t size) noexcept;
        bool fetch_pad(size_t len) noexcept;
    };

    class tl_ostream {
    public:
        explicit tl_ostream(tl_output_connector* provider);
        tl_ostream(const tl_ostream&) = delete;
        tl_ostream& operator=(const tl_ostream&) = delete;

        tl_ostream(tl_ostream&&) = delete;
        tl_ostream& operator=(tl_ostream&&) = delete;

        ~tl_ostream() { last_release(); };

        bool nat_write(uint32_t value) {
            if (ptr + TL_UINT32_SIZE > end_block) [[unlikely]] {
                return store_data2(&value, TL_UINT32_SIZE);
            }
            std::memcpy(ptr, reinterpret_cast<const char *>(&value), TL_UINT32_SIZE);
            ptr += TL_UINT32_SIZE;
            return true;
        };
        bool int_write(int32_t value) {
            if (ptr + TL_INT32_SIZE > end_block) [[unlikely]] {
                return store_data2(&value, TL_INT32_SIZE);
            }
            std::memcpy(ptr, reinterpret_cast<const char *>(&value), TL_INT32_SIZE);
            ptr += TL_INT32_SIZE;
            return true;
        };
        bool long_write(int64_t value) {
            if (ptr + TL_INT64_SIZE > end_block) [[unlikely]] {
                return store_data2(&value, TL_INT64_SIZE);
            }
            std::memcpy(ptr, reinterpret_cast<const char *>(&value), TL_INT64_SIZE);
            ptr += TL_INT64_SIZE;
            return true;
        };
        bool float_write(float value) {
            if (ptr + TL_FLOAT32_SIZE > end_block) [[unlikely]] {
                return store_data2(&value, TL_FLOAT32_SIZE);
            }
            std::memcpy(ptr, reinterpret_cast<const char *>(&value), TL_FLOAT32_SIZE);
            ptr += TL_FLOAT32_SIZE;
            return true;
        };
        bool double_write(double value) {
            if (ptr + TL_FLOAT64_SIZE > end_block) [[unlikely]] {
                return store_data2(&value, TL_FLOAT64_SIZE);
            }
            std::memcpy(ptr, reinterpret_cast<const char *>(&value), TL_FLOAT64_SIZE);
            ptr += TL_FLOAT64_SIZE;
            return true;
        };
        bool string_write(const std::string& value);

        void last_release() noexcept;

        [[nodiscard]] bool has_error() const;
        [[nodiscard]] std::optional<tl_stream_error>& get_error();

        bool set_error(tl_error_type type, const char* e);
        bool set_error_eof();
        bool set_error_sequence_length();
    private:
        tl_output_connector* provider;
        std::optional<tl_stream_error> error;

        std::byte* start_block{};
        std::byte* ptr{};
        std::byte* end_block{};

        void grow_buffer();
        bool store_data(const void* vdata, size_t size);
        bool store_data2(const void* vdata, size_t size);
        bool store_pad(size_t size);
    };
}

#endif //BASICTL_CPP_IO_STREAMS_H
