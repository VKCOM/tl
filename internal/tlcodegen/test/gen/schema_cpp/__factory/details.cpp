#include "../__meta/headers.hpp"
#include "headers.hpp"

#include "../__common_namespace/types/withFloat.hpp"
#include "../__common_namespace/types/vector.hpp"
#include "../unique/functions/unique.stringToInt.hpp"
#include "../unique/functions/unique.get.hpp"
#include "../__common_namespace/types/tuple.hpp"
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
#include "../__common_namespace/types/string.hpp"
#include "../__common_namespace/types/statOne.hpp"
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
#include "../__common_namespace/types/rpcInvokeReqExtra.hpp"
#include "../__common_namespace/types/true.hpp"
#include "../pkg2/types/pkg2.t1.hpp"
#include "../pkg2/types/pkg2.foo.hpp"
#include "../__common_namespace/types/nonOptNat.hpp"
#include "../__common_namespace/types/myTwoDicts.hpp"
#include "../__common_namespace/types/myMcValueVector.hpp"
#include "../__common_namespace/types/myMcValueTuple.hpp"
#include "../__common_namespace/types/myMcValue.hpp"
#include "../__common_namespace/types/myBoxedVectorSlice.hpp"
#include "../__common_namespace/types/long.hpp"
#include "../__common_namespace/types/issue3498.hpp"
#include "../service6/types/service6.findResultRow.hpp"
#include "../service6/types/service6.error.hpp"
#include "../__common_namespace/types/int.hpp"
#include "../__common_namespace/functions/getStats.hpp"
#include "../tasks/types/tasks.queueTypeStats.hpp"
#include "../__common_namespace/functions/getNonOptNat.hpp"
#include "../__common_namespace/functions/getMyValue.hpp"
#include "../__common_namespace/types/MyValue.hpp"
#include "../__common_namespace/types/myString.hpp"
#include "../__common_namespace/types/myInt.hpp"
#include "../__common_namespace/functions/getMyDouble.hpp"
#include "../__common_namespace/types/myDouble.hpp"
#include "../__common_namespace/functions/getMyDictOfInt.hpp"
#include "../__common_namespace/types/myDictOfInt.hpp"
#include "../__common_namespace/functions/getMaybeIface.hpp"
#include "../service1/types/service1.Value.hpp"
#include "../service1/types/service1.strvalueWithTime.hpp"
#include "../service1/types/service1.strvalue.hpp"
#include "../service1/types/service1.not_found.hpp"
#include "../service1/types/service1.longvalueWithTime.hpp"
#include "../service1/types/service1.longvalue.hpp"
#include "../__common_namespace/functions/getFloat.hpp"
#include "../__common_namespace/functions/getDouble.hpp"
#include "../__common_namespace/functions/get_arrays.hpp"
#include "../__common_namespace/types/float.hpp"
#include "../__common_namespace/types/fieldConflict4.hpp"
#include "../__common_namespace/types/fieldConflict3.hpp"
#include "../__common_namespace/types/fieldConflict2.hpp"
#include "../__common_namespace/types/fieldConflict1.hpp"
#include "../__common_namespace/types/double.hpp"
#include "../__common_namespace/types/dictionary.hpp"
#include "../__common_namespace/functions/boxedVector64.hpp"
#include "../__common_namespace/functions/boxedVector32BoxedElem.hpp"
#include "../__common_namespace/functions/boxedVector32.hpp"
#include "../__common_namespace/functions/boxedTupleSlice3.hpp"
#include "../__common_namespace/functions/boxedTupleSlice2.hpp"
#include "../__common_namespace/types/myBoxedTupleSlice.hpp"
#include "../__common_namespace/functions/boxedTupleSlice1.hpp"
#include "../__common_namespace/functions/boxedTuple.hpp"
#include "../__common_namespace/functions/boxedString.hpp"
#include "../__common_namespace/functions/boxedInt.hpp"
#include "../__common_namespace/functions/boxedArray.hpp"
#include "../__common_namespace/types/myBoxedArray.hpp"
#include "../__common_namespace/types/boolStat.hpp"
#include "../__common_namespace/types/Bool.hpp"
#include "../__common_namespace/types/benchObject.hpp"
#include "../__common_namespace/types/integer.hpp"
#include "../antispam/types/antispam.PatternFull.hpp"
#include "../antispam/types/antispam.patternNotFound.hpp"
#include "../antispam/types/antispam.patternFound.hpp"
#include "../antispam/functions/antispam.getPattern.hpp"

void tl2::factory::set_all_factories() {

	struct tl2_antispam_GetPattern_tl_function : public tl2::meta::tl_function {
        tl2::antispam::GetPattern object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			tl2::antispam::PatternFull result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("antispam.getPattern", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_antispam_GetPattern_tl_function>();
	});

	tl2::meta::set_create_function_by_name("antispam.getPattern", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_antispam_GetPattern_tl_function>();
	});

	struct tl2_antispam_PatternFound_tl_object : public tl2::meta::tl_object {
        tl2::antispam::PatternFound object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("antispam.patternFound", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_antispam_PatternFound_tl_object>();
	});

	struct tl2_antispam_PatternNotFound_tl_object : public tl2::meta::tl_object {
        tl2::antispam::PatternNotFound object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("antispam.patternNotFound", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_antispam_PatternNotFound_tl_object>();
	});

	struct tl2_BenchObject_tl_object : public tl2::meta::tl_object {
        tl2::BenchObject object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("benchObject", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_BenchObject_tl_object>();
	});

	struct tl2_BoolStat_tl_object : public tl2::meta::tl_object {
        tl2::BoolStat object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("boolStat", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_BoolStat_tl_object>();
	});

	struct tl2_BoxedArray_tl_function : public tl2::meta::tl_function {
        tl2::BoxedArray object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			tl2::MyBoxedArray result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("boxedArray", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_BoxedArray_tl_function>();
	});

	tl2::meta::set_create_function_by_name("boxedArray", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_BoxedArray_tl_function>();
	});

	struct tl2_BoxedInt_tl_function : public tl2::meta::tl_function {
        tl2::BoxedInt object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			int32_t result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("boxedInt", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_BoxedInt_tl_function>();
	});

	tl2::meta::set_create_function_by_name("boxedInt", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_BoxedInt_tl_function>();
	});

	struct tl2_BoxedString_tl_function : public tl2::meta::tl_function {
        tl2::BoxedString object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			std::string result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("boxedString", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_BoxedString_tl_function>();
	});

	tl2::meta::set_create_function_by_name("boxedString", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_BoxedString_tl_function>();
	});

	struct tl2_BoxedTuple_tl_function : public tl2::meta::tl_function {
        tl2::BoxedTuple object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			std::array<int32_t, 3> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("boxedTuple", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_BoxedTuple_tl_function>();
	});

	tl2::meta::set_create_function_by_name("boxedTuple", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_BoxedTuple_tl_function>();
	});

	struct tl2_BoxedTupleSlice1_tl_function : public tl2::meta::tl_function {
        tl2::BoxedTupleSlice1 object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			std::vector<int32_t> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("boxedTupleSlice1", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_BoxedTupleSlice1_tl_function>();
	});

	tl2::meta::set_create_function_by_name("boxedTupleSlice1", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_BoxedTupleSlice1_tl_function>();
	});

	struct tl2_BoxedTupleSlice2_tl_function : public tl2::meta::tl_function {
        tl2::BoxedTupleSlice2 object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			tl2::MyBoxedTupleSlice result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("boxedTupleSlice2", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_BoxedTupleSlice2_tl_function>();
	});

	tl2::meta::set_create_function_by_name("boxedTupleSlice2", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_BoxedTupleSlice2_tl_function>();
	});

	struct tl2_BoxedTupleSlice3_tl_function : public tl2::meta::tl_function {
        tl2::BoxedTupleSlice3 object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			std::vector<int32_t> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("boxedTupleSlice3", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_BoxedTupleSlice3_tl_function>();
	});

	tl2::meta::set_create_function_by_name("boxedTupleSlice3", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_BoxedTupleSlice3_tl_function>();
	});

	struct tl2_BoxedVector32_tl_function : public tl2::meta::tl_function {
        tl2::BoxedVector32 object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			std::vector<int32_t> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("boxedVector32", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_BoxedVector32_tl_function>();
	});

	tl2::meta::set_create_function_by_name("boxedVector32", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_BoxedVector32_tl_function>();
	});

	struct tl2_BoxedVector32BoxedElem_tl_function : public tl2::meta::tl_function {
        tl2::BoxedVector32BoxedElem object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			std::vector<int32_t> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("boxedVector32BoxedElem", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_BoxedVector32BoxedElem_tl_function>();
	});

	tl2::meta::set_create_function_by_name("boxedVector32BoxedElem", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_BoxedVector32BoxedElem_tl_function>();
	});

	struct tl2_BoxedVector64_tl_function : public tl2::meta::tl_function {
        tl2::BoxedVector64 object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			std::vector<int64_t> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("boxedVector64", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_BoxedVector64_tl_function>();
	});

	tl2::meta::set_create_function_by_name("boxedVector64", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_BoxedVector64_tl_function>();
	});

	struct tl2_FieldConflict1_tl_object : public tl2::meta::tl_object {
        tl2::FieldConflict1 object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("fieldConflict1", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_FieldConflict1_tl_object>();
	});

	struct tl2_FieldConflict2_tl_object : public tl2::meta::tl_object {
        tl2::FieldConflict2 object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("fieldConflict2", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_FieldConflict2_tl_object>();
	});

	struct tl2_FieldConflict3_tl_object : public tl2::meta::tl_object {
        tl2::FieldConflict3 object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("fieldConflict3", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_FieldConflict3_tl_object>();
	});

	struct tl2_FieldConflict4_tl_object : public tl2::meta::tl_object {
        tl2::FieldConflict4 object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("fieldConflict4", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_FieldConflict4_tl_object>();
	});

	struct tl2_Get_arrays_tl_function : public tl2::meta::tl_function {
        tl2::Get_arrays object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			std::array<int32_t, 5> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("get_arrays", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_Get_arrays_tl_function>();
	});

	tl2::meta::set_create_function_by_name("get_arrays", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_Get_arrays_tl_function>();
	});

	struct tl2_GetDouble_tl_function : public tl2::meta::tl_function {
        tl2::GetDouble object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			double result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("getDouble", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_GetDouble_tl_function>();
	});

	tl2::meta::set_create_function_by_name("getDouble", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_GetDouble_tl_function>();
	});

	struct tl2_GetFloat_tl_function : public tl2::meta::tl_function {
        tl2::GetFloat object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			float result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("getFloat", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_GetFloat_tl_function>();
	});

	tl2::meta::set_create_function_by_name("getFloat", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_GetFloat_tl_function>();
	});

	struct tl2_GetMaybeIface_tl_function : public tl2::meta::tl_function {
        tl2::GetMaybeIface object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			std::optional<::tl2::service1::Value> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("getMaybeIface", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_GetMaybeIface_tl_function>();
	});

	tl2::meta::set_create_function_by_name("getMaybeIface", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_GetMaybeIface_tl_function>();
	});

	struct tl2_GetMyDictOfInt_tl_function : public tl2::meta::tl_function {
        tl2::GetMyDictOfInt object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			tl2::MyDictOfInt result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("getMyDictOfInt", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_GetMyDictOfInt_tl_function>();
	});

	tl2::meta::set_create_function_by_name("getMyDictOfInt", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_GetMyDictOfInt_tl_function>();
	});

	struct tl2_GetMyDouble_tl_function : public tl2::meta::tl_function {
        tl2::GetMyDouble object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			tl2::MyDouble result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("getMyDouble", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_GetMyDouble_tl_function>();
	});

	tl2::meta::set_create_function_by_name("getMyDouble", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_GetMyDouble_tl_function>();
	});

	struct tl2_GetMyValue_tl_function : public tl2::meta::tl_function {
        tl2::GetMyValue object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			tl2::MyValue result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("getMyValue", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_GetMyValue_tl_function>();
	});

	tl2::meta::set_create_function_by_name("getMyValue", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_GetMyValue_tl_function>();
	});

	struct tl2_GetNonOptNat_tl_function : public tl2::meta::tl_function {
        tl2::GetNonOptNat object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			std::vector<int32_t> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("getNonOptNat", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_GetNonOptNat_tl_function>();
	});

	tl2::meta::set_create_function_by_name("getNonOptNat", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_GetNonOptNat_tl_function>();
	});

	struct tl2_GetStats_tl_function : public tl2::meta::tl_function {
        tl2::GetStats object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			tl2::tasks::QueueTypeStats result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("getStats", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_GetStats_tl_function>();
	});

	tl2::meta::set_create_function_by_name("getStats", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_GetStats_tl_function>();
	});

	struct tl2_Integer_tl_object : public tl2::meta::tl_object {
        tl2::Integer object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("integer", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_Integer_tl_object>();
	});

	struct tl2_Issue3498_tl_object : public tl2::meta::tl_object {
        tl2::Issue3498 object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("issue3498", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_Issue3498_tl_object>();
	});

	struct tl2_MyBoxedArray_tl_object : public tl2::meta::tl_object {
        tl2::MyBoxedArray object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("myBoxedArray", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_MyBoxedArray_tl_object>();
	});

	struct tl2_MyBoxedTupleSlice_tl_object : public tl2::meta::tl_object {
        tl2::MyBoxedTupleSlice object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("myBoxedTupleSlice", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_MyBoxedTupleSlice_tl_object>();
	});

	struct tl2_MyBoxedVectorSlice_tl_object : public tl2::meta::tl_object {
        tl2::MyBoxedVectorSlice object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("myBoxedVectorSlice", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_MyBoxedVectorSlice_tl_object>();
	});

	struct tl2_MyInt_tl_object : public tl2::meta::tl_object {
        tl2::MyInt object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("myInt", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_MyInt_tl_object>();
	});

	struct tl2_MyMcValue_tl_object : public tl2::meta::tl_object {
        tl2::MyMcValue object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("myMcValue", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_MyMcValue_tl_object>();
	});

	struct tl2_MyMcValueTuple_tl_object : public tl2::meta::tl_object {
        tl2::MyMcValueTuple object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("myMcValueTuple", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_MyMcValueTuple_tl_object>();
	});

	struct tl2_MyMcValueVector_tl_object : public tl2::meta::tl_object {
        tl2::MyMcValueVector object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("myMcValueVector", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_MyMcValueVector_tl_object>();
	});

	struct tl2_MyString_tl_object : public tl2::meta::tl_object {
        tl2::MyString object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("myString", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_MyString_tl_object>();
	});

	struct tl2_MyTwoDicts_tl_object : public tl2::meta::tl_object {
        tl2::MyTwoDicts object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("myTwoDicts", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_MyTwoDicts_tl_object>();
	});

	struct tl2_NonOptNat_tl_object : public tl2::meta::tl_object {
        tl2::NonOptNat object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("nonOptNat", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_NonOptNat_tl_object>();
	});

	struct tl2_pkg2_Foo_tl_object : public tl2::meta::tl_object {
        tl2::pkg2::Foo object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("pkg2.foo", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_pkg2_Foo_tl_object>();
	});

	struct tl2_pkg2_T1_tl_object : public tl2::meta::tl_object {
        tl2::pkg2::T1 object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("pkg2.t1", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_pkg2_T1_tl_object>();
	});

	struct tl2_RpcInvokeReqExtra_tl_object : public tl2::meta::tl_object {
        tl2::RpcInvokeReqExtra object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("rpcInvokeReqExtra", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_RpcInvokeReqExtra_tl_object>();
	});

	struct tl2_service1_Add_tl_function : public tl2::meta::tl_function {
        tl2::service1::Add object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.add", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Add_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service1.add", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_Add_tl_function>();
	});

	struct tl2_service1_AddOrGet_tl_function : public tl2::meta::tl_function {
        tl2::service1::AddOrGet object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			tl2::service1::Value result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.addOrGet", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_AddOrGet_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service1.addOrGet", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_AddOrGet_tl_function>();
	});

	struct tl2_service1_AddOrIncr_tl_function : public tl2::meta::tl_function {
        tl2::service1::AddOrIncr object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			tl2::service1::Value result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.addOrIncr", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_AddOrIncr_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service1.addOrIncr", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_AddOrIncr_tl_function>();
	});

	struct tl2_service1_Append_tl_function : public tl2::meta::tl_function {
        tl2::service1::Append object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.append", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Append_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service1.append", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_Append_tl_function>();
	});

	struct tl2_service1_Cas_tl_function : public tl2::meta::tl_function {
        tl2::service1::Cas object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.cas", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Cas_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service1.cas", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_Cas_tl_function>();
	});

	struct tl2_service1_Decr_tl_function : public tl2::meta::tl_function {
        tl2::service1::Decr object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			tl2::service1::Value result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.decr", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Decr_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service1.decr", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_Decr_tl_function>();
	});

	struct tl2_service1_Delete_tl_function : public tl2::meta::tl_function {
        tl2::service1::Delete object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.delete", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Delete_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service1.delete", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_Delete_tl_function>();
	});

	struct tl2_service1_DisableExpiration_tl_function : public tl2::meta::tl_function {
        tl2::service1::DisableExpiration object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.disableExpiration", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_DisableExpiration_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service1.disableExpiration", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_DisableExpiration_tl_function>();
	});

	struct tl2_service1_DisableKeysStat_tl_function : public tl2::meta::tl_function {
        tl2::service1::DisableKeysStat object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.disableKeysStat", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_DisableKeysStat_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service1.disableKeysStat", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_DisableKeysStat_tl_function>();
	});

	struct tl2_service1_EnableExpiration_tl_function : public tl2::meta::tl_function {
        tl2::service1::EnableExpiration object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.enableExpiration", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_EnableExpiration_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service1.enableExpiration", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_EnableExpiration_tl_function>();
	});

	struct tl2_service1_EnableKeysStat_tl_function : public tl2::meta::tl_function {
        tl2::service1::EnableKeysStat object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.enableKeysStat", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_EnableKeysStat_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service1.enableKeysStat", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_EnableKeysStat_tl_function>();
	});

	struct tl2_service1_Exists_tl_function : public tl2::meta::tl_function {
        tl2::service1::Exists object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.exists", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Exists_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service1.exists", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_Exists_tl_function>();
	});

	struct tl2_service1_Get_tl_function : public tl2::meta::tl_function {
        tl2::service1::Get object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			tl2::service1::Value result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.get", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Get_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service1.get", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_Get_tl_function>();
	});

	struct tl2_service1_GetExpireTime_tl_function : public tl2::meta::tl_function {
        tl2::service1::GetExpireTime object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			std::optional<int32_t> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.getExpireTime", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_GetExpireTime_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service1.getExpireTime", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_GetExpireTime_tl_function>();
	});

	struct tl2_service1_GetKeysStat_tl_function : public tl2::meta::tl_function {
        tl2::service1::GetKeysStat object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			std::optional<::tl2::service1::KeysStat> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.getKeysStat", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_GetKeysStat_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service1.getKeysStat", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_GetKeysStat_tl_function>();
	});

	struct tl2_service1_GetKeysStatPeriods_tl_function : public tl2::meta::tl_function {
        tl2::service1::GetKeysStatPeriods object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			std::vector<int32_t> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.getKeysStatPeriods", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_GetKeysStatPeriods_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service1.getKeysStatPeriods", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_GetKeysStatPeriods_tl_function>();
	});

	struct tl2_service1_GetWildcard_tl_function : public tl2::meta::tl_function {
        tl2::service1::GetWildcard object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			std::vector<::tl2::Map<std::string, std::string>> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.getWildcard", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_GetWildcard_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service1.getWildcard", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_GetWildcard_tl_function>();
	});

	struct tl2_service1_GetWildcardDict_tl_function : public tl2::meta::tl_function {
        tl2::service1::GetWildcardDict object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			tl2::Dictionary<std::string> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.getWildcardDict", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_GetWildcardDict_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service1.getWildcardDict", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_GetWildcardDict_tl_function>();
	});

	struct tl2_service1_GetWildcardList_tl_function : public tl2::meta::tl_function {
        tl2::service1::GetWildcardList object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			std::vector<std::string> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.getWildcardList", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_GetWildcardList_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service1.getWildcardList", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_GetWildcardList_tl_function>();
	});

	struct tl2_service1_GetWildcardWithFlags_tl_function : public tl2::meta::tl_function {
        tl2::service1::GetWildcardWithFlags object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			tl2::Dictionary<::tl2::service1::Value> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.getWildcardWithFlags", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_GetWildcardWithFlags_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service1.getWildcardWithFlags", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_GetWildcardWithFlags_tl_function>();
	});

	struct tl2_service1_Incr_tl_function : public tl2::meta::tl_function {
        tl2::service1::Incr object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			tl2::service1::Value result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.incr", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Incr_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service1.incr", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_Incr_tl_function>();
	});

	struct tl2_service1_KeysStat_tl_object : public tl2::meta::tl_object {
        tl2::service1::KeysStat object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("service1.keysStat", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_KeysStat_tl_object>();
	});

	struct tl2_service1_Longvalue_tl_object : public tl2::meta::tl_object {
        tl2::service1::Longvalue object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("service1.longvalue", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Longvalue_tl_object>();
	});

	struct tl2_service1_LongvalueWithTime_tl_object : public tl2::meta::tl_object {
        tl2::service1::LongvalueWithTime object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("service1.longvalueWithTime", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_LongvalueWithTime_tl_object>();
	});

	struct tl2_service1_Not_found_tl_object : public tl2::meta::tl_object {
        tl2::service1::Not_found object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("service1.not_found", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Not_found_tl_object>();
	});

	struct tl2_service1_Replace_tl_function : public tl2::meta::tl_function {
        tl2::service1::Replace object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.replace", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Replace_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service1.replace", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_Replace_tl_function>();
	});

	struct tl2_service1_ReplaceOrIncr_tl_function : public tl2::meta::tl_function {
        tl2::service1::ReplaceOrIncr object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			tl2::service1::Value result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.replaceOrIncr", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_ReplaceOrIncr_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service1.replaceOrIncr", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_ReplaceOrIncr_tl_function>();
	});

	struct tl2_service1_Set_tl_function : public tl2::meta::tl_function {
        tl2::service1::Set object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.set", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Set_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service1.set", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_Set_tl_function>();
	});

	struct tl2_service1_SetOrIncr_tl_function : public tl2::meta::tl_function {
        tl2::service1::SetOrIncr object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			tl2::service1::Value result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.setOrIncr", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_SetOrIncr_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service1.setOrIncr", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_SetOrIncr_tl_function>();
	});

	struct tl2_service1_Strvalue_tl_object : public tl2::meta::tl_object {
        tl2::service1::Strvalue object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("service1.strvalue", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Strvalue_tl_object>();
	});

	struct tl2_service1_StrvalueWithTime_tl_object : public tl2::meta::tl_object {
        tl2::service1::StrvalueWithTime object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("service1.strvalueWithTime", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_StrvalueWithTime_tl_object>();
	});

	struct tl2_service1_Touch_tl_function : public tl2::meta::tl_function {
        tl2::service1::Touch object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service1.touch", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service1_Touch_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service1.touch", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service1_Touch_tl_function>();
	});

	struct tl2_service2_AddOrIncrMany_tl_function : public tl2::meta::tl_function {
        tl2::service2::AddOrIncrMany object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			std::vector<::tl2::service2::CounterSet> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service2.addOrIncrMany", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service2_AddOrIncrMany_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service2.addOrIncrMany", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service2_AddOrIncrMany_tl_function>();
	});

	struct tl2_service2_Set_tl_function : public tl2::meta::tl_function {
        tl2::service2::Set object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			tl2::True result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service2.set", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service2_Set_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service2.set", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service2_Set_tl_function>();
	});

	struct tl2_service2_SetObjectTtl_tl_function : public tl2::meta::tl_function {
        tl2::service2::SetObjectTtl object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			tl2::True result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service2.setObjectTtl", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service2_SetObjectTtl_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service2.setObjectTtl", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service2_SetObjectTtl_tl_function>();
	});

	struct tl2_service3_CreateProduct_tl_function : public tl2::meta::tl_function {
        tl2::service3::CreateProduct object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.createProduct", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_CreateProduct_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service3.createProduct", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_CreateProduct_tl_function>();
	});

	struct tl2_service3_DeleteAllProducts_tl_function : public tl2::meta::tl_function {
        tl2::service3::DeleteAllProducts object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.deleteAllProducts", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_DeleteAllProducts_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service3.deleteAllProducts", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_DeleteAllProducts_tl_function>();
	});

	struct tl2_service3_DeleteGroupedProducts_tl_function : public tl2::meta::tl_function {
        tl2::service3::DeleteGroupedProducts object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.deleteGroupedProducts", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_DeleteGroupedProducts_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service3.deleteGroupedProducts", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_DeleteGroupedProducts_tl_function>();
	});

	struct tl2_service3_DeleteProduct_tl_function : public tl2::meta::tl_function {
        tl2::service3::DeleteProduct object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.deleteProduct", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_DeleteProduct_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service3.deleteProduct", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_DeleteProduct_tl_function>();
	});

	struct tl2_service3_GetLastVisitTimestamp_tl_function : public tl2::meta::tl_function {
        tl2::service3::GetLastVisitTimestamp object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			std::optional<int32_t> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.getLastVisitTimestamp", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_GetLastVisitTimestamp_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service3.getLastVisitTimestamp", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_GetLastVisitTimestamp_tl_function>();
	});

	struct tl2_service3_GetLimits_tl_function : public tl2::meta::tl_function {
        tl2::service3::GetLimits object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			tl2::service3::Limits result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.getLimits", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_GetLimits_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service3.getLimits", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_GetLimits_tl_function>();
	});

	struct tl2_service3_GetProductStats_tl_function : public tl2::meta::tl_function {
        tl2::service3::GetProductStats object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			std::optional<std::vector<::tl2::service3::ProductStatsOld>> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.getProductStats", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_GetProductStats_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service3.getProductStats", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_GetProductStats_tl_function>();
	});

	struct tl2_service3_GetProducts_tl_function : public tl2::meta::tl_function {
        tl2::service3::GetProducts object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			std::optional<std::vector<::tl2::service3::Product>> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.getProducts", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_GetProducts_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service3.getProducts", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_GetProducts_tl_function>();
	});

	struct tl2_service3_GetScheduledProducts_tl_function : public tl2::meta::tl_function {
        tl2::service3::GetScheduledProducts object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			std::optional<std::vector<::tl2::service3::Productmode<0>>> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.getScheduledProducts", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_GetScheduledProducts_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service3.getScheduledProducts", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_GetScheduledProducts_tl_function>();
	});

	struct tl2_service3_GroupCountLimit_tl_object : public tl2::meta::tl_object {
        tl2::service3::GroupCountLimit object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("service3.groupCountLimit", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_GroupCountLimit_tl_object>();
	});

	struct tl2_service3_GroupSizeLimit_tl_object : public tl2::meta::tl_object {
        tl2::service3::GroupSizeLimit object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("service3.groupSizeLimit", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_GroupSizeLimit_tl_object>();
	});

	struct tl2_service3_Limits_tl_object : public tl2::meta::tl_object {
        tl2::service3::Limits object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("service3.limits", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_Limits_tl_object>();
	});

	struct tl2_service3_ProductStatsOld_tl_object : public tl2::meta::tl_object {
        tl2::service3::ProductStatsOld object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("service3.productStatsOld", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_ProductStatsOld_tl_object>();
	});

	struct tl2_service3_RestoreAllProducts_tl_function : public tl2::meta::tl_function {
        tl2::service3::RestoreAllProducts object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.restoreAllProducts", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_RestoreAllProducts_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service3.restoreAllProducts", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_RestoreAllProducts_tl_function>();
	});

	struct tl2_service3_RestoreGroupedProducts_tl_function : public tl2::meta::tl_function {
        tl2::service3::RestoreGroupedProducts object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.restoreGroupedProducts", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_RestoreGroupedProducts_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service3.restoreGroupedProducts", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_RestoreGroupedProducts_tl_function>();
	});

	struct tl2_service3_RestoreProduct_tl_function : public tl2::meta::tl_function {
        tl2::service3::RestoreProduct object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.restoreProduct", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_RestoreProduct_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service3.restoreProduct", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_RestoreProduct_tl_function>();
	});

	struct tl2_service3_SetLastVisitTimestamp_tl_function : public tl2::meta::tl_function {
        tl2::service3::SetLastVisitTimestamp object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			bool result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.setLastVisitTimestamp", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_SetLastVisitTimestamp_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service3.setLastVisitTimestamp", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_SetLastVisitTimestamp_tl_function>();
	});

	struct tl2_service3_SetLimits_tl_function : public tl2::meta::tl_function {
        tl2::service3::SetLimits object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			tl2::BoolStat result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service3.setLimits", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service3_SetLimits_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service3.setLimits", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service3_SetLimits_tl_function>();
	});

	struct tl2_service4_ModifiedNewsEntry_tl_object : public tl2::meta::tl_object {
        tl2::service4::ModifiedNewsEntry object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("service4.modifiedNewsEntry", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service4_ModifiedNewsEntry_tl_object>();
	});

	struct tl2_service4_Object_tl_object : public tl2::meta::tl_object {
        tl2::service4::Object object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("service4.object", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service4_Object_tl_object>();
	});

	struct tl2_service5_EmptyOutput_tl_object : public tl2::meta::tl_object {
        tl2::service5::EmptyOutput object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("service5.emptyOutput", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service5_EmptyOutput_tl_object>();
	});

	struct tl2_service5_Insert_tl_function : public tl2::meta::tl_function {
        tl2::service5::Insert object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			tl2::service5::Output result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service5.insert", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service5_Insert_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service5.insert", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service5_Insert_tl_function>();
	});

	struct tl2_service5_Params_tl_object : public tl2::meta::tl_object {
        tl2::service5::Params object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("service5.params", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service5_Params_tl_object>();
	});

	struct tl2_service5_PerformQuery_tl_function : public tl2::meta::tl_function {
        tl2::service5::PerformQuery object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			tl2::service5::Output result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service5.performQuery", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service5_PerformQuery_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service5.performQuery", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service5_PerformQuery_tl_function>();
	});

	struct tl2_service5_Query_tl_function : public tl2::meta::tl_function {
        tl2::service5::Query object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			tl2::service5::Output result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service5.query", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service5_Query_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service5.query", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service5_Query_tl_function>();
	});

	struct tl2_service5_StringOutput_tl_object : public tl2::meta::tl_object {
        tl2::service5::StringOutput object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("service5.stringOutput", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service5_StringOutput_tl_object>();
	});

	struct tl2_service6_Error_tl_object : public tl2::meta::tl_object {
        tl2::service6::Error object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("service6.error", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service6_Error_tl_object>();
	});

	struct tl2_service6_FindResultRow_tl_object : public tl2::meta::tl_object {
        tl2::service6::FindResultRow object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("service6.findResultRow", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service6_FindResultRow_tl_object>();
	});

	struct tl2_service6_FindWithBoundsResult_tl_object : public tl2::meta::tl_object {
        tl2::service6::FindWithBoundsResult object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("service6.findWithBoundsResult", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service6_FindWithBoundsResult_tl_object>();
	});

	struct tl2_service6_MultiFind_tl_function : public tl2::meta::tl_function {
        tl2::service6::MultiFind object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service6.multiFind", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service6_MultiFind_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service6.multiFind", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service6_MultiFind_tl_function>();
	});

	struct tl2_service6_MultiFindWithBounds_tl_function : public tl2::meta::tl_function {
        tl2::service6::MultiFindWithBounds object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("service6.multiFindWithBounds", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_service6_MultiFindWithBounds_tl_function>();
	});

	tl2::meta::set_create_function_by_name("service6.multiFindWithBounds", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_service6_MultiFindWithBounds_tl_function>();
	});

	struct tl2_StatOne_tl_object : public tl2::meta::tl_object {
        tl2::StatOne object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("statOne", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_StatOne_tl_object>();
	});

	struct tl2_tasks_AddTask_tl_function : public tl2::meta::tl_function {
        tl2::tasks::AddTask object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			int64_t result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("tasks.addTask", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_AddTask_tl_function>();
	});

	tl2::meta::set_create_function_by_name("tasks.addTask", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_tasks_AddTask_tl_function>();
	});

	struct tl2_tasks_CronTask_tl_object : public tl2::meta::tl_object {
        tl2::tasks::CronTask object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("tasks.cronTask", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_CronTask_tl_object>();
	});

	struct tl2_tasks_CronTaskWithId_tl_object : public tl2::meta::tl_object {
        tl2::tasks::CronTaskWithId object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("tasks.cronTaskWithId", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_CronTaskWithId_tl_object>();
	});

	struct tl2_tasks_CronTime_tl_object : public tl2::meta::tl_object {
        tl2::tasks::CronTime object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("tasks.cronTime", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_CronTime_tl_object>();
	});

	struct tl2_tasks_GetAnyTask_tl_function : public tl2::meta::tl_function {
        tl2::tasks::GetAnyTask object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			std::optional<::tl2::tasks::TaskInfo> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("tasks.getAnyTask", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_GetAnyTask_tl_function>();
	});

	tl2::meta::set_create_function_by_name("tasks.getAnyTask", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_tasks_GetAnyTask_tl_function>();
	});

	struct tl2_tasks_GetQueueSize_tl_function : public tl2::meta::tl_function {
        tl2::tasks::GetQueueSize object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			tl2::tasks::QueueStats result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("tasks.getQueueSize", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_GetQueueSize_tl_function>();
	});

	tl2::meta::set_create_function_by_name("tasks.getQueueSize", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_tasks_GetQueueSize_tl_function>();
	});

	struct tl2_tasks_GetQueueTypes_tl_function : public tl2::meta::tl_function {
        tl2::tasks::GetQueueTypes object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			std::vector<::tl2::tasks::QueueTypeInfo> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("tasks.getQueueTypes", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_GetQueueTypes_tl_function>();
	});

	tl2::meta::set_create_function_by_name("tasks.getQueueTypes", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_tasks_GetQueueTypes_tl_function>();
	});

	struct tl2_tasks_GetTaskFromQueue_tl_function : public tl2::meta::tl_function {
        tl2::tasks::GetTaskFromQueue object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			std::optional<::tl2::tasks::TaskInfo> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("tasks.getTaskFromQueue", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_GetTaskFromQueue_tl_function>();
	});

	tl2::meta::set_create_function_by_name("tasks.getTaskFromQueue", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_tasks_GetTaskFromQueue_tl_function>();
	});

	struct tl2_tasks_QueueTypeInfo_tl_object : public tl2::meta::tl_object {
        tl2::tasks::QueueTypeInfo object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("tasks.queueTypeInfo", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_QueueTypeInfo_tl_object>();
	});

	struct tl2_tasks_QueueTypeSettings_tl_object : public tl2::meta::tl_object {
        tl2::tasks::QueueTypeSettings object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("tasks.queueTypeSettings", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_QueueTypeSettings_tl_object>();
	});

	struct tl2_tasks_QueueTypeStats_tl_object : public tl2::meta::tl_object {
        tl2::tasks::QueueTypeStats object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("tasks.queueTypeStats", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_QueueTypeStats_tl_object>();
	});

	struct tl2_tasks_Task_tl_object : public tl2::meta::tl_object {
        tl2::tasks::Task object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("tasks.task", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_Task_tl_object>();
	});

	struct tl2_tasks_TaskInfo_tl_object : public tl2::meta::tl_object {
        tl2::tasks::TaskInfo object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("tasks.taskInfo", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_TaskInfo_tl_object>();
	});

	struct tl2_tasks_TaskStatusInProgress_tl_object : public tl2::meta::tl_object {
        tl2::tasks::TaskStatusInProgress object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("tasks.taskStatusInProgress", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_TaskStatusInProgress_tl_object>();
	});

	struct tl2_tasks_TaskStatusNotCurrentlyInEngine_tl_object : public tl2::meta::tl_object {
        tl2::tasks::TaskStatusNotCurrentlyInEngine object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("tasks.taskStatusNotCurrentlyInEngine", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_TaskStatusNotCurrentlyInEngine_tl_object>();
	});

	struct tl2_tasks_TaskStatusScheduled_tl_object : public tl2::meta::tl_object {
        tl2::tasks::TaskStatusScheduled object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("tasks.taskStatusScheduled", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_TaskStatusScheduled_tl_object>();
	});

	struct tl2_tasks_TaskStatusWaiting_tl_object : public tl2::meta::tl_object {
        tl2::tasks::TaskStatusWaiting object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("tasks.taskStatusWaiting", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tasks_TaskStatusWaiting_tl_object>();
	});

	struct tl2_tree_stats_ObjectLimitValueLong_tl_object : public tl2::meta::tl_object {
        tl2::tree_stats::ObjectLimitValueLong object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("tree_stats.objectLimitValueLong", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_tree_stats_ObjectLimitValueLong_tl_object>();
	});

	struct tl2_True_tl_object : public tl2::meta::tl_object {
        tl2::True object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("true", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_True_tl_object>();
	});

	struct tl2_unique_Get_tl_function : public tl2::meta::tl_function {
        tl2::unique::Get object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			std::optional<int32_t> result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("unique.get", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_unique_Get_tl_function>();
	});

	tl2::meta::set_create_function_by_name("unique.get", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_unique_Get_tl_function>();
	});

	struct tl2_unique_StringToInt_tl_function : public tl2::meta::tl_function {
        tl2::unique::StringToInt object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

		bool read_write_result(basictl::tl_istream &in, basictl::tl_ostream &out) override {
			int32_t result;
			bool read_result = this->object.read_result(in, result);
			if (!read_result) {
				return false;
			}
			return this->object.write_result(out, result);
		}

    };
	tl2::meta::set_create_object_by_name("unique.stringToInt", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_unique_StringToInt_tl_function>();
	});

	tl2::meta::set_create_function_by_name("unique.stringToInt", []() -> std::unique_ptr<tl2::meta::tl_function> {
        return std::make_unique<tl2_unique_StringToInt_tl_function>();
	});

	struct tl2_WithFloat_tl_object : public tl2::meta::tl_object {
        tl2::WithFloat object;

        bool read(basictl::tl_istream &s) override {return object.read(s);}
        bool write(basictl::tl_ostream &s) override {return object.write(s);}
        
		bool read_boxed(basictl::tl_istream &s) override {return object.read_boxed(s);}
        bool write_boxed(basictl::tl_ostream &s) override {return object.write_boxed(s);}
		
		bool write_json(std::ostream &s) override {return object.write_json(s);}

    };
	tl2::meta::set_create_object_by_name("withFloat", []() -> std::unique_ptr<tl2::meta::tl_object> {
        return std::make_unique<tl2_WithFloat_tl_object>();
	});

}
