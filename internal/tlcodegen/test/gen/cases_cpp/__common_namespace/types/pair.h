// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"


namespace tlgen { 
template<typename X, typename Y>
struct Pair {
  X x{};
  Y y{};

  std::string_view tl_name() const { return "pair"; }
  uint32_t tl_tag() const { return 0xf01604df; }
};

} // namespace tlgen

