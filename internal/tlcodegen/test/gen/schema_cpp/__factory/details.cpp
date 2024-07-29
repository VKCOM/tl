#include "../__meta/headers.hpp"
#include "headers.hpp"

#include "../__common/types/withFloat.hpp"
#include "../__common/types/vector.hpp"
#include "../unique/functions/unique.stringToInt.hpp"
#include "../unique/functions/unique.get.hpp"
#include "../__common/types/tuple.hpp"
#include "../tree_stats/types/tree_stats.objectLimitValueLong.hpp"
#include "../tasks/types/tasks.TaskStatusItems.hpp"
#include "../tasks/types/tasks.taskInfo.hpp"
#include "../tasks/types/tasks.queueTypeInfo.hpp"
#include "../tasks/types/tasks.queueTypeSettings.hpp"
#include "../tasks/types/tasks.queueStats.hpp"
#include "../tasks/functions/tasks.getTaskFromQueue.hpp"
#include "../tasks/functions/tasks.getQueueTypes.hpp"
#include "../tasks/functions/tasks.getQueueSize.hpp"
#include "../tasks/functions/tasks.getAnyTask.hpp"
#include "../tasks/types/tasks.cronTaskWithId.hpp"
#include "../tasks/types/tasks.cronTask.hpp"
#include "../tasks/types/tasks.cronTime.hpp"
#include "../tasks/functions/tasks.addTask.hpp"
#include "../tasks/types/tasks.task.hpp"
#include "../__common/types/string.hpp"
#include "../__common/types/statOne.hpp"
#include "../service6/functions/service6.multiFindWithBounds.hpp"
#include "../service6/functions/service6.multiFind.hpp"
#include "../service6/types/service6.findWithBoundsResult.hpp"
#include "../service5/functions/service5.query.hpp"
#include "../service5/functions/service5.performQuery.hpp"
#include "../service5/types/service5.params.hpp"
#include "../service5/types/service5.Output.hpp"
#include "../service5/types/service5.stringOutput.hpp"
#include "../service5/functions/service5.insert.hpp"
#include "../service5/types/service5.emptyOutput.hpp"
#include "../service4/types/service4.modifiedNewsEntry.hpp"
#include "../service4/types/service4.object.hpp"
#include "../service3/functions/service3.setLimits.hpp"
#include "../service3/functions/service3.setLastVisitTimestamp.hpp"
#include "../service3/functions/service3.restoreProduct.hpp"
#include "../service3/functions/service3.restoreGroupedProducts.hpp"
#include "../service3/functions/service3.restoreAllProducts.hpp"
#include "../service3/types/service3.productStatsOld.hpp"
#include "../service3/types/service3.limits.hpp"
#include "../service3/types/service3.groupSizeLimit.hpp"
#include "../service3/types/service3.groupCountLimit.hpp"
#include "../service3/functions/service3.getScheduledProducts.hpp"
#include "../service3/functions/service3.getProducts.hpp"
#include "../service3/functions/service3.getProductStats.hpp"
#include "../service3/functions/service3.getLimits.hpp"
#include "../service3/functions/service3.getLastVisitTimestamp.hpp"
#include "../service3/functions/service3.deleteProduct.hpp"
#include "../service3/functions/service3.deleteGroupedProducts.hpp"
#include "../service3/functions/service3.deleteAllProducts.hpp"
#include "../service3/functions/service3.createProduct.hpp"
#include "../service2/functions/service2.setObjectTtl.hpp"
#include "../service2/functions/service2.set.hpp"
#include "../service2/functions/service2.addOrIncrMany.hpp"
#include "../service1/functions/service1.touch.hpp"
#include "../service1/functions/service1.setOrIncr.hpp"
#include "../service1/functions/service1.set.hpp"
#include "../service1/functions/service1.replaceOrIncr.hpp"
#include "../service1/functions/service1.replace.hpp"
#include "../service1/types/service1.keysStat.hpp"
#include "../service1/functions/service1.incr.hpp"
#include "../service1/functions/service1.getWildcardWithFlags.hpp"
#include "../service1/functions/service1.getWildcardList.hpp"
#include "../service1/functions/service1.getWildcardDict.hpp"
#include "../service1/functions/service1.getWildcard.hpp"
#include "../service1/functions/service1.getKeysStatPeriods.hpp"
#include "../service1/functions/service1.getKeysStat.hpp"
#include "../service1/functions/service1.getExpireTime.hpp"
#include "../service1/functions/service1.get.hpp"
#include "../service1/functions/service1.exists.hpp"
#include "../service1/functions/service1.enableKeysStat.hpp"
#include "../service1/functions/service1.enableExpiration.hpp"
#include "../service1/functions/service1.disableKeysStat.hpp"
#include "../service1/functions/service1.disableExpiration.hpp"
#include "../service1/functions/service1.delete.hpp"
#include "../service1/functions/service1.decr.hpp"
#include "../service1/functions/service1.cas.hpp"
#include "../service1/functions/service1.append.hpp"
#include "../service1/functions/service1.addOrIncr.hpp"
#include "../service1/functions/service1.addOrGet.hpp"
#include "../service1/functions/service1.add.hpp"
#include "../__common/types/rpcInvokeReqExtra.hpp"
#include "../__common/types/true.hpp"
#include "../pkg2/types/pkg2.t1.hpp"
#include "../pkg2/types/pkg2.foo.hpp"
#include "../__common/types/nonOptNat.hpp"
#include "../__common/types/myTwoDicts.hpp"
#include "../__common/types/myMcValueVector.hpp"
#include "../__common/types/myMcValueTuple.hpp"
#include "../__common/types/myMcValue.hpp"
#include "../__common/types/myBoxedVectorSlice.hpp"
#include "../__common/types/long.hpp"
#include "../__common/types/issue3498.hpp"
#include "../service6/types/service6.findResultRow.hpp"
#include "../service6/types/service6.error.hpp"
#include "../__common/types/int.hpp"
#include "../__common/functions/getStats.hpp"
#include "../tasks/types/tasks.queueTypeStats.hpp"
#include "../__common/functions/getNonOptNat.hpp"
#include "../__common/functions/getMyValue.hpp"
#include "../__common/types/MyValue.hpp"
#include "../__common/types/myString.hpp"
#include "../__common/types/myInt.hpp"
#include "../__common/functions/getMyDouble.hpp"
#include "../__common/types/myDouble.hpp"
#include "../__common/functions/getMyDictOfInt.hpp"
#include "../__common/types/myDictOfInt.hpp"
#include "../__common/functions/getMaybeIface.hpp"
#include "../service1/types/service1.Value.hpp"
#include "../service1/types/service1.strvalueWithTime.hpp"
#include "../service1/types/service1.strvalue.hpp"
#include "../service1/types/service1.not_found.hpp"
#include "../service1/types/service1.longvalueWithTime.hpp"
#include "../service1/types/service1.longvalue.hpp"
#include "../__common/functions/getFloat.hpp"
#include "../__common/functions/getDouble.hpp"
#include "../__common/functions/get_arrays.hpp"
#include "../__common/types/float.hpp"
#include "../__common/types/fieldConflict4.hpp"
#include "../__common/types/fieldConflict3.hpp"
#include "../__common/types/fieldConflict2.hpp"
#include "../__common/types/fieldConflict1.hpp"
#include "../__common/types/double.hpp"
#include "../__common/types/dictionary.hpp"
#include "../__common/functions/boxedVector64.hpp"
#include "../__common/functions/boxedVector32BoxedElem.hpp"
#include "../__common/functions/boxedVector32.hpp"
#include "../__common/functions/boxedTupleSlice3.hpp"
#include "../__common/functions/boxedTupleSlice2.hpp"
#include "../__common/types/myBoxedTupleSlice.hpp"
#include "../__common/functions/boxedTupleSlice1.hpp"
#include "../__common/functions/boxedTuple.hpp"
#include "../__common/functions/boxedString.hpp"
#include "../__common/functions/boxedInt.hpp"
#include "../__common/functions/boxedArray.hpp"
#include "../__common/types/myBoxedArray.hpp"
#include "../__common/types/boolStat.hpp"
#include "../__common/types/Bool.hpp"
#include "../__common/types/benchObject.hpp"
#include "../__common/types/integer.hpp"
#include "../antispam/types/antispam.PatternFull.hpp"
#include "../antispam/types/antispam.patternNotFound.hpp"
#include "../antispam/types/antispam.patternFound.hpp"
#include "../antispam/functions/antispam.getPattern.hpp"

void tl2::factory::set_all_factories() {

	struct tl2_antispam_GetPattern_tl_function : public tl2::meta::tl_function {
        tl2::antispam::GetPattern object;
        explicit tl2_antispam_GetPattern_tl_function(tl2::antispam::GetPattern o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			tl2::antispam::PatternFull result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("antispam.getPattern", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_antispam_GetPattern_tl_function>(tl2::antispam::GetPattern{});
	});

	tl2::meta::set_create_function_by_name("antispam.getPattern", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_antispam_GetPattern_tl_function>(tl2::antispam::GetPattern{});
	});

	struct tl2_antispam_PatternFound_tl_object : public tl2::meta::tl_object {
        tl2::antispam::PatternFound object;
        explicit tl2_antispam_PatternFound_tl_object(tl2::antispam::PatternFound o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("antispam.patternFound", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_antispam_PatternFound_tl_object>(tl2::antispam::PatternFound{});
	});

	struct tl2_antispam_PatternNotFound_tl_object : public tl2::meta::tl_object {
        tl2::antispam::PatternNotFound object;
        explicit tl2_antispam_PatternNotFound_tl_object(tl2::antispam::PatternNotFound o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("antispam.patternNotFound", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_antispam_PatternNotFound_tl_object>(tl2::antispam::PatternNotFound{});
	});

	struct tl2_BenchObject_tl_object : public tl2::meta::tl_object {
        tl2::BenchObject object;
        explicit tl2_BenchObject_tl_object(tl2::BenchObject o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("benchObject", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_BenchObject_tl_object>(tl2::BenchObject{});
	});

	struct tl2_BoolStat_tl_object : public tl2::meta::tl_object {
        tl2::BoolStat object;
        explicit tl2_BoolStat_tl_object(tl2::BoolStat o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("boolStat", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_BoolStat_tl_object>(tl2::BoolStat{});
	});

	struct tl2_BoxedArray_tl_function : public tl2::meta::tl_function {
        tl2::BoxedArray object;
        explicit tl2_BoxedArray_tl_function(tl2::BoxedArray o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			tl2::MyBoxedArray result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("boxedArray", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_BoxedArray_tl_function>(tl2::BoxedArray{});
	});

	tl2::meta::set_create_function_by_name("boxedArray", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_BoxedArray_tl_function>(tl2::BoxedArray{});
	});

	struct tl2_BoxedInt_tl_function : public tl2::meta::tl_function {
        tl2::BoxedInt object;
        explicit tl2_BoxedInt_tl_function(tl2::BoxedInt o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			int32_t result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("boxedInt", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_BoxedInt_tl_function>(tl2::BoxedInt{});
	});

	tl2::meta::set_create_function_by_name("boxedInt", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_BoxedInt_tl_function>(tl2::BoxedInt{});
	});

	struct tl2_BoxedString_tl_function : public tl2::meta::tl_function {
        tl2::BoxedString object;
        explicit tl2_BoxedString_tl_function(tl2::BoxedString o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			std::string result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("boxedString", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_BoxedString_tl_function>(tl2::BoxedString{});
	});

	tl2::meta::set_create_function_by_name("boxedString", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_BoxedString_tl_function>(tl2::BoxedString{});
	});

	struct tl2_BoxedTuple_tl_function : public tl2::meta::tl_function {
        tl2::BoxedTuple object;
        explicit tl2_BoxedTuple_tl_function(tl2::BoxedTuple o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			std::array<int32_t, 3> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("boxedTuple", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_BoxedTuple_tl_function>(tl2::BoxedTuple{});
	});

	tl2::meta::set_create_function_by_name("boxedTuple", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_BoxedTuple_tl_function>(tl2::BoxedTuple{});
	});

	struct tl2_BoxedTupleSlice1_tl_function : public tl2::meta::tl_function {
        tl2::BoxedTupleSlice1 object;
        explicit tl2_BoxedTupleSlice1_tl_function(tl2::BoxedTupleSlice1 o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			std::vector<int32_t> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("boxedTupleSlice1", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_BoxedTupleSlice1_tl_function>(tl2::BoxedTupleSlice1{});
	});

	tl2::meta::set_create_function_by_name("boxedTupleSlice1", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_BoxedTupleSlice1_tl_function>(tl2::BoxedTupleSlice1{});
	});

	struct tl2_BoxedTupleSlice2_tl_function : public tl2::meta::tl_function {
        tl2::BoxedTupleSlice2 object;
        explicit tl2_BoxedTupleSlice2_tl_function(tl2::BoxedTupleSlice2 o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			tl2::MyBoxedTupleSlice result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("boxedTupleSlice2", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_BoxedTupleSlice2_tl_function>(tl2::BoxedTupleSlice2{});
	});

	tl2::meta::set_create_function_by_name("boxedTupleSlice2", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_BoxedTupleSlice2_tl_function>(tl2::BoxedTupleSlice2{});
	});

	struct tl2_BoxedTupleSlice3_tl_function : public tl2::meta::tl_function {
        tl2::BoxedTupleSlice3 object;
        explicit tl2_BoxedTupleSlice3_tl_function(tl2::BoxedTupleSlice3 o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			std::vector<int32_t> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("boxedTupleSlice3", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_BoxedTupleSlice3_tl_function>(tl2::BoxedTupleSlice3{});
	});

	tl2::meta::set_create_function_by_name("boxedTupleSlice3", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_BoxedTupleSlice3_tl_function>(tl2::BoxedTupleSlice3{});
	});

	struct tl2_BoxedVector32_tl_function : public tl2::meta::tl_function {
        tl2::BoxedVector32 object;
        explicit tl2_BoxedVector32_tl_function(tl2::BoxedVector32 o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			std::vector<int32_t> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("boxedVector32", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_BoxedVector32_tl_function>(tl2::BoxedVector32{});
	});

	tl2::meta::set_create_function_by_name("boxedVector32", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_BoxedVector32_tl_function>(tl2::BoxedVector32{});
	});

	struct tl2_BoxedVector32BoxedElem_tl_function : public tl2::meta::tl_function {
        tl2::BoxedVector32BoxedElem object;
        explicit tl2_BoxedVector32BoxedElem_tl_function(tl2::BoxedVector32BoxedElem o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			std::vector<int32_t> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("boxedVector32BoxedElem", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_BoxedVector32BoxedElem_tl_function>(tl2::BoxedVector32BoxedElem{});
	});

	tl2::meta::set_create_function_by_name("boxedVector32BoxedElem", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_BoxedVector32BoxedElem_tl_function>(tl2::BoxedVector32BoxedElem{});
	});

	struct tl2_BoxedVector64_tl_function : public tl2::meta::tl_function {
        tl2::BoxedVector64 object;
        explicit tl2_BoxedVector64_tl_function(tl2::BoxedVector64 o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			std::vector<int64_t> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("boxedVector64", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_BoxedVector64_tl_function>(tl2::BoxedVector64{});
	});

	tl2::meta::set_create_function_by_name("boxedVector64", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_BoxedVector64_tl_function>(tl2::BoxedVector64{});
	});

	struct tl2_FieldConflict1_tl_object : public tl2::meta::tl_object {
        tl2::FieldConflict1 object;
        explicit tl2_FieldConflict1_tl_object(tl2::FieldConflict1 o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("fieldConflict1", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_FieldConflict1_tl_object>(tl2::FieldConflict1{});
	});

	struct tl2_FieldConflict2_tl_object : public tl2::meta::tl_object {
        tl2::FieldConflict2 object;
        explicit tl2_FieldConflict2_tl_object(tl2::FieldConflict2 o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("fieldConflict2", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_FieldConflict2_tl_object>(tl2::FieldConflict2{});
	});

	struct tl2_FieldConflict3_tl_object : public tl2::meta::tl_object {
        tl2::FieldConflict3 object;
        explicit tl2_FieldConflict3_tl_object(tl2::FieldConflict3 o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("fieldConflict3", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_FieldConflict3_tl_object>(tl2::FieldConflict3{});
	});

	struct tl2_FieldConflict4_tl_object : public tl2::meta::tl_object {
        tl2::FieldConflict4 object;
        explicit tl2_FieldConflict4_tl_object(tl2::FieldConflict4 o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("fieldConflict4", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_FieldConflict4_tl_object>(tl2::FieldConflict4{});
	});

	struct tl2_Get_arrays_tl_function : public tl2::meta::tl_function {
        tl2::Get_arrays object;
        explicit tl2_Get_arrays_tl_function(tl2::Get_arrays o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			std::array<int32_t, 5> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("get_arrays", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_Get_arrays_tl_function>(tl2::Get_arrays{});
	});

	tl2::meta::set_create_function_by_name("get_arrays", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_Get_arrays_tl_function>(tl2::Get_arrays{});
	});

	struct tl2_GetDouble_tl_function : public tl2::meta::tl_function {
        tl2::GetDouble object;
        explicit tl2_GetDouble_tl_function(tl2::GetDouble o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			double result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("getDouble", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_GetDouble_tl_function>(tl2::GetDouble{});
	});

	tl2::meta::set_create_function_by_name("getDouble", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_GetDouble_tl_function>(tl2::GetDouble{});
	});

	struct tl2_GetFloat_tl_function : public tl2::meta::tl_function {
        tl2::GetFloat object;
        explicit tl2_GetFloat_tl_function(tl2::GetFloat o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			float result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("getFloat", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_GetFloat_tl_function>(tl2::GetFloat{});
	});

	tl2::meta::set_create_function_by_name("getFloat", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_GetFloat_tl_function>(tl2::GetFloat{});
	});

	struct tl2_GetMaybeIface_tl_function : public tl2::meta::tl_function {
        tl2::GetMaybeIface object;
        explicit tl2_GetMaybeIface_tl_function(tl2::GetMaybeIface o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			std::optional<::tl2::service1::Value> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("getMaybeIface", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_GetMaybeIface_tl_function>(tl2::GetMaybeIface{});
	});

	tl2::meta::set_create_function_by_name("getMaybeIface", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_GetMaybeIface_tl_function>(tl2::GetMaybeIface{});
	});

	struct tl2_GetMyDictOfInt_tl_function : public tl2::meta::tl_function {
        tl2::GetMyDictOfInt object;
        explicit tl2_GetMyDictOfInt_tl_function(tl2::GetMyDictOfInt o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			tl2::MyDictOfInt result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("getMyDictOfInt", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_GetMyDictOfInt_tl_function>(tl2::GetMyDictOfInt{});
	});

	tl2::meta::set_create_function_by_name("getMyDictOfInt", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_GetMyDictOfInt_tl_function>(tl2::GetMyDictOfInt{});
	});

	struct tl2_GetMyDouble_tl_function : public tl2::meta::tl_function {
        tl2::GetMyDouble object;
        explicit tl2_GetMyDouble_tl_function(tl2::GetMyDouble o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			tl2::MyDouble result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("getMyDouble", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_GetMyDouble_tl_function>(tl2::GetMyDouble{});
	});

	tl2::meta::set_create_function_by_name("getMyDouble", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_GetMyDouble_tl_function>(tl2::GetMyDouble{});
	});

	struct tl2_GetMyValue_tl_function : public tl2::meta::tl_function {
        tl2::GetMyValue object;
        explicit tl2_GetMyValue_tl_function(tl2::GetMyValue o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			tl2::MyValue result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("getMyValue", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_GetMyValue_tl_function>(tl2::GetMyValue{});
	});

	tl2::meta::set_create_function_by_name("getMyValue", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_GetMyValue_tl_function>(tl2::GetMyValue{});
	});

	struct tl2_GetNonOptNat_tl_function : public tl2::meta::tl_function {
        tl2::GetNonOptNat object;
        explicit tl2_GetNonOptNat_tl_function(tl2::GetNonOptNat o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			std::vector<int32_t> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("getNonOptNat", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_GetNonOptNat_tl_function>(tl2::GetNonOptNat{});
	});

	tl2::meta::set_create_function_by_name("getNonOptNat", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_GetNonOptNat_tl_function>(tl2::GetNonOptNat{});
	});

	struct tl2_GetStats_tl_function : public tl2::meta::tl_function {
        tl2::GetStats object;
        explicit tl2_GetStats_tl_function(tl2::GetStats o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			tl2::tasks::QueueTypeStats result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("getStats", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_GetStats_tl_function>(tl2::GetStats{});
	});

	tl2::meta::set_create_function_by_name("getStats", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_GetStats_tl_function>(tl2::GetStats{});
	});

	struct tl2_Integer_tl_object : public tl2::meta::tl_object {
        tl2::Integer object;
        explicit tl2_Integer_tl_object(tl2::Integer o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("integer", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_Integer_tl_object>(tl2::Integer{});
	});

	struct tl2_Issue3498_tl_object : public tl2::meta::tl_object {
        tl2::Issue3498 object;
        explicit tl2_Issue3498_tl_object(tl2::Issue3498 o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("issue3498", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_Issue3498_tl_object>(tl2::Issue3498{});
	});

	struct tl2_MyBoxedArray_tl_object : public tl2::meta::tl_object {
        tl2::MyBoxedArray object;
        explicit tl2_MyBoxedArray_tl_object(tl2::MyBoxedArray o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("myBoxedArray", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_MyBoxedArray_tl_object>(tl2::MyBoxedArray{});
	});

	struct tl2_MyBoxedTupleSlice_tl_object : public tl2::meta::tl_object {
        tl2::MyBoxedTupleSlice object;
        explicit tl2_MyBoxedTupleSlice_tl_object(tl2::MyBoxedTupleSlice o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("myBoxedTupleSlice", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_MyBoxedTupleSlice_tl_object>(tl2::MyBoxedTupleSlice{});
	});

	struct tl2_MyBoxedVectorSlice_tl_object : public tl2::meta::tl_object {
        tl2::MyBoxedVectorSlice object;
        explicit tl2_MyBoxedVectorSlice_tl_object(tl2::MyBoxedVectorSlice o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("myBoxedVectorSlice", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_MyBoxedVectorSlice_tl_object>(tl2::MyBoxedVectorSlice{});
	});

	struct tl2_MyInt_tl_object : public tl2::meta::tl_object {
        tl2::MyInt object;
        explicit tl2_MyInt_tl_object(tl2::MyInt o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("myInt", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_MyInt_tl_object>(tl2::MyInt{});
	});

	struct tl2_MyMcValue_tl_object : public tl2::meta::tl_object {
        tl2::MyMcValue object;
        explicit tl2_MyMcValue_tl_object(tl2::MyMcValue o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("myMcValue", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_MyMcValue_tl_object>(tl2::MyMcValue{});
	});

	struct tl2_MyMcValueTuple_tl_object : public tl2::meta::tl_object {
        tl2::MyMcValueTuple object;
        explicit tl2_MyMcValueTuple_tl_object(tl2::MyMcValueTuple o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("myMcValueTuple", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_MyMcValueTuple_tl_object>(tl2::MyMcValueTuple{});
	});

	struct tl2_MyMcValueVector_tl_object : public tl2::meta::tl_object {
        tl2::MyMcValueVector object;
        explicit tl2_MyMcValueVector_tl_object(tl2::MyMcValueVector o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("myMcValueVector", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_MyMcValueVector_tl_object>(tl2::MyMcValueVector{});
	});

	struct tl2_MyString_tl_object : public tl2::meta::tl_object {
        tl2::MyString object;
        explicit tl2_MyString_tl_object(tl2::MyString o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("myString", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_MyString_tl_object>(tl2::MyString{});
	});

	struct tl2_MyTwoDicts_tl_object : public tl2::meta::tl_object {
        tl2::MyTwoDicts object;
        explicit tl2_MyTwoDicts_tl_object(tl2::MyTwoDicts o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("myTwoDicts", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_MyTwoDicts_tl_object>(tl2::MyTwoDicts{});
	});

	struct tl2_NonOptNat_tl_object : public tl2::meta::tl_object {
        tl2::NonOptNat object;
        explicit tl2_NonOptNat_tl_object(tl2::NonOptNat o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("nonOptNat", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_NonOptNat_tl_object>(tl2::NonOptNat{});
	});

	struct tl2_pkg2_Foo_tl_object : public tl2::meta::tl_object {
        tl2::pkg2::Foo object;
        explicit tl2_pkg2_Foo_tl_object(tl2::pkg2::Foo o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("pkg2.foo", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_pkg2_Foo_tl_object>(tl2::pkg2::Foo{});
	});

	struct tl2_pkg2_T1_tl_object : public tl2::meta::tl_object {
        tl2::pkg2::T1 object;
        explicit tl2_pkg2_T1_tl_object(tl2::pkg2::T1 o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("pkg2.t1", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_pkg2_T1_tl_object>(tl2::pkg2::T1{});
	});

	struct tl2_RpcInvokeReqExtra_tl_object : public tl2::meta::tl_object {
        tl2::RpcInvokeReqExtra object;
        explicit tl2_RpcInvokeReqExtra_tl_object(tl2::RpcInvokeReqExtra o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("rpcInvokeReqExtra", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_RpcInvokeReqExtra_tl_object>(tl2::RpcInvokeReqExtra{});
	});

	struct tl2_service1_Add_tl_function : public tl2::meta::tl_function {
        tl2::service1::Add object;
        explicit tl2_service1_Add_tl_function(tl2::service1::Add o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.add", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Add_tl_function>(tl2::service1::Add{});
	});

	tl2::meta::set_create_function_by_name("service1.add", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_Add_tl_function>(tl2::service1::Add{});
	});

	struct tl2_service1_AddOrGet_tl_function : public tl2::meta::tl_function {
        tl2::service1::AddOrGet object;
        explicit tl2_service1_AddOrGet_tl_function(tl2::service1::AddOrGet o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			tl2::service1::Value result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.addOrGet", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_AddOrGet_tl_function>(tl2::service1::AddOrGet{});
	});

	tl2::meta::set_create_function_by_name("service1.addOrGet", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_AddOrGet_tl_function>(tl2::service1::AddOrGet{});
	});

	struct tl2_service1_AddOrIncr_tl_function : public tl2::meta::tl_function {
        tl2::service1::AddOrIncr object;
        explicit tl2_service1_AddOrIncr_tl_function(tl2::service1::AddOrIncr o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			tl2::service1::Value result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.addOrIncr", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_AddOrIncr_tl_function>(tl2::service1::AddOrIncr{});
	});

	tl2::meta::set_create_function_by_name("service1.addOrIncr", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_AddOrIncr_tl_function>(tl2::service1::AddOrIncr{});
	});

	struct tl2_service1_Append_tl_function : public tl2::meta::tl_function {
        tl2::service1::Append object;
        explicit tl2_service1_Append_tl_function(tl2::service1::Append o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.append", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Append_tl_function>(tl2::service1::Append{});
	});

	tl2::meta::set_create_function_by_name("service1.append", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_Append_tl_function>(tl2::service1::Append{});
	});

	struct tl2_service1_Cas_tl_function : public tl2::meta::tl_function {
        tl2::service1::Cas object;
        explicit tl2_service1_Cas_tl_function(tl2::service1::Cas o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.cas", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Cas_tl_function>(tl2::service1::Cas{});
	});

	tl2::meta::set_create_function_by_name("service1.cas", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_Cas_tl_function>(tl2::service1::Cas{});
	});

	struct tl2_service1_Decr_tl_function : public tl2::meta::tl_function {
        tl2::service1::Decr object;
        explicit tl2_service1_Decr_tl_function(tl2::service1::Decr o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			tl2::service1::Value result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.decr", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Decr_tl_function>(tl2::service1::Decr{});
	});

	tl2::meta::set_create_function_by_name("service1.decr", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_Decr_tl_function>(tl2::service1::Decr{});
	});

	struct tl2_service1_Delete_tl_function : public tl2::meta::tl_function {
        tl2::service1::Delete object;
        explicit tl2_service1_Delete_tl_function(tl2::service1::Delete o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.delete", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Delete_tl_function>(tl2::service1::Delete{});
	});

	tl2::meta::set_create_function_by_name("service1.delete", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_Delete_tl_function>(tl2::service1::Delete{});
	});

	struct tl2_service1_DisableExpiration_tl_function : public tl2::meta::tl_function {
        tl2::service1::DisableExpiration object;
        explicit tl2_service1_DisableExpiration_tl_function(tl2::service1::DisableExpiration o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.disableExpiration", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_DisableExpiration_tl_function>(tl2::service1::DisableExpiration{});
	});

	tl2::meta::set_create_function_by_name("service1.disableExpiration", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_DisableExpiration_tl_function>(tl2::service1::DisableExpiration{});
	});

	struct tl2_service1_DisableKeysStat_tl_function : public tl2::meta::tl_function {
        tl2::service1::DisableKeysStat object;
        explicit tl2_service1_DisableKeysStat_tl_function(tl2::service1::DisableKeysStat o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.disableKeysStat", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_DisableKeysStat_tl_function>(tl2::service1::DisableKeysStat{});
	});

	tl2::meta::set_create_function_by_name("service1.disableKeysStat", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_DisableKeysStat_tl_function>(tl2::service1::DisableKeysStat{});
	});

	struct tl2_service1_EnableExpiration_tl_function : public tl2::meta::tl_function {
        tl2::service1::EnableExpiration object;
        explicit tl2_service1_EnableExpiration_tl_function(tl2::service1::EnableExpiration o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.enableExpiration", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_EnableExpiration_tl_function>(tl2::service1::EnableExpiration{});
	});

	tl2::meta::set_create_function_by_name("service1.enableExpiration", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_EnableExpiration_tl_function>(tl2::service1::EnableExpiration{});
	});

	struct tl2_service1_EnableKeysStat_tl_function : public tl2::meta::tl_function {
        tl2::service1::EnableKeysStat object;
        explicit tl2_service1_EnableKeysStat_tl_function(tl2::service1::EnableKeysStat o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.enableKeysStat", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_EnableKeysStat_tl_function>(tl2::service1::EnableKeysStat{});
	});

	tl2::meta::set_create_function_by_name("service1.enableKeysStat", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_EnableKeysStat_tl_function>(tl2::service1::EnableKeysStat{});
	});

	struct tl2_service1_Exists_tl_function : public tl2::meta::tl_function {
        tl2::service1::Exists object;
        explicit tl2_service1_Exists_tl_function(tl2::service1::Exists o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.exists", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Exists_tl_function>(tl2::service1::Exists{});
	});

	tl2::meta::set_create_function_by_name("service1.exists", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_Exists_tl_function>(tl2::service1::Exists{});
	});

	struct tl2_service1_Get_tl_function : public tl2::meta::tl_function {
        tl2::service1::Get object;
        explicit tl2_service1_Get_tl_function(tl2::service1::Get o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			tl2::service1::Value result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.get", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Get_tl_function>(tl2::service1::Get{});
	});

	tl2::meta::set_create_function_by_name("service1.get", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_Get_tl_function>(tl2::service1::Get{});
	});

	struct tl2_service1_GetExpireTime_tl_function : public tl2::meta::tl_function {
        tl2::service1::GetExpireTime object;
        explicit tl2_service1_GetExpireTime_tl_function(tl2::service1::GetExpireTime o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			std::optional<int32_t> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.getExpireTime", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_GetExpireTime_tl_function>(tl2::service1::GetExpireTime{});
	});

	tl2::meta::set_create_function_by_name("service1.getExpireTime", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_GetExpireTime_tl_function>(tl2::service1::GetExpireTime{});
	});

	struct tl2_service1_GetKeysStat_tl_function : public tl2::meta::tl_function {
        tl2::service1::GetKeysStat object;
        explicit tl2_service1_GetKeysStat_tl_function(tl2::service1::GetKeysStat o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			std::optional<::tl2::service1::KeysStat> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.getKeysStat", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_GetKeysStat_tl_function>(tl2::service1::GetKeysStat{});
	});

	tl2::meta::set_create_function_by_name("service1.getKeysStat", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_GetKeysStat_tl_function>(tl2::service1::GetKeysStat{});
	});

	struct tl2_service1_GetKeysStatPeriods_tl_function : public tl2::meta::tl_function {
        tl2::service1::GetKeysStatPeriods object;
        explicit tl2_service1_GetKeysStatPeriods_tl_function(tl2::service1::GetKeysStatPeriods o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			std::vector<int32_t> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.getKeysStatPeriods", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_GetKeysStatPeriods_tl_function>(tl2::service1::GetKeysStatPeriods{});
	});

	tl2::meta::set_create_function_by_name("service1.getKeysStatPeriods", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_GetKeysStatPeriods_tl_function>(tl2::service1::GetKeysStatPeriods{});
	});

	struct tl2_service1_GetWildcard_tl_function : public tl2::meta::tl_function {
        tl2::service1::GetWildcard object;
        explicit tl2_service1_GetWildcard_tl_function(tl2::service1::GetWildcard o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			std::vector<::tl2::Map<std::string, std::string>> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.getWildcard", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_GetWildcard_tl_function>(tl2::service1::GetWildcard{});
	});

	tl2::meta::set_create_function_by_name("service1.getWildcard", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_GetWildcard_tl_function>(tl2::service1::GetWildcard{});
	});

	struct tl2_service1_GetWildcardDict_tl_function : public tl2::meta::tl_function {
        tl2::service1::GetWildcardDict object;
        explicit tl2_service1_GetWildcardDict_tl_function(tl2::service1::GetWildcardDict o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			tl2::Dictionary<std::string> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.getWildcardDict", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_GetWildcardDict_tl_function>(tl2::service1::GetWildcardDict{});
	});

	tl2::meta::set_create_function_by_name("service1.getWildcardDict", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_GetWildcardDict_tl_function>(tl2::service1::GetWildcardDict{});
	});

	struct tl2_service1_GetWildcardList_tl_function : public tl2::meta::tl_function {
        tl2::service1::GetWildcardList object;
        explicit tl2_service1_GetWildcardList_tl_function(tl2::service1::GetWildcardList o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			std::vector<std::string> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.getWildcardList", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_GetWildcardList_tl_function>(tl2::service1::GetWildcardList{});
	});

	tl2::meta::set_create_function_by_name("service1.getWildcardList", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_GetWildcardList_tl_function>(tl2::service1::GetWildcardList{});
	});

	struct tl2_service1_GetWildcardWithFlags_tl_function : public tl2::meta::tl_function {
        tl2::service1::GetWildcardWithFlags object;
        explicit tl2_service1_GetWildcardWithFlags_tl_function(tl2::service1::GetWildcardWithFlags o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			tl2::Dictionary<::tl2::service1::Value> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.getWildcardWithFlags", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_GetWildcardWithFlags_tl_function>(tl2::service1::GetWildcardWithFlags{});
	});

	tl2::meta::set_create_function_by_name("service1.getWildcardWithFlags", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_GetWildcardWithFlags_tl_function>(tl2::service1::GetWildcardWithFlags{});
	});

	struct tl2_service1_Incr_tl_function : public tl2::meta::tl_function {
        tl2::service1::Incr object;
        explicit tl2_service1_Incr_tl_function(tl2::service1::Incr o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			tl2::service1::Value result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.incr", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Incr_tl_function>(tl2::service1::Incr{});
	});

	tl2::meta::set_create_function_by_name("service1.incr", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_Incr_tl_function>(tl2::service1::Incr{});
	});

	struct tl2_service1_KeysStat_tl_object : public tl2::meta::tl_object {
        tl2::service1::KeysStat object;
        explicit tl2_service1_KeysStat_tl_object(tl2::service1::KeysStat o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("service1.keysStat", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_KeysStat_tl_object>(tl2::service1::KeysStat{});
	});

	struct tl2_service1_Longvalue_tl_object : public tl2::meta::tl_object {
        tl2::service1::Longvalue object;
        explicit tl2_service1_Longvalue_tl_object(tl2::service1::Longvalue o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("service1.longvalue", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Longvalue_tl_object>(tl2::service1::Longvalue{});
	});

	struct tl2_service1_LongvalueWithTime_tl_object : public tl2::meta::tl_object {
        tl2::service1::LongvalueWithTime object;
        explicit tl2_service1_LongvalueWithTime_tl_object(tl2::service1::LongvalueWithTime o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("service1.longvalueWithTime", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_LongvalueWithTime_tl_object>(tl2::service1::LongvalueWithTime{});
	});

	struct tl2_service1_Not_found_tl_object : public tl2::meta::tl_object {
        tl2::service1::Not_found object;
        explicit tl2_service1_Not_found_tl_object(tl2::service1::Not_found o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("service1.not_found", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Not_found_tl_object>(tl2::service1::Not_found{});
	});

	struct tl2_service1_Replace_tl_function : public tl2::meta::tl_function {
        tl2::service1::Replace object;
        explicit tl2_service1_Replace_tl_function(tl2::service1::Replace o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.replace", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Replace_tl_function>(tl2::service1::Replace{});
	});

	tl2::meta::set_create_function_by_name("service1.replace", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_Replace_tl_function>(tl2::service1::Replace{});
	});

	struct tl2_service1_ReplaceOrIncr_tl_function : public tl2::meta::tl_function {
        tl2::service1::ReplaceOrIncr object;
        explicit tl2_service1_ReplaceOrIncr_tl_function(tl2::service1::ReplaceOrIncr o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			tl2::service1::Value result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.replaceOrIncr", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_ReplaceOrIncr_tl_function>(tl2::service1::ReplaceOrIncr{});
	});

	tl2::meta::set_create_function_by_name("service1.replaceOrIncr", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_ReplaceOrIncr_tl_function>(tl2::service1::ReplaceOrIncr{});
	});

	struct tl2_service1_Set_tl_function : public tl2::meta::tl_function {
        tl2::service1::Set object;
        explicit tl2_service1_Set_tl_function(tl2::service1::Set o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.set", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Set_tl_function>(tl2::service1::Set{});
	});

	tl2::meta::set_create_function_by_name("service1.set", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_Set_tl_function>(tl2::service1::Set{});
	});

	struct tl2_service1_SetOrIncr_tl_function : public tl2::meta::tl_function {
        tl2::service1::SetOrIncr object;
        explicit tl2_service1_SetOrIncr_tl_function(tl2::service1::SetOrIncr o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			tl2::service1::Value result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.setOrIncr", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_SetOrIncr_tl_function>(tl2::service1::SetOrIncr{});
	});

	tl2::meta::set_create_function_by_name("service1.setOrIncr", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_SetOrIncr_tl_function>(tl2::service1::SetOrIncr{});
	});

	struct tl2_service1_Strvalue_tl_object : public tl2::meta::tl_object {
        tl2::service1::Strvalue object;
        explicit tl2_service1_Strvalue_tl_object(tl2::service1::Strvalue o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("service1.strvalue", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Strvalue_tl_object>(tl2::service1::Strvalue{});
	});

	struct tl2_service1_StrvalueWithTime_tl_object : public tl2::meta::tl_object {
        tl2::service1::StrvalueWithTime object;
        explicit tl2_service1_StrvalueWithTime_tl_object(tl2::service1::StrvalueWithTime o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("service1.strvalueWithTime", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_StrvalueWithTime_tl_object>(tl2::service1::StrvalueWithTime{});
	});

	struct tl2_service1_Touch_tl_function : public tl2::meta::tl_function {
        tl2::service1::Touch object;
        explicit tl2_service1_Touch_tl_function(tl2::service1::Touch o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.touch", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Touch_tl_function>(tl2::service1::Touch{});
	});

	tl2::meta::set_create_function_by_name("service1.touch", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_Touch_tl_function>(tl2::service1::Touch{});
	});

	struct tl2_service2_AddOrIncrMany_tl_function : public tl2::meta::tl_function {
        tl2::service2::AddOrIncrMany object;
        explicit tl2_service2_AddOrIncrMany_tl_function(tl2::service2::AddOrIncrMany o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			std::vector<::tl2::service2::CounterSet> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service2.addOrIncrMany", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service2_AddOrIncrMany_tl_function>(tl2::service2::AddOrIncrMany{});
	});

	tl2::meta::set_create_function_by_name("service2.addOrIncrMany", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service2_AddOrIncrMany_tl_function>(tl2::service2::AddOrIncrMany{});
	});

	struct tl2_service2_Set_tl_function : public tl2::meta::tl_function {
        tl2::service2::Set object;
        explicit tl2_service2_Set_tl_function(tl2::service2::Set o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			tl2::True result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service2.set", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service2_Set_tl_function>(tl2::service2::Set{});
	});

	tl2::meta::set_create_function_by_name("service2.set", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service2_Set_tl_function>(tl2::service2::Set{});
	});

	struct tl2_service2_SetObjectTtl_tl_function : public tl2::meta::tl_function {
        tl2::service2::SetObjectTtl object;
        explicit tl2_service2_SetObjectTtl_tl_function(tl2::service2::SetObjectTtl o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			tl2::True result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service2.setObjectTtl", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service2_SetObjectTtl_tl_function>(tl2::service2::SetObjectTtl{});
	});

	tl2::meta::set_create_function_by_name("service2.setObjectTtl", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service2_SetObjectTtl_tl_function>(tl2::service2::SetObjectTtl{});
	});

	struct tl2_service3_CreateProduct_tl_function : public tl2::meta::tl_function {
        tl2::service3::CreateProduct object;
        explicit tl2_service3_CreateProduct_tl_function(tl2::service3::CreateProduct o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.createProduct", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_CreateProduct_tl_function>(tl2::service3::CreateProduct{});
	});

	tl2::meta::set_create_function_by_name("service3.createProduct", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_CreateProduct_tl_function>(tl2::service3::CreateProduct{});
	});

	struct tl2_service3_DeleteAllProducts_tl_function : public tl2::meta::tl_function {
        tl2::service3::DeleteAllProducts object;
        explicit tl2_service3_DeleteAllProducts_tl_function(tl2::service3::DeleteAllProducts o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.deleteAllProducts", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_DeleteAllProducts_tl_function>(tl2::service3::DeleteAllProducts{});
	});

	tl2::meta::set_create_function_by_name("service3.deleteAllProducts", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_DeleteAllProducts_tl_function>(tl2::service3::DeleteAllProducts{});
	});

	struct tl2_service3_DeleteGroupedProducts_tl_function : public tl2::meta::tl_function {
        tl2::service3::DeleteGroupedProducts object;
        explicit tl2_service3_DeleteGroupedProducts_tl_function(tl2::service3::DeleteGroupedProducts o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.deleteGroupedProducts", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_DeleteGroupedProducts_tl_function>(tl2::service3::DeleteGroupedProducts{});
	});

	tl2::meta::set_create_function_by_name("service3.deleteGroupedProducts", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_DeleteGroupedProducts_tl_function>(tl2::service3::DeleteGroupedProducts{});
	});

	struct tl2_service3_DeleteProduct_tl_function : public tl2::meta::tl_function {
        tl2::service3::DeleteProduct object;
        explicit tl2_service3_DeleteProduct_tl_function(tl2::service3::DeleteProduct o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.deleteProduct", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_DeleteProduct_tl_function>(tl2::service3::DeleteProduct{});
	});

	tl2::meta::set_create_function_by_name("service3.deleteProduct", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_DeleteProduct_tl_function>(tl2::service3::DeleteProduct{});
	});

	struct tl2_service3_GetLastVisitTimestamp_tl_function : public tl2::meta::tl_function {
        tl2::service3::GetLastVisitTimestamp object;
        explicit tl2_service3_GetLastVisitTimestamp_tl_function(tl2::service3::GetLastVisitTimestamp o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			std::optional<int32_t> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.getLastVisitTimestamp", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_GetLastVisitTimestamp_tl_function>(tl2::service3::GetLastVisitTimestamp{});
	});

	tl2::meta::set_create_function_by_name("service3.getLastVisitTimestamp", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_GetLastVisitTimestamp_tl_function>(tl2::service3::GetLastVisitTimestamp{});
	});

	struct tl2_service3_GetLimits_tl_function : public tl2::meta::tl_function {
        tl2::service3::GetLimits object;
        explicit tl2_service3_GetLimits_tl_function(tl2::service3::GetLimits o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			tl2::service3::Limits result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.getLimits", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_GetLimits_tl_function>(tl2::service3::GetLimits{});
	});

	tl2::meta::set_create_function_by_name("service3.getLimits", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_GetLimits_tl_function>(tl2::service3::GetLimits{});
	});

	struct tl2_service3_GetProductStats_tl_function : public tl2::meta::tl_function {
        tl2::service3::GetProductStats object;
        explicit tl2_service3_GetProductStats_tl_function(tl2::service3::GetProductStats o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			std::optional<std::vector<::tl2::service3::ProductStatsOld>> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.getProductStats", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_GetProductStats_tl_function>(tl2::service3::GetProductStats{});
	});

	tl2::meta::set_create_function_by_name("service3.getProductStats", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_GetProductStats_tl_function>(tl2::service3::GetProductStats{});
	});

	struct tl2_service3_GetProducts_tl_function : public tl2::meta::tl_function {
        tl2::service3::GetProducts object;
        explicit tl2_service3_GetProducts_tl_function(tl2::service3::GetProducts o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			std::optional<std::vector<::tl2::service3::Product>> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.getProducts", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_GetProducts_tl_function>(tl2::service3::GetProducts{});
	});

	tl2::meta::set_create_function_by_name("service3.getProducts", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_GetProducts_tl_function>(tl2::service3::GetProducts{});
	});

	struct tl2_service3_GetScheduledProducts_tl_function : public tl2::meta::tl_function {
        tl2::service3::GetScheduledProducts object;
        explicit tl2_service3_GetScheduledProducts_tl_function(tl2::service3::GetScheduledProducts o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			std::optional<std::vector<::tl2::service3::Productmode<0>>> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.getScheduledProducts", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_GetScheduledProducts_tl_function>(tl2::service3::GetScheduledProducts{});
	});

	tl2::meta::set_create_function_by_name("service3.getScheduledProducts", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_GetScheduledProducts_tl_function>(tl2::service3::GetScheduledProducts{});
	});

	struct tl2_service3_GroupCountLimit_tl_object : public tl2::meta::tl_object {
        tl2::service3::GroupCountLimit object;
        explicit tl2_service3_GroupCountLimit_tl_object(tl2::service3::GroupCountLimit o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("service3.groupCountLimit", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_GroupCountLimit_tl_object>(tl2::service3::GroupCountLimit{});
	});

	struct tl2_service3_GroupSizeLimit_tl_object : public tl2::meta::tl_object {
        tl2::service3::GroupSizeLimit object;
        explicit tl2_service3_GroupSizeLimit_tl_object(tl2::service3::GroupSizeLimit o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("service3.groupSizeLimit", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_GroupSizeLimit_tl_object>(tl2::service3::GroupSizeLimit{});
	});

	struct tl2_service3_Limits_tl_object : public tl2::meta::tl_object {
        tl2::service3::Limits object;
        explicit tl2_service3_Limits_tl_object(tl2::service3::Limits o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("service3.limits", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_Limits_tl_object>(tl2::service3::Limits{});
	});

	struct tl2_service3_ProductStatsOld_tl_object : public tl2::meta::tl_object {
        tl2::service3::ProductStatsOld object;
        explicit tl2_service3_ProductStatsOld_tl_object(tl2::service3::ProductStatsOld o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("service3.productStatsOld", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_ProductStatsOld_tl_object>(tl2::service3::ProductStatsOld{});
	});

	struct tl2_service3_RestoreAllProducts_tl_function : public tl2::meta::tl_function {
        tl2::service3::RestoreAllProducts object;
        explicit tl2_service3_RestoreAllProducts_tl_function(tl2::service3::RestoreAllProducts o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.restoreAllProducts", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_RestoreAllProducts_tl_function>(tl2::service3::RestoreAllProducts{});
	});

	tl2::meta::set_create_function_by_name("service3.restoreAllProducts", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_RestoreAllProducts_tl_function>(tl2::service3::RestoreAllProducts{});
	});

	struct tl2_service3_RestoreGroupedProducts_tl_function : public tl2::meta::tl_function {
        tl2::service3::RestoreGroupedProducts object;
        explicit tl2_service3_RestoreGroupedProducts_tl_function(tl2::service3::RestoreGroupedProducts o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.restoreGroupedProducts", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_RestoreGroupedProducts_tl_function>(tl2::service3::RestoreGroupedProducts{});
	});

	tl2::meta::set_create_function_by_name("service3.restoreGroupedProducts", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_RestoreGroupedProducts_tl_function>(tl2::service3::RestoreGroupedProducts{});
	});

	struct tl2_service3_RestoreProduct_tl_function : public tl2::meta::tl_function {
        tl2::service3::RestoreProduct object;
        explicit tl2_service3_RestoreProduct_tl_function(tl2::service3::RestoreProduct o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.restoreProduct", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_RestoreProduct_tl_function>(tl2::service3::RestoreProduct{});
	});

	tl2::meta::set_create_function_by_name("service3.restoreProduct", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_RestoreProduct_tl_function>(tl2::service3::RestoreProduct{});
	});

	struct tl2_service3_SetLastVisitTimestamp_tl_function : public tl2::meta::tl_function {
        tl2::service3::SetLastVisitTimestamp object;
        explicit tl2_service3_SetLastVisitTimestamp_tl_function(tl2::service3::SetLastVisitTimestamp o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.setLastVisitTimestamp", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_SetLastVisitTimestamp_tl_function>(tl2::service3::SetLastVisitTimestamp{});
	});

	tl2::meta::set_create_function_by_name("service3.setLastVisitTimestamp", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_SetLastVisitTimestamp_tl_function>(tl2::service3::SetLastVisitTimestamp{});
	});

	struct tl2_service3_SetLimits_tl_function : public tl2::meta::tl_function {
        tl2::service3::SetLimits object;
        explicit tl2_service3_SetLimits_tl_function(tl2::service3::SetLimits o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			tl2::BoolStat result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.setLimits", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_SetLimits_tl_function>(tl2::service3::SetLimits{});
	});

	tl2::meta::set_create_function_by_name("service3.setLimits", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_SetLimits_tl_function>(tl2::service3::SetLimits{});
	});

	struct tl2_service4_ModifiedNewsEntry_tl_object : public tl2::meta::tl_object {
        tl2::service4::ModifiedNewsEntry object;
        explicit tl2_service4_ModifiedNewsEntry_tl_object(tl2::service4::ModifiedNewsEntry o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("service4.modifiedNewsEntry", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service4_ModifiedNewsEntry_tl_object>(tl2::service4::ModifiedNewsEntry{});
	});

	struct tl2_service4_Object_tl_object : public tl2::meta::tl_object {
        tl2::service4::Object object;
        explicit tl2_service4_Object_tl_object(tl2::service4::Object o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("service4.object", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service4_Object_tl_object>(tl2::service4::Object{});
	});

	struct tl2_service5_EmptyOutput_tl_object : public tl2::meta::tl_object {
        tl2::service5::EmptyOutput object;
        explicit tl2_service5_EmptyOutput_tl_object(tl2::service5::EmptyOutput o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("service5.emptyOutput", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service5_EmptyOutput_tl_object>(tl2::service5::EmptyOutput{});
	});

	struct tl2_service5_Insert_tl_function : public tl2::meta::tl_function {
        tl2::service5::Insert object;
        explicit tl2_service5_Insert_tl_function(tl2::service5::Insert o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			tl2::service5::Output result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service5.insert", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service5_Insert_tl_function>(tl2::service5::Insert{});
	});

	tl2::meta::set_create_function_by_name("service5.insert", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service5_Insert_tl_function>(tl2::service5::Insert{});
	});

	struct tl2_service5_Params_tl_object : public tl2::meta::tl_object {
        tl2::service5::Params object;
        explicit tl2_service5_Params_tl_object(tl2::service5::Params o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("service5.params", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service5_Params_tl_object>(tl2::service5::Params{});
	});

	struct tl2_service5_PerformQuery_tl_function : public tl2::meta::tl_function {
        tl2::service5::PerformQuery object;
        explicit tl2_service5_PerformQuery_tl_function(tl2::service5::PerformQuery o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			tl2::service5::Output result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service5.performQuery", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service5_PerformQuery_tl_function>(tl2::service5::PerformQuery{});
	});

	tl2::meta::set_create_function_by_name("service5.performQuery", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service5_PerformQuery_tl_function>(tl2::service5::PerformQuery{});
	});

	struct tl2_service5_Query_tl_function : public tl2::meta::tl_function {
        tl2::service5::Query object;
        explicit tl2_service5_Query_tl_function(tl2::service5::Query o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			tl2::service5::Output result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service5.query", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service5_Query_tl_function>(tl2::service5::Query{});
	});

	tl2::meta::set_create_function_by_name("service5.query", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service5_Query_tl_function>(tl2::service5::Query{});
	});

	struct tl2_service5_StringOutput_tl_object : public tl2::meta::tl_object {
        tl2::service5::StringOutput object;
        explicit tl2_service5_StringOutput_tl_object(tl2::service5::StringOutput o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("service5.stringOutput", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service5_StringOutput_tl_object>(tl2::service5::StringOutput{});
	});

	struct tl2_service6_Error_tl_object : public tl2::meta::tl_object {
        tl2::service6::Error object;
        explicit tl2_service6_Error_tl_object(tl2::service6::Error o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("service6.error", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service6_Error_tl_object>(tl2::service6::Error{});
	});

	struct tl2_service6_FindResultRow_tl_object : public tl2::meta::tl_object {
        tl2::service6::FindResultRow object;
        explicit tl2_service6_FindResultRow_tl_object(tl2::service6::FindResultRow o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("service6.findResultRow", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service6_FindResultRow_tl_object>(tl2::service6::FindResultRow{});
	});

	struct tl2_service6_FindWithBoundsResult_tl_object : public tl2::meta::tl_object {
        tl2::service6::FindWithBoundsResult object;
        explicit tl2_service6_FindWithBoundsResult_tl_object(tl2::service6::FindWithBoundsResult o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("service6.findWithBoundsResult", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service6_FindWithBoundsResult_tl_object>(tl2::service6::FindWithBoundsResult{});
	});

	struct tl2_service6_MultiFind_tl_function : public tl2::meta::tl_function {
        tl2::service6::MultiFind object;
        explicit tl2_service6_MultiFind_tl_function(tl2::service6::MultiFind o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service6.multiFind", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service6_MultiFind_tl_function>(tl2::service6::MultiFind{});
	});

	tl2::meta::set_create_function_by_name("service6.multiFind", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service6_MultiFind_tl_function>(tl2::service6::MultiFind{});
	});

	struct tl2_service6_MultiFindWithBounds_tl_function : public tl2::meta::tl_function {
        tl2::service6::MultiFindWithBounds object;
        explicit tl2_service6_MultiFindWithBounds_tl_function(tl2::service6::MultiFindWithBounds o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service6.multiFindWithBounds", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service6_MultiFindWithBounds_tl_function>(tl2::service6::MultiFindWithBounds{});
	});

	tl2::meta::set_create_function_by_name("service6.multiFindWithBounds", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service6_MultiFindWithBounds_tl_function>(tl2::service6::MultiFindWithBounds{});
	});

	struct tl2_StatOne_tl_object : public tl2::meta::tl_object {
        tl2::StatOne object;
        explicit tl2_StatOne_tl_object(tl2::StatOne o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("statOne", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_StatOne_tl_object>(tl2::StatOne{});
	});

	struct tl2_tasks_AddTask_tl_function : public tl2::meta::tl_function {
        tl2::tasks::AddTask object;
        explicit tl2_tasks_AddTask_tl_function(tl2::tasks::AddTask o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			int64_t result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("tasks.addTask", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_AddTask_tl_function>(tl2::tasks::AddTask{});
	});

	tl2::meta::set_create_function_by_name("tasks.addTask", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_tasks_AddTask_tl_function>(tl2::tasks::AddTask{});
	});

	struct tl2_tasks_CronTask_tl_object : public tl2::meta::tl_object {
        tl2::tasks::CronTask object;
        explicit tl2_tasks_CronTask_tl_object(tl2::tasks::CronTask o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("tasks.cronTask", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_CronTask_tl_object>(tl2::tasks::CronTask{});
	});

	struct tl2_tasks_CronTaskWithId_tl_object : public tl2::meta::tl_object {
        tl2::tasks::CronTaskWithId object;
        explicit tl2_tasks_CronTaskWithId_tl_object(tl2::tasks::CronTaskWithId o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("tasks.cronTaskWithId", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_CronTaskWithId_tl_object>(tl2::tasks::CronTaskWithId{});
	});

	struct tl2_tasks_CronTime_tl_object : public tl2::meta::tl_object {
        tl2::tasks::CronTime object;
        explicit tl2_tasks_CronTime_tl_object(tl2::tasks::CronTime o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("tasks.cronTime", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_CronTime_tl_object>(tl2::tasks::CronTime{});
	});

	struct tl2_tasks_GetAnyTask_tl_function : public tl2::meta::tl_function {
        tl2::tasks::GetAnyTask object;
        explicit tl2_tasks_GetAnyTask_tl_function(tl2::tasks::GetAnyTask o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			std::optional<::tl2::tasks::TaskInfo> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("tasks.getAnyTask", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_GetAnyTask_tl_function>(tl2::tasks::GetAnyTask{});
	});

	tl2::meta::set_create_function_by_name("tasks.getAnyTask", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_tasks_GetAnyTask_tl_function>(tl2::tasks::GetAnyTask{});
	});

	struct tl2_tasks_GetQueueSize_tl_function : public tl2::meta::tl_function {
        tl2::tasks::GetQueueSize object;
        explicit tl2_tasks_GetQueueSize_tl_function(tl2::tasks::GetQueueSize o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			tl2::tasks::QueueStats result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("tasks.getQueueSize", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_GetQueueSize_tl_function>(tl2::tasks::GetQueueSize{});
	});

	tl2::meta::set_create_function_by_name("tasks.getQueueSize", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_tasks_GetQueueSize_tl_function>(tl2::tasks::GetQueueSize{});
	});

	struct tl2_tasks_GetQueueTypes_tl_function : public tl2::meta::tl_function {
        tl2::tasks::GetQueueTypes object;
        explicit tl2_tasks_GetQueueTypes_tl_function(tl2::tasks::GetQueueTypes o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			std::vector<::tl2::tasks::QueueTypeInfo> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("tasks.getQueueTypes", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_GetQueueTypes_tl_function>(tl2::tasks::GetQueueTypes{});
	});

	tl2::meta::set_create_function_by_name("tasks.getQueueTypes", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_tasks_GetQueueTypes_tl_function>(tl2::tasks::GetQueueTypes{});
	});

	struct tl2_tasks_GetTaskFromQueue_tl_function : public tl2::meta::tl_function {
        tl2::tasks::GetTaskFromQueue object;
        explicit tl2_tasks_GetTaskFromQueue_tl_function(tl2::tasks::GetTaskFromQueue o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			std::optional<::tl2::tasks::TaskInfo> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("tasks.getTaskFromQueue", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_GetTaskFromQueue_tl_function>(tl2::tasks::GetTaskFromQueue{});
	});

	tl2::meta::set_create_function_by_name("tasks.getTaskFromQueue", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_tasks_GetTaskFromQueue_tl_function>(tl2::tasks::GetTaskFromQueue{});
	});

	struct tl2_tasks_QueueTypeInfo_tl_object : public tl2::meta::tl_object {
        tl2::tasks::QueueTypeInfo object;
        explicit tl2_tasks_QueueTypeInfo_tl_object(tl2::tasks::QueueTypeInfo o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("tasks.queueTypeInfo", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_QueueTypeInfo_tl_object>(tl2::tasks::QueueTypeInfo{});
	});

	struct tl2_tasks_QueueTypeSettings_tl_object : public tl2::meta::tl_object {
        tl2::tasks::QueueTypeSettings object;
        explicit tl2_tasks_QueueTypeSettings_tl_object(tl2::tasks::QueueTypeSettings o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("tasks.queueTypeSettings", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_QueueTypeSettings_tl_object>(tl2::tasks::QueueTypeSettings{});
	});

	struct tl2_tasks_QueueTypeStats_tl_object : public tl2::meta::tl_object {
        tl2::tasks::QueueTypeStats object;
        explicit tl2_tasks_QueueTypeStats_tl_object(tl2::tasks::QueueTypeStats o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("tasks.queueTypeStats", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_QueueTypeStats_tl_object>(tl2::tasks::QueueTypeStats{});
	});

	struct tl2_tasks_Task_tl_object : public tl2::meta::tl_object {
        tl2::tasks::Task object;
        explicit tl2_tasks_Task_tl_object(tl2::tasks::Task o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("tasks.task", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_Task_tl_object>(tl2::tasks::Task{});
	});

	struct tl2_tasks_TaskInfo_tl_object : public tl2::meta::tl_object {
        tl2::tasks::TaskInfo object;
        explicit tl2_tasks_TaskInfo_tl_object(tl2::tasks::TaskInfo o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("tasks.taskInfo", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_TaskInfo_tl_object>(tl2::tasks::TaskInfo{});
	});

	struct tl2_tasks_TaskStatusInProgress_tl_object : public tl2::meta::tl_object {
        tl2::tasks::TaskStatusInProgress object;
        explicit tl2_tasks_TaskStatusInProgress_tl_object(tl2::tasks::TaskStatusInProgress o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("tasks.taskStatusInProgress", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_TaskStatusInProgress_tl_object>(tl2::tasks::TaskStatusInProgress{});
	});

	struct tl2_tasks_TaskStatusNotCurrentlyInEngine_tl_object : public tl2::meta::tl_object {
        tl2::tasks::TaskStatusNotCurrentlyInEngine object;
        explicit tl2_tasks_TaskStatusNotCurrentlyInEngine_tl_object(tl2::tasks::TaskStatusNotCurrentlyInEngine o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("tasks.taskStatusNotCurrentlyInEngine", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_TaskStatusNotCurrentlyInEngine_tl_object>(tl2::tasks::TaskStatusNotCurrentlyInEngine{});
	});

	struct tl2_tasks_TaskStatusScheduled_tl_object : public tl2::meta::tl_object {
        tl2::tasks::TaskStatusScheduled object;
        explicit tl2_tasks_TaskStatusScheduled_tl_object(tl2::tasks::TaskStatusScheduled o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("tasks.taskStatusScheduled", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_TaskStatusScheduled_tl_object>(tl2::tasks::TaskStatusScheduled{});
	});

	struct tl2_tasks_TaskStatusWaiting_tl_object : public tl2::meta::tl_object {
        tl2::tasks::TaskStatusWaiting object;
        explicit tl2_tasks_TaskStatusWaiting_tl_object(tl2::tasks::TaskStatusWaiting o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("tasks.taskStatusWaiting", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_TaskStatusWaiting_tl_object>(tl2::tasks::TaskStatusWaiting{});
	});

	struct tl2_tree_stats_ObjectLimitValueLong_tl_object : public tl2::meta::tl_object {
        tl2::tree_stats::ObjectLimitValueLong object;
        explicit tl2_tree_stats_ObjectLimitValueLong_tl_object(tl2::tree_stats::ObjectLimitValueLong o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("tree_stats.objectLimitValueLong", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tree_stats_ObjectLimitValueLong_tl_object>(tl2::tree_stats::ObjectLimitValueLong{});
	});

	struct tl2_True_tl_object : public tl2::meta::tl_object {
        tl2::True object;
        explicit tl2_True_tl_object(tl2::True o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("true", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_True_tl_object>(tl2::True{});
	});

	struct tl2_unique_Get_tl_function : public tl2::meta::tl_function {
        tl2::unique::Get object;
        explicit tl2_unique_Get_tl_function(tl2::unique::Get o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			std::optional<int32_t> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("unique.get", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_unique_Get_tl_function>(tl2::unique::Get{});
	});

	tl2::meta::set_create_function_by_name("unique.get", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_unique_Get_tl_function>(tl2::unique::Get{});
	});

	struct tl2_unique_StringToInt_tl_function : public tl2::meta::tl_function {
        tl2::unique::StringToInt object;
        explicit tl2_unique_StringToInt_tl_function(tl2::unique::StringToInt o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) {
			int32_t result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("unique.stringToInt", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_unique_StringToInt_tl_function>(tl2::unique::StringToInt{});
	});

	tl2::meta::set_create_function_by_name("unique.stringToInt", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_unique_StringToInt_tl_function>(tl2::unique::StringToInt{});
	});

	struct tl2_WithFloat_tl_object : public tl2::meta::tl_object {
        tl2::WithFloat object;
        explicit tl2_WithFloat_tl_object(tl2::WithFloat o) : object(std::move(o)) {}

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}

    };
	tl2::meta::set_create_object_by_name("withFloat", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_WithFloat_tl_object>(tl2::WithFloat{});
	});

}
