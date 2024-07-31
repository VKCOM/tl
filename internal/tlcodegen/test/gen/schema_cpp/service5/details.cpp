#include "headers/service5.query.hpp"
#include "headers/service5.performQuery.hpp"
#include "headers/service5.params.hpp"
#include "headers/service5.Output.hpp"
#include "headers/service5.stringOutput.hpp"
#include "headers/service5.insert.hpp"
#include "headers/service5.emptyOutput.hpp"


bool tl2::service5::EmptyOutput::write_json(std::ostream& s)const {
	if (!::tl2::details::Service5EmptyOutputWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service5::EmptyOutput::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service5EmptyOutputRead(s, *this)) { return false; }
	return true;
}

bool tl2::service5::EmptyOutput::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service5EmptyOutputWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service5::EmptyOutput::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service5EmptyOutputReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service5::EmptyOutput::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service5EmptyOutputWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service5EmptyOutputReset(::tl2::service5::EmptyOutput& item) {
}

bool tl2::details::Service5EmptyOutputWriteJSON(std::ostream& s, const ::tl2::service5::EmptyOutput& item) {
	s << "{";
	s << "}";
	return true;
}

bool tl2::details::Service5EmptyOutputRead(::basictl::tl_istream & s, ::tl2::service5::EmptyOutput& item) {
	return true;
}

bool tl2::details::Service5EmptyOutputWrite(::basictl::tl_ostream & s, const ::tl2::service5::EmptyOutput& item) {
	return true;
}

bool tl2::details::Service5EmptyOutputReadBoxed(::basictl::tl_istream & s, ::tl2::service5::EmptyOutput& item) {
	if (!s.nat_read_exact_tag(0x11e46879)) { return false; }
	return tl2::details::Service5EmptyOutputRead(s, item);
}

bool tl2::details::Service5EmptyOutputWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service5::EmptyOutput& item) {
	if (!s.nat_write(0x11e46879)) { return false; }
	return tl2::details::Service5EmptyOutputWrite(s, item);
}

bool tl2::service5::Insert::write_json(std::ostream& s)const {
	if (!::tl2::details::Service5InsertWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service5::Insert::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service5InsertRead(s, *this)) { return false; }
	return true;
}

bool tl2::service5::Insert::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service5InsertWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service5::Insert::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service5InsertReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service5::Insert::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service5InsertWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service5InsertReset(::tl2::service5::Insert& item) {
	item.table.clear();
	item.data.clear();
}

bool tl2::details::Service5InsertWriteJSON(std::ostream& s, const ::tl2::service5::Insert& item) {
	s << "{";
	s << "\"table\":";
	s << "\"" << item.table << "\"";
	s << ",";
	s << "\"data\":";
	s << "\"" << item.data << "\"";
	s << "}";
	return true;
}

bool tl2::details::Service5InsertRead(::basictl::tl_istream & s, ::tl2::service5::Insert& item) {
	if (!s.string_read(item.table)) { return false; }
	if (!s.string_read(item.data)) { return false; }
	return true;
}

bool tl2::details::Service5InsertWrite(::basictl::tl_ostream & s, const ::tl2::service5::Insert& item) {
	if (!s.string_write(item.table)) { return false;}
	if (!s.string_write(item.data)) { return false;}
	return true;
}

bool tl2::details::Service5InsertReadBoxed(::basictl::tl_istream & s, ::tl2::service5::Insert& item) {
	if (!s.nat_read_exact_tag(0xc911ee2c)) { return false; }
	return tl2::details::Service5InsertRead(s, item);
}

bool tl2::details::Service5InsertWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service5::Insert& item) {
	if (!s.nat_write(0xc911ee2c)) { return false; }
	return tl2::details::Service5InsertWrite(s, item);
}

bool tl2::details::Service5InsertReadResult(::basictl::tl_istream & s, tl2::service5::Insert& item, ::tl2::service5::Output& result) {
	if (!::tl2::details::Service5OutputReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service5InsertWriteResult(::basictl::tl_ostream & s, tl2::service5::Insert& item, ::tl2::service5::Output& result) {
	if (!::tl2::details::Service5OutputWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service5::Insert::read_result(::basictl::tl_istream & s, ::tl2::service5::Output & result) {
	return tl2::details::Service5InsertReadResult(s, *this, result);
}
bool tl2::service5::Insert::write_result(::basictl::tl_ostream & s, ::tl2::service5::Output & result) {
	return tl2::details::Service5InsertWriteResult(s, *this, result);
}

static const std::string_view Service5Output_tbl_tl_name[]{"service5.emptyOutput", "service5.stringOutput"};
static const uint32_t Service5Output_tbl_tl_tag[]{0x11e46879, 0x179e9863};

bool tl2::service5::Output::write_json(std::ostream & s)const {
	if (!::tl2::details::Service5OutputWriteJSON(s, *this)) { return false; }
	return true;
}
bool tl2::service5::Output::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service5OutputReadBoxed(s, *this)) { return false; }
	return true;
}
bool tl2::service5::Output::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service5OutputWriteBoxed(s, *this)) { return false; }
	return true;
}
std::string_view tl2::service5::Output::tl_name() const {
	return Service5Output_tbl_tl_name[value.index()];
}
uint32_t tl2::service5::Output::tl_tag() const {
	return Service5Output_tbl_tl_tag[value.index()];
}


void tl2::details::Service5OutputReset(::tl2::service5::Output& item) {
	item.value.emplace<0>(); // TODO - optimize, if already 0, call Reset function
}

bool tl2::details::Service5OutputWriteJSON(std::ostream & s, const ::tl2::service5::Output& item) {
	s << "{";
	s << "\"type\":";
	s << Service5Output_tbl_tl_tag[item.value.index()];
	switch (item.value.index()) {
	case 1:
		s << ",\"value\":";
		if (!::tl2::details::Service5StringOutputWriteJSON(s, std::get<1>(item.value))) { return false; }
		break;
	}
	s << "}";
	return true;
}
bool tl2::details::Service5OutputReadBoxed(::basictl::tl_istream & s, ::tl2::service5::Output& item) {
	uint32_t nat;
	s.nat_read(nat);
	switch (nat) {
	case 0x11e46879:
		if (item.value.index() != 0) { item.value.emplace<0>(); }
		break;
	case 0x179e9863:
		if (item.value.index() != 1) { item.value.emplace<1>(); }
		if (!::tl2::details::Service5StringOutputRead(s, std::get<1>(item.value))) { return false; }
		break;
	default:
		return s.set_error_union_tag();
    }
	return true;
}

bool tl2::details::Service5OutputWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service5::Output& item) {
	s.nat_write(Service5Output_tbl_tl_tag[item.value.index()]);
	switch (item.value.index()) {
	case 1:
		if (!::tl2::details::Service5StringOutputWrite(s, std::get<1>(item.value))) { return false; }
		break;
	}
	return true;
}

bool tl2::service5::Params::write_json(std::ostream& s)const {
	if (!::tl2::details::Service5ParamsWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service5::Params::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service5ParamsRead(s, *this)) { return false; }
	return true;
}

bool tl2::service5::Params::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service5ParamsWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service5::Params::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service5ParamsReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service5::Params::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service5ParamsWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service5ParamsReset(::tl2::service5::Params& item) {
	item.fields_mask = 0;
	item.max_execution_speed = 0;
	item.max_execution_speed_bytes = 0;
}

bool tl2::details::Service5ParamsWriteJSON(std::ostream& s, const ::tl2::service5::Params& item) {
	s << "{";
	s << "\"fields_mask\":";
	s << item.fields_mask;
	if ((item.fields_mask & (1<<0)) != 0) {
		s << ",";
		s << "\"max_execution_speed\":";
		s << item.max_execution_speed;
	}
	if ((item.fields_mask & (1<<1)) != 0) {
		s << ",";
		s << "\"max_execution_speed_bytes\":";
		s << item.max_execution_speed_bytes;
	}
	s << "}";
	return true;
}

bool tl2::details::Service5ParamsRead(::basictl::tl_istream & s, ::tl2::service5::Params& item) {
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

bool tl2::details::Service5ParamsWrite(::basictl::tl_ostream & s, const ::tl2::service5::Params& item) {
	if (!s.nat_write(item.fields_mask)) { return false;}
	if ((item.fields_mask & (1<<0)) != 0) {
			if (!s.int_write(item.max_execution_speed)) { return false;}
	}
	if ((item.fields_mask & (1<<1)) != 0) {
			if (!s.int_write(item.max_execution_speed_bytes)) { return false;}
	}
	return true;
}

bool tl2::details::Service5ParamsReadBoxed(::basictl::tl_istream & s, ::tl2::service5::Params& item) {
	if (!s.nat_read_exact_tag(0x12ae5cb5)) { return false; }
	return tl2::details::Service5ParamsRead(s, item);
}

bool tl2::details::Service5ParamsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service5::Params& item) {
	if (!s.nat_write(0x12ae5cb5)) { return false; }
	return tl2::details::Service5ParamsWrite(s, item);
}

bool tl2::service5::PerformQuery::write_json(std::ostream& s)const {
	if (!::tl2::details::Service5PerformQueryWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service5::PerformQuery::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service5PerformQueryRead(s, *this)) { return false; }
	return true;
}

bool tl2::service5::PerformQuery::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service5PerformQueryWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service5::PerformQuery::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service5PerformQueryReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service5::PerformQuery::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service5PerformQueryWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service5PerformQueryReset(::tl2::service5::PerformQuery& item) {
	item.query.clear();
}

bool tl2::details::Service5PerformQueryWriteJSON(std::ostream& s, const ::tl2::service5::PerformQuery& item) {
	s << "{";
	s << "\"query\":";
	s << "\"" << item.query << "\"";
	s << "}";
	return true;
}

bool tl2::details::Service5PerformQueryRead(::basictl::tl_istream & s, ::tl2::service5::PerformQuery& item) {
	if (!s.string_read(item.query)) { return false; }
	return true;
}

bool tl2::details::Service5PerformQueryWrite(::basictl::tl_ostream & s, const ::tl2::service5::PerformQuery& item) {
	if (!s.string_write(item.query)) { return false;}
	return true;
}

bool tl2::details::Service5PerformQueryReadBoxed(::basictl::tl_istream & s, ::tl2::service5::PerformQuery& item) {
	if (!s.nat_read_exact_tag(0x019d80a5)) { return false; }
	return tl2::details::Service5PerformQueryRead(s, item);
}

bool tl2::details::Service5PerformQueryWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service5::PerformQuery& item) {
	if (!s.nat_write(0x019d80a5)) { return false; }
	return tl2::details::Service5PerformQueryWrite(s, item);
}

bool tl2::details::Service5PerformQueryReadResult(::basictl::tl_istream & s, tl2::service5::PerformQuery& item, ::tl2::service5::Output& result) {
	if (!::tl2::details::Service5OutputReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service5PerformQueryWriteResult(::basictl::tl_ostream & s, tl2::service5::PerformQuery& item, ::tl2::service5::Output& result) {
	if (!::tl2::details::Service5OutputWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service5::PerformQuery::read_result(::basictl::tl_istream & s, ::tl2::service5::Output & result) {
	return tl2::details::Service5PerformQueryReadResult(s, *this, result);
}
bool tl2::service5::PerformQuery::write_result(::basictl::tl_ostream & s, ::tl2::service5::Output & result) {
	return tl2::details::Service5PerformQueryWriteResult(s, *this, result);
}

bool tl2::service5::Query::write_json(std::ostream& s)const {
	if (!::tl2::details::Service5QueryWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service5::Query::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service5QueryRead(s, *this)) { return false; }
	return true;
}

bool tl2::service5::Query::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service5QueryWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service5::Query::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service5QueryReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service5::Query::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service5QueryWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service5QueryReset(::tl2::service5::Query& item) {
	item.query.clear();
	::tl2::details::Service5ParamsReset(item.params);
}

bool tl2::details::Service5QueryWriteJSON(std::ostream& s, const ::tl2::service5::Query& item) {
	s << "{";
	s << "\"query\":";
	s << "\"" << item.query << "\"";
	s << ",";
	s << "\"params\":";
	if (!::tl2::details::Service5ParamsWriteJSON(s, item.params)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::Service5QueryRead(::basictl::tl_istream & s, ::tl2::service5::Query& item) {
	if (!s.string_read(item.query)) { return false; }
	if (!::tl2::details::Service5ParamsRead(s, item.params)) { return false; }
	return true;
}

bool tl2::details::Service5QueryWrite(::basictl::tl_ostream & s, const ::tl2::service5::Query& item) {
	if (!s.string_write(item.query)) { return false;}
	if (!::tl2::details::Service5ParamsWrite(s, item.params)) { return false; }
	return true;
}

bool tl2::details::Service5QueryReadBoxed(::basictl::tl_istream & s, ::tl2::service5::Query& item) {
	if (!s.nat_read_exact_tag(0xb3b62513)) { return false; }
	return tl2::details::Service5QueryRead(s, item);
}

bool tl2::details::Service5QueryWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service5::Query& item) {
	if (!s.nat_write(0xb3b62513)) { return false; }
	return tl2::details::Service5QueryWrite(s, item);
}

bool tl2::details::Service5QueryReadResult(::basictl::tl_istream & s, tl2::service5::Query& item, ::tl2::service5::Output& result) {
	if (!::tl2::details::Service5OutputReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service5QueryWriteResult(::basictl::tl_ostream & s, tl2::service5::Query& item, ::tl2::service5::Output& result) {
	if (!::tl2::details::Service5OutputWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service5::Query::read_result(::basictl::tl_istream & s, ::tl2::service5::Output & result) {
	return tl2::details::Service5QueryReadResult(s, *this, result);
}
bool tl2::service5::Query::write_result(::basictl::tl_ostream & s, ::tl2::service5::Output & result) {
	return tl2::details::Service5QueryWriteResult(s, *this, result);
}

bool tl2::service5::StringOutput::write_json(std::ostream& s)const {
	if (!::tl2::details::Service5StringOutputWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service5::StringOutput::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service5StringOutputRead(s, *this)) { return false; }
	return true;
}

bool tl2::service5::StringOutput::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service5StringOutputWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service5::StringOutput::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service5StringOutputReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service5::StringOutput::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service5StringOutputWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service5StringOutputReset(::tl2::service5::StringOutput& item) {
	item.http_code = 0;
	item.response.clear();
}

bool tl2::details::Service5StringOutputWriteJSON(std::ostream& s, const ::tl2::service5::StringOutput& item) {
	s << "{";
	s << "\"http_code\":";
	s << item.http_code;
	s << ",";
	s << "\"response\":";
	s << "\"" << item.response << "\"";
	s << "}";
	return true;
}

bool tl2::details::Service5StringOutputRead(::basictl::tl_istream & s, ::tl2::service5::StringOutput& item) {
	if (!s.int_read(item.http_code)) { return false; }
	if (!s.string_read(item.response)) { return false; }
	return true;
}

bool tl2::details::Service5StringOutputWrite(::basictl::tl_ostream & s, const ::tl2::service5::StringOutput& item) {
	if (!s.int_write(item.http_code)) { return false;}
	if (!s.string_write(item.response)) { return false;}
	return true;
}

bool tl2::details::Service5StringOutputReadBoxed(::basictl::tl_istream & s, ::tl2::service5::StringOutput& item) {
	if (!s.nat_read_exact_tag(0x179e9863)) { return false; }
	return tl2::details::Service5StringOutputRead(s, item);
}

bool tl2::details::Service5StringOutputWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service5::StringOutput& item) {
	if (!s.nat_write(0x179e9863)) { return false; }
	return tl2::details::Service5StringOutputWrite(s, item);
}
