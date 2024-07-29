#include "headers/service6.multiFindWithBounds.hpp"
#include "headers/service6.multiFind.hpp"
#include "headers/service6.findWithBoundsResult.hpp"
#include "headers/service6.findResultRow.hpp"
#include "headers/service6.error.hpp"
#include "../__common_namespace/headers/int.hpp"
#include "../__common_namespace/headers/Either.hpp"


void tl2::details::BuiltinVectorService6FindResultRowReset(std::vector<::tl2::service6::FindResultRow>& item) {
	item.resize(0); // TODO - unwrap
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

bool tl2::service6::Error::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service6ErrorRead(s, *this)) { return false; }
	return true;
}

bool tl2::service6::Error::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service6ErrorWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service6::Error::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service6ErrorReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service6::Error::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service6ErrorWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service6ErrorReset(::tl2::service6::Error& item) {
	item.code = 0;
}

bool tl2::details::Service6ErrorRead(::basictl::tl_istream & s, ::tl2::service6::Error& item) {
	if (!s.int_read(item.code)) { return false; }
	return true;
}

bool tl2::details::Service6ErrorWrite(::basictl::tl_ostream & s, const ::tl2::service6::Error& item) {
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

bool tl2::service6::FindResultRow::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service6FindResultRowRead(s, *this)) { return false; }
	return true;
}

bool tl2::service6::FindResultRow::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service6FindResultRowWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service6::FindResultRow::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service6FindResultRowReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service6::FindResultRow::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service6FindResultRowWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service6FindResultRowReset(::tl2::service6::FindResultRow& item) {
	item.x = 0;
}

bool tl2::details::Service6FindResultRowRead(::basictl::tl_istream & s, ::tl2::service6::FindResultRow& item) {
	if (!s.int_read(item.x)) { return false; }
	return true;
}

bool tl2::details::Service6FindResultRowWrite(::basictl::tl_ostream & s, const ::tl2::service6::FindResultRow& item) {
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

bool tl2::service6::FindWithBoundsResult::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service6FindWithBoundsResultRead(s, *this)) { return false; }
	return true;
}

bool tl2::service6::FindWithBoundsResult::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service6FindWithBoundsResultWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service6::FindWithBoundsResult::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service6FindWithBoundsResultReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service6::FindWithBoundsResult::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service6FindWithBoundsResultWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service6FindWithBoundsResultReset(::tl2::service6::FindWithBoundsResult& item) {
	item.x = 0;
}

bool tl2::details::Service6FindWithBoundsResultRead(::basictl::tl_istream & s, ::tl2::service6::FindWithBoundsResult& item) {
	if (!s.int_read(item.x)) { return false; }
	return true;
}

bool tl2::details::Service6FindWithBoundsResultWrite(::basictl::tl_ostream & s, const ::tl2::service6::FindWithBoundsResult& item) {
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

bool tl2::service6::MultiFind::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service6MultiFindRead(s, *this)) { return false; }
	return true;
}

bool tl2::service6::MultiFind::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service6MultiFindWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service6::MultiFind::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service6MultiFindReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service6::MultiFind::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service6MultiFindWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service6MultiFindReset(::tl2::service6::MultiFind& item) {
	item.clusters.clear();
	item.limit = 0;
	item.eq_threshold = 0;
}

bool tl2::details::Service6MultiFindRead(::basictl::tl_istream & s, ::tl2::service6::MultiFind& item) {
	if (!::tl2::details::BuiltinVectorIntRead(s, item.clusters)) { return false; }
	if (!s.int_read(item.limit)) { return false; }
	if (!s.double_read(item.eq_threshold)) { return false; }
	return true;
}

bool tl2::details::Service6MultiFindWrite(::basictl::tl_ostream & s, const ::tl2::service6::MultiFind& item) {
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

bool tl2::service6::MultiFind::read_result(::basictl::tl_istream & s, std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>> & result) {
	return tl2::details::Service6MultiFindReadResult(s, *this, result);
}
bool tl2::service6::MultiFind::write_result(::basictl::tl_ostream & s, std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>> & result) {
	return tl2::details::Service6MultiFindWriteResult(s, *this, result);
}

bool tl2::service6::MultiFindWithBounds::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service6MultiFindWithBoundsRead(s, *this)) { return false; }
	return true;
}

bool tl2::service6::MultiFindWithBounds::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service6MultiFindWithBoundsWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service6::MultiFindWithBounds::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service6MultiFindWithBoundsReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service6::MultiFindWithBounds::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service6MultiFindWithBoundsWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service6MultiFindWithBoundsReset(::tl2::service6::MultiFindWithBounds& item) {
	item.clusters.clear();
}

bool tl2::details::Service6MultiFindWithBoundsRead(::basictl::tl_istream & s, ::tl2::service6::MultiFindWithBounds& item) {
	if (!::tl2::details::BuiltinVectorIntRead(s, item.clusters)) { return false; }
	return true;
}

bool tl2::details::Service6MultiFindWithBoundsWrite(::basictl::tl_ostream & s, const ::tl2::service6::MultiFindWithBounds& item) {
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

bool tl2::service6::MultiFindWithBounds::read_result(::basictl::tl_istream & s, std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>> & result) {
	return tl2::details::Service6MultiFindWithBoundsReadResult(s, *this, result);
}
bool tl2::service6::MultiFindWithBounds::write_result(::basictl::tl_ostream & s, std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>> & result) {
	return tl2::details::Service6MultiFindWithBoundsWriteResult(s, *this, result);
}
