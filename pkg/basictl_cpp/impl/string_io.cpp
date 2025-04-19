/** TLGEN: CPP INCLUDES */
#include "string_io.h"
/** TLGEN: CPP INCLUDES END */

namespace basictl {
    tl_connector_result<std::span<const std::byte>> tl_istream_string::get_buffer() noexcept {
        return tl_connector_result(std::span<const std::byte>{reinterpret_cast<const std::byte*>(buffer.data()) + used_size, buffer.size() - used_size});
    }

    void tl_istream_string::release_buffer(size_t size) noexcept {
        used_size += size;
    }

    std::span<const std::byte> tl_istream_string::used_buffer() {
        return {reinterpret_cast<const std::byte*>(buffer.data()), used_size};
    }

    tl_connector_result<std::span<std::byte>> tl_ostream_string::get_buffer() noexcept {
        return tl_connector_result(std::span<std::byte>{reinterpret_cast<std::byte*>(buffer.data()) + used_size, buffer.size() - used_size});
    }

    void tl_ostream_string::release_buffer(size_t size) noexcept {
        used_size += size;
        if (used_size == buffer.size()) {
            buffer.resize(buffer.size() * 3 / 2 + 1024);
        }
    }

    std::span<std::byte> tl_ostream_string::used_buffer() {
        return {reinterpret_cast<std::byte*>(buffer.data()), used_size};
    }
}