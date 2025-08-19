# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-08-19  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	120072897	         9.989 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2579050	       471.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11184746	       107.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   55866	     21568 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	506801295	         2.200 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	274483040	         4.369 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	274749352	         4.372 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	350480760	         3.426 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21593821	        57.68 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1095 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1314199	       914.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   69091	     17329 ns/op	   15864 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	296285926	         4.121 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	428126787	         2.804 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11118754	       106.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13678702	        87.89 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	385088278	         3.113 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11296191	       105.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160469461	         7.477 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	10037544	       119.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4809189	       249.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   72867	     16382 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	382773147	         3.125 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11399760	       105.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	482208097	         2.494 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	10886516	       110.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	202381434	         5.923 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8344782	       144.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	37831711	        31.77 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8783367	       138.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4184787	       285.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   73072	     16474 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	192753782	         6.227 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8370351	       144.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	52069323	        23.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	 9865832	       121.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4208649	       285.7 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	123311271	         9.737 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13769491	        87.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9393757	       128.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   71030	     16782 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	296372073	         4.046 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 7998414	       148.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	550315350	         2.178 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   76184	     15692 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20504811	        58.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2485089	       481.9 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  102056	     11720 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   70605	     16930 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	24474222	        49.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2702041	       446.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3503182	       345.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   70041	     16862 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.494 / 0 allocs | 110.5 / 0 allocs | **44.3x** | N/A | N/A | N/A | N/A |
| Enum | 4.121 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 57.68 / 0 allocs | 1095 / 89 B / 5 allocs | **19.0x** | 914.2 / 0 allocs | **15.8x** | 17329 / 15864 B / 76 allocs | **300.4x** |
| GTE | 3.113 / 0 allocs | 105.7 / 0 allocs | **34.0x** | N/A | N/A | N/A | N/A |
| MinLength | 23.06 / 0 allocs | 121.9 / 0 allocs | **5.3x** | 285.7 / 32 B / 2 allocs | **12.4x** | N/A | N/A |
| UUID | 49.23 / 0 allocs | 446.1 / 0 allocs | **9.1x** | 345.1 / 0 allocs | **7.0x** | 16862 / 15542 B / 76 allocs | **342.5x** |
| MaxItems | 5.923 / 0 allocs | 144.0 / 0 allocs | **24.3x** | N/A | N/A | N/A | N/A |
| MaxLength | 31.77 / 0 allocs | 138.1 / 0 allocs | **4.3x** | 285.0 / 32 B / 2 allocs | **9.0x** | 16474 / 15648 B / 81 allocs | **518.5x** |
| LT | 3.125 / 0 allocs | 105.1 / 0 allocs | **33.6x** | N/A | N/A | N/A | N/A |
| MinItems | 6.227 / 0 allocs | 144.0 / 0 allocs | **23.1x** | N/A | N/A | N/A | N/A |
| Alpha | 9.989 / 0 allocs | 471.7 / 0 allocs | **47.2x** | 107.8 / 0 allocs | **10.8x** | 21568 / 16937 B / 101 allocs | **2159.2x** |
| Required | 4.046 / 0 allocs | 148.9 / 0 allocs | **36.8x** | 2.178 / 0 allocs | **0.5x** | 15692 / 15488 B / 73 allocs | **3878.4x** |
| Length | 7.477 / 0 allocs | 119.6 / 0 allocs | **16.0x** | 249.2 / 32 B / 2 allocs | **33.3x** | 16382 / 15616 B / 79 allocs | **2191.0x** |
| URL | 58.20 / 0 allocs | 481.9 / 144 B / 1 allocs | **8.3x** | 11720 / 145 B / 1 allocs | **201.4x** | 16930 / 15681 B / 77 allocs | **290.9x** |
| Numeric | 9.737 / 0 allocs | 87.12 / 0 allocs | **8.9x** | 128.2 / 0 allocs | **13.2x** | 16782 / 15574 B / 78 allocs | **1723.5x** |
| GT | 2.804 / 0 allocs | 106.4 / 0 allocs | **37.9x** | 87.89 / 0 allocs | **31.3x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 2.200 | 0 allocs |
| CELMultipleExpressions | 4.369 | 0 allocs |
| CELBasic | 4.372 | 0 allocs |
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
