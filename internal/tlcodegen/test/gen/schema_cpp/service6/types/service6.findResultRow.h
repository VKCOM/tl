// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"


namespace tlgen { namespace service6 { 
struct FindResultRow {
  int32_t x = 0;

  // tl type info
  static constexpr uint32_t TL_TAG = 0xbd3946e3;
  static constexpr std::string_view TL_NAME = "service6.findResultRow";

  uint32_t tl_tag() const { return 0xbd3946e3; }
  std::string_view tl_name() const { return "service6.findResultRow"; }

  // basic serialization methods 
  bool write_json(std::ostream& s) const;

  bool read(::tlgen::basictl::tl_istream & s) noexcept;
  bool write(::tlgen::basictl::tl_ostream & s) const noexcept;

  void read(::tlgen::basictl::tl_throwable_istream & s);
  void write(::tlgen::basictl::tl_throwable_ostream & s) const;

  bool read_boxed(::tlgen::basictl::tl_istream & s) noexcept;
  bool write_boxed(::tlgen::basictl::tl_ostream & s) const noexcept;
  
  void read_boxed(::tlgen::basictl::tl_throwable_istream & s);
  void write_boxed(::tlgen::basictl::tl_throwable_ostream & s) const;

  friend std::ostream& operator<<(std::ostream& s, const FindResultRow& rhs) {
    rhs.write_json(s);
    return s;
  }
};

}} // namespace tlgen::service6

