// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"


namespace tlgen { namespace service3 { 
struct Product {
  int32_t type = 0;
  std::vector<int32_t> id;
  std::vector<int32_t> info;
  int32_t date = 0;
  int32_t expiration_date = 0;
  bool removed = false;

  // tl type info
  static constexpr uint32_t TL_TAG = 0x461f4ce2;
  static constexpr std::string_view TL_NAME = "service3.product";

  uint32_t tl_tag() const { return 0x461f4ce2; }
  std::string_view tl_name() const { return "service3.product"; }

  // basic serialization methods 
  bool write_json(std::ostream& s, [[maybe_unused]] uint32_t nat_mode) const;

  bool read(::tlgen::basictl::tl_istream & s, [[maybe_unused]] uint32_t nat_mode) noexcept;
  bool write(::tlgen::basictl::tl_ostream & s, [[maybe_unused]] uint32_t nat_mode) const noexcept;

  void read(::tlgen::basictl::tl_throwable_istream & s, [[maybe_unused]] uint32_t nat_mode);
  void write(::tlgen::basictl::tl_throwable_ostream & s, [[maybe_unused]] uint32_t nat_mode) const;

  bool read_boxed(::tlgen::basictl::tl_istream & s, [[maybe_unused]] uint32_t nat_mode) noexcept;
  bool write_boxed(::tlgen::basictl::tl_ostream & s, [[maybe_unused]] uint32_t nat_mode) const noexcept;
  
  void read_boxed(::tlgen::basictl::tl_throwable_istream & s, [[maybe_unused]] uint32_t nat_mode);
  void write_boxed(::tlgen::basictl::tl_throwable_ostream & s, [[maybe_unused]] uint32_t nat_mode) const;
};

}} // namespace tlgen::service3

namespace tlgen { namespace service3 { 
template<uint32_t mode>
struct Productmode {
  int32_t type = 0;
  std::vector<int32_t> id;
  std::vector<int32_t> info;
  int32_t date = 0;
  int32_t expiration_date = 0;
  bool removed = false;

  // tl type info
  static constexpr uint32_t TL_TAG = 0x461f4ce2;
  static constexpr std::string_view TL_NAME = "service3.product";

  uint32_t tl_tag() const { return 0x461f4ce2; }
  std::string_view tl_name() const { return "service3.product"; }
};

}} // namespace tlgen::service3

