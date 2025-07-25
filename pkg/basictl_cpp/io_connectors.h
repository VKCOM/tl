#ifndef BASIC_CPP_IO_CONNECTORS_H
#define BASIC_CPP_IO_CONNECTORS_H

#include <variant>
#include <span>

/** TLGEN: CPP INCLUDES */
#include "errors.h"
/** TLGEN: CPP INCLUDES END */

namespace basictl {
    template<typename Type>
    class tl_connector_result {
    public:

        explicit tl_connector_result(Type buffer)
                : result_(buffer) {}

        explicit tl_connector_result(basictl::tl_connector_error error)
                : result_(std::move(error)) {}

        explicit operator bool() const noexcept {
            return std::holds_alternative<Type>(result_);
        }

        bool has_value() const noexcept {
            return std::holds_alternative<Type>(result_);
        }

        Type value() const noexcept {
            return std::get<Type>(result_);
        }

        [[nodiscard]] const tl_connector_error &error() const & noexcept {
            return std::get<tl_connector_error>(result_);
        }

        tl_connector_error error() && noexcept {
            return std::get<tl_connector_error>(std::move(result_));
        }

    private:
        std::variant<Type, basictl::tl_connector_error> result_;
    };

    class tl_input_connector {
    public:
        virtual ~tl_input_connector() = default;

        virtual tl_connector_result<std::span<const std::byte>> get_buffer() noexcept = 0;

        virtual void advance(size_t size) noexcept = 0;
    };

    class tl_output_connector {
    public:
        virtual ~tl_output_connector() = default;

        virtual tl_connector_result<std::span<std::byte>> get_buffer() noexcept = 0;

        virtual void advance(size_t size) noexcept = 0;
    };
}

#endif //BASIC_CPP_IO_CONNECTORS_H
