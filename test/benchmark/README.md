# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-09-06  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124448221	         9.642 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2600636	       471.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11217044	       108.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   55819	     21177 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	644110057	         1.860 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	296226940	         4.051 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	296634061	         4.047 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	350588444	         3.423 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21750153	        59.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1089 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1300735	       922.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   69387	     17265 ns/op	   15866 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	275259950	         4.356 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	482145103	         2.490 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11204766	       107.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13626266	        88.96 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	482284027	         2.487 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11377929	       105.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160625215	         7.470 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9643060	       124.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4874464	       246.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   74336	     16241 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	479149375	         2.491 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11242962	       106.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	426631258	         2.806 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	10947242	       109.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	183479544	         6.538 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8313145	       144.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	36037034	        33.31 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8602761	       139.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4217984	       288.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   73473	     16272 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	201829662	         5.930 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8251519	       144.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	52114898	        23.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10141256	       118.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4260697	       281.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	158234509	         7.580 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	12110196	        98.94 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9392428	       128.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   71101	     16653 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	321389732	         3.735 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 7962834	       151.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642157938	         1.867 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   76740	     15578 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20797177	        57.63 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2486210	       482.7 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	   95487	     12284 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   71370	     16827 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	24699199	        48.61 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2693541	       454.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3579867	       336.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   70966	     16927 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.806 / 0 allocs | 109.9 / 0 allocs | **39.2x** | N/A | N/A | N/A | N/A |
| Enum | 4.356 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 59.15 / 0 allocs | 1089 / 89 B / 5 allocs | **18.4x** | 922.1 / 0 allocs | **15.6x** | 17265 / 15866 B / 76 allocs | **291.9x** |
| GTE | 2.487 / 0 allocs | 105.6 / 0 allocs | **42.5x** | N/A | N/A | N/A | N/A |
| MinLength | 23.01 / 0 allocs | 118.3 / 0 allocs | **5.1x** | 281.4 / 32 B / 2 allocs | **12.2x** | N/A | N/A |
| UUID | 48.61 / 0 allocs | 454.9 / 0 allocs | **9.4x** | 336.0 / 0 allocs | **6.9x** | 16927 / 15542 B / 76 allocs | **348.2x** |
| MaxItems | 6.538 / 0 allocs | 144.6 / 0 allocs | **22.1x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.31 / 0 allocs | 139.8 / 0 allocs | **4.2x** | 288.6 / 32 B / 2 allocs | **8.7x** | 16272 / 15648 B / 81 allocs | **488.5x** |
| LT | 2.491 / 0 allocs | 106.4 / 0 allocs | **42.7x** | N/A | N/A | N/A | N/A |
| MinItems | 5.930 / 0 allocs | 144.5 / 0 allocs | **24.4x** | N/A | N/A | N/A | N/A |
| Alpha | 9.642 / 0 allocs | 471.2 / 0 allocs | **48.9x** | 108.1 / 0 allocs | **11.2x** | 21177 / 16937 B / 101 allocs | **2196.3x** |
| Required | 3.735 / 0 allocs | 151.6 / 0 allocs | **40.6x** | 1.867 / 0 allocs | **0.5x** | 15578 / 15488 B / 73 allocs | **4170.8x** |
| Length | 7.470 / 0 allocs | 124.5 / 0 allocs | **16.7x** | 246.6 / 32 B / 2 allocs | **33.0x** | 16241 / 15616 B / 79 allocs | **2174.2x** |
| URL | 57.63 / 0 allocs | 482.7 / 144 B / 1 allocs | **8.4x** | 12284 / 146 B / 1 allocs | **213.2x** | 16827 / 15681 B / 77 allocs | **292.0x** |
| Numeric | 7.580 / 0 allocs | 98.94 / 0 allocs | **13.1x** | 128.3 / 0 allocs | **16.9x** | 16653 / 15574 B / 78 allocs | **2197.0x** |
| GT | 2.490 / 0 allocs | 107.0 / 0 allocs | **43.0x** | 88.96 / 0 allocs | **35.7x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.860 | 0 allocs |
| CELMultipleExpressions | 4.051 | 0 allocs |
| CELBasic | 4.047 | 0 allocs |
| CELCrossField | 3.423 | 0 allocs |
| CELStringLength | 1.000 | 0 allocs |
| CELNumericComparison | 1.000 | 0 allocs |

CEL (Common Expression Language) support allows complex runtime expressions with near-zero overhead.

## Running Benchmarks

```bash
# Update all benchmark documentation
make sync-benchmarks

# Run benchmarks manually
cd test
go test ./benchmark/ -bench=. -benchmem -benchtime=10s

# Run specific validator benchmarks
go test ./benchmark/ -bench=BenchmarkGoValid{ValidatorName} -benchmem
```
