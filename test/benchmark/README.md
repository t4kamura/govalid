# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-08-16  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	120587086	         9.952 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2565640	       471.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11230387	       107.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   56906	     21054 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	547633404	         2.193 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	275265040	         4.360 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	275247038	         4.363 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	349937134	         3.431 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21691807	        55.81 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1080 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1330288	       897.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   70838	     17102 ns/op	   15872 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	296579954	         4.044 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	427024551	         2.806 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11312290	       106.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13736905	        87.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	385486027	         3.115 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11366497	       105.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160630404	         7.468 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	10070362	       119.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4873629	       245.7 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   74427	     16162 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	385406454	         3.113 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11378107	       105.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	481680627	         2.492 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	10883443	       111.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	202758931	         5.927 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8323197	       144.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	37751620	        31.75 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8787621	       136.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4255162	       280.8 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   74248	     16138 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	192755456	         6.220 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8395002	       142.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	52081052	        23.04 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	 9878415	       122.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4241484	       281.3 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	124365764	         9.648 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13754833	        86.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9460560	       127.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   71635	     16696 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	296681864	         4.047 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8024377	       148.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	551228768	         2.178 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   75736	     15590 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20643026	        57.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2525745	       476.3 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  101060	     11790 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   71356	     16759 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	24658245	        48.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2689501	       447.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3561416	       337.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   71746	     16760 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.492 / 0 allocs | 111.3 / 0 allocs | **44.7x** | N/A | N/A | N/A | N/A |
| Enum | 4.044 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 55.81 / 0 allocs | 1080 / 89 B / 5 allocs | **19.4x** | 897.2 / 0 allocs | **16.1x** | 17102 / 15872 B / 76 allocs | **306.4x** |
| GTE | 3.115 / 0 allocs | 105.6 / 0 allocs | **33.9x** | N/A | N/A | N/A | N/A |
| MinLength | 23.04 / 0 allocs | 122.1 / 0 allocs | **5.3x** | 281.3 / 32 B / 2 allocs | **12.2x** | N/A | N/A |
| UUID | 48.57 / 0 allocs | 447.9 / 0 allocs | **9.2x** | 337.3 / 0 allocs | **6.9x** | 16760 / 15542 B / 76 allocs | **345.1x** |
| MaxItems | 5.927 / 0 allocs | 144.5 / 0 allocs | **24.4x** | N/A | N/A | N/A | N/A |
| MaxLength | 31.75 / 0 allocs | 136.9 / 0 allocs | **4.3x** | 280.8 / 32 B / 2 allocs | **8.8x** | 16138 / 15648 B / 81 allocs | **508.3x** |
| LT | 3.113 / 0 allocs | 105.4 / 0 allocs | **33.9x** | N/A | N/A | N/A | N/A |
| MinItems | 6.220 / 0 allocs | 142.5 / 0 allocs | **22.9x** | N/A | N/A | N/A | N/A |
| Alpha | 9.952 / 0 allocs | 471.0 / 0 allocs | **47.3x** | 107.7 / 0 allocs | **10.8x** | 21054 / 16937 B / 101 allocs | **2115.6x** |
| Required | 4.047 / 0 allocs | 148.8 / 0 allocs | **36.8x** | 2.178 / 0 allocs | **0.5x** | 15590 / 15488 B / 73 allocs | **3852.2x** |
| Length | 7.468 / 0 allocs | 119.2 / 0 allocs | **16.0x** | 245.7 / 32 B / 2 allocs | **32.9x** | 16162 / 15616 B / 79 allocs | **2164.2x** |
| URL | 57.95 / 0 allocs | 476.3 / 144 B / 1 allocs | **8.2x** | 11790 / 146 B / 1 allocs | **203.5x** | 16759 / 15681 B / 77 allocs | **289.2x** |
| Numeric | 9.648 / 0 allocs | 86.74 / 0 allocs | **9.0x** | 127.9 / 0 allocs | **13.3x** | 16696 / 15574 B / 78 allocs | **1730.5x** |
| GT | 2.806 / 0 allocs | 106.2 / 0 allocs | **37.8x** | 87.87 / 0 allocs | **31.3x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 2.193 | 0 allocs |
| CELMultipleExpressions | 4.360 | 0 allocs |
| CELBasic | 4.363 | 0 allocs |
| CELCrossField | 3.431 | 0 allocs |
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
