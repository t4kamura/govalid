# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-08-29  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124437501	         9.645 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2588953	       468.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11227578	       108.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   57170	     20979 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	645415204	         1.863 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	296133271	         4.054 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	296162587	         4.054 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	350229806	         3.426 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21719246	        55.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1085 ns/op	      88 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1326657	       905.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   70304	     17209 ns/op	   15856 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	275300907	         4.358 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	482039652	         2.490 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11200474	       107.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13716912	        89.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	481298295	         2.493 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11342349	       105.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160641926	         7.470 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9570769	       124.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4468275	       255.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   69744	     16275 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	482423085	         2.488 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11358307	       105.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	428614407	         2.803 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	10974268	       110.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	183529161	         6.535 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8308698	       144.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	36033322	        33.29 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8655680	       138.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4195110	       283.8 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   73774	     16282 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	201640836	         5.942 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8336640	       144.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	52000114	        23.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10151724	       118.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4210930	       283.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	157943289	         7.566 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	12184207	        99.63 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9471967	       128.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   71686	     16644 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	320959957	         3.736 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8036719	       148.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642108894	         1.869 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   76684	     15658 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20742380	        57.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2482161	       482.9 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  103380	     11671 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   71278	     16734 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	24640891	        49.18 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2682717	       448.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3583071	       335.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   71115	     16903 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.803 / 0 allocs | 110.1 / 0 allocs | **39.3x** | N/A | N/A | N/A | N/A |
| Enum | 4.358 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 55.58 / 0 allocs | 1085 / 88 B / 5 allocs | **19.5x** | 905.3 / 0 allocs | **16.3x** | 17209 / 15856 B / 76 allocs | **309.6x** |
| GTE | 2.493 / 0 allocs | 105.9 / 0 allocs | **42.5x** | N/A | N/A | N/A | N/A |
| MinLength | 23.00 / 0 allocs | 118.2 / 0 allocs | **5.1x** | 283.2 / 32 B / 2 allocs | **12.3x** | N/A | N/A |
| UUID | 49.18 / 0 allocs | 448.2 / 0 allocs | **9.1x** | 335.9 / 0 allocs | **6.8x** | 16903 / 15542 B / 76 allocs | **343.7x** |
| MaxItems | 6.535 / 0 allocs | 144.8 / 0 allocs | **22.2x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.29 / 0 allocs | 138.9 / 0 allocs | **4.2x** | 283.8 / 32 B / 2 allocs | **8.5x** | 16282 / 15648 B / 81 allocs | **489.1x** |
| LT | 2.488 / 0 allocs | 105.5 / 0 allocs | **42.4x** | N/A | N/A | N/A | N/A |
| MinItems | 5.942 / 0 allocs | 144.3 / 0 allocs | **24.3x** | N/A | N/A | N/A | N/A |
| Alpha | 9.645 / 0 allocs | 468.3 / 0 allocs | **48.6x** | 108.5 / 0 allocs | **11.2x** | 20979 / 16937 B / 101 allocs | **2175.1x** |
| Required | 3.736 / 0 allocs | 148.9 / 0 allocs | **39.9x** | 1.869 / 0 allocs | **0.5x** | 15658 / 15488 B / 73 allocs | **4191.1x** |
| Length | 7.470 / 0 allocs | 124.9 / 0 allocs | **16.7x** | 255.9 / 32 B / 2 allocs | **34.3x** | 16275 / 15616 B / 79 allocs | **2178.7x** |
| URL | 57.69 / 0 allocs | 482.9 / 144 B / 1 allocs | **8.4x** | 11671 / 145 B / 1 allocs | **202.3x** | 16734 / 15681 B / 77 allocs | **290.1x** |
| Numeric | 7.566 / 0 allocs | 99.63 / 0 allocs | **13.2x** | 128.1 / 0 allocs | **16.9x** | 16644 / 15574 B / 78 allocs | **2199.8x** |
| GT | 2.490 / 0 allocs | 107.4 / 0 allocs | **43.1x** | 89.59 / 0 allocs | **36.0x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.863 | 0 allocs |
| CELMultipleExpressions | 4.054 | 0 allocs |
| CELBasic | 4.054 | 0 allocs |
| CELCrossField | 3.426 | 0 allocs |
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
