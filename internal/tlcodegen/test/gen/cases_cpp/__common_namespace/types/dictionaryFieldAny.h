// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"


namespace tlgen { 
template<typename k, typename v>
struct DictionaryFieldAny {
  k key{};
  v value{};

  std::string_view tl_name() const { return "dictionaryFieldAny"; }
  uint32_t tl_tag() const { return 0x2c43a65b; }
};

} // namespace tlgen

