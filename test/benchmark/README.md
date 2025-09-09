# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-09-09  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124423137	         9.663 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2605066	       465.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11253799	       109.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   56974	     20884 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	559996090	         2.108 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	296203453	         4.050 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	293538878	         4.072 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	348980149	         3.432 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21765828	        55.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1111 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1353634	       886.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   68907	     17222 ns/op	   15873 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	275364184	         4.357 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	481444286	         2.519 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11259484	       106.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13663078	        88.17 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	482209663	         2.492 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11390151	       106.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160395621	         7.475 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9644949	       124.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4850661	       247.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   74767	     16240 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	480890370	         2.491 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11295063	       105.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	428277296	         2.805 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11032057	       109.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	183679644	         6.541 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8340045	       144.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	35685355	        33.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8565288	       140.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4224396	       283.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   72915	     16356 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	202908286	         5.912 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8302192	       144.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	51868174	        23.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10101247	       119.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4237912	       282.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	157926294	         7.595 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	12144838	        98.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9219182	       128.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   72568	     16658 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	321004024	         3.741 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8008600	       150.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	641623330	         1.867 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   75730	     15689 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20735096	        57.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2473405	       483.8 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	   95420	     11950 ns/op	     148 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   71646	     16797 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	24828404	        48.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2685124	       448.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3568594	       336.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   71559	     16748 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.805 / 0 allocs | 109.5 / 0 allocs | **39.0x** | N/A | N/A | N/A | N/A |
| Enum | 4.357 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 55.30 / 0 allocs | 1111 / 89 B / 5 allocs | **20.1x** | 886.5 / 0 allocs | **16.0x** | 17222 / 15873 B / 76 allocs | **311.4x** |
| GTE | 2.492 / 0 allocs | 106.6 / 0 allocs | **42.8x** | N/A | N/A | N/A | N/A |
| MinLength | 23.05 / 0 allocs | 119.0 / 0 allocs | **5.2x** | 282.9 / 32 B / 2 allocs | **12.3x** | N/A | N/A |
| UUID | 48.49 / 0 allocs | 448.2 / 0 allocs | **9.2x** | 336.5 / 0 allocs | **6.9x** | 16748 / 15542 B / 76 allocs | **345.4x** |
| MaxItems | 6.541 / 0 allocs | 144.8 / 0 allocs | **22.1x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.36 / 0 allocs | 140.5 / 0 allocs | **4.2x** | 283.9 / 32 B / 2 allocs | **8.5x** | 16356 / 15648 B / 81 allocs | **490.3x** |
| LT | 2.491 / 0 allocs | 105.1 / 0 allocs | **42.2x** | N/A | N/A | N/A | N/A |
| MinItems | 5.912 / 0 allocs | 144.0 / 0 allocs | **24.4x** | N/A | N/A | N/A | N/A |
| Alpha | 9.663 / 0 allocs | 465.6 / 0 allocs | **48.2x** | 109.2 / 0 allocs | **11.3x** | 20884 / 16937 B / 101 allocs | **2161.2x** |
| Required | 3.741 / 0 allocs | 150.1 / 0 allocs | **40.1x** | 1.867 / 0 allocs | **0.5x** | 15689 / 15488 B / 73 allocs | **4193.8x** |
| Length | 7.475 / 0 allocs | 124.4 / 0 allocs | **16.6x** | 247.1 / 32 B / 2 allocs | **33.1x** | 16240 / 15616 B / 79 allocs | **2172.6x** |
| URL | 57.53 / 0 allocs | 483.8 / 144 B / 1 allocs | **8.4x** | 11950 / 148 B / 1 allocs | **207.7x** | 16797 / 15681 B / 77 allocs | **292.0x** |
| Numeric | 7.595 / 0 allocs | 98.60 / 0 allocs | **13.0x** | 128.6 / 0 allocs | **16.9x** | 16658 / 15574 B / 78 allocs | **2193.3x** |
| GT | 2.519 / 0 allocs | 106.3 / 0 allocs | **42.2x** | 88.17 / 0 allocs | **35.0x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 2.108 | 0 allocs |
| CELMultipleExpressions | 4.050 | 0 allocs |
| CELBasic | 4.072 | 0 allocs |
| CELCrossField | 3.432 | 0 allocs |
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
