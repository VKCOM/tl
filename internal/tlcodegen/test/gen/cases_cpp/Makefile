# compiler options
CC = g++
CFLAGS = -std=c++20 -O3 -Wno-noexcept-type -g -Wall -Wextra -Werror -Wunused-parameter

# compile all object files together
all: __build/main.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o __build/benchmarks.o __build/cases.o __build/casesTL2.o __build/cases_bytes.o 
	@mkdir -p __build
	$(CC) $(CFLAGS) -o all __build/main.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o __build/benchmarks.o __build/cases.o __build/casesTL2.o __build/cases_bytes.o 

__build/main.o: main.cpp
	@mkdir -p __build
	$(CC) $(CFLAGS) -c main.cpp -o __build/main.o

# compile meta data collection
__build/__meta.o: __meta/headers.h __meta/details.cpp __build
	$(CC) $(CFLAGS) -I. -o __build/__meta.o -c __meta/details.cpp

# compile objects factories
__build/__factory.o: __factory/headers.h __factory/details.cpp __common_namespace/types/vector.h __common_namespace/types/tuple.h __common_namespace/types/int.h cases/types/cases.testVector.h cases/types/cases.testUnionContainer.h cases/types/cases.testUnion2.h cases/types/cases.testUnion1.h cases/types/cases.testTuple.h cases/types/cases.testRecursiveFieldMask.h cases/types/cases.testOutFieldMaskContainer.h cases/types/cases.testMaybe.h cases/types/cases.testLocalFieldmask.h cases/types/cases.testInplaceStructArgs2.h cases/types/cases.testInplaceStructArgs.h cases/types/cases.testEnumContainer.h cases/types/cases.testDictString.h cases/types/cases.testDictInt.h cases/types/cases.testDictAny.h cases/types/cases.testBeforeReadBitValidation.h cases/types/cases.testArray.h cases/types/cases.testAllPossibleFieldConfigsContainer.h casesTL2/types/casesTL2.testVectorBool.h casesTL2/types/casesTL2.testObject.h casesTL2/functions/casesTL2.testFunctionWithDep2.h casesTL2/functions/casesTL2.testFunctionWithDep1.h casesTL2/functions/casesTL2.testFunctionNoDep5.h casesTL2/functions/casesTL2.testFunctionNoDep4.h casesTL2/functions/casesTL2.testFunctionNoDep3.h casesTL2/functions/casesTL2.testFunctionNoDep2.h casesTL2/functions/casesTL2.testFunctionNoDep1.h casesTL2/types/casesTL2.testArrayFlexibleBool.h casesTL2/types/casesTL2.testArrayFixedBool.h cases/types/cases.replace7plusplus.h cases/types/cases.replace7plus.h cases/types/cases.replace7.h cases/types/cases.myCycle1.h cases/types/cases.myCycle2.h cases/types/cases.myCycle3.h cases_bytes/types/cases_bytes.testVector.h cases_bytes/types/cases_bytes.testTuple.h cases_bytes/types/cases_bytes.testEnumContainer.h cases/types/cases.TestEnumItems.h cases_bytes/types/cases_bytes.TestEnumItems.h cases_bytes/types/cases_bytes.testDictStringString.h cases_bytes/types/cases_bytes.testDictString.h cases_bytes/types/cases_bytes.testDictInt.h cases_bytes/types/cases_bytes.testDictAny.h cases_bytes/types/cases_bytes.testArray.h __common_namespace/types/Bool.h benchmarks/types/benchmarks.vrutoyTopLevelContainerWithDependency.h benchmarks/types/benchmarks.vrutoyTopLevelContainer.h benchmarks/types/benchmarks.vrutoytopLevelUnionEmpty.h benchmarks/types/benchmarks.vrutoytopLevelUnionBig.h benchmarks/types/benchmarks.vrutoyPositions.h benchmarks/types/benchmarks.vruposition.h __common_namespace/types/true.h benchmarks/types/benchmarks.vruhash.h __build
	$(CC) $(CFLAGS) -I. -o __build/__factory.o -c __factory/details.cpp

# compile streams which are used to work with io
__build/io_streams.o: basictl/constants.h basictl/errors.h basictl/io_connectors.h basictl/io_streams.cpp basictl/io_streams.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/io_streams.o -c basictl/io_streams.cpp

__build/io_throwable_streams.o: basictl/constants.h basictl/errors.h basictl/io_connectors.h basictl/io_throwable_streams.cpp basictl/io_throwable_streams.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/io_throwable_streams.o -c basictl/io_throwable_streams.cpp

__build/string_io.o: basictl/io_connectors.h basictl/impl/string_io.cpp basictl/impl/string_io.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/string_io.o -c basictl/impl/string_io.cpp

# build object files for individual namespaces
__build/__common_namespace.o: __common_namespace/details.cpp  __common_namespace/headers/dictionary.h __common_namespace/headers/dictionaryAny.h __common_namespace/headers/dictionaryField.h __common_namespace/headers/dictionaryFieldAny.h __common_namespace/headers/int.h __common_namespace/headers/int32.h __common_namespace/headers/int64.h __common_namespace/headers/true.h __common_namespace/headers/vector.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/__common_namespace.o -c __common_namespace/details.cpp
__build/benchmarks.o: benchmarks/details.cpp  benchmarks/headers/benchmarks.VrutoyTopLevelUnion.h benchmarks/headers/benchmarks.vruhash.h benchmarks/headers/benchmarks.vruposition.h benchmarks/headers/benchmarks.vrutoyPositions.h benchmarks/headers/benchmarks.vrutoyTopLevelContainer.h benchmarks/headers/benchmarks.vrutoyTopLevelContainerWithDependency.h benchmarks/headers/benchmarks.vrutoytopLevelUnionBig.h benchmarks/headers/benchmarks.vrutoytopLevelUnionEmpty.h benchmarks/headers/benchmarks_vector.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/benchmarks.o -c benchmarks/details.cpp
__build/cases.o: cases/details.cpp  cases/headers/cases.TestEnum.h cases/headers/cases.TestEnumItems.h cases/headers/cases.TestUnion.h cases/headers/cases.inplace1.h cases/headers/cases.inplace2.h cases/headers/cases.inplace3.h cases/headers/cases.myCycle1.h cases/headers/cases.myCycle2.h cases/headers/cases.myCycle3.h cases/headers/cases.replace7.h cases/headers/cases.replace7plus.h cases/headers/cases.replace7plusplus.h cases/headers/cases.testAllPossibleFieldConfigs.h cases/headers/cases.testAllPossibleFieldConfigsContainer.h cases/headers/cases.testArray.h cases/headers/cases.testBeforeReadBitValidation.h cases/headers/cases.testDictAny.h cases/headers/cases.testDictInt.h cases/headers/cases.testDictString.h cases/headers/cases.testEnumContainer.h cases/headers/cases.testInplaceStructArgs.h cases/headers/cases.testInplaceStructArgs2.h cases/headers/cases.testLocalFieldmask.h cases/headers/cases.testMaybe.h cases/headers/cases.testOutFieldMask.h cases/headers/cases.testOutFieldMaskContainer.h cases/headers/cases.testRecursiveFieldMask.h cases/headers/cases.testTuple.h cases/headers/cases.testUnion1.h cases/headers/cases.testUnion2.h cases/headers/cases.testUnionContainer.h cases/headers/cases.testVector.h cases/headers/cases_int.h cases/headers/cases_pair.h cases/headers/cases_tuple.h cases/headers/cases_vector.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/cases.o -c cases/details.cpp
__build/casesTL2.o: casesTL2/details.cpp  casesTL2/headers/casesTL2.testArrayFixedBool.h casesTL2/headers/casesTL2.testArrayFlexibleBool.h casesTL2/headers/casesTL2.testFunctionNoDep1.h casesTL2/headers/casesTL2.testFunctionNoDep2.h casesTL2/headers/casesTL2.testFunctionNoDep3.h casesTL2/headers/casesTL2.testFunctionNoDep4.h casesTL2/headers/casesTL2.testFunctionNoDep5.h casesTL2/headers/casesTL2.testFunctionWithDep1.h casesTL2/headers/casesTL2.testFunctionWithDep2.h casesTL2/headers/casesTL2.testObject.h casesTL2/headers/casesTL2.testVectorBool.h casesTL2/headers/casesTL2_Bool.h casesTL2/headers/casesTL2_benchmarks.VrutoyTopLevelUnion.h casesTL2/headers/casesTL2_true.h casesTL2/headers/casesTL2_vector.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/casesTL2.o -c casesTL2/details.cpp
__build/cases_bytes.o: cases_bytes/details.cpp  cases_bytes/headers/cases_bytes.TestEnum.h cases_bytes/headers/cases_bytes.TestEnumItems.h cases_bytes/headers/cases_bytes.testArray.h cases_bytes/headers/cases_bytes.testDictAny.h cases_bytes/headers/cases_bytes.testDictInt.h cases_bytes/headers/cases_bytes.testDictString.h cases_bytes/headers/cases_bytes.testDictStringString.h cases_bytes/headers/cases_bytes.testEnumContainer.h cases_bytes/headers/cases_bytes.testTuple.h cases_bytes/headers/cases_bytes.testVector.h cases_bytes/headers/cases_bytes_dictionary.h cases_bytes/headers/cases_bytes_dictionaryField.h cases_bytes/headers/cases_bytes_string.h cases_bytes/headers/cases_bytes_tuple.h cases_bytes/headers/cases_bytes_vector.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/cases_bytes.o -c cases_bytes/details.cpp
