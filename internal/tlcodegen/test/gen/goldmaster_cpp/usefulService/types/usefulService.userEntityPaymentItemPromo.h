// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"


namespace tlgen { namespace usefulService { 
struct UserEntityPaymentItemPromo {
  std::string content;

  std::string_view tl_name() const { return "usefulService.userEntityPaymentItemPromo"; }
  uint32_t tl_tag() const { return 0x24c7ec9f; }

  bool write_json(std::ostream& s, [[maybe_unused]] uint32_t nat_fields_mask) const;

  bool read(::tlgen::basictl::tl_istream & s, [[maybe_unused]] uint32_t nat_fields_mask) noexcept;
  bool write(::tlgen::basictl::tl_ostream & s, [[maybe_unused]] uint32_t nat_fields_mask) const noexcept;

  void read(::tlgen::basictl::tl_throwable_istream & s, [[maybe_unused]] uint32_t nat_fields_mask);
  void write(::tlgen::basictl::tl_throwable_ostream & s, [[maybe_unused]] uint32_t nat_fields_mask) const;

  bool read_boxed(::tlgen::basictl::tl_istream & s, [[maybe_unused]] uint32_t nat_fields_mask) noexcept;
  bool write_boxed(::tlgen::basictl::tl_ostream & s, [[maybe_unused]] uint32_t nat_fields_mask)const noexcept;
  
  void read_boxed(::tlgen::basictl::tl_throwable_istream & s, [[maybe_unused]] uint32_t nat_fields_mask);
  void write_boxed(::tlgen::basictl::tl_throwable_ostream & s, [[maybe_unused]] uint32_t nat_fields_mask)const;
};

}} // namespace tlgen::usefulService

