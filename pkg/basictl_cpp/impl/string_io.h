#ifndef BASICTL_CPP_STRING_IO_H
#define BASICTL_CPP_STRING_IO_H

/** TLGEN: CPP INCLUDES */
#include "../io_connectors.h"
/** TLGEN: CPP INCLUDES END */

namespace basictl {
    class tl_istream_string : public tl_input_connector {
    public:
        explicit tl_istream_string(const std::string & buffer) : buffer(buffer) {}

        tl_connector_result<std::span<const std::byte>> get_buffer() noexcept override;
        void release_buffer(size_t size) noexcept override;

        std::span<const std::byte> used_buffer();
    private:
        const std::string & buffer;
        size_t used_size = 0;
    };

    class tl_ostream_string : public tl_output_connector {
    public:
        explicit tl_ostream_string(std::string & buffer) : buffer(buffer) {}

        tl_connector_result<std::span<std::byte>> get_buffer() noexcept override;
        void release_buffer(size_t size) noexcept override;

        std::span<std::byte> used_buffer();
    private:
        std::string & buffer;
        size_t used_size = 0;
    };
};

#endif //BASICTL_CPP_STRING_IO_H
