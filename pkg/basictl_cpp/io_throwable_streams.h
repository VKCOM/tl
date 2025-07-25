#ifndef BASICTL_CPP_IO_THROWABLE_STREAMS_H
#define BASICTL_CPP_IO_THROWABLE_STREAMS_H

/** TLGEN: CPP INCLUDES */
#include "dependencies.h"
#include "constants.h"
#include "io_connectors.h"
/** TLGEN: CPP INCLUDES END */

namespace basictl {
    class tl_istream;

    class tl_throwable_istream {
    public:
        explicit tl_throwable_istream(tl_input_connector &provider);

        tl_throwable_istream(const tl_throwable_istream &) = delete;

        tl_throwable_istream &operator=(const tl_throwable_istream &) = delete;

        tl_throwable_istream(tl_throwable_istream &&) = delete;

        tl_throwable_istream &operator=(tl_throwable_istream &&) = delete;

        ~tl_throwable_istream() { sync(); };

        friend class tl_istream;

        void pass_data(tl_istream &to) noexcept;

        uint32_t nat_read() {
            uint32_t result;
            nat_read(result);
            return result;
        }

        void nat_read(uint32_t &value) {
            if (ptr + basictl::TL_UINT32_SIZE > end_block) [[unlikely]] {
                return fetch_data2(&value, basictl::TL_UINT32_SIZE);
            }
            std::memcpy(reinterpret_cast<char *>(&value), ptr, basictl::TL_UINT32_SIZE);
            ptr += basictl::TL_UINT32_SIZE;
        }

        void nat_read_exact_tag(uint32_t tag) {
            uint32_t actual_tag = 0;
            nat_read(actual_tag);
            if (tag != actual_tag) [[unlikely]] {
                throw tl_stream_error(tl_error_type::UNEXPECTED_TAG, "unexpected tag");
            }
        }

        int32_t int_read() {
            int32_t result;
            int_read(result);
            return result;
        }

        void int_read(int32_t &value) {
            if (ptr + basictl::TL_INT32_SIZE > end_block) [[unlikely]] {
                return fetch_data2(&value, basictl::TL_INT32_SIZE);
            }
            std::memcpy(reinterpret_cast<char *>(&value), ptr, basictl::TL_INT32_SIZE);
            ptr += basictl::TL_INT32_SIZE;
        };

        int64_t long_read() {
            int64_t result;
            long_read(result);
            return result;
        }

        void long_read(int64_t &value) {
            if (ptr + basictl::TL_INT64_SIZE > end_block) [[unlikely]] {
                return fetch_data2(&value, basictl::TL_INT64_SIZE);
            }
            std::memcpy(reinterpret_cast<char *>(&value), ptr, basictl::TL_INT64_SIZE);
            ptr += basictl::TL_INT64_SIZE;
        };

        float float_read() {
            float result;
            float_read(result);
            return result;
        }

        void float_read(float &value) {
            if (ptr + basictl::TL_FLOAT32_SIZE > end_block) [[unlikely]] {
                return fetch_data2(&value, basictl::TL_FLOAT32_SIZE);
            }
            std::memcpy(reinterpret_cast<char *>(&value), ptr, basictl::TL_FLOAT32_SIZE);
            ptr += basictl::TL_FLOAT32_SIZE;
        };

        double double_read() {
            double result;
            double_read(result);
            return result;
        }

        void double_read(double &value) {
            if (ptr + basictl::TL_FLOAT64_SIZE > end_block) [[unlikely]] {
                return fetch_data2(&value, basictl::TL_FLOAT64_SIZE);
            }
            std::memcpy(reinterpret_cast<char *>(&value), ptr, basictl::TL_FLOAT64_SIZE);
            ptr += basictl::TL_FLOAT64_SIZE;
        }

        bool bool_read(uint32_t f, uint32_t t) {
            bool result;
            bool_read(result, f, t);
            return result;
        }

        void bool_read(bool &value, uint32_t f, uint32_t t) {
            uint32_t tag = 0;
            nat_read(tag);
            if (tag == t) {
                value = true;
                return;
            }
            if (tag != f) [[unlikely]] {
                throw tl_stream_error(tl_error_type::UNEXPECTED_TAG, "unexpected bool tag");
            }
            value = false;
        }

        void string_read(std::string &value);

        void sync() noexcept;

    private:
        tl_input_connector *provider;

        const std::byte *start_block{};
        const std::byte *ptr{};
        const std::byte *end_block{};

        void grow_buffer();

        void ensure_byte();

        void fetch_data(void *vdata, size_t size);

        void fetch_data2(void *vdata, size_t size);

        void fetch_data_append(std::string &value, size_t size);

        void fetch_pad(size_t len);
    };

    class tl_ostream;

    class tl_throwable_ostream { // TODO - prohibit copy/move
    public:
        explicit tl_throwable_ostream(tl_output_connector &provider);

        tl_throwable_ostream(const tl_throwable_ostream &) = delete;

        tl_throwable_ostream &operator=(const tl_throwable_ostream &) = delete;

        tl_throwable_ostream(tl_throwable_ostream &&) = delete;

        tl_throwable_ostream &operator=(tl_throwable_ostream &&) = delete;

        ~tl_throwable_ostream() { sync(); };

        friend class tl_ostream;

        void pass_data(tl_ostream &to) noexcept;

        void nat_write(uint32_t value) {
            if (ptr + basictl::TL_UINT32_SIZE > end_block) [[unlikely]] {
                return store_data2(&value, basictl::TL_UINT32_SIZE);
            }
            std::memcpy(ptr, reinterpret_cast<const char *>(&value), basictl::TL_UINT32_SIZE);
            ptr += basictl::TL_UINT32_SIZE;
        };

        void int_write(int32_t value) {
            if (ptr + basictl::TL_INT32_SIZE > end_block) [[unlikely]] {
                return store_data2(&value, basictl::TL_INT32_SIZE);
            }
            std::memcpy(ptr, reinterpret_cast<const char *>(&value), basictl::TL_INT32_SIZE);
            ptr += basictl::TL_INT32_SIZE;
        };

        void long_write(int64_t value) {
            if (ptr + basictl::TL_INT64_SIZE > end_block) [[unlikely]] {
                return store_data2(&value, basictl::TL_INT64_SIZE);
            }
            std::memcpy(ptr, reinterpret_cast<const char *>(&value), basictl::TL_INT64_SIZE);
            ptr += basictl::TL_INT64_SIZE;
        };

        void float_write(float value) {
            if (ptr + basictl::TL_FLOAT32_SIZE > end_block) [[unlikely]] {
                return store_data2(&value, basictl::TL_FLOAT32_SIZE);
            }
            std::memcpy(ptr, reinterpret_cast<const char *>(&value), basictl::TL_FLOAT32_SIZE);
            ptr += basictl::TL_FLOAT32_SIZE;
        };

        void double_write(double value) {
            if (ptr + basictl::TL_FLOAT64_SIZE > end_block) [[unlikely]] {
                return store_data2(&value, basictl::TL_FLOAT64_SIZE);
            }
            std::memcpy(ptr, reinterpret_cast<const char *>(&value), basictl::TL_FLOAT64_SIZE);
            ptr += basictl::TL_FLOAT64_SIZE;
        };

        void string_write(const std::string &value);

        void sync() noexcept;

    private:
        tl_output_connector *provider;

        std::byte *start_block{};
        std::byte *ptr{};
        std::byte *end_block{};

        void grow_buffer();

        void store_data(const void *vdata, size_t size);

        void store_data2(const void *vdata, size_t size);

        void store_pad(size_t size);
    };
}


#endif //BASICTL_CPP_IO_THROWABLE_STREAMS_H
