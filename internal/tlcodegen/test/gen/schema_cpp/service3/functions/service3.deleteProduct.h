// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"


namespace tlgen { namespace service3 { 
struct DeleteProduct {
  int32_t user_id = 0;
  int32_t type = 0;
  std::vector<int32_t> id;
  std::vector<int32_t> info;

  std::string_view tl_name() const { return "service3.deleteProduct"; }
  uint32_t tl_tag() const { return 0x6867e707; }

  bool write_json(std::ostream& s) const;

  bool read(::tlgen::basictl::tl_istream & s) noexcept;
  bool write(::tlgen::basictl::tl_ostream & s) const noexcept;

  void read(::tlgen::basictl::tl_throwable_istream & s);
  void write(::tlgen::basictl::tl_throwable_ostream & s) const;

  bool read_boxed(::tlgen::basictl::tl_istream & s) noexcept;
  bool write_boxed(::tlgen::basictl::tl_ostream & s)const noexcept;
  
  void read_boxed(::tlgen::basictl::tl_throwable_istream & s);
  void write_boxed(::tlgen::basictl::tl_throwable_ostream & s)const;

  bool read_result(::tlgen::basictl::tl_istream & s, bool & result) noexcept;
  bool write_result(::tlgen::basictl::tl_ostream & s, bool & result) noexcept;

  void read_result(::tlgen::basictl::tl_throwable_istream & s, bool & result);
  void write_result(::tlgen::basictl::tl_throwable_ostream & s, bool & result);

  friend std::ostream& operator<<(std::ostream& s, const DeleteProduct& rhs) {
    rhs.write_json(s);
    return s;
  }
};

}} // namespace tlgen::service3

