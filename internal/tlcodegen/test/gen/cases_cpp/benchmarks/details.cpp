#include "headers/benchmarks_vector.h"
#include "headers/benchmarks.vrutoyTopLevelContainerWithDependency.h"
#include "headers/benchmarks.vrutoyTopLevelContainer.h"
#include "headers/benchmarks.VrutoyTopLevelUnion.h"
#include "headers/benchmarks.vrutoytopLevelUnionEmpty.h"
#include "headers/benchmarks.vrutoytopLevelUnionBig.h"
#include "headers/benchmarks.vruposition.h"
#include "headers/benchmarks.vrutoyPositions.h"
#include "../__common_namespace/headers/true.h"
#include "headers/benchmarks.vruhash.h"


bool tl2::benchmarks::Vruhash::write_json(std::ostream& s)const {
	if (!::tl2::details::BenchmarksVruHashWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::benchmarks::Vruhash::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BenchmarksVruHashRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::benchmarks::Vruhash::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BenchmarksVruHashWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::benchmarks::Vruhash::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::benchmarks::Vruhash::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::benchmarks::Vruhash::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BenchmarksVruHashReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::benchmarks::Vruhash::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BenchmarksVruHashWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::benchmarks::Vruhash::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::benchmarks::Vruhash::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::BenchmarksVruHashReset(::tl2::benchmarks::Vruhash& item) noexcept {
	item.low = 0;
	item.high = 0;
}

bool tl2::details::BenchmarksVruHashWriteJSON(std::ostream& s, const ::tl2::benchmarks::Vruhash& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.low != 0) {
		add_comma = true;
		s << "\"low\":";
		s << item.low;
	}
	if (item.high != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"high\":";
		s << item.high;
	}
	s << "}";
	return true;
}

bool tl2::details::BenchmarksVruHashRead(::basictl::tl_istream & s, ::tl2::benchmarks::Vruhash& item) noexcept {
	if (!s.long_read(item.low)) { return false; }
	if (!s.long_read(item.high)) { return false; }
	return true;
}

bool tl2::details::BenchmarksVruHashWrite(::basictl::tl_ostream & s, const ::tl2::benchmarks::Vruhash& item) noexcept {
	if (!s.long_write(item.low)) { return false;}
	if (!s.long_write(item.high)) { return false;}
	return true;
}

bool tl2::details::BenchmarksVruHashReadBoxed(::basictl::tl_istream & s, ::tl2::benchmarks::Vruhash& item) {
	if (!s.nat_read_exact_tag(0xd31bd0fd)) { return false; }
	return tl2::details::BenchmarksVruHashRead(s, item);
}

bool tl2::details::BenchmarksVruHashWriteBoxed(::basictl::tl_ostream & s, const ::tl2::benchmarks::Vruhash& item) {
	if (!s.nat_write(0xd31bd0fd)) { return false; }
	return tl2::details::BenchmarksVruHashWrite(s, item);
}

bool tl2::benchmarks::Vruposition::write_json(std::ostream& s)const {
	if (!::tl2::details::BenchmarksVruPositionWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::benchmarks::Vruposition::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BenchmarksVruPositionRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::benchmarks::Vruposition::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BenchmarksVruPositionWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::benchmarks::Vruposition::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::benchmarks::Vruposition::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::benchmarks::Vruposition::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BenchmarksVruPositionReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::benchmarks::Vruposition::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BenchmarksVruPositionWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::benchmarks::Vruposition::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::benchmarks::Vruposition::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::BenchmarksVruPositionReset(::tl2::benchmarks::Vruposition& item) noexcept {
	item.fields_mask = 0;
	::tl2::details::TrueReset(item.commit_bit);
	::tl2::details::TrueReset(item.meta_block);
	::tl2::details::TrueReset(item.split_payload);
	::tl2::details::TrueReset(item.rotation_block);
	::tl2::details::TrueReset(item.canonical_hash);
	item.payload_offset = 0;
	item.block_time_nano = 0;
	::tl2::details::BenchmarksVruHashReset(item.hash);
	item.file_offset = 0;
	item.seq_number = 0;
}

bool tl2::details::BenchmarksVruPositionWriteJSON(std::ostream& s, const ::tl2::benchmarks::Vruposition& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.fields_mask != 0) {
		add_comma = true;
		s << "\"fields_mask\":";
		s << item.fields_mask;
	}
	if ((item.fields_mask & (1<<0)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"commit_bit\":";
		if (!::tl2::details::TrueWriteJSON(s, item.commit_bit)) { return false; }
	}
	if ((item.fields_mask & (1<<1)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"meta_block\":";
		if (!::tl2::details::TrueWriteJSON(s, item.meta_block)) { return false; }
	}
	if ((item.fields_mask & (1<<3)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"split_payload\":";
		if (!::tl2::details::TrueWriteJSON(s, item.split_payload)) { return false; }
	}
	if ((item.fields_mask & (1<<5)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"rotation_block\":";
		if (!::tl2::details::TrueWriteJSON(s, item.rotation_block)) { return false; }
	}
	if ((item.fields_mask & (1<<15)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"canonical_hash\":";
		if (!::tl2::details::TrueWriteJSON(s, item.canonical_hash)) { return false; }
	}
	if (item.payload_offset != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"payload_offset\":";
		s << item.payload_offset;
	}
	if (item.block_time_nano != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"block_time_nano\":";
		s << item.block_time_nano;
	}
	if (add_comma) {
		s << ",";
	}
	add_comma = true;
	s << "\"hash\":";
	if (!::tl2::details::BenchmarksVruHashWriteJSON(s, item.hash)) { return false; }
	if (item.file_offset != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"file_offset\":";
		s << item.file_offset;
	}
	if ((item.fields_mask & (1<<14)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"seq_number\":";
		s << item.seq_number;
	}
	s << "}";
	return true;
}

bool tl2::details::BenchmarksVruPositionRead(::basictl::tl_istream & s, ::tl2::benchmarks::Vruposition& item) noexcept {
	if (!s.nat_read(item.fields_mask)) { return false; }
	if ((item.fields_mask & (1<<0)) != 0) {
		if (!::tl2::details::TrueRead(s, item.commit_bit)) { return false; }
	} else {
			::tl2::details::TrueReset(item.commit_bit);
	}
	if ((item.fields_mask & (1<<1)) != 0) {
		if (!::tl2::details::TrueRead(s, item.meta_block)) { return false; }
	} else {
			::tl2::details::TrueReset(item.meta_block);
	}
	if ((item.fields_mask & (1<<3)) != 0) {
		if (!::tl2::details::TrueRead(s, item.split_payload)) { return false; }
	} else {
			::tl2::details::TrueReset(item.split_payload);
	}
	if ((item.fields_mask & (1<<5)) != 0) {
		if (!::tl2::details::TrueRead(s, item.rotation_block)) { return false; }
	} else {
			::tl2::details::TrueReset(item.rotation_block);
	}
	if ((item.fields_mask & (1<<15)) != 0) {
		if (!::tl2::details::TrueRead(s, item.canonical_hash)) { return false; }
	} else {
			::tl2::details::TrueReset(item.canonical_hash);
	}
	if (!s.long_read(item.payload_offset)) { return false; }
	if (!s.long_read(item.block_time_nano)) { return false; }
	if (!::tl2::details::BenchmarksVruHashRead(s, item.hash)) { return false; }
	if (!s.long_read(item.file_offset)) { return false; }
	if ((item.fields_mask & (1<<14)) != 0) {
		if (!s.long_read(item.seq_number)) { return false; }
	} else {
			item.seq_number = 0;
	}
	return true;
}

bool tl2::details::BenchmarksVruPositionWrite(::basictl::tl_ostream & s, const ::tl2::benchmarks::Vruposition& item) noexcept {
	if (!s.nat_write(item.fields_mask)) { return false;}
	if ((item.fields_mask & (1<<0)) != 0) {
			if (!::tl2::details::TrueWrite(s, item.commit_bit)) { return false; }
	}
	if ((item.fields_mask & (1<<1)) != 0) {
			if (!::tl2::details::TrueWrite(s, item.meta_block)) { return false; }
	}
	if ((item.fields_mask & (1<<3)) != 0) {
			if (!::tl2::details::TrueWrite(s, item.split_payload)) { return false; }
	}
	if ((item.fields_mask & (1<<5)) != 0) {
			if (!::tl2::details::TrueWrite(s, item.rotation_block)) { return false; }
	}
	if ((item.fields_mask & (1<<15)) != 0) {
			if (!::tl2::details::TrueWrite(s, item.canonical_hash)) { return false; }
	}
	if (!s.long_write(item.payload_offset)) { return false;}
	if (!s.long_write(item.block_time_nano)) { return false;}
	if (!::tl2::details::BenchmarksVruHashWrite(s, item.hash)) { return false; }
	if (!s.long_write(item.file_offset)) { return false;}
	if ((item.fields_mask & (1<<14)) != 0) {
			if (!s.long_write(item.seq_number)) { return false;}
	}
	return true;
}

bool tl2::details::BenchmarksVruPositionReadBoxed(::basictl::tl_istream & s, ::tl2::benchmarks::Vruposition& item) {
	if (!s.nat_read_exact_tag(0x32792c04)) { return false; }
	return tl2::details::BenchmarksVruPositionRead(s, item);
}

bool tl2::details::BenchmarksVruPositionWriteBoxed(::basictl::tl_ostream & s, const ::tl2::benchmarks::Vruposition& item) {
	if (!s.nat_write(0x32792c04)) { return false; }
	return tl2::details::BenchmarksVruPositionWrite(s, item);
}

bool tl2::benchmarks::VrutoyPositions::write_json(std::ostream& s, uint32_t nat_n)const {
	if (!::tl2::details::BenchmarksVrutoyPositionsWriteJSON(s, *this, nat_n)) { return false; }
	return true;
}

bool tl2::benchmarks::VrutoyPositions::read(::basictl::tl_istream & s, uint32_t nat_n) noexcept {
	if (!::tl2::details::BenchmarksVrutoyPositionsRead(s, *this, nat_n)) { return false; }
	s.last_release();
	return true;
}

bool tl2::benchmarks::VrutoyPositions::write(::basictl::tl_ostream & s, uint32_t nat_n)const noexcept {
	if (!::tl2::details::BenchmarksVrutoyPositionsWrite(s, *this, nat_n)) { return false; }
	s.last_release();
	return true;
}

void tl2::benchmarks::VrutoyPositions::read_or_throw(::basictl::tl_throwable_istream & s, uint32_t nat_n) {
	::basictl::tl_istream s2(s);
	this->read(s2, nat_n);
	s2.pass_data(s);
}

void tl2::benchmarks::VrutoyPositions::write_or_throw(::basictl::tl_throwable_ostream & s, uint32_t nat_n)const {
	::basictl::tl_ostream s2(s);
	this->write(s2, nat_n);
	s2.pass_data(s);
}

bool tl2::benchmarks::VrutoyPositions::read_boxed(::basictl::tl_istream & s, uint32_t nat_n) noexcept {
	if (!::tl2::details::BenchmarksVrutoyPositionsReadBoxed(s, *this, nat_n)) { return false; }
	s.last_release();
	return true;
}

bool tl2::benchmarks::VrutoyPositions::write_boxed(::basictl::tl_ostream & s, uint32_t nat_n)const noexcept {
	if (!::tl2::details::BenchmarksVrutoyPositionsWriteBoxed(s, *this, nat_n)) { return false; }
	s.last_release();
	return true;
}

void tl2::benchmarks::VrutoyPositions::read_boxed_or_throw(::basictl::tl_throwable_istream & s, uint32_t nat_n) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2, nat_n);
	s2.pass_data(s);
}

void tl2::benchmarks::VrutoyPositions::write_boxed_or_throw(::basictl::tl_throwable_ostream & s, uint32_t nat_n)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2, nat_n);
	s2.pass_data(s);
}

void tl2::details::BenchmarksVrutoyPositionsReset(::tl2::benchmarks::VrutoyPositions& item) noexcept {
	item.next_positions.clear();
}

bool tl2::details::BenchmarksVrutoyPositionsWriteJSON(std::ostream& s, const ::tl2::benchmarks::VrutoyPositions& item, uint32_t nat_n) noexcept {
	s << "{";
	if ((item.next_positions.size() != 0) || (nat_n != 0)) {
		s << "\"next_positions\":";
		if (!::tl2::details::BuiltinTupleBenchmarksVruPositionWriteJSON(s, item.next_positions, nat_n)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::BenchmarksVrutoyPositionsRead(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoyPositions& item, uint32_t nat_n) noexcept {
	if (!::tl2::details::BuiltinTupleBenchmarksVruPositionRead(s, item.next_positions, nat_n)) { return false; }
	return true;
}

bool tl2::details::BenchmarksVrutoyPositionsWrite(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoyPositions& item, uint32_t nat_n) noexcept {
	if (!::tl2::details::BuiltinTupleBenchmarksVruPositionWrite(s, item.next_positions, nat_n)) { return false; }
	return true;
}

bool tl2::details::BenchmarksVrutoyPositionsReadBoxed(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoyPositions& item, uint32_t nat_n) {
	if (!s.nat_read_exact_tag(0xb6003de0)) { return false; }
	return tl2::details::BenchmarksVrutoyPositionsRead(s, item, nat_n);
}

bool tl2::details::BenchmarksVrutoyPositionsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoyPositions& item, uint32_t nat_n) {
	if (!s.nat_write(0xb6003de0)) { return false; }
	return tl2::details::BenchmarksVrutoyPositionsWrite(s, item, nat_n);
}

bool tl2::benchmarks::VrutoyTopLevelContainer::write_json(std::ostream& s)const {
	if (!::tl2::details::BenchmarksVrutoyTopLevelContainerWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::benchmarks::VrutoyTopLevelContainer::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BenchmarksVrutoyTopLevelContainerRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::benchmarks::VrutoyTopLevelContainer::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BenchmarksVrutoyTopLevelContainerWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::benchmarks::VrutoyTopLevelContainer::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::benchmarks::VrutoyTopLevelContainer::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::benchmarks::VrutoyTopLevelContainer::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BenchmarksVrutoyTopLevelContainerReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::benchmarks::VrutoyTopLevelContainer::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BenchmarksVrutoyTopLevelContainerWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::benchmarks::VrutoyTopLevelContainer::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::benchmarks::VrutoyTopLevelContainer::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::BenchmarksVrutoyTopLevelContainerReset(::tl2::benchmarks::VrutoyTopLevelContainer& item) noexcept {
	::tl2::details::BenchmarksVrutoyTopLevelUnionReset(item.value);
}

bool tl2::details::BenchmarksVrutoyTopLevelContainerWriteJSON(std::ostream& s, const ::tl2::benchmarks::VrutoyTopLevelContainer& item) noexcept {
	s << "{";
	s << "\"value\":";
	if (!::tl2::details::BenchmarksVrutoyTopLevelUnionWriteJSON(s, item.value)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::BenchmarksVrutoyTopLevelContainerRead(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoyTopLevelContainer& item) noexcept {
	if (!::tl2::details::BenchmarksVrutoyTopLevelUnionReadBoxed(s, item.value)) { return false; }
	return true;
}

bool tl2::details::BenchmarksVrutoyTopLevelContainerWrite(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoyTopLevelContainer& item) noexcept {
	if (!::tl2::details::BenchmarksVrutoyTopLevelUnionWriteBoxed(s, item.value)) { return false; }
	return true;
}

bool tl2::details::BenchmarksVrutoyTopLevelContainerReadBoxed(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoyTopLevelContainer& item) {
	if (!s.nat_read_exact_tag(0xfb442ca5)) { return false; }
	return tl2::details::BenchmarksVrutoyTopLevelContainerRead(s, item);
}

bool tl2::details::BenchmarksVrutoyTopLevelContainerWriteBoxed(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoyTopLevelContainer& item) {
	if (!s.nat_write(0xfb442ca5)) { return false; }
	return tl2::details::BenchmarksVrutoyTopLevelContainerWrite(s, item);
}

bool tl2::benchmarks::VrutoyTopLevelContainerWithDependency::write_json(std::ostream& s)const {
	if (!::tl2::details::BenchmarksVrutoyTopLevelContainerWithDependencyWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::benchmarks::VrutoyTopLevelContainerWithDependency::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BenchmarksVrutoyTopLevelContainerWithDependencyRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::benchmarks::VrutoyTopLevelContainerWithDependency::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BenchmarksVrutoyTopLevelContainerWithDependencyWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::benchmarks::VrutoyTopLevelContainerWithDependency::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::benchmarks::VrutoyTopLevelContainerWithDependency::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::benchmarks::VrutoyTopLevelContainerWithDependency::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BenchmarksVrutoyTopLevelContainerWithDependencyReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::benchmarks::VrutoyTopLevelContainerWithDependency::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BenchmarksVrutoyTopLevelContainerWithDependencyWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::benchmarks::VrutoyTopLevelContainerWithDependency::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::benchmarks::VrutoyTopLevelContainerWithDependency::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::BenchmarksVrutoyTopLevelContainerWithDependencyReset(::tl2::benchmarks::VrutoyTopLevelContainerWithDependency& item) noexcept {
	item.n = 0;
	::tl2::details::BenchmarksVrutoyPositionsReset(item.value);
}

bool tl2::details::BenchmarksVrutoyTopLevelContainerWithDependencyWriteJSON(std::ostream& s, const ::tl2::benchmarks::VrutoyTopLevelContainerWithDependency& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.n != 0) {
		add_comma = true;
		s << "\"n\":";
		s << item.n;
	}
	if (add_comma) {
		s << ",";
	}
	add_comma = true;
	s << "\"value\":";
	if (!::tl2::details::BenchmarksVrutoyPositionsWriteJSON(s, item.value, item.n)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::BenchmarksVrutoyTopLevelContainerWithDependencyRead(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoyTopLevelContainerWithDependency& item) noexcept {
	if (!s.nat_read(item.n)) { return false; }
	if (!::tl2::details::BenchmarksVrutoyPositionsRead(s, item.value, item.n)) { return false; }
	return true;
}

bool tl2::details::BenchmarksVrutoyTopLevelContainerWithDependencyWrite(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoyTopLevelContainerWithDependency& item) noexcept {
	if (!s.nat_write(item.n)) { return false;}
	if (!::tl2::details::BenchmarksVrutoyPositionsWrite(s, item.value, item.n)) { return false; }
	return true;
}

bool tl2::details::BenchmarksVrutoyTopLevelContainerWithDependencyReadBoxed(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoyTopLevelContainerWithDependency& item) {
	if (!s.nat_read_exact_tag(0xc176008e)) { return false; }
	return tl2::details::BenchmarksVrutoyTopLevelContainerWithDependencyRead(s, item);
}

bool tl2::details::BenchmarksVrutoyTopLevelContainerWithDependencyWriteBoxed(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoyTopLevelContainerWithDependency& item) {
	if (!s.nat_write(0xc176008e)) { return false; }
	return tl2::details::BenchmarksVrutoyTopLevelContainerWithDependencyWrite(s, item);
}

static const std::string_view BenchmarksVrutoyTopLevelUnion_tbl_tl_name[]{"benchmarks.vrutoytopLevelUnionBig", "benchmarks.vrutoytopLevelUnionEmpty"};
static const uint32_t BenchmarksVrutoyTopLevelUnion_tbl_tl_tag[]{0xef556bee, 0xce27c770};

bool tl2::benchmarks::VrutoyTopLevelUnion::write_json(std::ostream & s)const {
	if (!::tl2::details::BenchmarksVrutoyTopLevelUnionWriteJSON(s, *this)) { return false; }
	return true;
}
bool tl2::benchmarks::VrutoyTopLevelUnion::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BenchmarksVrutoyTopLevelUnionReadBoxed(s, *this)) { return false; }
	return true;
}
bool tl2::benchmarks::VrutoyTopLevelUnion::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BenchmarksVrutoyTopLevelUnionWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::benchmarks::VrutoyTopLevelUnion::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::benchmarks::VrutoyTopLevelUnion::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

std::string_view tl2::benchmarks::VrutoyTopLevelUnion::tl_name() const {
	return BenchmarksVrutoyTopLevelUnion_tbl_tl_name[value.index()];
}
uint32_t tl2::benchmarks::VrutoyTopLevelUnion::tl_tag() const {
	return BenchmarksVrutoyTopLevelUnion_tbl_tl_tag[value.index()];
}


void tl2::details::BenchmarksVrutoyTopLevelUnionReset(::tl2::benchmarks::VrutoyTopLevelUnion& item) noexcept{
	item.value.emplace<0>(); // TODO - optimize, if already 0, call Reset function
}

bool tl2::details::BenchmarksVrutoyTopLevelUnionWriteJSON(std::ostream & s, const ::tl2::benchmarks::VrutoyTopLevelUnion& item) noexcept {
	s << "{";
	s << "\"type\":";
	s << "\"" << BenchmarksVrutoyTopLevelUnion_tbl_tl_name[item.value.index()] << "\"";
	switch (item.value.index()) {
	case 0:
		s << ",\"value\":";
		if (!::tl2::details::BenchmarksVrutoytopLevelUnionBigWriteJSON(s, std::get<0>(item.value))) { return false; }
		break;
	}
	s << "}";
	return true;
}
bool tl2::details::BenchmarksVrutoyTopLevelUnionReadBoxed(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoyTopLevelUnion& item) noexcept {
	uint32_t nat;
	s.nat_read(nat);
	switch (nat) {
	case 0xef556bee:
		if (item.value.index() != 0) { item.value.emplace<0>(); }
		if (!::tl2::details::BenchmarksVrutoytopLevelUnionBigRead(s, std::get<0>(item.value))) { return false; }
		break;
	case 0xce27c770:
		if (item.value.index() != 1) { item.value.emplace<1>(); }
		break;
	default:
		return s.set_error_union_tag();
    }
	return true;
}

bool tl2::details::BenchmarksVrutoyTopLevelUnionWriteBoxed(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoyTopLevelUnion& item) noexcept{
	s.nat_write(BenchmarksVrutoyTopLevelUnion_tbl_tl_tag[item.value.index()]);
	switch (item.value.index()) {
	case 0:
		if (!::tl2::details::BenchmarksVrutoytopLevelUnionBigWrite(s, std::get<0>(item.value))) { return false; }
		break;
	}
	return true;
}

bool tl2::benchmarks::VrutoytopLevelUnionBig::write_json(std::ostream& s)const {
	if (!::tl2::details::BenchmarksVrutoytopLevelUnionBigWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::benchmarks::VrutoytopLevelUnionBig::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BenchmarksVrutoytopLevelUnionBigRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::benchmarks::VrutoytopLevelUnionBig::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BenchmarksVrutoytopLevelUnionBigWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::benchmarks::VrutoytopLevelUnionBig::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::benchmarks::VrutoytopLevelUnionBig::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::benchmarks::VrutoytopLevelUnionBig::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BenchmarksVrutoytopLevelUnionBigReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::benchmarks::VrutoytopLevelUnionBig::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BenchmarksVrutoytopLevelUnionBigWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::benchmarks::VrutoytopLevelUnionBig::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::benchmarks::VrutoytopLevelUnionBig::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::BenchmarksVrutoytopLevelUnionBigReset(::tl2::benchmarks::VrutoytopLevelUnionBig& item) noexcept {
	item.next_positions.clear();
}

bool tl2::details::BenchmarksVrutoytopLevelUnionBigWriteJSON(std::ostream& s, const ::tl2::benchmarks::VrutoytopLevelUnionBig& item) noexcept {
	s << "{";
	if (item.next_positions.size() != 0) {
		s << "\"next_positions\":";
		if (!::tl2::details::BuiltinVectorBenchmarksVruPositionWriteJSON(s, item.next_positions)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::BenchmarksVrutoytopLevelUnionBigRead(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoytopLevelUnionBig& item) noexcept {
	if (!::tl2::details::BuiltinVectorBenchmarksVruPositionRead(s, item.next_positions)) { return false; }
	return true;
}

bool tl2::details::BenchmarksVrutoytopLevelUnionBigWrite(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoytopLevelUnionBig& item) noexcept {
	if (!::tl2::details::BuiltinVectorBenchmarksVruPositionWrite(s, item.next_positions)) { return false; }
	return true;
}

bool tl2::details::BenchmarksVrutoytopLevelUnionBigReadBoxed(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoytopLevelUnionBig& item) {
	if (!s.nat_read_exact_tag(0xef556bee)) { return false; }
	return tl2::details::BenchmarksVrutoytopLevelUnionBigRead(s, item);
}

bool tl2::details::BenchmarksVrutoytopLevelUnionBigWriteBoxed(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoytopLevelUnionBig& item) {
	if (!s.nat_write(0xef556bee)) { return false; }
	return tl2::details::BenchmarksVrutoytopLevelUnionBigWrite(s, item);
}

bool tl2::benchmarks::VrutoytopLevelUnionEmpty::write_json(std::ostream& s)const {
	if (!::tl2::details::BenchmarksVrutoytopLevelUnionEmptyWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::benchmarks::VrutoytopLevelUnionEmpty::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BenchmarksVrutoytopLevelUnionEmptyRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::benchmarks::VrutoytopLevelUnionEmpty::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BenchmarksVrutoytopLevelUnionEmptyWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::benchmarks::VrutoytopLevelUnionEmpty::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::benchmarks::VrutoytopLevelUnionEmpty::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::benchmarks::VrutoytopLevelUnionEmpty::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BenchmarksVrutoytopLevelUnionEmptyReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::benchmarks::VrutoytopLevelUnionEmpty::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BenchmarksVrutoytopLevelUnionEmptyWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::benchmarks::VrutoytopLevelUnionEmpty::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::benchmarks::VrutoytopLevelUnionEmpty::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::BenchmarksVrutoytopLevelUnionEmptyReset(::tl2::benchmarks::VrutoytopLevelUnionEmpty& item) noexcept {
}

bool tl2::details::BenchmarksVrutoytopLevelUnionEmptyWriteJSON(std::ostream& s, const ::tl2::benchmarks::VrutoytopLevelUnionEmpty& item) noexcept {
	s << "true";
	return true;
}

bool tl2::details::BenchmarksVrutoytopLevelUnionEmptyRead(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoytopLevelUnionEmpty& item) noexcept {
	return true;
}

bool tl2::details::BenchmarksVrutoytopLevelUnionEmptyWrite(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoytopLevelUnionEmpty& item) noexcept {
	return true;
}

bool tl2::details::BenchmarksVrutoytopLevelUnionEmptyReadBoxed(::basictl::tl_istream & s, ::tl2::benchmarks::VrutoytopLevelUnionEmpty& item) {
	if (!s.nat_read_exact_tag(0xce27c770)) { return false; }
	return tl2::details::BenchmarksVrutoytopLevelUnionEmptyRead(s, item);
}

bool tl2::details::BenchmarksVrutoytopLevelUnionEmptyWriteBoxed(::basictl::tl_ostream & s, const ::tl2::benchmarks::VrutoytopLevelUnionEmpty& item) {
	if (!s.nat_write(0xce27c770)) { return false; }
	return tl2::details::BenchmarksVrutoytopLevelUnionEmptyWrite(s, item);
}

void tl2::details::BuiltinTupleBenchmarksVruPositionReset(std::vector<::tl2::benchmarks::Vruposition>& item) {
	item.resize(0);
}

bool tl2::details::BuiltinTupleBenchmarksVruPositionWriteJSON(std::ostream & s, const std::vector<::tl2::benchmarks::Vruposition>& item, uint32_t nat_n) {
	if (item.size() != nat_n) {
		// TODO add exception
		return false;
	}
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		if (!::tl2::details::BenchmarksVruPositionWriteJSON(s, el)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinTupleBenchmarksVruPositionRead(::basictl::tl_istream & s, std::vector<::tl2::benchmarks::Vruposition>& item, uint32_t nat_n) {
	// TODO - check length sanity
	item.resize(nat_n);
	for(auto && el : item) {
		if (!::tl2::details::BenchmarksVruPositionRead(s, el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinTupleBenchmarksVruPositionWrite(::basictl::tl_ostream & s, const std::vector<::tl2::benchmarks::Vruposition>& item, uint32_t nat_n) {
	if (item.size() != nat_n)
		return s.set_error_sequence_length();
	for(const auto & el : item) {
		if (!::tl2::details::BenchmarksVruPositionWrite(s, el)) { return false; }
	}
	return true;
}

void tl2::details::BuiltinVectorBenchmarksVruPositionReset(std::vector<::tl2::benchmarks::Vruposition>& item) {
	item.resize(0); // TODO - unwrap
}

bool tl2::details::BuiltinVectorBenchmarksVruPositionWriteJSON(std::ostream & s, const std::vector<::tl2::benchmarks::Vruposition>& item) {
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		if (!::tl2::details::BenchmarksVruPositionWriteJSON(s, el)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinVectorBenchmarksVruPositionRead(::basictl::tl_istream & s, std::vector<::tl2::benchmarks::Vruposition>& item) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	// TODO - check length sanity
	item.resize(len);
	for(auto && el : item) {
		if (!::tl2::details::BenchmarksVruPositionRead(s, el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinVectorBenchmarksVruPositionWrite(::basictl::tl_ostream & s, const std::vector<::tl2::benchmarks::Vruposition>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!::tl2::details::BenchmarksVruPositionWrite(s, el)) { return false; }
	}
	return true;
}

void tl2::details::VectorBenchmarksVruPositionReset(std::vector<::tl2::benchmarks::Vruposition>& item) noexcept {
	item.clear();
}

bool tl2::details::VectorBenchmarksVruPositionWriteJSON(std::ostream& s, const std::vector<::tl2::benchmarks::Vruposition>& item) noexcept {
	if (!::tl2::details::BuiltinVectorBenchmarksVruPositionWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorBenchmarksVruPositionRead(::basictl::tl_istream & s, std::vector<::tl2::benchmarks::Vruposition>& item) noexcept {
	if (!::tl2::details::BuiltinVectorBenchmarksVruPositionRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorBenchmarksVruPositionWrite(::basictl::tl_ostream & s, const std::vector<::tl2::benchmarks::Vruposition>& item) noexcept {
	if (!::tl2::details::BuiltinVectorBenchmarksVruPositionWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorBenchmarksVruPositionReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::benchmarks::Vruposition>& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	return tl2::details::VectorBenchmarksVruPositionRead(s, item);
}

bool tl2::details::VectorBenchmarksVruPositionWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::benchmarks::Vruposition>& item) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	return tl2::details::VectorBenchmarksVruPositionWrite(s, item);
}
