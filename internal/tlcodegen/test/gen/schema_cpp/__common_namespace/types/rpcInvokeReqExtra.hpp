#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "true.hpp"


namespace tl2 { 
struct RpcInvokeReqExtra {
	uint32_t fields_mask = 0;
	::tl2::True query{};
	::tl2::True sort{};
	::tl2::True sort_reverse{};
	int64_t wait_binlog_pos = 0;
	std::vector<std::string> string_forward_keys;

	std::string_view tl_name() const { return "rpcInvokeReqExtra"; }
	uint32_t tl_tag() const { return 0xf3ef81a9; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const RpcInvokeReqExtra& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

