#include "headers/service6_vector.h"
#include "headers/service6.multiFindWithBounds.h"
#include "headers/service6.multiFind.h"
#include "headers/service6.findWithBoundsResult.h"
#include "headers/service6.findResultRow.h"
#include "headers/service6.error.h"
#include "headers/service6_Either.h"
#include "headers/service6_right.h"
#include "headers/service6_left.h"
#include "../__common_namespace/headers/Either.h"
#include "../__common_namespace/headers/int.h"


void tl2::details::BuiltinVectorEitherIntVectorService6FindWithBoundsResultReset(std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item) {
	item.resize(0); // TODO - unwrap
}

bool tl2::details::BuiltinVectorEitherIntVectorService6FindWithBoundsResultWriteJSON(std::ostream & s, const std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item) {
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		if (!::tl2::details::EitherIntVectorService6FindWithBoundsResultWriteJSON(s, el)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinVectorEitherIntVectorService6FindWithBoundsResultRead(::basictl::tl_istream & s, std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	// TODO - check length sanity
	item.resize(len);
	for(auto && el : item) {
		if (!::tl2::details::EitherIntVectorService6FindWithBoundsResultReadBoxed(s, el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinVectorEitherIntVectorService6FindWithBoundsResultWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!::tl2::details::EitherIntVectorService6FindWithBoundsResultWriteBoxed(s, el)) { return false; }
	}
	return true;
}

void tl2::details::BuiltinVectorService6FindResultRowReset(std::vector<::tl2::service6::FindResultRow>& item) {
	item.resize(0); // TODO - unwrap
}

bool tl2::details::BuiltinVectorService6FindResultRowWriteJSON(std::ostream & s, const std::vector<::tl2::service6::FindResultRow>& item) {
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		if (!::tl2::details::Service6FindResultRowWriteJSON(s, el)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinVectorService6FindResultRowRead(::basictl::tl_istream & s, std::vector<::tl2::service6::FindResultRow>& item) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	// TODO - check length sanity
	item.resize(len);
	for(auto && el : item) {
		if (!::tl2::details::Service6FindResultRowRead(s, el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinVectorService6FindResultRowWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service6::FindResultRow>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!::tl2::details::Service6FindResultRowWrite(s, el)) { return false; }
	}
	return true;
}

void tl2::details::BuiltinVectorService6FindWithBoundsResultReset(std::vector<::tl2::service6::FindWithBoundsResult>& item) {
	item.resize(0); // TODO - unwrap
}

bool tl2::details::BuiltinVectorService6FindWithBoundsResultWriteJSON(std::ostream & s, const std::vector<::tl2::service6::FindWithBoundsResult>& item) {
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		if (!::tl2::details::Service6FindWithBoundsResultWriteJSON(s, el)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinVectorService6FindWithBoundsResultRead(::basictl::tl_istream & s, std::vector<::tl2::service6::FindWithBoundsResult>& item) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	// TODO - check length sanity
	item.resize(len);
	for(auto && el : item) {
		if (!::tl2::details::Service6FindWithBoundsResultRead(s, el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinVectorService6FindWithBoundsResultWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service6::FindWithBoundsResult>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!::tl2::details::Service6FindWithBoundsResultWrite(s, el)) { return false; }
	}
	return true;
}

static const std::string_view EitherIntVectorService6FindWithBoundsResult_tbl_tl_name[]{"left", "right"};
static const uint32_t EitherIntVectorService6FindWithBoundsResult_tbl_tl_tag[]{0x0a29cd5d, 0xdf3ecb3b};

void tl2::details::EitherIntVectorService6FindWithBoundsResultReset(::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) noexcept{
	item.value.emplace<0>(); // TODO - optimize, if already 0, call Reset function
}

bool tl2::details::EitherIntVectorService6FindWithBoundsResultWriteJSON(std::ostream & s, const ::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) noexcept {
	s << "{";
	s << "\"type\":";
	s << "\"" << EitherIntVectorService6FindWithBoundsResult_tbl_tl_name[item.value.index()] << "\"";
	switch (item.value.index()) {
	case 0:
		s << ",\"value\":";
		if (!::tl2::details::LeftIntVectorService6FindWithBoundsResultWriteJSON(s, std::get<0>(item.value))) { return false; }
		break;
	case 1:
		s << ",\"value\":";
		if (!::tl2::details::RightIntVectorService6FindWithBoundsResultWriteJSON(s, std::get<1>(item.value))) { return false; }
		break;
	}
	s << "}";
	return true;
}
bool tl2::details::EitherIntVectorService6FindWithBoundsResultReadBoxed(::basictl::tl_istream & s, ::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) noexcept {
	uint32_t nat;
	s.nat_read(nat);
	switch (nat) {
	case 0x0a29cd5d:
		if (item.value.index() != 0) { item.value.emplace<0>(); }
		if (!::tl2::details::LeftIntVectorService6FindWithBoundsResultRead(s, std::get<0>(item.value))) { return false; }
		break;
	case 0xdf3ecb3b:
		if (item.value.index() != 1) { item.value.emplace<1>(); }
		if (!::tl2::details::RightIntVectorService6FindWithBoundsResultRead(s, std::get<1>(item.value))) { return false; }
		break;
	default:
		return s.set_error_union_tag();
    }
	return true;
}

bool tl2::details::EitherIntVectorService6FindWithBoundsResultWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) noexcept{
	s.nat_write(EitherIntVectorService6FindWithBoundsResult_tbl_tl_tag[item.value.index()]);
	switch (item.value.index()) {
	case 0:
		if (!::tl2::details::LeftIntVectorService6FindWithBoundsResultWrite(s, std::get<0>(item.value))) { return false; }
		break;
	case 1:
		if (!::tl2::details::RightIntVectorService6FindWithBoundsResultWrite(s, std::get<1>(item.value))) { return false; }
		break;
	}
	return true;
}

void tl2::details::LeftIntVectorService6FindWithBoundsResultReset(::tl2::Left<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) noexcept {
	item.value = 0;
}

bool tl2::details::LeftIntVectorService6FindWithBoundsResultWriteJSON(std::ostream& s, const ::tl2::Left<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) noexcept {
	s << "{";
	if (item.value != 0) {
		s << "\"value\":";
		s << item.value;
	}
	s << "}";
	return true;
}

bool tl2::details::LeftIntVectorService6FindWithBoundsResultRead(::basictl::tl_istream & s, ::tl2::Left<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) noexcept {
	if (!s.int_read(item.value)) { return false; }
	return true;
}

bool tl2::details::LeftIntVectorService6FindWithBoundsResultWrite(::basictl::tl_ostream & s, const ::tl2::Left<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) noexcept {
	if (!s.int_write(item.value)) { return false;}
	return true;
}

bool tl2::details::LeftIntVectorService6FindWithBoundsResultReadBoxed(::basictl::tl_istream & s, ::tl2::Left<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) {
	if (!s.nat_read_exact_tag(0x0a29cd5d)) { return false; }
	return tl2::details::LeftIntVectorService6FindWithBoundsResultRead(s, item);
}

bool tl2::details::LeftIntVectorService6FindWithBoundsResultWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Left<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) {
	if (!s.nat_write(0x0a29cd5d)) { return false; }
	return tl2::details::LeftIntVectorService6FindWithBoundsResultWrite(s, item);
}

void tl2::details::RightIntVectorService6FindWithBoundsResultReset(::tl2::Right<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) noexcept {
	item.value.clear();
}

bool tl2::details::RightIntVectorService6FindWithBoundsResultWriteJSON(std::ostream& s, const ::tl2::Right<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) noexcept {
	s << "{";
	if (item.value.size() != 0) {
		s << "\"value\":";
		if (!::tl2::details::BuiltinVectorService6FindWithBoundsResultWriteJSON(s, item.value)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::RightIntVectorService6FindWithBoundsResultRead(::basictl::tl_istream & s, ::tl2::Right<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) noexcept {
	if (!::tl2::details::BuiltinVectorService6FindWithBoundsResultRead(s, item.value)) { return false; }
	return true;
}

bool tl2::details::RightIntVectorService6FindWithBoundsResultWrite(::basictl::tl_ostream & s, const ::tl2::Right<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) noexcept {
	if (!::tl2::details::BuiltinVectorService6FindWithBoundsResultWrite(s, item.value)) { return false; }
	return true;
}

bool tl2::details::RightIntVectorService6FindWithBoundsResultReadBoxed(::basictl::tl_istream & s, ::tl2::Right<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) {
	if (!s.nat_read_exact_tag(0xdf3ecb3b)) { return false; }
	return tl2::details::RightIntVectorService6FindWithBoundsResultRead(s, item);
}

bool tl2::details::RightIntVectorService6FindWithBoundsResultWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Right<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) {
	if (!s.nat_write(0xdf3ecb3b)) { return false; }
	return tl2::details::RightIntVectorService6FindWithBoundsResultWrite(s, item);
}

bool tl2::service6::Error::write_json(std::ostream& s)const {
	if (!::tl2::details::Service6ErrorWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service6::Error::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service6ErrorRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::service6::Error::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service6ErrorWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::service6::Error::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service6::Error::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service6::Error::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service6ErrorReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::service6::Error::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service6ErrorWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::service6::Error::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service6::Error::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service6ErrorReset(::tl2::service6::Error& item) noexcept {
	item.code = 0;
}

bool tl2::details::Service6ErrorWriteJSON(std::ostream& s, const ::tl2::service6::Error& item) noexcept {
	s << "{";
	if (item.code != 0) {
		s << "\"code\":";
		s << item.code;
	}
	s << "}";
	return true;
}

bool tl2::details::Service6ErrorRead(::basictl::tl_istream & s, ::tl2::service6::Error& item) noexcept {
	if (!s.int_read(item.code)) { return false; }
	return true;
}

bool tl2::details::Service6ErrorWrite(::basictl::tl_ostream & s, const ::tl2::service6::Error& item) noexcept {
	if (!s.int_write(item.code)) { return false;}
	return true;
}

bool tl2::details::Service6ErrorReadBoxed(::basictl::tl_istream & s, ::tl2::service6::Error& item) {
	if (!s.nat_read_exact_tag(0x738553ef)) { return false; }
	return tl2::details::Service6ErrorRead(s, item);
}

bool tl2::details::Service6ErrorWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service6::Error& item) {
	if (!s.nat_write(0x738553ef)) { return false; }
	return tl2::details::Service6ErrorWrite(s, item);
}

bool tl2::service6::FindResultRow::write_json(std::ostream& s)const {
	if (!::tl2::details::Service6FindResultRowWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service6::FindResultRow::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service6FindResultRowRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::service6::FindResultRow::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service6FindResultRowWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::service6::FindResultRow::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service6::FindResultRow::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service6::FindResultRow::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service6FindResultRowReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::service6::FindResultRow::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service6FindResultRowWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::service6::FindResultRow::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service6::FindResultRow::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service6FindResultRowReset(::tl2::service6::FindResultRow& item) noexcept {
	item.x = 0;
}

bool tl2::details::Service6FindResultRowWriteJSON(std::ostream& s, const ::tl2::service6::FindResultRow& item) noexcept {
	s << "{";
	if (item.x != 0) {
		s << "\"x\":";
		s << item.x;
	}
	s << "}";
	return true;
}

bool tl2::details::Service6FindResultRowRead(::basictl::tl_istream & s, ::tl2::service6::FindResultRow& item) noexcept {
	if (!s.int_read(item.x)) { return false; }
	return true;
}

bool tl2::details::Service6FindResultRowWrite(::basictl::tl_ostream & s, const ::tl2::service6::FindResultRow& item) noexcept {
	if (!s.int_write(item.x)) { return false;}
	return true;
}

bool tl2::details::Service6FindResultRowReadBoxed(::basictl::tl_istream & s, ::tl2::service6::FindResultRow& item) {
	if (!s.nat_read_exact_tag(0xbd3946e3)) { return false; }
	return tl2::details::Service6FindResultRowRead(s, item);
}

bool tl2::details::Service6FindResultRowWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service6::FindResultRow& item) {
	if (!s.nat_write(0xbd3946e3)) { return false; }
	return tl2::details::Service6FindResultRowWrite(s, item);
}

bool tl2::service6::FindWithBoundsResult::write_json(std::ostream& s)const {
	if (!::tl2::details::Service6FindWithBoundsResultWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service6::FindWithBoundsResult::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service6FindWithBoundsResultRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::service6::FindWithBoundsResult::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service6FindWithBoundsResultWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::service6::FindWithBoundsResult::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service6::FindWithBoundsResult::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service6::FindWithBoundsResult::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service6FindWithBoundsResultReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::service6::FindWithBoundsResult::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service6FindWithBoundsResultWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::service6::FindWithBoundsResult::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service6::FindWithBoundsResult::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service6FindWithBoundsResultReset(::tl2::service6::FindWithBoundsResult& item) noexcept {
	item.x = 0;
}

bool tl2::details::Service6FindWithBoundsResultWriteJSON(std::ostream& s, const ::tl2::service6::FindWithBoundsResult& item) noexcept {
	s << "{";
	if (item.x != 0) {
		s << "\"x\":";
		s << item.x;
	}
	s << "}";
	return true;
}

bool tl2::details::Service6FindWithBoundsResultRead(::basictl::tl_istream & s, ::tl2::service6::FindWithBoundsResult& item) noexcept {
	if (!s.int_read(item.x)) { return false; }
	return true;
}

bool tl2::details::Service6FindWithBoundsResultWrite(::basictl::tl_ostream & s, const ::tl2::service6::FindWithBoundsResult& item) noexcept {
	if (!s.int_write(item.x)) { return false;}
	return true;
}

bool tl2::details::Service6FindWithBoundsResultReadBoxed(::basictl::tl_istream & s, ::tl2::service6::FindWithBoundsResult& item) {
	if (!s.nat_read_exact_tag(0x3ded850a)) { return false; }
	return tl2::details::Service6FindWithBoundsResultRead(s, item);
}

bool tl2::details::Service6FindWithBoundsResultWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service6::FindWithBoundsResult& item) {
	if (!s.nat_write(0x3ded850a)) { return false; }
	return tl2::details::Service6FindWithBoundsResultWrite(s, item);
}

bool tl2::service6::MultiFind::write_json(std::ostream& s)const {
	if (!::tl2::details::Service6MultiFindWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service6::MultiFind::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service6MultiFindRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::service6::MultiFind::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service6MultiFindWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::service6::MultiFind::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service6::MultiFind::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service6::MultiFind::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service6MultiFindReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::service6::MultiFind::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service6MultiFindWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::service6::MultiFind::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service6::MultiFind::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service6MultiFindReset(::tl2::service6::MultiFind& item) noexcept {
	item.clusters.clear();
	item.limit = 0;
	item.eq_threshold = 0;
}

bool tl2::details::Service6MultiFindWriteJSON(std::ostream& s, const ::tl2::service6::MultiFind& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.clusters.size() != 0) {
		add_comma = true;
		s << "\"clusters\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.clusters)) { return false; }
	}
	if (item.limit != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"limit\":";
		s << item.limit;
	}
	if (item.eq_threshold != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"eq_threshold\":";
		s << item.eq_threshold;
	}
	s << "}";
	return true;
}

bool tl2::details::Service6MultiFindRead(::basictl::tl_istream & s, ::tl2::service6::MultiFind& item) noexcept {
	if (!::tl2::details::BuiltinVectorIntRead(s, item.clusters)) { return false; }
	if (!s.int_read(item.limit)) { return false; }
	if (!s.double_read(item.eq_threshold)) { return false; }
	return true;
}

bool tl2::details::Service6MultiFindWrite(::basictl::tl_ostream & s, const ::tl2::service6::MultiFind& item) noexcept {
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.clusters)) { return false; }
	if (!s.int_write(item.limit)) { return false;}
	if (!s.double_write(item.eq_threshold)) { return false;}
	return true;
}

bool tl2::details::Service6MultiFindReadBoxed(::basictl::tl_istream & s, ::tl2::service6::MultiFind& item) {
	if (!s.nat_read_exact_tag(0xe62178d8)) { return false; }
	return tl2::details::Service6MultiFindRead(s, item);
}

bool tl2::details::Service6MultiFindWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service6::MultiFind& item) {
	if (!s.nat_write(0xe62178d8)) { return false; }
	return tl2::details::Service6MultiFindWrite(s, item);
}

bool tl2::details::Service6MultiFindReadResult(::basictl::tl_istream & s, tl2::service6::MultiFind& item, std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& result) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false;}
	if (!::tl2::details::BuiltinVectorEitherService6ErrorVectorService6FindResultRowRead(s, result)) { return false; }
	return true;
}
bool tl2::details::Service6MultiFindWriteResult(::basictl::tl_ostream & s, tl2::service6::MultiFind& item, std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& result) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorEitherService6ErrorVectorService6FindResultRowWrite(s, result)) { return false; }
	return true;
}

bool tl2::service6::MultiFind::read_result(::basictl::tl_istream & s, std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>> & result) noexcept {
	bool success = tl2::details::Service6MultiFindReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service6::MultiFind::write_result(::basictl::tl_ostream & s, std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>> & result) noexcept {
	bool success = tl2::details::Service6MultiFindWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service6::MultiFind::read_result_or_throw(::basictl::tl_throwable_istream & s, std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>> & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service6::MultiFind::write_result_or_throw(::basictl::tl_throwable_ostream & s, std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>> & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::service6::MultiFindWithBounds::write_json(std::ostream& s)const {
	if (!::tl2::details::Service6MultiFindWithBoundsWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service6::MultiFindWithBounds::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service6MultiFindWithBoundsRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::service6::MultiFindWithBounds::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service6MultiFindWithBoundsWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::service6::MultiFindWithBounds::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service6::MultiFindWithBounds::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service6::MultiFindWithBounds::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service6MultiFindWithBoundsReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::service6::MultiFindWithBounds::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service6MultiFindWithBoundsWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::service6::MultiFindWithBounds::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service6::MultiFindWithBounds::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service6MultiFindWithBoundsReset(::tl2::service6::MultiFindWithBounds& item) noexcept {
	item.clusters.clear();
}

bool tl2::details::Service6MultiFindWithBoundsWriteJSON(std::ostream& s, const ::tl2::service6::MultiFindWithBounds& item) noexcept {
	s << "{";
	if (item.clusters.size() != 0) {
		s << "\"clusters\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.clusters)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::Service6MultiFindWithBoundsRead(::basictl::tl_istream & s, ::tl2::service6::MultiFindWithBounds& item) noexcept {
	if (!::tl2::details::BuiltinVectorIntRead(s, item.clusters)) { return false; }
	return true;
}

bool tl2::details::Service6MultiFindWithBoundsWrite(::basictl::tl_ostream & s, const ::tl2::service6::MultiFindWithBounds& item) noexcept {
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.clusters)) { return false; }
	return true;
}

bool tl2::details::Service6MultiFindWithBoundsReadBoxed(::basictl::tl_istream & s, ::tl2::service6::MultiFindWithBounds& item) {
	if (!s.nat_read_exact_tag(0x84b168cf)) { return false; }
	return tl2::details::Service6MultiFindWithBoundsRead(s, item);
}

bool tl2::details::Service6MultiFindWithBoundsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service6::MultiFindWithBounds& item) {
	if (!s.nat_write(0x84b168cf)) { return false; }
	return tl2::details::Service6MultiFindWithBoundsWrite(s, item);
}

bool tl2::details::Service6MultiFindWithBoundsReadResult(::basictl::tl_istream & s, tl2::service6::MultiFindWithBounds& item, std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& result) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false;}
	if (!::tl2::details::BuiltinVectorEitherIntVectorService6FindWithBoundsResultRead(s, result)) { return false; }
	return true;
}
bool tl2::details::Service6MultiFindWithBoundsWriteResult(::basictl::tl_ostream & s, tl2::service6::MultiFindWithBounds& item, std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& result) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorEitherIntVectorService6FindWithBoundsResultWrite(s, result)) { return false; }
	return true;
}

bool tl2::service6::MultiFindWithBounds::read_result(::basictl::tl_istream & s, std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>> & result) noexcept {
	bool success = tl2::details::Service6MultiFindWithBoundsReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service6::MultiFindWithBounds::write_result(::basictl::tl_ostream & s, std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>> & result) noexcept {
	bool success = tl2::details::Service6MultiFindWithBoundsWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service6::MultiFindWithBounds::read_result_or_throw(::basictl::tl_throwable_istream & s, std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>> & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service6::MultiFindWithBounds::write_result_or_throw(::basictl::tl_throwable_ostream & s, std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>> & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

void tl2::details::VectorService6FindWithBoundsResultReset(std::vector<::tl2::service6::FindWithBoundsResult>& item) noexcept {
	item.clear();
}

bool tl2::details::VectorService6FindWithBoundsResultWriteJSON(std::ostream& s, const std::vector<::tl2::service6::FindWithBoundsResult>& item) noexcept {
	if (!::tl2::details::BuiltinVectorService6FindWithBoundsResultWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService6FindWithBoundsResultRead(::basictl::tl_istream & s, std::vector<::tl2::service6::FindWithBoundsResult>& item) noexcept {
	if (!::tl2::details::BuiltinVectorService6FindWithBoundsResultRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService6FindWithBoundsResultWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service6::FindWithBoundsResult>& item) noexcept {
	if (!::tl2::details::BuiltinVectorService6FindWithBoundsResultWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService6FindWithBoundsResultReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::service6::FindWithBoundsResult>& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	return tl2::details::VectorService6FindWithBoundsResultRead(s, item);
}

bool tl2::details::VectorService6FindWithBoundsResultWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::service6::FindWithBoundsResult>& item) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	return tl2::details::VectorService6FindWithBoundsResultWrite(s, item);
}
