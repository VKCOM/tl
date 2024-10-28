#pragma once

#include "../../../basics/basictl.h"


namespace tl2 { namespace unique { 
struct Get {
	// tl magic for function
	static constexpr uint32_t MAGIC() { return 0xce89bbf2; }

	std::string key;

	std::string_view tl_name() const { return "unique.get"; }
	uint32_t tl_tag() const { return 0xce89bbf2; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::optional<int32_t> & result);
	bool write_result(::basictl::tl_ostream & s, std::optional<int32_t> & result);

	friend std::ostream& operator<<(std::ostream& s, const Get& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::unique
