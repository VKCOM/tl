#pragma once

#include "../../basictl/io_streams.h"
#include "../../__common_namespace/types/vector.h"
#include "../types/service3.productStatsOld.h"
#include "../types/service3.product.h"
#include "../types/service3.groupSizeLimit.h"
#include "../types/service3.groupCountLimit.h"

namespace tl2 { namespace details { 

void VectorService3GroupCountLimitReset(std::vector<::tl2::service3::GroupCountLimit>& item);

bool VectorService3GroupCountLimitWriteJSON(std::ostream& s, const std::vector<::tl2::service3::GroupCountLimit>& item);
bool VectorService3GroupCountLimitRead(::basictl::tl_istream & s, std::vector<::tl2::service3::GroupCountLimit>& item);
bool VectorService3GroupCountLimitWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service3::GroupCountLimit>& item);
bool VectorService3GroupCountLimitReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::service3::GroupCountLimit>& item);
bool VectorService3GroupCountLimitWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::service3::GroupCountLimit>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void VectorService3GroupSizeLimitReset(std::vector<::tl2::service3::GroupSizeLimit>& item);

bool VectorService3GroupSizeLimitWriteJSON(std::ostream& s, const std::vector<::tl2::service3::GroupSizeLimit>& item);
bool VectorService3GroupSizeLimitRead(::basictl::tl_istream & s, std::vector<::tl2::service3::GroupSizeLimit>& item);
bool VectorService3GroupSizeLimitWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service3::GroupSizeLimit>& item);
bool VectorService3GroupSizeLimitReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::service3::GroupSizeLimit>& item);
bool VectorService3GroupSizeLimitWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::service3::GroupSizeLimit>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void VectorService3ProductReset(std::vector<::tl2::service3::Product>& item);

bool VectorService3ProductWriteJSON(std::ostream& s, const std::vector<::tl2::service3::Product>& item, uint32_t nat_t);
bool VectorService3ProductRead(::basictl::tl_istream & s, std::vector<::tl2::service3::Product>& item, uint32_t nat_t);
bool VectorService3ProductWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service3::Product>& item, uint32_t nat_t);
bool VectorService3ProductReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::service3::Product>& item, uint32_t nat_t);
bool VectorService3ProductWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::service3::Product>& item, uint32_t nat_t);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void VectorService3Product0Reset(std::vector<::tl2::service3::Productmode<0>>& item);

bool VectorService3Product0WriteJSON(std::ostream& s, const std::vector<::tl2::service3::Productmode<0>>& item);
bool VectorService3Product0Read(::basictl::tl_istream & s, std::vector<::tl2::service3::Productmode<0>>& item);
bool VectorService3Product0Write(::basictl::tl_ostream & s, const std::vector<::tl2::service3::Productmode<0>>& item);
bool VectorService3Product0ReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::service3::Productmode<0>>& item);
bool VectorService3Product0WriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::service3::Productmode<0>>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

bool VectorService3Product0MaybeWriteJSON(std::ostream & s, const std::optional<std::vector<::tl2::service3::Productmode<0>>>& item);

bool VectorService3Product0MaybeReadBoxed(::basictl::tl_istream & s, std::optional<std::vector<::tl2::service3::Productmode<0>>>& item);
bool VectorService3Product0MaybeWriteBoxed(::basictl::tl_ostream & s, const std::optional<std::vector<::tl2::service3::Productmode<0>>>& item);


}} // namespace tl2::details

namespace tl2 { namespace details { 

bool VectorService3ProductMaybeWriteJSON(std::ostream & s, const std::optional<std::vector<::tl2::service3::Product>>& item, uint32_t nat_t);

bool VectorService3ProductMaybeReadBoxed(::basictl::tl_istream & s, std::optional<std::vector<::tl2::service3::Product>>& item, uint32_t nat_t);
bool VectorService3ProductMaybeWriteBoxed(::basictl::tl_ostream & s, const std::optional<std::vector<::tl2::service3::Product>>& item, uint32_t nat_t);


}} // namespace tl2::details

namespace tl2 { namespace details { 

void VectorService3ProductStatsOldReset(std::vector<::tl2::service3::ProductStatsOld>& item);

bool VectorService3ProductStatsOldWriteJSON(std::ostream& s, const std::vector<::tl2::service3::ProductStatsOld>& item);
bool VectorService3ProductStatsOldRead(::basictl::tl_istream & s, std::vector<::tl2::service3::ProductStatsOld>& item);
bool VectorService3ProductStatsOldWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service3::ProductStatsOld>& item);
bool VectorService3ProductStatsOldReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::service3::ProductStatsOld>& item);
bool VectorService3ProductStatsOldWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::service3::ProductStatsOld>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

bool VectorService3ProductStatsOldMaybeWriteJSON(std::ostream & s, const std::optional<std::vector<::tl2::service3::ProductStatsOld>>& item);

bool VectorService3ProductStatsOldMaybeReadBoxed(::basictl::tl_istream & s, std::optional<std::vector<::tl2::service3::ProductStatsOld>>& item);
bool VectorService3ProductStatsOldMaybeWriteBoxed(::basictl::tl_ostream & s, const std::optional<std::vector<::tl2::service3::ProductStatsOld>>& item);


}} // namespace tl2::details

