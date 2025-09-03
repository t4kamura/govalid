# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-09-03  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124417839	         9.646 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2604793	       465.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11201521	       108.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   56191	     21107 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	645400677	         1.862 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	292573257	         4.070 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	295936942	         4.053 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	349672831	         3.434 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21708608	        58.94 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1096 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1310360	       916.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   70045	     17210 ns/op	   15857 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	272936988	         4.374 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	479252360	         2.496 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11281459	       106.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13746662	        88.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	481770405	         2.489 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11397022	       105.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160420890	         7.474 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9616960	       124.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4849602	       247.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   74324	     16205 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	479780576	         2.491 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11344315	       105.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	428104407	         2.802 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11004387	       109.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	182398582	         6.560 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8351438	       144.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	36075828	        33.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8616219	       139.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4200902	       284.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   73996	     16225 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	202878552	         5.914 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8311624	       144.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	52044813	        23.02 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10095406	       119.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4215235	       283.3 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	158007918	         7.640 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	12200952	        98.46 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9449515	       128.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   72055	     16557 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	320792942	         3.736 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 7994203	       149.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642545600	         1.867 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   77242	     15437 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20798468	        57.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2470570	       486.0 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  102946	     11722 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   71587	     16744 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	24752619	        48.51 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2677819	       449.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3556492	       336.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   70358	     16733 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.802 / 0 allocs | 109.9 / 0 allocs | **39.2x** | N/A | N/A | N/A | N/A |
| Enum | 4.374 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 58.94 / 0 allocs | 1096 / 89 B / 5 allocs | **18.6x** | 916.0 / 0 allocs | **15.5x** | 17210 / 15857 B / 76 allocs | **292.0x** |
| GTE | 2.489 / 0 allocs | 105.4 / 0 allocs | **42.3x** | N/A | N/A | N/A | N/A |
| MinLength | 23.02 / 0 allocs | 119.2 / 0 allocs | **5.2x** | 283.3 / 32 B / 2 allocs | **12.3x** | N/A | N/A |
| UUID | 48.51 / 0 allocs | 449.0 / 0 allocs | **9.3x** | 336.7 / 0 allocs | **6.9x** | 16733 / 15542 B / 76 allocs | **344.9x** |
| MaxItems | 6.560 / 0 allocs | 144.1 / 0 allocs | **22.0x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.30 / 0 allocs | 139.3 / 0 allocs | **4.2x** | 284.4 / 32 B / 2 allocs | **8.5x** | 16225 / 15648 B / 81 allocs | **487.2x** |
| LT | 2.491 / 0 allocs | 105.4 / 0 allocs | **42.3x** | N/A | N/A | N/A | N/A |
| MinItems | 5.914 / 0 allocs | 144.5 / 0 allocs | **24.4x** | N/A | N/A | N/A | N/A |
| Alpha | 9.646 / 0 allocs | 465.8 / 0 allocs | **48.3x** | 108.0 / 0 allocs | **11.2x** | 21107 / 16937 B / 101 allocs | **2188.2x** |
| Required | 3.736 / 0 allocs | 149.2 / 0 allocs | **39.9x** | 1.867 / 0 allocs | **0.5x** | 15437 / 15488 B / 73 allocs | **4132.0x** |
| Length | 7.474 / 0 allocs | 124.8 / 0 allocs | **16.7x** | 247.1 / 32 B / 2 allocs | **33.1x** | 16205 / 15616 B / 79 allocs | **2168.2x** |
| URL | 57.65 / 0 allocs | 486.0 / 144 B / 1 allocs | **8.4x** | 11722 / 146 B / 1 allocs | **203.3x** | 16744 / 15681 B / 77 allocs | **290.4x** |
| Numeric | 7.640 / 0 allocs | 98.46 / 0 allocs | **12.9x** | 128.2 / 0 allocs | **16.8x** | 16557 / 15574 B / 78 allocs | **2167.1x** |
| GT | 2.496 / 0 allocs | 106.8 / 0 allocs | **42.8x** | 88.07 / 0 allocs | **35.3x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.862 | 0 allocs |
| CELMultipleExpressions | 4.070 | 0 allocs |
| CELBasic | 4.053 | 0 allocs |
| CELCrossField | 3.434 | 0 allocs |
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
