#include "headers/service3_vector.hpp"
#include "headers/service3.setLimits.hpp"
#include "headers/service3.setLastVisitTimestamp.hpp"
#include "headers/service3.restoreProduct.hpp"
#include "headers/service3.restoreGroupedProducts.hpp"
#include "headers/service3.restoreAllProducts.hpp"
#include "headers/service3.productStatsOld.hpp"
#include "headers/service3.product.hpp"
#include "headers/service3.limits.hpp"
#include "headers/service3.groupSizeLimit.hpp"
#include "headers/service3.groupCountLimit.hpp"
#include "headers/service3.getScheduledProducts.hpp"
#include "headers/service3.getProducts.hpp"
#include "headers/service3.getProductStats.hpp"
#include "headers/service3.getLimits.hpp"
#include "headers/service3.getLastVisitTimestamp.hpp"
#include "headers/service3.deleteProduct.hpp"
#include "headers/service3.deleteGroupedProducts.hpp"
#include "headers/service3.deleteAllProducts.hpp"
#include "headers/service3.createProduct.hpp"
#include "../__common_namespace/headers/int.hpp"
#include "headers/service3_boolStat.hpp"
#include "../__common_namespace/headers/Bool.hpp"


bool tl2::BoolStat::write_json(std::ostream& s)const {
	if (!::tl2::details::BoolStatWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::BoolStat::read(::basictl::tl_istream & s) {
	if (!::tl2::details::BoolStatRead(s, *this)) { return false; }
	return true;
}

bool tl2::BoolStat::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::BoolStatWrite(s, *this)) { return false; }
	return true;
}

bool tl2::BoolStat::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::BoolStatReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::BoolStat::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::BoolStatWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::BoolStatReset(::tl2::BoolStat& item) {
	item.statTrue = 0;
	item.statFalse = 0;
	item.statUnknown = 0;
}

bool tl2::details::BoolStatWriteJSON(std::ostream& s, const ::tl2::BoolStat& item) {
	auto add_comma = false;
	s << "{";
	if (item.statTrue != 0) {
		add_comma = true;
		s << "\"statTrue\":";
		s << item.statTrue;
	}
	if (item.statFalse != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"statFalse\":";
		s << item.statFalse;
	}
	if (item.statUnknown != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"statUnknown\":";
		s << item.statUnknown;
	}
	s << "}";
	return true;
}

bool tl2::details::BoolStatRead(::basictl::tl_istream & s, ::tl2::BoolStat& item) {
	if (!s.int_read(item.statTrue)) { return false; }
	if (!s.int_read(item.statFalse)) { return false; }
	if (!s.int_read(item.statUnknown)) { return false; }
	return true;
}

bool tl2::details::BoolStatWrite(::basictl::tl_ostream & s, const ::tl2::BoolStat& item) {
	if (!s.int_write(item.statTrue)) { return false;}
	if (!s.int_write(item.statFalse)) { return false;}
	if (!s.int_write(item.statUnknown)) { return false;}
	return true;
}

bool tl2::details::BoolStatReadBoxed(::basictl::tl_istream & s, ::tl2::BoolStat& item) {
	if (!s.nat_read_exact_tag(0x92cbcbfa)) { return false; }
	return tl2::details::BoolStatRead(s, item);
}

bool tl2::details::BoolStatWriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoolStat& item) {
	if (!s.nat_write(0x92cbcbfa)) { return false; }
	return tl2::details::BoolStatWrite(s, item);
}

void tl2::details::BuiltinVectorService3GroupCountLimitReset(std::vector<::tl2::service3::GroupCountLimit>& item) {
	item.resize(0); // TODO - unwrap
}

bool tl2::details::BuiltinVectorService3GroupCountLimitWriteJSON(std::ostream & s, const std::vector<::tl2::service3::GroupCountLimit>& item) {
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		if (!::tl2::details::Service3GroupCountLimitWriteJSON(s, el)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinVectorService3GroupCountLimitRead(::basictl::tl_istream & s, std::vector<::tl2::service3::GroupCountLimit>& item) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	// TODO - check length sanity
	item.resize(len);
	for(auto && el : item) {
		if (!::tl2::details::Service3GroupCountLimitRead(s, el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinVectorService3GroupCountLimitWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service3::GroupCountLimit>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!::tl2::details::Service3GroupCountLimitWrite(s, el)) { return false; }
	}
	return true;
}

void tl2::details::BuiltinVectorService3GroupSizeLimitReset(std::vector<::tl2::service3::GroupSizeLimit>& item) {
	item.resize(0); // TODO - unwrap
}

bool tl2::details::BuiltinVectorService3GroupSizeLimitWriteJSON(std::ostream & s, const std::vector<::tl2::service3::GroupSizeLimit>& item) {
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		if (!::tl2::details::Service3GroupSizeLimitWriteJSON(s, el)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinVectorService3GroupSizeLimitRead(::basictl::tl_istream & s, std::vector<::tl2::service3::GroupSizeLimit>& item) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	// TODO - check length sanity
	item.resize(len);
	for(auto && el : item) {
		if (!::tl2::details::Service3GroupSizeLimitRead(s, el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinVectorService3GroupSizeLimitWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service3::GroupSizeLimit>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!::tl2::details::Service3GroupSizeLimitWrite(s, el)) { return false; }
	}
	return true;
}

void tl2::details::BuiltinVectorService3ProductReset(std::vector<::tl2::service3::Product>& item) {
	item.resize(0); // TODO - unwrap
}

bool tl2::details::BuiltinVectorService3ProductWriteJSON(std::ostream & s, const std::vector<::tl2::service3::Product>& item, uint32_t nat_t) {
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		if (!::tl2::details::Service3ProductWriteJSON(s, el, nat_t)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinVectorService3ProductRead(::basictl::tl_istream & s, std::vector<::tl2::service3::Product>& item, uint32_t nat_t) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	// TODO - check length sanity
	item.resize(len);
	for(auto && el : item) {
		if (!::tl2::details::Service3ProductRead(s, el, nat_t)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinVectorService3ProductWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service3::Product>& item, uint32_t nat_t) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!::tl2::details::Service3ProductWrite(s, el, nat_t)) { return false; }
	}
	return true;
}

void tl2::details::BuiltinVectorService3Product0Reset(std::vector<::tl2::service3::Productmode<0>>& item) {
	item.resize(0); // TODO - unwrap
}

bool tl2::details::BuiltinVectorService3Product0WriteJSON(std::ostream & s, const std::vector<::tl2::service3::Productmode<0>>& item) {
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		if (!::tl2::details::Service3Product0WriteJSON(s, el)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinVectorService3Product0Read(::basictl::tl_istream & s, std::vector<::tl2::service3::Productmode<0>>& item) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	// TODO - check length sanity
	item.resize(len);
	for(auto && el : item) {
		if (!::tl2::details::Service3Product0Read(s, el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinVectorService3Product0Write(::basictl::tl_ostream & s, const std::vector<::tl2::service3::Productmode<0>>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!::tl2::details::Service3Product0Write(s, el)) { return false; }
	}
	return true;
}

void tl2::details::BuiltinVectorService3ProductStatsOldReset(std::vector<::tl2::service3::ProductStatsOld>& item) {
	item.resize(0); // TODO - unwrap
}

bool tl2::details::BuiltinVectorService3ProductStatsOldWriteJSON(std::ostream & s, const std::vector<::tl2::service3::ProductStatsOld>& item) {
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		if (!::tl2::details::Service3ProductStatsOldWriteJSON(s, el)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinVectorService3ProductStatsOldRead(::basictl::tl_istream & s, std::vector<::tl2::service3::ProductStatsOld>& item) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	// TODO - check length sanity
	item.resize(len);
	for(auto && el : item) {
		if (!::tl2::details::Service3ProductStatsOldRead(s, el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinVectorService3ProductStatsOldWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service3::ProductStatsOld>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!::tl2::details::Service3ProductStatsOldWrite(s, el)) { return false; }
	}
	return true;
}

bool tl2::service3::CreateProduct::write_json(std::ostream& s)const {
	if (!::tl2::details::Service3CreateProductWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service3::CreateProduct::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3CreateProductRead(s, *this)) { return false; }
	return true;
}

bool tl2::service3::CreateProduct::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3CreateProductWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service3::CreateProduct::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3CreateProductReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service3::CreateProduct::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3CreateProductWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service3CreateProductReset(::tl2::service3::CreateProduct& item) {
	item.user_id = 0;
	item.type = 0;
	item.id.clear();
	item.info.clear();
	item.date = 0;
	item.expiration_date = 0;
}

bool tl2::details::Service3CreateProductWriteJSON(std::ostream& s, const ::tl2::service3::CreateProduct& item) {
	auto add_comma = false;
	s << "{";
	if (item.user_id != 0) {
		add_comma = true;
		s << "\"user_id\":";
		s << item.user_id;
	}
	if (item.type != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"type\":";
		s << item.type;
	}
	if (item.id.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"id\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.id)) { return false; }
	}
	if (item.info.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"info\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.info)) { return false; }
	}
	if (item.date != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"date\":";
		s << item.date;
	}
	if (item.expiration_date != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"expiration_date\":";
		s << item.expiration_date;
	}
	s << "}";
	return true;
}

bool tl2::details::Service3CreateProductRead(::basictl::tl_istream & s, ::tl2::service3::CreateProduct& item) {
	if (!s.int_read(item.user_id)) { return false; }
	if (!s.int_read(item.type)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, item.id)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, item.info)) { return false; }
	if (!s.int_read(item.date)) { return false; }
	if (!s.int_read(item.expiration_date)) { return false; }
	return true;
}

bool tl2::details::Service3CreateProductWrite(::basictl::tl_ostream & s, const ::tl2::service3::CreateProduct& item) {
	if (!s.int_write(item.user_id)) { return false;}
	if (!s.int_write(item.type)) { return false;}
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.id)) { return false; }
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.info)) { return false; }
	if (!s.int_write(item.date)) { return false;}
	if (!s.int_write(item.expiration_date)) { return false;}
	return true;
}

bool tl2::details::Service3CreateProductReadBoxed(::basictl::tl_istream & s, ::tl2::service3::CreateProduct& item) {
	if (!s.nat_read_exact_tag(0xb7d92bd9)) { return false; }
	return tl2::details::Service3CreateProductRead(s, item);
}

bool tl2::details::Service3CreateProductWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::CreateProduct& item) {
	if (!s.nat_write(0xb7d92bd9)) { return false; }
	return tl2::details::Service3CreateProductWrite(s, item);
}

bool tl2::details::Service3CreateProductReadResult(::basictl::tl_istream & s, tl2::service3::CreateProduct& item, bool& result) {
	if (!::tl2::details::BoolReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service3CreateProductWriteResult(::basictl::tl_ostream & s, tl2::service3::CreateProduct& item, bool& result) {
	if (!::tl2::details::BoolWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service3::CreateProduct::read_result(::basictl::tl_istream & s, bool & result) {
	return tl2::details::Service3CreateProductReadResult(s, *this, result);
}
bool tl2::service3::CreateProduct::write_result(::basictl::tl_ostream & s, bool & result) {
	return tl2::details::Service3CreateProductWriteResult(s, *this, result);
}

bool tl2::service3::DeleteAllProducts::write_json(std::ostream& s)const {
	if (!::tl2::details::Service3DeleteAllProductsWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service3::DeleteAllProducts::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3DeleteAllProductsRead(s, *this)) { return false; }
	return true;
}

bool tl2::service3::DeleteAllProducts::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3DeleteAllProductsWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service3::DeleteAllProducts::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3DeleteAllProductsReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service3::DeleteAllProducts::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3DeleteAllProductsWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service3DeleteAllProductsReset(::tl2::service3::DeleteAllProducts& item) {
	item.user_id = 0;
	item.type = 0;
	item.start_date = 0;
	item.end_date = 0;
}

bool tl2::details::Service3DeleteAllProductsWriteJSON(std::ostream& s, const ::tl2::service3::DeleteAllProducts& item) {
	auto add_comma = false;
	s << "{";
	if (item.user_id != 0) {
		add_comma = true;
		s << "\"user_id\":";
		s << item.user_id;
	}
	if (item.type != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"type\":";
		s << item.type;
	}
	if (item.start_date != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"start_date\":";
		s << item.start_date;
	}
	if (item.end_date != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"end_date\":";
		s << item.end_date;
	}
	s << "}";
	return true;
}

bool tl2::details::Service3DeleteAllProductsRead(::basictl::tl_istream & s, ::tl2::service3::DeleteAllProducts& item) {
	if (!s.int_read(item.user_id)) { return false; }
	if (!s.int_read(item.type)) { return false; }
	if (!s.int_read(item.start_date)) { return false; }
	if (!s.int_read(item.end_date)) { return false; }
	return true;
}

bool tl2::details::Service3DeleteAllProductsWrite(::basictl::tl_ostream & s, const ::tl2::service3::DeleteAllProducts& item) {
	if (!s.int_write(item.user_id)) { return false;}
	if (!s.int_write(item.type)) { return false;}
	if (!s.int_write(item.start_date)) { return false;}
	if (!s.int_write(item.end_date)) { return false;}
	return true;
}

bool tl2::details::Service3DeleteAllProductsReadBoxed(::basictl::tl_istream & s, ::tl2::service3::DeleteAllProducts& item) {
	if (!s.nat_read_exact_tag(0x4494acc2)) { return false; }
	return tl2::details::Service3DeleteAllProductsRead(s, item);
}

bool tl2::details::Service3DeleteAllProductsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::DeleteAllProducts& item) {
	if (!s.nat_write(0x4494acc2)) { return false; }
	return tl2::details::Service3DeleteAllProductsWrite(s, item);
}

bool tl2::details::Service3DeleteAllProductsReadResult(::basictl::tl_istream & s, tl2::service3::DeleteAllProducts& item, bool& result) {
	if (!::tl2::details::BoolReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service3DeleteAllProductsWriteResult(::basictl::tl_ostream & s, tl2::service3::DeleteAllProducts& item, bool& result) {
	if (!::tl2::details::BoolWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service3::DeleteAllProducts::read_result(::basictl::tl_istream & s, bool & result) {
	return tl2::details::Service3DeleteAllProductsReadResult(s, *this, result);
}
bool tl2::service3::DeleteAllProducts::write_result(::basictl::tl_ostream & s, bool & result) {
	return tl2::details::Service3DeleteAllProductsWriteResult(s, *this, result);
}

bool tl2::service3::DeleteGroupedProducts::write_json(std::ostream& s)const {
	if (!::tl2::details::Service3DeleteGroupedProductsWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service3::DeleteGroupedProducts::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3DeleteGroupedProductsRead(s, *this)) { return false; }
	return true;
}

bool tl2::service3::DeleteGroupedProducts::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3DeleteGroupedProductsWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service3::DeleteGroupedProducts::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3DeleteGroupedProductsReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service3::DeleteGroupedProducts::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3DeleteGroupedProductsWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service3DeleteGroupedProductsReset(::tl2::service3::DeleteGroupedProducts& item) {
	item.user_id = 0;
	item.type = 0;
	item.id.clear();
	item.start_date = 0;
	item.end_date = 0;
}

bool tl2::details::Service3DeleteGroupedProductsWriteJSON(std::ostream& s, const ::tl2::service3::DeleteGroupedProducts& item) {
	auto add_comma = false;
	s << "{";
	if (item.user_id != 0) {
		add_comma = true;
		s << "\"user_id\":";
		s << item.user_id;
	}
	if (item.type != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"type\":";
		s << item.type;
	}
	if (item.id.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"id\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.id)) { return false; }
	}
	if (item.start_date != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"start_date\":";
		s << item.start_date;
	}
	if (item.end_date != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"end_date\":";
		s << item.end_date;
	}
	s << "}";
	return true;
}

bool tl2::details::Service3DeleteGroupedProductsRead(::basictl::tl_istream & s, ::tl2::service3::DeleteGroupedProducts& item) {
	if (!s.int_read(item.user_id)) { return false; }
	if (!s.int_read(item.type)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, item.id)) { return false; }
	if (!s.int_read(item.start_date)) { return false; }
	if (!s.int_read(item.end_date)) { return false; }
	return true;
}

bool tl2::details::Service3DeleteGroupedProductsWrite(::basictl::tl_ostream & s, const ::tl2::service3::DeleteGroupedProducts& item) {
	if (!s.int_write(item.user_id)) { return false;}
	if (!s.int_write(item.type)) { return false;}
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.id)) { return false; }
	if (!s.int_write(item.start_date)) { return false;}
	if (!s.int_write(item.end_date)) { return false;}
	return true;
}

bool tl2::details::Service3DeleteGroupedProductsReadBoxed(::basictl::tl_istream & s, ::tl2::service3::DeleteGroupedProducts& item) {
	if (!s.nat_read_exact_tag(0xe468e614)) { return false; }
	return tl2::details::Service3DeleteGroupedProductsRead(s, item);
}

bool tl2::details::Service3DeleteGroupedProductsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::DeleteGroupedProducts& item) {
	if (!s.nat_write(0xe468e614)) { return false; }
	return tl2::details::Service3DeleteGroupedProductsWrite(s, item);
}

bool tl2::details::Service3DeleteGroupedProductsReadResult(::basictl::tl_istream & s, tl2::service3::DeleteGroupedProducts& item, bool& result) {
	if (!::tl2::details::BoolReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service3DeleteGroupedProductsWriteResult(::basictl::tl_ostream & s, tl2::service3::DeleteGroupedProducts& item, bool& result) {
	if (!::tl2::details::BoolWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service3::DeleteGroupedProducts::read_result(::basictl::tl_istream & s, bool & result) {
	return tl2::details::Service3DeleteGroupedProductsReadResult(s, *this, result);
}
bool tl2::service3::DeleteGroupedProducts::write_result(::basictl::tl_ostream & s, bool & result) {
	return tl2::details::Service3DeleteGroupedProductsWriteResult(s, *this, result);
}

bool tl2::service3::DeleteProduct::write_json(std::ostream& s)const {
	if (!::tl2::details::Service3DeleteProductWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service3::DeleteProduct::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3DeleteProductRead(s, *this)) { return false; }
	return true;
}

bool tl2::service3::DeleteProduct::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3DeleteProductWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service3::DeleteProduct::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3DeleteProductReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service3::DeleteProduct::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3DeleteProductWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service3DeleteProductReset(::tl2::service3::DeleteProduct& item) {
	item.user_id = 0;
	item.type = 0;
	item.id.clear();
	item.info.clear();
}

bool tl2::details::Service3DeleteProductWriteJSON(std::ostream& s, const ::tl2::service3::DeleteProduct& item) {
	auto add_comma = false;
	s << "{";
	if (item.user_id != 0) {
		add_comma = true;
		s << "\"user_id\":";
		s << item.user_id;
	}
	if (item.type != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"type\":";
		s << item.type;
	}
	if (item.id.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"id\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.id)) { return false; }
	}
	if (item.info.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"info\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.info)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::Service3DeleteProductRead(::basictl::tl_istream & s, ::tl2::service3::DeleteProduct& item) {
	if (!s.int_read(item.user_id)) { return false; }
	if (!s.int_read(item.type)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, item.id)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, item.info)) { return false; }
	return true;
}

bool tl2::details::Service3DeleteProductWrite(::basictl::tl_ostream & s, const ::tl2::service3::DeleteProduct& item) {
	if (!s.int_write(item.user_id)) { return false;}
	if (!s.int_write(item.type)) { return false;}
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.id)) { return false; }
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.info)) { return false; }
	return true;
}

bool tl2::details::Service3DeleteProductReadBoxed(::basictl::tl_istream & s, ::tl2::service3::DeleteProduct& item) {
	if (!s.nat_read_exact_tag(0x6867e707)) { return false; }
	return tl2::details::Service3DeleteProductRead(s, item);
}

bool tl2::details::Service3DeleteProductWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::DeleteProduct& item) {
	if (!s.nat_write(0x6867e707)) { return false; }
	return tl2::details::Service3DeleteProductWrite(s, item);
}

bool tl2::details::Service3DeleteProductReadResult(::basictl::tl_istream & s, tl2::service3::DeleteProduct& item, bool& result) {
	if (!::tl2::details::BoolReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service3DeleteProductWriteResult(::basictl::tl_ostream & s, tl2::service3::DeleteProduct& item, bool& result) {
	if (!::tl2::details::BoolWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service3::DeleteProduct::read_result(::basictl::tl_istream & s, bool & result) {
	return tl2::details::Service3DeleteProductReadResult(s, *this, result);
}
bool tl2::service3::DeleteProduct::write_result(::basictl::tl_ostream & s, bool & result) {
	return tl2::details::Service3DeleteProductWriteResult(s, *this, result);
}

bool tl2::service3::GetLastVisitTimestamp::write_json(std::ostream& s)const {
	if (!::tl2::details::Service3GetLastVisitTimestampWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GetLastVisitTimestamp::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3GetLastVisitTimestampRead(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GetLastVisitTimestamp::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3GetLastVisitTimestampWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GetLastVisitTimestamp::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3GetLastVisitTimestampReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GetLastVisitTimestamp::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3GetLastVisitTimestampWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service3GetLastVisitTimestampReset(::tl2::service3::GetLastVisitTimestamp& item) {
	item.user_id = 0;
}

bool tl2::details::Service3GetLastVisitTimestampWriteJSON(std::ostream& s, const ::tl2::service3::GetLastVisitTimestamp& item) {
	s << "{";
	if (item.user_id != 0) {
		s << "\"user_id\":";
		s << item.user_id;
	}
	s << "}";
	return true;
}

bool tl2::details::Service3GetLastVisitTimestampRead(::basictl::tl_istream & s, ::tl2::service3::GetLastVisitTimestamp& item) {
	if (!s.int_read(item.user_id)) { return false; }
	return true;
}

bool tl2::details::Service3GetLastVisitTimestampWrite(::basictl::tl_ostream & s, const ::tl2::service3::GetLastVisitTimestamp& item) {
	if (!s.int_write(item.user_id)) { return false;}
	return true;
}

bool tl2::details::Service3GetLastVisitTimestampReadBoxed(::basictl::tl_istream & s, ::tl2::service3::GetLastVisitTimestamp& item) {
	if (!s.nat_read_exact_tag(0x9a4c788d)) { return false; }
	return tl2::details::Service3GetLastVisitTimestampRead(s, item);
}

bool tl2::details::Service3GetLastVisitTimestampWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::GetLastVisitTimestamp& item) {
	if (!s.nat_write(0x9a4c788d)) { return false; }
	return tl2::details::Service3GetLastVisitTimestampWrite(s, item);
}

bool tl2::details::Service3GetLastVisitTimestampReadResult(::basictl::tl_istream & s, tl2::service3::GetLastVisitTimestamp& item, std::optional<int32_t>& result) {
	if (!::tl2::details::IntMaybeReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service3GetLastVisitTimestampWriteResult(::basictl::tl_ostream & s, tl2::service3::GetLastVisitTimestamp& item, std::optional<int32_t>& result) {
	if (!::tl2::details::IntMaybeWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service3::GetLastVisitTimestamp::read_result(::basictl::tl_istream & s, std::optional<int32_t> & result) {
	return tl2::details::Service3GetLastVisitTimestampReadResult(s, *this, result);
}
bool tl2::service3::GetLastVisitTimestamp::write_result(::basictl::tl_ostream & s, std::optional<int32_t> & result) {
	return tl2::details::Service3GetLastVisitTimestampWriteResult(s, *this, result);
}

bool tl2::service3::GetLimits::write_json(std::ostream& s)const {
	if (!::tl2::details::Service3GetLimitsWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GetLimits::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3GetLimitsRead(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GetLimits::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3GetLimitsWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GetLimits::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3GetLimitsReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GetLimits::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3GetLimitsWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service3GetLimitsReset(::tl2::service3::GetLimits& item) {
}

bool tl2::details::Service3GetLimitsWriteJSON(std::ostream& s, const ::tl2::service3::GetLimits& item) {
	s << "true";
	return true;
}

bool tl2::details::Service3GetLimitsRead(::basictl::tl_istream & s, ::tl2::service3::GetLimits& item) {
	return true;
}

bool tl2::details::Service3GetLimitsWrite(::basictl::tl_ostream & s, const ::tl2::service3::GetLimits& item) {
	return true;
}

bool tl2::details::Service3GetLimitsReadBoxed(::basictl::tl_istream & s, ::tl2::service3::GetLimits& item) {
	if (!s.nat_read_exact_tag(0xeb399467)) { return false; }
	return tl2::details::Service3GetLimitsRead(s, item);
}

bool tl2::details::Service3GetLimitsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::GetLimits& item) {
	if (!s.nat_write(0xeb399467)) { return false; }
	return tl2::details::Service3GetLimitsWrite(s, item);
}

bool tl2::details::Service3GetLimitsReadResult(::basictl::tl_istream & s, tl2::service3::GetLimits& item, ::tl2::service3::Limits& result) {
	if (!::tl2::details::Service3LimitsReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service3GetLimitsWriteResult(::basictl::tl_ostream & s, tl2::service3::GetLimits& item, ::tl2::service3::Limits& result) {
	if (!::tl2::details::Service3LimitsWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service3::GetLimits::read_result(::basictl::tl_istream & s, ::tl2::service3::Limits & result) {
	return tl2::details::Service3GetLimitsReadResult(s, *this, result);
}
bool tl2::service3::GetLimits::write_result(::basictl::tl_ostream & s, ::tl2::service3::Limits & result) {
	return tl2::details::Service3GetLimitsWriteResult(s, *this, result);
}

bool tl2::service3::GetProductStats::write_json(std::ostream& s)const {
	if (!::tl2::details::Service3GetProductStatsWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GetProductStats::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3GetProductStatsRead(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GetProductStats::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3GetProductStatsWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GetProductStats::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3GetProductStatsReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GetProductStats::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3GetProductStatsWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service3GetProductStatsReset(::tl2::service3::GetProductStats& item) {
	item.user_id = 0;
	item.types.clear();
}

bool tl2::details::Service3GetProductStatsWriteJSON(std::ostream& s, const ::tl2::service3::GetProductStats& item) {
	auto add_comma = false;
	s << "{";
	if (item.user_id != 0) {
		add_comma = true;
		s << "\"user_id\":";
		s << item.user_id;
	}
	if (item.types.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"types\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.types)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::Service3GetProductStatsRead(::basictl::tl_istream & s, ::tl2::service3::GetProductStats& item) {
	if (!s.int_read(item.user_id)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, item.types)) { return false; }
	return true;
}

bool tl2::details::Service3GetProductStatsWrite(::basictl::tl_ostream & s, const ::tl2::service3::GetProductStats& item) {
	if (!s.int_write(item.user_id)) { return false;}
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.types)) { return false; }
	return true;
}

bool tl2::details::Service3GetProductStatsReadBoxed(::basictl::tl_istream & s, ::tl2::service3::GetProductStats& item) {
	if (!s.nat_read_exact_tag(0x261f6898)) { return false; }
	return tl2::details::Service3GetProductStatsRead(s, item);
}

bool tl2::details::Service3GetProductStatsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::GetProductStats& item) {
	if (!s.nat_write(0x261f6898)) { return false; }
	return tl2::details::Service3GetProductStatsWrite(s, item);
}

bool tl2::details::Service3GetProductStatsReadResult(::basictl::tl_istream & s, tl2::service3::GetProductStats& item, std::optional<std::vector<::tl2::service3::ProductStatsOld>>& result) {
	if (!::tl2::details::VectorService3ProductStatsOldMaybeReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service3GetProductStatsWriteResult(::basictl::tl_ostream & s, tl2::service3::GetProductStats& item, std::optional<std::vector<::tl2::service3::ProductStatsOld>>& result) {
	if (!::tl2::details::VectorService3ProductStatsOldMaybeWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service3::GetProductStats::read_result(::basictl::tl_istream & s, std::optional<std::vector<::tl2::service3::ProductStatsOld>> & result) {
	return tl2::details::Service3GetProductStatsReadResult(s, *this, result);
}
bool tl2::service3::GetProductStats::write_result(::basictl::tl_ostream & s, std::optional<std::vector<::tl2::service3::ProductStatsOld>> & result) {
	return tl2::details::Service3GetProductStatsWriteResult(s, *this, result);
}

bool tl2::service3::GetProducts::write_json(std::ostream& s)const {
	if (!::tl2::details::Service3GetProductsWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GetProducts::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3GetProductsRead(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GetProducts::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3GetProductsWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GetProducts::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3GetProductsReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GetProducts::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3GetProductsWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service3GetProductsReset(::tl2::service3::GetProducts& item) {
	item.user_id = 0;
	item.mode = 0;
	item.types.clear();
	item.start_date = 0;
	item.end_date = 0;
	item.offset = 0;
	item.limit = 0;
	item.allowed_info0.clear();
}

bool tl2::details::Service3GetProductsWriteJSON(std::ostream& s, const ::tl2::service3::GetProducts& item) {
	auto add_comma = false;
	s << "{";
	if (item.user_id != 0) {
		add_comma = true;
		s << "\"user_id\":";
		s << item.user_id;
	}
	if (item.mode != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"mode\":";
		s << item.mode;
	}
	if (item.types.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"types\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.types)) { return false; }
	}
	if (item.start_date != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"start_date\":";
		s << item.start_date;
	}
	if (item.end_date != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"end_date\":";
		s << item.end_date;
	}
	if (item.offset != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"offset\":";
		s << item.offset;
	}
	if (item.limit != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"limit\":";
		s << item.limit;
	}
	if (item.allowed_info0.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"allowed_info0\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.allowed_info0)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::Service3GetProductsRead(::basictl::tl_istream & s, ::tl2::service3::GetProducts& item) {
	if (!s.int_read(item.user_id)) { return false; }
	if (!s.nat_read(item.mode)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, item.types)) { return false; }
	if (!s.int_read(item.start_date)) { return false; }
	if (!s.int_read(item.end_date)) { return false; }
	if (!s.int_read(item.offset)) { return false; }
	if (!s.int_read(item.limit)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, item.allowed_info0)) { return false; }
	return true;
}

bool tl2::details::Service3GetProductsWrite(::basictl::tl_ostream & s, const ::tl2::service3::GetProducts& item) {
	if (!s.int_write(item.user_id)) { return false;}
	if (!s.nat_write(item.mode)) { return false;}
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.types)) { return false; }
	if (!s.int_write(item.start_date)) { return false;}
	if (!s.int_write(item.end_date)) { return false;}
	if (!s.int_write(item.offset)) { return false;}
	if (!s.int_write(item.limit)) { return false;}
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.allowed_info0)) { return false; }
	return true;
}

bool tl2::details::Service3GetProductsReadBoxed(::basictl::tl_istream & s, ::tl2::service3::GetProducts& item) {
	if (!s.nat_read_exact_tag(0xeb306233)) { return false; }
	return tl2::details::Service3GetProductsRead(s, item);
}

bool tl2::details::Service3GetProductsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::GetProducts& item) {
	if (!s.nat_write(0xeb306233)) { return false; }
	return tl2::details::Service3GetProductsWrite(s, item);
}

bool tl2::details::Service3GetProductsReadResult(::basictl::tl_istream & s, tl2::service3::GetProducts& item, std::optional<std::vector<::tl2::service3::Product>>& result) {
	if (!::tl2::details::VectorService3ProductMaybeReadBoxed(s, result, item.mode)) { return false; }
	return true;
}
bool tl2::details::Service3GetProductsWriteResult(::basictl::tl_ostream & s, tl2::service3::GetProducts& item, std::optional<std::vector<::tl2::service3::Product>>& result) {
	if (!::tl2::details::VectorService3ProductMaybeWriteBoxed(s, result, item.mode)) { return false; }
	return true;
}

bool tl2::service3::GetProducts::read_result(::basictl::tl_istream & s, std::optional<std::vector<::tl2::service3::Product>> & result) {
	return tl2::details::Service3GetProductsReadResult(s, *this, result);
}
bool tl2::service3::GetProducts::write_result(::basictl::tl_ostream & s, std::optional<std::vector<::tl2::service3::Product>> & result) {
	return tl2::details::Service3GetProductsWriteResult(s, *this, result);
}

bool tl2::service3::GetScheduledProducts::write_json(std::ostream& s)const {
	if (!::tl2::details::Service3GetScheduledProductsWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GetScheduledProducts::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3GetScheduledProductsRead(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GetScheduledProducts::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3GetScheduledProductsWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GetScheduledProducts::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3GetScheduledProductsReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GetScheduledProducts::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3GetScheduledProductsWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service3GetScheduledProductsReset(::tl2::service3::GetScheduledProducts& item) {
	item.user_id = 0;
	item.types.clear();
}

bool tl2::details::Service3GetScheduledProductsWriteJSON(std::ostream& s, const ::tl2::service3::GetScheduledProducts& item) {
	auto add_comma = false;
	s << "{";
	if (item.user_id != 0) {
		add_comma = true;
		s << "\"user_id\":";
		s << item.user_id;
	}
	if (item.types.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"types\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.types)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::Service3GetScheduledProductsRead(::basictl::tl_istream & s, ::tl2::service3::GetScheduledProducts& item) {
	if (!s.int_read(item.user_id)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, item.types)) { return false; }
	return true;
}

bool tl2::details::Service3GetScheduledProductsWrite(::basictl::tl_ostream & s, const ::tl2::service3::GetScheduledProducts& item) {
	if (!s.int_write(item.user_id)) { return false;}
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.types)) { return false; }
	return true;
}

bool tl2::details::Service3GetScheduledProductsReadBoxed(::basictl::tl_istream & s, ::tl2::service3::GetScheduledProducts& item) {
	if (!s.nat_read_exact_tag(0xf53ad7bd)) { return false; }
	return tl2::details::Service3GetScheduledProductsRead(s, item);
}

bool tl2::details::Service3GetScheduledProductsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::GetScheduledProducts& item) {
	if (!s.nat_write(0xf53ad7bd)) { return false; }
	return tl2::details::Service3GetScheduledProductsWrite(s, item);
}

bool tl2::details::Service3GetScheduledProductsReadResult(::basictl::tl_istream & s, tl2::service3::GetScheduledProducts& item, std::optional<std::vector<::tl2::service3::Productmode<0>>>& result) {
	if (!::tl2::details::VectorService3Product0MaybeReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service3GetScheduledProductsWriteResult(::basictl::tl_ostream & s, tl2::service3::GetScheduledProducts& item, std::optional<std::vector<::tl2::service3::Productmode<0>>>& result) {
	if (!::tl2::details::VectorService3Product0MaybeWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service3::GetScheduledProducts::read_result(::basictl::tl_istream & s, std::optional<std::vector<::tl2::service3::Productmode<0>>> & result) {
	return tl2::details::Service3GetScheduledProductsReadResult(s, *this, result);
}
bool tl2::service3::GetScheduledProducts::write_result(::basictl::tl_ostream & s, std::optional<std::vector<::tl2::service3::Productmode<0>>> & result) {
	return tl2::details::Service3GetScheduledProductsWriteResult(s, *this, result);
}

bool tl2::service3::GroupCountLimit::write_json(std::ostream& s)const {
	if (!::tl2::details::Service3GroupCountLimitWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GroupCountLimit::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3GroupCountLimitRead(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GroupCountLimit::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3GroupCountLimitWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GroupCountLimit::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3GroupCountLimitReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GroupCountLimit::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3GroupCountLimitWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service3GroupCountLimitReset(::tl2::service3::GroupCountLimit& item) {
	item.types.clear();
	item.limit = 0;
}

bool tl2::details::Service3GroupCountLimitWriteJSON(std::ostream& s, const ::tl2::service3::GroupCountLimit& item) {
	auto add_comma = false;
	s << "{";
	if (item.types.size() != 0) {
		add_comma = true;
		s << "\"types\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.types)) { return false; }
	}
	if (item.limit != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"limit\":";
		s << item.limit;
	}
	s << "}";
	return true;
}

bool tl2::details::Service3GroupCountLimitRead(::basictl::tl_istream & s, ::tl2::service3::GroupCountLimit& item) {
	if (!::tl2::details::BuiltinVectorIntRead(s, item.types)) { return false; }
	if (!s.int_read(item.limit)) { return false; }
	return true;
}

bool tl2::details::Service3GroupCountLimitWrite(::basictl::tl_ostream & s, const ::tl2::service3::GroupCountLimit& item) {
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.types)) { return false; }
	if (!s.int_write(item.limit)) { return false;}
	return true;
}

bool tl2::details::Service3GroupCountLimitReadBoxed(::basictl::tl_istream & s, ::tl2::service3::GroupCountLimit& item) {
	if (!s.nat_read_exact_tag(0x8c04ea7f)) { return false; }
	return tl2::details::Service3GroupCountLimitRead(s, item);
}

bool tl2::details::Service3GroupCountLimitWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::GroupCountLimit& item) {
	if (!s.nat_write(0x8c04ea7f)) { return false; }
	return tl2::details::Service3GroupCountLimitWrite(s, item);
}

bool tl2::service3::GroupSizeLimit::write_json(std::ostream& s)const {
	if (!::tl2::details::Service3GroupSizeLimitWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GroupSizeLimit::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3GroupSizeLimitRead(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GroupSizeLimit::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3GroupSizeLimitWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GroupSizeLimit::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3GroupSizeLimitReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service3::GroupSizeLimit::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3GroupSizeLimitWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service3GroupSizeLimitReset(::tl2::service3::GroupSizeLimit& item) {
	item.type = 0;
	item.limit = 0;
}

bool tl2::details::Service3GroupSizeLimitWriteJSON(std::ostream& s, const ::tl2::service3::GroupSizeLimit& item) {
	auto add_comma = false;
	s << "{";
	if (item.type != 0) {
		add_comma = true;
		s << "\"type\":";
		s << item.type;
	}
	if (item.limit != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"limit\":";
		s << item.limit;
	}
	s << "}";
	return true;
}

bool tl2::details::Service3GroupSizeLimitRead(::basictl::tl_istream & s, ::tl2::service3::GroupSizeLimit& item) {
	if (!s.int_read(item.type)) { return false; }
	if (!s.int_read(item.limit)) { return false; }
	return true;
}

bool tl2::details::Service3GroupSizeLimitWrite(::basictl::tl_ostream & s, const ::tl2::service3::GroupSizeLimit& item) {
	if (!s.int_write(item.type)) { return false;}
	if (!s.int_write(item.limit)) { return false;}
	return true;
}

bool tl2::details::Service3GroupSizeLimitReadBoxed(::basictl::tl_istream & s, ::tl2::service3::GroupSizeLimit& item) {
	if (!s.nat_read_exact_tag(0x90e59396)) { return false; }
	return tl2::details::Service3GroupSizeLimitRead(s, item);
}

bool tl2::details::Service3GroupSizeLimitWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::GroupSizeLimit& item) {
	if (!s.nat_write(0x90e59396)) { return false; }
	return tl2::details::Service3GroupSizeLimitWrite(s, item);
}

bool tl2::service3::Limits::write_json(std::ostream& s)const {
	if (!::tl2::details::Service3LimitsWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service3::Limits::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3LimitsRead(s, *this)) { return false; }
	return true;
}

bool tl2::service3::Limits::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3LimitsWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service3::Limits::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3LimitsReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service3::Limits::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3LimitsWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service3LimitsReset(::tl2::service3::Limits& item) {
	item.default_group_size_limit = 0;
	item.custom_group_size_limits.clear();
	item.default_group_count_limit = 0;
	item.custom_group_count_limits.clear();
}

bool tl2::details::Service3LimitsWriteJSON(std::ostream& s, const ::tl2::service3::Limits& item) {
	auto add_comma = false;
	s << "{";
	if (item.default_group_size_limit != 0) {
		add_comma = true;
		s << "\"default_group_size_limit\":";
		s << item.default_group_size_limit;
	}
	if (item.custom_group_size_limits.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"custom_group_size_limits\":";
		if (!::tl2::details::BuiltinVectorService3GroupSizeLimitWriteJSON(s, item.custom_group_size_limits)) { return false; }
	}
	if (item.default_group_count_limit != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"default_group_count_limit\":";
		s << item.default_group_count_limit;
	}
	if (item.custom_group_count_limits.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"custom_group_count_limits\":";
		if (!::tl2::details::BuiltinVectorService3GroupCountLimitWriteJSON(s, item.custom_group_count_limits)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::Service3LimitsRead(::basictl::tl_istream & s, ::tl2::service3::Limits& item) {
	if (!s.int_read(item.default_group_size_limit)) { return false; }
	if (!::tl2::details::BuiltinVectorService3GroupSizeLimitRead(s, item.custom_group_size_limits)) { return false; }
	if (!s.int_read(item.default_group_count_limit)) { return false; }
	if (!::tl2::details::BuiltinVectorService3GroupCountLimitRead(s, item.custom_group_count_limits)) { return false; }
	return true;
}

bool tl2::details::Service3LimitsWrite(::basictl::tl_ostream & s, const ::tl2::service3::Limits& item) {
	if (!s.int_write(item.default_group_size_limit)) { return false;}
	if (!::tl2::details::BuiltinVectorService3GroupSizeLimitWrite(s, item.custom_group_size_limits)) { return false; }
	if (!s.int_write(item.default_group_count_limit)) { return false;}
	if (!::tl2::details::BuiltinVectorService3GroupCountLimitWrite(s, item.custom_group_count_limits)) { return false; }
	return true;
}

bool tl2::details::Service3LimitsReadBoxed(::basictl::tl_istream & s, ::tl2::service3::Limits& item) {
	if (!s.nat_read_exact_tag(0x80ee61ca)) { return false; }
	return tl2::details::Service3LimitsRead(s, item);
}

bool tl2::details::Service3LimitsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::Limits& item) {
	if (!s.nat_write(0x80ee61ca)) { return false; }
	return tl2::details::Service3LimitsWrite(s, item);
}

bool tl2::service3::Product::write_json(std::ostream& s, uint32_t nat_mode)const {
	if (!::tl2::details::Service3ProductWriteJSON(s, *this, nat_mode)) { return false; }
	return true;
}

bool tl2::service3::Product::read(::basictl::tl_istream & s, uint32_t nat_mode) {
	if (!::tl2::details::Service3ProductRead(s, *this, nat_mode)) { return false; }
	return true;
}

bool tl2::service3::Product::write(::basictl::tl_ostream & s, uint32_t nat_mode)const {
	if (!::tl2::details::Service3ProductWrite(s, *this, nat_mode)) { return false; }
	return true;
}

bool tl2::service3::Product::read_boxed(::basictl::tl_istream & s, uint32_t nat_mode) {
	if (!::tl2::details::Service3ProductReadBoxed(s, *this, nat_mode)) { return false; }
	return true;
}

bool tl2::service3::Product::write_boxed(::basictl::tl_ostream & s, uint32_t nat_mode)const {
	if (!::tl2::details::Service3ProductWriteBoxed(s, *this, nat_mode)) { return false; }
	return true;
}

void tl2::details::Service3ProductReset(::tl2::service3::Product& item) {
	item.type = 0;
	item.id.clear();
	item.info.clear();
	item.date = 0;
	item.expiration_date = 0;
	item.removed = false;
}

bool tl2::details::Service3ProductWriteJSON(std::ostream& s, const ::tl2::service3::Product& item, uint32_t nat_mode) {
	auto add_comma = false;
	s << "{";
	if (item.type != 0) {
		add_comma = true;
		s << "\"type\":";
		s << item.type;
	}
	if (item.id.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"id\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.id)) { return false; }
	}
	if (item.info.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"info\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.info)) { return false; }
	}
	if (item.date != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"date\":";
		s << item.date;
	}
	if (item.expiration_date != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"expiration_date\":";
		s << item.expiration_date;
	}
	if ((nat_mode & (1<<0)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"removed\":";
		if (!::tl2::details::BoolWriteJSON(s, item.removed)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::Service3ProductRead(::basictl::tl_istream & s, ::tl2::service3::Product& item, uint32_t nat_mode) {
	if (!s.int_read(item.type)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, item.id)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, item.info)) { return false; }
	if (!s.int_read(item.date)) { return false; }
	if (!s.int_read(item.expiration_date)) { return false; }
	if ((nat_mode & (1<<0)) != 0) {
		if (!::tl2::details::BoolReadBoxed(s, item.removed)) { return false; }
	} else {
			item.removed = false;
	}
	return true;
}

bool tl2::details::Service3ProductWrite(::basictl::tl_ostream & s, const ::tl2::service3::Product& item, uint32_t nat_mode) {
	if (!s.int_write(item.type)) { return false;}
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.id)) { return false; }
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.info)) { return false; }
	if (!s.int_write(item.date)) { return false;}
	if (!s.int_write(item.expiration_date)) { return false;}
	if ((nat_mode & (1<<0)) != 0) {
			if (!::tl2::details::BoolWriteBoxed(s, item.removed)) { return false; }
	}
	return true;
}

bool tl2::details::Service3ProductReadBoxed(::basictl::tl_istream & s, ::tl2::service3::Product& item, uint32_t nat_mode) {
	if (!s.nat_read_exact_tag(0x461f4ce2)) { return false; }
	return tl2::details::Service3ProductRead(s, item, nat_mode);
}

bool tl2::details::Service3ProductWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::Product& item, uint32_t nat_mode) {
	if (!s.nat_write(0x461f4ce2)) { return false; }
	return tl2::details::Service3ProductWrite(s, item, nat_mode);
}

void tl2::details::Service3Product0Reset(::tl2::service3::Productmode<0>& item) {
	item.type = 0;
	item.id.clear();
	item.info.clear();
	item.date = 0;
	item.expiration_date = 0;
	item.removed = false;
}

bool tl2::details::Service3Product0WriteJSON(std::ostream& s, const ::tl2::service3::Productmode<0>& item) {
	auto add_comma = false;
	s << "{";
	if (item.type != 0) {
		add_comma = true;
		s << "\"type\":";
		s << item.type;
	}
	if (item.id.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"id\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.id)) { return false; }
	}
	if (item.info.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"info\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.info)) { return false; }
	}
	if (item.date != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"date\":";
		s << item.date;
	}
	if (item.expiration_date != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"expiration_date\":";
		s << item.expiration_date;
	}
	if ((0 & (1<<0)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"removed\":";
		if (!::tl2::details::BoolWriteJSON(s, item.removed)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::Service3Product0Read(::basictl::tl_istream & s, ::tl2::service3::Productmode<0>& item) {
	if (!s.int_read(item.type)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, item.id)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, item.info)) { return false; }
	if (!s.int_read(item.date)) { return false; }
	if (!s.int_read(item.expiration_date)) { return false; }
	if ((0 & (1<<0)) != 0) {
		if (!::tl2::details::BoolReadBoxed(s, item.removed)) { return false; }
	} else {
			item.removed = false;
	}
	return true;
}

bool tl2::details::Service3Product0Write(::basictl::tl_ostream & s, const ::tl2::service3::Productmode<0>& item) {
	if (!s.int_write(item.type)) { return false;}
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.id)) { return false; }
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.info)) { return false; }
	if (!s.int_write(item.date)) { return false;}
	if (!s.int_write(item.expiration_date)) { return false;}
	if ((0 & (1<<0)) != 0) {
			if (!::tl2::details::BoolWriteBoxed(s, item.removed)) { return false; }
	}
	return true;
}

bool tl2::details::Service3Product0ReadBoxed(::basictl::tl_istream & s, ::tl2::service3::Productmode<0>& item) {
	if (!s.nat_read_exact_tag(0x461f4ce2)) { return false; }
	return tl2::details::Service3Product0Read(s, item);
}

bool tl2::details::Service3Product0WriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::Productmode<0>& item) {
	if (!s.nat_write(0x461f4ce2)) { return false; }
	return tl2::details::Service3Product0Write(s, item);
}

bool tl2::service3::ProductStatsOld::write_json(std::ostream& s)const {
	if (!::tl2::details::Service3ProductStatsOldWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service3::ProductStatsOld::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3ProductStatsOldRead(s, *this)) { return false; }
	return true;
}

bool tl2::service3::ProductStatsOld::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3ProductStatsOldWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service3::ProductStatsOld::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3ProductStatsOldReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service3::ProductStatsOld::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3ProductStatsOldWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service3ProductStatsOldReset(::tl2::service3::ProductStatsOld& item) {
	item.type = 0;
	item.count_new = 0;
	item.count_total = 0;
	item.count_scheduled = 0;
	item.next_scheduled_at = 0;
}

bool tl2::details::Service3ProductStatsOldWriteJSON(std::ostream& s, const ::tl2::service3::ProductStatsOld& item) {
	auto add_comma = false;
	s << "{";
	if (item.type != 0) {
		add_comma = true;
		s << "\"type\":";
		s << item.type;
	}
	if (item.count_new != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"count_new\":";
		s << item.count_new;
	}
	if (item.count_total != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"count_total\":";
		s << item.count_total;
	}
	if (item.count_scheduled != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"count_scheduled\":";
		s << item.count_scheduled;
	}
	if (item.next_scheduled_at != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"next_scheduled_at\":";
		s << item.next_scheduled_at;
	}
	s << "}";
	return true;
}

bool tl2::details::Service3ProductStatsOldRead(::basictl::tl_istream & s, ::tl2::service3::ProductStatsOld& item) {
	if (!s.int_read(item.type)) { return false; }
	if (!s.int_read(item.count_new)) { return false; }
	if (!s.int_read(item.count_total)) { return false; }
	if (!s.int_read(item.count_scheduled)) { return false; }
	if (!s.int_read(item.next_scheduled_at)) { return false; }
	return true;
}

bool tl2::details::Service3ProductStatsOldWrite(::basictl::tl_ostream & s, const ::tl2::service3::ProductStatsOld& item) {
	if (!s.int_write(item.type)) { return false;}
	if (!s.int_write(item.count_new)) { return false;}
	if (!s.int_write(item.count_total)) { return false;}
	if (!s.int_write(item.count_scheduled)) { return false;}
	if (!s.int_write(item.next_scheduled_at)) { return false;}
	return true;
}

bool tl2::details::Service3ProductStatsOldReadBoxed(::basictl::tl_istream & s, ::tl2::service3::ProductStatsOld& item) {
	if (!s.nat_read_exact_tag(0x6319810b)) { return false; }
	return tl2::details::Service3ProductStatsOldRead(s, item);
}

bool tl2::details::Service3ProductStatsOldWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::ProductStatsOld& item) {
	if (!s.nat_write(0x6319810b)) { return false; }
	return tl2::details::Service3ProductStatsOldWrite(s, item);
}

bool tl2::service3::RestoreAllProducts::write_json(std::ostream& s)const {
	if (!::tl2::details::Service3RestoreAllProductsWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service3::RestoreAllProducts::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3RestoreAllProductsRead(s, *this)) { return false; }
	return true;
}

bool tl2::service3::RestoreAllProducts::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3RestoreAllProductsWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service3::RestoreAllProducts::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3RestoreAllProductsReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service3::RestoreAllProducts::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3RestoreAllProductsWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service3RestoreAllProductsReset(::tl2::service3::RestoreAllProducts& item) {
	item.user_id = 0;
	item.type = 0;
	item.start_date = 0;
	item.end_date = 0;
}

bool tl2::details::Service3RestoreAllProductsWriteJSON(std::ostream& s, const ::tl2::service3::RestoreAllProducts& item) {
	auto add_comma = false;
	s << "{";
	if (item.user_id != 0) {
		add_comma = true;
		s << "\"user_id\":";
		s << item.user_id;
	}
	if (item.type != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"type\":";
		s << item.type;
	}
	if (item.start_date != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"start_date\":";
		s << item.start_date;
	}
	if (item.end_date != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"end_date\":";
		s << item.end_date;
	}
	s << "}";
	return true;
}

bool tl2::details::Service3RestoreAllProductsRead(::basictl::tl_istream & s, ::tl2::service3::RestoreAllProducts& item) {
	if (!s.int_read(item.user_id)) { return false; }
	if (!s.int_read(item.type)) { return false; }
	if (!s.int_read(item.start_date)) { return false; }
	if (!s.int_read(item.end_date)) { return false; }
	return true;
}

bool tl2::details::Service3RestoreAllProductsWrite(::basictl::tl_ostream & s, const ::tl2::service3::RestoreAllProducts& item) {
	if (!s.int_write(item.user_id)) { return false;}
	if (!s.int_write(item.type)) { return false;}
	if (!s.int_write(item.start_date)) { return false;}
	if (!s.int_write(item.end_date)) { return false;}
	return true;
}

bool tl2::details::Service3RestoreAllProductsReadBoxed(::basictl::tl_istream & s, ::tl2::service3::RestoreAllProducts& item) {
	if (!s.nat_read_exact_tag(0x4d839ed0)) { return false; }
	return tl2::details::Service3RestoreAllProductsRead(s, item);
}

bool tl2::details::Service3RestoreAllProductsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::RestoreAllProducts& item) {
	if (!s.nat_write(0x4d839ed0)) { return false; }
	return tl2::details::Service3RestoreAllProductsWrite(s, item);
}

bool tl2::details::Service3RestoreAllProductsReadResult(::basictl::tl_istream & s, tl2::service3::RestoreAllProducts& item, bool& result) {
	if (!::tl2::details::BoolReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service3RestoreAllProductsWriteResult(::basictl::tl_ostream & s, tl2::service3::RestoreAllProducts& item, bool& result) {
	if (!::tl2::details::BoolWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service3::RestoreAllProducts::read_result(::basictl::tl_istream & s, bool & result) {
	return tl2::details::Service3RestoreAllProductsReadResult(s, *this, result);
}
bool tl2::service3::RestoreAllProducts::write_result(::basictl::tl_ostream & s, bool & result) {
	return tl2::details::Service3RestoreAllProductsWriteResult(s, *this, result);
}

bool tl2::service3::RestoreGroupedProducts::write_json(std::ostream& s)const {
	if (!::tl2::details::Service3RestoreGroupedProductsWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service3::RestoreGroupedProducts::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3RestoreGroupedProductsRead(s, *this)) { return false; }
	return true;
}

bool tl2::service3::RestoreGroupedProducts::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3RestoreGroupedProductsWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service3::RestoreGroupedProducts::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3RestoreGroupedProductsReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service3::RestoreGroupedProducts::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3RestoreGroupedProductsWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service3RestoreGroupedProductsReset(::tl2::service3::RestoreGroupedProducts& item) {
	item.user_id = 0;
	item.type = 0;
	item.id.clear();
	item.start_date = 0;
	item.end_date = 0;
}

bool tl2::details::Service3RestoreGroupedProductsWriteJSON(std::ostream& s, const ::tl2::service3::RestoreGroupedProducts& item) {
	auto add_comma = false;
	s << "{";
	if (item.user_id != 0) {
		add_comma = true;
		s << "\"user_id\":";
		s << item.user_id;
	}
	if (item.type != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"type\":";
		s << item.type;
	}
	if (item.id.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"id\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.id)) { return false; }
	}
	if (item.start_date != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"start_date\":";
		s << item.start_date;
	}
	if (item.end_date != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"end_date\":";
		s << item.end_date;
	}
	s << "}";
	return true;
}

bool tl2::details::Service3RestoreGroupedProductsRead(::basictl::tl_istream & s, ::tl2::service3::RestoreGroupedProducts& item) {
	if (!s.int_read(item.user_id)) { return false; }
	if (!s.int_read(item.type)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, item.id)) { return false; }
	if (!s.int_read(item.start_date)) { return false; }
	if (!s.int_read(item.end_date)) { return false; }
	return true;
}

bool tl2::details::Service3RestoreGroupedProductsWrite(::basictl::tl_ostream & s, const ::tl2::service3::RestoreGroupedProducts& item) {
	if (!s.int_write(item.user_id)) { return false;}
	if (!s.int_write(item.type)) { return false;}
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.id)) { return false; }
	if (!s.int_write(item.start_date)) { return false;}
	if (!s.int_write(item.end_date)) { return false;}
	return true;
}

bool tl2::details::Service3RestoreGroupedProductsReadBoxed(::basictl::tl_istream & s, ::tl2::service3::RestoreGroupedProducts& item) {
	if (!s.nat_read_exact_tag(0x1f17bfac)) { return false; }
	return tl2::details::Service3RestoreGroupedProductsRead(s, item);
}

bool tl2::details::Service3RestoreGroupedProductsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::RestoreGroupedProducts& item) {
	if (!s.nat_write(0x1f17bfac)) { return false; }
	return tl2::details::Service3RestoreGroupedProductsWrite(s, item);
}

bool tl2::details::Service3RestoreGroupedProductsReadResult(::basictl::tl_istream & s, tl2::service3::RestoreGroupedProducts& item, bool& result) {
	if (!::tl2::details::BoolReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service3RestoreGroupedProductsWriteResult(::basictl::tl_ostream & s, tl2::service3::RestoreGroupedProducts& item, bool& result) {
	if (!::tl2::details::BoolWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service3::RestoreGroupedProducts::read_result(::basictl::tl_istream & s, bool & result) {
	return tl2::details::Service3RestoreGroupedProductsReadResult(s, *this, result);
}
bool tl2::service3::RestoreGroupedProducts::write_result(::basictl::tl_ostream & s, bool & result) {
	return tl2::details::Service3RestoreGroupedProductsWriteResult(s, *this, result);
}

bool tl2::service3::RestoreProduct::write_json(std::ostream& s)const {
	if (!::tl2::details::Service3RestoreProductWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service3::RestoreProduct::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3RestoreProductRead(s, *this)) { return false; }
	return true;
}

bool tl2::service3::RestoreProduct::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3RestoreProductWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service3::RestoreProduct::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3RestoreProductReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service3::RestoreProduct::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3RestoreProductWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service3RestoreProductReset(::tl2::service3::RestoreProduct& item) {
	item.user_id = 0;
	item.type = 0;
	item.id.clear();
	item.info.clear();
}

bool tl2::details::Service3RestoreProductWriteJSON(std::ostream& s, const ::tl2::service3::RestoreProduct& item) {
	auto add_comma = false;
	s << "{";
	if (item.user_id != 0) {
		add_comma = true;
		s << "\"user_id\":";
		s << item.user_id;
	}
	if (item.type != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"type\":";
		s << item.type;
	}
	if (item.id.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"id\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.id)) { return false; }
	}
	if (item.info.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"info\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.info)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::Service3RestoreProductRead(::basictl::tl_istream & s, ::tl2::service3::RestoreProduct& item) {
	if (!s.int_read(item.user_id)) { return false; }
	if (!s.int_read(item.type)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, item.id)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, item.info)) { return false; }
	return true;
}

bool tl2::details::Service3RestoreProductWrite(::basictl::tl_ostream & s, const ::tl2::service3::RestoreProduct& item) {
	if (!s.int_write(item.user_id)) { return false;}
	if (!s.int_write(item.type)) { return false;}
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.id)) { return false; }
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.info)) { return false; }
	return true;
}

bool tl2::details::Service3RestoreProductReadBoxed(::basictl::tl_istream & s, ::tl2::service3::RestoreProduct& item) {
	if (!s.nat_read_exact_tag(0x6170d515)) { return false; }
	return tl2::details::Service3RestoreProductRead(s, item);
}

bool tl2::details::Service3RestoreProductWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::RestoreProduct& item) {
	if (!s.nat_write(0x6170d515)) { return false; }
	return tl2::details::Service3RestoreProductWrite(s, item);
}

bool tl2::details::Service3RestoreProductReadResult(::basictl::tl_istream & s, tl2::service3::RestoreProduct& item, bool& result) {
	if (!::tl2::details::BoolReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service3RestoreProductWriteResult(::basictl::tl_ostream & s, tl2::service3::RestoreProduct& item, bool& result) {
	if (!::tl2::details::BoolWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service3::RestoreProduct::read_result(::basictl::tl_istream & s, bool & result) {
	return tl2::details::Service3RestoreProductReadResult(s, *this, result);
}
bool tl2::service3::RestoreProduct::write_result(::basictl::tl_ostream & s, bool & result) {
	return tl2::details::Service3RestoreProductWriteResult(s, *this, result);
}

bool tl2::service3::SetLastVisitTimestamp::write_json(std::ostream& s)const {
	if (!::tl2::details::Service3SetLastVisitTimestampWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service3::SetLastVisitTimestamp::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3SetLastVisitTimestampRead(s, *this)) { return false; }
	return true;
}

bool tl2::service3::SetLastVisitTimestamp::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3SetLastVisitTimestampWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service3::SetLastVisitTimestamp::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3SetLastVisitTimestampReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service3::SetLastVisitTimestamp::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3SetLastVisitTimestampWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service3SetLastVisitTimestampReset(::tl2::service3::SetLastVisitTimestamp& item) {
	item.user_id = 0;
	item.timestamp = 0;
}

bool tl2::details::Service3SetLastVisitTimestampWriteJSON(std::ostream& s, const ::tl2::service3::SetLastVisitTimestamp& item) {
	auto add_comma = false;
	s << "{";
	if (item.user_id != 0) {
		add_comma = true;
		s << "\"user_id\":";
		s << item.user_id;
	}
	if (item.timestamp != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"timestamp\":";
		s << item.timestamp;
	}
	s << "}";
	return true;
}

bool tl2::details::Service3SetLastVisitTimestampRead(::basictl::tl_istream & s, ::tl2::service3::SetLastVisitTimestamp& item) {
	if (!s.int_read(item.user_id)) { return false; }
	if (!s.int_read(item.timestamp)) { return false; }
	return true;
}

bool tl2::details::Service3SetLastVisitTimestampWrite(::basictl::tl_ostream & s, const ::tl2::service3::SetLastVisitTimestamp& item) {
	if (!s.int_write(item.user_id)) { return false;}
	if (!s.int_write(item.timestamp)) { return false;}
	return true;
}

bool tl2::details::Service3SetLastVisitTimestampReadBoxed(::basictl::tl_istream & s, ::tl2::service3::SetLastVisitTimestamp& item) {
	if (!s.nat_read_exact_tag(0x7909b020)) { return false; }
	return tl2::details::Service3SetLastVisitTimestampRead(s, item);
}

bool tl2::details::Service3SetLastVisitTimestampWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::SetLastVisitTimestamp& item) {
	if (!s.nat_write(0x7909b020)) { return false; }
	return tl2::details::Service3SetLastVisitTimestampWrite(s, item);
}

bool tl2::details::Service3SetLastVisitTimestampReadResult(::basictl::tl_istream & s, tl2::service3::SetLastVisitTimestamp& item, bool& result) {
	if (!::tl2::details::BoolReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service3SetLastVisitTimestampWriteResult(::basictl::tl_ostream & s, tl2::service3::SetLastVisitTimestamp& item, bool& result) {
	if (!::tl2::details::BoolWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service3::SetLastVisitTimestamp::read_result(::basictl::tl_istream & s, bool & result) {
	return tl2::details::Service3SetLastVisitTimestampReadResult(s, *this, result);
}
bool tl2::service3::SetLastVisitTimestamp::write_result(::basictl::tl_ostream & s, bool & result) {
	return tl2::details::Service3SetLastVisitTimestampWriteResult(s, *this, result);
}

bool tl2::service3::SetLimits::write_json(std::ostream& s)const {
	if (!::tl2::details::Service3SetLimitsWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service3::SetLimits::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3SetLimitsRead(s, *this)) { return false; }
	return true;
}

bool tl2::service3::SetLimits::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3SetLimitsWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service3::SetLimits::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service3SetLimitsReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service3::SetLimits::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service3SetLimitsWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service3SetLimitsReset(::tl2::service3::SetLimits& item) {
	::tl2::details::Service3LimitsReset(item.limits);
}

bool tl2::details::Service3SetLimitsWriteJSON(std::ostream& s, const ::tl2::service3::SetLimits& item) {
	s << "{";
	s << "\"limits\":";
	if (!::tl2::details::Service3LimitsWriteJSON(s, item.limits)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::Service3SetLimitsRead(::basictl::tl_istream & s, ::tl2::service3::SetLimits& item) {
	if (!::tl2::details::Service3LimitsRead(s, item.limits)) { return false; }
	return true;
}

bool tl2::details::Service3SetLimitsWrite(::basictl::tl_ostream & s, const ::tl2::service3::SetLimits& item) {
	if (!::tl2::details::Service3LimitsWrite(s, item.limits)) { return false; }
	return true;
}

bool tl2::details::Service3SetLimitsReadBoxed(::basictl::tl_istream & s, ::tl2::service3::SetLimits& item) {
	if (!s.nat_read_exact_tag(0x3ad5c19c)) { return false; }
	return tl2::details::Service3SetLimitsRead(s, item);
}

bool tl2::details::Service3SetLimitsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service3::SetLimits& item) {
	if (!s.nat_write(0x3ad5c19c)) { return false; }
	return tl2::details::Service3SetLimitsWrite(s, item);
}

bool tl2::details::Service3SetLimitsReadResult(::basictl::tl_istream & s, tl2::service3::SetLimits& item, ::tl2::BoolStat& result) {
	if (!::tl2::details::BoolStatReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service3SetLimitsWriteResult(::basictl::tl_ostream & s, tl2::service3::SetLimits& item, ::tl2::BoolStat& result) {
	if (!::tl2::details::BoolStatWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service3::SetLimits::read_result(::basictl::tl_istream & s, ::tl2::BoolStat & result) {
	return tl2::details::Service3SetLimitsReadResult(s, *this, result);
}
bool tl2::service3::SetLimits::write_result(::basictl::tl_ostream & s, ::tl2::BoolStat & result) {
	return tl2::details::Service3SetLimitsWriteResult(s, *this, result);
}

void tl2::details::VectorService3GroupCountLimitReset(std::vector<::tl2::service3::GroupCountLimit>& item) {
	item.clear();
}

bool tl2::details::VectorService3GroupCountLimitWriteJSON(std::ostream& s, const std::vector<::tl2::service3::GroupCountLimit>& item) {
	if (!::tl2::details::BuiltinVectorService3GroupCountLimitWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService3GroupCountLimitRead(::basictl::tl_istream & s, std::vector<::tl2::service3::GroupCountLimit>& item) {
	if (!::tl2::details::BuiltinVectorService3GroupCountLimitRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService3GroupCountLimitWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service3::GroupCountLimit>& item) {
	if (!::tl2::details::BuiltinVectorService3GroupCountLimitWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService3GroupCountLimitReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::service3::GroupCountLimit>& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	return tl2::details::VectorService3GroupCountLimitRead(s, item);
}

bool tl2::details::VectorService3GroupCountLimitWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::service3::GroupCountLimit>& item) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	return tl2::details::VectorService3GroupCountLimitWrite(s, item);
}

void tl2::details::VectorService3GroupSizeLimitReset(std::vector<::tl2::service3::GroupSizeLimit>& item) {
	item.clear();
}

bool tl2::details::VectorService3GroupSizeLimitWriteJSON(std::ostream& s, const std::vector<::tl2::service3::GroupSizeLimit>& item) {
	if (!::tl2::details::BuiltinVectorService3GroupSizeLimitWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService3GroupSizeLimitRead(::basictl::tl_istream & s, std::vector<::tl2::service3::GroupSizeLimit>& item) {
	if (!::tl2::details::BuiltinVectorService3GroupSizeLimitRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService3GroupSizeLimitWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service3::GroupSizeLimit>& item) {
	if (!::tl2::details::BuiltinVectorService3GroupSizeLimitWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService3GroupSizeLimitReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::service3::GroupSizeLimit>& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	return tl2::details::VectorService3GroupSizeLimitRead(s, item);
}

bool tl2::details::VectorService3GroupSizeLimitWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::service3::GroupSizeLimit>& item) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	return tl2::details::VectorService3GroupSizeLimitWrite(s, item);
}

void tl2::details::VectorService3ProductReset(std::vector<::tl2::service3::Product>& item) {
	item.clear();
}

bool tl2::details::VectorService3ProductWriteJSON(std::ostream& s, const std::vector<::tl2::service3::Product>& item, uint32_t nat_t) {
	if (!::tl2::details::BuiltinVectorService3ProductWriteJSON(s, item, nat_t)) { return false; }
	return true;
}

bool tl2::details::VectorService3ProductRead(::basictl::tl_istream & s, std::vector<::tl2::service3::Product>& item, uint32_t nat_t) {
	if (!::tl2::details::BuiltinVectorService3ProductRead(s, item, nat_t)) { return false; }
	return true;
}

bool tl2::details::VectorService3ProductWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service3::Product>& item, uint32_t nat_t) {
	if (!::tl2::details::BuiltinVectorService3ProductWrite(s, item, nat_t)) { return false; }
	return true;
}

bool tl2::details::VectorService3ProductReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::service3::Product>& item, uint32_t nat_t) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	return tl2::details::VectorService3ProductRead(s, item, nat_t);
}

bool tl2::details::VectorService3ProductWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::service3::Product>& item, uint32_t nat_t) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	return tl2::details::VectorService3ProductWrite(s, item, nat_t);
}

void tl2::details::VectorService3Product0Reset(std::vector<::tl2::service3::Productmode<0>>& item) {
	item.clear();
}

bool tl2::details::VectorService3Product0WriteJSON(std::ostream& s, const std::vector<::tl2::service3::Productmode<0>>& item) {
	if (!::tl2::details::BuiltinVectorService3Product0WriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService3Product0Read(::basictl::tl_istream & s, std::vector<::tl2::service3::Productmode<0>>& item) {
	if (!::tl2::details::BuiltinVectorService3Product0Read(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService3Product0Write(::basictl::tl_ostream & s, const std::vector<::tl2::service3::Productmode<0>>& item) {
	if (!::tl2::details::BuiltinVectorService3Product0Write(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService3Product0ReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::service3::Productmode<0>>& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	return tl2::details::VectorService3Product0Read(s, item);
}

bool tl2::details::VectorService3Product0WriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::service3::Productmode<0>>& item) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	return tl2::details::VectorService3Product0Write(s, item);
}

bool tl2::details::VectorService3Product0MaybeWriteJSON(std::ostream & s, const std::optional<std::vector<::tl2::service3::Productmode<0>>>& item) {
	s << "{";
	if (item) {
		s << "\"ok\":true";
		if((*item).size() != 0) {
			s << ",\"value\":";
			if (!::tl2::details::BuiltinVectorService3Product0WriteJSON(s, *item)) { return false; }
		}
	}
	s << "}";
	return true;
}
bool tl2::details::VectorService3Product0MaybeReadBoxed(::basictl::tl_istream & s, std::optional<std::vector<::tl2::service3::Productmode<0>>>& item) {
	bool has_item = false;
	if (!s.bool_read(has_item, 0x27930a7b, 0x3f9c8ef8)) { return false; }
	if (has_item) {
		if (!item) {
			item.emplace();
		}
		if (!::tl2::details::BuiltinVectorService3Product0Read(s, *item)) { return false; }
		return true;
	}
	item.reset();
	return true;
}

bool tl2::details::VectorService3Product0MaybeWriteBoxed(::basictl::tl_ostream & s, const std::optional<std::vector<::tl2::service3::Productmode<0>>>& item) {
	if (!s.nat_write(item ? 0x3f9c8ef8 : 0x27930a7b)) { return false; }
	if (item) {
		if (!::tl2::details::BuiltinVectorService3Product0Write(s, *item)) { return false; }
	}
	return true;
}

bool tl2::details::VectorService3ProductMaybeWriteJSON(std::ostream & s, const std::optional<std::vector<::tl2::service3::Product>>& item, uint32_t nat_t) {
	s << "{";
	if (item) {
		s << "\"ok\":true";
		if((*item).size() != 0) {
			s << ",\"value\":";
			if (!::tl2::details::BuiltinVectorService3ProductWriteJSON(s, *item, nat_t)) { return false; }
		}
	}
	s << "}";
	return true;
}
bool tl2::details::VectorService3ProductMaybeReadBoxed(::basictl::tl_istream & s, std::optional<std::vector<::tl2::service3::Product>>& item, uint32_t nat_t) {
	bool has_item = false;
	if (!s.bool_read(has_item, 0x27930a7b, 0x3f9c8ef8)) { return false; }
	if (has_item) {
		if (!item) {
			item.emplace();
		}
		if (!::tl2::details::BuiltinVectorService3ProductRead(s, *item, nat_t)) { return false; }
		return true;
	}
	item.reset();
	return true;
}

bool tl2::details::VectorService3ProductMaybeWriteBoxed(::basictl::tl_ostream & s, const std::optional<std::vector<::tl2::service3::Product>>& item, uint32_t nat_t) {
	if (!s.nat_write(item ? 0x3f9c8ef8 : 0x27930a7b)) { return false; }
	if (item) {
		if (!::tl2::details::BuiltinVectorService3ProductWrite(s, *item, nat_t)) { return false; }
	}
	return true;
}

void tl2::details::VectorService3ProductStatsOldReset(std::vector<::tl2::service3::ProductStatsOld>& item) {
	item.clear();
}

bool tl2::details::VectorService3ProductStatsOldWriteJSON(std::ostream& s, const std::vector<::tl2::service3::ProductStatsOld>& item) {
	if (!::tl2::details::BuiltinVectorService3ProductStatsOldWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService3ProductStatsOldRead(::basictl::tl_istream & s, std::vector<::tl2::service3::ProductStatsOld>& item) {
	if (!::tl2::details::BuiltinVectorService3ProductStatsOldRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService3ProductStatsOldWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service3::ProductStatsOld>& item) {
	if (!::tl2::details::BuiltinVectorService3ProductStatsOldWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService3ProductStatsOldReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::service3::ProductStatsOld>& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	return tl2::details::VectorService3ProductStatsOldRead(s, item);
}

bool tl2::details::VectorService3ProductStatsOldWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::service3::ProductStatsOld>& item) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	return tl2::details::VectorService3ProductStatsOldWrite(s, item);
}

bool tl2::details::VectorService3ProductStatsOldMaybeWriteJSON(std::ostream & s, const std::optional<std::vector<::tl2::service3::ProductStatsOld>>& item) {
	s << "{";
	if (item) {
		s << "\"ok\":true";
		if((*item).size() != 0) {
			s << ",\"value\":";
			if (!::tl2::details::BuiltinVectorService3ProductStatsOldWriteJSON(s, *item)) { return false; }
		}
	}
	s << "}";
	return true;
}
bool tl2::details::VectorService3ProductStatsOldMaybeReadBoxed(::basictl::tl_istream & s, std::optional<std::vector<::tl2::service3::ProductStatsOld>>& item) {
	bool has_item = false;
	if (!s.bool_read(has_item, 0x27930a7b, 0x3f9c8ef8)) { return false; }
	if (has_item) {
		if (!item) {
			item.emplace();
		}
		if (!::tl2::details::BuiltinVectorService3ProductStatsOldRead(s, *item)) { return false; }
		return true;
	}
	item.reset();
	return true;
}

bool tl2::details::VectorService3ProductStatsOldMaybeWriteBoxed(::basictl::tl_ostream & s, const std::optional<std::vector<::tl2::service3::ProductStatsOld>>& item) {
	if (!s.nat_write(item ? 0x3f9c8ef8 : 0x27930a7b)) { return false; }
	if (item) {
		if (!::tl2::details::BuiltinVectorService3ProductStatsOldWrite(s, *item)) { return false; }
	}
	return true;
}
