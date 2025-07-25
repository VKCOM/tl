// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#include "service5/headers/service5.query.h"
#include "service5/headers/service5.performQuery.h"
#include "service5/headers/service5.params.h"
#include "service5/headers/service5.Output.h"
#include "service5/headers/service5.stringOutput.h"
#include "service5/headers/service5.insert.h"
#include "service5/headers/service5.emptyOutput.h"


bool tlgen::service5::EmptyOutput::write_json(std::ostream& s)const {
  if (!::tlgen::details::Service5EmptyOutputWriteJSON(s, *this)) { return false; }
  return true;
}

bool tlgen::service5::EmptyOutput::read(::tlgen::basictl::tl_istream & s) noexcept {
  if (!::tlgen::details::Service5EmptyOutputRead(s, *this)) { return s.set_error_unknown_scenario(); }
  s.sync();
  return true;
}

bool tlgen::service5::EmptyOutput::write(::tlgen::basictl::tl_ostream & s) const noexcept {
  if (!::tlgen::details::Service5EmptyOutputWrite(s, *this)) { return s.set_error_unknown_scenario(); }
  s.sync();
  return true;
}

void tlgen::service5::EmptyOutput::read(::tlgen::basictl::tl_throwable_istream & s) {
  ::tlgen::basictl::tl_istream s2(s);
  this->read(s2);
  s2.pass_data(s);
}

void tlgen::service5::EmptyOutput::write(::tlgen::basictl::tl_throwable_ostream & s) const {
  ::tlgen::basictl::tl_ostream s2(s);
  this->write(s2);
  s2.pass_data(s);
}

bool tlgen::service5::EmptyOutput::read_boxed(::tlgen::basictl::tl_istream & s) noexcept {
  if (!::tlgen::details::Service5EmptyOutputReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
  s.sync();
  return true;
}

bool tlgen::service5::EmptyOutput::write_boxed(::tlgen::basictl::tl_ostream & s) const noexcept {
  if (!::tlgen::details::Service5EmptyOutputWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
  s.sync();
  return true;
}

void tlgen::service5::EmptyOutput::read_boxed(::tlgen::basictl::tl_throwable_istream & s) {
  ::tlgen::basictl::tl_istream s2(s);
  this->read_boxed(s2);
  s2.pass_data(s);
}

void tlgen::service5::EmptyOutput::write_boxed(::tlgen::basictl::tl_throwable_ostream & s) const {
  ::tlgen::basictl::tl_ostream s2(s);
  this->write_boxed(s2);
  s2.pass_data(s);
}

void tlgen::details::Service5EmptyOutputReset(::tlgen::service5::EmptyOutput& item) noexcept {
  (void)item;
}

bool tlgen::details::Service5EmptyOutputWriteJSON(std::ostream& s, const ::tlgen::service5::EmptyOutput& item) noexcept {
  (void)s;
  (void)item;
  s << "true";
  return true;
}

bool tlgen::details::Service5EmptyOutputRead(::tlgen::basictl::tl_istream & s, ::tlgen::service5::EmptyOutput& item) noexcept {
  (void)s;
  (void)item;
  return true;
}

bool tlgen::details::Service5EmptyOutputWrite(::tlgen::basictl::tl_ostream & s, const ::tlgen::service5::EmptyOutput& item) noexcept {
  (void)s;
  (void)item;
  return true;
}

bool tlgen::details::Service5EmptyOutputReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::service5::EmptyOutput& item) {
  if (!s.nat_read_exact_tag(0x11e46879)) { return false; }
  return tlgen::details::Service5EmptyOutputRead(s, item);
}

bool tlgen::details::Service5EmptyOutputWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::service5::EmptyOutput& item) {
  if (!s.nat_write(0x11e46879)) { return false; }
  return tlgen::details::Service5EmptyOutputWrite(s, item);
}

bool tlgen::service5::Insert::write_json(std::ostream& s)const {
  if (!::tlgen::details::Service5InsertWriteJSON(s, *this)) { return false; }
  return true;
}

bool tlgen::service5::Insert::read(::tlgen::basictl::tl_istream & s) noexcept {
  if (!::tlgen::details::Service5InsertRead(s, *this)) { return s.set_error_unknown_scenario(); }
  s.sync();
  return true;
}

bool tlgen::service5::Insert::write(::tlgen::basictl::tl_ostream & s) const noexcept {
  if (!::tlgen::details::Service5InsertWrite(s, *this)) { return s.set_error_unknown_scenario(); }
  s.sync();
  return true;
}

void tlgen::service5::Insert::read(::tlgen::basictl::tl_throwable_istream & s) {
  ::tlgen::basictl::tl_istream s2(s);
  this->read(s2);
  s2.pass_data(s);
}

void tlgen::service5::Insert::write(::tlgen::basictl::tl_throwable_ostream & s) const {
  ::tlgen::basictl::tl_ostream s2(s);
  this->write(s2);
  s2.pass_data(s);
}

bool tlgen::service5::Insert::read_boxed(::tlgen::basictl::tl_istream & s) noexcept {
  if (!::tlgen::details::Service5InsertReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
  s.sync();
  return true;
}

bool tlgen::service5::Insert::write_boxed(::tlgen::basictl::tl_ostream & s) const noexcept {
  if (!::tlgen::details::Service5InsertWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
  s.sync();
  return true;
}

void tlgen::service5::Insert::read_boxed(::tlgen::basictl::tl_throwable_istream & s) {
  ::tlgen::basictl::tl_istream s2(s);
  this->read_boxed(s2);
  s2.pass_data(s);
}

void tlgen::service5::Insert::write_boxed(::tlgen::basictl::tl_throwable_ostream & s) const {
  ::tlgen::basictl::tl_ostream s2(s);
  this->write_boxed(s2);
  s2.pass_data(s);
}

void tlgen::details::Service5InsertReset(::tlgen::service5::Insert& item) noexcept {
  (void)item;
  item.table.clear();
  item.data.clear();
}

bool tlgen::details::Service5InsertWriteJSON(std::ostream& s, const ::tlgen::service5::Insert& item) noexcept {
  (void)s;
  (void)item;
  auto add_comma = false;
  s << "{";
  if (item.table.size() != 0) {
    add_comma = true;
    s << "\"table\":";
    s << "\"" << item.table << "\"";
  }
  if (item.data.size() != 0) {
    if (add_comma) {
      s << ",";
    }
    add_comma = true;
    s << "\"data\":";
    s << "\"" << item.data << "\"";
  }
  s << "}";
  return true;
}

bool tlgen::details::Service5InsertRead(::tlgen::basictl::tl_istream & s, ::tlgen::service5::Insert& item) noexcept {
  (void)s;
  (void)item;
  if (!s.string_read(item.table)) { return false; }
  if (!s.string_read(item.data)) { return false; }
  return true;
}

bool tlgen::details::Service5InsertWrite(::tlgen::basictl::tl_ostream & s, const ::tlgen::service5::Insert& item) noexcept {
  (void)s;
  (void)item;
  if (!s.string_write(item.table)) { return false;}
  if (!s.string_write(item.data)) { return false;}
  return true;
}

bool tlgen::details::Service5InsertReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::service5::Insert& item) {
  if (!s.nat_read_exact_tag(0xc911ee2c)) { return false; }
  return tlgen::details::Service5InsertRead(s, item);
}

bool tlgen::details::Service5InsertWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::service5::Insert& item) {
  if (!s.nat_write(0xc911ee2c)) { return false; }
  return tlgen::details::Service5InsertWrite(s, item);
}

bool tlgen::details::Service5InsertReadResult(::tlgen::basictl::tl_istream & s, const tlgen::service5::Insert& item, ::tlgen::service5::Output& result) {
  (void)s;
  (void)item;
  (void)result;
  if (!::tlgen::details::Service5OutputReadBoxed(s, result)) { return s.set_error_unknown_scenario(); }
  return true;
}
bool tlgen::details::Service5InsertWriteResult(::tlgen::basictl::tl_ostream & s, const tlgen::service5::Insert& item, const ::tlgen::service5::Output& result) {
  (void)s;
  (void)item;
  (void)result;
  if (!::tlgen::details::Service5OutputWriteBoxed(s, result)) { return s.set_error_unknown_scenario(); }
  return true;
}

bool tlgen::service5::Insert::read_result(::tlgen::basictl::tl_istream & s, ::tlgen::service5::Output & result) const noexcept {
  bool success = tlgen::details::Service5InsertReadResult(s, *this, result);
  s.sync();
  return success;
}
bool tlgen::service5::Insert::write_result(::tlgen::basictl::tl_ostream & s, const ::tlgen::service5::Output & result) const noexcept {
  bool success = tlgen::details::Service5InsertWriteResult(s, *this, result);
  s.sync();
  return success;
}

void tlgen::service5::Insert::read_result(::tlgen::basictl::tl_throwable_istream & s, ::tlgen::service5::Output & result) const {
  ::tlgen::basictl::tl_istream s2(s);
  this->read_result(s2, result);
  s2.pass_data(s);
}
void tlgen::service5::Insert::write_result(::tlgen::basictl::tl_throwable_ostream & s, const ::tlgen::service5::Output & result) const {
  ::tlgen::basictl::tl_ostream s2(s);
  this->write_result(s2, result);
  s2.pass_data(s);
}

static const std::string_view Service5Output_tbl_tl_name[]{"service5.emptyOutput", "service5.stringOutput"};
static const uint32_t Service5Output_tbl_tl_tag[]{0x11e46879, 0x179e9863};

bool tlgen::service5::Output::write_json(std::ostream & s)const {
  if (!::tlgen::details::Service5OutputWriteJSON(s, *this)) { return false; }
  return true;
}
bool tlgen::service5::Output::read_boxed(::tlgen::basictl::tl_istream & s) noexcept {
  if (!::tlgen::details::Service5OutputReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
  return true;
}
bool tlgen::service5::Output::write_boxed(::tlgen::basictl::tl_ostream & s)const noexcept {
  if (!::tlgen::details::Service5OutputWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
  return true;
}

void tlgen::service5::Output::read_boxed(::tlgen::basictl::tl_throwable_istream & s) {
  ::tlgen::basictl::tl_istream s2(s);
  this->read_boxed(s2);
  s2.pass_data(s);
}

void tlgen::service5::Output::write_boxed(::tlgen::basictl::tl_throwable_ostream & s)const {
  ::tlgen::basictl::tl_ostream s2(s);
  this->write_boxed(s2);
  s2.pass_data(s);
}

std::string_view tlgen::service5::Output::tl_name() const {
  return Service5Output_tbl_tl_name[value.index()];
}
uint32_t tlgen::service5::Output::tl_tag() const {
  return Service5Output_tbl_tl_tag[value.index()];
}


void tlgen::details::Service5OutputReset(::tlgen::service5::Output& item) noexcept{
  item.value.emplace<0>(); // TODO - optimize, if already 0, call Reset function
}

bool tlgen::details::Service5OutputWriteJSON(std::ostream & s, const ::tlgen::service5::Output& item) noexcept {
  s << "{";
  s << "\"type\":";
  s << "\"" << Service5Output_tbl_tl_name[item.value.index()] << "\"";
  switch (item.value.index()) {
  case 1:
    s << ",\"value\":";
    if (!::tlgen::details::Service5StringOutputWriteJSON(s, std::get<1>(item.value))) { return false; }
    break;
  }
  s << "}";
  return true;
}
bool tlgen::details::Service5OutputReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::service5::Output& item) noexcept {
  uint32_t nat;
  if (!s.nat_read(nat)) { return false; }
  switch (nat) {
  case 0x11e46879:
    if (item.value.index() != 0) { item.value.emplace<0>(); }
    break;
  case 0x179e9863:
    if (item.value.index() != 1) { item.value.emplace<1>(); }
    if (!::tlgen::details::Service5StringOutputRead(s, std::get<1>(item.value))) { return s.set_error_unknown_scenario(); }
    break;
  default:
    return s.set_error_union_tag();
    }
  return true;
}

bool tlgen::details::Service5OutputWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::service5::Output& item) noexcept{
  if (!s.nat_write(Service5Output_tbl_tl_tag[item.value.index()])) { return false; }
  switch (item.value.index()) {
  case 1:
    if (!::tlgen::details::Service5StringOutputWrite(s, std::get<1>(item.value))) { return s.set_error_unknown_scenario(); }
    break;
  }
  return true;
}

bool tlgen::service5::Params::write_json(std::ostream& s)const {
  if (!::tlgen::details::Service5ParamsWriteJSON(s, *this)) { return false; }
  return true;
}

bool tlgen::service5::Params::read(::tlgen::basictl::tl_istream & s) noexcept {
  if (!::tlgen::details::Service5ParamsRead(s, *this)) { return s.set_error_unknown_scenario(); }
  s.sync();
  return true;
}

bool tlgen::service5::Params::write(::tlgen::basictl::tl_ostream & s) const noexcept {
  if (!::tlgen::details::Service5ParamsWrite(s, *this)) { return s.set_error_unknown_scenario(); }
  s.sync();
  return true;
}

void tlgen::service5::Params::read(::tlgen::basictl::tl_throwable_istream & s) {
  ::tlgen::basictl::tl_istream s2(s);
  this->read(s2);
  s2.pass_data(s);
}

void tlgen::service5::Params::write(::tlgen::basictl::tl_throwable_ostream & s) const {
  ::tlgen::basictl::tl_ostream s2(s);
  this->write(s2);
  s2.pass_data(s);
}

bool tlgen::service5::Params::read_boxed(::tlgen::basictl::tl_istream & s) noexcept {
  if (!::tlgen::details::Service5ParamsReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
  s.sync();
  return true;
}

bool tlgen::service5::Params::write_boxed(::tlgen::basictl::tl_ostream & s) const noexcept {
  if (!::tlgen::details::Service5ParamsWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
  s.sync();
  return true;
}

void tlgen::service5::Params::read_boxed(::tlgen::basictl::tl_throwable_istream & s) {
  ::tlgen::basictl::tl_istream s2(s);
  this->read_boxed(s2);
  s2.pass_data(s);
}

void tlgen::service5::Params::write_boxed(::tlgen::basictl::tl_throwable_ostream & s) const {
  ::tlgen::basictl::tl_ostream s2(s);
  this->write_boxed(s2);
  s2.pass_data(s);
}

void tlgen::details::Service5ParamsReset(::tlgen::service5::Params& item) noexcept {
  (void)item;
  item.fields_mask = 0;
  item.max_execution_speed = 0;
  item.max_execution_speed_bytes = 0;
}

bool tlgen::details::Service5ParamsWriteJSON(std::ostream& s, const ::tlgen::service5::Params& item) noexcept {
  (void)s;
  (void)item;
  auto add_comma = false;
  s << "{";
  if (item.fields_mask != 0) {
    add_comma = true;
    s << "\"fields_mask\":";
    s << item.fields_mask;
  }
  if ((item.fields_mask & (1<<0)) != 0) {
    if (add_comma) {
      s << ",";
    }
    add_comma = true;
    s << "\"max_execution_speed\":";
    s << item.max_execution_speed;
  }
  if ((item.fields_mask & (1<<1)) != 0) {
    if (add_comma) {
      s << ",";
    }
    add_comma = true;
    s << "\"max_execution_speed_bytes\":";
    s << item.max_execution_speed_bytes;
  }
  s << "}";
  return true;
}

bool tlgen::details::Service5ParamsRead(::tlgen::basictl::tl_istream & s, ::tlgen::service5::Params& item) noexcept {
  (void)s;
  (void)item;
  if (!s.nat_read(item.fields_mask)) { return false; }
  if ((item.fields_mask & (1<<0)) != 0) {
    if (!s.int_read(item.max_execution_speed)) { return false; }
  } else {
      item.max_execution_speed = 0;
  }
  if ((item.fields_mask & (1<<1)) != 0) {
    if (!s.int_read(item.max_execution_speed_bytes)) { return false; }
  } else {
      item.max_execution_speed_bytes = 0;
  }
  return true;
}

bool tlgen::details::Service5ParamsWrite(::tlgen::basictl::tl_ostream & s, const ::tlgen::service5::Params& item) noexcept {
  (void)s;
  (void)item;
  if (!s.nat_write(item.fields_mask)) { return false;}
  if ((item.fields_mask & (1<<0)) != 0) {
      if (!s.int_write(item.max_execution_speed)) { return false;}
  }
  if ((item.fields_mask & (1<<1)) != 0) {
      if (!s.int_write(item.max_execution_speed_bytes)) { return false;}
  }
  return true;
}

bool tlgen::details::Service5ParamsReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::service5::Params& item) {
  if (!s.nat_read_exact_tag(0x12ae5cb5)) { return false; }
  return tlgen::details::Service5ParamsRead(s, item);
}

bool tlgen::details::Service5ParamsWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::service5::Params& item) {
  if (!s.nat_write(0x12ae5cb5)) { return false; }
  return tlgen::details::Service5ParamsWrite(s, item);
}

bool tlgen::service5::PerformQuery::write_json(std::ostream& s)const {
  if (!::tlgen::details::Service5PerformQueryWriteJSON(s, *this)) { return false; }
  return true;
}

bool tlgen::service5::PerformQuery::read(::tlgen::basictl::tl_istream & s) noexcept {
  if (!::tlgen::details::Service5PerformQueryRead(s, *this)) { return s.set_error_unknown_scenario(); }
  s.sync();
  return true;
}

bool tlgen::service5::PerformQuery::write(::tlgen::basictl::tl_ostream & s) const noexcept {
  if (!::tlgen::details::Service5PerformQueryWrite(s, *this)) { return s.set_error_unknown_scenario(); }
  s.sync();
  return true;
}

void tlgen::service5::PerformQuery::read(::tlgen::basictl::tl_throwable_istream & s) {
  ::tlgen::basictl::tl_istream s2(s);
  this->read(s2);
  s2.pass_data(s);
}

void tlgen::service5::PerformQuery::write(::tlgen::basictl::tl_throwable_ostream & s) const {
  ::tlgen::basictl::tl_ostream s2(s);
  this->write(s2);
  s2.pass_data(s);
}

bool tlgen::service5::PerformQuery::read_boxed(::tlgen::basictl::tl_istream & s) noexcept {
  if (!::tlgen::details::Service5PerformQueryReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
  s.sync();
  return true;
}

bool tlgen::service5::PerformQuery::write_boxed(::tlgen::basictl::tl_ostream & s) const noexcept {
  if (!::tlgen::details::Service5PerformQueryWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
  s.sync();
  return true;
}

void tlgen::service5::PerformQuery::read_boxed(::tlgen::basictl::tl_throwable_istream & s) {
  ::tlgen::basictl::tl_istream s2(s);
  this->read_boxed(s2);
  s2.pass_data(s);
}

void tlgen::service5::PerformQuery::write_boxed(::tlgen::basictl::tl_throwable_ostream & s) const {
  ::tlgen::basictl::tl_ostream s2(s);
  this->write_boxed(s2);
  s2.pass_data(s);
}

void tlgen::details::Service5PerformQueryReset(::tlgen::service5::PerformQuery& item) noexcept {
  (void)item;
  item.query.clear();
}

bool tlgen::details::Service5PerformQueryWriteJSON(std::ostream& s, const ::tlgen::service5::PerformQuery& item) noexcept {
  (void)s;
  (void)item;
  s << "{";
  if (item.query.size() != 0) {
    s << "\"query\":";
    s << "\"" << item.query << "\"";
  }
  s << "}";
  return true;
}

bool tlgen::details::Service5PerformQueryRead(::tlgen::basictl::tl_istream & s, ::tlgen::service5::PerformQuery& item) noexcept {
  (void)s;
  (void)item;
  if (!s.string_read(item.query)) { return false; }
  return true;
}

bool tlgen::details::Service5PerformQueryWrite(::tlgen::basictl::tl_ostream & s, const ::tlgen::service5::PerformQuery& item) noexcept {
  (void)s;
  (void)item;
  if (!s.string_write(item.query)) { return false;}
  return true;
}

bool tlgen::details::Service5PerformQueryReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::service5::PerformQuery& item) {
  if (!s.nat_read_exact_tag(0x019d80a5)) { return false; }
  return tlgen::details::Service5PerformQueryRead(s, item);
}

bool tlgen::details::Service5PerformQueryWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::service5::PerformQuery& item) {
  if (!s.nat_write(0x019d80a5)) { return false; }
  return tlgen::details::Service5PerformQueryWrite(s, item);
}

bool tlgen::details::Service5PerformQueryReadResult(::tlgen::basictl::tl_istream & s, const tlgen::service5::PerformQuery& item, ::tlgen::service5::Output& result) {
  (void)s;
  (void)item;
  (void)result;
  if (!::tlgen::details::Service5OutputReadBoxed(s, result)) { return s.set_error_unknown_scenario(); }
  return true;
}
bool tlgen::details::Service5PerformQueryWriteResult(::tlgen::basictl::tl_ostream & s, const tlgen::service5::PerformQuery& item, const ::tlgen::service5::Output& result) {
  (void)s;
  (void)item;
  (void)result;
  if (!::tlgen::details::Service5OutputWriteBoxed(s, result)) { return s.set_error_unknown_scenario(); }
  return true;
}

bool tlgen::service5::PerformQuery::read_result(::tlgen::basictl::tl_istream & s, ::tlgen::service5::Output & result) const noexcept {
  bool success = tlgen::details::Service5PerformQueryReadResult(s, *this, result);
  s.sync();
  return success;
}
bool tlgen::service5::PerformQuery::write_result(::tlgen::basictl::tl_ostream & s, const ::tlgen::service5::Output & result) const noexcept {
  bool success = tlgen::details::Service5PerformQueryWriteResult(s, *this, result);
  s.sync();
  return success;
}

void tlgen::service5::PerformQuery::read_result(::tlgen::basictl::tl_throwable_istream & s, ::tlgen::service5::Output & result) const {
  ::tlgen::basictl::tl_istream s2(s);
  this->read_result(s2, result);
  s2.pass_data(s);
}
void tlgen::service5::PerformQuery::write_result(::tlgen::basictl::tl_throwable_ostream & s, const ::tlgen::service5::Output & result) const {
  ::tlgen::basictl::tl_ostream s2(s);
  this->write_result(s2, result);
  s2.pass_data(s);
}

bool tlgen::service5::Query::write_json(std::ostream& s)const {
  if (!::tlgen::details::Service5QueryWriteJSON(s, *this)) { return false; }
  return true;
}

bool tlgen::service5::Query::read(::tlgen::basictl::tl_istream & s) noexcept {
  if (!::tlgen::details::Service5QueryRead(s, *this)) { return s.set_error_unknown_scenario(); }
  s.sync();
  return true;
}

bool tlgen::service5::Query::write(::tlgen::basictl::tl_ostream & s) const noexcept {
  if (!::tlgen::details::Service5QueryWrite(s, *this)) { return s.set_error_unknown_scenario(); }
  s.sync();
  return true;
}

void tlgen::service5::Query::read(::tlgen::basictl::tl_throwable_istream & s) {
  ::tlgen::basictl::tl_istream s2(s);
  this->read(s2);
  s2.pass_data(s);
}

void tlgen::service5::Query::write(::tlgen::basictl::tl_throwable_ostream & s) const {
  ::tlgen::basictl::tl_ostream s2(s);
  this->write(s2);
  s2.pass_data(s);
}

bool tlgen::service5::Query::read_boxed(::tlgen::basictl::tl_istream & s) noexcept {
  if (!::tlgen::details::Service5QueryReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
  s.sync();
  return true;
}

bool tlgen::service5::Query::write_boxed(::tlgen::basictl::tl_ostream & s) const noexcept {
  if (!::tlgen::details::Service5QueryWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
  s.sync();
  return true;
}

void tlgen::service5::Query::read_boxed(::tlgen::basictl::tl_throwable_istream & s) {
  ::tlgen::basictl::tl_istream s2(s);
  this->read_boxed(s2);
  s2.pass_data(s);
}

void tlgen::service5::Query::write_boxed(::tlgen::basictl::tl_throwable_ostream & s) const {
  ::tlgen::basictl::tl_ostream s2(s);
  this->write_boxed(s2);
  s2.pass_data(s);
}

void tlgen::details::Service5QueryReset(::tlgen::service5::Query& item) noexcept {
  (void)item;
  item.query.clear();
  ::tlgen::details::Service5ParamsReset(item.params);
}

bool tlgen::details::Service5QueryWriteJSON(std::ostream& s, const ::tlgen::service5::Query& item) noexcept {
  (void)s;
  (void)item;
  auto add_comma = false;
  s << "{";
  if (item.query.size() != 0) {
    add_comma = true;
    s << "\"query\":";
    s << "\"" << item.query << "\"";
  }
  if (add_comma) {
    s << ",";
  }
  add_comma = true;
  s << "\"params\":";
  if (!::tlgen::details::Service5ParamsWriteJSON(s, item.params)) { return false; }
  s << "}";
  return true;
}

bool tlgen::details::Service5QueryRead(::tlgen::basictl::tl_istream & s, ::tlgen::service5::Query& item) noexcept {
  (void)s;
  (void)item;
  if (!s.string_read(item.query)) { return false; }
  if (!::tlgen::details::Service5ParamsRead(s, item.params)) { return s.set_error_unknown_scenario(); }
  return true;
}

bool tlgen::details::Service5QueryWrite(::tlgen::basictl::tl_ostream & s, const ::tlgen::service5::Query& item) noexcept {
  (void)s;
  (void)item;
  if (!s.string_write(item.query)) { return false;}
  if (!::tlgen::details::Service5ParamsWrite(s, item.params)) { return s.set_error_unknown_scenario(); }
  return true;
}

bool tlgen::details::Service5QueryReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::service5::Query& item) {
  if (!s.nat_read_exact_tag(0xb3b62513)) { return false; }
  return tlgen::details::Service5QueryRead(s, item);
}

bool tlgen::details::Service5QueryWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::service5::Query& item) {
  if (!s.nat_write(0xb3b62513)) { return false; }
  return tlgen::details::Service5QueryWrite(s, item);
}

bool tlgen::details::Service5QueryReadResult(::tlgen::basictl::tl_istream & s, const tlgen::service5::Query& item, ::tlgen::service5::Output& result) {
  (void)s;
  (void)item;
  (void)result;
  if (!::tlgen::details::Service5OutputReadBoxed(s, result)) { return s.set_error_unknown_scenario(); }
  return true;
}
bool tlgen::details::Service5QueryWriteResult(::tlgen::basictl::tl_ostream & s, const tlgen::service5::Query& item, const ::tlgen::service5::Output& result) {
  (void)s;
  (void)item;
  (void)result;
  if (!::tlgen::details::Service5OutputWriteBoxed(s, result)) { return s.set_error_unknown_scenario(); }
  return true;
}

bool tlgen::service5::Query::read_result(::tlgen::basictl::tl_istream & s, ::tlgen::service5::Output & result) const noexcept {
  bool success = tlgen::details::Service5QueryReadResult(s, *this, result);
  s.sync();
  return success;
}
bool tlgen::service5::Query::write_result(::tlgen::basictl::tl_ostream & s, const ::tlgen::service5::Output & result) const noexcept {
  bool success = tlgen::details::Service5QueryWriteResult(s, *this, result);
  s.sync();
  return success;
}

void tlgen::service5::Query::read_result(::tlgen::basictl::tl_throwable_istream & s, ::tlgen::service5::Output & result) const {
  ::tlgen::basictl::tl_istream s2(s);
  this->read_result(s2, result);
  s2.pass_data(s);
}
void tlgen::service5::Query::write_result(::tlgen::basictl::tl_throwable_ostream & s, const ::tlgen::service5::Output & result) const {
  ::tlgen::basictl::tl_ostream s2(s);
  this->write_result(s2, result);
  s2.pass_data(s);
}

bool tlgen::service5::StringOutput::write_json(std::ostream& s)const {
  if (!::tlgen::details::Service5StringOutputWriteJSON(s, *this)) { return false; }
  return true;
}

bool tlgen::service5::StringOutput::read(::tlgen::basictl::tl_istream & s) noexcept {
  if (!::tlgen::details::Service5StringOutputRead(s, *this)) { return s.set_error_unknown_scenario(); }
  s.sync();
  return true;
}

bool tlgen::service5::StringOutput::write(::tlgen::basictl::tl_ostream & s) const noexcept {
  if (!::tlgen::details::Service5StringOutputWrite(s, *this)) { return s.set_error_unknown_scenario(); }
  s.sync();
  return true;
}

void tlgen::service5::StringOutput::read(::tlgen::basictl::tl_throwable_istream & s) {
  ::tlgen::basictl::tl_istream s2(s);
  this->read(s2);
  s2.pass_data(s);
}

void tlgen::service5::StringOutput::write(::tlgen::basictl::tl_throwable_ostream & s) const {
  ::tlgen::basictl::tl_ostream s2(s);
  this->write(s2);
  s2.pass_data(s);
}

bool tlgen::service5::StringOutput::read_boxed(::tlgen::basictl::tl_istream & s) noexcept {
  if (!::tlgen::details::Service5StringOutputReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
  s.sync();
  return true;
}

bool tlgen::service5::StringOutput::write_boxed(::tlgen::basictl::tl_ostream & s) const noexcept {
  if (!::tlgen::details::Service5StringOutputWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
  s.sync();
  return true;
}

void tlgen::service5::StringOutput::read_boxed(::tlgen::basictl::tl_throwable_istream & s) {
  ::tlgen::basictl::tl_istream s2(s);
  this->read_boxed(s2);
  s2.pass_data(s);
}

void tlgen::service5::StringOutput::write_boxed(::tlgen::basictl::tl_throwable_ostream & s) const {
  ::tlgen::basictl::tl_ostream s2(s);
  this->write_boxed(s2);
  s2.pass_data(s);
}

void tlgen::details::Service5StringOutputReset(::tlgen::service5::StringOutput& item) noexcept {
  (void)item;
  item.http_code = 0;
  item.response.clear();
}

bool tlgen::details::Service5StringOutputWriteJSON(std::ostream& s, const ::tlgen::service5::StringOutput& item) noexcept {
  (void)s;
  (void)item;
  auto add_comma = false;
  s << "{";
  if (item.http_code != 0) {
    add_comma = true;
    s << "\"http_code\":";
    s << item.http_code;
  }
  if (item.response.size() != 0) {
    if (add_comma) {
      s << ",";
    }
    add_comma = true;
    s << "\"response\":";
    s << "\"" << item.response << "\"";
  }
  s << "}";
  return true;
}

bool tlgen::details::Service5StringOutputRead(::tlgen::basictl::tl_istream & s, ::tlgen::service5::StringOutput& item) noexcept {
  (void)s;
  (void)item;
  if (!s.int_read(item.http_code)) { return false; }
  if (!s.string_read(item.response)) { return false; }
  return true;
}

bool tlgen::details::Service5StringOutputWrite(::tlgen::basictl::tl_ostream & s, const ::tlgen::service5::StringOutput& item) noexcept {
  (void)s;
  (void)item;
  if (!s.int_write(item.http_code)) { return false;}
  if (!s.string_write(item.response)) { return false;}
  return true;
}

bool tlgen::details::Service5StringOutputReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::service5::StringOutput& item) {
  if (!s.nat_read_exact_tag(0x179e9863)) { return false; }
  return tlgen::details::Service5StringOutputRead(s, item);
}

bool tlgen::details::Service5StringOutputWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::service5::StringOutput& item) {
  if (!s.nat_write(0x179e9863)) { return false; }
  return tlgen::details::Service5StringOutputWrite(s, item);
}
