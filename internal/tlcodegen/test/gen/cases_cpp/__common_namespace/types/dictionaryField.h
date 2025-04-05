#pragma once

#include "../../basictl/io_streams.h"


namespace tl2 { 
template<typename t>
struct DictionaryField {
	std::string key;
	t value{};

	std::string_view tl_name() const { return "dictionaryField"; }
	uint32_t tl_tag() const { return 0x239c1b62; }
};

} // namespace tl2

