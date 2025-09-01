# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-09-01  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124432663	         9.645 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2574128	       469.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11274555	       108.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   55958	     21220 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	645985629	         1.861 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	296078730	         4.058 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	295645357	         4.057 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	350152804	         3.429 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21616075	        57.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1101 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1309911	       915.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   69096	     17410 ns/op	   15872 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	275161912	         4.360 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	481450934	         2.491 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11331394	       106.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13717003	        89.99 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	480513830	         2.494 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11372782	       105.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160733698	         7.496 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9665860	       125.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4838700	       248.5 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   72868	     16398 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	481499220	         2.491 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11303779	       105.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	428091456	         2.802 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11010656	       109.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	183747669	         6.539 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8340793	       144.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	36030624	        33.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8627188	       139.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4187655	       285.7 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   71712	     16508 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	202397036	         5.921 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8286237	       144.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	52158578	        23.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10031962	       119.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4238160	       284.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	158479702	         7.560 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	12126556	        98.75 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9428397	       128.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   70882	     16870 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	320411356	         3.739 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8050390	       149.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	643033843	         1.867 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   74907	     16031 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20760453	        57.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2472481	       487.3 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  102091	     11794 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   71007	     17150 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	24539461	        48.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2645286	       461.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3588681	       341.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   70436	     16996 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.802 / 0 allocs | 109.6 / 0 allocs | **39.1x** | N/A | N/A | N/A | N/A |
| Enum | 4.360 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 57.62 / 0 allocs | 1101 / 89 B / 5 allocs | **19.1x** | 915.4 / 0 allocs | **15.9x** | 17410 / 15872 B / 76 allocs | **302.2x** |
| GTE | 2.494 / 0 allocs | 105.7 / 0 allocs | **42.4x** | N/A | N/A | N/A | N/A |
| MinLength | 23.03 / 0 allocs | 119.5 / 0 allocs | **5.2x** | 284.0 / 32 B / 2 allocs | **12.3x** | N/A | N/A |
| UUID | 48.65 / 0 allocs | 461.8 / 0 allocs | **9.5x** | 341.3 / 0 allocs | **7.0x** | 16996 / 15542 B / 76 allocs | **349.4x** |
| MaxItems | 6.539 / 0 allocs | 144.8 / 0 allocs | **22.1x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.32 / 0 allocs | 139.5 / 0 allocs | **4.2x** | 285.7 / 32 B / 2 allocs | **8.6x** | 16508 / 15648 B / 81 allocs | **495.4x** |
| LT | 2.491 / 0 allocs | 105.6 / 0 allocs | **42.4x** | N/A | N/A | N/A | N/A |
| MinItems | 5.921 / 0 allocs | 144.4 / 0 allocs | **24.4x** | N/A | N/A | N/A | N/A |
| Alpha | 9.645 / 0 allocs | 469.9 / 0 allocs | **48.7x** | 108.6 / 0 allocs | **11.3x** | 21220 / 16937 B / 101 allocs | **2200.1x** |
| Required | 3.739 / 0 allocs | 149.8 / 0 allocs | **40.1x** | 1.867 / 0 allocs | **0.5x** | 16031 / 15488 B / 73 allocs | **4287.5x** |
| Length | 7.496 / 0 allocs | 125.0 / 0 allocs | **16.7x** | 248.5 / 32 B / 2 allocs | **33.2x** | 16398 / 15616 B / 79 allocs | **2187.6x** |
| URL | 57.66 / 0 allocs | 487.3 / 144 B / 1 allocs | **8.5x** | 11794 / 146 B / 1 allocs | **204.5x** | 17150 / 15681 B / 77 allocs | **297.4x** |
| Numeric | 7.560 / 0 allocs | 98.75 / 0 allocs | **13.1x** | 128.3 / 0 allocs | **17.0x** | 16870 / 15574 B / 78 allocs | **2231.5x** |
| GT | 2.491 / 0 allocs | 106.7 / 0 allocs | **42.8x** | 89.99 / 0 allocs | **36.1x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.861 | 0 allocs |
| CELMultipleExpressions | 4.058 | 0 allocs |
| CELBasic | 4.057 | 0 allocs |
| CELCrossField | 3.429 | 0 allocs |
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
