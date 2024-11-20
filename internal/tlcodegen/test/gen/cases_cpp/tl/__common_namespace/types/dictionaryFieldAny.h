#pragma once

#include "../../../basics/basictl.h"


namespace tl2 { 
template<typename k, typename v>
struct DictionaryFieldAny {
	k key{};
	v value{};

	std::string_view tl_name() const { return "dictionaryFieldAny"; }
	uint32_t tl_tag() const { return 0x2c43a65b; }
};

} // namespace tl2

