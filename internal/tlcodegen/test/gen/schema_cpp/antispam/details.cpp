// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#include "antispam/headers/antispam.PatternFull.h"
#include "antispam/headers/antispam.patternNotFound.h"
#include "antispam/headers/antispam.patternFound.h"
#include "antispam/headers/antispam.getPattern.h"


bool tl2::antispam::GetPattern::write_json(std::ostream& s)const {
	if (!::tl2::details::AntispamGetPatternWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::antispam::GetPattern::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::AntispamGetPatternRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::antispam::GetPattern::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::AntispamGetPatternWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::antispam::GetPattern::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::antispam::GetPattern::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::antispam::GetPattern::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::AntispamGetPatternReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::antispam::GetPattern::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::AntispamGetPatternWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::antispam::GetPattern::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::antispam::GetPattern::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::AntispamGetPatternReset(::tl2::antispam::GetPattern& item) noexcept {
	item.id = 0;
}

bool tl2::details::AntispamGetPatternWriteJSON(std::ostream& s, const ::tl2::antispam::GetPattern& item) noexcept {
	s << "{";
	if (item.id != 0) {
		s << "\"id\":";
		s << item.id;
	}
	s << "}";
	return true;
}

bool tl2::details::AntispamGetPatternRead(::basictl::tl_istream & s, ::tl2::antispam::GetPattern& item) noexcept {
	if (!s.int_read(item.id)) { return false; }
	return true;
}

bool tl2::details::AntispamGetPatternWrite(::basictl::tl_ostream & s, const ::tl2::antispam::GetPattern& item) noexcept {
	if (!s.int_write(item.id)) { return false;}
	return true;
}

bool tl2::details::AntispamGetPatternReadBoxed(::basictl::tl_istream & s, ::tl2::antispam::GetPattern& item) {
	if (!s.nat_read_exact_tag(0x3de14136)) { return false; }
	return tl2::details::AntispamGetPatternRead(s, item);
}

bool tl2::details::AntispamGetPatternWriteBoxed(::basictl::tl_ostream & s, const ::tl2::antispam::GetPattern& item) {
	if (!s.nat_write(0x3de14136)) { return false; }
	return tl2::details::AntispamGetPatternWrite(s, item);
}

bool tl2::details::AntispamGetPatternReadResult(::basictl::tl_istream & s, tl2::antispam::GetPattern& item, ::tl2::antispam::PatternFull& result) {
	if (!::tl2::details::AntispamPatternFullReadBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}
bool tl2::details::AntispamGetPatternWriteResult(::basictl::tl_ostream & s, tl2::antispam::GetPattern& item, ::tl2::antispam::PatternFull& result) {
	if (!::tl2::details::AntispamPatternFullWriteBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::antispam::GetPattern::read_result(::basictl::tl_istream & s, ::tl2::antispam::PatternFull & result) noexcept {
	bool success = tl2::details::AntispamGetPatternReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::antispam::GetPattern::write_result(::basictl::tl_ostream & s, ::tl2::antispam::PatternFull & result) noexcept {
	bool success = tl2::details::AntispamGetPatternWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::antispam::GetPattern::read_result_or_throw(::basictl::tl_throwable_istream & s, ::tl2::antispam::PatternFull & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::antispam::GetPattern::write_result_or_throw(::basictl::tl_throwable_ostream & s, ::tl2::antispam::PatternFull & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::antispam::PatternFound::write_json(std::ostream& s)const {
	if (!::tl2::details::AntispamPatternFoundWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::antispam::PatternFound::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::AntispamPatternFoundRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::antispam::PatternFound::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::AntispamPatternFoundWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::antispam::PatternFound::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::antispam::PatternFound::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::antispam::PatternFound::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::AntispamPatternFoundReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::antispam::PatternFound::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::AntispamPatternFoundWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::antispam::PatternFound::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::antispam::PatternFound::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::AntispamPatternFoundReset(::tl2::antispam::PatternFound& item) noexcept {
	item.ip = 0;
	item.uahash = 0;
	item.flags = 0;
	item.type = 0;
	item.text.clear();
}

bool tl2::details::AntispamPatternFoundWriteJSON(std::ostream& s, const ::tl2::antispam::PatternFound& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.ip != 0) {
		add_comma = true;
		s << "\"ip\":";
		s << item.ip;
	}
	if (item.uahash != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"uahash\":";
		s << item.uahash;
	}
	if (item.flags != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"flags\":";
		s << item.flags;
	}
	if (item.type != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"type\":";
		s << item.type;
	}
	if (item.text.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"text\":";
		s << "\"" << item.text << "\"";
	}
	s << "}";
	return true;
}

bool tl2::details::AntispamPatternFoundRead(::basictl::tl_istream & s, ::tl2::antispam::PatternFound& item) noexcept {
	if (!s.int_read(item.ip)) { return false; }
	if (!s.int_read(item.uahash)) { return false; }
	if (!s.int_read(item.flags)) { return false; }
	if (!s.int_read(item.type)) { return false; }
	if (!s.string_read(item.text)) { return false; }
	return true;
}

bool tl2::details::AntispamPatternFoundWrite(::basictl::tl_ostream & s, const ::tl2::antispam::PatternFound& item) noexcept {
	if (!s.int_write(item.ip)) { return false;}
	if (!s.int_write(item.uahash)) { return false;}
	if (!s.int_write(item.flags)) { return false;}
	if (!s.int_write(item.type)) { return false;}
	if (!s.string_write(item.text)) { return false;}
	return true;
}

bool tl2::details::AntispamPatternFoundReadBoxed(::basictl::tl_istream & s, ::tl2::antispam::PatternFound& item) {
	if (!s.nat_read_exact_tag(0xa7688492)) { return false; }
	return tl2::details::AntispamPatternFoundRead(s, item);
}

bool tl2::details::AntispamPatternFoundWriteBoxed(::basictl::tl_ostream & s, const ::tl2::antispam::PatternFound& item) {
	if (!s.nat_write(0xa7688492)) { return false; }
	return tl2::details::AntispamPatternFoundWrite(s, item);
}

static const std::string_view AntispamPatternFull_tbl_tl_name[]{"antispam.patternFound", "antispam.patternNotFound"};
static const uint32_t AntispamPatternFull_tbl_tl_tag[]{0xa7688492, 0x2c22e225};

bool tl2::antispam::PatternFull::write_json(std::ostream & s)const {
	if (!::tl2::details::AntispamPatternFullWriteJSON(s, *this)) { return false; }
	return true;
}
bool tl2::antispam::PatternFull::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::AntispamPatternFullReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	return true;
}
bool tl2::antispam::PatternFull::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::AntispamPatternFullWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	return true;
}

void tl2::antispam::PatternFull::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::antispam::PatternFull::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

std::string_view tl2::antispam::PatternFull::tl_name() const {
	return AntispamPatternFull_tbl_tl_name[value.index()];
}
uint32_t tl2::antispam::PatternFull::tl_tag() const {
	return AntispamPatternFull_tbl_tl_tag[value.index()];
}


void tl2::details::AntispamPatternFullReset(::tl2::antispam::PatternFull& item) noexcept{
	item.value.emplace<0>(); // TODO - optimize, if already 0, call Reset function
}

bool tl2::details::AntispamPatternFullWriteJSON(std::ostream & s, const ::tl2::antispam::PatternFull& item) noexcept {
	s << "{";
	s << "\"type\":";
	s << "\"" << AntispamPatternFull_tbl_tl_name[item.value.index()] << "\"";
	switch (item.value.index()) {
	case 0:
		s << ",\"value\":";
		if (!::tl2::details::AntispamPatternFoundWriteJSON(s, std::get<0>(item.value))) { return false; }
		break;
	}
	s << "}";
	return true;
}
bool tl2::details::AntispamPatternFullReadBoxed(::basictl::tl_istream & s, ::tl2::antispam::PatternFull& item) noexcept {
	uint32_t nat;
	if (!s.nat_read(nat)) { return false; }
	switch (nat) {
	case 0xa7688492:
		if (item.value.index() != 0) { item.value.emplace<0>(); }
		if (!::tl2::details::AntispamPatternFoundRead(s, std::get<0>(item.value))) { return s.set_error_unknown_scenario(); }
		break;
	case 0x2c22e225:
		if (item.value.index() != 1) { item.value.emplace<1>(); }
		break;
	default:
		return s.set_error_union_tag();
    }
	return true;
}

bool tl2::details::AntispamPatternFullWriteBoxed(::basictl::tl_ostream & s, const ::tl2::antispam::PatternFull& item) noexcept{
	if (!s.nat_write(AntispamPatternFull_tbl_tl_tag[item.value.index()])) { return false; }
	switch (item.value.index()) {
	case 0:
		if (!::tl2::details::AntispamPatternFoundWrite(s, std::get<0>(item.value))) { return s.set_error_unknown_scenario(); }
		break;
	}
	return true;
}

bool tl2::antispam::PatternNotFound::write_json(std::ostream& s)const {
	if (!::tl2::details::AntispamPatternNotFoundWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::antispam::PatternNotFound::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::AntispamPatternNotFoundRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::antispam::PatternNotFound::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::AntispamPatternNotFoundWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::antispam::PatternNotFound::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::antispam::PatternNotFound::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::antispam::PatternNotFound::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::AntispamPatternNotFoundReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::antispam::PatternNotFound::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::AntispamPatternNotFoundWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::antispam::PatternNotFound::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::antispam::PatternNotFound::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::AntispamPatternNotFoundReset(::tl2::antispam::PatternNotFound& item) noexcept {
}

bool tl2::details::AntispamPatternNotFoundWriteJSON(std::ostream& s, const ::tl2::antispam::PatternNotFound& item) noexcept {
	s << "true";
	return true;
}

bool tl2::details::AntispamPatternNotFoundRead(::basictl::tl_istream & s, ::tl2::antispam::PatternNotFound& item) noexcept {
	return true;
}

bool tl2::details::AntispamPatternNotFoundWrite(::basictl::tl_ostream & s, const ::tl2::antispam::PatternNotFound& item) noexcept {
	return true;
}

bool tl2::details::AntispamPatternNotFoundReadBoxed(::basictl::tl_istream & s, ::tl2::antispam::PatternNotFound& item) {
	if (!s.nat_read_exact_tag(0x2c22e225)) { return false; }
	return tl2::details::AntispamPatternNotFoundRead(s, item);
}

bool tl2::details::AntispamPatternNotFoundWriteBoxed(::basictl::tl_ostream & s, const ::tl2::antispam::PatternNotFound& item) {
	if (!s.nat_write(0x2c22e225)) { return false; }
	return tl2::details::AntispamPatternNotFoundWrite(s, item);
}
