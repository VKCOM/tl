// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"


namespace tlgen { namespace cyc2 { 
struct MyCycle;
}} // namespace tlgen::cyc2

namespace tlgen { namespace cyc1 { 
struct MyCycle {
  uint32_t fields_mask = 0;
  std::shared_ptr<::tlgen::cyc2::MyCycle> a{};

  ~MyCycle() {}

  // tl type info
  static constexpr uint32_t TL_TAG = 0x136ecc9e;
  static constexpr std::string_view TL_NAME = "cyc1.myCycle";

  uint32_t tl_tag() const { return 0x136ecc9e; }
  std::string_view tl_name() const { return "cyc1.myCycle"; }

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

  friend std::ostream& operator<<(std::ostream& s, const MyCycle& rhs) {
    rhs.write_json(s);
    return s;
  }
};

}} // namespace tlgen::cyc1

