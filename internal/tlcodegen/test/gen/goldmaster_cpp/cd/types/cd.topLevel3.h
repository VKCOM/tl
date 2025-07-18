// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "__common_namespace/types/halfStr.h"
#include "__common_namespace/types/useStr.h"


namespace tlgen { namespace cd { 
struct TopLevel3 {
  ::tlgen::UseStr a{};
  ::tlgen::HalfStr b{};

  // tl type info
  static constexpr uint32_t TL_TAG = 0x5cd1ca89;
  static constexpr std::string_view TL_NAME = "cd.topLevel3";

  uint32_t tl_tag() const { return 0x5cd1ca89; }
  std::string_view tl_name() const { return "cd.topLevel3"; }

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

  friend std::ostream& operator<<(std::ostream& s, const TopLevel3& rhs) {
    rhs.write_json(s);
    return s;
  }
};

}} // namespace tlgen::cd

