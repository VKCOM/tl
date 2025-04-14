
CC = g++
CFLAGS = -std=c++20 -O3 -Wno-noexcept-type -g -Wall -Wextra -Werror=return-type -Wno-unused-parameter
# compile all object files together
all: main.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o __build/benchmarks.o __build/cases.o __build/cases_bytes.o 
	$(CC) $(CFLAGS) -o all main.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o __build/benchmarks.o __build/cases.o __build/cases_bytes.o 

main.o: main.cpp
	$(CC) $(CFLAGS) -c main.cpp

# compile meta sources such as meta data collection and objects factories
__build/__meta.o: __meta/headers.h __meta/details.cpp
	$(CC) $(CFLAGS) -o __build/__meta.o -c __meta/details.cpp
__build/__factory.o: __factory/headers.h __factory/details.cpp cases/types/cases.testVector.h cases/types/cases.testUnionContainer.h cases/types/cases.testUnion2.h cases/types/cases.testUnion1.h cases/types/cases.testTuple.h cases/types/cases.testRecursiveFieldMask.h cases/types/cases.testOutFieldMaskContainer.h cases/types/cases.testMaybe.h cases/types/cases.testLocalFieldmask.h cases/types/cases.testInplaceStructArgs2.h cases/types/cases.testInplaceStructArgs.h cases/types/cases.testEnumContainer.h cases/types/cases.testDictString.h cases/types/cases.testDictInt.h cases/types/cases.testDictAny.h cases/types/cases.testBeforeReadBitValidation.h cases/types/cases.testArray.h cases/types/cases.testAllPossibleFieldConfigsContainer.h cases/types/cases.replace7plusplus.h cases/types/cases.replace7plus.h cases/types/cases.replace7.h cases/types/cases.myCycle1.h cases/types/cases.myCycle2.h cases/types/cases.myCycle3.h cases_bytes/types/cases_bytes.testVector.h cases_bytes/types/cases_bytes.testTuple.h cases_bytes/types/cases_bytes.testEnumContainer.h cases/types/cases.TestEnumItems.h cases_bytes/types/cases_bytes.TestEnumItems.h cases_bytes/types/cases_bytes.testDictStringString.h cases_bytes/types/cases_bytes.testDictString.h cases_bytes/types/cases_bytes.testDictInt.h cases_bytes/types/cases_bytes.testDictAny.h cases_bytes/types/cases_bytes.testArray.h benchmarks/types/benchmarks.vrutoyTopLevelContainerWithDependency.h benchmarks/types/benchmarks.vrutoyTopLevelContainer.h benchmarks/types/benchmarks.vrutoytopLevelUnionEmpty.h benchmarks/types/benchmarks.vrutoytopLevelUnionBig.h benchmarks/types/benchmarks.vruposition.h __common_namespace/types/true.h benchmarks/types/benchmarks.vruhash.h
	$(CC) $(CFLAGS) -o __build/__factory.o -c __factory/details.cpp

# compile streams which are used to work with io
__build/io_streams.o: basictl/constants.h basictl/errors.h basictl/io_connectors.h basictl/io_streams.cpp basictl/io_streams.h
	$(CC) $(CFLAGS) -o __build/io_streams.o -c basictl/io_streams.cpp

__build/io_throwable_streams.o: basictl/constants.h basictl/errors.h basictl/io_connectors.h basictl/io_throwable_streams.cpp basictl/io_throwable_streams.h
	$(CC) $(CFLAGS) -o __build/io_throwable_streams.o -c basictl/io_throwable_streams.cpp

__build/string_io.o: basictl/io_connectors.h basictl/impl/string_io.cpp basictl/impl/string_io.h
	$(CC) $(CFLAGS) -o __build/string_io.o -c basictl/impl/string_io.cpp

# compile individual namespaces
__build/__common_namespace.o: __common_namespace/details.cpp __common_namespace/details.cpp __common_namespace/headers/dictionary.h __common_namespace/headers/dictionaryAny.h __common_namespace/headers/dictionaryField.h __common_namespace/headers/dictionaryFieldAny.h __common_namespace/headers/int32.h __common_namespace/headers/int64.h __common_namespace/headers/true.h __common_namespace/headers/vector.h
	$(CC) $(CFLAGS) -o __build/__common_namespace.o -c __common_namespace/details.cpp
__build/benchmarks.o: benchmarks/details.cpp benchmarks/details.cpp benchmarks/headers/benchmarks.VrutoyTopLevelUnion.h benchmarks/headers/benchmarks.vruhash.h benchmarks/headers/benchmarks.vruposition.h benchmarks/headers/benchmarks.vrutoyPositions.h benchmarks/headers/benchmarks.vrutoyTopLevelContainer.h benchmarks/headers/benchmarks.vrutoyTopLevelContainerWithDependency.h benchmarks/headers/benchmarks.vrutoytopLevelUnionBig.h benchmarks/headers/benchmarks.vrutoytopLevelUnionEmpty.h benchmarks/headers/benchmarks_vector.h
	$(CC) $(CFLAGS) -o __build/benchmarks.o -c benchmarks/details.cpp
__build/cases.o: cases/details.cpp cases/details.cpp cases/headers/cases.TestEnum.h cases/headers/cases.TestEnumItems.h cases/headers/cases.TestUnion.h cases/headers/cases.inplace1.h cases/headers/cases.inplace2.h cases/headers/cases.inplace3.h cases/headers/cases.myCycle1.h cases/headers/cases.myCycle2.h cases/headers/cases.myCycle3.h cases/headers/cases.replace7.h cases/headers/cases.replace7plus.h cases/headers/cases.replace7plusplus.h cases/headers/cases.testAllPossibleFieldConfigs.h cases/headers/cases.testAllPossibleFieldConfigsContainer.h cases/headers/cases.testArray.h cases/headers/cases.testBeforeReadBitValidation.h cases/headers/cases.testDictAny.h cases/headers/cases.testDictInt.h cases/headers/cases.testDictString.h cases/headers/cases.testEnumContainer.h cases/headers/cases.testInplaceStructArgs.h cases/headers/cases.testInplaceStructArgs2.h cases/headers/cases.testLocalFieldmask.h cases/headers/cases.testMaybe.h cases/headers/cases.testOutFieldMask.h cases/headers/cases.testOutFieldMaskContainer.h cases/headers/cases.testRecursiveFieldMask.h cases/headers/cases.testTuple.h cases/headers/cases.testUnion1.h cases/headers/cases.testUnion2.h cases/headers/cases.testUnionContainer.h cases/headers/cases.testVector.h cases/headers/cases_int.h cases/headers/cases_pair.h cases/headers/cases_tuple.h cases/headers/cases_vector.h
	$(CC) $(CFLAGS) -o __build/cases.o -c cases/details.cpp
__build/cases_bytes.o: cases_bytes/details.cpp cases_bytes/details.cpp cases_bytes/headers/cases_bytes.TestEnum.h cases_bytes/headers/cases_bytes.TestEnumItems.h cases_bytes/headers/cases_bytes.testArray.h cases_bytes/headers/cases_bytes.testDictAny.h cases_bytes/headers/cases_bytes.testDictInt.h cases_bytes/headers/cases_bytes.testDictString.h cases_bytes/headers/cases_bytes.testDictStringString.h cases_bytes/headers/cases_bytes.testEnumContainer.h cases_bytes/headers/cases_bytes.testTuple.h cases_bytes/headers/cases_bytes.testVector.h cases_bytes/headers/cases_bytes_dictionary.h cases_bytes/headers/cases_bytes_dictionaryField.h cases_bytes/headers/cases_bytes_string.h cases_bytes/headers/cases_bytes_tuple.h cases_bytes/headers/cases_bytes_vector.h
	$(CC) $(CFLAGS) -o __build/cases_bytes.o -c cases_bytes/details.cpp
