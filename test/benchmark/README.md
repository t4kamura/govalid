# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-09-03  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124419230	         9.645 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2613654	       466.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11316974	       108.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   56302	     21478 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	642744628	         1.866 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	296598542	         4.049 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	296305714	         4.049 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	350334048	         3.425 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21779134	        55.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1104 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1351773	       887.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   69955	     17232 ns/op	   15878 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	275241488	         4.361 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	481888122	         2.490 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11323578	       106.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13630080	        88.29 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	481938559	         2.490 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11364541	       106.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160604052	         7.468 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9665653	       124.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4844137	       246.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   74618	     16029 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	482322352	         2.516 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11366935	       105.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	428434497	         2.802 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	10987646	       109.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	183611649	         6.534 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8347116	       144.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	35820655	        33.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8607754	       139.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4167285	       284.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   73933	     16252 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	202807056	         5.920 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8272192	       144.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	51847780	        23.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10096666	       120.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4232017	       283.7 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	158763394	         7.560 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	12052304	       100.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9460755	       128.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   71398	     16775 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	321214708	         3.741 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8036763	       149.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642413748	         1.867 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   76534	     15635 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20790434	        57.64 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2461256	       487.0 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  103214	     12159 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   71305	     16775 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	24684392	        48.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2665218	       450.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3581176	       336.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   71060	     16682 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.802 / 0 allocs | 109.8 / 0 allocs | **39.2x** | N/A | N/A | N/A | N/A |
| Enum | 4.361 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 55.22 / 0 allocs | 1104 / 89 B / 5 allocs | **20.0x** | 887.8 / 0 allocs | **16.1x** | 17232 / 15878 B / 76 allocs | **312.1x** |
| GTE | 2.490 / 0 allocs | 106.2 / 0 allocs | **42.7x** | N/A | N/A | N/A | N/A |
| MinLength | 23.43 / 0 allocs | 120.6 / 0 allocs | **5.1x** | 283.7 / 32 B / 2 allocs | **12.1x** | N/A | N/A |
| UUID | 48.60 / 0 allocs | 450.3 / 0 allocs | **9.3x** | 336.0 / 0 allocs | **6.9x** | 16682 / 15542 B / 76 allocs | **343.3x** |
| MaxItems | 6.534 / 0 allocs | 144.0 / 0 allocs | **22.0x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.34 / 0 allocs | 139.2 / 0 allocs | **4.2x** | 284.9 / 32 B / 2 allocs | **8.5x** | 16252 / 15648 B / 81 allocs | **487.5x** |
| LT | 2.516 / 0 allocs | 105.7 / 0 allocs | **42.0x** | N/A | N/A | N/A | N/A |
| MinItems | 5.920 / 0 allocs | 144.8 / 0 allocs | **24.5x** | N/A | N/A | N/A | N/A |
| Alpha | 9.645 / 0 allocs | 466.1 / 0 allocs | **48.3x** | 108.0 / 0 allocs | **11.2x** | 21478 / 16937 B / 101 allocs | **2226.9x** |
| Required | 3.741 / 0 allocs | 149.0 / 0 allocs | **39.8x** | 1.867 / 0 allocs | **0.5x** | 15635 / 15488 B / 73 allocs | **4179.4x** |
| Length | 7.468 / 0 allocs | 124.5 / 0 allocs | **16.7x** | 246.4 / 32 B / 2 allocs | **33.0x** | 16029 / 15616 B / 79 allocs | **2146.4x** |
| URL | 57.64 / 0 allocs | 487.0 / 144 B / 1 allocs | **8.4x** | 12159 / 146 B / 1 allocs | **210.9x** | 16775 / 15681 B / 77 allocs | **291.0x** |
| Numeric | 7.560 / 0 allocs | 100.1 / 0 allocs | **13.2x** | 128.2 / 0 allocs | **17.0x** | 16775 / 15574 B / 78 allocs | **2218.9x** |
| GT | 2.490 / 0 allocs | 106.4 / 0 allocs | **42.7x** | 88.29 / 0 allocs | **35.5x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.866 | 0 allocs |
| CELMultipleExpressions | 4.049 | 0 allocs |
| CELBasic | 4.049 | 0 allocs |
| CELCrossField | 3.425 | 0 allocs |
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
