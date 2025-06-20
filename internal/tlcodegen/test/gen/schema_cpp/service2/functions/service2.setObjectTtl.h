// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service2/types/service2.objectId.h"
#include "__common_namespace/types/true.h"


namespace tlgen { namespace service2 { 
struct SetObjectTtl {
  uint32_t objectIdLength = 0;
  ::tlgen::service2::ObjectId objectId{};
  int32_t ttl = 0;

  std::string_view tl_name() const { return "service2.setObjectTtl"; }
  uint32_t tl_tag() const { return 0x6f98f025; }

  bool write_json(std::ostream& s) const;

  bool read(::tlgen::basictl::tl_istream & s) noexcept;
  bool write(::tlgen::basictl::tl_ostream & s) const noexcept;

  void read(::tlgen::basictl::tl_throwable_istream & s);
  void write(::tlgen::basictl::tl_throwable_ostream & s) const;

  bool read_boxed(::tlgen::basictl::tl_istream & s) noexcept;
  bool write_boxed(::tlgen::basictl::tl_ostream & s)const noexcept;
  
  void read_boxed(::tlgen::basictl::tl_throwable_istream & s);
  void write_boxed(::tlgen::basictl::tl_throwable_ostream & s)const;

  bool read_result(::tlgen::basictl::tl_istream & s, ::tlgen::True & result) noexcept;
  bool write_result(::tlgen::basictl::tl_ostream & s, ::tlgen::True & result) noexcept;

  void read_result(::tlgen::basictl::tl_throwable_istream & s, ::tlgen::True & result);
  void write_result(::tlgen::basictl::tl_throwable_ostream & s, ::tlgen::True & result);

  friend std::ostream& operator<<(std::ostream& s, const SetObjectTtl& rhs) {
    rhs.write_json(s);
    return s;
  }
};

}} // namespace tlgen::service2

