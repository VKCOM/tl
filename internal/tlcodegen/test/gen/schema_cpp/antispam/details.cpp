#include "headers/antispam.PatternFull.hpp"
#include "headers/antispam.patternNotFound.hpp"
#include "headers/antispam.patternFound.hpp"
#include "headers/antispam.getPattern.hpp"


bool tl2::antispam::GetPattern::write_json(std::ostream& s)const {
	if (!::tl2::details::AntispamGetPatternWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::antispam::GetPattern::read(::basictl::tl_istream & s) {
	if (!::tl2::details::AntispamGetPatternRead(s, *this)) { return false; }
	return true;
}

bool tl2::antispam::GetPattern::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::AntispamGetPatternWrite(s, *this)) { return false; }
	return true;
}

bool tl2::antispam::GetPattern::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::AntispamGetPatternReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::antispam::GetPattern::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::AntispamGetPatternWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::AntispamGetPatternReset(::tl2::antispam::GetPattern& item) {
	item.id = 0;
}

bool tl2::details::AntispamGetPatternWriteJSON(std::ostream& s, const ::tl2::antispam::GetPattern& item) {
	s << "{";
	if (item.id != 0) {
		s << "\"id\":";
		s << item.id;
	}
	s << "}";
	return true;
}

bool tl2::details::AntispamGetPatternRead(::basictl::tl_istream & s, ::tl2::antispam::GetPattern& item) {
	if (!s.int_read(item.id)) { return false; }
	return true;
}

bool tl2::details::AntispamGetPatternWrite(::basictl::tl_ostream & s, const ::tl2::antispam::GetPattern& item) {
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
	if (!::tl2::details::AntispamPatternFullReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::AntispamGetPatternWriteResult(::basictl::tl_ostream & s, tl2::antispam::GetPattern& item, ::tl2::antispam::PatternFull& result) {
	if (!::tl2::details::AntispamPatternFullWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::antispam::GetPattern::read_result(::basictl::tl_istream & s, ::tl2::antispam::PatternFull & result) {
	return tl2::details::AntispamGetPatternReadResult(s, *this, result);
}
bool tl2::antispam::GetPattern::write_result(::basictl::tl_ostream & s, ::tl2::antispam::PatternFull & result) {
	return tl2::details::AntispamGetPatternWriteResult(s, *this, result);
}

bool tl2::antispam::PatternFound::write_json(std::ostream& s)const {
	if (!::tl2::details::AntispamPatternFoundWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::antispam::PatternFound::read(::basictl::tl_istream & s) {
	if (!::tl2::details::AntispamPatternFoundRead(s, *this)) { return false; }
	return true;
}

bool tl2::antispam::PatternFound::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::AntispamPatternFoundWrite(s, *this)) { return false; }
	return true;
}

bool tl2::antispam::PatternFound::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::AntispamPatternFoundReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::antispam::PatternFound::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::AntispamPatternFoundWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::AntispamPatternFoundReset(::tl2::antispam::PatternFound& item) {
	item.ip = 0;
	item.uahash = 0;
	item.flags = 0;
	item.type = 0;
	item.text.clear();
}

bool tl2::details::AntispamPatternFoundWriteJSON(std::ostream& s, const ::tl2::antispam::PatternFound& item) {
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

bool tl2::details::AntispamPatternFoundRead(::basictl::tl_istream & s, ::tl2::antispam::PatternFound& item) {
	if (!s.int_read(item.ip)) { return false; }
	if (!s.int_read(item.uahash)) { return false; }
	if (!s.int_read(item.flags)) { return false; }
	if (!s.int_read(item.type)) { return false; }
	if (!s.string_read(item.text)) { return false; }
	return true;
}

bool tl2::details::AntispamPatternFoundWrite(::basictl::tl_ostream & s, const ::tl2::antispam::PatternFound& item) {
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
bool tl2::antispam::PatternFull::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::AntispamPatternFullReadBoxed(s, *this)) { return false; }
	return true;
}
bool tl2::antispam::PatternFull::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::AntispamPatternFullWriteBoxed(s, *this)) { return false; }
	return true;
}
std::string_view tl2::antispam::PatternFull::tl_name() const {
	return AntispamPatternFull_tbl_tl_name[value.index()];
}
uint32_t tl2::antispam::PatternFull::tl_tag() const {
	return AntispamPatternFull_tbl_tl_tag[value.index()];
}


void tl2::details::AntispamPatternFullReset(::tl2::antispam::PatternFull& item) {
	item.value.emplace<0>(); // TODO - optimize, if already 0, call Reset function
}

bool tl2::details::AntispamPatternFullWriteJSON(std::ostream & s, const ::tl2::antispam::PatternFull& item) {
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
bool tl2::details::AntispamPatternFullReadBoxed(::basictl::tl_istream & s, ::tl2::antispam::PatternFull& item) {
	uint32_t nat;
	s.nat_read(nat);
	switch (nat) {
	case 0xa7688492:
		if (item.value.index() != 0) { item.value.emplace<0>(); }
		if (!::tl2::details::AntispamPatternFoundRead(s, std::get<0>(item.value))) { return false; }
		break;
	case 0x2c22e225:
		if (item.value.index() != 1) { item.value.emplace<1>(); }
		break;
	default:
		return s.set_error_union_tag();
    }
	return true;
}

bool tl2::details::AntispamPatternFullWriteBoxed(::basictl::tl_ostream & s, const ::tl2::antispam::PatternFull& item) {
	s.nat_write(AntispamPatternFull_tbl_tl_tag[item.value.index()]);
	switch (item.value.index()) {
	case 0:
		if (!::tl2::details::AntispamPatternFoundWrite(s, std::get<0>(item.value))) { return false; }
		break;
	}
	return true;
}

bool tl2::antispam::PatternNotFound::write_json(std::ostream& s)const {
	if (!::tl2::details::AntispamPatternNotFoundWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::antispam::PatternNotFound::read(::basictl::tl_istream & s) {
	if (!::tl2::details::AntispamPatternNotFoundRead(s, *this)) { return false; }
	return true;
}

bool tl2::antispam::PatternNotFound::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::AntispamPatternNotFoundWrite(s, *this)) { return false; }
	return true;
}

bool tl2::antispam::PatternNotFound::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::AntispamPatternNotFoundReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::antispam::PatternNotFound::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::AntispamPatternNotFoundWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::AntispamPatternNotFoundReset(::tl2::antispam::PatternNotFound& item) {
}

bool tl2::details::AntispamPatternNotFoundWriteJSON(std::ostream& s, const ::tl2::antispam::PatternNotFound& item) {
	s << "true";
	return true;
}

bool tl2::details::AntispamPatternNotFoundRead(::basictl::tl_istream & s, ::tl2::antispam::PatternNotFound& item) {
	return true;
}

bool tl2::details::AntispamPatternNotFoundWrite(::basictl::tl_ostream & s, const ::tl2::antispam::PatternNotFound& item) {
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
