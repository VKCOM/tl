// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "casesTL2/types/casesTL2.testObject.h"


namespace tlgen { namespace casesTL2 { 
struct TestFunctionNoDep5 {
  int32_t x = 0;

  // tl type info
  static constexpr uint32_t TL_TAG = 0x2b47b925;
  static constexpr std::string_view TL_NAME = "casesTL2.testFunctionNoDep5";

  uint32_t tl_tag() const { return 0x2b47b925; }
  std::string_view tl_name() const { return "casesTL2.testFunctionNoDep5"; }

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

  // function methods and properties
  using ResultType = ::tlgen::casesTL2::TestObject;

  bool read_result(::tlgen::basictl::tl_istream & s, ::tlgen::casesTL2::TestObject & result) const noexcept;
  bool write_result(::tlgen::basictl::tl_ostream & s, const ::tlgen::casesTL2::TestObject & result) const noexcept;

  void read_result(::tlgen::basictl::tl_throwable_istream & s, ::tlgen::casesTL2::TestObject & result) const;
  void write_result(::tlgen::basictl::tl_throwable_ostream & s, const ::tlgen::casesTL2::TestObject & result) const;

  friend std::ostream& operator<<(std::ostream& s, const TestFunctionNoDep5& rhs) {
    rhs.write_json(s);
    return s;
  }
};

}} // namespace tlgen::casesTL2

