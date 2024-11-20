#include "headers/service4.modifiedNewsEntry.h"
#include "headers/service4.object.h"
#include "../__common_namespace/headers/int.h"
#include "../__common_namespace/headers/Bool.h"


bool tl2::service4::ModifiedNewsEntry::write_json(std::ostream& s)const {
	if (!::tl2::details::Service4ModifiedNewsEntryWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service4::ModifiedNewsEntry::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service4ModifiedNewsEntryRead(s, *this)) { return false; }
	return true;
}

bool tl2::service4::ModifiedNewsEntry::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service4ModifiedNewsEntryWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service4::ModifiedNewsEntry::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service4ModifiedNewsEntryReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service4::ModifiedNewsEntry::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service4ModifiedNewsEntryWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service4ModifiedNewsEntryReset(::tl2::service4::ModifiedNewsEntry& item) {
	::tl2::details::Service4ObjectReset(item.object);
	item.creation_date = 0;
	item.fields_mask = 0;
	item.restoration_date = 0;
	item.deletion_date = 0;
	item.hidden_by_privacy = false;
}

bool tl2::details::Service4ModifiedNewsEntryWriteJSON(std::ostream& s, const ::tl2::service4::ModifiedNewsEntry& item) {
	auto add_comma = false;
	s << "{";
	add_comma = true;
	s << "\"object\":";
	if (!::tl2::details::Service4ObjectWriteJSON(s, item.object)) { return false; }
	if (item.creation_date != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"creation_date\":";
		s << item.creation_date;
	}
	if (item.fields_mask != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"fields_mask\":";
		s << item.fields_mask;
	}
	if ((item.fields_mask & (1<<0)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"restoration_date\":";
		s << item.restoration_date;
	}
	if ((item.fields_mask & (1<<1)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"deletion_date\":";
		s << item.deletion_date;
	}
	if ((item.fields_mask & (1<<16)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"hidden_by_privacy\":";
		if (!::tl2::details::BoolWriteJSON(s, item.hidden_by_privacy)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::Service4ModifiedNewsEntryRead(::basictl::tl_istream & s, ::tl2::service4::ModifiedNewsEntry& item) {
	if (!::tl2::details::Service4ObjectRead(s, item.object)) { return false; }
	if (!s.int_read(item.creation_date)) { return false; }
	if (!s.nat_read(item.fields_mask)) { return false; }
	if ((item.fields_mask & (1<<0)) != 0) {
		if (!s.int_read(item.restoration_date)) { return false; }
	} else {
			item.restoration_date = 0;
	}
	if ((item.fields_mask & (1<<1)) != 0) {
		if (!s.int_read(item.deletion_date)) { return false; }
	} else {
			item.deletion_date = 0;
	}
	if ((item.fields_mask & (1<<16)) != 0) {
		if (!::tl2::details::BoolReadBoxed(s, item.hidden_by_privacy)) { return false; }
	} else {
			item.hidden_by_privacy = false;
	}
	return true;
}

bool tl2::details::Service4ModifiedNewsEntryWrite(::basictl::tl_ostream & s, const ::tl2::service4::ModifiedNewsEntry& item) {
	if (!::tl2::details::Service4ObjectWrite(s, item.object)) { return false; }
	if (!s.int_write(item.creation_date)) { return false;}
	if (!s.nat_write(item.fields_mask)) { return false;}
	if ((item.fields_mask & (1<<0)) != 0) {
			if (!s.int_write(item.restoration_date)) { return false;}
	}
	if ((item.fields_mask & (1<<1)) != 0) {
			if (!s.int_write(item.deletion_date)) { return false;}
	}
	if ((item.fields_mask & (1<<16)) != 0) {
			if (!::tl2::details::BoolWriteBoxed(s, item.hidden_by_privacy)) { return false; }
	}
	return true;
}

bool tl2::details::Service4ModifiedNewsEntryReadBoxed(::basictl::tl_istream & s, ::tl2::service4::ModifiedNewsEntry& item) {
	if (!s.nat_read_exact_tag(0xda19832a)) { return false; }
	return tl2::details::Service4ModifiedNewsEntryRead(s, item);
}

bool tl2::details::Service4ModifiedNewsEntryWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service4::ModifiedNewsEntry& item) {
	if (!s.nat_write(0xda19832a)) { return false; }
	return tl2::details::Service4ModifiedNewsEntryWrite(s, item);
}

bool tl2::service4::Object::write_json(std::ostream& s)const {
	if (!::tl2::details::Service4ObjectWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service4::Object::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service4ObjectRead(s, *this)) { return false; }
	return true;
}

bool tl2::service4::Object::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service4ObjectWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service4::Object::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service4ObjectReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service4::Object::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service4ObjectWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service4ObjectReset(::tl2::service4::Object& item) {
	item.type = 0;
	item.joint_id.clear();
	item.object_id.clear();
}

bool tl2::details::Service4ObjectWriteJSON(std::ostream& s, const ::tl2::service4::Object& item) {
	auto add_comma = false;
	s << "{";
	if (item.type != 0) {
		add_comma = true;
		s << "\"type\":";
		s << item.type;
	}
	if (item.joint_id.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"joint_id\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.joint_id)) { return false; }
	}
	if (item.object_id.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"object_id\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.object_id)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::Service4ObjectRead(::basictl::tl_istream & s, ::tl2::service4::Object& item) {
	if (!s.int_read(item.type)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, item.joint_id)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, item.object_id)) { return false; }
	return true;
}

bool tl2::details::Service4ObjectWrite(::basictl::tl_ostream & s, const ::tl2::service4::Object& item) {
	if (!s.int_write(item.type)) { return false;}
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.joint_id)) { return false; }
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.object_id)) { return false; }
	return true;
}

bool tl2::details::Service4ObjectReadBoxed(::basictl::tl_istream & s, ::tl2::service4::Object& item) {
	if (!s.nat_read_exact_tag(0xa6eeca4f)) { return false; }
	return tl2::details::Service4ObjectRead(s, item);
}

bool tl2::details::Service4ObjectWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service4::Object& item) {
	if (!s.nat_write(0xa6eeca4f)) { return false; }
	return tl2::details::Service4ObjectWrite(s, item);
}
