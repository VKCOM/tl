// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"


namespace tlgen { 
struct Replace7 {
  uint32_t n = 0;
  uint32_t m = 0;
  std::vector<std::vector<int32_t>> a;

  // tl type info
  static constexpr uint32_t TL_TAG = 0xf4c66d9f;
  static constexpr std::string_view TL_NAME = "replace7";

  uint32_t tl_tag() const { return 0xf4c66d9f; }
  std::string_view tl_name() const { return "replace7"; }

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

  friend std::ostream& operator<<(std::ostream& s, const Replace7& rhs) {
    rhs.write_json(s);
    return s;
  }
};

} // namespace tlgen

