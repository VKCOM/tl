
#pragma once

#include <string>
#include <functional>

#include "../a_tlgen_helpers_code.hpp"

namespace tl2 {
    namespace meta {
        struct tl_object {
            virtual bool read(::basictl::tl_istream &s) = 0;
            virtual bool write(::basictl::tl_ostream &s) = 0;

            virtual bool read_boxed(::basictl::tl_istream &s) = 0;
            virtual bool write_boxed(::basictl::tl_ostream &s) = 0;

            virtual ~tl_object() = default;
        };

        struct tl_function : public tl_object {
            virtual bool read_write_result(::basictl::tl_istream &in, ::basictl::tl_ostream &out) = 0;

            virtual ~tl_function() = default;
        };

        struct tl_item {
            uint32_t tag{};
            uint32_t annotations{};
            std::string name;

            std::function<std::unique_ptr<tl2::meta::tl_object>()> create_object;
            std::function<std::unique_ptr<tl2::meta::tl_function>()> create_function;
        };

		tl_item get_item_by_name(std::string &&s);

		void set_create_object_by_name(std::string &&s, std::function<std::unique_ptr<tl2::meta::tl_object>()> &&factory);
		void set_create_function_by_name(std::string &&s, std::function<std::unique_ptr<tl2::meta::tl_function>()> &&factory);
        
    }
}