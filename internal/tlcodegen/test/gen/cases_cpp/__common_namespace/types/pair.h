// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"


namespace tlgen { 
template<typename X, typename Y>
struct Pair {
  X x{};
  Y y{};

  // tl type info
  static constexpr uint32_t TL_TAG = 0xf01604df;
  static constexpr std::string_view TL_NAME = "pair";

  uint32_t tl_tag() const { return 0xf01604df; }
  std::string_view tl_name() const { return "pair"; }
};

} // namespace tlgen

