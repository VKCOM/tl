#ifndef BASICTL_CPP_ERRORS_H
#define BASICTL_CPP_ERRORS_H

#include <exception>
#include <string_view>
#include <cstring>
#include <string>

namespace basictl {
    enum class tl_error_type {
        STREAM_EOF,
        INCORRECT_SEQUENCE_LENGTH,
        INCORRECT_STRING_PADDING,
        UNEXPECTED_TAG,
    };

    template <typename Type>
    class basic_error final : public std::exception
    {
    public:
        basic_error(Type type, std::string message)
                : type_(type)
                , message_(std::move(message))
        {}

        [[nodiscard]] Type type() const noexcept {
            return type_;
        }

        [[nodiscard]] std::string_view message() const noexcept {
            return message_;
        }

        [[nodiscard]] const char * what() const noexcept override {
            return message_.c_str();
        }

    private:
        Type type_;
        std::string message_;
    };

    using tl_error = basic_error<tl_error_type>;
    using tl_connector_error = basic_error<std::uint32_t>;
    using tl_stream_error = std::variant<tl_error, tl_connector_error>;

    std::exception static exception_from_tl_stream_error(tl_stream_error & error) {
        switch (error.index()) {
            case 0: return std::get<0>(error);
            case 1: return std::get<1>(error);
            default: return {};
        }
    }
}

#endif //BASICTL_CPP_ERRORS_H
