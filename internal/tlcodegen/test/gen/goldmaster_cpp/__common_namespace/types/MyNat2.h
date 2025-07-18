// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/myPlus.h"
#include "__common_namespace/types/myZero.h"


namespace tlgen { 
struct MyPlus;
} // namespace tlgen

namespace tlgen { 
struct MyNat2 {
  std::variant<::tlgen::MyZero, ::tlgen::MyPlus> value;

  bool is_myZero() const { return value.index() == 0; }
  bool is_myPlus() const { return value.index() == 1; }

  void set_myZero() { value.emplace<0>(); }

  std::string_view tl_name() const;
  uint32_t tl_tag() const;

  bool write_json(std::ostream& s)const;

  bool read_boxed(::tlgen::basictl::tl_istream & s) noexcept;
  bool write_boxed(::tlgen::basictl::tl_ostream & s)const noexcept;
  
  void read_boxed(::tlgen::basictl::tl_throwable_istream & s);
  void write_boxed(::tlgen::basictl::tl_throwable_ostream & s)const;
};

} // namespace tlgen

