#include "headers/tree_stats.ObjectLimitValue.hpp"
#include "headers/tree_stats.objectLimitValueLong.hpp"
#include "headers/tree_stats.objectLimitValueDouble.hpp"


static const std::string_view TreeStatsObjectLimitValue_tbl_tl_name[]{"tree_stats.objectLimitValueLong", "tree_stats.objectLimitValueDouble"};
static const uint32_t TreeStatsObjectLimitValue_tbl_tl_tag[]{0x73111993, 0x5dfb8816};

bool tl2::tree_stats::ObjectLimitValue::write_json(std::ostream & s)const {
	if (!::tl2::details::TreeStatsObjectLimitValueWriteJSON(s, *this)) { return false; }
	return true;
}
bool tl2::tree_stats::ObjectLimitValue::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::TreeStatsObjectLimitValueReadBoxed(s, *this)) { return false; }
	return true;
}
bool tl2::tree_stats::ObjectLimitValue::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::TreeStatsObjectLimitValueWriteBoxed(s, *this)) { return false; }
	return true;
}
std::string_view tl2::tree_stats::ObjectLimitValue::tl_name() const {
	return TreeStatsObjectLimitValue_tbl_tl_name[value.index()];
}
uint32_t tl2::tree_stats::ObjectLimitValue::tl_tag() const {
	return TreeStatsObjectLimitValue_tbl_tl_tag[value.index()];
}


void tl2::details::TreeStatsObjectLimitValueReset(::tl2::tree_stats::ObjectLimitValue& item) {
	item.value.emplace<0>(); // TODO - optimize, if already 0, call Reset function
}

bool tl2::details::TreeStatsObjectLimitValueWriteJSON(std::ostream & s, const ::tl2::tree_stats::ObjectLimitValue& item) {
	s << "{";
	s << "\"type\":";
	s << TreeStatsObjectLimitValue_tbl_tl_tag[item.value.index()];
	switch (item.value.index()) {
	case 1:
		s << ",\"value\":";
		if (!::tl2::details::TreeStatsObjectLimitValueDoubleWriteJSON(s, std::get<1>(item.value))) { return false; }
		break;
	}
	s << "}";
	return true;
}
bool tl2::details::TreeStatsObjectLimitValueReadBoxed(::basictl::tl_istream & s, ::tl2::tree_stats::ObjectLimitValue& item) {
	uint32_t nat;
	s.nat_read(nat);
	switch (nat) {
	case 0x73111993:
		if (item.value.index() != 0) { item.value.emplace<0>(); }
		break;
	case 0x5dfb8816:
		if (item.value.index() != 1) { item.value.emplace<1>(); }
		if (!::tl2::details::TreeStatsObjectLimitValueDoubleRead(s, std::get<1>(item.value))) { return false; }
		break;
	default:
		return s.set_error_union_tag();
    }
	return true;
}

bool tl2::details::TreeStatsObjectLimitValueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tree_stats::ObjectLimitValue& item) {
	s.nat_write(TreeStatsObjectLimitValue_tbl_tl_tag[item.value.index()]);
	switch (item.value.index()) {
	case 1:
		if (!::tl2::details::TreeStatsObjectLimitValueDoubleWrite(s, std::get<1>(item.value))) { return false; }
		break;
	}
	return true;
}

void tl2::details::TreeStatsObjectLimitValueDoubleReset(::tl2::tree_stats::ObjectLimitValueDouble& item) {
	item = 0;
}

bool tl2::details::TreeStatsObjectLimitValueDoubleWriteJSON(std::ostream& s, const ::tl2::tree_stats::ObjectLimitValueDouble& item) {
	s << item;
	return true;
}

bool tl2::details::TreeStatsObjectLimitValueDoubleRead(::basictl::tl_istream & s, ::tl2::tree_stats::ObjectLimitValueDouble& item) {
	if (!s.double_read(item)) { return false; }
	return true;
}

bool tl2::details::TreeStatsObjectLimitValueDoubleWrite(::basictl::tl_ostream & s, const ::tl2::tree_stats::ObjectLimitValueDouble& item) {
	if (!s.double_write(item)) { return false;}
	return true;
}

bool tl2::details::TreeStatsObjectLimitValueDoubleReadBoxed(::basictl::tl_istream & s, ::tl2::tree_stats::ObjectLimitValueDouble& item) {
	if (!s.nat_read_exact_tag(0x5dfb8816)) { return false; }
	return tl2::details::TreeStatsObjectLimitValueDoubleRead(s, item);
}

bool tl2::details::TreeStatsObjectLimitValueDoubleWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tree_stats::ObjectLimitValueDouble& item) {
	if (!s.nat_write(0x5dfb8816)) { return false; }
	return tl2::details::TreeStatsObjectLimitValueDoubleWrite(s, item);
}

bool tl2::tree_stats::ObjectLimitValueLong::write_json(std::ostream& s)const {
	if (!::tl2::details::TreeStatsObjectLimitValueLongWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::tree_stats::ObjectLimitValueLong::read(::basictl::tl_istream & s) {
	if (!::tl2::details::TreeStatsObjectLimitValueLongRead(s, *this)) { return false; }
	return true;
}

bool tl2::tree_stats::ObjectLimitValueLong::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::TreeStatsObjectLimitValueLongWrite(s, *this)) { return false; }
	return true;
}

bool tl2::tree_stats::ObjectLimitValueLong::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::TreeStatsObjectLimitValueLongReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::tree_stats::ObjectLimitValueLong::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::TreeStatsObjectLimitValueLongWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::TreeStatsObjectLimitValueLongReset(::tl2::tree_stats::ObjectLimitValueLong& item) {
}

bool tl2::details::TreeStatsObjectLimitValueLongWriteJSON(std::ostream& s, const ::tl2::tree_stats::ObjectLimitValueLong& item) {
	s << "{";
	s << "}";
	return true;
}

bool tl2::details::TreeStatsObjectLimitValueLongRead(::basictl::tl_istream & s, ::tl2::tree_stats::ObjectLimitValueLong& item) {
	return true;
}

bool tl2::details::TreeStatsObjectLimitValueLongWrite(::basictl::tl_ostream & s, const ::tl2::tree_stats::ObjectLimitValueLong& item) {
	return true;
}

bool tl2::details::TreeStatsObjectLimitValueLongReadBoxed(::basictl::tl_istream & s, ::tl2::tree_stats::ObjectLimitValueLong& item) {
	if (!s.nat_read_exact_tag(0x73111993)) { return false; }
	return tl2::details::TreeStatsObjectLimitValueLongRead(s, item);
}

bool tl2::details::TreeStatsObjectLimitValueLongWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tree_stats::ObjectLimitValueLong& item) {
	if (!s.nat_write(0x73111993)) { return false; }
	return tl2::details::TreeStatsObjectLimitValueLongWrite(s, item);
}
